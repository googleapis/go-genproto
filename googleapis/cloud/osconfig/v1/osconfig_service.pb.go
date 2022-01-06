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
// source: google/cloud/osconfig/v1/osconfig_service.proto

package osconfig

import (
	context "context"
	reflect "reflect"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_google_cloud_osconfig_v1_osconfig_service_proto protoreflect.FileDescriptor

var file_google_cloud_osconfig_v1_osconfig_service_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f,
	0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x73, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x73, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x73,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x74, 0x63, 0x68, 0x5f,
	0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xbe, 0x0d, 0x0a, 0x0f, 0x4f, 0x73, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9d, 0x01, 0x0a, 0x0f,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x50, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x12,
	0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f,
	0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x65, 0x50, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x74,
	0x63, 0x68, 0x4a, 0x6f, 0x62, 0x22, 0x34, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x22, 0x29, 0x2f,
	0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x70, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x73,
	0x3a, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x91, 0x01, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x12, 0x2c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x73, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x74, 0x63, 0x68, 0x4a,
	0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x22, 0x30, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x23, 0x12, 0x21, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65,
	0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x70, 0x61, 0x74, 0x63,
	0x68, 0x4a, 0x6f, 0x62, 0x73, 0x2f, 0x2a, 0x7d, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x9a, 0x01, 0x0a, 0x0e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x50, 0x61, 0x74, 0x63, 0x68, 0x4a,
	0x6f, 0x62, 0x12, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x50, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x22, 0x33, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x22,
	0x28, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x70, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x73, 0x2f,
	0x2a, 0x7d, 0x3a, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x3a, 0x01, 0x2a, 0x12, 0xa4, 0x01, 0x0a,
	0x0d, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x2e,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x73,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61,
	0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x73,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61,
	0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x12, 0x21, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x7d,
	0x2f, 0x70, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x73, 0xda, 0x41, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x12, 0xe0, 0x01, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x74, 0x63,
	0x68, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x12, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x3d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x44, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x35, 0x12, 0x33, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a,
	0x2f, 0x70, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0xda, 0x41, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0xec, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x22, 0x70, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3c, 0x22, 0x28, 0x2f, 0x76, 0x31,
	0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x70, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x10, 0x70, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0xda, 0x41, 0x2b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x2c, 0x70, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2c, 0x70, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x12, 0xad, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x61, 0x74,
	0x63, 0x68, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x33, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x73, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x74, 0x63, 0x68,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x74,
	0x63, 0x68, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x37, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x2a, 0x12, 0x28, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0xda, 0x41,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0xc0, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61,
	0x74, 0x63, 0x68, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x35,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x73,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61,
	0x74, 0x63, 0x68, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x12, 0x28, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0xda,
	0x41, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0xa0, 0x01, 0x0a, 0x15, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x37, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x2a, 0x28, 0x2f, 0x76, 0x31, 0x2f,
	0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a,
	0x2f, 0x70, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x2f, 0x2a, 0x7d, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x4b, 0xca, 0x41, 0x17,
	0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0xdc, 0x02, 0x0a, 0x1c, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x73,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x4f, 0x73, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76,
	0x31, 0x3b, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0xaa, 0x02, 0x18, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4f, 0x73, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x18, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x4f, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5c, 0x56, 0x31,
	0xea, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x3a, 0x3a, 0x4f, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a, 0x3a, 0x56, 0x31, 0xea, 0x41,
	0x95, 0x01, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x34, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x2f, 0x7b, 0x7a,
	0x6f, 0x6e, 0x65, 0x7d, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x7b,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x7d, 0x12, 0x3c, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x7d, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x7d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_google_cloud_osconfig_v1_osconfig_service_proto_goTypes = []interface{}{
	(*ExecutePatchJobRequest)(nil),              // 0: google.cloud.osconfig.v1.ExecutePatchJobRequest
	(*GetPatchJobRequest)(nil),                  // 1: google.cloud.osconfig.v1.GetPatchJobRequest
	(*CancelPatchJobRequest)(nil),               // 2: google.cloud.osconfig.v1.CancelPatchJobRequest
	(*ListPatchJobsRequest)(nil),                // 3: google.cloud.osconfig.v1.ListPatchJobsRequest
	(*ListPatchJobInstanceDetailsRequest)(nil),  // 4: google.cloud.osconfig.v1.ListPatchJobInstanceDetailsRequest
	(*CreatePatchDeploymentRequest)(nil),        // 5: google.cloud.osconfig.v1.CreatePatchDeploymentRequest
	(*GetPatchDeploymentRequest)(nil),           // 6: google.cloud.osconfig.v1.GetPatchDeploymentRequest
	(*ListPatchDeploymentsRequest)(nil),         // 7: google.cloud.osconfig.v1.ListPatchDeploymentsRequest
	(*DeletePatchDeploymentRequest)(nil),        // 8: google.cloud.osconfig.v1.DeletePatchDeploymentRequest
	(*PatchJob)(nil),                            // 9: google.cloud.osconfig.v1.PatchJob
	(*ListPatchJobsResponse)(nil),               // 10: google.cloud.osconfig.v1.ListPatchJobsResponse
	(*ListPatchJobInstanceDetailsResponse)(nil), // 11: google.cloud.osconfig.v1.ListPatchJobInstanceDetailsResponse
	(*PatchDeployment)(nil),                     // 12: google.cloud.osconfig.v1.PatchDeployment
	(*ListPatchDeploymentsResponse)(nil),        // 13: google.cloud.osconfig.v1.ListPatchDeploymentsResponse
	(*emptypb.Empty)(nil),                       // 14: google.protobuf.Empty
}
var file_google_cloud_osconfig_v1_osconfig_service_proto_depIdxs = []int32{
	0,  // 0: google.cloud.osconfig.v1.OsConfigService.ExecutePatchJob:input_type -> google.cloud.osconfig.v1.ExecutePatchJobRequest
	1,  // 1: google.cloud.osconfig.v1.OsConfigService.GetPatchJob:input_type -> google.cloud.osconfig.v1.GetPatchJobRequest
	2,  // 2: google.cloud.osconfig.v1.OsConfigService.CancelPatchJob:input_type -> google.cloud.osconfig.v1.CancelPatchJobRequest
	3,  // 3: google.cloud.osconfig.v1.OsConfigService.ListPatchJobs:input_type -> google.cloud.osconfig.v1.ListPatchJobsRequest
	4,  // 4: google.cloud.osconfig.v1.OsConfigService.ListPatchJobInstanceDetails:input_type -> google.cloud.osconfig.v1.ListPatchJobInstanceDetailsRequest
	5,  // 5: google.cloud.osconfig.v1.OsConfigService.CreatePatchDeployment:input_type -> google.cloud.osconfig.v1.CreatePatchDeploymentRequest
	6,  // 6: google.cloud.osconfig.v1.OsConfigService.GetPatchDeployment:input_type -> google.cloud.osconfig.v1.GetPatchDeploymentRequest
	7,  // 7: google.cloud.osconfig.v1.OsConfigService.ListPatchDeployments:input_type -> google.cloud.osconfig.v1.ListPatchDeploymentsRequest
	8,  // 8: google.cloud.osconfig.v1.OsConfigService.DeletePatchDeployment:input_type -> google.cloud.osconfig.v1.DeletePatchDeploymentRequest
	9,  // 9: google.cloud.osconfig.v1.OsConfigService.ExecutePatchJob:output_type -> google.cloud.osconfig.v1.PatchJob
	9,  // 10: google.cloud.osconfig.v1.OsConfigService.GetPatchJob:output_type -> google.cloud.osconfig.v1.PatchJob
	9,  // 11: google.cloud.osconfig.v1.OsConfigService.CancelPatchJob:output_type -> google.cloud.osconfig.v1.PatchJob
	10, // 12: google.cloud.osconfig.v1.OsConfigService.ListPatchJobs:output_type -> google.cloud.osconfig.v1.ListPatchJobsResponse
	11, // 13: google.cloud.osconfig.v1.OsConfigService.ListPatchJobInstanceDetails:output_type -> google.cloud.osconfig.v1.ListPatchJobInstanceDetailsResponse
	12, // 14: google.cloud.osconfig.v1.OsConfigService.CreatePatchDeployment:output_type -> google.cloud.osconfig.v1.PatchDeployment
	12, // 15: google.cloud.osconfig.v1.OsConfigService.GetPatchDeployment:output_type -> google.cloud.osconfig.v1.PatchDeployment
	13, // 16: google.cloud.osconfig.v1.OsConfigService.ListPatchDeployments:output_type -> google.cloud.osconfig.v1.ListPatchDeploymentsResponse
	14, // 17: google.cloud.osconfig.v1.OsConfigService.DeletePatchDeployment:output_type -> google.protobuf.Empty
	9,  // [9:18] is the sub-list for method output_type
	0,  // [0:9] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_osconfig_v1_osconfig_service_proto_init() }
func file_google_cloud_osconfig_v1_osconfig_service_proto_init() {
	if File_google_cloud_osconfig_v1_osconfig_service_proto != nil {
		return
	}
	file_google_cloud_osconfig_v1_patch_deployments_proto_init()
	file_google_cloud_osconfig_v1_patch_jobs_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_osconfig_v1_osconfig_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_osconfig_v1_osconfig_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_osconfig_v1_osconfig_service_proto_depIdxs,
	}.Build()
	File_google_cloud_osconfig_v1_osconfig_service_proto = out.File
	file_google_cloud_osconfig_v1_osconfig_service_proto_rawDesc = nil
	file_google_cloud_osconfig_v1_osconfig_service_proto_goTypes = nil
	file_google_cloud_osconfig_v1_osconfig_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OsConfigServiceClient is the client API for OsConfigService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OsConfigServiceClient interface {
	// Patch VM instances by creating and running a patch job.
	ExecutePatchJob(ctx context.Context, in *ExecutePatchJobRequest, opts ...grpc.CallOption) (*PatchJob, error)
	// Get the patch job. This can be used to track the progress of an
	// ongoing patch job or review the details of completed jobs.
	GetPatchJob(ctx context.Context, in *GetPatchJobRequest, opts ...grpc.CallOption) (*PatchJob, error)
	// Cancel a patch job. The patch job must be active. Canceled patch jobs
	// cannot be restarted.
	CancelPatchJob(ctx context.Context, in *CancelPatchJobRequest, opts ...grpc.CallOption) (*PatchJob, error)
	// Get a list of patch jobs.
	ListPatchJobs(ctx context.Context, in *ListPatchJobsRequest, opts ...grpc.CallOption) (*ListPatchJobsResponse, error)
	// Get a list of instance details for a given patch job.
	ListPatchJobInstanceDetails(ctx context.Context, in *ListPatchJobInstanceDetailsRequest, opts ...grpc.CallOption) (*ListPatchJobInstanceDetailsResponse, error)
	// Create an OS Config patch deployment.
	CreatePatchDeployment(ctx context.Context, in *CreatePatchDeploymentRequest, opts ...grpc.CallOption) (*PatchDeployment, error)
	// Get an OS Config patch deployment.
	GetPatchDeployment(ctx context.Context, in *GetPatchDeploymentRequest, opts ...grpc.CallOption) (*PatchDeployment, error)
	// Get a page of OS Config patch deployments.
	ListPatchDeployments(ctx context.Context, in *ListPatchDeploymentsRequest, opts ...grpc.CallOption) (*ListPatchDeploymentsResponse, error)
	// Delete an OS Config patch deployment.
	DeletePatchDeployment(ctx context.Context, in *DeletePatchDeploymentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type osConfigServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOsConfigServiceClient(cc grpc.ClientConnInterface) OsConfigServiceClient {
	return &osConfigServiceClient{cc}
}

func (c *osConfigServiceClient) ExecutePatchJob(ctx context.Context, in *ExecutePatchJobRequest, opts ...grpc.CallOption) (*PatchJob, error) {
	out := new(PatchJob)
	err := c.cc.Invoke(ctx, "/google.cloud.osconfig.v1.OsConfigService/ExecutePatchJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *osConfigServiceClient) GetPatchJob(ctx context.Context, in *GetPatchJobRequest, opts ...grpc.CallOption) (*PatchJob, error) {
	out := new(PatchJob)
	err := c.cc.Invoke(ctx, "/google.cloud.osconfig.v1.OsConfigService/GetPatchJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *osConfigServiceClient) CancelPatchJob(ctx context.Context, in *CancelPatchJobRequest, opts ...grpc.CallOption) (*PatchJob, error) {
	out := new(PatchJob)
	err := c.cc.Invoke(ctx, "/google.cloud.osconfig.v1.OsConfigService/CancelPatchJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *osConfigServiceClient) ListPatchJobs(ctx context.Context, in *ListPatchJobsRequest, opts ...grpc.CallOption) (*ListPatchJobsResponse, error) {
	out := new(ListPatchJobsResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.osconfig.v1.OsConfigService/ListPatchJobs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *osConfigServiceClient) ListPatchJobInstanceDetails(ctx context.Context, in *ListPatchJobInstanceDetailsRequest, opts ...grpc.CallOption) (*ListPatchJobInstanceDetailsResponse, error) {
	out := new(ListPatchJobInstanceDetailsResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.osconfig.v1.OsConfigService/ListPatchJobInstanceDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *osConfigServiceClient) CreatePatchDeployment(ctx context.Context, in *CreatePatchDeploymentRequest, opts ...grpc.CallOption) (*PatchDeployment, error) {
	out := new(PatchDeployment)
	err := c.cc.Invoke(ctx, "/google.cloud.osconfig.v1.OsConfigService/CreatePatchDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *osConfigServiceClient) GetPatchDeployment(ctx context.Context, in *GetPatchDeploymentRequest, opts ...grpc.CallOption) (*PatchDeployment, error) {
	out := new(PatchDeployment)
	err := c.cc.Invoke(ctx, "/google.cloud.osconfig.v1.OsConfigService/GetPatchDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *osConfigServiceClient) ListPatchDeployments(ctx context.Context, in *ListPatchDeploymentsRequest, opts ...grpc.CallOption) (*ListPatchDeploymentsResponse, error) {
	out := new(ListPatchDeploymentsResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.osconfig.v1.OsConfigService/ListPatchDeployments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *osConfigServiceClient) DeletePatchDeployment(ctx context.Context, in *DeletePatchDeploymentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/google.cloud.osconfig.v1.OsConfigService/DeletePatchDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OsConfigServiceServer is the server API for OsConfigService service.
type OsConfigServiceServer interface {
	// Patch VM instances by creating and running a patch job.
	ExecutePatchJob(context.Context, *ExecutePatchJobRequest) (*PatchJob, error)
	// Get the patch job. This can be used to track the progress of an
	// ongoing patch job or review the details of completed jobs.
	GetPatchJob(context.Context, *GetPatchJobRequest) (*PatchJob, error)
	// Cancel a patch job. The patch job must be active. Canceled patch jobs
	// cannot be restarted.
	CancelPatchJob(context.Context, *CancelPatchJobRequest) (*PatchJob, error)
	// Get a list of patch jobs.
	ListPatchJobs(context.Context, *ListPatchJobsRequest) (*ListPatchJobsResponse, error)
	// Get a list of instance details for a given patch job.
	ListPatchJobInstanceDetails(context.Context, *ListPatchJobInstanceDetailsRequest) (*ListPatchJobInstanceDetailsResponse, error)
	// Create an OS Config patch deployment.
	CreatePatchDeployment(context.Context, *CreatePatchDeploymentRequest) (*PatchDeployment, error)
	// Get an OS Config patch deployment.
	GetPatchDeployment(context.Context, *GetPatchDeploymentRequest) (*PatchDeployment, error)
	// Get a page of OS Config patch deployments.
	ListPatchDeployments(context.Context, *ListPatchDeploymentsRequest) (*ListPatchDeploymentsResponse, error)
	// Delete an OS Config patch deployment.
	DeletePatchDeployment(context.Context, *DeletePatchDeploymentRequest) (*emptypb.Empty, error)
}

// UnimplementedOsConfigServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOsConfigServiceServer struct {
}

func (*UnimplementedOsConfigServiceServer) ExecutePatchJob(context.Context, *ExecutePatchJobRequest) (*PatchJob, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecutePatchJob not implemented")
}
func (*UnimplementedOsConfigServiceServer) GetPatchJob(context.Context, *GetPatchJobRequest) (*PatchJob, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPatchJob not implemented")
}
func (*UnimplementedOsConfigServiceServer) CancelPatchJob(context.Context, *CancelPatchJobRequest) (*PatchJob, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelPatchJob not implemented")
}
func (*UnimplementedOsConfigServiceServer) ListPatchJobs(context.Context, *ListPatchJobsRequest) (*ListPatchJobsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPatchJobs not implemented")
}
func (*UnimplementedOsConfigServiceServer) ListPatchJobInstanceDetails(context.Context, *ListPatchJobInstanceDetailsRequest) (*ListPatchJobInstanceDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPatchJobInstanceDetails not implemented")
}
func (*UnimplementedOsConfigServiceServer) CreatePatchDeployment(context.Context, *CreatePatchDeploymentRequest) (*PatchDeployment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePatchDeployment not implemented")
}
func (*UnimplementedOsConfigServiceServer) GetPatchDeployment(context.Context, *GetPatchDeploymentRequest) (*PatchDeployment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPatchDeployment not implemented")
}
func (*UnimplementedOsConfigServiceServer) ListPatchDeployments(context.Context, *ListPatchDeploymentsRequest) (*ListPatchDeploymentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPatchDeployments not implemented")
}
func (*UnimplementedOsConfigServiceServer) DeletePatchDeployment(context.Context, *DeletePatchDeploymentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePatchDeployment not implemented")
}

func RegisterOsConfigServiceServer(s *grpc.Server, srv OsConfigServiceServer) {
	s.RegisterService(&_OsConfigService_serviceDesc, srv)
}

func _OsConfigService_ExecutePatchJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecutePatchJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OsConfigServiceServer).ExecutePatchJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.osconfig.v1.OsConfigService/ExecutePatchJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OsConfigServiceServer).ExecutePatchJob(ctx, req.(*ExecutePatchJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OsConfigService_GetPatchJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPatchJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OsConfigServiceServer).GetPatchJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.osconfig.v1.OsConfigService/GetPatchJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OsConfigServiceServer).GetPatchJob(ctx, req.(*GetPatchJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OsConfigService_CancelPatchJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelPatchJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OsConfigServiceServer).CancelPatchJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.osconfig.v1.OsConfigService/CancelPatchJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OsConfigServiceServer).CancelPatchJob(ctx, req.(*CancelPatchJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OsConfigService_ListPatchJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPatchJobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OsConfigServiceServer).ListPatchJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.osconfig.v1.OsConfigService/ListPatchJobs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OsConfigServiceServer).ListPatchJobs(ctx, req.(*ListPatchJobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OsConfigService_ListPatchJobInstanceDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPatchJobInstanceDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OsConfigServiceServer).ListPatchJobInstanceDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.osconfig.v1.OsConfigService/ListPatchJobInstanceDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OsConfigServiceServer).ListPatchJobInstanceDetails(ctx, req.(*ListPatchJobInstanceDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OsConfigService_CreatePatchDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePatchDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OsConfigServiceServer).CreatePatchDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.osconfig.v1.OsConfigService/CreatePatchDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OsConfigServiceServer).CreatePatchDeployment(ctx, req.(*CreatePatchDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OsConfigService_GetPatchDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPatchDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OsConfigServiceServer).GetPatchDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.osconfig.v1.OsConfigService/GetPatchDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OsConfigServiceServer).GetPatchDeployment(ctx, req.(*GetPatchDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OsConfigService_ListPatchDeployments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPatchDeploymentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OsConfigServiceServer).ListPatchDeployments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.osconfig.v1.OsConfigService/ListPatchDeployments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OsConfigServiceServer).ListPatchDeployments(ctx, req.(*ListPatchDeploymentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OsConfigService_DeletePatchDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePatchDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OsConfigServiceServer).DeletePatchDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.osconfig.v1.OsConfigService/DeletePatchDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OsConfigServiceServer).DeletePatchDeployment(ctx, req.(*DeletePatchDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OsConfigService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.osconfig.v1.OsConfigService",
	HandlerType: (*OsConfigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecutePatchJob",
			Handler:    _OsConfigService_ExecutePatchJob_Handler,
		},
		{
			MethodName: "GetPatchJob",
			Handler:    _OsConfigService_GetPatchJob_Handler,
		},
		{
			MethodName: "CancelPatchJob",
			Handler:    _OsConfigService_CancelPatchJob_Handler,
		},
		{
			MethodName: "ListPatchJobs",
			Handler:    _OsConfigService_ListPatchJobs_Handler,
		},
		{
			MethodName: "ListPatchJobInstanceDetails",
			Handler:    _OsConfigService_ListPatchJobInstanceDetails_Handler,
		},
		{
			MethodName: "CreatePatchDeployment",
			Handler:    _OsConfigService_CreatePatchDeployment_Handler,
		},
		{
			MethodName: "GetPatchDeployment",
			Handler:    _OsConfigService_GetPatchDeployment_Handler,
		},
		{
			MethodName: "ListPatchDeployments",
			Handler:    _OsConfigService_ListPatchDeployments_Handler,
		},
		{
			MethodName: "DeletePatchDeployment",
			Handler:    _OsConfigService_DeletePatchDeployment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/osconfig/v1/osconfig_service.proto",
}
