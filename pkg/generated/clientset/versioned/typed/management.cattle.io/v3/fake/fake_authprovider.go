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

// FakeAuthProviders implements AuthProviderInterface
type FakeAuthProviders struct {
	Fake *FakeManagementV3
}

var authprovidersResource = schema.GroupVersionResource{Group: "management.cattle.io", Version: "v3", Resource: "authproviders"}

var authprovidersKind = schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "AuthProvider"}

// Get takes name of the authProvider, and returns the corresponding authProvider object, and an error if there is any.
func (c *FakeAuthProviders) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.AuthProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(authprovidersResource, name), &v3.AuthProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.AuthProvider), err
}

// List takes label and field selectors, and returns the list of AuthProviders that match those selectors.
func (c *FakeAuthProviders) List(ctx context.Context, opts v1.ListOptions) (result *v3.AuthProviderList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(authprovidersResource, authprovidersKind, opts), &v3.AuthProviderList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.AuthProviderList{ListMeta: obj.(*v3.AuthProviderList).ListMeta}
	for _, item := range obj.(*v3.AuthProviderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested authProviders.
func (c *FakeAuthProviders) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(authprovidersResource, opts))
}

// Create takes the representation of a authProvider and creates it.  Returns the server's representation of the authProvider, and an error, if there is any.
func (c *FakeAuthProviders) Create(ctx context.Context, authProvider *v3.AuthProvider, opts v1.CreateOptions) (result *v3.AuthProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(authprovidersResource, authProvider), &v3.AuthProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.AuthProvider), err
}

// Update takes the representation of a authProvider and updates it. Returns the server's representation of the authProvider, and an error, if there is any.
func (c *FakeAuthProviders) Update(ctx context.Context, authProvider *v3.AuthProvider, opts v1.UpdateOptions) (result *v3.AuthProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(authprovidersResource, authProvider), &v3.AuthProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.AuthProvider), err
}

// Delete takes name of the authProvider and deletes it. Returns an error if one occurs.
func (c *FakeAuthProviders) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(authprovidersResource, name, opts), &v3.AuthProvider{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAuthProviders) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(authprovidersResource, listOpts)

	_, err := c.Fake.Invokes(action, &v3.AuthProviderList{})
	return err
}

// Patch applies the patch and returns the patched authProvider.
func (c *FakeAuthProviders) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.AuthProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(authprovidersResource, name, pt, data, subresources...), &v3.AuthProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.AuthProvider), err
}
