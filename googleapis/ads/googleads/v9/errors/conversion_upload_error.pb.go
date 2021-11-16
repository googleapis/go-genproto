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
// source: google/ads/googleads/v9/errors/conversion_upload_error.proto

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

// Enum describing possible conversion upload errors.
type ConversionUploadErrorEnum_ConversionUploadError int32

const (
	// Enum unspecified.
	ConversionUploadErrorEnum_UNSPECIFIED ConversionUploadErrorEnum_ConversionUploadError = 0
	// The received error code is not known in this version.
	ConversionUploadErrorEnum_UNKNOWN ConversionUploadErrorEnum_ConversionUploadError = 1
	// The request contained more than 2000 conversions.
	ConversionUploadErrorEnum_TOO_MANY_CONVERSIONS_IN_REQUEST ConversionUploadErrorEnum_ConversionUploadError = 2
	// The specified gclid could not be decoded.
	ConversionUploadErrorEnum_UNPARSEABLE_GCLID ConversionUploadErrorEnum_ConversionUploadError = 3
	// The specified conversion_date_time is before the event time
	// associated with the given identifier or iOS URL parameter.
	ConversionUploadErrorEnum_CONVERSION_PRECEDES_EVENT ConversionUploadErrorEnum_ConversionUploadError = 42
	// The click associated with the given identifier or iOS URL parameter is
	// either too old to be imported or occurred outside of the click through
	// lookback window for the specified conversion action.
	ConversionUploadErrorEnum_EXPIRED_EVENT ConversionUploadErrorEnum_ConversionUploadError = 43
	// The click associated with the given identifier or iOS URL parameter
	// occurred too recently. Please try uploading again after 6 hours have
	// passed since the click occurred.
	ConversionUploadErrorEnum_TOO_RECENT_EVENT ConversionUploadErrorEnum_ConversionUploadError = 44
	// The click associated with the given identifier or iOS URL parameter could
	// not be found in the system. This can happen if the identifier or iOS URL
	// parameter are collected for non Google Ads clicks.
	ConversionUploadErrorEnum_EVENT_NOT_FOUND ConversionUploadErrorEnum_ConversionUploadError = 45
	// The click associated with the given identifier or iOS URL parameter is
	// owned by a customer account that the uploading customer does not manage.
	ConversionUploadErrorEnum_UNAUTHORIZED_CUSTOMER ConversionUploadErrorEnum_ConversionUploadError = 8
	// No upload eligible conversion action that matches the provided
	// information can be found for the customer.
	ConversionUploadErrorEnum_INVALID_CONVERSION_ACTION ConversionUploadErrorEnum_ConversionUploadError = 9
	// The specified conversion action was created too recently.
	// Please try the upload again after 4-6 hours have passed since the
	// conversion action was created.
	ConversionUploadErrorEnum_TOO_RECENT_CONVERSION_ACTION ConversionUploadErrorEnum_ConversionUploadError = 10
	// The click associated with the given identifier does not contain
	// conversion tracking information.
	ConversionUploadErrorEnum_CONVERSION_TRACKING_NOT_ENABLED_AT_IMPRESSION_TIME ConversionUploadErrorEnum_ConversionUploadError = 11
	// The specified conversion action does not use an external attribution
	// model, but external_attribution_data was set.
	ConversionUploadErrorEnum_EXTERNAL_ATTRIBUTION_DATA_SET_FOR_NON_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION ConversionUploadErrorEnum_ConversionUploadError = 12
	// The specified conversion action uses an external attribution model, but
	// external_attribution_data or one of its contained fields was not set.
	// Both external_attribution_credit and external_attribution_model must be
	// set for externally attributed conversion actions.
	ConversionUploadErrorEnum_EXTERNAL_ATTRIBUTION_DATA_NOT_SET_FOR_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION ConversionUploadErrorEnum_ConversionUploadError = 13
	// Order IDs are not supported for conversion actions which use an external
	// attribution model.
	ConversionUploadErrorEnum_ORDER_ID_NOT_PERMITTED_FOR_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION ConversionUploadErrorEnum_ConversionUploadError = 14
	// A conversion with the same order id and conversion action combination
	// already exists in our system.
	ConversionUploadErrorEnum_ORDER_ID_ALREADY_IN_USE ConversionUploadErrorEnum_ConversionUploadError = 15
	// The request contained two or more conversions with the same order id and
	// conversion action combination.
	ConversionUploadErrorEnum_DUPLICATE_ORDER_ID ConversionUploadErrorEnum_ConversionUploadError = 16
	// The call occurred too recently. Please try uploading again after 12 hours
	// have passed since the call occurred.
	ConversionUploadErrorEnum_TOO_RECENT_CALL ConversionUploadErrorEnum_ConversionUploadError = 17
	// The click that initiated the call is too old for this conversion to be
	// imported.
	ConversionUploadErrorEnum_EXPIRED_CALL ConversionUploadErrorEnum_ConversionUploadError = 18
	// The call or the click leading to the call was not found.
	ConversionUploadErrorEnum_CALL_NOT_FOUND ConversionUploadErrorEnum_ConversionUploadError = 19
	// The specified conversion_date_time is before the call_start_date_time.
	ConversionUploadErrorEnum_CONVERSION_PRECEDES_CALL ConversionUploadErrorEnum_ConversionUploadError = 20
	// The click associated with the call does not contain conversion tracking
	// information.
	ConversionUploadErrorEnum_CONVERSION_TRACKING_NOT_ENABLED_AT_CALL_TIME ConversionUploadErrorEnum_ConversionUploadError = 21
	// The caller’s phone number cannot be parsed. It should be formatted either
	// as E.164 "+16502531234", International "+64 3-331 6005" or US national
	// number "6502531234".
	ConversionUploadErrorEnum_UNPARSEABLE_CALLERS_PHONE_NUMBER ConversionUploadErrorEnum_ConversionUploadError = 22
	// A conversion with this timestamp already exists for this click. To upload
	// another conversion, please use a different timestamp.
	ConversionUploadErrorEnum_CLICK_CONVERSION_ALREADY_EXISTS ConversionUploadErrorEnum_ConversionUploadError = 23
	// A conversion with this timestamp already exists for this call. To upload
	// another conversion, please use a different timestamp.
	ConversionUploadErrorEnum_CALL_CONVERSION_ALREADY_EXISTS ConversionUploadErrorEnum_ConversionUploadError = 24
	// This conversion has the same click and timestamp as another conversion in
	// the request. To upload another conversion for this click, please use a
	// different timestamp.
	ConversionUploadErrorEnum_DUPLICATE_CLICK_CONVERSION_IN_REQUEST ConversionUploadErrorEnum_ConversionUploadError = 25
	// This conversion has the same call and timestamp as another conversion in
	// the request. To upload another conversion for this call, please use a
	// different timestamp.
	ConversionUploadErrorEnum_DUPLICATE_CALL_CONVERSION_IN_REQUEST ConversionUploadErrorEnum_ConversionUploadError = 26
	// The custom variable is not enabled.
	ConversionUploadErrorEnum_CUSTOM_VARIABLE_NOT_ENABLED ConversionUploadErrorEnum_ConversionUploadError = 28
	// The value of the custom variable contains private customer data, such
	// as email addresses or phone numbers.
	ConversionUploadErrorEnum_CUSTOM_VARIABLE_VALUE_CONTAINS_PII ConversionUploadErrorEnum_ConversionUploadError = 29
	// The click associated with the given identifier or iOS URL parameter isn't
	// from the account where conversion tracking is set up.
	ConversionUploadErrorEnum_INVALID_CUSTOMER_FOR_CLICK ConversionUploadErrorEnum_ConversionUploadError = 30
	// The click associated with the given call isn't from the account where
	// conversion tracking is set up.
	ConversionUploadErrorEnum_INVALID_CUSTOMER_FOR_CALL ConversionUploadErrorEnum_ConversionUploadError = 31
	// The conversion can't be uploaded because the conversion source didn't
	// comply with the App Tracking Transparency (ATT) policy or the person who
	// converted didn't consent to tracking.
	ConversionUploadErrorEnum_CONVERSION_NOT_COMPLIANT_WITH_ATT_POLICY ConversionUploadErrorEnum_ConversionUploadError = 32
	// No click was found for the provided user identifiers that could be
	// applied to the specified conversion action.
	ConversionUploadErrorEnum_CLICK_NOT_FOUND ConversionUploadErrorEnum_ConversionUploadError = 33
	// The provided user identifier is not a SHA-256 hash. It is either unhashed
	// or hashed using a different hash function.
	ConversionUploadErrorEnum_INVALID_USER_IDENTIFIER ConversionUploadErrorEnum_ConversionUploadError = 34
	// Conversion actions which use an external attribution model cannot be used
	// with UserIdentifier.
	ConversionUploadErrorEnum_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION_NOT_PERMITTED_WITH_USER_IDENTIFIER ConversionUploadErrorEnum_ConversionUploadError = 35
	// The provided user identifier is not supported. ConversionUploadService
	// only supports hashed_email and hashed_phone_number.
	ConversionUploadErrorEnum_UNSUPPORTED_USER_IDENTIFIER ConversionUploadErrorEnum_ConversionUploadError = 36
	// gbraid and wbraid are both set in the request. Only one is allowed.
	ConversionUploadErrorEnum_GBRAID_WBRAID_BOTH_SET ConversionUploadErrorEnum_ConversionUploadError = 38
	// The specified wbraid could not be decoded.
	ConversionUploadErrorEnum_UNPARSEABLE_WBRAID ConversionUploadErrorEnum_ConversionUploadError = 39
	// The specified gbraid could not be decoded.
	ConversionUploadErrorEnum_UNPARSEABLE_GBRAID ConversionUploadErrorEnum_ConversionUploadError = 40
	// Conversion types which use an external attribution model cannot be used
	// with gbraid or wbraid.
	ConversionUploadErrorEnum_EXTERNALLY_ATTRIBUTED_CONVERSION_TYPE_NOT_PERMITTED_WITH_BRAID ConversionUploadErrorEnum_ConversionUploadError = 41
	// Conversion actions which use the one-per-click counting type cannot be
	// used with gbraid or wbraid.
	ConversionUploadErrorEnum_ONE_PER_CLICK_CONVERSION_ACTION_NOT_PERMITTED_WITH_BRAID ConversionUploadErrorEnum_ConversionUploadError = 46
)

// Enum value maps for ConversionUploadErrorEnum_ConversionUploadError.
var (
	ConversionUploadErrorEnum_ConversionUploadError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "TOO_MANY_CONVERSIONS_IN_REQUEST",
		3:  "UNPARSEABLE_GCLID",
		42: "CONVERSION_PRECEDES_EVENT",
		43: "EXPIRED_EVENT",
		44: "TOO_RECENT_EVENT",
		45: "EVENT_NOT_FOUND",
		8:  "UNAUTHORIZED_CUSTOMER",
		9:  "INVALID_CONVERSION_ACTION",
		10: "TOO_RECENT_CONVERSION_ACTION",
		11: "CONVERSION_TRACKING_NOT_ENABLED_AT_IMPRESSION_TIME",
		12: "EXTERNAL_ATTRIBUTION_DATA_SET_FOR_NON_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION",
		13: "EXTERNAL_ATTRIBUTION_DATA_NOT_SET_FOR_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION",
		14: "ORDER_ID_NOT_PERMITTED_FOR_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION",
		15: "ORDER_ID_ALREADY_IN_USE",
		16: "DUPLICATE_ORDER_ID",
		17: "TOO_RECENT_CALL",
		18: "EXPIRED_CALL",
		19: "CALL_NOT_FOUND",
		20: "CONVERSION_PRECEDES_CALL",
		21: "CONVERSION_TRACKING_NOT_ENABLED_AT_CALL_TIME",
		22: "UNPARSEABLE_CALLERS_PHONE_NUMBER",
		23: "CLICK_CONVERSION_ALREADY_EXISTS",
		24: "CALL_CONVERSION_ALREADY_EXISTS",
		25: "DUPLICATE_CLICK_CONVERSION_IN_REQUEST",
		26: "DUPLICATE_CALL_CONVERSION_IN_REQUEST",
		28: "CUSTOM_VARIABLE_NOT_ENABLED",
		29: "CUSTOM_VARIABLE_VALUE_CONTAINS_PII",
		30: "INVALID_CUSTOMER_FOR_CLICK",
		31: "INVALID_CUSTOMER_FOR_CALL",
		32: "CONVERSION_NOT_COMPLIANT_WITH_ATT_POLICY",
		33: "CLICK_NOT_FOUND",
		34: "INVALID_USER_IDENTIFIER",
		35: "EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION_NOT_PERMITTED_WITH_USER_IDENTIFIER",
		36: "UNSUPPORTED_USER_IDENTIFIER",
		38: "GBRAID_WBRAID_BOTH_SET",
		39: "UNPARSEABLE_WBRAID",
		40: "UNPARSEABLE_GBRAID",
		41: "EXTERNALLY_ATTRIBUTED_CONVERSION_TYPE_NOT_PERMITTED_WITH_BRAID",
		46: "ONE_PER_CLICK_CONVERSION_ACTION_NOT_PERMITTED_WITH_BRAID",
	}
	ConversionUploadErrorEnum_ConversionUploadError_value = map[string]int32{
		"UNSPECIFIED":                     0,
		"UNKNOWN":                         1,
		"TOO_MANY_CONVERSIONS_IN_REQUEST": 2,
		"UNPARSEABLE_GCLID":               3,
		"CONVERSION_PRECEDES_EVENT":       42,
		"EXPIRED_EVENT":                   43,
		"TOO_RECENT_EVENT":                44,
		"EVENT_NOT_FOUND":                 45,
		"UNAUTHORIZED_CUSTOMER":           8,
		"INVALID_CONVERSION_ACTION":       9,
		"TOO_RECENT_CONVERSION_ACTION":    10,
		"CONVERSION_TRACKING_NOT_ENABLED_AT_IMPRESSION_TIME":                            11,
		"EXTERNAL_ATTRIBUTION_DATA_SET_FOR_NON_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION": 12,
		"EXTERNAL_ATTRIBUTION_DATA_NOT_SET_FOR_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION": 13,
		"ORDER_ID_NOT_PERMITTED_FOR_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION":            14,
		"ORDER_ID_ALREADY_IN_USE":                      15,
		"DUPLICATE_ORDER_ID":                           16,
		"TOO_RECENT_CALL":                              17,
		"EXPIRED_CALL":                                 18,
		"CALL_NOT_FOUND":                               19,
		"CONVERSION_PRECEDES_CALL":                     20,
		"CONVERSION_TRACKING_NOT_ENABLED_AT_CALL_TIME": 21,
		"UNPARSEABLE_CALLERS_PHONE_NUMBER":             22,
		"CLICK_CONVERSION_ALREADY_EXISTS":              23,
		"CALL_CONVERSION_ALREADY_EXISTS":               24,
		"DUPLICATE_CLICK_CONVERSION_IN_REQUEST":        25,
		"DUPLICATE_CALL_CONVERSION_IN_REQUEST":         26,
		"CUSTOM_VARIABLE_NOT_ENABLED":                  28,
		"CUSTOM_VARIABLE_VALUE_CONTAINS_PII":           29,
		"INVALID_CUSTOMER_FOR_CLICK":                   30,
		"INVALID_CUSTOMER_FOR_CALL":                    31,
		"CONVERSION_NOT_COMPLIANT_WITH_ATT_POLICY":     32,
		"CLICK_NOT_FOUND":                              33,
		"INVALID_USER_IDENTIFIER":                      34,
		"EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION_NOT_PERMITTED_WITH_USER_IDENTIFIER": 35,
		"UNSUPPORTED_USER_IDENTIFIER": 36,
		"GBRAID_WBRAID_BOTH_SET":      38,
		"UNPARSEABLE_WBRAID":          39,
		"UNPARSEABLE_GBRAID":          40,
		"EXTERNALLY_ATTRIBUTED_CONVERSION_TYPE_NOT_PERMITTED_WITH_BRAID": 41,
		"ONE_PER_CLICK_CONVERSION_ACTION_NOT_PERMITTED_WITH_BRAID":       46,
	}
)

func (x ConversionUploadErrorEnum_ConversionUploadError) Enum() *ConversionUploadErrorEnum_ConversionUploadError {
	p := new(ConversionUploadErrorEnum_ConversionUploadError)
	*p = x
	return p
}

func (x ConversionUploadErrorEnum_ConversionUploadError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConversionUploadErrorEnum_ConversionUploadError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v9_errors_conversion_upload_error_proto_enumTypes[0].Descriptor()
}

func (ConversionUploadErrorEnum_ConversionUploadError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v9_errors_conversion_upload_error_proto_enumTypes[0]
}

func (x ConversionUploadErrorEnum_ConversionUploadError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConversionUploadErrorEnum_ConversionUploadError.Descriptor instead.
func (ConversionUploadErrorEnum_ConversionUploadError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v9_errors_conversion_upload_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible conversion upload errors.
type ConversionUploadErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConversionUploadErrorEnum) Reset() {
	*x = ConversionUploadErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v9_errors_conversion_upload_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConversionUploadErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversionUploadErrorEnum) ProtoMessage() {}

func (x *ConversionUploadErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v9_errors_conversion_upload_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversionUploadErrorEnum.ProtoReflect.Descriptor instead.
func (*ConversionUploadErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v9_errors_conversion_upload_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v9_errors_conversion_upload_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v9_errors_conversion_upload_error_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x0c, 0x0a,
	0x19, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xfb, 0x0b, 0x0a, 0x15, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x01, 0x12, 0x23, 0x0a, 0x1f, 0x54, 0x4f, 0x4f, 0x5f, 0x4d, 0x41, 0x4e, 0x59, 0x5f, 0x43,
	0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x53, 0x5f, 0x49, 0x4e, 0x5f, 0x52, 0x45,
	0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x4e, 0x50, 0x41, 0x52,
	0x53, 0x45, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x47, 0x43, 0x4c, 0x49, 0x44, 0x10, 0x03, 0x12, 0x1d,
	0x0a, 0x19, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x52, 0x45,
	0x43, 0x45, 0x44, 0x45, 0x53, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x10, 0x2a, 0x12, 0x11, 0x0a,
	0x0d, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x10, 0x2b,
	0x12, 0x14, 0x0a, 0x10, 0x54, 0x4f, 0x4f, 0x5f, 0x52, 0x45, 0x43, 0x45, 0x4e, 0x54, 0x5f, 0x45,
	0x56, 0x45, 0x4e, 0x54, 0x10, 0x2c, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f,
	0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x2d, 0x12, 0x19, 0x0a, 0x15, 0x55,
	0x4e, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x45, 0x44, 0x5f, 0x43, 0x55, 0x53, 0x54,
	0x4f, 0x4d, 0x45, 0x52, 0x10, 0x08, 0x12, 0x1d, 0x0a, 0x19, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x09, 0x12, 0x20, 0x0a, 0x1c, 0x54, 0x4f, 0x4f, 0x5f, 0x52, 0x45, 0x43,
	0x45, 0x4e, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x41,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0a, 0x12, 0x36, 0x0a, 0x32, 0x43, 0x4f, 0x4e, 0x56, 0x45,
	0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x4e,
	0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x5f, 0x41, 0x54, 0x5f, 0x49, 0x4d,
	0x50, 0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x0b, 0x12,
	0x51, 0x0a, 0x4d, 0x45, 0x58, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x41, 0x54, 0x54, 0x52,
	0x49, 0x42, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x53, 0x45, 0x54,
	0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x4e, 0x4f, 0x4e, 0x5f, 0x45, 0x58, 0x54, 0x45, 0x52, 0x4e, 0x41,
	0x4c, 0x4c, 0x59, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x44, 0x5f, 0x43,
	0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x10, 0x0c, 0x12, 0x51, 0x0a, 0x4d, 0x45, 0x58, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x41,
	0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f,
	0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x45, 0x58, 0x54, 0x45,
	0x52, 0x4e, 0x41, 0x4c, 0x4c, 0x59, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45,
	0x44, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x0d, 0x12, 0x46, 0x0a, 0x42, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x49,
	0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x54, 0x54, 0x45, 0x44, 0x5f,
	0x46, 0x4f, 0x52, 0x5f, 0x45, 0x58, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x4c, 0x59, 0x5f, 0x41,
	0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x44, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52,
	0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0e, 0x12, 0x1b, 0x0a,
	0x17, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x49, 0x44, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44,
	0x59, 0x5f, 0x49, 0x4e, 0x5f, 0x55, 0x53, 0x45, 0x10, 0x0f, 0x12, 0x16, 0x0a, 0x12, 0x44, 0x55,
	0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x49, 0x44,
	0x10, 0x10, 0x12, 0x13, 0x0a, 0x0f, 0x54, 0x4f, 0x4f, 0x5f, 0x52, 0x45, 0x43, 0x45, 0x4e, 0x54,
	0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x10, 0x11, 0x12, 0x10, 0x0a, 0x0c, 0x45, 0x58, 0x50, 0x49, 0x52,
	0x45, 0x44, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x10, 0x12, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x41, 0x4c,
	0x4c, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x13, 0x12, 0x1c, 0x0a,
	0x18, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x52, 0x45, 0x43,
	0x45, 0x44, 0x45, 0x53, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x10, 0x14, 0x12, 0x30, 0x0a, 0x2c, 0x43,
	0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x49,
	0x4e, 0x47, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x5f, 0x41,
	0x54, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x15, 0x12, 0x24, 0x0a,
	0x20, 0x55, 0x4e, 0x50, 0x41, 0x52, 0x53, 0x45, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x43, 0x41, 0x4c,
	0x4c, 0x45, 0x52, 0x53, 0x5f, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45,
	0x52, 0x10, 0x16, 0x12, 0x23, 0x0a, 0x1f, 0x43, 0x4c, 0x49, 0x43, 0x4b, 0x5f, 0x43, 0x4f, 0x4e,
	0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f,
	0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x17, 0x12, 0x22, 0x0a, 0x1e, 0x43, 0x41, 0x4c, 0x4c,
	0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x4c, 0x52, 0x45,
	0x41, 0x44, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x18, 0x12, 0x29, 0x0a, 0x25,
	0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x4c, 0x49, 0x43, 0x4b, 0x5f,
	0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x5f, 0x52, 0x45,
	0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x19, 0x12, 0x28, 0x0a, 0x24, 0x44, 0x55, 0x50, 0x4c, 0x49,
	0x43, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52,
	0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10,
	0x1a, 0x12, 0x1f, 0x0a, 0x1b, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x5f, 0x56, 0x41, 0x52, 0x49,
	0x41, 0x42, 0x4c, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44,
	0x10, 0x1c, 0x12, 0x26, 0x0a, 0x22, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x5f, 0x56, 0x41, 0x52,
	0x49, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x54,
	0x41, 0x49, 0x4e, 0x53, 0x5f, 0x50, 0x49, 0x49, 0x10, 0x1d, 0x12, 0x1e, 0x0a, 0x1a, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x46,
	0x4f, 0x52, 0x5f, 0x43, 0x4c, 0x49, 0x43, 0x4b, 0x10, 0x1e, 0x12, 0x1d, 0x0a, 0x19, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x46,
	0x4f, 0x52, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x10, 0x1f, 0x12, 0x2c, 0x0a, 0x28, 0x43, 0x4f, 0x4e,
	0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x4f, 0x4d, 0x50,
	0x4c, 0x49, 0x41, 0x4e, 0x54, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x41, 0x54, 0x54, 0x5f, 0x50,
	0x4f, 0x4c, 0x49, 0x43, 0x59, 0x10, 0x20, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4c, 0x49, 0x43, 0x4b,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x21, 0x12, 0x1b, 0x0a, 0x17,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x44, 0x45,
	0x4e, 0x54, 0x49, 0x46, 0x49, 0x45, 0x52, 0x10, 0x22, 0x12, 0x4e, 0x0a, 0x4a, 0x45, 0x58, 0x54,
	0x45, 0x52, 0x4e, 0x41, 0x4c, 0x4c, 0x59, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54,
	0x45, 0x44, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x54, 0x54,
	0x45, 0x44, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x44, 0x45,
	0x4e, 0x54, 0x49, 0x46, 0x49, 0x45, 0x52, 0x10, 0x23, 0x12, 0x1f, 0x0a, 0x1b, 0x55, 0x4e, 0x53,
	0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x44,
	0x45, 0x4e, 0x54, 0x49, 0x46, 0x49, 0x45, 0x52, 0x10, 0x24, 0x12, 0x1a, 0x0a, 0x16, 0x47, 0x42,
	0x52, 0x41, 0x49, 0x44, 0x5f, 0x57, 0x42, 0x52, 0x41, 0x49, 0x44, 0x5f, 0x42, 0x4f, 0x54, 0x48,
	0x5f, 0x53, 0x45, 0x54, 0x10, 0x26, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x4e, 0x50, 0x41, 0x52, 0x53,
	0x45, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x57, 0x42, 0x52, 0x41, 0x49, 0x44, 0x10, 0x27, 0x12, 0x16,
	0x0a, 0x12, 0x55, 0x4e, 0x50, 0x41, 0x52, 0x53, 0x45, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x47, 0x42,
	0x52, 0x41, 0x49, 0x44, 0x10, 0x28, 0x12, 0x42, 0x0a, 0x3e, 0x45, 0x58, 0x54, 0x45, 0x52, 0x4e,
	0x41, 0x4c, 0x4c, 0x59, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x44, 0x5f,
	0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x4e, 0x4f, 0x54, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x54, 0x54, 0x45, 0x44, 0x5f, 0x57, 0x49,
	0x54, 0x48, 0x5f, 0x42, 0x52, 0x41, 0x49, 0x44, 0x10, 0x29, 0x12, 0x3c, 0x0a, 0x38, 0x4f, 0x4e,
	0x45, 0x5f, 0x50, 0x45, 0x52, 0x5f, 0x43, 0x4c, 0x49, 0x43, 0x4b, 0x5f, 0x43, 0x4f, 0x4e, 0x56,
	0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x54, 0x54, 0x45, 0x44, 0x5f, 0x57, 0x49, 0x54, 0x48,
	0x5f, 0x42, 0x52, 0x41, 0x49, 0x44, 0x10, 0x2e, 0x42, 0xf5, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42,
	0x1a, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x2e, 0x56, 0x39, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x5c, 0x56, 0x39, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x39, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v9_errors_conversion_upload_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v9_errors_conversion_upload_error_proto_rawDescData = file_google_ads_googleads_v9_errors_conversion_upload_error_proto_rawDesc
)

func file_google_ads_googleads_v9_errors_conversion_upload_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v9_errors_conversion_upload_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v9_errors_conversion_upload_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v9_errors_conversion_upload_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v9_errors_conversion_upload_error_proto_rawDescData
}

var file_google_ads_googleads_v9_errors_conversion_upload_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v9_errors_conversion_upload_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v9_errors_conversion_upload_error_proto_goTypes = []interface{}{
	(ConversionUploadErrorEnum_ConversionUploadError)(0), // 0: google.ads.googleads.v9.errors.ConversionUploadErrorEnum.ConversionUploadError
	(*ConversionUploadErrorEnum)(nil),                    // 1: google.ads.googleads.v9.errors.ConversionUploadErrorEnum
}
var file_google_ads_googleads_v9_errors_conversion_upload_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v9_errors_conversion_upload_error_proto_init() }
func file_google_ads_googleads_v9_errors_conversion_upload_error_proto_init() {
	if File_google_ads_googleads_v9_errors_conversion_upload_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v9_errors_conversion_upload_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConversionUploadErrorEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v9_errors_conversion_upload_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v9_errors_conversion_upload_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v9_errors_conversion_upload_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v9_errors_conversion_upload_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v9_errors_conversion_upload_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v9_errors_conversion_upload_error_proto = out.File
	file_google_ads_googleads_v9_errors_conversion_upload_error_proto_rawDesc = nil
	file_google_ads_googleads_v9_errors_conversion_upload_error_proto_goTypes = nil
	file_google_ads_googleads_v9_errors_conversion_upload_error_proto_depIdxs = nil
}
