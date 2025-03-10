// Copyright 2025 Google LLC
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
// 	protoc-gen-go v1.26.0
// 	protoc        v4.24.4
// source: google/apps/drive/activity/v2/actor.proto

package activity

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

// The types of system events that may trigger activity.
type SystemEvent_Type int32

const (
	// The event type is unspecified.
	SystemEvent_TYPE_UNSPECIFIED SystemEvent_Type = 0
	// The event is a consequence of a user account being deleted.
	SystemEvent_USER_DELETION SystemEvent_Type = 1
	// The event is due to the system automatically purging trash.
	SystemEvent_TRASH_AUTO_PURGE SystemEvent_Type = 2
)

// Enum value maps for SystemEvent_Type.
var (
	SystemEvent_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "USER_DELETION",
		2: "TRASH_AUTO_PURGE",
	}
	SystemEvent_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"USER_DELETION":    1,
		"TRASH_AUTO_PURGE": 2,
	}
)

func (x SystemEvent_Type) Enum() *SystemEvent_Type {
	p := new(SystemEvent_Type)
	*p = x
	return p
}

func (x SystemEvent_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SystemEvent_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_google_apps_drive_activity_v2_actor_proto_enumTypes[0].Descriptor()
}

func (SystemEvent_Type) Type() protoreflect.EnumType {
	return &file_google_apps_drive_activity_v2_actor_proto_enumTypes[0]
}

func (x SystemEvent_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SystemEvent_Type.Descriptor instead.
func (SystemEvent_Type) EnumDescriptor() ([]byte, []int) {
	return file_google_apps_drive_activity_v2_actor_proto_rawDescGZIP(), []int{4, 0}
}

// The actor of a Drive activity.
type Actor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of actor.
	//
	// Types that are assignable to Type:
	//
	//	*Actor_User
	//	*Actor_Anonymous
	//	*Actor_Impersonation
	//	*Actor_System
	//	*Actor_Administrator
	Type isActor_Type `protobuf_oneof:"type"`
}

func (x *Actor) Reset() {
	*x = Actor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_apps_drive_activity_v2_actor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Actor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Actor) ProtoMessage() {}

func (x *Actor) ProtoReflect() protoreflect.Message {
	mi := &file_google_apps_drive_activity_v2_actor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Actor.ProtoReflect.Descriptor instead.
func (*Actor) Descriptor() ([]byte, []int) {
	return file_google_apps_drive_activity_v2_actor_proto_rawDescGZIP(), []int{0}
}

func (m *Actor) GetType() isActor_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Actor) GetUser() *User {
	if x, ok := x.GetType().(*Actor_User); ok {
		return x.User
	}
	return nil
}

func (x *Actor) GetAnonymous() *AnonymousUser {
	if x, ok := x.GetType().(*Actor_Anonymous); ok {
		return x.Anonymous
	}
	return nil
}

func (x *Actor) GetImpersonation() *Impersonation {
	if x, ok := x.GetType().(*Actor_Impersonation); ok {
		return x.Impersonation
	}
	return nil
}

func (x *Actor) GetSystem() *SystemEvent {
	if x, ok := x.GetType().(*Actor_System); ok {
		return x.System
	}
	return nil
}

func (x *Actor) GetAdministrator() *Administrator {
	if x, ok := x.GetType().(*Actor_Administrator); ok {
		return x.Administrator
	}
	return nil
}

type isActor_Type interface {
	isActor_Type()
}

type Actor_User struct {
	// An end user.
	User *User `protobuf:"bytes,1,opt,name=user,proto3,oneof"`
}

type Actor_Anonymous struct {
	// An anonymous user.
	Anonymous *AnonymousUser `protobuf:"bytes,2,opt,name=anonymous,proto3,oneof"`
}

type Actor_Impersonation struct {
	// An account acting on behalf of another.
	Impersonation *Impersonation `protobuf:"bytes,3,opt,name=impersonation,proto3,oneof"`
}

type Actor_System struct {
	// A non-user actor (i.e. system triggered).
	System *SystemEvent `protobuf:"bytes,4,opt,name=system,proto3,oneof"`
}

type Actor_Administrator struct {
	// An administrator.
	Administrator *Administrator `protobuf:"bytes,5,opt,name=administrator,proto3,oneof"`
}

func (*Actor_User) isActor_Type() {}

func (*Actor_Anonymous) isActor_Type() {}

func (*Actor_Impersonation) isActor_Type() {}

func (*Actor_System) isActor_Type() {}

func (*Actor_Administrator) isActor_Type() {}

// Information about an end user.
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of user, such as known, unknown, and deleted.
	//
	// Types that are assignable to Type:
	//
	//	*User_KnownUser_
	//	*User_DeletedUser_
	//	*User_UnknownUser_
	Type isUser_Type `protobuf_oneof:"type"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_apps_drive_activity_v2_actor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_google_apps_drive_activity_v2_actor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_google_apps_drive_activity_v2_actor_proto_rawDescGZIP(), []int{1}
}

func (m *User) GetType() isUser_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *User) GetKnownUser() *User_KnownUser {
	if x, ok := x.GetType().(*User_KnownUser_); ok {
		return x.KnownUser
	}
	return nil
}

func (x *User) GetDeletedUser() *User_DeletedUser {
	if x, ok := x.GetType().(*User_DeletedUser_); ok {
		return x.DeletedUser
	}
	return nil
}

func (x *User) GetUnknownUser() *User_UnknownUser {
	if x, ok := x.GetType().(*User_UnknownUser_); ok {
		return x.UnknownUser
	}
	return nil
}

type isUser_Type interface {
	isUser_Type()
}

type User_KnownUser_ struct {
	// A known user.
	KnownUser *User_KnownUser `protobuf:"bytes,2,opt,name=known_user,json=knownUser,proto3,oneof"`
}

type User_DeletedUser_ struct {
	// A user whose account has since been deleted.
	DeletedUser *User_DeletedUser `protobuf:"bytes,3,opt,name=deleted_user,json=deletedUser,proto3,oneof"`
}

type User_UnknownUser_ struct {
	// A user about whom nothing is currently known.
	UnknownUser *User_UnknownUser `protobuf:"bytes,4,opt,name=unknown_user,json=unknownUser,proto3,oneof"`
}

func (*User_KnownUser_) isUser_Type() {}

func (*User_DeletedUser_) isUser_Type() {}

func (*User_UnknownUser_) isUser_Type() {}

// Empty message representing an anonymous user or indicating the authenticated
// user should be anonymized.
type AnonymousUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AnonymousUser) Reset() {
	*x = AnonymousUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_apps_drive_activity_v2_actor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnonymousUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnonymousUser) ProtoMessage() {}

func (x *AnonymousUser) ProtoReflect() protoreflect.Message {
	mi := &file_google_apps_drive_activity_v2_actor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnonymousUser.ProtoReflect.Descriptor instead.
func (*AnonymousUser) Descriptor() ([]byte, []int) {
	return file_google_apps_drive_activity_v2_actor_proto_rawDescGZIP(), []int{2}
}

// Information about an impersonation, where an admin acts on behalf of an end
// user. Information about the acting admin is not currently available.
type Impersonation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The impersonated user.
	ImpersonatedUser *User `protobuf:"bytes,1,opt,name=impersonated_user,json=impersonatedUser,proto3" json:"impersonated_user,omitempty"`
}

func (x *Impersonation) Reset() {
	*x = Impersonation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_apps_drive_activity_v2_actor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Impersonation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Impersonation) ProtoMessage() {}

func (x *Impersonation) ProtoReflect() protoreflect.Message {
	mi := &file_google_apps_drive_activity_v2_actor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Impersonation.ProtoReflect.Descriptor instead.
func (*Impersonation) Descriptor() ([]byte, []int) {
	return file_google_apps_drive_activity_v2_actor_proto_rawDescGZIP(), []int{3}
}

func (x *Impersonation) GetImpersonatedUser() *User {
	if x != nil {
		return x.ImpersonatedUser
	}
	return nil
}

// Event triggered by system operations instead of end users.
type SystemEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of the system event that may triggered activity.
	Type SystemEvent_Type `protobuf:"varint,1,opt,name=type,proto3,enum=google.apps.drive.activity.v2.SystemEvent_Type" json:"type,omitempty"`
}

func (x *SystemEvent) Reset() {
	*x = SystemEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_apps_drive_activity_v2_actor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemEvent) ProtoMessage() {}

func (x *SystemEvent) ProtoReflect() protoreflect.Message {
	mi := &file_google_apps_drive_activity_v2_actor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemEvent.ProtoReflect.Descriptor instead.
func (*SystemEvent) Descriptor() ([]byte, []int) {
	return file_google_apps_drive_activity_v2_actor_proto_rawDescGZIP(), []int{4}
}

func (x *SystemEvent) GetType() SystemEvent_Type {
	if x != nil {
		return x.Type
	}
	return SystemEvent_TYPE_UNSPECIFIED
}

// Empty message representing an administrator.
type Administrator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Administrator) Reset() {
	*x = Administrator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_apps_drive_activity_v2_actor_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Administrator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Administrator) ProtoMessage() {}

func (x *Administrator) ProtoReflect() protoreflect.Message {
	mi := &file_google_apps_drive_activity_v2_actor_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Administrator.ProtoReflect.Descriptor instead.
func (*Administrator) Descriptor() ([]byte, []int) {
	return file_google_apps_drive_activity_v2_actor_proto_rawDescGZIP(), []int{5}
}

// A known user.
type User_KnownUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The identifier for this user that can be used with the People API to get
	// more information. The format is `people/ACCOUNT_ID`. See
	// https://developers.google.com/people/.
	PersonName string `protobuf:"bytes,1,opt,name=person_name,json=personName,proto3" json:"person_name,omitempty"`
	// True if this is the user making the request.
	IsCurrentUser bool `protobuf:"varint,2,opt,name=is_current_user,json=isCurrentUser,proto3" json:"is_current_user,omitempty"`
}

func (x *User_KnownUser) Reset() {
	*x = User_KnownUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_apps_drive_activity_v2_actor_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User_KnownUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User_KnownUser) ProtoMessage() {}

func (x *User_KnownUser) ProtoReflect() protoreflect.Message {
	mi := &file_google_apps_drive_activity_v2_actor_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User_KnownUser.ProtoReflect.Descriptor instead.
func (*User_KnownUser) Descriptor() ([]byte, []int) {
	return file_google_apps_drive_activity_v2_actor_proto_rawDescGZIP(), []int{1, 0}
}

func (x *User_KnownUser) GetPersonName() string {
	if x != nil {
		return x.PersonName
	}
	return ""
}

func (x *User_KnownUser) GetIsCurrentUser() bool {
	if x != nil {
		return x.IsCurrentUser
	}
	return false
}

// A user whose account has since been deleted.
type User_DeletedUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *User_DeletedUser) Reset() {
	*x = User_DeletedUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_apps_drive_activity_v2_actor_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User_DeletedUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User_DeletedUser) ProtoMessage() {}

func (x *User_DeletedUser) ProtoReflect() protoreflect.Message {
	mi := &file_google_apps_drive_activity_v2_actor_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User_DeletedUser.ProtoReflect.Descriptor instead.
func (*User_DeletedUser) Descriptor() ([]byte, []int) {
	return file_google_apps_drive_activity_v2_actor_proto_rawDescGZIP(), []int{1, 1}
}

// A user about whom nothing is currently known.
type User_UnknownUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *User_UnknownUser) Reset() {
	*x = User_UnknownUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_apps_drive_activity_v2_actor_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User_UnknownUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User_UnknownUser) ProtoMessage() {}

func (x *User_UnknownUser) ProtoReflect() protoreflect.Message {
	mi := &file_google_apps_drive_activity_v2_actor_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User_UnknownUser.ProtoReflect.Descriptor instead.
func (*User_UnknownUser) Descriptor() ([]byte, []int) {
	return file_google_apps_drive_activity_v2_actor_proto_rawDescGZIP(), []int{1, 2}
}

var File_google_apps_drive_activity_v2_actor_proto protoreflect.FileDescriptor

var file_google_apps_drive_activity_v2_actor_proto_rawDesc = []byte{
	0x0a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x22, 0x8a, 0x03, 0x0a, 0x05, 0x41,
	0x63, 0x74, 0x6f, 0x72, 0x12, 0x39, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73,
	0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e,
	0x76, 0x32, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x48, 0x00, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12,
	0x4c, 0x0a, 0x09, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73,
	0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e,
	0x76, 0x32, 0x2e, 0x41, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x55, 0x73, 0x65, 0x72,
	0x48, 0x00, 0x52, 0x09, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x12, 0x54, 0x0a,
	0x0d, 0x69, 0x6d, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70,
	0x70, 0x73, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6d, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0d, 0x69, 0x6d, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70,
	0x73, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x2e, 0x76, 0x32, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48,
	0x00, 0x52, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x54, 0x0a, 0x0d, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32,
	0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x48, 0x00,
	0x52, 0x0d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x42,
	0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xfe, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x4e, 0x0a, 0x0a, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70,
	0x70, 0x73, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x55,
	0x73, 0x65, 0x72, 0x48, 0x00, 0x52, 0x09, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x54, 0x0a, 0x0c, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x70, 0x70, 0x73, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x54, 0x0a, 0x0c, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77,
	0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x48, 0x00, 0x52,
	0x0b, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x54, 0x0a, 0x09,
	0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x69, 0x73,
	0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x1a, 0x0d, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x1a, 0x0d, 0x0a, 0x0b, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x41, 0x6e, 0x6f, 0x6e,
	0x79, 0x6d, 0x6f, 0x75, 0x73, 0x55, 0x73, 0x65, 0x72, 0x22, 0x61, 0x0a, 0x0d, 0x49, 0x6d, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x50, 0x0a, 0x11, 0x69, 0x6d,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x70, 0x70, 0x73, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x10, 0x69, 0x6d, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x22, 0x99, 0x01, 0x0a,
	0x0b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x43, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x22, 0x45, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x11, 0x0a, 0x0d, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4f, 0x4e,
	0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x52, 0x41, 0x53, 0x48, 0x5f, 0x41, 0x55, 0x54, 0x4f,
	0x5f, 0x50, 0x55, 0x52, 0x47, 0x45, 0x10, 0x02, 0x22, 0x0f, 0x0a, 0x0d, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x42, 0xbf, 0x01, 0x0a, 0x21, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x42,
	0x0a, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2f,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x3b, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0xa2, 0x02, 0x04, 0x47, 0x41, 0x44, 0x41, 0xaa, 0x02, 0x1d, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x73, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x2e,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x1d, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x70, 0x70, 0x73, 0x5c, 0x44, 0x72, 0x69, 0x76, 0x65, 0x5c,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5c, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_apps_drive_activity_v2_actor_proto_rawDescOnce sync.Once
	file_google_apps_drive_activity_v2_actor_proto_rawDescData = file_google_apps_drive_activity_v2_actor_proto_rawDesc
)

func file_google_apps_drive_activity_v2_actor_proto_rawDescGZIP() []byte {
	file_google_apps_drive_activity_v2_actor_proto_rawDescOnce.Do(func() {
		file_google_apps_drive_activity_v2_actor_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_apps_drive_activity_v2_actor_proto_rawDescData)
	})
	return file_google_apps_drive_activity_v2_actor_proto_rawDescData
}

var file_google_apps_drive_activity_v2_actor_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_apps_drive_activity_v2_actor_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_google_apps_drive_activity_v2_actor_proto_goTypes = []interface{}{
	(SystemEvent_Type)(0),    // 0: google.apps.drive.activity.v2.SystemEvent.Type
	(*Actor)(nil),            // 1: google.apps.drive.activity.v2.Actor
	(*User)(nil),             // 2: google.apps.drive.activity.v2.User
	(*AnonymousUser)(nil),    // 3: google.apps.drive.activity.v2.AnonymousUser
	(*Impersonation)(nil),    // 4: google.apps.drive.activity.v2.Impersonation
	(*SystemEvent)(nil),      // 5: google.apps.drive.activity.v2.SystemEvent
	(*Administrator)(nil),    // 6: google.apps.drive.activity.v2.Administrator
	(*User_KnownUser)(nil),   // 7: google.apps.drive.activity.v2.User.KnownUser
	(*User_DeletedUser)(nil), // 8: google.apps.drive.activity.v2.User.DeletedUser
	(*User_UnknownUser)(nil), // 9: google.apps.drive.activity.v2.User.UnknownUser
}
var file_google_apps_drive_activity_v2_actor_proto_depIdxs = []int32{
	2,  // 0: google.apps.drive.activity.v2.Actor.user:type_name -> google.apps.drive.activity.v2.User
	3,  // 1: google.apps.drive.activity.v2.Actor.anonymous:type_name -> google.apps.drive.activity.v2.AnonymousUser
	4,  // 2: google.apps.drive.activity.v2.Actor.impersonation:type_name -> google.apps.drive.activity.v2.Impersonation
	5,  // 3: google.apps.drive.activity.v2.Actor.system:type_name -> google.apps.drive.activity.v2.SystemEvent
	6,  // 4: google.apps.drive.activity.v2.Actor.administrator:type_name -> google.apps.drive.activity.v2.Administrator
	7,  // 5: google.apps.drive.activity.v2.User.known_user:type_name -> google.apps.drive.activity.v2.User.KnownUser
	8,  // 6: google.apps.drive.activity.v2.User.deleted_user:type_name -> google.apps.drive.activity.v2.User.DeletedUser
	9,  // 7: google.apps.drive.activity.v2.User.unknown_user:type_name -> google.apps.drive.activity.v2.User.UnknownUser
	2,  // 8: google.apps.drive.activity.v2.Impersonation.impersonated_user:type_name -> google.apps.drive.activity.v2.User
	0,  // 9: google.apps.drive.activity.v2.SystemEvent.type:type_name -> google.apps.drive.activity.v2.SystemEvent.Type
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_google_apps_drive_activity_v2_actor_proto_init() }
func file_google_apps_drive_activity_v2_actor_proto_init() {
	if File_google_apps_drive_activity_v2_actor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_apps_drive_activity_v2_actor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Actor); i {
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
		file_google_apps_drive_activity_v2_actor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_google_apps_drive_activity_v2_actor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnonymousUser); i {
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
		file_google_apps_drive_activity_v2_actor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Impersonation); i {
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
		file_google_apps_drive_activity_v2_actor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemEvent); i {
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
		file_google_apps_drive_activity_v2_actor_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Administrator); i {
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
		file_google_apps_drive_activity_v2_actor_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User_KnownUser); i {
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
		file_google_apps_drive_activity_v2_actor_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User_DeletedUser); i {
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
		file_google_apps_drive_activity_v2_actor_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User_UnknownUser); i {
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
	file_google_apps_drive_activity_v2_actor_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Actor_User)(nil),
		(*Actor_Anonymous)(nil),
		(*Actor_Impersonation)(nil),
		(*Actor_System)(nil),
		(*Actor_Administrator)(nil),
	}
	file_google_apps_drive_activity_v2_actor_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*User_KnownUser_)(nil),
		(*User_DeletedUser_)(nil),
		(*User_UnknownUser_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_apps_drive_activity_v2_actor_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_apps_drive_activity_v2_actor_proto_goTypes,
		DependencyIndexes: file_google_apps_drive_activity_v2_actor_proto_depIdxs,
		EnumInfos:         file_google_apps_drive_activity_v2_actor_proto_enumTypes,
		MessageInfos:      file_google_apps_drive_activity_v2_actor_proto_msgTypes,
	}.Build()
	File_google_apps_drive_activity_v2_actor_proto = out.File
	file_google_apps_drive_activity_v2_actor_proto_rawDesc = nil
	file_google_apps_drive_activity_v2_actor_proto_goTypes = nil
	file_google_apps_drive_activity_v2_actor_proto_depIdxs = nil
}
