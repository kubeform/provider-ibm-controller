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

type Firewall struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallSpec   `json:"spec,omitempty"`
	Status            FirewallStatus `json:"status,omitempty"`
}

type FirewallSpecAccessRuleConfiguration struct {
	// Target type
	Target *string `json:"target" tf:"target"`
	// Target value
	Value *string `json:"value" tf:"value"`
}

type FirewallSpecAccessRule struct {
	// access rule firewall identifier
	// +optional
	AccessRuleID  *string                              `json:"accessRuleID,omitempty" tf:"access_rule_id"`
	Configuration *FirewallSpecAccessRuleConfiguration `json:"configuration" tf:"configuration"`
	// Access rule mode
	Mode *string `json:"mode" tf:"mode"`
	// description
	// +optional
	Notes *string `json:"notes,omitempty" tf:"notes"`
}

type FirewallSpecLockdownConfigurations struct {
	// Target type
	Target *string `json:"target" tf:"target"`
	// Target value
	Value *string `json:"value" tf:"value"`
}

type FirewallSpecLockdown struct {
	// +kubebuilder:validation:MinItems=1
	Configurations []FirewallSpecLockdownConfigurations `json:"configurations" tf:"configurations"`
	// description
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// firewall identifier
	// +optional
	LockdownID *string `json:"lockdownID,omitempty" tf:"lockdown_id"`
	// Firewall rule paused or enabled
	// +optional
	Paused *bool `json:"paused,omitempty" tf:"paused"`
	// Firewall priority
	// +optional
	Priority *int64 `json:"priority,omitempty" tf:"priority"`
	// URL in which firewall rule is applied
	Urls []string `json:"urls" tf:"urls"`
}

type FirewallSpecUaRuleConfiguration struct {
	// Target type
	Target *string `json:"target" tf:"target"`
	// Target value
	Value *string `json:"value" tf:"value"`
}

type FirewallSpecUaRule struct {
	Configuration *FirewallSpecUaRuleConfiguration `json:"configuration" tf:"configuration"`
	// description
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// user agent rule mode
	Mode *string `json:"mode" tf:"mode"`
	// Rule whether paused or not
	// +optional
	Paused *bool `json:"paused,omitempty" tf:"paused"`
	// User Agent firewall identifier
	// +optional
	UaRuleID *string `json:"uaRuleID,omitempty" tf:"ua_rule_id"`
}

type FirewallSpec struct {
	State *FirewallSpecResource `json:"state,omitempty" tf:"-"`

	Resource FirewallSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type FirewallSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Access Rule Data
	// +optional
	AccessRule *FirewallSpecAccessRule `json:"accessRule,omitempty" tf:"access_rule"`
	// CIS object id
	CisID *string `json:"cisID" tf:"cis_id"`
	// Associated CIS domain
	DomainID *string `json:"domainID" tf:"domain_id"`
	// Type of firewall.Allowable values are access-rules,ua-rules,lockdowns
	FirewallType *string `json:"firewallType" tf:"firewall_type"`
	// Lockdown Data
	// +optional
	Lockdown *FirewallSpecLockdown `json:"lockdown,omitempty" tf:"lockdown"`
	// User Agent Rule Data
	// +optional
	UaRule *FirewallSpecUaRule `json:"uaRule,omitempty" tf:"ua_rule"`
}

type FirewallStatus struct {
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

// FirewallList is a list of Firewalls
type FirewallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Firewall CRD objects
	Items []Firewall `json:"items,omitempty"`
}
