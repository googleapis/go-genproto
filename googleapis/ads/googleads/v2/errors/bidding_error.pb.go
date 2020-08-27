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
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: google/ads/googleads/v2/errors/bidding_error.proto

package errors

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Enum describing possible bidding errors.
type BiddingErrorEnum_BiddingError int32

const (
	// Enum unspecified.
	BiddingErrorEnum_UNSPECIFIED BiddingErrorEnum_BiddingError = 0
	// The received error code is not known in this version.
	BiddingErrorEnum_UNKNOWN BiddingErrorEnum_BiddingError = 1
	// Cannot transition to new bidding strategy.
	BiddingErrorEnum_BIDDING_STRATEGY_TRANSITION_NOT_ALLOWED BiddingErrorEnum_BiddingError = 2
	// Cannot attach bidding strategy to campaign.
	BiddingErrorEnum_CANNOT_ATTACH_BIDDING_STRATEGY_TO_CAMPAIGN BiddingErrorEnum_BiddingError = 7
	// Bidding strategy is not supported or cannot be used as anonymous.
	BiddingErrorEnum_INVALID_ANONYMOUS_BIDDING_STRATEGY_TYPE BiddingErrorEnum_BiddingError = 10
	// The type does not match the named strategy's type.
	BiddingErrorEnum_INVALID_BIDDING_STRATEGY_TYPE BiddingErrorEnum_BiddingError = 14
	// The bid is invalid.
	BiddingErrorEnum_INVALID_BID BiddingErrorEnum_BiddingError = 17
	// Bidding strategy is not available for the account type.
	BiddingErrorEnum_BIDDING_STRATEGY_NOT_AVAILABLE_FOR_ACCOUNT_TYPE BiddingErrorEnum_BiddingError = 18
	// Conversion tracking is not enabled in the campaign that has YouTube
	// Video Builder transitions.
	BiddingErrorEnum_CONVERSION_TRACKING_NOT_ENABLED BiddingErrorEnum_BiddingError = 19
	// Not enough conversions tracked for YouTube Video Builder transitions.
	BiddingErrorEnum_NOT_ENOUGH_CONVERSIONS BiddingErrorEnum_BiddingError = 20
	// Campaign can not be created with given bidding strategy. It can be
	// transitioned to the strategy, once eligible.
	BiddingErrorEnum_CANNOT_CREATE_CAMPAIGN_WITH_BIDDING_STRATEGY BiddingErrorEnum_BiddingError = 21
	// Cannot target content network only as campaign uses Page One Promoted
	// bidding strategy.
	BiddingErrorEnum_CANNOT_TARGET_CONTENT_NETWORK_ONLY_WITH_CAMPAIGN_LEVEL_POP_BIDDING_STRATEGY BiddingErrorEnum_BiddingError = 23
	// Budget Optimizer and Target Spend bidding strategies are not supported
	// for campaigns with AdSchedule targeting.
	BiddingErrorEnum_BIDDING_STRATEGY_NOT_SUPPORTED_WITH_AD_SCHEDULE BiddingErrorEnum_BiddingError = 24
	// Pay per conversion is not available to all the customer, only few
	// customers on the allow-list can use this.
	BiddingErrorEnum_PAY_PER_CONVERSION_NOT_AVAILABLE_FOR_CUSTOMER BiddingErrorEnum_BiddingError = 25
	// Pay per conversion is not allowed with Target CPA.
	BiddingErrorEnum_PAY_PER_CONVERSION_NOT_ALLOWED_WITH_TARGET_CPA BiddingErrorEnum_BiddingError = 26
	// Cannot set bidding strategy to Manual CPM for search network only
	// campaigns.
	BiddingErrorEnum_BIDDING_STRATEGY_NOT_ALLOWED_FOR_SEARCH_ONLY_CAMPAIGNS BiddingErrorEnum_BiddingError = 27
	// The bidding strategy is not supported for use in drafts or experiments.
	BiddingErrorEnum_BIDDING_STRATEGY_NOT_SUPPORTED_IN_DRAFTS_OR_EXPERIMENTS BiddingErrorEnum_BiddingError = 28
	// Bidding strategy type does not support product type ad group criterion.
	BiddingErrorEnum_BIDDING_STRATEGY_TYPE_DOES_NOT_SUPPORT_PRODUCT_TYPE_ADGROUP_CRITERION BiddingErrorEnum_BiddingError = 29
	// Bid amount is too small.
	BiddingErrorEnum_BID_TOO_SMALL BiddingErrorEnum_BiddingError = 30
	// Bid amount is too big.
	BiddingErrorEnum_BID_TOO_BIG BiddingErrorEnum_BiddingError = 31
	// Bid has too many fractional digit precision.
	BiddingErrorEnum_BID_TOO_MANY_FRACTIONAL_DIGITS BiddingErrorEnum_BiddingError = 32
	// Invalid domain name specified.
	BiddingErrorEnum_INVALID_DOMAIN_NAME BiddingErrorEnum_BiddingError = 33
	// The field is not compatible with the payment mode.
	BiddingErrorEnum_NOT_COMPATIBLE_WITH_PAYMENT_MODE BiddingErrorEnum_BiddingError = 34
	// The field is not compatible with the budget type.
	BiddingErrorEnum_NOT_COMPATIBLE_WITH_BUDGET_TYPE BiddingErrorEnum_BiddingError = 35
	// The field is not compatible with the bidding strategy type.
	BiddingErrorEnum_NOT_COMPATIBLE_WITH_BIDDING_STRATEGY_TYPE BiddingErrorEnum_BiddingError = 36
)

// Enum value maps for BiddingErrorEnum_BiddingError.
var (
	BiddingErrorEnum_BiddingError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "BIDDING_STRATEGY_TRANSITION_NOT_ALLOWED",
		7:  "CANNOT_ATTACH_BIDDING_STRATEGY_TO_CAMPAIGN",
		10: "INVALID_ANONYMOUS_BIDDING_STRATEGY_TYPE",
		14: "INVALID_BIDDING_STRATEGY_TYPE",
		17: "INVALID_BID",
		18: "BIDDING_STRATEGY_NOT_AVAILABLE_FOR_ACCOUNT_TYPE",
		19: "CONVERSION_TRACKING_NOT_ENABLED",
		20: "NOT_ENOUGH_CONVERSIONS",
		21: "CANNOT_CREATE_CAMPAIGN_WITH_BIDDING_STRATEGY",
		23: "CANNOT_TARGET_CONTENT_NETWORK_ONLY_WITH_CAMPAIGN_LEVEL_POP_BIDDING_STRATEGY",
		24: "BIDDING_STRATEGY_NOT_SUPPORTED_WITH_AD_SCHEDULE",
		25: "PAY_PER_CONVERSION_NOT_AVAILABLE_FOR_CUSTOMER",
		26: "PAY_PER_CONVERSION_NOT_ALLOWED_WITH_TARGET_CPA",
		27: "BIDDING_STRATEGY_NOT_ALLOWED_FOR_SEARCH_ONLY_CAMPAIGNS",
		28: "BIDDING_STRATEGY_NOT_SUPPORTED_IN_DRAFTS_OR_EXPERIMENTS",
		29: "BIDDING_STRATEGY_TYPE_DOES_NOT_SUPPORT_PRODUCT_TYPE_ADGROUP_CRITERION",
		30: "BID_TOO_SMALL",
		31: "BID_TOO_BIG",
		32: "BID_TOO_MANY_FRACTIONAL_DIGITS",
		33: "INVALID_DOMAIN_NAME",
		34: "NOT_COMPATIBLE_WITH_PAYMENT_MODE",
		35: "NOT_COMPATIBLE_WITH_BUDGET_TYPE",
		36: "NOT_COMPATIBLE_WITH_BIDDING_STRATEGY_TYPE",
	}
	BiddingErrorEnum_BiddingError_value = map[string]int32{
		"UNSPECIFIED": 0,
		"UNKNOWN":     1,
		"BIDDING_STRATEGY_TRANSITION_NOT_ALLOWED":                                     2,
		"CANNOT_ATTACH_BIDDING_STRATEGY_TO_CAMPAIGN":                                  7,
		"INVALID_ANONYMOUS_BIDDING_STRATEGY_TYPE":                                     10,
		"INVALID_BIDDING_STRATEGY_TYPE":                                               14,
		"INVALID_BID":                                                                 17,
		"BIDDING_STRATEGY_NOT_AVAILABLE_FOR_ACCOUNT_TYPE":                             18,
		"CONVERSION_TRACKING_NOT_ENABLED":                                             19,
		"NOT_ENOUGH_CONVERSIONS":                                                      20,
		"CANNOT_CREATE_CAMPAIGN_WITH_BIDDING_STRATEGY":                                21,
		"CANNOT_TARGET_CONTENT_NETWORK_ONLY_WITH_CAMPAIGN_LEVEL_POP_BIDDING_STRATEGY": 23,
		"BIDDING_STRATEGY_NOT_SUPPORTED_WITH_AD_SCHEDULE":                             24,
		"PAY_PER_CONVERSION_NOT_AVAILABLE_FOR_CUSTOMER":                               25,
		"PAY_PER_CONVERSION_NOT_ALLOWED_WITH_TARGET_CPA":                              26,
		"BIDDING_STRATEGY_NOT_ALLOWED_FOR_SEARCH_ONLY_CAMPAIGNS":                      27,
		"BIDDING_STRATEGY_NOT_SUPPORTED_IN_DRAFTS_OR_EXPERIMENTS":                     28,
		"BIDDING_STRATEGY_TYPE_DOES_NOT_SUPPORT_PRODUCT_TYPE_ADGROUP_CRITERION":       29,
		"BID_TOO_SMALL":                             30,
		"BID_TOO_BIG":                               31,
		"BID_TOO_MANY_FRACTIONAL_DIGITS":            32,
		"INVALID_DOMAIN_NAME":                       33,
		"NOT_COMPATIBLE_WITH_PAYMENT_MODE":          34,
		"NOT_COMPATIBLE_WITH_BUDGET_TYPE":           35,
		"NOT_COMPATIBLE_WITH_BIDDING_STRATEGY_TYPE": 36,
	}
)

func (x BiddingErrorEnum_BiddingError) Enum() *BiddingErrorEnum_BiddingError {
	p := new(BiddingErrorEnum_BiddingError)
	*p = x
	return p
}

func (x BiddingErrorEnum_BiddingError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BiddingErrorEnum_BiddingError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v2_errors_bidding_error_proto_enumTypes[0].Descriptor()
}

func (BiddingErrorEnum_BiddingError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v2_errors_bidding_error_proto_enumTypes[0]
}

func (x BiddingErrorEnum_BiddingError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BiddingErrorEnum_BiddingError.Descriptor instead.
func (BiddingErrorEnum_BiddingError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v2_errors_bidding_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible bidding errors.
type BiddingErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BiddingErrorEnum) Reset() {
	*x = BiddingErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v2_errors_bidding_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BiddingErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BiddingErrorEnum) ProtoMessage() {}

func (x *BiddingErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v2_errors_bidding_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BiddingErrorEnum.ProtoReflect.Descriptor instead.
func (*BiddingErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v2_errors_bidding_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v2_errors_bidding_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v2_errors_bidding_error_proto_rawDesc = []byte{
	0x0a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x2f, 0x62, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb3, 0x08, 0x0a, 0x10, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x9e, 0x08, 0x0a, 0x0c, 0x42, 0x69, 0x64, 0x64,
	0x69, 0x6e, 0x67, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x2b, 0x0a, 0x27, 0x42, 0x49, 0x44, 0x44, 0x49, 0x4e,
	0x47, 0x5f, 0x53, 0x54, 0x52, 0x41, 0x54, 0x45, 0x47, 0x59, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53,
	0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45,
	0x44, 0x10, 0x02, 0x12, 0x2e, 0x0a, 0x2a, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x54,
	0x54, 0x41, 0x43, 0x48, 0x5f, 0x42, 0x49, 0x44, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x52,
	0x41, 0x54, 0x45, 0x47, 0x59, 0x5f, 0x54, 0x4f, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47,
	0x4e, 0x10, 0x07, 0x12, 0x2b, 0x0a, 0x27, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x41,
	0x4e, 0x4f, 0x4e, 0x59, 0x4d, 0x4f, 0x55, 0x53, 0x5f, 0x42, 0x49, 0x44, 0x44, 0x49, 0x4e, 0x47,
	0x5f, 0x53, 0x54, 0x52, 0x41, 0x54, 0x45, 0x47, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x0a,
	0x12, 0x21, 0x0a, 0x1d, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x42, 0x49, 0x44, 0x44,
	0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x52, 0x41, 0x54, 0x45, 0x47, 0x59, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x10, 0x0e, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x42,
	0x49, 0x44, 0x10, 0x11, 0x12, 0x33, 0x0a, 0x2f, 0x42, 0x49, 0x44, 0x44, 0x49, 0x4e, 0x47, 0x5f,
	0x53, 0x54, 0x52, 0x41, 0x54, 0x45, 0x47, 0x59, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x56, 0x41,
	0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55,
	0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x12, 0x12, 0x23, 0x0a, 0x1f, 0x43, 0x4f, 0x4e,
	0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x49, 0x4e, 0x47,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x13, 0x12, 0x1a,
	0x0a, 0x16, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x4f, 0x55, 0x47, 0x48, 0x5f, 0x43, 0x4f, 0x4e,
	0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x14, 0x12, 0x30, 0x0a, 0x2c, 0x43, 0x41,
	0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x41, 0x4d, 0x50,
	0x41, 0x49, 0x47, 0x4e, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x42, 0x49, 0x44, 0x44, 0x49, 0x4e,
	0x47, 0x5f, 0x53, 0x54, 0x52, 0x41, 0x54, 0x45, 0x47, 0x59, 0x10, 0x15, 0x12, 0x4f, 0x0a, 0x4b,
	0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x43, 0x4f,
	0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x4f, 0x4e,
	0x4c, 0x59, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e,
	0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x50, 0x4f, 0x50, 0x5f, 0x42, 0x49, 0x44, 0x44, 0x49,
	0x4e, 0x47, 0x5f, 0x53, 0x54, 0x52, 0x41, 0x54, 0x45, 0x47, 0x59, 0x10, 0x17, 0x12, 0x33, 0x0a,
	0x2f, 0x42, 0x49, 0x44, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x52, 0x41, 0x54, 0x45, 0x47,
	0x59, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f,
	0x57, 0x49, 0x54, 0x48, 0x5f, 0x41, 0x44, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45,
	0x10, 0x18, 0x12, 0x31, 0x0a, 0x2d, 0x50, 0x41, 0x59, 0x5f, 0x50, 0x45, 0x52, 0x5f, 0x43, 0x4f,
	0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x56, 0x41,
	0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f,
	0x4d, 0x45, 0x52, 0x10, 0x19, 0x12, 0x32, 0x0a, 0x2e, 0x50, 0x41, 0x59, 0x5f, 0x50, 0x45, 0x52,
	0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x54, 0x41, 0x52,
	0x47, 0x45, 0x54, 0x5f, 0x43, 0x50, 0x41, 0x10, 0x1a, 0x12, 0x3a, 0x0a, 0x36, 0x42, 0x49, 0x44,
	0x44, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x52, 0x41, 0x54, 0x45, 0x47, 0x59, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x53, 0x45,
	0x41, 0x52, 0x43, 0x48, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49,
	0x47, 0x4e, 0x53, 0x10, 0x1b, 0x12, 0x3b, 0x0a, 0x37, 0x42, 0x49, 0x44, 0x44, 0x49, 0x4e, 0x47,
	0x5f, 0x53, 0x54, 0x52, 0x41, 0x54, 0x45, 0x47, 0x59, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x55,
	0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x49, 0x4e, 0x5f, 0x44, 0x52, 0x41, 0x46, 0x54,
	0x53, 0x5f, 0x4f, 0x52, 0x5f, 0x45, 0x58, 0x50, 0x45, 0x52, 0x49, 0x4d, 0x45, 0x4e, 0x54, 0x53,
	0x10, 0x1c, 0x12, 0x49, 0x0a, 0x45, 0x42, 0x49, 0x44, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54,
	0x52, 0x41, 0x54, 0x45, 0x47, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x4f, 0x45, 0x53,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x50, 0x52, 0x4f,
	0x44, 0x55, 0x43, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x44, 0x47, 0x52, 0x4f, 0x55,
	0x50, 0x5f, 0x43, 0x52, 0x49, 0x54, 0x45, 0x52, 0x49, 0x4f, 0x4e, 0x10, 0x1d, 0x12, 0x11, 0x0a,
	0x0d, 0x42, 0x49, 0x44, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x53, 0x4d, 0x41, 0x4c, 0x4c, 0x10, 0x1e,
	0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x49, 0x44, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x42, 0x49, 0x47, 0x10,
	0x1f, 0x12, 0x22, 0x0a, 0x1e, 0x42, 0x49, 0x44, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x4d, 0x41, 0x4e,
	0x59, 0x5f, 0x46, 0x52, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x5f, 0x44, 0x49, 0x47,
	0x49, 0x54, 0x53, 0x10, 0x20, 0x12, 0x17, 0x0a, 0x13, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x5f, 0x44, 0x4f, 0x4d, 0x41, 0x49, 0x4e, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x21, 0x12, 0x24,
	0x0a, 0x20, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x54, 0x49, 0x42, 0x4c, 0x45,
	0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x4f,
	0x44, 0x45, 0x10, 0x22, 0x12, 0x23, 0x0a, 0x1f, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x4f, 0x4d, 0x50,
	0x41, 0x54, 0x49, 0x42, 0x4c, 0x45, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x42, 0x55, 0x44, 0x47,
	0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x23, 0x12, 0x2d, 0x0a, 0x29, 0x4e, 0x4f, 0x54,
	0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x54, 0x49, 0x42, 0x4c, 0x45, 0x5f, 0x57, 0x49, 0x54, 0x48,
	0x5f, 0x42, 0x49, 0x44, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x52, 0x41, 0x54, 0x45, 0x47,
	0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x24, 0x42, 0xec, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42,
	0x11, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41,
	0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x32, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x32, 0x5c, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x32, 0x3a,
	0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v2_errors_bidding_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v2_errors_bidding_error_proto_rawDescData = file_google_ads_googleads_v2_errors_bidding_error_proto_rawDesc
)

func file_google_ads_googleads_v2_errors_bidding_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v2_errors_bidding_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v2_errors_bidding_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v2_errors_bidding_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v2_errors_bidding_error_proto_rawDescData
}

var file_google_ads_googleads_v2_errors_bidding_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v2_errors_bidding_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v2_errors_bidding_error_proto_goTypes = []interface{}{
	(BiddingErrorEnum_BiddingError)(0), // 0: google.ads.googleads.v2.errors.BiddingErrorEnum.BiddingError
	(*BiddingErrorEnum)(nil),           // 1: google.ads.googleads.v2.errors.BiddingErrorEnum
}
var file_google_ads_googleads_v2_errors_bidding_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v2_errors_bidding_error_proto_init() }
func file_google_ads_googleads_v2_errors_bidding_error_proto_init() {
	if File_google_ads_googleads_v2_errors_bidding_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v2_errors_bidding_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BiddingErrorEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v2_errors_bidding_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v2_errors_bidding_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v2_errors_bidding_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v2_errors_bidding_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v2_errors_bidding_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v2_errors_bidding_error_proto = out.File
	file_google_ads_googleads_v2_errors_bidding_error_proto_rawDesc = nil
	file_google_ads_googleads_v2_errors_bidding_error_proto_goTypes = nil
	file_google_ads_googleads_v2_errors_bidding_error_proto_depIdxs = nil
}
