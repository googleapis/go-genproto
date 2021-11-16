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
// 	protoc        v3.13.0
// source: google/ads/googleads/v9/errors/setting_error.proto

package errors

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Enum describing possible setting errors.
type SettingErrorEnum_SettingError int32

const (
	// Enum unspecified.
	SettingErrorEnum_UNSPECIFIED SettingErrorEnum_SettingError = 0
	// The received error code is not known in this version.
	SettingErrorEnum_UNKNOWN SettingErrorEnum_SettingError = 1
	// The campaign setting is not available for this Google Ads account.
	SettingErrorEnum_SETTING_TYPE_IS_NOT_AVAILABLE SettingErrorEnum_SettingError = 3
	// The setting is not compatible with the campaign.
	SettingErrorEnum_SETTING_TYPE_IS_NOT_COMPATIBLE_WITH_CAMPAIGN SettingErrorEnum_SettingError = 4
	// The supplied TargetingSetting contains an invalid CriterionTypeGroup. See
	// CriterionTypeGroup documentation for CriterionTypeGroups allowed
	// in Campaign or AdGroup TargetingSettings.
	SettingErrorEnum_TARGETING_SETTING_CONTAINS_INVALID_CRITERION_TYPE_GROUP SettingErrorEnum_SettingError = 5
	// TargetingSetting must not explicitly
	// set any of the Demographic CriterionTypeGroups (AGE_RANGE, GENDER,
	// PARENT, INCOME_RANGE) to false (it's okay to not set them at all, in
	// which case the system will set them to true automatically).
	SettingErrorEnum_TARGETING_SETTING_DEMOGRAPHIC_CRITERION_TYPE_GROUPS_MUST_BE_SET_TO_TARGET_ALL SettingErrorEnum_SettingError = 6
	// TargetingSetting cannot change any of
	// the Demographic CriterionTypeGroups (AGE_RANGE, GENDER, PARENT,
	// INCOME_RANGE) from true to false.
	SettingErrorEnum_TARGETING_SETTING_CANNOT_CHANGE_TARGET_ALL_TO_FALSE_FOR_DEMOGRAPHIC_CRITERION_TYPE_GROUP SettingErrorEnum_SettingError = 7
	// At least one feed id should be present.
	SettingErrorEnum_DYNAMIC_SEARCH_ADS_SETTING_AT_LEAST_ONE_FEED_ID_MUST_BE_PRESENT SettingErrorEnum_SettingError = 8
	// The supplied DynamicSearchAdsSetting contains an invalid domain name.
	SettingErrorEnum_DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_INVALID_DOMAIN_NAME SettingErrorEnum_SettingError = 9
	// The supplied DynamicSearchAdsSetting contains a subdomain name.
	SettingErrorEnum_DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_SUBDOMAIN_NAME SettingErrorEnum_SettingError = 10
	// The supplied DynamicSearchAdsSetting contains an invalid language code.
	SettingErrorEnum_DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_INVALID_LANGUAGE_CODE SettingErrorEnum_SettingError = 11
	// TargetingSettings in search campaigns should not have
	// CriterionTypeGroup.PLACEMENT set to targetAll.
	SettingErrorEnum_TARGET_ALL_IS_NOT_ALLOWED_FOR_PLACEMENT_IN_SEARCH_CAMPAIGN SettingErrorEnum_SettingError = 12
	// The setting value is not compatible with the campaign type.
	SettingErrorEnum_SETTING_VALUE_NOT_COMPATIBLE_WITH_CAMPAIGN SettingErrorEnum_SettingError = 20
)

// Enum value maps for SettingErrorEnum_SettingError.
var (
	SettingErrorEnum_SettingError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		3:  "SETTING_TYPE_IS_NOT_AVAILABLE",
		4:  "SETTING_TYPE_IS_NOT_COMPATIBLE_WITH_CAMPAIGN",
		5:  "TARGETING_SETTING_CONTAINS_INVALID_CRITERION_TYPE_GROUP",
		6:  "TARGETING_SETTING_DEMOGRAPHIC_CRITERION_TYPE_GROUPS_MUST_BE_SET_TO_TARGET_ALL",
		7:  "TARGETING_SETTING_CANNOT_CHANGE_TARGET_ALL_TO_FALSE_FOR_DEMOGRAPHIC_CRITERION_TYPE_GROUP",
		8:  "DYNAMIC_SEARCH_ADS_SETTING_AT_LEAST_ONE_FEED_ID_MUST_BE_PRESENT",
		9:  "DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_INVALID_DOMAIN_NAME",
		10: "DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_SUBDOMAIN_NAME",
		11: "DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_INVALID_LANGUAGE_CODE",
		12: "TARGET_ALL_IS_NOT_ALLOWED_FOR_PLACEMENT_IN_SEARCH_CAMPAIGN",
		20: "SETTING_VALUE_NOT_COMPATIBLE_WITH_CAMPAIGN",
	}
	SettingErrorEnum_SettingError_value = map[string]int32{
		"UNSPECIFIED":                   0,
		"UNKNOWN":                       1,
		"SETTING_TYPE_IS_NOT_AVAILABLE": 3,
		"SETTING_TYPE_IS_NOT_COMPATIBLE_WITH_CAMPAIGN":                                             4,
		"TARGETING_SETTING_CONTAINS_INVALID_CRITERION_TYPE_GROUP":                                  5,
		"TARGETING_SETTING_DEMOGRAPHIC_CRITERION_TYPE_GROUPS_MUST_BE_SET_TO_TARGET_ALL":            6,
		"TARGETING_SETTING_CANNOT_CHANGE_TARGET_ALL_TO_FALSE_FOR_DEMOGRAPHIC_CRITERION_TYPE_GROUP": 7,
		"DYNAMIC_SEARCH_ADS_SETTING_AT_LEAST_ONE_FEED_ID_MUST_BE_PRESENT":                          8,
		"DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_INVALID_DOMAIN_NAME":                                  9,
		"DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_SUBDOMAIN_NAME":                                       10,
		"DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_INVALID_LANGUAGE_CODE":                                11,
		"TARGET_ALL_IS_NOT_ALLOWED_FOR_PLACEMENT_IN_SEARCH_CAMPAIGN":                               12,
		"SETTING_VALUE_NOT_COMPATIBLE_WITH_CAMPAIGN":                                               20,
	}
)

func (x SettingErrorEnum_SettingError) Enum() *SettingErrorEnum_SettingError {
	p := new(SettingErrorEnum_SettingError)
	*p = x
	return p
}

func (x SettingErrorEnum_SettingError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SettingErrorEnum_SettingError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v9_errors_setting_error_proto_enumTypes[0].Descriptor()
}

func (SettingErrorEnum_SettingError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v9_errors_setting_error_proto_enumTypes[0]
}

func (x SettingErrorEnum_SettingError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SettingErrorEnum_SettingError.Descriptor instead.
func (SettingErrorEnum_SettingError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v9_errors_setting_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible setting errors.
type SettingErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SettingErrorEnum) Reset() {
	*x = SettingErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v9_errors_setting_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SettingErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SettingErrorEnum) ProtoMessage() {}

func (x *SettingErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v9_errors_setting_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SettingErrorEnum.ProtoReflect.Descriptor instead.
func (*SettingErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v9_errors_setting_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v9_errors_setting_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v9_errors_setting_error_proto_rawDesc = []byte{
	0x0a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xed, 0x05, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xd8, 0x05, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x21, 0x0a, 0x1d, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e,
	0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x56,
	0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x03, 0x12, 0x30, 0x0a, 0x2c, 0x53, 0x45, 0x54,
	0x54, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x53, 0x5f, 0x4e, 0x4f, 0x54,
	0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x54, 0x49, 0x42, 0x4c, 0x45, 0x5f, 0x57, 0x49, 0x54, 0x48,
	0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x10, 0x04, 0x12, 0x3b, 0x0a, 0x37, 0x54,
	0x41, 0x52, 0x47, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47,
	0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x5f, 0x43, 0x52, 0x49, 0x54, 0x45, 0x52, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x10, 0x05, 0x12, 0x51, 0x0a, 0x4d, 0x54, 0x41, 0x52, 0x47,
	0x45, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x44, 0x45,
	0x4d, 0x4f, 0x47, 0x52, 0x41, 0x50, 0x48, 0x49, 0x43, 0x5f, 0x43, 0x52, 0x49, 0x54, 0x45, 0x52,
	0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x53, 0x5f,
	0x4d, 0x55, 0x53, 0x54, 0x5f, 0x42, 0x45, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x54, 0x4f, 0x5f, 0x54,
	0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x06, 0x12, 0x5c, 0x0a, 0x58, 0x54,
	0x41, 0x52, 0x47, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47,
	0x5f, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x54,
	0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x5f, 0x54, 0x4f, 0x5f, 0x46, 0x41, 0x4c,
	0x53, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x44, 0x45, 0x4d, 0x4f, 0x47, 0x52, 0x41, 0x50, 0x48,
	0x49, 0x43, 0x5f, 0x43, 0x52, 0x49, 0x54, 0x45, 0x52, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x10, 0x07, 0x12, 0x43, 0x0a, 0x3f, 0x44, 0x59, 0x4e,
	0x41, 0x4d, 0x49, 0x43, 0x5f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x41, 0x44, 0x53, 0x5f,
	0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x54, 0x5f, 0x4c, 0x45, 0x41, 0x53, 0x54,
	0x5f, 0x4f, 0x4e, 0x45, 0x5f, 0x46, 0x45, 0x45, 0x44, 0x5f, 0x49, 0x44, 0x5f, 0x4d, 0x55, 0x53,
	0x54, 0x5f, 0x42, 0x45, 0x5f, 0x50, 0x52, 0x45, 0x53, 0x45, 0x4e, 0x54, 0x10, 0x08, 0x12, 0x3b,
	0x0a, 0x37, 0x44, 0x59, 0x4e, 0x41, 0x4d, 0x49, 0x43, 0x5f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48,
	0x5f, 0x41, 0x44, 0x53, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x4f, 0x4e,
	0x54, 0x41, 0x49, 0x4e, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x44, 0x4f,
	0x4d, 0x41, 0x49, 0x4e, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x09, 0x12, 0x36, 0x0a, 0x32, 0x44,
	0x59, 0x4e, 0x41, 0x4d, 0x49, 0x43, 0x5f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x41, 0x44,
	0x53, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49,
	0x4e, 0x53, 0x5f, 0x53, 0x55, 0x42, 0x44, 0x4f, 0x4d, 0x41, 0x49, 0x4e, 0x5f, 0x4e, 0x41, 0x4d,
	0x45, 0x10, 0x0a, 0x12, 0x3d, 0x0a, 0x39, 0x44, 0x59, 0x4e, 0x41, 0x4d, 0x49, 0x43, 0x5f, 0x53,
	0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x41, 0x44, 0x53, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e,
	0x47, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x5f, 0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41, 0x47, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45,
	0x10, 0x0b, 0x12, 0x3e, 0x0a, 0x3a, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x41, 0x4c, 0x4c,
	0x5f, 0x49, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x5f,
	0x46, 0x4f, 0x52, 0x5f, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x49, 0x4e,
	0x5f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e,
	0x10, 0x0c, 0x12, 0x2e, 0x0a, 0x2a, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x56, 0x41,
	0x4c, 0x55, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x54, 0x49, 0x42,
	0x4c, 0x45, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e,
	0x10, 0x14, 0x42, 0xec, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x39, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x11, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x2e, 0x56, 0x39, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x5c, 0x56, 0x39, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x22, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x39, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v9_errors_setting_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v9_errors_setting_error_proto_rawDescData = file_google_ads_googleads_v9_errors_setting_error_proto_rawDesc
)

func file_google_ads_googleads_v9_errors_setting_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v9_errors_setting_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v9_errors_setting_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v9_errors_setting_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v9_errors_setting_error_proto_rawDescData
}

var file_google_ads_googleads_v9_errors_setting_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v9_errors_setting_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v9_errors_setting_error_proto_goTypes = []interface{}{
	(SettingErrorEnum_SettingError)(0), // 0: google.ads.googleads.v9.errors.SettingErrorEnum.SettingError
	(*SettingErrorEnum)(nil),           // 1: google.ads.googleads.v9.errors.SettingErrorEnum
}
var file_google_ads_googleads_v9_errors_setting_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v9_errors_setting_error_proto_init() }
func file_google_ads_googleads_v9_errors_setting_error_proto_init() {
	if File_google_ads_googleads_v9_errors_setting_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v9_errors_setting_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SettingErrorEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v9_errors_setting_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v9_errors_setting_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v9_errors_setting_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v9_errors_setting_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v9_errors_setting_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v9_errors_setting_error_proto = out.File
	file_google_ads_googleads_v9_errors_setting_error_proto_rawDesc = nil
	file_google_ads_googleads_v9_errors_setting_error_proto_goTypes = nil
	file_google_ads_googleads_v9_errors_setting_error_proto_depIdxs = nil
}
