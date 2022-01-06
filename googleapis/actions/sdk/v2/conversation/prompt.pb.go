// Copyright 2020 Google LLC
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
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: google/actions/sdk/v2/conversation/prompt/prompt.proto

package conversation

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Represent a response to a user.
type Prompt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. Mode for how this messages should be merged with previously
	// defined messages.
	// "false" will clear all previously defined messages (first and last
	// simple, content, suggestions link and canvas) and add messages defined in
	// this prompt.
	// "true" will add messages defined in this prompt to messages defined in
	// previous responses. Setting this field to "true" will also enable appending
	// to some fields inside Simple prompts, the Suggestion prompt and the Canvas
	// prompt (part of the Content prompt). The Content and Link messages will
	// always be overwritten if defined in the prompt.
	// Default value is "false".
	//
	// Deprecated: Do not use.
	Append bool `protobuf:"varint,1,opt,name=append,proto3" json:"append,omitempty"`
	// Optional. Mode for how this messages should be merged with previously
	// defined messages.
	// "true" clears all previously defined messages (first and last
	// simple, content, suggestions link and canvas) and adds messages defined in
	// this prompt.
	// "false" adds messages defined in this prompt to messages defined in
	// previous responses. Leaving this field to "false" also enables
	// appending to some fields inside Simple prompts, the Suggestions prompt,
	// and the Canvas prompt (part of the Content prompt). The Content and Link
	// messages are always overwritten if defined in the prompt. Default
	// value is "false".
	Override bool `protobuf:"varint,8,opt,name=override,proto3" json:"override,omitempty"`
	// Optional. The first voice and text-only response.
	FirstSimple *Simple `protobuf:"bytes,2,opt,name=first_simple,json=firstSimple,proto3" json:"first_simple,omitempty"`
	// Optional. A content like a card, list or media to display to the user.
	Content *Content `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	// Optional. The last voice and text-only response.
	LastSimple *Simple `protobuf:"bytes,4,opt,name=last_simple,json=lastSimple,proto3" json:"last_simple,omitempty"`
	// Optional. Suggestions to be displayed to the user which will always appear
	// at the end of the response.
	// If the "override" field in the containing prompt is "false", the titles
	// defined in this field will be added to titles defined in any previously
	// defined suggestions prompts and duplicate values will be removed.
	Suggestions []*Suggestion `protobuf:"bytes,5,rep,name=suggestions,proto3" json:"suggestions,omitempty"`
	// Optional. An additional suggestion chip that can link out to the associated app
	// or site.
	// The chip will be rendered with the title "Open <name>". Max 20 chars.
	Link *Link `protobuf:"bytes,6,opt,name=link,proto3" json:"link,omitempty"`
	// Optional. Represents a Interactive Canvas response to be sent to the user.
	Canvas *Canvas `protobuf:"bytes,9,opt,name=canvas,proto3" json:"canvas,omitempty"`
}

func (x *Prompt) Reset() {
	*x = Prompt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_conversation_prompt_prompt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Prompt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Prompt) ProtoMessage() {}

func (x *Prompt) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_conversation_prompt_prompt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Prompt.ProtoReflect.Descriptor instead.
func (*Prompt) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_conversation_prompt_prompt_proto_rawDescGZIP(), []int{0}
}

// Deprecated: Do not use.
func (x *Prompt) GetAppend() bool {
	if x != nil {
		return x.Append
	}
	return false
}

func (x *Prompt) GetOverride() bool {
	if x != nil {
		return x.Override
	}
	return false
}

func (x *Prompt) GetFirstSimple() *Simple {
	if x != nil {
		return x.FirstSimple
	}
	return nil
}

func (x *Prompt) GetContent() *Content {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Prompt) GetLastSimple() *Simple {
	if x != nil {
		return x.LastSimple
	}
	return nil
}

func (x *Prompt) GetSuggestions() []*Suggestion {
	if x != nil {
		return x.Suggestions
	}
	return nil
}

func (x *Prompt) GetLink() *Link {
	if x != nil {
		return x.Link
	}
	return nil
}

func (x *Prompt) GetCanvas() *Canvas {
	if x != nil {
		return x.Canvas
	}
	return nil
}

var File_google_actions_sdk_v2_conversation_prompt_prompt_proto protoreflect.FileDescriptor

var file_google_actions_sdk_v2_conversation_prompt_prompt_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x6d,
	0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e,
	0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x3e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b,
	0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f,
	0x63, 0x61, 0x6e, 0x76, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b,
	0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64,
	0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f,
	0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2f, 0x73,
	0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xf7, 0x03, 0x0a, 0x06, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12, 0x1a, 0x0a, 0x06, 0x61, 0x70,
	0x70, 0x65, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x42, 0x02, 0x18, 0x01, 0x52, 0x06,
	0x61, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69,
	0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69,
	0x64, 0x65, 0x12, 0x4d, 0x0a, 0x0c, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x73, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32,
	0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x52, 0x0b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x53, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x12, 0x45, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x4b, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73,
	0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x53,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x50, 0x0a, 0x0b, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e,
	0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x73, 0x75, 0x67, 0x67,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3c, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52,
	0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x42, 0x0a, 0x06, 0x63, 0x61, 0x6e, 0x76, 0x61, 0x73, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x61, 0x6e, 0x76, 0x61,
	0x73, 0x52, 0x06, 0x63, 0x61, 0x6e, 0x76, 0x61, 0x73, 0x42, 0x87, 0x01, 0x0a, 0x26, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0b, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x4e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_actions_sdk_v2_conversation_prompt_prompt_proto_rawDescOnce sync.Once
	file_google_actions_sdk_v2_conversation_prompt_prompt_proto_rawDescData = file_google_actions_sdk_v2_conversation_prompt_prompt_proto_rawDesc
)

func file_google_actions_sdk_v2_conversation_prompt_prompt_proto_rawDescGZIP() []byte {
	file_google_actions_sdk_v2_conversation_prompt_prompt_proto_rawDescOnce.Do(func() {
		file_google_actions_sdk_v2_conversation_prompt_prompt_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_actions_sdk_v2_conversation_prompt_prompt_proto_rawDescData)
	})
	return file_google_actions_sdk_v2_conversation_prompt_prompt_proto_rawDescData
}

var file_google_actions_sdk_v2_conversation_prompt_prompt_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_actions_sdk_v2_conversation_prompt_prompt_proto_goTypes = []interface{}{
	(*Prompt)(nil),     // 0: google.actions.sdk.v2.conversation.Prompt
	(*Simple)(nil),     // 1: google.actions.sdk.v2.conversation.Simple
	(*Content)(nil),    // 2: google.actions.sdk.v2.conversation.Content
	(*Suggestion)(nil), // 3: google.actions.sdk.v2.conversation.Suggestion
	(*Link)(nil),       // 4: google.actions.sdk.v2.conversation.Link
	(*Canvas)(nil),     // 5: google.actions.sdk.v2.conversation.Canvas
}
var file_google_actions_sdk_v2_conversation_prompt_prompt_proto_depIdxs = []int32{
	1, // 0: google.actions.sdk.v2.conversation.Prompt.first_simple:type_name -> google.actions.sdk.v2.conversation.Simple
	2, // 1: google.actions.sdk.v2.conversation.Prompt.content:type_name -> google.actions.sdk.v2.conversation.Content
	1, // 2: google.actions.sdk.v2.conversation.Prompt.last_simple:type_name -> google.actions.sdk.v2.conversation.Simple
	3, // 3: google.actions.sdk.v2.conversation.Prompt.suggestions:type_name -> google.actions.sdk.v2.conversation.Suggestion
	4, // 4: google.actions.sdk.v2.conversation.Prompt.link:type_name -> google.actions.sdk.v2.conversation.Link
	5, // 5: google.actions.sdk.v2.conversation.Prompt.canvas:type_name -> google.actions.sdk.v2.conversation.Canvas
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_actions_sdk_v2_conversation_prompt_prompt_proto_init() }
func file_google_actions_sdk_v2_conversation_prompt_prompt_proto_init() {
	if File_google_actions_sdk_v2_conversation_prompt_prompt_proto != nil {
		return
	}
	file_google_actions_sdk_v2_conversation_prompt_content_canvas_proto_init()
	file_google_actions_sdk_v2_conversation_prompt_content_content_proto_init()
	file_google_actions_sdk_v2_conversation_prompt_content_link_proto_init()
	file_google_actions_sdk_v2_conversation_prompt_simple_proto_init()
	file_google_actions_sdk_v2_conversation_prompt_suggestion_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_actions_sdk_v2_conversation_prompt_prompt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Prompt); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_actions_sdk_v2_conversation_prompt_prompt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_actions_sdk_v2_conversation_prompt_prompt_proto_goTypes,
		DependencyIndexes: file_google_actions_sdk_v2_conversation_prompt_prompt_proto_depIdxs,
		MessageInfos:      file_google_actions_sdk_v2_conversation_prompt_prompt_proto_msgTypes,
	}.Build()
	File_google_actions_sdk_v2_conversation_prompt_prompt_proto = out.File
	file_google_actions_sdk_v2_conversation_prompt_prompt_proto_rawDesc = nil
	file_google_actions_sdk_v2_conversation_prompt_prompt_proto_goTypes = nil
	file_google_actions_sdk_v2_conversation_prompt_prompt_proto_depIdxs = nil
}
