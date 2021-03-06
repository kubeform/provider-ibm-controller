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

type AccountGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AccountGroupSpec   `json:"spec,omitempty"`
	Status            AccountGroupStatus `json:"status,omitempty"`
}

type AccountGroupSpec struct {
	State *AccountGroupSpecResource `json:"state,omitempty" tf:"-"`

	Resource AccountGroupSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type AccountGroupSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The time stamp at which the account group was created.
	// +optional
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at"`
	// The IAM ID of the user or service that created the account group.
	// +optional
	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by"`
	// The Cloud Resource Name (CRN) of the account group.
	// +optional
	Crn *string `json:"crn,omitempty" tf:"crn"`
	// The enterprise account ID.
	// +optional
	EnterpriseAccountID *string `json:"enterpriseAccountID,omitempty" tf:"enterprise_account_id"`
	// The enterprise ID that the account group is a part of.
	// +optional
	EnterpriseID *string `json:"enterpriseID,omitempty" tf:"enterprise_id"`
	// The path from the enterprise to this particular account group.
	// +optional
	EnterprisePath *string `json:"enterprisePath,omitempty" tf:"enterprise_path"`
	// The name of the account group. This field must have 3 - 60 characters.
	Name *string `json:"name" tf:"name"`
	// The CRN of the parent under which the account group will be created. The parent can be an existing account group or the enterprise itself.
	Parent *string `json:"parent" tf:"parent"`
	// The email address of the primary contact of the account group.
	// +optional
	PrimaryContactEmail *string `json:"primaryContactEmail,omitempty" tf:"primary_contact_email"`
	// The IAM ID of the primary contact for this account group, such as `IBMid-0123ABC`. The IAM ID must already exist.
	PrimaryContactIamID *string `json:"primaryContactIamID" tf:"primary_contact_iam_id"`
	// The state of the account group.
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// The time stamp at which the account group was last updated.
	// +optional
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at"`
	// The IAM ID of the user or service that updated the account group.
	// +optional
	UpdatedBy *string `json:"updatedBy,omitempty" tf:"updated_by"`
	// The URL of the account group.
	// +optional
	Url *string `json:"url,omitempty" tf:"url"`
}

type AccountGroupStatus struct {
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

// AccountGroupList is a list of AccountGroups
type AccountGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AccountGroup CRD objects
	Items []AccountGroup `json:"items,omitempty"`
}
