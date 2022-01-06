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
// source: google/cloud/integrations/v1alpha/product.proto

package integrations

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

// Enum Product.
type Product int32

const (
	// Default value.
	Product_PRODUCT_UNSPECIFIED Product = 0
	// Integration Platform.
	Product_IP Product = 1
	// Apigee.
	Product_APIGEE Product = 2
	// Security Command Center.
	Product_SECURITY Product = 3
)

// Enum value maps for Product.
var (
	Product_name = map[int32]string{
		0: "PRODUCT_UNSPECIFIED",
		1: "IP",
		2: "APIGEE",
		3: "SECURITY",
	}
	Product_value = map[string]int32{
		"PRODUCT_UNSPECIFIED": 0,
		"IP":                  1,
		"APIGEE":              2,
		"SECURITY":            3,
	}
)

func (x Product) Enum() *Product {
	p := new(Product)
	*p = x
	return p
}

func (x Product) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Product) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_integrations_v1alpha_product_proto_enumTypes[0].Descriptor()
}

func (Product) Type() protoreflect.EnumType {
	return &file_google_cloud_integrations_v1alpha_product_proto_enumTypes[0]
}

func (x Product) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Product.Descriptor instead.
func (Product) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_integrations_v1alpha_product_proto_rawDescGZIP(), []int{0}
}

var File_google_cloud_integrations_v1alpha_product_proto protoreflect.FileDescriptor

var file_google_cloud_integrations_v1alpha_product_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2a, 0x44, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12,
	0x17, 0x0a, 0x13, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x49, 0x50, 0x10, 0x01,
	0x12, 0x0a, 0x0a, 0x06, 0x41, 0x50, 0x49, 0x47, 0x45, 0x45, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08,
	0x53, 0x45, 0x43, 0x55, 0x52, 0x49, 0x54, 0x59, 0x10, 0x03, 0x42, 0x86, 0x01, 0x0a, 0x25, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x42, 0x0c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_integrations_v1alpha_product_proto_rawDescOnce sync.Once
	file_google_cloud_integrations_v1alpha_product_proto_rawDescData = file_google_cloud_integrations_v1alpha_product_proto_rawDesc
)

func file_google_cloud_integrations_v1alpha_product_proto_rawDescGZIP() []byte {
	file_google_cloud_integrations_v1alpha_product_proto_rawDescOnce.Do(func() {
		file_google_cloud_integrations_v1alpha_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_integrations_v1alpha_product_proto_rawDescData)
	})
	return file_google_cloud_integrations_v1alpha_product_proto_rawDescData
}

var file_google_cloud_integrations_v1alpha_product_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_integrations_v1alpha_product_proto_goTypes = []interface{}{
	(Product)(0), // 0: google.cloud.integrations.v1alpha.Product
}
var file_google_cloud_integrations_v1alpha_product_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_integrations_v1alpha_product_proto_init() }
func file_google_cloud_integrations_v1alpha_product_proto_init() {
	if File_google_cloud_integrations_v1alpha_product_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_integrations_v1alpha_product_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_integrations_v1alpha_product_proto_goTypes,
		DependencyIndexes: file_google_cloud_integrations_v1alpha_product_proto_depIdxs,
		EnumInfos:         file_google_cloud_integrations_v1alpha_product_proto_enumTypes,
	}.Build()
	File_google_cloud_integrations_v1alpha_product_proto = out.File
	file_google_cloud_integrations_v1alpha_product_proto_rawDesc = nil
	file_google_cloud_integrations_v1alpha_product_proto_goTypes = nil
	file_google_cloud_integrations_v1alpha_product_proto_depIdxs = nil
}
