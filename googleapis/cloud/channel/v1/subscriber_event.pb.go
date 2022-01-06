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
// source: google/cloud/channel/v1/subscriber_event.proto

package channel

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

// Type of customer event.
type CustomerEvent_Type int32

const (
	// Default value. This state doesn't show unless an error occurs.
	CustomerEvent_TYPE_UNSPECIFIED CustomerEvent_Type = 0
	// Primary domain for customer was changed.
	CustomerEvent_PRIMARY_DOMAIN_CHANGED CustomerEvent_Type = 1
	// Primary domain of the customer has been verified.
	CustomerEvent_PRIMARY_DOMAIN_VERIFIED CustomerEvent_Type = 2
)

// Enum value maps for CustomerEvent_Type.
var (
	CustomerEvent_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "PRIMARY_DOMAIN_CHANGED",
		2: "PRIMARY_DOMAIN_VERIFIED",
	}
	CustomerEvent_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED":        0,
		"PRIMARY_DOMAIN_CHANGED":  1,
		"PRIMARY_DOMAIN_VERIFIED": 2,
	}
)

func (x CustomerEvent_Type) Enum() *CustomerEvent_Type {
	p := new(CustomerEvent_Type)
	*p = x
	return p
}

func (x CustomerEvent_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CustomerEvent_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_channel_v1_subscriber_event_proto_enumTypes[0].Descriptor()
}

func (CustomerEvent_Type) Type() protoreflect.EnumType {
	return &file_google_cloud_channel_v1_subscriber_event_proto_enumTypes[0]
}

func (x CustomerEvent_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CustomerEvent_Type.Descriptor instead.
func (CustomerEvent_Type) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_channel_v1_subscriber_event_proto_rawDescGZIP(), []int{0, 0}
}

// Type of entitlement event.
type EntitlementEvent_Type int32

const (
	// Default value. This state doesn't show unless an error occurs.
	EntitlementEvent_TYPE_UNSPECIFIED EntitlementEvent_Type = 0
	// A new entitlement was created.
	EntitlementEvent_CREATED EntitlementEvent_Type = 1
	// The offer type associated with an entitlement was changed.
	// This is not triggered if an entitlement converts from a commit offer to a
	// flexible offer as part of a renewal.
	EntitlementEvent_PRICE_PLAN_SWITCHED EntitlementEvent_Type = 3
	// Annual commitment for a commit plan was changed.
	EntitlementEvent_COMMITMENT_CHANGED EntitlementEvent_Type = 4
	// An annual entitlement was renewed.
	EntitlementEvent_RENEWED EntitlementEvent_Type = 5
	// Entitlement was suspended.
	EntitlementEvent_SUSPENDED EntitlementEvent_Type = 6
	// Entitlement was unsuspended.
	EntitlementEvent_ACTIVATED EntitlementEvent_Type = 7
	// Entitlement was cancelled.
	EntitlementEvent_CANCELLED EntitlementEvent_Type = 8
	// Entitlement was upgraded or downgraded (e.g. from Google Workspace
	// Business Standard to Google Workspace Business Plus).
	EntitlementEvent_SKU_CHANGED EntitlementEvent_Type = 9
	// The renewal settings of an entitlement has changed.
	EntitlementEvent_RENEWAL_SETTING_CHANGED EntitlementEvent_Type = 10
	// Paid service has started on trial entitlement.
	EntitlementEvent_PAID_SERVICE_STARTED EntitlementEvent_Type = 11
	// License was assigned to or revoked from a user.
	EntitlementEvent_LICENSE_ASSIGNMENT_CHANGED EntitlementEvent_Type = 12
	// License cap was changed for the entitlement.
	EntitlementEvent_LICENSE_CAP_CHANGED EntitlementEvent_Type = 13
)

// Enum value maps for EntitlementEvent_Type.
var (
	EntitlementEvent_Type_name = map[int32]string{
		0:  "TYPE_UNSPECIFIED",
		1:  "CREATED",
		3:  "PRICE_PLAN_SWITCHED",
		4:  "COMMITMENT_CHANGED",
		5:  "RENEWED",
		6:  "SUSPENDED",
		7:  "ACTIVATED",
		8:  "CANCELLED",
		9:  "SKU_CHANGED",
		10: "RENEWAL_SETTING_CHANGED",
		11: "PAID_SERVICE_STARTED",
		12: "LICENSE_ASSIGNMENT_CHANGED",
		13: "LICENSE_CAP_CHANGED",
	}
	EntitlementEvent_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED":           0,
		"CREATED":                    1,
		"PRICE_PLAN_SWITCHED":        3,
		"COMMITMENT_CHANGED":         4,
		"RENEWED":                    5,
		"SUSPENDED":                  6,
		"ACTIVATED":                  7,
		"CANCELLED":                  8,
		"SKU_CHANGED":                9,
		"RENEWAL_SETTING_CHANGED":    10,
		"PAID_SERVICE_STARTED":       11,
		"LICENSE_ASSIGNMENT_CHANGED": 12,
		"LICENSE_CAP_CHANGED":        13,
	}
)

func (x EntitlementEvent_Type) Enum() *EntitlementEvent_Type {
	p := new(EntitlementEvent_Type)
	*p = x
	return p
}

func (x EntitlementEvent_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EntitlementEvent_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_channel_v1_subscriber_event_proto_enumTypes[1].Descriptor()
}

func (EntitlementEvent_Type) Type() protoreflect.EnumType {
	return &file_google_cloud_channel_v1_subscriber_event_proto_enumTypes[1]
}

func (x EntitlementEvent_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EntitlementEvent_Type.Descriptor instead.
func (EntitlementEvent_Type) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_channel_v1_subscriber_event_proto_rawDescGZIP(), []int{1, 0}
}

// Represents Pub/Sub message content describing customer update.
type CustomerEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resource name of the customer.
	// Format: accounts/{account_id}/customers/{customer_id}
	Customer string `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
	// Type of event which happened on the customer.
	EventType CustomerEvent_Type `protobuf:"varint,2,opt,name=event_type,json=eventType,proto3,enum=google.cloud.channel.v1.CustomerEvent_Type" json:"event_type,omitempty"`
}

func (x *CustomerEvent) Reset() {
	*x = CustomerEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_channel_v1_subscriber_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerEvent) ProtoMessage() {}

func (x *CustomerEvent) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_channel_v1_subscriber_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerEvent.ProtoReflect.Descriptor instead.
func (*CustomerEvent) Descriptor() ([]byte, []int) {
	return file_google_cloud_channel_v1_subscriber_event_proto_rawDescGZIP(), []int{0}
}

func (x *CustomerEvent) GetCustomer() string {
	if x != nil {
		return x.Customer
	}
	return ""
}

func (x *CustomerEvent) GetEventType() CustomerEvent_Type {
	if x != nil {
		return x.EventType
	}
	return CustomerEvent_TYPE_UNSPECIFIED
}

// Represents Pub/Sub message content describing entitlement update.
type EntitlementEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resource name of an entitlement of the form:
	// accounts/{account_id}/customers/{customer_id}/entitlements/{entitlement_id}
	Entitlement string `protobuf:"bytes,1,opt,name=entitlement,proto3" json:"entitlement,omitempty"`
	// Type of event which happened on the entitlement.
	EventType EntitlementEvent_Type `protobuf:"varint,2,opt,name=event_type,json=eventType,proto3,enum=google.cloud.channel.v1.EntitlementEvent_Type" json:"event_type,omitempty"`
}

func (x *EntitlementEvent) Reset() {
	*x = EntitlementEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_channel_v1_subscriber_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntitlementEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntitlementEvent) ProtoMessage() {}

func (x *EntitlementEvent) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_channel_v1_subscriber_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntitlementEvent.ProtoReflect.Descriptor instead.
func (*EntitlementEvent) Descriptor() ([]byte, []int) {
	return file_google_cloud_channel_v1_subscriber_event_proto_rawDescGZIP(), []int{1}
}

func (x *EntitlementEvent) GetEntitlement() string {
	if x != nil {
		return x.Entitlement
	}
	return ""
}

func (x *EntitlementEvent) GetEventType() EntitlementEvent_Type {
	if x != nil {
		return x.EventType
	}
	return EntitlementEvent_TYPE_UNSPECIFIED
}

// Represents information which resellers will get as part of notification from
// Cloud Pub/Sub.
type SubscriberEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies the Pub/Sub event provided to the partners.
	// This is a required field.
	//
	// Types that are assignable to Event:
	//	*SubscriberEvent_CustomerEvent
	//	*SubscriberEvent_EntitlementEvent
	Event isSubscriberEvent_Event `protobuf_oneof:"event"`
}

func (x *SubscriberEvent) Reset() {
	*x = SubscriberEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_channel_v1_subscriber_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriberEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriberEvent) ProtoMessage() {}

func (x *SubscriberEvent) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_channel_v1_subscriber_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriberEvent.ProtoReflect.Descriptor instead.
func (*SubscriberEvent) Descriptor() ([]byte, []int) {
	return file_google_cloud_channel_v1_subscriber_event_proto_rawDescGZIP(), []int{2}
}

func (m *SubscriberEvent) GetEvent() isSubscriberEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *SubscriberEvent) GetCustomerEvent() *CustomerEvent {
	if x, ok := x.GetEvent().(*SubscriberEvent_CustomerEvent); ok {
		return x.CustomerEvent
	}
	return nil
}

func (x *SubscriberEvent) GetEntitlementEvent() *EntitlementEvent {
	if x, ok := x.GetEvent().(*SubscriberEvent_EntitlementEvent); ok {
		return x.EntitlementEvent
	}
	return nil
}

type isSubscriberEvent_Event interface {
	isSubscriberEvent_Event()
}

type SubscriberEvent_CustomerEvent struct {
	// Customer event send as part of Pub/Sub event to partners.
	CustomerEvent *CustomerEvent `protobuf:"bytes,1,opt,name=customer_event,json=customerEvent,proto3,oneof"`
}

type SubscriberEvent_EntitlementEvent struct {
	// Entitlement event send as part of Pub/Sub event to partners.
	EntitlementEvent *EntitlementEvent `protobuf:"bytes,2,opt,name=entitlement_event,json=entitlementEvent,proto3,oneof"`
}

func (*SubscriberEvent_CustomerEvent) isSubscriberEvent_Event() {}

func (*SubscriberEvent_EntitlementEvent) isSubscriberEvent_Event() {}

var File_google_cloud_channel_v1_subscriber_event_proto protoreflect.FileDescriptor

var file_google_cloud_channel_v1_subscriber_event_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x72, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xf9, 0x01, 0x0a, 0x0d, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x45, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xfa, 0x41, 0x26, 0x0a, 0x24, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x4a, 0x0a, 0x0a, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x55, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x52, 0x49, 0x4d, 0x41, 0x52, 0x59,
	0x5f, 0x44, 0x4f, 0x4d, 0x41, 0x49, 0x4e, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x52, 0x49, 0x4d, 0x41, 0x52, 0x59, 0x5f, 0x44, 0x4f, 0x4d,
	0x41, 0x49, 0x4e, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x02, 0x22, 0xc9,
	0x03, 0x0a, 0x10, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x4e, 0x0a, 0x0b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2c, 0xfa, 0x41, 0x29, 0x0a, 0x27, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x4d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x95, 0x02, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x17,
	0x0a, 0x13, 0x50, 0x52, 0x49, 0x43, 0x45, 0x5f, 0x50, 0x4c, 0x41, 0x4e, 0x5f, 0x53, 0x57, 0x49,
	0x54, 0x43, 0x48, 0x45, 0x44, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x4f, 0x4d, 0x4d, 0x49,
	0x54, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0x04, 0x12,
	0x0b, 0x0a, 0x07, 0x52, 0x45, 0x4e, 0x45, 0x57, 0x45, 0x44, 0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09,
	0x53, 0x55, 0x53, 0x50, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x06, 0x12, 0x0d, 0x0a, 0x09, 0x41,
	0x43, 0x54, 0x49, 0x56, 0x41, 0x54, 0x45, 0x44, 0x10, 0x07, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x41,
	0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x08, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x4b, 0x55,
	0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0x09, 0x12, 0x1b, 0x0a, 0x17, 0x52, 0x45,
	0x4e, 0x45, 0x57, 0x41, 0x4c, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x48,
	0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0x0a, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x41, 0x49, 0x44, 0x5f,
	0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10,
	0x0b, 0x12, 0x1e, 0x0a, 0x1a, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x5f, 0x41, 0x53, 0x53,
	0x49, 0x47, 0x4e, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10,
	0x0c, 0x12, 0x17, 0x0a, 0x13, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x5f, 0x43, 0x41, 0x50,
	0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0x0d, 0x22, 0xc5, 0x01, 0x0a, 0x0f, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x4f,
	0x0a, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00,
	0x52, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x58, 0x0a, 0x11, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x10, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x42, 0x75, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76,
	0x31, 0x42, 0x14, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x76,
	0x31, 0x3b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_cloud_channel_v1_subscriber_event_proto_rawDescOnce sync.Once
	file_google_cloud_channel_v1_subscriber_event_proto_rawDescData = file_google_cloud_channel_v1_subscriber_event_proto_rawDesc
)

func file_google_cloud_channel_v1_subscriber_event_proto_rawDescGZIP() []byte {
	file_google_cloud_channel_v1_subscriber_event_proto_rawDescOnce.Do(func() {
		file_google_cloud_channel_v1_subscriber_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_channel_v1_subscriber_event_proto_rawDescData)
	})
	return file_google_cloud_channel_v1_subscriber_event_proto_rawDescData
}

var file_google_cloud_channel_v1_subscriber_event_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_cloud_channel_v1_subscriber_event_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_channel_v1_subscriber_event_proto_goTypes = []interface{}{
	(CustomerEvent_Type)(0),    // 0: google.cloud.channel.v1.CustomerEvent.Type
	(EntitlementEvent_Type)(0), // 1: google.cloud.channel.v1.EntitlementEvent.Type
	(*CustomerEvent)(nil),      // 2: google.cloud.channel.v1.CustomerEvent
	(*EntitlementEvent)(nil),   // 3: google.cloud.channel.v1.EntitlementEvent
	(*SubscriberEvent)(nil),    // 4: google.cloud.channel.v1.SubscriberEvent
}
var file_google_cloud_channel_v1_subscriber_event_proto_depIdxs = []int32{
	0, // 0: google.cloud.channel.v1.CustomerEvent.event_type:type_name -> google.cloud.channel.v1.CustomerEvent.Type
	1, // 1: google.cloud.channel.v1.EntitlementEvent.event_type:type_name -> google.cloud.channel.v1.EntitlementEvent.Type
	2, // 2: google.cloud.channel.v1.SubscriberEvent.customer_event:type_name -> google.cloud.channel.v1.CustomerEvent
	3, // 3: google.cloud.channel.v1.SubscriberEvent.entitlement_event:type_name -> google.cloud.channel.v1.EntitlementEvent
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_channel_v1_subscriber_event_proto_init() }
func file_google_cloud_channel_v1_subscriber_event_proto_init() {
	if File_google_cloud_channel_v1_subscriber_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_channel_v1_subscriber_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerEvent); i {
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
		file_google_cloud_channel_v1_subscriber_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntitlementEvent); i {
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
		file_google_cloud_channel_v1_subscriber_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriberEvent); i {
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
	file_google_cloud_channel_v1_subscriber_event_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*SubscriberEvent_CustomerEvent)(nil),
		(*SubscriberEvent_EntitlementEvent)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_channel_v1_subscriber_event_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_channel_v1_subscriber_event_proto_goTypes,
		DependencyIndexes: file_google_cloud_channel_v1_subscriber_event_proto_depIdxs,
		EnumInfos:         file_google_cloud_channel_v1_subscriber_event_proto_enumTypes,
		MessageInfos:      file_google_cloud_channel_v1_subscriber_event_proto_msgTypes,
	}.Build()
	File_google_cloud_channel_v1_subscriber_event_proto = out.File
	file_google_cloud_channel_v1_subscriber_event_proto_rawDesc = nil
	file_google_cloud_channel_v1_subscriber_event_proto_goTypes = nil
	file_google_cloud_channel_v1_subscriber_event_proto_depIdxs = nil
}
