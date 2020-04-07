package provider

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"sigs.k8s.io/yaml"
)

const awsconfigsYAML = `
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: (devel)
  creationTimestamp: null
  name: awsconfigs.provider.giantswarm.io
spec:
  group: provider.giantswarm.io
  names:
    kind: AWSConfig
    listKind: AWSConfigList
    plural: awsconfigs
    singular: awsconfig
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
            aws:
              properties:
                api:
                  description: AWSConfigSpecAWSAPI deprecated since aws-operator v12
                    resources.
                  properties:
                    elb:
                      description: AWSConfigSpecAWSAPIELB deprecated since aws-operator
                        v12 resources.
                      properties:
                        idleTimeoutSeconds:
                          type: integer
                      required:
                      - idleTimeoutSeconds
                      type: object
                    hostedZones:
                      type: string
                  required:
                  - elb
                  - hostedZones
                  type: object
                availabilityZones:
                  description: "AvailabilityZones is the number of AWS availability
                    zones used to spread the tenant cluster's worker nodes across.
                    There are limitations on availability zone settings due to binary
                    IP range splitting and provider specific region capabilities.
                    When for instance choosing 3 availability zones, the configured
                    IP range will be split into 4 ranges and thus one of it will not
                    be able to be utilized. Such limitations have to be considered
                    when designing the network topology and configuring tenant cluster
                    HA via AvailabilityZones. \n The selection and usage of the actual
                    availability zones for the created tenant cluster is randomized.
                    In case there are 4 availability zones provided in the used region
                    and the user selects 2 availability zones, the actually used availability
                    zones in which tenant cluster workload is put into will tend to
                    be different across tenant cluster creations. This is done in
                    order to provide more HA during single availability zone failures.
                    In case a specific availability zone fails, not all tenant clusters
                    will be affected due to the described selection process."
                  type: integer
                az:
                  description: "TODO remove the deprecated AZ field due to AvailabilityZones.
                    \n     https://github.com/giantswarm/giantswarm/issues/4507"
                  type: string
                credentialSecret:
                  properties:
                    name:
                      type: string
                    namespace:
                      type: string
                  required:
                  - name
                  - namespace
                  type: object
                etcd:
                  description: AWSConfigSpecAWSEtcd deprecated since aws-operator
                    v12 resources.
                  properties:
                    elb:
                      description: AWSConfigSpecAWSEtcdELB deprecated since aws-operator
                        v12 resources.
                      properties:
                        idleTimeoutSeconds:
                          type: integer
                      required:
                      - idleTimeoutSeconds
                      type: object
                    hostedZones:
                      type: string
                  required:
                  - elb
                  - hostedZones
                  type: object
                hostedZones:
                  description: "HostedZones is AWS hosted zones names in the host
                    cluster account. For each zone there will be \"CLUSTER_ID.k8s\"
                    NS record created in the host cluster account. Then for each created
                    NS record there will be a zone created in the guest account. After
                    that component specific records under those zones: \t- api.CLUSTER_ID.k8s.{{
                    .Spec.AWS.HostedZones.API.Name }} \t- etcd.CLUSTER_ID.k8s.{{ .Spec.AWS.HostedZones.Etcd.Name
                    }} \t- ingress.CLUSTER_ID.k8s.{{ .Spec.AWS.HostedZones.Ingress.Name
                    }} \t- *.CLUSTER_ID.k8s.{{ .Spec.AWS.HostedZones.Ingress.Name
                    }}"
                  properties:
                    api:
                      properties:
                        name:
                          type: string
                      required:
                      - name
                      type: object
                    etcd:
                      properties:
                        name:
                          type: string
                      required:
                      - name
                      type: object
                    ingress:
                      properties:
                        name:
                          type: string
                      required:
                      - name
                      type: object
                  required:
                  - api
                  - etcd
                  - ingress
                  type: object
                ingress:
                  description: AWSConfigSpecAWSIngress deprecated since aws-operator
                    v12 resources.
                  properties:
                    elb:
                      description: AWSConfigSpecAWSIngressELB deprecated since aws-operator
                        v12 resources.
                      properties:
                        idleTimeoutSeconds:
                          type: integer
                      required:
                      - idleTimeoutSeconds
                      type: object
                    hostedZones:
                      type: string
                  required:
                  - elb
                  - hostedZones
                  type: object
                masters:
                  items:
                    properties:
                      dockerVolumeSizeGB:
                        type: integer
                      imageID:
                        type: string
                      instanceType:
                        type: string
                    required:
                    - dockerVolumeSizeGB
                    - imageID
                    - instanceType
                    type: object
                  type: array
                region:
                  type: string
                vpc:
                  properties:
                    cidr:
                      type: string
                    peerId:
                      type: string
                    privateSubnetCidr:
                      type: string
                    publicSubnetCidr:
                      type: string
                    routeTableNames:
                      items:
                        type: string
                      type: array
                  required:
                  - cidr
                  - peerId
                  - privateSubnetCidr
                  - publicSubnetCidr
                  - routeTableNames
                  type: object
                workers:
                  items:
                    properties:
                      dockerVolumeSizeGB:
                        type: integer
                      imageID:
                        type: string
                      instanceType:
                        type: string
                    required:
                    - dockerVolumeSizeGB
                    - imageID
                    - instanceType
                    type: object
                  type: array
              required:
              - api
              - availabilityZones
              - az
              - credentialSecret
              - etcd
              - hostedZones
              - ingress
              - masters
              - region
              - vpc
              - workers
              type: object
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
            versionBundle:
              properties:
                version:
                  type: string
              required:
              - version
              type: object
          required:
          - aws
          - cluster
          - versionBundle
          type: object
        status:
          properties:
            aws:
              properties:
                autoScalingGroup:
                  properties:
                    name:
                      type: string
                  required:
                  - name
                  type: object
                availabilityZones:
                  items:
                    properties:
                      name:
                        type: string
                      subnet:
                        properties:
                          private:
                            properties:
                              cidr:
                                type: string
                            required:
                            - cidr
                            type: object
                          public:
                            properties:
                              cidr:
                                type: string
                            required:
                            - cidr
                            type: object
                        required:
                        - private
                        - public
                        type: object
                    required:
                    - name
                    - subnet
                    type: object
                  type: array
              required:
              - autoScalingGroup
              - availabilityZones
              type: object
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
          required:
          - aws
          - cluster
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

func NewAWSConfigCRD() *v1beta1.CustomResourceDefinition {
	var crd v1beta1.CustomResourceDefinition
	_ = yaml.Unmarshal([]byte(awsconfigsYAML), &crd)
	return &crd
}
