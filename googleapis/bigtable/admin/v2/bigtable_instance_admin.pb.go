// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/bigtable/admin/v2/bigtable_instance_admin.proto
// DO NOT EDIT!

/*
Package google_bigtable_admin_v2 is a generated protocol buffer package.

It is generated from these files:
	google.golang.org/genproto/googleapis/bigtable/admin/v2/bigtable_instance_admin.proto
	google.golang.org/genproto/googleapis/bigtable/admin/v2/bigtable_table_admin.proto
	google.golang.org/genproto/googleapis/bigtable/admin/v2/common.proto
	google.golang.org/genproto/googleapis/bigtable/admin/v2/instance.proto
	google.golang.org/genproto/googleapis/bigtable/admin/v2/table.proto

It has these top-level messages:
	CreateInstanceRequest
	GetInstanceRequest
	ListInstancesRequest
	ListInstancesResponse
	DeleteInstanceRequest
	CreateClusterRequest
	GetClusterRequest
	ListClustersRequest
	ListClustersResponse
	DeleteClusterRequest
	CreateInstanceMetadata
	UpdateClusterMetadata
	CreateTableRequest
	DropRowRangeRequest
	ListTablesRequest
	ListTablesResponse
	GetTableRequest
	DeleteTableRequest
	ModifyColumnFamiliesRequest
	Instance
	Cluster
	Table
	ColumnFamily
	GcRule
*/
package google_bigtable_admin_v2 // import "google.golang.org/genproto/googleapis/bigtable/admin/v2"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"
import google_longrunning "google.golang.org/genproto/googleapis/longrunning"
import google_protobuf3 "github.com/golang/protobuf/ptypes/empty"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

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

// Request message for BigtableInstanceAdmin.CreateInstance.
type CreateInstanceRequest struct {
	// The unique name of the project in which to create the new instance.
	// Values are of the form projects/<project>
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	// The id to be used when referring to the new instance within its project,
	// e.g. just the "myinstance" section of the full name
	// "projects/myproject/instances/myinstance"
	InstanceId string `protobuf:"bytes,2,opt,name=instance_id,json=instanceId" json:"instance_id,omitempty"`
	// The instance to create.
	// Fields marked "@OutputOnly" must be left blank.
	Instance *Instance `protobuf:"bytes,3,opt,name=instance" json:"instance,omitempty"`
	// The clusters to be created within the instance, mapped by desired
	// cluster ID (e.g. just the "mycluster" part of the full name
	// "projects/myproject/instances/myinstance/clusters/mycluster").
	// Fields marked "@OutputOnly" must be left blank.
	// Currently exactly one cluster must be specified.
	Clusters map[string]*Cluster `protobuf:"bytes,4,rep,name=clusters" json:"clusters,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *CreateInstanceRequest) Reset()                    { *m = CreateInstanceRequest{} }
func (m *CreateInstanceRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateInstanceRequest) ProtoMessage()               {}
func (*CreateInstanceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateInstanceRequest) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

func (m *CreateInstanceRequest) GetClusters() map[string]*Cluster {
	if m != nil {
		return m.Clusters
	}
	return nil
}

// Request message for BigtableInstanceAdmin.GetInstance.
type GetInstanceRequest struct {
	// The unique name of the requested instance. Values are of the form
	// projects/<project>/instances/<instance>
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetInstanceRequest) Reset()                    { *m = GetInstanceRequest{} }
func (m *GetInstanceRequest) String() string            { return proto.CompactTextString(m) }
func (*GetInstanceRequest) ProtoMessage()               {}
func (*GetInstanceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// Request message for BigtableInstanceAdmin.ListInstances.
type ListInstancesRequest struct {
	// The unique name of the project for which a list of instances is requested.
	// Values are of the form projects/<project>
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	// The value of `next_page_token` returned by a previous call.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListInstancesRequest) Reset()                    { *m = ListInstancesRequest{} }
func (m *ListInstancesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListInstancesRequest) ProtoMessage()               {}
func (*ListInstancesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// Response message for BigtableInstanceAdmin.ListInstances.
type ListInstancesResponse struct {
	// The list of requested instances.
	Instances []*Instance `protobuf:"bytes,1,rep,name=instances" json:"instances,omitempty"`
	// Locations from which Instance information could not be retrieved,
	// due to an outage or some other transient condition.
	// Instances whose Clusters are all in one of the failed locations
	// may be missing from 'instances', and Instances with at least one
	// Cluster in a failed location may only have partial information returned.
	FailedLocations []string `protobuf:"bytes,2,rep,name=failed_locations,json=failedLocations" json:"failed_locations,omitempty"`
	// Set if not all instances could be returned in a single response.
	// Pass this value to `page_token` in another request to get the next
	// page of results.
	NextPageToken string `protobuf:"bytes,3,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListInstancesResponse) Reset()                    { *m = ListInstancesResponse{} }
func (m *ListInstancesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListInstancesResponse) ProtoMessage()               {}
func (*ListInstancesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ListInstancesResponse) GetInstances() []*Instance {
	if m != nil {
		return m.Instances
	}
	return nil
}

// Request message for BigtableInstanceAdmin.DeleteInstance.
type DeleteInstanceRequest struct {
	// The unique name of the instance to be deleted.
	// Values are of the form projects/<project>/instances/<instance>
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DeleteInstanceRequest) Reset()                    { *m = DeleteInstanceRequest{} }
func (m *DeleteInstanceRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteInstanceRequest) ProtoMessage()               {}
func (*DeleteInstanceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// Request message for BigtableInstanceAdmin.CreateCluster.
type CreateClusterRequest struct {
	// The unique name of the instance in which to create the new cluster.
	// Values are of the form
	// projects/<project>/instances/<instance>/clusters/[a-z][-a-z0-9]*
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	// The id to be used when referring to the new cluster within its instance,
	// e.g. just the "mycluster" section of the full name
	// "projects/myproject/instances/myinstance/clusters/mycluster"
	ClusterId string `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId" json:"cluster_id,omitempty"`
	// The cluster to be created.
	// Fields marked "@OutputOnly" must be left blank.
	Cluster *Cluster `protobuf:"bytes,3,opt,name=cluster" json:"cluster,omitempty"`
}

func (m *CreateClusterRequest) Reset()                    { *m = CreateClusterRequest{} }
func (m *CreateClusterRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateClusterRequest) ProtoMessage()               {}
func (*CreateClusterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CreateClusterRequest) GetCluster() *Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

// Request message for BigtableInstanceAdmin.GetCluster.
type GetClusterRequest struct {
	// The unique name of the requested cluster. Values are of the form
	// projects/<project>/instances/<instance>/clusters/<cluster>
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetClusterRequest) Reset()                    { *m = GetClusterRequest{} }
func (m *GetClusterRequest) String() string            { return proto.CompactTextString(m) }
func (*GetClusterRequest) ProtoMessage()               {}
func (*GetClusterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

// Request message for BigtableInstanceAdmin.ListClusters.
type ListClustersRequest struct {
	// The unique name of the instance for which a list of clusters is requested.
	// Values are of the form projects/<project>/instances/<instance>
	// Use <instance> = '-' to list Clusters for all Instances in a project,
	// for example "projects/myproject/instances/-"
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	// The value of `next_page_token` returned by a previous call.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListClustersRequest) Reset()                    { *m = ListClustersRequest{} }
func (m *ListClustersRequest) String() string            { return proto.CompactTextString(m) }
func (*ListClustersRequest) ProtoMessage()               {}
func (*ListClustersRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// Response message for BigtableInstanceAdmin.ListClusters.
type ListClustersResponse struct {
	// The list of requested clusters.
	Clusters []*Cluster `protobuf:"bytes,1,rep,name=clusters" json:"clusters,omitempty"`
	// Locations from which Cluster information could not be retrieved,
	// due to an outage or some other transient condition.
	// Clusters from these locations may be missing from 'clusters',
	// or may only have partial information returned.
	FailedLocations []string `protobuf:"bytes,2,rep,name=failed_locations,json=failedLocations" json:"failed_locations,omitempty"`
	// Set if not all clusters could be returned in a single response.
	// Pass this value to `page_token` in another request to get the next
	// page of results.
	NextPageToken string `protobuf:"bytes,3,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListClustersResponse) Reset()                    { *m = ListClustersResponse{} }
func (m *ListClustersResponse) String() string            { return proto.CompactTextString(m) }
func (*ListClustersResponse) ProtoMessage()               {}
func (*ListClustersResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ListClustersResponse) GetClusters() []*Cluster {
	if m != nil {
		return m.Clusters
	}
	return nil
}

// Request message for BigtableInstanceAdmin.DeleteCluster.
type DeleteClusterRequest struct {
	// The unique name of the cluster to be deleted. Values are of the form
	// projects/<project>/instances/<instance>/clusters/<cluster>
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DeleteClusterRequest) Reset()                    { *m = DeleteClusterRequest{} }
func (m *DeleteClusterRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteClusterRequest) ProtoMessage()               {}
func (*DeleteClusterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

// The metadata for the Operation returned by CreateInstance.
type CreateInstanceMetadata struct {
	// The request that prompted the initiation of this CreateInstance operation.
	OriginalRequest *CreateInstanceRequest `protobuf:"bytes,1,opt,name=original_request,json=originalRequest" json:"original_request,omitempty"`
	// The time at which the original request was received.
	RequestTime *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=request_time,json=requestTime" json:"request_time,omitempty"`
	// The time at which the operation failed or was completed successfully.
	FinishTime *google_protobuf1.Timestamp `protobuf:"bytes,3,opt,name=finish_time,json=finishTime" json:"finish_time,omitempty"`
}

func (m *CreateInstanceMetadata) Reset()                    { *m = CreateInstanceMetadata{} }
func (m *CreateInstanceMetadata) String() string            { return proto.CompactTextString(m) }
func (*CreateInstanceMetadata) ProtoMessage()               {}
func (*CreateInstanceMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *CreateInstanceMetadata) GetOriginalRequest() *CreateInstanceRequest {
	if m != nil {
		return m.OriginalRequest
	}
	return nil
}

func (m *CreateInstanceMetadata) GetRequestTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.RequestTime
	}
	return nil
}

func (m *CreateInstanceMetadata) GetFinishTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.FinishTime
	}
	return nil
}

// The metadata for the Operation returned by UpdateCluster.
type UpdateClusterMetadata struct {
	// The request that prompted the initiation of this UpdateCluster operation.
	OriginalRequest *Cluster `protobuf:"bytes,1,opt,name=original_request,json=originalRequest" json:"original_request,omitempty"`
	// The time at which the original request was received.
	RequestTime *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=request_time,json=requestTime" json:"request_time,omitempty"`
	// The time at which the operation failed or was completed successfully.
	FinishTime *google_protobuf1.Timestamp `protobuf:"bytes,3,opt,name=finish_time,json=finishTime" json:"finish_time,omitempty"`
}

func (m *UpdateClusterMetadata) Reset()                    { *m = UpdateClusterMetadata{} }
func (m *UpdateClusterMetadata) String() string            { return proto.CompactTextString(m) }
func (*UpdateClusterMetadata) ProtoMessage()               {}
func (*UpdateClusterMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *UpdateClusterMetadata) GetOriginalRequest() *Cluster {
	if m != nil {
		return m.OriginalRequest
	}
	return nil
}

func (m *UpdateClusterMetadata) GetRequestTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.RequestTime
	}
	return nil
}

func (m *UpdateClusterMetadata) GetFinishTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.FinishTime
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateInstanceRequest)(nil), "google.bigtable.admin.v2.CreateInstanceRequest")
	proto.RegisterType((*GetInstanceRequest)(nil), "google.bigtable.admin.v2.GetInstanceRequest")
	proto.RegisterType((*ListInstancesRequest)(nil), "google.bigtable.admin.v2.ListInstancesRequest")
	proto.RegisterType((*ListInstancesResponse)(nil), "google.bigtable.admin.v2.ListInstancesResponse")
	proto.RegisterType((*DeleteInstanceRequest)(nil), "google.bigtable.admin.v2.DeleteInstanceRequest")
	proto.RegisterType((*CreateClusterRequest)(nil), "google.bigtable.admin.v2.CreateClusterRequest")
	proto.RegisterType((*GetClusterRequest)(nil), "google.bigtable.admin.v2.GetClusterRequest")
	proto.RegisterType((*ListClustersRequest)(nil), "google.bigtable.admin.v2.ListClustersRequest")
	proto.RegisterType((*ListClustersResponse)(nil), "google.bigtable.admin.v2.ListClustersResponse")
	proto.RegisterType((*DeleteClusterRequest)(nil), "google.bigtable.admin.v2.DeleteClusterRequest")
	proto.RegisterType((*CreateInstanceMetadata)(nil), "google.bigtable.admin.v2.CreateInstanceMetadata")
	proto.RegisterType((*UpdateClusterMetadata)(nil), "google.bigtable.admin.v2.UpdateClusterMetadata")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BigtableInstanceAdmin service

type BigtableInstanceAdminClient interface {
	// Create an instance within a project.
	CreateInstance(ctx context.Context, in *CreateInstanceRequest, opts ...grpc.CallOption) (*google_longrunning.Operation, error)
	// Gets information about an instance.
	GetInstance(ctx context.Context, in *GetInstanceRequest, opts ...grpc.CallOption) (*Instance, error)
	// Lists information about instances in a project.
	ListInstances(ctx context.Context, in *ListInstancesRequest, opts ...grpc.CallOption) (*ListInstancesResponse, error)
	// Updates an instance within a project.
	UpdateInstance(ctx context.Context, in *Instance, opts ...grpc.CallOption) (*Instance, error)
	// Delete an instance from a project.
	DeleteInstance(ctx context.Context, in *DeleteInstanceRequest, opts ...grpc.CallOption) (*google_protobuf3.Empty, error)
	// Creates a cluster within an instance.
	CreateCluster(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*google_longrunning.Operation, error)
	// Gets information about a cluster.
	GetCluster(ctx context.Context, in *GetClusterRequest, opts ...grpc.CallOption) (*Cluster, error)
	// Lists information about clusters in an instance.
	ListClusters(ctx context.Context, in *ListClustersRequest, opts ...grpc.CallOption) (*ListClustersResponse, error)
	// Updates a cluster within an instance.
	UpdateCluster(ctx context.Context, in *Cluster, opts ...grpc.CallOption) (*google_longrunning.Operation, error)
	// Deletes a cluster from an instance.
	DeleteCluster(ctx context.Context, in *DeleteClusterRequest, opts ...grpc.CallOption) (*google_protobuf3.Empty, error)
}

type bigtableInstanceAdminClient struct {
	cc *grpc.ClientConn
}

func NewBigtableInstanceAdminClient(cc *grpc.ClientConn) BigtableInstanceAdminClient {
	return &bigtableInstanceAdminClient{cc}
}

func (c *bigtableInstanceAdminClient) CreateInstance(ctx context.Context, in *CreateInstanceRequest, opts ...grpc.CallOption) (*google_longrunning.Operation, error) {
	out := new(google_longrunning.Operation)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.v2.BigtableInstanceAdmin/CreateInstance", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableInstanceAdminClient) GetInstance(ctx context.Context, in *GetInstanceRequest, opts ...grpc.CallOption) (*Instance, error) {
	out := new(Instance)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.v2.BigtableInstanceAdmin/GetInstance", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableInstanceAdminClient) ListInstances(ctx context.Context, in *ListInstancesRequest, opts ...grpc.CallOption) (*ListInstancesResponse, error) {
	out := new(ListInstancesResponse)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.v2.BigtableInstanceAdmin/ListInstances", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableInstanceAdminClient) UpdateInstance(ctx context.Context, in *Instance, opts ...grpc.CallOption) (*Instance, error) {
	out := new(Instance)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.v2.BigtableInstanceAdmin/UpdateInstance", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableInstanceAdminClient) DeleteInstance(ctx context.Context, in *DeleteInstanceRequest, opts ...grpc.CallOption) (*google_protobuf3.Empty, error) {
	out := new(google_protobuf3.Empty)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.v2.BigtableInstanceAdmin/DeleteInstance", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableInstanceAdminClient) CreateCluster(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*google_longrunning.Operation, error) {
	out := new(google_longrunning.Operation)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.v2.BigtableInstanceAdmin/CreateCluster", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableInstanceAdminClient) GetCluster(ctx context.Context, in *GetClusterRequest, opts ...grpc.CallOption) (*Cluster, error) {
	out := new(Cluster)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.v2.BigtableInstanceAdmin/GetCluster", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableInstanceAdminClient) ListClusters(ctx context.Context, in *ListClustersRequest, opts ...grpc.CallOption) (*ListClustersResponse, error) {
	out := new(ListClustersResponse)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.v2.BigtableInstanceAdmin/ListClusters", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableInstanceAdminClient) UpdateCluster(ctx context.Context, in *Cluster, opts ...grpc.CallOption) (*google_longrunning.Operation, error) {
	out := new(google_longrunning.Operation)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.v2.BigtableInstanceAdmin/UpdateCluster", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableInstanceAdminClient) DeleteCluster(ctx context.Context, in *DeleteClusterRequest, opts ...grpc.CallOption) (*google_protobuf3.Empty, error) {
	out := new(google_protobuf3.Empty)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.v2.BigtableInstanceAdmin/DeleteCluster", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BigtableInstanceAdmin service

type BigtableInstanceAdminServer interface {
	// Create an instance within a project.
	CreateInstance(context.Context, *CreateInstanceRequest) (*google_longrunning.Operation, error)
	// Gets information about an instance.
	GetInstance(context.Context, *GetInstanceRequest) (*Instance, error)
	// Lists information about instances in a project.
	ListInstances(context.Context, *ListInstancesRequest) (*ListInstancesResponse, error)
	// Updates an instance within a project.
	UpdateInstance(context.Context, *Instance) (*Instance, error)
	// Delete an instance from a project.
	DeleteInstance(context.Context, *DeleteInstanceRequest) (*google_protobuf3.Empty, error)
	// Creates a cluster within an instance.
	CreateCluster(context.Context, *CreateClusterRequest) (*google_longrunning.Operation, error)
	// Gets information about a cluster.
	GetCluster(context.Context, *GetClusterRequest) (*Cluster, error)
	// Lists information about clusters in an instance.
	ListClusters(context.Context, *ListClustersRequest) (*ListClustersResponse, error)
	// Updates a cluster within an instance.
	UpdateCluster(context.Context, *Cluster) (*google_longrunning.Operation, error)
	// Deletes a cluster from an instance.
	DeleteCluster(context.Context, *DeleteClusterRequest) (*google_protobuf3.Empty, error)
}

func RegisterBigtableInstanceAdminServer(s *grpc.Server, srv BigtableInstanceAdminServer) {
	s.RegisterService(&_BigtableInstanceAdmin_serviceDesc, srv)
}

func _BigtableInstanceAdmin_CreateInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableInstanceAdminServer).CreateInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.v2.BigtableInstanceAdmin/CreateInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableInstanceAdminServer).CreateInstance(ctx, req.(*CreateInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableInstanceAdmin_GetInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableInstanceAdminServer).GetInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.v2.BigtableInstanceAdmin/GetInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableInstanceAdminServer).GetInstance(ctx, req.(*GetInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableInstanceAdmin_ListInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableInstanceAdminServer).ListInstances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.v2.BigtableInstanceAdmin/ListInstances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableInstanceAdminServer).ListInstances(ctx, req.(*ListInstancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableInstanceAdmin_UpdateInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Instance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableInstanceAdminServer).UpdateInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.v2.BigtableInstanceAdmin/UpdateInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableInstanceAdminServer).UpdateInstance(ctx, req.(*Instance))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableInstanceAdmin_DeleteInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableInstanceAdminServer).DeleteInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.v2.BigtableInstanceAdmin/DeleteInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableInstanceAdminServer).DeleteInstance(ctx, req.(*DeleteInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableInstanceAdmin_CreateCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableInstanceAdminServer).CreateCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.v2.BigtableInstanceAdmin/CreateCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableInstanceAdminServer).CreateCluster(ctx, req.(*CreateClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableInstanceAdmin_GetCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableInstanceAdminServer).GetCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.v2.BigtableInstanceAdmin/GetCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableInstanceAdminServer).GetCluster(ctx, req.(*GetClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableInstanceAdmin_ListClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableInstanceAdminServer).ListClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.v2.BigtableInstanceAdmin/ListClusters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableInstanceAdminServer).ListClusters(ctx, req.(*ListClustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableInstanceAdmin_UpdateCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cluster)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableInstanceAdminServer).UpdateCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.v2.BigtableInstanceAdmin/UpdateCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableInstanceAdminServer).UpdateCluster(ctx, req.(*Cluster))
	}
	return interceptor(ctx, in, info, handler)
}

func _BigtableInstanceAdmin_DeleteCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigtableInstanceAdminServer).DeleteCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bigtable.admin.v2.BigtableInstanceAdmin/DeleteCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigtableInstanceAdminServer).DeleteCluster(ctx, req.(*DeleteClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BigtableInstanceAdmin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.bigtable.admin.v2.BigtableInstanceAdmin",
	HandlerType: (*BigtableInstanceAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateInstance",
			Handler:    _BigtableInstanceAdmin_CreateInstance_Handler,
		},
		{
			MethodName: "GetInstance",
			Handler:    _BigtableInstanceAdmin_GetInstance_Handler,
		},
		{
			MethodName: "ListInstances",
			Handler:    _BigtableInstanceAdmin_ListInstances_Handler,
		},
		{
			MethodName: "UpdateInstance",
			Handler:    _BigtableInstanceAdmin_UpdateInstance_Handler,
		},
		{
			MethodName: "DeleteInstance",
			Handler:    _BigtableInstanceAdmin_DeleteInstance_Handler,
		},
		{
			MethodName: "CreateCluster",
			Handler:    _BigtableInstanceAdmin_CreateCluster_Handler,
		},
		{
			MethodName: "GetCluster",
			Handler:    _BigtableInstanceAdmin_GetCluster_Handler,
		},
		{
			MethodName: "ListClusters",
			Handler:    _BigtableInstanceAdmin_ListClusters_Handler,
		},
		{
			MethodName: "UpdateCluster",
			Handler:    _BigtableInstanceAdmin_UpdateCluster_Handler,
		},
		{
			MethodName: "DeleteCluster",
			Handler:    _BigtableInstanceAdmin_DeleteCluster_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google.golang.org/genproto/googleapis/bigtable/admin/v2/bigtable_instance_admin.proto",
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/bigtable/admin/v2/bigtable_instance_admin.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 1006 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0x41, 0x6f, 0x1b, 0x45,
	0x14, 0xd6, 0xd8, 0xa5, 0xd4, 0xcf, 0x75, 0x12, 0x86, 0x38, 0xb2, 0x56, 0x54, 0xa4, 0x8b, 0x54,
	0x5c, 0x37, 0xec, 0x0a, 0x17, 0x29, 0x28, 0xc1, 0x08, 0x52, 0x4a, 0x15, 0x29, 0x15, 0x91, 0xd5,
	0x1e, 0xe0, 0x80, 0x35, 0xb6, 0x27, 0xdb, 0xa1, 0xeb, 0x99, 0x65, 0x77, 0x1c, 0x11, 0xa1, 0x5e,
	0x10, 0xe2, 0x10, 0x09, 0x0e, 0x70, 0x44, 0x9c, 0xb8, 0x70, 0xe0, 0x9f, 0x70, 0xe4, 0xc8, 0x09,
	0x89, 0x1f, 0xc0, 0x4f, 0x40, 0xbb, 0x33, 0xb3, 0xf6, 0x3a, 0x6b, 0xef, 0x06, 0x84, 0xd4, 0x8b,
	0xb5, 0xfb, 0x66, 0xde, 0x9b, 0x6f, 0xbe, 0xef, 0xdb, 0xf7, 0x0c, 0x8f, 0x3d, 0x21, 0x3c, 0x9f,
	0x3a, 0x9e, 0xf0, 0x09, 0xf7, 0x1c, 0x11, 0x7a, 0xae, 0x47, 0x79, 0x10, 0x0a, 0x29, 0x5c, 0xb5,
	0x44, 0x02, 0x16, 0xb9, 0x43, 0xe6, 0x49, 0x32, 0xf4, 0xa9, 0x4b, 0xc6, 0x13, 0xc6, 0xdd, 0xd3,
	0x6e, 0x1a, 0x19, 0x30, 0x1e, 0x49, 0xc2, 0x47, 0x74, 0x90, 0x2c, 0x39, 0x49, 0x2a, 0x6e, 0xe9,
	0xb2, 0x66, 0x97, 0xa3, 0x16, 0x4f, 0xbb, 0xd6, 0x61, 0xb9, 0x03, 0x49, 0xc0, 0xdc, 0x88, 0x86,
	0xa7, 0x6c, 0x44, 0x47, 0x82, 0x9f, 0x30, 0xcf, 0x25, 0x9c, 0x0b, 0x49, 0x24, 0x13, 0x3c, 0x52,
	0x87, 0x58, 0x1f, 0xfe, 0x5b, 0xec, 0x06, 0xb2, 0xae, 0x73, 0x50, 0xae, 0x8e, 0x2f, 0xb8, 0x17,
	0x4e, 0x39, 0x67, 0xdc, 0x73, 0x45, 0x40, 0xc3, 0x0c, 0x96, 0xbb, 0x1e, 0x93, 0x4f, 0xa6, 0x43,
	0x67, 0x24, 0x26, 0xae, 0xaa, 0xe3, 0x26, 0x0b, 0xc3, 0xe9, 0x89, 0x1b, 0xc8, 0xb3, 0x80, 0x46,
	0x2e, 0x9d, 0x04, 0xf2, 0x4c, 0xfd, 0xea, 0xa4, 0xfd, 0xe2, 0x24, 0xc9, 0x26, 0x34, 0x92, 0x64,
	0x12, 0xcc, 0x9e, 0x54, 0xb2, 0xfd, 0x5b, 0x05, 0x9a, 0xf7, 0x42, 0x4a, 0x24, 0x3d, 0xd4, 0xd7,
	0xe9, 0xd3, 0xcf, 0xa7, 0x34, 0x92, 0x78, 0x0b, 0xae, 0x06, 0x24, 0xa4, 0x5c, 0xb6, 0xd0, 0x36,
	0x6a, 0xd7, 0xfa, 0xfa, 0x0d, 0xbf, 0x0a, 0xf5, 0x54, 0x2c, 0x36, 0x6e, 0x55, 0x92, 0x45, 0x30,
	0xa1, 0xc3, 0x31, 0x7e, 0x17, 0xae, 0x99, 0xb7, 0x56, 0x75, 0x1b, 0xb5, 0xeb, 0x5d, 0xdb, 0x59,
	0x26, 0xa4, 0x93, 0x9e, 0x9a, 0xe6, 0xe0, 0x8f, 0xe1, 0xda, 0xc8, 0x9f, 0x46, 0x92, 0x86, 0x51,
	0xeb, 0xca, 0x76, 0xb5, 0x5d, 0xef, 0xf6, 0x96, 0xe7, 0xe7, 0x62, 0x77, 0xee, 0xe9, 0xfc, 0xfb,
	0x5c, 0x86, 0x67, 0xfd, 0xb4, 0x9c, 0xf5, 0x29, 0x34, 0x32, 0x4b, 0x78, 0x03, 0xaa, 0x4f, 0xe9,
	0x99, 0xbe, 0x61, 0xfc, 0x88, 0x77, 0xe1, 0x85, 0x53, 0xe2, 0x4f, 0x69, 0x72, 0xb1, 0x7a, 0xf7,
	0xe6, 0x8a, 0xa3, 0x55, 0xa5, 0xbe, 0xda, 0xbf, 0x57, 0x79, 0x1b, 0xd9, 0x6d, 0xc0, 0x0f, 0xa8,
	0x5c, 0x64, 0x12, 0xc3, 0x15, 0x4e, 0x26, 0x54, 0x9f, 0x92, 0x3c, 0xdb, 0x0f, 0x61, 0xf3, 0x88,
	0x45, 0xe9, 0xd6, 0xa8, 0x88, 0xf5, 0x1b, 0x00, 0x01, 0xf1, 0xe8, 0x40, 0x8a, 0xa7, 0x94, 0x6b,
	0xd2, 0x6b, 0x71, 0xe4, 0x51, 0x1c, 0xb0, 0x7f, 0x45, 0xd0, 0x5c, 0xa8, 0x17, 0x05, 0x82, 0x47,
	0x14, 0xbf, 0x07, 0x35, 0xc3, 0x6c, 0xd4, 0x42, 0x09, 0x9d, 0x65, 0xe4, 0x98, 0x25, 0xe1, 0xdb,
	0xb0, 0x71, 0x42, 0x98, 0x4f, 0xc7, 0x03, 0x5f, 0x8c, 0x94, 0x5d, 0x5b, 0x95, 0xed, 0x6a, 0xbb,
	0xd6, 0x5f, 0x57, 0xf1, 0x23, 0x13, 0xc6, 0xb7, 0x60, 0x9d, 0xd3, 0x2f, 0xe4, 0x60, 0x0e, 0x6a,
	0x35, 0x81, 0xda, 0x88, 0xc3, 0xc7, 0x29, 0xdc, 0x3b, 0xd0, 0xfc, 0x80, 0xfa, 0xf4, 0xa2, 0xe9,
	0xf2, 0xa8, 0x3a, 0x47, 0xb0, 0xa9, 0x64, 0x36, 0x8c, 0x17, 0x73, 0xa5, 0x15, 0x9f, 0x19, 0xb4,
	0xa6, 0x23, 0x87, 0x63, 0xbc, 0x0f, 0x2f, 0xea, 0x17, 0x6d, 0xcf, 0x12, 0x1a, 0x9b, 0x0c, 0xfb,
	0x75, 0x78, 0xe9, 0x01, 0x95, 0x0b, 0x40, 0xf2, 0x50, 0x1f, 0xc1, 0xcb, 0xb1, 0x20, 0xc6, 0x6e,
	0xff, 0x51, 0xdf, 0x5f, 0x90, 0xf2, 0xcb, 0xac, 0x9c, 0x96, 0xb7, 0x37, 0xf7, 0xb1, 0x28, 0x75,
	0x4b, 0xdc, 0x26, 0x4d, 0xf9, 0x3f, 0xb4, 0xed, 0xc0, 0xa6, 0xd2, 0xb6, 0x04, 0x49, 0x7f, 0x23,
	0xd8, 0xca, 0x7e, 0xc1, 0x0f, 0xa9, 0x24, 0x63, 0x22, 0x09, 0xfe, 0x04, 0x36, 0x44, 0xc8, 0x3c,
	0xc6, 0x89, 0x3f, 0x08, 0x55, 0x89, 0x24, 0xb5, 0xde, 0x75, 0x2f, 0xd9, 0x0d, 0xfa, 0xeb, 0xa6,
	0x90, 0x81, 0xd2, 0x83, 0xeb, 0xba, 0xe4, 0x20, 0xee, 0x87, 0xfa, 0x53, 0xb7, 0x4c, 0x5d, 0xd3,
	0x3d, 0x9d, 0x47, 0xa6, 0x59, 0xf6, 0xeb, 0x7a, 0x7f, 0x1c, 0xc1, 0xfb, 0x50, 0x3f, 0x61, 0x9c,
	0x45, 0x4f, 0x54, 0x76, 0xb5, 0x30, 0x1b, 0xd4, 0xf6, 0x38, 0x60, 0xff, 0x89, 0xa0, 0xf9, 0x38,
	0x18, 0xcf, 0xdc, 0x9c, 0xde, 0xf8, 0x68, 0xe9, 0x8d, 0x4b, 0x48, 0xfa, 0x3c, 0xdd, 0xb1, 0xfb,
	0x47, 0x1d, 0x9a, 0x07, 0x1a, 0xaa, 0x11, 0xe3, 0xfd, 0x18, 0x31, 0xfe, 0x0e, 0xc1, 0x5a, 0x56,
	0x24, 0x7c, 0x59, 0x39, 0xad, 0x1b, 0x26, 0x61, 0x6e, 0x94, 0x3a, 0x1f, 0x99, 0x51, 0x6a, 0xef,
	0x7c, 0xf5, 0xfb, 0x5f, 0x3f, 0x54, 0x6e, 0xd9, 0x37, 0xe3, 0x19, 0xfd, 0xa5, 0xfa, 0xbc, 0x7a,
	0x41, 0x28, 0x3e, 0xa3, 0x23, 0x19, 0xb9, 0x9d, 0x67, 0xe9, 0xdc, 0x8e, 0xf6, 0x50, 0x07, 0x9f,
	0x23, 0xa8, 0xcf, 0xb5, 0x6c, 0xbc, 0xb3, 0x1c, 0xcd, 0xc5, 0xce, 0x6e, 0x95, 0xe8, 0xa4, 0xf6,
	0xed, 0x04, 0xcf, 0x6b, 0x58, 0xe1, 0x89, 0x6d, 0x3f, 0x87, 0x66, 0x06, 0xc6, 0xed, 0x3c, 0xc3,
	0x3f, 0x22, 0x68, 0x64, 0xba, 0x38, 0x76, 0x96, 0x1f, 0x90, 0x37, 0x3e, 0x2c, 0xb7, 0xf4, 0x7e,
	0xd5, 0x3f, 0x16, 0xd0, 0xad, 0x62, 0x0b, 0x7f, 0x83, 0x60, 0x4d, 0x39, 0x37, 0x65, 0xab, 0xc4,
	0xfd, 0x4b, 0x71, 0xa4, 0x35, 0xb3, 0x8a, 0x39, 0x8a, 0x35, 0xfb, 0x1a, 0xc1, 0x5a, 0x76, 0x7c,
	0xac, 0x32, 0x51, 0xee, 0xa0, 0xb1, 0xb6, 0x2e, 0x58, 0xf9, 0x7e, 0xfc, 0x97, 0xca, 0xf0, 0xd1,
	0x29, 0xa1, 0xd6, 0x4f, 0x08, 0x1a, 0x99, 0xb9, 0xb4, 0x4a, 0xad, 0xbc, 0x01, 0x56, 0xe4, 0xe4,
	0x5e, 0x82, 0x65, 0xd7, 0xde, 0xc9, 0xd7, 0x26, 0x83, 0xc6, 0x35, 0x2d, 0x7d, 0xcf, 0x8c, 0x2a,
	0xfc, 0x3d, 0x02, 0x98, 0xcd, 0x2a, 0x7c, 0x67, 0xa5, 0xb3, 0x17, 0x90, 0x15, 0x77, 0x1c, 0xfb,
	0xad, 0x04, 0x9d, 0x83, 0x77, 0x8a, 0x98, 0x4a, 0xa1, 0xc5, 0xa4, 0xfd, 0x8c, 0xe0, 0xfa, 0xfc,
	0x20, 0xc3, 0x6f, 0xac, 0x76, 0xec, 0xc2, 0xfc, 0xb4, 0x9c, 0xb2, 0xdb, 0xb5, 0xbf, 0xb3, 0x28,
	0x4b, 0x72, 0x18, 0x77, 0x85, 0x46, 0xa6, 0x49, 0xe3, 0x62, 0x42, 0x8a, 0xd4, 0xdc, 0x4d, 0x90,
	0xbc, 0x69, 0x5d, 0x8a, 0xaf, 0xd8, 0xee, 0xdf, 0x22, 0x68, 0x64, 0x26, 0xea, 0x2a, 0x9f, 0xe5,
	0x8d, 0xde, 0xa5, 0x66, 0xd7, 0xe4, 0x74, 0x2e, 0x05, 0xe9, 0xe0, 0x1d, 0x78, 0x65, 0x24, 0x26,
	0x4b, 0x21, 0x1c, 0x58, 0xb9, 0xad, 0xff, 0x38, 0x3e, 0xfa, 0x18, 0x0d, 0xaf, 0x26, 0x18, 0xee,
	0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0xb4, 0xae, 0x02, 0xb9, 0x33, 0x0e, 0x00, 0x00,
}
