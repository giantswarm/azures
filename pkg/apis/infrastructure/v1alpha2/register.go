package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	group   = "infrastructure.giantswarm.io"
	version = "v1alpha2"
)

// knownTypes is the full list of objects to register with the scheme. It
// should contain pointers of zero values of all custom objects and custom
// object lists in the group version.
var knownTypes = []runtime.Object{
	&AWSCluster{},
	&AWSClusterList{},
	&AWSMachineDeployment{},
	&AWSMachineDeploymentList{},
	&G8sControlPlane{},
	&G8sControlPlaneList{},
}

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = schema.GroupVersion{
	Group:   group,
	Version: version,
}

var (
	schemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)

	// AddToScheme is used by the generated client.
	AddToScheme = schemeBuilder.AddToScheme
)

// Adds the list of known types to api.Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion, knownTypes...)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
