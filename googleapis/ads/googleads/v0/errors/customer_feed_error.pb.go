// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/customer_feed_error.proto

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

// Enum describing possible customer feed errors.
type CustomerFeedErrorEnum_CustomerFeedError int32

const (
	// Enum unspecified.
	CustomerFeedErrorEnum_UNSPECIFIED CustomerFeedErrorEnum_CustomerFeedError = 0
	// The received error code is not known in this version.
	CustomerFeedErrorEnum_UNKNOWN CustomerFeedErrorEnum_CustomerFeedError = 1
	// An active feed already exists for this customer and place holder type.
	CustomerFeedErrorEnum_FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE CustomerFeedErrorEnum_CustomerFeedError = 2
	// The specified feed is removed.
	CustomerFeedErrorEnum_CANNOT_CREATE_FOR_REMOVED_FEED CustomerFeedErrorEnum_CustomerFeedError = 3
	// The CustomerFeed already exists. Update should be used to modify the
	// existing CustomerFeed.
	CustomerFeedErrorEnum_CANNOT_CREATE_ALREADY_EXISTING_CUSTOMER_FEED CustomerFeedErrorEnum_CustomerFeedError = 4
	// Cannot update removed customer feed.
	CustomerFeedErrorEnum_CANNOT_MODIFY_REMOVED_CUSTOMER_FEED CustomerFeedErrorEnum_CustomerFeedError = 5
	// Invalid placeholder type.
	CustomerFeedErrorEnum_INVALID_PLACEHOLDER_TYPE CustomerFeedErrorEnum_CustomerFeedError = 6
	// Feed mapping for this placeholder type does not exist.
	CustomerFeedErrorEnum_MISSING_FEEDMAPPING_FOR_PLACEHOLDER_TYPE CustomerFeedErrorEnum_CustomerFeedError = 7
	// Placeholder not allowed at the account level.
	CustomerFeedErrorEnum_PLACEHOLDER_TYPE_NOT_ALLOWED_ON_CUSTOMER_FEED CustomerFeedErrorEnum_CustomerFeedError = 8
)

var CustomerFeedErrorEnum_CustomerFeedError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE",
	3: "CANNOT_CREATE_FOR_REMOVED_FEED",
	4: "CANNOT_CREATE_ALREADY_EXISTING_CUSTOMER_FEED",
	5: "CANNOT_MODIFY_REMOVED_CUSTOMER_FEED",
	6: "INVALID_PLACEHOLDER_TYPE",
	7: "MISSING_FEEDMAPPING_FOR_PLACEHOLDER_TYPE",
	8: "PLACEHOLDER_TYPE_NOT_ALLOWED_ON_CUSTOMER_FEED",
}
var CustomerFeedErrorEnum_CustomerFeedError_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE":      2,
	"CANNOT_CREATE_FOR_REMOVED_FEED":                3,
	"CANNOT_CREATE_ALREADY_EXISTING_CUSTOMER_FEED":  4,
	"CANNOT_MODIFY_REMOVED_CUSTOMER_FEED":           5,
	"INVALID_PLACEHOLDER_TYPE":                      6,
	"MISSING_FEEDMAPPING_FOR_PLACEHOLDER_TYPE":      7,
	"PLACEHOLDER_TYPE_NOT_ALLOWED_ON_CUSTOMER_FEED": 8,
}

func (x CustomerFeedErrorEnum_CustomerFeedError) String() string {
	return proto.EnumName(CustomerFeedErrorEnum_CustomerFeedError_name, int32(x))
}
func (CustomerFeedErrorEnum_CustomerFeedError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_customer_feed_error_7c655ec838c7ee5d, []int{0, 0}
}

// Container for enum describing possible customer feed errors.
type CustomerFeedErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomerFeedErrorEnum) Reset()         { *m = CustomerFeedErrorEnum{} }
func (m *CustomerFeedErrorEnum) String() string { return proto.CompactTextString(m) }
func (*CustomerFeedErrorEnum) ProtoMessage()    {}
func (*CustomerFeedErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_feed_error_7c655ec838c7ee5d, []int{0}
}
func (m *CustomerFeedErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerFeedErrorEnum.Unmarshal(m, b)
}
func (m *CustomerFeedErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerFeedErrorEnum.Marshal(b, m, deterministic)
}
func (dst *CustomerFeedErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerFeedErrorEnum.Merge(dst, src)
}
func (m *CustomerFeedErrorEnum) XXX_Size() int {
	return xxx_messageInfo_CustomerFeedErrorEnum.Size(m)
}
func (m *CustomerFeedErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerFeedErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerFeedErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CustomerFeedErrorEnum)(nil), "google.ads.googleads.v0.errors.CustomerFeedErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.CustomerFeedErrorEnum_CustomerFeedError", CustomerFeedErrorEnum_CustomerFeedError_name, CustomerFeedErrorEnum_CustomerFeedError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/customer_feed_error.proto", fileDescriptor_customer_feed_error_7c655ec838c7ee5d)
}

var fileDescriptor_customer_feed_error_7c655ec838c7ee5d = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x89, 0x0b, 0x2d, 0xda, 0x1e, 0x08, 0x2b, 0x81, 0x38, 0xa0, 0x1c, 0xcc, 0x01, 0x0e,
	0x65, 0x6d, 0xc4, 0x05, 0x99, 0xd3, 0xd6, 0x3b, 0x0e, 0x16, 0xf6, 0xae, 0x65, 0x3b, 0x2e, 0x41,
	0x91, 0x56, 0xa1, 0x36, 0x16, 0x52, 0x93, 0xad, 0xbc, 0x6d, 0x1f, 0x88, 0x23, 0x8f, 0xc2, 0x3b,
	0xf0, 0x02, 0xdc, 0x38, 0x71, 0x45, 0xeb, 0x4d, 0x2d, 0x25, 0x01, 0x4e, 0xfe, 0x35, 0xfe, 0xfe,
	0x99, 0x59, 0xfd, 0x83, 0xde, 0xb4, 0x4a, 0xb5, 0x17, 0x8d, 0xb7, 0xac, 0xb5, 0x67, 0xa5, 0x51,
	0x37, 0xbe, 0xd7, 0x74, 0x9d, 0xea, 0xb4, 0x77, 0x7e, 0xad, 0xaf, 0xd4, 0xaa, 0xe9, 0xe4, 0xe7,
	0xa6, 0xa9, 0x65, 0x5f, 0x24, 0x97, 0x9d, 0xba, 0x52, 0x78, 0x62, 0x71, 0xb2, 0xac, 0x35, 0x19,
	0x9c, 0xe4, 0xc6, 0x27, 0xd6, 0xe9, 0xfe, 0x76, 0xd0, 0xa3, 0x70, 0xe3, 0x8e, 0x9a, 0xa6, 0x06,
	0x53, 0x86, 0xf5, 0xf5, 0xca, 0xfd, 0xe1, 0xa0, 0x87, 0x7b, 0x7f, 0xf0, 0x03, 0x74, 0x3c, 0xe3,
	0x45, 0x06, 0x61, 0x1c, 0xc5, 0xc0, 0xc6, 0x77, 0xf0, 0x31, 0x3a, 0x9a, 0xf1, 0xf7, 0x5c, 0x9c,
	0xf1, 0xf1, 0x08, 0x9f, 0xa0, 0x17, 0x11, 0x00, 0x93, 0x34, 0xc9, 0x81, 0xb2, 0xb9, 0x84, 0x0f,
	0x71, 0x51, 0x16, 0x32, 0x12, 0xb9, 0xcc, 0x12, 0x1a, 0xc2, 0x3b, 0x91, 0x30, 0xc8, 0x65, 0x39,
	0xcf, 0x60, 0xec, 0x60, 0x17, 0x4d, 0x42, 0xca, 0xb9, 0x28, 0x65, 0x98, 0x03, 0x2d, 0xa1, 0xe7,
	0x72, 0x48, 0x45, 0x05, 0x4c, 0x9a, 0x3e, 0xe3, 0x03, 0xec, 0xa3, 0x93, 0x6d, 0x66, 0xab, 0x75,
	0xcc, 0xa7, 0x32, 0x9c, 0x15, 0xa5, 0x48, 0x21, 0xb7, 0x8e, 0xbb, 0xf8, 0x39, 0x7a, 0xb6, 0x71,
	0xa4, 0x82, 0xc5, 0xd1, 0x7c, 0xe8, 0xb8, 0x0d, 0xde, 0xc3, 0x4f, 0xd1, 0x93, 0x98, 0x57, 0x34,
	0x89, 0xd9, 0xfe, 0x72, 0x87, 0xe6, 0x29, 0x69, 0x5c, 0x14, 0x66, 0x82, 0xe1, 0x53, 0x9a, 0x65,
	0xbd, 0xfe, 0xdb, 0x53, 0x8e, 0xf0, 0x2b, 0xf4, 0x72, 0xb7, 0x2a, 0xcd, 0x0a, 0x34, 0x49, 0xc4,
	0x19, 0x30, 0x29, 0xf8, 0xce, 0xf8, 0xfb, 0xa7, 0xbf, 0x46, 0xc8, 0x3d, 0x57, 0x2b, 0xf2, 0xff,
	0x80, 0x4e, 0x1f, 0xef, 0x65, 0x90, 0x99, 0x60, 0xb3, 0xd1, 0x47, 0xb6, 0x71, 0xb6, 0xea, 0x62,
	0xb9, 0x6e, 0x89, 0xea, 0x5a, 0xaf, 0x6d, 0xd6, 0x7d, 0xec, 0xb7, 0x47, 0x72, 0xf9, 0x45, 0xff,
	0xeb, 0x66, 0xde, 0xda, 0xcf, 0x57, 0xe7, 0x60, 0x4a, 0xe9, 0x37, 0x67, 0x32, 0xb5, 0xcd, 0x68,
	0xad, 0x89, 0x95, 0x46, 0x55, 0x3e, 0xe9, 0x47, 0xea, 0xef, 0xb7, 0xc0, 0x82, 0xd6, 0x7a, 0x31,
	0x00, 0x8b, 0xca, 0x5f, 0x58, 0xe0, 0xa7, 0xe3, 0xda, 0x6a, 0x10, 0xd0, 0x5a, 0x07, 0xc1, 0x80,
	0x04, 0x41, 0xe5, 0x07, 0x81, 0x85, 0x3e, 0x1d, 0xf6, 0xdb, 0xbd, 0xfe, 0x13, 0x00, 0x00, 0xff,
	0xff, 0x29, 0x1f, 0x8c, 0xe2, 0xd0, 0x02, 0x00, 0x00,
}
