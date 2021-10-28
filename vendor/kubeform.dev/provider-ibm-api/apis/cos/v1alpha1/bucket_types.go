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

type Bucket struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketSpec   `json:"spec,omitempty"`
	Status            BucketStatus `json:"status,omitempty"`
}

type BucketSpecActivityTracking struct {
	// The instance of Activity Tracker that will receive object event data
	ActivityTrackerCrn *string `json:"activityTrackerCrn" tf:"activity_tracker_crn"`
	// If set to true, all object read events will be sent to Activity Tracker.
	// +optional
	ReadDataEvents *bool `json:"readDataEvents,omitempty" tf:"read_data_events"`
	// If set to true, all object write events will be sent to Activity Tracker.
	// +optional
	WriteDataEvents *bool `json:"writeDataEvents,omitempty" tf:"write_data_events"`
}

type BucketSpecArchiveRule struct {
	// Specifies the number of days when the specific rule action takes effect.
	Days *int64 `json:"days" tf:"days"`
	// Enable or disable an archive rule for a bucket
	Enable *bool `json:"enable" tf:"enable"`
	// Unique identifier for the rule.Archive rules allow you to set a specific time frame after which objects transition to the archive. Set Rule ID for cos bucket
	// +optional
	RuleID *string `json:"ruleID,omitempty" tf:"rule_id"`
	// Specifies the storage class/archive type to which you want the object to transition. It can be Glacier or Accelerated
	Type *string `json:"type" tf:"type"`
}

type BucketSpecExpireRule struct {
	// Specifies the number of days when the specific rule action takes effect.
	Days *int64 `json:"days" tf:"days"`
	// Enable or disable an expire rule for a bucket
	Enable *bool `json:"enable" tf:"enable"`
	// The rule applies to any objects with keys that match this prefix
	// +optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix"`
	// Unique identifier for the rule.Expire rules allow you to set a specific time frame after which objects are deleted. Set Rule ID for cos bucket
	// +optional
	RuleID *string `json:"ruleID,omitempty" tf:"rule_id"`
}

type BucketSpecMetricsMonitoring struct {
	// Instance of IBM Cloud Monitoring that will receive the bucket metrics.
	MetricsMonitoringCrn *string `json:"metricsMonitoringCrn" tf:"metrics_monitoring_crn"`
	// Request metrics will be sent to the monitoring service.
	// +optional
	RequestMetricsEnabled *bool `json:"requestMetricsEnabled,omitempty" tf:"request_metrics_enabled"`
	// Usage metrics will be sent to the monitoring service.
	// +optional
	UsageMetricsEnabled *bool `json:"usageMetricsEnabled,omitempty" tf:"usage_metrics_enabled"`
}

type BucketSpecObjectVersioning struct {
	// Enable or suspend the versioning for objects in the bucket
	// +optional
	Enable *bool `json:"enable,omitempty" tf:"enable"`
}

type BucketSpecRetentionRule struct {
	// If an object is stored in the bucket without specifying a custom retention period.
	Default *int64 `json:"default" tf:"default"`
	// Maximum duration of time an object can be kept unmodified in the bucket.
	Maximum *int64 `json:"maximum" tf:"maximum"`
	// Minimum duration of time an object must be kept unmodified in the bucket
	Minimum *int64 `json:"minimum" tf:"minimum"`
	// Enable or disable the permanent retention policy on the bucket
	// +optional
	Permanent *bool `json:"permanent,omitempty" tf:"permanent"`
}

type BucketSpec struct {
	State *BucketSpecResource `json:"state,omitempty" tf:"-"`

	Resource BucketSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type BucketSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Enables sending log data to Activity Tracker and LogDNA to provide visibility into object read and write events
	// +optional
	ActivityTracking *BucketSpecActivityTracking `json:"activityTracking,omitempty" tf:"activity_tracking"`
	// List of IPv4 or IPv6 addresses
	// +optional
	AllowedIP []string `json:"allowedIP,omitempty" tf:"allowed_ip"`
	// Enable configuration archive_rule (glacier/accelerated) to COS Bucket after a defined period of time
	// +optional
	ArchiveRule *BucketSpecArchiveRule `json:"archiveRule,omitempty" tf:"archive_rule"`
	// COS Bucket name
	BucketName *string `json:"bucketName" tf:"bucket_name"`
	// CRN of resource instance
	// +optional
	Crn *string `json:"crn,omitempty" tf:"crn"`
	// Cros region location info
	// +optional
	CrossRegionLocation *string `json:"crossRegionLocation,omitempty" tf:"cross_region_location"`
	// public or private
	// +optional
	EndpointType *string `json:"endpointType,omitempty" tf:"endpoint_type"`
	// Enable configuration expire_rule to COS Bucket after a defined period of time
	// +optional
	// +kubebuilder:validation:MaxItems=1000
	ExpireRule []BucketSpecExpireRule `json:"expireRule,omitempty" tf:"expire_rule"`
	// COS buckets need to be empty before they can be deleted. force_delete option empty the bucket and delete it.
	// +optional
	ForceDelete *bool `json:"forceDelete,omitempty" tf:"force_delete"`
	// sets a maximum amount of storage (in bytes) available for a bucket
	// +optional
	HardQuota *int64 `json:"hardQuota,omitempty" tf:"hard_quota"`
	// CRN of the key you want to use data at rest encryption
	// +optional
	KeyProtect *string `json:"keyProtect,omitempty" tf:"key_protect"`
	// Enables sending metrics to IBM Cloud Monitoring.
	// +optional
	MetricsMonitoring *BucketSpecMetricsMonitoring `json:"metricsMonitoring,omitempty" tf:"metrics_monitoring"`
	// Protect objects from accidental deletion or overwrites. Versioning allows you to keep multiple versions of an object protecting from unintentional data loss.
	// +optional
	ObjectVersioning *BucketSpecObjectVersioning `json:"objectVersioning,omitempty" tf:"object_versioning"`
	// Region Location info.
	// +optional
	RegionLocation *string `json:"regionLocation,omitempty" tf:"region_location"`
	// resource instance ID
	ResourceInstanceID *string `json:"resourceInstanceID" tf:"resource_instance_id"`
	// A retention policy is enabled at the IBM Cloud Object Storage bucket level. Minimum, maximum and default retention period are defined by this policy and apply to all objects in the bucket.
	// +optional
	RetentionRule *BucketSpecRetentionRule `json:"retentionRule,omitempty" tf:"retention_rule"`
	// Private endpoint for the COS bucket
	// +optional
	S3EndpointPrivate *string `json:"s3EndpointPrivate,omitempty" tf:"s3_endpoint_private"`
	// Public endpoint for the COS bucket
	// +optional
	S3EndpointPublic *string `json:"s3EndpointPublic,omitempty" tf:"s3_endpoint_public"`
	// single site location info
	// +optional
	SingleSiteLocation *string `json:"singleSiteLocation,omitempty" tf:"single_site_location"`
	// Storage class info
	StorageClass *string `json:"storageClass" tf:"storage_class"`
}

type BucketStatus struct {
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

// BucketList is a list of Buckets
type BucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Bucket CRD objects
	Items []Bucket `json:"items,omitempty"`
}
