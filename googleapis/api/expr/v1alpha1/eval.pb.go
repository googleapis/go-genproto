// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/expr/v1alpha1/eval.proto

package expr

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	status "google.golang.org/genproto/googleapis/rpc/status"
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

// The state of an evaluation.
//
// Can represent an inital, partial, or completed state of evaluation.
type EvalState struct {
	// The unique values referenced in this message.
	Values []*ExprValue `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	// An ordered list of results.
	//
	// Tracks the flow of evaluation through the expression.
	// May be sparse.
	Results              []*EvalState_Result `protobuf:"bytes,3,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *EvalState) Reset()         { *m = EvalState{} }
func (m *EvalState) String() string { return proto.CompactTextString(m) }
func (*EvalState) ProtoMessage()    {}
func (*EvalState) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e95f32326d4b8b7, []int{0}
}

func (m *EvalState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvalState.Unmarshal(m, b)
}
func (m *EvalState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvalState.Marshal(b, m, deterministic)
}
func (m *EvalState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvalState.Merge(m, src)
}
func (m *EvalState) XXX_Size() int {
	return xxx_messageInfo_EvalState.Size(m)
}
func (m *EvalState) XXX_DiscardUnknown() {
	xxx_messageInfo_EvalState.DiscardUnknown(m)
}

var xxx_messageInfo_EvalState proto.InternalMessageInfo

func (m *EvalState) GetValues() []*ExprValue {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *EvalState) GetResults() []*EvalState_Result {
	if m != nil {
		return m.Results
	}
	return nil
}

// A single evalution result.
type EvalState_Result struct {
	// The id of the expression this result if for.
	Expr int64 `protobuf:"varint,1,opt,name=expr,proto3" json:"expr,omitempty"`
	// The index in `values` of the resulting value.
	Value                int64    `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EvalState_Result) Reset()         { *m = EvalState_Result{} }
func (m *EvalState_Result) String() string { return proto.CompactTextString(m) }
func (*EvalState_Result) ProtoMessage()    {}
func (*EvalState_Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e95f32326d4b8b7, []int{0, 0}
}

func (m *EvalState_Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvalState_Result.Unmarshal(m, b)
}
func (m *EvalState_Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvalState_Result.Marshal(b, m, deterministic)
}
func (m *EvalState_Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvalState_Result.Merge(m, src)
}
func (m *EvalState_Result) XXX_Size() int {
	return xxx_messageInfo_EvalState_Result.Size(m)
}
func (m *EvalState_Result) XXX_DiscardUnknown() {
	xxx_messageInfo_EvalState_Result.DiscardUnknown(m)
}

var xxx_messageInfo_EvalState_Result proto.InternalMessageInfo

func (m *EvalState_Result) GetExpr() int64 {
	if m != nil {
		return m.Expr
	}
	return 0
}

func (m *EvalState_Result) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// The value of an evaluated expression.
type ExprValue struct {
	// An expression can resolve to a value, error or unknown.
	//
	// Types that are valid to be assigned to Kind:
	//	*ExprValue_Value
	//	*ExprValue_Error
	//	*ExprValue_Unknown
	Kind                 isExprValue_Kind `protobuf_oneof:"kind"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ExprValue) Reset()         { *m = ExprValue{} }
func (m *ExprValue) String() string { return proto.CompactTextString(m) }
func (*ExprValue) ProtoMessage()    {}
func (*ExprValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e95f32326d4b8b7, []int{1}
}

func (m *ExprValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExprValue.Unmarshal(m, b)
}
func (m *ExprValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExprValue.Marshal(b, m, deterministic)
}
func (m *ExprValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExprValue.Merge(m, src)
}
func (m *ExprValue) XXX_Size() int {
	return xxx_messageInfo_ExprValue.Size(m)
}
func (m *ExprValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ExprValue.DiscardUnknown(m)
}

var xxx_messageInfo_ExprValue proto.InternalMessageInfo

type isExprValue_Kind interface {
	isExprValue_Kind()
}

type ExprValue_Value struct {
	Value *Value `protobuf:"bytes,1,opt,name=value,proto3,oneof"`
}

type ExprValue_Error struct {
	Error *ErrorSet `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

type ExprValue_Unknown struct {
	Unknown *UnknownSet `protobuf:"bytes,3,opt,name=unknown,proto3,oneof"`
}

func (*ExprValue_Value) isExprValue_Kind() {}

func (*ExprValue_Error) isExprValue_Kind() {}

func (*ExprValue_Unknown) isExprValue_Kind() {}

func (m *ExprValue) GetKind() isExprValue_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *ExprValue) GetValue() *Value {
	if x, ok := m.GetKind().(*ExprValue_Value); ok {
		return x.Value
	}
	return nil
}

func (m *ExprValue) GetError() *ErrorSet {
	if x, ok := m.GetKind().(*ExprValue_Error); ok {
		return x.Error
	}
	return nil
}

func (m *ExprValue) GetUnknown() *UnknownSet {
	if x, ok := m.GetKind().(*ExprValue_Unknown); ok {
		return x.Unknown
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ExprValue) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ExprValue_Value)(nil),
		(*ExprValue_Error)(nil),
		(*ExprValue_Unknown)(nil),
	}
}

// A set of errors.
//
// The errors included depend on the context. See `ExprValue.error`.
type ErrorSet struct {
	// The errors in the set.
	Errors               []*status.Status `protobuf:"bytes,1,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ErrorSet) Reset()         { *m = ErrorSet{} }
func (m *ErrorSet) String() string { return proto.CompactTextString(m) }
func (*ErrorSet) ProtoMessage()    {}
func (*ErrorSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e95f32326d4b8b7, []int{2}
}

func (m *ErrorSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorSet.Unmarshal(m, b)
}
func (m *ErrorSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorSet.Marshal(b, m, deterministic)
}
func (m *ErrorSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorSet.Merge(m, src)
}
func (m *ErrorSet) XXX_Size() int {
	return xxx_messageInfo_ErrorSet.Size(m)
}
func (m *ErrorSet) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorSet.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorSet proto.InternalMessageInfo

func (m *ErrorSet) GetErrors() []*status.Status {
	if m != nil {
		return m.Errors
	}
	return nil
}

// A set of expressions for which the value is unknown.
//
// The unknowns included depend on the context. See `ExprValue.unknown`.
type UnknownSet struct {
	// The ids of the expressions with unknown values.
	Exprs                []int64  `protobuf:"varint,1,rep,packed,name=exprs,proto3" json:"exprs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnknownSet) Reset()         { *m = UnknownSet{} }
func (m *UnknownSet) String() string { return proto.CompactTextString(m) }
func (*UnknownSet) ProtoMessage()    {}
func (*UnknownSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e95f32326d4b8b7, []int{3}
}

func (m *UnknownSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnknownSet.Unmarshal(m, b)
}
func (m *UnknownSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnknownSet.Marshal(b, m, deterministic)
}
func (m *UnknownSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnknownSet.Merge(m, src)
}
func (m *UnknownSet) XXX_Size() int {
	return xxx_messageInfo_UnknownSet.Size(m)
}
func (m *UnknownSet) XXX_DiscardUnknown() {
	xxx_messageInfo_UnknownSet.DiscardUnknown(m)
}

var xxx_messageInfo_UnknownSet proto.InternalMessageInfo

func (m *UnknownSet) GetExprs() []int64 {
	if m != nil {
		return m.Exprs
	}
	return nil
}

func init() {
	proto.RegisterType((*EvalState)(nil), "google.api.expr.v1alpha1.EvalState")
	proto.RegisterType((*EvalState_Result)(nil), "google.api.expr.v1alpha1.EvalState.Result")
	proto.RegisterType((*ExprValue)(nil), "google.api.expr.v1alpha1.ExprValue")
	proto.RegisterType((*ErrorSet)(nil), "google.api.expr.v1alpha1.ErrorSet")
	proto.RegisterType((*UnknownSet)(nil), "google.api.expr.v1alpha1.UnknownSet")
}

func init() {
	proto.RegisterFile("google/api/expr/v1alpha1/eval.proto", fileDescriptor_1e95f32326d4b8b7)
}

var fileDescriptor_1e95f32326d4b8b7 = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x4b, 0xeb, 0x40,
	0x10, 0xc7, 0x5f, 0x5e, 0xda, 0xf4, 0xbd, 0xe9, 0x6d, 0x11, 0x0c, 0x45, 0xb0, 0xa4, 0x3d, 0x94,
	0x1e, 0x36, 0x34, 0x82, 0x82, 0xf5, 0x20, 0xc5, 0x82, 0xc7, 0x92, 0xa2, 0x07, 0x6f, 0x6b, 0x5d,
	0x62, 0xe8, 0x9a, 0x5d, 0x36, 0x3f, 0xec, 0xdf, 0xe7, 0xd1, 0xbf, 0xc8, 0xa3, 0xec, 0x6c, 0x16,
	0x0f, 0x92, 0xde, 0x3a, 0xbb, 0x9f, 0xcf, 0x77, 0xa6, 0xd9, 0x81, 0x49, 0x26, 0x65, 0x26, 0x78,
	0xcc, 0x54, 0x1e, 0xf3, 0x83, 0xd2, 0x71, 0xb3, 0x60, 0x42, 0xbd, 0xb2, 0x45, 0xcc, 0x1b, 0x26,
	0xa8, 0xd2, 0xb2, 0x92, 0x24, 0xb4, 0x10, 0x65, 0x2a, 0xa7, 0x06, 0xa2, 0x0e, 0x1a, 0x4d, 0x3b,
	0xf5, 0x86, 0x89, 0x9a, 0x5b, 0x7f, 0x74, 0xda, 0x52, 0x5a, 0xed, 0xe2, 0xb2, 0x62, 0x55, 0x5d,
	0xda, 0x8b, 0xe8, 0xc3, 0x83, 0xff, 0xeb, 0x86, 0x89, 0x6d, 0xc5, 0x2a, 0x4e, 0x96, 0x10, 0xa0,
	0x55, 0x86, 0xde, 0xd8, 0x9f, 0x0d, 0x93, 0x09, 0xed, 0xea, 0x4b, 0xd7, 0x07, 0xa5, 0x1f, 0x0d,
	0x9b, 0xb6, 0x0a, 0xb9, 0x83, 0x81, 0xe6, 0x65, 0x2d, 0xaa, 0x32, 0xf4, 0xd1, 0x9e, 0x1f, 0xb1,
	0x5d, 0x4b, 0x9a, 0xa2, 0x92, 0x3a, 0x75, 0x94, 0x40, 0x60, 0x8f, 0x08, 0x81, 0x9e, 0x91, 0x42,
	0x6f, 0xec, 0xcd, 0xfc, 0x14, 0x7f, 0x93, 0x13, 0xe8, 0x63, 0xb7, 0xf0, 0x2f, 0x1e, 0xda, 0x22,
	0xfa, 0x34, 0x7f, 0xc2, 0xcd, 0x43, 0xae, 0x1c, 0x63, 0xc4, 0x61, 0x72, 0xde, 0x3d, 0x05, 0xf2,
	0xf7, 0x7f, 0xda, 0x18, 0x72, 0x0d, 0x7d, 0xae, 0xb5, 0xd4, 0x18, 0x3e, 0x4c, 0xa2, 0x23, 0xe3,
	0x1b, 0x6c, 0xcb, 0x2b, 0xe3, 0xa2, 0x42, 0x6e, 0x61, 0x50, 0x17, 0xfb, 0x42, 0xbe, 0x17, 0xa1,
	0x8f, 0xf6, 0xb4, 0xdb, 0x7e, 0xb0, 0xa0, 0xf5, 0x9d, 0xb6, 0x0a, 0xa0, 0xb7, 0xcf, 0x8b, 0x97,
	0xe8, 0x12, 0xfe, 0xb9, 0x78, 0x32, 0x87, 0x00, 0xe3, 0xdd, 0x7b, 0x10, 0x17, 0xaa, 0xd5, 0x8e,
	0x6e, 0xf1, 0x1d, 0xd3, 0x96, 0x88, 0x22, 0x80, 0x9f, 0x60, 0xf3, 0xa1, 0x4c, 0x53, 0x2b, 0xfa,
	0xa9, 0x2d, 0x56, 0x02, 0xce, 0x76, 0xf2, 0xad, 0x73, 0xb2, 0x15, 0xae, 0xc2, 0xc6, 0x2c, 0xc6,
	0xc6, 0x7b, 0xba, 0x69, 0xb1, 0x4c, 0x0a, 0x56, 0x64, 0x54, 0xea, 0x2c, 0xce, 0x78, 0x81, 0x6b,
	0x13, 0xdb, 0x2b, 0xa6, 0xf2, 0xf2, 0xf7, 0xe2, 0x2d, 0x4d, 0xf5, 0xe5, 0x79, 0xcf, 0x01, 0xb2,
	0x17, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x62, 0xde, 0x1d, 0xe2, 0x02, 0x00, 0x00,
}
