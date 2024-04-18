/*
Copyright The KubeStellar Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	v1alpha1 "github.com/kubestellar/kubestellar/api/control/v1alpha1"
	scheme "github.com/kubestellar/kubestellar/pkg/generated/clientset/versioned/scheme"
)

// BindingsGetter has a method to return a BindingInterface.
// A group's client should implement this interface.
type BindingsGetter interface {
	Bindings() BindingInterface
}

// BindingInterface has methods to work with Binding resources.
type BindingInterface interface {
	Create(ctx context.Context, binding *v1alpha1.Binding, opts v1.CreateOptions) (*v1alpha1.Binding, error)
	Update(ctx context.Context, binding *v1alpha1.Binding, opts v1.UpdateOptions) (*v1alpha1.Binding, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Binding, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.BindingList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Binding, err error)
	BindingExpansion
}

// bindings implements BindingInterface
type bindings struct {
	client rest.Interface
}

// newBindings returns a Bindings
func newBindings(c *ControlV1alpha1Client) *bindings {
	return &bindings{
		client: c.RESTClient(),
	}
}

// Get takes name of the binding, and returns the corresponding binding object, and an error if there is any.
func (c *bindings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Binding, err error) {
	result = &v1alpha1.Binding{}
	err = c.client.Get().
		Resource("bindings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Bindings that match those selectors.
func (c *bindings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BindingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.BindingList{}
	err = c.client.Get().
		Resource("bindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested bindings.
func (c *bindings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("bindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a binding and creates it.  Returns the server's representation of the binding, and an error, if there is any.
func (c *bindings) Create(ctx context.Context, binding *v1alpha1.Binding, opts v1.CreateOptions) (result *v1alpha1.Binding, err error) {
	result = &v1alpha1.Binding{}
	err = c.client.Post().
		Resource("bindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(binding).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a binding and updates it. Returns the server's representation of the binding, and an error, if there is any.
func (c *bindings) Update(ctx context.Context, binding *v1alpha1.Binding, opts v1.UpdateOptions) (result *v1alpha1.Binding, err error) {
	result = &v1alpha1.Binding{}
	err = c.client.Put().
		Resource("bindings").
		Name(binding.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(binding).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the binding and deletes it. Returns an error if one occurs.
func (c *bindings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("bindings").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *bindings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("bindings").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched binding.
func (c *bindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Binding, err error) {
	result = &v1alpha1.Binding{}
	err = c.client.Patch(pt).
		Resource("bindings").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}