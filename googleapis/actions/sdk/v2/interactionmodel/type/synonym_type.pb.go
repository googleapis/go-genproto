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
// source: google/actions/sdk/v2/interactionmodel/type/synonym_type.proto

package _type

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The type of matching that entries in this type will use. This will ensure
// all of the types use the same matching method and allow variation of
// matching for synonym matching (i.e. fuzzy versus exact). If the value is
// `UNSPECIFIED` it will be defaulted to `EXACT_MATCH`.
type SynonymType_MatchType int32

const (
	// Defaults to `EXACT_MATCH`.
	SynonymType_UNSPECIFIED SynonymType_MatchType = 0
	// Looks for an exact match of the synonym or name.
	SynonymType_EXACT_MATCH SynonymType_MatchType = 1
	// Looser than `EXACT_MATCH`. Looks for similar matches as well as exact
	// matches.
	SynonymType_FUZZY_MATCH SynonymType_MatchType = 2
)

// Enum value maps for SynonymType_MatchType.
var (
	SynonymType_MatchType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "EXACT_MATCH",
		2: "FUZZY_MATCH",
	}
	SynonymType_MatchType_value = map[string]int32{
		"UNSPECIFIED": 0,
		"EXACT_MATCH": 1,
		"FUZZY_MATCH": 2,
	}
)

func (x SynonymType_MatchType) Enum() *SynonymType_MatchType {
	p := new(SynonymType_MatchType)
	*p = x
	return p
}

func (x SynonymType_MatchType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SynonymType_MatchType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_actions_sdk_v2_interactionmodel_type_synonym_type_proto_enumTypes[0].Descriptor()
}

func (SynonymType_MatchType) Type() protoreflect.EnumType {
	return &file_google_actions_sdk_v2_interactionmodel_type_synonym_type_proto_enumTypes[0]
}

func (x SynonymType_MatchType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SynonymType_MatchType.Descriptor instead.
func (SynonymType_MatchType) EnumDescriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_interactionmodel_type_synonym_type_proto_rawDescGZIP(), []int{0, 0}
}

// Type that matches text by set of synonyms.
type SynonymType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The match type for the synonym.
	MatchType SynonymType_MatchType `protobuf:"varint,1,opt,name=match_type,json=matchType,proto3,enum=google.actions.sdk.v2.interactionmodel.type.SynonymType_MatchType" json:"match_type,omitempty"`
	// Optional. When set to true this will match unknown words or phrases based on
	// surrounding input and intent training data, such as items that might be
	// added to a grocery list.
	AcceptUnknownValues bool `protobuf:"varint,3,opt,name=accept_unknown_values,json=acceptUnknownValues,proto3" json:"accept_unknown_values,omitempty"`
	// Required. Named map of synonym entities.
	Entities map[string]*SynonymType_Entity `protobuf:"bytes,2,rep,name=entities,proto3" json:"entities,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SynonymType) Reset() {
	*x = SynonymType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_interactionmodel_type_synonym_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SynonymType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SynonymType) ProtoMessage() {}

func (x *SynonymType) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_interactionmodel_type_synonym_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SynonymType.ProtoReflect.Descriptor instead.
func (*SynonymType) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_interactionmodel_type_synonym_type_proto_rawDescGZIP(), []int{0}
}

func (x *SynonymType) GetMatchType() SynonymType_MatchType {
	if x != nil {
		return x.MatchType
	}
	return SynonymType_UNSPECIFIED
}

func (x *SynonymType) GetAcceptUnknownValues() bool {
	if x != nil {
		return x.AcceptUnknownValues
	}
	return false
}

func (x *SynonymType) GetEntities() map[string]*SynonymType_Entity {
	if x != nil {
		return x.Entities
	}
	return nil
}

// Represents a synonym entity field that contains the details of a single
// entry inside the type.
type SynonymType_Entity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The entity display details.
	Display *EntityDisplay `protobuf:"bytes,1,opt,name=display,proto3" json:"display,omitempty"`
	// Optional. The list of synonyms for the entity.
	// **This field is localizable.**
	Synonyms []string `protobuf:"bytes,2,rep,name=synonyms,proto3" json:"synonyms,omitempty"`
}

func (x *SynonymType_Entity) Reset() {
	*x = SynonymType_Entity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_interactionmodel_type_synonym_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SynonymType_Entity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SynonymType_Entity) ProtoMessage() {}

func (x *SynonymType_Entity) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_interactionmodel_type_synonym_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SynonymType_Entity.ProtoReflect.Descriptor instead.
func (*SynonymType_Entity) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_interactionmodel_type_synonym_type_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SynonymType_Entity) GetDisplay() *EntityDisplay {
	if x != nil {
		return x.Display
	}
	return nil
}

func (x *SynonymType_Entity) GetSynonyms() []string {
	if x != nil {
		return x.Synonyms
	}
	return nil
}

var File_google_actions_sdk_v2_interactionmodel_type_synonym_type_proto protoreflect.FileDescriptor

var file_google_actions_sdk_v2_interactionmodel_type_synonym_type_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x73, 0x79,
	0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x1a, 0x40, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64,
	0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xdc, 0x04, 0x0a, 0x0b, 0x53, 0x79, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x66, 0x0a, 0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x42, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x53, 0x79, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x09, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x37, 0x0a, 0x15, 0x61, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x5f, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x13, 0x61, 0x63,
	0x63, 0x65, 0x70, 0x74, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x12, 0x67, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x46, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x53, 0x79, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x1a, 0x84, 0x01, 0x0a, 0x06, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x59, 0x0a, 0x07, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x44, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x07, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x12, 0x1f, 0x0a, 0x08, 0x73, 0x79, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x08, 0x73, 0x79, 0x6e, 0x6f, 0x6e, 0x79, 0x6d,
	0x73, 0x1a, 0x7c, 0x0a, 0x0d, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x55, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x53, 0x79, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x3e, 0x0a, 0x09, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a,
	0x0b, 0x45, 0x58, 0x41, 0x43, 0x54, 0x5f, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x10, 0x01, 0x12, 0x0f,
	0x0a, 0x0b, 0x46, 0x55, 0x5a, 0x5a, 0x59, 0x5f, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x10, 0x02, 0x42,
	0x91, 0x01, 0x0a, 0x2f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x42, 0x10, 0x53, 0x79, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x54, 0x79, 0x70, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_actions_sdk_v2_interactionmodel_type_synonym_type_proto_rawDescOnce sync.Once
	file_google_actions_sdk_v2_interactionmodel_type_synonym_type_proto_rawDescData = file_google_actions_sdk_v2_interactionmodel_type_synonym_type_proto_rawDesc
)

func file_google_actions_sdk_v2_interactionmodel_type_synonym_type_proto_rawDescGZIP() []byte {
	file_google_actions_sdk_v2_interactionmodel_type_synonym_type_proto_rawDescOnce.Do(func() {
		file_google_actions_sdk_v2_interactionmodel_type_synonym_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_actions_sdk_v2_interactionmodel_type_synonym_type_proto_rawDescData)
	})
	return file_google_actions_sdk_v2_interactionmodel_type_synonym_type_proto_rawDescData
}

var file_google_actions_sdk_v2_interactionmodel_type_synonym_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_actions_sdk_v2_interactionmodel_type_synonym_type_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_actions_sdk_v2_interactionmodel_type_synonym_type_proto_goTypes = []interface{}{
	(SynonymType_MatchType)(0), // 0: google.actions.sdk.v2.interactionmodel.type.SynonymType.MatchType
	(*SynonymType)(nil),        // 1: google.actions.sdk.v2.interactionmodel.type.SynonymType
	(*SynonymType_Entity)(nil), // 2: google.actions.sdk.v2.interactionmodel.type.SynonymType.Entity
	nil,                        // 3: google.actions.sdk.v2.interactionmodel.type.SynonymType.EntitiesEntry
	(*EntityDisplay)(nil),      // 4: google.actions.sdk.v2.interactionmodel.type.EntityDisplay
}
var file_google_actions_sdk_v2_interactionmodel_type_synonym_type_proto_depIdxs = []int32{
	0, // 0: google.actions.sdk.v2.interactionmodel.type.SynonymType.match_type:type_name -> google.actions.sdk.v2.interactionmodel.type.SynonymType.MatchType
	3, // 1: google.actions.sdk.v2.interactionmodel.type.SynonymType.entities:type_name -> google.actions.sdk.v2.interactionmodel.type.SynonymType.EntitiesEntry
	4, // 2: google.actions.sdk.v2.interactionmodel.type.SynonymType.Entity.display:type_name -> google.actions.sdk.v2.interactionmodel.type.EntityDisplay
	2, // 3: google.actions.sdk.v2.interactionmodel.type.SynonymType.EntitiesEntry.value:type_name -> google.actions.sdk.v2.interactionmodel.type.SynonymType.Entity
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_actions_sdk_v2_interactionmodel_type_synonym_type_proto_init() }
func file_google_actions_sdk_v2_interactionmodel_type_synonym_type_proto_init() {
	if File_google_actions_sdk_v2_interactionmodel_type_synonym_type_proto != nil {
		return
	}
	file_google_actions_sdk_v2_interactionmodel_type_entity_display_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_actions_sdk_v2_interactionmodel_type_synonym_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SynonymType); i {
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
		file_google_actions_sdk_v2_interactionmodel_type_synonym_type_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SynonymType_Entity); i {
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
			RawDescriptor: file_google_actions_sdk_v2_interactionmodel_type_synonym_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_actions_sdk_v2_interactionmodel_type_synonym_type_proto_goTypes,
		DependencyIndexes: file_google_actions_sdk_v2_interactionmodel_type_synonym_type_proto_depIdxs,
		EnumInfos:         file_google_actions_sdk_v2_interactionmodel_type_synonym_type_proto_enumTypes,
		MessageInfos:      file_google_actions_sdk_v2_interactionmodel_type_synonym_type_proto_msgTypes,
	}.Build()
	File_google_actions_sdk_v2_interactionmodel_type_synonym_type_proto = out.File
	file_google_actions_sdk_v2_interactionmodel_type_synonym_type_proto_rawDesc = nil
	file_google_actions_sdk_v2_interactionmodel_type_synonym_type_proto_goTypes = nil
	file_google_actions_sdk_v2_interactionmodel_type_synonym_type_proto_depIdxs = nil
}
