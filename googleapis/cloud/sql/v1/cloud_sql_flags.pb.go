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
// source: google/cloud/sql/v1/cloud_sql_flags.proto

package sql

import (
	context "context"
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SqlFlagType int32

const (
	// This is an unknown flag type.
	SqlFlagType_SQL_FLAG_TYPE_UNSPECIFIED SqlFlagType = 0
	// Boolean type flag.
	SqlFlagType_BOOLEAN SqlFlagType = 1
	// String type flag.
	SqlFlagType_STRING SqlFlagType = 2
	// Integer type flag.
	SqlFlagType_INTEGER SqlFlagType = 3
	// Flag type used for a server startup option.
	SqlFlagType_NONE SqlFlagType = 4
	// Type introduced specially for MySQL TimeZone offset. Accept a string value
	// with the format [-12:59, 13:00].
	SqlFlagType_MYSQL_TIMEZONE_OFFSET SqlFlagType = 5
	// Float type flag.
	SqlFlagType_FLOAT SqlFlagType = 6
	// Comma-separated list of the strings in a SqlFlagType enum.
	SqlFlagType_REPEATED_STRING SqlFlagType = 7
)

// Enum value maps for SqlFlagType.
var (
	SqlFlagType_name = map[int32]string{
		0: "SQL_FLAG_TYPE_UNSPECIFIED",
		1: "BOOLEAN",
		2: "STRING",
		3: "INTEGER",
		4: "NONE",
		5: "MYSQL_TIMEZONE_OFFSET",
		6: "FLOAT",
		7: "REPEATED_STRING",
	}
	SqlFlagType_value = map[string]int32{
		"SQL_FLAG_TYPE_UNSPECIFIED": 0,
		"BOOLEAN":                   1,
		"STRING":                    2,
		"INTEGER":                   3,
		"NONE":                      4,
		"MYSQL_TIMEZONE_OFFSET":     5,
		"FLOAT":                     6,
		"REPEATED_STRING":           7,
	}
)

func (x SqlFlagType) Enum() *SqlFlagType {
	p := new(SqlFlagType)
	*p = x
	return p
}

func (x SqlFlagType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SqlFlagType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_sql_v1_cloud_sql_flags_proto_enumTypes[0].Descriptor()
}

func (SqlFlagType) Type() protoreflect.EnumType {
	return &file_google_cloud_sql_v1_cloud_sql_flags_proto_enumTypes[0]
}

func (x SqlFlagType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SqlFlagType.Descriptor instead.
func (SqlFlagType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_sql_v1_cloud_sql_flags_proto_rawDescGZIP(), []int{0}
}

// Flags list request.
type SqlFlagsListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Database type and version you want to retrieve flags for. By default, this
	// method returns flags for all database types and versions.
	DatabaseVersion string `protobuf:"bytes,1,opt,name=database_version,json=databaseVersion,proto3" json:"database_version,omitempty"`
}

func (x *SqlFlagsListRequest) Reset() {
	*x = SqlFlagsListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_sql_v1_cloud_sql_flags_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SqlFlagsListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SqlFlagsListRequest) ProtoMessage() {}

func (x *SqlFlagsListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_sql_v1_cloud_sql_flags_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SqlFlagsListRequest.ProtoReflect.Descriptor instead.
func (*SqlFlagsListRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_sql_v1_cloud_sql_flags_proto_rawDescGZIP(), []int{0}
}

func (x *SqlFlagsListRequest) GetDatabaseVersion() string {
	if x != nil {
		return x.DatabaseVersion
	}
	return ""
}

// Flags list response.
type FlagsListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This is always **sql#flagsList**.
	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	// List of flags.
	Items []*Flag `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *FlagsListResponse) Reset() {
	*x = FlagsListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_sql_v1_cloud_sql_flags_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlagsListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlagsListResponse) ProtoMessage() {}

func (x *FlagsListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_sql_v1_cloud_sql_flags_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlagsListResponse.ProtoReflect.Descriptor instead.
func (*FlagsListResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_sql_v1_cloud_sql_flags_proto_rawDescGZIP(), []int{1}
}

func (x *FlagsListResponse) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *FlagsListResponse) GetItems() []*Flag {
	if x != nil {
		return x.Items
	}
	return nil
}

// A flag resource.
type Flag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This is the name of the flag. Flag names always use underscores, not
	// hyphens, for example: **max_allowed_packet**
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The type of the flag. Flags are typed to being **BOOLEAN**, **STRING**,
	// **INTEGER** or **NONE**. **NONE** is used for flags which do not take a
	// value, such as **skip_grant_tables**.
	Type SqlFlagType `protobuf:"varint,2,opt,name=type,proto3,enum=google.cloud.sql.v1.SqlFlagType" json:"type,omitempty"`
	// The database version this flag applies to. Can be **MYSQL_8_0**,
	// **MYSQL_5_6**, or **MYSQL_5_7**.
	AppliesTo []SqlDatabaseVersion `protobuf:"varint,3,rep,packed,name=applies_to,json=appliesTo,proto3,enum=google.cloud.sql.v1.SqlDatabaseVersion" json:"applies_to,omitempty"`
	// For **STRING** flags, a list of strings that the value can be set to.
	AllowedStringValues []string `protobuf:"bytes,4,rep,name=allowed_string_values,json=allowedStringValues,proto3" json:"allowed_string_values,omitempty"`
	// For **INTEGER** flags, the minimum allowed value.
	MinValue *wrapperspb.Int64Value `protobuf:"bytes,5,opt,name=min_value,json=minValue,proto3" json:"min_value,omitempty"`
	// For **INTEGER** flags, the maximum allowed value.
	MaxValue *wrapperspb.Int64Value `protobuf:"bytes,6,opt,name=max_value,json=maxValue,proto3" json:"max_value,omitempty"`
	// Indicates whether changing this flag will trigger a database restart. Only
	// applicable to Second Generation instances.
	RequiresRestart *wrapperspb.BoolValue `protobuf:"bytes,7,opt,name=requires_restart,json=requiresRestart,proto3" json:"requires_restart,omitempty"`
	// This is always **sql#flag**.
	Kind string `protobuf:"bytes,8,opt,name=kind,proto3" json:"kind,omitempty"`
	// Whether or not the flag is considered in beta.
	InBeta *wrapperspb.BoolValue `protobuf:"bytes,9,opt,name=in_beta,json=inBeta,proto3" json:"in_beta,omitempty"`
	// Use this field if only certain integers are accepted. Can be combined
	// with min_value and max_value to add additional values.
	AllowedIntValues []int64 `protobuf:"varint,10,rep,packed,name=allowed_int_values,json=allowedIntValues,proto3" json:"allowed_int_values,omitempty"`
}

func (x *Flag) Reset() {
	*x = Flag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_sql_v1_cloud_sql_flags_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Flag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Flag) ProtoMessage() {}

func (x *Flag) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_sql_v1_cloud_sql_flags_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Flag.ProtoReflect.Descriptor instead.
func (*Flag) Descriptor() ([]byte, []int) {
	return file_google_cloud_sql_v1_cloud_sql_flags_proto_rawDescGZIP(), []int{2}
}

func (x *Flag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Flag) GetType() SqlFlagType {
	if x != nil {
		return x.Type
	}
	return SqlFlagType_SQL_FLAG_TYPE_UNSPECIFIED
}

func (x *Flag) GetAppliesTo() []SqlDatabaseVersion {
	if x != nil {
		return x.AppliesTo
	}
	return nil
}

func (x *Flag) GetAllowedStringValues() []string {
	if x != nil {
		return x.AllowedStringValues
	}
	return nil
}

func (x *Flag) GetMinValue() *wrapperspb.Int64Value {
	if x != nil {
		return x.MinValue
	}
	return nil
}

func (x *Flag) GetMaxValue() *wrapperspb.Int64Value {
	if x != nil {
		return x.MaxValue
	}
	return nil
}

func (x *Flag) GetRequiresRestart() *wrapperspb.BoolValue {
	if x != nil {
		return x.RequiresRestart
	}
	return nil
}

func (x *Flag) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Flag) GetInBeta() *wrapperspb.BoolValue {
	if x != nil {
		return x.InBeta
	}
	return nil
}

func (x *Flag) GetAllowedIntValues() []int64 {
	if x != nil {
		return x.AllowedIntValues
	}
	return nil
}

var File_google_cloud_sql_v1_cloud_sql_flags_proto protoreflect.FileDescriptor

var file_google_cloud_sql_v1_cloud_sql_flags_proto_rawDesc = []byte{
	0x0a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x71, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x73, 0x71, 0x6c, 0x5f,
	0x66, 0x6c, 0x61, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x71, 0x6c,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x73, 0x71, 0x6c, 0x5f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x13, 0x53, 0x71, 0x6c, 0x46, 0x6c, 0x61,
	0x67, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a,
	0x10, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x58, 0x0a, 0x11, 0x46, 0x6c, 0x61, 0x67,
	0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e,
	0x64, 0x12, 0x2f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x22, 0xfe, 0x03, 0x0a, 0x04, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x34, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x71, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x71, 0x6c, 0x46, 0x6c, 0x61, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x46, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x73,
	0x5f, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x71, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x09, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x54, 0x6f, 0x12, 0x32, 0x0a,
	0x15, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x13, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x12, 0x38, 0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x6d,
	0x61, 0x78, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x6d, 0x61, 0x78,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x45, 0x0a, 0x10, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x73, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0f, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x12, 0x33, 0x0a, 0x07, 0x69, 0x6e, 0x5f, 0x62, 0x65, 0x74, 0x61, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x69,
	0x6e, 0x42, 0x65, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x12, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64,
	0x5f, 0x69, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x10, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x2a, 0x97, 0x01, 0x0a, 0x0b, 0x53, 0x71, 0x6c, 0x46, 0x6c, 0x61, 0x67, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x53, 0x51, 0x4c, 0x5f, 0x46, 0x4c, 0x41, 0x47, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x4f, 0x4f, 0x4c, 0x45, 0x41, 0x4e, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x49,
	0x4e, 0x54, 0x45, 0x47, 0x45, 0x52, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x04, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x59, 0x53, 0x51, 0x4c, 0x5f, 0x54, 0x49, 0x4d, 0x45,
	0x5a, 0x4f, 0x4e, 0x45, 0x5f, 0x4f, 0x46, 0x46, 0x53, 0x45, 0x54, 0x10, 0x05, 0x12, 0x09, 0x0a,
	0x05, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x10, 0x06, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x45, 0x50, 0x45,
	0x41, 0x54, 0x45, 0x44, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x07, 0x32, 0xfc, 0x01,
	0x0a, 0x0f, 0x53, 0x71, 0x6c, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x6b, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x71, 0x6c, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x1a, 0x7c,
	0xca, 0x41, 0x17, 0x73, 0x71, 0x6c, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x5f, 0x68, 0x74, 0x74,
	0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2c, 0x68, 0x74, 0x74, 0x70,
	0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x73, 0x71, 0x6c, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x42, 0x67, 0x0a, 0x17,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x42, 0x12, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x71,
	0x6c, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x36, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x71, 0x6c, 0x2f, 0x76,
	0x31, 0x3b, 0x73, 0x71, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_sql_v1_cloud_sql_flags_proto_rawDescOnce sync.Once
	file_google_cloud_sql_v1_cloud_sql_flags_proto_rawDescData = file_google_cloud_sql_v1_cloud_sql_flags_proto_rawDesc
)

func file_google_cloud_sql_v1_cloud_sql_flags_proto_rawDescGZIP() []byte {
	file_google_cloud_sql_v1_cloud_sql_flags_proto_rawDescOnce.Do(func() {
		file_google_cloud_sql_v1_cloud_sql_flags_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_sql_v1_cloud_sql_flags_proto_rawDescData)
	})
	return file_google_cloud_sql_v1_cloud_sql_flags_proto_rawDescData
}

var file_google_cloud_sql_v1_cloud_sql_flags_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_sql_v1_cloud_sql_flags_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_sql_v1_cloud_sql_flags_proto_goTypes = []interface{}{
	(SqlFlagType)(0),              // 0: google.cloud.sql.v1.SqlFlagType
	(*SqlFlagsListRequest)(nil),   // 1: google.cloud.sql.v1.SqlFlagsListRequest
	(*FlagsListResponse)(nil),     // 2: google.cloud.sql.v1.FlagsListResponse
	(*Flag)(nil),                  // 3: google.cloud.sql.v1.Flag
	(SqlDatabaseVersion)(0),       // 4: google.cloud.sql.v1.SqlDatabaseVersion
	(*wrapperspb.Int64Value)(nil), // 5: google.protobuf.Int64Value
	(*wrapperspb.BoolValue)(nil),  // 6: google.protobuf.BoolValue
}
var file_google_cloud_sql_v1_cloud_sql_flags_proto_depIdxs = []int32{
	3, // 0: google.cloud.sql.v1.FlagsListResponse.items:type_name -> google.cloud.sql.v1.Flag
	0, // 1: google.cloud.sql.v1.Flag.type:type_name -> google.cloud.sql.v1.SqlFlagType
	4, // 2: google.cloud.sql.v1.Flag.applies_to:type_name -> google.cloud.sql.v1.SqlDatabaseVersion
	5, // 3: google.cloud.sql.v1.Flag.min_value:type_name -> google.protobuf.Int64Value
	5, // 4: google.cloud.sql.v1.Flag.max_value:type_name -> google.protobuf.Int64Value
	6, // 5: google.cloud.sql.v1.Flag.requires_restart:type_name -> google.protobuf.BoolValue
	6, // 6: google.cloud.sql.v1.Flag.in_beta:type_name -> google.protobuf.BoolValue
	1, // 7: google.cloud.sql.v1.SqlFlagsService.List:input_type -> google.cloud.sql.v1.SqlFlagsListRequest
	2, // 8: google.cloud.sql.v1.SqlFlagsService.List:output_type -> google.cloud.sql.v1.FlagsListResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_google_cloud_sql_v1_cloud_sql_flags_proto_init() }
func file_google_cloud_sql_v1_cloud_sql_flags_proto_init() {
	if File_google_cloud_sql_v1_cloud_sql_flags_proto != nil {
		return
	}
	file_google_cloud_sql_v1_cloud_sql_resources_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_sql_v1_cloud_sql_flags_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SqlFlagsListRequest); i {
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
		file_google_cloud_sql_v1_cloud_sql_flags_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlagsListResponse); i {
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
		file_google_cloud_sql_v1_cloud_sql_flags_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Flag); i {
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
			RawDescriptor: file_google_cloud_sql_v1_cloud_sql_flags_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_sql_v1_cloud_sql_flags_proto_goTypes,
		DependencyIndexes: file_google_cloud_sql_v1_cloud_sql_flags_proto_depIdxs,
		EnumInfos:         file_google_cloud_sql_v1_cloud_sql_flags_proto_enumTypes,
		MessageInfos:      file_google_cloud_sql_v1_cloud_sql_flags_proto_msgTypes,
	}.Build()
	File_google_cloud_sql_v1_cloud_sql_flags_proto = out.File
	file_google_cloud_sql_v1_cloud_sql_flags_proto_rawDesc = nil
	file_google_cloud_sql_v1_cloud_sql_flags_proto_goTypes = nil
	file_google_cloud_sql_v1_cloud_sql_flags_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SqlFlagsServiceClient is the client API for SqlFlagsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SqlFlagsServiceClient interface {
	// Lists all available database flags for Cloud SQL instances.
	List(ctx context.Context, in *SqlFlagsListRequest, opts ...grpc.CallOption) (*FlagsListResponse, error)
}

type sqlFlagsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSqlFlagsServiceClient(cc grpc.ClientConnInterface) SqlFlagsServiceClient {
	return &sqlFlagsServiceClient{cc}
}

func (c *sqlFlagsServiceClient) List(ctx context.Context, in *SqlFlagsListRequest, opts ...grpc.CallOption) (*FlagsListResponse, error) {
	out := new(FlagsListResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.sql.v1.SqlFlagsService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SqlFlagsServiceServer is the server API for SqlFlagsService service.
type SqlFlagsServiceServer interface {
	// Lists all available database flags for Cloud SQL instances.
	List(context.Context, *SqlFlagsListRequest) (*FlagsListResponse, error)
}

// UnimplementedSqlFlagsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSqlFlagsServiceServer struct {
}

func (*UnimplementedSqlFlagsServiceServer) List(context.Context, *SqlFlagsListRequest) (*FlagsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterSqlFlagsServiceServer(s *grpc.Server, srv SqlFlagsServiceServer) {
	s.RegisterService(&_SqlFlagsService_serviceDesc, srv)
}

func _SqlFlagsService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SqlFlagsListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SqlFlagsServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.sql.v1.SqlFlagsService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SqlFlagsServiceServer).List(ctx, req.(*SqlFlagsListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SqlFlagsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.sql.v1.SqlFlagsService",
	HandlerType: (*SqlFlagsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _SqlFlagsService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/sql/v1/cloud_sql_flags.proto",
}
