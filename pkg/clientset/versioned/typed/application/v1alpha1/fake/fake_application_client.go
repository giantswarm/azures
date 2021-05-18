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

package fake

import (
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"

	v1alpha1 "github.com/giantswarm/apiextensions/v3/pkg/clientset/versioned/typed/application/v1alpha1"
)

type FakeApplicationV1alpha1 struct {
	*testing.Fake
}

func (c *FakeApplicationV1alpha1) Apps(namespace string) v1alpha1.AppInterface {
	return &FakeApps{c, namespace}
}

func (c *FakeApplicationV1alpha1) AppCatalogs() v1alpha1.AppCatalogInterface {
	return &FakeAppCatalogs{c}
}

func (c *FakeApplicationV1alpha1) AppCatalogEntries(namespace string) v1alpha1.AppCatalogEntryInterface {
	return &FakeAppCatalogEntries{c, namespace}
}

func (c *FakeApplicationV1alpha1) Catalogs(namespace string) v1alpha1.CatalogInterface {
	return &FakeCatalogs{c, namespace}
}

func (c *FakeApplicationV1alpha1) Charts(namespace string) v1alpha1.ChartInterface {
	return &FakeCharts{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeApplicationV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
