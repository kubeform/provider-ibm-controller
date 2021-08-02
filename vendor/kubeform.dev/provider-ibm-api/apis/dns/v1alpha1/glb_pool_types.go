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

type GlbPool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlbPoolSpec   `json:"spec,omitempty"`
	Status            GlbPoolStatus `json:"status,omitempty"`
}

type GlbPoolSpecOrigins struct {
	// The address of the origin server. It can be a hostname or an IP address.
	Address *string `json:"address" tf:"address"`
	// Description of the origin server.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// Whether the origin server is enabled.
	Enabled *bool `json:"enabled" tf:"enabled"`
	// Whether the health is `true` or `false`.
	// +optional
	Health *bool `json:"health,omitempty" tf:"health"`
	// The Reason for health check failure
	// +optional
	HealthFailureReason *string `json:"healthFailureReason,omitempty" tf:"health_failure_reason"`
	// The name of the origin server.
	Name *string `json:"name" tf:"name"`
}

type GlbPoolSpec struct {
	State *GlbPoolSpecResource `json:"state,omitempty" tf:"-"`

	Resource GlbPoolSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type GlbPoolSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The time when a load balancer pool is created.
	// +optional
	CreatedOn *string `json:"createdOn,omitempty" tf:"created_on"`
	// Descriptive text of the load balancer pool
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// Whether the load balancer pool is enabled
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
	// Whether the load balancer pool is enabled
	// +optional
	Health *string `json:"health,omitempty" tf:"health"`
	// Health check region of VSIs
	// +optional
	HealthcheckRegion *string `json:"healthcheckRegion,omitempty" tf:"healthcheck_region"`
	// Health check subnet crn of VSIs
	// +optional
	HealthcheckSubnets []string `json:"healthcheckSubnets,omitempty" tf:"healthcheck_subnets"`
	// The minimum number of origins that must be healthy for this pool to serve traffic
	// +optional
	HealthyOriginsThreshold *int64 `json:"healthyOriginsThreshold,omitempty" tf:"healthy_origins_threshold"`
	// Instance Id
	InstanceID *string `json:"instanceID" tf:"instance_id"`
	// The recent time when a load balancer pool is modified.
	// +optional
	ModifiedOn *string `json:"modifiedOn,omitempty" tf:"modified_on"`
	// The ID of the load balancer monitor to be associated to this pool
	// +optional
	Monitor *string `json:"monitor,omitempty" tf:"monitor"`
	// The unique identifier of a service instance.
	Name *string `json:"name" tf:"name"`
	// The notification channel,It is a webhook url
	// +optional
	NotificationChannel *string `json:"notificationChannel,omitempty" tf:"notification_channel"`
	// Origins info
	Origins []GlbPoolSpecOrigins `json:"origins" tf:"origins"`
	// Pool Id
	// +optional
	PoolID *string `json:"poolID,omitempty" tf:"pool_id"`
}

type GlbPoolStatus struct {
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

// GlbPoolList is a list of GlbPools
type GlbPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GlbPool CRD objects
	Items []GlbPool `json:"items,omitempty"`
}