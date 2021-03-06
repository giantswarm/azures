/*
Copyright 2021 Giant Swarm GmbH.

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

	v1alpha1 "github.com/giantswarm/apiextensions/v3/pkg/apis/backup/v1alpha1"
	scheme "github.com/giantswarm/apiextensions/v3/pkg/clientset/versioned/scheme"
)

// ETCDBackupsGetter has a method to return a ETCDBackupInterface.
// A group's client should implement this interface.
type ETCDBackupsGetter interface {
	ETCDBackups() ETCDBackupInterface
}

// ETCDBackupInterface has methods to work with ETCDBackup resources.
type ETCDBackupInterface interface {
	Create(ctx context.Context, eTCDBackup *v1alpha1.ETCDBackup, opts v1.CreateOptions) (*v1alpha1.ETCDBackup, error)
	Update(ctx context.Context, eTCDBackup *v1alpha1.ETCDBackup, opts v1.UpdateOptions) (*v1alpha1.ETCDBackup, error)
	UpdateStatus(ctx context.Context, eTCDBackup *v1alpha1.ETCDBackup, opts v1.UpdateOptions) (*v1alpha1.ETCDBackup, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ETCDBackup, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ETCDBackupList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ETCDBackup, err error)
	ETCDBackupExpansion
}

// eTCDBackups implements ETCDBackupInterface
type eTCDBackups struct {
	client rest.Interface
}

// newETCDBackups returns a ETCDBackups
func newETCDBackups(c *BackupV1alpha1Client) *eTCDBackups {
	return &eTCDBackups{
		client: c.RESTClient(),
	}
}

// Get takes name of the eTCDBackup, and returns the corresponding eTCDBackup object, and an error if there is any.
func (c *eTCDBackups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ETCDBackup, err error) {
	result = &v1alpha1.ETCDBackup{}
	err = c.client.Get().
		Resource("etcdbackups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ETCDBackups that match those selectors.
func (c *eTCDBackups) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ETCDBackupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ETCDBackupList{}
	err = c.client.Get().
		Resource("etcdbackups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested eTCDBackups.
func (c *eTCDBackups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("etcdbackups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a eTCDBackup and creates it.  Returns the server's representation of the eTCDBackup, and an error, if there is any.
func (c *eTCDBackups) Create(ctx context.Context, eTCDBackup *v1alpha1.ETCDBackup, opts v1.CreateOptions) (result *v1alpha1.ETCDBackup, err error) {
	result = &v1alpha1.ETCDBackup{}
	err = c.client.Post().
		Resource("etcdbackups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(eTCDBackup).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a eTCDBackup and updates it. Returns the server's representation of the eTCDBackup, and an error, if there is any.
func (c *eTCDBackups) Update(ctx context.Context, eTCDBackup *v1alpha1.ETCDBackup, opts v1.UpdateOptions) (result *v1alpha1.ETCDBackup, err error) {
	result = &v1alpha1.ETCDBackup{}
	err = c.client.Put().
		Resource("etcdbackups").
		Name(eTCDBackup.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(eTCDBackup).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *eTCDBackups) UpdateStatus(ctx context.Context, eTCDBackup *v1alpha1.ETCDBackup, opts v1.UpdateOptions) (result *v1alpha1.ETCDBackup, err error) {
	result = &v1alpha1.ETCDBackup{}
	err = c.client.Put().
		Resource("etcdbackups").
		Name(eTCDBackup.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(eTCDBackup).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the eTCDBackup and deletes it. Returns an error if one occurs.
func (c *eTCDBackups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("etcdbackups").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *eTCDBackups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("etcdbackups").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched eTCDBackup.
func (c *eTCDBackups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ETCDBackup, err error) {
	result = &v1alpha1.ETCDBackup{}
	err = c.client.Patch(pt).
		Resource("etcdbackups").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
