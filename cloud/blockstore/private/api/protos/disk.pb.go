// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.0
// source: cloud/blockstore/private/api/protos/disk.proto

package protos

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

type TDiskRegistryChangeStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ChangeState:
	//	*TDiskRegistryChangeStateRequest_ChangeDeviceState
	//	*TDiskRegistryChangeStateRequest_ChangeAgentState
	//	*TDiskRegistryChangeStateRequest_DisableAgent
	ChangeState isTDiskRegistryChangeStateRequest_ChangeState `protobuf_oneof:"ChangeState"`
	Message     string                                        `protobuf:"bytes,101,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *TDiskRegistryChangeStateRequest) Reset() {
	*x = TDiskRegistryChangeStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_blockstore_private_api_protos_disk_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TDiskRegistryChangeStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TDiskRegistryChangeStateRequest) ProtoMessage() {}

func (x *TDiskRegistryChangeStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_blockstore_private_api_protos_disk_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TDiskRegistryChangeStateRequest.ProtoReflect.Descriptor instead.
func (*TDiskRegistryChangeStateRequest) Descriptor() ([]byte, []int) {
	return file_cloud_blockstore_private_api_protos_disk_proto_rawDescGZIP(), []int{0}
}

func (m *TDiskRegistryChangeStateRequest) GetChangeState() isTDiskRegistryChangeStateRequest_ChangeState {
	if m != nil {
		return m.ChangeState
	}
	return nil
}

func (x *TDiskRegistryChangeStateRequest) GetChangeDeviceState() *TDiskRegistryChangeStateRequest_TChangeDeviceState {
	if x, ok := x.GetChangeState().(*TDiskRegistryChangeStateRequest_ChangeDeviceState); ok {
		return x.ChangeDeviceState
	}
	return nil
}

func (x *TDiskRegistryChangeStateRequest) GetChangeAgentState() *TDiskRegistryChangeStateRequest_TChangeAgentState {
	if x, ok := x.GetChangeState().(*TDiskRegistryChangeStateRequest_ChangeAgentState); ok {
		return x.ChangeAgentState
	}
	return nil
}

func (x *TDiskRegistryChangeStateRequest) GetDisableAgent() *TDiskRegistryChangeStateRequest_TDisableAgent {
	if x, ok := x.GetChangeState().(*TDiskRegistryChangeStateRequest_DisableAgent); ok {
		return x.DisableAgent
	}
	return nil
}

func (x *TDiskRegistryChangeStateRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type isTDiskRegistryChangeStateRequest_ChangeState interface {
	isTDiskRegistryChangeStateRequest_ChangeState()
}

type TDiskRegistryChangeStateRequest_ChangeDeviceState struct {
	ChangeDeviceState *TDiskRegistryChangeStateRequest_TChangeDeviceState `protobuf:"bytes,1,opt,name=ChangeDeviceState,proto3,oneof"`
}

type TDiskRegistryChangeStateRequest_ChangeAgentState struct {
	ChangeAgentState *TDiskRegistryChangeStateRequest_TChangeAgentState `protobuf:"bytes,2,opt,name=ChangeAgentState,proto3,oneof"`
}

type TDiskRegistryChangeStateRequest_DisableAgent struct {
	DisableAgent *TDiskRegistryChangeStateRequest_TDisableAgent `protobuf:"bytes,3,opt,name=DisableAgent,proto3,oneof"`
}

func (*TDiskRegistryChangeStateRequest_ChangeDeviceState) isTDiskRegistryChangeStateRequest_ChangeState() {
}

func (*TDiskRegistryChangeStateRequest_ChangeAgentState) isTDiskRegistryChangeStateRequest_ChangeState() {
}

func (*TDiskRegistryChangeStateRequest_DisableAgent) isTDiskRegistryChangeStateRequest_ChangeState() {
}

type TDiskRegistryChangeStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TDiskRegistryChangeStateResponse) Reset() {
	*x = TDiskRegistryChangeStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_blockstore_private_api_protos_disk_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TDiskRegistryChangeStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TDiskRegistryChangeStateResponse) ProtoMessage() {}

func (x *TDiskRegistryChangeStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_blockstore_private_api_protos_disk_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TDiskRegistryChangeStateResponse.ProtoReflect.Descriptor instead.
func (*TDiskRegistryChangeStateResponse) Descriptor() ([]byte, []int) {
	return file_cloud_blockstore_private_api_protos_disk_proto_rawDescGZIP(), []int{1}
}

type TReassignDiskRegistryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SystemKind string `protobuf:"bytes,1,opt,name=SystemKind,proto3" json:"SystemKind,omitempty"`
	LogKind    string `protobuf:"bytes,2,opt,name=LogKind,proto3" json:"LogKind,omitempty"`
	IndexKind  string `protobuf:"bytes,3,opt,name=IndexKind,proto3" json:"IndexKind,omitempty"`
}

func (x *TReassignDiskRegistryRequest) Reset() {
	*x = TReassignDiskRegistryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_blockstore_private_api_protos_disk_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TReassignDiskRegistryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TReassignDiskRegistryRequest) ProtoMessage() {}

func (x *TReassignDiskRegistryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_blockstore_private_api_protos_disk_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TReassignDiskRegistryRequest.ProtoReflect.Descriptor instead.
func (*TReassignDiskRegistryRequest) Descriptor() ([]byte, []int) {
	return file_cloud_blockstore_private_api_protos_disk_proto_rawDescGZIP(), []int{2}
}

func (x *TReassignDiskRegistryRequest) GetSystemKind() string {
	if x != nil {
		return x.SystemKind
	}
	return ""
}

func (x *TReassignDiskRegistryRequest) GetLogKind() string {
	if x != nil {
		return x.LogKind
	}
	return ""
}

func (x *TReassignDiskRegistryRequest) GetIndexKind() string {
	if x != nil {
		return x.IndexKind
	}
	return ""
}

type TReassignDiskRegistryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TReassignDiskRegistryResponse) Reset() {
	*x = TReassignDiskRegistryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_blockstore_private_api_protos_disk_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TReassignDiskRegistryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TReassignDiskRegistryResponse) ProtoMessage() {}

func (x *TReassignDiskRegistryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_blockstore_private_api_protos_disk_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TReassignDiskRegistryResponse.ProtoReflect.Descriptor instead.
func (*TReassignDiskRegistryResponse) Descriptor() ([]byte, []int) {
	return file_cloud_blockstore_private_api_protos_disk_proto_rawDescGZIP(), []int{3}
}

type TGetDiskRegistryTabletInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TGetDiskRegistryTabletInfoRequest) Reset() {
	*x = TGetDiskRegistryTabletInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_blockstore_private_api_protos_disk_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TGetDiskRegistryTabletInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TGetDiskRegistryTabletInfoRequest) ProtoMessage() {}

func (x *TGetDiskRegistryTabletInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_blockstore_private_api_protos_disk_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TGetDiskRegistryTabletInfoRequest.ProtoReflect.Descriptor instead.
func (*TGetDiskRegistryTabletInfoRequest) Descriptor() ([]byte, []int) {
	return file_cloud_blockstore_private_api_protos_disk_proto_rawDescGZIP(), []int{4}
}

type TGetDiskRegistryTabletInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TabletId uint64 `protobuf:"varint,1,opt,name=TabletId,proto3" json:"TabletId,omitempty"`
}

func (x *TGetDiskRegistryTabletInfoResponse) Reset() {
	*x = TGetDiskRegistryTabletInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_blockstore_private_api_protos_disk_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TGetDiskRegistryTabletInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TGetDiskRegistryTabletInfoResponse) ProtoMessage() {}

func (x *TGetDiskRegistryTabletInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_blockstore_private_api_protos_disk_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TGetDiskRegistryTabletInfoResponse.ProtoReflect.Descriptor instead.
func (*TGetDiskRegistryTabletInfoResponse) Descriptor() ([]byte, []int) {
	return file_cloud_blockstore_private_api_protos_disk_proto_rawDescGZIP(), []int{5}
}

func (x *TGetDiskRegistryTabletInfoResponse) GetTabletId() uint64 {
	if x != nil {
		return x.TabletId
	}
	return 0
}

type TDiskRegistryChangeStateRequest_TChangeDeviceState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceUUID string `protobuf:"bytes,1,opt,name=DeviceUUID,proto3" json:"DeviceUUID,omitempty"`
	State      uint32 `protobuf:"varint,2,opt,name=State,proto3" json:"State,omitempty"`
}

func (x *TDiskRegistryChangeStateRequest_TChangeDeviceState) Reset() {
	*x = TDiskRegistryChangeStateRequest_TChangeDeviceState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_blockstore_private_api_protos_disk_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TDiskRegistryChangeStateRequest_TChangeDeviceState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TDiskRegistryChangeStateRequest_TChangeDeviceState) ProtoMessage() {}

func (x *TDiskRegistryChangeStateRequest_TChangeDeviceState) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_blockstore_private_api_protos_disk_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TDiskRegistryChangeStateRequest_TChangeDeviceState.ProtoReflect.Descriptor instead.
func (*TDiskRegistryChangeStateRequest_TChangeDeviceState) Descriptor() ([]byte, []int) {
	return file_cloud_blockstore_private_api_protos_disk_proto_rawDescGZIP(), []int{0, 0}
}

func (x *TDiskRegistryChangeStateRequest_TChangeDeviceState) GetDeviceUUID() string {
	if x != nil {
		return x.DeviceUUID
	}
	return ""
}

func (x *TDiskRegistryChangeStateRequest_TChangeDeviceState) GetState() uint32 {
	if x != nil {
		return x.State
	}
	return 0
}

type TDiskRegistryChangeStateRequest_TChangeAgentState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentId string `protobuf:"bytes,1,opt,name=AgentId,proto3" json:"AgentId,omitempty"`
	State   uint32 `protobuf:"varint,2,opt,name=State,proto3" json:"State,omitempty"`
}

func (x *TDiskRegistryChangeStateRequest_TChangeAgentState) Reset() {
	*x = TDiskRegistryChangeStateRequest_TChangeAgentState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_blockstore_private_api_protos_disk_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TDiskRegistryChangeStateRequest_TChangeAgentState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TDiskRegistryChangeStateRequest_TChangeAgentState) ProtoMessage() {}

func (x *TDiskRegistryChangeStateRequest_TChangeAgentState) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_blockstore_private_api_protos_disk_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TDiskRegistryChangeStateRequest_TChangeAgentState.ProtoReflect.Descriptor instead.
func (*TDiskRegistryChangeStateRequest_TChangeAgentState) Descriptor() ([]byte, []int) {
	return file_cloud_blockstore_private_api_protos_disk_proto_rawDescGZIP(), []int{0, 1}
}

func (x *TDiskRegistryChangeStateRequest_TChangeAgentState) GetAgentId() string {
	if x != nil {
		return x.AgentId
	}
	return ""
}

func (x *TDiskRegistryChangeStateRequest_TChangeAgentState) GetState() uint32 {
	if x != nil {
		return x.State
	}
	return 0
}

type TDiskRegistryChangeStateRequest_TDisableAgent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentId     string   `protobuf:"bytes,1,opt,name=AgentId,proto3" json:"AgentId,omitempty"`
	DeviceUUIDs []string `protobuf:"bytes,2,rep,name=DeviceUUIDs,proto3" json:"DeviceUUIDs,omitempty"`
}

func (x *TDiskRegistryChangeStateRequest_TDisableAgent) Reset() {
	*x = TDiskRegistryChangeStateRequest_TDisableAgent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloud_blockstore_private_api_protos_disk_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TDiskRegistryChangeStateRequest_TDisableAgent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TDiskRegistryChangeStateRequest_TDisableAgent) ProtoMessage() {}

func (x *TDiskRegistryChangeStateRequest_TDisableAgent) ProtoReflect() protoreflect.Message {
	mi := &file_cloud_blockstore_private_api_protos_disk_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TDiskRegistryChangeStateRequest_TDisableAgent.ProtoReflect.Descriptor instead.
func (*TDiskRegistryChangeStateRequest_TDisableAgent) Descriptor() ([]byte, []int) {
	return file_cloud_blockstore_private_api_protos_disk_proto_rawDescGZIP(), []int{0, 2}
}

func (x *TDiskRegistryChangeStateRequest_TDisableAgent) GetAgentId() string {
	if x != nil {
		return x.AgentId
	}
	return ""
}

func (x *TDiskRegistryChangeStateRequest_TDisableAgent) GetDeviceUUIDs() []string {
	if x != nil {
		return x.DeviceUUIDs
	}
	return nil
}

var File_cloud_blockstore_private_api_protos_disk_proto protoreflect.FileDescriptor

var file_cloud_blockstore_private_api_protos_disk_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x20, 0x4e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xaa, 0x05, 0x0a, 0x1f, 0x54, 0x44, 0x69, 0x73, 0x6b, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x84, 0x01, 0x0a, 0x11, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x54, 0x2e, 0x4e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4e, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x44, 0x69, 0x73, 0x6b, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x11, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x81, 0x01,
	0x0a, 0x10, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x53, 0x2e, 0x4e, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x4e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4e, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x44, 0x69, 0x73,
	0x6b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52,
	0x10, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x75, 0x0a, 0x0c, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4f, 0x2e, 0x4e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x4e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4e, 0x50, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x44, 0x69, 0x73, 0x6b,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x44, 0x69, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0c, 0x44, 0x69, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x1a, 0x4a, 0x0a, 0x12, 0x54, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x55, 0x55, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x1a, 0x43,
	0x0a, 0x11, 0x54, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x1a, 0x4b, 0x0a, 0x0d, 0x54, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x55, 0x55, 0x49, 0x44, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x55, 0x55, 0x49, 0x44, 0x73,
	0x42, 0x0d, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22,
	0x22, 0x0a, 0x20, 0x54, 0x44, 0x69, 0x73, 0x6b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x76, 0x0a, 0x1c, 0x54, 0x52, 0x65, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x44, 0x69, 0x73, 0x6b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4b, 0x69, 0x6e,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4b,
	0x69, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x4c, 0x6f, 0x67, 0x4b, 0x69, 0x6e, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4c, 0x6f, 0x67, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4b, 0x69, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4b, 0x69, 0x6e, 0x64, 0x22, 0x1f, 0x0a, 0x1d, 0x54,
	0x52, 0x65, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x44, 0x69, 0x73, 0x6b, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x0a, 0x21,
	0x54, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x6b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x40, 0x0a, 0x22, 0x54, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x6b, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x74, 0x49, 0x64, 0x42, 0x36, 0x5a, 0x34, 0x61, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d,
	0x74, 0x65, 0x61, 0x6d, 0x2e, 0x72, 0x75, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_cloud_blockstore_private_api_protos_disk_proto_rawDescOnce sync.Once
	file_cloud_blockstore_private_api_protos_disk_proto_rawDescData = file_cloud_blockstore_private_api_protos_disk_proto_rawDesc
)

func file_cloud_blockstore_private_api_protos_disk_proto_rawDescGZIP() []byte {
	file_cloud_blockstore_private_api_protos_disk_proto_rawDescOnce.Do(func() {
		file_cloud_blockstore_private_api_protos_disk_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloud_blockstore_private_api_protos_disk_proto_rawDescData)
	})
	return file_cloud_blockstore_private_api_protos_disk_proto_rawDescData
}

var file_cloud_blockstore_private_api_protos_disk_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_cloud_blockstore_private_api_protos_disk_proto_goTypes = []interface{}{
	(*TDiskRegistryChangeStateRequest)(nil),                    // 0: NCloud.NBlockStore.NPrivateProto.TDiskRegistryChangeStateRequest
	(*TDiskRegistryChangeStateResponse)(nil),                   // 1: NCloud.NBlockStore.NPrivateProto.TDiskRegistryChangeStateResponse
	(*TReassignDiskRegistryRequest)(nil),                       // 2: NCloud.NBlockStore.NPrivateProto.TReassignDiskRegistryRequest
	(*TReassignDiskRegistryResponse)(nil),                      // 3: NCloud.NBlockStore.NPrivateProto.TReassignDiskRegistryResponse
	(*TGetDiskRegistryTabletInfoRequest)(nil),                  // 4: NCloud.NBlockStore.NPrivateProto.TGetDiskRegistryTabletInfoRequest
	(*TGetDiskRegistryTabletInfoResponse)(nil),                 // 5: NCloud.NBlockStore.NPrivateProto.TGetDiskRegistryTabletInfoResponse
	(*TDiskRegistryChangeStateRequest_TChangeDeviceState)(nil), // 6: NCloud.NBlockStore.NPrivateProto.TDiskRegistryChangeStateRequest.TChangeDeviceState
	(*TDiskRegistryChangeStateRequest_TChangeAgentState)(nil),  // 7: NCloud.NBlockStore.NPrivateProto.TDiskRegistryChangeStateRequest.TChangeAgentState
	(*TDiskRegistryChangeStateRequest_TDisableAgent)(nil),      // 8: NCloud.NBlockStore.NPrivateProto.TDiskRegistryChangeStateRequest.TDisableAgent
}
var file_cloud_blockstore_private_api_protos_disk_proto_depIdxs = []int32{
	6, // 0: NCloud.NBlockStore.NPrivateProto.TDiskRegistryChangeStateRequest.ChangeDeviceState:type_name -> NCloud.NBlockStore.NPrivateProto.TDiskRegistryChangeStateRequest.TChangeDeviceState
	7, // 1: NCloud.NBlockStore.NPrivateProto.TDiskRegistryChangeStateRequest.ChangeAgentState:type_name -> NCloud.NBlockStore.NPrivateProto.TDiskRegistryChangeStateRequest.TChangeAgentState
	8, // 2: NCloud.NBlockStore.NPrivateProto.TDiskRegistryChangeStateRequest.DisableAgent:type_name -> NCloud.NBlockStore.NPrivateProto.TDiskRegistryChangeStateRequest.TDisableAgent
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_cloud_blockstore_private_api_protos_disk_proto_init() }
func file_cloud_blockstore_private_api_protos_disk_proto_init() {
	if File_cloud_blockstore_private_api_protos_disk_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloud_blockstore_private_api_protos_disk_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TDiskRegistryChangeStateRequest); i {
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
		file_cloud_blockstore_private_api_protos_disk_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TDiskRegistryChangeStateResponse); i {
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
		file_cloud_blockstore_private_api_protos_disk_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TReassignDiskRegistryRequest); i {
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
		file_cloud_blockstore_private_api_protos_disk_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TReassignDiskRegistryResponse); i {
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
		file_cloud_blockstore_private_api_protos_disk_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TGetDiskRegistryTabletInfoRequest); i {
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
		file_cloud_blockstore_private_api_protos_disk_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TGetDiskRegistryTabletInfoResponse); i {
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
		file_cloud_blockstore_private_api_protos_disk_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TDiskRegistryChangeStateRequest_TChangeDeviceState); i {
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
		file_cloud_blockstore_private_api_protos_disk_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TDiskRegistryChangeStateRequest_TChangeAgentState); i {
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
		file_cloud_blockstore_private_api_protos_disk_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TDiskRegistryChangeStateRequest_TDisableAgent); i {
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
	file_cloud_blockstore_private_api_protos_disk_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*TDiskRegistryChangeStateRequest_ChangeDeviceState)(nil),
		(*TDiskRegistryChangeStateRequest_ChangeAgentState)(nil),
		(*TDiskRegistryChangeStateRequest_DisableAgent)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloud_blockstore_private_api_protos_disk_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloud_blockstore_private_api_protos_disk_proto_goTypes,
		DependencyIndexes: file_cloud_blockstore_private_api_protos_disk_proto_depIdxs,
		MessageInfos:      file_cloud_blockstore_private_api_protos_disk_proto_msgTypes,
	}.Build()
	File_cloud_blockstore_private_api_protos_disk_proto = out.File
	file_cloud_blockstore_private_api_protos_disk_proto_rawDesc = nil
	file_cloud_blockstore_private_api_protos_disk_proto_goTypes = nil
	file_cloud_blockstore_private_api_protos_disk_proto_depIdxs = nil
}