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
// 	protoc-gen-go v1.22.0
// 	protoc        v3.12.3
// source: google/ads/googleads/v2/resources/billing_setup.proto

package resources

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v2/enums"
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

// A billing setup, which associates a payments account and an advertiser. A
// billing setup is specific to one advertiser.
type BillingSetup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the billing setup.
	// BillingSetup resource names have the form:
	//
	// `customers/{customer_id}/billingSetups/{billing_setup_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the billing setup.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Output only. The status of the billing setup.
	Status enums.BillingSetupStatusEnum_BillingSetupStatus `protobuf:"varint,3,opt,name=status,proto3,enum=google.ads.googleads.v2.enums.BillingSetupStatusEnum_BillingSetupStatus" json:"status,omitempty"`
	// Immutable. The resource name of the payments account associated with this billing
	// setup. Payments resource names have the form:
	//
	// `customers/{customer_id}/paymentsAccounts/{payments_account_id}`
	// When setting up billing, this is used to signup with an existing payments
	// account (and then payments_account_info should not be set).
	// When getting a billing setup, this and payments_account_info will be
	// populated.
	PaymentsAccount *wrappers.StringValue `protobuf:"bytes,11,opt,name=payments_account,json=paymentsAccount,proto3" json:"payments_account,omitempty"`
	// Immutable. The payments account information associated with this billing setup.
	// When setting up billing, this is used to signup with a new payments account
	// (and then payments_account should not be set).
	// When getting a billing setup, this and payments_account will be
	// populated.
	PaymentsAccountInfo *BillingSetup_PaymentsAccountInfo `protobuf:"bytes,12,opt,name=payments_account_info,json=paymentsAccountInfo,proto3" json:"payments_account_info,omitempty"`
	// When creating a new billing setup, this is when the setup should take
	// effect. NOW is the only acceptable start time if the customer doesn't have
	// any approved setups.
	//
	// When fetching an existing billing setup, this is the requested start time.
	// However, if the setup was approved (see status) after the requested start
	// time, then this is the approval time.
	//
	// Types that are assignable to StartTime:
	//	*BillingSetup_StartDateTime
	//	*BillingSetup_StartTimeType
	StartTime isBillingSetup_StartTime `protobuf_oneof:"start_time"`
	// When the billing setup ends / ended. This is either FOREVER or the start
	// time of the next scheduled billing setup.
	//
	// Types that are assignable to EndTime:
	//	*BillingSetup_EndDateTime
	//	*BillingSetup_EndTimeType
	EndTime isBillingSetup_EndTime `protobuf_oneof:"end_time"`
}

func (x *BillingSetup) Reset() {
	*x = BillingSetup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v2_resources_billing_setup_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BillingSetup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillingSetup) ProtoMessage() {}

func (x *BillingSetup) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v2_resources_billing_setup_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillingSetup.ProtoReflect.Descriptor instead.
func (*BillingSetup) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v2_resources_billing_setup_proto_rawDescGZIP(), []int{0}
}

func (x *BillingSetup) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *BillingSetup) GetId() *wrappers.Int64Value {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *BillingSetup) GetStatus() enums.BillingSetupStatusEnum_BillingSetupStatus {
	if x != nil {
		return x.Status
	}
	return enums.BillingSetupStatusEnum_UNSPECIFIED
}

func (x *BillingSetup) GetPaymentsAccount() *wrappers.StringValue {
	if x != nil {
		return x.PaymentsAccount
	}
	return nil
}

func (x *BillingSetup) GetPaymentsAccountInfo() *BillingSetup_PaymentsAccountInfo {
	if x != nil {
		return x.PaymentsAccountInfo
	}
	return nil
}

func (m *BillingSetup) GetStartTime() isBillingSetup_StartTime {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (x *BillingSetup) GetStartDateTime() *wrappers.StringValue {
	if x, ok := x.GetStartTime().(*BillingSetup_StartDateTime); ok {
		return x.StartDateTime
	}
	return nil
}

func (x *BillingSetup) GetStartTimeType() enums.TimeTypeEnum_TimeType {
	if x, ok := x.GetStartTime().(*BillingSetup_StartTimeType); ok {
		return x.StartTimeType
	}
	return enums.TimeTypeEnum_UNSPECIFIED
}

func (m *BillingSetup) GetEndTime() isBillingSetup_EndTime {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (x *BillingSetup) GetEndDateTime() *wrappers.StringValue {
	if x, ok := x.GetEndTime().(*BillingSetup_EndDateTime); ok {
		return x.EndDateTime
	}
	return nil
}

func (x *BillingSetup) GetEndTimeType() enums.TimeTypeEnum_TimeType {
	if x, ok := x.GetEndTime().(*BillingSetup_EndTimeType); ok {
		return x.EndTimeType
	}
	return enums.TimeTypeEnum_UNSPECIFIED
}

type isBillingSetup_StartTime interface {
	isBillingSetup_StartTime()
}

type BillingSetup_StartDateTime struct {
	// Immutable. The start date time in yyyy-MM-dd or yyyy-MM-dd HH:mm:ss format. Only a
	// future time is allowed.
	StartDateTime *wrappers.StringValue `protobuf:"bytes,9,opt,name=start_date_time,json=startDateTime,proto3,oneof"`
}

type BillingSetup_StartTimeType struct {
	// Immutable. The start time as a type. Only NOW is allowed.
	StartTimeType enums.TimeTypeEnum_TimeType `protobuf:"varint,10,opt,name=start_time_type,json=startTimeType,proto3,enum=google.ads.googleads.v2.enums.TimeTypeEnum_TimeType,oneof"`
}

func (*BillingSetup_StartDateTime) isBillingSetup_StartTime() {}

func (*BillingSetup_StartTimeType) isBillingSetup_StartTime() {}

type isBillingSetup_EndTime interface {
	isBillingSetup_EndTime()
}

type BillingSetup_EndDateTime struct {
	// Output only. The end date time in yyyy-MM-dd or yyyy-MM-dd HH:mm:ss format.
	EndDateTime *wrappers.StringValue `protobuf:"bytes,13,opt,name=end_date_time,json=endDateTime,proto3,oneof"`
}

type BillingSetup_EndTimeType struct {
	// Output only. The end time as a type.  The only possible value is FOREVER.
	EndTimeType enums.TimeTypeEnum_TimeType `protobuf:"varint,14,opt,name=end_time_type,json=endTimeType,proto3,enum=google.ads.googleads.v2.enums.TimeTypeEnum_TimeType,oneof"`
}

func (*BillingSetup_EndDateTime) isBillingSetup_EndTime() {}

func (*BillingSetup_EndTimeType) isBillingSetup_EndTime() {}

// Container of payments account information for this billing.
type BillingSetup_PaymentsAccountInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. A 16 digit id used to identify the payments account associated with the
	// billing setup.
	//
	// This must be passed as a string with dashes, e.g. "1234-5678-9012-3456".
	PaymentsAccountId *wrappers.StringValue `protobuf:"bytes,1,opt,name=payments_account_id,json=paymentsAccountId,proto3" json:"payments_account_id,omitempty"`
	// Immutable. The name of the payments account associated with the billing setup.
	//
	// This enables the user to specify a meaningful name for a payments account
	// to aid in reconciling monthly invoices.
	//
	// This name will be printed in the monthly invoices.
	PaymentsAccountName *wrappers.StringValue `protobuf:"bytes,2,opt,name=payments_account_name,json=paymentsAccountName,proto3" json:"payments_account_name,omitempty"`
	// Immutable. A 12 digit id used to identify the payments profile associated with the
	// billing setup.
	//
	// This must be passed in as a string with dashes, e.g. "1234-5678-9012".
	PaymentsProfileId *wrappers.StringValue `protobuf:"bytes,3,opt,name=payments_profile_id,json=paymentsProfileId,proto3" json:"payments_profile_id,omitempty"`
	// Output only. The name of the payments profile associated with the billing setup.
	PaymentsProfileName *wrappers.StringValue `protobuf:"bytes,4,opt,name=payments_profile_name,json=paymentsProfileName,proto3" json:"payments_profile_name,omitempty"`
	// Output only. A secondary payments profile id present in uncommon situations, e.g.
	// when a sequential liability agreement has been arranged.
	SecondaryPaymentsProfileId *wrappers.StringValue `protobuf:"bytes,5,opt,name=secondary_payments_profile_id,json=secondaryPaymentsProfileId,proto3" json:"secondary_payments_profile_id,omitempty"`
}

func (x *BillingSetup_PaymentsAccountInfo) Reset() {
	*x = BillingSetup_PaymentsAccountInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v2_resources_billing_setup_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BillingSetup_PaymentsAccountInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillingSetup_PaymentsAccountInfo) ProtoMessage() {}

func (x *BillingSetup_PaymentsAccountInfo) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v2_resources_billing_setup_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillingSetup_PaymentsAccountInfo.ProtoReflect.Descriptor instead.
func (*BillingSetup_PaymentsAccountInfo) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v2_resources_billing_setup_proto_rawDescGZIP(), []int{0, 0}
}

func (x *BillingSetup_PaymentsAccountInfo) GetPaymentsAccountId() *wrappers.StringValue {
	if x != nil {
		return x.PaymentsAccountId
	}
	return nil
}

func (x *BillingSetup_PaymentsAccountInfo) GetPaymentsAccountName() *wrappers.StringValue {
	if x != nil {
		return x.PaymentsAccountName
	}
	return nil
}

func (x *BillingSetup_PaymentsAccountInfo) GetPaymentsProfileId() *wrappers.StringValue {
	if x != nil {
		return x.PaymentsProfileId
	}
	return nil
}

func (x *BillingSetup_PaymentsAccountInfo) GetPaymentsProfileName() *wrappers.StringValue {
	if x != nil {
		return x.PaymentsProfileName
	}
	return nil
}

func (x *BillingSetup_PaymentsAccountInfo) GetSecondaryPaymentsProfileId() *wrappers.StringValue {
	if x != nil {
		return x.SecondaryPaymentsProfileId
	}
	return nil
}

var File_google_ads_googleads_v2_resources_billing_setup_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v2_resources_billing_setup_proto_rawDesc = []byte{
	0x0a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x74, 0x75,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x38, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2f, 0x76, 0x32, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x5f, 0x73, 0x65, 0x74, 0x75, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x0b,
	0x0a, 0x0c, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x75, 0x70, 0x12, 0x52,
	0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2d, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x27, 0x0a, 0x25, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x53,
	0x65, 0x74, 0x75, 0x70, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x30, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x65, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x48, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x75,
	0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x42, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x79, 0x0a, 0x10, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x30, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x2a, 0x0a, 0x28, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x7c, 0x0a, 0x15, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x53, 0x65, 0x74, 0x75, 0x70, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52,
	0x13, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x4b, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x05,
	0x48, 0x00, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x63, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x32, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x47, 0x0a, 0x0d, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x48, 0x01, 0x52, 0x0b, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x5f, 0x0a, 0x0d, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x45,
	0x6e, 0x75, 0x6d, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x48, 0x01, 0x52, 0x0b, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x1a, 0xcf, 0x03, 0x0a, 0x13, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x51, 0x0a, 0x13, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x11, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x55, 0x0a, 0x15, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x13, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x51, 0x0a, 0x13, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0,
	0x41, 0x05, 0x52, 0x11, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x55, 0x0a, 0x15, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x13, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x64, 0x0a, 0x1d,
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x1a, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72,
	0x79, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x49, 0x64, 0x3a, 0x5e, 0xea, 0x41, 0x5b, 0x0a, 0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x75, 0x70, 0x12, 0x32,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x7d, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x75,
	0x70, 0x73, 0x2f, 0x7b, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x74, 0x75,
	0x70, 0x7d, 0x42, 0x0c, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x42, 0x0a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0xfe, 0x01, 0x0a,
	0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x11, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x53,
	0x65, 0x74, 0x75, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x21,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x32, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x32, 0x5c, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a,
	0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x56, 0x32, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v2_resources_billing_setup_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v2_resources_billing_setup_proto_rawDescData = file_google_ads_googleads_v2_resources_billing_setup_proto_rawDesc
)

func file_google_ads_googleads_v2_resources_billing_setup_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v2_resources_billing_setup_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v2_resources_billing_setup_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v2_resources_billing_setup_proto_rawDescData)
	})
	return file_google_ads_googleads_v2_resources_billing_setup_proto_rawDescData
}

var file_google_ads_googleads_v2_resources_billing_setup_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_ads_googleads_v2_resources_billing_setup_proto_goTypes = []interface{}{
	(*BillingSetup)(nil),                                 // 0: google.ads.googleads.v2.resources.BillingSetup
	(*BillingSetup_PaymentsAccountInfo)(nil),             // 1: google.ads.googleads.v2.resources.BillingSetup.PaymentsAccountInfo
	(*wrappers.Int64Value)(nil),                          // 2: google.protobuf.Int64Value
	(enums.BillingSetupStatusEnum_BillingSetupStatus)(0), // 3: google.ads.googleads.v2.enums.BillingSetupStatusEnum.BillingSetupStatus
	(*wrappers.StringValue)(nil),                         // 4: google.protobuf.StringValue
	(enums.TimeTypeEnum_TimeType)(0),                     // 5: google.ads.googleads.v2.enums.TimeTypeEnum.TimeType
}
var file_google_ads_googleads_v2_resources_billing_setup_proto_depIdxs = []int32{
	2,  // 0: google.ads.googleads.v2.resources.BillingSetup.id:type_name -> google.protobuf.Int64Value
	3,  // 1: google.ads.googleads.v2.resources.BillingSetup.status:type_name -> google.ads.googleads.v2.enums.BillingSetupStatusEnum.BillingSetupStatus
	4,  // 2: google.ads.googleads.v2.resources.BillingSetup.payments_account:type_name -> google.protobuf.StringValue
	1,  // 3: google.ads.googleads.v2.resources.BillingSetup.payments_account_info:type_name -> google.ads.googleads.v2.resources.BillingSetup.PaymentsAccountInfo
	4,  // 4: google.ads.googleads.v2.resources.BillingSetup.start_date_time:type_name -> google.protobuf.StringValue
	5,  // 5: google.ads.googleads.v2.resources.BillingSetup.start_time_type:type_name -> google.ads.googleads.v2.enums.TimeTypeEnum.TimeType
	4,  // 6: google.ads.googleads.v2.resources.BillingSetup.end_date_time:type_name -> google.protobuf.StringValue
	5,  // 7: google.ads.googleads.v2.resources.BillingSetup.end_time_type:type_name -> google.ads.googleads.v2.enums.TimeTypeEnum.TimeType
	4,  // 8: google.ads.googleads.v2.resources.BillingSetup.PaymentsAccountInfo.payments_account_id:type_name -> google.protobuf.StringValue
	4,  // 9: google.ads.googleads.v2.resources.BillingSetup.PaymentsAccountInfo.payments_account_name:type_name -> google.protobuf.StringValue
	4,  // 10: google.ads.googleads.v2.resources.BillingSetup.PaymentsAccountInfo.payments_profile_id:type_name -> google.protobuf.StringValue
	4,  // 11: google.ads.googleads.v2.resources.BillingSetup.PaymentsAccountInfo.payments_profile_name:type_name -> google.protobuf.StringValue
	4,  // 12: google.ads.googleads.v2.resources.BillingSetup.PaymentsAccountInfo.secondary_payments_profile_id:type_name -> google.protobuf.StringValue
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v2_resources_billing_setup_proto_init() }
func file_google_ads_googleads_v2_resources_billing_setup_proto_init() {
	if File_google_ads_googleads_v2_resources_billing_setup_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v2_resources_billing_setup_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BillingSetup); i {
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
		file_google_ads_googleads_v2_resources_billing_setup_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BillingSetup_PaymentsAccountInfo); i {
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
	file_google_ads_googleads_v2_resources_billing_setup_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*BillingSetup_StartDateTime)(nil),
		(*BillingSetup_StartTimeType)(nil),
		(*BillingSetup_EndDateTime)(nil),
		(*BillingSetup_EndTimeType)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v2_resources_billing_setup_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v2_resources_billing_setup_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v2_resources_billing_setup_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v2_resources_billing_setup_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v2_resources_billing_setup_proto = out.File
	file_google_ads_googleads_v2_resources_billing_setup_proto_rawDesc = nil
	file_google_ads_googleads_v2_resources_billing_setup_proto_goTypes = nil
	file_google_ads_googleads_v2_resources_billing_setup_proto_depIdxs = nil
}
