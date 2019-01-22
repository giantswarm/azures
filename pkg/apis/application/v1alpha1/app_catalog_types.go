package v1alpha1

import (
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewAppCatalogCRD returns a new custom resource definition for AppCatalog.
// This might look something like the following.
//
//     apiVersion: apiextensions.k8s.io/v1beta1
//     kind: CustomResourceDefinition
//     metadata:
//       name: appcatalog.application.giantswarm.io
//     spec:
//       group: application.giantswarm.io
//       scope: Cluster
//       version: v1alpha1
//       names:
//         kind: AppCatalog
//         plural: appcatalogs
//         singular: appcatalog
//
func NewAppCatalogCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return &apiextensionsv1beta1.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: apiextensionsv1beta1.SchemeGroupVersion.String(),
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "appcatalogs.application.giantswarm.io",
		},
		Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
			Group:   "application.giantswarm.io",
			Scope:   "Cluster",
			Version: "v1alpha1",
			Names: apiextensionsv1beta1.CustomResourceDefinitionNames{
				Kind:     "AppCatalog",
				Plural:   "appcatalogs",
				Singular: "appcatalog",
			},
		},
	}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AppCatalog struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              AppCatalogSpec `json:"spec"`
}

type AppCatalogSpec struct {
	// Title is the name of the app catalog for this CR
	// e.g. Catalog of Apps by Giant Swarm
	Title       string `json:"title" yaml:"title"`
	Description string `json:"description" yaml:"description"`
	// CatalogStorage references a map containing values that should be
	// applied to the appcatalog.
	CatalogStorage AppCatalogSpecCatalogStorage `json:"catalogStorage" yaml:"catalogStorage"`
	// Config is the config to be applied when apps belonging to this
	// catalog are deployed.
	Config AppCatalogSpecConfig  `json:"config" yaml:"config"`
	// LogoURL contains the links for logo image file for this app catalog
	LogoURL string `json:"logoURL" yaml:"logoURL"`
}

type AppCatalogSpecCatalogStorage struct {
	// Type indicates which repository type would be used for this AppCatalog.
	// e.g. helm
	Type string `json:"type" yaml:"type"`
	// URL is the link to where this AppCatalog's repository is located
	// e.g. https://giantswarm.github.com/app-catalog/.
	URL string `json:"URL" yaml:"URL"`
}

type AppCatalogSpecConfig struct {
	// ConfigMap references a config map containing catalog values that
	// should be applied to apps in this catalog.
	ConfigMap AppCatalogSpecConfigConfigMap `json:"configMap" yaml:"configMap"`
	// Secret references a secret containing catalog values that should be
	// applied to apps in this catalog.
	Secret AppCatalogSpecConfigSecret `json:"secret" yaml:"secret"`
}

type AppCatalogSpecConfigConfigMap struct {
	// Name is the name of the config map containing catalog values to
	// apply, e.g. app-catalog-values.
	Name string `json:"name" yaml:"name"`
	// Namespace is the namespace of the catalog values config map,
	// e.g. giantswarm.
	Namespace string `json:"namespace" yaml:"namespace"`
}

type AppCatalogSpecConfigSecret struct {
	// Name is the name of the secret containing catalog values to apply,
	// e.g. app-catalog-secret.
	Name string `json:"name" yaml:"name"`
	// Namespace is the namespace of the secret,
	// e.g. giantswarm.
	Namespace string `json:"namespace" yaml:"namespace"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AppCatalogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []AppCatalog `json:"items"`
}
