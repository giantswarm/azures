package v1alpha2

import (
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

const (
	kindAWSMachineDeployment = "AWSMachineDeployment"
)

const awsMachineDeploymentCRDYAML = `
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: awsmachinedeployments.infrastructure.giantswarm.io
spec:
  group: infrastructure.giantswarm.io
  names:
    kind: AWSMachineDeployment
    plural: awsmachinedeployments
    singular: awsmachinedeployment
  scope: Namespaced
  subresources:
    status: {}
  versions:
  - name: v1alpha1
    served: false
    storage: false
    schema:
      openAPIV3Schema:
        type: object
        properties: {}
  - name: v1alpha2
    served: true
    storage: true
    schema:
      openAPIV3Schema:
        type: object
        properties:
          spec:
            type: object
            properties:
              nodePool:
                description: |
                  General node pool configuration.
                type: object
                properties:
                  description:
                    description: |
                      User-friendly description of the node pool, e. g. to inform about the purpose.
                    type: string
                    maxLength: 100
                  machine:
                    type: object
                    properties:
                      dockerVolumeSizeGB:
                        description: |
                          Size of the volume reserved for Docker images and overlay file systems
                          of Docker containers. Unit: 1 GB = 1,000,000,000 Bytes.
                        format: int32
                        type: integer
                      kubeletVolumeSizeGB:
                        description: |
                          Size of the volume reserved for the kubelet, which can be used by Pods
                          via volumes of type EmptyDir. Unit: 1 GB = 1,000,000,000 Bytes.
                        format: int32
                        type: integer
                  scaling:
                    description: |
                      Configures the size of the node pool, in terms of the number of worker
                      nodes it can have, as a range. The actual number of worker nodes will
                      be determined by cluster-autoscaler within the configured range.
                      Setting the minimum and maximum to the same number effectively disables
                      autoscaling.
                    type: object
                    properties:
                      max:
                        description: |
                          Maximum number of worker nodes in this node pool.
                        format: int32
                        type: integer
                      min:
                        description: |
                          Minimum number of worker nodes in this node pool.
                        format: int32
                        type: integer
              provider:
                description: |
                  Configuration specific to AWS.
                type: object
                properties:
                  availabilityZones:
                    description: |
                      Name(s) of the availability zone(s) to use for worker nodes. Using multiple
                      availability zones results in higher resilience but can also result in
                      higher cost due to network traffic between availability zones.
                    items:
                      type: string
                    type: array
                  instanceDistribution:
                    description: |
                      Attributes defining the instance distribution in a node pool.
                    properties:
                      onDemandBaseCapacity:
                        default: 0
                        description: |
                          Base capacity of on-demand instances to use for worker nodes.
                        format: int32
                        minimum: 0
                        type: integer
                      onDemandPercentageAboveBaseCapacity:
                        default: 100
                        description: |
                          Percentage of on-demand instances to use for worker nodes, for instances exceeding onDemandBaseCapacity.
                        format: int32
                        maximum: 100
                        minimum: 0
                        type: integer
                    type: object
                  worker:
                    type: object
                    description: |
                      Specification of worker nodes.
                    properties:
                      instanceType:
                        description: |
                          AWS EC2 instance type name to use for the worker nodes in this node pool.
                        type: string
                      useAlikeInstanceTypes:
                        description: |
                          If true, certain instance types with specs similar to instanceType will be used.
                        type: boolean
          status:
            type: object
            properties:
              provider:
                description: |
                  Status specific to AWS.
                type: object
                properties:
                  worker:
                    type: object
                    description: |
                      Status of worker nodes.
                    properties:
                      instanceTypes:
                        description: |
                          AWS EC2 instance types used for the worker nodes in this node pool.
                        items:
                          type: string
                        type: array
                      spotInstances:
                        description: |
                          Number of spot instances used in this node pool.
                        type: integer
  conversion:
    strategy: None
`

var awsMachineDeploymentCRD *apiextensionsv1beta1.CustomResourceDefinition

func init() {
	err := yaml.Unmarshal([]byte(awsMachineDeploymentCRDYAML), &awsMachineDeploymentCRD)
	if err != nil {
		panic(err)
	}
}

func NewAWSMachineDeploymentCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return awsMachineDeploymentCRD.DeepCopy()
}

func NewAWSMachineDeploymentTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: SchemeGroupVersion.String(),
		Kind:       kindAWSMachineDeployment,
	}
}

// NewAWSMachineDeploymentCR returns an AWSMachineDeployment Custom Resource.
func NewAWSMachineDeploymentCR() *AWSMachineDeployment {
	return &AWSMachineDeployment{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				crDocsAnnotation: awsClusterDocumentationLink,
			},
		},
		TypeMeta: NewAWSMachineDeploymentTypeMeta(),
	}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AWSMachineDeployment is the infrastructure provider referenced in upstream
// CAPI MachineDeployment CRs.
type AWSMachineDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AWSMachineDeploymentSpec   `json:"spec" yaml:"spec"`
	Status            AWSMachineDeploymentStatus `json:"status" yaml:"status"`
}

type AWSMachineDeploymentSpec struct {
	NodePool AWSMachineDeploymentSpecNodePool `json:"nodePool" yaml:"nodePool"`
	Provider AWSMachineDeploymentSpecProvider `json:"provider" yaml:"provider"`
}

type AWSMachineDeploymentSpecNodePool struct {
	Description string                                  `json:"description" yaml:"description"`
	Machine     AWSMachineDeploymentSpecNodePoolMachine `json:"machine" yaml:"machine"`
	Scaling     AWSMachineDeploymentSpecNodePoolScaling `json:"scaling" yaml:"scaling"`
}

type AWSMachineDeploymentSpecNodePoolMachine struct {
	DockerVolumeSizeGB  int `json:"dockerVolumeSizeGB" yaml:"dockerVolumeSizeGB"`
	KubeletVolumeSizeGB int `json:"kubeletVolumeSizeGB" yaml:"kubeletVolumeSizeGB"`
}

type AWSMachineDeploymentSpecNodePoolScaling struct {
	Max int `json:"max" yaml:"max"`
	Min int `json:"min" yaml:"min"`
}

type AWSMachineDeploymentSpecProvider struct {
	AvailabilityZones    []string                                     `json:"availabilityZones" yaml:"availabilityZones"`
	InstanceDistribution AWSMachineDeploymentSpecInstanceDistribution `json:"instanceDistribution" yaml:"instanceDistribution"`
	Worker               AWSMachineDeploymentSpecProviderWorker       `json:"worker" yaml:"worker"`
}

type AWSMachineDeploymentSpecInstanceDistribution struct {
	// +kubebuilder:default=0
	// +kubebuilder:validation:Minimum=0
	OnDemandBaseCapacity int `json:"onDemandBaseCapacity" yaml:"onDemandBaseCapacity"`
	// +kubebuilder:default=100
	// +kubebuilder:validation:Maximum=100
	// +kubebuilder:validation:Minimum=0
	OnDemandPercentageAboveBaseCapacity int `json:"onDemandPercentageAboveBaseCapacity" yaml:"onDemandPercentageAboveBaseCapacity"`
}

type AWSMachineDeploymentSpecProviderWorker struct {
	InstanceType          string `json:"instanceType" yaml:"instanceType"`
	UseAlikeInstanceTypes bool   `json:"useAlikeInstanceTypes" yaml:"useAlikeInstanceTypes"`
}

type AWSMachineDeploymentStatus struct {
	Provider AWSMachineDeploymentStatusProvider `json:"provider" yaml:"provider"`
}

type AWSMachineDeploymentStatusProvider struct {
	Worker AWSMachineDeploymentStatusProviderWorker `json:"worker" yaml:"worker"`
}

type AWSMachineDeploymentStatusProviderWorker struct {
	InstanceTypes []string `json:"instanceTypes" yaml:"instanceTypes"`
	SpotInstances int      `json:"spotInstances" yaml:"spotInstances"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AWSMachineDeploymentList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata" yaml:"metadata"`
	Items           []AWSMachineDeployment `json:"items" yaml:"items"`
}
