// Generated by *go generate* - DO NOT EDIT
/*
Copyright 2022. projectsveltos.io. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package crd

var ClassifierFile = "../../config/crd/bases/lib.projectsveltos.io_classifiers.yaml"
var ClassifierCRD = []byte(`---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.8.0
  creationTimestamp: null
  name: classifiers.lib.projectsveltos.io
spec:
  group: lib.projectsveltos.io
  names:
    kind: Classifier
    listKind: ClassifierList
    plural: classifiers
    singular: classifier
  scope: Cluster
  versions:
  - name: v1alpha1
    schema:
      openAPIV3Schema:
        description: Classifier is the Schema for the classifiers API
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
            description: ClassifierSpec defines the desired state of Classifier
            properties:
              classifierLabels:
                description: ClassifierLabels is set of labels, key,value pair, that
                  will be added to each cluster matching Classifier instance
                items:
                  properties:
                    key:
                      description: Key is the label key
                      type: string
                    value:
                      description: Value is the label value
                      type: string
                  required:
                  - key
                  - value
                  type: object
                type: array
              deployedResources:
                description: DeployedResources allows to classify based on current
                  deployed resources
                items:
                  properties:
                    fieldFilters:
                      description: 'FieldFilters allows to filter resources based
                        on current field values. Internally uses FieldSelector so
                        only fields supported by FieldSelector can be used. Current
                        list: https://github.com/kubernetes/kubernetes/blob/9d577d8a29893062dfbd669997396dbd01ab0e47/pkg/apis/core/v1/conversion.go#L33'
                      items:
                        properties:
                          field:
                            description: Field is the field
                            type: string
                          operation:
                            description: Operation is the comparison operation
                            enum:
                            - Equal
                            - Different
                            type: string
                          value:
                            description: Value is the field value
                            type: string
                        required:
                        - field
                        - operation
                        - value
                        type: object
                      type: array
                    group:
                      description: Group of the resource deployed in the CAPI Cluster.
                      type: string
                    kind:
                      description: Kind of the resource deployed in the CAPI Cluster.
                      minLength: 1
                      type: string
                    labelFilters:
                      description: LabelFilters allows to filter resources based on
                        current labels.
                      items:
                        properties:
                          key:
                            description: Key is the label key
                            type: string
                          operation:
                            description: Operation is the comparison operation
                            enum:
                            - Equal
                            - Different
                            type: string
                          value:
                            description: Value is the label value
                            type: string
                        required:
                        - key
                        - operation
                        - value
                        type: object
                      type: array
                    maxCount:
                      description: MaxCount is the maximun number of resources to
                        match
                      type: integer
                    minCount:
                      description: MinCount is the minimum number of resources to
                        match
                      type: integer
                    namespace:
                      description: Namespace of the resource deployed in the CAPI
                        Cluster. Empty for resources scoped at cluster level.
                      type: string
                    version:
                      description: Version of the resource deployed in the CAPI Cluster.
                      type: string
                  required:
                  - group
                  - kind
                  - version
                  type: object
                type: array
              kubernetesVersionConstraints:
                description: KubernetesVersionConstraints allows to classify based
                  on current kubernetes version
                items:
                  properties:
                    comparison:
                      description: Comparison indicate how to compare cluster kubernetes
                        version with the specified version
                      enum:
                      - Equal
                      - NotEqual
                      - GreaterThan
                      - LessThan
                      - GreaterThanOrEqualTo
                      - LessThanOrEqualTo
                      type: string
                    version:
                      description: Version is the kubernetes version
                      type: string
                  required:
                  - comparison
                  - version
                  type: object
                type: array
            required:
            - classifierLabels
            type: object
          status:
            description: ClassifierStatus defines the observed state of Classifier
            properties:
              clusterInfo:
                description: ClusterInfo reference all the cluster-api Cluster where
                  Classifier has been/is being deployed
                items:
                  properties:
                    cluster:
                      description: Cluster references the CAPI Cluster
                      properties:
                        apiVersion:
                          description: API version of the referent.
                          type: string
                        fieldPath:
                          description: 'If referring to a piece of an object instead
                            of an entire object, this string should contain a valid
                            JSON/Go field access statement, such as desiredState.manifest.containers[2].
                            For example, if the object reference is to a container
                            within a pod, this would take on a value like: "spec.containers{name}"
                            (where "name" refers to the name of the container that
                            triggered the event) or if no container name is specified
                            "spec.containers[2]" (container with index 2 in this pod).
                            This syntax is chosen only to have some well-defined way
                            of referencing a part of an object. TODO: this design
                            is not final and this field is subject to change in the
                            future.'
                          type: string
                        kind:
                          description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
                          type: string
                        name:
                          description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                          type: string
                        namespace:
                          description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                          type: string
                        resourceVersion:
                          description: 'Specific resourceVersion to which this reference
                            is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency'
                          type: string
                        uid:
                          description: 'UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids'
                          type: string
                      type: object
                    failureMessage:
                      description: FailureMessage provides more information about
                        the error.
                      type: string
                    hash:
                      description: Hash represents the hash of the Classifier currently
                        deployed in the CAPI Cluster
                      format: byte
                      type: string
                    status:
                      description: Status represents the state of the feature in the
                        workload cluster
                      enum:
                      - Provisioning
                      - Provisioned
                      - Failed
                      - Removing
                      - Removed
                      type: string
                  required:
                  - cluster
                  - hash
                  - status
                  type: object
                type: array
              machingClusterStatuses:
                description: MatchingClusterRefs reference all the cluster-api Cluster
                  currently matching Classifier
                items:
                  properties:
                    clusterRef:
                      description: ClusterRef references the matching Cluster
                      properties:
                        apiVersion:
                          description: API version of the referent.
                          type: string
                        fieldPath:
                          description: 'If referring to a piece of an object instead
                            of an entire object, this string should contain a valid
                            JSON/Go field access statement, such as desiredState.manifest.containers[2].
                            For example, if the object reference is to a container
                            within a pod, this would take on a value like: "spec.containers{name}"
                            (where "name" refers to the name of the container that
                            triggered the event) or if no container name is specified
                            "spec.containers[2]" (container with index 2 in this pod).
                            This syntax is chosen only to have some well-defined way
                            of referencing a part of an object. TODO: this design
                            is not final and this field is subject to change in the
                            future.'
                          type: string
                        kind:
                          description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
                          type: string
                        name:
                          description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                          type: string
                        namespace:
                          description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                          type: string
                        resourceVersion:
                          description: 'Specific resourceVersion to which this reference
                            is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency'
                          type: string
                        uid:
                          description: 'UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids'
                          type: string
                      type: object
                    managedLabels:
                      description: ManagedLabels indicates the labels being managed
                        on the cluster by this Classifier instance
                      items:
                        type: string
                      type: array
                    unManagedLabels:
                      description: UnManagedLabel indicates the labels this Classifier
                        instance would like to manage but cannot because different
                        instance is already managing it
                      items:
                        properties:
                          failureMessage:
                            description: FailureMessage is a human consumable message
                              explaining the misconfiguration
                            type: string
                          key:
                            description: Key represents a label Classifier would like
                              to manage but cannot because currently managed by different
                              instance
                            type: string
                        required:
                        - key
                        type: object
                      type: array
                  required:
                  - clusterRef
                  type: object
                type: array
            type: object
        type: object
    served: true
    storage: true
    subresources:
      status: {}
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`)
