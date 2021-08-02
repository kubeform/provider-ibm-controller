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

type VirtualConnection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualConnectionSpec   `json:"spec,omitempty"`
	Status            VirtualConnectionStatus `json:"status,omitempty"`
}

type VirtualConnectionSpec struct {
	State *VirtualConnectionSpecResource `json:"state,omitempty" tf:"-"`

	Resource VirtualConnectionSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type VirtualConnectionSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The date and time resource was created
	// +optional
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at"`
	// The Direct Link gateway identifier
	Gateway *string `json:"gateway" tf:"gateway"`
	// The user-defined name for this virtual connection. Virtualconnection names are unique within a gateway. This is the name of thevirtual connection itself, the network being connected may have its ownname attribute
	Name *string `json:"name" tf:"name"`
	// For virtual connections across two different IBM Cloud Accounts network_account indicates the account that owns the target network.
	// +optional
	NetworkAccount *string `json:"networkAccount,omitempty" tf:"network_account"`
	// Unique identifier of the target network. For type=vpc virtual connections this is the CRN of the target VPC. This field does not apply to type=classic connections.
	// +optional
	NetworkID *string `json:"networkID,omitempty" tf:"network_id"`
	// The crn of the Direct link gateway
	// +optional
	RelatedCrn *string `json:"relatedCrn,omitempty" tf:"related_crn"`
	// Status of the virtual connection.Possible values: [pending,attached,approval_pending,rejected,expired,deleting,detached_by_network_pending,detached_by_network]
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// The type of virtual connection.Allowable values (classic,vpc)
	Type *string `json:"type" tf:"type"`
	// The Direct Gateway virtual connection identifier
	// +optional
	VirtualConnectionID *string `json:"virtualConnectionID,omitempty" tf:"virtual_connection_id"`
}

type VirtualConnectionStatus struct {
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

// VirtualConnectionList is a list of VirtualConnections
type VirtualConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VirtualConnection CRD objects
	Items []VirtualConnection `json:"items,omitempty"`
}
