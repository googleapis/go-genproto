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
// source: google/cloud/dataqna/v1alpha/auto_suggestion_service.proto

package dataqna

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

// The type of suggestion.
type SuggestionType int32

const (
	// No suggestiont type is specified.
	SuggestionType_SUGGESTION_TYPE_UNSPECIFIED SuggestionType = 0
	// Entity suggestion type. Suggestions are for single entities.
	SuggestionType_ENTITY SuggestionType = 1
	// Template suggestion type. Suggestions are for full sentences.
	SuggestionType_TEMPLATE SuggestionType = 2
)

// Enum value maps for SuggestionType.
var (
	SuggestionType_name = map[int32]string{
		0: "SUGGESTION_TYPE_UNSPECIFIED",
		1: "ENTITY",
		2: "TEMPLATE",
	}
	SuggestionType_value = map[string]int32{
		"SUGGESTION_TYPE_UNSPECIFIED": 0,
		"ENTITY":                      1,
		"TEMPLATE":                    2,
	}
)

func (x SuggestionType) Enum() *SuggestionType {
	p := new(SuggestionType)
	*p = x
	return p
}

func (x SuggestionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SuggestionType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_enumTypes[0].Descriptor()
}

func (SuggestionType) Type() protoreflect.EnumType {
	return &file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_enumTypes[0]
}

func (x SuggestionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SuggestionType.Descriptor instead.
func (SuggestionType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_rawDescGZIP(), []int{0}
}

// Request for query suggestions.
type SuggestQueriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The parent of the suggestion query is the resource denoting the project and
	// location.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The scopes to which this search is restricted. The only supported scope
	// pattern is
	// `//bigquery.googleapis.com/projects/{GCP-PROJECT-ID}/datasets/{DATASET-ID}/tables/{TABLE-ID}`.
	Scopes []string `protobuf:"bytes,2,rep,name=scopes,proto3" json:"scopes,omitempty"`
	// User query for which to generate suggestions. If the query is empty, zero
	// state suggestions are returned. This allows UIs to display suggestions
	// right away, helping the user to get a sense of what a query might look
	// like.
	Query string `protobuf:"bytes,3,opt,name=query,proto3" json:"query,omitempty"`
	// The requested suggestion type. Multiple suggestion types can be
	// requested, but there is no guarantee that the service will return
	// suggestions for each type. Suggestions for a requested type might rank
	// lower than suggestions for other types and the service may decide to cut
	// these suggestions off.
	SuggestionTypes []SuggestionType `protobuf:"varint,4,rep,packed,name=suggestion_types,json=suggestionTypes,proto3,enum=google.cloud.dataqna.v1alpha.SuggestionType" json:"suggestion_types,omitempty"`
}

func (x *SuggestQueriesRequest) Reset() {
	*x = SuggestQueriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuggestQueriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestQueriesRequest) ProtoMessage() {}

func (x *SuggestQueriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestQueriesRequest.ProtoReflect.Descriptor instead.
func (*SuggestQueriesRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_rawDescGZIP(), []int{0}
}

func (x *SuggestQueriesRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *SuggestQueriesRequest) GetScopes() []string {
	if x != nil {
		return x.Scopes
	}
	return nil
}

func (x *SuggestQueriesRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *SuggestQueriesRequest) GetSuggestionTypes() []SuggestionType {
	if x != nil {
		return x.SuggestionTypes
	}
	return nil
}

// A suggestion for a query with a ranking score.
type Suggestion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Detailed information about the suggestion.
	SuggestionInfo *SuggestionInfo `protobuf:"bytes,1,opt,name=suggestion_info,json=suggestionInfo,proto3" json:"suggestion_info,omitempty"`
	// The score of the suggestion. This can be used to define ordering in UI.
	// The score represents confidence in the suggestion where higher is better.
	// All score values must be in the range [0, 1).
	RankingScore float64 `protobuf:"fixed64,2,opt,name=ranking_score,json=rankingScore,proto3" json:"ranking_score,omitempty"`
	// The type of the suggestion.
	SuggestionType SuggestionType `protobuf:"varint,3,opt,name=suggestion_type,json=suggestionType,proto3,enum=google.cloud.dataqna.v1alpha.SuggestionType" json:"suggestion_type,omitempty"`
}

func (x *Suggestion) Reset() {
	*x = Suggestion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Suggestion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Suggestion) ProtoMessage() {}

func (x *Suggestion) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Suggestion.ProtoReflect.Descriptor instead.
func (*Suggestion) Descriptor() ([]byte, []int) {
	return file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_rawDescGZIP(), []int{1}
}

func (x *Suggestion) GetSuggestionInfo() *SuggestionInfo {
	if x != nil {
		return x.SuggestionInfo
	}
	return nil
}

func (x *Suggestion) GetRankingScore() float64 {
	if x != nil {
		return x.RankingScore
	}
	return 0
}

func (x *Suggestion) GetSuggestionType() SuggestionType {
	if x != nil {
		return x.SuggestionType
	}
	return SuggestionType_SUGGESTION_TYPE_UNSPECIFIED
}

// Detailed information about the suggestion.
type SuggestionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Annotations for the suggestion. This provides information about which part
	// of the suggestion corresponds to what semantic meaning (e.g. a metric).
	AnnotatedSuggestion *AnnotatedString `protobuf:"bytes,1,opt,name=annotated_suggestion,json=annotatedSuggestion,proto3" json:"annotated_suggestion,omitempty"`
	// Matches between user query and the annotated string.
	QueryMatches []*SuggestionInfo_MatchInfo `protobuf:"bytes,2,rep,name=query_matches,json=queryMatches,proto3" json:"query_matches,omitempty"`
}

func (x *SuggestionInfo) Reset() {
	*x = SuggestionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuggestionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestionInfo) ProtoMessage() {}

func (x *SuggestionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestionInfo.ProtoReflect.Descriptor instead.
func (*SuggestionInfo) Descriptor() ([]byte, []int) {
	return file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_rawDescGZIP(), []int{2}
}

func (x *SuggestionInfo) GetAnnotatedSuggestion() *AnnotatedString {
	if x != nil {
		return x.AnnotatedSuggestion
	}
	return nil
}

func (x *SuggestionInfo) GetQueryMatches() []*SuggestionInfo_MatchInfo {
	if x != nil {
		return x.QueryMatches
	}
	return nil
}

// Response to SuggestQueries.
type SuggestQueriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of suggestions.
	Suggestions []*Suggestion `protobuf:"bytes,1,rep,name=suggestions,proto3" json:"suggestions,omitempty"`
}

func (x *SuggestQueriesResponse) Reset() {
	*x = SuggestQueriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuggestQueriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestQueriesResponse) ProtoMessage() {}

func (x *SuggestQueriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestQueriesResponse.ProtoReflect.Descriptor instead.
func (*SuggestQueriesResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_rawDescGZIP(), []int{3}
}

func (x *SuggestQueriesResponse) GetSuggestions() []*Suggestion {
	if x != nil {
		return x.Suggestions
	}
	return nil
}

// MatchInfo describes which part of suggestion matched with data in user
// typed query. This can be used to highlight matching parts in the UI. This
// is different from the annotations provided in annotated_suggestion. The
// annotated_suggestion provides information about the semantic meaning, while
// this provides information about how it relates to the input.
//
// Example:
// user query: `top products`
//
// ```
// annotated_suggestion {
//  text_formatted = "top product_group"
//  html_formatted = "top <b>product_group</b>"
//  markups {
//   {type: TEXT, start_char_index: 0, length: 3}
//   {type: DIMENSION, start_char_index: 4, length: 13}
//  }
// }
//
// query_matches {
//  { start_char_index: 0, length: 3 }
//  { start_char_index: 4, length: 7}
// }
// ```
type SuggestionInfo_MatchInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unicode character index of the string annotation.
	StartCharIndex int32 `protobuf:"varint,1,opt,name=start_char_index,json=startCharIndex,proto3" json:"start_char_index,omitempty"`
	// Count of unicode characters of this substring.
	Length int32 `protobuf:"varint,2,opt,name=length,proto3" json:"length,omitempty"`
}

func (x *SuggestionInfo_MatchInfo) Reset() {
	*x = SuggestionInfo_MatchInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuggestionInfo_MatchInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestionInfo_MatchInfo) ProtoMessage() {}

func (x *SuggestionInfo_MatchInfo) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestionInfo_MatchInfo.ProtoReflect.Descriptor instead.
func (*SuggestionInfo_MatchInfo) Descriptor() ([]byte, []int) {
	return file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_rawDescGZIP(), []int{2, 0}
}

func (x *SuggestionInfo_MatchInfo) GetStartCharIndex() int32 {
	if x != nil {
		return x.StartCharIndex
	}
	return 0
}

func (x *SuggestionInfo_MatchInfo) GetLength() int32 {
	if x != nil {
		return x.Length
	}
	return 0
}

var File_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto protoreflect.FileDescriptor

var file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x71, 0x6e, 0x61, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61,
	0x75, 0x74, 0x6f, 0x5f, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x71,
	0x6e, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x71, 0x6e, 0x61, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xe1, 0x01, 0x0a, 0x15, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x51, 0x75,
	0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xe0, 0x41,
	0x02, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x57, 0x0a,
	0x10, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x71, 0x6e, 0x61, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x22, 0xdf, 0x01, 0x0a, 0x0a, 0x53, 0x75, 0x67, 0x67, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x55, 0x0a, 0x0f, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x71, 0x6e, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x75,
	0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x73, 0x75,
	0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x23, 0x0a, 0x0d,
	0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0c, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x12, 0x55, 0x0a, 0x0f, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x71, 0x6e,
	0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x9e, 0x02, 0x0a, 0x0e, 0x53, 0x75, 0x67,
	0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x60, 0x0a, 0x14, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x71, 0x6e, 0x61,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x65, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x13, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x65, 0x64, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5b, 0x0a,
	0x0d, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x71, 0x6e, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x1a, 0x4d, 0x0a, 0x09, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x63, 0x68, 0x61, 0x72, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x43, 0x68, 0x61, 0x72, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0x64, 0x0a, 0x16, 0x53, 0x75, 0x67,
	0x67, 0x65, 0x73, 0x74, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x71, 0x6e, 0x61, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0b, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2a,
	0x4b, 0x0a, 0x0e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1f, 0x0a, 0x1b, 0x53, 0x55, 0x47, 0x47, 0x45, 0x53, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x10, 0x01, 0x12, 0x0c,
	0x0a, 0x08, 0x54, 0x45, 0x4d, 0x50, 0x4c, 0x41, 0x54, 0x45, 0x10, 0x02, 0x32, 0xa5, 0x02, 0x0a,
	0x15, 0x41, 0x75, 0x74, 0x6f, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xbf, 0x01, 0x0a, 0x0e, 0x53, 0x75, 0x67, 0x67, 0x65,
	0x73, 0x74, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x71, 0x6e, 0x61,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74,
	0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x71, 0x6e, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x75,
	0x67, 0x67, 0x65, 0x73, 0x74, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x42, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3c, 0x22, 0x37, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x51, 0x75,
	0x65, 0x72, 0x69, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x1a, 0x4a, 0xca, 0x41, 0x16, 0x64, 0x61, 0x74,
	0x61, 0x71, 0x6e, 0x61, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77,
	0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x42, 0xe5, 0x01, 0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x71, 0x6e,
	0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x1a, 0x41, 0x75, 0x74, 0x6f, 0x53,
	0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x71, 0x6e, 0x61, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x3b, 0x64, 0x61, 0x74, 0x61, 0x71, 0x6e, 0x61, 0xaa, 0x02, 0x1c, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x51, 0x6e, 0x41, 0x2e, 0x56, 0x31, 0x41, 0x6c, 0x70, 0x68, 0x61, 0xca, 0x02, 0x1c, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x61, 0x74, 0x61, 0x51,
	0x6e, 0x41, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xea, 0x02, 0x1f, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x61, 0x74, 0x61,
	0x51, 0x6e, 0x41, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_rawDescOnce sync.Once
	file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_rawDescData = file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_rawDesc
)

func file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_rawDescGZIP() []byte {
	file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_rawDescOnce.Do(func() {
		file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_rawDescData)
	})
	return file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_rawDescData
}

var file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_goTypes = []interface{}{
	(SuggestionType)(0),              // 0: google.cloud.dataqna.v1alpha.SuggestionType
	(*SuggestQueriesRequest)(nil),    // 1: google.cloud.dataqna.v1alpha.SuggestQueriesRequest
	(*Suggestion)(nil),               // 2: google.cloud.dataqna.v1alpha.Suggestion
	(*SuggestionInfo)(nil),           // 3: google.cloud.dataqna.v1alpha.SuggestionInfo
	(*SuggestQueriesResponse)(nil),   // 4: google.cloud.dataqna.v1alpha.SuggestQueriesResponse
	(*SuggestionInfo_MatchInfo)(nil), // 5: google.cloud.dataqna.v1alpha.SuggestionInfo.MatchInfo
	(*AnnotatedString)(nil),          // 6: google.cloud.dataqna.v1alpha.AnnotatedString
}
var file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_depIdxs = []int32{
	0, // 0: google.cloud.dataqna.v1alpha.SuggestQueriesRequest.suggestion_types:type_name -> google.cloud.dataqna.v1alpha.SuggestionType
	3, // 1: google.cloud.dataqna.v1alpha.Suggestion.suggestion_info:type_name -> google.cloud.dataqna.v1alpha.SuggestionInfo
	0, // 2: google.cloud.dataqna.v1alpha.Suggestion.suggestion_type:type_name -> google.cloud.dataqna.v1alpha.SuggestionType
	6, // 3: google.cloud.dataqna.v1alpha.SuggestionInfo.annotated_suggestion:type_name -> google.cloud.dataqna.v1alpha.AnnotatedString
	5, // 4: google.cloud.dataqna.v1alpha.SuggestionInfo.query_matches:type_name -> google.cloud.dataqna.v1alpha.SuggestionInfo.MatchInfo
	2, // 5: google.cloud.dataqna.v1alpha.SuggestQueriesResponse.suggestions:type_name -> google.cloud.dataqna.v1alpha.Suggestion
	1, // 6: google.cloud.dataqna.v1alpha.AutoSuggestionService.SuggestQueries:input_type -> google.cloud.dataqna.v1alpha.SuggestQueriesRequest
	4, // 7: google.cloud.dataqna.v1alpha.AutoSuggestionService.SuggestQueries:output_type -> google.cloud.dataqna.v1alpha.SuggestQueriesResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_init() }
func file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_init() {
	if File_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto != nil {
		return
	}
	file_google_cloud_dataqna_v1alpha_annotated_string_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuggestQueriesRequest); i {
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
		file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Suggestion); i {
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
		file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuggestionInfo); i {
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
		file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuggestQueriesResponse); i {
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
		file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuggestionInfo_MatchInfo); i {
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
			RawDescriptor: file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_depIdxs,
		EnumInfos:         file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_enumTypes,
		MessageInfos:      file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_msgTypes,
	}.Build()
	File_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto = out.File
	file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_rawDesc = nil
	file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_goTypes = nil
	file_google_cloud_dataqna_v1alpha_auto_suggestion_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AutoSuggestionServiceClient is the client API for AutoSuggestionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AutoSuggestionServiceClient interface {
	// Gets a list of suggestions based on a prefix string.
	// AutoSuggestion tolerance should be less than 1 second.
	SuggestQueries(ctx context.Context, in *SuggestQueriesRequest, opts ...grpc.CallOption) (*SuggestQueriesResponse, error)
}

type autoSuggestionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAutoSuggestionServiceClient(cc grpc.ClientConnInterface) AutoSuggestionServiceClient {
	return &autoSuggestionServiceClient{cc}
}

func (c *autoSuggestionServiceClient) SuggestQueries(ctx context.Context, in *SuggestQueriesRequest, opts ...grpc.CallOption) (*SuggestQueriesResponse, error) {
	out := new(SuggestQueriesResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.dataqna.v1alpha.AutoSuggestionService/SuggestQueries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AutoSuggestionServiceServer is the server API for AutoSuggestionService service.
type AutoSuggestionServiceServer interface {
	// Gets a list of suggestions based on a prefix string.
	// AutoSuggestion tolerance should be less than 1 second.
	SuggestQueries(context.Context, *SuggestQueriesRequest) (*SuggestQueriesResponse, error)
}

// UnimplementedAutoSuggestionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAutoSuggestionServiceServer struct {
}

func (*UnimplementedAutoSuggestionServiceServer) SuggestQueries(context.Context, *SuggestQueriesRequest) (*SuggestQueriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SuggestQueries not implemented")
}

func RegisterAutoSuggestionServiceServer(s *grpc.Server, srv AutoSuggestionServiceServer) {
	s.RegisterService(&_AutoSuggestionService_serviceDesc, srv)
}

func _AutoSuggestionService_SuggestQueries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuggestQueriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoSuggestionServiceServer).SuggestQueries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dataqna.v1alpha.AutoSuggestionService/SuggestQueries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoSuggestionServiceServer).SuggestQueries(ctx, req.(*SuggestQueriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AutoSuggestionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.dataqna.v1alpha.AutoSuggestionService",
	HandlerType: (*AutoSuggestionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SuggestQueries",
			Handler:    _AutoSuggestionService_SuggestQueries_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/dataqna/v1alpha/auto_suggestion_service.proto",
}
