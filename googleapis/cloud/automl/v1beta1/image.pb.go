// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/image.proto

package automl

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Dataset metadata that is specific to image classification.
type ImageClassificationDatasetMetadata struct {
	// Required.
	// Type of the classification problem.
	ClassificationType   ClassificationType `protobuf:"varint,1,opt,name=classification_type,json=classificationType,proto3,enum=google.cloud.automl.v1beta1.ClassificationType" json:"classification_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ImageClassificationDatasetMetadata) Reset()         { *m = ImageClassificationDatasetMetadata{} }
func (m *ImageClassificationDatasetMetadata) String() string { return proto.CompactTextString(m) }
func (*ImageClassificationDatasetMetadata) ProtoMessage()    {}
func (*ImageClassificationDatasetMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_29b9f2bc900da869, []int{0}
}

func (m *ImageClassificationDatasetMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageClassificationDatasetMetadata.Unmarshal(m, b)
}
func (m *ImageClassificationDatasetMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageClassificationDatasetMetadata.Marshal(b, m, deterministic)
}
func (m *ImageClassificationDatasetMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageClassificationDatasetMetadata.Merge(m, src)
}
func (m *ImageClassificationDatasetMetadata) XXX_Size() int {
	return xxx_messageInfo_ImageClassificationDatasetMetadata.Size(m)
}
func (m *ImageClassificationDatasetMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageClassificationDatasetMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ImageClassificationDatasetMetadata proto.InternalMessageInfo

func (m *ImageClassificationDatasetMetadata) GetClassificationType() ClassificationType {
	if m != nil {
		return m.ClassificationType
	}
	return ClassificationType_CLASSIFICATION_TYPE_UNSPECIFIED
}

// Model metadata for image classification.
type ImageClassificationModelMetadata struct {
	// Optional. The ID of the `base` model. If it is specified, the new model
	// will be created based on the `base` model. Otherwise, the new model will be
	// created from scratch. The `base` model is expected to be in the same
	// `project` and `location` as the new model to create.
	BaseModelId string `protobuf:"bytes,1,opt,name=base_model_id,json=baseModelId,proto3" json:"base_model_id,omitempty"`
	// Required. The train budget of creating this model. The actual
	// `train_cost` will be equal or less than this value.
	TrainBudget int64 `protobuf:"varint,2,opt,name=train_budget,json=trainBudget,proto3" json:"train_budget,omitempty"`
	// Output only. The actual train cost of creating this model. If this
	// model is created from a `base` model, the train cost used to create the
	// `base` model are not included.
	TrainCost int64 `protobuf:"varint,3,opt,name=train_cost,json=trainCost,proto3" json:"train_cost,omitempty"`
	// Output only. The reason that this create model operation stopped,
	// e.g. BUDGET_REACHED, CONVERGED.
	StopReason           string   `protobuf:"bytes,5,opt,name=stop_reason,json=stopReason,proto3" json:"stop_reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageClassificationModelMetadata) Reset()         { *m = ImageClassificationModelMetadata{} }
func (m *ImageClassificationModelMetadata) String() string { return proto.CompactTextString(m) }
func (*ImageClassificationModelMetadata) ProtoMessage()    {}
func (*ImageClassificationModelMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_29b9f2bc900da869, []int{1}
}

func (m *ImageClassificationModelMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageClassificationModelMetadata.Unmarshal(m, b)
}
func (m *ImageClassificationModelMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageClassificationModelMetadata.Marshal(b, m, deterministic)
}
func (m *ImageClassificationModelMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageClassificationModelMetadata.Merge(m, src)
}
func (m *ImageClassificationModelMetadata) XXX_Size() int {
	return xxx_messageInfo_ImageClassificationModelMetadata.Size(m)
}
func (m *ImageClassificationModelMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageClassificationModelMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ImageClassificationModelMetadata proto.InternalMessageInfo

func (m *ImageClassificationModelMetadata) GetBaseModelId() string {
	if m != nil {
		return m.BaseModelId
	}
	return ""
}

func (m *ImageClassificationModelMetadata) GetTrainBudget() int64 {
	if m != nil {
		return m.TrainBudget
	}
	return 0
}

func (m *ImageClassificationModelMetadata) GetTrainCost() int64 {
	if m != nil {
		return m.TrainCost
	}
	return 0
}

func (m *ImageClassificationModelMetadata) GetStopReason() string {
	if m != nil {
		return m.StopReason
	}
	return ""
}

func init() {
	proto.RegisterType((*ImageClassificationDatasetMetadata)(nil), "google.cloud.automl.v1beta1.ImageClassificationDatasetMetadata")
	proto.RegisterType((*ImageClassificationModelMetadata)(nil), "google.cloud.automl.v1beta1.ImageClassificationModelMetadata")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/image.proto", fileDescriptor_29b9f2bc900da869)
}

var fileDescriptor_29b9f2bc900da869 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xb1, 0x6a, 0xe3, 0x40,
	0x10, 0x86, 0xd1, 0x99, 0x3b, 0xf0, 0xfa, 0xee, 0x0a, 0xa5, 0x11, 0x4e, 0x82, 0x1d, 0x35, 0x71,
	0xa5, 0x8d, 0x93, 0x32, 0x55, 0xec, 0x34, 0x2e, 0x0c, 0x46, 0xa4, 0x4a, 0xa3, 0x8c, 0xa4, 0xf5,
	0xb2, 0x20, 0x69, 0x16, 0xed, 0x28, 0xe0, 0x17, 0xc8, 0xbb, 0xe4, 0x2d, 0x83, 0x66, 0x45, 0xc0,
	0xc4, 0xb8, 0xdc, 0x4f, 0xdf, 0x3f, 0xfb, 0x8f, 0x56, 0xdc, 0x6a, 0x44, 0x5d, 0x29, 0x59, 0x54,
	0xd8, 0x95, 0x12, 0x3a, 0xc2, 0xba, 0x92, 0xef, 0xcb, 0x5c, 0x11, 0x2c, 0xa5, 0xa9, 0x41, 0xab,
	0xc4, 0xb6, 0x48, 0x18, 0x5e, 0x7a, 0x31, 0x61, 0x31, 0xf1, 0x62, 0x32, 0x88, 0xd3, 0xab, 0x61,
	0x0a, 0x58, 0x23, 0xa1, 0x69, 0x90, 0x80, 0x0c, 0x36, 0xce, 0x47, 0xa7, 0x77, 0xe7, 0xee, 0x28,
	0x2a, 0x70, 0xce, 0xec, 0x4d, 0xc1, 0x91, 0x21, 0x31, 0x1b, 0x12, 0x7c, 0xca, 0xbb, 0xbd, 0x24,
	0x53, 0x2b, 0x47, 0x50, 0x5b, 0x2f, 0xc4, 0x1f, 0x81, 0x88, 0x37, 0x7d, 0xbb, 0xf5, 0x51, 0xfc,
	0x19, 0x08, 0x9c, 0xa2, 0xad, 0x22, 0x28, 0x81, 0x20, 0x7c, 0x13, 0x17, 0xc7, 0xf3, 0x33, 0x3a,
	0x58, 0x15, 0x05, 0xf3, 0x60, 0xf1, 0xff, 0x5e, 0x26, 0x67, 0x56, 0x4a, 0x8e, 0x07, 0xbf, 0x1c,
	0xac, 0x4a, 0xc3, 0xe2, 0x07, 0x8b, 0x3f, 0x03, 0x31, 0x3f, 0x51, 0x64, 0x8b, 0xa5, 0xaa, 0xbe,
	0x6b, 0xc4, 0xe2, 0x5f, 0x0e, 0x4e, 0x65, 0x75, 0x4f, 0x33, 0x53, 0x72, 0x81, 0x71, 0x3a, 0xe9,
	0x21, 0x9b, 0x9b, 0x32, 0xbc, 0x11, 0x7f, 0xa9, 0x05, 0xd3, 0x64, 0x79, 0x57, 0x6a, 0x45, 0xd1,
	0xaf, 0x79, 0xb0, 0x18, 0xa5, 0x13, 0x66, 0x2b, 0x46, 0xe1, 0xb5, 0x10, 0x5e, 0x29, 0xd0, 0x51,
	0x34, 0x62, 0x61, 0xcc, 0x64, 0x8d, 0x8e, 0xc2, 0x99, 0x98, 0x38, 0x42, 0x9b, 0xb5, 0x0a, 0x1c,
	0x36, 0xd1, 0x6f, 0xbe, 0x43, 0xf4, 0x28, 0x65, 0xb2, 0x6a, 0xc5, 0xac, 0xc0, 0xfa, 0xdc, 0xd6,
	0x2b, 0xc1, 0xbb, 0xec, 0xfa, 0x7f, 0xbc, 0x0b, 0x5e, 0x9f, 0x06, 0x55, 0x63, 0x05, 0x8d, 0x4e,
	0xb0, 0xd5, 0x52, 0xab, 0x86, 0x5f, 0x40, 0xfa, 0x4f, 0x60, 0x8d, 0x3b, 0xf9, 0xae, 0x8f, 0xfe,
	0x98, 0xff, 0x61, 0xfb, 0xe1, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xde, 0x9b, 0x5b, 0x1e, 0x68, 0x02,
	0x00, 0x00,
}
