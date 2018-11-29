// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/securitycenter/v1beta1/finding.proto

package securitycenter

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
	return fileDescriptor_da6ee7073afacba3, []int{0, 0}
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
	State Finding_State `protobuf:"varint,4,opt,name=state,proto3,enum=google.cloud.securitycenter.v1beta1.Finding_State" json:"state,omitempty"`
	// The additional taxonomy group within findings from a give source.
	// This field is immutable after creation time.
	// Example: "XSS_FLASH_INJECTION"
	Category string `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty"`
	// The URI that, if available, points to a web page outside of Cloud SCC
	// where additional information about the finding can be found. This field is
	// guaranteed to be either empty or a well formed URL.
	ExternalUri string `protobuf:"bytes,6,opt,name=external_uri,json=externalUri,proto3" json:"external_uri,omitempty"`
	// Source specific properties. These properties are managed by the source
	// that writes the finding.
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
	return fileDescriptor_da6ee7073afacba3, []int{0}
}

func (m *Finding) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Finding.Unmarshal(m, b)
}
func (m *Finding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Finding.Marshal(b, m, deterministic)
}
func (m *Finding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Finding.Merge(m, src)
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
	proto.RegisterEnum("google.cloud.securitycenter.v1beta1.Finding_State", Finding_State_name, Finding_State_value)
	proto.RegisterType((*Finding)(nil), "google.cloud.securitycenter.v1beta1.Finding")
	proto.RegisterMapType((map[string]*_struct.Value)(nil), "google.cloud.securitycenter.v1beta1.Finding.SourcePropertiesEntry")
}

func init() {
	proto.RegisterFile("google/cloud/securitycenter/v1beta1/finding.proto", fileDescriptor_da6ee7073afacba3)
}

var fileDescriptor_da6ee7073afacba3 = []byte{
	// 510 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x41, 0x6f, 0xda, 0x30,
	0x14, 0xc7, 0x17, 0x28, 0x14, 0x1e, 0xb4, 0xa2, 0x96, 0x5a, 0x45, 0x68, 0xd2, 0x58, 0x7b, 0x18,
	0x87, 0x29, 0x11, 0xec, 0xc2, 0xd6, 0x53, 0xdb, 0x51, 0x8d, 0xc3, 0x10, 0x0b, 0xb4, 0xd2, 0xb6,
	0x03, 0x32, 0xe9, 0x6b, 0x64, 0x35, 0xb1, 0x23, 0xdb, 0x41, 0xe3, 0xb2, 0x8f, 0xb6, 0xcf, 0x36,
	0xc5, 0x71, 0xaa, 0xd1, 0x4d, 0x1b, 0xbb, 0xe5, 0xfd, 0xdf, 0xfb, 0xfd, 0xfd, 0xfc, 0x9e, 0x03,
	0x83, 0x48, 0x88, 0x28, 0x46, 0x3f, 0x8c, 0x45, 0x76, 0xe7, 0x2b, 0x0c, 0x33, 0xc9, 0xf4, 0x26,
	0x44, 0xae, 0x51, 0xfa, 0xeb, 0xc1, 0x0a, 0x35, 0x1d, 0xf8, 0xf7, 0x8c, 0xdf, 0x31, 0x1e, 0x79,
	0xa9, 0x14, 0x5a, 0x90, 0xb3, 0x02, 0xf1, 0x0c, 0xe2, 0x6d, 0x23, 0x9e, 0x45, 0xba, 0xcf, 0xad,
	0x2f, 0x4d, 0x99, 0x4f, 0x39, 0x17, 0x9a, 0x6a, 0x26, 0xb8, 0x2a, 0x2c, 0xba, 0xa3, 0x5d, 0x4e,
	0x2d, 0xe5, 0x65, 0x42, 0xe5, 0x43, 0x49, 0x96, 0xbe, 0x26, 0x5a, 0x65, 0xf7, 0xbe, 0xd2, 0x32,
	0x0b, 0xb5, 0xcd, 0xbe, 0x78, 0x9a, 0xd5, 0x2c, 0x41, 0xa5, 0x69, 0x92, 0x16, 0x05, 0xa7, 0x3f,
	0x6a, 0xb0, 0x7f, 0x5d, 0xdc, 0x86, 0x10, 0xd8, 0xe3, 0x34, 0x41, 0xd7, 0xe9, 0x39, 0xfd, 0x66,
	0x60, 0xbe, 0xc9, 0x09, 0xd4, 0x53, 0x2a, 0x91, 0x6b, 0xb7, 0x62, 0x54, 0x1b, 0x91, 0x33, 0x38,
	0x90, 0xa8, 0x44, 0x26, 0x43, 0x5c, 0x1a, 0xa8, 0x6a, 0xd2, 0xed, 0x52, 0x9c, 0xe6, 0xf0, 0x07,
	0xa8, 0x29, 0x4d, 0x35, 0xba, 0x7b, 0x3d, 0xa7, 0x7f, 0x38, 0x1c, 0x7a, 0x3b, 0x0c, 0xca, 0xb3,
	0xdd, 0x78, 0xf3, 0x9c, 0x0c, 0x0a, 0x03, 0xd2, 0x85, 0x46, 0x48, 0x35, 0x46, 0x42, 0x6e, 0xdc,
	0x9a, 0x39, 0xe9, 0x31, 0x26, 0x2f, 0xa1, 0x8d, 0xdf, 0x34, 0x4a, 0x4e, 0xe3, 0x65, 0x26, 0x99,
	0x5b, 0x37, 0xf9, 0x56, 0xa9, 0xdd, 0x48, 0x46, 0x04, 0x1c, 0xd9, 0x5e, 0x53, 0x29, 0x52, 0x94,
	0x9a, 0xa1, 0x72, 0xf7, 0x7b, 0xd5, 0x7e, 0x6b, 0x78, 0xf9, 0x7f, 0x4d, 0x19, 0x97, 0xd9, 0xa3,
	0xc9, 0x98, 0x6b, 0xb9, 0x09, 0x3a, 0xea, 0x89, 0x4c, 0x3e, 0xc3, 0xe1, 0xf6, 0xb6, 0xdc, 0x46,
	0xcf, 0xe9, 0xb7, 0x76, 0x1c, 0xc1, 0xdc, 0xca, 0x1f, 0x73, 0x32, 0x38, 0x50, 0xbf, 0x86, 0xe4,
	0x2d, 0x00, 0xae, 0x91, 0xeb, 0x65, 0xbe, 0x4a, 0xb7, 0x69, 0x6c, 0xbb, 0xa5, 0x6d, 0xb9, 0x67,
	0x6f, 0x51, 0xee, 0x39, 0x68, 0x9a, 0xea, 0x3c, 0x26, 0xe7, 0xd0, 0x0a, 0x25, 0x52, 0x8d, 0x05,
	0x0b, 0xff, 0x64, 0xa1, 0x28, 0xcf, 0x85, 0xee, 0x57, 0x38, 0xfe, 0xe3, 0xed, 0x49, 0x07, 0xaa,
	0x0f, 0xb8, 0xb1, 0xaf, 0x26, 0xff, 0x24, 0xaf, 0xa1, 0xb6, 0xa6, 0x71, 0x86, 0xe6, 0xcd, 0xb4,
	0x86, 0x27, 0xbf, 0x9d, 0x70, 0x9b, 0x67, 0x83, 0xa2, 0xe8, 0x5d, 0x65, 0xe4, 0x9c, 0x8e, 0xa0,
	0x66, 0xf6, 0x4d, 0x8e, 0xe1, 0x68, 0xbe, 0xb8, 0x58, 0x8c, 0x97, 0x37, 0xd3, 0xf9, 0x6c, 0x7c,
	0x35, 0xb9, 0x9e, 0x8c, 0xdf, 0x77, 0x9e, 0x11, 0x80, 0xfa, 0xc5, 0xd5, 0x62, 0x72, 0x3b, 0xee,
	0x38, 0xa4, 0x0d, 0x8d, 0xc9, 0xd4, 0x46, 0x95, 0xcb, 0xef, 0xf0, 0x2a, 0x14, 0xc9, 0x2e, 0x63,
	0x9d, 0x39, 0x5f, 0x3e, 0xd9, 0xb2, 0x48, 0xc4, 0x94, 0x47, 0x9e, 0x90, 0x91, 0x1f, 0x21, 0x37,
	0x6d, 0xf9, 0x45, 0x8a, 0xa6, 0x4c, 0xfd, 0xf5, 0x2f, 0x3c, 0xdf, 0x96, 0x57, 0x75, 0x43, 0xbf,
	0xf9, 0x19, 0x00, 0x00, 0xff, 0xff, 0xf7, 0xda, 0xfd, 0x34, 0x38, 0x04, 0x00, 0x00,
}
