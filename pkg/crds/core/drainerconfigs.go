package core

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"sigs.k8s.io/yaml"
)

const drainerconfigsYAML = `
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: (devel)
  creationTimestamp: null
  name: drainerconfigs.core.giantswarm.io
spec:
  group: core.giantswarm.io
  names:
    kind: DrainerConfig
    listKind: DrainerConfigList
    plural: drainerconfigs
    singular: drainerconfig
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
                cluster:
                  properties:
                    api:
                      properties:
                        endpoint:
                          description: Endpoint is the guest cluster API endpoint.
                          type: string
                      required:
                      - endpoint
                      type: object
                    id:
                      description: ID is the guest cluster ID of which a node should
                        be drained.
                      type: string
                  required:
                  - api
                  - id
                  type: object
                node:
                  properties:
                    name:
                      description: "Name is the identifier of the guest cluster's
                        master and worker nodes. In Kubernetes/Kubectl they are represented
                        as node names. The names are manage in an abstracted way because
                        of provider specific differences. \n     AWS: EC2 instance
                        DNS.     Azure: VM name.     KVM: host cluster pod name."
                      type: string
                  required:
                  - name
                  type: object
              required:
              - cluster
              - node
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
        status:
          properties:
            conditions:
              items:
                description: DrainerConfigStatusCondition expresses a condition in
                  which a node may is.
                properties:
                  lastHeartbeatTime:
                    description: LastHeartbeatTime is the last time we got an update
                      on a given condition.
                    format: date-time
                    type: string
                  lastTransitionTime:
                    description: LastTransitionTime is the last time the condition
                      transitioned from one status to another.
                    format: date-time
                    type: string
                  status:
                    description: Status may be True, False or Unknown.
                    type: string
                  type:
                    description: Type may be Pending, Ready, Draining, Drained.
                    type: string
                required:
                - lastHeartbeatTime
                - lastTransitionTime
                - status
                - type
                type: object
              type: array
          required:
          - conditions
          type: object
      required:
      - metadata
      - spec
      - status
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

func NewDrainerConfigCRD() *v1beta1.CustomResourceDefinition {
	var crd v1beta1.CustomResourceDefinition
	_ = yaml.Unmarshal([]byte(drainerconfigsYAML), &crd)
	return &crd
}
