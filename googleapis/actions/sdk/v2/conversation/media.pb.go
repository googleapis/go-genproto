// Copyright 2024 Google LLC
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.24.4
// source: google/actions/sdk/v2/conversation/prompt/content/media.proto

package conversation

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Media type of this response.
type Media_MediaType int32

const (
	// Unspecified media type.
	Media_MEDIA_TYPE_UNSPECIFIED Media_MediaType = 0
	// Audio file.
	Media_AUDIO Media_MediaType = 1
	// Response to acknowledge a media status report.
	Media_MEDIA_STATUS_ACK Media_MediaType = 2
)

// Enum value maps for Media_MediaType.
var (
	Media_MediaType_name = map[int32]string{
		0: "MEDIA_TYPE_UNSPECIFIED",
		1: "AUDIO",
		2: "MEDIA_STATUS_ACK",
	}
	Media_MediaType_value = map[string]int32{
		"MEDIA_TYPE_UNSPECIFIED": 0,
		"AUDIO":                  1,
		"MEDIA_STATUS_ACK":       2,
	}
)

func (x Media_MediaType) Enum() *Media_MediaType {
	p := new(Media_MediaType)
	*p = x
	return p
}

func (x Media_MediaType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Media_MediaType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_actions_sdk_v2_conversation_prompt_content_media_proto_enumTypes[0].Descriptor()
}

func (Media_MediaType) Type() protoreflect.EnumType {
	return &file_google_actions_sdk_v2_conversation_prompt_content_media_proto_enumTypes[0]
}

func (x Media_MediaType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Media_MediaType.Descriptor instead.
func (Media_MediaType) EnumDescriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_conversation_prompt_content_media_proto_rawDescGZIP(), []int{0, 0}
}

// Optional media control types the media response can support
type Media_OptionalMediaControls int32

const (
	// Unspecified value
	Media_OPTIONAL_MEDIA_CONTROLS_UNSPECIFIED Media_OptionalMediaControls = 0
	// Paused event. Triggered when user pauses the media.
	Media_PAUSED Media_OptionalMediaControls = 1
	// Stopped event. Triggered when user exits out of 3p session during media
	// play.
	Media_STOPPED Media_OptionalMediaControls = 2
)

// Enum value maps for Media_OptionalMediaControls.
var (
	Media_OptionalMediaControls_name = map[int32]string{
		0: "OPTIONAL_MEDIA_CONTROLS_UNSPECIFIED",
		1: "PAUSED",
		2: "STOPPED",
	}
	Media_OptionalMediaControls_value = map[string]int32{
		"OPTIONAL_MEDIA_CONTROLS_UNSPECIFIED": 0,
		"PAUSED":                              1,
		"STOPPED":                             2,
	}
)

func (x Media_OptionalMediaControls) Enum() *Media_OptionalMediaControls {
	p := new(Media_OptionalMediaControls)
	*p = x
	return p
}

func (x Media_OptionalMediaControls) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Media_OptionalMediaControls) Descriptor() protoreflect.EnumDescriptor {
	return file_google_actions_sdk_v2_conversation_prompt_content_media_proto_enumTypes[1].Descriptor()
}

func (Media_OptionalMediaControls) Type() protoreflect.EnumType {
	return &file_google_actions_sdk_v2_conversation_prompt_content_media_proto_enumTypes[1]
}

func (x Media_OptionalMediaControls) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Media_OptionalMediaControls.Descriptor instead.
func (Media_OptionalMediaControls) EnumDescriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_conversation_prompt_content_media_proto_rawDescGZIP(), []int{0, 1}
}

// Represents one media object.
// Contains information about the media, such as name, description, url, etc.
type Media struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Media type.
	MediaType Media_MediaType `protobuf:"varint,8,opt,name=media_type,json=mediaType,proto3,enum=google.actions.sdk.v2.conversation.Media_MediaType" json:"media_type,omitempty"`
	// Start offset of the first media object.
	StartOffset *durationpb.Duration `protobuf:"bytes,5,opt,name=start_offset,json=startOffset,proto3" json:"start_offset,omitempty"`
	// Optional media control types this media response session can support.
	// If set, request will be made to 3p when a certain media event happens.
	// If not set, 3p must still handle two default control type, FINISHED and
	// FAILED.
	OptionalMediaControls []Media_OptionalMediaControls `protobuf:"varint,6,rep,packed,name=optional_media_controls,json=optionalMediaControls,proto3,enum=google.actions.sdk.v2.conversation.Media_OptionalMediaControls" json:"optional_media_controls,omitempty"`
	// List of Media Objects
	MediaObjects []*MediaObject `protobuf:"bytes,7,rep,name=media_objects,json=mediaObjects,proto3" json:"media_objects,omitempty"`
}

func (x *Media) Reset() {
	*x = Media{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_conversation_prompt_content_media_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Media) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Media) ProtoMessage() {}

func (x *Media) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_conversation_prompt_content_media_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Media.ProtoReflect.Descriptor instead.
func (*Media) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_conversation_prompt_content_media_proto_rawDescGZIP(), []int{0}
}

func (x *Media) GetMediaType() Media_MediaType {
	if x != nil {
		return x.MediaType
	}
	return Media_MEDIA_TYPE_UNSPECIFIED
}

func (x *Media) GetStartOffset() *durationpb.Duration {
	if x != nil {
		return x.StartOffset
	}
	return nil
}

func (x *Media) GetOptionalMediaControls() []Media_OptionalMediaControls {
	if x != nil {
		return x.OptionalMediaControls
	}
	return nil
}

func (x *Media) GetMediaObjects() []*MediaObject {
	if x != nil {
		return x.MediaObjects
	}
	return nil
}

// Represents a single media object
type MediaObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of this media object.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Description of this media object.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// The url pointing to the media content.
	Url string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	// Image to show with the media card.
	Image *MediaImage `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *MediaObject) Reset() {
	*x = MediaObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_conversation_prompt_content_media_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaObject) ProtoMessage() {}

func (x *MediaObject) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_conversation_prompt_content_media_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaObject.ProtoReflect.Descriptor instead.
func (*MediaObject) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_conversation_prompt_content_media_proto_rawDescGZIP(), []int{1}
}

func (x *MediaObject) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MediaObject) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MediaObject) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *MediaObject) GetImage() *MediaImage {
	if x != nil {
		return x.Image
	}
	return nil
}

// Image to show with the media card.
type MediaImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Image.
	//
	// Types that are assignable to Image:
	//
	//	*MediaImage_Large
	//	*MediaImage_Icon
	Image isMediaImage_Image `protobuf_oneof:"image"`
}

func (x *MediaImage) Reset() {
	*x = MediaImage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_conversation_prompt_content_media_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaImage) ProtoMessage() {}

func (x *MediaImage) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_conversation_prompt_content_media_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaImage.ProtoReflect.Descriptor instead.
func (*MediaImage) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_conversation_prompt_content_media_proto_rawDescGZIP(), []int{2}
}

func (m *MediaImage) GetImage() isMediaImage_Image {
	if m != nil {
		return m.Image
	}
	return nil
}

func (x *MediaImage) GetLarge() *Image {
	if x, ok := x.GetImage().(*MediaImage_Large); ok {
		return x.Large
	}
	return nil
}

func (x *MediaImage) GetIcon() *Image {
	if x, ok := x.GetImage().(*MediaImage_Icon); ok {
		return x.Icon
	}
	return nil
}

type isMediaImage_Image interface {
	isMediaImage_Image()
}

type MediaImage_Large struct {
	// A large image, such as the cover of the album, etc.
	Large *Image `protobuf:"bytes,1,opt,name=large,proto3,oneof"`
}

type MediaImage_Icon struct {
	// A small image icon displayed on the right from the title.
	// It's resized to 36x36 dp.
	Icon *Image `protobuf:"bytes,2,opt,name=icon,proto3,oneof"`
}

func (*MediaImage_Large) isMediaImage_Image() {}

func (*MediaImage_Icon) isMediaImage_Image() {}

var File_google_actions_sdk_v2_conversation_prompt_content_media_proto protoreflect.FileDescriptor

var file_google_actions_sdk_v2_conversation_prompt_content_media_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x8d, 0x04, 0x0a, 0x05, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x12, 0x52, 0x0a, 0x0a,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x3c, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x77,
	0x0a, 0x17, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73,
	0x52, 0x15, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x12, 0x54, 0x0a, 0x0d, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x0c, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22, 0x48, 0x0a,
	0x09, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x4d, 0x45,
	0x44, 0x49, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x55, 0x44, 0x49, 0x4f, 0x10,
	0x01, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x41, 0x43, 0x4b, 0x10, 0x02, 0x22, 0x59, 0x0a, 0x15, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73,
	0x12, 0x27, 0x0a, 0x23, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x5f, 0x4d, 0x45, 0x44,
	0x49, 0x41, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x41, 0x55,
	0x53, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54, 0x4f, 0x50, 0x50, 0x45, 0x44,
	0x10, 0x02, 0x22, 0x9b, 0x01, 0x0a, 0x0b, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x44, 0x0a, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76,
	0x32, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d,
	0x65, 0x64, 0x69, 0x61, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x22, 0x99, 0x01, 0x0a, 0x0a, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x41, 0x0a, 0x05, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x61, 0x72,
	0x67, 0x65, 0x12, 0x3f, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x04, 0x69,
	0x63, 0x6f, 0x6e, 0x42, 0x07, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x86, 0x01, 0x0a,
	0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_actions_sdk_v2_conversation_prompt_content_media_proto_rawDescOnce sync.Once
	file_google_actions_sdk_v2_conversation_prompt_content_media_proto_rawDescData = file_google_actions_sdk_v2_conversation_prompt_content_media_proto_rawDesc
)

func file_google_actions_sdk_v2_conversation_prompt_content_media_proto_rawDescGZIP() []byte {
	file_google_actions_sdk_v2_conversation_prompt_content_media_proto_rawDescOnce.Do(func() {
		file_google_actions_sdk_v2_conversation_prompt_content_media_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_actions_sdk_v2_conversation_prompt_content_media_proto_rawDescData)
	})
	return file_google_actions_sdk_v2_conversation_prompt_content_media_proto_rawDescData
}

var file_google_actions_sdk_v2_conversation_prompt_content_media_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_actions_sdk_v2_conversation_prompt_content_media_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_actions_sdk_v2_conversation_prompt_content_media_proto_goTypes = []interface{}{
	(Media_MediaType)(0),             // 0: google.actions.sdk.v2.conversation.Media.MediaType
	(Media_OptionalMediaControls)(0), // 1: google.actions.sdk.v2.conversation.Media.OptionalMediaControls
	(*Media)(nil),                    // 2: google.actions.sdk.v2.conversation.Media
	(*MediaObject)(nil),              // 3: google.actions.sdk.v2.conversation.MediaObject
	(*MediaImage)(nil),               // 4: google.actions.sdk.v2.conversation.MediaImage
	(*durationpb.Duration)(nil),      // 5: google.protobuf.Duration
	(*Image)(nil),                    // 6: google.actions.sdk.v2.conversation.Image
}
var file_google_actions_sdk_v2_conversation_prompt_content_media_proto_depIdxs = []int32{
	0, // 0: google.actions.sdk.v2.conversation.Media.media_type:type_name -> google.actions.sdk.v2.conversation.Media.MediaType
	5, // 1: google.actions.sdk.v2.conversation.Media.start_offset:type_name -> google.protobuf.Duration
	1, // 2: google.actions.sdk.v2.conversation.Media.optional_media_controls:type_name -> google.actions.sdk.v2.conversation.Media.OptionalMediaControls
	3, // 3: google.actions.sdk.v2.conversation.Media.media_objects:type_name -> google.actions.sdk.v2.conversation.MediaObject
	4, // 4: google.actions.sdk.v2.conversation.MediaObject.image:type_name -> google.actions.sdk.v2.conversation.MediaImage
	6, // 5: google.actions.sdk.v2.conversation.MediaImage.large:type_name -> google.actions.sdk.v2.conversation.Image
	6, // 6: google.actions.sdk.v2.conversation.MediaImage.icon:type_name -> google.actions.sdk.v2.conversation.Image
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_google_actions_sdk_v2_conversation_prompt_content_media_proto_init() }
func file_google_actions_sdk_v2_conversation_prompt_content_media_proto_init() {
	if File_google_actions_sdk_v2_conversation_prompt_content_media_proto != nil {
		return
	}
	file_google_actions_sdk_v2_conversation_prompt_content_image_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_actions_sdk_v2_conversation_prompt_content_media_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Media); i {
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
		file_google_actions_sdk_v2_conversation_prompt_content_media_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaObject); i {
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
		file_google_actions_sdk_v2_conversation_prompt_content_media_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaImage); i {
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
	file_google_actions_sdk_v2_conversation_prompt_content_media_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*MediaImage_Large)(nil),
		(*MediaImage_Icon)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_actions_sdk_v2_conversation_prompt_content_media_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_actions_sdk_v2_conversation_prompt_content_media_proto_goTypes,
		DependencyIndexes: file_google_actions_sdk_v2_conversation_prompt_content_media_proto_depIdxs,
		EnumInfos:         file_google_actions_sdk_v2_conversation_prompt_content_media_proto_enumTypes,
		MessageInfos:      file_google_actions_sdk_v2_conversation_prompt_content_media_proto_msgTypes,
	}.Build()
	File_google_actions_sdk_v2_conversation_prompt_content_media_proto = out.File
	file_google_actions_sdk_v2_conversation_prompt_content_media_proto_rawDesc = nil
	file_google_actions_sdk_v2_conversation_prompt_content_media_proto_goTypes = nil
	file_google_actions_sdk_v2_conversation_prompt_content_media_proto_depIdxs = nil
}
