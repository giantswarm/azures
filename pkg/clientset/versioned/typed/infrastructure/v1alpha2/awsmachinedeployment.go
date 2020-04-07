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

// AWSMachineDeploymentsGetter has a method to return a AWSMachineDeploymentInterface.
// A group's client should implement this interface.
type AWSMachineDeploymentsGetter interface {
	AWSMachineDeployments(namespace string) AWSMachineDeploymentInterface
}

// AWSMachineDeploymentInterface has methods to work with AWSMachineDeployment resources.
type AWSMachineDeploymentInterface interface {
	Create(*v1alpha2.AWSMachineDeployment) (*v1alpha2.AWSMachineDeployment, error)
	Update(*v1alpha2.AWSMachineDeployment) (*v1alpha2.AWSMachineDeployment, error)
	UpdateStatus(*v1alpha2.AWSMachineDeployment) (*v1alpha2.AWSMachineDeployment, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha2.AWSMachineDeployment, error)
	List(opts v1.ListOptions) (*v1alpha2.AWSMachineDeploymentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.AWSMachineDeployment, err error)
	AWSMachineDeploymentExpansion
}

// aWSMachineDeployments implements AWSMachineDeploymentInterface
type aWSMachineDeployments struct {
	client rest.Interface
	ns     string
}

// newAWSMachineDeployments returns a AWSMachineDeployments
func newAWSMachineDeployments(c *InfrastructureV1alpha2Client, namespace string) *aWSMachineDeployments {
	return &aWSMachineDeployments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the aWSMachineDeployment, and returns the corresponding aWSMachineDeployment object, and an error if there is any.
func (c *aWSMachineDeployments) Get(name string, options v1.GetOptions) (result *v1alpha2.AWSMachineDeployment, err error) {
	result = &v1alpha2.AWSMachineDeployment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsmachinedeployments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AWSMachineDeployments that match those selectors.
func (c *aWSMachineDeployments) List(opts v1.ListOptions) (result *v1alpha2.AWSMachineDeploymentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.AWSMachineDeploymentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsmachinedeployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested aWSMachineDeployments.
func (c *aWSMachineDeployments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsmachinedeployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a aWSMachineDeployment and creates it.  Returns the server's representation of the aWSMachineDeployment, and an error, if there is any.
func (c *aWSMachineDeployments) Create(aWSMachineDeployment *v1alpha2.AWSMachineDeployment) (result *v1alpha2.AWSMachineDeployment, err error) {
	result = &v1alpha2.AWSMachineDeployment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsmachinedeployments").
		Body(aWSMachineDeployment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a aWSMachineDeployment and updates it. Returns the server's representation of the aWSMachineDeployment, and an error, if there is any.
func (c *aWSMachineDeployments) Update(aWSMachineDeployment *v1alpha2.AWSMachineDeployment) (result *v1alpha2.AWSMachineDeployment, err error) {
	result = &v1alpha2.AWSMachineDeployment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsmachinedeployments").
		Name(aWSMachineDeployment.Name).
		Body(aWSMachineDeployment).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *aWSMachineDeployments) UpdateStatus(aWSMachineDeployment *v1alpha2.AWSMachineDeployment) (result *v1alpha2.AWSMachineDeployment, err error) {
	result = &v1alpha2.AWSMachineDeployment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsmachinedeployments").
		Name(aWSMachineDeployment.Name).
		SubResource("status").
		Body(aWSMachineDeployment).
		Do().
		Into(result)
	return
}

// Delete takes name of the aWSMachineDeployment and deletes it. Returns an error if one occurs.
func (c *aWSMachineDeployments) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsmachinedeployments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *aWSMachineDeployments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsmachinedeployments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched aWSMachineDeployment.
func (c *aWSMachineDeployments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.AWSMachineDeployment, err error) {
	result = &v1alpha2.AWSMachineDeployment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsmachinedeployments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
