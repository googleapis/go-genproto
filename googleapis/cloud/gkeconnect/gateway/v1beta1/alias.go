// Copyright 2024 Google LLC
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

// Package gateway aliases all exported identifiers in package
// "cloud.google.com/go/gkeconnect/gateway/apiv1beta1/gatewaypb".
//
// Deprecated: Please use types in: cloud.google.com/go/gkeconnect/gateway/apiv1beta1/gatewaypb.
// Please read https://github.com/googleapis/google-cloud-go/blob/main/migration.md
// for more details.
package gateway

import (
	src "cloud.google.com/go/gkeconnect/gateway/apiv1beta1/gatewaypb"
	grpc "google.golang.org/grpc"
)

// Deprecated: Please use consts in: cloud.google.com/go/gkeconnect/gateway/apiv1beta1/gatewaypb
const (
	GenerateCredentialsRequest_OPERATING_SYSTEM_UNSPECIFIED = src.GenerateCredentialsRequest_OPERATING_SYSTEM_UNSPECIFIED
	GenerateCredentialsRequest_OPERATING_SYSTEM_WINDOWS     = src.GenerateCredentialsRequest_OPERATING_SYSTEM_WINDOWS
)

// Deprecated: Please use vars in: cloud.google.com/go/gkeconnect/gateway/apiv1beta1/gatewaypb
var (
	File_google_cloud_gkeconnect_gateway_v1beta1_control_proto = src.File_google_cloud_gkeconnect_gateway_v1beta1_control_proto
	GenerateCredentialsRequest_OperatingSystem_name            = src.GenerateCredentialsRequest_OperatingSystem_name
	GenerateCredentialsRequest_OperatingSystem_value           = src.GenerateCredentialsRequest_OperatingSystem_value
)

// GatewayControlClient is the client API for GatewayControl service. For
// semantics around ctx use and closing/ending streaming RPCs, please refer to
// https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/gkeconnect/gateway/apiv1beta1/gatewaypb
type GatewayControlClient = src.GatewayControlClient

// GatewayControlServer is the server API for GatewayControl service.
//
// Deprecated: Please use types in: cloud.google.com/go/gkeconnect/gateway/apiv1beta1/gatewaypb
type GatewayControlServer = src.GatewayControlServer

// A request for connection information for a particular membership.
//
// Deprecated: Please use types in: cloud.google.com/go/gkeconnect/gateway/apiv1beta1/gatewaypb
type GenerateCredentialsRequest = src.GenerateCredentialsRequest

// Operating systems requiring specialized kubeconfigs.
//
// Deprecated: Please use types in: cloud.google.com/go/gkeconnect/gateway/apiv1beta1/gatewaypb
type GenerateCredentialsRequest_OperatingSystem = src.GenerateCredentialsRequest_OperatingSystem

// Connection information for a particular membership.
//
// Deprecated: Please use types in: cloud.google.com/go/gkeconnect/gateway/apiv1beta1/gatewaypb
type GenerateCredentialsResponse = src.GenerateCredentialsResponse

// UnimplementedGatewayControlServer can be embedded to have forward
// compatible implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/gkeconnect/gateway/apiv1beta1/gatewaypb
type UnimplementedGatewayControlServer = src.UnimplementedGatewayControlServer

// Deprecated: Please use funcs in: cloud.google.com/go/gkeconnect/gateway/apiv1beta1/gatewaypb
func NewGatewayControlClient(cc grpc.ClientConnInterface) GatewayControlClient {
	return src.NewGatewayControlClient(cc)
}

// Deprecated: Please use funcs in: cloud.google.com/go/gkeconnect/gateway/apiv1beta1/gatewaypb
func RegisterGatewayControlServer(s *grpc.Server, srv GatewayControlServer) {
	src.RegisterGatewayControlServer(s, srv)
}
