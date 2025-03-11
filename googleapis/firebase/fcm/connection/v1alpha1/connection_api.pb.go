// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.24.4
// source: google/firebase/fcm/connection/v1alpha1/connection_api.proto

package connection

import (
	context "context"
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request sent to FCM from the connected client.
type UpstreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of request the client is making to FCM.
	//
	// Types that are assignable to RequestType:
	//
	//	*UpstreamRequest_Ack
	RequestType isUpstreamRequest_RequestType `protobuf_oneof:"request_type"`
}

func (x *UpstreamRequest) Reset() {
	*x = UpstreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpstreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpstreamRequest) ProtoMessage() {}

func (x *UpstreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpstreamRequest.ProtoReflect.Descriptor instead.
func (*UpstreamRequest) Descriptor() ([]byte, []int) {
	return file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_rawDescGZIP(), []int{0}
}

func (m *UpstreamRequest) GetRequestType() isUpstreamRequest_RequestType {
	if m != nil {
		return m.RequestType
	}
	return nil
}

func (x *UpstreamRequest) GetAck() *Ack {
	if x, ok := x.GetRequestType().(*UpstreamRequest_Ack); ok {
		return x.Ack
	}
	return nil
}

type isUpstreamRequest_RequestType interface {
	isUpstreamRequest_RequestType()
}

type UpstreamRequest_Ack struct {
	// Message acknowledgement.
	Ack *Ack `protobuf:"bytes,1,opt,name=ack,proto3,oneof"`
}

func (*UpstreamRequest_Ack) isUpstreamRequest_RequestType() {}

// Response sent to the connected client from FCM.
type DownstreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of response FCM is sending to the client.
	//
	// Types that are assignable to ResponseType:
	//
	//	*DownstreamResponse_Message
	ResponseType isDownstreamResponse_ResponseType `protobuf_oneof:"response_type"`
}

func (x *DownstreamResponse) Reset() {
	*x = DownstreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownstreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownstreamResponse) ProtoMessage() {}

func (x *DownstreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownstreamResponse.ProtoReflect.Descriptor instead.
func (*DownstreamResponse) Descriptor() ([]byte, []int) {
	return file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_rawDescGZIP(), []int{1}
}

func (m *DownstreamResponse) GetResponseType() isDownstreamResponse_ResponseType {
	if m != nil {
		return m.ResponseType
	}
	return nil
}

func (x *DownstreamResponse) GetMessage() *Message {
	if x, ok := x.GetResponseType().(*DownstreamResponse_Message); ok {
		return x.Message
	}
	return nil
}

type isDownstreamResponse_ResponseType interface {
	isDownstreamResponse_ResponseType()
}

type DownstreamResponse_Message struct {
	// Message sent to FCM via the [Send
	// API](https://firebase.google.com/docs/cloud-messaging/send-message)
	// targeting this client.
	Message *Message `protobuf:"bytes,1,opt,name=message,proto3,oneof"`
}

func (*DownstreamResponse_Message) isDownstreamResponse_ResponseType() {}

// Acknowledgement to indicate a client successfully received an FCM message.
//
// If a message is not acked, FCM will continously resend the message until
// it expires. Duplicate delivery in this case is working as intended.
type Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id of message being acknowledged
	MessageId string `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
}

func (x *Ack) Reset() {
	*x = Ack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_rawDescGZIP(), []int{2}
}

func (x *Ack) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

// Message created through the [Send
// API](https://firebase.google.com/docs/reference/fcm/rest/v1/projects.messages#resource-message).
type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The identifier of the message. Used to ack the message.
	MessageId string `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	// Time the message was received in FCM.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Expiry time of the message. Currently it is always 4 weeks.
	ExpireTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
	// The arbitrary payload set in the [Send
	// API](https://firebase.google.com/docs/reference/fcm/rest/v1/projects.messages#resource-message).
	Data map[string]string `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_rawDescGZIP(), []int{3}
}

func (x *Message) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *Message) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Message) GetExpireTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpireTime
	}
	return nil
}

func (x *Message) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_google_firebase_fcm_connection_v1alpha1_connection_api_proto protoreflect.FileDescriptor

var file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73,
	0x65, 0x2f, 0x66, 0x63, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x27,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x66, 0x63, 0x6d, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a, 0x0f, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x03, 0x61, 0x63, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x66, 0x63, 0x6d, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x41, 0x63, 0x6b, 0x48, 0x00, 0x52, 0x03, 0x61, 0x63, 0x6b, 0x42, 0x0e, 0x0a, 0x0c, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x73, 0x0a, 0x12, 0x44,
	0x6f, 0x77, 0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4c, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x69, 0x72, 0x65,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x66, 0x63, 0x6d, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42,
	0x0f, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x24, 0x0a, 0x03, 0x41, 0x63, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0xab, 0x02, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b,
	0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x4e, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x66, 0x63, 0x6d, 0x2e,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x44,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x32, 0x98, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x41, 0x70, 0x69, 0x12, 0x86, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x12, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x69, 0x72, 0x65,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x66, 0x63, 0x6d, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x66,
	0x63, 0x6d, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42,
	0x82, 0x01, 0x0a, 0x2b, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66,
	0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x66, 0x63, 0x6d, 0x2e, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x50,
	0x01, 0x5a, 0x51, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73,
	0x65, 0x2f, 0x66, 0x63, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_rawDescOnce sync.Once
	file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_rawDescData = file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_rawDesc
)

func file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_rawDescGZIP() []byte {
	file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_rawDescOnce.Do(func() {
		file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_rawDescData)
	})
	return file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_rawDescData
}

var file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_goTypes = []interface{}{
	(*UpstreamRequest)(nil),       // 0: google.firebase.fcm.connection.v1alpha1.UpstreamRequest
	(*DownstreamResponse)(nil),    // 1: google.firebase.fcm.connection.v1alpha1.DownstreamResponse
	(*Ack)(nil),                   // 2: google.firebase.fcm.connection.v1alpha1.Ack
	(*Message)(nil),               // 3: google.firebase.fcm.connection.v1alpha1.Message
	nil,                           // 4: google.firebase.fcm.connection.v1alpha1.Message.DataEntry
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_depIdxs = []int32{
	2, // 0: google.firebase.fcm.connection.v1alpha1.UpstreamRequest.ack:type_name -> google.firebase.fcm.connection.v1alpha1.Ack
	3, // 1: google.firebase.fcm.connection.v1alpha1.DownstreamResponse.message:type_name -> google.firebase.fcm.connection.v1alpha1.Message
	5, // 2: google.firebase.fcm.connection.v1alpha1.Message.create_time:type_name -> google.protobuf.Timestamp
	5, // 3: google.firebase.fcm.connection.v1alpha1.Message.expire_time:type_name -> google.protobuf.Timestamp
	4, // 4: google.firebase.fcm.connection.v1alpha1.Message.data:type_name -> google.firebase.fcm.connection.v1alpha1.Message.DataEntry
	0, // 5: google.firebase.fcm.connection.v1alpha1.ConnectionApi.Connect:input_type -> google.firebase.fcm.connection.v1alpha1.UpstreamRequest
	1, // 6: google.firebase.fcm.connection.v1alpha1.ConnectionApi.Connect:output_type -> google.firebase.fcm.connection.v1alpha1.DownstreamResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_init() }
func file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_init() {
	if File_google_firebase_fcm_connection_v1alpha1_connection_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpstreamRequest); i {
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
		file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownstreamResponse); i {
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
		file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ack); i {
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
		file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
	file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*UpstreamRequest_Ack)(nil),
	}
	file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*DownstreamResponse_Message)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_goTypes,
		DependencyIndexes: file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_depIdxs,
		MessageInfos:      file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_msgTypes,
	}.Build()
	File_google_firebase_fcm_connection_v1alpha1_connection_api_proto = out.File
	file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_rawDesc = nil
	file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_goTypes = nil
	file_google_firebase_fcm_connection_v1alpha1_connection_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ConnectionApiClient is the client API for ConnectionApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConnectionApiClient interface {
	// Creates a streaming connection with FCM to send messages and their
	// respective ACKs.
	//
	// The client credentials need to be passed in the [gRPC
	// Metadata](https://grpc.io/docs/guides/concepts.html#metadata). The Format
	// of the header is:
	//
	//	Key: "authorization"
	//	Value: "Checkin [client_id:secret]"
	//
	// The project's API key also needs to be sent to authorize the project.
	// That can be set in the X-Goog-Api-Key Metadata header.
	Connect(ctx context.Context, opts ...grpc.CallOption) (ConnectionApi_ConnectClient, error)
}

type connectionApiClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectionApiClient(cc grpc.ClientConnInterface) ConnectionApiClient {
	return &connectionApiClient{cc}
}

func (c *connectionApiClient) Connect(ctx context.Context, opts ...grpc.CallOption) (ConnectionApi_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ConnectionApi_serviceDesc.Streams[0], "/google.firebase.fcm.connection.v1alpha1.ConnectionApi/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &connectionApiConnectClient{stream}
	return x, nil
}

type ConnectionApi_ConnectClient interface {
	Send(*UpstreamRequest) error
	Recv() (*DownstreamResponse, error)
	grpc.ClientStream
}

type connectionApiConnectClient struct {
	grpc.ClientStream
}

func (x *connectionApiConnectClient) Send(m *UpstreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *connectionApiConnectClient) Recv() (*DownstreamResponse, error) {
	m := new(DownstreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ConnectionApiServer is the server API for ConnectionApi service.
type ConnectionApiServer interface {
	// Creates a streaming connection with FCM to send messages and their
	// respective ACKs.
	//
	// The client credentials need to be passed in the [gRPC
	// Metadata](https://grpc.io/docs/guides/concepts.html#metadata). The Format
	// of the header is:
	//
	//	Key: "authorization"
	//	Value: "Checkin [client_id:secret]"
	//
	// The project's API key also needs to be sent to authorize the project.
	// That can be set in the X-Goog-Api-Key Metadata header.
	Connect(ConnectionApi_ConnectServer) error
}

// UnimplementedConnectionApiServer can be embedded to have forward compatible implementations.
type UnimplementedConnectionApiServer struct {
}

func (*UnimplementedConnectionApiServer) Connect(ConnectionApi_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}

func RegisterConnectionApiServer(s *grpc.Server, srv ConnectionApiServer) {
	s.RegisterService(&_ConnectionApi_serviceDesc, srv)
}

func _ConnectionApi_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ConnectionApiServer).Connect(&connectionApiConnectServer{stream})
}

type ConnectionApi_ConnectServer interface {
	Send(*DownstreamResponse) error
	Recv() (*UpstreamRequest, error)
	grpc.ServerStream
}

type connectionApiConnectServer struct {
	grpc.ServerStream
}

func (x *connectionApiConnectServer) Send(m *DownstreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *connectionApiConnectServer) Recv() (*UpstreamRequest, error) {
	m := new(UpstreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ConnectionApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.firebase.fcm.connection.v1alpha1.ConnectionApi",
	HandlerType: (*ConnectionApiServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _ConnectionApi_Connect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "google/firebase/fcm/connection/v1alpha1/connection_api.proto",
}
