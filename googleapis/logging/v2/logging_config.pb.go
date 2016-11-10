// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/logging/v2/logging_config.proto
// DO NOT EDIT!

package v2 // import "google.golang.org/genproto/googleapis/logging/v2"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"
import google_protobuf5 "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Available log entry formats. Log entries can be written to Cloud
// Logging in either format and can be exported in either format.
// Version 2 is the preferred format.
type LogSink_VersionFormat int32

const (
	// An unspecified version format will default to V2.
	LogSink_VERSION_FORMAT_UNSPECIFIED LogSink_VersionFormat = 0
	// `LogEntry` version 2 format.
	LogSink_V2 LogSink_VersionFormat = 1
	// `LogEntry` version 1 format.
	LogSink_V1 LogSink_VersionFormat = 2
)

var LogSink_VersionFormat_name = map[int32]string{
	0: "VERSION_FORMAT_UNSPECIFIED",
	1: "V2",
	2: "V1",
}
var LogSink_VersionFormat_value = map[string]int32{
	"VERSION_FORMAT_UNSPECIFIED": 0,
	"V2": 1,
	"V1": 2,
}

func (x LogSink_VersionFormat) String() string {
	return proto.EnumName(LogSink_VersionFormat_name, int32(x))
}
func (LogSink_VersionFormat) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

// Describes a sink used to export log entries outside Stackdriver Logging.
type LogSink struct {
	// Required. The client-assigned sink identifier, unique within the
	// project. Example: `"my-syslog-errors-to-pubsub"`.  Sink identifiers are
	// limited to 1000 characters and can include only the following characters:
	// `A-Z`, `a-z`, `0-9`, and the special characters `_-.`.  The maximum length
	// of the name is 100 characters.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Required. The export destination. See
	// [Exporting Logs With Sinks](/logging/docs/api/tasks/exporting-logs).
	// Examples:
	//
	//     "storage.googleapis.com/my-gcs-bucket"
	//     "bigquery.googleapis.com/projects/my-project-id/datasets/my-dataset"
	//     "pubsub.googleapis.com/projects/my-project/topics/my-topic"
	Destination string `protobuf:"bytes,3,opt,name=destination" json:"destination,omitempty"`
	// Optional. An [advanced logs filter](/logging/docs/view/advanced_filters).
	// Only log entries matching the filter are exported. The filter
	// must be consistent with the log entry format specified by the
	// `outputVersionFormat` parameter, regardless of the format of the
	// log entry that was originally written to Stackdriver Logging.
	// Example filter (V2 format):
	//
	//     logName=projects/my-projectid/logs/syslog AND severity>=ERROR
	Filter string `protobuf:"bytes,5,opt,name=filter" json:"filter,omitempty"`
	// Optional. The log entry version to use for this sink's exported log
	// entries.  This version does not have to correspond to the version of the
	// log entry that was written to Stackdriver Logging. If omitted, the
	// v2 format is used.
	OutputVersionFormat LogSink_VersionFormat `protobuf:"varint,6,opt,name=output_version_format,json=outputVersionFormat,enum=google.logging.v2.LogSink_VersionFormat" json:"output_version_format,omitempty"`
	// Output only. The iam identity to which the destination needs to grant write
	// access.  This may be a service account or a group.
	// Examples (Do not assume these specific values):
	//    "serviceAccount:cloud-logs@system.gserviceaccount.com"
	//    "group:cloud-logs@google.com"
	//
	//   For GCS destinations, the role "roles/owner" is required on the bucket
	//   For Cloud Pubsub destinations, the role "roles/pubsub.publisher" is
	//     required on the topic
	//   For BigQuery, the role "roles/editor" is required on the dataset
	WriterIdentity string `protobuf:"bytes,8,opt,name=writer_identity,json=writerIdentity" json:"writer_identity,omitempty"`
}

func (m *LogSink) Reset()                    { *m = LogSink{} }
func (m *LogSink) String() string            { return proto.CompactTextString(m) }
func (*LogSink) ProtoMessage()               {}
func (*LogSink) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

// The parameters to `ListSinks`.
type ListSinksRequest struct {
	// Required. The cloud resource containing the sinks.
	// Example: `"projects/my-logging-project"`.
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	// Optional. If present, then retrieve the next batch of results from the
	// preceding call to this method.  `pageToken` must be the value of
	// `nextPageToken` from the previous response.  The values of other method
	// parameters should be identical to those in the previous call.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
	// Optional. The maximum number of results to return from this request.
	// Non-positive values are ignored.  The presence of `nextPageToken` in the
	// response indicates that more results might be available.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
}

func (m *ListSinksRequest) Reset()                    { *m = ListSinksRequest{} }
func (m *ListSinksRequest) String() string            { return proto.CompactTextString(m) }
func (*ListSinksRequest) ProtoMessage()               {}
func (*ListSinksRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

// Result returned from `ListSinks`.
type ListSinksResponse struct {
	// A list of sinks.
	Sinks []*LogSink `protobuf:"bytes,1,rep,name=sinks" json:"sinks,omitempty"`
	// If there might be more results than appear in this response, then
	// `nextPageToken` is included.  To get the next set of results, call the same
	// method again using the value of `nextPageToken` as `pageToken`.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListSinksResponse) Reset()                    { *m = ListSinksResponse{} }
func (m *ListSinksResponse) String() string            { return proto.CompactTextString(m) }
func (*ListSinksResponse) ProtoMessage()               {}
func (*ListSinksResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *ListSinksResponse) GetSinks() []*LogSink {
	if m != nil {
		return m.Sinks
	}
	return nil
}

// The parameters to `GetSink`.
type GetSinkRequest struct {
	// Required. The resource name of the sink to return.
	// Example: `"projects/my-project-id/sinks/my-sink-id"`.
	SinkName string `protobuf:"bytes,1,opt,name=sink_name,json=sinkName" json:"sink_name,omitempty"`
}

func (m *GetSinkRequest) Reset()                    { *m = GetSinkRequest{} }
func (m *GetSinkRequest) String() string            { return proto.CompactTextString(m) }
func (*GetSinkRequest) ProtoMessage()               {}
func (*GetSinkRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

// The parameters to `CreateSink`.
type CreateSinkRequest struct {
	// Required. The resource in which to create the sink.
	// Example: `"projects/my-project-id"`.
	// The new sink must be provided in the request.
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	// Required. The new sink, whose `name` parameter is a sink identifier that
	// is not already in use.
	Sink *LogSink `protobuf:"bytes,2,opt,name=sink" json:"sink,omitempty"`
}

func (m *CreateSinkRequest) Reset()                    { *m = CreateSinkRequest{} }
func (m *CreateSinkRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateSinkRequest) ProtoMessage()               {}
func (*CreateSinkRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *CreateSinkRequest) GetSink() *LogSink {
	if m != nil {
		return m.Sink
	}
	return nil
}

// The parameters to `UpdateSink`.
type UpdateSinkRequest struct {
	// Required. The resource name of the sink to update, including the parent
	// resource and the sink identifier.  If the sink does not exist, this method
	// creates the sink.  Example: `"projects/my-project-id/sinks/my-sink-id"`.
	SinkName string `protobuf:"bytes,1,opt,name=sink_name,json=sinkName" json:"sink_name,omitempty"`
	// Required. The updated sink, whose name is the same identifier that appears
	// as part of `sinkName`.  If `sinkName` does not exist, then
	// this method creates a new sink.
	Sink *LogSink `protobuf:"bytes,2,opt,name=sink" json:"sink,omitempty"`
}

func (m *UpdateSinkRequest) Reset()                    { *m = UpdateSinkRequest{} }
func (m *UpdateSinkRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateSinkRequest) ProtoMessage()               {}
func (*UpdateSinkRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *UpdateSinkRequest) GetSink() *LogSink {
	if m != nil {
		return m.Sink
	}
	return nil
}

// The parameters to `DeleteSink`.
type DeleteSinkRequest struct {
	// Required. The resource name of the sink to delete, including the parent
	// resource and the sink identifier.  Example:
	// `"projects/my-project-id/sinks/my-sink-id"`.  It is an error if the sink
	// does not exist.
	SinkName string `protobuf:"bytes,1,opt,name=sink_name,json=sinkName" json:"sink_name,omitempty"`
}

func (m *DeleteSinkRequest) Reset()                    { *m = DeleteSinkRequest{} }
func (m *DeleteSinkRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteSinkRequest) ProtoMessage()               {}
func (*DeleteSinkRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func init() {
	proto.RegisterType((*LogSink)(nil), "google.logging.v2.LogSink")
	proto.RegisterType((*ListSinksRequest)(nil), "google.logging.v2.ListSinksRequest")
	proto.RegisterType((*ListSinksResponse)(nil), "google.logging.v2.ListSinksResponse")
	proto.RegisterType((*GetSinkRequest)(nil), "google.logging.v2.GetSinkRequest")
	proto.RegisterType((*CreateSinkRequest)(nil), "google.logging.v2.CreateSinkRequest")
	proto.RegisterType((*UpdateSinkRequest)(nil), "google.logging.v2.UpdateSinkRequest")
	proto.RegisterType((*DeleteSinkRequest)(nil), "google.logging.v2.DeleteSinkRequest")
	proto.RegisterEnum("google.logging.v2.LogSink_VersionFormat", LogSink_VersionFormat_name, LogSink_VersionFormat_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ConfigServiceV2 service

type ConfigServiceV2Client interface {
	// Lists sinks.
	ListSinks(ctx context.Context, in *ListSinksRequest, opts ...grpc.CallOption) (*ListSinksResponse, error)
	// Gets a sink.
	GetSink(ctx context.Context, in *GetSinkRequest, opts ...grpc.CallOption) (*LogSink, error)
	// Creates a sink.
	CreateSink(ctx context.Context, in *CreateSinkRequest, opts ...grpc.CallOption) (*LogSink, error)
	// Updates or creates a sink.
	UpdateSink(ctx context.Context, in *UpdateSinkRequest, opts ...grpc.CallOption) (*LogSink, error)
	// Deletes a sink.
	DeleteSink(ctx context.Context, in *DeleteSinkRequest, opts ...grpc.CallOption) (*google_protobuf5.Empty, error)
}

type configServiceV2Client struct {
	cc *grpc.ClientConn
}

func NewConfigServiceV2Client(cc *grpc.ClientConn) ConfigServiceV2Client {
	return &configServiceV2Client{cc}
}

func (c *configServiceV2Client) ListSinks(ctx context.Context, in *ListSinksRequest, opts ...grpc.CallOption) (*ListSinksResponse, error) {
	out := new(ListSinksResponse)
	err := grpc.Invoke(ctx, "/google.logging.v2.ConfigServiceV2/ListSinks", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceV2Client) GetSink(ctx context.Context, in *GetSinkRequest, opts ...grpc.CallOption) (*LogSink, error) {
	out := new(LogSink)
	err := grpc.Invoke(ctx, "/google.logging.v2.ConfigServiceV2/GetSink", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceV2Client) CreateSink(ctx context.Context, in *CreateSinkRequest, opts ...grpc.CallOption) (*LogSink, error) {
	out := new(LogSink)
	err := grpc.Invoke(ctx, "/google.logging.v2.ConfigServiceV2/CreateSink", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceV2Client) UpdateSink(ctx context.Context, in *UpdateSinkRequest, opts ...grpc.CallOption) (*LogSink, error) {
	out := new(LogSink)
	err := grpc.Invoke(ctx, "/google.logging.v2.ConfigServiceV2/UpdateSink", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceV2Client) DeleteSink(ctx context.Context, in *DeleteSinkRequest, opts ...grpc.CallOption) (*google_protobuf5.Empty, error) {
	out := new(google_protobuf5.Empty)
	err := grpc.Invoke(ctx, "/google.logging.v2.ConfigServiceV2/DeleteSink", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ConfigServiceV2 service

type ConfigServiceV2Server interface {
	// Lists sinks.
	ListSinks(context.Context, *ListSinksRequest) (*ListSinksResponse, error)
	// Gets a sink.
	GetSink(context.Context, *GetSinkRequest) (*LogSink, error)
	// Creates a sink.
	CreateSink(context.Context, *CreateSinkRequest) (*LogSink, error)
	// Updates or creates a sink.
	UpdateSink(context.Context, *UpdateSinkRequest) (*LogSink, error)
	// Deletes a sink.
	DeleteSink(context.Context, *DeleteSinkRequest) (*google_protobuf5.Empty, error)
}

func RegisterConfigServiceV2Server(s *grpc.Server, srv ConfigServiceV2Server) {
	s.RegisterService(&_ConfigServiceV2_serviceDesc, srv)
}

func _ConfigServiceV2_ListSinks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSinksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceV2Server).ListSinks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.logging.v2.ConfigServiceV2/ListSinks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceV2Server).ListSinks(ctx, req.(*ListSinksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigServiceV2_GetSink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceV2Server).GetSink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.logging.v2.ConfigServiceV2/GetSink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceV2Server).GetSink(ctx, req.(*GetSinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigServiceV2_CreateSink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceV2Server).CreateSink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.logging.v2.ConfigServiceV2/CreateSink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceV2Server).CreateSink(ctx, req.(*CreateSinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigServiceV2_UpdateSink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceV2Server).UpdateSink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.logging.v2.ConfigServiceV2/UpdateSink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceV2Server).UpdateSink(ctx, req.(*UpdateSinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigServiceV2_DeleteSink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceV2Server).DeleteSink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.logging.v2.ConfigServiceV2/DeleteSink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceV2Server).DeleteSink(ctx, req.(*DeleteSinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConfigServiceV2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.logging.v2.ConfigServiceV2",
	HandlerType: (*ConfigServiceV2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListSinks",
			Handler:    _ConfigServiceV2_ListSinks_Handler,
		},
		{
			MethodName: "GetSink",
			Handler:    _ConfigServiceV2_GetSink_Handler,
		},
		{
			MethodName: "CreateSink",
			Handler:    _ConfigServiceV2_CreateSink_Handler,
		},
		{
			MethodName: "UpdateSink",
			Handler:    _ConfigServiceV2_UpdateSink_Handler,
		},
		{
			MethodName: "DeleteSink",
			Handler:    _ConfigServiceV2_DeleteSink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google.golang.org/genproto/googleapis/logging/v2/logging_config.proto",
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/logging/v2/logging_config.proto", fileDescriptor1)
}

var fileDescriptor1 = []byte{
	// 724 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xef, 0x4e, 0x1a, 0x4b,
	0x14, 0xbf, 0xa0, 0xa2, 0x1c, 0xa3, 0xc2, 0xdc, 0x68, 0xc8, 0x1a, 0xef, 0xe5, 0xee, 0xd5, 0x7b,
	0x09, 0x37, 0x77, 0x97, 0xe2, 0xb7, 0x36, 0x4d, 0x53, 0x15, 0x1b, 0x12, 0xab, 0x06, 0x94, 0x0f,
	0xb6, 0xc9, 0x76, 0xc5, 0xc3, 0x76, 0x2a, 0x3b, 0xb3, 0xdd, 0x19, 0x68, 0xd5, 0x98, 0x34, 0x7d,
	0x85, 0x3e, 0x40, 0x1f, 0xaa, 0xaf, 0xd0, 0xe7, 0x68, 0x9a, 0x9d, 0x59, 0x04, 0x05, 0x29, 0x7e,
	0xd1, 0x99, 0xdf, 0xf9, 0xf3, 0x3b, 0xe7, 0x77, 0xce, 0x0e, 0x50, 0xf1, 0x38, 0xf7, 0xda, 0x68,
	0x79, 0xbc, 0xed, 0x32, 0xcf, 0xe2, 0xa1, 0x67, 0x7b, 0xc8, 0x82, 0x90, 0x4b, 0x6e, 0x6b, 0x93,
	0x1b, 0x50, 0x61, 0xb7, 0xb9, 0xe7, 0x51, 0xe6, 0xd9, 0xdd, 0x72, 0xef, 0xe8, 0x34, 0x39, 0x6b,
	0x51, 0xcf, 0x52, 0xae, 0x24, 0x1b, 0xa7, 0x89, 0x8d, 0x56, 0xb7, 0x6c, 0x54, 0x27, 0xcb, 0xec,
	0x06, 0xd4, 0x16, 0x18, 0x76, 0x69, 0x13, 0x75, 0x46, 0xdb, 0x65, 0x8c, 0x4b, 0x57, 0x52, 0xce,
	0x84, 0xce, 0x6e, 0x6c, 0x7a, 0x54, 0xbe, 0xed, 0x9c, 0x5a, 0x4d, 0xee, 0xdb, 0x3a, 0x9d, 0xad,
	0x0c, 0xa7, 0x9d, 0x96, 0x1d, 0xc8, 0x8b, 0x00, 0x85, 0x8d, 0x7e, 0x20, 0x2f, 0xf4, 0xdf, 0x38,
	0xe8, 0xc9, 0xaf, 0x83, 0x24, 0xf5, 0x51, 0x48, 0xd7, 0x0f, 0xfa, 0x27, 0x1d, 0x6c, 0x7e, 0x4d,
	0xc2, 0xec, 0x1e, 0xf7, 0xea, 0x94, 0x9d, 0x13, 0x02, 0xd3, 0xcc, 0xf5, 0x31, 0x97, 0xc8, 0x27,
	0x0a, 0xe9, 0x9a, 0x3a, 0x93, 0x3c, 0xcc, 0x9f, 0xa1, 0x90, 0x94, 0xa9, 0x3a, 0x73, 0x53, 0xca,
	0x34, 0x08, 0x91, 0x15, 0x48, 0xb5, 0x68, 0x5b, 0x62, 0x98, 0x9b, 0x51, 0xc6, 0xf8, 0x46, 0x5e,
	0xc3, 0x32, 0xef, 0xc8, 0xa0, 0x23, 0x9d, 0x2e, 0x86, 0x82, 0x72, 0xe6, 0xb4, 0x78, 0xe8, 0xbb,
	0x32, 0x97, 0xca, 0x27, 0x0a, 0x8b, 0xe5, 0x82, 0x35, 0xa4, 0xa4, 0x15, 0x17, 0x62, 0x35, 0x74,
	0xc0, 0xae, 0xf2, 0xaf, 0xfd, 0xae, 0xd3, 0xdc, 0x02, 0xc9, 0xbf, 0xb0, 0xf4, 0x21, 0xa4, 0x12,
	0x43, 0x87, 0x9e, 0x21, 0x93, 0x54, 0x5e, 0xe4, 0xe6, 0x14, 0xfd, 0xa2, 0x86, 0xab, 0x31, 0x6a,
	0x3e, 0x83, 0x85, 0xdb, 0x91, 0x7f, 0x80, 0xd1, 0xa8, 0xd4, 0xea, 0xd5, 0x83, 0x7d, 0x67, 0xf7,
	0xa0, 0xf6, 0xf2, 0xf9, 0x91, 0x73, 0xbc, 0x5f, 0x3f, 0xac, 0x6c, 0x57, 0x77, 0xab, 0x95, 0x9d,
	0xcc, 0x6f, 0x24, 0x05, 0xc9, 0x46, 0x39, 0x93, 0x50, 0xff, 0x1f, 0x65, 0x92, 0x66, 0x0b, 0x32,
	0x7b, 0x54, 0xc8, 0xa8, 0x30, 0x51, 0xc3, 0xf7, 0x1d, 0x14, 0x32, 0xea, 0x39, 0x70, 0x43, 0x64,
	0x32, 0xd6, 0x2a, 0xbe, 0x91, 0x35, 0x80, 0xc0, 0xf5, 0xd0, 0x91, 0xfc, 0x1c, 0x59, 0x2e, 0xa9,
	0x6c, 0xe9, 0x08, 0x39, 0x8a, 0x00, 0xb2, 0x0a, 0xea, 0xe2, 0x08, 0x7a, 0x89, 0x4a, 0xca, 0x99,
	0xda, 0x5c, 0x04, 0xd4, 0xe9, 0x25, 0x9a, 0x3e, 0x64, 0x07, 0x78, 0x44, 0xc0, 0x99, 0x40, 0x52,
	0x82, 0x19, 0x11, 0x01, 0xb9, 0x44, 0x7e, 0xaa, 0x30, 0x5f, 0x36, 0xee, 0x17, 0xad, 0xa6, 0x1d,
	0xc9, 0x3f, 0xb0, 0xc4, 0xf0, 0xa3, 0x74, 0x86, 0xea, 0x58, 0x88, 0xe0, 0xc3, 0x5e, 0x2d, 0xe6,
	0xff, 0xb0, 0xf8, 0x02, 0x15, 0x5b, 0xaf, 0xa9, 0x55, 0x48, 0x47, 0x29, 0x9c, 0x81, 0x1d, 0x98,
	0x8b, 0x80, 0x7d, 0xd7, 0x47, 0xf3, 0x15, 0x64, 0xb7, 0x43, 0x74, 0x25, 0x0e, 0x46, 0xdc, 0x27,
	0x83, 0x05, 0xd3, 0x51, 0xa0, 0x22, 0x1e, 0x5f, 0xb4, 0xf2, 0x33, 0xdf, 0x40, 0xf6, 0x38, 0x38,
	0xbb, 0x93, 0x7c, 0x5c, 0x39, 0x0f, 0x66, 0x28, 0x41, 0x76, 0x07, 0xdb, 0x38, 0x39, 0x43, 0xf9,
	0xc7, 0x34, 0x2c, 0x6d, 0xab, 0xef, 0xb4, 0xae, 0x3f, 0xda, 0x46, 0x99, 0x5c, 0x43, 0xfa, 0x66,
	0x44, 0xe4, 0xef, 0x51, 0xa4, 0x77, 0x16, 0xc5, 0x58, 0x1f, 0xef, 0xa4, 0xa7, 0x6c, 0x6e, 0x7c,
	0xfe, 0xf6, 0xfd, 0x4b, 0xf2, 0x4f, 0xb2, 0x16, 0x3d, 0x3b, 0x57, 0x5a, 0xc4, 0xa7, 0x41, 0xc8,
	0xdf, 0x61, 0x53, 0x0a, 0xbb, 0x78, 0x6d, 0xeb, 0xd1, 0x4a, 0x98, 0x8d, 0x47, 0x46, 0xfe, 0x1a,
	0x91, 0xf7, 0xf6, 0x38, 0x8d, 0x31, 0xa2, 0x98, 0x45, 0x45, 0xb8, 0x4e, 0x4c, 0x45, 0x78, 0x23,
	0xc2, 0x00, 0xa7, 0xa6, 0xb4, 0x8b, 0xd7, 0xe4, 0x0a, 0xa0, 0x3f, 0x79, 0x32, 0xaa, 0xa1, 0xa1,
	0xc5, 0x18, 0xcb, 0xfd, 0x9f, 0xe2, 0xde, 0x30, 0xc7, 0x37, 0xfb, 0x58, 0xcd, 0x8d, 0x7c, 0x4a,
	0x00, 0xf4, 0x57, 0x63, 0x24, 0xfb, 0xd0, 0xe6, 0x8c, 0x65, 0x2f, 0x29, 0xf6, 0xa2, 0x31, 0x41,
	0xe7, 0x71, 0x09, 0x5d, 0x80, 0xfe, 0xea, 0x8c, 0xac, 0x60, 0x68, 0xb3, 0x8c, 0x95, 0x9e, 0x57,
	0xef, 0x21, 0xb6, 0x2a, 0xd1, 0x83, 0xdd, 0xd3, 0xbd, 0x38, 0x01, 0xfb, 0xd6, 0x09, 0x2c, 0x37,
	0xb9, 0x3f, 0x4c, 0xb7, 0xb5, 0xb0, 0xa7, 0xcf, 0x7a, 0x3b, 0x0f, 0x13, 0x27, 0xa5, 0x87, 0xfe,
	0xb4, 0x9d, 0xa6, 0x94, 0x71, 0xf3, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x52, 0x86, 0xdf, 0xea,
	0x15, 0x07, 0x00, 0x00,
}
