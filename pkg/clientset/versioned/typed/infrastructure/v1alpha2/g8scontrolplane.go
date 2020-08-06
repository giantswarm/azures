/*
Copyright 2020 Giant Swarm GmbH.

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

package v1alpha2

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	v1alpha2 "github.com/giantswarm/apiextensions/v2/pkg/apis/infrastructure/v1alpha2"
	scheme "github.com/giantswarm/apiextensions/v2/pkg/clientset/versioned/scheme"
)

// G8sControlPlanesGetter has a method to return a G8sControlPlaneInterface.
// A group's client should implement this interface.
type G8sControlPlanesGetter interface {
	G8sControlPlanes(namespace string) G8sControlPlaneInterface
}

// G8sControlPlaneInterface has methods to work with G8sControlPlane resources.
type G8sControlPlaneInterface interface {
	Create(ctx context.Context, g8sControlPlane *v1alpha2.G8sControlPlane, opts v1.CreateOptions) (*v1alpha2.G8sControlPlane, error)
	Update(ctx context.Context, g8sControlPlane *v1alpha2.G8sControlPlane, opts v1.UpdateOptions) (*v1alpha2.G8sControlPlane, error)
	UpdateStatus(ctx context.Context, g8sControlPlane *v1alpha2.G8sControlPlane, opts v1.UpdateOptions) (*v1alpha2.G8sControlPlane, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha2.G8sControlPlane, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha2.G8sControlPlaneList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.G8sControlPlane, err error)
	G8sControlPlaneExpansion
}

// g8sControlPlanes implements G8sControlPlaneInterface
type g8sControlPlanes struct {
	client rest.Interface
	ns     string
}

// newG8sControlPlanes returns a G8sControlPlanes
func newG8sControlPlanes(c *InfrastructureV1alpha2Client, namespace string) *g8sControlPlanes {
	return &g8sControlPlanes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the g8sControlPlane, and returns the corresponding g8sControlPlane object, and an error if there is any.
func (c *g8sControlPlanes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.G8sControlPlane, err error) {
	result = &v1alpha2.G8sControlPlane{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("g8scontrolplanes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of G8sControlPlanes that match those selectors.
func (c *g8sControlPlanes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.G8sControlPlaneList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.G8sControlPlaneList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("g8scontrolplanes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested g8sControlPlanes.
func (c *g8sControlPlanes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("g8scontrolplanes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a g8sControlPlane and creates it.  Returns the server's representation of the g8sControlPlane, and an error, if there is any.
func (c *g8sControlPlanes) Create(ctx context.Context, g8sControlPlane *v1alpha2.G8sControlPlane, opts v1.CreateOptions) (result *v1alpha2.G8sControlPlane, err error) {
	result = &v1alpha2.G8sControlPlane{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("g8scontrolplanes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(g8sControlPlane).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a g8sControlPlane and updates it. Returns the server's representation of the g8sControlPlane, and an error, if there is any.
func (c *g8sControlPlanes) Update(ctx context.Context, g8sControlPlane *v1alpha2.G8sControlPlane, opts v1.UpdateOptions) (result *v1alpha2.G8sControlPlane, err error) {
	result = &v1alpha2.G8sControlPlane{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("g8scontrolplanes").
		Name(g8sControlPlane.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(g8sControlPlane).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *g8sControlPlanes) UpdateStatus(ctx context.Context, g8sControlPlane *v1alpha2.G8sControlPlane, opts v1.UpdateOptions) (result *v1alpha2.G8sControlPlane, err error) {
	result = &v1alpha2.G8sControlPlane{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("g8scontrolplanes").
		Name(g8sControlPlane.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(g8sControlPlane).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the g8sControlPlane and deletes it. Returns an error if one occurs.
func (c *g8sControlPlanes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("g8scontrolplanes").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *g8sControlPlanes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("g8scontrolplanes").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched g8sControlPlane.
func (c *g8sControlPlanes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.G8sControlPlane, err error) {
	result = &v1alpha2.G8sControlPlane{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("g8scontrolplanes").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
