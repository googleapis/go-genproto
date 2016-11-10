// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/rpc/status/status.proto
// DO NOT EDIT!

/*
Package status is a generated protocol buffer package.

It is generated from these files:
	google.golang.org/genproto/googleapis/rpc/status/status.proto

It has these top-level messages:
	Status
*/
package status // import "google.golang.org/genproto/googleapis/rpc/status"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The `Status` type defines a logical error model that is suitable for different
// programming environments, including REST APIs and RPC APIs. It is used by
// [gRPC](https://github.com/grpc). The error model is designed to be:
//
// - Simple to use and understand for most users
// - Flexible enough to meet unexpected needs
//
// # Overview
//
// The `Status` message contains three pieces of data: error code, error message,
// and error details. The error code should be an enum value of
// [google.rpc.Code][google.rpc.Code], but it may accept additional error codes if needed.  The
// error message should be a developer-facing English message that helps
// developers *understand* and *resolve* the error. If a localized user-facing
// error message is needed, put the localized message in the error details or
// localize it in the client. The optional error details may contain arbitrary
// information about the error. There is a predefined set of error detail types
// in the package `google.rpc` which can be used for common error conditions.
//
// # Language mapping
//
// The `Status` message is the logical representation of the error model, but it
// is not necessarily the actual wire format. When the `Status` message is
// exposed in different client libraries and different wire protocols, it can be
// mapped differently. For example, it will likely be mapped to some exceptions
// in Java, but more likely mapped to some error codes in C.
//
// # Other uses
//
// The error model and the `Status` message can be used in a variety of
// environments, either with or without APIs, to provide a
// consistent developer experience across different environments.
//
// Example uses of this error model include:
//
// - Partial errors. If a service needs to return partial errors to the client,
//     it may embed the `Status` in the normal response to indicate the partial
//     errors.
//
// - Workflow errors. A typical workflow has multiple steps. Each step may
//     have a `Status` message for error reporting purpose.
//
// - Batch operations. If a client uses batch request and batch response, the
//     `Status` message should be used directly inside batch response, one for
//     each error sub-response.
//
// - Asynchronous operations. If an API call embeds asynchronous operation
//     results in its response, the status of those operations should be
//     represented directly using the `Status` message.
//
// - Logging. If some API errors are stored in logs, the message `Status` could
//     be used directly after any stripping needed for security/privacy reasons.
type Status struct {
	// The status code, which should be an enum value of [google.rpc.Code][google.rpc.Code].
	Code int32 `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	// A developer-facing error message, which should be in English. Any
	// user-facing error message should be localized and sent in the
	// [google.rpc.Status.details][google.rpc.Status.details] field, or localized by the client.
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	// A list of messages that carry the error details.  There will be a
	// common set of message types for APIs to use.
	Details []*google_protobuf.Any `protobuf:"bytes,3,rep,name=details" json:"details,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Status) GetDetails() []*google_protobuf.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

func init() {
	proto.RegisterType((*Status)(nil), "google.rpc.Status")
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/rpc/status/status.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xbf, 0x4b, 0xc4, 0x30,
	0x14, 0xc7, 0xe9, 0x55, 0xef, 0x30, 0x07, 0x0e, 0xc1, 0xa1, 0x38, 0x15, 0xa7, 0x4e, 0xef, 0x81,
	0x0e, 0x0e, 0xe2, 0xe0, 0xf9, 0x0f, 0x94, 0xba, 0x39, 0x08, 0x69, 0x2e, 0xf7, 0x2c, 0xb4, 0x7d,
	0x21, 0x49, 0x87, 0xfe, 0x3b, 0xfe, 0xa5, 0xd2, 0xa4, 0xc5, 0xd9, 0x21, 0xbc, 0x84, 0x7c, 0xbe,
	0x3f, 0x78, 0xe2, 0x95, 0x98, 0xa9, 0x37, 0x40, 0xdc, 0xab, 0x91, 0x80, 0x1d, 0x21, 0x99, 0xd1,
	0x3a, 0x0e, 0x8c, 0xe9, 0x4b, 0xd9, 0xce, 0xa3, 0xb3, 0x1a, 0x7d, 0x50, 0x61, 0xf2, 0xeb, 0x80,
	0x88, 0x48, 0xb1, 0xca, 0x9d, 0xd5, 0xf7, 0x48, 0x5d, 0xf8, 0x9e, 0x5a, 0xd0, 0x3c, 0x60, 0xb2,
	0xc3, 0x08, 0xb5, 0xd3, 0x05, 0x6d, 0x98, 0xad, 0xf1, 0xa8, 0xc6, 0x79, 0x39, 0x49, 0xfc, 0x70,
	0x11, 0xfb, 0x8f, 0x68, 0x26, 0xa5, 0xb8, 0xd2, 0x7c, 0x36, 0x45, 0x56, 0x66, 0xd5, 0x75, 0x13,
	0xef, 0xb2, 0x10, 0x87, 0xc1, 0x78, 0xaf, 0xc8, 0x14, 0xbb, 0x32, 0xab, 0x6e, 0x9a, 0xed, 0x29,
	0x41, 0x1c, 0xce, 0x26, 0xa8, 0xae, 0xf7, 0x45, 0x5e, 0xe6, 0xd5, 0xf1, 0xf1, 0x0e, 0xd6, 0x1a,
	0x5b, 0x1e, 0xbc, 0x8d, 0x73, 0xb3, 0x41, 0xa7, 0x2f, 0x71, 0xab, 0x79, 0x80, 0xbf, 0xaa, 0xa7,
	0x63, 0xca, 0xad, 0x17, 0xbc, 0xce, 0x3e, 0x9f, 0xff, 0xbb, 0x84, 0x97, 0x34, 0x7e, 0x76, 0x79,
	0x53, 0xbf, 0xb7, 0xfb, 0x48, 0x3e, 0xfd, 0x06, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x94, 0xe0, 0x71,
	0x4c, 0x01, 0x00, 0x00,
}
