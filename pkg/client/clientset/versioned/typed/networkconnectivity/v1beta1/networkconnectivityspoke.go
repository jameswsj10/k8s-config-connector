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

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/networkconnectivity/v1beta1"
	scheme "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NetworkConnectivitySpokesGetter has a method to return a NetworkConnectivitySpokeInterface.
// A group's client should implement this interface.
type NetworkConnectivitySpokesGetter interface {
	NetworkConnectivitySpokes(namespace string) NetworkConnectivitySpokeInterface
}

// NetworkConnectivitySpokeInterface has methods to work with NetworkConnectivitySpoke resources.
type NetworkConnectivitySpokeInterface interface {
	Create(ctx context.Context, networkConnectivitySpoke *v1beta1.NetworkConnectivitySpoke, opts v1.CreateOptions) (*v1beta1.NetworkConnectivitySpoke, error)
	Update(ctx context.Context, networkConnectivitySpoke *v1beta1.NetworkConnectivitySpoke, opts v1.UpdateOptions) (*v1beta1.NetworkConnectivitySpoke, error)
	UpdateStatus(ctx context.Context, networkConnectivitySpoke *v1beta1.NetworkConnectivitySpoke, opts v1.UpdateOptions) (*v1beta1.NetworkConnectivitySpoke, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.NetworkConnectivitySpoke, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.NetworkConnectivitySpokeList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.NetworkConnectivitySpoke, err error)
	NetworkConnectivitySpokeExpansion
}

// networkConnectivitySpokes implements NetworkConnectivitySpokeInterface
type networkConnectivitySpokes struct {
	client rest.Interface
	ns     string
}

// newNetworkConnectivitySpokes returns a NetworkConnectivitySpokes
func newNetworkConnectivitySpokes(c *NetworkconnectivityV1beta1Client, namespace string) *networkConnectivitySpokes {
	return &networkConnectivitySpokes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the networkConnectivitySpoke, and returns the corresponding networkConnectivitySpoke object, and an error if there is any.
func (c *networkConnectivitySpokes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.NetworkConnectivitySpoke, err error) {
	result = &v1beta1.NetworkConnectivitySpoke{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networkconnectivityspokes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NetworkConnectivitySpokes that match those selectors.
func (c *networkConnectivitySpokes) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.NetworkConnectivitySpokeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.NetworkConnectivitySpokeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networkconnectivityspokes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested networkConnectivitySpokes.
func (c *networkConnectivitySpokes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("networkconnectivityspokes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a networkConnectivitySpoke and creates it.  Returns the server's representation of the networkConnectivitySpoke, and an error, if there is any.
func (c *networkConnectivitySpokes) Create(ctx context.Context, networkConnectivitySpoke *v1beta1.NetworkConnectivitySpoke, opts v1.CreateOptions) (result *v1beta1.NetworkConnectivitySpoke, err error) {
	result = &v1beta1.NetworkConnectivitySpoke{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("networkconnectivityspokes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(networkConnectivitySpoke).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a networkConnectivitySpoke and updates it. Returns the server's representation of the networkConnectivitySpoke, and an error, if there is any.
func (c *networkConnectivitySpokes) Update(ctx context.Context, networkConnectivitySpoke *v1beta1.NetworkConnectivitySpoke, opts v1.UpdateOptions) (result *v1beta1.NetworkConnectivitySpoke, err error) {
	result = &v1beta1.NetworkConnectivitySpoke{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("networkconnectivityspokes").
		Name(networkConnectivitySpoke.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(networkConnectivitySpoke).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *networkConnectivitySpokes) UpdateStatus(ctx context.Context, networkConnectivitySpoke *v1beta1.NetworkConnectivitySpoke, opts v1.UpdateOptions) (result *v1beta1.NetworkConnectivitySpoke, err error) {
	result = &v1beta1.NetworkConnectivitySpoke{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("networkconnectivityspokes").
		Name(networkConnectivitySpoke.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(networkConnectivitySpoke).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the networkConnectivitySpoke and deletes it. Returns an error if one occurs.
func (c *networkConnectivitySpokes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("networkconnectivityspokes").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *networkConnectivitySpokes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("networkconnectivityspokes").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched networkConnectivitySpoke.
func (c *networkConnectivitySpokes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.NetworkConnectivitySpoke, err error) {
	result = &v1beta1.NetworkConnectivitySpoke{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("networkconnectivityspokes").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
