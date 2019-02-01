// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/customer_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v0/errors"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Set of errors that are related to requests dealing with Customer.
// Next id: 26
type CustomerErrorEnum_CustomerError int32

const (
	// Enum unspecified.
	CustomerErrorEnum_UNSPECIFIED CustomerErrorEnum_CustomerError = 0
	// The received error code is not known in this version.
	CustomerErrorEnum_UNKNOWN CustomerErrorEnum_CustomerError = 1
	// Customer status is not allowed to be changed from DRAFT and CLOSED.
	// Currency code and at least one of country code and time zone needs to be
	// set when status is changed to ENABLED.
	CustomerErrorEnum_STATUS_CHANGE_DISALLOWED CustomerErrorEnum_CustomerError = 2
	// CustomerService cannot get a customer that has not been fully set up.
	CustomerErrorEnum_ACCOUNT_NOT_SET_UP CustomerErrorEnum_CustomerError = 3
)

var CustomerErrorEnum_CustomerError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "STATUS_CHANGE_DISALLOWED",
	3: "ACCOUNT_NOT_SET_UP",
}
var CustomerErrorEnum_CustomerError_value = map[string]int32{
	"UNSPECIFIED":              0,
	"UNKNOWN":                  1,
	"STATUS_CHANGE_DISALLOWED": 2,
	"ACCOUNT_NOT_SET_UP":       3,
}

func (x CustomerErrorEnum_CustomerError) String() string {
	return proto.EnumName(CustomerErrorEnum_CustomerError_name, int32(x))
}
func (CustomerErrorEnum_CustomerError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_customer_error_350560db7ca35ecc, []int{0, 0}
}

// Container for enum describing possible customer errors.
type CustomerErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomerErrorEnum) Reset()         { *m = CustomerErrorEnum{} }
func (m *CustomerErrorEnum) String() string { return proto.CompactTextString(m) }
func (*CustomerErrorEnum) ProtoMessage()    {}
func (*CustomerErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_error_350560db7ca35ecc, []int{0}
}
func (m *CustomerErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerErrorEnum.Unmarshal(m, b)
}
func (m *CustomerErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerErrorEnum.Marshal(b, m, deterministic)
}
func (dst *CustomerErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerErrorEnum.Merge(dst, src)
}
func (m *CustomerErrorEnum) XXX_Size() int {
	return xxx_messageInfo_CustomerErrorEnum.Size(m)
}
func (m *CustomerErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CustomerErrorEnum)(nil), "google.ads.googleads.v0.errors.CustomerErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.CustomerErrorEnum_CustomerError", CustomerErrorEnum_CustomerError_name, CustomerErrorEnum_CustomerError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/customer_error.proto", fileDescriptor_customer_error_350560db7ca35ecc)
}

var fileDescriptor_customer_error_350560db7ca35ecc = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xdf, 0x4a, 0xfb, 0x30,
	0x1c, 0xc5, 0x7f, 0xeb, 0xe0, 0x27, 0x64, 0x88, 0x35, 0x17, 0xe2, 0x85, 0xec, 0xa2, 0x0f, 0x90,
	0x16, 0x76, 0x17, 0xaf, 0xb2, 0x36, 0xce, 0xe1, 0x48, 0x0b, 0xfd, 0x33, 0x90, 0x42, 0xa8, 0x6d,
	0x09, 0xc2, 0xda, 0x8c, 0x64, 0x1b, 0x3e, 0x8f, 0x97, 0x3e, 0x8a, 0x8f, 0x22, 0xf8, 0x0e, 0xd2,
	0x66, 0x2b, 0xec, 0x42, 0xaf, 0x72, 0x72, 0xf8, 0x9c, 0xe4, 0x7c, 0xbf, 0x60, 0x26, 0xa4, 0x14,
	0x9b, 0xda, 0x2d, 0x2a, 0xed, 0x1a, 0xd9, 0xa9, 0x83, 0xe7, 0xd6, 0x4a, 0x49, 0xa5, 0xdd, 0x72,
	0xaf, 0x77, 0xb2, 0xa9, 0x15, 0xef, 0xef, 0x68, 0xab, 0xe4, 0x4e, 0xc2, 0xa9, 0x21, 0x51, 0x51,
	0x69, 0x34, 0x84, 0xd0, 0xc1, 0x43, 0x26, 0xe4, 0xbc, 0x81, 0x6b, 0xff, 0x98, 0xa3, 0x9d, 0x43,
	0xdb, 0x7d, 0xe3, 0x94, 0xe0, 0xf2, 0xcc, 0x84, 0x57, 0x60, 0x92, 0xb2, 0x38, 0xa2, 0xfe, 0xf2,
	0x61, 0x49, 0x03, 0xfb, 0x1f, 0x9c, 0x80, 0x8b, 0x94, 0x3d, 0xb1, 0x70, 0xcd, 0xec, 0x11, 0xbc,
	0x03, 0xb7, 0x71, 0x42, 0x92, 0x34, 0xe6, 0xfe, 0x23, 0x61, 0x0b, 0xca, 0x83, 0x65, 0x4c, 0x56,
	0xab, 0x70, 0x4d, 0x03, 0xdb, 0x82, 0x37, 0x00, 0x12, 0xdf, 0x0f, 0x53, 0x96, 0x70, 0x16, 0x26,
	0x3c, 0xa6, 0x09, 0x4f, 0x23, 0x7b, 0x3c, 0xff, 0x1e, 0x01, 0xa7, 0x94, 0x0d, 0xfa, 0xbb, 0xe0,
	0x1c, 0x9e, 0x35, 0x89, 0xba, 0xa1, 0xa2, 0xd1, 0x73, 0x70, 0x4c, 0x09, 0xb9, 0x29, 0x5a, 0x81,
	0xa4, 0x12, 0xae, 0xa8, 0xdb, 0x7e, 0xe4, 0xd3, 0x6e, 0xb6, 0xaf, 0xfa, 0xb7, 0x55, 0xdd, 0x9b,
	0xe3, 0xdd, 0x1a, 0x2f, 0x08, 0xf9, 0xb0, 0xa6, 0x0b, 0xf3, 0x18, 0xa9, 0x34, 0x32, 0xb2, 0x53,
	0x99, 0x87, 0xfa, 0x2f, 0xf5, 0xe7, 0x09, 0xc8, 0x49, 0xa5, 0xf3, 0x01, 0xc8, 0x33, 0x2f, 0x37,
	0xc0, 0x97, 0xe5, 0x18, 0x17, 0x63, 0x52, 0x69, 0x8c, 0x07, 0x04, 0xe3, 0xcc, 0xc3, 0xd8, 0x40,
	0x2f, 0xff, 0xfb, 0x76, 0xb3, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd4, 0xb9, 0xcb, 0xa3, 0xc7,
	0x01, 0x00, 0x00,
}
