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

type Capture struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CaptureSpec   `json:"spec,omitempty"`
	Status            CaptureStatus `json:"status,omitempty"`
}

type CaptureSpec struct {
	State *CaptureSpecResource `json:"state,omitempty" tf:"-"`

	Resource CaptureSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type CaptureSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of Cloud Storage Access Key
	// +optional
	PiCaptureCloudStorageAccessKey *string `json:"piCaptureCloudStorageAccessKey,omitempty" tf:"pi_capture_cloud_storage_access_key"`
	// List of Regions to use
	// +optional
	PiCaptureCloudStorageRegion *string `json:"piCaptureCloudStorageRegion,omitempty" tf:"pi_capture_cloud_storage_region"`
	// Name of the Cloud Storage Secret Key
	// +optional
	PiCaptureCloudStorageSecretKey *string `json:"piCaptureCloudStorageSecretKey,omitempty" tf:"pi_capture_cloud_storage_secret_key"`
	// Name of destination to store the image capture to
	PiCaptureDestination *string `json:"piCaptureDestination" tf:"pi_capture_destination"`
	// Name of the capture to create. Note : this must be unique
	PiCaptureName *string `json:"piCaptureName" tf:"pi_capture_name"`
	// Name of the Image Path
	// +optional
	PiCaptureStorageImagePath *string `json:"piCaptureStorageImagePath,omitempty" tf:"pi_capture_storage_image_path"`
	// List of volume names that need to be passed in the input
	// +optional
	PiCaptureVolumeIDS *string `json:"piCaptureVolumeIDS,omitempty" tf:"pi_capture_volume_ids"`
	//  Cloud Instance ID - This is the service_instance_id.
	PiCloudInstanceID *string `json:"piCloudInstanceID" tf:"pi_cloud_instance_id"`
	// Instance Name of the Power VM
	PiInstanceName *string `json:"piInstanceName" tf:"pi_instance_name"`
}

type CaptureStatus struct {
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

// CaptureList is a list of Captures
type CaptureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Capture CRD objects
	Items []Capture `json:"items,omitempty"`
}
