// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ComputeTargetHTTPSProxySpec struct {
	/* Immutable. An optional description of this resource. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Location represents the geographical location of the ComputeTargetHTTPSProxy. Specify a region name or "global" for global resources. Reference: GCP definition of regions/zones (https://cloud.google.com/compute/docs/regions-zones/) */
	Location string `json:"location"`

	/* Immutable. This field only applies when the forwarding rule that references
	this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED. */
	// +optional
	ProxyBind *bool `json:"proxyBind,omitempty"`

	/* Specifies the QUIC override policy for this resource. This determines
	whether the load balancer will attempt to negotiate QUIC with clients
	or not. Can specify one of NONE, ENABLE, or DISABLE. If NONE is
	specified, uses the QUIC policy with no user overrides, which is
	equivalent to DISABLE. Default value: "NONE" Possible values: ["NONE", "ENABLE", "DISABLE"]. */
	// +optional
	QuicOverride *string `json:"quicOverride,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/*  */
	SslCertificates []v1alpha1.ResourceRef `json:"sslCertificates"`

	/* A reference to the ComputeSSLPolicy resource that will be
	associated with the ComputeTargetHTTPSProxy resource. If not set,
	the ComputeTargetHTTPSProxy resource will not have any SSL policy
	configured. */
	// +optional
	SslPolicyRef *v1alpha1.ResourceRef `json:"sslPolicyRef,omitempty"`

	/* A reference to the ComputeURLMap resource that defines the mapping
	from URL to the BackendService. */
	UrlMapRef v1alpha1.ResourceRef `json:"urlMapRef"`
}

type ComputeTargetHTTPSProxyStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeTargetHTTPSProxy's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Creation timestamp in RFC3339 text format. */
	CreationTimestamp string `json:"creationTimestamp,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration int `json:"observedGeneration,omitempty"`
	/* The unique identifier for the resource. */
	ProxyId int `json:"proxyId,omitempty"`
	/*  */
	SelfLink string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeTargetHTTPSProxy is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeTargetHTTPSProxy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeTargetHTTPSProxySpec   `json:"spec,omitempty"`
	Status ComputeTargetHTTPSProxyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeTargetHTTPSProxyList contains a list of ComputeTargetHTTPSProxy
type ComputeTargetHTTPSProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeTargetHTTPSProxy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeTargetHTTPSProxy{}, &ComputeTargetHTTPSProxyList{})
}
