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
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: google/cloud/vision/v1p4beta1/product_search.proto

package vision

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Parameters for a product search request.
type ProductSearchParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The bounding polygon around the area of interest in the image.
	// If it is not specified, system discretion will be applied.
	BoundingPoly *BoundingPoly `protobuf:"bytes,9,opt,name=bounding_poly,json=boundingPoly,proto3" json:"bounding_poly,omitempty"`
	// The resource name of a
	// [ProductSet][google.cloud.vision.v1p4beta1.ProductSet] to be searched for
	// similar images.
	//
	// Format is:
	// `projects/PROJECT_ID/locations/LOC_ID/productSets/PRODUCT_SET_ID`.
	ProductSet string `protobuf:"bytes,6,opt,name=product_set,json=productSet,proto3" json:"product_set,omitempty"`
	// The list of product categories to search in. Currently, we only consider
	// the first category, and either "homegoods-v2", "apparel-v2", "toys-v2",
	// "packagedgoods-v1", or "general-v1" should be specified. The legacy
	// categories "homegoods", "apparel", and "toys" are still supported but will
	// be deprecated. For new products, please use "homegoods-v2", "apparel-v2",
	// or "toys-v2" for better product search accuracy. It is recommended to
	// migrate existing products to these categories as well.
	ProductCategories []string `protobuf:"bytes,7,rep,name=product_categories,json=productCategories,proto3" json:"product_categories,omitempty"`
	// The filtering expression. This can be used to restrict search results based
	// on Product labels. We currently support an AND of OR of key-value
	// expressions, where each expression within an OR must have the same key. An
	// '=' should be used to connect the key and value.
	//
	// For example, "(color = red OR color = blue) AND brand = Google" is
	// acceptable, but "(color = red OR brand = Google)" is not acceptable.
	// "color: red" is not acceptable because it uses a ':' instead of an '='.
	Filter string `protobuf:"bytes,8,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ProductSearchParams) Reset() {
	*x = ProductSearchParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_vision_v1p4beta1_product_search_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductSearchParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductSearchParams) ProtoMessage() {}

func (x *ProductSearchParams) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_vision_v1p4beta1_product_search_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductSearchParams.ProtoReflect.Descriptor instead.
func (*ProductSearchParams) Descriptor() ([]byte, []int) {
	return file_google_cloud_vision_v1p4beta1_product_search_proto_rawDescGZIP(), []int{0}
}

func (x *ProductSearchParams) GetBoundingPoly() *BoundingPoly {
	if x != nil {
		return x.BoundingPoly
	}
	return nil
}

func (x *ProductSearchParams) GetProductSet() string {
	if x != nil {
		return x.ProductSet
	}
	return ""
}

func (x *ProductSearchParams) GetProductCategories() []string {
	if x != nil {
		return x.ProductCategories
	}
	return nil
}

func (x *ProductSearchParams) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

// Results for a product search request.
type ProductSearchResults struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Timestamp of the index which provided these results. Products added to the
	// product set and products removed from the product set after this time are
	// not reflected in the current results.
	IndexTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=index_time,json=indexTime,proto3" json:"index_time,omitempty"`
	// List of results, one for each product match.
	Results []*ProductSearchResults_Result `protobuf:"bytes,5,rep,name=results,proto3" json:"results,omitempty"`
	// List of results grouped by products detected in the query image. Each entry
	// corresponds to one bounding polygon in the query image, and contains the
	// matching products specific to that region. There may be duplicate product
	// matches in the union of all the per-product results.
	ProductGroupedResults []*ProductSearchResults_GroupedResult `protobuf:"bytes,6,rep,name=product_grouped_results,json=productGroupedResults,proto3" json:"product_grouped_results,omitempty"`
}

func (x *ProductSearchResults) Reset() {
	*x = ProductSearchResults{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_vision_v1p4beta1_product_search_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductSearchResults) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductSearchResults) ProtoMessage() {}

func (x *ProductSearchResults) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_vision_v1p4beta1_product_search_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductSearchResults.ProtoReflect.Descriptor instead.
func (*ProductSearchResults) Descriptor() ([]byte, []int) {
	return file_google_cloud_vision_v1p4beta1_product_search_proto_rawDescGZIP(), []int{1}
}

func (x *ProductSearchResults) GetIndexTime() *timestamppb.Timestamp {
	if x != nil {
		return x.IndexTime
	}
	return nil
}

func (x *ProductSearchResults) GetResults() []*ProductSearchResults_Result {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *ProductSearchResults) GetProductGroupedResults() []*ProductSearchResults_GroupedResult {
	if x != nil {
		return x.ProductGroupedResults
	}
	return nil
}

// Information about a product.
type ProductSearchResults_Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Product.
	Product *Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	// A confidence level on the match, ranging from 0 (no confidence) to
	// 1 (full confidence).
	Score float32 `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
	// The resource name of the image from the product that is the closest match
	// to the query.
	Image string `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *ProductSearchResults_Result) Reset() {
	*x = ProductSearchResults_Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_vision_v1p4beta1_product_search_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductSearchResults_Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductSearchResults_Result) ProtoMessage() {}

func (x *ProductSearchResults_Result) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_vision_v1p4beta1_product_search_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductSearchResults_Result.ProtoReflect.Descriptor instead.
func (*ProductSearchResults_Result) Descriptor() ([]byte, []int) {
	return file_google_cloud_vision_v1p4beta1_product_search_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ProductSearchResults_Result) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

func (x *ProductSearchResults_Result) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *ProductSearchResults_Result) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

// Prediction for what the object in the bounding box is.
type ProductSearchResults_ObjectAnnotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Object ID that should align with EntityAnnotation mid.
	Mid string `protobuf:"bytes,1,opt,name=mid,proto3" json:"mid,omitempty"`
	// The BCP-47 language code, such as "en-US" or "sr-Latn". For more
	// information, see
	// http://www.unicode.org/reports/tr35/#Unicode_locale_identifier.
	LanguageCode string `protobuf:"bytes,2,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	// Object name, expressed in its `language_code` language.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Score of the result. Range [0, 1].
	Score float32 `protobuf:"fixed32,4,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *ProductSearchResults_ObjectAnnotation) Reset() {
	*x = ProductSearchResults_ObjectAnnotation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_vision_v1p4beta1_product_search_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductSearchResults_ObjectAnnotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductSearchResults_ObjectAnnotation) ProtoMessage() {}

func (x *ProductSearchResults_ObjectAnnotation) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_vision_v1p4beta1_product_search_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductSearchResults_ObjectAnnotation.ProtoReflect.Descriptor instead.
func (*ProductSearchResults_ObjectAnnotation) Descriptor() ([]byte, []int) {
	return file_google_cloud_vision_v1p4beta1_product_search_proto_rawDescGZIP(), []int{1, 1}
}

func (x *ProductSearchResults_ObjectAnnotation) GetMid() string {
	if x != nil {
		return x.Mid
	}
	return ""
}

func (x *ProductSearchResults_ObjectAnnotation) GetLanguageCode() string {
	if x != nil {
		return x.LanguageCode
	}
	return ""
}

func (x *ProductSearchResults_ObjectAnnotation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductSearchResults_ObjectAnnotation) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

// Information about the products similar to a single product in a query
// image.
type ProductSearchResults_GroupedResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The bounding polygon around the product detected in the query image.
	BoundingPoly *BoundingPoly `protobuf:"bytes,1,opt,name=bounding_poly,json=boundingPoly,proto3" json:"bounding_poly,omitempty"`
	// List of results, one for each product match.
	Results []*ProductSearchResults_Result `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	// List of generic predictions for the object in the bounding box.
	ObjectAnnotations []*ProductSearchResults_ObjectAnnotation `protobuf:"bytes,3,rep,name=object_annotations,json=objectAnnotations,proto3" json:"object_annotations,omitempty"`
}

func (x *ProductSearchResults_GroupedResult) Reset() {
	*x = ProductSearchResults_GroupedResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_vision_v1p4beta1_product_search_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductSearchResults_GroupedResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductSearchResults_GroupedResult) ProtoMessage() {}

func (x *ProductSearchResults_GroupedResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_vision_v1p4beta1_product_search_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductSearchResults_GroupedResult.ProtoReflect.Descriptor instead.
func (*ProductSearchResults_GroupedResult) Descriptor() ([]byte, []int) {
	return file_google_cloud_vision_v1p4beta1_product_search_proto_rawDescGZIP(), []int{1, 2}
}

func (x *ProductSearchResults_GroupedResult) GetBoundingPoly() *BoundingPoly {
	if x != nil {
		return x.BoundingPoly
	}
	return nil
}

func (x *ProductSearchResults_GroupedResult) GetResults() []*ProductSearchResults_Result {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *ProductSearchResults_GroupedResult) GetObjectAnnotations() []*ProductSearchResults_ObjectAnnotation {
	if x != nil {
		return x.ObjectAnnotations
	}
	return nil
}

var File_google_cloud_vision_v1p4beta1_product_search_proto protoreflect.FileDescriptor

var file_google_cloud_vision_v1p4beta1_product_search_proto_rawDesc = []byte{
	0x0a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x70, 0x34, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x70, 0x34, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x70, 0x34, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x67, 0x65, 0x6f, 0x6d,
	0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x70, 0x34, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf6, 0x01, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12,
	0x50, 0x0a, 0x0d, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x79,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x70,
	0x34, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x50,
	0x6f, 0x6c, 0x79, 0x52, 0x0c, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c,
	0x79, 0x12, 0x46, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x25, 0xfa, 0x41, 0x22, 0x0a, 0x20, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x74, 0x52, 0x0a, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x74, 0x12, 0x2d, 0x0a, 0x12, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x22, 0xbe, 0x06, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x54, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x70, 0x34,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x79, 0x0a, 0x17, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x65, 0x64, 0x5f, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x70, 0x34, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x65, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x15,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x65, 0x64, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x1a, 0x76, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x40, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x70, 0x34, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x1a, 0x73, 0x0a,
	0x10, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x1a, 0xac, 0x02, 0x0a, 0x0d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x65, 0x64, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x50, 0x0a, 0x0d, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x5f, 0x70, 0x6f, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x70, 0x34, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x42, 0x6f, 0x75, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x79, 0x52, 0x0c, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x79, 0x12, 0x54, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x70, 0x34, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x2e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x73, 0x0a, 0x12,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x70, 0x34, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x2e, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x42, 0x88, 0x01, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x70, 0x34, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x12, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x70, 0x34, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x04, 0x47, 0x43, 0x56, 0x4e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_vision_v1p4beta1_product_search_proto_rawDescOnce sync.Once
	file_google_cloud_vision_v1p4beta1_product_search_proto_rawDescData = file_google_cloud_vision_v1p4beta1_product_search_proto_rawDesc
)

func file_google_cloud_vision_v1p4beta1_product_search_proto_rawDescGZIP() []byte {
	file_google_cloud_vision_v1p4beta1_product_search_proto_rawDescOnce.Do(func() {
		file_google_cloud_vision_v1p4beta1_product_search_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_vision_v1p4beta1_product_search_proto_rawDescData)
	})
	return file_google_cloud_vision_v1p4beta1_product_search_proto_rawDescData
}

var file_google_cloud_vision_v1p4beta1_product_search_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_cloud_vision_v1p4beta1_product_search_proto_goTypes = []interface{}{
	(*ProductSearchParams)(nil),                   // 0: google.cloud.vision.v1p4beta1.ProductSearchParams
	(*ProductSearchResults)(nil),                  // 1: google.cloud.vision.v1p4beta1.ProductSearchResults
	(*ProductSearchResults_Result)(nil),           // 2: google.cloud.vision.v1p4beta1.ProductSearchResults.Result
	(*ProductSearchResults_ObjectAnnotation)(nil), // 3: google.cloud.vision.v1p4beta1.ProductSearchResults.ObjectAnnotation
	(*ProductSearchResults_GroupedResult)(nil),    // 4: google.cloud.vision.v1p4beta1.ProductSearchResults.GroupedResult
	(*BoundingPoly)(nil),                          // 5: google.cloud.vision.v1p4beta1.BoundingPoly
	(*timestamppb.Timestamp)(nil),                 // 6: google.protobuf.Timestamp
	(*Product)(nil),                               // 7: google.cloud.vision.v1p4beta1.Product
}
var file_google_cloud_vision_v1p4beta1_product_search_proto_depIdxs = []int32{
	5, // 0: google.cloud.vision.v1p4beta1.ProductSearchParams.bounding_poly:type_name -> google.cloud.vision.v1p4beta1.BoundingPoly
	6, // 1: google.cloud.vision.v1p4beta1.ProductSearchResults.index_time:type_name -> google.protobuf.Timestamp
	2, // 2: google.cloud.vision.v1p4beta1.ProductSearchResults.results:type_name -> google.cloud.vision.v1p4beta1.ProductSearchResults.Result
	4, // 3: google.cloud.vision.v1p4beta1.ProductSearchResults.product_grouped_results:type_name -> google.cloud.vision.v1p4beta1.ProductSearchResults.GroupedResult
	7, // 4: google.cloud.vision.v1p4beta1.ProductSearchResults.Result.product:type_name -> google.cloud.vision.v1p4beta1.Product
	5, // 5: google.cloud.vision.v1p4beta1.ProductSearchResults.GroupedResult.bounding_poly:type_name -> google.cloud.vision.v1p4beta1.BoundingPoly
	2, // 6: google.cloud.vision.v1p4beta1.ProductSearchResults.GroupedResult.results:type_name -> google.cloud.vision.v1p4beta1.ProductSearchResults.Result
	3, // 7: google.cloud.vision.v1p4beta1.ProductSearchResults.GroupedResult.object_annotations:type_name -> google.cloud.vision.v1p4beta1.ProductSearchResults.ObjectAnnotation
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_google_cloud_vision_v1p4beta1_product_search_proto_init() }
func file_google_cloud_vision_v1p4beta1_product_search_proto_init() {
	if File_google_cloud_vision_v1p4beta1_product_search_proto != nil {
		return
	}
	file_google_cloud_vision_v1p4beta1_geometry_proto_init()
	file_google_cloud_vision_v1p4beta1_product_search_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_vision_v1p4beta1_product_search_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductSearchParams); i {
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
		file_google_cloud_vision_v1p4beta1_product_search_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductSearchResults); i {
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
		file_google_cloud_vision_v1p4beta1_product_search_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductSearchResults_Result); i {
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
		file_google_cloud_vision_v1p4beta1_product_search_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductSearchResults_ObjectAnnotation); i {
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
		file_google_cloud_vision_v1p4beta1_product_search_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductSearchResults_GroupedResult); i {
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
			RawDescriptor: file_google_cloud_vision_v1p4beta1_product_search_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_vision_v1p4beta1_product_search_proto_goTypes,
		DependencyIndexes: file_google_cloud_vision_v1p4beta1_product_search_proto_depIdxs,
		MessageInfos:      file_google_cloud_vision_v1p4beta1_product_search_proto_msgTypes,
	}.Build()
	File_google_cloud_vision_v1p4beta1_product_search_proto = out.File
	file_google_cloud_vision_v1p4beta1_product_search_proto_rawDesc = nil
	file_google_cloud_vision_v1p4beta1_product_search_proto_goTypes = nil
	file_google_cloud_vision_v1p4beta1_product_search_proto_depIdxs = nil
}
