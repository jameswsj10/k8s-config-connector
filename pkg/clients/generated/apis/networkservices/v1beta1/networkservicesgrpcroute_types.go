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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type GrpcrouteAbort struct {
	/* The HTTP status code used to abort the request. The value must be between 200 and 599 inclusive. */
	// +optional
	HttpStatus *int `json:"httpStatus,omitempty"`

	/* The percentage of traffic which will be aborted. The value must be between [0, 100] */
	// +optional
	Percentage *int `json:"percentage,omitempty"`
}

type GrpcrouteAction struct {
	/* Optional. The destination services to which traffic should be forwarded. If multiple destinations are specified, traffic will be split between Backend Service(s) according to the weight field of these destinations. */
	// +optional
	Destinations []GrpcrouteDestinations `json:"destinations,omitempty"`

	/* Optional. The specification for fault injection introduced into traffic to test the resiliency of clients to destination service failure. As part of fault injection, when clients send requests to a destination, delays can be introduced on a percentage of requests before sending those requests to the destination service. Similarly requests from clients can be aborted by for a percentage of requests. timeout and retry_policy will be ignored by clients that are configured with a fault_injection_policy */
	// +optional
	FaultInjectionPolicy *GrpcrouteFaultInjectionPolicy `json:"faultInjectionPolicy,omitempty"`

	/* Optional. Specifies the retry policy associated with this route. */
	// +optional
	RetryPolicy *GrpcrouteRetryPolicy `json:"retryPolicy,omitempty"`

	/* Optional. Specifies the timeout for selected route. Timeout is computed from the time the request has been fully processed (i.e. end of stream) up until the response has been completely processed. Timeout includes all retries. */
	// +optional
	Timeout *string `json:"timeout,omitempty"`
}

type GrpcrouteDelay struct {
	/* Specify a fixed delay before forwarding the request. */
	// +optional
	FixedDelay *string `json:"fixedDelay,omitempty"`

	/* The percentage of traffic on which delay will be injected. The value must be between [0, 100] */
	// +optional
	Percentage *int `json:"percentage,omitempty"`
}

type GrpcrouteDestinations struct {
	/*  */
	ServiceRef v1alpha1.ResourceRef `json:"serviceRef"`

	/* Optional. Specifies the proportion of requests forwarded to the backend referenced by the serviceName field. This is computed as: weight/Sum(weights in this destination list). For non-zero values, there may be some epsilon from the exact proportion defined here depending on the precision an implementation supports. If only one serviceName is specified and it has a weight greater than 0, 100% of the traffic is forwarded to that backend. If weights are specified for any one service name, they need to be specified for all of them. If weights are unspecified for all services, then, traffic is distributed in equal proportions to all of them. */
	// +optional
	Weight *int `json:"weight,omitempty"`
}

type GrpcrouteFaultInjectionPolicy struct {
	/* The specification for aborting to client requests. */
	// +optional
	Abort *GrpcrouteAbort `json:"abort,omitempty"`

	/* The specification for injecting delay to client requests. */
	// +optional
	Delay *GrpcrouteDelay `json:"delay,omitempty"`
}

type GrpcrouteHeaders struct {
	/* Required. The key of the header. */
	Key string `json:"key"`

	/* Optional. Specifies how to match against the value of the header. If not specified, a default value of EXACT is used. Possible values: MATCH_TYPE_UNSPECIFIED, MATCH_ANY, MATCH_ALL */
	// +optional
	Type *string `json:"type,omitempty"`

	/* Required. The value of the header. */
	Value string `json:"value"`
}

type GrpcrouteMatches struct {
	/* Optional. Specifies a collection of headers to match. */
	// +optional
	Headers []GrpcrouteHeaders `json:"headers,omitempty"`

	/* Optional. A gRPC method to match against. If this field is empty or omitted, will match all methods. */
	// +optional
	Method *GrpcrouteMethod `json:"method,omitempty"`
}

type GrpcrouteMethod struct {
	/* Optional. Specifies that matches are case sensitive. The default value is true. case_sensitive must not be used with a type of REGULAR_EXPRESSION. */
	// +optional
	CaseSensitive *bool `json:"caseSensitive,omitempty"`

	/* Required. Name of the method to match against. If unspecified, will match all methods. */
	GrpcMethod string `json:"grpcMethod"`

	/* Required. Name of the service to match against. If unspecified, will match all services. */
	GrpcService string `json:"grpcService"`

	/* Optional. Specifies how to match against the name. If not specified, a default value of "EXACT" is used. Possible values: TYPE_UNSPECIFIED, EXACT, REGULAR_EXPRESSION */
	// +optional
	Type *string `json:"type,omitempty"`
}

type GrpcrouteRetryPolicy struct {
	/* Specifies the allowed number of retries. This number must be > 0. If not specpfied, default to 1. */
	// +optional
	NumRetries *int `json:"numRetries,omitempty"`

	/* - connect-failure: Router will retry on failures connecting to Backend Services, for example due to connection timeouts. - refused-stream: Router will retry if the backend service resets the stream with a REFUSED_STREAM error code. This reset type indicates that it is safe to retry. - cancelled: Router will retry if the gRPC status code in the response header is set to cancelled - deadline-exceeded: Router will retry if the gRPC status code in the response header is set to deadline-exceeded - resource-exhausted: Router will retry if the gRPC status code in the response header is set to resource-exhausted - unavailable: Router will retry if the gRPC status code in the response header is set to unavailable */
	// +optional
	RetryConditions []string `json:"retryConditions,omitempty"`
}

type GrpcrouteRules struct {
	/* Required. A detailed rule defining how to route traffic. This field is required. */
	Action GrpcrouteAction `json:"action"`

	/* Optional. Matches define conditions used for matching the rule against incoming gRPC requests. Each match is independent, i.e. this rule will be matched if ANY one of the matches is satisfied. If no matches field is specified, this rule will unconditionally match traffic. */
	// +optional
	Matches []GrpcrouteMatches `json:"matches,omitempty"`
}

type NetworkServicesGRPCRouteSpec struct {
	/* Optional. A free-text description of the resource. Max length 1024 characters. */
	// +optional
	Description *string `json:"description,omitempty"`

	/*  */
	// +optional
	Gateways []v1alpha1.ResourceRef `json:"gateways,omitempty"`

	/* Required. Service hostnames with an optional port for which this route describes traffic. Format: [:] Hostname is the fully qualified domain name of a network host. This matches the RFC 1123 definition of a hostname with 2 notable exceptions: - IPs are not allowed. - A hostname may be prefixed with a wildcard label (*.). The wildcard label must appear by itself as the first label. Hostname can be “precise” which is a domain name without the terminating dot of a network host (e.g. “foo.example.com”) or “wildcard”, which is a domain name prefixed with a single wildcard label (e.g. *.example.com). Note that as per RFC1035 and RFC1123, a label must consist of lower case alphanumeric characters or ‘-’, and must start and end with an alphanumeric character. No other punctuation is allowed. The routes associated with a Router must have unique hostnames. If you attempt to attach multiple routes with conflicting hostnames, the configuration will be rejected. For example, while it is acceptable for routes for the hostnames "*.foo.bar.com" and "*.bar.com" to be associated with the same route, it is not possible to associate two routes both with "*.bar.com" or both with "bar.com". In the case that multiple routes match the hostname, the most specific match will be selected. For example, "foo.bar.baz.com" will take precedence over "*.bar.baz.com" and "*.bar.baz.com" will take precedence over "*.baz.com". If a port is specified, then gRPC clients must use the channel URI with the port to match this rule (i.e. "xds:///service:123"), otherwise they must supply the URI without a port (i.e. "xds:///service"). */
	Hostnames []string `json:"hostnames"`

	/* Immutable. The location for the resource */
	Location string `json:"location"`

	/*  */
	// +optional
	Meshes []v1alpha1.ResourceRef `json:"meshes,omitempty"`

	/* Immutable. The Project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Required. A list of detailed rules defining how to route traffic. Within a single GrpcRoute, the GrpcRoute.RouteAction associated with the first matching GrpcRoute.RouteRule will be executed. At least one rule must be supplied. */
	Rules []GrpcrouteRules `json:"rules"`
}

type NetworkServicesGRPCRouteStatus struct {
	/* Conditions represent the latest available observations of the
	   NetworkServicesGRPCRoute's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Output only. The timestamp when the resource was created. */
	CreateTime string `json:"createTime,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration int `json:"observedGeneration,omitempty"`
	/* Output only. Server-defined URL of this resource */
	SelfLink string `json:"selfLink,omitempty"`
	/* Output only. The timestamp when the resource was updated. */
	UpdateTime string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetworkServicesGRPCRoute is the Schema for the networkservices API
// +k8s:openapi-gen=true
type NetworkServicesGRPCRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkServicesGRPCRouteSpec   `json:"spec,omitempty"`
	Status NetworkServicesGRPCRouteStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetworkServicesGRPCRouteList contains a list of NetworkServicesGRPCRoute
type NetworkServicesGRPCRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkServicesGRPCRoute `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkServicesGRPCRoute{}, &NetworkServicesGRPCRouteList{})
}
