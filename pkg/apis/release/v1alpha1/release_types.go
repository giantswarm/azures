package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/pkg/apis/release"
	"github.com/giantswarm/apiextensions/pkg/key"
)

type ReleaseState string

var (
	stateActive     ReleaseState = "active"     // nolint
	stateDeprecated ReleaseState = "deprecated" // nolint
	stateWIP        ReleaseState = "wip"        // nolint
)

func (r ReleaseState) String() string {
	return string(r)
}

func NewReleaseCR(name string, spec ReleaseSpec) *Release {
	cr := Release{
		Spec: spec,
	}
	groupVersionKind := metav1.GroupVersionKind{
		Group:   release.Group,
		Version: version,
		Kind:    release.KindRelease,
	}
	meta := key.NewCustomResourceMeta(groupVersionKind, name, "")
	cr.ObjectMeta = meta.ObjectMeta
	cr.TypeMeta = meta.TypeMeta
	return &cr
}

// +kubebuilder:printcolumn:name="Kubernetes version",type=string,JSONPath=`.spec.components[?(@.name=="kubernetes")].version`,description="Version of the kubernetes component in this release"
// +kubebuilder:printcolumn:name="State",type=string,JSONPath=`.spec.state`,description="State of the release"
// +kubebuilder:printcolumn:name="Age",type=string,JSONPath=`.spec.date`,description="Time since release creation"
// +kubebuilder:resource:scope=Cluster,categories=common;giantswarm
// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:storageversion

// Release is a Kubernetes resource (CR) representing a Giant Swarm tenant cluster release.
type Release struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              ReleaseSpec `json:"spec"`
}

type ReleaseSpec struct {
	// Apps describes apps used in this release.
	Apps []ReleaseSpecApp `json:"apps"`
	// +kubebuilder:validation:MinItems=1
	// Components describes components used in this release.
	Components []ReleaseSpecComponent `json:"components"`
	// Date that the release became active.
	Date *metav1.Time `json:"date"`
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern=`^(active|deprecated|wip)$`
	// State indicates the availability of the release: deprecated, active, or wip.
	State ReleaseState `json:"state"`
}

type ReleaseSpecComponent struct {
	// Name of the component.
	Name string `json:"name"`
	// +kubebuilder:validation:Pattern=`^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`
	// Version of the component.
	Version string `json:"version"`
}

type ReleaseSpecApp struct {
	// Version of the upstream component used in the app.
	ComponentVersion string `json:"componentVersion,omitempty"`
	// Name of the app.
	Name string `json:"name"`
	// +kubebuilder:validation:Pattern=`^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`
	// Version of the app.
	Version string `json:"version"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ReleaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Release `json:"items"`
}
