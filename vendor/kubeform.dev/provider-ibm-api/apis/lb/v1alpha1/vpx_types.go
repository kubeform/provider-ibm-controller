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

type Vpx struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpxSpec   `json:"spec,omitempty"`
	Status            VpxStatus `json:"status,omitempty"`
}

type VpxSpec struct {
	State *VpxSpecResource `json:"state,omitempty" tf:"-"`

	Resource VpxSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type VpxSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Datacenter name
	Datacenter *string `json:"datacenter" tf:"datacenter"`
	// IP address count
	IpCount *int64 `json:"ipCount" tf:"ip_count"`
	// management IP address
	// +optional
	ManagementIPAddress *string `json:"managementIPAddress,omitempty" tf:"management_ip_address"`
	// Name
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// Plan info
	Plan *string `json:"plan" tf:"plan"`
	// Private subnet
	// +optional
	PrivateSubnet *string `json:"privateSubnet,omitempty" tf:"private_subnet"`
	// Private VLAN id
	// +optional
	PrivateVLANID *int64 `json:"privateVLANID,omitempty" tf:"private_vlan_id"`
	// Public subnet
	// +optional
	PublicSubnet *string `json:"publicSubnet,omitempty" tf:"public_subnet"`
	// Piblic VLAN id
	// +optional
	PublicVLANID *int64 `json:"publicVLANID,omitempty" tf:"public_vlan_id"`
	// Speed value
	Speed *int64 `json:"speed" tf:"speed"`
	// List of the tags
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags"`
	// Type of the VPX
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
	// version info
	Version *string `json:"version" tf:"version"`
	// List of VIP ids
	// +optional
	VipPool []string `json:"vipPool,omitempty" tf:"vip_pool"`
}

type VpxStatus struct {
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

// VpxList is a list of Vpxs
type VpxList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Vpx CRD objects
	Items []Vpx `json:"items,omitempty"`
}
