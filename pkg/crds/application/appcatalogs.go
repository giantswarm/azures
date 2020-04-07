package application

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"sigs.k8s.io/yaml"
)

const appcatalogsYAML = `
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: (devel)
  creationTimestamp: null
  name: appcatalogs.application.giantswarm.io
spec:
  group: application.giantswarm.io
  names:
    kind: AppCatalog
    listKind: AppCatalogList
    plural: appcatalogs
    singular: appcatalog
  scope: Namespaced
  validation:
    openAPIV3Schema:
      description: "AppCatalog CRs might look something like the following. \n    apiVersion:
        application.giantswarm.io/v1alpha1    kind: AppCatalog    metadata:      name:
        \"giantswarm\"      labels:        app-operator.giantswarm.io/version: \"1.0.0\"
        \n    spec:      title: \"Giant Swarm\"      description: \"Catalog of Apps
        by Giant Swarm\"      config:        configMap:          name: \"app-catalog-values\"
        \         namespace: \"giantswarm\"        secret:          name: \"app-catalog-secrets\"
        \         namespace: \"giantswarm\"      logoURL: \"/images/repo_icons/incubator.png\"
        \     storage:        type: \"helm\"        URL: \"https://giantswarm.github.com/app-catalog/\""
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
            config:
              description: Config is the config to be applied when apps belonging
                to this catalog are deployed.
              properties:
                configMap:
                  description: ConfigMap references a config map containing catalog
                    values that should be applied to apps in this catalog.
                  properties:
                    name:
                      description: Name is the name of the config map containing catalog
                        values to apply, e.g. app-catalog-values.
                      type: string
                    namespace:
                      description: Namespace is the namespace of the catalog values
                        config map, e.g. giantswarm.
                      type: string
                  required:
                  - name
                  - namespace
                  type: object
                secret:
                  description: Secret references a secret containing catalog values
                    that should be applied to apps in this catalog.
                  properties:
                    name:
                      description: Name is the name of the secret containing catalog
                        values to apply, e.g. app-catalog-secret.
                      type: string
                    namespace:
                      description: Namespace is the namespace of the secret, e.g.
                        giantswarm.
                      type: string
                  required:
                  - name
                  - namespace
                  type: object
              required:
              - configMap
              - secret
              type: object
            description:
              type: string
            logoURL:
              description: LogoURL contains the links for logo image file for this
                app catalog
              type: string
            storage:
              description: Storage references a map containing values that should
                be applied to the appcatalog.
              properties:
                URL:
                  description: URL is the link to where this AppCatalog's repository
                    is located e.g. https://giantswarm.github.com/app-catalog/.
                  type: string
                type:
                  description: Type indicates which repository type would be used
                    for this AppCatalog. e.g. helm
                  type: string
              required:
              - URL
              - type
              type: object
            title:
              description: Title is the name of the app catalog for this CR e.g. Catalog
                of Apps by Giant Swarm
              type: string
          required:
          - config
          - description
          - logoURL
          - storage
          - title
          type: object
      required:
      - metadata
      - spec
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

func NewAppCatalogCRD() *v1beta1.CustomResourceDefinition {
	var crd v1beta1.CustomResourceDefinition
	_ = yaml.Unmarshal([]byte(appcatalogsYAML), &crd)
	return &crd
}
