// Copyright 2021 Google LLC
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
// source: google/ads/googleads/v7/common/tag_snippet.proto

package common

import (
	reflect "reflect"
	sync "sync"

	enums "google.golang.org/genproto/googleapis/ads/googleads/v7/enums"
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

// The site tag and event snippet pair for a TrackingCodeType.
type TagSnippet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of the generated tag snippets for tracking conversions.
	Type enums.TrackingCodeTypeEnum_TrackingCodeType `protobuf:"varint,1,opt,name=type,proto3,enum=google.ads.googleads.v7.enums.TrackingCodeTypeEnum_TrackingCodeType" json:"type,omitempty"`
	// The format of the web page where the tracking tag and snippet will be
	// installed, e.g. HTML.
	PageFormat enums.TrackingCodePageFormatEnum_TrackingCodePageFormat `protobuf:"varint,2,opt,name=page_format,json=pageFormat,proto3,enum=google.ads.googleads.v7.enums.TrackingCodePageFormatEnum_TrackingCodePageFormat" json:"page_format,omitempty"`
	// The site tag that adds visitors to your basic remarketing lists and sets
	// new cookies on your domain.
	GlobalSiteTag *string `protobuf:"bytes,5,opt,name=global_site_tag,json=globalSiteTag,proto3,oneof" json:"global_site_tag,omitempty"`
	// The event snippet that works with the site tag to track actions that
	// should be counted as conversions.
	EventSnippet *string `protobuf:"bytes,6,opt,name=event_snippet,json=eventSnippet,proto3,oneof" json:"event_snippet,omitempty"`
}

func (x *TagSnippet) Reset() {
	*x = TagSnippet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v7_common_tag_snippet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagSnippet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagSnippet) ProtoMessage() {}

func (x *TagSnippet) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v7_common_tag_snippet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagSnippet.ProtoReflect.Descriptor instead.
func (*TagSnippet) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v7_common_tag_snippet_proto_rawDescGZIP(), []int{0}
}

func (x *TagSnippet) GetType() enums.TrackingCodeTypeEnum_TrackingCodeType {
	if x != nil {
		return x.Type
	}
	return enums.TrackingCodeTypeEnum_TrackingCodeType(0)
}

func (x *TagSnippet) GetPageFormat() enums.TrackingCodePageFormatEnum_TrackingCodePageFormat {
	if x != nil {
		return x.PageFormat
	}
	return enums.TrackingCodePageFormatEnum_TrackingCodePageFormat(0)
}

func (x *TagSnippet) GetGlobalSiteTag() string {
	if x != nil && x.GlobalSiteTag != nil {
		return *x.GlobalSiteTag
	}
	return ""
}

func (x *TagSnippet) GetEventSnippet() string {
	if x != nil && x.EventSnippet != nil {
		return *x.EventSnippet
	}
	return ""
}

var File_google_ads_googleads_v7_common_tag_snippet_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v7_common_tag_snippet_proto_rawDesc = []byte{
	0x0a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x37, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x74, 0x61, 0x67, 0x5f, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x1a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x37, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x37, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6, 0x02, 0x0a, 0x0a, 0x54, 0x61, 0x67, 0x53,
	0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x12, 0x58, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x44, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x37, 0x2e, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x64,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69,
	0x6e, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x71, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x50, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x37, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x43, 0x6f,
	0x64, 0x65, 0x50, 0x61, 0x67, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x45, 0x6e, 0x75, 0x6d,
	0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x61, 0x67,
	0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x46, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x12, 0x2b, 0x0a, 0x0f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x73, 0x69,
	0x74, 0x65, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0d,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x53, 0x69, 0x74, 0x65, 0x54, 0x61, 0x67, 0x88, 0x01, 0x01,
	0x12, 0x28, 0x0a, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x88, 0x01, 0x01, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x67,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x61, 0x67, 0x42, 0x10,
	0x0a, 0x0e, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74,
	0x42, 0xea, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x37,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x0f, 0x54, 0x61, 0x67, 0x53, 0x6e, 0x69, 0x70,
	0x70, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f,
	0x76, 0x37, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x37,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56,
	0x37, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x3a, 0x3a, 0x56, 0x37, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v7_common_tag_snippet_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v7_common_tag_snippet_proto_rawDescData = file_google_ads_googleads_v7_common_tag_snippet_proto_rawDesc
)

func file_google_ads_googleads_v7_common_tag_snippet_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v7_common_tag_snippet_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v7_common_tag_snippet_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v7_common_tag_snippet_proto_rawDescData)
	})
	return file_google_ads_googleads_v7_common_tag_snippet_proto_rawDescData
}

var file_google_ads_googleads_v7_common_tag_snippet_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v7_common_tag_snippet_proto_goTypes = []interface{}{
	(*TagSnippet)(nil), // 0: google.ads.googleads.v7.common.TagSnippet
	(enums.TrackingCodeTypeEnum_TrackingCodeType)(0),             // 1: google.ads.googleads.v7.enums.TrackingCodeTypeEnum.TrackingCodeType
	(enums.TrackingCodePageFormatEnum_TrackingCodePageFormat)(0), // 2: google.ads.googleads.v7.enums.TrackingCodePageFormatEnum.TrackingCodePageFormat
}
var file_google_ads_googleads_v7_common_tag_snippet_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v7.common.TagSnippet.type:type_name -> google.ads.googleads.v7.enums.TrackingCodeTypeEnum.TrackingCodeType
	2, // 1: google.ads.googleads.v7.common.TagSnippet.page_format:type_name -> google.ads.googleads.v7.enums.TrackingCodePageFormatEnum.TrackingCodePageFormat
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v7_common_tag_snippet_proto_init() }
func file_google_ads_googleads_v7_common_tag_snippet_proto_init() {
	if File_google_ads_googleads_v7_common_tag_snippet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v7_common_tag_snippet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagSnippet); i {
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
	file_google_ads_googleads_v7_common_tag_snippet_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v7_common_tag_snippet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v7_common_tag_snippet_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v7_common_tag_snippet_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v7_common_tag_snippet_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v7_common_tag_snippet_proto = out.File
	file_google_ads_googleads_v7_common_tag_snippet_proto_rawDesc = nil
	file_google_ads_googleads_v7_common_tag_snippet_proto_goTypes = nil
	file_google_ads_googleads_v7_common_tag_snippet_proto_depIdxs = nil
}
