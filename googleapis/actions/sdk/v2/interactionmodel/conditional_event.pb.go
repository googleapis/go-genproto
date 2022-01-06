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
// source: google/actions/sdk/v2/interactionmodel/conditional_event.proto

package interactionmodel

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Registers events that trigger as the result of a true condition.
type ConditionalEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Filter condition for this event to trigger. If condition is evaluated to
	// true then the associated `handler` will be triggered.
	// The following variable references are supported:
	//   `$session` - To reference data in session storage.
	//   `$user` - To reference data in user storage.
	// The following boolean operators are supported (with examples):
	//   `&&` - `session.params.counter > 0 && session.params.counter < 100`
	//   `||` - `session.params.foo == "John" || session.params.counter == "Adam"`
	//   `!`  - `!(session.params.counter == 5)`
	// The following comparisons are supported:
	//   `==`, `!=`, `<`, `>`, `<=`, `>=`
	// The following list and string operators are supported (with examples):
	//   `in`        - "Watermelon" in `session.params.fruitList`
	//   `size`      - `size(session.params.fruitList) > 2`
	//   `substring` - `session.params.fullName.contains("John")`
	Condition string `protobuf:"bytes,1,opt,name=condition,proto3" json:"condition,omitempty"`
	// Optional. Destination scene which the conversation should jump to when the associated
	// condition is evaluated to true. The state of the current scene is destroyed
	// on the transition.
	TransitionToScene string `protobuf:"bytes,2,opt,name=transition_to_scene,json=transitionToScene,proto3" json:"transition_to_scene,omitempty"`
	// Optional. Event handler which is triggered when the associated condition is evaluated
	// to `true`. Should execute before transitioning to the destination scene.
	// Useful to generate Prompts in response to events.
	Handler *EventHandler `protobuf:"bytes,3,opt,name=handler,proto3" json:"handler,omitempty"`
}

func (x *ConditionalEvent) Reset() {
	*x = ConditionalEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_interactionmodel_conditional_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConditionalEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConditionalEvent) ProtoMessage() {}

func (x *ConditionalEvent) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_interactionmodel_conditional_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConditionalEvent.ProtoReflect.Descriptor instead.
func (*ConditionalEvent) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_interactionmodel_conditional_event_proto_rawDescGZIP(), []int{0}
}

func (x *ConditionalEvent) GetCondition() string {
	if x != nil {
		return x.Condition
	}
	return ""
}

func (x *ConditionalEvent) GetTransitionToScene() string {
	if x != nil {
		return x.TransitionToScene
	}
	return ""
}

func (x *ConditionalEvent) GetHandler() *EventHandler {
	if x != nil {
		return x.Handler
	}
	return nil
}

var File_google_actions_sdk_v2_interactionmodel_conditional_event_proto protoreflect.FileDescriptor

var file_google_actions_sdk_v2_interactionmodel_conditional_event_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x01, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x09, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a,
	0x13, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x5f, 0x73,
	0x63, 0x65, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52,
	0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x53, 0x63, 0x65,
	0x6e, 0x65, 0x12, 0x53, 0x0a, 0x07, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x07,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x42, 0x9d, 0x01, 0x0a, 0x2a, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73,
	0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x15, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x56, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73,
	0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x3b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_actions_sdk_v2_interactionmodel_conditional_event_proto_rawDescOnce sync.Once
	file_google_actions_sdk_v2_interactionmodel_conditional_event_proto_rawDescData = file_google_actions_sdk_v2_interactionmodel_conditional_event_proto_rawDesc
)

func file_google_actions_sdk_v2_interactionmodel_conditional_event_proto_rawDescGZIP() []byte {
	file_google_actions_sdk_v2_interactionmodel_conditional_event_proto_rawDescOnce.Do(func() {
		file_google_actions_sdk_v2_interactionmodel_conditional_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_actions_sdk_v2_interactionmodel_conditional_event_proto_rawDescData)
	})
	return file_google_actions_sdk_v2_interactionmodel_conditional_event_proto_rawDescData
}

var file_google_actions_sdk_v2_interactionmodel_conditional_event_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_actions_sdk_v2_interactionmodel_conditional_event_proto_goTypes = []interface{}{
	(*ConditionalEvent)(nil), // 0: google.actions.sdk.v2.interactionmodel.ConditionalEvent
	(*EventHandler)(nil),     // 1: google.actions.sdk.v2.interactionmodel.EventHandler
}
var file_google_actions_sdk_v2_interactionmodel_conditional_event_proto_depIdxs = []int32{
	1, // 0: google.actions.sdk.v2.interactionmodel.ConditionalEvent.handler:type_name -> google.actions.sdk.v2.interactionmodel.EventHandler
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_actions_sdk_v2_interactionmodel_conditional_event_proto_init() }
func file_google_actions_sdk_v2_interactionmodel_conditional_event_proto_init() {
	if File_google_actions_sdk_v2_interactionmodel_conditional_event_proto != nil {
		return
	}
	file_google_actions_sdk_v2_interactionmodel_event_handler_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_actions_sdk_v2_interactionmodel_conditional_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConditionalEvent); i {
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
			RawDescriptor: file_google_actions_sdk_v2_interactionmodel_conditional_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_actions_sdk_v2_interactionmodel_conditional_event_proto_goTypes,
		DependencyIndexes: file_google_actions_sdk_v2_interactionmodel_conditional_event_proto_depIdxs,
		MessageInfos:      file_google_actions_sdk_v2_interactionmodel_conditional_event_proto_msgTypes,
	}.Build()
	File_google_actions_sdk_v2_interactionmodel_conditional_event_proto = out.File
	file_google_actions_sdk_v2_interactionmodel_conditional_event_proto_rawDesc = nil
	file_google_actions_sdk_v2_interactionmodel_conditional_event_proto_goTypes = nil
	file_google_actions_sdk_v2_interactionmodel_conditional_event_proto_depIdxs = nil
}
