// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/firestore/admin/v1beta2/firestore_admin.proto

package admin // import "google.golang.org/genproto/googleapis/firestore/admin/v1beta2"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import longrunning "google.golang.org/genproto/googleapis/longrunning"
import field_mask "google.golang.org/genproto/protobuf/field_mask"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// The request for [FirestoreAdmin.UpdateField][google.firestore.admin.v1beta2.FirestoreAdmin.UpdateField].
type UpdateFieldRequest struct {
	// The field to be updated.
	Field *Field `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	// A mask, relative to the field. If specified, only configuration specified
	// by this field_mask will be updated in the field.
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateFieldRequest) Reset()         { *m = UpdateFieldRequest{} }
func (m *UpdateFieldRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateFieldRequest) ProtoMessage()    {}
func (*UpdateFieldRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d93f4454f4162ee9, []int{0}
}
func (m *UpdateFieldRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateFieldRequest.Unmarshal(m, b)
}
func (m *UpdateFieldRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateFieldRequest.Marshal(b, m, deterministic)
}
func (m *UpdateFieldRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateFieldRequest.Merge(m, src)
}
func (m *UpdateFieldRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateFieldRequest.Size(m)
}
func (m *UpdateFieldRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateFieldRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateFieldRequest proto.InternalMessageInfo

func (m *UpdateFieldRequest) GetField() *Field {
	if m != nil {
		return m.Field
	}
	return nil
}

func (m *UpdateFieldRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

// The request for [FirestoreAdmin.GetField][google.firestore.admin.v1beta2.FirestoreAdmin.GetField].
type GetFieldRequest struct {
	// A name of the form
	// `/projects/{project_id}/databases/{database_id}/collectionGroups/{collection_id}/fields/{field_id}`
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFieldRequest) Reset()         { *m = GetFieldRequest{} }
func (m *GetFieldRequest) String() string { return proto.CompactTextString(m) }
func (*GetFieldRequest) ProtoMessage()    {}
func (*GetFieldRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d93f4454f4162ee9, []int{1}
}
func (m *GetFieldRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFieldRequest.Unmarshal(m, b)
}
func (m *GetFieldRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFieldRequest.Marshal(b, m, deterministic)
}
func (m *GetFieldRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFieldRequest.Merge(m, src)
}
func (m *GetFieldRequest) XXX_Size() int {
	return xxx_messageInfo_GetFieldRequest.Size(m)
}
func (m *GetFieldRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFieldRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFieldRequest proto.InternalMessageInfo

func (m *GetFieldRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The request for [FirestoreAdmin.ListFields][google.firestore.admin.v1beta2.FirestoreAdmin.ListFields].
type ListFieldsRequest struct {
	// A parent name of the form
	// `/projects/{project_id}/databases/{database_id}/collectionGroups/{collection_id}`
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The filter to apply to list results. Currently,
	// [FirestoreAdmin.ListFields][google.firestore.admin.v1beta2.FirestoreAdmin.ListFields] only supports listing fields
	// that have been explicitly overridden. To issue this query, call
	// [FirestoreAdmin.ListFields][google.firestore.admin.v1beta2.FirestoreAdmin.ListFields] with the filter set to
	// `indexConfig.usesAncestorConfig:false`.
	Filter string `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	// The number of results to return.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A page token, returned from a previous call to
	// [FirestoreAdmin.ListFields][google.firestore.admin.v1beta2.FirestoreAdmin.ListFields], that may be used to get the next
	// page of results.
	PageToken            string   `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListFieldsRequest) Reset()         { *m = ListFieldsRequest{} }
func (m *ListFieldsRequest) String() string { return proto.CompactTextString(m) }
func (*ListFieldsRequest) ProtoMessage()    {}
func (*ListFieldsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d93f4454f4162ee9, []int{2}
}
func (m *ListFieldsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFieldsRequest.Unmarshal(m, b)
}
func (m *ListFieldsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFieldsRequest.Marshal(b, m, deterministic)
}
func (m *ListFieldsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFieldsRequest.Merge(m, src)
}
func (m *ListFieldsRequest) XXX_Size() int {
	return xxx_messageInfo_ListFieldsRequest.Size(m)
}
func (m *ListFieldsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFieldsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListFieldsRequest proto.InternalMessageInfo

func (m *ListFieldsRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListFieldsRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *ListFieldsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListFieldsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// The response for [FirestoreAdmin.ListFields][google.firestore.admin.v1beta2.FirestoreAdmin.ListFields].
type ListFieldsResponse struct {
	// The requested fields.
	Fields []*Field `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty"`
	// A page token that may be used to request another page of results. If blank,
	// this is the last page.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListFieldsResponse) Reset()         { *m = ListFieldsResponse{} }
func (m *ListFieldsResponse) String() string { return proto.CompactTextString(m) }
func (*ListFieldsResponse) ProtoMessage()    {}
func (*ListFieldsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d93f4454f4162ee9, []int{3}
}
func (m *ListFieldsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFieldsResponse.Unmarshal(m, b)
}
func (m *ListFieldsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFieldsResponse.Marshal(b, m, deterministic)
}
func (m *ListFieldsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFieldsResponse.Merge(m, src)
}
func (m *ListFieldsResponse) XXX_Size() int {
	return xxx_messageInfo_ListFieldsResponse.Size(m)
}
func (m *ListFieldsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFieldsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListFieldsResponse proto.InternalMessageInfo

func (m *ListFieldsResponse) GetFields() []*Field {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *ListFieldsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func init() {
	proto.RegisterType((*UpdateFieldRequest)(nil), "google.firestore.admin.v1beta2.UpdateFieldRequest")
	proto.RegisterType((*GetFieldRequest)(nil), "google.firestore.admin.v1beta2.GetFieldRequest")
	proto.RegisterType((*ListFieldsRequest)(nil), "google.firestore.admin.v1beta2.ListFieldsRequest")
	proto.RegisterType((*ListFieldsResponse)(nil), "google.firestore.admin.v1beta2.ListFieldsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FirestoreAdminClient is the client API for FirestoreAdmin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FirestoreAdminClient interface {
	// Gets the metadata and configuration for a Field.
	GetField(ctx context.Context, in *GetFieldRequest, opts ...grpc.CallOption) (*Field, error)
	// Updates a field configuration. Currently, field updates apply only to
	// single field index configuration. However, calls to
	// [FirestoreAdmin.UpdateField][google.firestore.admin.v1beta2.FirestoreAdmin.UpdateField] should provide a field mask to avoid
	// changing any configuration that the caller isn't aware of. The field mask
	// should be specified as: `{ paths: "index_config" }`.
	//
	// This call returns a [google.longrunning.Operation][google.longrunning.Operation] which may be used to
	// track the status of the field update. The metadata for
	// the operation will be the type [FieldOperationMetadata][google.firestore.admin.v1beta2.FieldOperationMetadata].
	//
	// To configure the default field settings for the database, use
	// the special `Field` with resource name:
	// `/projects/{project_id}/databases/{database_id}/collectionGroups/__default__/fields/*`.
	UpdateField(ctx context.Context, in *UpdateFieldRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
	// Lists the field configuration and metadata for this database.
	//
	// Currently, [FirestoreAdmin.ListFields][google.firestore.admin.v1beta2.FirestoreAdmin.ListFields] only supports listing fields
	// that have been explicitly overridden. To issue this query, call
	// [FirestoreAdmin.ListFields][google.firestore.admin.v1beta2.FirestoreAdmin.ListFields] with the filter set to
	// `indexConfig.usesAncestorConfig:false`.
	ListFields(ctx context.Context, in *ListFieldsRequest, opts ...grpc.CallOption) (*ListFieldsResponse, error)
}

type firestoreAdminClient struct {
	cc *grpc.ClientConn
}

func NewFirestoreAdminClient(cc *grpc.ClientConn) FirestoreAdminClient {
	return &firestoreAdminClient{cc}
}

func (c *firestoreAdminClient) GetField(ctx context.Context, in *GetFieldRequest, opts ...grpc.CallOption) (*Field, error) {
	out := new(Field)
	err := c.cc.Invoke(ctx, "/google.firestore.admin.v1beta2.FirestoreAdmin/GetField", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firestoreAdminClient) UpdateField(ctx context.Context, in *UpdateFieldRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/google.firestore.admin.v1beta2.FirestoreAdmin/UpdateField", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firestoreAdminClient) ListFields(ctx context.Context, in *ListFieldsRequest, opts ...grpc.CallOption) (*ListFieldsResponse, error) {
	out := new(ListFieldsResponse)
	err := c.cc.Invoke(ctx, "/google.firestore.admin.v1beta2.FirestoreAdmin/ListFields", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FirestoreAdminServer is the server API for FirestoreAdmin service.
type FirestoreAdminServer interface {
	// Gets the metadata and configuration for a Field.
	GetField(context.Context, *GetFieldRequest) (*Field, error)
	// Updates a field configuration. Currently, field updates apply only to
	// single field index configuration. However, calls to
	// [FirestoreAdmin.UpdateField][google.firestore.admin.v1beta2.FirestoreAdmin.UpdateField] should provide a field mask to avoid
	// changing any configuration that the caller isn't aware of. The field mask
	// should be specified as: `{ paths: "index_config" }`.
	//
	// This call returns a [google.longrunning.Operation][google.longrunning.Operation] which may be used to
	// track the status of the field update. The metadata for
	// the operation will be the type [FieldOperationMetadata][google.firestore.admin.v1beta2.FieldOperationMetadata].
	//
	// To configure the default field settings for the database, use
	// the special `Field` with resource name:
	// `/projects/{project_id}/databases/{database_id}/collectionGroups/__default__/fields/*`.
	UpdateField(context.Context, *UpdateFieldRequest) (*longrunning.Operation, error)
	// Lists the field configuration and metadata for this database.
	//
	// Currently, [FirestoreAdmin.ListFields][google.firestore.admin.v1beta2.FirestoreAdmin.ListFields] only supports listing fields
	// that have been explicitly overridden. To issue this query, call
	// [FirestoreAdmin.ListFields][google.firestore.admin.v1beta2.FirestoreAdmin.ListFields] with the filter set to
	// `indexConfig.usesAncestorConfig:false`.
	ListFields(context.Context, *ListFieldsRequest) (*ListFieldsResponse, error)
}

func RegisterFirestoreAdminServer(s *grpc.Server, srv FirestoreAdminServer) {
	s.RegisterService(&_FirestoreAdmin_serviceDesc, srv)
}

func _FirestoreAdmin_GetField_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFieldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirestoreAdminServer).GetField(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.firestore.admin.v1beta2.FirestoreAdmin/GetField",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirestoreAdminServer).GetField(ctx, req.(*GetFieldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FirestoreAdmin_UpdateField_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFieldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirestoreAdminServer).UpdateField(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.firestore.admin.v1beta2.FirestoreAdmin/UpdateField",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirestoreAdminServer).UpdateField(ctx, req.(*UpdateFieldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FirestoreAdmin_ListFields_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFieldsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirestoreAdminServer).ListFields(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.firestore.admin.v1beta2.FirestoreAdmin/ListFields",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirestoreAdminServer).ListFields(ctx, req.(*ListFieldsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FirestoreAdmin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.firestore.admin.v1beta2.FirestoreAdmin",
	HandlerType: (*FirestoreAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetField",
			Handler:    _FirestoreAdmin_GetField_Handler,
		},
		{
			MethodName: "UpdateField",
			Handler:    _FirestoreAdmin_UpdateField_Handler,
		},
		{
			MethodName: "ListFields",
			Handler:    _FirestoreAdmin_ListFields_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/firestore/admin/v1beta2/firestore_admin.proto",
}

func init() {
	proto.RegisterFile("google/firestore/admin/v1beta2/firestore_admin.proto", fileDescriptor_d93f4454f4162ee9)
}

var fileDescriptor_d93f4454f4162ee9 = []byte{
	// 610 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x95, 0xd3, 0xb4, 0x6a, 0xa7, 0xfa, 0xbe, 0x8a, 0x41, 0x42, 0x91, 0x4b, 0x51, 0x64, 0x28,
	0xaa, 0xba, 0xf0, 0x28, 0x86, 0x15, 0x51, 0x17, 0x24, 0x28, 0x41, 0x08, 0x44, 0x48, 0xf9, 0x91,
	0xd8, 0x44, 0x93, 0xf8, 0xc6, 0x32, 0x71, 0x66, 0x8c, 0x67, 0x8c, 0x4a, 0xab, 0x0a, 0x89, 0x07,
	0x60, 0xc3, 0x1b, 0xb0, 0x64, 0xd1, 0x3d, 0xab, 0x3e, 0x04, 0xaf, 0xc0, 0x83, 0xa0, 0xf9, 0x71,
	0x7e, 0x5a, 0x20, 0x81, 0xdd, 0xcc, 0xbd, 0xe7, 0xdc, 0x7b, 0xe6, 0xfa, 0x5c, 0xa3, 0xbb, 0x11,
	0xe7, 0x51, 0x02, 0x64, 0x18, 0x67, 0x20, 0x24, 0xcf, 0x80, 0xd0, 0x70, 0x1c, 0x33, 0xf2, 0xae,
	0xd6, 0x07, 0x49, 0x83, 0x69, 0xbc, 0xa7, 0xe3, 0x7e, 0x9a, 0x71, 0xc9, 0xf1, 0x0d, 0xc3, 0xf2,
	0x27, 0x59, 0xdf, 0x64, 0x2d, 0xcb, 0xbd, 0x6e, 0xab, 0xd2, 0x34, 0x26, 0x94, 0x31, 0x2e, 0xa9,
	0x8c, 0x39, 0x13, 0x86, 0xed, 0xee, 0x2f, 0xec, 0x09, 0x49, 0xb8, 0x24, 0x36, 0x66, 0x21, 0x1c,
	0x59, 0xec, 0x4d, 0x8b, 0x4d, 0x38, 0x8b, 0xb2, 0x9c, 0xb1, 0x98, 0x45, 0x84, 0xa7, 0x90, 0xcd,
	0x35, 0xdf, 0xb6, 0x20, 0x7d, 0xeb, 0xe7, 0x43, 0x02, 0xe3, 0x54, 0xbe, 0xb7, 0xc9, 0xea, 0xc5,
	0xa4, 0x96, 0xd2, 0x1b, 0x53, 0x31, 0x32, 0x08, 0xef, 0x93, 0x83, 0xf0, 0x8b, 0x34, 0xa4, 0x12,
	0x5a, 0x2a, 0xd5, 0x85, 0xb7, 0x39, 0x08, 0x89, 0xeb, 0x68, 0x55, 0x43, 0x2b, 0x4e, 0xd5, 0xd9,
	0xdb, 0x0c, 0x76, 0xfd, 0x3f, 0x0f, 0xc8, 0x37, 0x64, 0xc3, 0xc1, 0x75, 0xb4, 0x99, 0xeb, 0x92,
	0xba, 0x51, 0xa5, 0xa4, 0x4b, 0xb8, 0x45, 0x89, 0x42, 0x8b, 0xe1, 0x3c, 0xa1, 0x62, 0xd4, 0x45,
	0x06, 0xae, 0xce, 0xde, 0x2e, 0xda, 0x6a, 0x83, 0x9c, 0x13, 0x83, 0x51, 0x99, 0xd1, 0x31, 0x68,
	0x2d, 0x1b, 0x5d, 0x7d, 0xf6, 0x3e, 0xa0, 0x2b, 0x8f, 0x63, 0x61, 0x70, 0xa2, 0x00, 0x5e, 0x43,
	0x6b, 0x29, 0xcd, 0x80, 0x49, 0x0b, 0xb5, 0x37, 0x15, 0x1f, 0xc6, 0x89, 0x84, 0x4c, 0x6b, 0xd9,
	0xe8, 0xda, 0x1b, 0xde, 0x46, 0x1b, 0x29, 0x8d, 0xa0, 0x27, 0xe2, 0x63, 0xa8, 0xac, 0x54, 0x9d,
	0xbd, 0xd5, 0xee, 0xba, 0x0a, 0x1c, 0xc6, 0xc7, 0x80, 0x77, 0x10, 0xd2, 0x49, 0xc9, 0x47, 0xc0,
	0x2a, 0x65, 0x4d, 0xd4, 0xf0, 0xe7, 0x2a, 0xe0, 0x9d, 0x20, 0x3c, 0x2b, 0x40, 0xa4, 0x9c, 0x09,
	0xc0, 0x07, 0xaa, 0x93, 0x8a, 0x54, 0x9c, 0xea, 0xca, 0xf2, 0x83, 0xb3, 0x24, 0x7c, 0x1b, 0x6d,
	0x31, 0x38, 0x92, 0xbd, 0x99, 0xc6, 0x46, 0xf1, 0x7f, 0x2a, 0xdc, 0x29, 0x9a, 0x07, 0xe7, 0x65,
	0xf4, 0x7f, 0xab, 0xa8, 0x78, 0x5f, 0x15, 0xc4, 0x67, 0x0e, 0x5a, 0x2f, 0x06, 0x87, 0xc9, 0xa2,
	0xb6, 0x17, 0x46, 0xec, 0x2e, 0xa7, 0xd3, 0x7b, 0xf4, 0xf1, 0xfb, 0x8f, 0xcf, 0xa5, 0x07, 0xb8,
	0x31, 0xf1, 0xeb, 0x89, 0xfa, 0x1a, 0x07, 0x69, 0xc6, 0xdf, 0xc0, 0x40, 0x0a, 0xb2, 0x4f, 0x42,
	0x2a, 0x69, 0x9f, 0x0a, 0x50, 0xe7, 0x01, 0x4f, 0x12, 0x18, 0x28, 0xb7, 0xb6, 0x33, 0x9e, 0xa7,
	0x2a, 0x64, 0x1e, 0x4a, 0xf6, 0x4f, 0xf1, 0x37, 0x07, 0x6d, 0xce, 0x38, 0x0f, 0x07, 0x8b, 0x24,
	0x5c, 0xb6, 0xa9, 0xbb, 0x53, 0x70, 0x66, 0x56, 0xc4, 0x7f, 0x5a, 0xac, 0x88, 0xf7, 0x4a, 0xcb,
	0x7d, 0x16, 0x3c, 0x9c, 0xca, 0x35, 0xbb, 0xf8, 0xaf, 0xa2, 0xef, 0x59, 0x87, 0x9f, 0x3b, 0x08,
	0x4d, 0xbf, 0x3e, 0xae, 0x2d, 0x92, 0x7e, 0xc9, 0xaa, 0x6e, 0xf0, 0x37, 0x14, 0x63, 0xae, 0x5f,
	0x4d, 0xdf, 0x18, 0x7c, 0xf9, 0xa7, 0x9c, 0xda, 0xb7, 0x34, 0xce, 0x1c, 0xe4, 0x0d, 0xf8, 0x78,
	0x81, 0x8a, 0xc6, 0xd5, 0x79, 0x97, 0x75, 0xd4, 0xee, 0x76, 0x9c, 0xd7, 0x4d, 0x4b, 0x8b, 0x78,
	0x42, 0x59, 0xe4, 0xf3, 0x2c, 0x22, 0x11, 0x30, 0xbd, 0xd9, 0xc4, 0xa4, 0x68, 0x1a, 0x8b, 0xdf,
	0xfd, 0xe4, 0xea, 0xfa, 0xf6, 0xa5, 0x54, 0x6e, 0x37, 0x5b, 0x87, 0x5f, 0x4b, 0xb7, 0xda, 0xa6,
	0x58, 0x33, 0xe1, 0x79, 0xe8, 0x4f, 0x1a, 0xfa, 0xba, 0xa3, 0xff, 0xb2, 0xd6, 0x50, 0x9c, 0xfe,
	0x9a, 0xae, 0x7e, 0xe7, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc6, 0xf6, 0x28, 0xdf, 0xe1, 0x05,
	0x00, 0x00,
}
