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

type Instance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec,omitempty"`
	Status            InstanceStatus `json:"status,omitempty"`
}

type InstanceSpecBootVolume struct {
	// +optional
	Encryption *string `json:"encryption,omitempty" tf:"encryption"`
	// +optional
	Iops *int64 `json:"iops,omitempty" tf:"iops"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Profile *string `json:"profile,omitempty" tf:"profile"`
	// +optional
	Size *int64 `json:"size,omitempty" tf:"size"`
	// +optional
	Snapshot *string `json:"snapshot,omitempty" tf:"snapshot"`
}

type InstanceSpecDisks struct {
	// The date and time that the disk was created.
	// +optional
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at"`
	// The URL for this instance disk.
	// +optional
	Href *string `json:"href,omitempty" tf:"href"`
	// The unique identifier for this instance disk.
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// The disk interface used for attaching the disk.The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.
	// +optional
	InterfaceType *string `json:"interfaceType,omitempty" tf:"interface_type"`
	// The user-defined name for this disk.
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// The resource type.
	// +optional
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type"`
	// The size of the disk in GB (gigabytes).
	// +optional
	Size *int64 `json:"size,omitempty" tf:"size"`
}

type InstanceSpecGpu struct {
	// +optional
	Cores *int64 `json:"cores,omitempty" tf:"cores"`
	// +optional
	Count *int64 `json:"count,omitempty" tf:"count"`
	// +optional
	Manufacturer *string `json:"manufacturer,omitempty" tf:"manufacturer"`
	// +optional
	Memory *int64 `json:"memory,omitempty" tf:"memory"`
	// +optional
	Model *string `json:"model,omitempty" tf:"model"`
}

type InstanceSpecNetworkInterfaces struct {
	// Indicates whether IP spoofing is allowed on this interface.
	// +optional
	AllowIPSpoofing *bool `json:"allowIPSpoofing,omitempty" tf:"allow_ip_spoofing"`
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	PrimaryIpv4Address *string `json:"primaryIpv4Address,omitempty" tf:"primary_ipv4_address"`
	// +optional
	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups"`
	Subnet         *string  `json:"subnet" tf:"subnet"`
}

type InstanceSpecPrimaryNetworkInterface struct {
	// Indicates whether IP spoofing is allowed on this interface.
	// +optional
	AllowIPSpoofing *bool `json:"allowIPSpoofing,omitempty" tf:"allow_ip_spoofing"`
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	// Deprecated
	PortSpeed *int64 `json:"portSpeed,omitempty" tf:"port_speed"`
	// +optional
	PrimaryIpv4Address *string `json:"primaryIpv4Address,omitempty" tf:"primary_ipv4_address"`
	// +optional
	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups"`
	Subnet         *string  `json:"subnet" tf:"subnet"`
}

type InstanceSpecStatusReasons struct {
	// A snake case string succinctly identifying the status reason
	// +optional
	Code *string `json:"code,omitempty" tf:"code"`
	// An explanation of the status reason
	// +optional
	Message *string `json:"message,omitempty" tf:"message"`
}

type InstanceSpecVcpu struct {
	// +optional
	Architecture *string `json:"architecture,omitempty" tf:"architecture"`
	// +optional
	Count *int64 `json:"count,omitempty" tf:"count"`
}

type InstanceSpecVolumeAttachments struct {
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	VolumeCrn *string `json:"volumeCrn,omitempty" tf:"volume_crn"`
	// +optional
	VolumeID *string `json:"volumeID,omitempty" tf:"volume_id"`
	// +optional
	VolumeName *string `json:"volumeName,omitempty" tf:"volume_name"`
}

type InstanceSpec struct {
	State *InstanceSpecResource `json:"state,omitempty" tf:"-"`

	Resource InstanceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type InstanceSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Auto delete volume along with instance
	// +optional
	AutoDeleteVolume *bool `json:"autoDeleteVolume,omitempty" tf:"auto_delete_volume"`
	// +optional
	BootVolume *InstanceSpecBootVolume `json:"bootVolume,omitempty" tf:"boot_volume"`
	// Unique Identifier of the Dedicated Host where the instance will be placed
	// +optional
	DedicatedHost *string `json:"dedicatedHost,omitempty" tf:"dedicated_host"`
	// Unique Identifier of the Dedicated Host Group where the instance will be placed
	// +optional
	DedicatedHostGroup *string `json:"dedicatedHostGroup,omitempty" tf:"dedicated_host_group"`
	// Collection of the instance's disks.
	// +optional
	Disks []InstanceSpecDisks `json:"disks,omitempty" tf:"disks"`
	// Define timeout to force the instances to start/stop in minutes.
	// +optional
	ForceRecoveryTime *int64 `json:"forceRecoveryTime,omitempty" tf:"force_recovery_time"`
	// +optional
	// Deprecated
	Gpu []InstanceSpecGpu `json:"gpu,omitempty" tf:"gpu"`
	// image id
	// +optional
	Image *string `json:"image,omitempty" tf:"image"`
	// Id of the instance template
	// +optional
	InstanceTemplate *string `json:"instanceTemplate,omitempty" tf:"instance_template"`
	// SSH key Ids for the instance
	// +optional
	Keys []string `json:"keys,omitempty" tf:"keys"`
	// Instance memory
	// +optional
	Memory *int64 `json:"memory,omitempty" tf:"memory"`
	// Instance name
	Name *string `json:"name" tf:"name"`
	// +optional
	NetworkInterfaces []InstanceSpecNetworkInterfaces `json:"networkInterfaces,omitempty" tf:"network_interfaces"`
	// Primary Network interface info
	// +optional
	PrimaryNetworkInterface *InstanceSpecPrimaryNetworkInterface `json:"primaryNetworkInterface,omitempty" tf:"primary_network_interface"`
	// Profile info
	// +optional
	Profile *string `json:"profile,omitempty" tf:"profile"`
	// The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance
	// +optional
	ResourceControllerURL *string `json:"resourceControllerURL,omitempty" tf:"resource_controller_url"`
	// The crn of the resource
	// +optional
	ResourceCrn *string `json:"resourceCrn,omitempty" tf:"resource_crn"`
	// Instance resource group
	// +optional
	ResourceGroup *string `json:"resourceGroup,omitempty" tf:"resource_group"`
	// The resource group name in which resource is provisioned
	// +optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name"`
	// The name of the resource
	// +optional
	ResourceName *string `json:"resourceName,omitempty" tf:"resource_name"`
	// The status of the resource
	// +optional
	ResourceStatus *string `json:"resourceStatus,omitempty" tf:"resource_status"`
	// instance status
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// The reasons for the current status (if any).
	// +optional
	StatusReasons []InstanceSpecStatusReasons `json:"statusReasons,omitempty" tf:"status_reasons"`
	// list of tags for the instance
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags"`
	// User data given for the instance
	// +optional
	UserData *string `json:"userData,omitempty" tf:"user_data"`
	// +optional
	Vcpu []InstanceSpecVcpu `json:"vcpu,omitempty" tf:"vcpu"`
	// +optional
	VolumeAttachments []InstanceSpecVolumeAttachments `json:"volumeAttachments,omitempty" tf:"volume_attachments"`
	// List of volumes
	// +optional
	Volumes []string `json:"volumes,omitempty" tf:"volumes"`
	// VPC id
	// +optional
	Vpc *string `json:"vpc,omitempty" tf:"vpc"`
	// Enables stopping of instance before deleting and waits till deletion is complete
	// +optional
	WaitBeforeDelete *bool `json:"waitBeforeDelete,omitempty" tf:"wait_before_delete"`
	// Zone name
	// +optional
	Zone *string `json:"zone,omitempty" tf:"zone"`
}

type InstanceStatus struct {
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

// InstanceList is a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Instance CRD objects
	Items []Instance `json:"items,omitempty"`
}
