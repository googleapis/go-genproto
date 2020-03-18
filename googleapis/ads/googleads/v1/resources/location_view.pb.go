// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/location_view.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A location view summarizes the performance of campaigns by
// Location criteria.
type LocationView struct {
	// Output only. The resource name of the location view.
	// Location view resource names have the form:
	//
	// `customers/{customer_id}/locationViews/{campaign_id}~{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocationView) Reset()         { *m = LocationView{} }
func (m *LocationView) String() string { return proto.CompactTextString(m) }
func (*LocationView) ProtoMessage()    {}
func (*LocationView) Descriptor() ([]byte, []int) {
	return fileDescriptor_3aa3d275dc2bd5c6, []int{0}
}

func (m *LocationView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocationView.Unmarshal(m, b)
}
func (m *LocationView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocationView.Marshal(b, m, deterministic)
}
func (m *LocationView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocationView.Merge(m, src)
}
func (m *LocationView) XXX_Size() int {
	return xxx_messageInfo_LocationView.Size(m)
}
func (m *LocationView) XXX_DiscardUnknown() {
	xxx_messageInfo_LocationView.DiscardUnknown(m)
}

var xxx_messageInfo_LocationView proto.InternalMessageInfo

func (m *LocationView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*LocationView)(nil), "google.ads.googleads.v1.resources.LocationView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/location_view.proto", fileDescriptor_3aa3d275dc2bd5c6)
}

var fileDescriptor_3aa3d275dc2bd5c6 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x31, 0x4b, 0xf3, 0x40,
	0x1c, 0xc6, 0x49, 0x0b, 0x2f, 0xbc, 0xa1, 0x0e, 0x76, 0xd2, 0x22, 0x68, 0x85, 0xa2, 0x8b, 0x77,
	0x44, 0x71, 0x39, 0xa7, 0xeb, 0x52, 0x10, 0x91, 0xd2, 0x21, 0x83, 0x06, 0xcb, 0x35, 0x39, 0xe3,
	0x41, 0x92, 0x7f, 0xb9, 0x4b, 0xd3, 0xa1, 0xf4, 0xcb, 0x38, 0xfa, 0x31, 0x1c, 0xfd, 0x14, 0xce,
	0xfd, 0x08, 0x0e, 0x22, 0xc9, 0xe5, 0xae, 0x71, 0x51, 0xb7, 0x07, 0xfe, 0xbf, 0xe7, 0xb9, 0x87,
	0xe7, 0xdc, 0xcb, 0x18, 0x20, 0x4e, 0x38, 0x66, 0x91, 0xc2, 0x5a, 0x96, 0xaa, 0xf0, 0xb0, 0xe4,
	0x0a, 0x16, 0x32, 0xe4, 0x0a, 0x27, 0x10, 0xb2, 0x5c, 0x40, 0x36, 0x2d, 0x04, 0x5f, 0xa2, 0xb9,
	0x84, 0x1c, 0xba, 0x7d, 0xcd, 0x22, 0x16, 0x29, 0x64, 0x6d, 0xa8, 0xf0, 0x90, 0xb5, 0xf5, 0x0e,
	0x4d, 0xf2, 0x5c, 0xe0, 0x47, 0xc1, 0x93, 0x68, 0x3a, 0xe3, 0x4f, 0xac, 0x10, 0x20, 0x75, 0x46,
	0x6f, 0xbf, 0x01, 0x18, 0x5b, 0x7d, 0x3a, 0x68, 0x9c, 0x58, 0x96, 0x41, 0x5e, 0x15, 0x50, 0xfa,
	0x7a, 0xfc, 0xea, 0xb8, 0x9d, 0x9b, 0xba, 0x94, 0x2f, 0xf8, 0xb2, 0x3b, 0x71, 0x77, 0x4c, 0xc0,
	0x34, 0x63, 0x29, 0xdf, 0x73, 0x8e, 0x9c, 0xd3, 0xff, 0xc3, 0xb3, 0x77, 0xda, 0xfe, 0xa0, 0x27,
	0xee, 0x60, 0xdb, 0xb0, 0x56, 0x73, 0xa1, 0x50, 0x08, 0x29, 0x6e, 0xa6, 0x4c, 0x3a, 0x26, 0xe3,
	0x96, 0xa5, 0x9c, 0x3c, 0x6c, 0xe8, 0xfd, 0x1f, 0x9d, 0xdd, 0xf3, 0x70, 0xa1, 0x72, 0x48, 0xb9,
	0x54, 0x78, 0x65, 0xe4, 0xda, 0xee, 0x56, 0x22, 0x0a, 0xaf, 0xbe, 0xcd, 0xb8, 0x1e, 0x7e, 0x3a,
	0xee, 0x20, 0x84, 0x14, 0xfd, 0x3a, 0xe4, 0x70, 0xb7, 0xf9, 0xd6, 0xb8, 0x5c, 0x60, 0xec, 0xdc,
	0x5d, 0xd7, 0xbe, 0x18, 0x12, 0x96, 0xc5, 0x08, 0x64, 0x8c, 0x63, 0x9e, 0x55, 0xfb, 0xe0, 0x6d,
	0xd5, 0x1f, 0xbe, 0xf5, 0xca, 0xaa, 0xe7, 0x56, 0x7b, 0x44, 0xe9, 0x4b, 0xab, 0x3f, 0xd2, 0x91,
	0x34, 0x52, 0x48, 0xcb, 0x52, 0xf9, 0x1e, 0x9a, 0x18, 0xf2, 0xcd, 0x30, 0x01, 0x8d, 0x54, 0x60,
	0x99, 0xc0, 0xf7, 0x02, 0xcb, 0x6c, 0x5a, 0x03, 0x7d, 0x20, 0x84, 0x46, 0x8a, 0x10, 0x4b, 0x11,
	0xe2, 0x7b, 0x84, 0x58, 0x6e, 0xf6, 0xaf, 0x2a, 0x7b, 0xf1, 0x15, 0x00, 0x00, 0xff, 0xff, 0xc9,
	0x12, 0x2c, 0x91, 0x82, 0x02, 0x00, 0x00,
}
