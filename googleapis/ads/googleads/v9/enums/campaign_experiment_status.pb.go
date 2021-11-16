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
// source: google/ads/googleads/v9/enums/campaign_experiment_status.proto

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

// Possible statuses of a campaign experiment.
type CampaignExperimentStatusEnum_CampaignExperimentStatus int32

const (
	// The status has not been specified.
	CampaignExperimentStatusEnum_UNSPECIFIED CampaignExperimentStatusEnum_CampaignExperimentStatus = 0
	// Used for return value only. Represents value unknown in this version.
	CampaignExperimentStatusEnum_UNKNOWN CampaignExperimentStatusEnum_CampaignExperimentStatus = 1
	// The experiment campaign is being initialized.
	CampaignExperimentStatusEnum_INITIALIZING CampaignExperimentStatusEnum_CampaignExperimentStatus = 2
	// Initialization of the experiment campaign failed.
	CampaignExperimentStatusEnum_INITIALIZATION_FAILED CampaignExperimentStatusEnum_CampaignExperimentStatus = 8
	// The experiment campaign is fully initialized. The experiment is currently
	// running, scheduled to run in the future or has ended based on its
	// end date. An experiment with the status INITIALIZING will be updated to
	// ENABLED when it is fully created.
	CampaignExperimentStatusEnum_ENABLED CampaignExperimentStatusEnum_CampaignExperimentStatus = 3
	// The experiment campaign was graduated to a stand-alone
	// campaign, existing independently of the experiment.
	CampaignExperimentStatusEnum_GRADUATED CampaignExperimentStatusEnum_CampaignExperimentStatus = 4
	// The experiment is removed.
	CampaignExperimentStatusEnum_REMOVED CampaignExperimentStatusEnum_CampaignExperimentStatus = 5
	// The experiment's changes are being applied to the original campaign.
	// The long running operation returned by the promote method can be polled
	// to see the status of the promotion.
	CampaignExperimentStatusEnum_PROMOTING CampaignExperimentStatusEnum_CampaignExperimentStatus = 6
	// Promote of the experiment campaign failed.
	CampaignExperimentStatusEnum_PROMOTION_FAILED CampaignExperimentStatusEnum_CampaignExperimentStatus = 9
	// The changes of the experiment are promoted to their original campaign.
	CampaignExperimentStatusEnum_PROMOTED CampaignExperimentStatusEnum_CampaignExperimentStatus = 7
	// The experiment was ended manually. It did not end based on its end date.
	CampaignExperimentStatusEnum_ENDED_MANUALLY CampaignExperimentStatusEnum_CampaignExperimentStatus = 10
)

// Enum value maps for CampaignExperimentStatusEnum_CampaignExperimentStatus.
var (
	CampaignExperimentStatusEnum_CampaignExperimentStatus_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "INITIALIZING",
		8:  "INITIALIZATION_FAILED",
		3:  "ENABLED",
		4:  "GRADUATED",
		5:  "REMOVED",
		6:  "PROMOTING",
		9:  "PROMOTION_FAILED",
		7:  "PROMOTED",
		10: "ENDED_MANUALLY",
	}
	CampaignExperimentStatusEnum_CampaignExperimentStatus_value = map[string]int32{
		"UNSPECIFIED":           0,
		"UNKNOWN":               1,
		"INITIALIZING":          2,
		"INITIALIZATION_FAILED": 8,
		"ENABLED":               3,
		"GRADUATED":             4,
		"REMOVED":               5,
		"PROMOTING":             6,
		"PROMOTION_FAILED":      9,
		"PROMOTED":              7,
		"ENDED_MANUALLY":        10,
	}
)

func (x CampaignExperimentStatusEnum_CampaignExperimentStatus) Enum() *CampaignExperimentStatusEnum_CampaignExperimentStatus {
	p := new(CampaignExperimentStatusEnum_CampaignExperimentStatus)
	*p = x
	return p
}

func (x CampaignExperimentStatusEnum_CampaignExperimentStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CampaignExperimentStatusEnum_CampaignExperimentStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v9_enums_campaign_experiment_status_proto_enumTypes[0].Descriptor()
}

func (CampaignExperimentStatusEnum_CampaignExperimentStatus) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v9_enums_campaign_experiment_status_proto_enumTypes[0]
}

func (x CampaignExperimentStatusEnum_CampaignExperimentStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CampaignExperimentStatusEnum_CampaignExperimentStatus.Descriptor instead.
func (CampaignExperimentStatusEnum_CampaignExperimentStatus) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v9_enums_campaign_experiment_status_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible statuses of a campaign experiment.
type CampaignExperimentStatusEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CampaignExperimentStatusEnum) Reset() {
	*x = CampaignExperimentStatusEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v9_enums_campaign_experiment_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CampaignExperimentStatusEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CampaignExperimentStatusEnum) ProtoMessage() {}

func (x *CampaignExperimentStatusEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v9_enums_campaign_experiment_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CampaignExperimentStatusEnum.ProtoReflect.Descriptor instead.
func (*CampaignExperimentStatusEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v9_enums_campaign_experiment_status_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v9_enums_campaign_experiment_status_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v9_enums_campaign_experiment_status_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf6, 0x01,
	0x0a, 0x1c, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xd5,
	0x01, 0x0a, 0x18, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x45, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x4e, 0x49,
	0x54, 0x49, 0x41, 0x4c, 0x49, 0x5a, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x49,
	0x4e, 0x49, 0x54, 0x49, 0x41, 0x4c, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x41,
	0x49, 0x4c, 0x45, 0x44, 0x10, 0x08, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45,
	0x44, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x47, 0x52, 0x41, 0x44, 0x55, 0x41, 0x54, 0x45, 0x44,
	0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x05, 0x12,
	0x0d, 0x0a, 0x09, 0x50, 0x52, 0x4f, 0x4d, 0x4f, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x06, 0x12, 0x14,
	0x0a, 0x10, 0x50, 0x52, 0x4f, 0x4d, 0x4f, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x41, 0x49, 0x4c,
	0x45, 0x44, 0x10, 0x09, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x52, 0x4f, 0x4d, 0x4f, 0x54, 0x45, 0x44,
	0x10, 0x07, 0x12, 0x12, 0x0a, 0x0e, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x5f, 0x4d, 0x41, 0x4e, 0x55,
	0x41, 0x4c, 0x4c, 0x59, 0x10, 0x0a, 0x42, 0xf2, 0x01, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x1d, 0x43, 0x61,
	0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67,
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
	file_google_ads_googleads_v9_enums_campaign_experiment_status_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v9_enums_campaign_experiment_status_proto_rawDescData = file_google_ads_googleads_v9_enums_campaign_experiment_status_proto_rawDesc
)

func file_google_ads_googleads_v9_enums_campaign_experiment_status_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v9_enums_campaign_experiment_status_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v9_enums_campaign_experiment_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v9_enums_campaign_experiment_status_proto_rawDescData)
	})
	return file_google_ads_googleads_v9_enums_campaign_experiment_status_proto_rawDescData
}

var file_google_ads_googleads_v9_enums_campaign_experiment_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v9_enums_campaign_experiment_status_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v9_enums_campaign_experiment_status_proto_goTypes = []interface{}{
	(CampaignExperimentStatusEnum_CampaignExperimentStatus)(0), // 0: google.ads.googleads.v9.enums.CampaignExperimentStatusEnum.CampaignExperimentStatus
	(*CampaignExperimentStatusEnum)(nil),                       // 1: google.ads.googleads.v9.enums.CampaignExperimentStatusEnum
}
var file_google_ads_googleads_v9_enums_campaign_experiment_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v9_enums_campaign_experiment_status_proto_init() }
func file_google_ads_googleads_v9_enums_campaign_experiment_status_proto_init() {
	if File_google_ads_googleads_v9_enums_campaign_experiment_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v9_enums_campaign_experiment_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CampaignExperimentStatusEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v9_enums_campaign_experiment_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v9_enums_campaign_experiment_status_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v9_enums_campaign_experiment_status_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v9_enums_campaign_experiment_status_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v9_enums_campaign_experiment_status_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v9_enums_campaign_experiment_status_proto = out.File
	file_google_ads_googleads_v9_enums_campaign_experiment_status_proto_rawDesc = nil
	file_google_ads_googleads_v9_enums_campaign_experiment_status_proto_goTypes = nil
	file_google_ads_googleads_v9_enums_campaign_experiment_status_proto_depIdxs = nil
}
