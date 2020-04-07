package application

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"sigs.k8s.io/yaml"
)

const chartsYAML = `
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: (devel)
  creationTimestamp: null
  name: charts.application.giantswarm.io
spec:
  group: application.giantswarm.io
  names:
    kind: Chart
    listKind: ChartList
    plural: charts
    singular: chart
  scope: Namespaced
  validation:
    openAPIV3Schema:
      description: "Chart CRs might look something like the following. \n    apiVersion:
        application.giantswarm.io/v1alpha1    kind: Chart    metadata:      name:
        \"prometheus\"      labels:        chart-operator.giantswarm.io/version: \"1.0.0\"
        \n    spec:      name: \"prometheus\"      namespace: \"monitoring\"      config:
        \       configMap:        name: \"prometheus-values\"        namespace: \"monitoring\"
        \       resourceVersion: \"\"      secret:        name: \"prometheus-secrets\"
        \       namespace: \"monitoring\"        resourceVersion: \"\"      tarballURL:
        \"https://giantswarm.github.com/app-catalog/prometheus-1-0-0.tgz\"      version:
        \"1.0.0\" \n    status:      appVersion: \"2.4.3\" # Optional value from Chart.yaml
        with the version of the deployed app.      release:        lastDeployed: \"2018-11-30T21:06:20Z\"
        \       status: \"DEPLOYED\"      version: \"1.1.0\" # Required value from
        Chart.yaml with the version of the chart."
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
              description: Config is the config to be applied when the chart is deployed.
              properties:
                configMap:
                  description: ConfigMap references a config map containing values
                    that should be applied to the chart.
                  properties:
                    name:
                      description: Name is the name of the config map containing chart
                        values to apply, e.g. prometheus-chart-values.
                      type: string
                    namespace:
                      description: Namespace is the namespace of the values config
                        map, e.g. monitoring.
                      type: string
                    resourceVersion:
                      description: ResourceVersion is the Kubernetes resource version
                        of the configmap. Used to detect if the configmap has changed,
                        e.g. 12345.
                      type: string
                  required:
                  - name
                  - namespace
                  - resourceVersion
                  type: object
                secret:
                  description: Secret references a secret containing secret values
                    that should be applied to the chart.
                  properties:
                    name:
                      description: Name is the name of the secret containing chart
                        values to apply, e.g. prometheus-chart-secret.
                      type: string
                    namespace:
                      description: Namespace is the namespace of the secret, e.g.
                        kube-system.
                      type: string
                    resourceVersion:
                      description: ResourceVersion is the Kubernetes resource version
                        of the secret. Used to detect if the secret has changed, e.g.
                        12345.
                      type: string
                  required:
                  - name
                  - namespace
                  - resourceVersion
                  type: object
              required:
              - configMap
              - secret
              type: object
            name:
              description: Name is the name of the Helm chart to be deployed. e.g.
                kubernetes-prometheus
              type: string
            namespace:
              description: Namespace is the namespace where the chart should be deployed.
                e.g. monitoring
              type: string
            tarballURL:
              description: TarballURL is the URL for the Helm chart tarball to be
                deployed. e.g. https://path/to/prom-1-0-0.tgz"
              type: string
            version:
              description: Version is the version of the chart that should be deployed.
                e.g. 1.0.0
              type: string
          required:
          - config
          - name
          - namespace
          - tarballURL
          - version
          type: object
        status:
          properties:
            appVersion:
              description: AppVersion is the value of the AppVersion field in the
                Chart.yaml of the deployed chart. This is an optional field with the
                version of the component being deployed. e.g. 0.21.0. https://docs.helm.sh/developing_charts/#the-chart-yaml-file
              type: string
            reason:
              description: Reason is the description of the last status of helm release
                when the chart is not installed successfully, e.g. deploy resource
                already exists.
              type: string
            release:
              description: Release is the status of the Helm release for the deployed
                chart.
              properties:
                lastDeployed:
                  description: LastDeployed is the time when the deployed chart was
                    last deployed.
                  format: date-time
                  type: string
                status:
                  description: Status is the status of the deployed chart, e.g. DEPLOYED.
                  type: string
              required:
              - lastDeployed
              - status
              type: object
            version:
              description: Version is the value of the Version field in the Chart.yaml
                of the deployed chart. e.g. 1.0.0.
              type: string
          required:
          - appVersion
          - release
          - version
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

func NewChartCRD() *v1beta1.CustomResourceDefinition {
	var crd v1beta1.CustomResourceDefinition
	_ = yaml.Unmarshal([]byte(chartsYAML), &crd)
	return &crd
}
