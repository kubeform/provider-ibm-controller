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

type Location struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LocationSpec   `json:"spec,omitempty"`
	Status            LocationStatus `json:"status,omitempty"`
}

type LocationSpecCosConfig struct {
	// +optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket"`
	// +optional
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint"`
	// +optional
	Region *string `json:"region,omitempty" tf:"region"`
}

type LocationSpecCosCredentials struct {
	// The HMAC secret access key ID
	// +optional
	AccessKeyID *string `json:"accessKeyID,omitempty" tf:"access_key_id"`
	// The HMAC secret access key
	// +optional
	SecretAccessKey *string `json:"secretAccessKey,omitempty" tf:"secret_access_key"`
}

type LocationSpec struct {
	State *LocationSpecResource `json:"state,omitempty" tf:"-"`

	Resource LocationSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type LocationSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// COSBucket - IBM Cloud Object Storage bucket configuration details
	// +optional
	CosConfig *LocationSpecCosConfig `json:"cosConfig,omitempty" tf:"cos_config"`
	// COSAuthorization - IBM Cloud Object Storage authorization keys
	// +optional
	CosCredentials *LocationSpecCosCredentials `json:"cosCredentials,omitempty" tf:"cos_credentials"`
	// Created Date
	// +optional
	CreatedOn *string `json:"createdOn,omitempty" tf:"created_on"`
	// Location CRN
	// +optional
	Crn *string `json:"crn,omitempty" tf:"crn"`
	// A description of the new Satellite location
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// The total number of hosts that are attached to the Satellite location.
	// +optional
	HostAttachedCount *int64 `json:"hostAttachedCount,omitempty" tf:"host_attached_count"`
	// The available number of hosts that can be assigned to a cluster resource in the Satellite location.
	// +optional
	HostAvailableCount *int64 `json:"hostAvailableCount,omitempty" tf:"host_available_count"`
	// +optional
	IngressHostname *string `json:"ingressHostname,omitempty" tf:"ingress_hostname"`
	// +optional
	IngressSecret *string `json:"-" sensitive:"true" tf:"ingress_secret"`
	// A unique name for the new Satellite location
	Location *string `json:"location" tf:"location"`
	// The account ID for IBM Log Analysis with LogDNA log forwarding
	// +optional
	LoggingAccountID *string `json:"loggingAccountID,omitempty" tf:"logging_account_id"`
	// The IBM Cloud metro from which the Satellite location is managed
	ManagedFrom *string `json:"managedFrom" tf:"managed_from"`
	// ID of the resource group.
	// +optional
	ResourceGroupID *string `json:"resourceGroupID,omitempty" tf:"resource_group_id"`
	// Name of the resource group
	// +optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name"`
	// List of tags associated with resource instance
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags"`
	// The names of at least three high availability zones to use for the location
	// +optional
	Zones []string `json:"zones,omitempty" tf:"zones"`
}

type LocationStatus struct {
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

// LocationList is a list of Locations
type LocationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Location CRD objects
	Items []Location `json:"items,omitempty"`
}
