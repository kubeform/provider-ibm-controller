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

type Image struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ImageSpec   `json:"spec,omitempty"`
	Status            ImageStatus `json:"status,omitempty"`
}

type ImageSpec struct {
	State *ImageSpecResource `json:"state,omitempty" tf:"-"`

	Resource ImageSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ImageSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The SHA256 checksum of this image
	// +optional
	Checksum *string `json:"checksum,omitempty" tf:"checksum"`
	// A base64-encoded, encrypted representation of the key that was used to encrypt the data for this image
	// +optional
	EncryptedDataKey *string `json:"encryptedDataKey,omitempty" tf:"encrypted_data_key"`
	// The type of encryption used on the image
	// +optional
	Encryption *string `json:"encryption,omitempty" tf:"encryption"`
	// The CRN of the Key Protect Root Key or Hyper Protect Crypto Service Root Key for this resource
	// +optional
	EncryptionKey *string `json:"encryptionKey,omitempty" tf:"encryption_key"`
	// Details for the stored image file
	// +optional
	File *int64 `json:"file,omitempty" tf:"file"`
	// Image Href value
	// +optional
	Href *string `json:"href,omitempty" tf:"href"`
	// Image name
	Name *string `json:"name" tf:"name"`
	// Image Operating system
	// +optional
	OperatingSystem *string `json:"operatingSystem,omitempty" tf:"operating_system"`
	// The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance
	// +optional
	ResourceControllerURL *string `json:"resourceControllerURL,omitempty" tf:"resource_controller_url"`
	// The crn of the resource
	// +optional
	ResourceCrn *string `json:"resourceCrn,omitempty" tf:"resource_crn"`
	// The resource group for this image
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
	// The minimum size (in gigabytes) of a volume onto which this image may be provisioned
	// +optional
	Size *int64 `json:"size,omitempty" tf:"size"`
	// Image volume id
	// +optional
	SourceVolume *string `json:"sourceVolume,omitempty" tf:"source_volume"`
	// The status of this image
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// Tags for the image
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags"`
	// Whether the image is publicly visible or private to the account
	// +optional
	Visibility *string `json:"visibility,omitempty" tf:"visibility"`
}

type ImageStatus struct {
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

// ImageList is a list of Images
type ImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Image CRD objects
	Items []Image `json:"items,omitempty"`
}
