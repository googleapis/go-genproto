// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/resultstore/v2/coverage.proto

package resultstore // import "google.golang.org/genproto/googleapis/devtools/resultstore/v2"

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

// Describes line coverage for a file
type LineCoverage struct {
	// Which source lines in the file represent the start of a statement that was
	// instrumented to detect whether it was executed by the test.
	//
	// This is a bitfield where i-th bit corresponds to the i-th line. Divide line
	// number by 8 to get index into byte array. Mod line number by 8 to get bit
	// number (0 = LSB, 7 = MSB).
	//
	// A 1 denotes the line was instrumented.
	// A 0 denotes the line was not instrumented.
	InstrumentedLines []byte `protobuf:"bytes,1,opt,name=instrumented_lines,json=instrumentedLines,proto3" json:"instrumented_lines,omitempty"`
	// Which of the instrumented source lines were executed by the test. Should
	// include lines that were not instrumented.
	//
	// This is a bitfield where i-th bit corresponds to the i-th line. Divide line
	// number by 8 to get index into byte array. Mod line number by 8 to get bit
	// number (0 = LSB, 7 = MSB).
	//
	// A 1 denotes the line was executed.
	// A 0 denotes the line was not executed.
	ExecutedLines        []byte   `protobuf:"bytes,2,opt,name=executed_lines,json=executedLines,proto3" json:"executed_lines,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LineCoverage) Reset()         { *m = LineCoverage{} }
func (m *LineCoverage) String() string { return proto.CompactTextString(m) }
func (*LineCoverage) ProtoMessage()    {}
func (*LineCoverage) Descriptor() ([]byte, []int) {
	return fileDescriptor_823b424facf00e45, []int{0}
}
func (m *LineCoverage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LineCoverage.Unmarshal(m, b)
}
func (m *LineCoverage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LineCoverage.Marshal(b, m, deterministic)
}
func (m *LineCoverage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LineCoverage.Merge(m, src)
}
func (m *LineCoverage) XXX_Size() int {
	return xxx_messageInfo_LineCoverage.Size(m)
}
func (m *LineCoverage) XXX_DiscardUnknown() {
	xxx_messageInfo_LineCoverage.DiscardUnknown(m)
}

var xxx_messageInfo_LineCoverage proto.InternalMessageInfo

func (m *LineCoverage) GetInstrumentedLines() []byte {
	if m != nil {
		return m.InstrumentedLines
	}
	return nil
}

func (m *LineCoverage) GetExecutedLines() []byte {
	if m != nil {
		return m.ExecutedLines
	}
	return nil
}

// Describes branch coverage for a file
type BranchCoverage struct {
	// The field branch_present denotes the lines containing at least one branch.
	//
	// This is a bitfield where i-th bit corresponds to the i-th line. Divide line
	// number by 8 to get index into byte array. Mod line number by 8 to get bit
	// number (0 = LSB, 7 = MSB).
	//
	// A 1 denotes the line contains at least one branch.
	// A 0 denotes the line contains no branches.
	BranchPresent []byte `protobuf:"bytes,1,opt,name=branch_present,json=branchPresent,proto3" json:"branch_present,omitempty"`
	// Contains the number of branches present, only for the lines which have the
	// corresponding bit set in branch_present, in a relative order ignoring
	// lines which do not have any branches.
	BranchesInLine []int32 `protobuf:"varint,2,rep,packed,name=branches_in_line,json=branchesInLine,proto3" json:"branches_in_line,omitempty"`
	// As each branch can have any one of the following three states: not
	// executed, executed but not taken, executed and taken.
	//
	// This is a bitfield where i-th bit corresponds to the i-th line. Divide line
	// number by 8 to get index into byte array. Mod line number by 8 to get bit
	// number (0 = LSB, 7 = MSB).
	//
	// i-th bit of the following two byte arrays are used to denote the above
	// mentioned states.
	//
	// not executed: i-th bit of executed == 0 && i-th bit of taken == 0
	// executed but not taken: i-th bit of executed == 1 && i-th bit of taken == 0
	// executed and taken: i-th bit of executed == 1 && i-th bit of taken == 1
	Executed []byte `protobuf:"bytes,3,opt,name=executed,proto3" json:"executed,omitempty"`
	// Described above.
	Taken                []byte   `protobuf:"bytes,4,opt,name=taken,proto3" json:"taken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BranchCoverage) Reset()         { *m = BranchCoverage{} }
func (m *BranchCoverage) String() string { return proto.CompactTextString(m) }
func (*BranchCoverage) ProtoMessage()    {}
func (*BranchCoverage) Descriptor() ([]byte, []int) {
	return fileDescriptor_823b424facf00e45, []int{1}
}
func (m *BranchCoverage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BranchCoverage.Unmarshal(m, b)
}
func (m *BranchCoverage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BranchCoverage.Marshal(b, m, deterministic)
}
func (m *BranchCoverage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BranchCoverage.Merge(m, src)
}
func (m *BranchCoverage) XXX_Size() int {
	return xxx_messageInfo_BranchCoverage.Size(m)
}
func (m *BranchCoverage) XXX_DiscardUnknown() {
	xxx_messageInfo_BranchCoverage.DiscardUnknown(m)
}

var xxx_messageInfo_BranchCoverage proto.InternalMessageInfo

func (m *BranchCoverage) GetBranchPresent() []byte {
	if m != nil {
		return m.BranchPresent
	}
	return nil
}

func (m *BranchCoverage) GetBranchesInLine() []int32 {
	if m != nil {
		return m.BranchesInLine
	}
	return nil
}

func (m *BranchCoverage) GetExecuted() []byte {
	if m != nil {
		return m.Executed
	}
	return nil
}

func (m *BranchCoverage) GetTaken() []byte {
	if m != nil {
		return m.Taken
	}
	return nil
}

// Describes code coverage for a particular file under test.
type FileCoverage struct {
	// Path of source file within the SourceContext of this Invocation.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Details of lines in a file required to calculate line coverage.
	LineCoverage *LineCoverage `protobuf:"bytes,2,opt,name=line_coverage,json=lineCoverage,proto3" json:"line_coverage,omitempty"`
	// Details of branches in a file required to calculate branch coverage.
	BranchCoverage       *BranchCoverage `protobuf:"bytes,3,opt,name=branch_coverage,json=branchCoverage,proto3" json:"branch_coverage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *FileCoverage) Reset()         { *m = FileCoverage{} }
func (m *FileCoverage) String() string { return proto.CompactTextString(m) }
func (*FileCoverage) ProtoMessage()    {}
func (*FileCoverage) Descriptor() ([]byte, []int) {
	return fileDescriptor_823b424facf00e45, []int{2}
}
func (m *FileCoverage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileCoverage.Unmarshal(m, b)
}
func (m *FileCoverage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileCoverage.Marshal(b, m, deterministic)
}
func (m *FileCoverage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileCoverage.Merge(m, src)
}
func (m *FileCoverage) XXX_Size() int {
	return xxx_messageInfo_FileCoverage.Size(m)
}
func (m *FileCoverage) XXX_DiscardUnknown() {
	xxx_messageInfo_FileCoverage.DiscardUnknown(m)
}

var xxx_messageInfo_FileCoverage proto.InternalMessageInfo

func (m *FileCoverage) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *FileCoverage) GetLineCoverage() *LineCoverage {
	if m != nil {
		return m.LineCoverage
	}
	return nil
}

func (m *FileCoverage) GetBranchCoverage() *BranchCoverage {
	if m != nil {
		return m.BranchCoverage
	}
	return nil
}

// Describes code coverage for a build or test Action. This is used to store
// baseline coverage for build Actions and test coverage for test Actions.
type ActionCoverage struct {
	// List of coverage info for all source files that the TestResult covers.
	FileCoverages        []*FileCoverage `protobuf:"bytes,2,rep,name=file_coverages,json=fileCoverages,proto3" json:"file_coverages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ActionCoverage) Reset()         { *m = ActionCoverage{} }
func (m *ActionCoverage) String() string { return proto.CompactTextString(m) }
func (*ActionCoverage) ProtoMessage()    {}
func (*ActionCoverage) Descriptor() ([]byte, []int) {
	return fileDescriptor_823b424facf00e45, []int{3}
}
func (m *ActionCoverage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActionCoverage.Unmarshal(m, b)
}
func (m *ActionCoverage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActionCoverage.Marshal(b, m, deterministic)
}
func (m *ActionCoverage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionCoverage.Merge(m, src)
}
func (m *ActionCoverage) XXX_Size() int {
	return xxx_messageInfo_ActionCoverage.Size(m)
}
func (m *ActionCoverage) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionCoverage.DiscardUnknown(m)
}

var xxx_messageInfo_ActionCoverage proto.InternalMessageInfo

func (m *ActionCoverage) GetFileCoverages() []*FileCoverage {
	if m != nil {
		return m.FileCoverages
	}
	return nil
}

func init() {
	proto.RegisterType((*LineCoverage)(nil), "google.devtools.resultstore.v2.LineCoverage")
	proto.RegisterType((*BranchCoverage)(nil), "google.devtools.resultstore.v2.BranchCoverage")
	proto.RegisterType((*FileCoverage)(nil), "google.devtools.resultstore.v2.FileCoverage")
	proto.RegisterType((*ActionCoverage)(nil), "google.devtools.resultstore.v2.ActionCoverage")
}

func init() {
	proto.RegisterFile("google/devtools/resultstore/v2/coverage.proto", fileDescriptor_823b424facf00e45)
}

var fileDescriptor_823b424facf00e45 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x6b, 0xe3, 0x30,
	0x10, 0xc5, 0xf9, 0x58, 0x76, 0x67, 0x6d, 0xef, 0xae, 0xd8, 0x83, 0xd9, 0xc3, 0x12, 0x0c, 0x81,
	0x1c, 0x1a, 0x19, 0xd2, 0x63, 0x4f, 0x4d, 0xa1, 0x10, 0xe8, 0x21, 0x75, 0x0f, 0x85, 0x5e, 0x8c,
	0xe3, 0x4c, 0x1c, 0x51, 0x45, 0x72, 0x25, 0xc5, 0xf4, 0x7f, 0xf4, 0x7f, 0xf5, 0x37, 0x95, 0xc8,
	0x1f, 0x71, 0x0f, 0x6d, 0x7a, 0xd3, 0xbc, 0x79, 0xf3, 0xde, 0xcc, 0x43, 0x30, 0xcd, 0xa5, 0xcc,
	0x39, 0x46, 0x6b, 0x2c, 0x8d, 0x94, 0x5c, 0x47, 0x0a, 0xf5, 0x9e, 0x1b, 0x6d, 0xa4, 0xc2, 0xa8,
	0x9c, 0x45, 0x99, 0x2c, 0x51, 0xa5, 0x39, 0xd2, 0x42, 0x49, 0x23, 0xc9, 0xff, 0x8a, 0x4e, 0x1b,
	0x3a, 0xed, 0xd0, 0x69, 0x39, 0x0b, 0xd7, 0xe0, 0xde, 0x30, 0x81, 0x57, 0xf5, 0x14, 0x99, 0x02,
	0x61, 0x42, 0x1b, 0xb5, 0xdf, 0xa1, 0x30, 0xb8, 0x4e, 0x38, 0x13, 0xa8, 0x03, 0x67, 0xe4, 0x4c,
	0xdc, 0xf8, 0x4f, 0xb7, 0x73, 0x98, 0xd2, 0x64, 0x0c, 0x3e, 0x3e, 0x63, 0xb6, 0x3f, 0x52, 0x7b,
	0x96, 0xea, 0x35, 0xa8, 0xa5, 0x85, 0x2f, 0x0e, 0xf8, 0x73, 0x95, 0x8a, 0x6c, 0xdb, 0x1a, 0x8d,
	0xc1, 0x5f, 0x59, 0x24, 0x29, 0x14, 0x6a, 0x14, 0xa6, 0x36, 0xf1, 0x2a, 0x74, 0x59, 0x81, 0x64,
	0x02, 0xbf, 0x2b, 0x00, 0x75, 0xc2, 0x84, 0xf5, 0x08, 0x7a, 0xa3, 0xfe, 0x64, 0x18, 0xfb, 0x0d,
	0xbe, 0x10, 0x07, 0x13, 0xf2, 0x0f, 0xbe, 0x37, 0xa6, 0x41, 0xdf, 0x4a, 0xb5, 0x35, 0xf9, 0x0b,
	0x43, 0x93, 0x3e, 0xa2, 0x08, 0x06, 0xb6, 0x51, 0x15, 0xe1, 0xab, 0x03, 0xee, 0x35, 0xe3, 0xc7,
	0xe3, 0x09, 0x0c, 0x8a, 0xd4, 0x6c, 0xed, 0x26, 0x3f, 0x62, 0xfb, 0x26, 0xb7, 0xe0, 0x1d, 0x4c,
	0x93, 0x26, 0x57, 0x7b, 0xe0, 0xcf, 0xd9, 0x19, 0xfd, 0x3c, 0x58, 0xda, 0x4d, 0x35, 0x76, 0x79,
	0x37, 0xe3, 0x7b, 0xf8, 0x55, 0x9f, 0xde, 0x8a, 0xf6, 0xad, 0x28, 0x3d, 0x25, 0xfa, 0x3e, 0xc3,
	0x26, 0x82, 0xa6, 0x0e, 0x11, 0xfc, 0xcb, 0xcc, 0x30, 0x29, 0x5a, 0xab, 0x3b, 0xf0, 0x37, 0x8c,
	0x1f, 0xb7, 0xd7, 0x36, 0xbc, 0x2f, 0xac, 0xdf, 0xcd, 0x25, 0xf6, 0x36, 0x9d, 0x4a, 0xcf, 0x9f,
	0x20, 0xcc, 0xe4, 0xee, 0x84, 0xc2, 0xd2, 0x79, 0x58, 0xd4, 0x8c, 0x5c, 0xf2, 0x54, 0xe4, 0x54,
	0xaa, 0x3c, 0xca, 0x51, 0xd8, 0x7f, 0x19, 0x55, 0xad, 0xb4, 0x60, 0xfa, 0xa3, 0x9f, 0x7c, 0xd1,
	0x29, 0x57, 0xdf, 0xec, 0xd4, 0xf9, 0x5b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x92, 0x57, 0x34,
	0xfe, 0x02, 0x00, 0x00,
}
