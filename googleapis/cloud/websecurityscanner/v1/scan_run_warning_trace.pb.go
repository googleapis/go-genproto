// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/websecurityscanner/v1/scan_run_warning_trace.proto

package websecurityscanner

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// Output only.
// Defines a warning message code.
// Next id: 6
type ScanRunWarningTrace_Code int32

const (
	// Default value is never used.
	ScanRunWarningTrace_CODE_UNSPECIFIED ScanRunWarningTrace_Code = 0
	// Indicates that a scan discovered an unexpectedly low number of URLs. This
	// is sometimes caused by complex navigation features or by using a single
	// URL for numerous pages.
	ScanRunWarningTrace_INSUFFICIENT_CRAWL_RESULTS ScanRunWarningTrace_Code = 1
	// Indicates that a scan discovered too many URLs to test, or excessive
	// redundant URLs.
	ScanRunWarningTrace_TOO_MANY_CRAWL_RESULTS ScanRunWarningTrace_Code = 2
	// Indicates that too many tests have been generated for the scan. Customer
	// should try reducing the number of starting URLs, increasing the QPS rate,
	// or narrowing down the scope of the scan using the excluded patterns.
	ScanRunWarningTrace_TOO_MANY_FUZZ_TASKS ScanRunWarningTrace_Code = 3
	// Indicates that a scan is blocked by IAP.
	ScanRunWarningTrace_BLOCKED_BY_IAP ScanRunWarningTrace_Code = 4
)

var ScanRunWarningTrace_Code_name = map[int32]string{
	0: "CODE_UNSPECIFIED",
	1: "INSUFFICIENT_CRAWL_RESULTS",
	2: "TOO_MANY_CRAWL_RESULTS",
	3: "TOO_MANY_FUZZ_TASKS",
	4: "BLOCKED_BY_IAP",
}

var ScanRunWarningTrace_Code_value = map[string]int32{
	"CODE_UNSPECIFIED":           0,
	"INSUFFICIENT_CRAWL_RESULTS": 1,
	"TOO_MANY_CRAWL_RESULTS":     2,
	"TOO_MANY_FUZZ_TASKS":        3,
	"BLOCKED_BY_IAP":             4,
}

func (x ScanRunWarningTrace_Code) String() string {
	return proto.EnumName(ScanRunWarningTrace_Code_name, int32(x))
}

func (ScanRunWarningTrace_Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_49d33550124a8683, []int{0, 0}
}

// Output only.
// Defines a warning trace message for ScanRun. Warning traces provide customers
// with useful information that helps make the scanning process more effective.
type ScanRunWarningTrace struct {
	// Output only. Indicates the warning code.
	Code                 ScanRunWarningTrace_Code `protobuf:"varint,1,opt,name=code,proto3,enum=google.cloud.websecurityscanner.v1.ScanRunWarningTrace_Code" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ScanRunWarningTrace) Reset()         { *m = ScanRunWarningTrace{} }
func (m *ScanRunWarningTrace) String() string { return proto.CompactTextString(m) }
func (*ScanRunWarningTrace) ProtoMessage()    {}
func (*ScanRunWarningTrace) Descriptor() ([]byte, []int) {
	return fileDescriptor_49d33550124a8683, []int{0}
}

func (m *ScanRunWarningTrace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanRunWarningTrace.Unmarshal(m, b)
}
func (m *ScanRunWarningTrace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanRunWarningTrace.Marshal(b, m, deterministic)
}
func (m *ScanRunWarningTrace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanRunWarningTrace.Merge(m, src)
}
func (m *ScanRunWarningTrace) XXX_Size() int {
	return xxx_messageInfo_ScanRunWarningTrace.Size(m)
}
func (m *ScanRunWarningTrace) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanRunWarningTrace.DiscardUnknown(m)
}

var xxx_messageInfo_ScanRunWarningTrace proto.InternalMessageInfo

func (m *ScanRunWarningTrace) GetCode() ScanRunWarningTrace_Code {
	if m != nil {
		return m.Code
	}
	return ScanRunWarningTrace_CODE_UNSPECIFIED
}

func init() {
	proto.RegisterEnum("google.cloud.websecurityscanner.v1.ScanRunWarningTrace_Code", ScanRunWarningTrace_Code_name, ScanRunWarningTrace_Code_value)
	proto.RegisterType((*ScanRunWarningTrace)(nil), "google.cloud.websecurityscanner.v1.ScanRunWarningTrace")
}

func init() {
	proto.RegisterFile("google/cloud/websecurityscanner/v1/scan_run_warning_trace.proto", fileDescriptor_49d33550124a8683)
}

var fileDescriptor_49d33550124a8683 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd1, 0xc1, 0x4a, 0xfb, 0x30,
	0x1c, 0x07, 0xf0, 0x7f, 0xfe, 0x0e, 0x0f, 0x39, 0x8c, 0x92, 0x89, 0x8e, 0x1d, 0x44, 0x76, 0x10,
	0x4f, 0x29, 0xd3, 0xa3, 0x82, 0xb4, 0x5d, 0x0b, 0x65, 0xb3, 0x2d, 0x4d, 0xca, 0xd8, 0x2e, 0x21,
	0xcb, 0x42, 0x18, 0xcc, 0x64, 0x64, 0xeb, 0x86, 0x0f, 0xe0, 0x4b, 0xf8, 0x92, 0xbe, 0x82, 0xb4,
	0x15, 0x0f, 0x6e, 0xa8, 0xb7, 0x24, 0x5f, 0xf2, 0x81, 0xdf, 0xf7, 0x07, 0x1f, 0x95, 0x31, 0x6a,
	0x25, 0x5d, 0xb1, 0x32, 0xe5, 0xc2, 0xdd, 0xcb, 0xf9, 0x46, 0x8a, 0xd2, 0x2e, 0xb7, 0x2f, 0x1b,
	0xc1, 0xb5, 0x96, 0xd6, 0xdd, 0x0d, 0xdc, 0xea, 0xc8, 0x6c, 0xa9, 0xd9, 0x9e, 0x5b, 0xbd, 0xd4,
	0x8a, 0x6d, 0x2d, 0x17, 0x12, 0xaf, 0xad, 0xd9, 0x1a, 0xd4, 0x6f, 0x00, 0x5c, 0x03, 0xf8, 0x10,
	0xc0, 0xbb, 0x41, 0xff, 0x1d, 0xc0, 0x0e, 0x11, 0x5c, 0xe7, 0xa5, 0x9e, 0x34, 0x04, 0xad, 0x04,
	0x94, 0xc1, 0x96, 0x30, 0x0b, 0xd9, 0x05, 0x57, 0xe0, 0xa6, 0x7d, 0xfb, 0x80, 0x7f, 0xa7, 0xf0,
	0x11, 0x06, 0x07, 0x66, 0x21, 0xf3, 0x5a, 0xea, 0xbf, 0x02, 0xd8, 0xaa, 0xae, 0xe8, 0x0c, 0x3a,
	0x41, 0x3a, 0x0c, 0x59, 0x91, 0x90, 0x2c, 0x0c, 0xe2, 0x28, 0x0e, 0x87, 0xce, 0x3f, 0x74, 0x09,
	0x7b, 0x71, 0x42, 0x8a, 0x28, 0x8a, 0x83, 0x38, 0x4c, 0x28, 0x0b, 0x72, 0x6f, 0x32, 0x66, 0x79,
	0x48, 0x8a, 0x31, 0x25, 0x0e, 0x40, 0x3d, 0x78, 0x4e, 0xd3, 0x94, 0x3d, 0x79, 0xc9, 0xf4, 0x5b,
	0xf6, 0x1f, 0x5d, 0xc0, 0xce, 0x57, 0x16, 0x15, 0xb3, 0x19, 0xa3, 0x1e, 0x19, 0x11, 0xe7, 0x04,
	0x21, 0xd8, 0xf6, 0xc7, 0x69, 0x30, 0x0a, 0x87, 0xcc, 0x9f, 0xb2, 0xd8, 0xcb, 0x9c, 0x96, 0xff,
	0x06, 0xe0, 0xb5, 0x30, 0xcf, 0x7f, 0x98, 0xc8, 0xef, 0x1e, 0x19, 0x29, 0xab, 0xaa, 0xcd, 0xc0,
	0x8c, 0x7e, 0xfe, 0x57, 0x66, 0xc5, 0xb5, 0xc2, 0xc6, 0x2a, 0x57, 0x49, 0x5d, 0x17, 0xef, 0x36,
	0x11, 0x5f, 0x2f, 0x37, 0x3f, 0x2d, 0xef, 0xfe, 0xf0, 0x75, 0x7e, 0x5a, 0x03, 0x77, 0x1f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x76, 0xfc, 0x6f, 0x15, 0xfc, 0x01, 0x00, 0x00,
}
