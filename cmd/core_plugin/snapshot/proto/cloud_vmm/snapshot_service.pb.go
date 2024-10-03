//  Copyright 2019 Google
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: snapshot_service.proto

package cloud_vmm

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OperationType int32

const (
	OperationType_NOT_SET       OperationType = 0
	OperationType_PRE_SNAPSHOT  OperationType = 1
	OperationType_POST_SNAPSHOT OperationType = 2
)

// Enum value maps for OperationType.
var (
	OperationType_name = map[int32]string{
		0: "NOT_SET",
		1: "PRE_SNAPSHOT",
		2: "POST_SNAPSHOT",
	}
	OperationType_value = map[string]int32{
		"NOT_SET":       0,
		"PRE_SNAPSHOT":  1,
		"POST_SNAPSHOT": 2,
	}
)

func (x OperationType) Enum() *OperationType {
	p := new(OperationType)
	*p = x
	return p
}

func (x OperationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperationType) Descriptor() protoreflect.EnumDescriptor {
	return file_snapshot_service_proto_enumTypes[0].Descriptor()
}

func (OperationType) Type() protoreflect.EnumType {
	return &file_snapshot_service_proto_enumTypes[0]
}

func (x OperationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperationType.Descriptor instead.
func (OperationType) EnumDescriptor() ([]byte, []int) {
	return file_snapshot_service_proto_rawDescGZIP(), []int{0}
}

type SupportedFeatures int32

const (
	SupportedFeatures_NONE      SupportedFeatures = 0
	SupportedFeatures_SNAPSHOTS SupportedFeatures = 1
)

// Enum value maps for SupportedFeatures.
var (
	SupportedFeatures_name = map[int32]string{
		0: "NONE",
		1: "SNAPSHOTS",
	}
	SupportedFeatures_value = map[string]int32{
		"NONE":      0,
		"SNAPSHOTS": 1,
	}
)

func (x SupportedFeatures) Enum() *SupportedFeatures {
	p := new(SupportedFeatures)
	*p = x
	return p
}

func (x SupportedFeatures) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SupportedFeatures) Descriptor() protoreflect.EnumDescriptor {
	return file_snapshot_service_proto_enumTypes[1].Descriptor()
}

func (SupportedFeatures) Type() protoreflect.EnumType {
	return &file_snapshot_service_proto_enumTypes[1]
}

func (x SupportedFeatures) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SupportedFeatures.Descriptor instead.
func (SupportedFeatures) EnumDescriptor() ([]byte, []int) {
	return file_snapshot_service_proto_rawDescGZIP(), []int{1}
}

type AgentErrorCode int32

const (
	AgentErrorCode_NO_ERROR AgentErrorCode = 0
	// The snapshot config was improper in some way.
	AgentErrorCode_INVALID_CONFIG AgentErrorCode = 1
	// The pre or post snapshot script was not found on disk.
	AgentErrorCode_SCRIPT_NOT_FOUND AgentErrorCode = 2
	// The pre or post snapshot script timed out.
	AgentErrorCode_SCRIPT_TIMED_OUT AgentErrorCode = 3
	// The pre or post snapshot script returned an error, but the "continue on
	// error" flag was not set.
	AgentErrorCode_UNHANDLED_SCRIPT_ERROR AgentErrorCode = 4
)

// Enum value maps for AgentErrorCode.
var (
	AgentErrorCode_name = map[int32]string{
		0: "NO_ERROR",
		1: "INVALID_CONFIG",
		2: "SCRIPT_NOT_FOUND",
		3: "SCRIPT_TIMED_OUT",
		4: "UNHANDLED_SCRIPT_ERROR",
	}
	AgentErrorCode_value = map[string]int32{
		"NO_ERROR":               0,
		"INVALID_CONFIG":         1,
		"SCRIPT_NOT_FOUND":       2,
		"SCRIPT_TIMED_OUT":       3,
		"UNHANDLED_SCRIPT_ERROR": 4,
	}
)

func (x AgentErrorCode) Enum() *AgentErrorCode {
	p := new(AgentErrorCode)
	*p = x
	return p
}

func (x AgentErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AgentErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_snapshot_service_proto_enumTypes[2].Descriptor()
}

func (AgentErrorCode) Type() protoreflect.EnumType {
	return &file_snapshot_service_proto_enumTypes[2]
}

func (x AgentErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AgentErrorCode.Descriptor instead.
func (AgentErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_snapshot_service_proto_rawDescGZIP(), []int{2}
}

type SnapshotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The operation id of the snapshot.
	OperationId int32 `protobuf:"varint,1,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	// A list of comma separated target/lun values, e.g. "1/2,3/4".
	DiskList string `protobuf:"bytes,2,opt,name=disk_list,json=diskList,proto3" json:"disk_list,omitempty"`
	// The operation type.
	Type OperationType `protobuf:"varint,3,opt,name=type,proto3,enum=cloud.vmm.OperationType" json:"type,omitempty"`
}

func (x *SnapshotRequest) Reset() {
	*x = SnapshotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snapshot_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotRequest) ProtoMessage() {}

func (x *SnapshotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_snapshot_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotRequest.ProtoReflect.Descriptor instead.
func (*SnapshotRequest) Descriptor() ([]byte, []int) {
	return file_snapshot_service_proto_rawDescGZIP(), []int{0}
}

func (x *SnapshotRequest) GetOperationId() int32 {
	if x != nil {
		return x.OperationId
	}
	return 0
}

func (x *SnapshotRequest) GetDiskList() string {
	if x != nil {
		return x.DiskList
	}
	return ""
}

func (x *SnapshotRequest) GetType() OperationType {
	if x != nil {
		return x.Type
	}
	return OperationType_NOT_SET
}

type ServerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SupportedFeatures []SupportedFeatures `protobuf:"varint,1,rep,packed,name=supported_features,json=supportedFeatures,proto3,enum=cloud.vmm.SupportedFeatures" json:"supported_features,omitempty"`
}

func (x *ServerInfo) Reset() {
	*x = ServerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snapshot_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerInfo) ProtoMessage() {}

func (x *ServerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_snapshot_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerInfo.ProtoReflect.Descriptor instead.
func (*ServerInfo) Descriptor() ([]byte, []int) {
	return file_snapshot_service_proto_rawDescGZIP(), []int{1}
}

func (x *ServerInfo) GetSupportedFeatures() []SupportedFeatures {
	if x != nil {
		return x.SupportedFeatures
	}
	return nil
}

type SnapshotResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The operation id of the snapshot.
	OperationId int32 `protobuf:"varint,1,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	// The return code of the scripts run by the guest. If this is non-zero, then
	// agent_return_code should be UNHANDLED_SCRIPT_ERROR.
	ScriptsReturnCode int32 `protobuf:"varint,2,opt,name=scripts_return_code,json=scriptsReturnCode,proto3" json:"scripts_return_code,omitempty"`
	// The agent return code.
	AgentReturnCode AgentErrorCode `protobuf:"varint,3,opt,name=agent_return_code,json=agentReturnCode,proto3,enum=cloud.vmm.AgentErrorCode" json:"agent_return_code,omitempty"`
	// The operation type.
	Type OperationType `protobuf:"varint,4,opt,name=type,proto3,enum=cloud.vmm.OperationType" json:"type,omitempty"`
}

func (x *SnapshotResponse) Reset() {
	*x = SnapshotResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snapshot_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotResponse) ProtoMessage() {}

func (x *SnapshotResponse) ProtoReflect() protoreflect.Message {
	mi := &file_snapshot_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotResponse.ProtoReflect.Descriptor instead.
func (*SnapshotResponse) Descriptor() ([]byte, []int) {
	return file_snapshot_service_proto_rawDescGZIP(), []int{2}
}

func (x *SnapshotResponse) GetOperationId() int32 {
	if x != nil {
		return x.OperationId
	}
	return 0
}

func (x *SnapshotResponse) GetScriptsReturnCode() int32 {
	if x != nil {
		return x.ScriptsReturnCode
	}
	return 0
}

func (x *SnapshotResponse) GetAgentReturnCode() AgentErrorCode {
	if x != nil {
		return x.AgentReturnCode
	}
	return AgentErrorCode_NO_ERROR
}

func (x *SnapshotResponse) GetType() OperationType {
	if x != nil {
		return x.Type
	}
	return OperationType_NOT_SET
}

type GuestReady struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestServerInfo bool `protobuf:"varint,1,opt,name=request_server_info,json=requestServerInfo,proto3" json:"request_server_info,omitempty"`
}

func (x *GuestReady) Reset() {
	*x = GuestReady{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snapshot_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GuestReady) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GuestReady) ProtoMessage() {}

func (x *GuestReady) ProtoReflect() protoreflect.Message {
	mi := &file_snapshot_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GuestReady.ProtoReflect.Descriptor instead.
func (*GuestReady) Descriptor() ([]byte, []int) {
	return file_snapshot_service_proto_rawDescGZIP(), []int{3}
}

func (x *GuestReady) GetRequestServerInfo() bool {
	if x != nil {
		return x.RequestServerInfo
	}
	return false
}

type GuestMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Msg:
	//
	//	*GuestMessage_SnapshotRequest
	//	*GuestMessage_ServerInfo
	Msg isGuestMessage_Msg `protobuf_oneof:"msg"`
}

func (x *GuestMessage) Reset() {
	*x = GuestMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snapshot_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GuestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GuestMessage) ProtoMessage() {}

func (x *GuestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_snapshot_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GuestMessage.ProtoReflect.Descriptor instead.
func (*GuestMessage) Descriptor() ([]byte, []int) {
	return file_snapshot_service_proto_rawDescGZIP(), []int{4}
}

func (m *GuestMessage) GetMsg() isGuestMessage_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *GuestMessage) GetSnapshotRequest() *SnapshotRequest {
	if x, ok := x.GetMsg().(*GuestMessage_SnapshotRequest); ok {
		return x.SnapshotRequest
	}
	return nil
}

func (x *GuestMessage) GetServerInfo() *ServerInfo {
	if x, ok := x.GetMsg().(*GuestMessage_ServerInfo); ok {
		return x.ServerInfo
	}
	return nil
}

type isGuestMessage_Msg interface {
	isGuestMessage_Msg()
}

type GuestMessage_SnapshotRequest struct {
	SnapshotRequest *SnapshotRequest `protobuf:"bytes,1,opt,name=snapshot_request,json=snapshotRequest,proto3,oneof"`
}

type GuestMessage_ServerInfo struct {
	ServerInfo *ServerInfo `protobuf:"bytes,2,opt,name=server_info,json=serverInfo,proto3,oneof"`
}

func (*GuestMessage_SnapshotRequest) isGuestMessage_Msg() {}

func (*GuestMessage_ServerInfo) isGuestMessage_Msg() {}

type ServerAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ServerAck) Reset() {
	*x = ServerAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snapshot_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerAck) ProtoMessage() {}

func (x *ServerAck) ProtoReflect() protoreflect.Message {
	mi := &file_snapshot_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerAck.ProtoReflect.Descriptor instead.
func (*ServerAck) Descriptor() ([]byte, []int) {
	return file_snapshot_service_proto_rawDescGZIP(), []int{5}
}

var File_snapshot_service_proto protoreflect.FileDescriptor

var file_snapshot_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x76, 0x6d, 0x6d, 0x22, 0x7f, 0x0a, 0x0f, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x69, 0x73,
	0x6b, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69,
	0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x6d, 0x6d,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x22, 0x59, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x4b, 0x0a, 0x12, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x5f,
	0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1c,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x6d, 0x6d, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x64, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x52, 0x11, 0x73, 0x75,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x22,
	0xda, 0x01, 0x0a, 0x10, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x73, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x52, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x45, 0x0a, 0x11, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x5f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x19, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x6d, 0x6d, 0x2e, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x0f, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2c,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x6d, 0x6d, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3c, 0x0a, 0x0a,
	0x47, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x61, 0x64, 0x79, 0x12, 0x2e, 0x0a, 0x13, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x98, 0x01, 0x0a, 0x0c, 0x47,
	0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x47, 0x0a, 0x10, 0x73,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x6d,
	0x6d, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x48, 0x00, 0x52, 0x0f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x6d, 0x6d, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x48, 0x00, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x05,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x0b, 0x0a, 0x09, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41,
	0x63, 0x6b, 0x2a, 0x41, 0x0a, 0x0d, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x10, 0x00,
	0x12, 0x10, 0x0a, 0x0c, 0x50, 0x52, 0x45, 0x5f, 0x53, 0x4e, 0x41, 0x50, 0x53, 0x48, 0x4f, 0x54,
	0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x4f, 0x53, 0x54, 0x5f, 0x53, 0x4e, 0x41, 0x50, 0x53,
	0x48, 0x4f, 0x54, 0x10, 0x02, 0x2a, 0x2c, 0x0a, 0x11, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x64, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f,
	0x4e, 0x45, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x4e, 0x41, 0x50, 0x53, 0x48, 0x4f, 0x54,
	0x53, 0x10, 0x01, 0x2a, 0x7a, 0x0a, 0x0e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x4f, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43,
	0x4f, 0x4e, 0x46, 0x49, 0x47, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x43, 0x52, 0x49, 0x50,
	0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x02, 0x12, 0x14, 0x0a,
	0x10, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x44, 0x5f, 0x4f, 0x55,
	0x54, 0x10, 0x03, 0x12, 0x1a, 0x0a, 0x16, 0x55, 0x4e, 0x48, 0x41, 0x4e, 0x44, 0x4c, 0x45, 0x44,
	0x5f, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04, 0x32,
	0xaa, 0x01, 0x0a, 0x0f, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x76, 0x6d, 0x6d, 0x2e, 0x47, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x61, 0x64, 0x79, 0x1a, 0x17,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x6d, 0x6d, 0x2e, 0x47, 0x75, 0x65, 0x73, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4f, 0x0a, 0x18, 0x48,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x46, 0x72,
	0x6f, 0x6d, 0x47, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x76, 0x6d, 0x6d, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x14, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x6d, 0x6d,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x63, 0x6b, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x76, 0x6d, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_snapshot_service_proto_rawDescOnce sync.Once
	file_snapshot_service_proto_rawDescData = file_snapshot_service_proto_rawDesc
)

func file_snapshot_service_proto_rawDescGZIP() []byte {
	file_snapshot_service_proto_rawDescOnce.Do(func() {
		file_snapshot_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_snapshot_service_proto_rawDescData)
	})
	return file_snapshot_service_proto_rawDescData
}

var file_snapshot_service_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_snapshot_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_snapshot_service_proto_goTypes = []any{
	(OperationType)(0),       // 0: cloud.vmm.OperationType
	(SupportedFeatures)(0),   // 1: cloud.vmm.SupportedFeatures
	(AgentErrorCode)(0),      // 2: cloud.vmm.AgentErrorCode
	(*SnapshotRequest)(nil),  // 3: cloud.vmm.SnapshotRequest
	(*ServerInfo)(nil),       // 4: cloud.vmm.ServerInfo
	(*SnapshotResponse)(nil), // 5: cloud.vmm.SnapshotResponse
	(*GuestReady)(nil),       // 6: cloud.vmm.GuestReady
	(*GuestMessage)(nil),     // 7: cloud.vmm.GuestMessage
	(*ServerAck)(nil),        // 8: cloud.vmm.ServerAck
}
var file_snapshot_service_proto_depIdxs = []int32{
	0, // 0: cloud.vmm.SnapshotRequest.type:type_name -> cloud.vmm.OperationType
	1, // 1: cloud.vmm.ServerInfo.supported_features:type_name -> cloud.vmm.SupportedFeatures
	2, // 2: cloud.vmm.SnapshotResponse.agent_return_code:type_name -> cloud.vmm.AgentErrorCode
	0, // 3: cloud.vmm.SnapshotResponse.type:type_name -> cloud.vmm.OperationType
	3, // 4: cloud.vmm.GuestMessage.snapshot_request:type_name -> cloud.vmm.SnapshotRequest
	4, // 5: cloud.vmm.GuestMessage.server_info:type_name -> cloud.vmm.ServerInfo
	6, // 6: cloud.vmm.SnapshotService.CreateConnection:input_type -> cloud.vmm.GuestReady
	5, // 7: cloud.vmm.SnapshotService.HandleResponsesFromGuest:input_type -> cloud.vmm.SnapshotResponse
	7, // 8: cloud.vmm.SnapshotService.CreateConnection:output_type -> cloud.vmm.GuestMessage
	8, // 9: cloud.vmm.SnapshotService.HandleResponsesFromGuest:output_type -> cloud.vmm.ServerAck
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_snapshot_service_proto_init() }
func file_snapshot_service_proto_init() {
	if File_snapshot_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_snapshot_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SnapshotRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_snapshot_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ServerInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_snapshot_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SnapshotResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_snapshot_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GuestReady); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_snapshot_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GuestMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_snapshot_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ServerAck); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_snapshot_service_proto_msgTypes[4].OneofWrappers = []any{
		(*GuestMessage_SnapshotRequest)(nil),
		(*GuestMessage_ServerInfo)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_snapshot_service_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_snapshot_service_proto_goTypes,
		DependencyIndexes: file_snapshot_service_proto_depIdxs,
		EnumInfos:         file_snapshot_service_proto_enumTypes,
		MessageInfos:      file_snapshot_service_proto_msgTypes,
	}.Build()
	File_snapshot_service_proto = out.File
	file_snapshot_service_proto_rawDesc = nil
	file_snapshot_service_proto_goTypes = nil
	file_snapshot_service_proto_depIdxs = nil
}
