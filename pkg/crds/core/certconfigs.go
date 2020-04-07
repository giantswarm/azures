package core

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"sigs.k8s.io/yaml"
)

const certconfigsYAML = `
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: (devel)
  creationTimestamp: null
  name: certconfigs.core.giantswarm.io
spec:
  group: core.giantswarm.io
  names:
    kind: CertConfig
    listKind: CertConfigList
    plural: certconfigs
    singular: certconfig
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
            cert:
              properties:
                allowBareDomains:
                  type: boolean
                altNames:
                  items:
                    type: string
                  type: array
                clusterComponent:
                  type: string
                clusterID:
                  type: string
                commonName:
                  type: string
                disableRegeneration:
                  type: boolean
                ipSans:
                  items:
                    type: string
                  type: array
                organizations:
                  items:
                    type: string
                  type: array
                ttl:
                  type: string
              required:
              - allowBareDomains
              - altNames
              - clusterComponent
              - clusterID
              - commonName
              - disableRegeneration
              - ipSans
              - organizations
              - ttl
              type: object
            versionBundle:
              properties:
                version:
                  type: string
              required:
              - version
              type: object
          required:
          - cert
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

func NewCertConfigCRD() *v1beta1.CustomResourceDefinition {
	var crd v1beta1.CustomResourceDefinition
	_ = yaml.Unmarshal([]byte(certconfigsYAML), &crd)
	return &crd
}
