// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/bigquery/storage/v1beta1/table_reference.proto

package storage

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Table reference that includes just the 3 strings needed to identify a table.
type TableReference struct {
	// The assigned project ID of the project.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// The ID of the dataset in the above project.
	DatasetId string `protobuf:"bytes,2,opt,name=dataset_id,json=datasetId,proto3" json:"dataset_id,omitempty"`
	// The ID of the table in the above dataset.
	TableId              string   `protobuf:"bytes,3,opt,name=table_id,json=tableId,proto3" json:"table_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TableReference) Reset()         { *m = TableReference{} }
func (m *TableReference) String() string { return proto.CompactTextString(m) }
func (*TableReference) ProtoMessage()    {}
func (*TableReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_00fafb38a9ed74ff, []int{0}
}

func (m *TableReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TableReference.Unmarshal(m, b)
}
func (m *TableReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TableReference.Marshal(b, m, deterministic)
}
func (m *TableReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TableReference.Merge(m, src)
}
func (m *TableReference) XXX_Size() int {
	return xxx_messageInfo_TableReference.Size(m)
}
func (m *TableReference) XXX_DiscardUnknown() {
	xxx_messageInfo_TableReference.DiscardUnknown(m)
}

var xxx_messageInfo_TableReference proto.InternalMessageInfo

func (m *TableReference) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *TableReference) GetDatasetId() string {
	if m != nil {
		return m.DatasetId
	}
	return ""
}

func (m *TableReference) GetTableId() string {
	if m != nil {
		return m.TableId
	}
	return ""
}

// All fields in this message optional.
type TableModifiers struct {
	// The snapshot time of the table. If not set, interpreted as now.
	SnapshotTime         *timestamp.Timestamp `protobuf:"bytes,1,opt,name=snapshot_time,json=snapshotTime,proto3" json:"snapshot_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TableModifiers) Reset()         { *m = TableModifiers{} }
func (m *TableModifiers) String() string { return proto.CompactTextString(m) }
func (*TableModifiers) ProtoMessage()    {}
func (*TableModifiers) Descriptor() ([]byte, []int) {
	return fileDescriptor_00fafb38a9ed74ff, []int{1}
}

func (m *TableModifiers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TableModifiers.Unmarshal(m, b)
}
func (m *TableModifiers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TableModifiers.Marshal(b, m, deterministic)
}
func (m *TableModifiers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TableModifiers.Merge(m, src)
}
func (m *TableModifiers) XXX_Size() int {
	return xxx_messageInfo_TableModifiers.Size(m)
}
func (m *TableModifiers) XXX_DiscardUnknown() {
	xxx_messageInfo_TableModifiers.DiscardUnknown(m)
}

var xxx_messageInfo_TableModifiers proto.InternalMessageInfo

func (m *TableModifiers) GetSnapshotTime() *timestamp.Timestamp {
	if m != nil {
		return m.SnapshotTime
	}
	return nil
}

func init() {
	proto.RegisterType((*TableReference)(nil), "google.cloud.bigquery.storage.v1beta1.TableReference")
	proto.RegisterType((*TableModifiers)(nil), "google.cloud.bigquery.storage.v1beta1.TableModifiers")
}

func init() {
	proto.RegisterFile("google/cloud/bigquery/storage/v1beta1/table_reference.proto", fileDescriptor_00fafb38a9ed74ff)
}

var fileDescriptor_00fafb38a9ed74ff = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0xa9, 0x82, 0xba, 0xf1, 0xcf, 0xa1, 0x5e, 0xd6, 0x05, 0x51, 0x16, 0x04, 0xbd, 0x24,
	0xac, 0x1e, 0xf7, 0x20, 0xec, 0xad, 0xa0, 0xa0, 0x65, 0x4f, 0x5e, 0x4a, 0xda, 0x4c, 0x63, 0xa4,
	0xed, 0xd4, 0x24, 0x15, 0xfc, 0x12, 0x7e, 0x66, 0xc9, 0xbf, 0x83, 0x27, 0xf7, 0x38, 0xf3, 0xf8,
	0xbd, 0xf7, 0x26, 0x21, 0x6b, 0x89, 0x28, 0x3b, 0x60, 0x4d, 0x87, 0x93, 0x60, 0xb5, 0x92, 0x9f,
	0x13, 0xe8, 0x6f, 0x66, 0x2c, 0x6a, 0x2e, 0x81, 0x7d, 0xad, 0x6a, 0xb0, 0x7c, 0xc5, 0x2c, 0xaf,
	0x3b, 0xa8, 0x34, 0xb4, 0xa0, 0x61, 0x68, 0x80, 0x8e, 0x1a, 0x2d, 0xe6, 0x37, 0x01, 0xa6, 0x1e,
	0xa6, 0x09, 0xa6, 0x11, 0xa6, 0x11, 0x5e, 0x5c, 0xc5, 0x0c, 0x0f, 0xd5, 0x53, 0xcb, 0xac, 0xea,
	0xc1, 0x58, 0xde, 0x8f, 0xc1, 0x67, 0xa9, 0xc8, 0xd9, 0xd6, 0x05, 0x94, 0xc9, 0x3f, 0xbf, 0x24,
	0x64, 0xd4, 0xf8, 0x01, 0x8d, 0xad, 0x94, 0x98, 0x67, 0xd7, 0xd9, 0xed, 0xac, 0x9c, 0xc5, 0x4d,
	0x21, 0x9c, 0x2c, 0xb8, 0xe5, 0x06, 0xbc, 0xbc, 0x17, 0xe4, 0xb8, 0x29, 0x44, 0x7e, 0x41, 0x8e,
	0x42, 0x61, 0x25, 0xe6, 0xfb, 0x5e, 0x3c, 0xf4, 0x73, 0x21, 0x96, 0xaf, 0x31, 0xea, 0x19, 0x85,
	0x6a, 0x15, 0x68, 0x93, 0x3f, 0x92, 0x53, 0x33, 0xf0, 0xd1, 0xbc, 0xa3, 0xad, 0x5c, 0x31, 0x9f,
	0x76, 0x7c, 0xbf, 0xa0, 0xf1, 0xb8, 0xd4, 0x9a, 0x6e, 0x53, 0xeb, 0xf2, 0x24, 0x01, 0x6e, 0xb5,
	0xf9, 0xc9, 0xc8, 0x5d, 0x83, 0x3d, 0xdd, 0xe9, 0x31, 0x36, 0xe7, 0x7f, 0x2f, 0x7d, 0x71, 0xee,
	0x6f, 0x4f, 0x91, 0x95, 0xd8, 0xf1, 0x41, 0x52, 0xd4, 0x92, 0x49, 0x18, 0x7c, 0x32, 0x0b, 0x12,
	0x1f, 0x95, 0xf9, 0xe7, 0x93, 0xd6, 0x71, 0xae, 0x0f, 0x3c, 0xf8, 0xf0, 0x1b, 0x00, 0x00, 0xff,
	0xff, 0xb3, 0xa1, 0xcb, 0xc3, 0xdc, 0x01, 0x00, 0x00,
}
