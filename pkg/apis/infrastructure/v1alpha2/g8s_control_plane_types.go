package v1alpha2

import (
	"github.com/ghodss/yaml"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/kubernetes/pkg/apis/core"
)

const (
	kindG8sControlPlane = "G8sControlPlane"
)

const g8sControlPlaneCRDYAML = `
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: g8sControlPlanes.core.giantswarm.io
spec:
  group: infrastructure.giantswarm.io
  scope: Namespaced
  version: v1alpha1
  names:
    kind: G8sControlPlane
    plural: g8sControlPlanes
    singular: g8sControlPlane
  subresources:
    status: {}
`

var g8sControlPlaneCRD *apiextensionsv1beta1.CustomResourceDefinition

func init() {
	err := yaml.Unmarshal([]byte(g8sControlPlaneCRDYAML), &g8sControlPlaneCRD)
	if err != nil {
		panic(err)
	}
}

func NewG8sControlPlaneCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return g8sControlPlaneCRD.DeepCopy()
}

func NewG8sControlPlaneTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: SchemeGroupVersion.String(),
		Kind:       kindG8sControlPlane,
	}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// G8sControlPlane defines the ControlPlane (Master nodes) of a
// Giant Swarm Tenant Cluster
//
//	apiVersion: infrastructure.giantswarm.io/v1alpha2
//	kind: G8sControlPlane
//	metadata:
//    labels:
//      label1: ????
//	  name: ????
//	spec:
//    replicas: 3
//    //TODO LH infrastructureRef ?
//  status:
//    replicas: 3
//    readyReplicas: 3
//
type G8sControlPlane struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              G8sControlPlaneSpec   `json:"spec"`
	Status            G8sControlPlaneStatus `json:"status"`
}

type G8sControlPlaneSpec struct {
	// Replicas is the number replicas of the master node.
	Replicas int `json:"replicas" yaml:"replicas"`
	// InfrastructureRef is a required reference to provider-specific
	// Infrastructure.
	InfrastructureRef corev1.ObjectReference `json:"infrastructureRef"`
}

// G8sControlPlaneStatus defines the observed state of G8sControlPlane.
type G8sControlPlaneStatus struct {
	// Total number of non-terminated machines targeted by this control plane
	// (their labels match the selector).
	// +optional
	Replicas int32 `json:"replicas,omitempty"`
	// Total number of fully running and ready control plane machines.
	// +optional
	ReadyReplicas int32 `json:"readyReplicas,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type G8sControlPlaneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []G8sControlPlane `json:"items"`
}
