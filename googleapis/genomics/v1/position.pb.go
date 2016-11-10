// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/genomics/v1/position.proto
// DO NOT EDIT!

package google_genomics_v1 // import "google.golang.org/genproto/googleapis/genomics/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// An abstraction for referring to a genomic position, in relation to some
// already known reference. For now, represents a genomic position as a
// reference name, a base number on that reference (0-based), and a
// determination of forward or reverse strand.
type Position struct {
	// The name of the reference in whatever reference set is being used.
	ReferenceName string `protobuf:"bytes,1,opt,name=reference_name,json=referenceName" json:"reference_name,omitempty"`
	// The 0-based offset from the start of the forward strand for that reference.
	Position int64 `protobuf:"varint,2,opt,name=position" json:"position,omitempty"`
	// Whether this position is on the reverse strand, as opposed to the forward
	// strand.
	ReverseStrand bool `protobuf:"varint,3,opt,name=reverse_strand,json=reverseStrand" json:"reverse_strand,omitempty"`
}

func (m *Position) Reset()                    { *m = Position{} }
func (m *Position) String() string            { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()               {}
func (*Position) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func init() {
	proto.RegisterType((*Position)(nil), "google.genomics.v1.Position")
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/genomics/v1/position.proto", fileDescriptor4)
}

var fileDescriptor4 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x8f, 0xb1, 0x6a, 0xc3, 0x30,
	0x10, 0x86, 0x51, 0x03, 0xc5, 0x15, 0xa4, 0x83, 0x87, 0x62, 0x32, 0x99, 0x42, 0xc1, 0x43, 0x91,
	0x08, 0x7d, 0x81, 0x92, 0xad, 0x4b, 0x31, 0xe9, 0x03, 0x84, 0xab, 0x7a, 0x11, 0x82, 0xf8, 0x4e,
	0x48, 0x42, 0xcf, 0xde, 0xb1, 0x48, 0x8e, 0xbd, 0x74, 0xc9, 0xa8, 0x4f, 0x3f, 0x9f, 0x3e, 0xc9,
	0x77, 0xcb, 0x6c, 0x2f, 0xa8, 0x2c, 0x5f, 0x80, 0xac, 0xe2, 0x60, 0xb5, 0x45, 0xf2, 0x81, 0x13,
	0xeb, 0xf9, 0x0a, 0xbc, 0x8b, 0x85, 0xf1, 0xe4, 0x4c, 0xd4, 0x79, 0xaf, 0x3d, 0x47, 0x97, 0x1c,
	0x93, 0xaa, 0xab, 0xb6, 0x5d, 0x0c, 0xd7, 0x89, 0xca, 0xfb, 0xdd, 0xc7, 0x6d, 0x56, 0xf0, 0x4e,
	0x47, 0x0c, 0xd9, 0x19, 0x34, 0x4c, 0x67, 0x67, 0x35, 0x10, 0x71, 0x82, 0x62, 0x8f, 0xb3, 0xfe,
	0x39, 0xc9, 0x66, 0xbc, 0x3e, 0xd8, 0xbe, 0xc8, 0xc7, 0x80, 0x67, 0x0c, 0x48, 0x06, 0x4f, 0x04,
	0x13, 0x76, 0xa2, 0x17, 0xc3, 0xc3, 0x71, 0xbb, 0xd2, 0x4f, 0x98, 0xb0, 0xdd, 0xc9, 0x66, 0x69,
	0xec, 0xee, 0x7a, 0x31, 0x6c, 0x8e, 0xeb, 0x79, 0x56, 0x64, 0x0c, 0x11, 0x4f, 0x31, 0x05, 0xa0,
	0x9f, 0x6e, 0xd3, 0x8b, 0xa1, 0x29, 0x8a, 0x4a, 0xbf, 0x2a, 0x3c, 0xbc, 0xca, 0x27, 0xc3, 0x93,
	0xfa, 0xff, 0xb5, 0xc3, 0x76, 0xa9, 0x19, 0x4b, 0xde, 0x28, 0x7e, 0x85, 0xf8, 0xbe, 0xaf, 0xa9,
	0x6f, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x54, 0xea, 0x40, 0x4d, 0x01, 0x00, 0x00,
}
