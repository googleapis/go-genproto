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
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: google/actions/sdk/v2/action.proto

package sdk

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Represents the list of Actions defined in a project.
type Actions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Map from intents to custom Actions to configure invocation for the project.
	// The invocation intents could either be system or custom intents defined
	// in the "custom/intents/" package. All intents defined here (system
	// intents & custom intents) must have a corresponding intent file in the
	// "custom/global/" package.
	Custom map[string]*Actions_CustomAction `protobuf:"bytes,3,rep,name=custom,proto3" json:"custom,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Actions) Reset() {
	*x = Actions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_action_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Actions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Actions) ProtoMessage() {}

func (x *Actions) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_action_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Actions.ProtoReflect.Descriptor instead.
func (*Actions) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_action_proto_rawDescGZIP(), []int{0}
}

func (x *Actions) GetCustom() map[string]*Actions_CustomAction {
	if x != nil {
		return x.Custom
	}
	return nil
}

// Defines the engagement mechanisms associated with this action. This
// allows end users to subscribe to push notification and daily update.
type Actions_Engagement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The title of the engagement that will be sent to end users asking for
	// their permission to receive updates. The prompt sent to end users for
	// daily updates will look like "What time would you like me to send your
	// daily {title}" and for push notifications will look like
	// "Is it ok if I send push notifications for {title}".
	// **This field is localizable.**
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// Push notification settings that this engagement supports.
	PushNotification *Actions_Engagement_PushNotification `protobuf:"bytes,2,opt,name=push_notification,json=pushNotification,proto3" json:"push_notification,omitempty"`
	// Recurring update settings that this engagement supports.
	//
	// Types that are assignable to RecurringUpdate:
	//	*Actions_Engagement_DailyUpdate_
	RecurringUpdate isActions_Engagement_RecurringUpdate `protobuf_oneof:"recurring_update"`
	// Link config for an action which determines whether sharing links is
	// enabled for the action and if so, contains the user friendly display name
	// for the link.
	// ActionLink is deprecated. Use AssistantLink instead.
	//
	// Deprecated: Do not use.
	ActionLink *Actions_Engagement_ActionLink `protobuf:"bytes,4,opt,name=action_link,json=actionLink,proto3" json:"action_link,omitempty"`
	// Link config for an action which determines whether sharing links is
	// enabled for the action and if so, contains the user friendly display name
	// for the link.
	AssistantLink *Actions_Engagement_AssistantLink `protobuf:"bytes,6,opt,name=assistant_link,json=assistantLink,proto3" json:"assistant_link,omitempty"`
}

func (x *Actions_Engagement) Reset() {
	*x = Actions_Engagement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_action_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Actions_Engagement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Actions_Engagement) ProtoMessage() {}

func (x *Actions_Engagement) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_action_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Actions_Engagement.ProtoReflect.Descriptor instead.
func (*Actions_Engagement) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_action_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Actions_Engagement) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Actions_Engagement) GetPushNotification() *Actions_Engagement_PushNotification {
	if x != nil {
		return x.PushNotification
	}
	return nil
}

func (m *Actions_Engagement) GetRecurringUpdate() isActions_Engagement_RecurringUpdate {
	if m != nil {
		return m.RecurringUpdate
	}
	return nil
}

func (x *Actions_Engagement) GetDailyUpdate() *Actions_Engagement_DailyUpdate {
	if x, ok := x.GetRecurringUpdate().(*Actions_Engagement_DailyUpdate_); ok {
		return x.DailyUpdate
	}
	return nil
}

// Deprecated: Do not use.
func (x *Actions_Engagement) GetActionLink() *Actions_Engagement_ActionLink {
	if x != nil {
		return x.ActionLink
	}
	return nil
}

func (x *Actions_Engagement) GetAssistantLink() *Actions_Engagement_AssistantLink {
	if x != nil {
		return x.AssistantLink
	}
	return nil
}

type isActions_Engagement_RecurringUpdate interface {
	isActions_Engagement_RecurringUpdate()
}

type Actions_Engagement_DailyUpdate_ struct {
	// Daily update settings that this engagement supports.
	DailyUpdate *Actions_Engagement_DailyUpdate `protobuf:"bytes,3,opt,name=daily_update,json=dailyUpdate,proto3,oneof"`
}

func (*Actions_Engagement_DailyUpdate_) isActions_Engagement_RecurringUpdate() {}

// Details regarding a custom action.
type Actions_CustomAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Engagement mechanisms associated with the action to help end users
	// subscribe to push notifications and daily updates.
	// Note that the intent name specified in daily updates/push notifications
	// slot config needs to match the intent corresponding to this action for
	// end users to subscribe to these updates.
	Engagement *Actions_Engagement `protobuf:"bytes,2,opt,name=engagement,proto3" json:"engagement,omitempty"`
}

func (x *Actions_CustomAction) Reset() {
	*x = Actions_CustomAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_action_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Actions_CustomAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Actions_CustomAction) ProtoMessage() {}

func (x *Actions_CustomAction) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_action_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Actions_CustomAction.ProtoReflect.Descriptor instead.
func (*Actions_CustomAction) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_action_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Actions_CustomAction) GetEngagement() *Actions_Engagement {
	if x != nil {
		return x.Engagement
	}
	return nil
}

// Defines push notification settings that this engagement supports.
type Actions_Engagement_PushNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Actions_Engagement_PushNotification) Reset() {
	*x = Actions_Engagement_PushNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_action_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Actions_Engagement_PushNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Actions_Engagement_PushNotification) ProtoMessage() {}

func (x *Actions_Engagement_PushNotification) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_action_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Actions_Engagement_PushNotification.ProtoReflect.Descriptor instead.
func (*Actions_Engagement_PushNotification) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_action_proto_rawDescGZIP(), []int{0, 0, 0}
}

// Defines daily update settings that this engagement supports.
type Actions_Engagement_DailyUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Actions_Engagement_DailyUpdate) Reset() {
	*x = Actions_Engagement_DailyUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_action_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Actions_Engagement_DailyUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Actions_Engagement_DailyUpdate) ProtoMessage() {}

func (x *Actions_Engagement_DailyUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_action_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Actions_Engagement_DailyUpdate.ProtoReflect.Descriptor instead.
func (*Actions_Engagement_DailyUpdate) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_action_proto_rawDescGZIP(), []int{0, 0, 1}
}

// Indicates whether sharing links is enabled for this action and the
// corresponding settings. Action links are used to deep link a user into a
// specific action.
// ActionLink is deprecated. Use AssistantLink instead.
//
// Deprecated: Do not use.
type Actions_Engagement_ActionLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User friendly display title for the link.
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *Actions_Engagement_ActionLink) Reset() {
	*x = Actions_Engagement_ActionLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_action_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Actions_Engagement_ActionLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Actions_Engagement_ActionLink) ProtoMessage() {}

func (x *Actions_Engagement_ActionLink) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_action_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Actions_Engagement_ActionLink.ProtoReflect.Descriptor instead.
func (*Actions_Engagement_ActionLink) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_action_proto_rawDescGZIP(), []int{0, 0, 2}
}

func (x *Actions_Engagement_ActionLink) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

// Indicates whether sharing links is enabled for this action and the
// corresponding settings. Assistant links are used to deep link a user into
// a specific action.
type Actions_Engagement_AssistantLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User friendly display title for the link.
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *Actions_Engagement_AssistantLink) Reset() {
	*x = Actions_Engagement_AssistantLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_action_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Actions_Engagement_AssistantLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Actions_Engagement_AssistantLink) ProtoMessage() {}

func (x *Actions_Engagement_AssistantLink) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_action_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Actions_Engagement_AssistantLink.ProtoReflect.Descriptor instead.
func (*Actions_Engagement_AssistantLink) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_action_proto_rawDescGZIP(), []int{0, 0, 3}
}

func (x *Actions_Engagement_AssistantLink) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

var File_google_actions_sdk_v2_action_proto protoreflect.FileDescriptor

var file_google_actions_sdk_v2_action_proto_rawDesc = []byte{
	0x0a, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x22, 0xbb, 0x06, 0x0a, 0x07,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x42, 0x0a, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x1a, 0xa8, 0x04, 0x0a, 0x0a,
	0x45, 0x6e, 0x67, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x67, 0x0a, 0x11, 0x70, 0x75, 0x73, 0x68, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b,
	0x2e, 0x76, 0x32, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x6e, 0x67, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x70, 0x75, 0x73, 0x68, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5a, 0x0a, 0x0c, 0x64, 0x61, 0x69,
	0x6c, 0x79, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x45, 0x6e, 0x67, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x61, 0x69, 0x6c, 0x79,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x59, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e,
	0x76, 0x32, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x6e, 0x67, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x6b,
	0x42, 0x02, 0x18, 0x01, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x6b,
	0x12, 0x5e, 0x0a, 0x0e, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x6c, 0x69,
	0x6e, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x6e, 0x67, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x4c, 0x69, 0x6e,
	0x6b, 0x52, 0x0d, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b,
	0x1a, 0x12, 0x0a, 0x10, 0x50, 0x75, 0x73, 0x68, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0d, 0x0a, 0x0b, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x1a, 0x26, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6e,
	0x6b, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x3a, 0x02, 0x18, 0x01, 0x1a, 0x25, 0x0a, 0x0d, 0x41,
	0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x42, 0x12, 0x0a, 0x10, 0x72, 0x65, 0x63, 0x75, 0x72, 0x72, 0x69, 0x6e, 0x67, 0x5f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x59, 0x0a, 0x0c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x0a, 0x65, 0x6e, 0x67, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e,
	0x76, 0x32, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x6e, 0x67, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x65, 0x6e, 0x67, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x1a, 0x66, 0x0a, 0x0b, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x41, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x64, 0x0a, 0x19, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x42, 0x0b, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x3b, 0x73, 0x64, 0x6b, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_actions_sdk_v2_action_proto_rawDescOnce sync.Once
	file_google_actions_sdk_v2_action_proto_rawDescData = file_google_actions_sdk_v2_action_proto_rawDesc
)

func file_google_actions_sdk_v2_action_proto_rawDescGZIP() []byte {
	file_google_actions_sdk_v2_action_proto_rawDescOnce.Do(func() {
		file_google_actions_sdk_v2_action_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_actions_sdk_v2_action_proto_rawDescData)
	})
	return file_google_actions_sdk_v2_action_proto_rawDescData
}

var file_google_actions_sdk_v2_action_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_google_actions_sdk_v2_action_proto_goTypes = []interface{}{
	(*Actions)(nil),              // 0: google.actions.sdk.v2.Actions
	(*Actions_Engagement)(nil),   // 1: google.actions.sdk.v2.Actions.Engagement
	(*Actions_CustomAction)(nil), // 2: google.actions.sdk.v2.Actions.CustomAction
	nil,                          // 3: google.actions.sdk.v2.Actions.CustomEntry
	(*Actions_Engagement_PushNotification)(nil), // 4: google.actions.sdk.v2.Actions.Engagement.PushNotification
	(*Actions_Engagement_DailyUpdate)(nil),      // 5: google.actions.sdk.v2.Actions.Engagement.DailyUpdate
	(*Actions_Engagement_ActionLink)(nil),       // 6: google.actions.sdk.v2.Actions.Engagement.ActionLink
	(*Actions_Engagement_AssistantLink)(nil),    // 7: google.actions.sdk.v2.Actions.Engagement.AssistantLink
}
var file_google_actions_sdk_v2_action_proto_depIdxs = []int32{
	3, // 0: google.actions.sdk.v2.Actions.custom:type_name -> google.actions.sdk.v2.Actions.CustomEntry
	4, // 1: google.actions.sdk.v2.Actions.Engagement.push_notification:type_name -> google.actions.sdk.v2.Actions.Engagement.PushNotification
	5, // 2: google.actions.sdk.v2.Actions.Engagement.daily_update:type_name -> google.actions.sdk.v2.Actions.Engagement.DailyUpdate
	6, // 3: google.actions.sdk.v2.Actions.Engagement.action_link:type_name -> google.actions.sdk.v2.Actions.Engagement.ActionLink
	7, // 4: google.actions.sdk.v2.Actions.Engagement.assistant_link:type_name -> google.actions.sdk.v2.Actions.Engagement.AssistantLink
	1, // 5: google.actions.sdk.v2.Actions.CustomAction.engagement:type_name -> google.actions.sdk.v2.Actions.Engagement
	2, // 6: google.actions.sdk.v2.Actions.CustomEntry.value:type_name -> google.actions.sdk.v2.Actions.CustomAction
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_google_actions_sdk_v2_action_proto_init() }
func file_google_actions_sdk_v2_action_proto_init() {
	if File_google_actions_sdk_v2_action_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_actions_sdk_v2_action_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Actions); i {
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
		file_google_actions_sdk_v2_action_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Actions_Engagement); i {
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
		file_google_actions_sdk_v2_action_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Actions_CustomAction); i {
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
		file_google_actions_sdk_v2_action_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Actions_Engagement_PushNotification); i {
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
		file_google_actions_sdk_v2_action_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Actions_Engagement_DailyUpdate); i {
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
		file_google_actions_sdk_v2_action_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Actions_Engagement_ActionLink); i {
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
		file_google_actions_sdk_v2_action_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Actions_Engagement_AssistantLink); i {
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
	file_google_actions_sdk_v2_action_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Actions_Engagement_DailyUpdate_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_actions_sdk_v2_action_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_actions_sdk_v2_action_proto_goTypes,
		DependencyIndexes: file_google_actions_sdk_v2_action_proto_depIdxs,
		MessageInfos:      file_google_actions_sdk_v2_action_proto_msgTypes,
	}.Build()
	File_google_actions_sdk_v2_action_proto = out.File
	file_google_actions_sdk_v2_action_proto_rawDesc = nil
	file_google_actions_sdk_v2_action_proto_goTypes = nil
	file_google_actions_sdk_v2_action_proto_depIdxs = nil
}
