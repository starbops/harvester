/*
Copyright 2023 Rancher Labs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNotifiers implements NotifierInterface
type FakeNotifiers struct {
	Fake *FakeManagementV3
	ns   string
}

var notifiersResource = schema.GroupVersionResource{Group: "management.cattle.io", Version: "v3", Resource: "notifiers"}

var notifiersKind = schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "Notifier"}

// Get takes name of the notifier, and returns the corresponding notifier object, and an error if there is any.
func (c *FakeNotifiers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.Notifier, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(notifiersResource, c.ns, name), &v3.Notifier{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.Notifier), err
}

// List takes label and field selectors, and returns the list of Notifiers that match those selectors.
func (c *FakeNotifiers) List(ctx context.Context, opts v1.ListOptions) (result *v3.NotifierList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(notifiersResource, notifiersKind, c.ns, opts), &v3.NotifierList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.NotifierList{ListMeta: obj.(*v3.NotifierList).ListMeta}
	for _, item := range obj.(*v3.NotifierList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested notifiers.
func (c *FakeNotifiers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(notifiersResource, c.ns, opts))

}

// Create takes the representation of a notifier and creates it.  Returns the server's representation of the notifier, and an error, if there is any.
func (c *FakeNotifiers) Create(ctx context.Context, notifier *v3.Notifier, opts v1.CreateOptions) (result *v3.Notifier, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(notifiersResource, c.ns, notifier), &v3.Notifier{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.Notifier), err
}

// Update takes the representation of a notifier and updates it. Returns the server's representation of the notifier, and an error, if there is any.
func (c *FakeNotifiers) Update(ctx context.Context, notifier *v3.Notifier, opts v1.UpdateOptions) (result *v3.Notifier, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(notifiersResource, c.ns, notifier), &v3.Notifier{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.Notifier), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNotifiers) UpdateStatus(ctx context.Context, notifier *v3.Notifier, opts v1.UpdateOptions) (*v3.Notifier, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(notifiersResource, "status", c.ns, notifier), &v3.Notifier{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.Notifier), err
}

// Delete takes name of the notifier and deletes it. Returns an error if one occurs.
func (c *FakeNotifiers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(notifiersResource, c.ns, name, opts), &v3.Notifier{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNotifiers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(notifiersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v3.NotifierList{})
	return err
}

// Patch applies the patch and returns the patched notifier.
func (c *FakeNotifiers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.Notifier, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(notifiersResource, c.ns, name, pt, data, subresources...), &v3.Notifier{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.Notifier), err
}
