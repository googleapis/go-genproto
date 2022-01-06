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
// source: google/cloud/integrations/v1alpha/value_type.proto

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

// The type of the parameter.
type ValueType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//	*ValueType_StringValue
	//	*ValueType_IntValue
	//	*ValueType_DoubleValue
	//	*ValueType_BooleanValue
	//	*ValueType_StringArray
	//	*ValueType_IntArray
	//	*ValueType_DoubleArray
	//	*ValueType_BooleanArray
	//	*ValueType_JsonValue
	Value isValueType_Value `protobuf_oneof:"value"`
}

func (x *ValueType) Reset() {
	*x = ValueType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_integrations_v1alpha_value_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValueType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValueType) ProtoMessage() {}

func (x *ValueType) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_integrations_v1alpha_value_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValueType.ProtoReflect.Descriptor instead.
func (*ValueType) Descriptor() ([]byte, []int) {
	return file_google_cloud_integrations_v1alpha_value_type_proto_rawDescGZIP(), []int{0}
}

func (m *ValueType) GetValue() isValueType_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *ValueType) GetStringValue() string {
	if x, ok := x.GetValue().(*ValueType_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (x *ValueType) GetIntValue() int64 {
	if x, ok := x.GetValue().(*ValueType_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (x *ValueType) GetDoubleValue() float64 {
	if x, ok := x.GetValue().(*ValueType_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (x *ValueType) GetBooleanValue() bool {
	if x, ok := x.GetValue().(*ValueType_BooleanValue); ok {
		return x.BooleanValue
	}
	return false
}

func (x *ValueType) GetStringArray() *StringParameterArray {
	if x, ok := x.GetValue().(*ValueType_StringArray); ok {
		return x.StringArray
	}
	return nil
}

func (x *ValueType) GetIntArray() *IntParameterArray {
	if x, ok := x.GetValue().(*ValueType_IntArray); ok {
		return x.IntArray
	}
	return nil
}

func (x *ValueType) GetDoubleArray() *DoubleParameterArray {
	if x, ok := x.GetValue().(*ValueType_DoubleArray); ok {
		return x.DoubleArray
	}
	return nil
}

func (x *ValueType) GetBooleanArray() *BooleanParameterArray {
	if x, ok := x.GetValue().(*ValueType_BooleanArray); ok {
		return x.BooleanArray
	}
	return nil
}

func (x *ValueType) GetJsonValue() string {
	if x, ok := x.GetValue().(*ValueType_JsonValue); ok {
		return x.JsonValue
	}
	return ""
}

type isValueType_Value interface {
	isValueType_Value()
}

type ValueType_StringValue struct {
	// String.
	StringValue string `protobuf:"bytes,1,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type ValueType_IntValue struct {
	// Integer.
	IntValue int64 `protobuf:"varint,2,opt,name=int_value,json=intValue,proto3,oneof"`
}

type ValueType_DoubleValue struct {
	// Double Number.
	DoubleValue float64 `protobuf:"fixed64,3,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type ValueType_BooleanValue struct {
	// Boolean.
	BooleanValue bool `protobuf:"varint,4,opt,name=boolean_value,json=booleanValue,proto3,oneof"`
}

type ValueType_StringArray struct {
	// String Array.
	StringArray *StringParameterArray `protobuf:"bytes,5,opt,name=string_array,json=stringArray,proto3,oneof"`
}

type ValueType_IntArray struct {
	// Integer Array.
	IntArray *IntParameterArray `protobuf:"bytes,6,opt,name=int_array,json=intArray,proto3,oneof"`
}

type ValueType_DoubleArray struct {
	// Double Number Array.
	DoubleArray *DoubleParameterArray `protobuf:"bytes,7,opt,name=double_array,json=doubleArray,proto3,oneof"`
}

type ValueType_BooleanArray struct {
	// Boolean Array.
	BooleanArray *BooleanParameterArray `protobuf:"bytes,8,opt,name=boolean_array,json=booleanArray,proto3,oneof"`
}

type ValueType_JsonValue struct {
	// Json.
	JsonValue string `protobuf:"bytes,9,opt,name=json_value,json=jsonValue,proto3,oneof"`
}

func (*ValueType_StringValue) isValueType_Value() {}

func (*ValueType_IntValue) isValueType_Value() {}

func (*ValueType_DoubleValue) isValueType_Value() {}

func (*ValueType_BooleanValue) isValueType_Value() {}

func (*ValueType_StringArray) isValueType_Value() {}

func (*ValueType_IntArray) isValueType_Value() {}

func (*ValueType_DoubleArray) isValueType_Value() {}

func (*ValueType_BooleanArray) isValueType_Value() {}

func (*ValueType_JsonValue) isValueType_Value() {}

// This message only contains a field of string array.
type StringParameterArray struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// String array.
	StringValues []string `protobuf:"bytes,1,rep,name=string_values,json=stringValues,proto3" json:"string_values,omitempty"`
}

func (x *StringParameterArray) Reset() {
	*x = StringParameterArray{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_integrations_v1alpha_value_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringParameterArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringParameterArray) ProtoMessage() {}

func (x *StringParameterArray) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_integrations_v1alpha_value_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringParameterArray.ProtoReflect.Descriptor instead.
func (*StringParameterArray) Descriptor() ([]byte, []int) {
	return file_google_cloud_integrations_v1alpha_value_type_proto_rawDescGZIP(), []int{1}
}

func (x *StringParameterArray) GetStringValues() []string {
	if x != nil {
		return x.StringValues
	}
	return nil
}

// This message only contains a field of integer array.
type IntParameterArray struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Integer array.
	IntValues []int64 `protobuf:"varint,1,rep,packed,name=int_values,json=intValues,proto3" json:"int_values,omitempty"`
}

func (x *IntParameterArray) Reset() {
	*x = IntParameterArray{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_integrations_v1alpha_value_type_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntParameterArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntParameterArray) ProtoMessage() {}

func (x *IntParameterArray) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_integrations_v1alpha_value_type_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntParameterArray.ProtoReflect.Descriptor instead.
func (*IntParameterArray) Descriptor() ([]byte, []int) {
	return file_google_cloud_integrations_v1alpha_value_type_proto_rawDescGZIP(), []int{2}
}

func (x *IntParameterArray) GetIntValues() []int64 {
	if x != nil {
		return x.IntValues
	}
	return nil
}

// This message only contains a field of double number array.
type DoubleParameterArray struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Double number array.
	DoubleValues []float64 `protobuf:"fixed64,1,rep,packed,name=double_values,json=doubleValues,proto3" json:"double_values,omitempty"`
}

func (x *DoubleParameterArray) Reset() {
	*x = DoubleParameterArray{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_integrations_v1alpha_value_type_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoubleParameterArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoubleParameterArray) ProtoMessage() {}

func (x *DoubleParameterArray) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_integrations_v1alpha_value_type_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoubleParameterArray.ProtoReflect.Descriptor instead.
func (*DoubleParameterArray) Descriptor() ([]byte, []int) {
	return file_google_cloud_integrations_v1alpha_value_type_proto_rawDescGZIP(), []int{3}
}

func (x *DoubleParameterArray) GetDoubleValues() []float64 {
	if x != nil {
		return x.DoubleValues
	}
	return nil
}

// This message only contains a field of boolean array.
type BooleanParameterArray struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Boolean array.
	BooleanValues []bool `protobuf:"varint,1,rep,packed,name=boolean_values,json=booleanValues,proto3" json:"boolean_values,omitempty"`
}

func (x *BooleanParameterArray) Reset() {
	*x = BooleanParameterArray{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_integrations_v1alpha_value_type_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BooleanParameterArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BooleanParameterArray) ProtoMessage() {}

func (x *BooleanParameterArray) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_integrations_v1alpha_value_type_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BooleanParameterArray.ProtoReflect.Descriptor instead.
func (*BooleanParameterArray) Descriptor() ([]byte, []int) {
	return file_google_cloud_integrations_v1alpha_value_type_proto_rawDescGZIP(), []int{4}
}

func (x *BooleanParameterArray) GetBooleanValues() []bool {
	if x != nil {
		return x.BooleanValues
	}
	return nil
}

var File_google_cloud_integrations_v1alpha_value_type_proto protoreflect.FileDescriptor

var file_google_cloud_integrations_v1alpha_value_type_proto_rawDesc = []byte{
	0x0a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x22, 0xb7, 0x04, 0x0a, 0x09, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x69, 0x6e,
	0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52,
	0x08, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x64, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x48,
	0x00, 0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x25,
	0x0a, 0x0d, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0c, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x5c, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f,
	0x61, 0x72, 0x72, 0x61, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x41,
	0x72, 0x72, 0x61, 0x79, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x41, 0x72,
	0x72, 0x61, 0x79, 0x12, 0x53, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x5f, 0x61, 0x72, 0x72, 0x61, 0x79,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x49, 0x6e, 0x74, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x41, 0x72, 0x72, 0x61, 0x79, 0x48, 0x00, 0x52, 0x08,
	0x69, 0x6e, 0x74, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x5c, 0x0a, 0x0c, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x5f, 0x61, 0x72, 0x72, 0x61, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x41, 0x72, 0x72, 0x61, 0x79, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x5f, 0x0a, 0x0d, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61,
	0x6e, 0x5f, 0x61, 0x72, 0x72, 0x61, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x41, 0x72, 0x72, 0x61, 0x79, 0x48, 0x00, 0x52, 0x0c, 0x62, 0x6f, 0x6f, 0x6c, 0x65,
	0x61, 0x6e, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x1f, 0x0a, 0x0a, 0x6a, 0x73, 0x6f, 0x6e, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x6a,
	0x73, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x3b, 0x0a, 0x14, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x32,
	0x0a, 0x11, 0x49, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x41, 0x72,
	0x72, 0x61, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x09, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x22, 0x3b, 0x0a, 0x14, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x01, 0x52, 0x0c, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22,
	0x3e, 0x0a, 0x15, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x6f, 0x6f, 0x6c,
	0x65, 0x61, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x08,
	0x52, 0x0d, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x42,
	0x88, 0x01, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x0e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4d, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x69, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_google_cloud_integrations_v1alpha_value_type_proto_rawDescOnce sync.Once
	file_google_cloud_integrations_v1alpha_value_type_proto_rawDescData = file_google_cloud_integrations_v1alpha_value_type_proto_rawDesc
)

func file_google_cloud_integrations_v1alpha_value_type_proto_rawDescGZIP() []byte {
	file_google_cloud_integrations_v1alpha_value_type_proto_rawDescOnce.Do(func() {
		file_google_cloud_integrations_v1alpha_value_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_integrations_v1alpha_value_type_proto_rawDescData)
	})
	return file_google_cloud_integrations_v1alpha_value_type_proto_rawDescData
}

var file_google_cloud_integrations_v1alpha_value_type_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_cloud_integrations_v1alpha_value_type_proto_goTypes = []interface{}{
	(*ValueType)(nil),             // 0: google.cloud.integrations.v1alpha.ValueType
	(*StringParameterArray)(nil),  // 1: google.cloud.integrations.v1alpha.StringParameterArray
	(*IntParameterArray)(nil),     // 2: google.cloud.integrations.v1alpha.IntParameterArray
	(*DoubleParameterArray)(nil),  // 3: google.cloud.integrations.v1alpha.DoubleParameterArray
	(*BooleanParameterArray)(nil), // 4: google.cloud.integrations.v1alpha.BooleanParameterArray
}
var file_google_cloud_integrations_v1alpha_value_type_proto_depIdxs = []int32{
	1, // 0: google.cloud.integrations.v1alpha.ValueType.string_array:type_name -> google.cloud.integrations.v1alpha.StringParameterArray
	2, // 1: google.cloud.integrations.v1alpha.ValueType.int_array:type_name -> google.cloud.integrations.v1alpha.IntParameterArray
	3, // 2: google.cloud.integrations.v1alpha.ValueType.double_array:type_name -> google.cloud.integrations.v1alpha.DoubleParameterArray
	4, // 3: google.cloud.integrations.v1alpha.ValueType.boolean_array:type_name -> google.cloud.integrations.v1alpha.BooleanParameterArray
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_integrations_v1alpha_value_type_proto_init() }
func file_google_cloud_integrations_v1alpha_value_type_proto_init() {
	if File_google_cloud_integrations_v1alpha_value_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_integrations_v1alpha_value_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValueType); i {
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
		file_google_cloud_integrations_v1alpha_value_type_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringParameterArray); i {
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
		file_google_cloud_integrations_v1alpha_value_type_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntParameterArray); i {
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
		file_google_cloud_integrations_v1alpha_value_type_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoubleParameterArray); i {
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
		file_google_cloud_integrations_v1alpha_value_type_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BooleanParameterArray); i {
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
	file_google_cloud_integrations_v1alpha_value_type_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ValueType_StringValue)(nil),
		(*ValueType_IntValue)(nil),
		(*ValueType_DoubleValue)(nil),
		(*ValueType_BooleanValue)(nil),
		(*ValueType_StringArray)(nil),
		(*ValueType_IntArray)(nil),
		(*ValueType_DoubleArray)(nil),
		(*ValueType_BooleanArray)(nil),
		(*ValueType_JsonValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_integrations_v1alpha_value_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_integrations_v1alpha_value_type_proto_goTypes,
		DependencyIndexes: file_google_cloud_integrations_v1alpha_value_type_proto_depIdxs,
		MessageInfos:      file_google_cloud_integrations_v1alpha_value_type_proto_msgTypes,
	}.Build()
	File_google_cloud_integrations_v1alpha_value_type_proto = out.File
	file_google_cloud_integrations_v1alpha_value_type_proto_rawDesc = nil
	file_google_cloud_integrations_v1alpha_value_type_proto_goTypes = nil
	file_google_cloud_integrations_v1alpha_value_type_proto_depIdxs = nil
}
