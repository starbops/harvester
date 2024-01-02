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

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/harvester/harvester/pkg/apis/harvesterhci.io/v1beta1"
	"github.com/rancher/lasso/pkg/client"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/generic"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

type VirtualMachineRestoreHandler func(string, *v1beta1.VirtualMachineRestore) (*v1beta1.VirtualMachineRestore, error)

type VirtualMachineRestoreController interface {
	generic.ControllerMeta
	VirtualMachineRestoreClient

	OnChange(ctx context.Context, name string, sync VirtualMachineRestoreHandler)
	OnRemove(ctx context.Context, name string, sync VirtualMachineRestoreHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() VirtualMachineRestoreCache
}

type VirtualMachineRestoreClient interface {
	Create(*v1beta1.VirtualMachineRestore) (*v1beta1.VirtualMachineRestore, error)
	Update(*v1beta1.VirtualMachineRestore) (*v1beta1.VirtualMachineRestore, error)

	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1beta1.VirtualMachineRestore, error)
	List(namespace string, opts metav1.ListOptions) (*v1beta1.VirtualMachineRestoreList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.VirtualMachineRestore, err error)
}

type VirtualMachineRestoreCache interface {
	Get(namespace, name string) (*v1beta1.VirtualMachineRestore, error)
	List(namespace string, selector labels.Selector) ([]*v1beta1.VirtualMachineRestore, error)

	AddIndexer(indexName string, indexer VirtualMachineRestoreIndexer)
	GetByIndex(indexName, key string) ([]*v1beta1.VirtualMachineRestore, error)
}

type VirtualMachineRestoreIndexer func(obj *v1beta1.VirtualMachineRestore) ([]string, error)

type virtualMachineRestoreController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewVirtualMachineRestoreController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) VirtualMachineRestoreController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &virtualMachineRestoreController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromVirtualMachineRestoreHandlerToHandler(sync VirtualMachineRestoreHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1beta1.VirtualMachineRestore
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1beta1.VirtualMachineRestore))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *virtualMachineRestoreController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1beta1.VirtualMachineRestore))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateVirtualMachineRestoreDeepCopyOnChange(client VirtualMachineRestoreClient, obj *v1beta1.VirtualMachineRestore, handler func(obj *v1beta1.VirtualMachineRestore) (*v1beta1.VirtualMachineRestore, error)) (*v1beta1.VirtualMachineRestore, error) {
	if obj == nil {
		return obj, nil
	}

	copyObj := obj.DeepCopy()
	newObj, err := handler(copyObj)
	if newObj != nil {
		copyObj = newObj
	}
	if obj.ResourceVersion == copyObj.ResourceVersion && !equality.Semantic.DeepEqual(obj, copyObj) {
		return client.Update(copyObj)
	}

	return copyObj, err
}

func (c *virtualMachineRestoreController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *virtualMachineRestoreController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *virtualMachineRestoreController) OnChange(ctx context.Context, name string, sync VirtualMachineRestoreHandler) {
	c.AddGenericHandler(ctx, name, FromVirtualMachineRestoreHandlerToHandler(sync))
}

func (c *virtualMachineRestoreController) OnRemove(ctx context.Context, name string, sync VirtualMachineRestoreHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromVirtualMachineRestoreHandlerToHandler(sync)))
}

func (c *virtualMachineRestoreController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *virtualMachineRestoreController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *virtualMachineRestoreController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *virtualMachineRestoreController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *virtualMachineRestoreController) Cache() VirtualMachineRestoreCache {
	return &virtualMachineRestoreCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *virtualMachineRestoreController) Create(obj *v1beta1.VirtualMachineRestore) (*v1beta1.VirtualMachineRestore, error) {
	result := &v1beta1.VirtualMachineRestore{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *virtualMachineRestoreController) Update(obj *v1beta1.VirtualMachineRestore) (*v1beta1.VirtualMachineRestore, error) {
	result := &v1beta1.VirtualMachineRestore{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *virtualMachineRestoreController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *virtualMachineRestoreController) Get(namespace, name string, options metav1.GetOptions) (*v1beta1.VirtualMachineRestore, error) {
	result := &v1beta1.VirtualMachineRestore{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *virtualMachineRestoreController) List(namespace string, opts metav1.ListOptions) (*v1beta1.VirtualMachineRestoreList, error) {
	result := &v1beta1.VirtualMachineRestoreList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *virtualMachineRestoreController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *virtualMachineRestoreController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v1beta1.VirtualMachineRestore, error) {
	result := &v1beta1.VirtualMachineRestore{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type virtualMachineRestoreCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *virtualMachineRestoreCache) Get(namespace, name string) (*v1beta1.VirtualMachineRestore, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1beta1.VirtualMachineRestore), nil
}

func (c *virtualMachineRestoreCache) List(namespace string, selector labels.Selector) (ret []*v1beta1.VirtualMachineRestore, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.VirtualMachineRestore))
	})

	return ret, err
}

func (c *virtualMachineRestoreCache) AddIndexer(indexName string, indexer VirtualMachineRestoreIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1beta1.VirtualMachineRestore))
		},
	}))
}

func (c *virtualMachineRestoreCache) GetByIndex(indexName, key string) (result []*v1beta1.VirtualMachineRestore, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1beta1.VirtualMachineRestore, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1beta1.VirtualMachineRestore))
	}
	return result, nil
}
