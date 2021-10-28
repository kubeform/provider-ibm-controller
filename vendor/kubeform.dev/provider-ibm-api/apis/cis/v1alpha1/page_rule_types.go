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

type PageRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PageRuleSpec   `json:"spec,omitempty"`
	Status            PageRuleStatus `json:"status,omitempty"`
}

type PageRuleSpecActions struct {
	// Page rule target url
	ID *string `json:"ID" tf:"id"`
	// Page rule actions status code
	// +optional
	StatusCode *int64 `json:"statusCode,omitempty" tf:"status_code"`
	// Page rule actions value url
	// +optional
	Url *string `json:"url,omitempty" tf:"url"`
	// Page rule target url
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type PageRuleSpecTargetsConstraint struct {
	// Constraint operator
	Operator *string `json:"operator" tf:"operator"`
	// Constraint value
	Value *string `json:"value" tf:"value"`
}

type PageRuleSpecTargets struct {
	// Page rule constraint
	Constraint *PageRuleSpecTargetsConstraint `json:"constraint" tf:"constraint"`
	// Page rule target url
	Target *string `json:"target" tf:"target"`
}

type PageRuleSpec struct {
	State *PageRuleSpecResource `json:"state,omitempty" tf:"-"`

	Resource PageRuleSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type PageRuleSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Page rule actions
	Actions []PageRuleSpecActions `json:"actions" tf:"actions"`
	// CIS instance crn
	CisID *string `json:"cisID" tf:"cis_id"`
	// Associated CIS domain
	DomainID *string `json:"domainID" tf:"domain_id"`
	// Page rule priority
	// +optional
	Priority *int64 `json:"priority,omitempty" tf:"priority"`
	// +optional
	RuleID *string `json:"ruleID,omitempty" tf:"rule_id"`
	// Page Rule status
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// Page rule targets
	Targets []PageRuleSpecTargets `json:"targets" tf:"targets"`
}

type PageRuleStatus struct {
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

// PageRuleList is a list of PageRules
type PageRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PageRule CRD objects
	Items []PageRule `json:"items,omitempty"`
}
