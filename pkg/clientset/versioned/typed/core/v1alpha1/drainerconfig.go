/*
Copyright 2018 Giant Swarm GmbH.

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
	v1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/core/v1alpha1"
	scheme "github.com/giantswarm/apiextensions/pkg/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DrainerConfigsGetter has a method to return a DrainerConfigInterface.
// A group's client should implement this interface.
type DrainerConfigsGetter interface {
	DrainerConfigs(namespace string) DrainerConfigInterface
}

// DrainerConfigInterface has methods to work with DrainerConfig resources.
type DrainerConfigInterface interface {
	Create(*v1alpha1.DrainerConfig) (*v1alpha1.DrainerConfig, error)
	Update(*v1alpha1.DrainerConfig) (*v1alpha1.DrainerConfig, error)
	UpdateStatus(*v1alpha1.DrainerConfig) (*v1alpha1.DrainerConfig, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DrainerConfig, error)
	List(opts v1.ListOptions) (*v1alpha1.DrainerConfigList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DrainerConfig, err error)
	DrainerConfigExpansion
}

// drainerConfigs implements DrainerConfigInterface
type drainerConfigs struct {
	client rest.Interface
	ns     string
}

// newDrainerConfigs returns a DrainerConfigs
func newDrainerConfigs(c *CoreV1alpha1Client, namespace string) *drainerConfigs {
	return &drainerConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the drainerConfig, and returns the corresponding drainerConfig object, and an error if there is any.
func (c *drainerConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.DrainerConfig, err error) {
	result = &v1alpha1.DrainerConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("drainerconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DrainerConfigs that match those selectors.
func (c *drainerConfigs) List(opts v1.ListOptions) (result *v1alpha1.DrainerConfigList, err error) {
	result = &v1alpha1.DrainerConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("drainerconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested drainerConfigs.
func (c *drainerConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("drainerconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a drainerConfig and creates it.  Returns the server's representation of the drainerConfig, and an error, if there is any.
func (c *drainerConfigs) Create(drainerConfig *v1alpha1.DrainerConfig) (result *v1alpha1.DrainerConfig, err error) {
	result = &v1alpha1.DrainerConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("drainerconfigs").
		Body(drainerConfig).
		Do().
		Into(result)
	return
}

// Update takes the representation of a drainerConfig and updates it. Returns the server's representation of the drainerConfig, and an error, if there is any.
func (c *drainerConfigs) Update(drainerConfig *v1alpha1.DrainerConfig) (result *v1alpha1.DrainerConfig, err error) {
	result = &v1alpha1.DrainerConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("drainerconfigs").
		Name(drainerConfig.Name).
		Body(drainerConfig).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *drainerConfigs) UpdateStatus(drainerConfig *v1alpha1.DrainerConfig) (result *v1alpha1.DrainerConfig, err error) {
	result = &v1alpha1.DrainerConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("drainerconfigs").
		Name(drainerConfig.Name).
		SubResource("status").
		Body(drainerConfig).
		Do().
		Into(result)
	return
}

// Delete takes name of the drainerConfig and deletes it. Returns an error if one occurs.
func (c *drainerConfigs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("drainerconfigs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *drainerConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("drainerconfigs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched drainerConfig.
func (c *drainerConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DrainerConfig, err error) {
	result = &v1alpha1.DrainerConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("drainerconfigs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
