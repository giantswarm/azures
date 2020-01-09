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

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/backup/v1alpha1"
	scheme "github.com/giantswarm/apiextensions/pkg/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// EtcdBackupsGetter has a method to return a EtcdBackupInterface.
// A group's client should implement this interface.
type EtcdBackupsGetter interface {
	EtcdBackups() EtcdBackupInterface
}

// EtcdBackupInterface has methods to work with EtcdBackup resources.
type EtcdBackupInterface interface {
	Create(*v1alpha1.EtcdBackup) (*v1alpha1.EtcdBackup, error)
	Update(*v1alpha1.EtcdBackup) (*v1alpha1.EtcdBackup, error)
	UpdateStatus(*v1alpha1.EtcdBackup) (*v1alpha1.EtcdBackup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.EtcdBackup, error)
	List(opts v1.ListOptions) (*v1alpha1.EtcdBackupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EtcdBackup, err error)
	EtcdBackupExpansion
}

// etcdBackups implements EtcdBackupInterface
type etcdBackups struct {
	client rest.Interface
}

// newEtcdBackups returns a EtcdBackups
func newEtcdBackups(c *BackupV1alpha1Client) *etcdBackups {
	return &etcdBackups{
		client: c.RESTClient(),
	}
}

// Get takes name of the etcdBackup, and returns the corresponding etcdBackup object, and an error if there is any.
func (c *etcdBackups) Get(name string, options v1.GetOptions) (result *v1alpha1.EtcdBackup, err error) {
	result = &v1alpha1.EtcdBackup{}
	err = c.client.Get().
		Resource("etcdbackups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EtcdBackups that match those selectors.
func (c *etcdBackups) List(opts v1.ListOptions) (result *v1alpha1.EtcdBackupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.EtcdBackupList{}
	err = c.client.Get().
		Resource("etcdbackups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested etcdBackups.
func (c *etcdBackups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("etcdbackups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a etcdBackup and creates it.  Returns the server's representation of the etcdBackup, and an error, if there is any.
func (c *etcdBackups) Create(etcdBackup *v1alpha1.EtcdBackup) (result *v1alpha1.EtcdBackup, err error) {
	result = &v1alpha1.EtcdBackup{}
	err = c.client.Post().
		Resource("etcdbackups").
		Body(etcdBackup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a etcdBackup and updates it. Returns the server's representation of the etcdBackup, and an error, if there is any.
func (c *etcdBackups) Update(etcdBackup *v1alpha1.EtcdBackup) (result *v1alpha1.EtcdBackup, err error) {
	result = &v1alpha1.EtcdBackup{}
	err = c.client.Put().
		Resource("etcdbackups").
		Name(etcdBackup.Name).
		Body(etcdBackup).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *etcdBackups) UpdateStatus(etcdBackup *v1alpha1.EtcdBackup) (result *v1alpha1.EtcdBackup, err error) {
	result = &v1alpha1.EtcdBackup{}
	err = c.client.Put().
		Resource("etcdbackups").
		Name(etcdBackup.Name).
		SubResource("status").
		Body(etcdBackup).
		Do().
		Into(result)
	return
}

// Delete takes name of the etcdBackup and deletes it. Returns an error if one occurs.
func (c *etcdBackups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("etcdbackups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *etcdBackups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("etcdbackups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched etcdBackup.
func (c *etcdBackups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EtcdBackup, err error) {
	result = &v1alpha1.EtcdBackup{}
	err = c.client.Patch(pt).
		Resource("etcdbackups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
