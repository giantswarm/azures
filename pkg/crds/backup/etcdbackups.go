package backup

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"sigs.k8s.io/yaml"
)

const etcdbackupsYAML = `
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: (devel)
  creationTimestamp: null
  name: etcdbackups.backup.giantswarm.io
spec:
  group: backup.giantswarm.io
  names:
    kind: ETCDBackup
    listKind: ETCDBackupList
    plural: etcdbackups
    singular: etcdbackup
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
            guestBackup:
              description: GuestBackup is a boolean indicating if the tenant clusters
                have to be backupped
              type: boolean
          required:
          - guestBackup
          type: object
        status:
          properties:
            finishedTimestamp:
              description: Timestamp when the last (final) attempt was made (when
                the Phase became either 'Completed' or 'Failed'
              format: date-time
              type: string
            instances:
              additionalProperties:
                properties:
                  name:
                    description: Name of the tenant cluster or 'Control Plane'
                    type: string
                  v2:
                    description: Status of the V2 backup for this instance
                    properties:
                      backupFileSize:
                        description: Size of the backup file
                        format: int64
                        type: integer
                      creationTime:
                        description: Time took by the backup creation process
                        format: int64
                        type: integer
                      encryptionTime:
                        description: Time took by the backup encryption process
                        format: int64
                        type: integer
                      finishedTimestamp:
                        description: Timestamp when the last (final) attempt was made
                          (when the Phase became either 'Completed' or 'Failed'
                        format: date-time
                        type: string
                      latestError:
                        description: Latest backup error message
                        type: string
                      startedTimestamp:
                        description: Timestamp when the first attempt was made
                        format: date-time
                        type: string
                      status:
                        description: Status of this isntance's backup job (can be
                          'Pending', 'Running'. 'Completed', 'Failed')
                        type: string
                      uploadTime:
                        description: Time took by the backup upload process
                        format: int64
                        type: integer
                    required:
                    - status
                    type: object
                  v3:
                    description: Status of the V3 backup for this instance
                    properties:
                      backupFileSize:
                        description: Size of the backup file
                        format: int64
                        type: integer
                      creationTime:
                        description: Time took by the backup creation process
                        format: int64
                        type: integer
                      encryptionTime:
                        description: Time took by the backup encryption process
                        format: int64
                        type: integer
                      finishedTimestamp:
                        description: Timestamp when the last (final) attempt was made
                          (when the Phase became either 'Completed' or 'Failed'
                        format: date-time
                        type: string
                      latestError:
                        description: Latest backup error message
                        type: string
                      startedTimestamp:
                        description: Timestamp when the first attempt was made
                        format: date-time
                        type: string
                      status:
                        description: Status of this isntance's backup job (can be
                          'Pending', 'Running'. 'Completed', 'Failed')
                        type: string
                      uploadTime:
                        description: Time took by the backup upload process
                        format: int64
                        type: integer
                    required:
                    - status
                    type: object
                required:
                - name
                - v2
                - v3
                type: object
              description: map containing the state of the backup for all instances
              type: object
            startedTimestamp:
              description: Timestamp when the first attempt was made
              format: date-time
              type: string
            status:
              description: Status of the whole backup job (can be 'Pending', 'Running'.
                'Completed', 'Failed')
              type: string
          required:
          - status
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

func NewETCDBackupCRD() *v1beta1.CustomResourceDefinition {
	var crd v1beta1.CustomResourceDefinition
	_ = yaml.Unmarshal([]byte(etcdbackupsYAML), &crd)
	return &crd
}
