// Copyright 2019 Google LLC.
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
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.12.3
// source: google/devtools/resultstore/v2/file_processing_error.proto

package resultstore

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
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

// Errors in file post-processing are categorized using this enum.
type FileProcessingErrorType int32

const (
	// Type unspecified or not listed here.
	FileProcessingErrorType_FILE_PROCESSING_ERROR_TYPE_UNSPECIFIED FileProcessingErrorType = 0
	// A read error occurred trying to read the file.
	FileProcessingErrorType_GENERIC_READ_ERROR FileProcessingErrorType = 1
	// There was an error trying to parse the file.
	FileProcessingErrorType_GENERIC_PARSE_ERROR FileProcessingErrorType = 2
	// File is exceeds size limit.
	FileProcessingErrorType_FILE_TOO_LARGE FileProcessingErrorType = 3
	// The result of parsing the file exceeded size limit.
	FileProcessingErrorType_OUTPUT_TOO_LARGE FileProcessingErrorType = 4
	// Read access to the file was denied by file system.
	FileProcessingErrorType_ACCESS_DENIED FileProcessingErrorType = 5
	// Deadline exceeded trying to read the file.
	FileProcessingErrorType_DEADLINE_EXCEEDED FileProcessingErrorType = 6
	// File not found.
	FileProcessingErrorType_NOT_FOUND FileProcessingErrorType = 7
	// File is empty but was expected to have content.
	FileProcessingErrorType_FILE_EMPTY FileProcessingErrorType = 8
)

// Enum value maps for FileProcessingErrorType.
var (
	FileProcessingErrorType_name = map[int32]string{
		0: "FILE_PROCESSING_ERROR_TYPE_UNSPECIFIED",
		1: "GENERIC_READ_ERROR",
		2: "GENERIC_PARSE_ERROR",
		3: "FILE_TOO_LARGE",
		4: "OUTPUT_TOO_LARGE",
		5: "ACCESS_DENIED",
		6: "DEADLINE_EXCEEDED",
		7: "NOT_FOUND",
		8: "FILE_EMPTY",
	}
	FileProcessingErrorType_value = map[string]int32{
		"FILE_PROCESSING_ERROR_TYPE_UNSPECIFIED": 0,
		"GENERIC_READ_ERROR":                     1,
		"GENERIC_PARSE_ERROR":                    2,
		"FILE_TOO_LARGE":                         3,
		"OUTPUT_TOO_LARGE":                       4,
		"ACCESS_DENIED":                          5,
		"DEADLINE_EXCEEDED":                      6,
		"NOT_FOUND":                              7,
		"FILE_EMPTY":                             8,
	}
)

func (x FileProcessingErrorType) Enum() *FileProcessingErrorType {
	p := new(FileProcessingErrorType)
	*p = x
	return p
}

func (x FileProcessingErrorType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FileProcessingErrorType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_devtools_resultstore_v2_file_processing_error_proto_enumTypes[0].Descriptor()
}

func (FileProcessingErrorType) Type() protoreflect.EnumType {
	return &file_google_devtools_resultstore_v2_file_processing_error_proto_enumTypes[0]
}

func (x FileProcessingErrorType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FileProcessingErrorType.Descriptor instead.
func (FileProcessingErrorType) EnumDescriptor() ([]byte, []int) {
	return file_google_devtools_resultstore_v2_file_processing_error_proto_rawDescGZIP(), []int{0}
}

// Stores errors reading or parsing a file during post-processing.
type FileProcessingErrors struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The uid of the File being read or parsed.
	FileUid string `protobuf:"bytes,1,opt,name=file_uid,json=fileUid,proto3" json:"file_uid,omitempty"`
	// What went wrong.
	FileProcessingErrors []*FileProcessingError `protobuf:"bytes,3,rep,name=file_processing_errors,json=fileProcessingErrors,proto3" json:"file_processing_errors,omitempty"`
}

func (x *FileProcessingErrors) Reset() {
	*x = FileProcessingErrors{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_resultstore_v2_file_processing_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileProcessingErrors) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileProcessingErrors) ProtoMessage() {}

func (x *FileProcessingErrors) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_resultstore_v2_file_processing_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileProcessingErrors.ProtoReflect.Descriptor instead.
func (*FileProcessingErrors) Descriptor() ([]byte, []int) {
	return file_google_devtools_resultstore_v2_file_processing_error_proto_rawDescGZIP(), []int{0}
}

func (x *FileProcessingErrors) GetFileUid() string {
	if x != nil {
		return x.FileUid
	}
	return ""
}

func (x *FileProcessingErrors) GetFileProcessingErrors() []*FileProcessingError {
	if x != nil {
		return x.FileProcessingErrors
	}
	return nil
}

// Stores an error reading or parsing a file during post-processing.
type FileProcessingError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of error that occurred.
	Type FileProcessingErrorType `protobuf:"varint,1,opt,name=type,proto3,enum=google.devtools.resultstore.v2.FileProcessingErrorType" json:"type,omitempty"`
	// Error message describing the problem.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *FileProcessingError) Reset() {
	*x = FileProcessingError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_resultstore_v2_file_processing_error_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileProcessingError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileProcessingError) ProtoMessage() {}

func (x *FileProcessingError) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_resultstore_v2_file_processing_error_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileProcessingError.ProtoReflect.Descriptor instead.
func (*FileProcessingError) Descriptor() ([]byte, []int) {
	return file_google_devtools_resultstore_v2_file_processing_error_proto_rawDescGZIP(), []int{1}
}

func (x *FileProcessingError) GetType() FileProcessingErrorType {
	if x != nil {
		return x.Type
	}
	return FileProcessingErrorType_FILE_PROCESSING_ERROR_TYPE_UNSPECIFIED
}

func (x *FileProcessingError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_google_devtools_resultstore_v2_file_processing_error_proto protoreflect.FileDescriptor

var file_google_devtools_resultstore_v2_file_processing_error_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x32,
	0x2f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67,
	0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x22, 0x9c, 0x01, 0x0a,
	0x14, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x55, 0x69, 0x64,
	0x12, 0x69, 0x0a, 0x16, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x69, 0x6e, 0x67, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f,
	0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x32, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x14, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x69, 0x6e, 0x67, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0x7c, 0x0a, 0x13, 0x46,
	0x69, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x4b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f,
	0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x32, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0xe9, 0x01, 0x0a, 0x17, 0x46, 0x69,
	0x6c, 0x65, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x26, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x50, 0x52,
	0x4f, 0x43, 0x45, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x16, 0x0a, 0x12, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x49, 0x43, 0x5f, 0x52, 0x45, 0x41,
	0x44, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x47, 0x45, 0x4e,
	0x45, 0x52, 0x49, 0x43, 0x5f, 0x50, 0x41, 0x52, 0x53, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x4c,
	0x41, 0x52, 0x47, 0x45, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x4f, 0x55, 0x54, 0x50, 0x55, 0x54,
	0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x4c, 0x41, 0x52, 0x47, 0x45, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d,
	0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x44, 0x45, 0x4e, 0x49, 0x45, 0x44, 0x10, 0x05, 0x12,
	0x15, 0x0a, 0x11, 0x44, 0x45, 0x41, 0x44, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x45, 0x58, 0x43, 0x45,
	0x45, 0x44, 0x45, 0x44, 0x10, 0x06, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f,
	0x55, 0x4e, 0x44, 0x10, 0x07, 0x12, 0x0e, 0x0a, 0x0a, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x45, 0x4d,
	0x50, 0x54, 0x59, 0x10, 0x08, 0x42, 0x71, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x50, 0x01, 0x5a, 0x49, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x32, 0x3b, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_devtools_resultstore_v2_file_processing_error_proto_rawDescOnce sync.Once
	file_google_devtools_resultstore_v2_file_processing_error_proto_rawDescData = file_google_devtools_resultstore_v2_file_processing_error_proto_rawDesc
)

func file_google_devtools_resultstore_v2_file_processing_error_proto_rawDescGZIP() []byte {
	file_google_devtools_resultstore_v2_file_processing_error_proto_rawDescOnce.Do(func() {
		file_google_devtools_resultstore_v2_file_processing_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_devtools_resultstore_v2_file_processing_error_proto_rawDescData)
	})
	return file_google_devtools_resultstore_v2_file_processing_error_proto_rawDescData
}

var file_google_devtools_resultstore_v2_file_processing_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_devtools_resultstore_v2_file_processing_error_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_devtools_resultstore_v2_file_processing_error_proto_goTypes = []interface{}{
	(FileProcessingErrorType)(0), // 0: google.devtools.resultstore.v2.FileProcessingErrorType
	(*FileProcessingErrors)(nil), // 1: google.devtools.resultstore.v2.FileProcessingErrors
	(*FileProcessingError)(nil),  // 2: google.devtools.resultstore.v2.FileProcessingError
}
var file_google_devtools_resultstore_v2_file_processing_error_proto_depIdxs = []int32{
	2, // 0: google.devtools.resultstore.v2.FileProcessingErrors.file_processing_errors:type_name -> google.devtools.resultstore.v2.FileProcessingError
	0, // 1: google.devtools.resultstore.v2.FileProcessingError.type:type_name -> google.devtools.resultstore.v2.FileProcessingErrorType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_devtools_resultstore_v2_file_processing_error_proto_init() }
func file_google_devtools_resultstore_v2_file_processing_error_proto_init() {
	if File_google_devtools_resultstore_v2_file_processing_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_devtools_resultstore_v2_file_processing_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileProcessingErrors); i {
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
		file_google_devtools_resultstore_v2_file_processing_error_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileProcessingError); i {
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
			RawDescriptor: file_google_devtools_resultstore_v2_file_processing_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_devtools_resultstore_v2_file_processing_error_proto_goTypes,
		DependencyIndexes: file_google_devtools_resultstore_v2_file_processing_error_proto_depIdxs,
		EnumInfos:         file_google_devtools_resultstore_v2_file_processing_error_proto_enumTypes,
		MessageInfos:      file_google_devtools_resultstore_v2_file_processing_error_proto_msgTypes,
	}.Build()
	File_google_devtools_resultstore_v2_file_processing_error_proto = out.File
	file_google_devtools_resultstore_v2_file_processing_error_proto_rawDesc = nil
	file_google_devtools_resultstore_v2_file_processing_error_proto_goTypes = nil
	file_google_devtools_resultstore_v2_file_processing_error_proto_depIdxs = nil
}
