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

type SubnetReservedIP struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubnetReservedIPSpec   `json:"spec,omitempty"`
	Status            SubnetReservedIPStatus `json:"status,omitempty"`
}

type SubnetReservedIPSpec struct {
	State *SubnetReservedIPSpecResource `json:"state,omitempty" tf:"-"`

	Resource SubnetReservedIPSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type SubnetReservedIPSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The user-defined or system-provided name for this reserved IP.
	// +optional
	Address *string `json:"address,omitempty" tf:"address"`
	// If set to true, this reserved IP will be automatically deleted
	// +optional
	AutoDelete *bool `json:"autoDelete,omitempty" tf:"auto_delete"`
	// The date and time that the reserved IP was created.
	// +optional
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at"`
	// The URL for this reserved IP.
	// +optional
	Href *string `json:"href,omitempty" tf:"href"`
	// The user-defined or system-provided name for this reserved IP.
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// The owner of a reserved IP, defining whether it is managed by the user or the provider.
	// +optional
	Owner *string `json:"owner,omitempty" tf:"owner"`
	// The unique identifier of the reserved IP.
	// +optional
	ReservedIP *string `json:"reservedIP,omitempty" tf:"reserved_ip"`
	// The resource type.
	// +optional
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type"`
	// The subnet identifier.
	Subnet *string `json:"subnet" tf:"subnet"`
	// The unique identifier for target.
	// +optional
	Target *string `json:"target,omitempty" tf:"target"`
}

type SubnetReservedIPStatus struct {
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

// SubnetReservedIPList is a list of SubnetReservedIPs
type SubnetReservedIPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SubnetReservedIP CRD objects
	Items []SubnetReservedIP `json:"items,omitempty"`
}