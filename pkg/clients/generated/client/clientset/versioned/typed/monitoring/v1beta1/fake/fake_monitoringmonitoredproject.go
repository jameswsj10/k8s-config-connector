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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/monitoring/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMonitoringMonitoredProjects implements MonitoringMonitoredProjectInterface
type FakeMonitoringMonitoredProjects struct {
	Fake *FakeMonitoringV1beta1
	ns   string
}

var monitoringmonitoredprojectsResource = schema.GroupVersionResource{Group: "monitoring.cnrm.cloud.google.com", Version: "v1beta1", Resource: "monitoringmonitoredprojects"}

var monitoringmonitoredprojectsKind = schema.GroupVersionKind{Group: "monitoring.cnrm.cloud.google.com", Version: "v1beta1", Kind: "MonitoringMonitoredProject"}

// Get takes name of the monitoringMonitoredProject, and returns the corresponding monitoringMonitoredProject object, and an error if there is any.
func (c *FakeMonitoringMonitoredProjects) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.MonitoringMonitoredProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(monitoringmonitoredprojectsResource, c.ns, name), &v1beta1.MonitoringMonitoredProject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MonitoringMonitoredProject), err
}

// List takes label and field selectors, and returns the list of MonitoringMonitoredProjects that match those selectors.
func (c *FakeMonitoringMonitoredProjects) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.MonitoringMonitoredProjectList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(monitoringmonitoredprojectsResource, monitoringmonitoredprojectsKind, c.ns, opts), &v1beta1.MonitoringMonitoredProjectList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.MonitoringMonitoredProjectList{ListMeta: obj.(*v1beta1.MonitoringMonitoredProjectList).ListMeta}
	for _, item := range obj.(*v1beta1.MonitoringMonitoredProjectList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested monitoringMonitoredProjects.
func (c *FakeMonitoringMonitoredProjects) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(monitoringmonitoredprojectsResource, c.ns, opts))

}

// Create takes the representation of a monitoringMonitoredProject and creates it.  Returns the server's representation of the monitoringMonitoredProject, and an error, if there is any.
func (c *FakeMonitoringMonitoredProjects) Create(ctx context.Context, monitoringMonitoredProject *v1beta1.MonitoringMonitoredProject, opts v1.CreateOptions) (result *v1beta1.MonitoringMonitoredProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(monitoringmonitoredprojectsResource, c.ns, monitoringMonitoredProject), &v1beta1.MonitoringMonitoredProject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MonitoringMonitoredProject), err
}

// Update takes the representation of a monitoringMonitoredProject and updates it. Returns the server's representation of the monitoringMonitoredProject, and an error, if there is any.
func (c *FakeMonitoringMonitoredProjects) Update(ctx context.Context, monitoringMonitoredProject *v1beta1.MonitoringMonitoredProject, opts v1.UpdateOptions) (result *v1beta1.MonitoringMonitoredProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(monitoringmonitoredprojectsResource, c.ns, monitoringMonitoredProject), &v1beta1.MonitoringMonitoredProject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MonitoringMonitoredProject), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMonitoringMonitoredProjects) UpdateStatus(ctx context.Context, monitoringMonitoredProject *v1beta1.MonitoringMonitoredProject, opts v1.UpdateOptions) (*v1beta1.MonitoringMonitoredProject, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(monitoringmonitoredprojectsResource, "status", c.ns, monitoringMonitoredProject), &v1beta1.MonitoringMonitoredProject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MonitoringMonitoredProject), err
}

// Delete takes name of the monitoringMonitoredProject and deletes it. Returns an error if one occurs.
func (c *FakeMonitoringMonitoredProjects) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(monitoringmonitoredprojectsResource, c.ns, name, opts), &v1beta1.MonitoringMonitoredProject{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMonitoringMonitoredProjects) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(monitoringmonitoredprojectsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.MonitoringMonitoredProjectList{})
	return err
}

// Patch applies the patch and returns the patched monitoringMonitoredProject.
func (c *FakeMonitoringMonitoredProjects) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.MonitoringMonitoredProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(monitoringmonitoredprojectsResource, c.ns, name, pt, data, subresources...), &v1beta1.MonitoringMonitoredProject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MonitoringMonitoredProject), err
}
