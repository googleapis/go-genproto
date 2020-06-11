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
// 	protoc-gen-go v1.22.0
// 	protoc        v3.12.3
// source: google/maps/routes/v1/compute_routes_request.proto

package routes

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// A set of values used to specify the mode of travel.
type RouteTravelMode int32

const (
	// No travel mode specified. Defaults to `DRIVE`.
	RouteTravelMode_TRAVEL_MODE_UNSPECIFIED RouteTravelMode = 0
	// Travel by passenger car.
	RouteTravelMode_DRIVE RouteTravelMode = 1
	// Travel by bicycle.
	RouteTravelMode_BICYCLE RouteTravelMode = 2
	// Travel by walking.
	RouteTravelMode_WALK RouteTravelMode = 3
	// Two-wheeled, motorized vehicle. For example, motorcycle. Note that this
	// differs from the `BICYCLE` travel mode which covers human-powered mode.
	RouteTravelMode_TWO_WHEELER RouteTravelMode = 4
)

// Enum value maps for RouteTravelMode.
var (
	RouteTravelMode_name = map[int32]string{
		0: "TRAVEL_MODE_UNSPECIFIED",
		1: "DRIVE",
		2: "BICYCLE",
		3: "WALK",
		4: "TWO_WHEELER",
	}
	RouteTravelMode_value = map[string]int32{
		"TRAVEL_MODE_UNSPECIFIED": 0,
		"DRIVE":                   1,
		"BICYCLE":                 2,
		"WALK":                    3,
		"TWO_WHEELER":             4,
	}
)

func (x RouteTravelMode) Enum() *RouteTravelMode {
	p := new(RouteTravelMode)
	*p = x
	return p
}

func (x RouteTravelMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RouteTravelMode) Descriptor() protoreflect.EnumDescriptor {
	return file_google_maps_routes_v1_compute_routes_request_proto_enumTypes[0].Descriptor()
}

func (RouteTravelMode) Type() protoreflect.EnumType {
	return &file_google_maps_routes_v1_compute_routes_request_proto_enumTypes[0]
}

func (x RouteTravelMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RouteTravelMode.Descriptor instead.
func (RouteTravelMode) EnumDescriptor() ([]byte, []int) {
	return file_google_maps_routes_v1_compute_routes_request_proto_rawDescGZIP(), []int{0}
}

// A set of values that specify factors to take into consideration when
// calculating the route.
type RoutingPreference int32

const (
	// No routing preference specified. Default to `TRAFFIC_AWARE`.
	RoutingPreference_ROUTING_PREFERENCE_UNSPECIFIED RoutingPreference = 0
	// Computes routes without taking traffic conditions into consideration.
	// Suitable when traffic conditions don't matter. Using this value produces
	// the lowest latency.
	RoutingPreference_TRAFFIC_UNAWARE RoutingPreference = 1
	// Calculates routes taking traffic conditions into consideration. In contrast
	// to `TRAFFIC_AWARE_OPTIMAL`, some optimizations are applied to significantly
	// reduce latency.
	RoutingPreference_TRAFFIC_AWARE RoutingPreference = 2
	// Calculates the routes taking traffic conditions into consideration,
	// without applying most performance optimizations. Using this value produces
	// the highest latency.
	RoutingPreference_TRAFFIC_AWARE_OPTIMAL RoutingPreference = 3
)

// Enum value maps for RoutingPreference.
var (
	RoutingPreference_name = map[int32]string{
		0: "ROUTING_PREFERENCE_UNSPECIFIED",
		1: "TRAFFIC_UNAWARE",
		2: "TRAFFIC_AWARE",
		3: "TRAFFIC_AWARE_OPTIMAL",
	}
	RoutingPreference_value = map[string]int32{
		"ROUTING_PREFERENCE_UNSPECIFIED": 0,
		"TRAFFIC_UNAWARE":                1,
		"TRAFFIC_AWARE":                  2,
		"TRAFFIC_AWARE_OPTIMAL":          3,
	}
)

func (x RoutingPreference) Enum() *RoutingPreference {
	p := new(RoutingPreference)
	*p = x
	return p
}

func (x RoutingPreference) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RoutingPreference) Descriptor() protoreflect.EnumDescriptor {
	return file_google_maps_routes_v1_compute_routes_request_proto_enumTypes[1].Descriptor()
}

func (RoutingPreference) Type() protoreflect.EnumType {
	return &file_google_maps_routes_v1_compute_routes_request_proto_enumTypes[1]
}

func (x RoutingPreference) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RoutingPreference.Descriptor instead.
func (RoutingPreference) EnumDescriptor() ([]byte, []int) {
	return file_google_maps_routes_v1_compute_routes_request_proto_rawDescGZIP(), []int{1}
}

// A set of values that specify the unit of measure used in the display.
type Units int32

const (
	// Units of measure not specified. Defaults to the unit of measure inferred
	// from the request.
	Units_UNITS_UNSPECIFIED Units = 0
	// Metric units of measure.
	Units_METRIC Units = 1
	// Imperial (English) units of measure.
	Units_IMPERIAL Units = 2
)

// Enum value maps for Units.
var (
	Units_name = map[int32]string{
		0: "UNITS_UNSPECIFIED",
		1: "METRIC",
		2: "IMPERIAL",
	}
	Units_value = map[string]int32{
		"UNITS_UNSPECIFIED": 0,
		"METRIC":            1,
		"IMPERIAL":          2,
	}
)

func (x Units) Enum() *Units {
	p := new(Units)
	*p = x
	return p
}

func (x Units) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Units) Descriptor() protoreflect.EnumDescriptor {
	return file_google_maps_routes_v1_compute_routes_request_proto_enumTypes[2].Descriptor()
}

func (Units) Type() protoreflect.EnumType {
	return &file_google_maps_routes_v1_compute_routes_request_proto_enumTypes[2]
}

func (x Units) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Units.Descriptor instead.
func (Units) EnumDescriptor() ([]byte, []int) {
	return file_google_maps_routes_v1_compute_routes_request_proto_rawDescGZIP(), []int{2}
}

// ComputeRoutes request message.
type ComputeRoutesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Origin waypoint.
	Origin *Waypoint `protobuf:"bytes,1,opt,name=origin,proto3" json:"origin,omitempty"`
	// Required. Destination waypoint.
	Destination *Waypoint `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
	// Optional. A set of waypoints along the route (excluding terminal points),
	// for either stopping at or passing by.
	Intermediates []*Waypoint `protobuf:"bytes,3,rep,name=intermediates,proto3" json:"intermediates,omitempty"`
	// Optional. Specifies the mode of transportation.
	TravelMode RouteTravelMode `protobuf:"varint,4,opt,name=travel_mode,json=travelMode,proto3,enum=google.maps.routes.v1.RouteTravelMode" json:"travel_mode,omitempty"`
	// Optional. Specifies how to compute the route. The server
	// attempts to use the selected routing preference to compute the route. If
	//  the routing preference results in an error or an extra long latency, then
	// an error is returned. In the future, we might implement a fallback
	// mechanism to use a different option when the preferred option does not give
	// a valid result. You can specify this option only when the `travel_mode` is
	// `DRIVE` or `TWO_WHEELER`, otherwise the request fails.
	RoutingPreference RoutingPreference `protobuf:"varint,5,opt,name=routing_preference,json=routingPreference,proto3,enum=google.maps.routes.v1.RoutingPreference" json:"routing_preference,omitempty"`
	// Optional. Specifies your preference for the quality of the polyline.
	PolylineQuality PolylineQuality `protobuf:"varint,6,opt,name=polyline_quality,json=polylineQuality,proto3,enum=google.maps.routes.v1.PolylineQuality" json:"polyline_quality,omitempty"`
	// Optional. The departure time. If you don't set this value, then this value
	// defaults to the time that you made the request. If you set this value to a
	// time that has already occurred, then the request fails.
	DepartureTime *timestamp.Timestamp `protobuf:"bytes,7,opt,name=departure_time,json=departureTime,proto3" json:"departure_time,omitempty"`
	// Specifies whether to calculate alternate routes in addition to the route.
	ComputeAlternativeRoutes bool `protobuf:"varint,8,opt,name=compute_alternative_routes,json=computeAlternativeRoutes,proto3" json:"compute_alternative_routes,omitempty"`
	// Optional. A set of conditions to satisfy that affect the way routes are
	// calculated.
	RouteModifiers *RouteModifiers `protobuf:"bytes,9,opt,name=route_modifiers,json=routeModifiers,proto3" json:"route_modifiers,omitempty"`
	// Optional. The BCP-47 language code, such as "en-US" or "sr-Latn". For more
	// information, see
	// http://www.unicode.org/reports/tr35/#Unicode_locale_identifier. See
	// [Language Support](https://developers.google.com/maps/faq#languagesupport)
	// for the list of supported languages. When you don't provide this value, the
	// display language is inferred from the location of the route request.
	LanguageCode string `protobuf:"bytes,10,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	// Optional. Specifies the units of measure for the display fields. This
	// includes the `instruction` field in `NavigationInstruction`. The units of
	// measure used for the route, leg, step distance, and duration are not
	// affected by this value. If you don't provide this value, then the display
	// units are inferred from the location of the request.
	Units Units `protobuf:"varint,11,opt,name=units,proto3,enum=google.maps.routes.v1.Units" json:"units,omitempty"`
}

func (x *ComputeRoutesRequest) Reset() {
	*x = ComputeRoutesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_routes_v1_compute_routes_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeRoutesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeRoutesRequest) ProtoMessage() {}

func (x *ComputeRoutesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_routes_v1_compute_routes_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeRoutesRequest.ProtoReflect.Descriptor instead.
func (*ComputeRoutesRequest) Descriptor() ([]byte, []int) {
	return file_google_maps_routes_v1_compute_routes_request_proto_rawDescGZIP(), []int{0}
}

func (x *ComputeRoutesRequest) GetOrigin() *Waypoint {
	if x != nil {
		return x.Origin
	}
	return nil
}

func (x *ComputeRoutesRequest) GetDestination() *Waypoint {
	if x != nil {
		return x.Destination
	}
	return nil
}

func (x *ComputeRoutesRequest) GetIntermediates() []*Waypoint {
	if x != nil {
		return x.Intermediates
	}
	return nil
}

func (x *ComputeRoutesRequest) GetTravelMode() RouteTravelMode {
	if x != nil {
		return x.TravelMode
	}
	return RouteTravelMode_TRAVEL_MODE_UNSPECIFIED
}

func (x *ComputeRoutesRequest) GetRoutingPreference() RoutingPreference {
	if x != nil {
		return x.RoutingPreference
	}
	return RoutingPreference_ROUTING_PREFERENCE_UNSPECIFIED
}

func (x *ComputeRoutesRequest) GetPolylineQuality() PolylineQuality {
	if x != nil {
		return x.PolylineQuality
	}
	return PolylineQuality_POLYLINE_QUALITY_UNSPECIFIED
}

func (x *ComputeRoutesRequest) GetDepartureTime() *timestamp.Timestamp {
	if x != nil {
		return x.DepartureTime
	}
	return nil
}

func (x *ComputeRoutesRequest) GetComputeAlternativeRoutes() bool {
	if x != nil {
		return x.ComputeAlternativeRoutes
	}
	return false
}

func (x *ComputeRoutesRequest) GetRouteModifiers() *RouteModifiers {
	if x != nil {
		return x.RouteModifiers
	}
	return nil
}

func (x *ComputeRoutesRequest) GetLanguageCode() string {
	if x != nil {
		return x.LanguageCode
	}
	return ""
}

func (x *ComputeRoutesRequest) GetUnits() Units {
	if x != nil {
		return x.Units
	}
	return Units_UNITS_UNSPECIFIED
}

// Encapsulates a set of optional conditions to satisfy when calculating the
// routes.
type RouteModifiers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies whether to avoid toll roads. Applies only to the `DRIVE` and
	// `TWO_WHEELER` travel modes.
	AvoidTolls bool `protobuf:"varint,1,opt,name=avoid_tolls,json=avoidTolls,proto3" json:"avoid_tolls,omitempty"`
	// Specifies whether to avoid highways. Applies only to the `DRIVE` and
	// `TWO_WHEELER` travel modes.
	AvoidHighways bool `protobuf:"varint,2,opt,name=avoid_highways,json=avoidHighways,proto3" json:"avoid_highways,omitempty"`
	// Specifies whether to avoid ferries. Applies only to the `DRIVE` and
	// `TWO_WHEELER` travel modes.
	AvoidFerries bool `protobuf:"varint,3,opt,name=avoid_ferries,json=avoidFerries,proto3" json:"avoid_ferries,omitempty"`
	// Specifies whether to avoid navigating indoors. Applies only to the `WALK`
	// travel mode.
	AvoidIndoor bool `protobuf:"varint,4,opt,name=avoid_indoor,json=avoidIndoor,proto3" json:"avoid_indoor,omitempty"`
	// Optional. Specifies the vehicle information.
	VehicleInfo *VehicleInfo `protobuf:"bytes,5,opt,name=vehicle_info,json=vehicleInfo,proto3" json:"vehicle_info,omitempty"`
	// Encapsulates information about toll passes.
	// TollPass is unset means no available pass.
	// Applies only to the DRIVE and TWO_WHEELER travel modes.
	TollPasses []TollPass `protobuf:"varint,6,rep,packed,name=toll_passes,json=tollPasses,proto3,enum=google.maps.routes.v1.TollPass" json:"toll_passes,omitempty"`
}

func (x *RouteModifiers) Reset() {
	*x = RouteModifiers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_routes_v1_compute_routes_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteModifiers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteModifiers) ProtoMessage() {}

func (x *RouteModifiers) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_routes_v1_compute_routes_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteModifiers.ProtoReflect.Descriptor instead.
func (*RouteModifiers) Descriptor() ([]byte, []int) {
	return file_google_maps_routes_v1_compute_routes_request_proto_rawDescGZIP(), []int{1}
}

func (x *RouteModifiers) GetAvoidTolls() bool {
	if x != nil {
		return x.AvoidTolls
	}
	return false
}

func (x *RouteModifiers) GetAvoidHighways() bool {
	if x != nil {
		return x.AvoidHighways
	}
	return false
}

func (x *RouteModifiers) GetAvoidFerries() bool {
	if x != nil {
		return x.AvoidFerries
	}
	return false
}

func (x *RouteModifiers) GetAvoidIndoor() bool {
	if x != nil {
		return x.AvoidIndoor
	}
	return false
}

func (x *RouteModifiers) GetVehicleInfo() *VehicleInfo {
	if x != nil {
		return x.VehicleInfo
	}
	return nil
}

func (x *RouteModifiers) GetTollPasses() []TollPass {
	if x != nil {
		return x.TollPasses
	}
	return nil
}

// Encapsulates the vehicle information, such as the license plate last
// character.
type VehicleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies the license plate last character. Could be a digit or a letter.
	LicensePlateLastCharacter string `protobuf:"bytes,1,opt,name=license_plate_last_character,json=licensePlateLastCharacter,proto3" json:"license_plate_last_character,omitempty"`
	// Describes the vehicle's emission type.
	// Applies only to the DRIVE travel mode.
	EmissionType VehicleEmissionType `protobuf:"varint,2,opt,name=emission_type,json=emissionType,proto3,enum=google.maps.routes.v1.VehicleEmissionType" json:"emission_type,omitempty"`
}

func (x *VehicleInfo) Reset() {
	*x = VehicleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_routes_v1_compute_routes_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VehicleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VehicleInfo) ProtoMessage() {}

func (x *VehicleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_routes_v1_compute_routes_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VehicleInfo.ProtoReflect.Descriptor instead.
func (*VehicleInfo) Descriptor() ([]byte, []int) {
	return file_google_maps_routes_v1_compute_routes_request_proto_rawDescGZIP(), []int{2}
}

func (x *VehicleInfo) GetLicensePlateLastCharacter() string {
	if x != nil {
		return x.LicensePlateLastCharacter
	}
	return ""
}

func (x *VehicleInfo) GetEmissionType() VehicleEmissionType {
	if x != nil {
		return x.EmissionType
	}
	return VehicleEmissionType_VEHICLE_EMISSION_TYPE_UNSPECIFIED
}

var File_google_maps_routes_v1_compute_routes_request_proto protoreflect.FileDescriptor

var file_google_maps_routes_v1_compute_routes_request_proto_rawDesc = []byte{
	0x0a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x5f,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70,
	0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x24, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x6c, 0x6c, 0x5f, 0x70, 0x61,
	0x73, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x65, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf8, 0x05, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a,
	0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x06,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x41, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x0d, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x73,
	0x12, 0x47, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d,
	0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x74,
	0x72, 0x61, 0x76, 0x65, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x57, 0x0a, 0x12, 0x72, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d,
	0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52,
	0x11, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x51, 0x0a, 0x10, 0x70, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x71,
	0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x51, 0x75, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x52, 0x0f, 0x70, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x51, 0x75,
	0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x41, 0x0a, 0x0e, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x75,
	0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x64, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x75, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x70,
	0x75, 0x74, 0x65, 0x5f, 0x61, 0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18, 0x63, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x4e, 0x0a, 0x0f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f,
	0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x52, 0x0e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x75,
	0x6e, 0x69, 0x74, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x52, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x22,
	0xa9, 0x02, 0x0a, 0x0e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x76, 0x6f, 0x69, 0x64, 0x5f, 0x74, 0x6f, 0x6c, 0x6c,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x61, 0x76, 0x6f, 0x69, 0x64, 0x54, 0x6f,
	0x6c, 0x6c, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x76, 0x6f, 0x69, 0x64, 0x5f, 0x68, 0x69, 0x67,
	0x68, 0x77, 0x61, 0x79, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x61, 0x76, 0x6f,
	0x69, 0x64, 0x48, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x76,
	0x6f, 0x69, 0x64, 0x5f, 0x66, 0x65, 0x72, 0x72, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0c, 0x61, 0x76, 0x6f, 0x69, 0x64, 0x46, 0x65, 0x72, 0x72, 0x69, 0x65, 0x73, 0x12,
	0x21, 0x0a, 0x0c, 0x61, 0x76, 0x6f, 0x69, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x6f, 0x6f, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x61, 0x76, 0x6f, 0x69, 0x64, 0x49, 0x6e, 0x64, 0x6f,
	0x6f, 0x72, 0x12, 0x45, 0x0a, 0x0c, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x76, 0x65,
	0x68, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x40, 0x0a, 0x0b, 0x74, 0x6f, 0x6c,
	0x6c, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x6c, 0x6c, 0x50, 0x61, 0x73, 0x73, 0x52,
	0x0a, 0x74, 0x6f, 0x6c, 0x6c, 0x50, 0x61, 0x73, 0x73, 0x65, 0x73, 0x22, 0x9f, 0x01, 0x0a, 0x0b,
	0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3f, 0x0a, 0x1c, 0x6c,
	0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x19, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x50, 0x6c, 0x61, 0x74, 0x65, 0x4c,
	0x61, 0x73, 0x74, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x12, 0x4f, 0x0a, 0x0d,
	0x65, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70,
	0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x68, 0x69,
	0x63, 0x6c, 0x65, 0x45, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0c, 0x65, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x2a, 0x61, 0x0a,
	0x0f, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x4d, 0x6f, 0x64, 0x65,
	0x12, 0x1b, 0x0a, 0x17, 0x54, 0x52, 0x41, 0x56, 0x45, 0x4c, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a,
	0x05, 0x44, 0x52, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x49, 0x43, 0x59,
	0x43, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x41, 0x4c, 0x4b, 0x10, 0x03, 0x12,
	0x0f, 0x0a, 0x0b, 0x54, 0x57, 0x4f, 0x5f, 0x57, 0x48, 0x45, 0x45, 0x4c, 0x45, 0x52, 0x10, 0x04,
	0x2a, 0x7a, 0x0a, 0x11, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x1e, 0x52, 0x4f, 0x55, 0x54, 0x49, 0x4e, 0x47,
	0x5f, 0x50, 0x52, 0x45, 0x46, 0x45, 0x52, 0x45, 0x4e, 0x43, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x54, 0x52, 0x41,
	0x46, 0x46, 0x49, 0x43, 0x5f, 0x55, 0x4e, 0x41, 0x57, 0x41, 0x52, 0x45, 0x10, 0x01, 0x12, 0x11,
	0x0a, 0x0d, 0x54, 0x52, 0x41, 0x46, 0x46, 0x49, 0x43, 0x5f, 0x41, 0x57, 0x41, 0x52, 0x45, 0x10,
	0x02, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x52, 0x41, 0x46, 0x46, 0x49, 0x43, 0x5f, 0x41, 0x57, 0x41,
	0x52, 0x45, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4d, 0x41, 0x4c, 0x10, 0x03, 0x2a, 0x38, 0x0a, 0x05,
	0x55, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x4e, 0x49, 0x54, 0x53, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4d, 0x50, 0x45,
	0x52, 0x49, 0x41, 0x4c, 0x10, 0x02, 0x42, 0xaf, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x42, 0x19, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x3b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0xf8, 0x01,
	0x01, 0xa2, 0x02, 0x04, 0x47, 0x4d, 0x52, 0x53, 0xaa, 0x02, 0x15, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x4d, 0x61, 0x70, 0x73, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x15, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x4d, 0x61, 0x70, 0x73, 0x5c, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_maps_routes_v1_compute_routes_request_proto_rawDescOnce sync.Once
	file_google_maps_routes_v1_compute_routes_request_proto_rawDescData = file_google_maps_routes_v1_compute_routes_request_proto_rawDesc
)

func file_google_maps_routes_v1_compute_routes_request_proto_rawDescGZIP() []byte {
	file_google_maps_routes_v1_compute_routes_request_proto_rawDescOnce.Do(func() {
		file_google_maps_routes_v1_compute_routes_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_maps_routes_v1_compute_routes_request_proto_rawDescData)
	})
	return file_google_maps_routes_v1_compute_routes_request_proto_rawDescData
}

var file_google_maps_routes_v1_compute_routes_request_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_google_maps_routes_v1_compute_routes_request_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_maps_routes_v1_compute_routes_request_proto_goTypes = []interface{}{
	(RouteTravelMode)(0),         // 0: google.maps.routes.v1.RouteTravelMode
	(RoutingPreference)(0),       // 1: google.maps.routes.v1.RoutingPreference
	(Units)(0),                   // 2: google.maps.routes.v1.Units
	(*ComputeRoutesRequest)(nil), // 3: google.maps.routes.v1.ComputeRoutesRequest
	(*RouteModifiers)(nil),       // 4: google.maps.routes.v1.RouteModifiers
	(*VehicleInfo)(nil),          // 5: google.maps.routes.v1.VehicleInfo
	(*Waypoint)(nil),             // 6: google.maps.routes.v1.Waypoint
	(PolylineQuality)(0),         // 7: google.maps.routes.v1.PolylineQuality
	(*timestamp.Timestamp)(nil),  // 8: google.protobuf.Timestamp
	(TollPass)(0),                // 9: google.maps.routes.v1.TollPass
	(VehicleEmissionType)(0),     // 10: google.maps.routes.v1.VehicleEmissionType
}
var file_google_maps_routes_v1_compute_routes_request_proto_depIdxs = []int32{
	6,  // 0: google.maps.routes.v1.ComputeRoutesRequest.origin:type_name -> google.maps.routes.v1.Waypoint
	6,  // 1: google.maps.routes.v1.ComputeRoutesRequest.destination:type_name -> google.maps.routes.v1.Waypoint
	6,  // 2: google.maps.routes.v1.ComputeRoutesRequest.intermediates:type_name -> google.maps.routes.v1.Waypoint
	0,  // 3: google.maps.routes.v1.ComputeRoutesRequest.travel_mode:type_name -> google.maps.routes.v1.RouteTravelMode
	1,  // 4: google.maps.routes.v1.ComputeRoutesRequest.routing_preference:type_name -> google.maps.routes.v1.RoutingPreference
	7,  // 5: google.maps.routes.v1.ComputeRoutesRequest.polyline_quality:type_name -> google.maps.routes.v1.PolylineQuality
	8,  // 6: google.maps.routes.v1.ComputeRoutesRequest.departure_time:type_name -> google.protobuf.Timestamp
	4,  // 7: google.maps.routes.v1.ComputeRoutesRequest.route_modifiers:type_name -> google.maps.routes.v1.RouteModifiers
	2,  // 8: google.maps.routes.v1.ComputeRoutesRequest.units:type_name -> google.maps.routes.v1.Units
	5,  // 9: google.maps.routes.v1.RouteModifiers.vehicle_info:type_name -> google.maps.routes.v1.VehicleInfo
	9,  // 10: google.maps.routes.v1.RouteModifiers.toll_passes:type_name -> google.maps.routes.v1.TollPass
	10, // 11: google.maps.routes.v1.VehicleInfo.emission_type:type_name -> google.maps.routes.v1.VehicleEmissionType
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_google_maps_routes_v1_compute_routes_request_proto_init() }
func file_google_maps_routes_v1_compute_routes_request_proto_init() {
	if File_google_maps_routes_v1_compute_routes_request_proto != nil {
		return
	}
	file_google_maps_routes_v1_polyline_proto_init()
	file_google_maps_routes_v1_toll_passes_proto_init()
	file_google_maps_routes_v1_vehicle_emission_type_proto_init()
	file_google_maps_routes_v1_waypoint_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_maps_routes_v1_compute_routes_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeRoutesRequest); i {
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
		file_google_maps_routes_v1_compute_routes_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteModifiers); i {
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
		file_google_maps_routes_v1_compute_routes_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VehicleInfo); i {
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
			RawDescriptor: file_google_maps_routes_v1_compute_routes_request_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_maps_routes_v1_compute_routes_request_proto_goTypes,
		DependencyIndexes: file_google_maps_routes_v1_compute_routes_request_proto_depIdxs,
		EnumInfos:         file_google_maps_routes_v1_compute_routes_request_proto_enumTypes,
		MessageInfos:      file_google_maps_routes_v1_compute_routes_request_proto_msgTypes,
	}.Build()
	File_google_maps_routes_v1_compute_routes_request_proto = out.File
	file_google_maps_routes_v1_compute_routes_request_proto_rawDesc = nil
	file_google_maps_routes_v1_compute_routes_request_proto_goTypes = nil
	file_google_maps_routes_v1_compute_routes_request_proto_depIdxs = nil
}
