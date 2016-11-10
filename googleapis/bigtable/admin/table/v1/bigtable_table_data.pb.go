// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/bigtable/admin/table/v1/bigtable_table_data.proto
// DO NOT EDIT!

/*
Package google_bigtable_admin_table_v1 is a generated protocol buffer package.

It is generated from these files:
	google.golang.org/genproto/googleapis/bigtable/admin/table/v1/bigtable_table_data.proto
	google.golang.org/genproto/googleapis/bigtable/admin/table/v1/bigtable_table_service_messages.proto
	google.golang.org/genproto/googleapis/bigtable/admin/table/v1/bigtable_table_service.proto

It has these top-level messages:
	Table
	ColumnFamily
	GcRule
	CreateTableRequest
	ListTablesRequest
	ListTablesResponse
	GetTableRequest
	DeleteTableRequest
	RenameTableRequest
	CreateColumnFamilyRequest
	DeleteColumnFamilyRequest
	BulkDeleteRowsRequest
*/
package google_bigtable_admin_table_v1 // import "google.golang.org/genproto/googleapis/bigtable/admin/table/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_longrunning "google.golang.org/genproto/googleapis/longrunning"
import google_protobuf3 "github.com/golang/protobuf/ptypes/duration"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Table_TimestampGranularity int32

const (
	Table_MILLIS Table_TimestampGranularity = 0
)

var Table_TimestampGranularity_name = map[int32]string{
	0: "MILLIS",
}
var Table_TimestampGranularity_value = map[string]int32{
	"MILLIS": 0,
}

func (x Table_TimestampGranularity) String() string {
	return proto.EnumName(Table_TimestampGranularity_name, int32(x))
}
func (Table_TimestampGranularity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

// A collection of user data indexed by row, column, and timestamp.
// Each table is served using the resources of its parent cluster.
type Table struct {
	// A unique identifier of the form
	// <cluster_name>/tables/[_a-zA-Z0-9][-_.a-zA-Z0-9]*
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// If this Table is in the process of being created, the Operation used to
	// track its progress. As long as this operation is present, the Table will
	// not accept any Table Admin or Read/Write requests.
	CurrentOperation *google_longrunning.Operation `protobuf:"bytes,2,opt,name=current_operation,json=currentOperation" json:"current_operation,omitempty"`
	// The column families configured for this table, mapped by column family id.
	ColumnFamilies map[string]*ColumnFamily `protobuf:"bytes,3,rep,name=column_families,json=columnFamilies" json:"column_families,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// The granularity (e.g. MILLIS, MICROS) at which timestamps are stored in
	// this table. Timestamps not matching the granularity will be rejected.
	// Cannot be changed once the table is created.
	Granularity Table_TimestampGranularity `protobuf:"varint,4,opt,name=granularity,enum=google.bigtable.admin.table.v1.Table_TimestampGranularity" json:"granularity,omitempty"`
}

func (m *Table) Reset()                    { *m = Table{} }
func (m *Table) String() string            { return proto.CompactTextString(m) }
func (*Table) ProtoMessage()               {}
func (*Table) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Table) GetCurrentOperation() *google_longrunning.Operation {
	if m != nil {
		return m.CurrentOperation
	}
	return nil
}

func (m *Table) GetColumnFamilies() map[string]*ColumnFamily {
	if m != nil {
		return m.ColumnFamilies
	}
	return nil
}

// A set of columns within a table which share a common configuration.
type ColumnFamily struct {
	// A unique identifier of the form <table_name>/columnFamilies/[-_.a-zA-Z0-9]+
	// The last segment is the same as the "name" field in
	// google.bigtable.v1.Family.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Garbage collection expression specified by the following grammar:
	//   GC = EXPR
	//      | "" ;
	//   EXPR = EXPR, "||", EXPR              (* lowest precedence *)
	//        | EXPR, "&&", EXPR
	//        | "(", EXPR, ")"                (* highest precedence *)
	//        | PROP ;
	//   PROP = "version() >", NUM32
	//        | "age() >", NUM64, [ UNIT ] ;
	//   NUM32 = non-zero-digit { digit } ;    (* # NUM32 <= 2^32 - 1 *)
	//   NUM64 = non-zero-digit { digit } ;    (* # NUM64 <= 2^63 - 1 *)
	//   UNIT =  "d" | "h" | "m"  (* d=days, h=hours, m=minutes, else micros *)
	// GC expressions can be up to 500 characters in length
	//
	// The different types of PROP are defined as follows:
	//   version() - cell index, counting from most recent and starting at 1
	//   age() - age of the cell (current time minus cell timestamp)
	//
	// Example: "version() > 3 || (age() > 3d && version() > 1)"
	//   drop cells beyond the most recent three, and drop cells older than three
	//   days unless they're the most recent cell in the row/column
	//
	// Garbage collection executes opportunistically in the background, and so
	// it's possible for reads to return a cell even if it matches the active GC
	// expression for its family.
	GcExpression string `protobuf:"bytes,2,opt,name=gc_expression,json=gcExpression" json:"gc_expression,omitempty"`
	// Garbage collection rule specified as a protobuf.
	// Supersedes `gc_expression`.
	// Must serialize to at most 500 bytes.
	//
	// NOTE: Garbage collection executes opportunistically in the background, and
	// so it's possible for reads to return a cell even if it matches the active
	// GC expression for its family.
	GcRule *GcRule `protobuf:"bytes,3,opt,name=gc_rule,json=gcRule" json:"gc_rule,omitempty"`
}

func (m *ColumnFamily) Reset()                    { *m = ColumnFamily{} }
func (m *ColumnFamily) String() string            { return proto.CompactTextString(m) }
func (*ColumnFamily) ProtoMessage()               {}
func (*ColumnFamily) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ColumnFamily) GetGcRule() *GcRule {
	if m != nil {
		return m.GcRule
	}
	return nil
}

// Rule for determining which cells to delete during garbage collection.
type GcRule struct {
	// Types that are valid to be assigned to Rule:
	//	*GcRule_MaxNumVersions
	//	*GcRule_MaxAge
	//	*GcRule_Intersection_
	//	*GcRule_Union_
	Rule isGcRule_Rule `protobuf_oneof:"rule"`
}

func (m *GcRule) Reset()                    { *m = GcRule{} }
func (m *GcRule) String() string            { return proto.CompactTextString(m) }
func (*GcRule) ProtoMessage()               {}
func (*GcRule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isGcRule_Rule interface {
	isGcRule_Rule()
}

type GcRule_MaxNumVersions struct {
	MaxNumVersions int32 `protobuf:"varint,1,opt,name=max_num_versions,json=maxNumVersions,oneof"`
}
type GcRule_MaxAge struct {
	MaxAge *google_protobuf3.Duration `protobuf:"bytes,2,opt,name=max_age,json=maxAge,oneof"`
}
type GcRule_Intersection_ struct {
	Intersection *GcRule_Intersection `protobuf:"bytes,3,opt,name=intersection,oneof"`
}
type GcRule_Union_ struct {
	Union *GcRule_Union `protobuf:"bytes,4,opt,name=union,oneof"`
}

func (*GcRule_MaxNumVersions) isGcRule_Rule() {}
func (*GcRule_MaxAge) isGcRule_Rule()         {}
func (*GcRule_Intersection_) isGcRule_Rule()  {}
func (*GcRule_Union_) isGcRule_Rule()         {}

func (m *GcRule) GetRule() isGcRule_Rule {
	if m != nil {
		return m.Rule
	}
	return nil
}

func (m *GcRule) GetMaxNumVersions() int32 {
	if x, ok := m.GetRule().(*GcRule_MaxNumVersions); ok {
		return x.MaxNumVersions
	}
	return 0
}

func (m *GcRule) GetMaxAge() *google_protobuf3.Duration {
	if x, ok := m.GetRule().(*GcRule_MaxAge); ok {
		return x.MaxAge
	}
	return nil
}

func (m *GcRule) GetIntersection() *GcRule_Intersection {
	if x, ok := m.GetRule().(*GcRule_Intersection_); ok {
		return x.Intersection
	}
	return nil
}

func (m *GcRule) GetUnion() *GcRule_Union {
	if x, ok := m.GetRule().(*GcRule_Union_); ok {
		return x.Union
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*GcRule) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _GcRule_OneofMarshaler, _GcRule_OneofUnmarshaler, _GcRule_OneofSizer, []interface{}{
		(*GcRule_MaxNumVersions)(nil),
		(*GcRule_MaxAge)(nil),
		(*GcRule_Intersection_)(nil),
		(*GcRule_Union_)(nil),
	}
}

func _GcRule_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*GcRule)
	// rule
	switch x := m.Rule.(type) {
	case *GcRule_MaxNumVersions:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.MaxNumVersions))
	case *GcRule_MaxAge:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.MaxAge); err != nil {
			return err
		}
	case *GcRule_Intersection_:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Intersection); err != nil {
			return err
		}
	case *GcRule_Union_:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Union); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("GcRule.Rule has unexpected type %T", x)
	}
	return nil
}

func _GcRule_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*GcRule)
	switch tag {
	case 1: // rule.max_num_versions
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Rule = &GcRule_MaxNumVersions{int32(x)}
		return true, err
	case 2: // rule.max_age
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf3.Duration)
		err := b.DecodeMessage(msg)
		m.Rule = &GcRule_MaxAge{msg}
		return true, err
	case 3: // rule.intersection
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GcRule_Intersection)
		err := b.DecodeMessage(msg)
		m.Rule = &GcRule_Intersection_{msg}
		return true, err
	case 4: // rule.union
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GcRule_Union)
		err := b.DecodeMessage(msg)
		m.Rule = &GcRule_Union_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _GcRule_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*GcRule)
	// rule
	switch x := m.Rule.(type) {
	case *GcRule_MaxNumVersions:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.MaxNumVersions))
	case *GcRule_MaxAge:
		s := proto.Size(x.MaxAge)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *GcRule_Intersection_:
		s := proto.Size(x.Intersection)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *GcRule_Union_:
		s := proto.Size(x.Union)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A GcRule which deletes cells matching all of the given rules.
type GcRule_Intersection struct {
	// Only delete cells which would be deleted by every element of `rules`.
	Rules []*GcRule `protobuf:"bytes,1,rep,name=rules" json:"rules,omitempty"`
}

func (m *GcRule_Intersection) Reset()                    { *m = GcRule_Intersection{} }
func (m *GcRule_Intersection) String() string            { return proto.CompactTextString(m) }
func (*GcRule_Intersection) ProtoMessage()               {}
func (*GcRule_Intersection) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

func (m *GcRule_Intersection) GetRules() []*GcRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// A GcRule which deletes cells matching any of the given rules.
type GcRule_Union struct {
	// Delete cells which would be deleted by any element of `rules`.
	Rules []*GcRule `protobuf:"bytes,1,rep,name=rules" json:"rules,omitempty"`
}

func (m *GcRule_Union) Reset()                    { *m = GcRule_Union{} }
func (m *GcRule_Union) String() string            { return proto.CompactTextString(m) }
func (*GcRule_Union) ProtoMessage()               {}
func (*GcRule_Union) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 1} }

func (m *GcRule_Union) GetRules() []*GcRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

func init() {
	proto.RegisterType((*Table)(nil), "google.bigtable.admin.table.v1.Table")
	proto.RegisterType((*ColumnFamily)(nil), "google.bigtable.admin.table.v1.ColumnFamily")
	proto.RegisterType((*GcRule)(nil), "google.bigtable.admin.table.v1.GcRule")
	proto.RegisterType((*GcRule_Intersection)(nil), "google.bigtable.admin.table.v1.GcRule.Intersection")
	proto.RegisterType((*GcRule_Union)(nil), "google.bigtable.admin.table.v1.GcRule.Union")
	proto.RegisterEnum("google.bigtable.admin.table.v1.Table_TimestampGranularity", Table_TimestampGranularity_name, Table_TimestampGranularity_value)
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/bigtable/admin/table/v1/bigtable_table_data.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 584 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xe1, 0x6a, 0x13, 0x4f,
	0x14, 0xc5, 0xb3, 0xdd, 0x24, 0xa5, 0xb7, 0xf9, 0xf7, 0x1f, 0x47, 0x91, 0x18, 0x50, 0x42, 0x04,
	0x09, 0x22, 0xb3, 0xb4, 0xf5, 0x83, 0x2d, 0xa2, 0x18, 0x5b, 0xdb, 0x4a, 0xd5, 0xb2, 0x56, 0x45,
	0x10, 0x96, 0xc9, 0x76, 0x3a, 0x2e, 0xee, 0xcc, 0x2c, 0xb3, 0x33, 0x21, 0x79, 0x03, 0x1f, 0xc5,
	0xa7, 0xf0, 0xd9, 0x64, 0x67, 0x76, 0xe3, 0x16, 0x8a, 0x89, 0xf8, 0x25, 0xb9, 0xdc, 0x99, 0xf3,
	0xbb, 0x87, 0x73, 0x67, 0xe1, 0x13, 0x93, 0x92, 0xa5, 0x14, 0x33, 0x99, 0x12, 0xc1, 0xb0, 0x54,
	0x2c, 0x60, 0x54, 0x64, 0x4a, 0x6a, 0x19, 0xb8, 0x23, 0x92, 0x25, 0x79, 0x30, 0x49, 0x98, 0x26,
	0x93, 0x94, 0x06, 0xe4, 0x82, 0x27, 0x22, 0x70, 0xf5, 0x74, 0x7b, 0xd1, 0x8f, 0xdc, 0xef, 0x05,
	0xd1, 0x04, 0x5b, 0x31, 0xba, 0x57, 0x82, 0xab, 0x1b, 0xd8, 0x2a, 0xb1, 0xab, 0xa7, 0xdb, 0xfd,
	0xf1, 0x6a, 0x83, 0x53, 0x29, 0x98, 0x32, 0x42, 0x24, 0x82, 0x05, 0x32, 0xa3, 0x8a, 0xe8, 0x44,
	0x8a, 0xdc, 0xcd, 0xe8, 0xef, 0xb1, 0x44, 0x7f, 0x35, 0x13, 0x1c, 0x4b, 0x1e, 0x38, 0x4e, 0x60,
	0x0f, 0x26, 0xe6, 0x32, 0xc8, 0xf4, 0x3c, 0xa3, 0x79, 0x70, 0x61, 0x9c, 0x64, 0x51, 0x38, 0xe9,
	0xf0, 0xa7, 0x0f, 0xad, 0xf3, 0xc2, 0x0b, 0x42, 0xd0, 0x14, 0x84, 0xd3, 0x9e, 0x37, 0xf0, 0x46,
	0x1b, 0xa1, 0xad, 0xd1, 0x6b, 0xb8, 0x11, 0x1b, 0xa5, 0xa8, 0xd0, 0xd1, 0x62, 0x68, 0x6f, 0x6d,
	0xe0, 0x8d, 0x36, 0x77, 0xee, 0xe2, 0xd2, 0x78, 0xcd, 0x19, 0x7e, 0x57, 0x5d, 0x0a, 0xbb, 0xa5,
	0x6e, 0xd1, 0x41, 0x13, 0xf8, 0x3f, 0x96, 0xa9, 0xe1, 0x22, 0xba, 0x24, 0x3c, 0x49, 0x13, 0x9a,
	0xf7, 0xfc, 0x81, 0x3f, 0xda, 0xdc, 0xd9, 0xc3, 0x7f, 0x8e, 0x08, 0x5b, 0x7f, 0xf8, 0xa5, 0x15,
	0xbf, 0x2a, 0xb5, 0x87, 0x42, 0xab, 0x79, 0xb8, 0x15, 0x5f, 0x69, 0xa2, 0x2f, 0xb0, 0xc9, 0x14,
	0x11, 0x26, 0x25, 0x2a, 0xd1, 0xf3, 0x5e, 0x73, 0xe0, 0x8d, 0xb6, 0x76, 0xf6, 0x57, 0xe3, 0x9f,
	0x27, 0x9c, 0xe6, 0x9a, 0xf0, 0xec, 0xe8, 0x37, 0x21, 0xac, 0xe3, 0xfa, 0x12, 0x6e, 0x5e, 0x63,
	0x02, 0x75, 0xc1, 0xff, 0x46, 0xe7, 0x65, 0x6e, 0x45, 0x89, 0xc6, 0xd0, 0x9a, 0x92, 0xd4, 0xd0,
	0x32, 0xaa, 0x47, 0xcb, 0x0c, 0xd4, 0xa8, 0xf3, 0xd0, 0x49, 0xf7, 0xd7, 0x9e, 0x78, 0xc3, 0x21,
	0xdc, 0xba, 0xce, 0x15, 0x02, 0x68, 0xbf, 0x39, 0x39, 0x3d, 0x3d, 0x79, 0xdf, 0x6d, 0x0c, 0xbf,
	0x7b, 0xd0, 0xa9, 0xeb, 0xaf, 0xdd, 0xe3, 0x7d, 0xf8, 0x8f, 0xc5, 0x11, 0x9d, 0x65, 0x8a, 0xe6,
	0x79, 0xb5, 0xc3, 0x8d, 0xb0, 0xc3, 0xe2, 0xc3, 0x45, 0x0f, 0x3d, 0x87, 0x75, 0x16, 0x47, 0xca,
	0xa4, 0xb4, 0xe7, 0x5b, 0xdf, 0x0f, 0x96, 0xf9, 0x3e, 0x8a, 0x43, 0x93, 0xd2, 0xb0, 0xcd, 0xec,
	0xff, 0xf0, 0x87, 0x0f, 0x6d, 0xd7, 0x42, 0x0f, 0xa1, 0xcb, 0xc9, 0x2c, 0x12, 0x86, 0x47, 0x53,
	0xaa, 0x0a, 0x7c, 0x6e, 0x0d, 0xb5, 0x8e, 0x1b, 0xe1, 0x16, 0x27, 0xb3, 0xb7, 0x86, 0x7f, 0x2c,
	0xfb, 0xe8, 0x31, 0xac, 0x17, 0x77, 0x09, 0xab, 0xf2, 0xba, 0x53, 0xcd, 0xad, 0x1e, 0x31, 0x3e,
	0x28, 0x1f, 0xed, 0x71, 0x23, 0x6c, 0x73, 0x32, 0x7b, 0xc1, 0x28, 0xfa, 0x0c, 0x9d, 0x44, 0x68,
	0xaa, 0x72, 0x1a, 0xdb, 0x57, 0xe9, 0x2c, 0xef, 0xae, 0x66, 0x19, 0x9f, 0xd4, 0xa4, 0xc7, 0x8d,
	0xf0, 0x0a, 0x0a, 0x1d, 0x40, 0xcb, 0x88, 0x82, 0xd9, 0x5c, 0x6d, 0x7d, 0x25, 0xf3, 0x83, 0x70,
	0x30, 0x27, 0xee, 0x9f, 0x42, 0xa7, 0x3e, 0x05, 0x3d, 0x85, 0x56, 0x91, 0x6d, 0x91, 0x83, 0xff,
	0x17, 0xe1, 0x3a, 0x51, 0xff, 0x10, 0x5a, 0x96, 0xff, 0x6f, 0x98, 0x71, 0x1b, 0x9a, 0x45, 0x31,
	0x7e, 0x06, 0xc3, 0x58, 0xf2, 0x25, 0xda, 0xf1, 0xed, 0x71, 0x79, 0x60, 0xbf, 0x90, 0x03, 0xa2,
	0xc9, 0x59, 0xb1, 0x91, 0x33, 0x6f, 0xd2, 0xb6, 0xab, 0xd9, 0xfd, 0x15, 0x00, 0x00, 0xff, 0xff,
	0x05, 0x58, 0xfe, 0xa5, 0x37, 0x05, 0x00, 0x00,
}
