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

type Glb struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlbSpec   `json:"spec,omitempty"`
	Status            GlbStatus `json:"status,omitempty"`
}

type GlbSpecAzPools struct {
	// Availability zone.
	AvailabilityZone *string `json:"availabilityZone" tf:"availability_zone"`
	// List of load balancer pools
	Pools []string `json:"pools" tf:"pools"`
}

type GlbSpec struct {
	State *GlbSpecResource `json:"state,omitempty" tf:"-"`

	Resource GlbSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type GlbSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Map availability zones to pool ID's.
	// +optional
	AzPools []GlbSpecAzPools `json:"azPools,omitempty" tf:"az_pools"`
	// GLB Load Balancer creation date
	// +optional
	CreatedOn *string `json:"createdOn,omitempty" tf:"created_on"`
	// A list of pool IDs ordered by their failover priority
	DefaultPools []string `json:"defaultPools" tf:"default_pools"`
	// Descriptive text of the load balancer
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// Whether the load balancer is enabled
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
	// The pool ID to use when all other pools are detected as unhealthy
	FallbackPool *string `json:"fallbackPool" tf:"fallback_pool"`
	// Load balancer Id
	// +optional
	GlbID *string `json:"glbID,omitempty" tf:"glb_id"`
	// Healthy state of the load balancer.
	// +optional
	Health *string `json:"health,omitempty" tf:"health"`
	// The GUID of the private DNS.
	InstanceID *string `json:"instanceID" tf:"instance_id"`
	// GLB Load Balancer Modification date
	// +optional
	ModifiedOn *string `json:"modifiedOn,omitempty" tf:"modified_on"`
	// Name of the load balancer
	Name *string `json:"name" tf:"name"`
	// Time to live in second
	// +optional
	Ttl *int64 `json:"ttl,omitempty" tf:"ttl"`
	// Zone Id
	ZoneID *string `json:"zoneID" tf:"zone_id"`
}

type GlbStatus struct {
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

// GlbList is a list of Glbs
type GlbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Glb CRD objects
	Items []Glb `json:"items,omitempty"`
}
