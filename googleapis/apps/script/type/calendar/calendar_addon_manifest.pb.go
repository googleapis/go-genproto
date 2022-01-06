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
// source: google/apps/script/type/calendar/calendar_addon_manifest.proto

package calendar

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	_type "google.golang.org/genproto/googleapis/apps/script/type"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// An enum defining the level of data access event triggers require.
type CalendarAddOnManifest_EventAccess int32

const (
	// Default value when nothing is set for EventAccess.
	CalendarAddOnManifest_UNSPECIFIED CalendarAddOnManifest_EventAccess = 0
	// METADATA gives event triggers the permission to access the metadata of
	// events such as event id and calendar id.
	CalendarAddOnManifest_METADATA CalendarAddOnManifest_EventAccess = 1
	// READ gives event triggers access to all provided event fields including
	// the metadata, attendees, and conference data.
	CalendarAddOnManifest_READ CalendarAddOnManifest_EventAccess = 3
	// WRITE gives event triggers access to the metadata of events and the
	// ability to perform all actions, including adding attendees and setting
	// conference data.
	CalendarAddOnManifest_WRITE CalendarAddOnManifest_EventAccess = 4
	// READ_WRITE gives event triggers access to all provided event fields
	// including the metadata, attendees, and conference data and the ability to
	// perform all actions.
	CalendarAddOnManifest_READ_WRITE CalendarAddOnManifest_EventAccess = 5
)

// Enum value maps for CalendarAddOnManifest_EventAccess.
var (
	CalendarAddOnManifest_EventAccess_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "METADATA",
		3: "READ",
		4: "WRITE",
		5: "READ_WRITE",
	}
	CalendarAddOnManifest_EventAccess_value = map[string]int32{
		"UNSPECIFIED": 0,
		"METADATA":    1,
		"READ":        3,
		"WRITE":       4,
		"READ_WRITE":  5,
	}
)

func (x CalendarAddOnManifest_EventAccess) Enum() *CalendarAddOnManifest_EventAccess {
	p := new(CalendarAddOnManifest_EventAccess)
	*p = x
	return p
}

func (x CalendarAddOnManifest_EventAccess) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CalendarAddOnManifest_EventAccess) Descriptor() protoreflect.EnumDescriptor {
	return file_google_apps_script_type_calendar_calendar_addon_manifest_proto_enumTypes[0].Descriptor()
}

func (CalendarAddOnManifest_EventAccess) Type() protoreflect.EnumType {
	return &file_google_apps_script_type_calendar_calendar_addon_manifest_proto_enumTypes[0]
}

func (x CalendarAddOnManifest_EventAccess) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CalendarAddOnManifest_EventAccess.Descriptor instead.
func (CalendarAddOnManifest_EventAccess) EnumDescriptor() ([]byte, []int) {
	return file_google_apps_script_type_calendar_calendar_addon_manifest_proto_rawDescGZIP(), []int{0, 0}
}

// Calendar add-on manifest.
type CalendarAddOnManifest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Defines an endpoint that will be executed contexts that don't
	// match a declared contextual trigger. Any cards generated by this function
	// will always be available to the user, but may be eclipsed by contextual
	// content when this add-on declares more targeted triggers.
	//
	// If present, this overrides the configuration from
	// `addOns.common.homepageTrigger`.
	HomepageTrigger *_type.HomepageExtensionPoint `protobuf:"bytes,6,opt,name=homepage_trigger,json=homepageTrigger,proto3" json:"homepage_trigger,omitempty"`
	// Defines conference solutions provided by this add-on.
	ConferenceSolution []*ConferenceSolution `protobuf:"bytes,3,rep,name=conference_solution,json=conferenceSolution,proto3" json:"conference_solution,omitempty"`
	// An endpoint to execute that creates a URL to the add-on's settings page.
	CreateSettingsUrlFunction string `protobuf:"bytes,5,opt,name=create_settings_url_function,json=createSettingsUrlFunction,proto3" json:"create_settings_url_function,omitempty"`
	// An endpoint to trigger when an event is opened (viewed/edited).
	EventOpenTrigger *CalendarExtensionPoint `protobuf:"bytes,10,opt,name=event_open_trigger,json=eventOpenTrigger,proto3" json:"event_open_trigger,omitempty"`
	// An endpoint to trigger when the open event is updated.
	EventUpdateTrigger *CalendarExtensionPoint `protobuf:"bytes,11,opt,name=event_update_trigger,json=eventUpdateTrigger,proto3" json:"event_update_trigger,omitempty"`
	// Define the level of data access when an event addon is triggered.
	CurrentEventAccess CalendarAddOnManifest_EventAccess `protobuf:"varint,12,opt,name=current_event_access,json=currentEventAccess,proto3,enum=google.apps.script.type.calendar.CalendarAddOnManifest_EventAccess" json:"current_event_access,omitempty"`
}

func (x *CalendarAddOnManifest) Reset() {
	*x = CalendarAddOnManifest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_apps_script_type_calendar_calendar_addon_manifest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalendarAddOnManifest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalendarAddOnManifest) ProtoMessage() {}

func (x *CalendarAddOnManifest) ProtoReflect() protoreflect.Message {
	mi := &file_google_apps_script_type_calendar_calendar_addon_manifest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalendarAddOnManifest.ProtoReflect.Descriptor instead.
func (*CalendarAddOnManifest) Descriptor() ([]byte, []int) {
	return file_google_apps_script_type_calendar_calendar_addon_manifest_proto_rawDescGZIP(), []int{0}
}

func (x *CalendarAddOnManifest) GetHomepageTrigger() *_type.HomepageExtensionPoint {
	if x != nil {
		return x.HomepageTrigger
	}
	return nil
}

func (x *CalendarAddOnManifest) GetConferenceSolution() []*ConferenceSolution {
	if x != nil {
		return x.ConferenceSolution
	}
	return nil
}

func (x *CalendarAddOnManifest) GetCreateSettingsUrlFunction() string {
	if x != nil {
		return x.CreateSettingsUrlFunction
	}
	return ""
}

func (x *CalendarAddOnManifest) GetEventOpenTrigger() *CalendarExtensionPoint {
	if x != nil {
		return x.EventOpenTrigger
	}
	return nil
}

func (x *CalendarAddOnManifest) GetEventUpdateTrigger() *CalendarExtensionPoint {
	if x != nil {
		return x.EventUpdateTrigger
	}
	return nil
}

func (x *CalendarAddOnManifest) GetCurrentEventAccess() CalendarAddOnManifest_EventAccess {
	if x != nil {
		return x.CurrentEventAccess
	}
	return CalendarAddOnManifest_UNSPECIFIED
}

// Defines conference related values.
type ConferenceSolution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The endpoint to call when ConferenceData should be created.
	OnCreateFunction string `protobuf:"bytes,1,opt,name=on_create_function,json=onCreateFunction,proto3" json:"on_create_function,omitempty"`
	// Required. IDs should be unique across ConferenceSolutions within one
	// add-on, but this is not strictly enforced. It is up to the add-on developer
	// to assign them uniquely, otherwise the wrong ConferenceSolution may be
	// used when the add-on is triggered. While the developer may change the
	// display name of an add-on, the ID should not be changed.
	Id string `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	// Required. The display name of the ConferenceSolution.
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The URL for the logo image of the ConferenceSolution.
	LogoUrl string `protobuf:"bytes,6,opt,name=logo_url,json=logoUrl,proto3" json:"logo_url,omitempty"`
}

func (x *ConferenceSolution) Reset() {
	*x = ConferenceSolution{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_apps_script_type_calendar_calendar_addon_manifest_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConferenceSolution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConferenceSolution) ProtoMessage() {}

func (x *ConferenceSolution) ProtoReflect() protoreflect.Message {
	mi := &file_google_apps_script_type_calendar_calendar_addon_manifest_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConferenceSolution.ProtoReflect.Descriptor instead.
func (*ConferenceSolution) Descriptor() ([]byte, []int) {
	return file_google_apps_script_type_calendar_calendar_addon_manifest_proto_rawDescGZIP(), []int{1}
}

func (x *ConferenceSolution) GetOnCreateFunction() string {
	if x != nil {
		return x.OnCreateFunction
	}
	return ""
}

func (x *ConferenceSolution) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ConferenceSolution) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ConferenceSolution) GetLogoUrl() string {
	if x != nil {
		return x.LogoUrl
	}
	return ""
}

// Common format for declaring a calendar add-on's triggers.
type CalendarExtensionPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The endpoint to execute when this extension point is
	// activated.
	RunFunction string `protobuf:"bytes,1,opt,name=run_function,json=runFunction,proto3" json:"run_function,omitempty"`
}

func (x *CalendarExtensionPoint) Reset() {
	*x = CalendarExtensionPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_apps_script_type_calendar_calendar_addon_manifest_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalendarExtensionPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalendarExtensionPoint) ProtoMessage() {}

func (x *CalendarExtensionPoint) ProtoReflect() protoreflect.Message {
	mi := &file_google_apps_script_type_calendar_calendar_addon_manifest_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalendarExtensionPoint.ProtoReflect.Descriptor instead.
func (*CalendarExtensionPoint) Descriptor() ([]byte, []int) {
	return file_google_apps_script_type_calendar_calendar_addon_manifest_proto_rawDescGZIP(), []int{2}
}

func (x *CalendarExtensionPoint) GetRunFunction() string {
	if x != nil {
		return x.RunFunction
	}
	return ""
}

var File_google_apps_script_type_calendar_calendar_addon_manifest_proto protoreflect.FileDescriptor

var file_google_apps_script_type_calendar_calendar_addon_manifest_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64,
	0x61, 0x72, 0x2f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x6f,
	0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64,
	0x61, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x73,
	0x2f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb9, 0x05, 0x0a, 0x15, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x41,
	0x64, 0x64, 0x4f, 0x6e, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x12, 0x5a, 0x0a, 0x10,
	0x68, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x70, 0x70, 0x73, 0x2e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0f, 0x68, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67,
	0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x65, 0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x70, 0x70, 0x73, 0x2e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x63, 0x6f, 0x6e,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x3f, 0x0a, 0x1c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x5f, 0x75, 0x72, 0x6c, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x19, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x55, 0x72, 0x6c, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x66, 0x0a, 0x12, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x74,
	0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e,
	0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x10, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4f, 0x70, 0x65,
	0x6e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x6a, 0x0a, 0x14, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x70, 0x70, 0x73, 0x2e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64,
	0x61, 0x72, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x12, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69,
	0x67, 0x67, 0x65, 0x72, 0x12, 0x75, 0x0a, 0x14, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x43, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73,
	0x2e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x63, 0x61, 0x6c,
	0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x41, 0x64,
	0x64, 0x4f, 0x6e, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x12, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x51, 0x0a, 0x0b, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4d,
	0x45, 0x54, 0x41, 0x44, 0x41, 0x54, 0x41, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x45, 0x41,
	0x44, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x57, 0x52, 0x49, 0x54, 0x45, 0x10, 0x04, 0x12, 0x0e,
	0x0a, 0x0a, 0x52, 0x45, 0x41, 0x44, 0x5f, 0x57, 0x52, 0x49, 0x54, 0x45, 0x10, 0x05, 0x22, 0x95,
	0x01, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x12, 0x6f, 0x6e, 0x5f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x10, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x6f, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x07, 0x6c,
	0x6f, 0x67, 0x6f, 0x55, 0x72, 0x6c, 0x22, 0x40, 0x0a, 0x16, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64,
	0x61, 0x72, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x26, 0x0a, 0x0c, 0x72, 0x75, 0x6e, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x72, 0x75, 0x6e,
	0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0xf2, 0x01, 0x0a, 0x24, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61,
	0x72, 0x42, 0x1a, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x41, 0x64, 0x64, 0x4f, 0x6e,
	0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x3f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72,
	0xaa, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x73, 0x2e, 0x53,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x61, 0x6c, 0x65, 0x6e,
	0x64, 0x61, 0x72, 0xca, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x70, 0x70,
	0x73, 0x5c, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5c, 0x54, 0x79, 0x70, 0x65, 0x5c, 0x43, 0x61,
	0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0xea, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a,
	0x3a, 0x41, 0x70, 0x70, 0x73, 0x3a, 0x3a, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x3a, 0x3a, 0x54,
	0x79, 0x70, 0x65, 0x3a, 0x3a, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_apps_script_type_calendar_calendar_addon_manifest_proto_rawDescOnce sync.Once
	file_google_apps_script_type_calendar_calendar_addon_manifest_proto_rawDescData = file_google_apps_script_type_calendar_calendar_addon_manifest_proto_rawDesc
)

func file_google_apps_script_type_calendar_calendar_addon_manifest_proto_rawDescGZIP() []byte {
	file_google_apps_script_type_calendar_calendar_addon_manifest_proto_rawDescOnce.Do(func() {
		file_google_apps_script_type_calendar_calendar_addon_manifest_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_apps_script_type_calendar_calendar_addon_manifest_proto_rawDescData)
	})
	return file_google_apps_script_type_calendar_calendar_addon_manifest_proto_rawDescData
}

var file_google_apps_script_type_calendar_calendar_addon_manifest_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_apps_script_type_calendar_calendar_addon_manifest_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_apps_script_type_calendar_calendar_addon_manifest_proto_goTypes = []interface{}{
	(CalendarAddOnManifest_EventAccess)(0), // 0: google.apps.script.type.calendar.CalendarAddOnManifest.EventAccess
	(*CalendarAddOnManifest)(nil),          // 1: google.apps.script.type.calendar.CalendarAddOnManifest
	(*ConferenceSolution)(nil),             // 2: google.apps.script.type.calendar.ConferenceSolution
	(*CalendarExtensionPoint)(nil),         // 3: google.apps.script.type.calendar.CalendarExtensionPoint
	(*_type.HomepageExtensionPoint)(nil),   // 4: google.apps.script.type.HomepageExtensionPoint
}
var file_google_apps_script_type_calendar_calendar_addon_manifest_proto_depIdxs = []int32{
	4, // 0: google.apps.script.type.calendar.CalendarAddOnManifest.homepage_trigger:type_name -> google.apps.script.type.HomepageExtensionPoint
	2, // 1: google.apps.script.type.calendar.CalendarAddOnManifest.conference_solution:type_name -> google.apps.script.type.calendar.ConferenceSolution
	3, // 2: google.apps.script.type.calendar.CalendarAddOnManifest.event_open_trigger:type_name -> google.apps.script.type.calendar.CalendarExtensionPoint
	3, // 3: google.apps.script.type.calendar.CalendarAddOnManifest.event_update_trigger:type_name -> google.apps.script.type.calendar.CalendarExtensionPoint
	0, // 4: google.apps.script.type.calendar.CalendarAddOnManifest.current_event_access:type_name -> google.apps.script.type.calendar.CalendarAddOnManifest.EventAccess
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_apps_script_type_calendar_calendar_addon_manifest_proto_init() }
func file_google_apps_script_type_calendar_calendar_addon_manifest_proto_init() {
	if File_google_apps_script_type_calendar_calendar_addon_manifest_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_apps_script_type_calendar_calendar_addon_manifest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalendarAddOnManifest); i {
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
		file_google_apps_script_type_calendar_calendar_addon_manifest_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConferenceSolution); i {
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
		file_google_apps_script_type_calendar_calendar_addon_manifest_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalendarExtensionPoint); i {
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
			RawDescriptor: file_google_apps_script_type_calendar_calendar_addon_manifest_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_apps_script_type_calendar_calendar_addon_manifest_proto_goTypes,
		DependencyIndexes: file_google_apps_script_type_calendar_calendar_addon_manifest_proto_depIdxs,
		EnumInfos:         file_google_apps_script_type_calendar_calendar_addon_manifest_proto_enumTypes,
		MessageInfos:      file_google_apps_script_type_calendar_calendar_addon_manifest_proto_msgTypes,
	}.Build()
	File_google_apps_script_type_calendar_calendar_addon_manifest_proto = out.File
	file_google_apps_script_type_calendar_calendar_addon_manifest_proto_rawDesc = nil
	file_google_apps_script_type_calendar_calendar_addon_manifest_proto_goTypes = nil
	file_google_apps_script_type_calendar_calendar_addon_manifest_proto_depIdxs = nil
}
