// Copyright 2022 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// GENERATED BY gen_go_data.go
// gen_go_data -package beta -var YAML_mesh blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/networkservices/beta/mesh.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/networkservices/beta/mesh.yaml
var YAML_mesh = []byte("info:\n  title: NetworkServices/Mesh\n  description: The NetworkServices Mesh resource\n  x-dcl-struct-name: Mesh\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Mesh\n    parameters:\n    - name: mesh\n      required: true\n      description: A full instance of a Mesh\n  apply:\n    description: The function used to apply information about a Mesh\n    parameters:\n    - name: mesh\n      required: true\n      description: A full instance of a Mesh\n  delete:\n    description: The function used to delete a Mesh\n    parameters:\n    - name: mesh\n      required: true\n      description: A full instance of a Mesh\n  deleteAll:\n    description: The function used to delete all Mesh\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Mesh\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Mesh:\n      title: Mesh\n      x-dcl-id: projects/{{project}}/locations/{{location}}/meshes/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-labels: labels\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - name\n      - project\n      - location\n      properties:\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The timestamp when the resource was created.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Optional. A free-text description of the resource. Max length\n            1024 characters.\n        interceptionPort:\n          type: integer\n          format: int64\n          x-dcl-go-name: InterceptionPort\n          description: Optional. If set to a valid TCP port (1-65535), instructs the\n            SIDECAR proxy to listen on the specified port of localhost (127.0.0.1)\n            address. The SIDECAR proxy will expect all traffic to be redirected to\n            this port regardless of its actual ip:port destination. If unset, a port\n            '15001' is used as the interception port. This field is only valid if\n            the type of Mesh is SIDECAR.\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: Optional. Set of label tags associated with the Mesh resource.\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Required. Name of the Mesh resource. It matches pattern `projects/*/locations/global/meshes/`.\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        selfLink:\n          type: string\n          x-dcl-go-name: SelfLink\n          readOnly: true\n          description: Output only. Server-defined URL of this resource\n          x-kubernetes-immutable: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. The timestamp when the resource was updated.\n          x-kubernetes-immutable: true\n")

// 3913 bytes
// MD5: a150cac5f2e7dd7d93b5aaca92bde491
