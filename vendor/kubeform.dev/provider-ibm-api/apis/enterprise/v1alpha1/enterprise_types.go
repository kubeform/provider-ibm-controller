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

type Enterprise struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EnterpriseSpec   `json:"spec,omitempty"`
	Status            EnterpriseStatus `json:"status,omitempty"`
}

type EnterpriseSpec struct {
	State *EnterpriseSpecResource `json:"state,omitempty" tf:"-"`

	Resource EnterpriseSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type EnterpriseSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The time stamp at which the enterprise was created.
	// +optional
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at"`
	// The IAM ID of the user or service that created the enterprise.
	// +optional
	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by"`
	// The Cloud Resource Name (CRN) of the enterprise.
	// +optional
	Crn *string `json:"crn,omitempty" tf:"crn"`
	// A domain or subdomain for the enterprise, such as `example.com` or `my.example.com`.
	// +optional
	Domain *string `json:"domain,omitempty" tf:"domain"`
	// The enterprise account ID.
	// +optional
	EnterpriseAccountID *string `json:"enterpriseAccountID,omitempty" tf:"enterprise_account_id"`
	// The name of the enterprise. This field must have 3 - 60 characters.
	Name *string `json:"name" tf:"name"`
	// The email of the primary contact of the enterprise.
	// +optional
	PrimaryContactEmail *string `json:"primaryContactEmail,omitempty" tf:"primary_contact_email"`
	// The IAM ID of the enterprise primary contact, such as `IBMid-0123ABC`. The IAM ID must already exist.
	PrimaryContactIamID *string `json:"primaryContactIamID" tf:"primary_contact_iam_id"`
	// The ID of the account that is used to create the enterprise.
	SourceAccountID *string `json:"sourceAccountID" tf:"source_account_id"`
	// The state of the enterprise.
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// The time stamp at which the enterprise was last updated.
	// +optional
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at"`
	// The IAM ID of the user or service that updated the enterprise.
	// +optional
	UpdatedBy *string `json:"updatedBy,omitempty" tf:"updated_by"`
	// The URL of the enterprise.
	// +optional
	Url *string `json:"url,omitempty" tf:"url"`
}

type EnterpriseStatus struct {
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

// EnterpriseList is a list of Enterprises
type EnterpriseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Enterprise CRD objects
	Items []Enterprise `json:"items,omitempty"`
}
