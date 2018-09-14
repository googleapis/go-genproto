// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/monitoring/v3/group.proto

package monitoring

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The description of a dynamic collection of monitored resources. Each group
// has a filter that is matched against monitored resources and their associated
// metadata. If a group's filter matches an available monitored resource, then
// that resource is a member of that group.  Groups can contain any number of
// monitored resources, and each monitored resource can be a member of any
// number of groups.
//
// Groups can be nested in parent-child hierarchies. The `parentName` field
// identifies an optional parent for each group.  If a group has a parent, then
// the only monitored resources available to be matched by the group's filter
// are the resources contained in the parent group.  In other words, a group
// contains the monitored resources that match its filter and the filters of all
// the group's ancestors.  A group without a parent can contain any monitored
// resource.
//
// For example, consider an infrastructure running a set of instances with two
// user-defined tags: `"environment"` and `"role"`. A parent group has a filter,
// `environment="production"`.  A child of that parent group has a filter,
// `role="transcoder"`.  The parent group contains all instances in the
// production environment, regardless of their roles.  The child group contains
// instances that have the transcoder role *and* are in the production
// environment.
//
// The monitored resources contained in a group can change at any moment,
// depending on what resources exist and what filters are associated with the
// group and its ancestors.
type Group struct {
	// Output only. The name of this group. The format is
	// `"projects/{project_id_or_number}/groups/{group_id}"`.
	// When creating a group, this field is ignored and a new name is created
	// consisting of the project specified in the call to `CreateGroup`
	// and a unique `{group_id}` that is generated automatically.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A user-assigned name for this group, used only for display purposes.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The name of the group's parent, if it has one.
	// The format is `"projects/{project_id_or_number}/groups/{group_id}"`.
	// For groups with no parent, `parentName` is the empty string, `""`.
	ParentName string `protobuf:"bytes,3,opt,name=parent_name,json=parentName,proto3" json:"parent_name,omitempty"`
	// The filter used to determine which monitored resources belong to this group.
	Filter string `protobuf:"bytes,5,opt,name=filter,proto3" json:"filter,omitempty"`
	// If true, the members of this group are considered to be a cluster.
	// The system can perform additional analysis on groups that are clusters.
	IsCluster            bool     `protobuf:"varint,6,opt,name=is_cluster,json=isCluster,proto3" json:"is_cluster,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Group) Reset()         { *m = Group{} }
func (m *Group) String() string { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()    {}
func (*Group) Descriptor() ([]byte, []int) {
	return fileDescriptor_907e30c1f087271d, []int{0}
}
func (m *Group) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Group.Unmarshal(m, b)
}
func (m *Group) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Group.Marshal(b, m, deterministic)
}
func (m *Group) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Group.Merge(m, src)
}
func (m *Group) XXX_Size() int {
	return xxx_messageInfo_Group.Size(m)
}
func (m *Group) XXX_DiscardUnknown() {
	xxx_messageInfo_Group.DiscardUnknown(m)
}

var xxx_messageInfo_Group proto.InternalMessageInfo

func (m *Group) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Group) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Group) GetParentName() string {
	if m != nil {
		return m.ParentName
	}
	return ""
}

func (m *Group) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *Group) GetIsCluster() bool {
	if m != nil {
		return m.IsCluster
	}
	return false
}

func init() {
	proto.RegisterType((*Group)(nil), "google.monitoring.v3.Group")
}

func init() { proto.RegisterFile("google/monitoring/v3/group.proto", fileDescriptor_907e30c1f087271d) }

var fileDescriptor_907e30c1f087271d = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcf, 0x4a, 0x2b, 0x31,
	0x14, 0x87, 0x49, 0xef, 0xed, 0x60, 0x4f, 0x5d, 0x0d, 0x22, 0x83, 0x20, 0x8e, 0xae, 0xba, 0xca,
	0x2c, 0xb2, 0x14, 0x5c, 0xb4, 0x8b, 0xae, 0x94, 0xd2, 0x45, 0x17, 0x32, 0x50, 0x62, 0x1b, 0x43,
	0x20, 0x93, 0x13, 0x92, 0x99, 0x82, 0x2f, 0xe2, 0x03, 0xb8, 0xf4, 0x51, 0x7c, 0x2a, 0x99, 0x93,
	0x91, 0x41, 0x70, 0x97, 0xf3, 0xfb, 0x3e, 0x72, 0xfe, 0x40, 0xa9, 0x11, 0xb5, 0x55, 0x55, 0x83,
	0xce, 0xb4, 0x18, 0x8c, 0xd3, 0xd5, 0x49, 0x54, 0x3a, 0x60, 0xe7, 0xb9, 0x0f, 0xd8, 0x62, 0x7e,
	0x91, 0x0c, 0x3e, 0x1a, 0xfc, 0x24, 0xee, 0xde, 0x19, 0x4c, 0xd7, 0xbd, 0x95, 0xe7, 0xf0, 0xdf,
	0xc9, 0x46, 0x15, 0xac, 0x64, 0x8b, 0xd9, 0x96, 0xde, 0xf9, 0x2d, 0x9c, 0x1f, 0x4d, 0xf4, 0x56,
	0xbe, 0xed, 0x89, 0x4d, 0x88, 0xcd, 0x87, 0xec, 0xa9, 0x57, 0x6e, 0x60, 0xee, 0x65, 0x50, 0xae,
	0x4d, 0xc6, 0x3f, 0x32, 0x20, 0x45, 0x24, 0x5c, 0x42, 0xf6, 0x6a, 0x6c, 0xab, 0x42, 0x31, 0x25,
	0x36, 0x54, 0xf9, 0x35, 0x80, 0x89, 0xfb, 0x83, 0xed, 0x62, 0xcf, 0xb2, 0x92, 0x2d, 0xce, 0xb6,
	0x33, 0x13, 0x57, 0x29, 0x58, 0x7e, 0x30, 0x28, 0x0e, 0xd8, 0xf0, 0xbf, 0xa6, 0x5e, 0x02, 0x8d,
	0xbc, 0xe9, 0xf7, 0xda, 0xb0, 0xe7, 0x87, 0xc1, 0xd1, 0x68, 0xa5, 0xd3, 0x1c, 0x83, 0xae, 0xb4,
	0x72, 0xb4, 0x75, 0x95, 0x90, 0xf4, 0x26, 0xfe, 0x3e, 0xcd, 0xfd, 0x58, 0x7d, 0x4e, 0xae, 0xd6,
	0xe9, 0x83, 0x95, 0xc5, 0xee, 0xc8, 0x1f, 0xc7, 0x56, 0x3b, 0xf1, 0xf5, 0x03, 0x6b, 0x82, 0xf5,
	0x08, 0xeb, 0x9d, 0x78, 0xc9, 0xa8, 0x89, 0xf8, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x95, 0xd1, 0xa1,
	0x34, 0x7e, 0x01, 0x00, 0x00,
}
