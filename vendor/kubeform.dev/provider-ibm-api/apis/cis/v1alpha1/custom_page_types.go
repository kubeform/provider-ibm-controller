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

type CustomPage struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CustomPageSpec   `json:"spec,omitempty"`
	Status            CustomPageStatus `json:"status,omitempty"`
}

type CustomPageSpec struct {
	State *CustomPageSpecResource `json:"state,omitempty" tf:"-"`

	Resource CustomPageSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type CustomPageSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// CIS instance crn
	CisID *string `json:"cisID" tf:"cis_id"`
	// Custom page created date
	// +optional
	CreatedOn *string `json:"createdOn,omitempty" tf:"created_on"`
	// Free text
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// Associated CIS domain
	DomainID *string `json:"domainID" tf:"domain_id"`
	// Custom page modified date
	// +optional
	ModifiedOn *string `json:"modifiedOn,omitempty" tf:"modified_on"`
	// Custom page identifier
	PageID *string `json:"pageID" tf:"page_id"`
	// Custom page preview target
	// +optional
	PreviewTarget *string `json:"previewTarget,omitempty" tf:"preview_target"`
	// Custom page state
	// +optional
	RequiredTokens []string `json:"requiredTokens,omitempty" tf:"required_tokens"`
	// Custom page state
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// Custom page url
	Url *string `json:"url" tf:"url"`
}

type CustomPageStatus struct {
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

// CustomPageList is a list of CustomPages
type CustomPageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CustomPage CRD objects
	Items []CustomPage `json:"items,omitempty"`
}
