/*
Copyright 2019 Giant Swarm GmbH.

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
	"time"

	v1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/example/v1alpha1"
	scheme "github.com/giantswarm/apiextensions/pkg/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ChaosMonkeyConfigsGetter has a method to return a ChaosMonkeyConfigInterface.
// A group's client should implement this interface.
type ChaosMonkeyConfigsGetter interface {
	ChaosMonkeyConfigs(namespace string) ChaosMonkeyConfigInterface
}

// ChaosMonkeyConfigInterface has methods to work with ChaosMonkeyConfig resources.
type ChaosMonkeyConfigInterface interface {
	Create(*v1alpha1.ChaosMonkeyConfig) (*v1alpha1.ChaosMonkeyConfig, error)
	Update(*v1alpha1.ChaosMonkeyConfig) (*v1alpha1.ChaosMonkeyConfig, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ChaosMonkeyConfig, error)
	List(opts v1.ListOptions) (*v1alpha1.ChaosMonkeyConfigList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ChaosMonkeyConfig, err error)
	ChaosMonkeyConfigExpansion
}

// chaosMonkeyConfigs implements ChaosMonkeyConfigInterface
type chaosMonkeyConfigs struct {
	client rest.Interface
	ns     string
}

// newChaosMonkeyConfigs returns a ChaosMonkeyConfigs
func newChaosMonkeyConfigs(c *ExampleV1alpha1Client, namespace string) *chaosMonkeyConfigs {
	return &chaosMonkeyConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the chaosMonkeyConfig, and returns the corresponding chaosMonkeyConfig object, and an error if there is any.
func (c *chaosMonkeyConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.ChaosMonkeyConfig, err error) {
	result = &v1alpha1.ChaosMonkeyConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("chaosmonkeyconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ChaosMonkeyConfigs that match those selectors.
func (c *chaosMonkeyConfigs) List(opts v1.ListOptions) (result *v1alpha1.ChaosMonkeyConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ChaosMonkeyConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("chaosmonkeyconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested chaosMonkeyConfigs.
func (c *chaosMonkeyConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("chaosmonkeyconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a chaosMonkeyConfig and creates it.  Returns the server's representation of the chaosMonkeyConfig, and an error, if there is any.
func (c *chaosMonkeyConfigs) Create(chaosMonkeyConfig *v1alpha1.ChaosMonkeyConfig) (result *v1alpha1.ChaosMonkeyConfig, err error) {
	result = &v1alpha1.ChaosMonkeyConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("chaosmonkeyconfigs").
		Body(chaosMonkeyConfig).
		Do().
		Into(result)
	return
}

// Update takes the representation of a chaosMonkeyConfig and updates it. Returns the server's representation of the chaosMonkeyConfig, and an error, if there is any.
func (c *chaosMonkeyConfigs) Update(chaosMonkeyConfig *v1alpha1.ChaosMonkeyConfig) (result *v1alpha1.ChaosMonkeyConfig, err error) {
	result = &v1alpha1.ChaosMonkeyConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("chaosmonkeyconfigs").
		Name(chaosMonkeyConfig.Name).
		Body(chaosMonkeyConfig).
		Do().
		Into(result)
	return
}

// Delete takes name of the chaosMonkeyConfig and deletes it. Returns an error if one occurs.
func (c *chaosMonkeyConfigs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("chaosmonkeyconfigs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *chaosMonkeyConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("chaosmonkeyconfigs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched chaosMonkeyConfig.
func (c *chaosMonkeyConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ChaosMonkeyConfig, err error) {
	result = &v1alpha1.ChaosMonkeyConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("chaosmonkeyconfigs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
