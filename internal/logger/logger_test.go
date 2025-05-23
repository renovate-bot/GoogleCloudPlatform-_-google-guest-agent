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

package logger

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/GoogleCloudPlatform/agentcommunication_client"
	"github.com/GoogleCloudPlatform/galog"
	"github.com/GoogleCloudPlatform/google-guest-agent/internal/events"
	"github.com/GoogleCloudPlatform/google-guest-agent/internal/metadata"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/exp/slices"
)

func TestNotAdditionalLoggers(t *testing.T) {
	opts := Options{
		Verbosity: 10,
		Level:     1,
	}
	ctx := context.Background()
	if err := Init(ctx, opts); err != nil {
		t.Fatalf("Init() failed: %v", err)
	}

	if galog.MinVerbosity() != opts.Verbosity {
		t.Errorf("MinVerbosity() = %d, want %d", galog.MinVerbosity(), opts.Verbosity)
	}

	if galog.CurrentLevel() != galog.ErrorLevel {
		t.Errorf("Level() = %s, want %d", galog.CurrentLevel(), opts.Level)
	}
}

func TestFileAdditionalLogger(t *testing.T) {
	tests := []struct {
		name      string
		logLevel  int
		wantError bool
	}{
		{
			name:      "valid-log-level-1",
			logLevel:  1,
			wantError: false,
		},
		{
			name:      "valid-log-level-2",
			logLevel:  2,
			wantError: false,
		},
		{
			name:      "valid-log-level-3",
			logLevel:  3,
			wantError: false,
		},
		{
			name:      "valid-log-level-4",
			logLevel:  4,
			wantError: false,
		},
		{
			name:      "invalid-log-level-5",
			logLevel:  5,
			wantError: true,
		},
	}

	ctx := context.Background()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			logFile := filepath.Join(t.TempDir(), "test.log")
			opts := Options{
				LogFile:           logFile,
				Verbosity:         10,
				Level:             tc.logLevel,
				LogToCloudLogging: true,
			}

			err := Init(ctx, opts)

			if (err == nil) == tc.wantError {
				t.Fatalf("Init() failed: %v", err)
			}

			ids := galog.RegisteredBackendIDs()
			fileBackendID := "log-backend,file"
			if !slices.Contains(ids, fileBackendID) {
				t.Errorf("RegisteredBackendIDs() = %v, want %v", ids, fileBackendID)
			}

			if !events.FetchManager().IsSubscribed(metadata.LongpollEvent, loggerMetadataSubscriberID) {
				t.Errorf("events.IsSubscribed(%q, %q) = false, want true", metadata.LongpollEvent, loggerMetadataSubscriberID)
			}
		})
	}
}

func TestStderrAdditionalLogger(t *testing.T) {
	opts := Options{
		LogToStderr: true,
		Verbosity:   10,
		Level:       1,
	}
	ctx := context.Background()
	if err := Init(ctx, opts); err != nil {
		t.Fatalf("Init() failed: %v", err)
	}

	ids := galog.RegisteredBackendIDs()
	stderrBackendID := "log-backend,stderr"
	if !slices.Contains(ids, stderrBackendID) {
		t.Errorf("RegisteredBackendIDs() = %v, want %v", ids, stderrBackendID)
	}
}

func TestInitCloudLogging(t *testing.T) {
	mdsJSON := `
	{
		"project":  {
			"attributes": {
				"project-id": "some-project-id"
			}
		}
	}
	`

	mds, err := metadata.UnmarshalDescriptor(mdsJSON)
	if err != nil {
		t.Fatalf("UnmarshalDescriptor(%q) returned an unexpected error: %v", mdsJSON, err)
	}

	ctx := context.Background()
	be, err := galog.NewCloudBackend(ctx, galog.CloudLoggingInitModeLazy, nil)
	if err != nil {
		t.Fatalf("NewCloudBackend() returned an unexpected error: %v", err)
	}

	tests := []struct {
		name     string
		data     any
		mds      any
		want     bool
		wantNoop bool
	}{
		{
			name:     "invalid-event-data-type",
			data:     "invalid-data-we-dont-want-string",
			want:     true,
			wantNoop: true,
		},
		{
			name: "invalid-mds",
			data: &Options{
				cloudLoggingBackend: be,
			},
			mds:      "invalid-mds-data-type-we-dont-want-string",
			want:     true,
			wantNoop: true,
		},
		{
			name: "success",
			data: &Options{
				CloudIdent:                        "test-ident",
				ProgramVersion:                    "test-version",
				cloudLoggingWithoutAuthentication: true,
				cloudLoggingBackend:               be,
			},
			wantNoop: false,
			mds:      mds,
			want:     false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			evData := &events.EventData{Data: tc.mds}
			got, gotNoop, err := initCloudLogging(ctx, "eventType", tc.data, evData)
			if (err != nil) != tc.want {
				t.Errorf("initCloudLogging() returned error: %v, want error: %t", err, tc.want)
			}
			if gotNoop != tc.wantNoop {
				t.Errorf("initCloudLogging() returned noop: %t, want noop: %t", gotNoop, tc.wantNoop)
			}
			if got != tc.want {
				t.Errorf("initCloudLogging() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestInitACSClientLogging(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name string
		opts Options
		want bool
	}{
		{
			name: "enable_logging",
			opts: Options{ACSClientDebugLogging: true, Level: 3},
			want: true,
		},
		{
			name: "disable_logging",
			opts: Options{ACSClientDebugLogging: false, Level: 3},
			want: false,
		},
		{
			name: "default_logging",
			opts: Options{Level: 3},
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if err := Init(ctx, tc.opts); err != nil {
				t.Fatalf("Init(ctx, %+v) failed unexpectedly: %v", tc.opts, err)
			}
			got := client.DebugLogging
			if got != tc.want {
				t.Errorf("Init(ctx, %+v) set acs.DebugLogging = %t, want %t", tc.opts, got, tc.want)
			}
		})
	}
}

func TestParseMIGLabels(t *testing.T) {
	tests := []struct {
		name    string
		mds     string
		want    map[string]string
		wantErr bool
	}{
		{
			name: "no-created-by",
			mds:  `{}`,
			want: map[string]string{},
		},
		{
			name: "zonal",
			mds: `
{
	"instance": {
		"attributes": {
			"created-by": "projects/test-project/zones/us-central1-a/instanceGroupManagers/test-mig"
		}
	}
}
`,
			want: map[string]string{
				migNameLabel: "test-mig",
				migZoneLabel: "us-central1-a",
			},
		},
		{
			name: "regional",
			mds: `
{
	"instance": {
		"attributes": {
			"created-by": "projects/test-project/regions/us-central1/instanceGroupManagers/test-mig"
		}
	}
}
`,
			want: map[string]string{
				migNameLabel:   "test-mig",
				migRegionLabel: "us-central1",
			},
		},
		{
			name: "invalid-created-by",
			mds: `
{
	"instance": {
		"attributes": {
			"created-by": "random-string"
		}
	}
}
`,
			want: map[string]string{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mds, err := metadata.UnmarshalDescriptor(tc.mds)
			if err != nil {
				t.Fatalf("UnmarshalDescriptor(%q) returned an unexpected error: %v", tc.mds, err)
			}
			got, err := parseMIGLabels(mds)
			if (err != nil) != tc.wantErr {
				t.Errorf("parseMIGLabels(%q) returned an unexpected error: %v", tc.mds, err)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("parseMIGLabels(%q) returned an unexpected diff (-want +got): %v", tc.mds, diff)
			}
		})
	}
}
