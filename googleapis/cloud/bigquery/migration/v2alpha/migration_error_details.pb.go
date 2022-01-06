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
// source: google/cloud/bigquery/migration/v2alpha/migration_error_details.proto

package migration

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	errdetails "google.golang.org/genproto/googleapis/rpc/errdetails"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Provides details for errors and the corresponding resources.
type ResourceErrorDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Information about the resource where the error is located.
	ResourceInfo *errdetails.ResourceInfo `protobuf:"bytes,1,opt,name=resource_info,json=resourceInfo,proto3" json:"resource_info,omitempty"`
	// Required. The error details for the resource.
	ErrorDetails []*ErrorDetail `protobuf:"bytes,2,rep,name=error_details,json=errorDetails,proto3" json:"error_details,omitempty"`
	// Required. How many errors there are in total for the resource. Truncation can be
	// indicated by having an `error_count` that is higher than the size of
	// `error_details`.
	ErrorCount int32 `protobuf:"varint,3,opt,name=error_count,json=errorCount,proto3" json:"error_count,omitempty"`
}

func (x *ResourceErrorDetail) Reset() {
	*x = ResourceErrorDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_bigquery_migration_v2alpha_migration_error_details_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceErrorDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceErrorDetail) ProtoMessage() {}

func (x *ResourceErrorDetail) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_bigquery_migration_v2alpha_migration_error_details_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceErrorDetail.ProtoReflect.Descriptor instead.
func (*ResourceErrorDetail) Descriptor() ([]byte, []int) {
	return file_google_cloud_bigquery_migration_v2alpha_migration_error_details_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceErrorDetail) GetResourceInfo() *errdetails.ResourceInfo {
	if x != nil {
		return x.ResourceInfo
	}
	return nil
}

func (x *ResourceErrorDetail) GetErrorDetails() []*ErrorDetail {
	if x != nil {
		return x.ErrorDetails
	}
	return nil
}

func (x *ResourceErrorDetail) GetErrorCount() int32 {
	if x != nil {
		return x.ErrorCount
	}
	return 0
}

// Provides details for errors, e.g. issues that where encountered when
// processing a subtask.
type ErrorDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The exact location within the resource (if applicable).
	Location *ErrorLocation `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	// Required. Describes the cause of the error with structured detail.
	ErrorInfo *errdetails.ErrorInfo `protobuf:"bytes,2,opt,name=error_info,json=errorInfo,proto3" json:"error_info,omitempty"`
}

func (x *ErrorDetail) Reset() {
	*x = ErrorDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_bigquery_migration_v2alpha_migration_error_details_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorDetail) ProtoMessage() {}

func (x *ErrorDetail) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_bigquery_migration_v2alpha_migration_error_details_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorDetail.ProtoReflect.Descriptor instead.
func (*ErrorDetail) Descriptor() ([]byte, []int) {
	return file_google_cloud_bigquery_migration_v2alpha_migration_error_details_proto_rawDescGZIP(), []int{1}
}

func (x *ErrorDetail) GetLocation() *ErrorLocation {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *ErrorDetail) GetErrorInfo() *errdetails.ErrorInfo {
	if x != nil {
		return x.ErrorInfo
	}
	return nil
}

// Holds information about where the error is located.
type ErrorLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. If applicable, denotes the line where the error occurred. A zero value
	// means that there is no line information.
	Line int32 `protobuf:"varint,1,opt,name=line,proto3" json:"line,omitempty"`
	// Optional. If applicable, denotes the column where the error occurred. A zero value
	// means that there is no columns information.
	Column int32 `protobuf:"varint,2,opt,name=column,proto3" json:"column,omitempty"`
}

func (x *ErrorLocation) Reset() {
	*x = ErrorLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_bigquery_migration_v2alpha_migration_error_details_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorLocation) ProtoMessage() {}

func (x *ErrorLocation) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_bigquery_migration_v2alpha_migration_error_details_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorLocation.ProtoReflect.Descriptor instead.
func (*ErrorLocation) Descriptor() ([]byte, []int) {
	return file_google_cloud_bigquery_migration_v2alpha_migration_error_details_proto_rawDescGZIP(), []int{2}
}

func (x *ErrorLocation) GetLine() int32 {
	if x != nil {
		return x.Line
	}
	return 0
}

func (x *ErrorLocation) GetColumn() int32 {
	if x != nil {
		return x.Column
	}
	return 0
}

var File_google_cloud_bigquery_migration_v2alpha_migration_error_details_proto protoreflect.FileDescriptor

var file_google_cloud_bigquery_migration_v2alpha_migration_error_details_proto_rawDesc = []byte{
	0x0a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62,
	0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6d,
	0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xdf, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x42, 0x0a, 0x0d, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x5e, 0x0a,
	0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6d, 0x69, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x24, 0x0a,
	0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0xa1, 0x01, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x57, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6d, 0x69,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0,
	0x41, 0x01, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x45, 0x0a, 0x0d, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x04, 0x6c, 0x69, 0x6e,
	0x65, 0x12, 0x1b, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x42, 0xf1,
	0x01, 0x0a, 0x2b, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6d, 0x69, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x1a,
	0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x3b, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0xaa, 0x02,
	0x27, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x42, 0x69,
	0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x56, 0x32, 0x41, 0x6c, 0x70, 0x68, 0x61, 0xca, 0x02, 0x27, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x5c, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_bigquery_migration_v2alpha_migration_error_details_proto_rawDescOnce sync.Once
	file_google_cloud_bigquery_migration_v2alpha_migration_error_details_proto_rawDescData = file_google_cloud_bigquery_migration_v2alpha_migration_error_details_proto_rawDesc
)

func file_google_cloud_bigquery_migration_v2alpha_migration_error_details_proto_rawDescGZIP() []byte {
	file_google_cloud_bigquery_migration_v2alpha_migration_error_details_proto_rawDescOnce.Do(func() {
		file_google_cloud_bigquery_migration_v2alpha_migration_error_details_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_bigquery_migration_v2alpha_migration_error_details_proto_rawDescData)
	})
	return file_google_cloud_bigquery_migration_v2alpha_migration_error_details_proto_rawDescData
}

var file_google_cloud_bigquery_migration_v2alpha_migration_error_details_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_bigquery_migration_v2alpha_migration_error_details_proto_goTypes = []interface{}{
	(*ResourceErrorDetail)(nil),     // 0: google.cloud.bigquery.migration.v2alpha.ResourceErrorDetail
	(*ErrorDetail)(nil),             // 1: google.cloud.bigquery.migration.v2alpha.ErrorDetail
	(*ErrorLocation)(nil),           // 2: google.cloud.bigquery.migration.v2alpha.ErrorLocation
	(*errdetails.ResourceInfo)(nil), // 3: google.rpc.ResourceInfo
	(*errdetails.ErrorInfo)(nil),    // 4: google.rpc.ErrorInfo
}
var file_google_cloud_bigquery_migration_v2alpha_migration_error_details_proto_depIdxs = []int32{
	3, // 0: google.cloud.bigquery.migration.v2alpha.ResourceErrorDetail.resource_info:type_name -> google.rpc.ResourceInfo
	1, // 1: google.cloud.bigquery.migration.v2alpha.ResourceErrorDetail.error_details:type_name -> google.cloud.bigquery.migration.v2alpha.ErrorDetail
	2, // 2: google.cloud.bigquery.migration.v2alpha.ErrorDetail.location:type_name -> google.cloud.bigquery.migration.v2alpha.ErrorLocation
	4, // 3: google.cloud.bigquery.migration.v2alpha.ErrorDetail.error_info:type_name -> google.rpc.ErrorInfo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_bigquery_migration_v2alpha_migration_error_details_proto_init() }
func file_google_cloud_bigquery_migration_v2alpha_migration_error_details_proto_init() {
	if File_google_cloud_bigquery_migration_v2alpha_migration_error_details_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_bigquery_migration_v2alpha_migration_error_details_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceErrorDetail); i {
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
		file_google_cloud_bigquery_migration_v2alpha_migration_error_details_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorDetail); i {
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
		file_google_cloud_bigquery_migration_v2alpha_migration_error_details_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorLocation); i {
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
			RawDescriptor: file_google_cloud_bigquery_migration_v2alpha_migration_error_details_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_bigquery_migration_v2alpha_migration_error_details_proto_goTypes,
		DependencyIndexes: file_google_cloud_bigquery_migration_v2alpha_migration_error_details_proto_depIdxs,
		MessageInfos:      file_google_cloud_bigquery_migration_v2alpha_migration_error_details_proto_msgTypes,
	}.Build()
	File_google_cloud_bigquery_migration_v2alpha_migration_error_details_proto = out.File
	file_google_cloud_bigquery_migration_v2alpha_migration_error_details_proto_rawDesc = nil
	file_google_cloud_bigquery_migration_v2alpha_migration_error_details_proto_goTypes = nil
	file_google_cloud_bigquery_migration_v2alpha_migration_error_details_proto_depIdxs = nil
}
