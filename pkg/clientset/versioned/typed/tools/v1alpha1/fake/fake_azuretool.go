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

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	v1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/tools/v1alpha1"
)

// FakeAzureTools implements AzureToolInterface
type FakeAzureTools struct {
	Fake *FakeToolsV1alpha1
	ns   string
}

var azuretoolsResource = schema.GroupVersionResource{Group: "tools.giantswarm.io", Version: "v1alpha1", Resource: "azuretools"}

var azuretoolsKind = schema.GroupVersionKind{Group: "tools.giantswarm.io", Version: "v1alpha1", Kind: "AzureTool"}

// Get takes name of the azureTool, and returns the corresponding azureTool object, and an error if there is any.
func (c *FakeAzureTools) Get(name string, options v1.GetOptions) (result *v1alpha1.AzureTool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(azuretoolsResource, c.ns, name), &v1alpha1.AzureTool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureTool), err
}

// List takes label and field selectors, and returns the list of AzureTools that match those selectors.
func (c *FakeAzureTools) List(opts v1.ListOptions) (result *v1alpha1.AzureToolList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(azuretoolsResource, azuretoolsKind, c.ns, opts), &v1alpha1.AzureToolList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzureToolList{ListMeta: obj.(*v1alpha1.AzureToolList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzureToolList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azureTools.
func (c *FakeAzureTools) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(azuretoolsResource, c.ns, opts))

}

// Create takes the representation of a azureTool and creates it.  Returns the server's representation of the azureTool, and an error, if there is any.
func (c *FakeAzureTools) Create(azureTool *v1alpha1.AzureTool) (result *v1alpha1.AzureTool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(azuretoolsResource, c.ns, azureTool), &v1alpha1.AzureTool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureTool), err
}

// Update takes the representation of a azureTool and updates it. Returns the server's representation of the azureTool, and an error, if there is any.
func (c *FakeAzureTools) Update(azureTool *v1alpha1.AzureTool) (result *v1alpha1.AzureTool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(azuretoolsResource, c.ns, azureTool), &v1alpha1.AzureTool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureTool), err
}

// Delete takes name of the azureTool and deletes it. Returns an error if one occurs.
func (c *FakeAzureTools) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(azuretoolsResource, c.ns, name), &v1alpha1.AzureTool{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzureTools) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(azuretoolsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzureToolList{})
	return err
}

// Patch applies the patch and returns the patched azureTool.
func (c *FakeAzureTools) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzureTool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(azuretoolsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AzureTool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureTool), err
}
