// Generated by *go generate* - DO NOT EDIT
/*
Copyright 2022-23. projectsveltos.io. All rights reserved.

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

var AddonComplianceFile = "../../config/crd/bases/lib.projectsveltos.io_addoncompliances.yaml"
var AddonComplianceCRD = []byte(`---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.12.0
  name: addoncompliances.lib.projectsveltos.io
spec:
  group: lib.projectsveltos.io
  names:
    kind: AddonCompliance
    listKind: AddonComplianceList
    plural: addoncompliances
    singular: addoncompliance
  scope: Cluster
  versions:
  - name: v1alpha1
    schema:
      openAPIV3Schema:
        description: AddonCompliance is the Schema for the AddonCompliance API
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
            description: AddonComplianceSpec defines the desired state of AddonCompliance
            properties:
              clusterRefs:
                description: ClusterRefs identifies clusters to associate to.
                items:
                  description: "ObjectReference contains enough information to let
                    you inspect or modify the referred object. --- New uses of this
                    type are discouraged because of difficulty describing its usage
                    when embedded in APIs. 1. Ignored fields.  It includes many fields
                    which are not generally honored.  For instance, ResourceVersion
                    and FieldPath are both very rarely valid in actual usage. 2. Invalid
                    usage help.  It is impossible to add specific help for individual
                    usage.  In most embedded usages, there are particular restrictions
                    like, \"must refer only to types A and B\" or \"UID not honored\"
                    or \"name must be restricted\". Those cannot be well described
                    when embedded. 3. Inconsistent validation.  Because the usages
                    are different, the validation rules are different by usage, which
                    makes it hard for users to predict what will happen. 4. The fields
                    are both imprecise and overly precise.  Kind is not a precise
                    mapping to a URL. This can produce ambiguity during interpretation
                    and require a REST mapping.  In most cases, the dependency is
                    on the group,resource tuple and the version of the actual struct
                    is irrelevant. 5. We cannot easily change it.  Because this type
                    is embedded in many locations, updates to this type will affect
                    numerous schemas.  Don't make new APIs embed an underspecified
                    API type they do not control. \n Instead of using this type, create
                    a locally provided and used type that is well-focused on your
                    reference. For example, ServiceReferences for admission registration:
                    https://github.com/kubernetes/api/blob/release-1.17/admissionregistration/v1/types.go#L533
                    ."
                  properties:
                    apiVersion:
                      description: API version of the referent.
                      type: string
                    fieldPath:
                      description: 'If referring to a piece of an object instead of
                        an entire object, this string should contain a valid JSON/Go
                        field access statement, such as desiredState.manifest.containers[2].
                        For example, if the object reference is to a container within
                        a pod, this would take on a value like: "spec.containers{name}"
                        (where "name" refers to the name of the container that triggered
                        the event) or if no container name is specified "spec.containers[2]"
                        (container with index 2 in this pod). This syntax is chosen
                        only to have some well-defined way of referencing a part of
                        an object. TODO: this design is not final and this field is
                        subject to change in the future.'
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
                  x-kubernetes-map-type: atomic
                type: array
              clusterSelector:
                description: ClusterSelector identifies clusters to associate to.
                type: string
              luaValidationRefs:
                description: LuaValidationRefs is a list of validations defined in
                  Lua language. In the matching clusters, add-ons will be deployed
                  only if all validations pass.
                items:
                  properties:
                    kind:
                      description: 'Kind of the resource. Supported kinds are: - flux
                        GitRepository;OCIRepository;Bucket - ConfigMap/Secret'
                      enum:
                      - GitRepository
                      - OCIRepository
                      - Bucket
                      - ConfigMap
                      - Secret
                      type: string
                    name:
                      description: Name of the referenced resource.
                      minLength: 1
                      type: string
                    namespace:
                      description: Namespace of the referenced resource.
                      minLength: 1
                      type: string
                    path:
                      description: Path to the directory containing the openapi validations.
                        Defaults to 'None', which translates to the root path of the
                        SourceRef. Ignored for ConfigMap/Secret.
                      type: string
                  required:
                  - kind
                  - name
                  - namespace
                  type: object
                type: array
              openAPIValidationRefs:
                description: OpenAPIValidationRefs is a list of OpenAPI validations.
                  In the matching clusters, add-ons will be deployed only if all validations
                  pass.
                items:
                  properties:
                    kind:
                      description: 'Kind of the resource. Supported kinds are: - flux
                        GitRepository;OCIRepository;Bucket - ConfigMap/Secret'
                      enum:
                      - GitRepository
                      - OCIRepository
                      - Bucket
                      - ConfigMap
                      - Secret
                      type: string
                    name:
                      description: Name of the referenced resource.
                      minLength: 1
                      type: string
                    namespace:
                      description: Namespace of the referenced resource.
                      minLength: 1
                      type: string
                    path:
                      description: Path to the directory containing the openapi validations.
                        Defaults to 'None', which translates to the root path of the
                        SourceRef. Ignored for ConfigMap/Secret.
                      type: string
                  required:
                  - kind
                  - name
                  - namespace
                  type: object
                type: array
            type: object
          status:
            description: AddonComplianceStatus defines the observed state of AddonCompliance
            properties:
              failureMessage:
                description: FailureMessage provides more information if an error
                  occurs.
                type: string
              luaValidations:
                additionalProperties:
                  format: byte
                  type: string
                description: LuaValidations contains all validations collected from
                  all existing referenced resources
                type: object
              matchingClusters:
                description: MatchingClusterRefs reference all the clusters currently
                  matching ClusterSelector
                items:
                  description: "ObjectReference contains enough information to let
                    you inspect or modify the referred object. --- New uses of this
                    type are discouraged because of difficulty describing its usage
                    when embedded in APIs. 1. Ignored fields.  It includes many fields
                    which are not generally honored.  For instance, ResourceVersion
                    and FieldPath are both very rarely valid in actual usage. 2. Invalid
                    usage help.  It is impossible to add specific help for individual
                    usage.  In most embedded usages, there are particular restrictions
                    like, \"must refer only to types A and B\" or \"UID not honored\"
                    or \"name must be restricted\". Those cannot be well described
                    when embedded. 3. Inconsistent validation.  Because the usages
                    are different, the validation rules are different by usage, which
                    makes it hard for users to predict what will happen. 4. The fields
                    are both imprecise and overly precise.  Kind is not a precise
                    mapping to a URL. This can produce ambiguity during interpretation
                    and require a REST mapping.  In most cases, the dependency is
                    on the group,resource tuple and the version of the actual struct
                    is irrelevant. 5. We cannot easily change it.  Because this type
                    is embedded in many locations, updates to this type will affect
                    numerous schemas.  Don't make new APIs embed an underspecified
                    API type they do not control. \n Instead of using this type, create
                    a locally provided and used type that is well-focused on your
                    reference. For example, ServiceReferences for admission registration:
                    https://github.com/kubernetes/api/blob/release-1.17/admissionregistration/v1/types.go#L533
                    ."
                  properties:
                    apiVersion:
                      description: API version of the referent.
                      type: string
                    fieldPath:
                      description: 'If referring to a piece of an object instead of
                        an entire object, this string should contain a valid JSON/Go
                        field access statement, such as desiredState.manifest.containers[2].
                        For example, if the object reference is to a container within
                        a pod, this would take on a value like: "spec.containers{name}"
                        (where "name" refers to the name of the container that triggered
                        the event) or if no container name is specified "spec.containers[2]"
                        (container with index 2 in this pod). This syntax is chosen
                        only to have some well-defined way of referencing a part of
                        an object. TODO: this design is not final and this field is
                        subject to change in the future.'
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
                  x-kubernetes-map-type: atomic
                type: array
              openapiValidations:
                additionalProperties:
                  format: byte
                  type: string
                description: OpenapiValidations contains all validations collected
                  from all existing referenced resources
                type: object
            type: object
        type: object
    served: true
    storage: true
    subresources:
      status: {}
`)
