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
// source: google/cloud/aiplatform/v1beta1/schema/predict/params/video_action_recognition.proto

package params

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

// Prediction model parameters for Video Action Recognition.
type VideoActionRecognitionPredictionParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Model only returns predictions with at least this confidence score.
	// Default value is 0.0
	ConfidenceThreshold float32 `protobuf:"fixed32,1,opt,name=confidence_threshold,json=confidenceThreshold,proto3" json:"confidence_threshold,omitempty"`
	// The model only returns up to that many top, by confidence score,
	// predictions per frame of the video. If this number is very high, the
	// Model may return fewer predictions per frame. Default value is 50.
	MaxPredictions int32 `protobuf:"varint,2,opt,name=max_predictions,json=maxPredictions,proto3" json:"max_predictions,omitempty"`
}

func (x *VideoActionRecognitionPredictionParams) Reset() {
	*x = VideoActionRecognitionPredictionParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_schema_predict_params_video_action_recognition_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoActionRecognitionPredictionParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoActionRecognitionPredictionParams) ProtoMessage() {}

func (x *VideoActionRecognitionPredictionParams) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_schema_predict_params_video_action_recognition_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoActionRecognitionPredictionParams.ProtoReflect.Descriptor instead.
func (*VideoActionRecognitionPredictionParams) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_schema_predict_params_video_action_recognition_proto_rawDescGZIP(), []int{0}
}

func (x *VideoActionRecognitionPredictionParams) GetConfidenceThreshold() float32 {
	if x != nil {
		return x.ConfidenceThreshold
	}
	return 0
}

func (x *VideoActionRecognitionPredictionParams) GetMaxPredictions() int32 {
	if x != nil {
		return x.MaxPredictions
	}
	return 0
}

var File_google_cloud_aiplatform_v1beta1_schema_predict_params_video_action_recognition_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1beta1_schema_predict_params_video_action_recognition_proto_rawDesc = []byte{
	0x0a, 0x54, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74,
	0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a, 0x26,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x67,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x31, 0x0a, 0x14, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64,
	0x65, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x13, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65,
	0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x61, 0x78,
	0x5f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0e, 0x6d, 0x61, 0x78, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x42, 0xc7, 0x01, 0x0a, 0x39, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x42, 0x2b, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63,
	0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x5b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x2f, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x3b, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1beta1_schema_predict_params_video_action_recognition_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1beta1_schema_predict_params_video_action_recognition_proto_rawDescData = file_google_cloud_aiplatform_v1beta1_schema_predict_params_video_action_recognition_proto_rawDesc
)

func file_google_cloud_aiplatform_v1beta1_schema_predict_params_video_action_recognition_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1beta1_schema_predict_params_video_action_recognition_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1beta1_schema_predict_params_video_action_recognition_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1beta1_schema_predict_params_video_action_recognition_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1beta1_schema_predict_params_video_action_recognition_proto_rawDescData
}

var file_google_cloud_aiplatform_v1beta1_schema_predict_params_video_action_recognition_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_aiplatform_v1beta1_schema_predict_params_video_action_recognition_proto_goTypes = []interface{}{
	(*VideoActionRecognitionPredictionParams)(nil), // 0: google.cloud.aiplatform.v1beta1.schema.predict.params.VideoActionRecognitionPredictionParams
}
var file_google_cloud_aiplatform_v1beta1_schema_predict_params_video_action_recognition_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_google_cloud_aiplatform_v1beta1_schema_predict_params_video_action_recognition_proto_init()
}
func file_google_cloud_aiplatform_v1beta1_schema_predict_params_video_action_recognition_proto_init() {
	if File_google_cloud_aiplatform_v1beta1_schema_predict_params_video_action_recognition_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_aiplatform_v1beta1_schema_predict_params_video_action_recognition_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoActionRecognitionPredictionParams); i {
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
			RawDescriptor: file_google_cloud_aiplatform_v1beta1_schema_predict_params_video_action_recognition_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1beta1_schema_predict_params_video_action_recognition_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1beta1_schema_predict_params_video_action_recognition_proto_depIdxs,
		MessageInfos:      file_google_cloud_aiplatform_v1beta1_schema_predict_params_video_action_recognition_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1beta1_schema_predict_params_video_action_recognition_proto = out.File
	file_google_cloud_aiplatform_v1beta1_schema_predict_params_video_action_recognition_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1beta1_schema_predict_params_video_action_recognition_proto_goTypes = nil
	file_google_cloud_aiplatform_v1beta1_schema_predict_params_video_action_recognition_proto_depIdxs = nil
}
