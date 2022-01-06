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
// source: google/cloud/recommender/logging/v1beta1/action_log.proto

package logging

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	v1beta1 "google.golang.org/genproto/googleapis/cloud/recommender/v1beta1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Log content of an action on a recommendation. This includes Mark* actions.
type ActionLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. User that executed this action. Eg, foo@gmail.com
	Actor string `protobuf:"bytes,1,opt,name=actor,proto3" json:"actor,omitempty"`
	// Required. State change that was made by the actor. Eg, SUCCEEDED.
	State v1beta1.RecommendationStateInfo_State `protobuf:"varint,2,opt,name=state,proto3,enum=google.cloud.recommender.v1beta1.RecommendationStateInfo_State" json:"state,omitempty"`
	// Optional. Metadata that was included with the action that was taken.
	StateMetadata map[string]string `protobuf:"bytes,3,rep,name=state_metadata,json=stateMetadata,proto3" json:"state_metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Required. Name of the recommendation which was acted on. Eg, :
	// 'projects/foo/locations/global/recommenders/roleReco/recommendations/r1'
	RecommendationName string `protobuf:"bytes,4,opt,name=recommendation_name,json=recommendationName,proto3" json:"recommendation_name,omitempty"`
}

func (x *ActionLog) Reset() {
	*x = ActionLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_recommender_logging_v1beta1_action_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionLog) ProtoMessage() {}

func (x *ActionLog) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_recommender_logging_v1beta1_action_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionLog.ProtoReflect.Descriptor instead.
func (*ActionLog) Descriptor() ([]byte, []int) {
	return file_google_cloud_recommender_logging_v1beta1_action_log_proto_rawDescGZIP(), []int{0}
}

func (x *ActionLog) GetActor() string {
	if x != nil {
		return x.Actor
	}
	return ""
}

func (x *ActionLog) GetState() v1beta1.RecommendationStateInfo_State {
	if x != nil {
		return x.State
	}
	return v1beta1.RecommendationStateInfo_State(0)
}

func (x *ActionLog) GetStateMetadata() map[string]string {
	if x != nil {
		return x.StateMetadata
	}
	return nil
}

func (x *ActionLog) GetRecommendationName() string {
	if x != nil {
		return x.RecommendationName
	}
	return ""
}

// Log content of an action on an insight. This includes Mark* actions.
type InsightActionLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. User that executed this action. Eg, foo@gmail.com
	Actor string `protobuf:"bytes,1,opt,name=actor,proto3" json:"actor,omitempty"`
	// Required. State change that was made by the actor. Eg, ACCEPTED.
	State v1beta1.InsightStateInfo_State `protobuf:"varint,2,opt,name=state,proto3,enum=google.cloud.recommender.v1beta1.InsightStateInfo_State" json:"state,omitempty"`
	// Optional. Metadata that was included with the action that was taken.
	StateMetadata map[string]string `protobuf:"bytes,3,rep,name=state_metadata,json=stateMetadata,proto3" json:"state_metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Required. Name of the insight which was acted on. Eg, :
	// 'projects/foo/locations/global/insightTypes/roleInsight/insights/i1'
	Insight string `protobuf:"bytes,4,opt,name=insight,proto3" json:"insight,omitempty"`
}

func (x *InsightActionLog) Reset() {
	*x = InsightActionLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_recommender_logging_v1beta1_action_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsightActionLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsightActionLog) ProtoMessage() {}

func (x *InsightActionLog) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_recommender_logging_v1beta1_action_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsightActionLog.ProtoReflect.Descriptor instead.
func (*InsightActionLog) Descriptor() ([]byte, []int) {
	return file_google_cloud_recommender_logging_v1beta1_action_log_proto_rawDescGZIP(), []int{1}
}

func (x *InsightActionLog) GetActor() string {
	if x != nil {
		return x.Actor
	}
	return ""
}

func (x *InsightActionLog) GetState() v1beta1.InsightStateInfo_State {
	if x != nil {
		return x.State
	}
	return v1beta1.InsightStateInfo_State(0)
}

func (x *InsightActionLog) GetStateMetadata() map[string]string {
	if x != nil {
		return x.StateMetadata
	}
	return nil
}

func (x *InsightActionLog) GetInsight() string {
	if x != nil {
		return x.Insight
	}
	return ""
}

var File_google_cloud_recommender_logging_v1beta1_action_log_proto protoreflect.FileDescriptor

var file_google_cloud_recommender_logging_v1beta1_action_log_proto_rawDesc = []byte{
	0x0a, 0x39, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x69,
	0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda, 0x02,
	0x0a, 0x09, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x12, 0x55, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x6d, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x46, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x6c, 0x6f, 0x67, 0x67,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x65, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2f, 0x0a, 0x13, 0x72, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x40, 0x0a, 0x12, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xde, 0x02, 0x0a, 0x10, 0x49,
	0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x12,
	0x19, 0x0a, 0x05, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x05, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x53, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x49, 0x6e, 0x73,
	0x69, 0x67, 0x68, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x79, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x6f, 0x67, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0d, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x07, 0x69, 0x6e,
	0x73, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x07, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x1a, 0x40, 0x0a, 0x12, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x91, 0x01, 0x0a, 0x2c,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x6c, 0x6f, 0x67,
	0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0e, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_recommender_logging_v1beta1_action_log_proto_rawDescOnce sync.Once
	file_google_cloud_recommender_logging_v1beta1_action_log_proto_rawDescData = file_google_cloud_recommender_logging_v1beta1_action_log_proto_rawDesc
)

func file_google_cloud_recommender_logging_v1beta1_action_log_proto_rawDescGZIP() []byte {
	file_google_cloud_recommender_logging_v1beta1_action_log_proto_rawDescOnce.Do(func() {
		file_google_cloud_recommender_logging_v1beta1_action_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_recommender_logging_v1beta1_action_log_proto_rawDescData)
	})
	return file_google_cloud_recommender_logging_v1beta1_action_log_proto_rawDescData
}

var file_google_cloud_recommender_logging_v1beta1_action_log_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_recommender_logging_v1beta1_action_log_proto_goTypes = []interface{}{
	(*ActionLog)(nil),        // 0: google.cloud.recommender.logging.v1beta1.ActionLog
	(*InsightActionLog)(nil), // 1: google.cloud.recommender.logging.v1beta1.InsightActionLog
	nil,                      // 2: google.cloud.recommender.logging.v1beta1.ActionLog.StateMetadataEntry
	nil,                      // 3: google.cloud.recommender.logging.v1beta1.InsightActionLog.StateMetadataEntry
	(v1beta1.RecommendationStateInfo_State)(0), // 4: google.cloud.recommender.v1beta1.RecommendationStateInfo.State
	(v1beta1.InsightStateInfo_State)(0),        // 5: google.cloud.recommender.v1beta1.InsightStateInfo.State
}
var file_google_cloud_recommender_logging_v1beta1_action_log_proto_depIdxs = []int32{
	4, // 0: google.cloud.recommender.logging.v1beta1.ActionLog.state:type_name -> google.cloud.recommender.v1beta1.RecommendationStateInfo.State
	2, // 1: google.cloud.recommender.logging.v1beta1.ActionLog.state_metadata:type_name -> google.cloud.recommender.logging.v1beta1.ActionLog.StateMetadataEntry
	5, // 2: google.cloud.recommender.logging.v1beta1.InsightActionLog.state:type_name -> google.cloud.recommender.v1beta1.InsightStateInfo.State
	3, // 3: google.cloud.recommender.logging.v1beta1.InsightActionLog.state_metadata:type_name -> google.cloud.recommender.logging.v1beta1.InsightActionLog.StateMetadataEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_recommender_logging_v1beta1_action_log_proto_init() }
func file_google_cloud_recommender_logging_v1beta1_action_log_proto_init() {
	if File_google_cloud_recommender_logging_v1beta1_action_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_recommender_logging_v1beta1_action_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionLog); i {
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
		file_google_cloud_recommender_logging_v1beta1_action_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsightActionLog); i {
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
			RawDescriptor: file_google_cloud_recommender_logging_v1beta1_action_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_recommender_logging_v1beta1_action_log_proto_goTypes,
		DependencyIndexes: file_google_cloud_recommender_logging_v1beta1_action_log_proto_depIdxs,
		MessageInfos:      file_google_cloud_recommender_logging_v1beta1_action_log_proto_msgTypes,
	}.Build()
	File_google_cloud_recommender_logging_v1beta1_action_log_proto = out.File
	file_google_cloud_recommender_logging_v1beta1_action_log_proto_rawDesc = nil
	file_google_cloud_recommender_logging_v1beta1_action_log_proto_goTypes = nil
	file_google_cloud_recommender_logging_v1beta1_action_log_proto_depIdxs = nil
}
