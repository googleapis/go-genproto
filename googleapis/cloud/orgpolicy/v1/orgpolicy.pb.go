// Copyright 2019 Google LLC.
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
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: google/cloud/orgpolicy/v1/orgpolicy.proto

package orgpolicy

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This enum can be used to set `Policies` that apply to all possible
// configuration values rather than specific values in `allowed_values` or
// `denied_values`.
//
// Settting this to `ALLOW` will mean this `Policy` allows all values.
// Similarly, setting it to `DENY` will mean no values are allowed. If
// set to either `ALLOW` or `DENY,  `allowed_values` and `denied_values`
// must be unset. Setting this to `ALL_VALUES_UNSPECIFIED` allows for
// setting `allowed_values` and `denied_values`.
type Policy_ListPolicy_AllValues int32

const (
	// Indicates that allowed_values or denied_values must be set.
	Policy_ListPolicy_ALL_VALUES_UNSPECIFIED Policy_ListPolicy_AllValues = 0
	// A policy with this set allows all values.
	Policy_ListPolicy_ALLOW Policy_ListPolicy_AllValues = 1
	// A policy with this set denies all values.
	Policy_ListPolicy_DENY Policy_ListPolicy_AllValues = 2
)

// Enum value maps for Policy_ListPolicy_AllValues.
var (
	Policy_ListPolicy_AllValues_name = map[int32]string{
		0: "ALL_VALUES_UNSPECIFIED",
		1: "ALLOW",
		2: "DENY",
	}
	Policy_ListPolicy_AllValues_value = map[string]int32{
		"ALL_VALUES_UNSPECIFIED": 0,
		"ALLOW":                  1,
		"DENY":                   2,
	}
)

func (x Policy_ListPolicy_AllValues) Enum() *Policy_ListPolicy_AllValues {
	p := new(Policy_ListPolicy_AllValues)
	*p = x
	return p
}

func (x Policy_ListPolicy_AllValues) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Policy_ListPolicy_AllValues) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_orgpolicy_v1_orgpolicy_proto_enumTypes[0].Descriptor()
}

func (Policy_ListPolicy_AllValues) Type() protoreflect.EnumType {
	return &file_google_cloud_orgpolicy_v1_orgpolicy_proto_enumTypes[0]
}

func (x Policy_ListPolicy_AllValues) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Policy_ListPolicy_AllValues.Descriptor instead.
func (Policy_ListPolicy_AllValues) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_orgpolicy_v1_orgpolicy_proto_rawDescGZIP(), []int{0, 0, 0}
}

// Defines a Cloud Organization `Policy` which is used to specify `Constraints`
// for configurations of Cloud Platform resources.
type Policy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Version of the `Policy`. Default version is 0;
	Version int32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// The name of the `Constraint` the `Policy` is configuring, for example,
	// `constraints/serviceuser.services`.
	//
	// Immutable after creation.
	Constraint string `protobuf:"bytes,2,opt,name=constraint,proto3" json:"constraint,omitempty"`
	// An opaque tag indicating the current version of the `Policy`, used for
	// concurrency control.
	//
	// When the `Policy` is returned from either a `GetPolicy` or a
	// `ListOrgPolicy` request, this `etag` indicates the version of the current
	// `Policy` to use when executing a read-modify-write loop.
	//
	// When the `Policy` is returned from a `GetEffectivePolicy` request, the
	// `etag` will be unset.
	//
	// When the `Policy` is used in a `SetOrgPolicy` method, use the `etag` value
	// that was returned from a `GetOrgPolicy` request as part of a
	// read-modify-write loop for concurrency control. Not setting the `etag`in a
	// `SetOrgPolicy` request will result in an unconditional write of the
	// `Policy`.
	Etag []byte `protobuf:"bytes,3,opt,name=etag,proto3" json:"etag,omitempty"`
	// The time stamp the `Policy` was previously updated. This is set by the
	// server, not specified by the caller, and represents the last time a call to
	// `SetOrgPolicy` was made for that `Policy`. Any value set by the client will
	// be ignored.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// The field to populate is based on the `constraint_type` value in the
	// `Constraint`.
	//   `list_constraint` => `list_policy`
	//   `boolean_constraint` => `boolean_policy`
	//
	//  A `restore_default` message may be used with any `Constraint` type.
	//
	// Providing a *_policy that is incompatible with the `constraint_type` will
	// result in an `invalid_argument` error.
	//
	// Attempting to set a `Policy` with a `policy_type` not set will result in an
	// `invalid_argument` error.
	//
	// Types that are assignable to PolicyType:
	//	*Policy_ListPolicy_
	//	*Policy_BooleanPolicy_
	//	*Policy_RestoreDefault_
	PolicyType isPolicy_PolicyType `protobuf_oneof:"policy_type"`
}

func (x *Policy) Reset() {
	*x = Policy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_orgpolicy_v1_orgpolicy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy) ProtoMessage() {}

func (x *Policy) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_orgpolicy_v1_orgpolicy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy.ProtoReflect.Descriptor instead.
func (*Policy) Descriptor() ([]byte, []int) {
	return file_google_cloud_orgpolicy_v1_orgpolicy_proto_rawDescGZIP(), []int{0}
}

func (x *Policy) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Policy) GetConstraint() string {
	if x != nil {
		return x.Constraint
	}
	return ""
}

func (x *Policy) GetEtag() []byte {
	if x != nil {
		return x.Etag
	}
	return nil
}

func (x *Policy) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (m *Policy) GetPolicyType() isPolicy_PolicyType {
	if m != nil {
		return m.PolicyType
	}
	return nil
}

func (x *Policy) GetListPolicy() *Policy_ListPolicy {
	if x, ok := x.GetPolicyType().(*Policy_ListPolicy_); ok {
		return x.ListPolicy
	}
	return nil
}

func (x *Policy) GetBooleanPolicy() *Policy_BooleanPolicy {
	if x, ok := x.GetPolicyType().(*Policy_BooleanPolicy_); ok {
		return x.BooleanPolicy
	}
	return nil
}

func (x *Policy) GetRestoreDefault() *Policy_RestoreDefault {
	if x, ok := x.GetPolicyType().(*Policy_RestoreDefault_); ok {
		return x.RestoreDefault
	}
	return nil
}

type isPolicy_PolicyType interface {
	isPolicy_PolicyType()
}

type Policy_ListPolicy_ struct {
	// List of values either allowed or disallowed.
	ListPolicy *Policy_ListPolicy `protobuf:"bytes,5,opt,name=list_policy,json=listPolicy,proto3,oneof"`
}

type Policy_BooleanPolicy_ struct {
	// For boolean `Constraints`, whether to enforce the `Constraint` or not.
	BooleanPolicy *Policy_BooleanPolicy `protobuf:"bytes,6,opt,name=boolean_policy,json=booleanPolicy,proto3,oneof"`
}

type Policy_RestoreDefault_ struct {
	// Restores the default behavior of the constraint; independent of
	// `Constraint` type.
	RestoreDefault *Policy_RestoreDefault `protobuf:"bytes,7,opt,name=restore_default,json=restoreDefault,proto3,oneof"`
}

func (*Policy_ListPolicy_) isPolicy_PolicyType() {}

func (*Policy_BooleanPolicy_) isPolicy_PolicyType() {}

func (*Policy_RestoreDefault_) isPolicy_PolicyType() {}

// Used in `policy_type` to specify how `list_policy` behaves at this
// resource.
//
// `ListPolicy` can define specific values and subtrees of Cloud Resource
// Manager resource hierarchy (`Organizations`, `Folders`, `Projects`) that
// are allowed or denied by setting the `allowed_values` and `denied_values`
// fields. This is achieved by using the `under:` and optional `is:` prefixes.
// The `under:` prefix is used to denote resource subtree values.
// The `is:` prefix is used to denote specific values, and is required only
// if the value contains a ":". Values prefixed with "is:" are treated the
// same as values with no prefix.
// Ancestry subtrees must be in one of the following formats:
//     - "projects/<project-id>", e.g. "projects/tokyo-rain-123"
//     - "folders/<folder-id>", e.g. "folders/1234"
//     - "organizations/<organization-id>", e.g. "organizations/1234"
// The `supports_under` field of the associated `Constraint`  defines whether
// ancestry prefixes can be used. You can set `allowed_values` and
// `denied_values` in the same `Policy` if `all_values` is
// `ALL_VALUES_UNSPECIFIED`. `ALLOW` or `DENY` are used to allow or deny all
// values. If `all_values` is set to either `ALLOW` or `DENY`,
// `allowed_values` and `denied_values` must be unset.
type Policy_ListPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of values allowed  at this resource. Can only be set if `all_values`
	// is set to `ALL_VALUES_UNSPECIFIED`.
	AllowedValues []string `protobuf:"bytes,1,rep,name=allowed_values,json=allowedValues,proto3" json:"allowed_values,omitempty"`
	// List of values denied at this resource. Can only be set if `all_values`
	// is set to `ALL_VALUES_UNSPECIFIED`.
	DeniedValues []string `protobuf:"bytes,2,rep,name=denied_values,json=deniedValues,proto3" json:"denied_values,omitempty"`
	// The policy all_values state.
	AllValues Policy_ListPolicy_AllValues `protobuf:"varint,3,opt,name=all_values,json=allValues,proto3,enum=google.cloud.orgpolicy.v1.Policy_ListPolicy_AllValues" json:"all_values,omitempty"`
	// Optional. The Google Cloud Console will try to default to a configuration
	// that matches the value specified in this `Policy`. If `suggested_value`
	// is not set, it will inherit the value specified higher in the hierarchy,
	// unless `inherit_from_parent` is `false`.
	SuggestedValue string `protobuf:"bytes,4,opt,name=suggested_value,json=suggestedValue,proto3" json:"suggested_value,omitempty"`
	// Determines the inheritance behavior for this `Policy`.
	//
	// By default, a `ListPolicy` set at a resource supercedes any `Policy` set
	// anywhere up the resource hierarchy. However, if `inherit_from_parent` is
	// set to `true`, then the values from the effective `Policy` of the parent
	// resource are inherited, meaning the values set in this `Policy` are
	// added to the values inherited up the hierarchy.
	//
	// Setting `Policy` hierarchies that inherit both allowed values and denied
	// values isn't recommended in most circumstances to keep the configuration
	// simple and understandable. However, it is possible to set a `Policy` with
	// `allowed_values` set that inherits a `Policy` with `denied_values` set.
	// In this case, the values that are allowed must be in `allowed_values` and
	// not present in `denied_values`.
	//
	// For example, suppose you have a `Constraint`
	// `constraints/serviceuser.services`, which has a `constraint_type` of
	// `list_constraint`, and with `constraint_default` set to `ALLOW`.
	// Suppose that at the Organization level, a `Policy` is applied that
	// restricts the allowed API activations to {`E1`, `E2`}. Then, if a
	// `Policy` is applied to a project below the Organization that has
	// `inherit_from_parent` set to `false` and field all_values set to DENY,
	// then an attempt to activate any API will be denied.
	//
	// The following examples demonstrate different possible layerings for
	// `projects/bar` parented by `organizations/foo`:
	//
	// Example 1 (no inherited values):
	//   `organizations/foo` has a `Policy` with values:
	//     {allowed_values: "E1" allowed_values:"E2"}
	//   `projects/bar` has `inherit_from_parent` `false` and values:
	//     {allowed_values: "E3" allowed_values: "E4"}
	// The accepted values at `organizations/foo` are `E1`, `E2`.
	// The accepted values at `projects/bar` are `E3`, and `E4`.
	//
	// Example 2 (inherited values):
	//   `organizations/foo` has a `Policy` with values:
	//     {allowed_values: "E1" allowed_values:"E2"}
	//   `projects/bar` has a `Policy` with values:
	//     {value: "E3" value: "E4" inherit_from_parent: true}
	// The accepted values at `organizations/foo` are `E1`, `E2`.
	// The accepted values at `projects/bar` are `E1`, `E2`, `E3`, and `E4`.
	//
	// Example 3 (inheriting both allowed and denied values):
	//   `organizations/foo` has a `Policy` with values:
	//     {allowed_values: "E1" allowed_values: "E2"}
	//   `projects/bar` has a `Policy` with:
	//     {denied_values: "E1"}
	// The accepted values at `organizations/foo` are `E1`, `E2`.
	// The value accepted at `projects/bar` is `E2`.
	//
	// Example 4 (RestoreDefault):
	//   `organizations/foo` has a `Policy` with values:
	//     {allowed_values: "E1" allowed_values:"E2"}
	//   `projects/bar` has a `Policy` with values:
	//     {RestoreDefault: {}}
	// The accepted values at `organizations/foo` are `E1`, `E2`.
	// The accepted values at `projects/bar` are either all or none depending on
	// the value of `constraint_default` (if `ALLOW`, all; if
	// `DENY`, none).
	//
	// Example 5 (no policy inherits parent policy):
	//   `organizations/foo` has no `Policy` set.
	//   `projects/bar` has no `Policy` set.
	// The accepted values at both levels are either all or none depending on
	// the value of `constraint_default` (if `ALLOW`, all; if
	// `DENY`, none).
	//
	// Example 6 (ListConstraint allowing all):
	//   `organizations/foo` has a `Policy` with values:
	//     {allowed_values: "E1" allowed_values: "E2"}
	//   `projects/bar` has a `Policy` with:
	//     {all: ALLOW}
	// The accepted values at `organizations/foo` are `E1`, E2`.
	// Any value is accepted at `projects/bar`.
	//
	// Example 7 (ListConstraint allowing none):
	//   `organizations/foo` has a `Policy` with values:
	//     {allowed_values: "E1" allowed_values: "E2"}
	//   `projects/bar` has a `Policy` with:
	//     {all: DENY}
	// The accepted values at `organizations/foo` are `E1`, E2`.
	// No value is accepted at `projects/bar`.
	//
	// Example 10 (allowed and denied subtrees of Resource Manager hierarchy):
	// Given the following resource hierarchy
	//   O1->{F1, F2}; F1->{P1}; F2->{P2, P3},
	//   `organizations/foo` has a `Policy` with values:
	//     {allowed_values: "under:organizations/O1"}
	//   `projects/bar` has a `Policy` with:
	//     {allowed_values: "under:projects/P3"}
	//     {denied_values: "under:folders/F2"}
	// The accepted values at `organizations/foo` are `organizations/O1`,
	//   `folders/F1`, `folders/F2`, `projects/P1`, `projects/P2`,
	//   `projects/P3`.
	// The accepted values at `projects/bar` are `organizations/O1`,
	//   `folders/F1`, `projects/P1`.
	InheritFromParent bool `protobuf:"varint,5,opt,name=inherit_from_parent,json=inheritFromParent,proto3" json:"inherit_from_parent,omitempty"`
}

func (x *Policy_ListPolicy) Reset() {
	*x = Policy_ListPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_orgpolicy_v1_orgpolicy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policy_ListPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy_ListPolicy) ProtoMessage() {}

func (x *Policy_ListPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_orgpolicy_v1_orgpolicy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy_ListPolicy.ProtoReflect.Descriptor instead.
func (*Policy_ListPolicy) Descriptor() ([]byte, []int) {
	return file_google_cloud_orgpolicy_v1_orgpolicy_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Policy_ListPolicy) GetAllowedValues() []string {
	if x != nil {
		return x.AllowedValues
	}
	return nil
}

func (x *Policy_ListPolicy) GetDeniedValues() []string {
	if x != nil {
		return x.DeniedValues
	}
	return nil
}

func (x *Policy_ListPolicy) GetAllValues() Policy_ListPolicy_AllValues {
	if x != nil {
		return x.AllValues
	}
	return Policy_ListPolicy_ALL_VALUES_UNSPECIFIED
}

func (x *Policy_ListPolicy) GetSuggestedValue() string {
	if x != nil {
		return x.SuggestedValue
	}
	return ""
}

func (x *Policy_ListPolicy) GetInheritFromParent() bool {
	if x != nil {
		return x.InheritFromParent
	}
	return false
}

// Used in `policy_type` to specify how `boolean_policy` will behave at this
// resource.
type Policy_BooleanPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If `true`, then the `Policy` is enforced. If `false`, then any
	// configuration is acceptable.
	//
	// Suppose you have a `Constraint`
	// `constraints/compute.disableSerialPortAccess` with `constraint_default`
	// set to `ALLOW`. A `Policy` for that `Constraint` exhibits the following
	// behavior:
	//   - If the `Policy` at this resource has enforced set to `false`, serial
	//     port connection attempts will be allowed.
	//   - If the `Policy` at this resource has enforced set to `true`, serial
	//     port connection attempts will be refused.
	//   - If the `Policy` at this resource is `RestoreDefault`, serial port
	//     connection attempts will be allowed.
	//   - If no `Policy` is set at this resource or anywhere higher in the
	//     resource hierarchy, serial port connection attempts will be allowed.
	//   - If no `Policy` is set at this resource, but one exists higher in the
	//     resource hierarchy, the behavior is as if the`Policy` were set at
	//     this resource.
	//
	// The following examples demonstrate the different possible layerings:
	//
	// Example 1 (nearest `Constraint` wins):
	//   `organizations/foo` has a `Policy` with:
	//     {enforced: false}
	//   `projects/bar` has no `Policy` set.
	// The constraint at `projects/bar` and `organizations/foo` will not be
	// enforced.
	//
	// Example 2 (enforcement gets replaced):
	//   `organizations/foo` has a `Policy` with:
	//     {enforced: false}
	//   `projects/bar` has a `Policy` with:
	//     {enforced: true}
	// The constraint at `organizations/foo` is not enforced.
	// The constraint at `projects/bar` is enforced.
	//
	// Example 3 (RestoreDefault):
	//   `organizations/foo` has a `Policy` with:
	//     {enforced: true}
	//   `projects/bar` has a `Policy` with:
	//     {RestoreDefault: {}}
	// The constraint at `organizations/foo` is enforced.
	// The constraint at `projects/bar` is not enforced, because
	// `constraint_default` for the `Constraint` is `ALLOW`.
	Enforced bool `protobuf:"varint,1,opt,name=enforced,proto3" json:"enforced,omitempty"`
}

func (x *Policy_BooleanPolicy) Reset() {
	*x = Policy_BooleanPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_orgpolicy_v1_orgpolicy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policy_BooleanPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy_BooleanPolicy) ProtoMessage() {}

func (x *Policy_BooleanPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_orgpolicy_v1_orgpolicy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy_BooleanPolicy.ProtoReflect.Descriptor instead.
func (*Policy_BooleanPolicy) Descriptor() ([]byte, []int) {
	return file_google_cloud_orgpolicy_v1_orgpolicy_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Policy_BooleanPolicy) GetEnforced() bool {
	if x != nil {
		return x.Enforced
	}
	return false
}

// Ignores policies set above this resource and restores the
// `constraint_default` enforcement behavior of the specific `Constraint` at
// this resource.
//
// Suppose that `constraint_default` is set to `ALLOW` for the
// `Constraint` `constraints/serviceuser.services`. Suppose that organization
// foo.com sets a `Policy` at their Organization resource node that restricts
// the allowed service activations to deny all service activations. They
// could then set a `Policy` with the `policy_type` `restore_default` on
// several experimental projects, restoring the `constraint_default`
// enforcement of the `Constraint` for only those projects, allowing those
// projects to have all services activated.
type Policy_RestoreDefault struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Policy_RestoreDefault) Reset() {
	*x = Policy_RestoreDefault{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_orgpolicy_v1_orgpolicy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policy_RestoreDefault) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy_RestoreDefault) ProtoMessage() {}

func (x *Policy_RestoreDefault) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_orgpolicy_v1_orgpolicy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy_RestoreDefault.ProtoReflect.Descriptor instead.
func (*Policy_RestoreDefault) Descriptor() ([]byte, []int) {
	return file_google_cloud_orgpolicy_v1_orgpolicy_proto_rawDescGZIP(), []int{0, 2}
}

var File_google_cloud_orgpolicy_v1_orgpolicy_proto protoreflect.FileDescriptor

var file_google_cloud_orgpolicy_v1_orgpolicy_proto_rawDesc = []byte{
	0x0a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f,
	0x72, 0x67, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x72, 0x67, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb2, 0x06, 0x0a, 0x06, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e,
	0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x74, 0x61, 0x67, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x65, 0x74, 0x61, 0x67, 0x12, 0x3b, 0x0a, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x4f, 0x0a, 0x0b, 0x6c, 0x69, 0x73, 0x74,
	0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x72, 0x67,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x48, 0x00, 0x52, 0x0a, 0x6c,
	0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x58, 0x0a, 0x0e, 0x62, 0x6f, 0x6f,
	0x6c, 0x65, 0x61, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6f, 0x72, 0x67, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x48, 0x00, 0x52, 0x0d, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x12, 0x5b, 0x0a, 0x0f, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x72, 0x67, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x48, 0x00,
	0x52, 0x0e, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x1a, 0xc6, 0x02, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12,
	0x25, 0x0a, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x6e, 0x69, 0x65, 0x64,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x64,
	0x65, 0x6e, 0x69, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x55, 0x0a, 0x0a, 0x61,
	0x6c, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f,
	0x72, 0x67, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x41, 0x6c,
	0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x09, 0x61, 0x6c, 0x6c, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x75, 0x67,
	0x67, 0x65, 0x73, 0x74, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x69,
	0x6e, 0x68, 0x65, 0x72, 0x69, 0x74, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x69, 0x6e, 0x68, 0x65, 0x72, 0x69,
	0x74, 0x46, 0x72, 0x6f, 0x6d, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x22, 0x3c, 0x0a, 0x09, 0x41,
	0x6c, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x16, 0x41, 0x4c, 0x4c, 0x5f,
	0x56, 0x41, 0x4c, 0x55, 0x45, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x10, 0x01, 0x12,
	0x08, 0x0a, 0x04, 0x44, 0x45, 0x4e, 0x59, 0x10, 0x02, 0x1a, 0x2b, 0x0a, 0x0d, 0x42, 0x6f, 0x6f,
	0x6c, 0x65, 0x61, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e,
	0x66, 0x6f, 0x72, 0x63, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x65, 0x6e,
	0x66, 0x6f, 0x72, 0x63, 0x65, 0x64, 0x1a, 0x10, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0xcc, 0x01, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x72, 0x67,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x4f, 0x72, 0x67, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x72, 0x67, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x6f, 0x72, 0x67, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0xaa,
	0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4f,
	0x72, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x19, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x4f, 0x72, 0x67, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x1c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x4f, 0x72, 0x67, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_orgpolicy_v1_orgpolicy_proto_rawDescOnce sync.Once
	file_google_cloud_orgpolicy_v1_orgpolicy_proto_rawDescData = file_google_cloud_orgpolicy_v1_orgpolicy_proto_rawDesc
)

func file_google_cloud_orgpolicy_v1_orgpolicy_proto_rawDescGZIP() []byte {
	file_google_cloud_orgpolicy_v1_orgpolicy_proto_rawDescOnce.Do(func() {
		file_google_cloud_orgpolicy_v1_orgpolicy_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_orgpolicy_v1_orgpolicy_proto_rawDescData)
	})
	return file_google_cloud_orgpolicy_v1_orgpolicy_proto_rawDescData
}

var file_google_cloud_orgpolicy_v1_orgpolicy_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_orgpolicy_v1_orgpolicy_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_orgpolicy_v1_orgpolicy_proto_goTypes = []interface{}{
	(Policy_ListPolicy_AllValues)(0), // 0: google.cloud.orgpolicy.v1.Policy.ListPolicy.AllValues
	(*Policy)(nil),                   // 1: google.cloud.orgpolicy.v1.Policy
	(*Policy_ListPolicy)(nil),        // 2: google.cloud.orgpolicy.v1.Policy.ListPolicy
	(*Policy_BooleanPolicy)(nil),     // 3: google.cloud.orgpolicy.v1.Policy.BooleanPolicy
	(*Policy_RestoreDefault)(nil),    // 4: google.cloud.orgpolicy.v1.Policy.RestoreDefault
	(*timestamppb.Timestamp)(nil),    // 5: google.protobuf.Timestamp
}
var file_google_cloud_orgpolicy_v1_orgpolicy_proto_depIdxs = []int32{
	5, // 0: google.cloud.orgpolicy.v1.Policy.update_time:type_name -> google.protobuf.Timestamp
	2, // 1: google.cloud.orgpolicy.v1.Policy.list_policy:type_name -> google.cloud.orgpolicy.v1.Policy.ListPolicy
	3, // 2: google.cloud.orgpolicy.v1.Policy.boolean_policy:type_name -> google.cloud.orgpolicy.v1.Policy.BooleanPolicy
	4, // 3: google.cloud.orgpolicy.v1.Policy.restore_default:type_name -> google.cloud.orgpolicy.v1.Policy.RestoreDefault
	0, // 4: google.cloud.orgpolicy.v1.Policy.ListPolicy.all_values:type_name -> google.cloud.orgpolicy.v1.Policy.ListPolicy.AllValues
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_cloud_orgpolicy_v1_orgpolicy_proto_init() }
func file_google_cloud_orgpolicy_v1_orgpolicy_proto_init() {
	if File_google_cloud_orgpolicy_v1_orgpolicy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_orgpolicy_v1_orgpolicy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policy); i {
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
		file_google_cloud_orgpolicy_v1_orgpolicy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policy_ListPolicy); i {
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
		file_google_cloud_orgpolicy_v1_orgpolicy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policy_BooleanPolicy); i {
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
		file_google_cloud_orgpolicy_v1_orgpolicy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policy_RestoreDefault); i {
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
	file_google_cloud_orgpolicy_v1_orgpolicy_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Policy_ListPolicy_)(nil),
		(*Policy_BooleanPolicy_)(nil),
		(*Policy_RestoreDefault_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_orgpolicy_v1_orgpolicy_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_orgpolicy_v1_orgpolicy_proto_goTypes,
		DependencyIndexes: file_google_cloud_orgpolicy_v1_orgpolicy_proto_depIdxs,
		EnumInfos:         file_google_cloud_orgpolicy_v1_orgpolicy_proto_enumTypes,
		MessageInfos:      file_google_cloud_orgpolicy_v1_orgpolicy_proto_msgTypes,
	}.Build()
	File_google_cloud_orgpolicy_v1_orgpolicy_proto = out.File
	file_google_cloud_orgpolicy_v1_orgpolicy_proto_rawDesc = nil
	file_google_cloud_orgpolicy_v1_orgpolicy_proto_goTypes = nil
	file_google_cloud_orgpolicy_v1_orgpolicy_proto_depIdxs = nil
}
