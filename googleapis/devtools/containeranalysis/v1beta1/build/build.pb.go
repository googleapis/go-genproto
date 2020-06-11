// Copyright 2018 The Grafeas Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.12.3
// source: google/devtools/containeranalysis/v1beta1/build/build.proto

package build

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	provenance "google.golang.org/genproto/googleapis/devtools/containeranalysis/v1beta1/provenance"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Public key formats.
type BuildSignature_KeyType int32

const (
	// `KeyType` is not set.
	BuildSignature_KEY_TYPE_UNSPECIFIED BuildSignature_KeyType = 0
	// `PGP ASCII Armored` public key.
	BuildSignature_PGP_ASCII_ARMORED BuildSignature_KeyType = 1
	// `PKIX PEM` public key.
	BuildSignature_PKIX_PEM BuildSignature_KeyType = 2
)

// Enum value maps for BuildSignature_KeyType.
var (
	BuildSignature_KeyType_name = map[int32]string{
		0: "KEY_TYPE_UNSPECIFIED",
		1: "PGP_ASCII_ARMORED",
		2: "PKIX_PEM",
	}
	BuildSignature_KeyType_value = map[string]int32{
		"KEY_TYPE_UNSPECIFIED": 0,
		"PGP_ASCII_ARMORED":    1,
		"PKIX_PEM":             2,
	}
)

func (x BuildSignature_KeyType) Enum() *BuildSignature_KeyType {
	p := new(BuildSignature_KeyType)
	*p = x
	return p
}

func (x BuildSignature_KeyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BuildSignature_KeyType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_devtools_containeranalysis_v1beta1_build_build_proto_enumTypes[0].Descriptor()
}

func (BuildSignature_KeyType) Type() protoreflect.EnumType {
	return &file_google_devtools_containeranalysis_v1beta1_build_build_proto_enumTypes[0]
}

func (x BuildSignature_KeyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BuildSignature_KeyType.Descriptor instead.
func (BuildSignature_KeyType) EnumDescriptor() ([]byte, []int) {
	return file_google_devtools_containeranalysis_v1beta1_build_build_proto_rawDescGZIP(), []int{1, 0}
}

// Note holding the version of the provider's builder and the signature of the
// provenance message in the build details occurrence.
type Build struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Immutable. Version of the builder which produced this build.
	BuilderVersion string `protobuf:"bytes,1,opt,name=builder_version,json=builderVersion,proto3" json:"builder_version,omitempty"`
	// Signature of the build in occurrences pointing to this build note
	// containing build details.
	Signature *BuildSignature `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Build) Reset() {
	*x = Build{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_containeranalysis_v1beta1_build_build_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Build) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Build) ProtoMessage() {}

func (x *Build) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_containeranalysis_v1beta1_build_build_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Build.ProtoReflect.Descriptor instead.
func (*Build) Descriptor() ([]byte, []int) {
	return file_google_devtools_containeranalysis_v1beta1_build_build_proto_rawDescGZIP(), []int{0}
}

func (x *Build) GetBuilderVersion() string {
	if x != nil {
		return x.BuilderVersion
	}
	return ""
}

func (x *Build) GetSignature() *BuildSignature {
	if x != nil {
		return x.Signature
	}
	return nil
}

// Message encapsulating the signature of the verified build.
type BuildSignature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Public key of the builder which can be used to verify that the related
	// findings are valid and unchanged. If `key_type` is empty, this defaults
	// to PEM encoded public keys.
	//
	// This field may be empty if `key_id` references an external key.
	//
	// For Cloud Build based signatures, this is a PEM encoded public
	// key. To verify the Cloud Build signature, place the contents of
	// this field into a file (public.pem). The signature field is base64-decoded
	// into its binary representation in signature.bin, and the provenance bytes
	// from `BuildDetails` are base64-decoded into a binary representation in
	// signed.bin. OpenSSL can then verify the signature:
	// `openssl sha256 -verify public.pem -signature signature.bin signed.bin`
	PublicKey string `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// Required. Signature of the related `BuildProvenance`. In JSON, this is
	// base-64 encoded.
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	// An ID for the key used to sign. This could be either an ID for the key
	// stored in `public_key` (such as the ID or fingerprint for a PGP key, or the
	// CN for a cert), or a reference to an external key (such as a reference to a
	// key in Cloud Key Management Service).
	KeyId string `protobuf:"bytes,3,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	// The type of the key, either stored in `public_key` or referenced in
	// `key_id`.
	KeyType BuildSignature_KeyType `protobuf:"varint,4,opt,name=key_type,json=keyType,proto3,enum=grafeas.v1beta1.build.BuildSignature_KeyType" json:"key_type,omitempty"`
}

func (x *BuildSignature) Reset() {
	*x = BuildSignature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_containeranalysis_v1beta1_build_build_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildSignature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildSignature) ProtoMessage() {}

func (x *BuildSignature) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_containeranalysis_v1beta1_build_build_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildSignature.ProtoReflect.Descriptor instead.
func (*BuildSignature) Descriptor() ([]byte, []int) {
	return file_google_devtools_containeranalysis_v1beta1_build_build_proto_rawDescGZIP(), []int{1}
}

func (x *BuildSignature) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *BuildSignature) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *BuildSignature) GetKeyId() string {
	if x != nil {
		return x.KeyId
	}
	return ""
}

func (x *BuildSignature) GetKeyType() BuildSignature_KeyType {
	if x != nil {
		return x.KeyType
	}
	return BuildSignature_KEY_TYPE_UNSPECIFIED
}

// Details of a build occurrence.
type Details struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The actual provenance for the build.
	Provenance *provenance.BuildProvenance `protobuf:"bytes,1,opt,name=provenance,proto3" json:"provenance,omitempty"`
	// Serialized JSON representation of the provenance, used in generating the
	// build signature in the corresponding build note. After verifying the
	// signature, `provenance_bytes` can be unmarshalled and compared to the
	// provenance to confirm that it is unchanged. A base64-encoded string
	// representation of the provenance bytes is used for the signature in order
	// to interoperate with openssl which expects this format for signature
	// verification.
	//
	// The serialized form is captured both to avoid ambiguity in how the
	// provenance is marshalled to json as well to prevent incompatibilities with
	// future changes.
	ProvenanceBytes string `protobuf:"bytes,2,opt,name=provenance_bytes,json=provenanceBytes,proto3" json:"provenance_bytes,omitempty"`
}

func (x *Details) Reset() {
	*x = Details{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_containeranalysis_v1beta1_build_build_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Details) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Details) ProtoMessage() {}

func (x *Details) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_containeranalysis_v1beta1_build_build_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Details.ProtoReflect.Descriptor instead.
func (*Details) Descriptor() ([]byte, []int) {
	return file_google_devtools_containeranalysis_v1beta1_build_build_proto_rawDescGZIP(), []int{2}
}

func (x *Details) GetProvenance() *provenance.BuildProvenance {
	if x != nil {
		return x.Provenance
	}
	return nil
}

func (x *Details) GetProvenanceBytes() string {
	if x != nil {
		return x.ProvenanceBytes
	}
	return ""
}

var File_google_devtools_containeranalysis_v1beta1_build_build_proto protoreflect.FileDescriptor

var file_google_devtools_containeranalysis_v1beta1_build_build_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67,
	0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x1a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x65, 0x76,
	0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f,
	0x70, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x65,
	0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a, 0x05, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x22, 0xf8, 0x01, 0x0a, 0x0e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x48, 0x0a, 0x08, 0x6b, 0x65, 0x79,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x67, 0x72,
	0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x2e, 0x4b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x22, 0x48, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18,
	0x0a, 0x14, 0x4b, 0x45, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x47, 0x50, 0x5f,
	0x41, 0x53, 0x43, 0x49, 0x49, 0x5f, 0x41, 0x52, 0x4d, 0x4f, 0x52, 0x45, 0x44, 0x10, 0x01, 0x12,
	0x0c, 0x0a, 0x08, 0x50, 0x4b, 0x49, 0x58, 0x5f, 0x50, 0x45, 0x4d, 0x10, 0x02, 0x22, 0x81, 0x01,
	0x0a, 0x07, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x4b, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e,
	0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x70, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x50, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76,
	0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x6e,
	0x61, 0x6e, 0x63, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x42, 0x78, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x50, 0x01, 0x5a,
	0x54, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69,
	0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x3b,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0xa2, 0x02, 0x03, 0x47, 0x52, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_devtools_containeranalysis_v1beta1_build_build_proto_rawDescOnce sync.Once
	file_google_devtools_containeranalysis_v1beta1_build_build_proto_rawDescData = file_google_devtools_containeranalysis_v1beta1_build_build_proto_rawDesc
)

func file_google_devtools_containeranalysis_v1beta1_build_build_proto_rawDescGZIP() []byte {
	file_google_devtools_containeranalysis_v1beta1_build_build_proto_rawDescOnce.Do(func() {
		file_google_devtools_containeranalysis_v1beta1_build_build_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_devtools_containeranalysis_v1beta1_build_build_proto_rawDescData)
	})
	return file_google_devtools_containeranalysis_v1beta1_build_build_proto_rawDescData
}

var file_google_devtools_containeranalysis_v1beta1_build_build_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_devtools_containeranalysis_v1beta1_build_build_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_devtools_containeranalysis_v1beta1_build_build_proto_goTypes = []interface{}{
	(BuildSignature_KeyType)(0),        // 0: grafeas.v1beta1.build.BuildSignature.KeyType
	(*Build)(nil),                      // 1: grafeas.v1beta1.build.Build
	(*BuildSignature)(nil),             // 2: grafeas.v1beta1.build.BuildSignature
	(*Details)(nil),                    // 3: grafeas.v1beta1.build.Details
	(*provenance.BuildProvenance)(nil), // 4: grafeas.v1beta1.provenance.BuildProvenance
}
var file_google_devtools_containeranalysis_v1beta1_build_build_proto_depIdxs = []int32{
	2, // 0: grafeas.v1beta1.build.Build.signature:type_name -> grafeas.v1beta1.build.BuildSignature
	0, // 1: grafeas.v1beta1.build.BuildSignature.key_type:type_name -> grafeas.v1beta1.build.BuildSignature.KeyType
	4, // 2: grafeas.v1beta1.build.Details.provenance:type_name -> grafeas.v1beta1.provenance.BuildProvenance
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_devtools_containeranalysis_v1beta1_build_build_proto_init() }
func file_google_devtools_containeranalysis_v1beta1_build_build_proto_init() {
	if File_google_devtools_containeranalysis_v1beta1_build_build_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_devtools_containeranalysis_v1beta1_build_build_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Build); i {
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
		file_google_devtools_containeranalysis_v1beta1_build_build_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildSignature); i {
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
		file_google_devtools_containeranalysis_v1beta1_build_build_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Details); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_devtools_containeranalysis_v1beta1_build_build_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_devtools_containeranalysis_v1beta1_build_build_proto_goTypes,
		DependencyIndexes: file_google_devtools_containeranalysis_v1beta1_build_build_proto_depIdxs,
		EnumInfos:         file_google_devtools_containeranalysis_v1beta1_build_build_proto_enumTypes,
		MessageInfos:      file_google_devtools_containeranalysis_v1beta1_build_build_proto_msgTypes,
	}.Build()
	File_google_devtools_containeranalysis_v1beta1_build_build_proto = out.File
	file_google_devtools_containeranalysis_v1beta1_build_build_proto_rawDesc = nil
	file_google_devtools_containeranalysis_v1beta1_build_build_proto_goTypes = nil
	file_google_devtools_containeranalysis_v1beta1_build_build_proto_depIdxs = nil
}
