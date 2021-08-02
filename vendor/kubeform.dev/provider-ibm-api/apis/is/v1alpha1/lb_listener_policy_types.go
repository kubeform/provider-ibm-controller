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

type LbListenerPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbListenerPolicySpec   `json:"spec,omitempty"`
	Status            LbListenerPolicyStatus `json:"status,omitempty"`
}

type LbListenerPolicySpecRules struct {
	// Condition of the rule
	Condition *string `json:"condition" tf:"condition"`
	// HTTP header field. This is only applicable to rule type.
	// +optional
	Field *string `json:"field,omitempty" tf:"field"`
	// Rule ID
	// +optional
	RuleID *string `json:"ruleID,omitempty" tf:"rule_id"`
	// Type of the rule
	Type *string `json:"type" tf:"type"`
	// Value to be matched for rule condition
	Value *string `json:"value" tf:"value"`
}

type LbListenerPolicySpec struct {
	State *LbListenerPolicySpecResource `json:"state,omitempty" tf:"-"`

	Resource LbListenerPolicySpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type LbListenerPolicySpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Policy Action
	Action *string `json:"action" tf:"action"`
	// Load Balancer Listener Policy
	Lb *string `json:"lb" tf:"lb"`
	// Listener ID
	Listener *string `json:"listener" tf:"listener"`
	// Policy name
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// Listener Policy ID
	// +optional
	PolicyID *string `json:"policyID,omitempty" tf:"policy_id"`
	// Listener Policy Priority
	Priority *int64 `json:"priority" tf:"priority"`
	// Listner Policy status
	// +optional
	ProvisioningStatus *string `json:"provisioningStatus,omitempty" tf:"provisioning_status"`
	// The crn of the LB resource
	// +optional
	RelatedCrn *string `json:"relatedCrn,omitempty" tf:"related_crn"`
	// Policy Rules
	// +optional
	Rules []LbListenerPolicySpecRules `json:"rules,omitempty" tf:"rules"`
	// Listener Policy target HTTPS Status code.
	// +optional
	TargetHTTPStatusCode *int64 `json:"targetHTTPStatusCode,omitempty" tf:"target_http_status_code"`
	// Listener Policy Target ID
	// +optional
	TargetID *string `json:"targetID,omitempty" tf:"target_id"`
	// Policy Target URL
	// +optional
	TargetURL *string `json:"targetURL,omitempty" tf:"target_url"`
}

type LbListenerPolicyStatus struct {
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

// LbListenerPolicyList is a list of LbListenerPolicys
type LbListenerPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LbListenerPolicy CRD objects
	Items []LbListenerPolicy `json:"items,omitempty"`
}
