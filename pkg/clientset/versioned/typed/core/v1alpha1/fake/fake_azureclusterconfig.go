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
//golint:ignore

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	v1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/core/v1alpha1"
)

// FakeAzureClusterConfigs implements AzureClusterConfigInterface
type FakeAzureClusterConfigs struct {
	Fake *FakeCoreV1alpha1
	ns   string
}

var azureclusterconfigsResource = schema.GroupVersionResource{Group: "core.giantswarm.io", Version: "v1alpha1", Resource: "azureclusterconfigs"}

var azureclusterconfigsKind = schema.GroupVersionKind{Group: "core.giantswarm.io", Version: "v1alpha1", Kind: "AzureClusterConfig"}

// Get takes name of the azureClusterConfig, and returns the corresponding azureClusterConfig object, and an error if there is any.
func (c *FakeAzureClusterConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.AzureClusterConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(azureclusterconfigsResource, c.ns, name), &v1alpha1.AzureClusterConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureClusterConfig), err
}

// List takes label and field selectors, and returns the list of AzureClusterConfigs that match those selectors.
func (c *FakeAzureClusterConfigs) List(opts v1.ListOptions) (result *v1alpha1.AzureClusterConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(azureclusterconfigsResource, azureclusterconfigsKind, c.ns, opts), &v1alpha1.AzureClusterConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzureClusterConfigList{ListMeta: obj.(*v1alpha1.AzureClusterConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzureClusterConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azureClusterConfigs.
func (c *FakeAzureClusterConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(azureclusterconfigsResource, c.ns, opts))

}

// Create takes the representation of a azureClusterConfig and creates it.  Returns the server's representation of the azureClusterConfig, and an error, if there is any.
func (c *FakeAzureClusterConfigs) Create(azureClusterConfig *v1alpha1.AzureClusterConfig) (result *v1alpha1.AzureClusterConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(azureclusterconfigsResource, c.ns, azureClusterConfig), &v1alpha1.AzureClusterConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureClusterConfig), err
}

// Update takes the representation of a azureClusterConfig and updates it. Returns the server's representation of the azureClusterConfig, and an error, if there is any.
func (c *FakeAzureClusterConfigs) Update(azureClusterConfig *v1alpha1.AzureClusterConfig) (result *v1alpha1.AzureClusterConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(azureclusterconfigsResource, c.ns, azureClusterConfig), &v1alpha1.AzureClusterConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureClusterConfig), err
}

// Delete takes name of the azureClusterConfig and deletes it. Returns an error if one occurs.
func (c *FakeAzureClusterConfigs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(azureclusterconfigsResource, c.ns, name), &v1alpha1.AzureClusterConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzureClusterConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(azureclusterconfigsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzureClusterConfigList{})
	return err
}

// Patch applies the patch and returns the patched azureClusterConfig.
func (c *FakeAzureClusterConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzureClusterConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(azureclusterconfigsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AzureClusterConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureClusterConfig), err
}
