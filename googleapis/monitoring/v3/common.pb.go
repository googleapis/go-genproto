// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/monitoring/v3/common.proto
// DO NOT EDIT!

package google_monitoring_v3 // import "google.golang.org/genproto/googleapis/monitoring/v3"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_api2 "google.golang.org/genproto/googleapis/api/distribution"
import google_protobuf3 "github.com/golang/protobuf/ptypes/duration"
import google_protobuf2 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The Aligner describes how to bring the data points in a single
// time series into temporal alignment.
type Aggregation_Aligner int32

const (
	// No alignment. Raw data is returned. Not valid if cross-time
	// series reduction is requested. The value type of the result is
	// the same as the value type of the input.
	Aggregation_ALIGN_NONE Aggregation_Aligner = 0
	// Align and convert to delta metric type. This alignment is valid
	// for cumulative metrics and delta metrics. Aligning an existing
	// delta metric to a delta metric requires that the alignment
	// period be increased. The value type of the result is the same
	// as the value type of the input.
	Aggregation_ALIGN_DELTA Aggregation_Aligner = 1
	// Align and convert to a rate. This alignment is valid for
	// cumulative metrics and delta metrics with numeric values. The output is a
	// gauge metric with value type
	// [DOUBLE][google.api.MetricDescriptor.ValueType.DOUBLE].
	Aggregation_ALIGN_RATE Aggregation_Aligner = 2
	// Align by interpolating between adjacent points around the
	// period boundary. This alignment is valid for gauge and delta
	// metrics with numeric values. The value type of the result is the same
	// as the value type of the input.
	Aggregation_ALIGN_INTERPOLATE Aggregation_Aligner = 3
	// Align by shifting the oldest data point before the period
	// boundary to the boundary. This alignment is valid for gauge
	// metrics. The value type of the result is the same as the
	// value type of the input.
	Aggregation_ALIGN_NEXT_OLDER Aggregation_Aligner = 4
	// Align time series via aggregation. The resulting data point in
	// the alignment period is the minimum of all data points in the
	// period. This alignment is valid for gauge and delta metrics with numeric
	// values. The value type of the result is the same as the value
	// type of the input.
	Aggregation_ALIGN_MIN Aggregation_Aligner = 10
	// Align time series via aggregation. The resulting data point in
	// the alignment period is the maximum of all data points in the
	// period. This alignment is valid for gauge and delta metrics with numeric
	// values. The value type of the result is the same as the value
	// type of the input.
	Aggregation_ALIGN_MAX Aggregation_Aligner = 11
	// Align time series via aggregation. The resulting data point in
	// the alignment period is the average or arithmetic mean of all
	// data points in the period. This alignment is valid for gauge and delta
	// metrics with numeric values. The value type of the output is
	// [DOUBLE][google.api.MetricDescriptor.ValueType.DOUBLE].
	Aggregation_ALIGN_MEAN Aggregation_Aligner = 12
	// Align time series via aggregation. The resulting data point in
	// the alignment period is the count of all data points in the
	// period. This alignment is valid for gauge and delta metrics with numeric
	// or Boolean values. The value type of the output is
	// [INT64][google.api.MetricDescriptor.ValueType.INT64].
	Aggregation_ALIGN_COUNT Aggregation_Aligner = 13
	// Align time series via aggregation. The resulting data point in
	// the alignment period is the sum of all data points in the
	// period. This alignment is valid for gauge and delta metrics with numeric
	// and distribution values. The value type of the output is the
	// same as the value type of the input.
	Aggregation_ALIGN_SUM Aggregation_Aligner = 14
	// Align time series via aggregation. The resulting data point in
	// the alignment period is the standard deviation of all data
	// points in the period. This alignment is valid for gauge and delta metrics
	// with numeric values. The value type of the output is
	// [DOUBLE][google.api.MetricDescriptor.ValueType.DOUBLE].
	Aggregation_ALIGN_STDDEV Aggregation_Aligner = 15
	// Align time series via aggregation. The resulting data point in
	// the alignment period is the count of True-valued data points in the
	// period. This alignment is valid for gauge metrics with
	// Boolean values. The value type of the output is
	// [INT64][google.api.MetricDescriptor.ValueType.INT64].
	Aggregation_ALIGN_COUNT_TRUE Aggregation_Aligner = 16
	// Align time series via aggregation. The resulting data point in
	// the alignment period is the fraction of True-valued data points in the
	// period. This alignment is valid for gauge metrics with Boolean values.
	// The output value is in the range [0, 1] and has value type
	// [DOUBLE][google.api.MetricDescriptor.ValueType.DOUBLE].
	Aggregation_ALIGN_FRACTION_TRUE Aggregation_Aligner = 17
	// Align time series via aggregation. The resulting data point in
	// the alignment period is the 99th percentile of all data
	// points in the period. This alignment is valid for gauge and delta metrics
	// with distribution values. The output is a gauge metric with value type
	// [DOUBLE][google.api.MetricDescriptor.ValueType.DOUBLE].
	Aggregation_ALIGN_PERCENTILE_99 Aggregation_Aligner = 18
	// Align time series via aggregation. The resulting data point in
	// the alignment period is the 95th percentile of all data
	// points in the period. This alignment is valid for gauge and delta metrics
	// with distribution values. The output is a gauge metric with value type
	// [DOUBLE][google.api.MetricDescriptor.ValueType.DOUBLE].
	Aggregation_ALIGN_PERCENTILE_95 Aggregation_Aligner = 19
	// Align time series via aggregation. The resulting data point in
	// the alignment period is the 50th percentile of all data
	// points in the period. This alignment is valid for gauge and delta metrics
	// with distribution values. The output is a gauge metric with value type
	// [DOUBLE][google.api.MetricDescriptor.ValueType.DOUBLE].
	Aggregation_ALIGN_PERCENTILE_50 Aggregation_Aligner = 20
	// Align time series via aggregation. The resulting data point in
	// the alignment period is the 5th percentile of all data
	// points in the period. This alignment is valid for gauge and delta metrics
	// with distribution values. The output is a gauge metric with value type
	// [DOUBLE][google.api.MetricDescriptor.ValueType.DOUBLE].
	Aggregation_ALIGN_PERCENTILE_05 Aggregation_Aligner = 21
)

var Aggregation_Aligner_name = map[int32]string{
	0:  "ALIGN_NONE",
	1:  "ALIGN_DELTA",
	2:  "ALIGN_RATE",
	3:  "ALIGN_INTERPOLATE",
	4:  "ALIGN_NEXT_OLDER",
	10: "ALIGN_MIN",
	11: "ALIGN_MAX",
	12: "ALIGN_MEAN",
	13: "ALIGN_COUNT",
	14: "ALIGN_SUM",
	15: "ALIGN_STDDEV",
	16: "ALIGN_COUNT_TRUE",
	17: "ALIGN_FRACTION_TRUE",
	18: "ALIGN_PERCENTILE_99",
	19: "ALIGN_PERCENTILE_95",
	20: "ALIGN_PERCENTILE_50",
	21: "ALIGN_PERCENTILE_05",
}
var Aggregation_Aligner_value = map[string]int32{
	"ALIGN_NONE":          0,
	"ALIGN_DELTA":         1,
	"ALIGN_RATE":          2,
	"ALIGN_INTERPOLATE":   3,
	"ALIGN_NEXT_OLDER":    4,
	"ALIGN_MIN":           10,
	"ALIGN_MAX":           11,
	"ALIGN_MEAN":          12,
	"ALIGN_COUNT":         13,
	"ALIGN_SUM":           14,
	"ALIGN_STDDEV":        15,
	"ALIGN_COUNT_TRUE":    16,
	"ALIGN_FRACTION_TRUE": 17,
	"ALIGN_PERCENTILE_99": 18,
	"ALIGN_PERCENTILE_95": 19,
	"ALIGN_PERCENTILE_50": 20,
	"ALIGN_PERCENTILE_05": 21,
}

func (x Aggregation_Aligner) String() string {
	return proto.EnumName(Aggregation_Aligner_name, int32(x))
}
func (Aggregation_Aligner) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{2, 0} }

// A Reducer describes how to aggregate data points from multiple
// time series into a single time series.
type Aggregation_Reducer int32

const (
	// No cross-time series reduction. The output of the aligner is
	// returned.
	Aggregation_REDUCE_NONE Aggregation_Reducer = 0
	// Reduce by computing the mean across time series for each
	// alignment period. This reducer is valid for delta and
	// gauge metrics with numeric or distribution values. The value type of the
	// output is [DOUBLE][google.api.MetricDescriptor.ValueType.DOUBLE].
	Aggregation_REDUCE_MEAN Aggregation_Reducer = 1
	// Reduce by computing the minimum across time series for each
	// alignment period. This reducer is valid for delta and
	// gauge metrics with numeric values. The value type of the output
	// is the same as the value type of the input.
	Aggregation_REDUCE_MIN Aggregation_Reducer = 2
	// Reduce by computing the maximum across time series for each
	// alignment period. This reducer is valid for delta and
	// gauge metrics with numeric values. The value type of the output
	// is the same as the value type of the input.
	Aggregation_REDUCE_MAX Aggregation_Reducer = 3
	// Reduce by computing the sum across time series for each
	// alignment period. This reducer is valid for delta and
	// gauge metrics with numeric and distribution values. The value type of
	// the output is the same as the value type of the input.
	Aggregation_REDUCE_SUM Aggregation_Reducer = 4
	// Reduce by computing the standard deviation across time series
	// for each alignment period. This reducer is valid for delta
	// and gauge metrics with numeric or distribution values. The value type of
	// the output is [DOUBLE][google.api.MetricDescriptor.ValueType.DOUBLE].
	Aggregation_REDUCE_STDDEV Aggregation_Reducer = 5
	// Reduce by computing the count of data points across time series
	// for each alignment period. This reducer is valid for delta
	// and gauge metrics of numeric, Boolean, distribution, and string value
	// type. The value type of the output is
	// [INT64][google.api.MetricDescriptor.ValueType.INT64].
	Aggregation_REDUCE_COUNT Aggregation_Reducer = 6
	// Reduce by computing the count of True-valued data points across time
	// series for each alignment period. This reducer is valid for delta
	// and gauge metrics of Boolean value type. The value type of
	// the output is [INT64][google.api.MetricDescriptor.ValueType.INT64].
	Aggregation_REDUCE_COUNT_TRUE Aggregation_Reducer = 7
	// Reduce by computing the fraction of True-valued data points across time
	// series for each alignment period. This reducer is valid for delta
	// and gauge metrics of Boolean value type. The output value is in the
	// range [0, 1] and has value type
	// [DOUBLE][google.api.MetricDescriptor.ValueType.DOUBLE].
	Aggregation_REDUCE_FRACTION_TRUE Aggregation_Reducer = 8
	// Reduce by computing 99th percentile of data points across time series
	// for each alignment period. This reducer is valid for gauge and delta
	// metrics of numeric and distribution type. The value of the output is
	// [DOUBLE][google.api.MetricDescriptor.ValueType.DOUBLE]
	Aggregation_REDUCE_PERCENTILE_99 Aggregation_Reducer = 9
	// Reduce by computing 95th percentile of data points across time series
	// for each alignment period. This reducer is valid for gauge and delta
	// metrics of numeric and distribution type. The value of the output is
	// [DOUBLE][google.api.MetricDescriptor.ValueType.DOUBLE]
	Aggregation_REDUCE_PERCENTILE_95 Aggregation_Reducer = 10
	// Reduce by computing 50th percentile of data points across time series
	// for each alignment period. This reducer is valid for gauge and delta
	// metrics of numeric and distribution type. The value of the output is
	// [DOUBLE][google.api.MetricDescriptor.ValueType.DOUBLE]
	Aggregation_REDUCE_PERCENTILE_50 Aggregation_Reducer = 11
	// Reduce by computing 5th percentile of data points across time series
	// for each alignment period. This reducer is valid for gauge and delta
	// metrics of numeric and distribution type. The value of the output is
	// [DOUBLE][google.api.MetricDescriptor.ValueType.DOUBLE]
	Aggregation_REDUCE_PERCENTILE_05 Aggregation_Reducer = 12
)

var Aggregation_Reducer_name = map[int32]string{
	0:  "REDUCE_NONE",
	1:  "REDUCE_MEAN",
	2:  "REDUCE_MIN",
	3:  "REDUCE_MAX",
	4:  "REDUCE_SUM",
	5:  "REDUCE_STDDEV",
	6:  "REDUCE_COUNT",
	7:  "REDUCE_COUNT_TRUE",
	8:  "REDUCE_FRACTION_TRUE",
	9:  "REDUCE_PERCENTILE_99",
	10: "REDUCE_PERCENTILE_95",
	11: "REDUCE_PERCENTILE_50",
	12: "REDUCE_PERCENTILE_05",
}
var Aggregation_Reducer_value = map[string]int32{
	"REDUCE_NONE":          0,
	"REDUCE_MEAN":          1,
	"REDUCE_MIN":           2,
	"REDUCE_MAX":           3,
	"REDUCE_SUM":           4,
	"REDUCE_STDDEV":        5,
	"REDUCE_COUNT":         6,
	"REDUCE_COUNT_TRUE":    7,
	"REDUCE_FRACTION_TRUE": 8,
	"REDUCE_PERCENTILE_99": 9,
	"REDUCE_PERCENTILE_95": 10,
	"REDUCE_PERCENTILE_50": 11,
	"REDUCE_PERCENTILE_05": 12,
}

func (x Aggregation_Reducer) String() string {
	return proto.EnumName(Aggregation_Reducer_name, int32(x))
}
func (Aggregation_Reducer) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{2, 1} }

// A single strongly-typed value.
type TypedValue struct {
	// The typed value field.
	//
	// Types that are valid to be assigned to Value:
	//	*TypedValue_BoolValue
	//	*TypedValue_Int64Value
	//	*TypedValue_DoubleValue
	//	*TypedValue_StringValue
	//	*TypedValue_DistributionValue
	Value isTypedValue_Value `protobuf_oneof:"value"`
}

func (m *TypedValue) Reset()                    { *m = TypedValue{} }
func (m *TypedValue) String() string            { return proto.CompactTextString(m) }
func (*TypedValue) ProtoMessage()               {}
func (*TypedValue) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type isTypedValue_Value interface {
	isTypedValue_Value()
}

type TypedValue_BoolValue struct {
	BoolValue bool `protobuf:"varint,1,opt,name=bool_value,json=boolValue,oneof"`
}
type TypedValue_Int64Value struct {
	Int64Value int64 `protobuf:"varint,2,opt,name=int64_value,json=int64Value,oneof"`
}
type TypedValue_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,3,opt,name=double_value,json=doubleValue,oneof"`
}
type TypedValue_StringValue struct {
	StringValue string `protobuf:"bytes,4,opt,name=string_value,json=stringValue,oneof"`
}
type TypedValue_DistributionValue struct {
	DistributionValue *google_api2.Distribution `protobuf:"bytes,5,opt,name=distribution_value,json=distributionValue,oneof"`
}

func (*TypedValue_BoolValue) isTypedValue_Value()         {}
func (*TypedValue_Int64Value) isTypedValue_Value()        {}
func (*TypedValue_DoubleValue) isTypedValue_Value()       {}
func (*TypedValue_StringValue) isTypedValue_Value()       {}
func (*TypedValue_DistributionValue) isTypedValue_Value() {}

func (m *TypedValue) GetValue() isTypedValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *TypedValue) GetBoolValue() bool {
	if x, ok := m.GetValue().(*TypedValue_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (m *TypedValue) GetInt64Value() int64 {
	if x, ok := m.GetValue().(*TypedValue_Int64Value); ok {
		return x.Int64Value
	}
	return 0
}

func (m *TypedValue) GetDoubleValue() float64 {
	if x, ok := m.GetValue().(*TypedValue_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (m *TypedValue) GetStringValue() string {
	if x, ok := m.GetValue().(*TypedValue_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *TypedValue) GetDistributionValue() *google_api2.Distribution {
	if x, ok := m.GetValue().(*TypedValue_DistributionValue); ok {
		return x.DistributionValue
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TypedValue) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TypedValue_OneofMarshaler, _TypedValue_OneofUnmarshaler, _TypedValue_OneofSizer, []interface{}{
		(*TypedValue_BoolValue)(nil),
		(*TypedValue_Int64Value)(nil),
		(*TypedValue_DoubleValue)(nil),
		(*TypedValue_StringValue)(nil),
		(*TypedValue_DistributionValue)(nil),
	}
}

func _TypedValue_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TypedValue)
	// value
	switch x := m.Value.(type) {
	case *TypedValue_BoolValue:
		t := uint64(0)
		if x.BoolValue {
			t = 1
		}
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *TypedValue_Int64Value:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Int64Value))
	case *TypedValue_DoubleValue:
		b.EncodeVarint(3<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.DoubleValue))
	case *TypedValue_StringValue:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.StringValue)
	case *TypedValue_DistributionValue:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DistributionValue); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("TypedValue.Value has unexpected type %T", x)
	}
	return nil
}

func _TypedValue_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TypedValue)
	switch tag {
	case 1: // value.bool_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &TypedValue_BoolValue{x != 0}
		return true, err
	case 2: // value.int64_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &TypedValue_Int64Value{int64(x)}
		return true, err
	case 3: // value.double_value
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Value = &TypedValue_DoubleValue{math.Float64frombits(x)}
		return true, err
	case 4: // value.string_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &TypedValue_StringValue{x}
		return true, err
	case 5: // value.distribution_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_api2.Distribution)
		err := b.DecodeMessage(msg)
		m.Value = &TypedValue_DistributionValue{msg}
		return true, err
	default:
		return false, nil
	}
}

func _TypedValue_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TypedValue)
	// value
	switch x := m.Value.(type) {
	case *TypedValue_BoolValue:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += 1
	case *TypedValue_Int64Value:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Int64Value))
	case *TypedValue_DoubleValue:
		n += proto.SizeVarint(3<<3 | proto.WireFixed64)
		n += 8
	case *TypedValue_StringValue:
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.StringValue)))
		n += len(x.StringValue)
	case *TypedValue_DistributionValue:
		s := proto.Size(x.DistributionValue)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A time interval extending just after a start time through an end time.
// If the start time is the same as the end time, then the interval
// represents a single point in time.
type TimeInterval struct {
	// Required. The end of the time interval.
	EndTime *google_protobuf2.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	// Optional. The beginning of the time interval.  The default value
	// for the start time is the end time. The start time must not be
	// later than the end time.
	StartTime *google_protobuf2.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
}

func (m *TimeInterval) Reset()                    { *m = TimeInterval{} }
func (m *TimeInterval) String() string            { return proto.CompactTextString(m) }
func (*TimeInterval) ProtoMessage()               {}
func (*TimeInterval) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *TimeInterval) GetEndTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *TimeInterval) GetStartTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

// Describes how to combine, or aggregate, multiple time series to
// provide different views of the data.
// See [Aggregation](/monitoring/api/learn_more#aggregation) for more details.
type Aggregation struct {
	// The alignment period for per-[time series][google.monitoring.v3.TimeSeries]
	// alignment. If present, `alignmentPeriod` must be at least 60
	// seconds.  After per-time series alignment, each time series will
	// contain data points only on the period boundaries. If
	// `perSeriesAligner` is not specified or equals `ALIGN_NONE`, then
	// this field is ignored. If `perSeriesAligner` is specified and
	// does not equal `ALIGN_NONE`, then this field must be defined;
	// otherwise an error is returned.
	AlignmentPeriod *google_protobuf3.Duration `protobuf:"bytes,1,opt,name=alignment_period,json=alignmentPeriod" json:"alignment_period,omitempty"`
	// The approach to be used to align individual time series. Not all
	// alignment functions may be applied to all time series, depending
	// on the metric type and value type of the original time
	// series. Alignment may change the metric type or the value type of
	// the time series.
	//
	// Time series data must be aligned in order to perform cross-time
	// series reduction. If `crossSeriesReducer` is specified, then
	// `perSeriesAligner` must be specified and not equal `ALIGN_NONE`
	// and `alignmentPeriod` must be specified; otherwise, an error is
	// returned.
	PerSeriesAligner Aggregation_Aligner `protobuf:"varint,2,opt,name=per_series_aligner,json=perSeriesAligner,enum=google.monitoring.v3.Aggregation_Aligner" json:"per_series_aligner,omitempty"`
	// The approach to be used to combine time series. Not all reducer
	// functions may be applied to all time series, depending on the
	// metric type and the value type of the original time
	// series. Reduction may change the metric type of value type of the
	// time series.
	//
	// Time series data must be aligned in order to perform cross-time
	// series reduction. If `crossSeriesReducer` is specified, then
	// `perSeriesAligner` must be specified and not equal `ALIGN_NONE`
	// and `alignmentPeriod` must be specified; otherwise, an error is
	// returned.
	CrossSeriesReducer Aggregation_Reducer `protobuf:"varint,4,opt,name=cross_series_reducer,json=crossSeriesReducer,enum=google.monitoring.v3.Aggregation_Reducer" json:"cross_series_reducer,omitempty"`
	// The set of fields to preserve when `crossSeriesReducer` is
	// specified. The `groupByFields` determine how the time series
	// are partitioned into subsets prior to applying the aggregation
	// function. Each subset contains time series that have the same
	// value for each of the grouping fields. Each individual time
	// series is a member of exactly one subset. The
	// `crossSeriesReducer` is applied to each subset of time series.
	// Fields not specified in `groupByFields` are aggregated away.
	// If `groupByFields` is not specified, the time series are
	// aggregated into a single output time series. If
	// `crossSeriesReducer` is not defined, this field is ignored.
	GroupByFields []string `protobuf:"bytes,5,rep,name=group_by_fields,json=groupByFields" json:"group_by_fields,omitempty"`
}

func (m *Aggregation) Reset()                    { *m = Aggregation{} }
func (m *Aggregation) String() string            { return proto.CompactTextString(m) }
func (*Aggregation) ProtoMessage()               {}
func (*Aggregation) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *Aggregation) GetAlignmentPeriod() *google_protobuf3.Duration {
	if m != nil {
		return m.AlignmentPeriod
	}
	return nil
}

func init() {
	proto.RegisterType((*TypedValue)(nil), "google.monitoring.v3.TypedValue")
	proto.RegisterType((*TimeInterval)(nil), "google.monitoring.v3.TimeInterval")
	proto.RegisterType((*Aggregation)(nil), "google.monitoring.v3.Aggregation")
	proto.RegisterEnum("google.monitoring.v3.Aggregation_Aligner", Aggregation_Aligner_name, Aggregation_Aligner_value)
	proto.RegisterEnum("google.monitoring.v3.Aggregation_Reducer", Aggregation_Reducer_name, Aggregation_Reducer_value)
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/monitoring/v3/common.proto", fileDescriptor2)
}

var fileDescriptor2 = []byte{
	// 775 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xcf, 0x6e, 0xe2, 0x46,
	0x18, 0xc7, 0x90, 0x2c, 0xe1, 0x33, 0x09, 0x93, 0x59, 0x56, 0xa5, 0xb9, 0x94, 0x52, 0xa9, 0x62,
	0x2f, 0x76, 0x94, 0x94, 0x4a, 0x51, 0x2f, 0x75, 0xc0, 0xdb, 0x58, 0x22, 0x06, 0xcd, 0x3a, 0xdb,
	0x48, 0x3d, 0x58, 0x06, 0xcf, 0xba, 0x23, 0x61, 0x8f, 0x35, 0x36, 0x91, 0x72, 0xeb, 0x6b, 0xf4,
	0x15, 0xfa, 0x48, 0x7d, 0x88, 0x3e, 0x43, 0xe5, 0x99, 0x61, 0x71, 0x54, 0xaa, 0xdd, 0x1b, 0xdf,
	0xef, 0xdf, 0x78, 0x7e, 0x9f, 0x2d, 0xe0, 0xe7, 0x84, 0xf3, 0x64, 0x43, 0xad, 0x84, 0x6f, 0xa2,
	0x2c, 0xb1, 0xb8, 0x48, 0xec, 0x84, 0x66, 0xb9, 0xe0, 0x25, 0xb7, 0x15, 0x15, 0xe5, 0xac, 0xb0,
	0x53, 0x9e, 0xb1, 0x92, 0x0b, 0x96, 0x25, 0xf6, 0xd3, 0xb5, 0xbd, 0xe6, 0x69, 0xca, 0x33, 0x4b,
	0xaa, 0x70, 0x5f, 0x27, 0xec, 0x25, 0xd6, 0xd3, 0xf5, 0x85, 0xf7, 0x65, 0xb9, 0x51, 0xce, 0xec,
	0x98, 0x15, 0xa5, 0x60, 0xab, 0x6d, 0xc9, 0x78, 0xf6, 0x62, 0x50, 0x07, 0x5c, 0xdc, 0x24, 0xac,
	0xfc, 0x7d, 0xbb, 0xb2, 0xd6, 0x3c, 0xb5, 0x55, 0x9c, 0x2d, 0x89, 0xd5, 0xf6, 0xa3, 0x9d, 0x97,
	0xcf, 0x39, 0x2d, 0xec, 0x78, 0x2b, 0x22, 0x65, 0xd7, 0x3f, 0xb4, 0xf5, 0xa7, 0xcf, 0x5b, 0x4b,
	0x96, 0xd2, 0xa2, 0x8c, 0xd2, 0x7c, 0xff, 0x4b, 0x99, 0x47, 0xff, 0x18, 0x00, 0xc1, 0x73, 0x4e,
	0xe3, 0x0f, 0xd1, 0x66, 0x4b, 0xf1, 0x37, 0x00, 0x2b, 0xce, 0x37, 0xe1, 0x53, 0x35, 0x0d, 0x8c,
	0xa1, 0x31, 0x3e, 0xb9, 0x6b, 0x90, 0x4e, 0x85, 0x29, 0xc1, 0xb7, 0x60, 0xb2, 0xac, 0xfc, 0xf1,
	0x07, 0xad, 0x68, 0x0e, 0x8d, 0x71, 0xeb, 0xae, 0x41, 0x40, 0x82, 0x4a, 0xf2, 0x1d, 0x74, 0x63,
	0xbe, 0x5d, 0x6d, 0xa8, 0xd6, 0xb4, 0x86, 0xc6, 0xd8, 0xb8, 0x6b, 0x10, 0x53, 0xa1, 0x9f, 0x44,
	0x55, 0x07, 0x59, 0xa2, 0x45, 0x47, 0x43, 0x63, 0xdc, 0xa9, 0x44, 0x0a, 0x55, 0x22, 0x0f, 0x70,
	0xbd, 0x2a, 0x2d, 0x3d, 0x1e, 0x1a, 0x63, 0xf3, 0x6a, 0x60, 0xe9, 0xf2, 0xa3, 0x9c, 0x59, 0xb3,
	0x9a, 0xea, 0xae, 0x41, 0xce, 0xeb, 0x2e, 0x19, 0x75, 0xdb, 0x86, 0x63, 0xe9, 0x1e, 0xfd, 0x61,
	0x40, 0x37, 0x60, 0x29, 0xf5, 0xb2, 0x92, 0x8a, 0xa7, 0x68, 0x83, 0x27, 0x70, 0x42, 0xb3, 0x38,
	0xac, 0x8a, 0x91, 0xd7, 0x31, 0xaf, 0x2e, 0x76, 0xd1, 0xbb, 0x1a, 0xad, 0x60, 0xd7, 0x1a, 0x69,
	0xd3, 0x2c, 0xae, 0x26, 0x7c, 0x03, 0x50, 0x94, 0x91, 0x28, 0x95, 0xd1, 0xf8, 0xac, 0xb1, 0x23,
	0xd5, 0xd5, 0x3c, 0xfa, 0xab, 0x0d, 0xa6, 0x93, 0x24, 0x82, 0x26, 0x72, 0x8d, 0x78, 0x06, 0x28,
	0xda, 0xb0, 0x24, 0x4b, 0x69, 0x56, 0x86, 0x39, 0x15, 0x8c, 0xc7, 0x3a, 0xf0, 0xeb, 0xff, 0x04,
	0xce, 0xf4, 0xee, 0x49, 0xef, 0x93, 0x65, 0x29, 0x1d, 0xf8, 0x57, 0xc0, 0x39, 0x15, 0x61, 0x41,
	0x05, 0xa3, 0x45, 0x28, 0x59, 0x2a, 0xe4, 0x8d, 0xce, 0xae, 0xde, 0x5a, 0x87, 0xde, 0x5f, 0xab,
	0xf6, 0x10, 0x96, 0xa3, 0x0c, 0x04, 0xe5, 0x54, 0xbc, 0x97, 0x19, 0x1a, 0xc1, 0xbf, 0x41, 0x7f,
	0x2d, 0x78, 0x51, 0xec, 0xa2, 0x05, 0x8d, 0xb7, 0x6b, 0x2a, 0xe4, 0xca, 0xbe, 0x28, 0x9a, 0x28,
	0x03, 0xc1, 0x32, 0x46, 0x85, 0x6b, 0x0c, 0x7f, 0x0f, 0xbd, 0x44, 0xf0, 0x6d, 0x1e, 0xae, 0x9e,
	0xc3, 0x8f, 0x8c, 0x6e, 0xe2, 0x62, 0x70, 0x3c, 0x6c, 0x8d, 0x3b, 0xe4, 0x54, 0xc2, 0xb7, 0xcf,
	0xef, 0x24, 0x38, 0xfa, 0xbb, 0x09, 0xed, 0xdd, 0x03, 0x9d, 0x01, 0x38, 0x73, 0xef, 0x17, 0x3f,
	0xf4, 0x17, 0xbe, 0x8b, 0x1a, 0xb8, 0x07, 0xa6, 0x9a, 0x67, 0xee, 0x3c, 0x70, 0x90, 0xb1, 0x17,
	0x10, 0x27, 0x70, 0x51, 0x13, 0xbf, 0x81, 0x73, 0x35, 0x7b, 0x7e, 0xe0, 0x92, 0xe5, 0x62, 0x5e,
	0xc1, 0x2d, 0xdc, 0x07, 0xa4, 0x73, 0xdc, 0xc7, 0x20, 0x5c, 0xcc, 0x67, 0x2e, 0x41, 0x47, 0xf8,
	0x14, 0x3a, 0x0a, 0xbd, 0xf7, 0x7c, 0x04, 0xb5, 0xd1, 0x79, 0x44, 0xe6, 0x3e, 0xfa, 0xde, 0x75,
	0x7c, 0xd4, 0xdd, 0x9f, 0x3d, 0x5d, 0x3c, 0xf8, 0x01, 0x3a, 0xdd, 0xeb, 0xdf, 0x3f, 0xdc, 0xa3,
	0x33, 0x8c, 0xa0, 0xab, 0xc7, 0x60, 0x36, 0x73, 0x3f, 0xa0, 0xde, 0xfe, 0x54, 0xe9, 0x08, 0x03,
	0xf2, 0xe0, 0x22, 0x84, 0xbf, 0x82, 0xd7, 0x0a, 0x7d, 0x47, 0x9c, 0x69, 0xe0, 0x2d, 0x7c, 0x45,
	0x9c, 0xef, 0x89, 0xa5, 0x4b, 0xa6, 0xae, 0x1f, 0x78, 0x73, 0x37, 0xbc, 0xb9, 0x41, 0xf8, 0x30,
	0x31, 0x41, 0xaf, 0x0f, 0x12, 0x93, 0x4b, 0xd4, 0x3f, 0x48, 0x5c, 0x4e, 0xd0, 0x9b, 0xd1, 0x9f,
	0x4d, 0x68, 0xef, 0x16, 0xd2, 0x03, 0x93, 0xb8, 0xb3, 0x87, 0xa9, 0x5b, 0x6b, 0x57, 0x03, 0xf2,
	0xca, 0xb2, 0xdd, 0x1d, 0xe0, 0xf9, 0xa8, 0x59, 0x9f, 0x9d, 0x47, 0xd4, 0xaa, 0xcd, 0x55, 0x05,
	0x47, 0xf8, 0x1c, 0x4e, 0x77, 0xb3, 0xea, 0xe0, 0xb8, 0x6a, 0x45, 0x43, 0xaa, 0xb6, 0x57, 0xd5,
	0x8a, 0xea, 0x88, 0xba, 0x7d, 0x1b, 0x0f, 0xa0, 0xaf, 0xe1, 0x97, 0xbd, 0x9c, 0xd4, 0x98, 0x97,
	0xc5, 0x74, 0xfe, 0x87, 0x99, 0x20, 0x38, 0xcc, 0x4c, 0x2e, 0x91, 0x79, 0x98, 0xb9, 0x9c, 0xa0,
	0xee, 0xed, 0x5b, 0x18, 0xac, 0x79, 0x7a, 0xf0, 0x25, 0xbf, 0x35, 0xa7, 0xf2, 0x3f, 0x62, 0x59,
	0x7d, 0x9c, 0x4b, 0x63, 0xf5, 0x4a, 0x7e, 0xa5, 0xd7, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xc5,
	0xb8, 0x84, 0x60, 0x6f, 0x06, 0x00, 0x00,
}
