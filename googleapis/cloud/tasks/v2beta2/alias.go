// Copyright 2022 Google LLC
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

// Code generated by aliasgen. DO NOT EDIT.

// Package cloudtasks aliases all exported identifiers in package
// "cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb".
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb.
// Please read https://github.com/googleapis/google-cloud-go/blob/main/migration.md
// for more details.
package cloudtasks

import (
	src "cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb"
	grpc "google.golang.org/grpc"
)

// Deprecated: Please use consts in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
const (
	HttpMethod_DELETE                  = src.HttpMethod_DELETE
	HttpMethod_GET                     = src.HttpMethod_GET
	HttpMethod_HEAD                    = src.HttpMethod_HEAD
	HttpMethod_HTTP_METHOD_UNSPECIFIED = src.HttpMethod_HTTP_METHOD_UNSPECIFIED
	HttpMethod_POST                    = src.HttpMethod_POST
	HttpMethod_PUT                     = src.HttpMethod_PUT
	Queue_DISABLED                     = src.Queue_DISABLED
	Queue_PAUSED                       = src.Queue_PAUSED
	Queue_RUNNING                      = src.Queue_RUNNING
	Queue_STATE_UNSPECIFIED            = src.Queue_STATE_UNSPECIFIED
	Task_BASIC                         = src.Task_BASIC
	Task_FULL                          = src.Task_FULL
	Task_VIEW_UNSPECIFIED              = src.Task_VIEW_UNSPECIFIED
)

// Deprecated: Please use vars in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
var (
	File_google_cloud_tasks_v2beta2_cloudtasks_proto = src.File_google_cloud_tasks_v2beta2_cloudtasks_proto
	File_google_cloud_tasks_v2beta2_queue_proto      = src.File_google_cloud_tasks_v2beta2_queue_proto
	File_google_cloud_tasks_v2beta2_target_proto     = src.File_google_cloud_tasks_v2beta2_target_proto
	File_google_cloud_tasks_v2beta2_task_proto       = src.File_google_cloud_tasks_v2beta2_task_proto
	HttpMethod_name                                  = src.HttpMethod_name
	HttpMethod_value                                 = src.HttpMethod_value
	Queue_State_name                                 = src.Queue_State_name
	Queue_State_value                                = src.Queue_State_value
	Task_View_name                                   = src.Task_View_name
	Task_View_value                                  = src.Task_View_value
)

// Request message for acknowledging a task using
// [AcknowledgeTask][google.cloud.tasks.v2beta2.CloudTasks.AcknowledgeTask].
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
type AcknowledgeTaskRequest = src.AcknowledgeTaskRequest

// App Engine HTTP request. The message defines the HTTP request that is sent
// to an App Engine app when the task is dispatched. This proto can only be
// used for tasks in a queue which has
// [app_engine_http_target][google.cloud.tasks.v2beta2.Queue.app_engine_http_target]
// set. Using
// [AppEngineHttpRequest][google.cloud.tasks.v2beta2.AppEngineHttpRequest]
// requires
// [`appengine.applications.get`](https://cloud.google.com/appengine/docs/admin-api/access-control)
// Google IAM permission for the project and the following scope:
// `https://www.googleapis.com/auth/cloud-platform` The task will be delivered
// to the App Engine app which belongs to the same project as the queue. For
// more information, see [How Requests are
// Routed](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed)
// and how routing is affected by [dispatch
// files](https://cloud.google.com/appengine/docs/python/config/dispatchref).
// Traffic is encrypted during transport and never leaves Google datacenters.
// Because this traffic is carried over a communication mechanism internal to
// Google, you cannot explicitly set the protocol (for example, HTTP or HTTPS).
// The request to the handler, however, will appear to have used the HTTP
// protocol. The
// [AppEngineRouting][google.cloud.tasks.v2beta2.AppEngineRouting] used to
// construct the URL that the task is delivered to can be set at the
// queue-level or task-level: - If set,
// [app_engine_routing_override][google.cloud.tasks.v2beta2.AppEngineHttpTarget.app_engine_routing_override]
// is used for all tasks in the queue, no matter what the setting is for the
// [task-level
// app_engine_routing][google.cloud.tasks.v2beta2.AppEngineHttpRequest.app_engine_routing].
// The `url` that the task will be sent to is: - `url =`
// [host][google.cloud.tasks.v2beta2.AppEngineRouting.host] `+`
// [relative_url][google.cloud.tasks.v2beta2.AppEngineHttpRequest.relative_url]
// Tasks can be dispatched to secure app handlers, unsecure app handlers, and
// URIs restricted with [`login:
// admin`](https://cloud.google.com/appengine/docs/standard/python/config/appref).
// Because tasks are not run as any user, they cannot be dispatched to URIs
// restricted with [`login:
// required`](https://cloud.google.com/appengine/docs/standard/python/config/appref)
// Task dispatches also do not follow redirects. The task attempt has succeeded
// if the app's request handler returns an HTTP response code in the range
// [`200` - `299`]. The task attempt has failed if the app's handler returns a
// non-2xx response code or Cloud Tasks does not receive response before the
// [deadline][Task.dispatch_deadline]. Failed tasks will be retried according
// to the [retry configuration][google.cloud.tasks.v2beta2.Queue.retry_config].
// `503` (Service Unavailable) is considered an App Engine system error instead
// of an application error and will cause Cloud Tasks' traffic congestion
// control to temporarily throttle the queue's dispatches. Unlike other types
// of task targets, a `429` (Too Many Requests) response from an app handler
// does not cause traffic congestion control to throttle the queue.
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
type AppEngineHttpRequest = src.AppEngineHttpRequest

// App Engine HTTP target. The task will be delivered to the App Engine
// application hostname specified by its
// [AppEngineHttpTarget][google.cloud.tasks.v2beta2.AppEngineHttpTarget] and
// [AppEngineHttpRequest][google.cloud.tasks.v2beta2.AppEngineHttpRequest]. The
// documentation for
// [AppEngineHttpRequest][google.cloud.tasks.v2beta2.AppEngineHttpRequest]
// explains how the task's host URL is constructed. Using
// [AppEngineHttpTarget][google.cloud.tasks.v2beta2.AppEngineHttpTarget]
// requires
// [`appengine.applications.get`](https://cloud.google.com/appengine/docs/admin-api/access-control)
// Google IAM permission for the project and the following scope:
// `https://www.googleapis.com/auth/cloud-platform`
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
type AppEngineHttpTarget = src.AppEngineHttpTarget

// App Engine Routing. Defines routing characteristics specific to App Engine
// - service, version, and instance. For more information about services,
// versions, and instances see [An Overview of App
// Engine](https://cloud.google.com/appengine/docs/python/an-overview-of-app-engine),
// [Microservices Architecture on Google App
// Engine](https://cloud.google.com/appengine/docs/python/microservices-on-app-engine),
// [App Engine Standard request
// routing](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed),
// and [App Engine Flex request
// routing](https://cloud.google.com/appengine/docs/flexible/python/how-requests-are-routed).
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
type AppEngineRouting = src.AppEngineRouting

// The status of a task attempt.
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
type AttemptStatus = src.AttemptStatus

// Request message for canceling a lease using
// [CancelLease][google.cloud.tasks.v2beta2.CloudTasks.CancelLease].
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
type CancelLeaseRequest = src.CancelLeaseRequest

// CloudTasksClient is the client API for CloudTasks service. For semantics
// around ctx use and closing/ending streaming RPCs, please refer to
// https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
type CloudTasksClient = src.CloudTasksClient

// CloudTasksServer is the server API for CloudTasks service.
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
type CloudTasksServer = src.CloudTasksServer

// Request message for
// [CreateQueue][google.cloud.tasks.v2beta2.CloudTasks.CreateQueue].
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
type CreateQueueRequest = src.CreateQueueRequest

// Request message for
// [CreateTask][google.cloud.tasks.v2beta2.CloudTasks.CreateTask].
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
type CreateTaskRequest = src.CreateTaskRequest

// Request message for
// [DeleteQueue][google.cloud.tasks.v2beta2.CloudTasks.DeleteQueue].
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
type DeleteQueueRequest = src.DeleteQueueRequest

// Request message for deleting a task using
// [DeleteTask][google.cloud.tasks.v2beta2.CloudTasks.DeleteTask].
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
type DeleteTaskRequest = src.DeleteTaskRequest

// Request message for
// [GetQueue][google.cloud.tasks.v2beta2.CloudTasks.GetQueue].
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
type GetQueueRequest = src.GetQueueRequest

// Request message for getting a task using
// [GetTask][google.cloud.tasks.v2beta2.CloudTasks.GetTask].
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
type GetTaskRequest = src.GetTaskRequest

// The HTTP method used to execute the task.
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
type HttpMethod = src.HttpMethod

// Request message for leasing tasks using
// [LeaseTasks][google.cloud.tasks.v2beta2.CloudTasks.LeaseTasks].
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
type LeaseTasksRequest = src.LeaseTasksRequest

// Response message for leasing tasks using
// [LeaseTasks][google.cloud.tasks.v2beta2.CloudTasks.LeaseTasks].
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
type LeaseTasksResponse = src.LeaseTasksResponse

// Request message for
// [ListQueues][google.cloud.tasks.v2beta2.CloudTasks.ListQueues].
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
type ListQueuesRequest = src.ListQueuesRequest

// Response message for
// [ListQueues][google.cloud.tasks.v2beta2.CloudTasks.ListQueues].
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
type ListQueuesResponse = src.ListQueuesResponse

// Request message for listing tasks using
// [ListTasks][google.cloud.tasks.v2beta2.CloudTasks.ListTasks].
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
type ListTasksRequest = src.ListTasksRequest

// Response message for listing tasks using
// [ListTasks][google.cloud.tasks.v2beta2.CloudTasks.ListTasks].
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
type ListTasksResponse = src.ListTasksResponse

// Request message for
// [PauseQueue][google.cloud.tasks.v2beta2.CloudTasks.PauseQueue].
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
type PauseQueueRequest = src.PauseQueueRequest

// The pull message contains data that can be used by the caller of
// [LeaseTasks][google.cloud.tasks.v2beta2.CloudTasks.LeaseTasks] to process
// the task. This proto can only be used for tasks in a queue which has
// [pull_target][google.cloud.tasks.v2beta2.Queue.pull_target] set.
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
type PullMessage = src.PullMessage

// Pull target.
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
type PullTarget = src.PullTarget

// Request message for
// [PurgeQueue][google.cloud.tasks.v2beta2.CloudTasks.PurgeQueue].
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
type PurgeQueueRequest = src.PurgeQueueRequest

// A queue is a container of related tasks. Queues are configured to manage
// how those tasks are dispatched. Configurable properties include rate limits,
// retry options, target types, and others.
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
type Queue = src.Queue

// Statistics for a queue.
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
type QueueStats = src.QueueStats
type Queue_AppEngineHttpTarget = src.Queue_AppEngineHttpTarget
type Queue_PullTarget = src.Queue_PullTarget

// State of the queue.
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
type Queue_State = src.Queue_State

// Rate limits. This message determines the maximum rate that tasks can be
// dispatched by a queue, regardless of whether the dispatch is a first task
// attempt or a retry. Note: The debugging command,
// [RunTask][google.cloud.tasks.v2beta2.CloudTasks.RunTask], will run a task
// even if the queue has reached its
// [RateLimits][google.cloud.tasks.v2beta2.RateLimits].
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
type RateLimits = src.RateLimits

// Request message for renewing a lease using
// [RenewLease][google.cloud.tasks.v2beta2.CloudTasks.RenewLease].
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
type RenewLeaseRequest = src.RenewLeaseRequest

// Request message for
// [ResumeQueue][google.cloud.tasks.v2beta2.CloudTasks.ResumeQueue].
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
type ResumeQueueRequest = src.ResumeQueueRequest

// Retry config. These settings determine how a failed task attempt is
// retried.
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
type RetryConfig = src.RetryConfig
type RetryConfig_MaxAttempts = src.RetryConfig_MaxAttempts
type RetryConfig_UnlimitedAttempts = src.RetryConfig_UnlimitedAttempts

// Request message for forcing a task to run now using
// [RunTask][google.cloud.tasks.v2beta2.CloudTasks.RunTask].
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
type RunTaskRequest = src.RunTaskRequest

// A unit of scheduled work.
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
type Task = src.Task

// Status of the task.
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
type TaskStatus = src.TaskStatus
type Task_AppEngineHttpRequest = src.Task_AppEngineHttpRequest
type Task_PullMessage = src.Task_PullMessage

// The view specifies a subset of [Task][google.cloud.tasks.v2beta2.Task]
// data. When a task is returned in a response, not all information is
// retrieved by default because some data, such as payloads, might be desirable
// to return only when needed because of its large size or because of the
// sensitivity of data that it contains.
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
type Task_View = src.Task_View

// UnimplementedCloudTasksServer can be embedded to have forward compatible
// implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
type UnimplementedCloudTasksServer = src.UnimplementedCloudTasksServer

// Request message for
// [UpdateQueue][google.cloud.tasks.v2beta2.CloudTasks.UpdateQueue].
//
// Deprecated: Please use types in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
type UpdateQueueRequest = src.UpdateQueueRequest

// Deprecated: Please use funcs in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
func NewCloudTasksClient(cc grpc.ClientConnInterface) CloudTasksClient {
	return src.NewCloudTasksClient(cc)
}

// Deprecated: Please use funcs in: cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb
func RegisterCloudTasksServer(s *grpc.Server, srv CloudTasksServer) {
	src.RegisterCloudTasksServer(s, srv)
}
