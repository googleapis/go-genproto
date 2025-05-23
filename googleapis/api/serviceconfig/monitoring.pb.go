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
// source: google/api/monitoring.proto

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

// Monitoring configuration of the service.
//
// The example below shows how to configure monitored resources and metrics
// for monitoring. In the example, a monitored resource and two metrics are
// defined. The `library.googleapis.com/book/returned_count` metric is sent
// to both producer and consumer projects, whereas the
// `library.googleapis.com/book/num_overdue` metric is only sent to the
// consumer project.
//
//	monitored_resources:
//	- type: library.googleapis.com/Branch
//	  display_name: "Library Branch"
//	  description: "A branch of a library."
//	  launch_stage: GA
//	  labels:
//	  - key: resource_container
//	    description: "The Cloud container (ie. project id) for the Branch."
//	  - key: location
//	    description: "The location of the library branch."
//	  - key: branch_id
//	    description: "The id of the branch."
//	metrics:
//	- name: library.googleapis.com/book/returned_count
//	  display_name: "Books Returned"
//	  description: "The count of books that have been returned."
//	  launch_stage: GA
//	  metric_kind: DELTA
//	  value_type: INT64
//	  unit: "1"
//	  labels:
//	  - key: customer_id
//	    description: "The id of the customer."
//	- name: library.googleapis.com/book/num_overdue
//	  display_name: "Books Overdue"
//	  description: "The current number of overdue books."
//	  launch_stage: GA
//	  metric_kind: GAUGE
//	  value_type: INT64
//	  unit: "1"
//	  labels:
//	  - key: customer_id
//	    description: "The id of the customer."
//	monitoring:
//	  producer_destinations:
//	  - monitored_resource: library.googleapis.com/Branch
//	    metrics:
//	    - library.googleapis.com/book/returned_count
//	  consumer_destinations:
//	  - monitored_resource: library.googleapis.com/Branch
//	    metrics:
//	    - library.googleapis.com/book/returned_count
//	    - library.googleapis.com/book/num_overdue
type Monitoring struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Monitoring configurations for sending metrics to the producer project.
	// There can be multiple producer destinations. A monitored resource type may
	// appear in multiple monitoring destinations if different aggregations are
	// needed for different sets of metrics associated with that monitored
	// resource type. A monitored resource and metric pair may only be used once
	// in the Monitoring configuration.
	ProducerDestinations []*Monitoring_MonitoringDestination `protobuf:"bytes,1,rep,name=producer_destinations,json=producerDestinations,proto3" json:"producer_destinations,omitempty"`
	// Monitoring configurations for sending metrics to the consumer project.
	// There can be multiple consumer destinations. A monitored resource type may
	// appear in multiple monitoring destinations if different aggregations are
	// needed for different sets of metrics associated with that monitored
	// resource type. A monitored resource and metric pair may only be used once
	// in the Monitoring configuration.
	ConsumerDestinations []*Monitoring_MonitoringDestination `protobuf:"bytes,2,rep,name=consumer_destinations,json=consumerDestinations,proto3" json:"consumer_destinations,omitempty"`
}

func (x *Monitoring) Reset() {
	*x = Monitoring{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_monitoring_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Monitoring) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Monitoring) ProtoMessage() {}

func (x *Monitoring) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_monitoring_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Monitoring.ProtoReflect.Descriptor instead.
func (*Monitoring) Descriptor() ([]byte, []int) {
	return file_google_api_monitoring_proto_rawDescGZIP(), []int{0}
}

func (x *Monitoring) GetProducerDestinations() []*Monitoring_MonitoringDestination {
	if x != nil {
		return x.ProducerDestinations
	}
	return nil
}

func (x *Monitoring) GetConsumerDestinations() []*Monitoring_MonitoringDestination {
	if x != nil {
		return x.ConsumerDestinations
	}
	return nil
}

// Configuration of a specific monitoring destination (the producer project
// or the consumer project).
type Monitoring_MonitoringDestination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The monitored resource type. The type must be defined in
	// [Service.monitored_resources][google.api.Service.monitored_resources]
	// section.
	MonitoredResource string `protobuf:"bytes,1,opt,name=monitored_resource,json=monitoredResource,proto3" json:"monitored_resource,omitempty"`
	// Types of the metrics to report to this monitoring destination.
	// Each type must be defined in
	// [Service.metrics][google.api.Service.metrics] section.
	Metrics []string `protobuf:"bytes,2,rep,name=metrics,proto3" json:"metrics,omitempty"`
}

func (x *Monitoring_MonitoringDestination) Reset() {
	*x = Monitoring_MonitoringDestination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_monitoring_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Monitoring_MonitoringDestination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Monitoring_MonitoringDestination) ProtoMessage() {}

func (x *Monitoring_MonitoringDestination) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_monitoring_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Monitoring_MonitoringDestination.ProtoReflect.Descriptor instead.
func (*Monitoring_MonitoringDestination) Descriptor() ([]byte, []int) {
	return file_google_api_monitoring_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Monitoring_MonitoringDestination) GetMonitoredResource() string {
	if x != nil {
		return x.MonitoredResource
	}
	return ""
}

func (x *Monitoring_MonitoringDestination) GetMetrics() []string {
	if x != nil {
		return x.Metrics
	}
	return nil
}

var File_google_api_monitoring_proto protoreflect.FileDescriptor

var file_google_api_monitoring_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x22, 0xb4, 0x02, 0x0a, 0x0a, 0x4d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x61, 0x0a, 0x15, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e,
	0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x14, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x44,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x61, 0x0a, 0x15, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x14, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x72, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x60,
	0x0a, 0x15, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x12, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x42, 0x71, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x42, 0x0f, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0xa2, 0x02, 0x04, 0x47,
	0x41, 0x50, 0x49, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_api_monitoring_proto_rawDescOnce sync.Once
	file_google_api_monitoring_proto_rawDescData = file_google_api_monitoring_proto_rawDesc
)

func file_google_api_monitoring_proto_rawDescGZIP() []byte {
	file_google_api_monitoring_proto_rawDescOnce.Do(func() {
		file_google_api_monitoring_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_api_monitoring_proto_rawDescData)
	})
	return file_google_api_monitoring_proto_rawDescData
}

var file_google_api_monitoring_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_api_monitoring_proto_goTypes = []interface{}{
	(*Monitoring)(nil),                       // 0: google.api.Monitoring
	(*Monitoring_MonitoringDestination)(nil), // 1: google.api.Monitoring.MonitoringDestination
}
var file_google_api_monitoring_proto_depIdxs = []int32{
	1, // 0: google.api.Monitoring.producer_destinations:type_name -> google.api.Monitoring.MonitoringDestination
	1, // 1: google.api.Monitoring.consumer_destinations:type_name -> google.api.Monitoring.MonitoringDestination
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_api_monitoring_proto_init() }
func file_google_api_monitoring_proto_init() {
	if File_google_api_monitoring_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_api_monitoring_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Monitoring); i {
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
		file_google_api_monitoring_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Monitoring_MonitoringDestination); i {
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
			RawDescriptor: file_google_api_monitoring_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_api_monitoring_proto_goTypes,
		DependencyIndexes: file_google_api_monitoring_proto_depIdxs,
		MessageInfos:      file_google_api_monitoring_proto_msgTypes,
	}.Build()
	File_google_api_monitoring_proto = out.File
	file_google_api_monitoring_proto_rawDesc = nil
	file_google_api_monitoring_proto_goTypes = nil
	file_google_api_monitoring_proto_depIdxs = nil
}
