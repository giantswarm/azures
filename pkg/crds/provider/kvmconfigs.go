package provider

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"sigs.k8s.io/yaml"
)

const kvmconfigsYAML = `
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: (devel)
  creationTimestamp: null
  name: kvmconfigs.provider.giantswarm.io
spec:
  group: provider.giantswarm.io
  names:
    kind: KVMConfig
    listKind: KVMConfigList
    plural: kvmconfigs
    singular: kvmconfig
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
            cluster:
              properties:
                calico:
                  properties:
                    cidr:
                      type: integer
                    mtu:
                      type: integer
                    subnet:
                      type: string
                  required:
                  - cidr
                  - mtu
                  - subnet
                  type: object
                customer:
                  properties:
                    id:
                      type: string
                  required:
                  - id
                  type: object
                docker:
                  properties:
                    daemon:
                      properties:
                        cidr:
                          type: string
                      required:
                      - cidr
                      type: object
                  required:
                  - daemon
                  type: object
                etcd:
                  properties:
                    altNames:
                      type: string
                    domain:
                      type: string
                    port:
                      type: integer
                    prefix:
                      type: string
                  required:
                  - altNames
                  - domain
                  - port
                  - prefix
                  type: object
                id:
                  type: string
                kubernetes:
                  properties:
                    api:
                      properties:
                        clusterIPRange:
                          type: string
                        domain:
                          type: string
                        securePort:
                          type: integer
                      required:
                      - clusterIPRange
                      - domain
                      - securePort
                      type: object
                    cloudProvider:
                      type: string
                    dns:
                      properties:
                        ip:
                          description: "An IP is a single IP address, a slice of bytes.
                            Functions in this package accept either 4-byte (IPv4)
                            or 16-byte (IPv6) slices as input. \n Note that in this
                            documentation, referring to an IP address as an IPv4 address
                            or an IPv6 address is a semantic property of the address,
                            not just the length of the byte slice: a 16-byte slice
                            can still be an IPv4 address."
                          format: byte
                          type: string
                      required:
                      - ip
                      type: object
                    domain:
                      type: string
                    ingressController:
                      properties:
                        docker:
                          properties:
                            image:
                              type: string
                          required:
                          - image
                          type: object
                        domain:
                          type: string
                        insecurePort:
                          type: integer
                        securePort:
                          type: integer
                        wildcardDomain:
                          type: string
                      required:
                      - docker
                      - domain
                      - insecurePort
                      - securePort
                      - wildcardDomain
                      type: object
                    kubelet:
                      properties:
                        altNames:
                          type: string
                        domain:
                          type: string
                        labels:
                          type: string
                        port:
                          type: integer
                      required:
                      - altNames
                      - domain
                      - labels
                      - port
                      type: object
                    networkSetup:
                      properties:
                        docker:
                          properties:
                            image:
                              type: string
                          required:
                          - image
                          type: object
                        kubeProxy:
                          description: ClusterKubernetesNetworkSetupKubeProxy describes
                            values passed to the kube-proxy running in a tenant cluster.
                          properties:
                            conntrackMaxPerCore:
                              description: Maximum number of NAT connections to track
                                per CPU core (0 to leave the limit as-is and ignore
                                conntrack-min). Passed to kube-proxy as --conntrack-max-per-core.
                              type: integer
                          required:
                          - conntrackMaxPerCore
                          type: object
                      required:
                      - docker
                      - kubeProxy
                      type: object
                    ssh:
                      properties:
                        userList:
                          items:
                            properties:
                              name:
                                type: string
                              publicKey:
                                type: string
                            required:
                            - name
                            - publicKey
                            type: object
                          type: array
                      required:
                      - userList
                      type: object
                  required:
                  - api
                  - cloudProvider
                  - dns
                  - domain
                  - ingressController
                  - kubelet
                  - networkSetup
                  - ssh
                  type: object
                masters:
                  items:
                    properties:
                      id:
                        type: string
                    required:
                    - id
                    type: object
                  type: array
                scaling:
                  properties:
                    max:
                      description: Max defines maximum number of worker nodes guest
                        cluster is allowed to have.
                      type: integer
                    min:
                      description: Min defines minimum number of worker nodes required
                        to be present in guest cluster.
                      type: integer
                  required:
                  - max
                  - min
                  type: object
                version:
                  description: Version is DEPRECATED and should just be dropped.
                  type: string
                workers:
                  items:
                    properties:
                      id:
                        type: string
                    required:
                    - id
                    type: object
                  type: array
              required:
              - calico
              - customer
              - docker
              - etcd
              - id
              - kubernetes
              - masters
              - scaling
              - version
              - workers
              type: object
            kvm:
              properties:
                endpointUpdater:
                  properties:
                    docker:
                      properties:
                        image:
                          type: string
                      required:
                      - image
                      type: object
                  required:
                  - docker
                  type: object
                k8sKVM:
                  properties:
                    docker:
                      properties:
                        image:
                          type: string
                      required:
                      - image
                      type: object
                    storageType:
                      type: string
                  required:
                  - docker
                  - storageType
                  type: object
                masters:
                  items:
                    properties:
                      cpus:
                        type: integer
                      disk:
                        type: string
                      dockerVolumeSizeGB:
                        type: integer
                      memory:
                        type: string
                    required:
                    - cpus
                    - disk
                    - dockerVolumeSizeGB
                    - memory
                    type: object
                  type: array
                network:
                  properties:
                    flannel:
                      properties:
                        vni:
                          type: integer
                      required:
                      - vni
                      type: object
                  required:
                  - flannel
                  type: object
                nodeController:
                  description: NOTE THIS IS DEPRECATED
                  properties:
                    docker:
                      description: NOTE THIS IS DEPRECATED
                      properties:
                        image:
                          type: string
                      required:
                      - image
                      type: object
                  required:
                  - docker
                  type: object
                portMappings:
                  items:
                    properties:
                      name:
                        type: string
                      nodePort:
                        type: integer
                      targetPort:
                        type: integer
                    required:
                    - name
                    - nodePort
                    - targetPort
                    type: object
                  type: array
                workers:
                  items:
                    properties:
                      cpus:
                        type: integer
                      disk:
                        type: string
                      dockerVolumeSizeGB:
                        type: integer
                      memory:
                        type: string
                    required:
                    - cpus
                    - disk
                    - dockerVolumeSizeGB
                    - memory
                    type: object
                  type: array
              required:
              - endpointUpdater
              - k8sKVM
              - masters
              - network
              - nodeController
              - portMappings
              - workers
              type: object
            versionBundle:
              properties:
                version:
                  type: string
              required:
              - version
              type: object
          required:
          - cluster
          - kvm
          - versionBundle
          type: object
        status:
          properties:
            cluster:
              properties:
                conditions:
                  description: Conditions is a list of status information expressing
                    the current conditional state of a guest cluster. This may reflect
                    the status of the guest cluster being updating or being up to
                    date.
                  items:
                    description: StatusClusterCondition expresses the conditions in
                      which a guest cluster may is.
                    properties:
                      lastTransitionTime:
                        description: LastTransitionTime is the last time the condition
                          transitioned from one status to another.
                        format: date-time
                        type: string
                      status:
                        description: Status may be True, False or Unknown.
                        type: string
                      type:
                        description: Type may be Creating, Created, Scaling, Scaled,
                          Draining, Drained, Updating, Updated, Deleting, Deleted.
                        type: string
                    required:
                    - lastTransitionTime
                    - status
                    - type
                    type: object
                  type: array
                network:
                  description: StatusClusterNetwork expresses the network segment
                    that is allocated for a guest cluster.
                  properties:
                    cidr:
                      type: string
                  required:
                  - cidr
                  type: object
                nodes:
                  description: Nodes is a list of guest cluster node information reflecting
                    the current state of the guest cluster nodes.
                  items:
                    description: StatusClusterNode holds information about a guest
                      cluster node.
                    properties:
                      labels:
                        additionalProperties:
                          type: string
                        description: Labels contains the kubernetes labels for corresponding
                          node.
                        type: object
                      lastTransitionTime:
                        description: LastTransitionTime is the last time the condition
                          transitioned from one status to another.
                        format: date-time
                        type: string
                      name:
                        description: Name referrs to a tenant cluster node name.
                        type: string
                      version:
                        description: Version referrs to the version used by the node
                          as mandated by the provider operator.
                        type: string
                    required:
                    - lastTransitionTime
                    - name
                    - version
                    type: object
                  type: array
                resources:
                  description: Resources is a list of arbitrary conditions of operatorkit
                    resource implementations.
                  items:
                    description: Resource is structure holding arbitrary conditions
                      of operatorkit resource implementations. Imagine an operator
                      implements an instance resource. This resource may operates
                      sequentially but has to operate based on a certain system state
                      it manages. So it tracks the status as needed here specific
                      to its own implementation and means in order to fulfil its premise.
                    properties:
                      conditions:
                        items:
                          description: StatusClusterResourceCondition expresses the
                            conditions in which an operatorkit resource may is.
                          properties:
                            lastTransitionTime:
                              description: LastTransitionTime is the last time the
                                condition transitioned from one status to another.
                              format: date-time
                              type: string
                            status:
                              description: Status may be True, False or Unknown.
                              type: string
                            type:
                              description: Type may be anything an operatorkit resource
                                may define.
                              type: string
                          required:
                          - lastTransitionTime
                          - status
                          - type
                          type: object
                        type: array
                      name:
                        type: string
                    required:
                    - conditions
                    - name
                    type: object
                  type: array
                scaling:
                  description: StatusClusterScaling expresses the current status of
                    desired number of worker nodes in guest cluster.
                  properties:
                    desiredCapacity:
                      type: integer
                  required:
                  - desiredCapacity
                  type: object
                versions:
                  description: Versions is a list that acts like a historical track
                    record of versions a guest cluster went through. A version is
                    only added to the list as soon as the guest cluster successfully
                    migrated to the version added here.
                  items:
                    description: StatusClusterVersion expresses the versions in which
                      a guest cluster was and may still be.
                    properties:
                      date:
                        description: "TODO date is deprecated due to LastTransitionTime
                          This can be removed ones the new properties are properly
                          used in all tenant clusters. \n     https://github.com/giantswarm/giantswarm/issues/3988"
                        format: date-time
                        type: string
                      lastTransitionTime:
                        description: LastTransitionTime is the last time the condition
                          transitioned from one status to another.
                        format: date-time
                        type: string
                      semver:
                        description: Semver is some semver version, e.g. 1.0.0.
                        type: string
                    required:
                    - date
                    - lastTransitionTime
                    - semver
                    type: object
                  type: array
              required:
              - conditions
              - network
              - nodes
              - resources
              - scaling
              - versions
              type: object
            kvm:
              properties:
                nodeIndexes:
                  additionalProperties:
                    type: integer
                  description: NodeIndexes is a map from nodeID -> nodeIndex. This
                    is used to create deterministic iSCSI initiator names.
                  type: object
              required:
              - nodeIndexes
              type: object
          required:
          - cluster
          - kvm
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

func NewKVMConfigCRD() *v1beta1.CustomResourceDefinition {
	var crd v1beta1.CustomResourceDefinition
	_ = yaml.Unmarshal([]byte(kvmconfigsYAML), &crd)
	return &crd
}
