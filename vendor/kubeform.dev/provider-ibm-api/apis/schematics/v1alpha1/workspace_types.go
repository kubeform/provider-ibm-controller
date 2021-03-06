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

type Workspace struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkspaceSpec   `json:"spec,omitempty"`
	Status            WorkspaceStatus `json:"status,omitempty"`
}

type WorkspaceSpecCatalogRef struct {
	// Dry run.
	// +optional
	DryRun *bool `json:"dryRun,omitempty" tf:"dry_run"`
	// The URL to the icon of the software template in the IBM Cloud catalog.
	// +optional
	ItemIconURL *string `json:"itemIconURL,omitempty" tf:"item_icon_url"`
	// The ID of the software template that you chose to install from the IBM Cloud catalog. This software is provisioned with Schematics.
	// +optional
	ItemID *string `json:"itemID,omitempty" tf:"item_id"`
	// The name of the software that you chose to install from the IBM Cloud catalog.
	// +optional
	ItemName *string `json:"itemName,omitempty" tf:"item_name"`
	// The URL to the readme file of the software template in the IBM Cloud catalog.
	// +optional
	ItemReadmeURL *string `json:"itemReadmeURL,omitempty" tf:"item_readme_url"`
	// The URL to the software template in the IBM Cloud catalog.
	// +optional
	ItemURL *string `json:"itemURL,omitempty" tf:"item_url"`
	// The URL to the dashboard to access your software.
	// +optional
	LaunchURL *string `json:"launchURL,omitempty" tf:"launch_url"`
	// The version of the software template that you chose to install from the IBM Cloud catalog.
	// +optional
	OfferingVersion *string `json:"offeringVersion,omitempty" tf:"offering_version"`
}

type WorkspaceSpecRuntimeData struct {
	// The command that was used to apply the Terraform template or IBM Cloud catalog software template.
	// +optional
	EngineCmd *string `json:"engineCmd,omitempty" tf:"engine_cmd"`
	// The provisioning engine that was used to apply the Terraform template or IBM Cloud catalog software template.
	// +optional
	EngineName *string `json:"engineName,omitempty" tf:"engine_name"`
	// The version of the provisioning engine that was used.
	// +optional
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version"`
	// The ID that was assigned to your Terraform template or IBM Cloud catalog software template.
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// The URL to access the logs that were created during the creation, update, or deletion of your IBM Cloud resources.
	// +optional
	LogStoreURL *string `json:"logStoreURL,omitempty" tf:"log_store_url"`
	// List of Output values.
	// +optional
	// List of resources.
	// +optional
	// The URL where the Terraform statefile (`terraform.tfstate`) is stored. You can use the statefile to find an overview of IBM Cloud resources that were created by Schematics. Schematics uses the statefile as an inventory list to determine future create, update, or deletion actions.
	// +optional
	StateStoreURL *string `json:"stateStoreURL,omitempty" tf:"state_store_url"`
}

type WorkspaceSpecSharedData struct {
	// Cluster created on.
	// +optional
	ClusterCreatedOn *string `json:"clusterCreatedOn,omitempty" tf:"cluster_created_on"`
	// The ID of the cluster where you want to provision the resources of all IBM Cloud catalog templates that are included in the catalog offering.
	// +optional
	ClusterID *string `json:"clusterID,omitempty" tf:"cluster_id"`
	// Cluster name.
	// +optional
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name"`
	// Cluster type.
	// +optional
	ClusterType *string `json:"clusterType,omitempty" tf:"cluster_type"`
	// The entitlement key that you want to use to install IBM Cloud entitled software.
	// +optional
	// The Kubernetes namespace or OpenShift project where the resources of all IBM Cloud catalog templates that are included in the catalog offering are deployed into.
	// +optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace"`
	// The IBM Cloud region that you want to use for the resources of all IBM Cloud catalog templates that are included in the catalog offering.
	// +optional
	Region *string `json:"region,omitempty" tf:"region"`
	// The ID of the resource group that you want to use for the resources of all IBM Cloud catalog templates that are included in the catalog offering.
	// +optional
	ResourceGroupID *string `json:"resourceGroupID,omitempty" tf:"resource_group_id"`
	// Cluster worker count.
	// +optional
	WorkerCount *int64 `json:"workerCount,omitempty" tf:"worker_count"`
	// Cluster worker type.
	// +optional
	WorkerMachineType *string `json:"workerMachineType,omitempty" tf:"worker_machine_type"`
}

type WorkspaceSpecTemplateInputs struct {
	// The description of your input variable.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// The name of the variable.
	Name *string `json:"name" tf:"name"`
	// If set to `true`, the value of your input variable is protected and not returned in your API response.
	// +optional
	Secure *bool `json:"secure,omitempty" tf:"secure"`
	// `Terraform v0.11` supports `string`, `list`, `map` data type. For more information, about the syntax, see [Configuring input variables](https://www.terraform.io/docs/configuration-0-11/variables.html). <br> `Terraform v0.12` additionally, supports `bool`, `number` and complex data types such as `list(type)`, `map(type)`, `object({attribute name=type,..})`, `set(type)`, `tuple([type])`. For more information, about the syntax to use the complex data type, see [Configuring variables](https://www.terraform.io/docs/configuration/variables.html#type-constraints).
	Type *string `json:"type" tf:"type"`
	// Variable uses default value; and is not over-ridden.
	// +optional
	UseDefault *bool `json:"useDefault,omitempty" tf:"use_default"`
	// Enter the value as a string for the primitive types such as `bool`, `number`, `string`, and `HCL` format for the complex variables, as you provide in a `.tfvars` file. **You need to enter escaped string of `HCL` format for the complex variable value**. For more information, about how to declare variables in a terraform configuration file and provide value to schematics, see [Providing values for the declared variables](/docs/schematics?topic=schematics-create-tf-config#declare-variable).
	Value *string `json:"value" tf:"value"`
}

type WorkspaceSpec struct {
	State *WorkspaceSpecResource `json:"state,omitempty" tf:"-"`

	Resource WorkspaceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type WorkspaceSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// List of applied shared dataset id.
	// +optional
	AppliedShareddataIDS []string `json:"appliedShareddataIDS,omitempty" tf:"applied_shareddata_ids"`
	// Information about the software template that you chose from the IBM Cloud catalog. This information is returned for IBM Cloud catalog offerings only.
	// +optional
	CatalogRef *WorkspaceSpecCatalogRef `json:"catalogRef,omitempty" tf:"catalog_ref"`
	// The timestamp when the workspace was created.
	// +optional
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at"`
	// The user ID that created the workspace.
	// +optional
	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by"`
	// Workspace CRN.
	// +optional
	Crn *string `json:"crn,omitempty" tf:"crn"`
	// The description of the workspace.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// If set to true, the workspace is frozen and changes to the workspace are disabled.
	// +optional
	Frozen *bool `json:"frozen,omitempty" tf:"frozen"`
	// The timestamp when the workspace was frozen.
	// +optional
	FrozenAt *string `json:"frozenAt,omitempty" tf:"frozen_at"`
	// The user ID that froze the workspace.
	// +optional
	FrozenBy *string `json:"frozenBy,omitempty" tf:"frozen_by"`
	// The timestamp when the last health check was performed by Schematics.
	// +optional
	LastHealthCheckAt *string `json:"lastHealthCheckAt,omitempty" tf:"last_health_check_at"`
	// The location where you want to create your Schematics workspace and run Schematics actions. The location that you enter must match the API endpoint that you use. For example, if you use the Frankfurt API endpoint, you must specify `eu-de` as your location. If you use an API endpoint for a geography and you do not specify a location, Schematics determines the location based on availability.
	// +optional
	Location *string `json:"location,omitempty" tf:"location"`
	// If set to true, the workspace is locked and disabled for changes.
	// +optional
	Locked *bool `json:"locked,omitempty" tf:"locked"`
	// The user ID that initiated a resource-related action, such as applying or destroying resources, that locked the workspace.
	// +optional
	LockedBy *string `json:"lockedBy,omitempty" tf:"locked_by"`
	// The timestamp when the workspace was locked.
	// +optional
	LockedTime *string `json:"lockedTime,omitempty" tf:"locked_time"`
	// The name of your workspace. The name can be up to 128 characters long and can include alphanumeric characters, spaces, dashes, and underscores. When you create a workspace for your own Terraform template, consider including the microservice component that you set up with your Terraform template and the IBM Cloud environment where you want to deploy your resources in your name.
	Name *string `json:"name" tf:"name"`
	// The ID of the resource group where you want to provision the workspace.
	// +optional
	ResourceGroup *string `json:"resourceGroup,omitempty" tf:"resource_group"`
	// Information about the provisioning engine, state file, and runtime logs.
	// +optional
	RuntimeData []WorkspaceSpecRuntimeData `json:"runtimeData,omitempty" tf:"runtime_data"`
	// Information that is shared across templates in IBM Cloud catalog offerings. This information is not provided when you create a workspace from your own Terraform template.
	// +optional
	SharedData *WorkspaceSpecSharedData `json:"sharedData,omitempty" tf:"shared_data"`
	// The status of the workspace.  **Active**: After you successfully ran your infrastructure code by applying your Terraform execution plan, the state of your workspace changes to `Active`.  **Connecting**: Schematics tries to connect to the template in your source repo. If successfully connected, the template is downloaded and metadata, such as input parameters, is extracted. After the template is downloaded, the state of the workspace changes to `Scanning`.  **Draft**: The workspace is created without a reference to a GitHub or GitLab repository.  **Failed**: If errors occur during the execution of your infrastructure code in IBM Cloud Schematics, your workspace status is set to `Failed`.  **Inactive**: The Terraform template was scanned successfully and the workspace creation is complete. You can now start running Schematics plan and apply actions to provision the IBM Cloud resources that you specified in your template. If you have an `Active` workspace and decide to remove all your resources, your workspace is set to `Inactive` after all your resources are removed.  **In progress**: When you instruct IBM Cloud Schematics to run your infrastructure code by applying your Terraform execution plan, the status of our workspace changes to `In progress`.  **Scanning**: The download of the Terraform template is complete and vulnerability scanning started. If the scan is successful, the workspace state changes to `Inactive`. If errors in your template are found, the state changes to `Template Error`.  **Stopped**: The Schematics plan, apply, or destroy action was cancelled manually.  **Template Error**: The Schematics template contains errors and cannot be processed.
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// The success or error code that was returned for the last plan, apply, or destroy action that ran against your workspace.
	// +optional
	StatusCode *string `json:"statusCode,omitempty" tf:"status_code"`
	// The success or error message that was returned for the last plan, apply, or destroy action that ran against your workspace.
	// +optional
	StatusMsg *string `json:"statusMsg,omitempty" tf:"status_msg"`
	// A list of tags that are associated with the workspace.
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags"`
	// A list of environment variables that you want to apply during the execution of a bash script or Terraform action. This field must be provided as a list of key-value pairs, for example, **TF_LOG=debug**. Each entry will be a map with one entry where `key is the environment variable name and value is value`. You can define environment variables for IBM Cloud catalog offerings that are provisioned by using a bash script.
	// +optional
	// The branch in GitHub where your Terraform template is stored.
	// +optional
	TemplateGitBranch *string `json:"templateGitBranch,omitempty" tf:"template_git_branch"`
	// The subfolder in your GitHub or GitLab repository where your Terraform template is stored.
	// +optional
	TemplateGitFolder *string `json:"templateGitFolder,omitempty" tf:"template_git_folder"`
	// Has uploaded git repo tar
	// +optional
	TemplateGitHasUploadedgitrepotar *bool `json:"templateGitHasUploadedgitrepotar,omitempty" tf:"template_git_has_uploadedgitrepotar"`
	// The release tag in GitHub of your Terraform template.
	// +optional
	TemplateGitRelease *string `json:"templateGitRelease,omitempty" tf:"template_git_release"`
	// Repo SHA value.
	// +optional
	TemplateGitRepoShaValue *string `json:"templateGitRepoShaValue,omitempty" tf:"template_git_repo_sha_value"`
	// The URL to the repository where the IBM Cloud catalog software template is stored.
	// +optional
	TemplateGitRepoURL *string `json:"templateGitRepoURL,omitempty" tf:"template_git_repo_url"`
	// The URL to the GitHub or GitLab repository where your Terraform and public bit bucket template is stored. For more information of the environment variable syntax, see [Create workspace new](/docs/schematics?topic=schematics-schematics-cli-reference#schematics-workspace-new).
	// +optional
	TemplateGitURL *string `json:"templateGitURL,omitempty" tf:"template_git_url"`
	// The content of an existing Terraform statefile that you want to import in to your workspace. To get the content of a Terraform statefile for a specific Terraform template in an existing workspace, run `ibmcloud terraform state pull --id <workspace_id> --template <template_id>`.
	// +optional
	TemplateInitStateFile *string `json:"templateInitStateFile,omitempty" tf:"template_init_state_file"`
	// VariablesRequest -.
	// +optional
	TemplateInputs []WorkspaceSpecTemplateInputs `json:"templateInputs,omitempty" tf:"template_inputs"`
	// Workspace template ref.
	// +optional
	TemplateRef *string `json:"templateRef,omitempty" tf:"template_ref"`
	// The Terraform version that you want to use to run your Terraform code. Enter `terraform_v0.12` to use Terraform version 0.12, and `terraform_v0.11` to use Terraform version 0.11. Make sure that your Terraform config files are compatible with the Terraform version that you select.
	TemplateType *string `json:"templateType" tf:"template_type"`
	// Uninstall script name.
	// +optional
	TemplateUninstallScriptName *string `json:"templateUninstallScriptName,omitempty" tf:"template_uninstall_script_name"`
	// A list of variable values that you want to apply during the Helm chart installation. The list must be provided in JSON format, such as `"autoscaling:  enabled: true  minReplicas: 2"`. The values that you define here override the default Helm chart values. This field is supported only for IBM Cloud catalog offerings that are provisioned by using the Terraform Helm provider.
	// +optional
	TemplateValues *string `json:"templateValues,omitempty" tf:"template_values"`
	// List of values metadata.
	// +optional
	// The timestamp when the workspace was last updated.
	// +optional
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at"`
	// The user ID that updated the workspace.
	// +optional
	UpdatedBy *string `json:"updatedBy,omitempty" tf:"updated_by"`
	// The personal access token to authenticate with your private GitHub or GitLab repository and access your Terraform template.
	// +optional
	XGithubToken *string `json:"xGithubToken,omitempty" tf:"x_github_token"`
}

type WorkspaceStatus struct {
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

// WorkspaceList is a list of Workspaces
type WorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Workspace CRD objects
	Items []Workspace `json:"items,omitempty"`
}
