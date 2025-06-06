//  Copyright 2024 Google LLC
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package handler

import (
	"context"
	"fmt"
	"runtime"
	"testing"

	acpb "github.com/GoogleCloudPlatform/agentcommunication_client/gapic/agentcommunicationpb"
	acmpb "github.com/GoogleCloudPlatform/google-guest-agent/internal/acp/proto/google_guest_agent/acp"
	"github.com/GoogleCloudPlatform/google-guest-agent/internal/acs/client"
	"github.com/GoogleCloudPlatform/google-guest-agent/internal/cfg"
	"github.com/GoogleCloudPlatform/google-guest-agent/internal/events"
	"github.com/GoogleCloudPlatform/google-guest-agent/internal/osinfo"
	"github.com/GoogleCloudPlatform/google-guest-agent/internal/plugin/manager"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
	apb "google.golang.org/protobuf/types/known/anypb"
)

type fakeConnection struct {
	seenMessage *acpb.MessageBody
	throwErr    bool
}

func (c *fakeConnection) Close() {}

func (c *fakeConnection) SendMessage(msg *acpb.MessageBody) error {
	if c.throwErr {
		return fmt.Errorf("test_error")
	}

	c.seenMessage = msg
	return nil
}

func (c *fakeConnection) Receive() (*acpb.MessageBody, error) {
	return nil, nil
}

func TestHandleMessage(t *testing.T) {
	if err := cfg.Load(nil); err != nil {
		t.Fatalf("cfg.Load(nil) = %v, want nil", err)
	}
	cfg.Retrieve().Core.ACSClient = true

	testMsg := &acmpb.ConfigurePluginStates{
		ConfigurePlugins: []*acmpb.ConfigurePluginStates_ConfigurePlugin{
			&acmpb.ConfigurePluginStates_ConfigurePlugin{
				Action: acmpb.ConfigurePluginStates_INSTALL,
				Plugin: &acmpb.ConfigurePluginStates_Plugin{
					Name:       "basic_plugin",
					RevisionId: "1",
				},
			},
		},
	}

	msgBytes, err := proto.Marshal(testMsg)
	if err != nil {
		t.Fatalf("proto.Marshal(%v) = %v, want nil", testMsg, err)
	}

	v := "1.0.0"
	info := osinfo.OSInfo{
		Architecture: "x86_64",
		OS:           "linux",
		VersionID:    "1.0.0",
	}

	wantOSInfo := &acmpb.OSInfo{
		Architecture:  info.Architecture,
		Type:          runtime.GOOS,
		Version:       info.VersionID,
		ShortName:     info.OS,
		LongName:      info.PrettyName,
		KernelRelease: info.KernelRelease,
		KernelVersion: info.KernelVersion,
	}

	wantAgentInfo := &acmpb.AgentInfo{
		Name:         "GCEGuestAgent",
		Architecture: runtime.GOARCH,
		AgentCapabilities: []acmpb.AgentInfo_AgentCapability{
			acmpb.AgentInfo_GET_AGENT_INFO,
			acmpb.AgentInfo_GET_OS_INFO,
			acmpb.AgentInfo_LIST_PLUGIN_STATES,
			acmpb.AgentInfo_CONFIGURE_PLUGIN_STATES,
		},
		Version: v,
	}

	f := dataFetchers{osInfoReader: func() osinfo.OSInfo { return info }, clientVersion: v, pluginManager: &manager.PluginManager{}}
	ctx := context.Background()

	tests := []struct {
		desc        string
		messageType string
		wantLabels  map[string]string
		want        proto.Message
		skipCall    bool
		eventError  error
		eventData   proto.Message
		throwErr    bool
	}{
		{
			desc:        "get_os_info",
			messageType: getOSInfoMsg,
			wantLabels:  map[string]string{messageTypeLabel: osInfoMsg},
			want:        wantOSInfo,
			eventData:   &acpb.MessageBody{Labels: map[string]string{messageTypeLabel: getOSInfoMsg}},
		},
		{
			desc:        "get_agent_info",
			messageType: getAgentInfoMsg,
			wantLabels:  map[string]string{messageTypeLabel: agentInfoMsg},
			want:        wantAgentInfo,
			eventData:   &acpb.MessageBody{Labels: map[string]string{messageTypeLabel: getAgentInfoMsg}},
		},
		{
			desc:        "list_plugin_states",
			messageType: listPluginStatesMsg,
			wantLabels:  map[string]string{messageTypeLabel: currentPluginStatesMsg},
			want:        &acmpb.CurrentPluginStates{},
			eventData:   &acpb.MessageBody{Labels: map[string]string{messageTypeLabel: listPluginStatesMsg}},
		},
		{
			desc:        "configure_plugin_states",
			messageType: configurePluginStatesMsg,
			want:        &acmpb.ConfigurePluginStates{},
			eventData:   &acpb.MessageBody{Labels: map[string]string{messageTypeLabel: configurePluginStatesMsg}},
		},
		{
			desc:        "configure_plugin_states_skip_call",
			messageType: configurePluginStatesMsg,
			want:        &acmpb.ConfigurePluginStates{},
			eventData:   &acpb.MessageBody{Labels: map[string]string{messageTypeLabel: configurePluginStatesMsg}, Body: &apb.Any{Value: msgBytes, TypeUrl: string(proto.MessageName(testMsg))}},
		},
		{
			desc:        "unknown_message_type",
			messageType: "get_agent_stack_trace",
			skipCall:    true,
			eventData:   &acpb.MessageBody{Labels: map[string]string{messageTypeLabel: "get_agent_stack_trace"}},
		},
		{
			desc:       "event_error",
			skipCall:   true,
			eventError: fmt.Errorf("test_error"),
		},
		{
			desc:     "invalid_event_data",
			skipCall: true,
		},
		{
			desc:      "send_error_continue",
			eventData: &acpb.MessageBody{Labels: map[string]string{"message_type": getAgentInfoMsg}},
			throwErr:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			d := &events.EventData{Error: tc.eventError, Data: tc.eventData}
			connection := &fakeConnection{throwErr: tc.throwErr}
			ctx := context.WithValue(ctx, client.OverrideConnection, connection)

			got, noop, err := f.handleMessage(ctx, "test-event", nil, d)
			if err != nil {
				t.Fatalf("handleMessage(ctx, %s, nil, %+v) error = %v, want nil", "test-event", d, err)
			}
			if !noop {
				t.Errorf("handleMessage(ctx, %s, nil, %+v) = %t, want false", "test-event", d, noop)
			}
			if !got {
				t.Errorf("handleMessage(ctx, %s, nil, %+v) = false, want true", "test-event", d)
			}
			f.worker.Wait()

			sentMsg := connection.seenMessage

			if tc.skipCall && sentMsg != nil {
				t.Fatalf("handleMessage(ctx, %s, nil, %+v) shouldn't have attempted to send a message", "test-event", d)
			}

			var msg proto.Message
			switch tc.messageType {
			case getOSInfoMsg:
				msg = new(acmpb.OSInfo)
			case getAgentInfoMsg:
				msg = new(acmpb.AgentInfo)
			case listPluginStatesMsg:
				msg = new(acmpb.CurrentPluginStates)
			default:
				// If known message type is not set in test run skip following checks.
				return
			}

			if err := sentMsg.GetBody().UnmarshalTo(msg); err != nil {
				t.Fatalf("Unmarshal message body failed with error: %v", err)
			}

			if diff := cmp.Diff(tc.want, msg, protocmp.Transform()); diff != "" {
				t.Errorf("Unexpected message body diff (-want +got):\n%s", diff)
			}

			if diff := cmp.Diff(tc.wantLabels, sentMsg.GetLabels()); diff != "" {
				t.Fatalf("Unexpected message labels diff (-want +got):\n%s", diff)
			}
		})
	}
}
