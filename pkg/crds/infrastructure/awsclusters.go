package infrastructure

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"sigs.k8s.io/yaml"
)

const awsclustersYAML = `
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: (devel)
  creationTimestamp: null
  name: awsclusters.infrastructure.giantswarm.io
spec:
  group: infrastructure.giantswarm.io
  names:
    kind: AWSCluster
    listKind: AWSClusterList
    plural: awsclusters
    singular: awscluster
  scope: Namespaced
  validation:
    openAPIV3Schema:
      description: "AWSCluster is the infrastructure provider referenced in upstream
        CAPI Cluster CRs. \n     apiVersion: infrastructure.giantswarm.io/v1alpha2
        \    kind: AWSCluster     metadata:       labels:         aws-operator.giantswarm.io/version:
        6.2.0         cluster-operator.giantswarm.io/version: 0.17.0         giantswarm.io/cluster:
        \"8y5kc\"         giantswarm.io/organization: \"giantswarm\"         release.giantswarm.io/version:
        7.3.1       name: 8y5kc     spec:       cluster:         description: my fancy
        cluster         dns:           domain: gauss.eu-central-1.aws.gigantic.io
        \        oidc:           claims:             username: email             groups:
        groups           clientID: foobar-dex-client           issuerURL: https://dex.gatekeeper.eu-central-1.aws.example.com
        \      provider:         credentialSecret:           name: credential-default
        \          namespace: giantswarm         master:           availabilityZone:
        eu-central-1a           instanceType: m4.large         region: eu-central-1
        \    status:       cluster:         conditions:         - lastTransitionTime:
        \"2019-03-25T17:10:09.333633991Z\"           type: Created         id: 8y5kc
        \        versions:         - lastTransitionTime: \"2019-03-25T17:10:09.995948706Z\"
        \          version: 4.9.0       provider:         network:           cidr:
        10.1.6.0/24           vpcID: vpc-1234567890abcdef0"
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
          description: metav1.ObjectMeta is standard Kubernetes resource metadata.
          type: object
        spec:
          description: AWSClusterSpec is the spec part for the AWSCluster resource.
          properties:
            cluster:
              description: Cluster provides cluster specification details.
              properties:
                description:
                  description: Description is a user-friendly description that should
                    explain the purpose of the cluster to humans.
                  type: string
                dns:
                  description: DNS holds DNS configuration details.
                  properties:
                    domain:
                      type: string
                  required:
                  - domain
                  type: object
                oidc:
                  description: OIDC holds configuration for OpenID Connect (OIDC)
                    authentication.
                  properties:
                    claims:
                      description: AWSClusterSpecClusterOIDCClaims defines OIDC claims.
                      properties:
                        groups:
                          type: string
                        username:
                          type: string
                      type: object
                    clientID:
                      type: string
                    issuerURL:
                      type: string
                  type: object
              required:
              - description
              - dns
              type: object
            provider:
              description: Provider holds provider-specific configuration details.
              properties:
                credentialSecret:
                  description: CredentialSecret specifies the location of the secret
                    providing the ARN of AWS IAM identity to use with this cluster.
                  properties:
                    name:
                      description: Name is the name of the provider credential resoure.
                      type: string
                    namespace:
                      description: Namespace is the kubernetes namespace that holds
                        the provider credential.
                      type: string
                  required:
                  - name
                  - namespace
                  type: object
                master:
                  description: Master holds master node configuration details.
                  properties:
                    availabilityZone:
                      description: AvailabilityZone is the AWS availability zone to
                        place the master node in.
                      type: string
                    instanceType:
                      description: InstanceType specifies the AWS EC2 instance type
                        to use for the master node.
                      type: string
                  required:
                  - availabilityZone
                  - instanceType
                  type: object
                region:
                  description: Region is the AWS region the cluster is to be running
                    in.
                  type: string
              required:
              - credentialSecret
              - master
              - region
              type: object
          required:
          - cluster
          - provider
          type: object
        status:
          description: AWSClusterStatus holds status information about the cluster,
            populated once the cluster is in creation or created.
          properties:
            cluster:
              description: Cluster provides cluster-specific status details, including
                conditions and versions.
              properties:
                conditions:
                  description: One or several conditions that are currently applicable
                    to the cluster.
                  items:
                    description: CommonClusterStatusCondition explains the current
                      condition(s) of the cluster.
                    properties:
                      condition:
                        description: Condition string, e. g. "Creating", "Created",
                          "Upgraded"
                        type: string
                      lastTransitionTime:
                        description: Time the condition occurred.
                        format: date-time
                        type: string
                    required:
                    - condition
                    - lastTransitionTime
                    type: object
                  type: array
                id:
                  description: Identifier of the cluster.
                  type: string
                versions:
                  description: Release versions the cluster used so far.
                  items:
                    description: CommonClusterStatusVersion informs which aws-operator
                      version was/responsible for this cluster.
                    properties:
                      lastTransitionTime:
                        description: Time the cluster assumed the given version.
                        format: date-time
                        type: string
                      version:
                        description: The aws-operator version responsible for handling
                          the cluster.
                        type: string
                    required:
                    - lastTransitionTime
                    - version
                    type: object
                  type: array
              required:
              - id
              type: object
            provider:
              description: Provider provides provider-specific status details.
              properties:
                network:
                  description: Network provides network-specific configuration details
                  properties:
                    cidr:
                      description: IPv4 address block used by the tenant cluster,
                        in CIDR notation.
                      type: string
                    vpcID:
                      description: VPCID contains the ID of the tenant cluster, e.g.
                        vpc-1234567890abcdef0.
                      type: string
                  required:
                  - cidr
                  - vpcID
                  type: object
              required:
              - network
              type: object
          type: object
      required:
      - spec
      type: object
  version: v1alpha2
  versions:
  - name: v1alpha2
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`

func NewAWSClusterCRD() *v1beta1.CustomResourceDefinition {
	var crd v1beta1.CustomResourceDefinition
	_ = yaml.Unmarshal([]byte(awsclustersYAML), &crd)
	return &crd
}
