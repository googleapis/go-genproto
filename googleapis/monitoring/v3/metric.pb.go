// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/monitoring/v3/metric.proto
// DO NOT EDIT!

package google_monitoring_v3

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_api5 "google.golang.org/genproto/googleapis/api/metric"
import google_api4 "google.golang.org/genproto/googleapis/api/monitoredres"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A single data point in a time series.
type Point struct {
	// The time interval to which the value applies.
	Interval *TimeInterval `protobuf:"bytes,1,opt,name=interval" json:"interval,omitempty"`
	// The value of the data point.
	Value *TypedValue `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (m *Point) String() string            { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *Point) GetInterval() *TimeInterval {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *Point) GetValue() *TypedValue {
	if m != nil {
		return m.Value
	}
	return nil
}

// A collection of data points that describes the time-varying nature
// of a metric. A time series is identified by a combination of a
// fully-specified monitored resource and a fully-specified metric.
type TimeSeries struct {
	// The fully-specified metric used to identify the time series.
	Metric *google_api5.Metric `protobuf:"bytes,1,opt,name=metric" json:"metric,omitempty"`
	// The fully-specified monitored resource used to identify the time series.
	Resource *google_api4.MonitoredResource `protobuf:"bytes,2,opt,name=resource" json:"resource,omitempty"`
	// The metric kind of the time series. This can be different than the metric
	// kind specified in [google.api.MetricDescriptor] because of alignment and
	// reduction operations on the data. This field is ignored when writing data;
	// the value specified in the descriptor is used instead.
	// @OutputOnly
	MetricKind google_api5.MetricDescriptor_MetricKind `protobuf:"varint,3,opt,name=metric_kind,json=metricKind,enum=google.api.MetricDescriptor_MetricKind" json:"metric_kind,omitempty"`
	// The value type of the time series. This can be different than the value
	// type specified in [google.api.MetricDescriptor] because of alignment and
	// reduction operations on the data. This field is ignored when writing data;
	// the value specified in the descriptor is used instead.
	// @OutputOnly
	ValueType google_api5.MetricDescriptor_ValueType `protobuf:"varint,4,opt,name=value_type,json=valueType,enum=google.api.MetricDescriptor_ValueType" json:"value_type,omitempty"`
	// The data points of this time series. When used as output, points will be
	// sorted by decreasing time order. When used as input, points could be
	// written in any orders.
	Points []*Point `protobuf:"bytes,5,rep,name=points" json:"points,omitempty"`
}

func (m *TimeSeries) Reset()                    { *m = TimeSeries{} }
func (m *TimeSeries) String() string            { return proto.CompactTextString(m) }
func (*TimeSeries) ProtoMessage()               {}
func (*TimeSeries) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *TimeSeries) GetMetric() *google_api5.Metric {
	if m != nil {
		return m.Metric
	}
	return nil
}

func (m *TimeSeries) GetResource() *google_api4.MonitoredResource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *TimeSeries) GetPoints() []*Point {
	if m != nil {
		return m.Points
	}
	return nil
}

func init() {
	proto.RegisterType((*Point)(nil), "google.monitoring.v3.Point")
	proto.RegisterType((*TimeSeries)(nil), "google.monitoring.v3.TimeSeries")
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/monitoring/v3/metric.proto", fileDescriptor5)
}

var fileDescriptor5 = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xe9, 0xe6, 0xc6, 0x4c, 0xc1, 0x43, 0xf0, 0x50, 0x26, 0x42, 0xd9, 0x41, 0xa7, 0x87,
	0x16, 0x56, 0x10, 0x3c, 0x28, 0x32, 0x14, 0x14, 0x11, 0x47, 0x14, 0xaf, 0xa3, 0xb6, 0x8f, 0x12,
	0x6c, 0xf3, 0x42, 0xda, 0x15, 0x76, 0xf2, 0x73, 0xf9, 0xed, 0xa4, 0x49, 0xba, 0x31, 0x9c, 0xa2,
	0xde, 0xd2, 0xe6, 0xf7, 0x7e, 0x2f, 0xef, 0xff, 0xc8, 0x55, 0x86, 0x98, 0xe5, 0x10, 0x64, 0x98,
	0xc7, 0x22, 0x0b, 0x50, 0x65, 0x61, 0x06, 0x42, 0x2a, 0xac, 0x30, 0x34, 0x57, 0xb1, 0xe4, 0x65,
	0x58, 0xa0, 0xe0, 0x15, 0x2a, 0x2e, 0xb2, 0xb0, 0x8e, 0xc2, 0x02, 0x2a, 0xc5, 0x93, 0x40, 0x53,
	0x74, 0xdf, 0x1a, 0xd6, 0x48, 0x50, 0x47, 0xc3, 0x8b, 0xdf, 0x79, 0x63, 0xc9, 0xad, 0x6d, 0x43,
	0x3a, 0x7c, 0xfc, 0x43, 0xb9, 0xe9, 0x0b, 0xa9, 0x82, 0x72, 0xfd, 0x31, 0x57, 0x50, 0xe2, 0x42,
	0x25, 0x60, 0x85, 0xff, 0x9a, 0x33, 0xc1, 0xa2, 0x40, 0x61, 0x0c, 0xa3, 0x77, 0xd2, 0x9b, 0x21,
	0x17, 0x15, 0xbd, 0x24, 0x03, 0x2e, 0x2a, 0x50, 0x75, 0x9c, 0x7b, 0x8e, 0xef, 0x8c, 0xdd, 0xc9,
	0x28, 0xd8, 0x96, 0x41, 0xf0, 0xcc, 0x0b, 0xb8, 0xb3, 0x24, 0x5b, 0xd5, 0xd0, 0x33, 0xd2, 0xab,
	0xe3, 0x7c, 0x01, 0x5e, 0x47, 0x17, 0xfb, 0xdf, 0x14, 0x2f, 0x25, 0xa4, 0x2f, 0x0d, 0xc7, 0x0c,
	0x3e, 0xfa, 0xe8, 0x10, 0xd2, 0x28, 0x9f, 0x40, 0x71, 0x28, 0xe9, 0x29, 0xe9, 0x9b, 0xc8, 0xec,
	0x23, 0x68, 0xeb, 0x89, 0x25, 0x0f, 0x1e, 0xf4, 0x0d, 0xb3, 0x04, 0x3d, 0x27, 0x83, 0x36, 0x0f,
	0xdb, 0xf5, 0x70, 0x83, 0x6e, 0x53, 0x63, 0x16, 0x62, 0x2b, 0x9c, 0xde, 0x12, 0xd7, 0x48, 0xe6,
	0x6f, 0x5c, 0xa4, 0x5e, 0xd7, 0x77, 0xc6, 0x7b, 0x93, 0xe3, 0xaf, 0xbd, 0xae, 0xa1, 0x4c, 0x14,
	0x97, 0x15, 0x2a, 0xfb, 0xe3, 0x9e, 0x8b, 0x94, 0x91, 0x62, 0x75, 0xa6, 0x37, 0x84, 0xe8, 0x41,
	0xe6, 0xd5, 0x52, 0x82, 0xb7, 0xa3, 0x45, 0x47, 0x3f, 0x8a, 0xf4, 0xf8, 0x4d, 0x10, 0x6c, 0xb7,
	0x6e, 0x8f, 0x34, 0x22, 0x7d, 0xd9, 0xec, 0xa1, 0xf4, 0x7a, 0x7e, 0x77, 0xec, 0x4e, 0x0e, 0xb6,
	0xe7, 0xa7, 0x77, 0xc5, 0x2c, 0x3a, 0x3d, 0x21, 0x5e, 0x82, 0xc5, 0x56, 0x72, 0xea, 0x9a, 0xbe,
	0xb3, 0x66, 0xcb, 0x33, 0xe7, 0xb5, 0xaf, 0xd7, 0x1d, 0x7d, 0x06, 0x00, 0x00, 0xff, 0xff, 0x42,
	0x89, 0xb9, 0x1f, 0x1a, 0x03, 0x00, 0x00,
}
