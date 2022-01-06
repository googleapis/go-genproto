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
// source: google/cloud/vision/v1/web_detection.proto

package vision

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

// Relevant information for the image from the Internet.
type WebDetection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deduced entities from similar images on the Internet.
	WebEntities []*WebDetection_WebEntity `protobuf:"bytes,1,rep,name=web_entities,json=webEntities,proto3" json:"web_entities,omitempty"`
	// Fully matching images from the Internet.
	// Can include resized copies of the query image.
	FullMatchingImages []*WebDetection_WebImage `protobuf:"bytes,2,rep,name=full_matching_images,json=fullMatchingImages,proto3" json:"full_matching_images,omitempty"`
	// Partial matching images from the Internet.
	// Those images are similar enough to share some key-point features. For
	// example an original image will likely have partial matching for its crops.
	PartialMatchingImages []*WebDetection_WebImage `protobuf:"bytes,3,rep,name=partial_matching_images,json=partialMatchingImages,proto3" json:"partial_matching_images,omitempty"`
	// Web pages containing the matching images from the Internet.
	PagesWithMatchingImages []*WebDetection_WebPage `protobuf:"bytes,4,rep,name=pages_with_matching_images,json=pagesWithMatchingImages,proto3" json:"pages_with_matching_images,omitempty"`
	// The visually similar image results.
	VisuallySimilarImages []*WebDetection_WebImage `protobuf:"bytes,6,rep,name=visually_similar_images,json=visuallySimilarImages,proto3" json:"visually_similar_images,omitempty"`
	// The service's best guess as to the topic of the request image.
	// Inferred from similar images on the open web.
	BestGuessLabels []*WebDetection_WebLabel `protobuf:"bytes,8,rep,name=best_guess_labels,json=bestGuessLabels,proto3" json:"best_guess_labels,omitempty"`
}

func (x *WebDetection) Reset() {
	*x = WebDetection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_vision_v1_web_detection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebDetection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebDetection) ProtoMessage() {}

func (x *WebDetection) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_vision_v1_web_detection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebDetection.ProtoReflect.Descriptor instead.
func (*WebDetection) Descriptor() ([]byte, []int) {
	return file_google_cloud_vision_v1_web_detection_proto_rawDescGZIP(), []int{0}
}

func (x *WebDetection) GetWebEntities() []*WebDetection_WebEntity {
	if x != nil {
		return x.WebEntities
	}
	return nil
}

func (x *WebDetection) GetFullMatchingImages() []*WebDetection_WebImage {
	if x != nil {
		return x.FullMatchingImages
	}
	return nil
}

func (x *WebDetection) GetPartialMatchingImages() []*WebDetection_WebImage {
	if x != nil {
		return x.PartialMatchingImages
	}
	return nil
}

func (x *WebDetection) GetPagesWithMatchingImages() []*WebDetection_WebPage {
	if x != nil {
		return x.PagesWithMatchingImages
	}
	return nil
}

func (x *WebDetection) GetVisuallySimilarImages() []*WebDetection_WebImage {
	if x != nil {
		return x.VisuallySimilarImages
	}
	return nil
}

func (x *WebDetection) GetBestGuessLabels() []*WebDetection_WebLabel {
	if x != nil {
		return x.BestGuessLabels
	}
	return nil
}

// Entity deduced from similar images on the Internet.
type WebDetection_WebEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Opaque entity ID.
	EntityId string `protobuf:"bytes,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	// Overall relevancy score for the entity.
	// Not normalized and not comparable across different image queries.
	Score float32 `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
	// Canonical description of the entity, in English.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *WebDetection_WebEntity) Reset() {
	*x = WebDetection_WebEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_vision_v1_web_detection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebDetection_WebEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebDetection_WebEntity) ProtoMessage() {}

func (x *WebDetection_WebEntity) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_vision_v1_web_detection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebDetection_WebEntity.ProtoReflect.Descriptor instead.
func (*WebDetection_WebEntity) Descriptor() ([]byte, []int) {
	return file_google_cloud_vision_v1_web_detection_proto_rawDescGZIP(), []int{0, 0}
}

func (x *WebDetection_WebEntity) GetEntityId() string {
	if x != nil {
		return x.EntityId
	}
	return ""
}

func (x *WebDetection_WebEntity) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *WebDetection_WebEntity) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// Metadata for online images.
type WebDetection_WebImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The result image URL.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// (Deprecated) Overall relevancy score for the image.
	Score float32 `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *WebDetection_WebImage) Reset() {
	*x = WebDetection_WebImage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_vision_v1_web_detection_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebDetection_WebImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebDetection_WebImage) ProtoMessage() {}

func (x *WebDetection_WebImage) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_vision_v1_web_detection_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebDetection_WebImage.ProtoReflect.Descriptor instead.
func (*WebDetection_WebImage) Descriptor() ([]byte, []int) {
	return file_google_cloud_vision_v1_web_detection_proto_rawDescGZIP(), []int{0, 1}
}

func (x *WebDetection_WebImage) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *WebDetection_WebImage) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

// Metadata for web pages.
type WebDetection_WebPage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The result web page URL.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// (Deprecated) Overall relevancy score for the web page.
	Score float32 `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
	// Title for the web page, may contain HTML markups.
	PageTitle string `protobuf:"bytes,3,opt,name=page_title,json=pageTitle,proto3" json:"page_title,omitempty"`
	// Fully matching images on the page.
	// Can include resized copies of the query image.
	FullMatchingImages []*WebDetection_WebImage `protobuf:"bytes,4,rep,name=full_matching_images,json=fullMatchingImages,proto3" json:"full_matching_images,omitempty"`
	// Partial matching images on the page.
	// Those images are similar enough to share some key-point features. For
	// example an original image will likely have partial matching for its
	// crops.
	PartialMatchingImages []*WebDetection_WebImage `protobuf:"bytes,5,rep,name=partial_matching_images,json=partialMatchingImages,proto3" json:"partial_matching_images,omitempty"`
}

func (x *WebDetection_WebPage) Reset() {
	*x = WebDetection_WebPage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_vision_v1_web_detection_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebDetection_WebPage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebDetection_WebPage) ProtoMessage() {}

func (x *WebDetection_WebPage) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_vision_v1_web_detection_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebDetection_WebPage.ProtoReflect.Descriptor instead.
func (*WebDetection_WebPage) Descriptor() ([]byte, []int) {
	return file_google_cloud_vision_v1_web_detection_proto_rawDescGZIP(), []int{0, 2}
}

func (x *WebDetection_WebPage) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *WebDetection_WebPage) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *WebDetection_WebPage) GetPageTitle() string {
	if x != nil {
		return x.PageTitle
	}
	return ""
}

func (x *WebDetection_WebPage) GetFullMatchingImages() []*WebDetection_WebImage {
	if x != nil {
		return x.FullMatchingImages
	}
	return nil
}

func (x *WebDetection_WebPage) GetPartialMatchingImages() []*WebDetection_WebImage {
	if x != nil {
		return x.PartialMatchingImages
	}
	return nil
}

// Label to provide extra metadata for the web detection.
type WebDetection_WebLabel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Label for extra metadata.
	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	// The BCP-47 language code for `label`, such as "en-US" or "sr-Latn".
	// For more information, see
	// http://www.unicode.org/reports/tr35/#Unicode_locale_identifier.
	LanguageCode string `protobuf:"bytes,2,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
}

func (x *WebDetection_WebLabel) Reset() {
	*x = WebDetection_WebLabel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_vision_v1_web_detection_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebDetection_WebLabel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebDetection_WebLabel) ProtoMessage() {}

func (x *WebDetection_WebLabel) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_vision_v1_web_detection_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebDetection_WebLabel.ProtoReflect.Descriptor instead.
func (*WebDetection_WebLabel) Descriptor() ([]byte, []int) {
	return file_google_cloud_vision_v1_web_detection_proto_rawDescGZIP(), []int{0, 3}
}

func (x *WebDetection_WebLabel) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *WebDetection_WebLabel) GetLanguageCode() string {
	if x != nil {
		return x.LanguageCode
	}
	return ""
}

var File_google_cloud_vision_v1_web_detection_proto protoreflect.FileDescriptor

var file_google_cloud_vision_v1_web_detection_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x65, 0x62, 0x5f, 0x64, 0x65, 0x74,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xce, 0x08, 0x0a, 0x0c, 0x57, 0x65, 0x62, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x51, 0x0a, 0x0c, 0x77, 0x65, 0x62, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x57, 0x65, 0x62, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x57, 0x65, 0x62, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x0b, 0x77, 0x65, 0x62, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x5f, 0x0a, 0x14, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65,
	0x62, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x57, 0x65, 0x62, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x52, 0x12, 0x66, 0x75, 0x6c, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e,
	0x67, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x65, 0x0a, 0x17, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x57, 0x65, 0x62, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x57,
	0x65, 0x62, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x15, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x69,
	0x0a, 0x1a, 0x70, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x62, 0x44,
	0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x57, 0x65, 0x62, 0x50, 0x61, 0x67, 0x65,
	0x52, 0x17, 0x70, 0x61, 0x67, 0x65, 0x73, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x69, 0x6e, 0x67, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x65, 0x0a, 0x17, 0x76, 0x69, 0x73,
	0x75, 0x61, 0x6c, 0x6c, 0x79, 0x5f, 0x73, 0x69, 0x6d, 0x69, 0x6c, 0x61, 0x72, 0x5f, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x62, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x57, 0x65, 0x62, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x15, 0x76, 0x69, 0x73, 0x75, 0x61,
	0x6c, 0x6c, 0x79, 0x53, 0x69, 0x6d, 0x69, 0x6c, 0x61, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73,
	0x12, 0x59, 0x0a, 0x11, 0x62, 0x65, 0x73, 0x74, 0x5f, 0x67, 0x75, 0x65, 0x73, 0x73, 0x5f, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x62, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x57, 0x65, 0x62, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x0f, 0x62, 0x65, 0x73, 0x74,
	0x47, 0x75, 0x65, 0x73, 0x73, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x60, 0x0a, 0x09, 0x57,
	0x65, 0x62, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x32, 0x0a,
	0x08, 0x57, 0x65, 0x62, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x1a, 0x98, 0x02, 0x0a, 0x07, 0x57, 0x65, 0x62, 0x50, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x5f, 0x0a, 0x14, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x62, 0x44,
	0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x57, 0x65, 0x62, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x52, 0x12, 0x66, 0x75, 0x6c, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x65, 0x0a, 0x17, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c,
	0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x57, 0x65, 0x62, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x57, 0x65, 0x62,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x15, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x45, 0x0a, 0x08,
	0x57, 0x65, 0x62, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x23,
	0x0a, 0x0d, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x43,
	0x6f, 0x64, 0x65, 0x42, 0x79, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x42, 0x11, 0x57, 0x65, 0x62, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x04, 0x47, 0x43, 0x56, 0x4e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_vision_v1_web_detection_proto_rawDescOnce sync.Once
	file_google_cloud_vision_v1_web_detection_proto_rawDescData = file_google_cloud_vision_v1_web_detection_proto_rawDesc
)

func file_google_cloud_vision_v1_web_detection_proto_rawDescGZIP() []byte {
	file_google_cloud_vision_v1_web_detection_proto_rawDescOnce.Do(func() {
		file_google_cloud_vision_v1_web_detection_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_vision_v1_web_detection_proto_rawDescData)
	})
	return file_google_cloud_vision_v1_web_detection_proto_rawDescData
}

var file_google_cloud_vision_v1_web_detection_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_cloud_vision_v1_web_detection_proto_goTypes = []interface{}{
	(*WebDetection)(nil),           // 0: google.cloud.vision.v1.WebDetection
	(*WebDetection_WebEntity)(nil), // 1: google.cloud.vision.v1.WebDetection.WebEntity
	(*WebDetection_WebImage)(nil),  // 2: google.cloud.vision.v1.WebDetection.WebImage
	(*WebDetection_WebPage)(nil),   // 3: google.cloud.vision.v1.WebDetection.WebPage
	(*WebDetection_WebLabel)(nil),  // 4: google.cloud.vision.v1.WebDetection.WebLabel
}
var file_google_cloud_vision_v1_web_detection_proto_depIdxs = []int32{
	1, // 0: google.cloud.vision.v1.WebDetection.web_entities:type_name -> google.cloud.vision.v1.WebDetection.WebEntity
	2, // 1: google.cloud.vision.v1.WebDetection.full_matching_images:type_name -> google.cloud.vision.v1.WebDetection.WebImage
	2, // 2: google.cloud.vision.v1.WebDetection.partial_matching_images:type_name -> google.cloud.vision.v1.WebDetection.WebImage
	3, // 3: google.cloud.vision.v1.WebDetection.pages_with_matching_images:type_name -> google.cloud.vision.v1.WebDetection.WebPage
	2, // 4: google.cloud.vision.v1.WebDetection.visually_similar_images:type_name -> google.cloud.vision.v1.WebDetection.WebImage
	4, // 5: google.cloud.vision.v1.WebDetection.best_guess_labels:type_name -> google.cloud.vision.v1.WebDetection.WebLabel
	2, // 6: google.cloud.vision.v1.WebDetection.WebPage.full_matching_images:type_name -> google.cloud.vision.v1.WebDetection.WebImage
	2, // 7: google.cloud.vision.v1.WebDetection.WebPage.partial_matching_images:type_name -> google.cloud.vision.v1.WebDetection.WebImage
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_google_cloud_vision_v1_web_detection_proto_init() }
func file_google_cloud_vision_v1_web_detection_proto_init() {
	if File_google_cloud_vision_v1_web_detection_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_vision_v1_web_detection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebDetection); i {
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
		file_google_cloud_vision_v1_web_detection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebDetection_WebEntity); i {
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
		file_google_cloud_vision_v1_web_detection_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebDetection_WebImage); i {
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
		file_google_cloud_vision_v1_web_detection_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebDetection_WebPage); i {
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
		file_google_cloud_vision_v1_web_detection_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebDetection_WebLabel); i {
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
			RawDescriptor: file_google_cloud_vision_v1_web_detection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_vision_v1_web_detection_proto_goTypes,
		DependencyIndexes: file_google_cloud_vision_v1_web_detection_proto_depIdxs,
		MessageInfos:      file_google_cloud_vision_v1_web_detection_proto_msgTypes,
	}.Build()
	File_google_cloud_vision_v1_web_detection_proto = out.File
	file_google_cloud_vision_v1_web_detection_proto_rawDesc = nil
	file_google_cloud_vision_v1_web_detection_proto_goTypes = nil
	file_google_cloud_vision_v1_web_detection_proto_depIdxs = nil
}
