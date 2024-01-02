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

package v1

import (
	"context"
	"time"

	"github.com/rancher/lasso/pkg/client"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/apply"
	"github.com/rancher/wrangler/pkg/condition"
	"github.com/rancher/wrangler/pkg/generic"
	"github.com/rancher/wrangler/pkg/kv"
	v1 "k8s.io/api/storage/v1"
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

type VolumeAttachmentHandler func(string, *v1.VolumeAttachment) (*v1.VolumeAttachment, error)

type VolumeAttachmentController interface {
	generic.ControllerMeta
	VolumeAttachmentClient

	OnChange(ctx context.Context, name string, sync VolumeAttachmentHandler)
	OnRemove(ctx context.Context, name string, sync VolumeAttachmentHandler)
	Enqueue(name string)
	EnqueueAfter(name string, duration time.Duration)

	Cache() VolumeAttachmentCache
}

type VolumeAttachmentClient interface {
	Create(*v1.VolumeAttachment) (*v1.VolumeAttachment, error)
	Update(*v1.VolumeAttachment) (*v1.VolumeAttachment, error)
	UpdateStatus(*v1.VolumeAttachment) (*v1.VolumeAttachment, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*v1.VolumeAttachment, error)
	List(opts metav1.ListOptions) (*v1.VolumeAttachmentList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.VolumeAttachment, err error)
}

type VolumeAttachmentCache interface {
	Get(name string) (*v1.VolumeAttachment, error)
	List(selector labels.Selector) ([]*v1.VolumeAttachment, error)

	AddIndexer(indexName string, indexer VolumeAttachmentIndexer)
	GetByIndex(indexName, key string) ([]*v1.VolumeAttachment, error)
}

type VolumeAttachmentIndexer func(obj *v1.VolumeAttachment) ([]string, error)

type volumeAttachmentController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewVolumeAttachmentController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) VolumeAttachmentController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &volumeAttachmentController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromVolumeAttachmentHandlerToHandler(sync VolumeAttachmentHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1.VolumeAttachment
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1.VolumeAttachment))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *volumeAttachmentController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1.VolumeAttachment))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateVolumeAttachmentDeepCopyOnChange(client VolumeAttachmentClient, obj *v1.VolumeAttachment, handler func(obj *v1.VolumeAttachment) (*v1.VolumeAttachment, error)) (*v1.VolumeAttachment, error) {
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

func (c *volumeAttachmentController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *volumeAttachmentController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *volumeAttachmentController) OnChange(ctx context.Context, name string, sync VolumeAttachmentHandler) {
	c.AddGenericHandler(ctx, name, FromVolumeAttachmentHandlerToHandler(sync))
}

func (c *volumeAttachmentController) OnRemove(ctx context.Context, name string, sync VolumeAttachmentHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromVolumeAttachmentHandlerToHandler(sync)))
}

func (c *volumeAttachmentController) Enqueue(name string) {
	c.controller.Enqueue("", name)
}

func (c *volumeAttachmentController) EnqueueAfter(name string, duration time.Duration) {
	c.controller.EnqueueAfter("", name, duration)
}

func (c *volumeAttachmentController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *volumeAttachmentController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *volumeAttachmentController) Cache() VolumeAttachmentCache {
	return &volumeAttachmentCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *volumeAttachmentController) Create(obj *v1.VolumeAttachment) (*v1.VolumeAttachment, error) {
	result := &v1.VolumeAttachment{}
	return result, c.client.Create(context.TODO(), "", obj, result, metav1.CreateOptions{})
}

func (c *volumeAttachmentController) Update(obj *v1.VolumeAttachment) (*v1.VolumeAttachment, error) {
	result := &v1.VolumeAttachment{}
	return result, c.client.Update(context.TODO(), "", obj, result, metav1.UpdateOptions{})
}

func (c *volumeAttachmentController) UpdateStatus(obj *v1.VolumeAttachment) (*v1.VolumeAttachment, error) {
	result := &v1.VolumeAttachment{}
	return result, c.client.UpdateStatus(context.TODO(), "", obj, result, metav1.UpdateOptions{})
}

func (c *volumeAttachmentController) Delete(name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), "", name, *options)
}

func (c *volumeAttachmentController) Get(name string, options metav1.GetOptions) (*v1.VolumeAttachment, error) {
	result := &v1.VolumeAttachment{}
	return result, c.client.Get(context.TODO(), "", name, result, options)
}

func (c *volumeAttachmentController) List(opts metav1.ListOptions) (*v1.VolumeAttachmentList, error) {
	result := &v1.VolumeAttachmentList{}
	return result, c.client.List(context.TODO(), "", result, opts)
}

func (c *volumeAttachmentController) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), "", opts)
}

func (c *volumeAttachmentController) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (*v1.VolumeAttachment, error) {
	result := &v1.VolumeAttachment{}
	return result, c.client.Patch(context.TODO(), "", name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type volumeAttachmentCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *volumeAttachmentCache) Get(name string) (*v1.VolumeAttachment, error) {
	obj, exists, err := c.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1.VolumeAttachment), nil
}

func (c *volumeAttachmentCache) List(selector labels.Selector) (ret []*v1.VolumeAttachment, err error) {

	err = cache.ListAll(c.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.VolumeAttachment))
	})

	return ret, err
}

func (c *volumeAttachmentCache) AddIndexer(indexName string, indexer VolumeAttachmentIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1.VolumeAttachment))
		},
	}))
}

func (c *volumeAttachmentCache) GetByIndex(indexName, key string) (result []*v1.VolumeAttachment, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1.VolumeAttachment, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1.VolumeAttachment))
	}
	return result, nil
}

type VolumeAttachmentStatusHandler func(obj *v1.VolumeAttachment, status v1.VolumeAttachmentStatus) (v1.VolumeAttachmentStatus, error)

type VolumeAttachmentGeneratingHandler func(obj *v1.VolumeAttachment, status v1.VolumeAttachmentStatus) ([]runtime.Object, v1.VolumeAttachmentStatus, error)

func RegisterVolumeAttachmentStatusHandler(ctx context.Context, controller VolumeAttachmentController, condition condition.Cond, name string, handler VolumeAttachmentStatusHandler) {
	statusHandler := &volumeAttachmentStatusHandler{
		client:    controller,
		condition: condition,
		handler:   handler,
	}
	controller.AddGenericHandler(ctx, name, FromVolumeAttachmentHandlerToHandler(statusHandler.sync))
}

func RegisterVolumeAttachmentGeneratingHandler(ctx context.Context, controller VolumeAttachmentController, apply apply.Apply,
	condition condition.Cond, name string, handler VolumeAttachmentGeneratingHandler, opts *generic.GeneratingHandlerOptions) {
	statusHandler := &volumeAttachmentGeneratingHandler{
		VolumeAttachmentGeneratingHandler: handler,
		apply:                             apply,
		name:                              name,
		gvk:                               controller.GroupVersionKind(),
	}
	if opts != nil {
		statusHandler.opts = *opts
	}
	controller.OnChange(ctx, name, statusHandler.Remove)
	RegisterVolumeAttachmentStatusHandler(ctx, controller, condition, name, statusHandler.Handle)
}

type volumeAttachmentStatusHandler struct {
	client    VolumeAttachmentClient
	condition condition.Cond
	handler   VolumeAttachmentStatusHandler
}

func (a *volumeAttachmentStatusHandler) sync(key string, obj *v1.VolumeAttachment) (*v1.VolumeAttachment, error) {
	if obj == nil {
		return obj, nil
	}

	origStatus := obj.Status.DeepCopy()
	obj = obj.DeepCopy()
	newStatus, err := a.handler(obj, obj.Status)
	if err != nil {
		// Revert to old status on error
		newStatus = *origStatus.DeepCopy()
	}

	if a.condition != "" {
		if errors.IsConflict(err) {
			a.condition.SetError(&newStatus, "", nil)
		} else {
			a.condition.SetError(&newStatus, "", err)
		}
	}
	if !equality.Semantic.DeepEqual(origStatus, &newStatus) {
		if a.condition != "" {
			// Since status has changed, update the lastUpdatedTime
			a.condition.LastUpdated(&newStatus, time.Now().UTC().Format(time.RFC3339))
		}

		var newErr error
		obj.Status = newStatus
		newObj, newErr := a.client.UpdateStatus(obj)
		if err == nil {
			err = newErr
		}
		if newErr == nil {
			obj = newObj
		}
	}
	return obj, err
}

type volumeAttachmentGeneratingHandler struct {
	VolumeAttachmentGeneratingHandler
	apply apply.Apply
	opts  generic.GeneratingHandlerOptions
	gvk   schema.GroupVersionKind
	name  string
}

func (a *volumeAttachmentGeneratingHandler) Remove(key string, obj *v1.VolumeAttachment) (*v1.VolumeAttachment, error) {
	if obj != nil {
		return obj, nil
	}

	obj = &v1.VolumeAttachment{}
	obj.Namespace, obj.Name = kv.RSplit(key, "/")
	obj.SetGroupVersionKind(a.gvk)

	return nil, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects()
}

func (a *volumeAttachmentGeneratingHandler) Handle(obj *v1.VolumeAttachment, status v1.VolumeAttachmentStatus) (v1.VolumeAttachmentStatus, error) {
	if !obj.DeletionTimestamp.IsZero() {
		return status, nil
	}

	objs, newStatus, err := a.VolumeAttachmentGeneratingHandler(obj, status)
	if err != nil {
		return newStatus, err
	}

	return newStatus, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects(objs...)
}
