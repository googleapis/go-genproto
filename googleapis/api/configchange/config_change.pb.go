// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/api/configchange/config_change.proto
// DO NOT EDIT!

/*
Package google_api is a generated protocol buffer package.

It is generated from these files:
	google.golang.org/genproto/googleapis/api/configchange/config_change.proto

It has these top-level messages:
	ConfigChange
	Advice
*/
package google_api // import "google.golang.org/genproto/googleapis/api/configchange"

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

// Classifies set of possible modifications to an object in the service
// configuration.
type ChangeType int32

const (
	// No value was provided.
	ChangeType_CHANGE_TYPE_UNSPECIFIED ChangeType = 0
	// The changed object exists in the 'new' service configuration, but not
	// in the 'old' service configuration.
	ChangeType_ADDED ChangeType = 1
	// The changed object exists in the 'old' service configuration, but not
	// in the 'new' service configuration.
	ChangeType_REMOVED ChangeType = 2
	// The changed object exists in both service configurations, but its value
	// is different.
	ChangeType_MODIFIED ChangeType = 3
)

var ChangeType_name = map[int32]string{
	0: "CHANGE_TYPE_UNSPECIFIED",
	1: "ADDED",
	2: "REMOVED",
	3: "MODIFIED",
}
var ChangeType_value = map[string]int32{
	"CHANGE_TYPE_UNSPECIFIED": 0,
	"ADDED":                   1,
	"REMOVED":                 2,
	"MODIFIED":                3,
}

func (x ChangeType) String() string {
	return proto.EnumName(ChangeType_name, int32(x))
}
func (ChangeType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Output generated from semantically comparing two versions of a service
// configuration.
//
// Includes detailed information about a field that have changed with
// applicable advice about potential consequences for the change, such as
// backwards-incompatibility.
type ConfigChange struct {
	// Object hierarchy path to the change, with levels separated by a '.'
	// character. For repeated fields, an applicable unique identifier field is
	// used for the index (usually selector, name, or id). For maps, the term
	// 'key' is used. If the field has no unique identifier, the numeric index
	// is used.
	// Examples:
	// - visibility.rules[selector=="google.LibraryService.CreateBook"].restriction
	// - quota.metric_rules[selector=="google"].metric_costs[key=="reads"].value
	// - logging.producer_destinations[0]
	Element string `protobuf:"bytes,1,opt,name=element" json:"element,omitempty"`
	// Value of the changed object in the old Service configuration,
	// in JSON format. This field will not be populated if ChangeType == ADDED.
	OldValue string `protobuf:"bytes,2,opt,name=old_value,json=oldValue" json:"old_value,omitempty"`
	// Value of the changed object in the new Service configuration,
	// in JSON format. This field will not be populated if ChangeType == REMOVED.
	NewValue string `protobuf:"bytes,3,opt,name=new_value,json=newValue" json:"new_value,omitempty"`
	// The type for this change, either ADDED, REMOVED, or MODIFIED.
	ChangeType ChangeType `protobuf:"varint,4,opt,name=change_type,json=changeType,enum=google.api.ChangeType" json:"change_type,omitempty"`
	// Collection of advice provided for this change, useful for determining the
	// possible impact of this change.
	Advices []*Advice `protobuf:"bytes,5,rep,name=advices" json:"advices,omitempty"`
}

func (m *ConfigChange) Reset()                    { *m = ConfigChange{} }
func (m *ConfigChange) String() string            { return proto.CompactTextString(m) }
func (*ConfigChange) ProtoMessage()               {}
func (*ConfigChange) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ConfigChange) GetAdvices() []*Advice {
	if m != nil {
		return m.Advices
	}
	return nil
}

// Generated advice about this change, used for providing more
// information about how a change will affect the existing service.
type Advice struct {
	// Useful description for why this advice was applied and what actions should
	// be taken to mitigate any implied risks.
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *Advice) Reset()                    { *m = Advice{} }
func (m *Advice) String() string            { return proto.CompactTextString(m) }
func (*Advice) ProtoMessage()               {}
func (*Advice) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*ConfigChange)(nil), "google.api.ConfigChange")
	proto.RegisterType((*Advice)(nil), "google.api.Advice")
	proto.RegisterEnum("google.api.ChangeType", ChangeType_name, ChangeType_value)
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/api/configchange/config_change.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0x4f, 0x6b, 0xc2, 0x30,
	0x18, 0xc6, 0x57, 0xff, 0xfb, 0x56, 0xc4, 0xe5, 0xb0, 0x15, 0xbc, 0x14, 0x4f, 0x22, 0xd2, 0x82,
	0x3b, 0xec, 0x5c, 0xdb, 0xce, 0x39, 0x50, 0x4b, 0xe7, 0x84, 0x9d, 0x4a, 0xd6, 0x66, 0x59, 0xa0,
	0x26, 0x41, 0x3b, 0xc5, 0xaf, 0xb3, 0x6f, 0xb3, 0x6f, 0x35, 0x6c, 0xea, 0xec, 0x2d, 0xcf, 0xfb,
	0x7b, 0x92, 0xf7, 0xe1, 0x09, 0xbc, 0x50, 0x21, 0x68, 0x4a, 0x2c, 0x2a, 0x52, 0xcc, 0xa9, 0x25,
	0x76, 0xd4, 0xa6, 0x84, 0xcb, 0x9d, 0xc8, 0x84, 0xad, 0x10, 0x96, 0x6c, 0x6f, 0x63, 0xc9, 0xec,
	0x58, 0xf0, 0x4f, 0x46, 0xe3, 0x2f, 0xcc, 0x29, 0x29, 0x44, 0xa4, 0x94, 0x95, 0xfb, 0x11, 0x14,
	0x6f, 0x61, 0xc9, 0x06, 0xbf, 0x1a, 0x74, 0xdc, 0xdc, 0xe3, 0xe6, 0x16, 0x64, 0x40, 0x93, 0xa4,
	0x64, 0x4b, 0x78, 0x66, 0x68, 0xa6, 0x36, 0x6c, 0x87, 0x17, 0x89, 0xfa, 0xd0, 0x16, 0x69, 0x12,
	0x1d, 0x70, 0xfa, 0x4d, 0x8c, 0x4a, 0xce, 0x5a, 0x22, 0x4d, 0x36, 0x67, 0x7d, 0x86, 0x9c, 0x1c,
	0x0b, 0x58, 0x55, 0x90, 0x93, 0xa3, 0x82, 0x8f, 0xa0, 0xab, 0x00, 0x51, 0x76, 0x92, 0xc4, 0xa8,
	0x99, 0xda, 0xb0, 0x3b, 0xb9, 0xb3, 0xae, 0x31, 0x2c, 0xb5, 0x7c, 0x7d, 0x92, 0x24, 0x84, 0xf8,
	0xff, 0x8c, 0xc6, 0xd0, 0xc4, 0xc9, 0x81, 0xc5, 0x64, 0x6f, 0xd4, 0xcd, 0xea, 0x50, 0x9f, 0xa0,
	0xf2, 0x25, 0x27, 0x47, 0xe1, 0xc5, 0x32, 0x18, 0x41, 0x43, 0x8d, 0x90, 0x09, 0x7a, 0x42, 0xf6,
	0xf1, 0x8e, 0xc9, 0x8c, 0x09, 0x5e, 0x84, 0x2d, 0x8f, 0x46, 0x2b, 0x80, 0xeb, 0x4e, 0xd4, 0x87,
	0x7b, 0xf7, 0xd9, 0x59, 0xce, 0xfc, 0x68, 0xfd, 0x1e, 0xf8, 0xd1, 0xdb, 0xf2, 0x35, 0xf0, 0xdd,
	0xf9, 0xd3, 0xdc, 0xf7, 0x7a, 0x37, 0xa8, 0x0d, 0x75, 0xc7, 0xf3, 0x7c, 0xaf, 0xa7, 0x21, 0x1d,
	0x9a, 0xa1, 0xbf, 0x58, 0x6d, 0x7c, 0xaf, 0x57, 0x41, 0x1d, 0x68, 0x2d, 0x56, 0x9e, 0x72, 0x55,
	0xa7, 0x63, 0xe8, 0xc6, 0x62, 0x5b, 0x8a, 0x37, 0xbd, 0x2d, 0xf7, 0x1a, 0x9c, 0x9b, 0x0f, 0xb4,
	0x9f, 0x4a, 0x6d, 0xe6, 0x04, 0xf3, 0x8f, 0x46, 0xfe, 0x13, 0x0f, 0x7f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x7c, 0x7e, 0x38, 0xa2, 0xd7, 0x01, 0x00, 0x00,
}
