// Copyright 2022 Google LLC
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
// 	protoc        v3.12.2
// source: google/cloud/functions/v1/operations.proto

package functions

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A type of an operation.
type OperationType int32

const (
	// Unknown operation type.
	OperationType_OPERATION_UNSPECIFIED OperationType = 0
	// Triggered by CreateFunction call
	OperationType_CREATE_FUNCTION OperationType = 1
	// Triggered by UpdateFunction call
	OperationType_UPDATE_FUNCTION OperationType = 2
	// Triggered by DeleteFunction call.
	OperationType_DELETE_FUNCTION OperationType = 3
)

// Enum value maps for OperationType.
var (
	OperationType_name = map[int32]string{
		0: "OPERATION_UNSPECIFIED",
		1: "CREATE_FUNCTION",
		2: "UPDATE_FUNCTION",
		3: "DELETE_FUNCTION",
	}
	OperationType_value = map[string]int32{
		"OPERATION_UNSPECIFIED": 0,
		"CREATE_FUNCTION":       1,
		"UPDATE_FUNCTION":       2,
		"DELETE_FUNCTION":       3,
	}
)

func (x OperationType) Enum() *OperationType {
	p := new(OperationType)
	*p = x
	return p
}

func (x OperationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperationType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_functions_v1_operations_proto_enumTypes[0].Descriptor()
}

func (OperationType) Type() protoreflect.EnumType {
	return &file_google_cloud_functions_v1_operations_proto_enumTypes[0]
}

func (x OperationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperationType.Descriptor instead.
func (OperationType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_functions_v1_operations_proto_rawDescGZIP(), []int{0}
}

// Metadata describing an [Operation][google.longrunning.Operation]
type OperationMetadataV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Target of the operation - for example
	// `projects/project-1/locations/region-1/functions/function-1`
	Target string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	// Type of operation.
	Type OperationType `protobuf:"varint,2,opt,name=type,proto3,enum=google.cloud.functions.v1.OperationType" json:"type,omitempty"`
	// The original request that started the operation.
	Request *anypb.Any `protobuf:"bytes,3,opt,name=request,proto3" json:"request,omitempty"`
	// Version id of the function created or updated by an API call.
	// This field is only populated for Create and Update operations.
	VersionId int64 `protobuf:"varint,4,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	// The last update timestamp of the operation.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// The Cloud Build ID of the function created or updated by an API call.
	// This field is only populated for Create and Update operations.
	BuildId string `protobuf:"bytes,6,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
	// An identifier for Firebase function sources. Disclaimer: This field is only
	// supported for Firebase function deployments.
	SourceToken string `protobuf:"bytes,7,opt,name=source_token,json=sourceToken,proto3" json:"source_token,omitempty"`
	// The Cloud Build Name of the function deployment.
	// This field is only populated for Create and Update operations.
	// `projects/<project-number>/locations/<region>/builds/<build-id>`.
	BuildName string `protobuf:"bytes,8,opt,name=build_name,json=buildName,proto3" json:"build_name,omitempty"`
}

func (x *OperationMetadataV1) Reset() {
	*x = OperationMetadataV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_functions_v1_operations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationMetadataV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationMetadataV1) ProtoMessage() {}

func (x *OperationMetadataV1) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_functions_v1_operations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationMetadataV1.ProtoReflect.Descriptor instead.
func (*OperationMetadataV1) Descriptor() ([]byte, []int) {
	return file_google_cloud_functions_v1_operations_proto_rawDescGZIP(), []int{0}
}

func (x *OperationMetadataV1) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *OperationMetadataV1) GetType() OperationType {
	if x != nil {
		return x.Type
	}
	return OperationType_OPERATION_UNSPECIFIED
}

func (x *OperationMetadataV1) GetRequest() *anypb.Any {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *OperationMetadataV1) GetVersionId() int64 {
	if x != nil {
		return x.VersionId
	}
	return 0
}

func (x *OperationMetadataV1) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *OperationMetadataV1) GetBuildId() string {
	if x != nil {
		return x.BuildId
	}
	return ""
}

func (x *OperationMetadataV1) GetSourceToken() string {
	if x != nil {
		return x.SourceToken
	}
	return ""
}

func (x *OperationMetadataV1) GetBuildName() string {
	if x != nil {
		return x.BuildName
	}
	return ""
}

var File_google_cloud_functions_v1_operations_proto protoreflect.FileDescriptor

var file_google_cloud_functions_v1_operations_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x66,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x02, 0x0a, 0x13, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x56, 0x31, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x12, 0x3c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x2e, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x2a, 0x69, 0x0a, 0x0d, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x4f,
	0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45,
	0x5f, 0x46, 0x55, 0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x55,
	0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x55, 0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02,
	0x12, 0x13, 0x0a, 0x0f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x46, 0x55, 0x4e, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x03, 0x42, 0x7f, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x18, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x42, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x66, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_functions_v1_operations_proto_rawDescOnce sync.Once
	file_google_cloud_functions_v1_operations_proto_rawDescData = file_google_cloud_functions_v1_operations_proto_rawDesc
)

func file_google_cloud_functions_v1_operations_proto_rawDescGZIP() []byte {
	file_google_cloud_functions_v1_operations_proto_rawDescOnce.Do(func() {
		file_google_cloud_functions_v1_operations_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_functions_v1_operations_proto_rawDescData)
	})
	return file_google_cloud_functions_v1_operations_proto_rawDescData
}

var file_google_cloud_functions_v1_operations_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_functions_v1_operations_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_functions_v1_operations_proto_goTypes = []interface{}{
	(OperationType)(0),            // 0: google.cloud.functions.v1.OperationType
	(*OperationMetadataV1)(nil),   // 1: google.cloud.functions.v1.OperationMetadataV1
	(*anypb.Any)(nil),             // 2: google.protobuf.Any
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_google_cloud_functions_v1_operations_proto_depIdxs = []int32{
	0, // 0: google.cloud.functions.v1.OperationMetadataV1.type:type_name -> google.cloud.functions.v1.OperationType
	2, // 1: google.cloud.functions.v1.OperationMetadataV1.request:type_name -> google.protobuf.Any
	3, // 2: google.cloud.functions.v1.OperationMetadataV1.update_time:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_cloud_functions_v1_operations_proto_init() }
func file_google_cloud_functions_v1_operations_proto_init() {
	if File_google_cloud_functions_v1_operations_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_functions_v1_operations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationMetadataV1); i {
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
			RawDescriptor: file_google_cloud_functions_v1_operations_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_functions_v1_operations_proto_goTypes,
		DependencyIndexes: file_google_cloud_functions_v1_operations_proto_depIdxs,
		EnumInfos:         file_google_cloud_functions_v1_operations_proto_enumTypes,
		MessageInfos:      file_google_cloud_functions_v1_operations_proto_msgTypes,
	}.Build()
	File_google_cloud_functions_v1_operations_proto = out.File
	file_google_cloud_functions_v1_operations_proto_rawDesc = nil
	file_google_cloud_functions_v1_operations_proto_goTypes = nil
	file_google_cloud_functions_v1_operations_proto_depIdxs = nil
}
