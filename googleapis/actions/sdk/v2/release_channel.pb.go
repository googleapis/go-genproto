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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.24.4
// source: google/actions/sdk/v2/release_channel.proto

package sdk

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

// Definition of release channel resource.
type ReleaseChannel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique name of the release channel in the following format.
	// `projects/{project}/releaseChannels/{release_channel}`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Version currently deployed to this release channel in the following format:
	// `projects/{project}/versions/{version}`.
	CurrentVersion string `protobuf:"bytes,2,opt,name=current_version,json=currentVersion,proto3" json:"current_version,omitempty"`
	// Version to be deployed to this release channel in the following format:
	// `projects/{project}/versions/{version}`.
	PendingVersion string `protobuf:"bytes,3,opt,name=pending_version,json=pendingVersion,proto3" json:"pending_version,omitempty"`
}

func (x *ReleaseChannel) Reset() {
	*x = ReleaseChannel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_release_channel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseChannel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseChannel) ProtoMessage() {}

func (x *ReleaseChannel) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_release_channel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseChannel.ProtoReflect.Descriptor instead.
func (*ReleaseChannel) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_release_channel_proto_rawDescGZIP(), []int{0}
}

func (x *ReleaseChannel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReleaseChannel) GetCurrentVersion() string {
	if x != nil {
		return x.CurrentVersion
	}
	return ""
}

func (x *ReleaseChannel) GetPendingVersion() string {
	if x != nil {
		return x.PendingVersion
	}
	return ""
}

var File_google_actions_sdk_v2_release_channel_proto protoreflect.FileDescriptor

var file_google_actions_sdk_v2_release_channel_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64,
	0x6b, 0x2e, 0x76, 0x32, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xd8, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x27, 0x0a, 0x0f, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x60, 0xea, 0x41, 0x5d, 0x0a, 0x25, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x34, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x2f, 0x7b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x7d, 0x42, 0x6c, 0x0a, 0x19, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x42, 0x13, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64,
	0x6b, 0x2f, 0x76, 0x32, 0x3b, 0x73, 0x64, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_actions_sdk_v2_release_channel_proto_rawDescOnce sync.Once
	file_google_actions_sdk_v2_release_channel_proto_rawDescData = file_google_actions_sdk_v2_release_channel_proto_rawDesc
)

func file_google_actions_sdk_v2_release_channel_proto_rawDescGZIP() []byte {
	file_google_actions_sdk_v2_release_channel_proto_rawDescOnce.Do(func() {
		file_google_actions_sdk_v2_release_channel_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_actions_sdk_v2_release_channel_proto_rawDescData)
	})
	return file_google_actions_sdk_v2_release_channel_proto_rawDescData
}

var file_google_actions_sdk_v2_release_channel_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_actions_sdk_v2_release_channel_proto_goTypes = []interface{}{
	(*ReleaseChannel)(nil), // 0: google.actions.sdk.v2.ReleaseChannel
}
var file_google_actions_sdk_v2_release_channel_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_actions_sdk_v2_release_channel_proto_init() }
func file_google_actions_sdk_v2_release_channel_proto_init() {
	if File_google_actions_sdk_v2_release_channel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_actions_sdk_v2_release_channel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleaseChannel); i {
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
			RawDescriptor: file_google_actions_sdk_v2_release_channel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_actions_sdk_v2_release_channel_proto_goTypes,
		DependencyIndexes: file_google_actions_sdk_v2_release_channel_proto_depIdxs,
		MessageInfos:      file_google_actions_sdk_v2_release_channel_proto_msgTypes,
	}.Build()
	File_google_actions_sdk_v2_release_channel_proto = out.File
	file_google_actions_sdk_v2_release_channel_proto_rawDesc = nil
	file_google_actions_sdk_v2_release_channel_proto_goTypes = nil
	file_google_actions_sdk_v2_release_channel_proto_depIdxs = nil
}
