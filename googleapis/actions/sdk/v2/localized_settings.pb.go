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
// source: google/actions/sdk/v2/localized_settings.proto

package sdk

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

// Represents settings of an Actions project that are specific to a user locale.
// In this instance, user means the end user who invokes your Actions.
// **This message is localizable.**
type LocalizedSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The default display name for this Actions project (if there is no
	// translation available)
	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Required. The pronunciation of the display name to invoke it within a voice
	// (spoken) context.
	Pronunciation string `protobuf:"bytes,2,opt,name=pronunciation,proto3" json:"pronunciation,omitempty"`
	// Required. The default short description for the Actions project (if there is no
	// translation available). 80 character limit.
	ShortDescription string `protobuf:"bytes,3,opt,name=short_description,json=shortDescription,proto3" json:"short_description,omitempty"`
	// Required. The default long description for the Actions project (if there is no
	// translation available). 4000 character limit.
	FullDescription string `protobuf:"bytes,4,opt,name=full_description,json=fullDescription,proto3" json:"full_description,omitempty"`
	// Required. Small square image, 192 x 192 px.
	// This should be specified as a reference to the corresponding image in the
	// `resources/images/` directory. For example, `$resources.images.foo` (without the
	// extension) for image in `resources/images/foo.jpg`
	// When working on a project pulled from Console, the Google-managed URL
	// pulled could be used. URLs from external sources are not allowed.
	SmallLogoImage string `protobuf:"bytes,5,opt,name=small_logo_image,json=smallLogoImage,proto3" json:"small_logo_image,omitempty"`
	// Optional. Large landscape image, 1920 x 1080 px.
	// This should be specified as a reference to the corresponding image in the
	// `resources/images/` directory. For example, `$resources.images.foo` (without the
	// extension) for image in `resources/images/foo.jpg`
	// When working on a project pulled from Console, the Google-managed URL
	// pulled could be used. URLs from external sources are not allowed.
	LargeBannerImage string `protobuf:"bytes,6,opt,name=large_banner_image,json=largeBannerImage,proto3" json:"large_banner_image,omitempty"`
	// Required. The name of the developer to be displayed to users.
	DeveloperName string `protobuf:"bytes,7,opt,name=developer_name,json=developerName,proto3" json:"developer_name,omitempty"`
	// Required. The contact email address for the developer.
	DeveloperEmail string `protobuf:"bytes,8,opt,name=developer_email,json=developerEmail,proto3" json:"developer_email,omitempty"`
	// Optional. The terms of service URL.
	TermsOfServiceUrl string `protobuf:"bytes,9,opt,name=terms_of_service_url,json=termsOfServiceUrl,proto3" json:"terms_of_service_url,omitempty"`
	// Required. The Google Assistant voice type that users hear when they interact with
	// your Actions. The supported values are "male_1", "male_2", "female_1", and
	// "female_2".
	Voice string `protobuf:"bytes,10,opt,name=voice,proto3" json:"voice,omitempty"`
	// Optional. The locale for the specified voice. If not specified, this resolves
	// to the user's Assistant locale. If specified, the voice locale must have
	// the same root language as the locale specified in LocalizedSettings.
	VoiceLocale string `protobuf:"bytes,14,opt,name=voice_locale,json=voiceLocale,proto3" json:"voice_locale,omitempty"`
	// Required. The privacy policy URL.
	PrivacyPolicyUrl string `protobuf:"bytes,11,opt,name=privacy_policy_url,json=privacyPolicyUrl,proto3" json:"privacy_policy_url,omitempty"`
	// Optional. Sample invocation phrases displayed as part of your Actions project's
	// description in the Assistant directory. This will help users learn how to
	// use it.
	SampleInvocations []string `protobuf:"bytes,12,rep,name=sample_invocations,json=sampleInvocations,proto3" json:"sample_invocations,omitempty"`
	// Optional. Theme customizations for visual components of your Actions.
	ThemeCustomization *ThemeCustomization `protobuf:"bytes,13,opt,name=theme_customization,json=themeCustomization,proto3" json:"theme_customization,omitempty"`
}

func (x *LocalizedSettings) Reset() {
	*x = LocalizedSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_localized_settings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalizedSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalizedSettings) ProtoMessage() {}

func (x *LocalizedSettings) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_localized_settings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalizedSettings.ProtoReflect.Descriptor instead.
func (*LocalizedSettings) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_localized_settings_proto_rawDescGZIP(), []int{0}
}

func (x *LocalizedSettings) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *LocalizedSettings) GetPronunciation() string {
	if x != nil {
		return x.Pronunciation
	}
	return ""
}

func (x *LocalizedSettings) GetShortDescription() string {
	if x != nil {
		return x.ShortDescription
	}
	return ""
}

func (x *LocalizedSettings) GetFullDescription() string {
	if x != nil {
		return x.FullDescription
	}
	return ""
}

func (x *LocalizedSettings) GetSmallLogoImage() string {
	if x != nil {
		return x.SmallLogoImage
	}
	return ""
}

func (x *LocalizedSettings) GetLargeBannerImage() string {
	if x != nil {
		return x.LargeBannerImage
	}
	return ""
}

func (x *LocalizedSettings) GetDeveloperName() string {
	if x != nil {
		return x.DeveloperName
	}
	return ""
}

func (x *LocalizedSettings) GetDeveloperEmail() string {
	if x != nil {
		return x.DeveloperEmail
	}
	return ""
}

func (x *LocalizedSettings) GetTermsOfServiceUrl() string {
	if x != nil {
		return x.TermsOfServiceUrl
	}
	return ""
}

func (x *LocalizedSettings) GetVoice() string {
	if x != nil {
		return x.Voice
	}
	return ""
}

func (x *LocalizedSettings) GetVoiceLocale() string {
	if x != nil {
		return x.VoiceLocale
	}
	return ""
}

func (x *LocalizedSettings) GetPrivacyPolicyUrl() string {
	if x != nil {
		return x.PrivacyPolicyUrl
	}
	return ""
}

func (x *LocalizedSettings) GetSampleInvocations() []string {
	if x != nil {
		return x.SampleInvocations
	}
	return nil
}

func (x *LocalizedSettings) GetThemeCustomization() *ThemeCustomization {
	if x != nil {
		return x.ThemeCustomization
	}
	return nil
}

var File_google_actions_sdk_v2_localized_settings_proto protoreflect.FileDescriptor

var file_google_actions_sdk_v2_localized_settings_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65,
	0x64, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x1a, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x74,
	0x68, 0x65, 0x6d, 0x65, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x05, 0x0a, 0x11, 0x4c, 0x6f,
	0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x26, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x6e, 0x75,
	0x6e, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x6e, 0x75, 0x6e, 0x63, 0x69, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x11, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x10, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x10, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x0f, 0x66, 0x75, 0x6c, 0x6c, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x10, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x5f, 0x6c, 0x6f,
	0x67, 0x6f, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x0e, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x4c, 0x6f, 0x67, 0x6f, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x31, 0x0a, 0x12, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x5f, 0x62, 0x61, 0x6e,
	0x6e, 0x65, 0x72, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x01, 0x52, 0x10, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x42, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x0e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f,
	0x70, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x0d, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x0f, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x5f,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x0e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x34, 0x0a, 0x14, 0x74, 0x65, 0x72, 0x6d, 0x73, 0x5f, 0x6f, 0x66, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x01, 0x52, 0x11, 0x74, 0x65, 0x72, 0x6d, 0x73, 0x4f, 0x66, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x19, 0x0a, 0x05, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x12, 0x26, 0x0a, 0x0c, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0b, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x31, 0x0a, 0x12, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x63, 0x79, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x10, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x63, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x32, 0x0a, 0x12,
	0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x11, 0x73,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x5f, 0x0a, 0x13, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73,
	0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x12, 0x74,
	0x68, 0x65, 0x6d, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x6f, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x42, 0x16,
	0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x3b, 0x73,
	0x64, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_actions_sdk_v2_localized_settings_proto_rawDescOnce sync.Once
	file_google_actions_sdk_v2_localized_settings_proto_rawDescData = file_google_actions_sdk_v2_localized_settings_proto_rawDesc
)

func file_google_actions_sdk_v2_localized_settings_proto_rawDescGZIP() []byte {
	file_google_actions_sdk_v2_localized_settings_proto_rawDescOnce.Do(func() {
		file_google_actions_sdk_v2_localized_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_actions_sdk_v2_localized_settings_proto_rawDescData)
	})
	return file_google_actions_sdk_v2_localized_settings_proto_rawDescData
}

var file_google_actions_sdk_v2_localized_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_actions_sdk_v2_localized_settings_proto_goTypes = []interface{}{
	(*LocalizedSettings)(nil),  // 0: google.actions.sdk.v2.LocalizedSettings
	(*ThemeCustomization)(nil), // 1: google.actions.sdk.v2.ThemeCustomization
}
var file_google_actions_sdk_v2_localized_settings_proto_depIdxs = []int32{
	1, // 0: google.actions.sdk.v2.LocalizedSettings.theme_customization:type_name -> google.actions.sdk.v2.ThemeCustomization
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_actions_sdk_v2_localized_settings_proto_init() }
func file_google_actions_sdk_v2_localized_settings_proto_init() {
	if File_google_actions_sdk_v2_localized_settings_proto != nil {
		return
	}
	file_google_actions_sdk_v2_theme_customization_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_actions_sdk_v2_localized_settings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalizedSettings); i {
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
			RawDescriptor: file_google_actions_sdk_v2_localized_settings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_actions_sdk_v2_localized_settings_proto_goTypes,
		DependencyIndexes: file_google_actions_sdk_v2_localized_settings_proto_depIdxs,
		MessageInfos:      file_google_actions_sdk_v2_localized_settings_proto_msgTypes,
	}.Build()
	File_google_actions_sdk_v2_localized_settings_proto = out.File
	file_google_actions_sdk_v2_localized_settings_proto_rawDesc = nil
	file_google_actions_sdk_v2_localized_settings_proto_goTypes = nil
	file_google_actions_sdk_v2_localized_settings_proto_depIdxs = nil
}
