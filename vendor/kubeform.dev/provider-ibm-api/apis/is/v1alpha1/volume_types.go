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

type VolumeSpecStatusReasons struct {
	// A snake case string succinctly identifying the status reason
	// +optional
	Code *string `json:"code,omitempty" tf:"code"`
	// An explanation of the status reason
	// +optional
	Message *string `json:"message,omitempty" tf:"message"`
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

	// Vloume capacity value
	// +optional
	Capacity *int64 `json:"capacity,omitempty" tf:"capacity"`
	// CRN value for the volume instance
	// +optional
	Crn *string `json:"crn,omitempty" tf:"crn"`
	// Deletes all snapshots created from this volume
	// +optional
	DeleteAllSnapshots *bool `json:"deleteAllSnapshots,omitempty" tf:"delete_all_snapshots"`
	// Volume encryption key info
	// +optional
	EncryptionKey *string `json:"encryptionKey,omitempty" tf:"encryption_key"`
	// IOPS value for the Volume
	// +optional
	Iops *int64 `json:"iops,omitempty" tf:"iops"`
	// Volume name
	Name *string `json:"name" tf:"name"`
	// Volume profile name
	Profile *string `json:"profile" tf:"profile"`
	// The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance
	// +optional
	ResourceControllerURL *string `json:"resourceControllerURL,omitempty" tf:"resource_controller_url"`
	// The crn of the resource
	// +optional
	ResourceCrn *string `json:"resourceCrn,omitempty" tf:"resource_crn"`
	// Resource group name
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
	// Identifier of the snapshot from which this volume was cloned
	// +optional
	SourceSnapshot *string `json:"sourceSnapshot,omitempty" tf:"source_snapshot"`
	// Volume status
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	StatusReasons []VolumeSpecStatusReasons `json:"statusReasons,omitempty" tf:"status_reasons"`
	// Tags for the volume instance
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags"`
	// Zone name
	Zone *string `json:"zone" tf:"zone"`
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
