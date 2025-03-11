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
// source: google/apps/script/type/extension_point.proto

package _type

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Common format for declaring a  menu item, or button, that appears within a
// host app.
type MenuItemExtensionPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The endpoint to execute when this extension point is
	// activated.
	RunFunction string `protobuf:"bytes,1,opt,name=run_function,json=runFunction,proto3" json:"run_function,omitempty"`
	// Required. User-visible text describing the action taken by activating this
	// extension point. For example, "Insert invoice".
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	// The URL for the logo image shown in the add-on toolbar.
	//
	// If not set, defaults to the add-on's primary logo URL.
	LogoUrl string `protobuf:"bytes,3,opt,name=logo_url,json=logoUrl,proto3" json:"logo_url,omitempty"`
}

func (x *MenuItemExtensionPoint) Reset() {
	*x = MenuItemExtensionPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_apps_script_type_extension_point_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuItemExtensionPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuItemExtensionPoint) ProtoMessage() {}

func (x *MenuItemExtensionPoint) ProtoReflect() protoreflect.Message {
	mi := &file_google_apps_script_type_extension_point_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuItemExtensionPoint.ProtoReflect.Descriptor instead.
func (*MenuItemExtensionPoint) Descriptor() ([]byte, []int) {
	return file_google_apps_script_type_extension_point_proto_rawDescGZIP(), []int{0}
}

func (x *MenuItemExtensionPoint) GetRunFunction() string {
	if x != nil {
		return x.RunFunction
	}
	return ""
}

func (x *MenuItemExtensionPoint) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *MenuItemExtensionPoint) GetLogoUrl() string {
	if x != nil {
		return x.LogoUrl
	}
	return ""
}

// Common format for declaring an add-on's home-page view.
type HomepageExtensionPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The endpoint to execute when this extension point is
	// activated.
	RunFunction string `protobuf:"bytes,1,opt,name=run_function,json=runFunction,proto3" json:"run_function,omitempty"`
	// Optional. If set to `false`, disable the home-page view in this context.
	//
	// Defaults to `true` if unset.
	//
	// If an add-ons custom home-page view is disabled, an autogenerated overview
	// card will be provided for users instead.
	Enabled *wrapperspb.BoolValue `protobuf:"bytes,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *HomepageExtensionPoint) Reset() {
	*x = HomepageExtensionPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_apps_script_type_extension_point_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomepageExtensionPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomepageExtensionPoint) ProtoMessage() {}

func (x *HomepageExtensionPoint) ProtoReflect() protoreflect.Message {
	mi := &file_google_apps_script_type_extension_point_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomepageExtensionPoint.ProtoReflect.Descriptor instead.
func (*HomepageExtensionPoint) Descriptor() ([]byte, []int) {
	return file_google_apps_script_type_extension_point_proto_rawDescGZIP(), []int{1}
}

func (x *HomepageExtensionPoint) GetRunFunction() string {
	if x != nil {
		return x.RunFunction
	}
	return ""
}

func (x *HomepageExtensionPoint) GetEnabled() *wrapperspb.BoolValue {
	if x != nil {
		return x.Enabled
	}
	return nil
}

// Format for declaring a universal action menu item extension point.
type UniversalActionExtensionPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. User-visible text describing the action taken by activating this
	// extension point, for example, "Add a new contact".
	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	// Required. The action type supported on a universal action menu item. It
	// could be either a link to open or an endpoint to execute.
	//
	// Types that are assignable to ActionType:
	//
	//	*UniversalActionExtensionPoint_OpenLink
	//	*UniversalActionExtensionPoint_RunFunction
	ActionType isUniversalActionExtensionPoint_ActionType `protobuf_oneof:"action_type"`
}

func (x *UniversalActionExtensionPoint) Reset() {
	*x = UniversalActionExtensionPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_apps_script_type_extension_point_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UniversalActionExtensionPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UniversalActionExtensionPoint) ProtoMessage() {}

func (x *UniversalActionExtensionPoint) ProtoReflect() protoreflect.Message {
	mi := &file_google_apps_script_type_extension_point_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UniversalActionExtensionPoint.ProtoReflect.Descriptor instead.
func (*UniversalActionExtensionPoint) Descriptor() ([]byte, []int) {
	return file_google_apps_script_type_extension_point_proto_rawDescGZIP(), []int{2}
}

func (x *UniversalActionExtensionPoint) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (m *UniversalActionExtensionPoint) GetActionType() isUniversalActionExtensionPoint_ActionType {
	if m != nil {
		return m.ActionType
	}
	return nil
}

func (x *UniversalActionExtensionPoint) GetOpenLink() string {
	if x, ok := x.GetActionType().(*UniversalActionExtensionPoint_OpenLink); ok {
		return x.OpenLink
	}
	return ""
}

func (x *UniversalActionExtensionPoint) GetRunFunction() string {
	if x, ok := x.GetActionType().(*UniversalActionExtensionPoint_RunFunction); ok {
		return x.RunFunction
	}
	return ""
}

type isUniversalActionExtensionPoint_ActionType interface {
	isUniversalActionExtensionPoint_ActionType()
}

type UniversalActionExtensionPoint_OpenLink struct {
	// URL to be opened by the UniversalAction.
	OpenLink string `protobuf:"bytes,2,opt,name=open_link,json=openLink,proto3,oneof"`
}

type UniversalActionExtensionPoint_RunFunction struct {
	// Endpoint to be run by the UniversalAction.
	RunFunction string `protobuf:"bytes,3,opt,name=run_function,json=runFunction,proto3,oneof"`
}

func (*UniversalActionExtensionPoint_OpenLink) isUniversalActionExtensionPoint_ActionType() {}

func (*UniversalActionExtensionPoint_RunFunction) isUniversalActionExtensionPoint_ActionType() {}

var File_google_apps_script_type_extension_point_proto protoreflect.FileDescriptor

var file_google_apps_script_type_extension_point_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x16, 0x4d, 0x65, 0x6e, 0x75,
	0x49, 0x74, 0x65, 0x6d, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x75, 0x6e, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x75, 0x6e, 0x46, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x6c,
	0x6f, 0x67, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c,
	0x6f, 0x67, 0x6f, 0x55, 0x72, 0x6c, 0x22, 0x71, 0x0a, 0x16, 0x48, 0x6f, 0x6d, 0x65, 0x70, 0x61,
	0x67, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x72, 0x75, 0x6e, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x75, 0x6e, 0x46, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x88, 0x01, 0x0a, 0x1d, 0x55, 0x6e,
	0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x12, 0x1d, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x6e, 0x4c, 0x69, 0x6e, 0x6b,
	0x12, 0x23, 0x0a, 0x0c, 0x72, 0x75, 0x6e, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x72, 0x75, 0x6e, 0x46, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0d, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x42, 0xa8, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70,
	0x70, 0x73, 0x2f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0xaa, 0x02,
	0x17, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x73, 0x2e, 0x53, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0xca, 0x02, 0x17, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x41, 0x70, 0x70, 0x73, 0x5c, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5c, 0x54, 0x79,
	0x70, 0x65, 0xea, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x70, 0x70,
	0x73, 0x3a, 0x3a, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x3a, 0x3a, 0x54, 0x79, 0x70, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_apps_script_type_extension_point_proto_rawDescOnce sync.Once
	file_google_apps_script_type_extension_point_proto_rawDescData = file_google_apps_script_type_extension_point_proto_rawDesc
)

func file_google_apps_script_type_extension_point_proto_rawDescGZIP() []byte {
	file_google_apps_script_type_extension_point_proto_rawDescOnce.Do(func() {
		file_google_apps_script_type_extension_point_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_apps_script_type_extension_point_proto_rawDescData)
	})
	return file_google_apps_script_type_extension_point_proto_rawDescData
}

var file_google_apps_script_type_extension_point_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_apps_script_type_extension_point_proto_goTypes = []interface{}{
	(*MenuItemExtensionPoint)(nil),        // 0: google.apps.script.type.MenuItemExtensionPoint
	(*HomepageExtensionPoint)(nil),        // 1: google.apps.script.type.HomepageExtensionPoint
	(*UniversalActionExtensionPoint)(nil), // 2: google.apps.script.type.UniversalActionExtensionPoint
	(*wrapperspb.BoolValue)(nil),          // 3: google.protobuf.BoolValue
}
var file_google_apps_script_type_extension_point_proto_depIdxs = []int32{
	3, // 0: google.apps.script.type.HomepageExtensionPoint.enabled:type_name -> google.protobuf.BoolValue
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_apps_script_type_extension_point_proto_init() }
func file_google_apps_script_type_extension_point_proto_init() {
	if File_google_apps_script_type_extension_point_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_apps_script_type_extension_point_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuItemExtensionPoint); i {
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
		file_google_apps_script_type_extension_point_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomepageExtensionPoint); i {
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
		file_google_apps_script_type_extension_point_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UniversalActionExtensionPoint); i {
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
	file_google_apps_script_type_extension_point_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*UniversalActionExtensionPoint_OpenLink)(nil),
		(*UniversalActionExtensionPoint_RunFunction)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_apps_script_type_extension_point_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_apps_script_type_extension_point_proto_goTypes,
		DependencyIndexes: file_google_apps_script_type_extension_point_proto_depIdxs,
		MessageInfos:      file_google_apps_script_type_extension_point_proto_msgTypes,
	}.Build()
	File_google_apps_script_type_extension_point_proto = out.File
	file_google_apps_script_type_extension_point_proto_rawDesc = nil
	file_google_apps_script_type_extension_point_proto_goTypes = nil
	file_google_apps_script_type_extension_point_proto_depIdxs = nil
}
