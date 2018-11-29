// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/logging/v2/log_entry.proto

package logging

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	monitoredres "google.golang.org/genproto/googleapis/api/monitoredres"
	_type "google.golang.org/genproto/googleapis/logging/type"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// An individual entry in a log.
type LogEntry struct {
	// Required. The resource name of the log to which this log entry belongs:
	//
	//     "projects/[PROJECT_ID]/logs/[LOG_ID]"
	//     "organizations/[ORGANIZATION_ID]/logs/[LOG_ID]"
	//     "billingAccounts/[BILLING_ACCOUNT_ID]/logs/[LOG_ID]"
	//     "folders/[FOLDER_ID]/logs/[LOG_ID]"
	//
	//  A project number may optionally be used in place of PROJECT_ID. The
	//  project number is translated to its corresponding PROJECT_ID internally
	//  and the `log_name` field will contain PROJECT_ID in queries and exports.
	//
	// `[LOG_ID]` must be URL-encoded within `log_name`. Example:
	// `"organizations/1234567890/logs/cloudresourcemanager.googleapis.com%2Factivity"`.
	// `[LOG_ID]` must be less than 512 characters long and can only include the
	// following characters: upper and lower case alphanumeric characters,
	// forward-slash, underscore, hyphen, and period.
	//
	// For backward compatibility, if `log_name` begins with a forward-slash, such
	// as `/projects/...`, then the log entry is ingested as usual but the
	// forward-slash is removed. Listing the log entry will not show the leading
	// slash and filtering for a log name with a leading slash will never return
	// any results.
	LogName string `protobuf:"bytes,12,opt,name=log_name,json=logName,proto3" json:"log_name,omitempty"`
	// Required. The primary monitored resource associated with this log entry.
	// Example: a log entry that reports a database error would be
	// associated with the monitored resource designating the particular
	// database that reported the error.
	Resource *monitoredres.MonitoredResource `protobuf:"bytes,8,opt,name=resource,proto3" json:"resource,omitempty"`
	// Optional. The log entry payload, which can be one of multiple types.
	//
	// Types that are valid to be assigned to Payload:
	//	*LogEntry_ProtoPayload
	//	*LogEntry_TextPayload
	//	*LogEntry_JsonPayload
	Payload isLogEntry_Payload `protobuf_oneof:"payload"`
	// Optional. The time the event described by the log entry occurred.
	// This time is used to compute the log entry's age and to enforce
	// the logs retention period. If this field is omitted in a new log
	// entry, then Logging assigns it the current time.
	// Timestamps have nanosecond accuracy, but trailing zeros in the fractional
	// seconds might be omitted when the timestamp is displayed.
	//
	// Incoming log entries should have timestamps that are no more than
	// the [logs retention period](/logging/quotas) in the past,
	// and no more than 24 hours in the future. Log entries outside those time
	// boundaries will not be available when calling `entries.list`, but
	// those log entries can still be exported with
	// [LogSinks](/logging/docs/api/tasks/exporting-logs).
	Timestamp *timestamp.Timestamp `protobuf:"bytes,9,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Output only. The time the log entry was received by Logging.
	ReceiveTimestamp *timestamp.Timestamp `protobuf:"bytes,24,opt,name=receive_timestamp,json=receiveTimestamp,proto3" json:"receive_timestamp,omitempty"`
	// Optional. The severity of the log entry. The default value is
	// `LogSeverity.DEFAULT`.
	Severity _type.LogSeverity `protobuf:"varint,10,opt,name=severity,proto3,enum=google.logging.type.LogSeverity" json:"severity,omitempty"`
	// Optional. A unique identifier for the log entry. If you provide a value,
	// then Logging considers other log entries in the same project,
	// with the same `timestamp`, and with the same `insert_id` to be duplicates
	// which can be removed.  If omitted in new log entries, then
	// Logging assigns its own unique identifier. The `insert_id` is also used
	// to order log entries that have the same `timestamp` value.
	InsertId string `protobuf:"bytes,4,opt,name=insert_id,json=insertId,proto3" json:"insert_id,omitempty"`
	// Optional. Information about the HTTP request associated with this
	// log entry, if applicable.
	HttpRequest *_type.HttpRequest `protobuf:"bytes,7,opt,name=http_request,json=httpRequest,proto3" json:"http_request,omitempty"`
	// Optional. A set of user-defined (key, value) data that provides additional
	// information about the log entry.
	Labels map[string]string `protobuf:"bytes,11,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Output only. Additional metadata about the monitored resource.
	// Only `k8s_container`, `k8s_pod`, and `k8s_node` MonitoredResources have
	// this field populated.
	Metadata *monitoredres.MonitoredResourceMetadata `protobuf:"bytes,25,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Optional. Information about an operation associated with the log entry, if
	// applicable.
	Operation *LogEntryOperation `protobuf:"bytes,15,opt,name=operation,proto3" json:"operation,omitempty"`
	// Optional. Resource name of the trace associated with the log entry, if any.
	// If it contains a relative resource name, the name is assumed to be relative
	// to `//tracing.googleapis.com`. Example:
	// `projects/my-projectid/traces/06796866738c859f2f19b7cfb3214824`
	Trace string `protobuf:"bytes,22,opt,name=trace,proto3" json:"trace,omitempty"`
	// Optional. The span ID within the trace associated with the log entry.
	// For Trace spans, this is the same format that the Trace
	// API v2 uses: a 16-character hexadecimal encoding of an 8-byte array, such
	// as <code>"000000000000004a"</code>.
	SpanId string `protobuf:"bytes,27,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	// Optional. The sampling decision of the trace associated with the log entry.
	// True means that the trace resource name in the `trace` field was sampled
	// for storage in a trace backend. False means that the trace was not sampled
	// for storage when this log entry was written, or the sampling decision was
	// unknown at the time. A non-sampled `trace` value is still useful as a
	// request correlation identifier. The default is False.
	TraceSampled bool `protobuf:"varint,30,opt,name=trace_sampled,json=traceSampled,proto3" json:"trace_sampled,omitempty"`
	// Optional. Source code location information associated with the log entry,
	// if any.
	SourceLocation       *LogEntrySourceLocation `protobuf:"bytes,23,opt,name=source_location,json=sourceLocation,proto3" json:"source_location,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *LogEntry) Reset()         { *m = LogEntry{} }
func (m *LogEntry) String() string { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()    {}
func (*LogEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ba2017251165146, []int{0}
}

func (m *LogEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogEntry.Unmarshal(m, b)
}
func (m *LogEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogEntry.Marshal(b, m, deterministic)
}
func (m *LogEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogEntry.Merge(m, src)
}
func (m *LogEntry) XXX_Size() int {
	return xxx_messageInfo_LogEntry.Size(m)
}
func (m *LogEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_LogEntry.DiscardUnknown(m)
}

var xxx_messageInfo_LogEntry proto.InternalMessageInfo

func (m *LogEntry) GetLogName() string {
	if m != nil {
		return m.LogName
	}
	return ""
}

func (m *LogEntry) GetResource() *monitoredres.MonitoredResource {
	if m != nil {
		return m.Resource
	}
	return nil
}

type isLogEntry_Payload interface {
	isLogEntry_Payload()
}

type LogEntry_ProtoPayload struct {
	ProtoPayload *any.Any `protobuf:"bytes,2,opt,name=proto_payload,json=protoPayload,proto3,oneof"`
}

type LogEntry_TextPayload struct {
	TextPayload string `protobuf:"bytes,3,opt,name=text_payload,json=textPayload,proto3,oneof"`
}

type LogEntry_JsonPayload struct {
	JsonPayload *_struct.Struct `protobuf:"bytes,6,opt,name=json_payload,json=jsonPayload,proto3,oneof"`
}

func (*LogEntry_ProtoPayload) isLogEntry_Payload() {}

func (*LogEntry_TextPayload) isLogEntry_Payload() {}

func (*LogEntry_JsonPayload) isLogEntry_Payload() {}

func (m *LogEntry) GetPayload() isLogEntry_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *LogEntry) GetProtoPayload() *any.Any {
	if x, ok := m.GetPayload().(*LogEntry_ProtoPayload); ok {
		return x.ProtoPayload
	}
	return nil
}

func (m *LogEntry) GetTextPayload() string {
	if x, ok := m.GetPayload().(*LogEntry_TextPayload); ok {
		return x.TextPayload
	}
	return ""
}

func (m *LogEntry) GetJsonPayload() *_struct.Struct {
	if x, ok := m.GetPayload().(*LogEntry_JsonPayload); ok {
		return x.JsonPayload
	}
	return nil
}

func (m *LogEntry) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *LogEntry) GetReceiveTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.ReceiveTimestamp
	}
	return nil
}

func (m *LogEntry) GetSeverity() _type.LogSeverity {
	if m != nil {
		return m.Severity
	}
	return _type.LogSeverity_DEFAULT
}

func (m *LogEntry) GetInsertId() string {
	if m != nil {
		return m.InsertId
	}
	return ""
}

func (m *LogEntry) GetHttpRequest() *_type.HttpRequest {
	if m != nil {
		return m.HttpRequest
	}
	return nil
}

func (m *LogEntry) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *LogEntry) GetMetadata() *monitoredres.MonitoredResourceMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *LogEntry) GetOperation() *LogEntryOperation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *LogEntry) GetTrace() string {
	if m != nil {
		return m.Trace
	}
	return ""
}

func (m *LogEntry) GetSpanId() string {
	if m != nil {
		return m.SpanId
	}
	return ""
}

func (m *LogEntry) GetTraceSampled() bool {
	if m != nil {
		return m.TraceSampled
	}
	return false
}

func (m *LogEntry) GetSourceLocation() *LogEntrySourceLocation {
	if m != nil {
		return m.SourceLocation
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*LogEntry) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*LogEntry_ProtoPayload)(nil),
		(*LogEntry_TextPayload)(nil),
		(*LogEntry_JsonPayload)(nil),
	}
}

// Additional information about a potentially long-running operation with which
// a log entry is associated.
type LogEntryOperation struct {
	// Optional. An arbitrary operation identifier. Log entries with the
	// same identifier are assumed to be part of the same operation.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Optional. An arbitrary producer identifier. The combination of
	// `id` and `producer` must be globally unique.  Examples for `producer`:
	// `"MyDivision.MyBigCompany.com"`, `"github.com/MyProject/MyApplication"`.
	Producer string `protobuf:"bytes,2,opt,name=producer,proto3" json:"producer,omitempty"`
	// Optional. Set this to True if this is the first log entry in the operation.
	First bool `protobuf:"varint,3,opt,name=first,proto3" json:"first,omitempty"`
	// Optional. Set this to True if this is the last log entry in the operation.
	Last                 bool     `protobuf:"varint,4,opt,name=last,proto3" json:"last,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogEntryOperation) Reset()         { *m = LogEntryOperation{} }
func (m *LogEntryOperation) String() string { return proto.CompactTextString(m) }
func (*LogEntryOperation) ProtoMessage()    {}
func (*LogEntryOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ba2017251165146, []int{1}
}

func (m *LogEntryOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogEntryOperation.Unmarshal(m, b)
}
func (m *LogEntryOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogEntryOperation.Marshal(b, m, deterministic)
}
func (m *LogEntryOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogEntryOperation.Merge(m, src)
}
func (m *LogEntryOperation) XXX_Size() int {
	return xxx_messageInfo_LogEntryOperation.Size(m)
}
func (m *LogEntryOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_LogEntryOperation.DiscardUnknown(m)
}

var xxx_messageInfo_LogEntryOperation proto.InternalMessageInfo

func (m *LogEntryOperation) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *LogEntryOperation) GetProducer() string {
	if m != nil {
		return m.Producer
	}
	return ""
}

func (m *LogEntryOperation) GetFirst() bool {
	if m != nil {
		return m.First
	}
	return false
}

func (m *LogEntryOperation) GetLast() bool {
	if m != nil {
		return m.Last
	}
	return false
}

// Additional information about the source code location that produced the log
// entry.
type LogEntrySourceLocation struct {
	// Optional. Source file name. Depending on the runtime environment, this
	// might be a simple name or a fully-qualified name.
	File string `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	// Optional. Line within the source file. 1-based; 0 indicates no line number
	// available.
	Line int64 `protobuf:"varint,2,opt,name=line,proto3" json:"line,omitempty"`
	// Optional. Human-readable name of the function or method being invoked, with
	// optional context such as the class or package name. This information may be
	// used in contexts such as the logs viewer, where a file and line number are
	// less meaningful. The format can vary by language. For example:
	// `qual.if.ied.Class.method` (Java), `dir/package.func` (Go), `function`
	// (Python).
	Function             string   `protobuf:"bytes,3,opt,name=function,proto3" json:"function,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogEntrySourceLocation) Reset()         { *m = LogEntrySourceLocation{} }
func (m *LogEntrySourceLocation) String() string { return proto.CompactTextString(m) }
func (*LogEntrySourceLocation) ProtoMessage()    {}
func (*LogEntrySourceLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ba2017251165146, []int{2}
}

func (m *LogEntrySourceLocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogEntrySourceLocation.Unmarshal(m, b)
}
func (m *LogEntrySourceLocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogEntrySourceLocation.Marshal(b, m, deterministic)
}
func (m *LogEntrySourceLocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogEntrySourceLocation.Merge(m, src)
}
func (m *LogEntrySourceLocation) XXX_Size() int {
	return xxx_messageInfo_LogEntrySourceLocation.Size(m)
}
func (m *LogEntrySourceLocation) XXX_DiscardUnknown() {
	xxx_messageInfo_LogEntrySourceLocation.DiscardUnknown(m)
}

var xxx_messageInfo_LogEntrySourceLocation proto.InternalMessageInfo

func (m *LogEntrySourceLocation) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

func (m *LogEntrySourceLocation) GetLine() int64 {
	if m != nil {
		return m.Line
	}
	return 0
}

func (m *LogEntrySourceLocation) GetFunction() string {
	if m != nil {
		return m.Function
	}
	return ""
}

func init() {
	proto.RegisterType((*LogEntry)(nil), "google.logging.v2.LogEntry")
	proto.RegisterMapType((map[string]string)(nil), "google.logging.v2.LogEntry.LabelsEntry")
	proto.RegisterType((*LogEntryOperation)(nil), "google.logging.v2.LogEntryOperation")
	proto.RegisterType((*LogEntrySourceLocation)(nil), "google.logging.v2.LogEntrySourceLocation")
}

func init() { proto.RegisterFile("google/logging/v2/log_entry.proto", fileDescriptor_8ba2017251165146) }

var fileDescriptor_8ba2017251165146 = []byte{
	// 772 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0xdd, 0x8e, 0xdb, 0x44,
	0x14, 0xae, 0x93, 0xb2, 0x71, 0x26, 0xd9, 0x6d, 0x77, 0x54, 0x1a, 0x6f, 0x5a, 0x20, 0x6c, 0xf9,
	0x09, 0x37, 0x8e, 0x14, 0x6e, 0xb6, 0xb4, 0x12, 0x6a, 0x2a, 0xd4, 0xad, 0x94, 0x42, 0x35, 0x41,
	0xbd, 0x40, 0x91, 0xac, 0x59, 0x7b, 0xe2, 0x0e, 0xd8, 0x33, 0x66, 0x3c, 0x8e, 0xc8, 0xa3, 0xf0,
	0x0a, 0x3c, 0x0a, 0x6f, 0xc0, 0x9b, 0x70, 0x89, 0xe6, 0xcc, 0xd8, 0x09, 0xc9, 0x2a, 0xbd, 0x3b,
	0x3f, 0xdf, 0x77, 0xce, 0xe7, 0x33, 0xe7, 0x24, 0xe8, 0xf3, 0x54, 0xca, 0x34, 0x63, 0x93, 0x4c,
	0xa6, 0x29, 0x17, 0xe9, 0x64, 0x3d, 0x35, 0x66, 0xc4, 0x84, 0x56, 0x9b, 0xb0, 0x50, 0x52, 0x4b,
	0x7c, 0x6e, 0x21, 0xa1, 0x83, 0x84, 0xeb, 0xe9, 0xf0, 0xb1, 0x63, 0xd1, 0x82, 0x4f, 0xa8, 0x10,
	0x52, 0x53, 0xcd, 0xa5, 0x28, 0x2d, 0x61, 0xf8, 0x64, 0x27, 0x9b, 0x4b, 0xc1, 0xb5, 0x54, 0x2c,
	0x89, 0x14, 0x2b, 0x65, 0xa5, 0x62, 0xe6, 0x40, 0x5f, 0xed, 0x35, 0xd6, 0x9b, 0x82, 0x4d, 0xde,
	0x6b, 0x5d, 0x44, 0x8a, 0xfd, 0x5e, 0xb1, 0x52, 0x1f, 0xc3, 0x19, 0x89, 0x25, 0x5b, 0x33, 0xc5,
	0xb5, 0x53, 0x39, 0xbc, 0x70, 0x38, 0xf0, 0x6e, 0xaa, 0xd5, 0x84, 0x8a, 0x3a, 0xf5, 0x78, 0x3f,
	0x55, 0x6a, 0x55, 0xc5, 0x75, 0x83, 0xcf, 0xf6, 0xb3, 0x9a, 0xe7, 0xac, 0xd4, 0x34, 0x2f, 0x2c,
	0xe0, 0xf2, 0x9f, 0x0e, 0xf2, 0xe7, 0x32, 0xfd, 0xc1, 0x8c, 0x04, 0x5f, 0x20, 0xdf, 0x34, 0x17,
	0x34, 0x67, 0x41, 0x7f, 0xe4, 0x8d, 0xbb, 0xa4, 0x93, 0xc9, 0xf4, 0x47, 0x9a, 0x33, 0xfc, 0x14,
	0xf9, 0xf5, 0x37, 0x06, 0xfe, 0xc8, 0x1b, 0xf7, 0xa6, 0x9f, 0x84, 0x6e, 0x74, 0xb4, 0xe0, 0xe1,
	0x9b, 0x7a, 0x12, 0xc4, 0x81, 0x48, 0x03, 0xc7, 0xcf, 0xd0, 0x29, 0xf4, 0x8a, 0x0a, 0xba, 0xc9,
	0x24, 0x4d, 0x82, 0x16, 0xf0, 0x1f, 0xd4, 0xfc, 0x5a, 0x5b, 0xf8, 0x42, 0x6c, 0xae, 0xef, 0x90,
	0x3e, 0xf8, 0x6f, 0x2d, 0x16, 0x3f, 0x41, 0x7d, 0xcd, 0xfe, 0xd0, 0x0d, 0xb7, 0x6d, 0x64, 0x5d,
	0xdf, 0x21, 0x3d, 0x13, 0xad, 0x41, 0xcf, 0x51, 0xff, 0xd7, 0x52, 0x8a, 0x06, 0x74, 0x02, 0x0d,
	0x06, 0x07, 0x0d, 0x16, 0x30, 0x1a, 0xc3, 0x36, 0xf0, 0x9a, 0x7d, 0x85, 0xba, 0xcd, 0x54, 0x82,
	0x2e, 0x50, 0x87, 0x07, 0xd4, 0x9f, 0x6b, 0x04, 0xd9, 0x82, 0xf1, 0x2b, 0x74, 0xae, 0x58, 0xcc,
	0xf8, 0x9a, 0x45, 0xdb, 0x0a, 0xc1, 0x07, 0x2b, 0xdc, 0x77, 0xa4, 0x26, 0x82, 0x9f, 0x23, 0xbf,
	0x7e, 0xf1, 0x00, 0x8d, 0xbc, 0xf1, 0xd9, 0x74, 0x14, 0xee, 0x2d, 0xa6, 0x59, 0x8d, 0x70, 0x2e,
	0xd3, 0x85, 0xc3, 0x91, 0x86, 0x81, 0x1f, 0xa1, 0x2e, 0x17, 0x25, 0x53, 0x3a, 0xe2, 0x49, 0x70,
	0x17, 0xde, 0xcd, 0xb7, 0x81, 0xd7, 0x09, 0x7e, 0x89, 0xfa, 0xbb, 0x8b, 0x17, 0x74, 0x40, 0xde,
	0xed, 0xe5, 0xaf, 0xb5, 0x2e, 0x88, 0xc5, 0x91, 0xde, 0xfb, 0xad, 0x83, 0xbf, 0x47, 0x27, 0x19,
	0xbd, 0x61, 0x59, 0x19, 0xf4, 0x46, 0xed, 0x71, 0x6f, 0xfa, 0x75, 0x78, 0x70, 0x36, 0x61, 0xbd,
	0x45, 0xe1, 0x1c, 0x90, 0x60, 0x13, 0x47, 0xc3, 0x2f, 0x90, 0x9f, 0x33, 0x4d, 0x13, 0xaa, 0x69,
	0x70, 0x01, 0x0a, 0xbe, 0x3c, 0xba, 0x3e, 0x6f, 0x1c, 0x98, 0x34, 0x34, 0x3c, 0x43, 0x5d, 0x59,
	0x30, 0x05, 0xc7, 0x18, 0xdc, 0x83, 0x1a, 0x5f, 0x1c, 0x91, 0xf1, 0x53, 0x8d, 0x25, 0x5b, 0x1a,
	0x7e, 0x80, 0x3e, 0xd2, 0x8a, 0xc6, 0x2c, 0x78, 0x08, 0x53, 0xb2, 0x0e, 0x1e, 0xa0, 0x4e, 0x59,
	0x50, 0x61, 0xa6, 0xf7, 0x08, 0xe2, 0x27, 0xc6, 0x7d, 0x6d, 0x96, 0xef, 0x14, 0x10, 0x51, 0x49,
	0xf3, 0x22, 0x63, 0x49, 0xf0, 0xe9, 0xc8, 0x1b, 0xfb, 0xa4, 0x0f, 0xc1, 0x85, 0x8d, 0x61, 0x82,
	0xee, 0x59, 0xcd, 0x51, 0x26, 0x63, 0xab, 0x6e, 0x00, 0xea, 0xbe, 0x39, 0xa2, 0x6e, 0x01, 0x8c,
	0xb9, 0x23, 0x90, 0xb3, 0xf2, 0x7f, 0xfe, 0xf0, 0x29, 0xea, 0xed, 0x4c, 0x11, 0xdf, 0x47, 0xed,
	0xdf, 0xd8, 0x26, 0xf0, 0x40, 0x9c, 0x31, 0xcd, 0x87, 0xac, 0x69, 0x56, 0x31, 0xb8, 0xa5, 0x2e,
	0xb1, 0xce, 0x77, 0xad, 0x2b, 0x6f, 0xd6, 0x45, 0x1d, 0x77, 0x06, 0x97, 0x1c, 0x9d, 0x1f, 0x4c,
	0x03, 0x9f, 0xa1, 0x16, 0x4f, 0x5c, 0xa9, 0x16, 0x4f, 0xf0, 0x10, 0xf9, 0x85, 0x92, 0x49, 0x15,
	0x33, 0xe5, 0x8a, 0x35, 0xbe, 0xe9, 0xb2, 0xe2, 0xaa, 0xd4, 0x70, 0x75, 0x3e, 0xb1, 0x0e, 0xc6,
	0xe8, 0x6e, 0x46, 0x4b, 0x0d, 0x9b, 0xe6, 0x13, 0xb0, 0x2f, 0x97, 0xe8, 0xe1, 0xed, 0x9f, 0x66,
	0xd0, 0x2b, 0x9e, 0x31, 0xd7, 0x11, 0x6c, 0xa8, 0xc0, 0x85, 0x15, 0xdf, 0x26, 0x60, 0x1b, 0x1d,
	0xab, 0x4a, 0xc4, 0x30, 0xbf, 0xb6, 0xd5, 0x51, 0xfb, 0xb3, 0x3f, 0x3d, 0xf4, 0x71, 0x2c, 0xf3,
	0xc3, 0x79, 0xce, 0x4e, 0xeb, 0xae, 0x6f, 0xe1, 0x47, 0xc3, 0xfb, 0xe5, 0xca, 0x61, 0x52, 0x99,
	0x51, 0x91, 0x86, 0x52, 0xa5, 0x93, 0x94, 0x09, 0x38, 0xc2, 0x89, 0x4d, 0xd1, 0x82, 0x97, 0x3b,
	0xff, 0x08, 0xcf, 0x9c, 0xf9, 0xaf, 0xe7, 0xfd, 0xd5, 0x1a, 0xbc, 0xb2, 0xec, 0x97, 0x99, 0xac,
	0x12, 0xf3, 0x58, 0xd0, 0xe7, 0xdd, 0xf4, 0xef, 0x3a, 0xb3, 0x84, 0xcc, 0xd2, 0x65, 0x96, 0xef,
	0xa6, 0x37, 0x27, 0x50, 0xfb, 0xdb, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x7b, 0xdc, 0xd5,
	0x6c, 0x06, 0x00, 0x00,
}
