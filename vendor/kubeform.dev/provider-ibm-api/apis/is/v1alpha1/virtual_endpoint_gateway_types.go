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

type VirtualEndpointGateway struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualEndpointGatewaySpec   `json:"spec,omitempty"`
	Status            VirtualEndpointGatewayStatus `json:"status,omitempty"`
}

type VirtualEndpointGatewaySpecIps struct {
	// The IP Address
	// +optional
	Address *string `json:"address,omitempty" tf:"address"`
	// The IPs id
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// The IPs name
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// The VPE Resource Type
	// +optional
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type"`
	// The Subnet id
	// +optional
	Subnet *string `json:"subnet,omitempty" tf:"subnet"`
}

type VirtualEndpointGatewaySpecTarget struct {
	// The target crn
	// +optional
	Crn *string `json:"crn,omitempty" tf:"crn"`
	// The target name
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// The target resource type
	ResourceType *string `json:"resourceType" tf:"resource_type"`
}

type VirtualEndpointGatewaySpec struct {
	State *VirtualEndpointGatewaySpecResource `json:"state,omitempty" tf:"-"`

	Resource VirtualEndpointGatewaySpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type VirtualEndpointGatewaySpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Endpoint gateway created date and time
	// +optional
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at"`
	// Endpoint gateway health state
	// +optional
	HealthState *string `json:"healthState,omitempty" tf:"health_state"`
	// Endpoint gateway IPs
	// +optional
	Ips []VirtualEndpointGatewaySpecIps `json:"ips,omitempty" tf:"ips"`
	// Endpoint gateway lifecycle state
	// +optional
	LifecycleState *string `json:"lifecycleState,omitempty" tf:"lifecycle_state"`
	// Endpoint gateway name
	Name *string `json:"name" tf:"name"`
	// The resource group id
	// +optional
	ResourceGroup *string `json:"resourceGroup,omitempty" tf:"resource_group"`
	// Endpoint gateway resource type
	// +optional
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type"`
	// List of tags for VPE
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags"`
	// Endpoint gateway target
	Target *VirtualEndpointGatewaySpecTarget `json:"target" tf:"target"`
	// The VPC id
	Vpc *string `json:"vpc" tf:"vpc"`
}

type VirtualEndpointGatewayStatus struct {
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

// VirtualEndpointGatewayList is a list of VirtualEndpointGateways
type VirtualEndpointGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VirtualEndpointGateway CRD objects
	Items []VirtualEndpointGateway `json:"items,omitempty"`
}
