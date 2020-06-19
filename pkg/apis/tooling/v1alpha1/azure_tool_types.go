package v1alpha1

import (
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/pkg/apis/tooling"
	"github.com/giantswarm/apiextensions/pkg/crd"
	"github.com/giantswarm/apiextensions/pkg/key"
)

// NewAzureToolCRD returns a CRD defining an AzureTool.
func NewAzureToolCRD() *v1.CustomResourceDefinition {
	return crd.LoadV1(tooling.Group, tooling.KindAzureTool)
}

// NewAzureToolCR returns an AzureTool custom resource.
func NewAzureToolCR(name, namespace string) *AzureTool {
	cr := AzureTool{}
	cr.TypeMeta, cr.ObjectMeta = key.NewMeta(SchemeGroupVersion, tooling.KindAzureTool, name, namespace)
	return &cr
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=azure;giantswarm
// +kubebuilder:storageversion
// +k8s:openapi-gen=true

type AzureTool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              AzureToolSpec `json:"spec"`
	// +kubebuilder:validation:Optional
	Status AzureToolStatus `json:"status"`
}

// +k8s:openapi-gen=true
type AzureToolSpec struct {
	// Workspace refers to the Azure Log Analytics Workspace.
	Workspace AzureToolWorkspace `json:"workspace" yaml:"workspace"`
}

// +k8s:openapi-gen=true
type AzureToolStatus struct {
	WorkspaceStatus string `json:"workspace_status"`
}

// +k8s:openapi-gen=true
type AzureToolWorkspace struct {
	// ID is the Workspace ID.
	ID string `json:"id" yaml:"id"`
	// Mode is the mode that the Workspace is running in.
	Mode string `json:"mode" yaml:"mode"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzureToolList is the type returned when listing AzureToolList resources.
type AzureToolList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata" yaml:"metadata"`
	Items           []AzureTool `json:"items" yaml:"items"`
}
