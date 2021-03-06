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

type Org struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrgSpec   `json:"spec,omitempty"`
	Status            OrgStatus `json:"status,omitempty"`
}

type OrgSpec struct {
	State *OrgSpecResource `json:"state,omitempty" tf:"-"`

	Resource OrgSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type OrgSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The IBMID of the users who will have auditor role in this org, ex - user@example.com
	// +optional
	Auditors []string `json:"auditors,omitempty" tf:"auditors"`
	// The IBMID of the users who will have billing manager role in this org, ex - user@example.com
	// +optional
	BillingManagers []string `json:"billingManagers,omitempty" tf:"billing_managers"`
	// The IBMID of the users who will have manager role in this org, ex - user@example.com
	// +optional
	Managers []string `json:"managers,omitempty" tf:"managers"`
	// Org name, for example myorg@domain
	Name *string `json:"name" tf:"name"`
	// Org quota guid
	// +optional
	OrgQuotaDefinitionGuid *string `json:"orgQuotaDefinitionGuid,omitempty" tf:"org_quota_definition_guid"`
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags"`
	// The IBMID of the users who will have user role in this org, ex - user@example.com
	// +optional
	Users []string `json:"users,omitempty" tf:"users"`
}

type OrgStatus struct {
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

// OrgList is a list of Orgs
type OrgList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Org CRD objects
	Items []Org `json:"items,omitempty"`
}
