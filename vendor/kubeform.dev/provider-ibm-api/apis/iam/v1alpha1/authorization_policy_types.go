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

type AuthorizationPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AuthorizationPolicySpec   `json:"spec,omitempty"`
	Status            AuthorizationPolicyStatus `json:"status,omitempty"`
}

type AuthorizationPolicySpec struct {
	State *AuthorizationPolicySpecResource `json:"state,omitempty" tf:"-"`

	Resource AuthorizationPolicySpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type AuthorizationPolicySpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Role names of the policy definition
	Roles []string `json:"roles" tf:"roles"`
	// The source resource group Id
	// +optional
	SourceResourceGroupID *string `json:"sourceResourceGroupID,omitempty" tf:"source_resource_group_id"`
	// The source resource instance Id
	// +optional
	SourceResourceInstanceID *string `json:"sourceResourceInstanceID,omitempty" tf:"source_resource_instance_id"`
	// Resource type of source service
	// +optional
	SourceResourceType *string `json:"sourceResourceType,omitempty" tf:"source_resource_type"`
	// Account GUID of source service
	// +optional
	SourceServiceAccount *string `json:"sourceServiceAccount,omitempty" tf:"source_service_account"`
	// The source service name
	SourceServiceName *string `json:"sourceServiceName" tf:"source_service_name"`
	// The target resource group Id
	// +optional
	TargetResourceGroupID *string `json:"targetResourceGroupID,omitempty" tf:"target_resource_group_id"`
	// The target resource instance Id
	// +optional
	TargetResourceInstanceID *string `json:"targetResourceInstanceID,omitempty" tf:"target_resource_instance_id"`
	// Resource type of target service
	// +optional
	TargetResourceType *string `json:"targetResourceType,omitempty" tf:"target_resource_type"`
	// The target service name
	TargetServiceName *string `json:"targetServiceName" tf:"target_service_name"`
	// +optional
	Version *string `json:"version,omitempty" tf:"version"`
}

type AuthorizationPolicyStatus struct {
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

// AuthorizationPolicyList is a list of AuthorizationPolicys
type AuthorizationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AuthorizationPolicy CRD objects
	Items []AuthorizationPolicy `json:"items,omitempty"`
}
