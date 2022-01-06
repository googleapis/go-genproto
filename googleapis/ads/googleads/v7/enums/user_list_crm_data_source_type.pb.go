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
// source: google/ads/googleads/v7/enums/user_list_crm_data_source_type.proto

package enums

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

// Enum describing possible user list crm data source type.
type UserListCrmDataSourceTypeEnum_UserListCrmDataSourceType int32

const (
	// Not specified.
	UserListCrmDataSourceTypeEnum_UNSPECIFIED UserListCrmDataSourceTypeEnum_UserListCrmDataSourceType = 0
	// Used for return value only. Represents value unknown in this version.
	UserListCrmDataSourceTypeEnum_UNKNOWN UserListCrmDataSourceTypeEnum_UserListCrmDataSourceType = 1
	// The uploaded data is first-party data.
	UserListCrmDataSourceTypeEnum_FIRST_PARTY UserListCrmDataSourceTypeEnum_UserListCrmDataSourceType = 2
	// The uploaded data is from a third-party credit bureau.
	UserListCrmDataSourceTypeEnum_THIRD_PARTY_CREDIT_BUREAU UserListCrmDataSourceTypeEnum_UserListCrmDataSourceType = 3
	// The uploaded data is from a third-party voter file.
	UserListCrmDataSourceTypeEnum_THIRD_PARTY_VOTER_FILE UserListCrmDataSourceTypeEnum_UserListCrmDataSourceType = 4
)

// Enum value maps for UserListCrmDataSourceTypeEnum_UserListCrmDataSourceType.
var (
	UserListCrmDataSourceTypeEnum_UserListCrmDataSourceType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "FIRST_PARTY",
		3: "THIRD_PARTY_CREDIT_BUREAU",
		4: "THIRD_PARTY_VOTER_FILE",
	}
	UserListCrmDataSourceTypeEnum_UserListCrmDataSourceType_value = map[string]int32{
		"UNSPECIFIED":               0,
		"UNKNOWN":                   1,
		"FIRST_PARTY":               2,
		"THIRD_PARTY_CREDIT_BUREAU": 3,
		"THIRD_PARTY_VOTER_FILE":    4,
	}
)

func (x UserListCrmDataSourceTypeEnum_UserListCrmDataSourceType) Enum() *UserListCrmDataSourceTypeEnum_UserListCrmDataSourceType {
	p := new(UserListCrmDataSourceTypeEnum_UserListCrmDataSourceType)
	*p = x
	return p
}

func (x UserListCrmDataSourceTypeEnum_UserListCrmDataSourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserListCrmDataSourceTypeEnum_UserListCrmDataSourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v7_enums_user_list_crm_data_source_type_proto_enumTypes[0].Descriptor()
}

func (UserListCrmDataSourceTypeEnum_UserListCrmDataSourceType) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v7_enums_user_list_crm_data_source_type_proto_enumTypes[0]
}

func (x UserListCrmDataSourceTypeEnum_UserListCrmDataSourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserListCrmDataSourceTypeEnum_UserListCrmDataSourceType.Descriptor instead.
func (UserListCrmDataSourceTypeEnum_UserListCrmDataSourceType) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v7_enums_user_list_crm_data_source_type_proto_rawDescGZIP(), []int{0, 0}
}

// Indicates source of Crm upload data.
type UserListCrmDataSourceTypeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserListCrmDataSourceTypeEnum) Reset() {
	*x = UserListCrmDataSourceTypeEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v7_enums_user_list_crm_data_source_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserListCrmDataSourceTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListCrmDataSourceTypeEnum) ProtoMessage() {}

func (x *UserListCrmDataSourceTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v7_enums_user_list_crm_data_source_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListCrmDataSourceTypeEnum.ProtoReflect.Descriptor instead.
func (*UserListCrmDataSourceTypeEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v7_enums_user_list_crm_data_source_type_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v7_enums_user_list_crm_data_source_type_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v7_enums_user_list_crm_data_source_type_proto_rawDesc = []byte{
	0x0a, 0x42, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x37, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x63, 0x72, 0x6d, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x37, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xa7, 0x01, 0x0a, 0x1d, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x72,
	0x6d, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x45,
	0x6e, 0x75, 0x6d, 0x22, 0x85, 0x01, 0x0a, 0x19, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x72, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12,
	0x0f, 0x0a, 0x0b, 0x46, 0x49, 0x52, 0x53, 0x54, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x59, 0x10, 0x02,
	0x12, 0x1d, 0x0a, 0x19, 0x54, 0x48, 0x49, 0x52, 0x44, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x59, 0x5f,
	0x43, 0x52, 0x45, 0x44, 0x49, 0x54, 0x5f, 0x42, 0x55, 0x52, 0x45, 0x41, 0x55, 0x10, 0x03, 0x12,
	0x1a, 0x0a, 0x16, 0x54, 0x48, 0x49, 0x52, 0x44, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x59, 0x5f, 0x56,
	0x4f, 0x54, 0x45, 0x52, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x04, 0x42, 0xf3, 0x01, 0x0a, 0x21,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x37, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x42, 0x1e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x72, 0x6d, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x37, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1d,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x37, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1d,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x37, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x21,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x37, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v7_enums_user_list_crm_data_source_type_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v7_enums_user_list_crm_data_source_type_proto_rawDescData = file_google_ads_googleads_v7_enums_user_list_crm_data_source_type_proto_rawDesc
)

func file_google_ads_googleads_v7_enums_user_list_crm_data_source_type_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v7_enums_user_list_crm_data_source_type_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v7_enums_user_list_crm_data_source_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v7_enums_user_list_crm_data_source_type_proto_rawDescData)
	})
	return file_google_ads_googleads_v7_enums_user_list_crm_data_source_type_proto_rawDescData
}

var file_google_ads_googleads_v7_enums_user_list_crm_data_source_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v7_enums_user_list_crm_data_source_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v7_enums_user_list_crm_data_source_type_proto_goTypes = []interface{}{
	(UserListCrmDataSourceTypeEnum_UserListCrmDataSourceType)(0), // 0: google.ads.googleads.v7.enums.UserListCrmDataSourceTypeEnum.UserListCrmDataSourceType
	(*UserListCrmDataSourceTypeEnum)(nil),                        // 1: google.ads.googleads.v7.enums.UserListCrmDataSourceTypeEnum
}
var file_google_ads_googleads_v7_enums_user_list_crm_data_source_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v7_enums_user_list_crm_data_source_type_proto_init() }
func file_google_ads_googleads_v7_enums_user_list_crm_data_source_type_proto_init() {
	if File_google_ads_googleads_v7_enums_user_list_crm_data_source_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v7_enums_user_list_crm_data_source_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserListCrmDataSourceTypeEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v7_enums_user_list_crm_data_source_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v7_enums_user_list_crm_data_source_type_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v7_enums_user_list_crm_data_source_type_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v7_enums_user_list_crm_data_source_type_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v7_enums_user_list_crm_data_source_type_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v7_enums_user_list_crm_data_source_type_proto = out.File
	file_google_ads_googleads_v7_enums_user_list_crm_data_source_type_proto_rawDesc = nil
	file_google_ads_googleads_v7_enums_user_list_crm_data_source_type_proto_goTypes = nil
	file_google_ads_googleads_v7_enums_user_list_crm_data_source_type_proto_depIdxs = nil
}
