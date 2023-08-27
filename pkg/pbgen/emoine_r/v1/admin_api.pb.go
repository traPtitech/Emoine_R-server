// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: emoine_r/v1/admin_api.proto

package emoine_rv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoId     string `protobuf:"bytes,1,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateEventRequest) Reset() {
	*x = CreateEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emoine_r_v1_admin_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventRequest) ProtoMessage() {}

func (x *CreateEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_emoine_r_v1_admin_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventRequest.ProtoReflect.Descriptor instead.
func (*CreateEventRequest) Descriptor() ([]byte, []int) {
	return file_emoine_r_v1_admin_api_proto_rawDescGZIP(), []int{0}
}

func (x *CreateEventRequest) GetVideoId() string {
	if x != nil {
		return x.VideoId
	}
	return ""
}

func (x *CreateEventRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *CreateEventResponse) Reset() {
	*x = CreateEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emoine_r_v1_admin_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventResponse) ProtoMessage() {}

func (x *CreateEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_emoine_r_v1_admin_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventResponse.ProtoReflect.Descriptor instead.
func (*CreateEventResponse) Descriptor() ([]byte, []int) {
	return file_emoine_r_v1_admin_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEventResponse) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

type UpdateEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId     string  `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	VideoId     *string `protobuf:"bytes,2,opt,name=video_id,json=videoId,proto3,oneof" json:"video_id,omitempty"`
	Description *string `protobuf:"bytes,3,opt,name=description,proto3,oneof" json:"description,omitempty"`
}

func (x *UpdateEventRequest) Reset() {
	*x = UpdateEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emoine_r_v1_admin_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEventRequest) ProtoMessage() {}

func (x *UpdateEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_emoine_r_v1_admin_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEventRequest.ProtoReflect.Descriptor instead.
func (*UpdateEventRequest) Descriptor() ([]byte, []int) {
	return file_emoine_r_v1_admin_api_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateEventRequest) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *UpdateEventRequest) GetVideoId() string {
	if x != nil && x.VideoId != nil {
		return *x.VideoId
	}
	return ""
}

func (x *UpdateEventRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

type DeleteEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId string `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
}

func (x *DeleteEventRequest) Reset() {
	*x = DeleteEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emoine_r_v1_admin_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEventRequest) ProtoMessage() {}

func (x *DeleteEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_emoine_r_v1_admin_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEventRequest.ProtoReflect.Descriptor instead.
func (*DeleteEventRequest) Descriptor() ([]byte, []int) {
	return file_emoine_r_v1_admin_api_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteEventRequest) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

type GetTokensRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId string `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
}

func (x *GetTokensRequest) Reset() {
	*x = GetTokensRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emoine_r_v1_admin_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokensRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokensRequest) ProtoMessage() {}

func (x *GetTokensRequest) ProtoReflect() protoreflect.Message {
	mi := &file_emoine_r_v1_admin_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTokensRequest.ProtoReflect.Descriptor instead.
func (*GetTokensRequest) Descriptor() ([]byte, []int) {
	return file_emoine_r_v1_admin_api_proto_rawDescGZIP(), []int{4}
}

func (x *GetTokensRequest) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

type GetTokensResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tokens []*Token `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty"`
}

func (x *GetTokensResponse) Reset() {
	*x = GetTokensResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emoine_r_v1_admin_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokensResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokensResponse) ProtoMessage() {}

func (x *GetTokensResponse) ProtoReflect() protoreflect.Message {
	mi := &file_emoine_r_v1_admin_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTokensResponse.ProtoReflect.Descriptor instead.
func (*GetTokensResponse) Descriptor() ([]byte, []int) {
	return file_emoine_r_v1_admin_api_proto_rawDescGZIP(), []int{5}
}

func (x *GetTokensResponse) GetTokens() []*Token {
	if x != nil {
		return x.Tokens
	}
	return nil
}

type GenerateTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId     string                 `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	Username    string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Description string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	ExpireAt    *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=expire_at,json=expireAt,proto3" json:"expire_at,omitempty"`
}

func (x *GenerateTokenRequest) Reset() {
	*x = GenerateTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emoine_r_v1_admin_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTokenRequest) ProtoMessage() {}

func (x *GenerateTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_emoine_r_v1_admin_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTokenRequest.ProtoReflect.Descriptor instead.
func (*GenerateTokenRequest) Descriptor() ([]byte, []int) {
	return file_emoine_r_v1_admin_api_proto_rawDescGZIP(), []int{6}
}

func (x *GenerateTokenRequest) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *GenerateTokenRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GenerateTokenRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GenerateTokenRequest) GetExpireAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpireAt
	}
	return nil
}

type GenerateTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token *Token `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GenerateTokenResponse) Reset() {
	*x = GenerateTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emoine_r_v1_admin_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTokenResponse) ProtoMessage() {}

func (x *GenerateTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_emoine_r_v1_admin_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTokenResponse.ProtoReflect.Descriptor instead.
func (*GenerateTokenResponse) Descriptor() ([]byte, []int) {
	return file_emoine_r_v1_admin_api_proto_rawDescGZIP(), []int{7}
}

func (x *GenerateTokenResponse) GetToken() *Token {
	if x != nil {
		return x.Token
	}
	return nil
}

type RevokeTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenId string `protobuf:"bytes,1,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	EventId string `protobuf:"bytes,2,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
}

func (x *RevokeTokenRequest) Reset() {
	*x = RevokeTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emoine_r_v1_admin_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevokeTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokeTokenRequest) ProtoMessage() {}

func (x *RevokeTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_emoine_r_v1_admin_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokeTokenRequest.ProtoReflect.Descriptor instead.
func (*RevokeTokenRequest) Descriptor() ([]byte, []int) {
	return file_emoine_r_v1_admin_api_proto_rawDescGZIP(), []int{8}
}

func (x *RevokeTokenRequest) GetTokenId() string {
	if x != nil {
		return x.TokenId
	}
	return ""
}

func (x *RevokeTokenRequest) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

var File_emoine_r_v1_admin_api_proto protoreflect.FileDescriptor

var file_emoine_r_v1_admin_api_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x65, 0x6d, 0x6f, 0x69, 0x6e, 0x65, 0x5f, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x65,
	0x6d, 0x6f, 0x69, 0x6e, 0x65, 0x5f, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x65, 0x6d, 0x6f, 0x69,
	0x6e, 0x65, 0x5f, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x51, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3f, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x6d,
	0x6f, 0x69, 0x6e, 0x65, 0x5f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x93, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42,
	0x0b, 0x0a, 0x09, 0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x42, 0x0e, 0x0a, 0x0c,
	0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2f, 0x0a, 0x12,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x2d, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x3f, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2a, 0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x65, 0x6d, 0x6f, 0x69, 0x6e, 0x65, 0x5f, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x22, 0xa8, 0x01,
	0x0a, 0x14, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x37, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x41, 0x74, 0x22, 0x41, 0x0a, 0x15, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x28, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x65, 0x6d, 0x6f, 0x69, 0x6e, 0x65, 0x5f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x4a, 0x0a, 0x12, 0x52,
	0x65, 0x76, 0x6f, 0x6b, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x32, 0xdf, 0x03, 0x0a, 0x0f, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x41, 0x50, 0x49, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x0b, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x2e, 0x65, 0x6d, 0x6f,
	0x69, 0x6e, 0x65, 0x5f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x65, 0x6d,
	0x6f, 0x69, 0x6e, 0x65, 0x5f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a,
	0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x2e, 0x65,
	0x6d, 0x6f, 0x69, 0x6e, 0x65, 0x5f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x46, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x2e, 0x65, 0x6d, 0x6f, 0x69, 0x6e, 0x65, 0x5f, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4a, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x1d, 0x2e, 0x65, 0x6d, 0x6f,
	0x69, 0x6e, 0x65, 0x5f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x65, 0x6d, 0x6f, 0x69,
	0x6e, 0x65, 0x5f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x0d, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x2e, 0x65, 0x6d, 0x6f,
	0x69, 0x6e, 0x65, 0x5f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x65, 0x6d, 0x6f, 0x69, 0x6e, 0x65, 0x5f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x46, 0x0a, 0x0b, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x1f, 0x2e, 0x65, 0x6d, 0x6f, 0x69, 0x6e, 0x65, 0x5f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x76, 0x6f, 0x6b, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0xaa, 0x01, 0x0a, 0x0f, 0x63, 0x6f,
	0x6d, 0x2e, 0x65, 0x6d, 0x6f, 0x69, 0x6e, 0x65, 0x5f, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x61, 0x50, 0x74,
	0x69, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x45, 0x6d, 0x6f, 0x69, 0x6e, 0x65, 0x5f, 0x52, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x62, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x6d, 0x6f, 0x69, 0x6e, 0x65, 0x5f,
	0x72, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x6d, 0x6f, 0x69, 0x6e, 0x65, 0x5f, 0x72, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x45, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x45, 0x6d, 0x6f, 0x69, 0x6e, 0x65, 0x52, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x0a, 0x45, 0x6d, 0x6f, 0x69, 0x6e, 0x65, 0x52, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x16, 0x45, 0x6d, 0x6f, 0x69, 0x6e, 0x65, 0x52, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x45, 0x6d, 0x6f, 0x69, 0x6e,
	0x65, 0x52, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_emoine_r_v1_admin_api_proto_rawDescOnce sync.Once
	file_emoine_r_v1_admin_api_proto_rawDescData = file_emoine_r_v1_admin_api_proto_rawDesc
)

func file_emoine_r_v1_admin_api_proto_rawDescGZIP() []byte {
	file_emoine_r_v1_admin_api_proto_rawDescOnce.Do(func() {
		file_emoine_r_v1_admin_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_emoine_r_v1_admin_api_proto_rawDescData)
	})
	return file_emoine_r_v1_admin_api_proto_rawDescData
}

var file_emoine_r_v1_admin_api_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_emoine_r_v1_admin_api_proto_goTypes = []interface{}{
	(*CreateEventRequest)(nil),    // 0: emoine_r.v1.CreateEventRequest
	(*CreateEventResponse)(nil),   // 1: emoine_r.v1.CreateEventResponse
	(*UpdateEventRequest)(nil),    // 2: emoine_r.v1.UpdateEventRequest
	(*DeleteEventRequest)(nil),    // 3: emoine_r.v1.DeleteEventRequest
	(*GetTokensRequest)(nil),      // 4: emoine_r.v1.GetTokensRequest
	(*GetTokensResponse)(nil),     // 5: emoine_r.v1.GetTokensResponse
	(*GenerateTokenRequest)(nil),  // 6: emoine_r.v1.GenerateTokenRequest
	(*GenerateTokenResponse)(nil), // 7: emoine_r.v1.GenerateTokenResponse
	(*RevokeTokenRequest)(nil),    // 8: emoine_r.v1.RevokeTokenRequest
	(*Event)(nil),                 // 9: emoine_r.v1.Event
	(*Token)(nil),                 // 10: emoine_r.v1.Token
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 12: google.protobuf.Empty
}
var file_emoine_r_v1_admin_api_proto_depIdxs = []int32{
	9,  // 0: emoine_r.v1.CreateEventResponse.event:type_name -> emoine_r.v1.Event
	10, // 1: emoine_r.v1.GetTokensResponse.tokens:type_name -> emoine_r.v1.Token
	11, // 2: emoine_r.v1.GenerateTokenRequest.expire_at:type_name -> google.protobuf.Timestamp
	10, // 3: emoine_r.v1.GenerateTokenResponse.token:type_name -> emoine_r.v1.Token
	0,  // 4: emoine_r.v1.AdminAPIService.CreateEvent:input_type -> emoine_r.v1.CreateEventRequest
	2,  // 5: emoine_r.v1.AdminAPIService.UpdateEvent:input_type -> emoine_r.v1.UpdateEventRequest
	3,  // 6: emoine_r.v1.AdminAPIService.DeleteEvent:input_type -> emoine_r.v1.DeleteEventRequest
	4,  // 7: emoine_r.v1.AdminAPIService.GetTokens:input_type -> emoine_r.v1.GetTokensRequest
	6,  // 8: emoine_r.v1.AdminAPIService.GenerateToken:input_type -> emoine_r.v1.GenerateTokenRequest
	8,  // 9: emoine_r.v1.AdminAPIService.RevokeToken:input_type -> emoine_r.v1.RevokeTokenRequest
	1,  // 10: emoine_r.v1.AdminAPIService.CreateEvent:output_type -> emoine_r.v1.CreateEventResponse
	12, // 11: emoine_r.v1.AdminAPIService.UpdateEvent:output_type -> google.protobuf.Empty
	12, // 12: emoine_r.v1.AdminAPIService.DeleteEvent:output_type -> google.protobuf.Empty
	5,  // 13: emoine_r.v1.AdminAPIService.GetTokens:output_type -> emoine_r.v1.GetTokensResponse
	7,  // 14: emoine_r.v1.AdminAPIService.GenerateToken:output_type -> emoine_r.v1.GenerateTokenResponse
	12, // 15: emoine_r.v1.AdminAPIService.RevokeToken:output_type -> google.protobuf.Empty
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_emoine_r_v1_admin_api_proto_init() }
func file_emoine_r_v1_admin_api_proto_init() {
	if File_emoine_r_v1_admin_api_proto != nil {
		return
	}
	file_emoine_r_v1_schema_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_emoine_r_v1_admin_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEventRequest); i {
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
		file_emoine_r_v1_admin_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEventResponse); i {
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
		file_emoine_r_v1_admin_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEventRequest); i {
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
		file_emoine_r_v1_admin_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEventRequest); i {
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
		file_emoine_r_v1_admin_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTokensRequest); i {
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
		file_emoine_r_v1_admin_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTokensResponse); i {
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
		file_emoine_r_v1_admin_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateTokenRequest); i {
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
		file_emoine_r_v1_admin_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateTokenResponse); i {
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
		file_emoine_r_v1_admin_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RevokeTokenRequest); i {
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
	file_emoine_r_v1_admin_api_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_emoine_r_v1_admin_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_emoine_r_v1_admin_api_proto_goTypes,
		DependencyIndexes: file_emoine_r_v1_admin_api_proto_depIdxs,
		MessageInfos:      file_emoine_r_v1_admin_api_proto_msgTypes,
	}.Build()
	File_emoine_r_v1_admin_api_proto = out.File
	file_emoine_r_v1_admin_api_proto_rawDesc = nil
	file_emoine_r_v1_admin_api_proto_goTypes = nil
	file_emoine_r_v1_admin_api_proto_depIdxs = nil
}
