// Copyright 2023 Ant Group Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/secretflow/kuscia/pkg/crd/apis/kuscia/v1alpha1"
	scheme "github.com/secretflow/kuscia/pkg/crd/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// InteropConfigsGetter has a method to return a InteropConfigInterface.
// A group's client should implement this interface.
type InteropConfigsGetter interface {
	InteropConfigs() InteropConfigInterface
}

// InteropConfigInterface has methods to work with InteropConfig resources.
type InteropConfigInterface interface {
	Create(ctx context.Context, interopConfig *v1alpha1.InteropConfig, opts v1.CreateOptions) (*v1alpha1.InteropConfig, error)
	Update(ctx context.Context, interopConfig *v1alpha1.InteropConfig, opts v1.UpdateOptions) (*v1alpha1.InteropConfig, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.InteropConfig, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.InteropConfigList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.InteropConfig, err error)
	InteropConfigExpansion
}

// interopConfigs implements InteropConfigInterface
type interopConfigs struct {
	client rest.Interface
}

// newInteropConfigs returns a InteropConfigs
func newInteropConfigs(c *KusciaV1alpha1Client) *interopConfigs {
	return &interopConfigs{
		client: c.RESTClient(),
	}
}

// Get takes name of the interopConfig, and returns the corresponding interopConfig object, and an error if there is any.
func (c *interopConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.InteropConfig, err error) {
	result = &v1alpha1.InteropConfig{}
	err = c.client.Get().
		Resource("interopconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of InteropConfigs that match those selectors.
func (c *interopConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.InteropConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.InteropConfigList{}
	err = c.client.Get().
		Resource("interopconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested interopConfigs.
func (c *interopConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("interopconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a interopConfig and creates it.  Returns the server's representation of the interopConfig, and an error, if there is any.
func (c *interopConfigs) Create(ctx context.Context, interopConfig *v1alpha1.InteropConfig, opts v1.CreateOptions) (result *v1alpha1.InteropConfig, err error) {
	result = &v1alpha1.InteropConfig{}
	err = c.client.Post().
		Resource("interopconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(interopConfig).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a interopConfig and updates it. Returns the server's representation of the interopConfig, and an error, if there is any.
func (c *interopConfigs) Update(ctx context.Context, interopConfig *v1alpha1.InteropConfig, opts v1.UpdateOptions) (result *v1alpha1.InteropConfig, err error) {
	result = &v1alpha1.InteropConfig{}
	err = c.client.Put().
		Resource("interopconfigs").
		Name(interopConfig.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(interopConfig).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the interopConfig and deletes it. Returns an error if one occurs.
func (c *interopConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("interopconfigs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *interopConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("interopconfigs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched interopConfig.
func (c *interopConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.InteropConfig, err error) {
	result = &v1alpha1.InteropConfig{}
	err = c.client.Patch(pt).
		Resource("interopconfigs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
