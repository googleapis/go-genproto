// Copyright 2021 Google LLC
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
// 	protoc        v3.12.2
// source: google/cloud/osconfig/v1alpha/instance_os_policies_compliance.proto

package osconfig

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This API resource represents the OS policies compliance data for a Compute
// Engine virtual machine (VM) instance at a given point in time.
//
// A Compute Engine VM can have multiple OS policy assignments, and each
// assignment can have multiple OS policies. As a result, multiple OS policies
// could be applied to a single VM.
//
// You can use this API resource to determine both the compliance state of your
// VM as well as the compliance state of an individual OS policy.
//
// For more information, see [View
// compliance](https://cloud.google.com/compute/docs/os-configuration-management/view-compliance).
//
// Deprecated: Do not use.
type InstanceOSPoliciesCompliance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The `InstanceOSPoliciesCompliance` API resource name.
	//
	// Format:
	// `projects/{project_number}/locations/{location}/instanceOSPoliciesCompliances/{instance_id}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The Compute Engine VM instance name.
	Instance string `protobuf:"bytes,2,opt,name=instance,proto3" json:"instance,omitempty"`
	// Output only. Compliance state of the VM.
	State OSPolicyComplianceState `protobuf:"varint,3,opt,name=state,proto3,enum=google.cloud.osconfig.v1alpha.OSPolicyComplianceState" json:"state,omitempty"`
	// Output only. Detailed compliance state of the VM.
	// This field is populated only when compliance state is `UNKNOWN`.
	//
	// It may contain one of the following values:
	//
	// * `no-compliance-data`: Compliance data is not available for this VM.
	// * `no-agent-detected`: OS Config agent is not detected for this VM.
	// * `config-not-supported-by-agent`: The version of the OS Config agent
	// running on this VM does not support configuration management.
	// * `inactive`: VM is not running.
	// * `internal-service-errors`: There were internal service errors encountered
	// while enforcing compliance.
	// * `agent-errors`: OS config agent encountered errors while enforcing
	// compliance.
	DetailedState string `protobuf:"bytes,4,opt,name=detailed_state,json=detailedState,proto3" json:"detailed_state,omitempty"`
	// Output only. The reason for the `detailed_state` of the VM (if any).
	DetailedStateReason string `protobuf:"bytes,5,opt,name=detailed_state_reason,json=detailedStateReason,proto3" json:"detailed_state_reason,omitempty"`
	// Output only. Compliance data for each `OSPolicy` that is applied to the VM.
	OsPolicyCompliances []*InstanceOSPoliciesCompliance_OSPolicyCompliance `protobuf:"bytes,6,rep,name=os_policy_compliances,json=osPolicyCompliances,proto3" json:"os_policy_compliances,omitempty"`
	// Output only. Timestamp of the last compliance check for the VM.
	LastComplianceCheckTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=last_compliance_check_time,json=lastComplianceCheckTime,proto3" json:"last_compliance_check_time,omitempty"`
	// Output only. Unique identifier for the last compliance run.
	// This id will be logged by the OS config agent during a compliance run and
	// can be used for debugging and tracing purpose.
	LastComplianceRunId string `protobuf:"bytes,8,opt,name=last_compliance_run_id,json=lastComplianceRunId,proto3" json:"last_compliance_run_id,omitempty"`
}

func (x *InstanceOSPoliciesCompliance) Reset() {
	*x = InstanceOSPoliciesCompliance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstanceOSPoliciesCompliance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstanceOSPoliciesCompliance) ProtoMessage() {}

func (x *InstanceOSPoliciesCompliance) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstanceOSPoliciesCompliance.ProtoReflect.Descriptor instead.
func (*InstanceOSPoliciesCompliance) Descriptor() ([]byte, []int) {
	return file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_rawDescGZIP(), []int{0}
}

func (x *InstanceOSPoliciesCompliance) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InstanceOSPoliciesCompliance) GetInstance() string {
	if x != nil {
		return x.Instance
	}
	return ""
}

func (x *InstanceOSPoliciesCompliance) GetState() OSPolicyComplianceState {
	if x != nil {
		return x.State
	}
	return OSPolicyComplianceState_OS_POLICY_COMPLIANCE_STATE_UNSPECIFIED
}

func (x *InstanceOSPoliciesCompliance) GetDetailedState() string {
	if x != nil {
		return x.DetailedState
	}
	return ""
}

func (x *InstanceOSPoliciesCompliance) GetDetailedStateReason() string {
	if x != nil {
		return x.DetailedStateReason
	}
	return ""
}

func (x *InstanceOSPoliciesCompliance) GetOsPolicyCompliances() []*InstanceOSPoliciesCompliance_OSPolicyCompliance {
	if x != nil {
		return x.OsPolicyCompliances
	}
	return nil
}

func (x *InstanceOSPoliciesCompliance) GetLastComplianceCheckTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastComplianceCheckTime
	}
	return nil
}

func (x *InstanceOSPoliciesCompliance) GetLastComplianceRunId() string {
	if x != nil {
		return x.LastComplianceRunId
	}
	return ""
}

// A request message for getting OS policies compliance data for the given
// Compute Engine VM instance.
//
// Deprecated: Do not use.
type GetInstanceOSPoliciesComplianceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. API resource name for instance OS policies compliance resource.
	//
	// Format:
	// `projects/{project}/locations/{location}/instanceOSPoliciesCompliances/{instance}`
	//
	// For `{project}`, either Compute Engine project-number or project-id can be
	// provided.
	// For `{instance}`, either Compute Engine VM instance-id or instance-name can
	// be provided.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetInstanceOSPoliciesComplianceRequest) Reset() {
	*x = GetInstanceOSPoliciesComplianceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInstanceOSPoliciesComplianceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInstanceOSPoliciesComplianceRequest) ProtoMessage() {}

func (x *GetInstanceOSPoliciesComplianceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInstanceOSPoliciesComplianceRequest.ProtoReflect.Descriptor instead.
func (*GetInstanceOSPoliciesComplianceRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_rawDescGZIP(), []int{1}
}

func (x *GetInstanceOSPoliciesComplianceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// A request message for listing OS policies compliance data for all Compute
// Engine VMs in the given location.
//
// Deprecated: Do not use.
type ListInstanceOSPoliciesCompliancesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The parent resource name.
	//
	// Format: `projects/{project}/locations/{location}`
	//
	// For `{project}`, either Compute Engine project-number or project-id can be
	// provided.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The maximum number of results to return.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A pagination token returned from a previous call to
	// `ListInstanceOSPoliciesCompliances` that indicates where this listing
	// should continue from.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// If provided, this field specifies the criteria that must be met by a
	// `InstanceOSPoliciesCompliance` API resource to be included in the response.
	Filter string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListInstanceOSPoliciesCompliancesRequest) Reset() {
	*x = ListInstanceOSPoliciesCompliancesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInstanceOSPoliciesCompliancesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInstanceOSPoliciesCompliancesRequest) ProtoMessage() {}

func (x *ListInstanceOSPoliciesCompliancesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInstanceOSPoliciesCompliancesRequest.ProtoReflect.Descriptor instead.
func (*ListInstanceOSPoliciesCompliancesRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_rawDescGZIP(), []int{2}
}

func (x *ListInstanceOSPoliciesCompliancesRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListInstanceOSPoliciesCompliancesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListInstanceOSPoliciesCompliancesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListInstanceOSPoliciesCompliancesRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

// A response message for listing OS policies compliance data for all Compute
// Engine VMs in the given location.
//
// Deprecated: Do not use.
type ListInstanceOSPoliciesCompliancesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of instance OS policies compliance objects.
	InstanceOsPoliciesCompliances []*InstanceOSPoliciesCompliance `protobuf:"bytes,1,rep,name=instance_os_policies_compliances,json=instanceOsPoliciesCompliances,proto3" json:"instance_os_policies_compliances,omitempty"`
	// The pagination token to retrieve the next page of instance OS policies
	// compliance objects.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListInstanceOSPoliciesCompliancesResponse) Reset() {
	*x = ListInstanceOSPoliciesCompliancesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInstanceOSPoliciesCompliancesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInstanceOSPoliciesCompliancesResponse) ProtoMessage() {}

func (x *ListInstanceOSPoliciesCompliancesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInstanceOSPoliciesCompliancesResponse.ProtoReflect.Descriptor instead.
func (*ListInstanceOSPoliciesCompliancesResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_rawDescGZIP(), []int{3}
}

func (x *ListInstanceOSPoliciesCompliancesResponse) GetInstanceOsPoliciesCompliances() []*InstanceOSPoliciesCompliance {
	if x != nil {
		return x.InstanceOsPoliciesCompliances
	}
	return nil
}

func (x *ListInstanceOSPoliciesCompliancesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// Compliance data for an OS policy
//
// Deprecated: Do not use.
type InstanceOSPoliciesCompliance_OSPolicyCompliance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The OS policy id
	OsPolicyId string `protobuf:"bytes,1,opt,name=os_policy_id,json=osPolicyId,proto3" json:"os_policy_id,omitempty"`
	// Reference to the `OSPolicyAssignment` API resource that the `OSPolicy`
	// belongs to.
	//
	// Format:
	// `projects/{project_number}/locations/{location}/osPolicyAssignments/{os_policy_assignment_id@revision_id}`
	OsPolicyAssignment string `protobuf:"bytes,2,opt,name=os_policy_assignment,json=osPolicyAssignment,proto3" json:"os_policy_assignment,omitempty"`
	// Compliance state of the OS policy.
	State OSPolicyComplianceState `protobuf:"varint,4,opt,name=state,proto3,enum=google.cloud.osconfig.v1alpha.OSPolicyComplianceState" json:"state,omitempty"`
	// Compliance data for each `OSPolicyResource` that is applied to the
	// VM.
	OsPolicyResourceCompliances []*OSPolicyResourceCompliance `protobuf:"bytes,5,rep,name=os_policy_resource_compliances,json=osPolicyResourceCompliances,proto3" json:"os_policy_resource_compliances,omitempty"`
}

func (x *InstanceOSPoliciesCompliance_OSPolicyCompliance) Reset() {
	*x = InstanceOSPoliciesCompliance_OSPolicyCompliance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstanceOSPoliciesCompliance_OSPolicyCompliance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstanceOSPoliciesCompliance_OSPolicyCompliance) ProtoMessage() {}

func (x *InstanceOSPoliciesCompliance_OSPolicyCompliance) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstanceOSPoliciesCompliance_OSPolicyCompliance.ProtoReflect.Descriptor instead.
func (*InstanceOSPoliciesCompliance_OSPolicyCompliance) Descriptor() ([]byte, []int) {
	return file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_rawDescGZIP(), []int{0, 0}
}

func (x *InstanceOSPoliciesCompliance_OSPolicyCompliance) GetOsPolicyId() string {
	if x != nil {
		return x.OsPolicyId
	}
	return ""
}

func (x *InstanceOSPoliciesCompliance_OSPolicyCompliance) GetOsPolicyAssignment() string {
	if x != nil {
		return x.OsPolicyAssignment
	}
	return ""
}

func (x *InstanceOSPoliciesCompliance_OSPolicyCompliance) GetState() OSPolicyComplianceState {
	if x != nil {
		return x.State
	}
	return OSPolicyComplianceState_OS_POLICY_COMPLIANCE_STATE_UNSPECIFIED
}

func (x *InstanceOSPoliciesCompliance_OSPolicyCompliance) GetOsPolicyResourceCompliances() []*OSPolicyResourceCompliance {
	if x != nil {
		return x.OsPolicyResourceCompliances
	}
	return nil
}

var File_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto protoreflect.FileDescriptor

var file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_rawDesc = []byte{
	0x0a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f,
	0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6f, 0x73, 0x5f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f,
	0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x08, 0x0a, 0x1c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x4f, 0x53, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x69, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f,
	0x0a, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x51, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x36,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x73,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4f,
	0x53, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x2a, 0x0a, 0x0e, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x0d, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x37,
	0x0a, 0x15, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x13, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x87, 0x01, 0x0a, 0x15, 0x6f, 0x73, 0x5f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x4f, 0x53, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x4f, 0x53, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x13, 0x6f, 0x73,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65,
	0x73, 0x12, 0x5c, 0x0a, 0x1a, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69,
	0x61, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x17, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x38, 0x0a, 0x16, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e,
	0x63, 0x65, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x13, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69,
	0x61, 0x6e, 0x63, 0x65, 0x52, 0x75, 0x6e, 0x49, 0x64, 0x1a, 0xeb, 0x02, 0x0a, 0x12, 0x4f, 0x53,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x20, 0x0a, 0x0c, 0x6f, 0x73, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x49, 0x64, 0x12, 0x61, 0x0a, 0x14, 0x6f, 0x73, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f,
	0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x2f, 0xfa, 0x41, 0x2c, 0x0a, 0x2a, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4f,
	0x53, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x12, 0x6f, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x41, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x4c, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4f, 0x53, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x7e, 0x0a, 0x1e, 0x6f, 0x73, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x73, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4f, 0x53, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x1b, 0x6f, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0x3a, 0x02, 0x18, 0x01, 0x3a, 0x8e, 0x01, 0x18, 0x01, 0xea, 0x41, 0x88, 0x01,
	0x0a, 0x34, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x4f, 0x53, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x50, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4f, 0x53, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69,
	0x65, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x7d, 0x22, 0x7e, 0x0a, 0x26, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4f, 0x53, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65,
	0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x50, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x3c, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x36, 0x0a, 0x34, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4f, 0x53, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x65, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x02, 0x18, 0x01, 0x22, 0xc5, 0x01, 0x0a, 0x28, 0x4c, 0x69, 0x73,
	0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4f, 0x53, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x69, 0x65, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x3a, 0x02, 0x18, 0x01,
	0x22, 0xde, 0x01, 0x0a, 0x29, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x4f, 0x53, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x69, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x84,
	0x01, 0x0a, 0x20, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6f, 0x73, 0x5f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x4f, 0x53, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x1d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x4f, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x3a, 0x02, 0x18,
	0x01, 0x42, 0xf2, 0x01, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x21, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x4f, 0x53, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x69, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x73, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x6f, 0x73, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0xaa, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x4f, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x56, 0x31, 0x41, 0x6c,
	0x70, 0x68, 0x61, 0xca, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x4f, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5c, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0xea, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x4f, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a, 0x3a, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_rawDescOnce sync.Once
	file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_rawDescData = file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_rawDesc
)

func file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_rawDescGZIP() []byte {
	file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_rawDescOnce.Do(func() {
		file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_rawDescData)
	})
	return file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_rawDescData
}

var file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_goTypes = []interface{}{
	(*InstanceOSPoliciesCompliance)(nil),                    // 0: google.cloud.osconfig.v1alpha.InstanceOSPoliciesCompliance
	(*GetInstanceOSPoliciesComplianceRequest)(nil),          // 1: google.cloud.osconfig.v1alpha.GetInstanceOSPoliciesComplianceRequest
	(*ListInstanceOSPoliciesCompliancesRequest)(nil),        // 2: google.cloud.osconfig.v1alpha.ListInstanceOSPoliciesCompliancesRequest
	(*ListInstanceOSPoliciesCompliancesResponse)(nil),       // 3: google.cloud.osconfig.v1alpha.ListInstanceOSPoliciesCompliancesResponse
	(*InstanceOSPoliciesCompliance_OSPolicyCompliance)(nil), // 4: google.cloud.osconfig.v1alpha.InstanceOSPoliciesCompliance.OSPolicyCompliance
	(OSPolicyComplianceState)(0),                            // 5: google.cloud.osconfig.v1alpha.OSPolicyComplianceState
	(*timestamppb.Timestamp)(nil),                           // 6: google.protobuf.Timestamp
	(*OSPolicyResourceCompliance)(nil),                      // 7: google.cloud.osconfig.v1alpha.OSPolicyResourceCompliance
}
var file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_depIdxs = []int32{
	5, // 0: google.cloud.osconfig.v1alpha.InstanceOSPoliciesCompliance.state:type_name -> google.cloud.osconfig.v1alpha.OSPolicyComplianceState
	4, // 1: google.cloud.osconfig.v1alpha.InstanceOSPoliciesCompliance.os_policy_compliances:type_name -> google.cloud.osconfig.v1alpha.InstanceOSPoliciesCompliance.OSPolicyCompliance
	6, // 2: google.cloud.osconfig.v1alpha.InstanceOSPoliciesCompliance.last_compliance_check_time:type_name -> google.protobuf.Timestamp
	0, // 3: google.cloud.osconfig.v1alpha.ListInstanceOSPoliciesCompliancesResponse.instance_os_policies_compliances:type_name -> google.cloud.osconfig.v1alpha.InstanceOSPoliciesCompliance
	5, // 4: google.cloud.osconfig.v1alpha.InstanceOSPoliciesCompliance.OSPolicyCompliance.state:type_name -> google.cloud.osconfig.v1alpha.OSPolicyComplianceState
	7, // 5: google.cloud.osconfig.v1alpha.InstanceOSPoliciesCompliance.OSPolicyCompliance.os_policy_resource_compliances:type_name -> google.cloud.osconfig.v1alpha.OSPolicyResourceCompliance
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_init() }
func file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_init() {
	if File_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto != nil {
		return
	}
	file_google_cloud_osconfig_v1alpha_config_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstanceOSPoliciesCompliance); i {
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
		file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInstanceOSPoliciesComplianceRequest); i {
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
		file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInstanceOSPoliciesCompliancesRequest); i {
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
		file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInstanceOSPoliciesCompliancesResponse); i {
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
		file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstanceOSPoliciesCompliance_OSPolicyCompliance); i {
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
			RawDescriptor: file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_goTypes,
		DependencyIndexes: file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_depIdxs,
		MessageInfos:      file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_msgTypes,
	}.Build()
	File_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto = out.File
	file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_rawDesc = nil
	file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_goTypes = nil
	file_google_cloud_osconfig_v1alpha_instance_os_policies_compliance_proto_depIdxs = nil
}
