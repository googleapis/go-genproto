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
// source: google/ads/googleads/v9/enums/payment_mode.proto

package enums

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

// Enum describing possible payment modes.
type PaymentModeEnum_PaymentMode int32

const (
	// Not specified.
	PaymentModeEnum_UNSPECIFIED PaymentModeEnum_PaymentMode = 0
	// Used for return value only. Represents value unknown in this version.
	PaymentModeEnum_UNKNOWN PaymentModeEnum_PaymentMode = 1
	// Pay per click.
	PaymentModeEnum_CLICKS PaymentModeEnum_PaymentMode = 4
	// Pay per conversion value. This mode is only supported by campaigns with
	// AdvertisingChannelType.HOTEL, BiddingStrategyType.COMMISSION, and
	// BudgetType.HOTEL_ADS_COMMISSION.
	PaymentModeEnum_CONVERSION_VALUE PaymentModeEnum_PaymentMode = 5
	// Pay per conversion. This mode is only supported by campaigns with
	// AdvertisingChannelType.DISPLAY (excluding
	// AdvertisingChannelSubType.DISPLAY_GMAIL), BiddingStrategyType.TARGET_CPA,
	// and BudgetType.FIXED_CPA. The customer must also be eligible for this
	// mode. See Customer.eligibility_failure_reasons for details.
	PaymentModeEnum_CONVERSIONS PaymentModeEnum_PaymentMode = 6
	// Pay per guest stay value. This mode is only supported by campaigns with
	// AdvertisingChannelType.HOTEL, BiddingStrategyType.COMMISSION, and
	// BudgetType.STANDARD.
	PaymentModeEnum_GUEST_STAY PaymentModeEnum_PaymentMode = 7
)

// Enum value maps for PaymentModeEnum_PaymentMode.
var (
	PaymentModeEnum_PaymentMode_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		4: "CLICKS",
		5: "CONVERSION_VALUE",
		6: "CONVERSIONS",
		7: "GUEST_STAY",
	}
	PaymentModeEnum_PaymentMode_value = map[string]int32{
		"UNSPECIFIED":      0,
		"UNKNOWN":          1,
		"CLICKS":           4,
		"CONVERSION_VALUE": 5,
		"CONVERSIONS":      6,
		"GUEST_STAY":       7,
	}
)

func (x PaymentModeEnum_PaymentMode) Enum() *PaymentModeEnum_PaymentMode {
	p := new(PaymentModeEnum_PaymentMode)
	*p = x
	return p
}

func (x PaymentModeEnum_PaymentMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PaymentModeEnum_PaymentMode) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v9_enums_payment_mode_proto_enumTypes[0].Descriptor()
}

func (PaymentModeEnum_PaymentMode) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v9_enums_payment_mode_proto_enumTypes[0]
}

func (x PaymentModeEnum_PaymentMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PaymentModeEnum_PaymentMode.Descriptor instead.
func (PaymentModeEnum_PaymentMode) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v9_enums_payment_mode_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible payment modes.
type PaymentModeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PaymentModeEnum) Reset() {
	*x = PaymentModeEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v9_enums_payment_mode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentModeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentModeEnum) ProtoMessage() {}

func (x *PaymentModeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v9_enums_payment_mode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentModeEnum.ProtoReflect.Descriptor instead.
func (*PaymentModeEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v9_enums_payment_mode_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v9_enums_payment_mode_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v9_enums_payment_mode_proto_rawDesc = []byte{
	0x0a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x81, 0x01, 0x0a, 0x0f, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x45,
	0x6e, 0x75, 0x6d, 0x22, 0x6e, 0x0a, 0x0b, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x6f,
	0x64, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01,
	0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4c, 0x49, 0x43, 0x4b, 0x53, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10,
	0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45,
	0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e,
	0x53, 0x10, 0x06, 0x12, 0x0e, 0x0a, 0x0a, 0x47, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x53, 0x54, 0x41,
	0x59, 0x10, 0x07, 0x42, 0xe5, 0x01, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x10, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56,
	0x39, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56,
	0x39, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x56, 0x39, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v9_enums_payment_mode_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v9_enums_payment_mode_proto_rawDescData = file_google_ads_googleads_v9_enums_payment_mode_proto_rawDesc
)

func file_google_ads_googleads_v9_enums_payment_mode_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v9_enums_payment_mode_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v9_enums_payment_mode_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v9_enums_payment_mode_proto_rawDescData)
	})
	return file_google_ads_googleads_v9_enums_payment_mode_proto_rawDescData
}

var file_google_ads_googleads_v9_enums_payment_mode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v9_enums_payment_mode_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v9_enums_payment_mode_proto_goTypes = []interface{}{
	(PaymentModeEnum_PaymentMode)(0), // 0: google.ads.googleads.v9.enums.PaymentModeEnum.PaymentMode
	(*PaymentModeEnum)(nil),          // 1: google.ads.googleads.v9.enums.PaymentModeEnum
}
var file_google_ads_googleads_v9_enums_payment_mode_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v9_enums_payment_mode_proto_init() }
func file_google_ads_googleads_v9_enums_payment_mode_proto_init() {
	if File_google_ads_googleads_v9_enums_payment_mode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v9_enums_payment_mode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentModeEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v9_enums_payment_mode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v9_enums_payment_mode_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v9_enums_payment_mode_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v9_enums_payment_mode_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v9_enums_payment_mode_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v9_enums_payment_mode_proto = out.File
	file_google_ads_googleads_v9_enums_payment_mode_proto_rawDesc = nil
	file_google_ads_googleads_v9_enums_payment_mode_proto_goTypes = nil
	file_google_ads_googleads_v9_enums_payment_mode_proto_depIdxs = nil
}
