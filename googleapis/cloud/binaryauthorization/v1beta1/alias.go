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

// Package binaryauthorization aliases all exported identifiers in package
// "cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb".
//
// Deprecated: Please use types in: cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb.
// Please read https://github.com/googleapis/google-cloud-go/blob/main/migration.md
// for more details.
package binaryauthorization

import (
	src "cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb"
	grpc "google.golang.org/grpc"
)

// Deprecated: Please use consts in: cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb
const (
	AdmissionRule_ALWAYS_ALLOW                                                                    = src.AdmissionRule_ALWAYS_ALLOW
	AdmissionRule_ALWAYS_DENY                                                                     = src.AdmissionRule_ALWAYS_DENY
	AdmissionRule_DRYRUN_AUDIT_LOG_ONLY                                                           = src.AdmissionRule_DRYRUN_AUDIT_LOG_ONLY
	AdmissionRule_ENFORCED_BLOCK_AND_AUDIT_LOG                                                    = src.AdmissionRule_ENFORCED_BLOCK_AND_AUDIT_LOG
	AdmissionRule_ENFORCEMENT_MODE_UNSPECIFIED                                                    = src.AdmissionRule_ENFORCEMENT_MODE_UNSPECIFIED
	AdmissionRule_EVALUATION_MODE_UNSPECIFIED                                                     = src.AdmissionRule_EVALUATION_MODE_UNSPECIFIED
	AdmissionRule_REQUIRE_ATTESTATION                                                             = src.AdmissionRule_REQUIRE_ATTESTATION
	ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails_ALLOW                     = src.ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails_ALLOW
	ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails_AUDIT_RESULT_UNSPECIFIED  = src.ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails_AUDIT_RESULT_UNSPECIFIED
	ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails_DENY                      = src.ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails_DENY
	ContinuousValidationEvent_ContinuousValidationPodEvent_POLICY_CONFORMANCE_VERDICT_UNSPECIFIED = src.ContinuousValidationEvent_ContinuousValidationPodEvent_POLICY_CONFORMANCE_VERDICT_UNSPECIFIED
	ContinuousValidationEvent_ContinuousValidationPodEvent_VIOLATES_POLICY                        = src.ContinuousValidationEvent_ContinuousValidationPodEvent_VIOLATES_POLICY
	PkixPublicKey_ECDSA_P256_SHA256                                                               = src.PkixPublicKey_ECDSA_P256_SHA256
	PkixPublicKey_ECDSA_P384_SHA384                                                               = src.PkixPublicKey_ECDSA_P384_SHA384
	PkixPublicKey_ECDSA_P521_SHA512                                                               = src.PkixPublicKey_ECDSA_P521_SHA512
	PkixPublicKey_EC_SIGN_P256_SHA256                                                             = src.PkixPublicKey_EC_SIGN_P256_SHA256
	PkixPublicKey_EC_SIGN_P384_SHA384                                                             = src.PkixPublicKey_EC_SIGN_P384_SHA384
	PkixPublicKey_EC_SIGN_P521_SHA512                                                             = src.PkixPublicKey_EC_SIGN_P521_SHA512
	PkixPublicKey_RSA_PSS_2048_SHA256                                                             = src.PkixPublicKey_RSA_PSS_2048_SHA256
	PkixPublicKey_RSA_PSS_3072_SHA256                                                             = src.PkixPublicKey_RSA_PSS_3072_SHA256
	PkixPublicKey_RSA_PSS_4096_SHA256                                                             = src.PkixPublicKey_RSA_PSS_4096_SHA256
	PkixPublicKey_RSA_PSS_4096_SHA512                                                             = src.PkixPublicKey_RSA_PSS_4096_SHA512
	PkixPublicKey_RSA_SIGN_PKCS1_2048_SHA256                                                      = src.PkixPublicKey_RSA_SIGN_PKCS1_2048_SHA256
	PkixPublicKey_RSA_SIGN_PKCS1_3072_SHA256                                                      = src.PkixPublicKey_RSA_SIGN_PKCS1_3072_SHA256
	PkixPublicKey_RSA_SIGN_PKCS1_4096_SHA256                                                      = src.PkixPublicKey_RSA_SIGN_PKCS1_4096_SHA256
	PkixPublicKey_RSA_SIGN_PKCS1_4096_SHA512                                                      = src.PkixPublicKey_RSA_SIGN_PKCS1_4096_SHA512
	PkixPublicKey_SIGNATURE_ALGORITHM_UNSPECIFIED                                                 = src.PkixPublicKey_SIGNATURE_ALGORITHM_UNSPECIFIED
	Policy_DISABLE                                                                                = src.Policy_DISABLE
	Policy_ENABLE                                                                                 = src.Policy_ENABLE
	Policy_GLOBAL_POLICY_EVALUATION_MODE_UNSPECIFIED                                              = src.Policy_GLOBAL_POLICY_EVALUATION_MODE_UNSPECIFIED
)

// Deprecated: Please use vars in: cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb
var (
	AdmissionRule_EnforcementMode_name                                                    = src.AdmissionRule_EnforcementMode_name
	AdmissionRule_EnforcementMode_value                                                   = src.AdmissionRule_EnforcementMode_value
	AdmissionRule_EvaluationMode_name                                                     = src.AdmissionRule_EvaluationMode_name
	AdmissionRule_EvaluationMode_value                                                    = src.AdmissionRule_EvaluationMode_value
	ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails_AuditResult_name  = src.ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails_AuditResult_name
	ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails_AuditResult_value = src.ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails_AuditResult_value
	ContinuousValidationEvent_ContinuousValidationPodEvent_PolicyConformanceVerdict_name  = src.ContinuousValidationEvent_ContinuousValidationPodEvent_PolicyConformanceVerdict_name
	ContinuousValidationEvent_ContinuousValidationPodEvent_PolicyConformanceVerdict_value = src.ContinuousValidationEvent_ContinuousValidationPodEvent_PolicyConformanceVerdict_value
	File_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto     = src.File_google_cloud_binaryauthorization_v1beta1_continuous_validation_logging_proto
	File_google_cloud_binaryauthorization_v1beta1_resources_proto                         = src.File_google_cloud_binaryauthorization_v1beta1_resources_proto
	File_google_cloud_binaryauthorization_v1beta1_service_proto                           = src.File_google_cloud_binaryauthorization_v1beta1_service_proto
	PkixPublicKey_SignatureAlgorithm_name                                                 = src.PkixPublicKey_SignatureAlgorithm_name
	PkixPublicKey_SignatureAlgorithm_value                                                = src.PkixPublicKey_SignatureAlgorithm_value
	Policy_GlobalPolicyEvaluationMode_name                                                = src.Policy_GlobalPolicyEvaluationMode_name
	Policy_GlobalPolicyEvaluationMode_value                                               = src.Policy_GlobalPolicyEvaluationMode_value
)

// An [admission rule][google.cloud.binaryauthorization.v1beta1.AdmissionRule]
// specifies either that all container images used in a pod creation request
// must be attested to by one or more
// [attestors][google.cloud.binaryauthorization.v1beta1.Attestor], that all pod
// creations will be allowed, or that all pod creations will be denied. Images
// matching an [admission allowlist
// pattern][google.cloud.binaryauthorization.v1beta1.AdmissionWhitelistPattern]
// are exempted from admission rules and will never block a pod creation.
//
// Deprecated: Please use types in: cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb
type AdmissionRule = src.AdmissionRule

// Defines the possible actions when a pod creation is denied by an admission
// rule.
//
// Deprecated: Please use types in: cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb
type AdmissionRule_EnforcementMode = src.AdmissionRule_EnforcementMode
type AdmissionRule_EvaluationMode = src.AdmissionRule_EvaluationMode

// An [admission allowlist
// pattern][google.cloud.binaryauthorization.v1beta1.AdmissionWhitelistPattern]
// exempts images from checks by [admission
// rules][google.cloud.binaryauthorization.v1beta1.AdmissionRule].
//
// Deprecated: Please use types in: cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb
type AdmissionWhitelistPattern = src.AdmissionWhitelistPattern

// An [attestor][google.cloud.binaryauthorization.v1beta1.Attestor] that
// attests to container image artifacts. An existing attestor cannot be
// modified except where indicated.
//
// Deprecated: Please use types in: cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb
type Attestor = src.Attestor

// An [attestor public
// key][google.cloud.binaryauthorization.v1beta1.AttestorPublicKey] that will
// be used to verify attestations signed by this attestor.
//
// Deprecated: Please use types in: cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb
type AttestorPublicKey = src.AttestorPublicKey
type AttestorPublicKey_AsciiArmoredPgpPublicKey = src.AttestorPublicKey_AsciiArmoredPgpPublicKey
type AttestorPublicKey_PkixPublicKey = src.AttestorPublicKey_PkixPublicKey
type Attestor_UserOwnedDrydockNote = src.Attestor_UserOwnedDrydockNote

// BinauthzManagementServiceV1Beta1Client is the client API for
// BinauthzManagementServiceV1Beta1 service. For semantics around ctx use and
// closing/ending streaming RPCs, please refer to
// https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb
type BinauthzManagementServiceV1Beta1Client = src.BinauthzManagementServiceV1Beta1Client

// BinauthzManagementServiceV1Beta1Server is the server API for
// BinauthzManagementServiceV1Beta1 service.
//
// Deprecated: Please use types in: cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb
type BinauthzManagementServiceV1Beta1Server = src.BinauthzManagementServiceV1Beta1Server

// Represents an auditing event from Continuous Validation.
//
// Deprecated: Please use types in: cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb
type ContinuousValidationEvent = src.ContinuousValidationEvent

// An auditing event for one Pod.
//
// Deprecated: Please use types in: cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb
type ContinuousValidationEvent_ContinuousValidationPodEvent = src.ContinuousValidationEvent_ContinuousValidationPodEvent

// Container image with auditing details.
//
// Deprecated: Please use types in: cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb
type ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails = src.ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails

// Result of the audit.
//
// Deprecated: Please use types in: cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb
type ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails_AuditResult = src.ContinuousValidationEvent_ContinuousValidationPodEvent_ImageDetails_AuditResult

// Audit time policy conformance verdict.
//
// Deprecated: Please use types in: cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb
type ContinuousValidationEvent_ContinuousValidationPodEvent_PolicyConformanceVerdict = src.ContinuousValidationEvent_ContinuousValidationPodEvent_PolicyConformanceVerdict
type ContinuousValidationEvent_PodEvent = src.ContinuousValidationEvent_PodEvent

// Request message for [BinauthzManagementService.CreateAttestor][].
//
// Deprecated: Please use types in: cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb
type CreateAttestorRequest = src.CreateAttestorRequest

// Request message for [BinauthzManagementService.DeleteAttestor][].
//
// Deprecated: Please use types in: cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb
type DeleteAttestorRequest = src.DeleteAttestorRequest

// Request message for [BinauthzManagementService.GetAttestor][].
//
// Deprecated: Please use types in: cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb
type GetAttestorRequest = src.GetAttestorRequest

// Request message for [BinauthzManagementService.GetPolicy][].
//
// Deprecated: Please use types in: cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb
type GetPolicyRequest = src.GetPolicyRequest

// Request to read the current system policy.
//
// Deprecated: Please use types in: cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb
type GetSystemPolicyRequest = src.GetSystemPolicyRequest

// Request message for [BinauthzManagementService.ListAttestors][].
//
// Deprecated: Please use types in: cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb
type ListAttestorsRequest = src.ListAttestorsRequest

// Response message for [BinauthzManagementService.ListAttestors][].
//
// Deprecated: Please use types in: cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb
type ListAttestorsResponse = src.ListAttestorsResponse

// A public key in the PkixPublicKey format (see
// https://tools.ietf.org/html/rfc5280#section-4.1.2.7 for details). Public
// keys of this type are typically textually encoded using the PEM format.
//
// Deprecated: Please use types in: cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb
type PkixPublicKey = src.PkixPublicKey

// Represents a signature algorithm and other information necessary to verify
// signatures with a given public key. This is based primarily on the public
// key types supported by Tink's PemKeyType, which is in turn based on KMS's
// supported signing algorithms. See
// https://cloud.google.com/kms/docs/algorithms. In the future, BinAuthz might
// support additional public key types independently of Tink and/or KMS.
//
// Deprecated: Please use types in: cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb
type PkixPublicKey_SignatureAlgorithm = src.PkixPublicKey_SignatureAlgorithm

// A [policy][google.cloud.binaryauthorization.v1beta1.Policy] for Binary
// Authorization.
//
// Deprecated: Please use types in: cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb
type Policy = src.Policy
type Policy_GlobalPolicyEvaluationMode = src.Policy_GlobalPolicyEvaluationMode

// SystemPolicyV1Beta1Client is the client API for SystemPolicyV1Beta1
// service. For semantics around ctx use and closing/ending streaming RPCs,
// please refer to
// https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb
type SystemPolicyV1Beta1Client = src.SystemPolicyV1Beta1Client

// SystemPolicyV1Beta1Server is the server API for SystemPolicyV1Beta1
// service.
//
// Deprecated: Please use types in: cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb
type SystemPolicyV1Beta1Server = src.SystemPolicyV1Beta1Server

// UnimplementedBinauthzManagementServiceV1Beta1Server can be embedded to have
// forward compatible implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb
type UnimplementedBinauthzManagementServiceV1Beta1Server = src.UnimplementedBinauthzManagementServiceV1Beta1Server

// UnimplementedSystemPolicyV1Beta1Server can be embedded to have forward
// compatible implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb
type UnimplementedSystemPolicyV1Beta1Server = src.UnimplementedSystemPolicyV1Beta1Server

// Request message for [BinauthzManagementService.UpdateAttestor][].
//
// Deprecated: Please use types in: cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb
type UpdateAttestorRequest = src.UpdateAttestorRequest

// Request message for [BinauthzManagementService.UpdatePolicy][].
//
// Deprecated: Please use types in: cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb
type UpdatePolicyRequest = src.UpdatePolicyRequest

// An [user owned drydock
// note][google.cloud.binaryauthorization.v1beta1.UserOwnedDrydockNote]
// references a Drydock ATTESTATION_AUTHORITY Note created by the user.
//
// Deprecated: Please use types in: cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb
type UserOwnedDrydockNote = src.UserOwnedDrydockNote

// Deprecated: Please use funcs in: cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb
func NewBinauthzManagementServiceV1Beta1Client(cc grpc.ClientConnInterface) BinauthzManagementServiceV1Beta1Client {
	return src.NewBinauthzManagementServiceV1Beta1Client(cc)
}

// Deprecated: Please use funcs in: cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb
func NewSystemPolicyV1Beta1Client(cc grpc.ClientConnInterface) SystemPolicyV1Beta1Client {
	return src.NewSystemPolicyV1Beta1Client(cc)
}

// Deprecated: Please use funcs in: cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb
func RegisterBinauthzManagementServiceV1Beta1Server(s *grpc.Server, srv BinauthzManagementServiceV1Beta1Server) {
	src.RegisterBinauthzManagementServiceV1Beta1Server(s, srv)
}

// Deprecated: Please use funcs in: cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb
func RegisterSystemPolicyV1Beta1Server(s *grpc.Server, srv SystemPolicyV1Beta1Server) {
	src.RegisterSystemPolicyV1Beta1Server(s, srv)
}
