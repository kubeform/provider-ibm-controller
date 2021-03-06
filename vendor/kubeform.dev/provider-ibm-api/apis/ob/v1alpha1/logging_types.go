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

type Logging struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoggingSpec   `json:"spec,omitempty"`
	Status            LoggingStatus `json:"status,omitempty"`
}

type LoggingSpec struct {
	State *LoggingSpecResource `json:"state,omitempty" tf:"-"`

	Resource LoggingSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type LoggingSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Agent key name
	// +optional
	AgentKey *string `json:"agentKey,omitempty" tf:"agent_key"`
	// Agent Namespace
	// +optional
	AgentNamespace *string `json:"agentNamespace,omitempty" tf:"agent_namespace"`
	// Name or ID of the cluster to be used.
	Cluster *string `json:"cluster" tf:"cluster"`
	// CRN
	// +optional
	Crn *string `json:"crn,omitempty" tf:"crn"`
	// Daemon Set Name
	// +optional
	DaemonsetName *string `json:"daemonsetName,omitempty" tf:"daemonset_name"`
	// Discovered agent
	// +optional
	DiscoveredAgent *bool `json:"discoveredAgent,omitempty" tf:"discovered_agent"`
	// ID of the LogDNA service instance to latch
	InstanceID *string `json:"instanceID" tf:"instance_id"`
	// LogDNA instance Name
	// +optional
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name"`
	// LogDNA ingestion key
	// +optional
	LogdnaIngestionKey *string `json:"logdnaIngestionKey,omitempty" tf:"logdna_ingestion_key"`
	// Namespace
	// +optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace"`
	// Add this option to connect to your LogDNA service instance through the private service endpoint
	// +optional
	PrivateEndpoint *bool `json:"privateEndpoint,omitempty" tf:"private_endpoint"`
}

type LoggingStatus struct {
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

// LoggingList is a list of Loggings
type LoggingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Logging CRD objects
	Items []Logging `json:"items,omitempty"`
}
