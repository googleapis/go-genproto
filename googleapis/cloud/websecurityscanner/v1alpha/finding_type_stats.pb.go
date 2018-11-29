// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/websecurityscanner/v1alpha/finding_type_stats.proto

package websecurityscanner

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// A FindingTypeStats resource represents stats regarding a specific FindingType
// of Findings under a given ScanRun.
type FindingTypeStats struct {
	// Output only.
	// The finding type associated with the stats.
	FindingType Finding_FindingType `protobuf:"varint,1,opt,name=finding_type,json=findingType,proto3,enum=google.cloud.websecurityscanner.v1alpha.Finding_FindingType" json:"finding_type,omitempty"`
	// Output only.
	// The count of findings belonging to this finding type.
	FindingCount         int32    `protobuf:"varint,2,opt,name=finding_count,json=findingCount,proto3" json:"finding_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindingTypeStats) Reset()         { *m = FindingTypeStats{} }
func (m *FindingTypeStats) String() string { return proto.CompactTextString(m) }
func (*FindingTypeStats) ProtoMessage()    {}
func (*FindingTypeStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_91da39ac488bf6ea, []int{0}
}

func (m *FindingTypeStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindingTypeStats.Unmarshal(m, b)
}
func (m *FindingTypeStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindingTypeStats.Marshal(b, m, deterministic)
}
func (m *FindingTypeStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindingTypeStats.Merge(m, src)
}
func (m *FindingTypeStats) XXX_Size() int {
	return xxx_messageInfo_FindingTypeStats.Size(m)
}
func (m *FindingTypeStats) XXX_DiscardUnknown() {
	xxx_messageInfo_FindingTypeStats.DiscardUnknown(m)
}

var xxx_messageInfo_FindingTypeStats proto.InternalMessageInfo

func (m *FindingTypeStats) GetFindingType() Finding_FindingType {
	if m != nil {
		return m.FindingType
	}
	return Finding_FINDING_TYPE_UNSPECIFIED
}

func (m *FindingTypeStats) GetFindingCount() int32 {
	if m != nil {
		return m.FindingCount
	}
	return 0
}

func init() {
	proto.RegisterType((*FindingTypeStats)(nil), "google.cloud.websecurityscanner.v1alpha.FindingTypeStats")
}

func init() {
	proto.RegisterFile("google/cloud/websecurityscanner/v1alpha/finding_type_stats.proto", fileDescriptor_91da39ac488bf6ea)
}

var fileDescriptor_91da39ac488bf6ea = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x48, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0xce, 0xc9, 0x2f, 0x4d, 0xd1, 0x2f, 0x4f, 0x4d, 0x2a, 0x4e, 0x4d, 0x2e,
	0x2d, 0xca, 0x2c, 0xa9, 0x2c, 0x4e, 0x4e, 0xcc, 0xcb, 0x4b, 0x2d, 0xd2, 0x2f, 0x33, 0x4c, 0xcc,
	0x29, 0xc8, 0x48, 0xd4, 0x4f, 0xcb, 0xcc, 0x4b, 0xc9, 0xcc, 0x4b, 0x8f, 0x2f, 0xa9, 0x2c, 0x48,
	0x8d, 0x2f, 0x2e, 0x49, 0x2c, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x87, 0x98,
	0xa0, 0x07, 0x36, 0x41, 0x0f, 0xd3, 0x04, 0x3d, 0xa8, 0x09, 0x52, 0x32, 0x50, 0xab, 0x12, 0x0b,
	0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0xa0, 0xc6, 0x48, 0x99,
	0x92, 0xe8, 0x10, 0x88, 0x36, 0xa5, 0x19, 0x8c, 0x5c, 0x02, 0x6e, 0x10, 0x91, 0x90, 0xca, 0x82,
	0xd4, 0x60, 0x90, 0xc3, 0x84, 0xe2, 0xb9, 0x78, 0x90, 0x9d, 0x2b, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1,
	0x67, 0x64, 0xa3, 0x47, 0xa4, 0x4b, 0xf5, 0xa0, 0x06, 0xea, 0x21, 0x19, 0x1c, 0xc4, 0x9d, 0x86,
	0xe0, 0x08, 0x29, 0x73, 0xf1, 0xc2, 0x2c, 0x48, 0xce, 0x2f, 0xcd, 0x2b, 0x91, 0x60, 0x52, 0x60,
	0xd4, 0x60, 0x0d, 0x82, 0xd9, 0xea, 0x0c, 0x12, 0x73, 0x5a, 0xc8, 0xc8, 0xa5, 0x9d, 0x9c, 0x9f,
	0x4b, 0xac, 0xad, 0x4e, 0xa2, 0xe8, 0xfe, 0x08, 0x00, 0xf9, 0x30, 0x80, 0x31, 0x2a, 0x12, 0x6a,
	0x42, 0x7a, 0x7e, 0x4e, 0x62, 0x5e, 0xba, 0x5e, 0x7e, 0x51, 0xba, 0x7e, 0x7a, 0x6a, 0x1e, 0xd8,
	0xff, 0xfa, 0x10, 0xa9, 0xc4, 0x82, 0xcc, 0x62, 0x82, 0x21, 0x67, 0x8d, 0x29, 0x95, 0xc4, 0x06,
	0x36, 0xc5, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xbd, 0xb4, 0x52, 0xba, 0x07, 0x02, 0x00, 0x00,
}
