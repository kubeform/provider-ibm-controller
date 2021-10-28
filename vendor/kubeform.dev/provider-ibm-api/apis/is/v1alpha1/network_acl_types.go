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

type NetworkACL struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkACLSpec   `json:"spec,omitempty"`
	Status            NetworkACLStatus `json:"status,omitempty"`
}

type NetworkACLSpecRulesIcmp struct {
	// +optional
	Code *int64 `json:"code,omitempty" tf:"code"`
	// +optional
	Type *int64 `json:"type,omitempty" tf:"type"`
}

type NetworkACLSpecRulesTcp struct {
	// +optional
	PortMax *int64 `json:"portMax,omitempty" tf:"port_max"`
	// +optional
	PortMin *int64 `json:"portMin,omitempty" tf:"port_min"`
	// +optional
	SourcePortMax *int64 `json:"sourcePortMax,omitempty" tf:"source_port_max"`
	// +optional
	SourcePortMin *int64 `json:"sourcePortMin,omitempty" tf:"source_port_min"`
}

type NetworkACLSpecRulesUdp struct {
	// +optional
	PortMax *int64 `json:"portMax,omitempty" tf:"port_max"`
	// +optional
	PortMin *int64 `json:"portMin,omitempty" tf:"port_min"`
	// +optional
	SourcePortMax *int64 `json:"sourcePortMax,omitempty" tf:"source_port_max"`
	// +optional
	SourcePortMin *int64 `json:"sourcePortMin,omitempty" tf:"source_port_min"`
}

type NetworkACLSpecRules struct {
	Action      *string `json:"action" tf:"action"`
	Destination *string `json:"destination" tf:"destination"`
	// Direction of traffic to enforce, either inbound or outbound
	Direction *string `json:"direction" tf:"direction"`
	// +optional
	Icmp *NetworkACLSpecRulesIcmp `json:"icmp,omitempty" tf:"icmp"`
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	IpVersion *string `json:"ipVersion,omitempty" tf:"ip_version"`
	Name      *string `json:"name" tf:"name"`
	Source    *string `json:"source" tf:"source"`
	// +optional
	Subnets *int64 `json:"subnets,omitempty" tf:"subnets"`
	// +optional
	Tcp *NetworkACLSpecRulesTcp `json:"tcp,omitempty" tf:"tcp"`
	// +optional
	Udp *NetworkACLSpecRulesUdp `json:"udp,omitempty" tf:"udp"`
}

type NetworkACLSpec struct {
	State *NetworkACLSpecResource `json:"state,omitempty" tf:"-"`

	Resource NetworkACLSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type NetworkACLSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The crn of the resource
	// +optional
	Crn *string `json:"crn,omitempty" tf:"crn"`
	// Network ACL name
	Name *string `json:"name" tf:"name"`
	// The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance
	// +optional
	ResourceControllerURL *string `json:"resourceControllerURL,omitempty" tf:"resource_controller_url"`
	// The crn of the resource
	// +optional
	ResourceCrn *string `json:"resourceCrn,omitempty" tf:"resource_crn"`
	// Resource group ID for the network ACL
	// +optional
	ResourceGroup *string `json:"resourceGroup,omitempty" tf:"resource_group"`
	// The resource group name in which resource is provisioned
	// +optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name"`
	// The name of the resource
	// +optional
	ResourceName *string `json:"resourceName,omitempty" tf:"resource_name"`
	// +optional
	Rules []NetworkACLSpecRules `json:"rules,omitempty" tf:"rules"`
	// List of tags
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags"`
	// Network ACL VPC name
	// +optional
	Vpc *string `json:"vpc,omitempty" tf:"vpc"`
}

type NetworkACLStatus struct {
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

// NetworkACLList is a list of NetworkACLs
type NetworkACLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NetworkACL CRD objects
	Items []NetworkACL `json:"items,omitempty"`
}
