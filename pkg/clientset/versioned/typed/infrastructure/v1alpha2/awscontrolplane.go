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

// Code generated by apiextensions/generator. DO NOT EDIT.

package v1alpha2

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	v1alpha2 "github.com/giantswarm/apiextensions/pkg/apis/infrastructure/v1alpha2"
	scheme "github.com/giantswarm/apiextensions/pkg/clientset/versioned/scheme"
)

// AWSControlPlanesGetter has a method to return a AWSControlPlaneInterface.
// A group's client should implement this interface.
type AWSControlPlanesGetter interface {
	AWSControlPlanes(namespace string) AWSControlPlaneInterface
}

// AWSControlPlaneInterface has methods to work with AWSControlPlane resources.
type AWSControlPlaneInterface interface {
	Create(*v1alpha2.AWSControlPlane) (*v1alpha2.AWSControlPlane, error)
	Update(*v1alpha2.AWSControlPlane) (*v1alpha2.AWSControlPlane, error)
	UpdateStatus(*v1alpha2.AWSControlPlane) (*v1alpha2.AWSControlPlane, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha2.AWSControlPlane, error)
	List(opts v1.ListOptions) (*v1alpha2.AWSControlPlaneList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.AWSControlPlane, err error)
	AWSControlPlaneExpansion
}

// aWSControlPlanes implements AWSControlPlaneInterface
type aWSControlPlanes struct {
	client rest.Interface
	ns     string
}

// newAWSControlPlanes returns a AWSControlPlanes
func newAWSControlPlanes(c *InfrastructureV1alpha2Client, namespace string) *aWSControlPlanes {
	return &aWSControlPlanes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the aWSControlPlane, and returns the corresponding aWSControlPlane object, and an error if there is any.
func (c *aWSControlPlanes) Get(name string, options v1.GetOptions) (result *v1alpha2.AWSControlPlane, err error) {
	result = &v1alpha2.AWSControlPlane{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awscontrolplanes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AWSControlPlanes that match those selectors.
func (c *aWSControlPlanes) List(opts v1.ListOptions) (result *v1alpha2.AWSControlPlaneList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.AWSControlPlaneList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awscontrolplanes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested aWSControlPlanes.
func (c *aWSControlPlanes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awscontrolplanes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a aWSControlPlane and creates it.  Returns the server's representation of the aWSControlPlane, and an error, if there is any.
func (c *aWSControlPlanes) Create(aWSControlPlane *v1alpha2.AWSControlPlane) (result *v1alpha2.AWSControlPlane, err error) {
	result = &v1alpha2.AWSControlPlane{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awscontrolplanes").
		Body(aWSControlPlane).
		Do().
		Into(result)
	return
}

// Update takes the representation of a aWSControlPlane and updates it. Returns the server's representation of the aWSControlPlane, and an error, if there is any.
func (c *aWSControlPlanes) Update(aWSControlPlane *v1alpha2.AWSControlPlane) (result *v1alpha2.AWSControlPlane, err error) {
	result = &v1alpha2.AWSControlPlane{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awscontrolplanes").
		Name(aWSControlPlane.Name).
		Body(aWSControlPlane).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *aWSControlPlanes) UpdateStatus(aWSControlPlane *v1alpha2.AWSControlPlane) (result *v1alpha2.AWSControlPlane, err error) {
	result = &v1alpha2.AWSControlPlane{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awscontrolplanes").
		Name(aWSControlPlane.Name).
		SubResource("status").
		Body(aWSControlPlane).
		Do().
		Into(result)
	return
}

// Delete takes name of the aWSControlPlane and deletes it. Returns an error if one occurs.
func (c *aWSControlPlanes) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awscontrolplanes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *aWSControlPlanes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awscontrolplanes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched aWSControlPlane.
func (c *aWSControlPlanes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.AWSControlPlane, err error) {
	result = &v1alpha2.AWSControlPlane{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awscontrolplanes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
