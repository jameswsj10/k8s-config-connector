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

type BackendserviceBackend struct {
	/* Specifies the balancing mode for this backend.

	For global HTTP(S) or TCP/SSL load balancing, the default is
	UTILIZATION. Valid values are UTILIZATION, RATE (for HTTP(S))
	and CONNECTION (for TCP/SSL). Default value: "UTILIZATION" Possible values: ["UTILIZATION", "RATE", "CONNECTION"]. */
	// +optional
	BalancingMode *string `json:"balancingMode,omitempty"`

	/* A multiplier applied to the group's maximum servicing capacity
	(based on UTILIZATION, RATE or CONNECTION).

	Default value is 1, which means the group will serve up to 100%
	of its configured capacity (depending on balancingMode). A
	setting of 0 means the group is completely drained, offering
	0% of its available Capacity. Valid range is [0.0,1.0]. */
	// +optional
	CapacityScaler *float64 `json:"capacityScaler,omitempty"`

	/* An optional description of this resource.
	Provide this property when you create the resource. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* This field designates whether this is a failover backend. More
	than one failover backend can be configured for a given RegionBackendService. */
	// +optional
	Failover *bool `json:"failover,omitempty"`

	/* Reference to a ComputeInstanceGroup or ComputeNetworkEndpointGroup
	resource. In case of instance group this defines the list of
	instances that serve traffic. Member virtual machine instances from
	each instance group must live in the same zone as the instance
	group itself. No two backends in a backend service are allowed to
	use same Instance Group resource.

	For Network Endpoint Groups this defines list of endpoints. All
	endpoints of Network Endpoint Group must be hosted on instances
	located in the same zone as the Network Endpoint Group.

	Backend services cannot mix Instance Group and Network Endpoint
	Group backends.

	When the 'load_balancing_scheme' is INTERNAL, only instance groups
	are supported. */
	Group BackendserviceGroup `json:"group"`

	/* The max number of simultaneous connections for the group. Can
	be used with either CONNECTION or UTILIZATION balancing modes.

	For CONNECTION mode, either maxConnections or one
	of maxConnectionsPerInstance or maxConnectionsPerEndpoint,
	as appropriate for group type, must be set. */
	// +optional
	MaxConnections *int `json:"maxConnections,omitempty"`

	/* The max number of simultaneous connections that a single backend
	network endpoint can handle. This is used to calculate the
	capacity of the group. Can be used in either CONNECTION or
	UTILIZATION balancing modes.

	For CONNECTION mode, either
	maxConnections or maxConnectionsPerEndpoint must be set. */
	// +optional
	MaxConnectionsPerEndpoint *int `json:"maxConnectionsPerEndpoint,omitempty"`

	/* The max number of simultaneous connections that a single
	backend instance can handle. This is used to calculate the
	capacity of the group. Can be used in either CONNECTION or
	UTILIZATION balancing modes.

	For CONNECTION mode, either maxConnections or
	maxConnectionsPerInstance must be set. */
	// +optional
	MaxConnectionsPerInstance *int `json:"maxConnectionsPerInstance,omitempty"`

	/* The max requests per second (RPS) of the group.

	Can be used with either RATE or UTILIZATION balancing modes,
	but required if RATE mode. For RATE mode, either maxRate or one
	of maxRatePerInstance or maxRatePerEndpoint, as appropriate for
	group type, must be set. */
	// +optional
	MaxRate *int `json:"maxRate,omitempty"`

	/* The max requests per second (RPS) that a single backend network
	endpoint can handle. This is used to calculate the capacity of
	the group. Can be used in either balancing mode. For RATE mode,
	either maxRate or maxRatePerEndpoint must be set. */
	// +optional
	MaxRatePerEndpoint *float64 `json:"maxRatePerEndpoint,omitempty"`

	/* The max requests per second (RPS) that a single backend
	instance can handle. This is used to calculate the capacity of
	the group. Can be used in either balancing mode. For RATE mode,
	either maxRate or maxRatePerInstance must be set. */
	// +optional
	MaxRatePerInstance *float64 `json:"maxRatePerInstance,omitempty"`

	/* Used when balancingMode is UTILIZATION. This ratio defines the
	CPU utilization target for the group. Valid range is [0.0, 1.0]. */
	// +optional
	MaxUtilization *float64 `json:"maxUtilization,omitempty"`
}

type BackendserviceBaseEjectionTime struct {
	/* Span of time that's a fraction of a second at nanosecond resolution. Durations
	less than one second are represented with a 0 'seconds' field and a positive
	'nanos' field. Must be from 0 to 999,999,999 inclusive. */
	// +optional
	Nanos *int `json:"nanos,omitempty"`

	/* Span of time at a resolution of a second. Must be from 0 to 315,576,000,000
	inclusive. */
	Seconds int `json:"seconds"`
}

type BackendserviceCacheKeyPolicy struct {
	/* If true requests to different hosts will be cached separately. */
	// +optional
	IncludeHost *bool `json:"includeHost,omitempty"`

	/* If true, http and https requests will be cached separately. */
	// +optional
	IncludeProtocol *bool `json:"includeProtocol,omitempty"`

	/* If true, include query string parameters in the cache key
	according to query_string_whitelist and
	query_string_blacklist. If neither is set, the entire query
	string will be included.

	If false, the query string will be excluded from the cache
	key entirely. */
	// +optional
	IncludeQueryString *bool `json:"includeQueryString,omitempty"`

	/* Names of query string parameters to exclude in cache keys.

	All other parameters will be included. Either specify
	query_string_whitelist or query_string_blacklist, not both.
	'&' and '=' will be percent encoded and not treated as
	delimiters. */
	// +optional
	QueryStringBlacklist []string `json:"queryStringBlacklist,omitempty"`

	/* Names of query string parameters to include in cache keys.

	All other parameters will be excluded. Either specify
	query_string_whitelist or query_string_blacklist, not both.
	'&' and '=' will be percent encoded and not treated as
	delimiters. */
	// +optional
	QueryStringWhitelist []string `json:"queryStringWhitelist,omitempty"`
}

type BackendserviceCdnPolicy struct {
	/* The CacheKeyPolicy for this CdnPolicy. */
	// +optional
	CacheKeyPolicy *BackendserviceCacheKeyPolicy `json:"cacheKeyPolicy,omitempty"`

	/* Specifies the cache setting for all responses from this backend.
	The possible values are: USE_ORIGIN_HEADERS, FORCE_CACHE_ALL and CACHE_ALL_STATIC Possible values: ["USE_ORIGIN_HEADERS", "FORCE_CACHE_ALL", "CACHE_ALL_STATIC"]. */
	// +optional
	CacheMode *string `json:"cacheMode,omitempty"`

	/* Specifies the maximum allowed TTL for cached content served by this origin. */
	// +optional
	ClientTtl *int `json:"clientTtl,omitempty"`

	/* Specifies the default TTL for cached content served by this origin for responses
	that do not have an existing valid TTL (max-age or s-max-age). */
	// +optional
	DefaultTtl *int `json:"defaultTtl,omitempty"`

	/* Specifies the maximum allowed TTL for cached content served by this origin. */
	// +optional
	MaxTtl *int `json:"maxTtl,omitempty"`

	/* Negative caching allows per-status code TTLs to be set, in order to apply fine-grained caching for common errors or redirects. */
	// +optional
	NegativeCaching *bool `json:"negativeCaching,omitempty"`

	/* Sets a cache TTL for the specified HTTP status code. negativeCaching must be enabled to configure negativeCachingPolicy.
	Omitting the policy and leaving negativeCaching enabled will use Cloud CDN's default cache TTLs. */
	// +optional
	NegativeCachingPolicy []BackendserviceNegativeCachingPolicy `json:"negativeCachingPolicy,omitempty"`

	/* Serve existing content from the cache (if available) when revalidating content with the origin, or when an error is encountered when refreshing the cache. */
	// +optional
	ServeWhileStale *int `json:"serveWhileStale,omitempty"`

	/* Maximum number of seconds the response to a signed URL request
	will be considered fresh, defaults to 1hr (3600s). After this
	time period, the response will be revalidated before
	being served.

	When serving responses to signed URL requests, Cloud CDN will
	internally behave as though all responses from this backend had a
	"Cache-Control: public, max-age=[TTL]" header, regardless of any
	existing Cache-Control header. The actual headers served in
	responses will not be altered. */
	// +optional
	SignedUrlCacheMaxAgeSec *int `json:"signedUrlCacheMaxAgeSec,omitempty"`
}

type BackendserviceCircuitBreakers struct {
	/* The timeout for new network connections to hosts. */
	// +optional
	ConnectTimeout *BackendserviceConnectTimeout `json:"connectTimeout,omitempty"`

	/* The maximum number of connections to the backend cluster.
	Defaults to 1024. */
	// +optional
	MaxConnections *int `json:"maxConnections,omitempty"`

	/* The maximum number of pending requests to the backend cluster.
	Defaults to 1024. */
	// +optional
	MaxPendingRequests *int `json:"maxPendingRequests,omitempty"`

	/* The maximum number of parallel requests to the backend cluster.
	Defaults to 1024. */
	// +optional
	MaxRequests *int `json:"maxRequests,omitempty"`

	/* Maximum requests for a single backend connection. This parameter
	is respected by both the HTTP/1.1 and HTTP/2 implementations. If
	not specified, there is no limit. Setting this parameter to 1
	will effectively disable keep alive. */
	// +optional
	MaxRequestsPerConnection *int `json:"maxRequestsPerConnection,omitempty"`

	/* The maximum number of parallel retries to the backend cluster.
	Defaults to 3. */
	// +optional
	MaxRetries *int `json:"maxRetries,omitempty"`
}

type BackendserviceConnectTimeout struct {
	/* Span of time that's a fraction of a second at nanosecond
	resolution. Durations less than one second are represented
	with a 0 seconds field and a positive nanos field. Must
	be from 0 to 999,999,999 inclusive. */
	// +optional
	Nanos *int `json:"nanos,omitempty"`

	/* Span of time at a resolution of a second.
	Must be from 0 to 315,576,000,000 inclusive. */
	Seconds int `json:"seconds"`
}

type BackendserviceConnectionTrackingPolicy struct {
	/* Specifies connection persistence when backends are unhealthy.

	If set to 'DEFAULT_FOR_PROTOCOL', the existing connections persist on
	unhealthy backends only for connection-oriented protocols (TCP and SCTP)
	and only if the Tracking Mode is PER_CONNECTION (default tracking mode)
	or the Session Affinity is configured for 5-tuple. They do not persist
	for UDP.

	If set to 'NEVER_PERSIST', after a backend becomes unhealthy, the existing
	connections on the unhealthy backend are never persisted on the unhealthy
	backend. They are always diverted to newly selected healthy backends
	(unless all backends are unhealthy).

	If set to 'ALWAYS_PERSIST', existing connections always persist on
	unhealthy backends regardless of protocol and session affinity. It is
	generally not recommended to use this mode overriding the default. Default value: "DEFAULT_FOR_PROTOCOL" Possible values: ["DEFAULT_FOR_PROTOCOL", "NEVER_PERSIST", "ALWAYS_PERSIST"]. */
	// +optional
	ConnectionPersistenceOnUnhealthyBackends *string `json:"connectionPersistenceOnUnhealthyBackends,omitempty"`

	/* Specifies how long to keep a Connection Tracking entry while there is
	no matching traffic (in seconds).

	For L4 ILB the minimum(default) is 10 minutes and maximum is 16 hours.

	For NLB the minimum(default) is 60 seconds and the maximum is 16 hours. */
	// +optional
	IdleTimeoutSec *int `json:"idleTimeoutSec,omitempty"`

	/* Specifies the key used for connection tracking. There are two options:
	'PER_CONNECTION': The Connection Tracking is performed as per the
	Connection Key (default Hash Method) for the specific protocol.

	'PER_SESSION': The Connection Tracking is performed as per the
	configured Session Affinity. It matches the configured Session Affinity. Default value: "PER_CONNECTION" Possible values: ["PER_CONNECTION", "PER_SESSION"]. */
	// +optional
	TrackingMode *string `json:"trackingMode,omitempty"`
}

type BackendserviceConsistentHash struct {
	/* Hash is based on HTTP Cookie. This field describes a HTTP cookie
	that will be used as the hash key for the consistent hash load
	balancer. If the cookie is not present, it will be generated.
	This field is applicable if the sessionAffinity is set to HTTP_COOKIE. */
	// +optional
	HttpCookie *BackendserviceHttpCookie `json:"httpCookie,omitempty"`

	/* The hash based on the value of the specified header field.
	This field is applicable if the sessionAffinity is set to HEADER_FIELD. */
	// +optional
	HttpHeaderName *string `json:"httpHeaderName,omitempty"`

	/* The minimum number of virtual nodes to use for the hash ring.
	Larger ring sizes result in more granular load
	distributions. If the number of hosts in the load balancing pool
	is larger than the ring size, each host will be assigned a single
	virtual node.
	Defaults to 1024. */
	// +optional
	MinimumRingSize *int `json:"minimumRingSize,omitempty"`
}

type BackendserviceFailoverPolicy struct {
	/* On failover or failback, this field indicates whether connection drain
	will be honored. Setting this to true has the following effect: connections
	to the old active pool are not drained. Connections to the new active pool
	use the timeout of 10 min (currently fixed). Setting to false has the
	following effect: both old and new connections will have a drain timeout
	of 10 min.
	This can be set to true only if the protocol is TCP.
	The default is false. */
	// +optional
	DisableConnectionDrainOnFailover *bool `json:"disableConnectionDrainOnFailover,omitempty"`

	/* This option is used only when no healthy VMs are detected in the primary
	and backup instance groups. When set to true, traffic is dropped. When
	set to false, new connections are sent across all VMs in the primary group.
	The default is false. */
	// +optional
	DropTrafficIfUnhealthy *bool `json:"dropTrafficIfUnhealthy,omitempty"`

	/* The value of the field must be in [0, 1]. If the ratio of the healthy
	VMs in the primary backend is at or below this number, traffic arriving
	at the load-balanced IP will be directed to the failover backend.
	In case where 'failoverRatio' is not set or all the VMs in the backup
	backend are unhealthy, the traffic will be directed back to the primary
	backend in the "force" mode, where traffic will be spread to the healthy
	VMs with the best effort, or to all VMs when no VM is healthy.
	This field is only used with l4 load balancing. */
	// +optional
	FailoverRatio *float64 `json:"failoverRatio,omitempty"`
}

type BackendserviceGroup struct {
	/*  */
	// +optional
	InstanceGroupRef *v1alpha1.ResourceRef `json:"instanceGroupRef,omitempty"`

	/*  */
	// +optional
	NetworkEndpointGroupRef *v1alpha1.ResourceRef `json:"networkEndpointGroupRef,omitempty"`
}

type BackendserviceHealthChecks struct {
	/*  */
	// +optional
	HealthCheckRef *v1alpha1.ResourceRef `json:"healthCheckRef,omitempty"`

	/*  */
	// +optional
	HttpHealthCheckRef *v1alpha1.ResourceRef `json:"httpHealthCheckRef,omitempty"`
}

type BackendserviceHttpCookie struct {
	/* Name of the cookie. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* Path to set for the cookie. */
	// +optional
	Path *string `json:"path,omitempty"`

	/* Lifetime of the cookie. */
	// +optional
	Ttl *BackendserviceTtl `json:"ttl,omitempty"`
}

type BackendserviceIap struct {
	/* OAuth2 Client ID for IAP. */
	Oauth2ClientId string `json:"oauth2ClientId"`

	/* OAuth2 Client Secret for IAP. */
	// +optional
	Oauth2ClientSecret *BackendserviceOauth2ClientSecret `json:"oauth2ClientSecret,omitempty"`

	/* OAuth2 Client Secret SHA-256 for IAP. */
	// +optional
	Oauth2ClientSecretSha256 *string `json:"oauth2ClientSecretSha256,omitempty"`
}

type BackendserviceInterval struct {
	/* Span of time that's a fraction of a second at nanosecond resolution. Durations
	less than one second are represented with a 0 'seconds' field and a positive
	'nanos' field. Must be from 0 to 999,999,999 inclusive. */
	// +optional
	Nanos *int `json:"nanos,omitempty"`

	/* Span of time at a resolution of a second. Must be from 0 to 315,576,000,000
	inclusive. */
	Seconds int `json:"seconds"`
}

type BackendserviceLogConfig struct {
	/* Whether to enable logging for the load balancer traffic served by this backend service. */
	// +optional
	Enable *bool `json:"enable,omitempty"`

	/* This field can only be specified if logging is enabled for this backend service. The value of
	the field must be in [0, 1]. This configures the sampling rate of requests to the load balancer
	where 1.0 means all logged requests are reported and 0.0 means no logged requests are reported.
	The default value is 1.0. */
	// +optional
	SampleRate *float64 `json:"sampleRate,omitempty"`
}

type BackendserviceNegativeCachingPolicy struct {
	/* The HTTP status code to define a TTL against. Only HTTP status codes 300, 301, 308, 404, 405, 410, 421, 451 and 501
	can be specified as values, and you cannot specify a status code more than once. */
	// +optional
	Code *int `json:"code,omitempty"`

	/* The TTL (in seconds) for which to cache responses with the corresponding status code. The maximum allowed value is 1800s
	(30 minutes), noting that infrequently accessed objects may be evicted from the cache before the defined TTL. */
	// +optional
	Ttl *int `json:"ttl,omitempty"`
}

type BackendserviceOauth2ClientSecret struct {
	/* Value of the field. Cannot be used if 'valueFrom' is specified. */
	// +optional
	Value *string `json:"value,omitempty"`

	/* Source for the field's value. Cannot be used if 'value' is specified. */
	// +optional
	ValueFrom *BackendserviceValueFrom `json:"valueFrom,omitempty"`
}

type BackendserviceOutlierDetection struct {
	/* The base time that a host is ejected for. The real time is equal to the base
	time multiplied by the number of times the host has been ejected. Defaults to
	30000ms or 30s. */
	// +optional
	BaseEjectionTime *BackendserviceBaseEjectionTime `json:"baseEjectionTime,omitempty"`

	/* Number of errors before a host is ejected from the connection pool. When the
	backend host is accessed over HTTP, a 5xx return code qualifies as an error.
	Defaults to 5. */
	// +optional
	ConsecutiveErrors *int `json:"consecutiveErrors,omitempty"`

	/* The number of consecutive gateway failures (502, 503, 504 status or connection
	errors that are mapped to one of those status codes) before a consecutive
	gateway failure ejection occurs. Defaults to 5. */
	// +optional
	ConsecutiveGatewayFailure *int `json:"consecutiveGatewayFailure,omitempty"`

	/* The percentage chance that a host will be actually ejected when an outlier
	status is detected through consecutive 5xx. This setting can be used to disable
	ejection or to ramp it up slowly. Defaults to 100. */
	// +optional
	EnforcingConsecutiveErrors *int `json:"enforcingConsecutiveErrors,omitempty"`

	/* The percentage chance that a host will be actually ejected when an outlier
	status is detected through consecutive gateway failures. This setting can be
	used to disable ejection or to ramp it up slowly. Defaults to 0. */
	// +optional
	EnforcingConsecutiveGatewayFailure *int `json:"enforcingConsecutiveGatewayFailure,omitempty"`

	/* The percentage chance that a host will be actually ejected when an outlier
	status is detected through success rate statistics. This setting can be used to
	disable ejection or to ramp it up slowly. Defaults to 100. */
	// +optional
	EnforcingSuccessRate *int `json:"enforcingSuccessRate,omitempty"`

	/* Time interval between ejection sweep analysis. This can result in both new
	ejections as well as hosts being returned to service. Defaults to 10 seconds. */
	// +optional
	Interval *BackendserviceInterval `json:"interval,omitempty"`

	/* Maximum percentage of hosts in the load balancing pool for the backend service
	that can be ejected. Defaults to 10%. */
	// +optional
	MaxEjectionPercent *int `json:"maxEjectionPercent,omitempty"`

	/* The number of hosts in a cluster that must have enough request volume to detect
	success rate outliers. If the number of hosts is less than this setting, outlier
	detection via success rate statistics is not performed for any host in the
	cluster. Defaults to 5. */
	// +optional
	SuccessRateMinimumHosts *int `json:"successRateMinimumHosts,omitempty"`

	/* The minimum number of total requests that must be collected in one interval (as
	defined by the interval duration above) to include this host in success rate
	based outlier detection. If the volume is lower than this setting, outlier
	detection via success rate statistics is not performed for that host. Defaults
	to 100. */
	// +optional
	SuccessRateRequestVolume *int `json:"successRateRequestVolume,omitempty"`

	/* This factor is used to determine the ejection threshold for success rate outlier
	ejection. The ejection threshold is the difference between the mean success
	rate, and the product of this factor and the standard deviation of the mean
	success rate: mean - (stdev * success_rate_stdev_factor). This factor is divided
	by a thousand to get a double. That is, if the desired factor is 1.9, the
	runtime value should be 1900. Defaults to 1900. */
	// +optional
	SuccessRateStdevFactor *int `json:"successRateStdevFactor,omitempty"`
}

type BackendserviceSecuritySettings struct {
	/* ClientTlsPolicy is a resource that specifies how a client should
	authenticate connections to backends of a service. This resource itself
	does not affect configuration unless it is attached to a backend
	service resource. *ConfigConnector only supports `external`
	references for this field.* */
	ClientTLSPolicyRef v1alpha1.ResourceRef `json:"clientTLSPolicyRef"`

	/* A list of alternate names to verify the subject identity in the certificate.
	If specified, the client will verify that the server certificate's subject
	alt name matches one of the specified values. */
	SubjectAltNames []string `json:"subjectAltNames"`
}

type BackendserviceTtl struct {
	/* Span of time that's a fraction of a second at nanosecond
	resolution. Durations less than one second are represented
	with a 0 seconds field and a positive nanos field. Must
	be from 0 to 999,999,999 inclusive. */
	// +optional
	Nanos *int `json:"nanos,omitempty"`

	/* Span of time at a resolution of a second.
	Must be from 0 to 315,576,000,000 inclusive. */
	Seconds int `json:"seconds"`
}

type BackendserviceValueFrom struct {
	/* Reference to a value with the given key in the given Secret in the resource's namespace. */
	// +optional
	SecretKeyRef *v1alpha1.ResourceRef `json:"secretKeyRef,omitempty"`
}

type ComputeBackendServiceSpec struct {
	/* Lifetime of cookies in seconds if session_affinity is
	GENERATED_COOKIE. If set to 0, the cookie is non-persistent and lasts
	only until the end of the browser session (or equivalent). The
	maximum allowed value for TTL is one day.

	When the load balancing scheme is INTERNAL, this field is not used. */
	// +optional
	AffinityCookieTtlSec *int `json:"affinityCookieTtlSec,omitempty"`

	/* The set of backends that serve this BackendService. */
	// +optional
	Backend []BackendserviceBackend `json:"backend,omitempty"`

	/* Cloud CDN configuration for this BackendService. */
	// +optional
	CdnPolicy *BackendserviceCdnPolicy `json:"cdnPolicy,omitempty"`

	/* Settings controlling the volume of connections to a backend service. This field
	is applicable only when the load_balancing_scheme is set to INTERNAL_SELF_MANAGED. */
	// +optional
	CircuitBreakers *BackendserviceCircuitBreakers `json:"circuitBreakers,omitempty"`

	/* Time for which instance will be drained (not accept new
	connections, but still work to finish started). */
	// +optional
	ConnectionDrainingTimeoutSec *int `json:"connectionDrainingTimeoutSec,omitempty"`

	/* Connection Tracking configuration for this BackendService.
	This is available only for Layer 4 Internal Load Balancing and
	Network Load Balancing. */
	// +optional
	ConnectionTrackingPolicy *BackendserviceConnectionTrackingPolicy `json:"connectionTrackingPolicy,omitempty"`

	/* Consistent Hash-based load balancing can be used to provide soft session
	affinity based on HTTP headers, cookies or other properties. This load balancing
	policy is applicable only for HTTP connections. The affinity to a particular
	destination host will be lost when one or more hosts are added/removed from the
	destination service. This field specifies parameters that control consistent
	hashing. This field only applies if the load_balancing_scheme is set to
	INTERNAL_SELF_MANAGED. This field is only applicable when locality_lb_policy is
	set to MAGLEV or RING_HASH. */
	// +optional
	ConsistentHash *BackendserviceConsistentHash `json:"consistentHash,omitempty"`

	/* Headers that the HTTP/S load balancer should add to proxied
	requests. */
	// +optional
	CustomRequestHeaders []string `json:"customRequestHeaders,omitempty"`

	/* Headers that the HTTP/S load balancer should add to proxied
	responses. */
	// +optional
	CustomResponseHeaders []string `json:"customResponseHeaders,omitempty"`

	/* An optional description of this resource. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* If true, enable Cloud CDN for this BackendService. */
	// +optional
	EnableCdn *bool `json:"enableCdn,omitempty"`

	/* Policy for failovers. */
	// +optional
	FailoverPolicy *BackendserviceFailoverPolicy `json:"failoverPolicy,omitempty"`

	/*  */
	// +optional
	HealthChecks []BackendserviceHealthChecks `json:"healthChecks,omitempty"`

	/* Settings for enabling Cloud Identity Aware Proxy. */
	// +optional
	Iap *BackendserviceIap `json:"iap,omitempty"`

	/* Immutable. Indicates whether the backend service will be used with internal or
	external load balancing. A backend service created for one type of
	load balancing cannot be used with the other. For more information, refer to
	[Choosing a load balancer](https://cloud.google.com/load-balancing/docs/backend-service). Default value: "EXTERNAL" Possible values: ["EXTERNAL", "INTERNAL_SELF_MANAGED", "EXTERNAL_MANAGED"]. */
	// +optional
	LoadBalancingScheme *string `json:"loadBalancingScheme,omitempty"`

	/* The load balancing algorithm used within the scope of the locality.
	The possible values are:

	* 'ROUND_ROBIN': This is a simple policy in which each healthy backend
	                 is selected in round robin order.

	* 'LEAST_REQUEST': An O(1) algorithm which selects two random healthy
	                   hosts and picks the host which has fewer active requests.

	* 'RING_HASH': The ring/modulo hash load balancer implements consistent
	               hashing to backends. The algorithm has the property that the
	               addition/removal of a host from a set of N hosts only affects
	               1/N of the requests.

	* 'RANDOM': The load balancer selects a random healthy host.

	* 'ORIGINAL_DESTINATION': Backend host is selected based on the client
	                          connection metadata, i.e., connections are opened
	                          to the same address as the destination address of
	                          the incoming connection before the connection
	                          was redirected to the load balancer.

	* 'MAGLEV': used as a drop in replacement for the ring hash load balancer.
	            Maglev is not as stable as ring hash but has faster table lookup
	            build times and host selection times. For more information about
	            Maglev, refer to https://ai.google/research/pubs/pub44824


	This field is applicable to either:

	* A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2,
	  and loadBalancingScheme set to INTERNAL_MANAGED.
	* A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED.


	If session_affinity is not NONE, and this field is not set to MAGLEV or RING_HASH,
	session affinity settings will not take effect.

	Only ROUND_ROBIN and RING_HASH are supported when the backend service is referenced
	by a URL map that is bound to target gRPC proxy that has validate_for_proxyless
	field set to true. Possible values: ["ROUND_ROBIN", "LEAST_REQUEST", "RING_HASH", "RANDOM", "ORIGINAL_DESTINATION", "MAGLEV"]. */
	// +optional
	LocalityLbPolicy *string `json:"localityLbPolicy,omitempty"`

	/* Location represents the geographical location of the ComputeBackendService. Specify a region name or "global" for global resources. Reference: GCP definition of regions/zones (https://cloud.google.com/compute/docs/regions-zones/) */
	Location string `json:"location"`

	/* This field denotes the logging options for the load balancer traffic served by this backend service.
	If logging is enabled, logs will be exported to Stackdriver. */
	// +optional
	LogConfig *BackendserviceLogConfig `json:"logConfig,omitempty"`

	/* The network to which this backend service belongs.  This field can
	only be specified when the load balancing scheme is set to
	INTERNAL. */
	// +optional
	NetworkRef *v1alpha1.ResourceRef `json:"networkRef,omitempty"`

	/* Settings controlling eviction of unhealthy hosts from the load balancing pool.
	This field is applicable only when the load_balancing_scheme is set
	to INTERNAL_SELF_MANAGED. */
	// +optional
	OutlierDetection *BackendserviceOutlierDetection `json:"outlierDetection,omitempty"`

	/* Name of backend port. The same name should appear in the instance
	groups referenced by this service. Required when the load balancing
	scheme is EXTERNAL. */
	// +optional
	PortName *string `json:"portName,omitempty"`

	/* The protocol this BackendService uses to communicate with backends.
	The default is HTTP. **NOTE**: HTTP2 is only valid for beta HTTP/2 load balancer
	types and may result in errors if used with the GA API. Possible values: ["HTTP", "HTTPS", "HTTP2", "TCP", "SSL", "GRPC"]. */
	// +optional
	Protocol *string `json:"protocol,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* The security policy associated with this backend service. */
	// +optional
	SecurityPolicyRef *v1alpha1.ResourceRef `json:"securityPolicyRef,omitempty"`

	/* The security settings that apply to this backend service. This field is applicable to either
	a regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and
	load_balancing_scheme set to INTERNAL_MANAGED; or a global backend service with the
	load_balancing_scheme set to INTERNAL_SELF_MANAGED. */
	// +optional
	SecuritySettings *BackendserviceSecuritySettings `json:"securitySettings,omitempty"`

	/* Type of session affinity to use. The default is NONE. Session affinity is
	not applicable if the protocol is UDP. Possible values: ["NONE", "CLIENT_IP", "CLIENT_IP_PORT_PROTO", "CLIENT_IP_PROTO", "GENERATED_COOKIE", "HEADER_FIELD", "HTTP_COOKIE"]. */
	// +optional
	SessionAffinity *string `json:"sessionAffinity,omitempty"`

	/* How many seconds to wait for the backend before considering it a
	failed request. Default is 30 seconds. Valid range is [1, 86400]. */
	// +optional
	TimeoutSec *int `json:"timeoutSec,omitempty"`
}

type ComputeBackendServiceStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeBackendService's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Creation timestamp in RFC3339 text format. */
	CreationTimestamp string `json:"creationTimestamp,omitempty"`
	/* Fingerprint of this resource. A hash of the contents stored in this
	object. This field is used in optimistic locking. */
	Fingerprint string `json:"fingerprint,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration int `json:"observedGeneration,omitempty"`
	/*  */
	SelfLink string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeBackendService is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeBackendService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeBackendServiceSpec   `json:"spec,omitempty"`
	Status ComputeBackendServiceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeBackendServiceList contains a list of ComputeBackendService
type ComputeBackendServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeBackendService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeBackendService{}, &ComputeBackendServiceList{})
}
