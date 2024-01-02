/*
Copyright 2024 Rancher Labs, Inc.

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

// FakeGoogleOAuthProviders implements GoogleOAuthProviderInterface
type FakeGoogleOAuthProviders struct {
	Fake *FakeManagementV3
}

var googleoauthprovidersResource = schema.GroupVersionResource{Group: "management.cattle.io", Version: "v3", Resource: "googleoauthproviders"}

var googleoauthprovidersKind = schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "GoogleOAuthProvider"}

// Get takes name of the googleOAuthProvider, and returns the corresponding googleOAuthProvider object, and an error if there is any.
func (c *FakeGoogleOAuthProviders) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.GoogleOAuthProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(googleoauthprovidersResource, name), &v3.GoogleOAuthProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.GoogleOAuthProvider), err
}

// List takes label and field selectors, and returns the list of GoogleOAuthProviders that match those selectors.
func (c *FakeGoogleOAuthProviders) List(ctx context.Context, opts v1.ListOptions) (result *v3.GoogleOAuthProviderList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(googleoauthprovidersResource, googleoauthprovidersKind, opts), &v3.GoogleOAuthProviderList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.GoogleOAuthProviderList{ListMeta: obj.(*v3.GoogleOAuthProviderList).ListMeta}
	for _, item := range obj.(*v3.GoogleOAuthProviderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested googleOAuthProviders.
func (c *FakeGoogleOAuthProviders) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(googleoauthprovidersResource, opts))
}

// Create takes the representation of a googleOAuthProvider and creates it.  Returns the server's representation of the googleOAuthProvider, and an error, if there is any.
func (c *FakeGoogleOAuthProviders) Create(ctx context.Context, googleOAuthProvider *v3.GoogleOAuthProvider, opts v1.CreateOptions) (result *v3.GoogleOAuthProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(googleoauthprovidersResource, googleOAuthProvider), &v3.GoogleOAuthProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.GoogleOAuthProvider), err
}

// Update takes the representation of a googleOAuthProvider and updates it. Returns the server's representation of the googleOAuthProvider, and an error, if there is any.
func (c *FakeGoogleOAuthProviders) Update(ctx context.Context, googleOAuthProvider *v3.GoogleOAuthProvider, opts v1.UpdateOptions) (result *v3.GoogleOAuthProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(googleoauthprovidersResource, googleOAuthProvider), &v3.GoogleOAuthProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.GoogleOAuthProvider), err
}

// Delete takes name of the googleOAuthProvider and deletes it. Returns an error if one occurs.
func (c *FakeGoogleOAuthProviders) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(googleoauthprovidersResource, name, opts), &v3.GoogleOAuthProvider{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGoogleOAuthProviders) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(googleoauthprovidersResource, listOpts)

	_, err := c.Fake.Invokes(action, &v3.GoogleOAuthProviderList{})
	return err
}

// Patch applies the patch and returns the patched googleOAuthProvider.
func (c *FakeGoogleOAuthProviders) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.GoogleOAuthProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(googleoauthprovidersResource, name, pt, data, subresources...), &v3.GoogleOAuthProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.GoogleOAuthProvider), err
}
