// Copyright 2020 Google LLC
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
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: google/appengine/logging/v1/request_log.proto

package logging

import (
	reflect "reflect"
	sync "sync"

	_type "google.golang.org/genproto/googleapis/logging/type"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Application log line emitted while processing a request.
type LogLine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Approximate time when this log entry was made.
	Time *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	// Severity of this log entry.
	Severity _type.LogSeverity `protobuf:"varint,2,opt,name=severity,proto3,enum=google.logging.type.LogSeverity" json:"severity,omitempty"`
	// App-provided log message.
	LogMessage string `protobuf:"bytes,3,opt,name=log_message,json=logMessage,proto3" json:"log_message,omitempty"`
	// Where in the source code this log message was written.
	SourceLocation *SourceLocation `protobuf:"bytes,4,opt,name=source_location,json=sourceLocation,proto3" json:"source_location,omitempty"`
}

func (x *LogLine) Reset() {
	*x = LogLine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_appengine_logging_v1_request_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogLine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogLine) ProtoMessage() {}

func (x *LogLine) ProtoReflect() protoreflect.Message {
	mi := &file_google_appengine_logging_v1_request_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogLine.ProtoReflect.Descriptor instead.
func (*LogLine) Descriptor() ([]byte, []int) {
	return file_google_appengine_logging_v1_request_log_proto_rawDescGZIP(), []int{0}
}

func (x *LogLine) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *LogLine) GetSeverity() _type.LogSeverity {
	if x != nil {
		return x.Severity
	}
	return _type.LogSeverity(0)
}

func (x *LogLine) GetLogMessage() string {
	if x != nil {
		return x.LogMessage
	}
	return ""
}

func (x *LogLine) GetSourceLocation() *SourceLocation {
	if x != nil {
		return x.SourceLocation
	}
	return nil
}

// Specifies a location in a source code file.
type SourceLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Source file name. Depending on the runtime environment, this might be a
	// simple name or a fully-qualified name.
	File string `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	// Line within the source file.
	Line int64 `protobuf:"varint,2,opt,name=line,proto3" json:"line,omitempty"`
	// Human-readable name of the function or method being invoked, with optional
	// context such as the class or package name. This information is used in
	// contexts such as the logs viewer, where a file and line number are less
	// meaningful. The format can vary by language. For example:
	// `qual.if.ied.Class.method` (Java), `dir/package.func` (Go), `function`
	// (Python).
	FunctionName string `protobuf:"bytes,3,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"`
}

func (x *SourceLocation) Reset() {
	*x = SourceLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_appengine_logging_v1_request_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourceLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceLocation) ProtoMessage() {}

func (x *SourceLocation) ProtoReflect() protoreflect.Message {
	mi := &file_google_appengine_logging_v1_request_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceLocation.ProtoReflect.Descriptor instead.
func (*SourceLocation) Descriptor() ([]byte, []int) {
	return file_google_appengine_logging_v1_request_log_proto_rawDescGZIP(), []int{1}
}

func (x *SourceLocation) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

func (x *SourceLocation) GetLine() int64 {
	if x != nil {
		return x.Line
	}
	return 0
}

func (x *SourceLocation) GetFunctionName() string {
	if x != nil {
		return x.FunctionName
	}
	return ""
}

// A reference to a particular snapshot of the source tree used to build and
// deploy an application.
type SourceReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. A URI string identifying the repository.
	// Example: "https://github.com/GoogleCloudPlatform/kubernetes.git"
	Repository string `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	// The canonical and persistent identifier of the deployed revision.
	// Example (git): "0035781c50ec7aa23385dc841529ce8a4b70db1b"
	RevisionId string `protobuf:"bytes,2,opt,name=revision_id,json=revisionId,proto3" json:"revision_id,omitempty"`
}

func (x *SourceReference) Reset() {
	*x = SourceReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_appengine_logging_v1_request_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourceReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceReference) ProtoMessage() {}

func (x *SourceReference) ProtoReflect() protoreflect.Message {
	mi := &file_google_appengine_logging_v1_request_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceReference.ProtoReflect.Descriptor instead.
func (*SourceReference) Descriptor() ([]byte, []int) {
	return file_google_appengine_logging_v1_request_log_proto_rawDescGZIP(), []int{2}
}

func (x *SourceReference) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *SourceReference) GetRevisionId() string {
	if x != nil {
		return x.RevisionId
	}
	return ""
}

// Complete log information about a single HTTP request to an App Engine
// application.
type RequestLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Application that handled this request.
	AppId string `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	// Module of the application that handled this request.
	ModuleId string `protobuf:"bytes,37,opt,name=module_id,json=moduleId,proto3" json:"module_id,omitempty"`
	// Version of the application that handled this request.
	VersionId string `protobuf:"bytes,2,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	// Globally unique identifier for a request, which is based on the request
	// start time.  Request IDs for requests which started later will compare
	// greater as strings than those for requests which started earlier.
	RequestId string `protobuf:"bytes,3,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Origin IP address.
	Ip string `protobuf:"bytes,4,opt,name=ip,proto3" json:"ip,omitempty"`
	// Time when the request started.
	StartTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Time when the request finished.
	EndTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Latency of the request.
	Latency *durationpb.Duration `protobuf:"bytes,8,opt,name=latency,proto3" json:"latency,omitempty"`
	// Number of CPU megacycles used to process request.
	MegaCycles int64 `protobuf:"varint,9,opt,name=mega_cycles,json=megaCycles,proto3" json:"mega_cycles,omitempty"`
	// Request method. Example: `"GET"`, `"HEAD"`, `"PUT"`, `"POST"`, `"DELETE"`.
	Method string `protobuf:"bytes,10,opt,name=method,proto3" json:"method,omitempty"`
	// Contains the path and query portion of the URL that was requested. For
	// example, if the URL was "http://example.com/app?name=val", the resource
	// would be "/app?name=val".  The fragment identifier, which is identified by
	// the `#` character, is not included.
	Resource string `protobuf:"bytes,11,opt,name=resource,proto3" json:"resource,omitempty"`
	// HTTP version of request. Example: `"HTTP/1.1"`.
	HttpVersion string `protobuf:"bytes,12,opt,name=http_version,json=httpVersion,proto3" json:"http_version,omitempty"`
	// HTTP response status code. Example: 200, 404.
	Status int32 `protobuf:"varint,13,opt,name=status,proto3" json:"status,omitempty"`
	// Size in bytes sent back to client by request.
	ResponseSize int64 `protobuf:"varint,14,opt,name=response_size,json=responseSize,proto3" json:"response_size,omitempty"`
	// Referrer URL of request.
	Referrer string `protobuf:"bytes,15,opt,name=referrer,proto3" json:"referrer,omitempty"`
	// User agent that made the request.
	UserAgent string `protobuf:"bytes,16,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	// The logged-in user who made the request.
	//
	// Most likely, this is the part of the user's email before the `@` sign.  The
	// field value is the same for different requests from the same user, but
	// different users can have similar names.  This information is also
	// available to the application via the App Engine Users API.
	//
	// This field will be populated starting with App Engine 1.9.21.
	Nickname string `protobuf:"bytes,40,opt,name=nickname,proto3" json:"nickname,omitempty"`
	// File or class that handled the request.
	UrlMapEntry string `protobuf:"bytes,17,opt,name=url_map_entry,json=urlMapEntry,proto3" json:"url_map_entry,omitempty"`
	// Internet host and port number of the resource being requested.
	Host string `protobuf:"bytes,20,opt,name=host,proto3" json:"host,omitempty"`
	// An indication of the relative cost of serving this request.
	Cost float64 `protobuf:"fixed64,21,opt,name=cost,proto3" json:"cost,omitempty"`
	// Queue name of the request, in the case of an offline request.
	TaskQueueName string `protobuf:"bytes,22,opt,name=task_queue_name,json=taskQueueName,proto3" json:"task_queue_name,omitempty"`
	// Task name of the request, in the case of an offline request.
	TaskName string `protobuf:"bytes,23,opt,name=task_name,json=taskName,proto3" json:"task_name,omitempty"`
	// Whether this was a loading request for the instance.
	WasLoadingRequest bool `protobuf:"varint,24,opt,name=was_loading_request,json=wasLoadingRequest,proto3" json:"was_loading_request,omitempty"`
	// Time this request spent in the pending request queue.
	PendingTime *durationpb.Duration `protobuf:"bytes,25,opt,name=pending_time,json=pendingTime,proto3" json:"pending_time,omitempty"`
	// If the instance processing this request belongs to a manually scaled
	// module, then this is the 0-based index of the instance. Otherwise, this
	// value is -1.
	InstanceIndex int32 `protobuf:"varint,26,opt,name=instance_index,json=instanceIndex,proto3" json:"instance_index,omitempty"`
	// Whether this request is finished or active.
	Finished bool `protobuf:"varint,27,opt,name=finished,proto3" json:"finished,omitempty"`
	// Whether this is the first `RequestLog` entry for this request.  If an
	// active request has several `RequestLog` entries written to Stackdriver
	// Logging, then this field will be set for one of them.
	First bool `protobuf:"varint,42,opt,name=first,proto3" json:"first,omitempty"`
	// An identifier for the instance that handled the request.
	InstanceId string `protobuf:"bytes,28,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	// A list of log lines emitted by the application while serving this request.
	Line []*LogLine `protobuf:"bytes,29,rep,name=line,proto3" json:"line,omitempty"`
	// App Engine release version.
	AppEngineRelease string `protobuf:"bytes,38,opt,name=app_engine_release,json=appEngineRelease,proto3" json:"app_engine_release,omitempty"`
	// Stackdriver Trace identifier for this request.
	TraceId string `protobuf:"bytes,39,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// If true, the value in the 'trace_id' field was sampled for storage in a
	// trace backend.
	TraceSampled bool `protobuf:"varint,43,opt,name=trace_sampled,json=traceSampled,proto3" json:"trace_sampled,omitempty"`
	// Source code for the application that handled this request. There can be
	// more than one source reference per deployed application if source code is
	// distributed among multiple repositories.
	SourceReference []*SourceReference `protobuf:"bytes,41,rep,name=source_reference,json=sourceReference,proto3" json:"source_reference,omitempty"`
}

func (x *RequestLog) Reset() {
	*x = RequestLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_appengine_logging_v1_request_log_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestLog) ProtoMessage() {}

func (x *RequestLog) ProtoReflect() protoreflect.Message {
	mi := &file_google_appengine_logging_v1_request_log_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestLog.ProtoReflect.Descriptor instead.
func (*RequestLog) Descriptor() ([]byte, []int) {
	return file_google_appengine_logging_v1_request_log_proto_rawDescGZIP(), []int{3}
}

func (x *RequestLog) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *RequestLog) GetModuleId() string {
	if x != nil {
		return x.ModuleId
	}
	return ""
}

func (x *RequestLog) GetVersionId() string {
	if x != nil {
		return x.VersionId
	}
	return ""
}

func (x *RequestLog) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *RequestLog) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *RequestLog) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *RequestLog) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *RequestLog) GetLatency() *durationpb.Duration {
	if x != nil {
		return x.Latency
	}
	return nil
}

func (x *RequestLog) GetMegaCycles() int64 {
	if x != nil {
		return x.MegaCycles
	}
	return 0
}

func (x *RequestLog) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *RequestLog) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *RequestLog) GetHttpVersion() string {
	if x != nil {
		return x.HttpVersion
	}
	return ""
}

func (x *RequestLog) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *RequestLog) GetResponseSize() int64 {
	if x != nil {
		return x.ResponseSize
	}
	return 0
}

func (x *RequestLog) GetReferrer() string {
	if x != nil {
		return x.Referrer
	}
	return ""
}

func (x *RequestLog) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

func (x *RequestLog) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *RequestLog) GetUrlMapEntry() string {
	if x != nil {
		return x.UrlMapEntry
	}
	return ""
}

func (x *RequestLog) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *RequestLog) GetCost() float64 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *RequestLog) GetTaskQueueName() string {
	if x != nil {
		return x.TaskQueueName
	}
	return ""
}

func (x *RequestLog) GetTaskName() string {
	if x != nil {
		return x.TaskName
	}
	return ""
}

func (x *RequestLog) GetWasLoadingRequest() bool {
	if x != nil {
		return x.WasLoadingRequest
	}
	return false
}

func (x *RequestLog) GetPendingTime() *durationpb.Duration {
	if x != nil {
		return x.PendingTime
	}
	return nil
}

func (x *RequestLog) GetInstanceIndex() int32 {
	if x != nil {
		return x.InstanceIndex
	}
	return 0
}

func (x *RequestLog) GetFinished() bool {
	if x != nil {
		return x.Finished
	}
	return false
}

func (x *RequestLog) GetFirst() bool {
	if x != nil {
		return x.First
	}
	return false
}

func (x *RequestLog) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *RequestLog) GetLine() []*LogLine {
	if x != nil {
		return x.Line
	}
	return nil
}

func (x *RequestLog) GetAppEngineRelease() string {
	if x != nil {
		return x.AppEngineRelease
	}
	return ""
}

func (x *RequestLog) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

func (x *RequestLog) GetTraceSampled() bool {
	if x != nil {
		return x.TraceSampled
	}
	return false
}

func (x *RequestLog) GetSourceReference() []*SourceReference {
	if x != nil {
		return x.SourceReference
	}
	return nil
}

var File_google_appengine_logging_v1_request_log_proto protoreflect.FileDescriptor

var file_google_appengine_logging_v1_request_log_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x26, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee, 0x01, 0x0a, 0x07, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x6e,
	0x65, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x3c, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6c, 0x6f, 0x67,
	0x67, 0x69, 0x6e, 0x67, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x76,
	0x65, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12,
	0x1f, 0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x54, 0x0a, 0x0f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x6c, 0x6f, 0x67,
	0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5d, 0x0a, 0x0e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x52, 0x0a, 0x0f, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xbb, 0x09, 0x0a, 0x0a, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x25, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x07,
	0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63,
	0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x67, 0x61, 0x5f, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x73,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6d, 0x65, 0x67, 0x61, 0x43, 0x79, 0x63, 0x6c,
	0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x68, 0x74,
	0x74, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72,
	0x65, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72,
	0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x28, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a,
	0x0d, 0x75, 0x72, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x72, 0x6c, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x61, 0x73,
	0x6b, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x16, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x74, 0x61, 0x73, 0x6b, 0x51, 0x75, 0x65, 0x75, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x17,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e,
	0x0a, 0x13, 0x77, 0x61, 0x73, 0x5f, 0x6c, 0x6f, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x18, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x77, 0x61, 0x73,
	0x4c, 0x6f, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c,
	0x0a, 0x0c, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x19,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0b, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x1a,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18,
	0x1b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x1d,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70,
	0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x6e, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x65,
	0x12, 0x2c, 0x0a, 0x12, 0x61, 0x70, 0x70, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5f, 0x72,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x18, 0x26, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x70,
	0x70, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x27, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x72, 0x61,
	0x63, 0x65, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x64, 0x18, 0x2b, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0c, 0x74, 0x72, 0x61, 0x63, 0x65, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x64, 0x12, 0x57,
	0x0a, 0x10, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x18, 0x29, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x6c, 0x6f, 0x67, 0x67,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x42, 0xe8, 0x01, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f,
	0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x6c, 0x6f, 0x67, 0x67, 0x69,
	0x6e, 0x67, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x67,
	0x69, 0x6e, 0x67, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5c,
	0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x70, 0x70, 0x45,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x3a, 0x3a, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_appengine_logging_v1_request_log_proto_rawDescOnce sync.Once
	file_google_appengine_logging_v1_request_log_proto_rawDescData = file_google_appengine_logging_v1_request_log_proto_rawDesc
)

func file_google_appengine_logging_v1_request_log_proto_rawDescGZIP() []byte {
	file_google_appengine_logging_v1_request_log_proto_rawDescOnce.Do(func() {
		file_google_appengine_logging_v1_request_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_appengine_logging_v1_request_log_proto_rawDescData)
	})
	return file_google_appengine_logging_v1_request_log_proto_rawDescData
}

var file_google_appengine_logging_v1_request_log_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_appengine_logging_v1_request_log_proto_goTypes = []interface{}{
	(*LogLine)(nil),               // 0: google.appengine.logging.v1.LogLine
	(*SourceLocation)(nil),        // 1: google.appengine.logging.v1.SourceLocation
	(*SourceReference)(nil),       // 2: google.appengine.logging.v1.SourceReference
	(*RequestLog)(nil),            // 3: google.appengine.logging.v1.RequestLog
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(_type.LogSeverity)(0),        // 5: google.logging.type.LogSeverity
	(*durationpb.Duration)(nil),   // 6: google.protobuf.Duration
}
var file_google_appengine_logging_v1_request_log_proto_depIdxs = []int32{
	4, // 0: google.appengine.logging.v1.LogLine.time:type_name -> google.protobuf.Timestamp
	5, // 1: google.appengine.logging.v1.LogLine.severity:type_name -> google.logging.type.LogSeverity
	1, // 2: google.appengine.logging.v1.LogLine.source_location:type_name -> google.appengine.logging.v1.SourceLocation
	4, // 3: google.appengine.logging.v1.RequestLog.start_time:type_name -> google.protobuf.Timestamp
	4, // 4: google.appengine.logging.v1.RequestLog.end_time:type_name -> google.protobuf.Timestamp
	6, // 5: google.appengine.logging.v1.RequestLog.latency:type_name -> google.protobuf.Duration
	6, // 6: google.appengine.logging.v1.RequestLog.pending_time:type_name -> google.protobuf.Duration
	0, // 7: google.appengine.logging.v1.RequestLog.line:type_name -> google.appengine.logging.v1.LogLine
	2, // 8: google.appengine.logging.v1.RequestLog.source_reference:type_name -> google.appengine.logging.v1.SourceReference
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_google_appengine_logging_v1_request_log_proto_init() }
func file_google_appengine_logging_v1_request_log_proto_init() {
	if File_google_appengine_logging_v1_request_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_appengine_logging_v1_request_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogLine); i {
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
		file_google_appengine_logging_v1_request_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourceLocation); i {
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
		file_google_appengine_logging_v1_request_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourceReference); i {
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
		file_google_appengine_logging_v1_request_log_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestLog); i {
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
			RawDescriptor: file_google_appengine_logging_v1_request_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_appengine_logging_v1_request_log_proto_goTypes,
		DependencyIndexes: file_google_appengine_logging_v1_request_log_proto_depIdxs,
		MessageInfos:      file_google_appengine_logging_v1_request_log_proto_msgTypes,
	}.Build()
	File_google_appengine_logging_v1_request_log_proto = out.File
	file_google_appengine_logging_v1_request_log_proto_rawDesc = nil
	file_google_appengine_logging_v1_request_log_proto_goTypes = nil
	file_google_appengine_logging_v1_request_log_proto_depIdxs = nil
}
