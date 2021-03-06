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

package v1alpha2

import (
	rest "k8s.io/client-go/rest"

	v1alpha2 "github.com/giantswarm/apiextensions/v3/pkg/apis/infrastructure/v1alpha2"
	"github.com/giantswarm/apiextensions/v3/pkg/clientset/versioned/scheme"
)

type InfrastructureV1alpha2Interface interface {
	RESTClient() rest.Interface
	AWSClustersGetter
	AWSControlPlanesGetter
	AWSMachineDeploymentsGetter
	G8sControlPlanesGetter
	NetworkPoolsGetter
}

// InfrastructureV1alpha2Client is used to interact with features provided by the infrastructure.giantswarm.io group.
type InfrastructureV1alpha2Client struct {
	restClient rest.Interface
}

func (c *InfrastructureV1alpha2Client) AWSClusters(namespace string) AWSClusterInterface {
	return newAWSClusters(c, namespace)
}

func (c *InfrastructureV1alpha2Client) AWSControlPlanes(namespace string) AWSControlPlaneInterface {
	return newAWSControlPlanes(c, namespace)
}

func (c *InfrastructureV1alpha2Client) AWSMachineDeployments(namespace string) AWSMachineDeploymentInterface {
	return newAWSMachineDeployments(c, namespace)
}

func (c *InfrastructureV1alpha2Client) G8sControlPlanes(namespace string) G8sControlPlaneInterface {
	return newG8sControlPlanes(c, namespace)
}

func (c *InfrastructureV1alpha2Client) NetworkPools(namespace string) NetworkPoolInterface {
	return newNetworkPools(c, namespace)
}

// NewForConfig creates a new InfrastructureV1alpha2Client for the given config.
func NewForConfig(c *rest.Config) (*InfrastructureV1alpha2Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &InfrastructureV1alpha2Client{client}, nil
}

// NewForConfigOrDie creates a new InfrastructureV1alpha2Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *InfrastructureV1alpha2Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new InfrastructureV1alpha2Client for the given RESTClient.
func New(c rest.Interface) *InfrastructureV1alpha2Client {
	return &InfrastructureV1alpha2Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha2.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *InfrastructureV1alpha2Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
