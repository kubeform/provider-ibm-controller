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

type Endpoint struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EndpointSpec   `json:"spec,omitempty"`
	Status            EndpointStatus `json:"status,omitempty"`
}

type EndpointSpec struct {
	State *EndpointSpecResource `json:"state,omitempty" tf:"-"`

	Resource EndpointSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type EndpointSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	//  Base path of an endpoint
	// +optional
	BasePath *string `json:"basePath,omitempty" tf:"base_path"`
	// Endpoint ID
	// +optional
	EndpointID *string `json:"endpointID,omitempty" tf:"endpoint_id"`
	// Managed indicates if endpoint is online or offline.
	// +optional
	Managed *bool `json:"managed,omitempty" tf:"managed"`
	// Endpoint name
	Name *string `json:"name" tf:"name"`
	// Json File path
	OpenAPIDocName *string `json:"openAPIDocName" tf:"open_api_doc_name"`
	//  Provider ID of an endpoint allowable values user-defined and whisk
	// +optional
	ProviderID *string `json:"providerID,omitempty" tf:"provider_id"`
	// Invokable routes for an endpoint
	// +optional
	Routes []string `json:"routes,omitempty" tf:"routes"`
	// Api Gateway Service Instance Crn
	ServiceInstanceCrn *string `json:"serviceInstanceCrn" tf:"service_instance_crn"`
	// The Shared status of an endpoint
	// +optional
	Shared *bool `json:"shared,omitempty" tf:"shared"`
	// Action type of Endpoint ALoowable values are share, unshare, manage, unmanage
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type EndpointStatus struct {
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

// EndpointList is a list of Endpoints
type EndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Endpoint CRD objects
	Items []Endpoint `json:"items,omitempty"`
}
