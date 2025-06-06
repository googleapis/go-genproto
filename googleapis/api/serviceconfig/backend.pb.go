// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.24.4
// source: google/api/backend.proto

package serviceconfig

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Path Translation specifies how to combine the backend address with the
// request path in order to produce the appropriate forwarding URL for the
// request.
//
// Path Translation is applicable only to HTTP-based backends. Backends which
// do not accept requests over HTTP/HTTPS should leave `path_translation`
// unspecified.
type BackendRule_PathTranslation int32

const (
	BackendRule_PATH_TRANSLATION_UNSPECIFIED BackendRule_PathTranslation = 0
	// Use the backend address as-is, with no modification to the path. If the
	// URL pattern contains variables, the variable names and values will be
	// appended to the query string. If a query string parameter and a URL
	// pattern variable have the same name, this may result in duplicate keys in
	// the query string.
	//
	// # Examples
	//
	// Given the following operation config:
	//
	//	Method path:        /api/company/{cid}/user/{uid}
	//	Backend address:    https://example.cloudfunctions.net/getUser
	//
	// Requests to the following request paths will call the backend at the
	// translated path:
	//
	//	Request path: /api/company/widgetworks/user/johndoe
	//	Translated:
	//	https://example.cloudfunctions.net/getUser?cid=widgetworks&uid=johndoe
	//
	//	Request path: /api/company/widgetworks/user/johndoe?timezone=EST
	//	Translated:
	//	https://example.cloudfunctions.net/getUser?timezone=EST&cid=widgetworks&uid=johndoe
	BackendRule_CONSTANT_ADDRESS BackendRule_PathTranslation = 1
	// The request path will be appended to the backend address.
	//
	// # Examples
	//
	// Given the following operation config:
	//
	//	Method path:        /api/company/{cid}/user/{uid}
	//	Backend address:    https://example.appspot.com
	//
	// Requests to the following request paths will call the backend at the
	// translated path:
	//
	//	Request path: /api/company/widgetworks/user/johndoe
	//	Translated:
	//	https://example.appspot.com/api/company/widgetworks/user/johndoe
	//
	//	Request path: /api/company/widgetworks/user/johndoe?timezone=EST
	//	Translated:
	//	https://example.appspot.com/api/company/widgetworks/user/johndoe?timezone=EST
	BackendRule_APPEND_PATH_TO_ADDRESS BackendRule_PathTranslation = 2
)

// Enum value maps for BackendRule_PathTranslation.
var (
	BackendRule_PathTranslation_name = map[int32]string{
		0: "PATH_TRANSLATION_UNSPECIFIED",
		1: "CONSTANT_ADDRESS",
		2: "APPEND_PATH_TO_ADDRESS",
	}
	BackendRule_PathTranslation_value = map[string]int32{
		"PATH_TRANSLATION_UNSPECIFIED": 0,
		"CONSTANT_ADDRESS":             1,
		"APPEND_PATH_TO_ADDRESS":       2,
	}
)

func (x BackendRule_PathTranslation) Enum() *BackendRule_PathTranslation {
	p := new(BackendRule_PathTranslation)
	*p = x
	return p
}

func (x BackendRule_PathTranslation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BackendRule_PathTranslation) Descriptor() protoreflect.EnumDescriptor {
	return file_google_api_backend_proto_enumTypes[0].Descriptor()
}

func (BackendRule_PathTranslation) Type() protoreflect.EnumType {
	return &file_google_api_backend_proto_enumTypes[0]
}

func (x BackendRule_PathTranslation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BackendRule_PathTranslation.Descriptor instead.
func (BackendRule_PathTranslation) EnumDescriptor() ([]byte, []int) {
	return file_google_api_backend_proto_rawDescGZIP(), []int{1, 0}
}

// `Backend` defines the backend configuration for a service.
type Backend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of API backend rules that apply to individual API methods.
	//
	// **NOTE:** All service configuration rules follow "last one wins" order.
	Rules []*BackendRule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (x *Backend) Reset() {
	*x = Backend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_backend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Backend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Backend) ProtoMessage() {}

func (x *Backend) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_backend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Backend.ProtoReflect.Descriptor instead.
func (*Backend) Descriptor() ([]byte, []int) {
	return file_google_api_backend_proto_rawDescGZIP(), []int{0}
}

func (x *Backend) GetRules() []*BackendRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

// A backend rule provides configuration for an individual API element.
type BackendRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Selects the methods to which this rule applies.
	//
	// Refer to [selector][google.api.DocumentationRule.selector] for syntax
	// details.
	Selector string `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	// The address of the API backend.
	//
	// The scheme is used to determine the backend protocol and security.
	// The following schemes are accepted:
	//
	//	SCHEME        PROTOCOL    SECURITY
	//	http://       HTTP        None
	//	https://      HTTP        TLS
	//	grpc://       gRPC        None
	//	grpcs://      gRPC        TLS
	//
	// It is recommended to explicitly include a scheme. Leaving out the scheme
	// may cause constrasting behaviors across platforms.
	//
	// If the port is unspecified, the default is:
	// - 80 for schemes without TLS
	// - 443 for schemes with TLS
	//
	// For HTTP backends, use [protocol][google.api.BackendRule.protocol]
	// to specify the protocol version.
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// The number of seconds to wait for a response from a request. The default
	// varies based on the request protocol and deployment environment.
	Deadline float64 `protobuf:"fixed64,3,opt,name=deadline,proto3" json:"deadline,omitempty"`
	// Deprecated, do not use.
	//
	// Deprecated: Do not use.
	MinDeadline float64 `protobuf:"fixed64,4,opt,name=min_deadline,json=minDeadline,proto3" json:"min_deadline,omitempty"`
	// The number of seconds to wait for the completion of a long running
	// operation. The default is no deadline.
	OperationDeadline float64                     `protobuf:"fixed64,5,opt,name=operation_deadline,json=operationDeadline,proto3" json:"operation_deadline,omitempty"`
	PathTranslation   BackendRule_PathTranslation `protobuf:"varint,6,opt,name=path_translation,json=pathTranslation,proto3,enum=google.api.BackendRule_PathTranslation" json:"path_translation,omitempty"`
	// Authentication settings used by the backend.
	//
	// These are typically used to provide service management functionality to
	// a backend served on a publicly-routable URL. The `authentication`
	// details should match the authentication behavior used by the backend.
	//
	// For example, specifying `jwt_audience` implies that the backend expects
	// authentication via a JWT.
	//
	// When authentication is unspecified, the resulting behavior is the same
	// as `disable_auth` set to `true`.
	//
	// Refer to https://developers.google.com/identity/protocols/OpenIDConnect for
	// JWT ID token.
	//
	// Types that are assignable to Authentication:
	//
	//	*BackendRule_JwtAudience
	//	*BackendRule_DisableAuth
	Authentication isBackendRule_Authentication `protobuf_oneof:"authentication"`
	// The protocol used for sending a request to the backend.
	// The supported values are "http/1.1" and "h2".
	//
	// The default value is inferred from the scheme in the
	// [address][google.api.BackendRule.address] field:
	//
	//	SCHEME        PROTOCOL
	//	http://       http/1.1
	//	https://      http/1.1
	//	grpc://       h2
	//	grpcs://      h2
	//
	// For secure HTTP backends (https://) that support HTTP/2, set this field
	// to "h2" for improved performance.
	//
	// Configuring this field to non-default values is only supported for secure
	// HTTP backends. This field will be ignored for all other backends.
	//
	// See
	// https://www.iana.org/assignments/tls-extensiontype-values/tls-extensiontype-values.xhtml#alpn-protocol-ids
	// for more details on the supported values.
	Protocol string `protobuf:"bytes,9,opt,name=protocol,proto3" json:"protocol,omitempty"`
	// The map between request protocol and the backend address.
	OverridesByRequestProtocol map[string]*BackendRule `protobuf:"bytes,10,rep,name=overrides_by_request_protocol,json=overridesByRequestProtocol,proto3" json:"overrides_by_request_protocol,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *BackendRule) Reset() {
	*x = BackendRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_backend_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackendRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackendRule) ProtoMessage() {}

func (x *BackendRule) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_backend_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackendRule.ProtoReflect.Descriptor instead.
func (*BackendRule) Descriptor() ([]byte, []int) {
	return file_google_api_backend_proto_rawDescGZIP(), []int{1}
}

func (x *BackendRule) GetSelector() string {
	if x != nil {
		return x.Selector
	}
	return ""
}

func (x *BackendRule) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *BackendRule) GetDeadline() float64 {
	if x != nil {
		return x.Deadline
	}
	return 0
}

// Deprecated: Do not use.
func (x *BackendRule) GetMinDeadline() float64 {
	if x != nil {
		return x.MinDeadline
	}
	return 0
}

func (x *BackendRule) GetOperationDeadline() float64 {
	if x != nil {
		return x.OperationDeadline
	}
	return 0
}

func (x *BackendRule) GetPathTranslation() BackendRule_PathTranslation {
	if x != nil {
		return x.PathTranslation
	}
	return BackendRule_PATH_TRANSLATION_UNSPECIFIED
}

func (m *BackendRule) GetAuthentication() isBackendRule_Authentication {
	if m != nil {
		return m.Authentication
	}
	return nil
}

func (x *BackendRule) GetJwtAudience() string {
	if x, ok := x.GetAuthentication().(*BackendRule_JwtAudience); ok {
		return x.JwtAudience
	}
	return ""
}

func (x *BackendRule) GetDisableAuth() bool {
	if x, ok := x.GetAuthentication().(*BackendRule_DisableAuth); ok {
		return x.DisableAuth
	}
	return false
}

func (x *BackendRule) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *BackendRule) GetOverridesByRequestProtocol() map[string]*BackendRule {
	if x != nil {
		return x.OverridesByRequestProtocol
	}
	return nil
}

type isBackendRule_Authentication interface {
	isBackendRule_Authentication()
}

type BackendRule_JwtAudience struct {
	// The JWT audience is used when generating a JWT ID token for the backend.
	// This ID token will be added in the HTTP "authorization" header, and sent
	// to the backend.
	JwtAudience string `protobuf:"bytes,7,opt,name=jwt_audience,json=jwtAudience,proto3,oneof"`
}

type BackendRule_DisableAuth struct {
	// When disable_auth is true, a JWT ID token won't be generated and the
	// original "Authorization" HTTP header will be preserved. If the header is
	// used to carry the original token and is expected by the backend, this
	// field must be set to true to preserve the header.
	DisableAuth bool `protobuf:"varint,8,opt,name=disable_auth,json=disableAuth,proto3,oneof"`
}

func (*BackendRule_JwtAudience) isBackendRule_Authentication() {}

func (*BackendRule_DisableAuth) isBackendRule_Authentication() {}

var File_google_api_backend_proto protoreflect.FileDescriptor

var file_google_api_backend_proto_rawDesc = []byte{
	0x0a, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x22, 0x38, 0x0a, 0x07, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x12, 0x2d, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73,
	0x22, 0xcc, 0x05, 0x0a, 0x0b, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x52, 0x75, 0x6c, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69,
	0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69,
	0x6e, 0x65, 0x12, 0x25, 0x0a, 0x0c, 0x6d, 0x69, 0x6e, 0x5f, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69,
	0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0b, 0x6d, 0x69,
	0x6e, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x52, 0x0a, 0x10, 0x70, 0x61, 0x74, 0x68,
	0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x50, 0x61, 0x74, 0x68,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x70, 0x61, 0x74,
	0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0c,
	0x6a, 0x77, 0x74, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x6a, 0x77, 0x74, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63,
	0x65, 0x12, 0x23, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x61, 0x75, 0x74,
	0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x41, 0x75, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x12, 0x7a, 0x0a, 0x1d, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x5f,
	0x62, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x52, 0x75,
	0x6c, 0x65, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x42, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x1a, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x42, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x1a, 0x66,
	0x0a, 0x1f, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x42, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x65, 0x0a, 0x0f, 0x50, 0x61, 0x74, 0x68, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x1c, 0x50, 0x41, 0x54,
	0x48, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x43,
	0x4f, 0x4e, 0x53, 0x54, 0x41, 0x4e, 0x54, 0x5f, 0x41, 0x44, 0x44, 0x52, 0x45, 0x53, 0x53, 0x10,
	0x01, 0x12, 0x1a, 0x0a, 0x16, 0x41, 0x50, 0x50, 0x45, 0x4e, 0x44, 0x5f, 0x50, 0x41, 0x54, 0x48,
	0x5f, 0x54, 0x4f, 0x5f, 0x41, 0x44, 0x44, 0x52, 0x45, 0x53, 0x53, 0x10, 0x02, 0x42, 0x10, 0x0a,
	0x0e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x6e, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x42, 0x0c, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0xa2, 0x02, 0x04, 0x47, 0x41, 0x50, 0x49, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_api_backend_proto_rawDescOnce sync.Once
	file_google_api_backend_proto_rawDescData = file_google_api_backend_proto_rawDesc
)

func file_google_api_backend_proto_rawDescGZIP() []byte {
	file_google_api_backend_proto_rawDescOnce.Do(func() {
		file_google_api_backend_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_api_backend_proto_rawDescData)
	})
	return file_google_api_backend_proto_rawDescData
}

var file_google_api_backend_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_api_backend_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_api_backend_proto_goTypes = []interface{}{
	(BackendRule_PathTranslation)(0), // 0: google.api.BackendRule.PathTranslation
	(*Backend)(nil),                  // 1: google.api.Backend
	(*BackendRule)(nil),              // 2: google.api.BackendRule
	nil,                              // 3: google.api.BackendRule.OverridesByRequestProtocolEntry
}
var file_google_api_backend_proto_depIdxs = []int32{
	2, // 0: google.api.Backend.rules:type_name -> google.api.BackendRule
	0, // 1: google.api.BackendRule.path_translation:type_name -> google.api.BackendRule.PathTranslation
	3, // 2: google.api.BackendRule.overrides_by_request_protocol:type_name -> google.api.BackendRule.OverridesByRequestProtocolEntry
	2, // 3: google.api.BackendRule.OverridesByRequestProtocolEntry.value:type_name -> google.api.BackendRule
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_api_backend_proto_init() }
func file_google_api_backend_proto_init() {
	if File_google_api_backend_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_api_backend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Backend); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_api_backend_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackendRule); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_google_api_backend_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*BackendRule_JwtAudience)(nil),
		(*BackendRule_DisableAuth)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_api_backend_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_api_backend_proto_goTypes,
		DependencyIndexes: file_google_api_backend_proto_depIdxs,
		EnumInfos:         file_google_api_backend_proto_enumTypes,
		MessageInfos:      file_google_api_backend_proto_msgTypes,
	}.Build()
	File_google_api_backend_proto = out.File
	file_google_api_backend_proto_rawDesc = nil
	file_google_api_backend_proto_goTypes = nil
	file_google_api_backend_proto_depIdxs = nil
}
