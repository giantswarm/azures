package core

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"sigs.k8s.io/yaml"
)

const clustersYAML = `
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: (devel)
  creationTimestamp: null
  name: clusters.core.giantswarm.io
spec:
  group: core.giantswarm.io
  names:
    kind: Cluster
    listKind: ClusterList
    plural: clusters
    singular: cluster
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
          description: ClusterSpec is the part of the interface available to users
            in order to request a tenant cluster creation by providing necessary configurations.
            Fields here are either mandatory or optional. Optional fields left blank
            will be filled with appropriate default values which are then propagated
            into the CR status.
          properties:
            description:
              description: Description is the optional cluster description users can
                provide. If left blank a cluster description will be generated. The
                cluster description is propagated into the CR status.
              type: string
            organization:
              description: Organization is the mandatory cluster organization in which
                a tenant cluster will be scoped into.
              type: string
            version:
              description: Version is the optional release version users can provide.
                If left blank the current default release version will be used. The
                release version is propagated into the CR status.
              type: string
          required:
          - description
          - organization
          - version
          type: object
        status:
          description: ClusterStatus is the part of the interface available to users
            in order to fetch a tenant cluster's status information after creation
            was requested. Fields here are automatically filled and can only ever
            be read. For instance the tenant cluster description will be generated
            if left blank upon cluster creation and made available here.
          properties:
            cluster:
              description: Cluster holds cluster specific status information.
              properties:
                description:
                  description: Description is the propagated cluster description users
                    can provide or the system generates automatically if left blank.
                  type: string
                id:
                  description: ID is the internal cluster ID automatically generated
                    upon cluster creation.
                  type: string
                version:
                  description: Version is the propagated release version users can
                    provide or the system sets to the current default release version.
                  type: string
              required:
              - description
              - id
              - version
              type: object
            conditions:
              description: Conditions is a list of status conditions.
              items:
                description: ClusterStatusCondition holds a specific status condition
                  describing certain aspects of the current state of the tenant cluster.
                properties:
                  status:
                    description: Status may be True, False or Unknown.
                    type: string
                  type:
                    description: Type may be Creating, Created, Updating, Updated,
                      or Deleting.
                    type: string
                required:
                - status
                - type
                type: object
              type: array
            lastHeartbeatTime:
              description: LastHeartbeatTime is the last time we got an update on
                a given condition.
              format: date-time
              type: string
            lastTransitionTime:
              description: LastTransitionTime is the last time the condition transitioned
                from one status to another.
              format: date-time
              type: string
          required:
          - cluster
          - conditions
          - lastHeartbeatTime
          - lastTransitionTime
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

func NewClusterCRD() *v1beta1.CustomResourceDefinition {
	var crd v1beta1.CustomResourceDefinition
	_ = yaml.Unmarshal([]byte(clustersYAML), &crd)
	return &crd
}
