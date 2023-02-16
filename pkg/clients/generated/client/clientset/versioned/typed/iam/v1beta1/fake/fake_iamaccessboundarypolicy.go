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

package fake

import (
	"context"

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/iam/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIAMAccessBoundaryPolicies implements IAMAccessBoundaryPolicyInterface
type FakeIAMAccessBoundaryPolicies struct {
	Fake *FakeIamV1beta1
	ns   string
}

var iamaccessboundarypoliciesResource = schema.GroupVersionResource{Group: "iam.cnrm.cloud.google.com", Version: "v1beta1", Resource: "iamaccessboundarypolicies"}

var iamaccessboundarypoliciesKind = schema.GroupVersionKind{Group: "iam.cnrm.cloud.google.com", Version: "v1beta1", Kind: "IAMAccessBoundaryPolicy"}

// Get takes name of the iAMAccessBoundaryPolicy, and returns the corresponding iAMAccessBoundaryPolicy object, and an error if there is any.
func (c *FakeIAMAccessBoundaryPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.IAMAccessBoundaryPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(iamaccessboundarypoliciesResource, c.ns, name), &v1beta1.IAMAccessBoundaryPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IAMAccessBoundaryPolicy), err
}

// List takes label and field selectors, and returns the list of IAMAccessBoundaryPolicies that match those selectors.
func (c *FakeIAMAccessBoundaryPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.IAMAccessBoundaryPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(iamaccessboundarypoliciesResource, iamaccessboundarypoliciesKind, c.ns, opts), &v1beta1.IAMAccessBoundaryPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.IAMAccessBoundaryPolicyList{ListMeta: obj.(*v1beta1.IAMAccessBoundaryPolicyList).ListMeta}
	for _, item := range obj.(*v1beta1.IAMAccessBoundaryPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iAMAccessBoundaryPolicies.
func (c *FakeIAMAccessBoundaryPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(iamaccessboundarypoliciesResource, c.ns, opts))

}

// Create takes the representation of a iAMAccessBoundaryPolicy and creates it.  Returns the server's representation of the iAMAccessBoundaryPolicy, and an error, if there is any.
func (c *FakeIAMAccessBoundaryPolicies) Create(ctx context.Context, iAMAccessBoundaryPolicy *v1beta1.IAMAccessBoundaryPolicy, opts v1.CreateOptions) (result *v1beta1.IAMAccessBoundaryPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(iamaccessboundarypoliciesResource, c.ns, iAMAccessBoundaryPolicy), &v1beta1.IAMAccessBoundaryPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IAMAccessBoundaryPolicy), err
}

// Update takes the representation of a iAMAccessBoundaryPolicy and updates it. Returns the server's representation of the iAMAccessBoundaryPolicy, and an error, if there is any.
func (c *FakeIAMAccessBoundaryPolicies) Update(ctx context.Context, iAMAccessBoundaryPolicy *v1beta1.IAMAccessBoundaryPolicy, opts v1.UpdateOptions) (result *v1beta1.IAMAccessBoundaryPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(iamaccessboundarypoliciesResource, c.ns, iAMAccessBoundaryPolicy), &v1beta1.IAMAccessBoundaryPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IAMAccessBoundaryPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIAMAccessBoundaryPolicies) UpdateStatus(ctx context.Context, iAMAccessBoundaryPolicy *v1beta1.IAMAccessBoundaryPolicy, opts v1.UpdateOptions) (*v1beta1.IAMAccessBoundaryPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(iamaccessboundarypoliciesResource, "status", c.ns, iAMAccessBoundaryPolicy), &v1beta1.IAMAccessBoundaryPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IAMAccessBoundaryPolicy), err
}

// Delete takes name of the iAMAccessBoundaryPolicy and deletes it. Returns an error if one occurs.
func (c *FakeIAMAccessBoundaryPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(iamaccessboundarypoliciesResource, c.ns, name, opts), &v1beta1.IAMAccessBoundaryPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIAMAccessBoundaryPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(iamaccessboundarypoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.IAMAccessBoundaryPolicyList{})
	return err
}

// Patch applies the patch and returns the patched iAMAccessBoundaryPolicy.
func (c *FakeIAMAccessBoundaryPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.IAMAccessBoundaryPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(iamaccessboundarypoliciesResource, c.ns, name, pt, data, subresources...), &v1beta1.IAMAccessBoundaryPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IAMAccessBoundaryPolicy), err
}
