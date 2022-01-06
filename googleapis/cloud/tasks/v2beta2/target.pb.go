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
// source: google/cloud/tasks/v2beta2/target.proto

package tasks

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The HTTP method used to execute the task.
type HttpMethod int32

const (
	// HTTP method unspecified
	HttpMethod_HTTP_METHOD_UNSPECIFIED HttpMethod = 0
	// HTTP POST
	HttpMethod_POST HttpMethod = 1
	// HTTP GET
	HttpMethod_GET HttpMethod = 2
	// HTTP HEAD
	HttpMethod_HEAD HttpMethod = 3
	// HTTP PUT
	HttpMethod_PUT HttpMethod = 4
	// HTTP DELETE
	HttpMethod_DELETE HttpMethod = 5
)

// Enum value maps for HttpMethod.
var (
	HttpMethod_name = map[int32]string{
		0: "HTTP_METHOD_UNSPECIFIED",
		1: "POST",
		2: "GET",
		3: "HEAD",
		4: "PUT",
		5: "DELETE",
	}
	HttpMethod_value = map[string]int32{
		"HTTP_METHOD_UNSPECIFIED": 0,
		"POST":                    1,
		"GET":                     2,
		"HEAD":                    3,
		"PUT":                     4,
		"DELETE":                  5,
	}
)

func (x HttpMethod) Enum() *HttpMethod {
	p := new(HttpMethod)
	*p = x
	return p
}

func (x HttpMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HttpMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_tasks_v2beta2_target_proto_enumTypes[0].Descriptor()
}

func (HttpMethod) Type() protoreflect.EnumType {
	return &file_google_cloud_tasks_v2beta2_target_proto_enumTypes[0]
}

func (x HttpMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HttpMethod.Descriptor instead.
func (HttpMethod) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_tasks_v2beta2_target_proto_rawDescGZIP(), []int{0}
}

// Pull target.
type PullTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PullTarget) Reset() {
	*x = PullTarget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_tasks_v2beta2_target_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullTarget) ProtoMessage() {}

func (x *PullTarget) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_tasks_v2beta2_target_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullTarget.ProtoReflect.Descriptor instead.
func (*PullTarget) Descriptor() ([]byte, []int) {
	return file_google_cloud_tasks_v2beta2_target_proto_rawDescGZIP(), []int{0}
}

// The pull message contains data that can be used by the caller of
// [LeaseTasks][google.cloud.tasks.v2beta2.CloudTasks.LeaseTasks] to process the task.
//
// This proto can only be used for tasks in a queue which has
// [pull_target][google.cloud.tasks.v2beta2.Queue.pull_target] set.
type PullMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A data payload consumed by the worker to execute the task.
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	// The task's tag.
	//
	// Tags allow similar tasks to be processed in a batch. If you label
	// tasks with a tag, your worker can
	// [lease tasks][google.cloud.tasks.v2beta2.CloudTasks.LeaseTasks] with the same tag using
	// [filter][google.cloud.tasks.v2beta2.LeaseTasksRequest.filter]. For example, if you want to
	// aggregate the events associated with a specific user once a day,
	// you could tag tasks with the user ID.
	//
	// The task's tag can only be set when the
	// [task is created][google.cloud.tasks.v2beta2.CloudTasks.CreateTask].
	//
	// The tag must be less than 500 characters.
	//
	// SDK compatibility: Although the SDK allows tags to be either
	// string or
	// [bytes](https://cloud.google.com/appengine/docs/standard/java/javadoc/com/google/appengine/api/taskqueue/TaskOptions.html#tag-byte:A-),
	// only UTF-8 encoded tags can be used in Cloud Tasks. If a tag isn't UTF-8
	// encoded, the tag will be empty when the task is returned by Cloud Tasks.
	Tag string `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *PullMessage) Reset() {
	*x = PullMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_tasks_v2beta2_target_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullMessage) ProtoMessage() {}

func (x *PullMessage) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_tasks_v2beta2_target_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullMessage.ProtoReflect.Descriptor instead.
func (*PullMessage) Descriptor() ([]byte, []int) {
	return file_google_cloud_tasks_v2beta2_target_proto_rawDescGZIP(), []int{1}
}

func (x *PullMessage) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *PullMessage) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

// App Engine HTTP target.
//
// The task will be delivered to the App Engine application hostname
// specified by its [AppEngineHttpTarget][google.cloud.tasks.v2beta2.AppEngineHttpTarget] and [AppEngineHttpRequest][google.cloud.tasks.v2beta2.AppEngineHttpRequest].
// The documentation for [AppEngineHttpRequest][google.cloud.tasks.v2beta2.AppEngineHttpRequest] explains how the
// task's host URL is constructed.
//
// Using [AppEngineHttpTarget][google.cloud.tasks.v2beta2.AppEngineHttpTarget] requires
// [`appengine.applications.get`](https://cloud.google.com/appengine/docs/admin-api/access-control)
// Google IAM permission for the project
// and the following scope:
//
// `https://www.googleapis.com/auth/cloud-platform`
type AppEngineHttpTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Overrides for the
	// [task-level app_engine_routing][google.cloud.tasks.v2beta2.AppEngineHttpRequest.app_engine_routing].
	//
	// If set, `app_engine_routing_override` is used for all tasks in
	// the queue, no matter what the setting is for the
	// [task-level app_engine_routing][google.cloud.tasks.v2beta2.AppEngineHttpRequest.app_engine_routing].
	AppEngineRoutingOverride *AppEngineRouting `protobuf:"bytes,1,opt,name=app_engine_routing_override,json=appEngineRoutingOverride,proto3" json:"app_engine_routing_override,omitempty"`
}

func (x *AppEngineHttpTarget) Reset() {
	*x = AppEngineHttpTarget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_tasks_v2beta2_target_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppEngineHttpTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppEngineHttpTarget) ProtoMessage() {}

func (x *AppEngineHttpTarget) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_tasks_v2beta2_target_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppEngineHttpTarget.ProtoReflect.Descriptor instead.
func (*AppEngineHttpTarget) Descriptor() ([]byte, []int) {
	return file_google_cloud_tasks_v2beta2_target_proto_rawDescGZIP(), []int{2}
}

func (x *AppEngineHttpTarget) GetAppEngineRoutingOverride() *AppEngineRouting {
	if x != nil {
		return x.AppEngineRoutingOverride
	}
	return nil
}

// App Engine HTTP request.
//
// The message defines the HTTP request that is sent to an App Engine app when
// the task is dispatched.
//
// This proto can only be used for tasks in a queue which has
// [app_engine_http_target][google.cloud.tasks.v2beta2.Queue.app_engine_http_target] set.
//
// Using [AppEngineHttpRequest][google.cloud.tasks.v2beta2.AppEngineHttpRequest] requires
// [`appengine.applications.get`](https://cloud.google.com/appengine/docs/admin-api/access-control)
// Google IAM permission for the project
// and the following scope:
//
// `https://www.googleapis.com/auth/cloud-platform`
//
// The task will be delivered to the App Engine app which belongs to the same
// project as the queue. For more information, see
// [How Requests are
// Routed](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed)
// and how routing is affected by
// [dispatch
// files](https://cloud.google.com/appengine/docs/python/config/dispatchref).
// Traffic is encrypted during transport and never leaves Google datacenters.
// Because this traffic is carried over a communication mechanism internal to
// Google, you cannot explicitly set the protocol (for example, HTTP or HTTPS).
// The request to the handler, however, will appear to have used the HTTP
// protocol.
//
// The [AppEngineRouting][google.cloud.tasks.v2beta2.AppEngineRouting] used to construct the URL that the task is
// delivered to can be set at the queue-level or task-level:
//
// * If set,
//   [app_engine_routing_override][google.cloud.tasks.v2beta2.AppEngineHttpTarget.app_engine_routing_override]
//   is used for all tasks in the queue, no matter what the setting
//   is for the
//   [task-level app_engine_routing][google.cloud.tasks.v2beta2.AppEngineHttpRequest.app_engine_routing].
//
//
// The `url` that the task will be sent to is:
//
// * `url =` [host][google.cloud.tasks.v2beta2.AppEngineRouting.host] `+`
//   [relative_url][google.cloud.tasks.v2beta2.AppEngineHttpRequest.relative_url]
//
// Tasks can be dispatched to secure app handlers, unsecure app handlers, and
// URIs restricted with
// [`login:
// admin`](https://cloud.google.com/appengine/docs/standard/python/config/appref).
// Because tasks are not run as any user, they cannot be dispatched to URIs
// restricted with
// [`login:
// required`](https://cloud.google.com/appengine/docs/standard/python/config/appref)
// Task dispatches also do not follow redirects.
//
// The task attempt has succeeded if the app's request handler returns an HTTP
// response code in the range [`200` - `299`]. The task attempt has failed if
// the app's handler returns a non-2xx response code or Cloud Tasks does
// not receive response before the [deadline][Task.dispatch_deadline]. Failed
// tasks will be retried according to the
// [retry configuration][google.cloud.tasks.v2beta2.Queue.retry_config]. `503` (Service Unavailable) is
// considered an App Engine system error instead of an application error and
// will cause Cloud Tasks' traffic congestion control to temporarily throttle
// the queue's dispatches. Unlike other types of task targets, a `429` (Too Many
// Requests) response from an app handler does not cause traffic congestion
// control to throttle the queue.
type AppEngineHttpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The HTTP method to use for the request. The default is POST.
	//
	// The app's request handler for the task's target URL must be able to handle
	// HTTP requests with this http_method, otherwise the task attempt fails with
	// error code 405 (Method Not Allowed). See [Writing a push task request
	// handler](https://cloud.google.com/appengine/docs/java/taskqueue/push/creating-handlers#writing_a_push_task_request_handler)
	// and the App Engine documentation for your runtime on [How Requests are
	// Handled](https://cloud.google.com/appengine/docs/standard/python3/how-requests-are-handled).
	HttpMethod HttpMethod `protobuf:"varint,1,opt,name=http_method,json=httpMethod,proto3,enum=google.cloud.tasks.v2beta2.HttpMethod" json:"http_method,omitempty"`
	// Task-level setting for App Engine routing.
	//
	// If set,
	// [app_engine_routing_override][google.cloud.tasks.v2beta2.AppEngineHttpTarget.app_engine_routing_override]
	// is used for all tasks in the queue, no matter what the setting is for the
	// [task-level app_engine_routing][google.cloud.tasks.v2beta2.AppEngineHttpRequest.app_engine_routing].
	AppEngineRouting *AppEngineRouting `protobuf:"bytes,2,opt,name=app_engine_routing,json=appEngineRouting,proto3" json:"app_engine_routing,omitempty"`
	// The relative URL.
	//
	// The relative URL must begin with "/" and must be a valid HTTP relative URL.
	// It can contain a path and query string arguments.
	// If the relative URL is empty, then the root path "/" will be used.
	// No spaces are allowed, and the maximum length allowed is 2083 characters.
	RelativeUrl string `protobuf:"bytes,3,opt,name=relative_url,json=relativeUrl,proto3" json:"relative_url,omitempty"`
	// HTTP request headers.
	//
	// This map contains the header field names and values.
	// Headers can be set when the
	// [task is created][google.cloud.tasks.v2beta2.CloudTasks.CreateTask].
	// Repeated headers are not supported but a header value can contain commas.
	//
	// Cloud Tasks sets some headers to default values:
	//
	// * `User-Agent`: By default, this header is
	//   `"AppEngine-Google; (+http://code.google.com/appengine)"`.
	//   This header can be modified, but Cloud Tasks will append
	//   `"AppEngine-Google; (+http://code.google.com/appengine)"` to the
	//   modified `User-Agent`.
	//
	// If the task has a [payload][google.cloud.tasks.v2beta2.AppEngineHttpRequest.payload], Cloud
	// Tasks sets the following headers:
	//
	// * `Content-Type`: By default, the `Content-Type` header is set to
	//   `"application/octet-stream"`. The default can be overridden by explicitly
	//   setting `Content-Type` to a particular media type when the
	//   [task is created][google.cloud.tasks.v2beta2.CloudTasks.CreateTask].
	//   For example, `Content-Type` can be set to `"application/json"`.
	// * `Content-Length`: This is computed by Cloud Tasks. This value is
	//   output only.   It cannot be changed.
	//
	// The headers below cannot be set or overridden:
	//
	// * `Host`
	// * `X-Google-*`
	// * `X-AppEngine-*`
	//
	// In addition, Cloud Tasks sets some headers when the task is dispatched,
	// such as headers containing information about the task; see
	// [request
	// headers](https://cloud.google.com/appengine/docs/python/taskqueue/push/creating-handlers#reading_request_headers).
	// These headers are set only when the task is dispatched, so they are not
	// visible when the task is returned in a Cloud Tasks response.
	//
	// Although there is no specific limit for the maximum number of headers or
	// the size, there is a limit on the maximum size of the [Task][google.cloud.tasks.v2beta2.Task]. For more
	// information, see the [CreateTask][google.cloud.tasks.v2beta2.CloudTasks.CreateTask] documentation.
	Headers map[string]string `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Payload.
	//
	// The payload will be sent as the HTTP message body. A message
	// body, and thus a payload, is allowed only if the HTTP method is
	// POST or PUT. It is an error to set a data payload on a task with
	// an incompatible [HttpMethod][google.cloud.tasks.v2beta2.HttpMethod].
	Payload []byte `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *AppEngineHttpRequest) Reset() {
	*x = AppEngineHttpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_tasks_v2beta2_target_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppEngineHttpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppEngineHttpRequest) ProtoMessage() {}

func (x *AppEngineHttpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_tasks_v2beta2_target_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppEngineHttpRequest.ProtoReflect.Descriptor instead.
func (*AppEngineHttpRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_tasks_v2beta2_target_proto_rawDescGZIP(), []int{3}
}

func (x *AppEngineHttpRequest) GetHttpMethod() HttpMethod {
	if x != nil {
		return x.HttpMethod
	}
	return HttpMethod_HTTP_METHOD_UNSPECIFIED
}

func (x *AppEngineHttpRequest) GetAppEngineRouting() *AppEngineRouting {
	if x != nil {
		return x.AppEngineRouting
	}
	return nil
}

func (x *AppEngineHttpRequest) GetRelativeUrl() string {
	if x != nil {
		return x.RelativeUrl
	}
	return ""
}

func (x *AppEngineHttpRequest) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *AppEngineHttpRequest) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

// App Engine Routing.
//
// Defines routing characteristics specific to App Engine - service, version,
// and instance.
//
// For more information about services, versions, and instances see
// [An Overview of App
// Engine](https://cloud.google.com/appengine/docs/python/an-overview-of-app-engine),
// [Microservices Architecture on Google App
// Engine](https://cloud.google.com/appengine/docs/python/microservices-on-app-engine),
// [App Engine Standard request
// routing](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed),
// and [App Engine Flex request
// routing](https://cloud.google.com/appengine/docs/flexible/python/how-requests-are-routed).
type AppEngineRouting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// App service.
	//
	// By default, the task is sent to the service which is the default
	// service when the task is attempted.
	//
	// For some queues or tasks which were created using the App Engine
	// Task Queue API, [host][google.cloud.tasks.v2beta2.AppEngineRouting.host] is not parsable
	// into [service][google.cloud.tasks.v2beta2.AppEngineRouting.service],
	// [version][google.cloud.tasks.v2beta2.AppEngineRouting.version], and
	// [instance][google.cloud.tasks.v2beta2.AppEngineRouting.instance]. For example, some tasks
	// which were created using the App Engine SDK use a custom domain
	// name; custom domains are not parsed by Cloud Tasks. If
	// [host][google.cloud.tasks.v2beta2.AppEngineRouting.host] is not parsable, then
	// [service][google.cloud.tasks.v2beta2.AppEngineRouting.service],
	// [version][google.cloud.tasks.v2beta2.AppEngineRouting.version], and
	// [instance][google.cloud.tasks.v2beta2.AppEngineRouting.instance] are the empty string.
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// App version.
	//
	// By default, the task is sent to the version which is the default
	// version when the task is attempted.
	//
	// For some queues or tasks which were created using the App Engine
	// Task Queue API, [host][google.cloud.tasks.v2beta2.AppEngineRouting.host] is not parsable
	// into [service][google.cloud.tasks.v2beta2.AppEngineRouting.service],
	// [version][google.cloud.tasks.v2beta2.AppEngineRouting.version], and
	// [instance][google.cloud.tasks.v2beta2.AppEngineRouting.instance]. For example, some tasks
	// which were created using the App Engine SDK use a custom domain
	// name; custom domains are not parsed by Cloud Tasks. If
	// [host][google.cloud.tasks.v2beta2.AppEngineRouting.host] is not parsable, then
	// [service][google.cloud.tasks.v2beta2.AppEngineRouting.service],
	// [version][google.cloud.tasks.v2beta2.AppEngineRouting.version], and
	// [instance][google.cloud.tasks.v2beta2.AppEngineRouting.instance] are the empty string.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// App instance.
	//
	// By default, the task is sent to an instance which is available when
	// the task is attempted.
	//
	// Requests can only be sent to a specific instance if
	// [manual scaling is used in App Engine
	// Standard](https://cloud.google.com/appengine/docs/python/an-overview-of-app-engine?hl=en_US#scaling_types_and_instance_classes).
	// App Engine Flex does not support instances. For more information, see
	// [App Engine Standard request
	// routing](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed)
	// and [App Engine Flex request
	// routing](https://cloud.google.com/appengine/docs/flexible/python/how-requests-are-routed).
	Instance string `protobuf:"bytes,3,opt,name=instance,proto3" json:"instance,omitempty"`
	// Output only. The host that the task is sent to.
	//
	// For more information, see
	// [How Requests are
	// Routed](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed).
	//
	// The host is constructed as:
	//
	//
	// * `host = [application_domain_name]`</br>
	//   `| [service] + '.' + [application_domain_name]`</br>
	//   `| [version] + '.' + [application_domain_name]`</br>
	//   `| [version_dot_service]+ '.' + [application_domain_name]`</br>
	//   `| [instance] + '.' + [application_domain_name]`</br>
	//   `| [instance_dot_service] + '.' + [application_domain_name]`</br>
	//   `| [instance_dot_version] + '.' + [application_domain_name]`</br>
	//   `| [instance_dot_version_dot_service] + '.' + [application_domain_name]`
	//
	// * `application_domain_name` = The domain name of the app, for
	//   example <app-id>.appspot.com, which is associated with the
	//   queue's project ID. Some tasks which were created using the App Engine
	//   SDK use a custom domain name.
	//
	// * `service =` [service][google.cloud.tasks.v2beta2.AppEngineRouting.service]
	//
	// * `version =` [version][google.cloud.tasks.v2beta2.AppEngineRouting.version]
	//
	// * `version_dot_service =`
	//   [version][google.cloud.tasks.v2beta2.AppEngineRouting.version] `+ '.' +`
	//   [service][google.cloud.tasks.v2beta2.AppEngineRouting.service]
	//
	// * `instance =` [instance][google.cloud.tasks.v2beta2.AppEngineRouting.instance]
	//
	// * `instance_dot_service =`
	//   [instance][google.cloud.tasks.v2beta2.AppEngineRouting.instance] `+ '.' +`
	//   [service][google.cloud.tasks.v2beta2.AppEngineRouting.service]
	//
	// * `instance_dot_version =`
	//   [instance][google.cloud.tasks.v2beta2.AppEngineRouting.instance] `+ '.' +`
	//   [version][google.cloud.tasks.v2beta2.AppEngineRouting.version]
	//
	// * `instance_dot_version_dot_service =`
	//   [instance][google.cloud.tasks.v2beta2.AppEngineRouting.instance] `+ '.' +`
	//   [version][google.cloud.tasks.v2beta2.AppEngineRouting.version] `+ '.' +`
	//   [service][google.cloud.tasks.v2beta2.AppEngineRouting.service]
	//
	// If [service][google.cloud.tasks.v2beta2.AppEngineRouting.service] is empty, then the task will be sent
	// to the service which is the default service when the task is attempted.
	//
	// If [version][google.cloud.tasks.v2beta2.AppEngineRouting.version] is empty, then the task will be sent
	// to the version which is the default version when the task is attempted.
	//
	// If [instance][google.cloud.tasks.v2beta2.AppEngineRouting.instance] is empty, then the task
	// will be sent to an instance which is available when the task is
	// attempted.
	//
	// If [service][google.cloud.tasks.v2beta2.AppEngineRouting.service],
	// [version][google.cloud.tasks.v2beta2.AppEngineRouting.version], or
	// [instance][google.cloud.tasks.v2beta2.AppEngineRouting.instance] is invalid, then the task
	// will be sent to the default version of the default service when
	// the task is attempted.
	Host string `protobuf:"bytes,4,opt,name=host,proto3" json:"host,omitempty"`
}

func (x *AppEngineRouting) Reset() {
	*x = AppEngineRouting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_tasks_v2beta2_target_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppEngineRouting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppEngineRouting) ProtoMessage() {}

func (x *AppEngineRouting) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_tasks_v2beta2_target_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppEngineRouting.ProtoReflect.Descriptor instead.
func (*AppEngineRouting) Descriptor() ([]byte, []int) {
	return file_google_cloud_tasks_v2beta2_target_proto_rawDescGZIP(), []int{4}
}

func (x *AppEngineRouting) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *AppEngineRouting) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *AppEngineRouting) GetInstance() string {
	if x != nil {
		return x.Instance
	}
	return ""
}

func (x *AppEngineRouting) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

var File_google_cloud_tasks_v2beta2_target_proto protoreflect.FileDescriptor

var file_google_cloud_tasks_v2beta2_target_proto_rawDesc = []byte{
	0x0a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2f, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x32,
	0x62, 0x65, 0x74, 0x61, 0x32, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x0c, 0x0a, 0x0a, 0x50, 0x75, 0x6c, 0x6c, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x22, 0x39, 0x0a, 0x0b, 0x50, 0x75, 0x6c, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x22, 0x82, 0x01, 0x0a,
	0x13, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x48, 0x74, 0x74, 0x70, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x12, 0x6b, 0x0a, 0x1b, 0x61, 0x70, 0x70, 0x5f, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x72,
	0x69, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76,
	0x32, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x18, 0x61, 0x70, 0x70, 0x45, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64,
	0x65, 0x22, 0x8d, 0x03, 0x0a, 0x14, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x48,
	0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x0b, 0x68, 0x74,
	0x74, 0x70, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e, 0x48, 0x74, 0x74,
	0x70, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x0a, 0x68, 0x74, 0x74, 0x70, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x12, 0x5a, 0x0a, 0x12, 0x61, 0x70, 0x70, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e, 0x41, 0x70, 0x70,
	0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x10, 0x61,
	0x70, 0x70, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x21, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x55,
	0x72, 0x6c, 0x12, 0x57, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x32,
	0x2e, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x76, 0x0a, 0x10, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x52, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x2a, 0x5b, 0x0a, 0x0a, 0x48, 0x74, 0x74,
	0x70, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1b, 0x0a, 0x17, 0x48, 0x54, 0x54, 0x50, 0x5f,
	0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x4f, 0x53, 0x54, 0x10, 0x01, 0x12, 0x07,
	0x0a, 0x03, 0x47, 0x45, 0x54, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x45, 0x41, 0x44, 0x10,
	0x03, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x55, 0x54, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45,
	0x4c, 0x45, 0x54, 0x45, 0x10, 0x05, 0x42, 0x70, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73,
	0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x32, 0x42, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74,
	0x61, 0x32, 0x3b, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_tasks_v2beta2_target_proto_rawDescOnce sync.Once
	file_google_cloud_tasks_v2beta2_target_proto_rawDescData = file_google_cloud_tasks_v2beta2_target_proto_rawDesc
)

func file_google_cloud_tasks_v2beta2_target_proto_rawDescGZIP() []byte {
	file_google_cloud_tasks_v2beta2_target_proto_rawDescOnce.Do(func() {
		file_google_cloud_tasks_v2beta2_target_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_tasks_v2beta2_target_proto_rawDescData)
	})
	return file_google_cloud_tasks_v2beta2_target_proto_rawDescData
}

var file_google_cloud_tasks_v2beta2_target_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_tasks_v2beta2_target_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_cloud_tasks_v2beta2_target_proto_goTypes = []interface{}{
	(HttpMethod)(0),              // 0: google.cloud.tasks.v2beta2.HttpMethod
	(*PullTarget)(nil),           // 1: google.cloud.tasks.v2beta2.PullTarget
	(*PullMessage)(nil),          // 2: google.cloud.tasks.v2beta2.PullMessage
	(*AppEngineHttpTarget)(nil),  // 3: google.cloud.tasks.v2beta2.AppEngineHttpTarget
	(*AppEngineHttpRequest)(nil), // 4: google.cloud.tasks.v2beta2.AppEngineHttpRequest
	(*AppEngineRouting)(nil),     // 5: google.cloud.tasks.v2beta2.AppEngineRouting
	nil,                          // 6: google.cloud.tasks.v2beta2.AppEngineHttpRequest.HeadersEntry
}
var file_google_cloud_tasks_v2beta2_target_proto_depIdxs = []int32{
	5, // 0: google.cloud.tasks.v2beta2.AppEngineHttpTarget.app_engine_routing_override:type_name -> google.cloud.tasks.v2beta2.AppEngineRouting
	0, // 1: google.cloud.tasks.v2beta2.AppEngineHttpRequest.http_method:type_name -> google.cloud.tasks.v2beta2.HttpMethod
	5, // 2: google.cloud.tasks.v2beta2.AppEngineHttpRequest.app_engine_routing:type_name -> google.cloud.tasks.v2beta2.AppEngineRouting
	6, // 3: google.cloud.tasks.v2beta2.AppEngineHttpRequest.headers:type_name -> google.cloud.tasks.v2beta2.AppEngineHttpRequest.HeadersEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_tasks_v2beta2_target_proto_init() }
func file_google_cloud_tasks_v2beta2_target_proto_init() {
	if File_google_cloud_tasks_v2beta2_target_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_tasks_v2beta2_target_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullTarget); i {
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
		file_google_cloud_tasks_v2beta2_target_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullMessage); i {
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
		file_google_cloud_tasks_v2beta2_target_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppEngineHttpTarget); i {
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
		file_google_cloud_tasks_v2beta2_target_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppEngineHttpRequest); i {
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
		file_google_cloud_tasks_v2beta2_target_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppEngineRouting); i {
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
			RawDescriptor: file_google_cloud_tasks_v2beta2_target_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_tasks_v2beta2_target_proto_goTypes,
		DependencyIndexes: file_google_cloud_tasks_v2beta2_target_proto_depIdxs,
		EnumInfos:         file_google_cloud_tasks_v2beta2_target_proto_enumTypes,
		MessageInfos:      file_google_cloud_tasks_v2beta2_target_proto_msgTypes,
	}.Build()
	File_google_cloud_tasks_v2beta2_target_proto = out.File
	file_google_cloud_tasks_v2beta2_target_proto_rawDesc = nil
	file_google_cloud_tasks_v2beta2_target_proto_goTypes = nil
	file_google_cloud_tasks_v2beta2_target_proto_depIdxs = nil
}
