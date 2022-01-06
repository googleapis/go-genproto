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
// source: google/cloud/sql/v1/cloud_sql_operations.proto

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
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Operations get request.
type SqlOperationsGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Instance operation ID.
	Operation string `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	// Project ID of the project that contains the instance.
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *SqlOperationsGetRequest) Reset() {
	*x = SqlOperationsGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_sql_v1_cloud_sql_operations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SqlOperationsGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SqlOperationsGetRequest) ProtoMessage() {}

func (x *SqlOperationsGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_sql_v1_cloud_sql_operations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SqlOperationsGetRequest.ProtoReflect.Descriptor instead.
func (*SqlOperationsGetRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_sql_v1_cloud_sql_operations_proto_rawDescGZIP(), []int{0}
}

func (x *SqlOperationsGetRequest) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *SqlOperationsGetRequest) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

// Operations list request.
type SqlOperationsListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Cloud SQL instance ID. This does not include the project ID.
	Instance string `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	// Maximum number of operations per response.
	MaxResults uint32 `protobuf:"varint,2,opt,name=max_results,json=maxResults,proto3" json:"max_results,omitempty"`
	// A previously-returned page token representing part of the larger set of
	// results to view.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Project ID of the project that contains the instance.
	Project string `protobuf:"bytes,4,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *SqlOperationsListRequest) Reset() {
	*x = SqlOperationsListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_sql_v1_cloud_sql_operations_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SqlOperationsListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SqlOperationsListRequest) ProtoMessage() {}

func (x *SqlOperationsListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_sql_v1_cloud_sql_operations_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SqlOperationsListRequest.ProtoReflect.Descriptor instead.
func (*SqlOperationsListRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_sql_v1_cloud_sql_operations_proto_rawDescGZIP(), []int{1}
}

func (x *SqlOperationsListRequest) GetInstance() string {
	if x != nil {
		return x.Instance
	}
	return ""
}

func (x *SqlOperationsListRequest) GetMaxResults() uint32 {
	if x != nil {
		return x.MaxResults
	}
	return 0
}

func (x *SqlOperationsListRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *SqlOperationsListRequest) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

// Operations list response.
type OperationsListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This is always **sql#operationsList**.
	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	// List of operation resources.
	Items []*Operation `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	// The continuation token, used to page through large result sets. Provide
	// this value in a subsequent request to return the next page of results.
	NextPageToken string `protobuf:"bytes,3,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *OperationsListResponse) Reset() {
	*x = OperationsListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_sql_v1_cloud_sql_operations_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationsListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationsListResponse) ProtoMessage() {}

func (x *OperationsListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_sql_v1_cloud_sql_operations_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationsListResponse.ProtoReflect.Descriptor instead.
func (*OperationsListResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_sql_v1_cloud_sql_operations_proto_rawDescGZIP(), []int{2}
}

func (x *OperationsListResponse) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *OperationsListResponse) GetItems() []*Operation {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *OperationsListResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_google_cloud_sql_v1_cloud_sql_operations_proto protoreflect.FileDescriptor

var file_google_cloud_sql_v1_cloud_sql_operations_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x71, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x73, 0x71, 0x6c, 0x5f,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x13, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73,
	0x71, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x73, 0x71, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x73,
	0x71, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x17, 0x53,
	0x71, 0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x90,
	0x01, 0x0a, 0x18, 0x53, 0x71, 0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x5f, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6d, 0x61,
	0x78, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x22, 0x8a, 0x01, 0x0a, 0x16, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x12, 0x34, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73,
	0x71, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xb1,
	0x03, 0x0a, 0x14, 0x53, 0x71, 0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8a, 0x01, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73,
	0x71, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x71, 0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x71, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x35, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x12, 0x2d, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x7d, 0x12, 0x8d, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x71, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x71, 0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x71, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x23, 0x12, 0x21, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x7c, 0xca, 0x41, 0x17, 0x73, 0x71, 0x6c, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0xd2, 0x41, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2c, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x73, 0x71, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x42, 0x6c, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x42, 0x17, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x53, 0x71, 0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x71, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x71, 0x6c,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_sql_v1_cloud_sql_operations_proto_rawDescOnce sync.Once
	file_google_cloud_sql_v1_cloud_sql_operations_proto_rawDescData = file_google_cloud_sql_v1_cloud_sql_operations_proto_rawDesc
)

func file_google_cloud_sql_v1_cloud_sql_operations_proto_rawDescGZIP() []byte {
	file_google_cloud_sql_v1_cloud_sql_operations_proto_rawDescOnce.Do(func() {
		file_google_cloud_sql_v1_cloud_sql_operations_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_sql_v1_cloud_sql_operations_proto_rawDescData)
	})
	return file_google_cloud_sql_v1_cloud_sql_operations_proto_rawDescData
}

var file_google_cloud_sql_v1_cloud_sql_operations_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_sql_v1_cloud_sql_operations_proto_goTypes = []interface{}{
	(*SqlOperationsGetRequest)(nil),  // 0: google.cloud.sql.v1.SqlOperationsGetRequest
	(*SqlOperationsListRequest)(nil), // 1: google.cloud.sql.v1.SqlOperationsListRequest
	(*OperationsListResponse)(nil),   // 2: google.cloud.sql.v1.OperationsListResponse
	(*Operation)(nil),                // 3: google.cloud.sql.v1.Operation
}
var file_google_cloud_sql_v1_cloud_sql_operations_proto_depIdxs = []int32{
	3, // 0: google.cloud.sql.v1.OperationsListResponse.items:type_name -> google.cloud.sql.v1.Operation
	0, // 1: google.cloud.sql.v1.SqlOperationsService.Get:input_type -> google.cloud.sql.v1.SqlOperationsGetRequest
	1, // 2: google.cloud.sql.v1.SqlOperationsService.List:input_type -> google.cloud.sql.v1.SqlOperationsListRequest
	3, // 3: google.cloud.sql.v1.SqlOperationsService.Get:output_type -> google.cloud.sql.v1.Operation
	2, // 4: google.cloud.sql.v1.SqlOperationsService.List:output_type -> google.cloud.sql.v1.OperationsListResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_sql_v1_cloud_sql_operations_proto_init() }
func file_google_cloud_sql_v1_cloud_sql_operations_proto_init() {
	if File_google_cloud_sql_v1_cloud_sql_operations_proto != nil {
		return
	}
	file_google_cloud_sql_v1_cloud_sql_resources_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_sql_v1_cloud_sql_operations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SqlOperationsGetRequest); i {
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
		file_google_cloud_sql_v1_cloud_sql_operations_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SqlOperationsListRequest); i {
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
		file_google_cloud_sql_v1_cloud_sql_operations_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationsListResponse); i {
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
			RawDescriptor: file_google_cloud_sql_v1_cloud_sql_operations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_sql_v1_cloud_sql_operations_proto_goTypes,
		DependencyIndexes: file_google_cloud_sql_v1_cloud_sql_operations_proto_depIdxs,
		MessageInfos:      file_google_cloud_sql_v1_cloud_sql_operations_proto_msgTypes,
	}.Build()
	File_google_cloud_sql_v1_cloud_sql_operations_proto = out.File
	file_google_cloud_sql_v1_cloud_sql_operations_proto_rawDesc = nil
	file_google_cloud_sql_v1_cloud_sql_operations_proto_goTypes = nil
	file_google_cloud_sql_v1_cloud_sql_operations_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SqlOperationsServiceClient is the client API for SqlOperationsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SqlOperationsServiceClient interface {
	// Retrieves an instance operation that has been performed on an instance.
	Get(ctx context.Context, in *SqlOperationsGetRequest, opts ...grpc.CallOption) (*Operation, error)
	// Lists all instance operations that have been performed on the given Cloud
	// SQL instance in the reverse chronological order of the start time.
	List(ctx context.Context, in *SqlOperationsListRequest, opts ...grpc.CallOption) (*OperationsListResponse, error)
}

type sqlOperationsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSqlOperationsServiceClient(cc grpc.ClientConnInterface) SqlOperationsServiceClient {
	return &sqlOperationsServiceClient{cc}
}

func (c *sqlOperationsServiceClient) Get(ctx context.Context, in *SqlOperationsGetRequest, opts ...grpc.CallOption) (*Operation, error) {
	out := new(Operation)
	err := c.cc.Invoke(ctx, "/google.cloud.sql.v1.SqlOperationsService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sqlOperationsServiceClient) List(ctx context.Context, in *SqlOperationsListRequest, opts ...grpc.CallOption) (*OperationsListResponse, error) {
	out := new(OperationsListResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.sql.v1.SqlOperationsService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SqlOperationsServiceServer is the server API for SqlOperationsService service.
type SqlOperationsServiceServer interface {
	// Retrieves an instance operation that has been performed on an instance.
	Get(context.Context, *SqlOperationsGetRequest) (*Operation, error)
	// Lists all instance operations that have been performed on the given Cloud
	// SQL instance in the reverse chronological order of the start time.
	List(context.Context, *SqlOperationsListRequest) (*OperationsListResponse, error)
}

// UnimplementedSqlOperationsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSqlOperationsServiceServer struct {
}

func (*UnimplementedSqlOperationsServiceServer) Get(context.Context, *SqlOperationsGetRequest) (*Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedSqlOperationsServiceServer) List(context.Context, *SqlOperationsListRequest) (*OperationsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterSqlOperationsServiceServer(s *grpc.Server, srv SqlOperationsServiceServer) {
	s.RegisterService(&_SqlOperationsService_serviceDesc, srv)
}

func _SqlOperationsService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SqlOperationsGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SqlOperationsServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.sql.v1.SqlOperationsService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SqlOperationsServiceServer).Get(ctx, req.(*SqlOperationsGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SqlOperationsService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SqlOperationsListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SqlOperationsServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.sql.v1.SqlOperationsService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SqlOperationsServiceServer).List(ctx, req.(*SqlOperationsListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SqlOperationsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.sql.v1.SqlOperationsService",
	HandlerType: (*SqlOperationsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _SqlOperationsService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _SqlOperationsService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/sql/v1/cloud_sql_operations.proto",
}
