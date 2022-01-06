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
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: google/appengine/v1/application.proto

package appengine

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Application_ServingStatus int32

const (
	// Serving status is unspecified.
	Application_UNSPECIFIED Application_ServingStatus = 0
	// Application is serving.
	Application_SERVING Application_ServingStatus = 1
	// Application has been disabled by the user.
	Application_USER_DISABLED Application_ServingStatus = 2
	// Application has been disabled by the system.
	Application_SYSTEM_DISABLED Application_ServingStatus = 3
)

// Enum value maps for Application_ServingStatus.
var (
	Application_ServingStatus_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "SERVING",
		2: "USER_DISABLED",
		3: "SYSTEM_DISABLED",
	}
	Application_ServingStatus_value = map[string]int32{
		"UNSPECIFIED":     0,
		"SERVING":         1,
		"USER_DISABLED":   2,
		"SYSTEM_DISABLED": 3,
	}
)

func (x Application_ServingStatus) Enum() *Application_ServingStatus {
	p := new(Application_ServingStatus)
	*p = x
	return p
}

func (x Application_ServingStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Application_ServingStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_google_appengine_v1_application_proto_enumTypes[0].Descriptor()
}

func (Application_ServingStatus) Type() protoreflect.EnumType {
	return &file_google_appengine_v1_application_proto_enumTypes[0]
}

func (x Application_ServingStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Application_ServingStatus.Descriptor instead.
func (Application_ServingStatus) EnumDescriptor() ([]byte, []int) {
	return file_google_appengine_v1_application_proto_rawDescGZIP(), []int{0, 0}
}

type Application_DatabaseType int32

const (
	// Database type is unspecified.
	Application_DATABASE_TYPE_UNSPECIFIED Application_DatabaseType = 0
	// Cloud Datastore
	Application_CLOUD_DATASTORE Application_DatabaseType = 1
	// Cloud Firestore Native
	Application_CLOUD_FIRESTORE Application_DatabaseType = 2
	// Cloud Firestore in Datastore Mode
	Application_CLOUD_DATASTORE_COMPATIBILITY Application_DatabaseType = 3
)

// Enum value maps for Application_DatabaseType.
var (
	Application_DatabaseType_name = map[int32]string{
		0: "DATABASE_TYPE_UNSPECIFIED",
		1: "CLOUD_DATASTORE",
		2: "CLOUD_FIRESTORE",
		3: "CLOUD_DATASTORE_COMPATIBILITY",
	}
	Application_DatabaseType_value = map[string]int32{
		"DATABASE_TYPE_UNSPECIFIED":     0,
		"CLOUD_DATASTORE":               1,
		"CLOUD_FIRESTORE":               2,
		"CLOUD_DATASTORE_COMPATIBILITY": 3,
	}
)

func (x Application_DatabaseType) Enum() *Application_DatabaseType {
	p := new(Application_DatabaseType)
	*p = x
	return p
}

func (x Application_DatabaseType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Application_DatabaseType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_appengine_v1_application_proto_enumTypes[1].Descriptor()
}

func (Application_DatabaseType) Type() protoreflect.EnumType {
	return &file_google_appengine_v1_application_proto_enumTypes[1]
}

func (x Application_DatabaseType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Application_DatabaseType.Descriptor instead.
func (Application_DatabaseType) EnumDescriptor() ([]byte, []int) {
	return file_google_appengine_v1_application_proto_rawDescGZIP(), []int{0, 1}
}

// An Application resource contains the top-level configuration of an App
// Engine application.
type Application struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Full path to the Application resource in the API.
	// Example: `apps/myapp`.
	//
	// @OutputOnly
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Identifier of the Application resource. This identifier is equivalent
	// to the project ID of the Google Cloud Platform project where you want to
	// deploy your application.
	// Example: `myapp`.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// HTTP path dispatch rules for requests to the application that do not
	// explicitly target a service or version. Rules are order-dependent.
	// Up to 20 dispatch rules can be supported.
	DispatchRules []*UrlDispatchRule `protobuf:"bytes,3,rep,name=dispatch_rules,json=dispatchRules,proto3" json:"dispatch_rules,omitempty"`
	// Google Apps authentication domain that controls which users can access
	// this application.
	//
	// Defaults to open access for any Google Account.
	AuthDomain string `protobuf:"bytes,6,opt,name=auth_domain,json=authDomain,proto3" json:"auth_domain,omitempty"`
	// Location from which this application runs. Application instances
	// run out of the data centers in the specified location, which is also where
	// all of the application's end user content is stored.
	//
	// Defaults to `us-central`.
	//
	// View the list of
	// [supported locations](https://cloud.google.com/appengine/docs/locations).
	LocationId string `protobuf:"bytes,7,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
	// Google Cloud Storage bucket that can be used for storing files
	// associated with this application. This bucket is associated with the
	// application and can be used by the gcloud deployment commands.
	//
	// @OutputOnly
	CodeBucket string `protobuf:"bytes,8,opt,name=code_bucket,json=codeBucket,proto3" json:"code_bucket,omitempty"`
	// Cookie expiration policy for this application.
	DefaultCookieExpiration *durationpb.Duration `protobuf:"bytes,9,opt,name=default_cookie_expiration,json=defaultCookieExpiration,proto3" json:"default_cookie_expiration,omitempty"`
	// Serving status of this application.
	ServingStatus Application_ServingStatus `protobuf:"varint,10,opt,name=serving_status,json=servingStatus,proto3,enum=google.appengine.v1.Application_ServingStatus" json:"serving_status,omitempty"`
	// Hostname used to reach this application, as resolved by App Engine.
	//
	// @OutputOnly
	DefaultHostname string `protobuf:"bytes,11,opt,name=default_hostname,json=defaultHostname,proto3" json:"default_hostname,omitempty"`
	// Google Cloud Storage bucket that can be used by this application to store
	// content.
	//
	// @OutputOnly
	DefaultBucket string                          `protobuf:"bytes,12,opt,name=default_bucket,json=defaultBucket,proto3" json:"default_bucket,omitempty"`
	Iap           *Application_IdentityAwareProxy `protobuf:"bytes,14,opt,name=iap,proto3" json:"iap,omitempty"`
	// The Google Container Registry domain used for storing managed build docker
	// images for this application.
	GcrDomain string `protobuf:"bytes,16,opt,name=gcr_domain,json=gcrDomain,proto3" json:"gcr_domain,omitempty"`
	// The type of the Cloud Firestore or Cloud Datastore database associated with
	// this application.
	DatabaseType Application_DatabaseType `protobuf:"varint,17,opt,name=database_type,json=databaseType,proto3,enum=google.appengine.v1.Application_DatabaseType" json:"database_type,omitempty"`
	// The feature specific settings to be used in the application.
	FeatureSettings *Application_FeatureSettings `protobuf:"bytes,18,opt,name=feature_settings,json=featureSettings,proto3" json:"feature_settings,omitempty"`
}

func (x *Application) Reset() {
	*x = Application{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_appengine_v1_application_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Application) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Application) ProtoMessage() {}

func (x *Application) ProtoReflect() protoreflect.Message {
	mi := &file_google_appengine_v1_application_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Application.ProtoReflect.Descriptor instead.
func (*Application) Descriptor() ([]byte, []int) {
	return file_google_appengine_v1_application_proto_rawDescGZIP(), []int{0}
}

func (x *Application) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Application) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Application) GetDispatchRules() []*UrlDispatchRule {
	if x != nil {
		return x.DispatchRules
	}
	return nil
}

func (x *Application) GetAuthDomain() string {
	if x != nil {
		return x.AuthDomain
	}
	return ""
}

func (x *Application) GetLocationId() string {
	if x != nil {
		return x.LocationId
	}
	return ""
}

func (x *Application) GetCodeBucket() string {
	if x != nil {
		return x.CodeBucket
	}
	return ""
}

func (x *Application) GetDefaultCookieExpiration() *durationpb.Duration {
	if x != nil {
		return x.DefaultCookieExpiration
	}
	return nil
}

func (x *Application) GetServingStatus() Application_ServingStatus {
	if x != nil {
		return x.ServingStatus
	}
	return Application_UNSPECIFIED
}

func (x *Application) GetDefaultHostname() string {
	if x != nil {
		return x.DefaultHostname
	}
	return ""
}

func (x *Application) GetDefaultBucket() string {
	if x != nil {
		return x.DefaultBucket
	}
	return ""
}

func (x *Application) GetIap() *Application_IdentityAwareProxy {
	if x != nil {
		return x.Iap
	}
	return nil
}

func (x *Application) GetGcrDomain() string {
	if x != nil {
		return x.GcrDomain
	}
	return ""
}

func (x *Application) GetDatabaseType() Application_DatabaseType {
	if x != nil {
		return x.DatabaseType
	}
	return Application_DATABASE_TYPE_UNSPECIFIED
}

func (x *Application) GetFeatureSettings() *Application_FeatureSettings {
	if x != nil {
		return x.FeatureSettings
	}
	return nil
}

// Rules to match an HTTP request and dispatch that request to a service.
type UrlDispatchRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Domain name to match against. The wildcard "`*`" is supported if
	// specified before a period: "`*.`".
	//
	// Defaults to matching all domains: "`*`".
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	// Pathname within the host. Must start with a "`/`". A
	// single "`*`" can be included at the end of the path.
	//
	// The sum of the lengths of the domain and path may not
	// exceed 100 characters.
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	// Resource ID of a service in this application that should
	// serve the matched request. The service must already
	// exist. Example: `default`.
	Service string `protobuf:"bytes,3,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *UrlDispatchRule) Reset() {
	*x = UrlDispatchRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_appengine_v1_application_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UrlDispatchRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UrlDispatchRule) ProtoMessage() {}

func (x *UrlDispatchRule) ProtoReflect() protoreflect.Message {
	mi := &file_google_appengine_v1_application_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UrlDispatchRule.ProtoReflect.Descriptor instead.
func (*UrlDispatchRule) Descriptor() ([]byte, []int) {
	return file_google_appengine_v1_application_proto_rawDescGZIP(), []int{1}
}

func (x *UrlDispatchRule) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *UrlDispatchRule) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *UrlDispatchRule) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

// Identity-Aware Proxy
type Application_IdentityAwareProxy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether the serving infrastructure will authenticate and
	// authorize all incoming requests.
	//
	// If true, the `oauth2_client_id` and `oauth2_client_secret`
	// fields must be non-empty.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// OAuth2 client ID to use for the authentication flow.
	Oauth2ClientId string `protobuf:"bytes,2,opt,name=oauth2_client_id,json=oauth2ClientId,proto3" json:"oauth2_client_id,omitempty"`
	// OAuth2 client secret to use for the authentication flow.
	//
	// For security reasons, this value cannot be retrieved via the API.
	// Instead, the SHA-256 hash of the value is returned in the
	// `oauth2_client_secret_sha256` field.
	//
	// @InputOnly
	Oauth2ClientSecret string `protobuf:"bytes,3,opt,name=oauth2_client_secret,json=oauth2ClientSecret,proto3" json:"oauth2_client_secret,omitempty"`
	// Hex-encoded SHA-256 hash of the client secret.
	//
	// @OutputOnly
	Oauth2ClientSecretSha256 string `protobuf:"bytes,4,opt,name=oauth2_client_secret_sha256,json=oauth2ClientSecretSha256,proto3" json:"oauth2_client_secret_sha256,omitempty"`
}

func (x *Application_IdentityAwareProxy) Reset() {
	*x = Application_IdentityAwareProxy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_appengine_v1_application_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Application_IdentityAwareProxy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Application_IdentityAwareProxy) ProtoMessage() {}

func (x *Application_IdentityAwareProxy) ProtoReflect() protoreflect.Message {
	mi := &file_google_appengine_v1_application_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Application_IdentityAwareProxy.ProtoReflect.Descriptor instead.
func (*Application_IdentityAwareProxy) Descriptor() ([]byte, []int) {
	return file_google_appengine_v1_application_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Application_IdentityAwareProxy) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Application_IdentityAwareProxy) GetOauth2ClientId() string {
	if x != nil {
		return x.Oauth2ClientId
	}
	return ""
}

func (x *Application_IdentityAwareProxy) GetOauth2ClientSecret() string {
	if x != nil {
		return x.Oauth2ClientSecret
	}
	return ""
}

func (x *Application_IdentityAwareProxy) GetOauth2ClientSecretSha256() string {
	if x != nil {
		return x.Oauth2ClientSecretSha256
	}
	return ""
}

// The feature specific settings to be used in the application. These define
// behaviors that are user configurable.
type Application_FeatureSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Boolean value indicating if split health checks should be used instead
	// of the legacy health checks. At an app.yaml level, this means defaulting
	// to 'readiness_check' and 'liveness_check' values instead of
	// 'health_check' ones. Once the legacy 'health_check' behavior is
	// deprecated, and this value is always true, this setting can
	// be removed.
	SplitHealthChecks bool `protobuf:"varint,1,opt,name=split_health_checks,json=splitHealthChecks,proto3" json:"split_health_checks,omitempty"`
	// If true, use [Container-Optimized OS](https://cloud.google.com/container-optimized-os/)
	// base image for VMs, rather than a base Debian image.
	UseContainerOptimizedOs bool `protobuf:"varint,2,opt,name=use_container_optimized_os,json=useContainerOptimizedOs,proto3" json:"use_container_optimized_os,omitempty"`
}

func (x *Application_FeatureSettings) Reset() {
	*x = Application_FeatureSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_appengine_v1_application_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Application_FeatureSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Application_FeatureSettings) ProtoMessage() {}

func (x *Application_FeatureSettings) ProtoReflect() protoreflect.Message {
	mi := &file_google_appengine_v1_application_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Application_FeatureSettings.ProtoReflect.Descriptor instead.
func (*Application_FeatureSettings) Descriptor() ([]byte, []int) {
	return file_google_appengine_v1_application_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Application_FeatureSettings) GetSplitHealthChecks() bool {
	if x != nil {
		return x.SplitHealthChecks
	}
	return false
}

func (x *Application_FeatureSettings) GetUseContainerOptimizedOs() bool {
	if x != nil {
		return x.UseContainerOptimizedOs
	}
	return false
}

var File_google_appengine_v1_application_proto protoreflect.FileDescriptor

var file_google_appengine_v1_application_proto_rawDesc = []byte{
	0x0a, 0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x0a, 0x0a, 0x0b, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x4b,
	0x0a, 0x0e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x72, 0x6c,
	0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x0d, 0x64, 0x69,
	0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1f, 0x0a, 0x0b,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x55,
	0x0a, 0x19, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65,
	0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x17, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x45, 0x78, 0x70, 0x69, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x55, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0d, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x29, 0x0a, 0x10,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x48,
	0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x45,
	0x0a, 0x03, 0x69, 0x61, 0x70, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41, 0x77, 0x61, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79,
	0x52, 0x03, 0x69, 0x61, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x63, 0x72, 0x5f, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x63, 0x72, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x12, 0x52, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x5b, 0x0a, 0x10, 0x66, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x12, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x52, 0x0f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x1a, 0xc9, 0x01, 0x0a, 0x12, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x41, 0x77, 0x61, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32,
	0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x30, 0x0a, 0x14, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12,
	0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x12, 0x3d, 0x0a, 0x1b, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x5f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x73, 0x68, 0x61, 0x32, 0x35,
	0x36, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x53, 0x68, 0x61, 0x32, 0x35,
	0x36, 0x1a, 0x7e, 0x0a, 0x0f, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x5f, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x11, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x73, 0x12, 0x3b, 0x0a, 0x1a, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x5f,
	0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x75, 0x73, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x65, 0x64, 0x4f,
	0x73, 0x22, 0x55, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x45, 0x52, 0x56, 0x49, 0x4e, 0x47, 0x10, 0x01,
	0x12, 0x11, 0x0a, 0x0d, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45,
	0x44, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x5f, 0x44, 0x49,
	0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x22, 0x7a, 0x0a, 0x0c, 0x44, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x44, 0x41, 0x54, 0x41,
	0x42, 0x41, 0x53, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4c, 0x4f, 0x55, 0x44,
	0x5f, 0x44, 0x41, 0x54, 0x41, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f,
	0x43, 0x4c, 0x4f, 0x55, 0x44, 0x5f, 0x46, 0x49, 0x52, 0x45, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x10,
	0x02, 0x12, 0x21, 0x0a, 0x1d, 0x43, 0x4c, 0x4f, 0x55, 0x44, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x53,
	0x54, 0x4f, 0x52, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x54, 0x49, 0x42, 0x49, 0x4c, 0x49,
	0x54, 0x59, 0x10, 0x03, 0x22, 0x57, 0x0a, 0x0f, 0x55, 0x72, 0x6c, 0x44, 0x69, 0x73, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0xc2, 0x01,
	0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x41, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76,
	0x31, 0x3b, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0xaa, 0x02, 0x19, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x70, 0x70, 0x45, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x5c, 0x56, 0x31, 0xea, 0x02, 0x1c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_appengine_v1_application_proto_rawDescOnce sync.Once
	file_google_appengine_v1_application_proto_rawDescData = file_google_appengine_v1_application_proto_rawDesc
)

func file_google_appengine_v1_application_proto_rawDescGZIP() []byte {
	file_google_appengine_v1_application_proto_rawDescOnce.Do(func() {
		file_google_appengine_v1_application_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_appengine_v1_application_proto_rawDescData)
	})
	return file_google_appengine_v1_application_proto_rawDescData
}

var file_google_appengine_v1_application_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_appengine_v1_application_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_appengine_v1_application_proto_goTypes = []interface{}{
	(Application_ServingStatus)(0),         // 0: google.appengine.v1.Application.ServingStatus
	(Application_DatabaseType)(0),          // 1: google.appengine.v1.Application.DatabaseType
	(*Application)(nil),                    // 2: google.appengine.v1.Application
	(*UrlDispatchRule)(nil),                // 3: google.appengine.v1.UrlDispatchRule
	(*Application_IdentityAwareProxy)(nil), // 4: google.appengine.v1.Application.IdentityAwareProxy
	(*Application_FeatureSettings)(nil),    // 5: google.appengine.v1.Application.FeatureSettings
	(*durationpb.Duration)(nil),            // 6: google.protobuf.Duration
}
var file_google_appengine_v1_application_proto_depIdxs = []int32{
	3, // 0: google.appengine.v1.Application.dispatch_rules:type_name -> google.appengine.v1.UrlDispatchRule
	6, // 1: google.appengine.v1.Application.default_cookie_expiration:type_name -> google.protobuf.Duration
	0, // 2: google.appengine.v1.Application.serving_status:type_name -> google.appengine.v1.Application.ServingStatus
	4, // 3: google.appengine.v1.Application.iap:type_name -> google.appengine.v1.Application.IdentityAwareProxy
	1, // 4: google.appengine.v1.Application.database_type:type_name -> google.appengine.v1.Application.DatabaseType
	5, // 5: google.appengine.v1.Application.feature_settings:type_name -> google.appengine.v1.Application.FeatureSettings
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_appengine_v1_application_proto_init() }
func file_google_appengine_v1_application_proto_init() {
	if File_google_appengine_v1_application_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_appengine_v1_application_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Application); i {
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
		file_google_appengine_v1_application_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UrlDispatchRule); i {
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
		file_google_appengine_v1_application_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Application_IdentityAwareProxy); i {
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
		file_google_appengine_v1_application_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Application_FeatureSettings); i {
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
			RawDescriptor: file_google_appengine_v1_application_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_appengine_v1_application_proto_goTypes,
		DependencyIndexes: file_google_appengine_v1_application_proto_depIdxs,
		EnumInfos:         file_google_appengine_v1_application_proto_enumTypes,
		MessageInfos:      file_google_appengine_v1_application_proto_msgTypes,
	}.Build()
	File_google_appengine_v1_application_proto = out.File
	file_google_appengine_v1_application_proto_rawDesc = nil
	file_google_appengine_v1_application_proto_goTypes = nil
	file_google_appengine_v1_application_proto_depIdxs = nil
}
