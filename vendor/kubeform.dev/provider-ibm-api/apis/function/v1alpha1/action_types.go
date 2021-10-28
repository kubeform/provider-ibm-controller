/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type Action struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ActionSpec   `json:"spec,omitempty"`
	Status            ActionStatus `json:"status,omitempty"`
}

type ActionSpecExec struct {
	// The code to execute.
	// +optional
	Code *string `json:"code,omitempty" tf:"code"`
	// The file path of code to execute.
	// +optional
	CodePath *string `json:"codePath,omitempty" tf:"code_path"`
	// The List of fully qualified action.
	// +optional
	Components []string `json:"components,omitempty" tf:"components"`
	// Container image name when kind is 'blackbox'.
	// +optional
	Image *string `json:"image,omitempty" tf:"image"`
	// Optional zipfile reference.
	// +optional
	Init *string `json:"init,omitempty" tf:"init"`
	// The type of action. Possible values can be found here (https://cloud.ibm.com/docs/openwhisk?topic=cloud-functions-runtimes)
	Kind *string `json:"kind" tf:"kind"`
	// The name of the action entry point (function or fully-qualified method name when applicable).
	// +optional
	Main *string `json:"main,omitempty" tf:"main"`
}

type ActionSpecLimits struct {
	// The maximum log size LIMIT in MB for the action.
	// +optional
	LogSize *int64 `json:"logSize,omitempty" tf:"log_size"`
	// The maximum memory LIMIT in MB for the action (default 256.
	// +optional
	Memory *int64 `json:"memory,omitempty" tf:"memory"`
	// The timeout LIMIT in milliseconds after which the action is terminated.
	// +optional
	Timeout *int64 `json:"timeout,omitempty" tf:"timeout"`
}

type ActionSpec struct {
	State *ActionSpecResource `json:"state,omitempty" tf:"-"`

	Resource ActionSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ActionSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ActionID *string `json:"actionID,omitempty" tf:"action_id"`
	// All annotations set on action by user and those set by the IBM Cloud Function backend/API.
	// +optional
	Annotations *string `json:"annotations,omitempty" tf:"annotations"`
	// Execution info
	Exec *ActionSpecExec `json:"exec" tf:"exec"`
	// +optional
	Limits *ActionSpecLimits `json:"limits,omitempty" tf:"limits"`
	// Name of action.
	Name *string `json:"name" tf:"name"`
	// IBM Cloud function namespace.
	Namespace *string `json:"namespace" tf:"namespace"`
	// All paramters set on action by user and those set by the IBM Cloud Function backend/API.
	// +optional
	Parameters *string `json:"parameters,omitempty" tf:"parameters"`
	// Action visibilty.
	// +optional
	Publish *bool `json:"publish,omitempty" tf:"publish"`
	// Action target endpoint URL.
	// +optional
	TargetEndpointURL *string `json:"targetEndpointURL,omitempty" tf:"target_endpoint_url"`
	// Annotation values in KEY VALUE format.
	// +optional
	UserDefinedAnnotations *string `json:"userDefinedAnnotations,omitempty" tf:"user_defined_annotations"`
	// Parameters values in KEY VALUE format. Parameter bindings included in the context passed to the action.
	// +optional
	UserDefinedParameters *string `json:"userDefinedParameters,omitempty" tf:"user_defined_parameters"`
	// Semantic version of the item.
	// +optional
	Version *string `json:"version,omitempty" tf:"version"`
}

type ActionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ActionList is a list of Actions
type ActionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Action CRD objects
	Items []Action `json:"items,omitempty"`
}
