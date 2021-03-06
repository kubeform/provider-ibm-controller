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

type Volume struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VolumeSpec   `json:"spec,omitempty"`
	Status            VolumeStatus `json:"status,omitempty"`
}

type VolumeSpec struct {
	State *VolumeSpecResource `json:"state,omitempty" tf:"-"`

	Resource VolumeSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type VolumeSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Should the volume be deleted during termination
	// +optional
	DeleteOnTermination *bool `json:"deleteOnTermination,omitempty" tf:"delete_on_termination"`
	// PVM Instance (ID or Name) to base volume affinity policy against;
	// +optional
	PiAffinityInstance *string `json:"piAffinityInstance,omitempty" tf:"pi_affinity_instance"`
	// Affinity policy for data volume being created
	// +optional
	PiAffinityPolicy *string `json:"piAffinityPolicy,omitempty" tf:"pi_affinity_policy"`
	// Volume (ID or Name) to base volume affinity policy against;
	// +optional
	PiAffinityVolume *string `json:"piAffinityVolume,omitempty" tf:"pi_affinity_volume"`
	// Cloud Instance ID - This is the service_instance_id.
	PiCloudInstanceID *string `json:"piCloudInstanceID" tf:"pi_cloud_instance_id"`
	// Volume Name to create
	PiVolumeName *string `json:"piVolumeName" tf:"pi_volume_name"`
	// Flag to indicate if the volume can be shared across multiple instances?
	// +optional
	PiVolumeShareable *bool `json:"piVolumeShareable,omitempty" tf:"pi_volume_shareable"`
	// Size of the volume in GB
	PiVolumeSize *float64 `json:"piVolumeSize" tf:"pi_volume_size"`
	// Volume type
	// +optional
	PiVolumeType *string `json:"piVolumeType,omitempty" tf:"pi_volume_type"`
	// Volume ID
	// +optional
	VolumeID *string `json:"volumeID,omitempty" tf:"volume_id"`
	// Volume status
	// +optional
	VolumeStatus *string `json:"volumeStatus,omitempty" tf:"volume_status"`
	// WWN Of the volume
	// +optional
	Wwn *string `json:"wwn,omitempty" tf:"wwn"`
}

type VolumeStatus struct {
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

// VolumeList is a list of Volumes
type VolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Volume CRD objects
	Items []Volume `json:"items,omitempty"`
}
