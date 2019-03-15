// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/securitycenter/v1/finding.proto

package securitycenter // import "google.golang.org/genproto/googleapis/cloud/securitycenter/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _struct "github.com/golang/protobuf/ptypes/struct"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The state of the finding.
type Finding_State int32

const (
	// Unspecified state.
	Finding_STATE_UNSPECIFIED Finding_State = 0
	// The finding requires attention and has not been addressed yet.
	Finding_ACTIVE Finding_State = 1
	// The finding has been fixed, triaged as a non-issue or otherwise addressed
	// and is no longer active.
	Finding_INACTIVE Finding_State = 2
)

var Finding_State_name = map[int32]string{
	0: "STATE_UNSPECIFIED",
	1: "ACTIVE",
	2: "INACTIVE",
}
var Finding_State_value = map[string]int32{
	"STATE_UNSPECIFIED": 0,
	"ACTIVE":            1,
	"INACTIVE":          2,
}

func (x Finding_State) String() string {
	return proto.EnumName(Finding_State_name, int32(x))
}
func (Finding_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_finding_e00b931465f548bd, []int{0, 0}
}

// Cloud Security Command Center (Cloud SCC) finding.
//
// A finding is a record of assessment data (security, risk, health or privacy)
// ingested into Cloud SCC for presentation, notification, analysis,
// policy testing, and enforcement. For example, an XSS vulnerability in an
// App Engine application is a finding.
type Finding struct {
	// The relative resource name of this finding. See:
	// https://cloud.google.com/apis/design/resource_names#relative_resource_name
	// Example:
	// "organizations/123/sources/456/findings/789"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The relative resource name of the source the finding belongs to. See:
	// https://cloud.google.com/apis/design/resource_names#relative_resource_name
	// This field is immutable after creation time.
	// For example:
	// "organizations/123/sources/456"
	Parent string `protobuf:"bytes,2,opt,name=parent,proto3" json:"parent,omitempty"`
	// The full resource name of the Google Cloud Platform (GCP) resource this
	// finding is for. See:
	// https://cloud.google.com/apis/design/resource_names#full_resource_name
	// This field is immutable after creation time.
	ResourceName string `protobuf:"bytes,3,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The state of the finding.
	State Finding_State `protobuf:"varint,4,opt,name=state,proto3,enum=google.cloud.securitycenter.v1.Finding_State" json:"state,omitempty"`
	// The additional taxonomy group within findings from a given source.
	// This field is immutable after creation time.
	// Example: "XSS_FLASH_INJECTION"
	Category string `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty"`
	// The URI that, if available, points to a web page outside of Cloud SCC
	// where additional information about the finding can be found. This field is
	// guaranteed to be either empty or a well formed URL.
	ExternalUri string `protobuf:"bytes,6,opt,name=external_uri,json=externalUri,proto3" json:"external_uri,omitempty"`
	// Source specific properties. These properties are managed by the source
	// that writes the finding. The key names in the source_properties map must be
	// between 1 and 255 characters, and must start with a letter and contain
	// alphanumeric characters or underscores only.
	SourceProperties map[string]*_struct.Value `protobuf:"bytes,7,rep,name=source_properties,json=sourceProperties,proto3" json:"source_properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Output only. User specified security marks. These marks are entirely
	// managed by the user and come from the SecurityMarks resource that belongs
	// to the finding.
	SecurityMarks *SecurityMarks `protobuf:"bytes,8,opt,name=security_marks,json=securityMarks,proto3" json:"security_marks,omitempty"`
	// The time at which the event took place. For example, if the finding
	// represents an open firewall it would capture the time the open firewall was
	// detected.
	EventTime *timestamp.Timestamp `protobuf:"bytes,9,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	// The time at which the finding was created in Cloud SCC.
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,10,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Finding) Reset()         { *m = Finding{} }
func (m *Finding) String() string { return proto.CompactTextString(m) }
func (*Finding) ProtoMessage()    {}
func (*Finding) Descriptor() ([]byte, []int) {
	return fileDescriptor_finding_e00b931465f548bd, []int{0}
}
func (m *Finding) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Finding.Unmarshal(m, b)
}
func (m *Finding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Finding.Marshal(b, m, deterministic)
}
func (dst *Finding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Finding.Merge(dst, src)
}
func (m *Finding) XXX_Size() int {
	return xxx_messageInfo_Finding.Size(m)
}
func (m *Finding) XXX_DiscardUnknown() {
	xxx_messageInfo_Finding.DiscardUnknown(m)
}

var xxx_messageInfo_Finding proto.InternalMessageInfo

func (m *Finding) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Finding) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *Finding) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *Finding) GetState() Finding_State {
	if m != nil {
		return m.State
	}
	return Finding_STATE_UNSPECIFIED
}

func (m *Finding) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *Finding) GetExternalUri() string {
	if m != nil {
		return m.ExternalUri
	}
	return ""
}

func (m *Finding) GetSourceProperties() map[string]*_struct.Value {
	if m != nil {
		return m.SourceProperties
	}
	return nil
}

func (m *Finding) GetSecurityMarks() *SecurityMarks {
	if m != nil {
		return m.SecurityMarks
	}
	return nil
}

func (m *Finding) GetEventTime() *timestamp.Timestamp {
	if m != nil {
		return m.EventTime
	}
	return nil
}

func (m *Finding) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func init() {
	proto.RegisterType((*Finding)(nil), "google.cloud.securitycenter.v1.Finding")
	proto.RegisterMapType((map[string]*_struct.Value)(nil), "google.cloud.securitycenter.v1.Finding.SourcePropertiesEntry")
	proto.RegisterEnum("google.cloud.securitycenter.v1.Finding_State", Finding_State_name, Finding_State_value)
}

func init() {
	proto.RegisterFile("google/cloud/securitycenter/v1/finding.proto", fileDescriptor_finding_e00b931465f548bd)
}

var fileDescriptor_finding_e00b931465f548bd = []byte{
	// 524 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x6f, 0xda, 0x40,
	0x10, 0xad, 0x49, 0x20, 0x30, 0x90, 0x88, 0xac, 0x94, 0xc8, 0x42, 0x55, 0x4a, 0xe9, 0x85, 0x43,
	0x6a, 0x0b, 0x72, 0x49, 0x1b, 0xf5, 0x90, 0x52, 0x52, 0x21, 0xb5, 0x08, 0x19, 0xc2, 0xa1, 0x3d,
	0xa0, 0x8d, 0x33, 0xb1, 0xb6, 0xc1, 0xbb, 0xd6, 0xee, 0x1a, 0x95, 0x1f, 0xd2, 0x3f, 0xd1, 0x4b,
	0xff, 0x62, 0xe5, 0x5d, 0x3b, 0x2a, 0xe9, 0x07, 0xbd, 0x79, 0xde, 0xbc, 0xf7, 0x66, 0x76, 0xdf,
	0x1a, 0x4e, 0x23, 0x21, 0xa2, 0x25, 0xfa, 0xe1, 0x52, 0xa4, 0xb7, 0xbe, 0xc2, 0x30, 0x95, 0x4c,
	0xaf, 0x43, 0xe4, 0x1a, 0xa5, 0xbf, 0xea, 0xf9, 0x77, 0x8c, 0xdf, 0x32, 0x1e, 0x79, 0x89, 0x14,
	0x5a, 0x90, 0x13, 0xcb, 0xf6, 0x0c, 0xdb, 0xdb, 0x64, 0x7b, 0xab, 0x5e, 0xeb, 0x69, 0xee, 0x46,
	0x13, 0xe6, 0x53, 0xce, 0x85, 0xa6, 0x9a, 0x09, 0xae, 0xac, 0xba, 0x75, 0xb6, 0x65, 0x56, 0x81,
	0x2c, 0x62, 0x2a, 0xef, 0x0b, 0x51, 0x61, 0x69, 0xaa, 0x9b, 0xf4, 0xce, 0x57, 0x5a, 0xa6, 0xa1,
	0xce, 0xbb, 0xcf, 0x1e, 0x77, 0x35, 0x8b, 0x51, 0x69, 0x1a, 0x27, 0x96, 0xd0, 0xf9, 0x51, 0x86,
	0xbd, 0x2b, 0x7b, 0x06, 0x42, 0x60, 0x97, 0xd3, 0x18, 0x5d, 0xa7, 0xed, 0x74, 0x6b, 0x81, 0xf9,
	0x26, 0xc7, 0x50, 0x49, 0xa8, 0x44, 0xae, 0xdd, 0x92, 0x41, 0xf3, 0x8a, 0xbc, 0x80, 0x7d, 0x89,
	0x4a, 0xa4, 0x32, 0xc4, 0x85, 0x11, 0xed, 0x98, 0x76, 0xa3, 0x00, 0xc7, 0x99, 0x78, 0x00, 0x65,
	0xa5, 0xa9, 0x46, 0x77, 0xb7, 0xed, 0x74, 0x0f, 0xfa, 0x2f, 0xbd, 0x7f, 0x5f, 0x8f, 0x97, 0x2f,
	0xe2, 0x4d, 0x33, 0x51, 0x60, 0xb5, 0xa4, 0x05, 0xd5, 0x90, 0x6a, 0x8c, 0x84, 0x5c, 0xbb, 0x65,
	0x33, 0xe4, 0xa1, 0x26, 0xcf, 0xa1, 0x81, 0x5f, 0x35, 0x4a, 0x4e, 0x97, 0x8b, 0x54, 0x32, 0xb7,
	0x62, 0xfa, 0xf5, 0x02, 0xbb, 0x96, 0x8c, 0x7c, 0x81, 0xc3, 0x7c, 0xcd, 0x44, 0x8a, 0x04, 0xa5,
	0x66, 0xa8, 0xdc, 0xbd, 0xf6, 0x4e, 0xb7, 0xde, 0x7f, 0xf3, 0xdf, 0xfb, 0x18, 0x83, 0xc9, 0x83,
	0x7e, 0xc8, 0xb5, 0x5c, 0x07, 0x4d, 0xf5, 0x08, 0x26, 0x33, 0x38, 0xd8, 0xcc, 0xc8, 0xad, 0xb6,
	0x9d, 0x6e, 0x7d, 0xfb, 0xc1, 0xa7, 0x39, 0xf2, 0x31, 0x13, 0x05, 0xfb, 0xea, 0xd7, 0x92, 0xbc,
	0x02, 0xc0, 0x15, 0x72, 0xbd, 0xc8, 0xb2, 0x73, 0x6b, 0xc6, 0xb1, 0x55, 0x38, 0x16, 0xc1, 0x7a,
	0xb3, 0x22, 0xd8, 0xa0, 0x66, 0xd8, 0x59, 0x4d, 0x2e, 0xa0, 0x1e, 0x4a, 0xa4, 0x1a, 0xad, 0x16,
	0xb6, 0x6a, 0xc1, 0xd2, 0x33, 0xa0, 0xf5, 0x19, 0x8e, 0xfe, 0x78, 0x70, 0xd2, 0x84, 0x9d, 0x7b,
	0x5c, 0xe7, 0xcf, 0x24, 0xfb, 0x24, 0xa7, 0x50, 0x5e, 0xd1, 0x65, 0x8a, 0xe6, 0x91, 0xd4, 0xfb,
	0xc7, 0xbf, 0x4d, 0x98, 0x67, 0xdd, 0xc0, 0x92, 0x5e, 0x97, 0xce, 0x9d, 0xce, 0x39, 0x94, 0x4d,
	0xca, 0xe4, 0x08, 0x0e, 0xa7, 0xb3, 0xcb, 0xd9, 0x70, 0x71, 0x3d, 0x9e, 0x4e, 0x86, 0x83, 0xd1,
	0xd5, 0x68, 0xf8, 0xae, 0xf9, 0x84, 0x00, 0x54, 0x2e, 0x07, 0xb3, 0xd1, 0x7c, 0xd8, 0x74, 0x48,
	0x03, 0xaa, 0xa3, 0x71, 0x5e, 0x95, 0xde, 0x7e, 0x73, 0xa0, 0x13, 0x8a, 0x78, 0xcb, 0x95, 0x4e,
	0x9c, 0x4f, 0x1f, 0x72, 0x46, 0x24, 0x96, 0x94, 0x47, 0x9e, 0x90, 0x91, 0x1f, 0x21, 0x37, 0x2b,
	0xf9, 0xb6, 0x45, 0x13, 0xa6, 0xfe, 0xf6, 0xb7, 0x5d, 0x6c, 0x22, 0xdf, 0x4b, 0x27, 0xef, 0xad,
	0xdd, 0xc0, 0x0c, 0x2c, 0x12, 0x1b, 0xd8, 0x81, 0xf3, 0xde, 0x4d, 0xc5, 0x38, 0x9f, 0xfd, 0x0c,
	0x00, 0x00, 0xff, 0xff, 0x9c, 0xb3, 0x08, 0xec, 0x32, 0x04, 0x00, 0x00,
}
