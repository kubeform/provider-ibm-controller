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

type DedicatedHost struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DedicatedHostSpec   `json:"spec,omitempty"`
	Status            DedicatedHostStatus `json:"status,omitempty"`
}

type DedicatedHostSpec struct {
	State *DedicatedHostSpecResource `json:"state,omitempty" tf:"-"`

	Resource DedicatedHostSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type DedicatedHostSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The capacity that the dedicated host's CPU allocation is restricted to.
	// +optional
	CpuCount *int64 `json:"cpuCount,omitempty" tf:"cpu_count"`
	// The data center in which the dedicatated host is to be provisioned.
	Datacenter *string `json:"datacenter" tf:"datacenter"`
	// The capacity that the dedicated host's disk allocation is restricted to.
	// +optional
	DiskCapacity *int64 `json:"diskCapacity,omitempty" tf:"disk_capacity"`
	// The domain of dedicatated host.
	Domain *string `json:"domain" tf:"domain"`
	// The flavor of the dedicatated host.
	// +optional
	Flavor *string `json:"flavor,omitempty" tf:"flavor"`
	// The host name of dedicatated host.
	Hostname *string `json:"hostname" tf:"hostname"`
	// The billing type for the dedicatated host.
	// +optional
	HourlyBilling *bool `json:"hourlyBilling,omitempty" tf:"hourly_billing"`
	// The capacity that the dedicated host's memory allocation is restricted to.
	// +optional
	MemoryCapacity *int64 `json:"memoryCapacity,omitempty" tf:"memory_capacity"`
	// The hostname of the primary router that the dedicated host is associated with.
	RouterHostname *string `json:"routerHostname" tf:"router_hostname"`
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags"`
	// +optional
	WaitTimeMinutes *int64 `json:"waitTimeMinutes,omitempty" tf:"wait_time_minutes"`
}

type DedicatedHostStatus struct {
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

// DedicatedHostList is a list of DedicatedHosts
type DedicatedHostList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DedicatedHost CRD objects
	Items []DedicatedHost `json:"items,omitempty"`
}
