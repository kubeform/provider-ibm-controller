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

type InstanceVolumeAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceVolumeAttachmentSpec   `json:"spec,omitempty"`
	Status            InstanceVolumeAttachmentStatus `json:"status,omitempty"`
}

type InstanceVolumeAttachmentSpec struct {
	State *InstanceVolumeAttachmentSpecResource `json:"state,omitempty" tf:"-"`

	Resource InstanceVolumeAttachmentSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type InstanceVolumeAttachmentSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The capacity of the volume in gigabytes. The specified minimum and maximum capacity values for creating or updating volumes may expand in the future.
	// +optional
	Capacity *int64 `json:"capacity,omitempty" tf:"capacity"`
	// If set to true, when deleting the attachment, the volume will also be deleted. Default value for this true.
	// +optional
	DeleteVolumeOnAttachmentDelete *bool `json:"deleteVolumeOnAttachmentDelete,omitempty" tf:"delete_volume_on_attachment_delete"`
	// If set to true, when deleting the instance the volume will also be deleted.
	// +optional
	DeleteVolumeOnInstanceDelete *bool `json:"deleteVolumeOnInstanceDelete,omitempty" tf:"delete_volume_on_instance_delete"`
	// A unique identifier for the device which is exposed to the instance operating system
	// +optional
	Device *string `json:"device,omitempty" tf:"device"`
	// The CRN of the [Key Protect Root Key](https://cloud.ibm.com/docs/key-protect?topic=key-protect-getting-started-tutorial) or [Hyper Protect Crypto Service Root Key](https://cloud.ibm.com/docs/hs-crypto?topic=hs-crypto-get-started) for this resource.
	// +optional
	EncryptionKey *string `json:"encryptionKey,omitempty" tf:"encryption_key"`
	// The URL for this volume attachment
	// +optional
	Href *string `json:"href,omitempty" tf:"href"`
	// Instance id
	Instance *string `json:"instance" tf:"instance"`
	// The maximum I/O operations per second (IOPS) for the volume.
	// +optional
	Iops *int64 `json:"iops,omitempty" tf:"iops"`
	// The user-defined name for this volume attachment.
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// The  globally unique name for the volume profile to use for this volume.
	// +optional
	Profile *string `json:"profile,omitempty" tf:"profile"`
	// The snapshot of the volume to be attached
	// +optional
	Snapshot *string `json:"snapshot,omitempty" tf:"snapshot"`
	// The status of this volume attachment, one of [ attached, attaching, deleting, detaching ]
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// The type of volume attachment one of [ boot, data ]
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
	// Instance id
	// +optional
	Volume *string `json:"volume,omitempty" tf:"volume"`
	// The unique identifier for this volume attachment
	// +optional
	VolumeAttachmentID *string `json:"volumeAttachmentID,omitempty" tf:"volume_attachment_id"`
	// The CRN for this volume
	// +optional
	VolumeCrn *string `json:"volumeCrn,omitempty" tf:"volume_crn"`
	// Link to documentation about deleted resources
	// +optional
	VolumeDeleted *string `json:"volumeDeleted,omitempty" tf:"volume_deleted"`
	// The URL for this volume
	// +optional
	VolumeHref *string `json:"volumeHref,omitempty" tf:"volume_href"`
	// The unique user-defined name for this volume
	// +optional
	VolumeName *string `json:"volumeName,omitempty" tf:"volume_name"`
}

type InstanceVolumeAttachmentStatus struct {
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

// InstanceVolumeAttachmentList is a list of InstanceVolumeAttachments
type InstanceVolumeAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of InstanceVolumeAttachment CRD objects
	Items []InstanceVolumeAttachment `json:"items,omitempty"`
}
