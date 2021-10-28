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

type DedicatedHostSpecAvailableVcpu struct {
	// The VCPU architecture.
	// +optional
	Architecture *string `json:"architecture,omitempty" tf:"architecture"`
	// The number of VCPUs assigned.
	// +optional
	Count *int64 `json:"count,omitempty" tf:"count"`
}

type DedicatedHostSpecDisksInstanceDisksDeleted struct {
	// Link to documentation about deleted resources.
	// +optional
	MoreInfo *string `json:"moreInfo,omitempty" tf:"more_info"`
}

type DedicatedHostSpecDisksInstanceDisks struct {
	// If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
	// +optional
	Deleted []DedicatedHostSpecDisksInstanceDisksDeleted `json:"deleted,omitempty" tf:"deleted"`
	// The URL for this instance disk.
	// +optional
	Href *string `json:"href,omitempty" tf:"href"`
	// The unique identifier for this instance disk.
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// The user-defined name for this disk.
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// The resource type.
	// +optional
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type"`
}

type DedicatedHostSpecDisks struct {
	// The remaining space left for instance placement in GB (gigabytes).
	// +optional
	Available *int64 `json:"available,omitempty" tf:"available"`
	// The date and time that the disk was created.
	// +optional
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at"`
	// The URL for this disk.
	// +optional
	Href *string `json:"href,omitempty" tf:"href"`
	// The unique identifier for this disk.
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// Instance disks that are on this dedicated host disk.
	// +optional
	InstanceDisks []DedicatedHostSpecDisksInstanceDisks `json:"instanceDisks,omitempty" tf:"instance_disks"`
	// The disk interface used for attaching the diskThe enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.
	// +optional
	InterfaceType *string `json:"interfaceType,omitempty" tf:"interface_type"`
	// The lifecycle state of this dedicated host disk.
	// +optional
	LifecycleState *string `json:"lifecycleState,omitempty" tf:"lifecycle_state"`
	// The user-defined or system-provided name for this disk.
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// Indicates whether this dedicated host disk is available for instance disk creation.
	// +optional
	Provisionable *bool `json:"provisionable,omitempty" tf:"provisionable"`
	// The type of resource referenced.
	// +optional
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type"`
	// The size of the disk in GB (gigabytes).
	// +optional
	Size *int64 `json:"size,omitempty" tf:"size"`
	// The instance disk interfaces supported for this dedicated host disk.
	// +optional
	SupportedInstanceInterfaceTypes []string `json:"supportedInstanceInterfaceTypes,omitempty" tf:"supported_instance_interface_types"`
}

type DedicatedHostSpecInstancesDeleted struct {
	// Link to documentation about deleted resources.
	// +optional
	MoreInfo *string `json:"moreInfo,omitempty" tf:"more_info"`
}

type DedicatedHostSpecInstances struct {
	// The CRN for this virtual server instance.
	// +optional
	Crn *string `json:"crn,omitempty" tf:"crn"`
	// If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
	// +optional
	Deleted []DedicatedHostSpecInstancesDeleted `json:"deleted,omitempty" tf:"deleted"`
	// The URL for this virtual server instance.
	// +optional
	Href *string `json:"href,omitempty" tf:"href"`
	// The unique identifier for this virtual server instance.
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// The user-defined name for this virtual server instance (and default system hostname).
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
}

type DedicatedHostSpecSupportedInstanceProfiles struct {
	// The URL for this virtual server instance profile.
	// +optional
	Href *string `json:"href,omitempty" tf:"href"`
	// The globally unique name for this virtual server instance profile.
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
}

type DedicatedHostSpecVcpu struct {
	// The VCPU architecture.
	// +optional
	Architecture *string `json:"architecture,omitempty" tf:"architecture"`
	// The number of VCPUs assigned.
	// +optional
	Count *int64 `json:"count,omitempty" tf:"count"`
}

type DedicatedHostSpec struct {
	State *DedicatedHostSpecResource `json:"state,omitempty" tf:"-"`

	Resource DedicatedHostSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type DedicatedHostSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The amount of memory in gibibytes that is currently available for instances.
	// +optional
	AvailableMemory *int64 `json:"availableMemory,omitempty" tf:"available_memory"`
	// The available VCPU for the dedicated host.
	// +optional
	AvailableVcpu []DedicatedHostSpecAvailableVcpu `json:"availableVcpu,omitempty" tf:"available_vcpu"`
	// The date and time that the dedicated host was created.
	// +optional
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at"`
	// The CRN for this dedicated host.
	// +optional
	Crn *string `json:"crn,omitempty" tf:"crn"`
	// Collection of the dedicated host's disks.
	// +optional
	Disks []DedicatedHostSpecDisks `json:"disks,omitempty" tf:"disks"`
	// The unique identifier of the dedicated host group for this dedicated host.
	HostGroup *string `json:"hostGroup" tf:"host_group"`
	// The URL for this dedicated host.
	// +optional
	Href *string `json:"href,omitempty" tf:"href"`
	// If set to true, instances can be placed on this dedicated host.
	// +optional
	InstancePlacementEnabled *bool `json:"instancePlacementEnabled,omitempty" tf:"instance_placement_enabled"`
	// Array of instances that are allocated to this dedicated host.
	// +optional
	Instances []DedicatedHostSpecInstances `json:"instances,omitempty" tf:"instances"`
	// The lifecycle state of the dedicated host resource.
	// +optional
	LifecycleState *string `json:"lifecycleState,omitempty" tf:"lifecycle_state"`
	// The total amount of memory in gibibytes for this host.
	// +optional
	Memory *int64 `json:"memory,omitempty" tf:"memory"`
	// The unique user-defined name for this dedicated host. If unspecified, the name will be a hyphenated list of randomly-selected words.
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// The Globally unique name of the dedicated host profile to use for this dedicated host.
	Profile *string `json:"profile" tf:"profile"`
	// Indicates whether this dedicated host is available for instance creation.
	// +optional
	Provisionable *bool `json:"provisionable,omitempty" tf:"provisionable"`
	// The unique identifier for the resource group to use. If unspecified, the account's [default resourcegroup](https://cloud.ibm.com/apidocs/resource-manager#introduction) is used.
	// +optional
	ResourceGroup *string `json:"resourceGroup,omitempty" tf:"resource_group"`
	// The type of resource referenced.
	// +optional
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type"`
	// The total number of sockets for this host.
	// +optional
	SocketCount *int64 `json:"socketCount,omitempty" tf:"socket_count"`
	// The administrative state of the dedicated host.The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the dedicated host on which the unexpected property value was encountered.
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// Array of instance profiles that can be used by instances placed on this dedicated host.
	// +optional
	SupportedInstanceProfiles []DedicatedHostSpecSupportedInstanceProfiles `json:"supportedInstanceProfiles,omitempty" tf:"supported_instance_profiles"`
	// The total VCPU of the dedicated host.
	// +optional
	Vcpu []DedicatedHostSpecVcpu `json:"vcpu,omitempty" tf:"vcpu"`
	// The globally unique name of the zone this dedicated host resides in.
	// +optional
	Zone *string `json:"zone,omitempty" tf:"zone"`
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
