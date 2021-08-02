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

type InstanceTemplate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceTemplateSpec   `json:"spec,omitempty"`
	Status            InstanceTemplateStatus `json:"status,omitempty"`
}

type InstanceTemplateSpecBootVolume struct {
	// +optional
	DeleteVolumeOnInstanceDelete *bool `json:"deleteVolumeOnInstanceDelete,omitempty" tf:"delete_volume_on_instance_delete"`
	// +optional
	Encryption *string `json:"encryption,omitempty" tf:"encryption"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Profile *string `json:"profile,omitempty" tf:"profile"`
	// +optional
	Size *int64 `json:"size,omitempty" tf:"size"`
}

type InstanceTemplateSpecNetworkInterfaces struct {
	// +optional
	AllowIPSpoofing *bool `json:"allowIPSpoofing,omitempty" tf:"allow_ip_spoofing"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	PrimaryIpv4Address *string `json:"primaryIpv4Address,omitempty" tf:"primary_ipv4_address"`
	// +optional
	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups"`
	Subnet         *string  `json:"subnet" tf:"subnet"`
}

type InstanceTemplateSpecPlacementTarget struct {
	// The CRN for this dedicated host.
	// +optional
	Crn *string `json:"crn,omitempty" tf:"crn"`
	// The URL for this dedicated host.
	// +optional
	Href *string `json:"href,omitempty" tf:"href"`
	// The unique identifier for this dedicated host.
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
}

type InstanceTemplateSpecPrimaryNetworkInterface struct {
	// +optional
	AllowIPSpoofing *bool `json:"allowIPSpoofing,omitempty" tf:"allow_ip_spoofing"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	PrimaryIpv4Address *string `json:"primaryIpv4Address,omitempty" tf:"primary_ipv4_address"`
	// +optional
	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups"`
	Subnet         *string  `json:"subnet" tf:"subnet"`
}

type InstanceTemplateSpecVolumeAttachmentsVolumePrototype struct {
	// The capacity of the volume in gigabytes. The specified minimum and maximum capacity values for creating or updating volumes may expand in the future.
	Capacity *int64 `json:"capacity" tf:"capacity"`
	// The CRN of the [Key Protect Root Key](https://cloud.ibm.com/docs/key-protect?topic=key-protect-getting-started-tutorial) or [Hyper Protect Crypto Service Root Key](https://cloud.ibm.com/docs/hs-crypto?topic=hs-crypto-get-started) for this resource.
	// +optional
	EncryptionKey *string `json:"encryptionKey,omitempty" tf:"encryption_key"`
	// The maximum I/O operations per second (IOPS) for the volume.
	// +optional
	Iops *int64 `json:"iops,omitempty" tf:"iops"`
	// The  globally unique name for the volume profile to use for this volume.
	Profile *string `json:"profile" tf:"profile"`
}

type InstanceTemplateSpecVolumeAttachments struct {
	// If set to true, when deleting the instance the volume will also be deleted.
	DeleteVolumeOnInstanceDelete *bool `json:"deleteVolumeOnInstanceDelete" tf:"delete_volume_on_instance_delete"`
	// The user-defined name for this volume attachment.
	Name *string `json:"name" tf:"name"`
	// The unique identifier for this volume.
	// +optional
	Volume *string `json:"volume,omitempty" tf:"volume"`
	// +optional
	VolumePrototype *InstanceTemplateSpecVolumeAttachmentsVolumePrototype `json:"volumePrototype,omitempty" tf:"volume_prototype"`
}

type InstanceTemplateSpec struct {
	State *InstanceTemplateSpecResource `json:"state,omitempty" tf:"-"`

	Resource InstanceTemplateSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type InstanceTemplateSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	BootVolume *InstanceTemplateSpecBootVolume `json:"bootVolume,omitempty" tf:"boot_volume"`
	// Unique Identifier of the Dedicated Host where the instance will be placed
	// +optional
	DedicatedHost *string `json:"dedicatedHost,omitempty" tf:"dedicated_host"`
	// Unique Identifier of the Dedicated Host Group where the instance will be placed
	// +optional
	DedicatedHostGroup *string `json:"dedicatedHostGroup,omitempty" tf:"dedicated_host_group"`
	// image name
	Image *string `json:"image" tf:"image"`
	// SSH key Ids for the instance template
	Keys []string `json:"keys" tf:"keys"`
	// Instance Template name
	Name *string `json:"name" tf:"name"`
	// +optional
	NetworkInterfaces []InstanceTemplateSpecNetworkInterfaces `json:"networkInterfaces,omitempty" tf:"network_interfaces"`
	// The placement restrictions to use for the virtual server instance.
	// +optional
	PlacementTarget []InstanceTemplateSpecPlacementTarget `json:"placementTarget,omitempty" tf:"placement_target"`
	// Primary Network interface info
	PrimaryNetworkInterface *InstanceTemplateSpecPrimaryNetworkInterface `json:"primaryNetworkInterface" tf:"primary_network_interface"`
	// Profile info
	Profile *string `json:"profile" tf:"profile"`
	// Instance template resource group
	// +optional
	ResourceGroup *string `json:"resourceGroup,omitempty" tf:"resource_group"`
	// User data given for the instance
	// +optional
	UserData *string `json:"userData,omitempty" tf:"user_data"`
	// +optional
	VolumeAttachments []InstanceTemplateSpecVolumeAttachments `json:"volumeAttachments,omitempty" tf:"volume_attachments"`
	// VPC id
	Vpc *string `json:"vpc" tf:"vpc"`
	// Zone name
	Zone *string `json:"zone" tf:"zone"`
}

type InstanceTemplateStatus struct {
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

// InstanceTemplateList is a list of InstanceTemplates
type InstanceTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of InstanceTemplate CRD objects
	Items []InstanceTemplate `json:"items,omitempty"`
}
