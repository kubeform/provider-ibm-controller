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

package provider

type IbmSpec struct {
	// The Bluemix API Key
	// +optional
	// Deprecated
	BluemixAPIKey *string `json:"bluemixAPIKey,omitempty" tf:"bluemix_api_key"`
	// The timeout (in seconds) to set for any Bluemix API calls made.
	// +optional
	// Deprecated
	BluemixTimeout *int64 `json:"bluemixTimeout,omitempty" tf:"bluemix_timeout"`
	// The IBM Cloud Function namespace
	// +optional
	// Deprecated
	FunctionNamespace *string `json:"functionNamespace,omitempty" tf:"function_namespace"`
	// Generation of Virtual Private Cloud. Default is 2
	// +optional
	// Deprecated
	Generation *int64 `json:"generation,omitempty" tf:"generation"`
	// The Classic Infrastructure API Key
	// +optional
	IaasClassicAPIKey *string `json:"iaasClassicAPIKey,omitempty" tf:"iaas_classic_api_key"`
	// The Classic Infrastructure Endpoint
	// +optional
	IaasClassicEndpointURL *string `json:"iaasClassicEndpointURL,omitempty" tf:"iaas_classic_endpoint_url"`
	// The timeout (in seconds) to set for any Classic Infrastructure API calls made.
	// +optional
	IaasClassicTimeout *int64 `json:"iaasClassicTimeout,omitempty" tf:"iaas_classic_timeout"`
	// The Classic Infrastructure API user name
	// +optional
	IaasClassicUsername *string `json:"iaasClassicUsername,omitempty" tf:"iaas_classic_username"`
	// IAM Authentication refresh token
	// +optional
	IamRefreshToken *string `json:"iamRefreshToken,omitempty" tf:"iam_refresh_token"`
	// IAM Authentication token
	// +optional
	IamToken *string `json:"iamToken,omitempty" tf:"iam_token"`
	// The IBM Cloud API Key
	// +optional
	IbmcloudAPIKey *string `json:"ibmcloudAPIKey,omitempty" tf:"ibmcloud_api_key"`
	// The timeout (in seconds) to set for any IBM Cloud API calls made.
	// +optional
	IbmcloudTimeout *int64 `json:"ibmcloudTimeout,omitempty" tf:"ibmcloud_timeout"`
	// The retry count to set for API calls.
	// +optional
	MaxRetries *int64 `json:"maxRetries,omitempty" tf:"max_retries"`
	// The IBM cloud Region (for example 'us-south').
	// +optional
	Region *string `json:"region,omitempty" tf:"region"`
	// The Resource group id.
	// +optional
	ResourceGroup *string `json:"resourceGroup,omitempty" tf:"resource_group"`
	// The next generation infrastructure service endpoint url.
	// +optional
	// Deprecated
	RiaasEndpoint *string `json:"riaasEndpoint,omitempty" tf:"riaas_endpoint"`
	// The SoftLayer API Key
	// +optional
	// Deprecated
	SoftlayerAPIKey *string `json:"softlayerAPIKey,omitempty" tf:"softlayer_api_key"`
	// The Softlayer Endpoint
	// +optional
	// Deprecated
	SoftlayerEndpointURL *string `json:"softlayerEndpointURL,omitempty" tf:"softlayer_endpoint_url"`
	// The timeout (in seconds) to set for any SoftLayer API calls made.
	// +optional
	// Deprecated
	SoftlayerTimeout *int64 `json:"softlayerTimeout,omitempty" tf:"softlayer_timeout"`
	// The SoftLayer user name
	// +optional
	// Deprecated
	SoftlayerUsername *string `json:"softlayerUsername,omitempty" tf:"softlayer_username"`
	// Visibility of the provider if it is private or public.
	// +optional
	Visibility *string `json:"visibility,omitempty" tf:"visibility"`
	// The IBM cloud Region zone (for example 'us-south-1') for power resources.
	// +optional
	Zone *string `json:"zone,omitempty" tf:"zone"`
}
