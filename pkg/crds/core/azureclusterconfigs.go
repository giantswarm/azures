package core

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"sigs.k8s.io/yaml"
)

const azureclusterconfigsYAML = `
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: (devel)
  creationTimestamp: null
  name: azureclusterconfigs.core.giantswarm.io
spec:
  group: core.giantswarm.io
  names:
    kind: AzureClusterConfig
    listKind: AzureClusterConfigList
    plural: azureclusterconfigs
    singular: azureclusterconfig
  scope: Namespaced
  validation:
    openAPIV3Schema:
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
          type: string
        metadata:
          type: object
        spec:
          properties:
            guest:
              properties:
                availabilityZones:
                  type: integer
                credentialSecret:
                  description: AzureClusterConfigSpecGuestCredentialSecret points
                    to the K8s Secret containing credentials for an Azure subscription
                    in which the tenant cluster should be created.
                  properties:
                    name:
                      type: string
                    namespace:
                      type: string
                  required:
                  - name
                  - namespace
                  type: object
                dnsZone:
                  description: DNSZone for guest cluster is supplemented with host
                    prefixes for specific services such as Kubernetes API or Etcd.
                    In general this DNS Zone should start with "k8s" like for example
                    "k8s.cluster.example.com.".
                  type: string
                id:
                  type: string
                masters:
                  items:
                    properties:
                      id:
                        type: string
                      vmSize:
                        type: string
                    required:
                    - id
                    type: object
                  type: array
                name:
                  type: string
                owner:
                  type: string
                releaseVersion:
                  type: string
                versionBundles:
                  items:
                    properties:
                      name:
                        type: string
                      version:
                        type: string
                    required:
                    - name
                    - version
                    type: object
                  type: array
                workers:
                  items:
                    properties:
                      id:
                        type: string
                      labels:
                        additionalProperties:
                          type: string
                        type: object
                      vmSize:
                        type: string
                    required:
                    - id
                    - labels
                    type: object
                  type: array
              required:
              - credentialSecret
              - dnsZone
              - id
              type: object
            versionBundle:
              properties:
                version:
                  type: string
              required:
              - version
              type: object
          required:
          - guest
          - versionBundle
          type: object
      required:
      - metadata
      - spec
      type: object
  version: v1alpha1
  versions:
  - name: v1alpha1
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`

func NewAzureClusterConfigCRD() *v1beta1.CustomResourceDefinition {
	var crd v1beta1.CustomResourceDefinition
	_ = yaml.Unmarshal([]byte(azureclusterconfigsYAML), &crd)
	return &crd
}