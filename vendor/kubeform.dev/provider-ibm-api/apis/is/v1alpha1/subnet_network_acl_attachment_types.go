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

type SubnetNetworkACLAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubnetNetworkACLAttachmentSpec   `json:"spec,omitempty"`
	Status            SubnetNetworkACLAttachmentStatus `json:"status,omitempty"`
}

type SubnetNetworkACLAttachmentSpecRulesIcmp struct {
	// The ICMP traffic code to allow
	// +optional
	Code *int64 `json:"code,omitempty" tf:"code"`
	// The ICMP traffic type to allow
	// +optional
	Type *int64 `json:"type,omitempty" tf:"type"`
}

type SubnetNetworkACLAttachmentSpecRulesTcp struct {
	// The inclusive upper bound of TCP destination port range
	// +optional
	PortMax *int64 `json:"portMax,omitempty" tf:"port_max"`
	// The inclusive lower bound of TCP destination port range
	// +optional
	PortMin *int64 `json:"portMin,omitempty" tf:"port_min"`
	// The inclusive upper bound of TCP source port range
	// +optional
	SourcePortMax *int64 `json:"sourcePortMax,omitempty" tf:"source_port_max"`
	// The inclusive lower bound of TCP source port range
	// +optional
	SourcePortMin *int64 `json:"sourcePortMin,omitempty" tf:"source_port_min"`
}

type SubnetNetworkACLAttachmentSpecRulesUdp struct {
	// The inclusive upper bound of UDP destination port range
	// +optional
	PortMax *int64 `json:"portMax,omitempty" tf:"port_max"`
	// The inclusive lower bound of UDP destination port range
	// +optional
	PortMin *int64 `json:"portMin,omitempty" tf:"port_min"`
	// The inclusive upper bound of UDP source port range
	// +optional
	SourcePortMax *int64 `json:"sourcePortMax,omitempty" tf:"source_port_max"`
	// The inclusive lower bound of UDP source port range
	// +optional
	SourcePortMin *int64 `json:"sourcePortMin,omitempty" tf:"source_port_min"`
}

type SubnetNetworkACLAttachmentSpecRules struct {
	// Whether to allow or deny matching traffic
	// +optional
	Action *string `json:"action,omitempty" tf:"action"`
	// The destination CIDR block
	// +optional
	Destination *string `json:"destination,omitempty" tf:"destination"`
	// Direction of traffic to enforce, either inbound or outbound
	// +optional
	Direction *string `json:"direction,omitempty" tf:"direction"`
	// +optional
	Icmp []SubnetNetworkACLAttachmentSpecRulesIcmp `json:"icmp,omitempty" tf:"icmp"`
	// The unique identifier for this Network ACL rule
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// The IP version for this rule
	// +optional
	IpVersion *string `json:"ipVersion,omitempty" tf:"ip_version"`
	// The user-defined name for this rule
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// The source CIDR block
	// +optional
	Source *string `json:"source,omitempty" tf:"source"`
	// +optional
	Tcp []SubnetNetworkACLAttachmentSpecRulesTcp `json:"tcp,omitempty" tf:"tcp"`
	// +optional
	Udp []SubnetNetworkACLAttachmentSpecRulesUdp `json:"udp,omitempty" tf:"udp"`
}

type SubnetNetworkACLAttachmentSpec struct {
	State *SubnetNetworkACLAttachmentSpecResource `json:"state,omitempty" tf:"-"`

	Resource SubnetNetworkACLAttachmentSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type SubnetNetworkACLAttachmentSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Network ACL name
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// The unique identifier of network ACL
	NetworkACL *string `json:"networkACL" tf:"network_acl"`
	// Resource group ID for the network ACL
	// +optional
	ResourceGroup *string `json:"resourceGroup,omitempty" tf:"resource_group"`
	// +optional
	Rules []SubnetNetworkACLAttachmentSpecRules `json:"rules,omitempty" tf:"rules"`
	// The subnet identifier
	Subnet *string `json:"subnet" tf:"subnet"`
	// Network ACL VPC
	// +optional
	Vpc *string `json:"vpc,omitempty" tf:"vpc"`
}

type SubnetNetworkACLAttachmentStatus struct {
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

// SubnetNetworkACLAttachmentList is a list of SubnetNetworkACLAttachments
type SubnetNetworkACLAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SubnetNetworkACLAttachment CRD objects
	Items []SubnetNetworkACLAttachment `json:"items,omitempty"`
}
