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

type SecurityGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecurityGroupSpec   `json:"spec,omitempty"`
	Status            SecurityGroupStatus `json:"status,omitempty"`
}

type SecurityGroupSpecRules struct {
	// +optional
	Code *int64 `json:"code,omitempty" tf:"code"`
	// Direction of traffic to enforce, either inbound or outbound
	// +optional
	Direction *string `json:"direction,omitempty" tf:"direction"`
	// IP version: ipv4
	// +optional
	IpVersion *string `json:"ipVersion,omitempty" tf:"ip_version"`
	// +optional
	PortMax *int64 `json:"portMax,omitempty" tf:"port_max"`
	// +optional
	PortMin *int64 `json:"portMin,omitempty" tf:"port_min"`
	// +optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol"`
	// Security group id: an IP address, a CIDR block, or a single security group identifier
	// +optional
	Remote *string `json:"remote,omitempty" tf:"remote"`
	// +optional
	Type *int64 `json:"type,omitempty" tf:"type"`
}

type SecurityGroupSpec struct {
	State *SecurityGroupSpecResource `json:"state,omitempty" tf:"-"`

	Resource SecurityGroupSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type SecurityGroupSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The crn of the resource
	// +optional
	Crn *string `json:"crn,omitempty" tf:"crn"`
	// Security group name
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance
	// +optional
	ResourceControllerURL *string `json:"resourceControllerURL,omitempty" tf:"resource_controller_url"`
	// The crn of the resource
	// +optional
	ResourceCrn *string `json:"resourceCrn,omitempty" tf:"resource_crn"`
	// Resource Group ID
	// +optional
	ResourceGroup *string `json:"resourceGroup,omitempty" tf:"resource_group"`
	// The resource group name in which resource is provisioned
	// +optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name"`
	// The name of the resource
	// +optional
	ResourceName *string `json:"resourceName,omitempty" tf:"resource_name"`
	// Security Rules
	// +optional
	Rules []SecurityGroupSpecRules `json:"rules,omitempty" tf:"rules"`
	// List of tags
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags"`
	// Security group's resource group id
	Vpc *string `json:"vpc" tf:"vpc"`
}

type SecurityGroupStatus struct {
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

// SecurityGroupList is a list of SecurityGroups
type SecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SecurityGroup CRD objects
	Items []SecurityGroup `json:"items,omitempty"`
}
