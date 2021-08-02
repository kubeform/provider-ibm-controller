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

type Certificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CertificateSpec   `json:"spec,omitempty"`
	Status            CertificateStatus `json:"status,omitempty"`
}

type CertificateSpecAdministrativeContactAdminAddress struct {
	// +optional
	AdminAddressLine1 *string `json:"adminAddressLine1,omitempty" tf:"admin_address_line1"`
	// +optional
	AdminAddressLine2 *string `json:"adminAddressLine2,omitempty" tf:"admin_address_line2"`
	// +optional
	AdminCity *string `json:"adminCity,omitempty" tf:"admin_city"`
	// +optional
	AdminCountryCode *string `json:"adminCountryCode,omitempty" tf:"admin_country_code"`
	// +optional
	AdminPostalCode *string `json:"adminPostalCode,omitempty" tf:"admin_postal_code"`
	// +optional
	AdminState *string `json:"adminState,omitempty" tf:"admin_state"`
}

type CertificateSpecAdministrativeContact struct {
	// +optional
	AdminAddress []CertificateSpecAdministrativeContactAdminAddress `json:"adminAddress,omitempty" tf:"admin_address"`
	// +optional
	AdminEmailAddress *string `json:"adminEmailAddress,omitempty" tf:"admin_email_address"`
	// +optional
	AdminFaxNumber *string `json:"adminFaxNumber,omitempty" tf:"admin_fax_number"`
	// +optional
	AdminFirstName *string `json:"adminFirstName,omitempty" tf:"admin_first_name"`
	// +optional
	AdminLastName *string `json:"adminLastName,omitempty" tf:"admin_last_name"`
	// +optional
	AdminOrganizationName *string `json:"adminOrganizationName,omitempty" tf:"admin_organization_name"`
	// +optional
	AdminPhoneNumber *string `json:"adminPhoneNumber,omitempty" tf:"admin_phone_number"`
	// +optional
	AdminTitle *string `json:"adminTitle,omitempty" tf:"admin_title"`
}

type CertificateSpecBillingContactBillingAddress struct {
	// +optional
	BillingAddressLine1 *string `json:"billingAddressLine1,omitempty" tf:"billing_address_line1"`
	// +optional
	BillingAddressLine2 *string `json:"billingAddressLine2,omitempty" tf:"billing_address_line2"`
	// +optional
	BillingCity *string `json:"billingCity,omitempty" tf:"billing_city"`
	// +optional
	BillingCountryCode *string `json:"billingCountryCode,omitempty" tf:"billing_country_code"`
	// +optional
	BillingPostalCode *string `json:"billingPostalCode,omitempty" tf:"billing_postal_code"`
	// +optional
	BillingState *string `json:"billingState,omitempty" tf:"billing_state"`
}

type CertificateSpecBillingContact struct {
	// +optional
	BillingAddress []CertificateSpecBillingContactBillingAddress `json:"billingAddress,omitempty" tf:"billing_address"`
	// +optional
	BillingEmailAddress *string `json:"billingEmailAddress,omitempty" tf:"billing_email_address"`
	// +optional
	BillingFaxNumber *string `json:"billingFaxNumber,omitempty" tf:"billing_fax_number"`
	// +optional
	BillingFirstName *string `json:"billingFirstName,omitempty" tf:"billing_first_name"`
	// +optional
	BillingLastName *string `json:"billingLastName,omitempty" tf:"billing_last_name"`
	// +optional
	BillingOrganizationName *string `json:"billingOrganizationName,omitempty" tf:"billing_organization_name"`
	// +optional
	BillingPhoneNumber *string `json:"billingPhoneNumber,omitempty" tf:"billing_phone_number"`
	// +optional
	BillingTitle *string `json:"billingTitle,omitempty" tf:"billing_title"`
}

type CertificateSpecOrganizationInformationOrgAddress struct {
	OrgAddressLine1 *string `json:"orgAddressLine1" tf:"org_address_line1"`
	// +optional
	OrgAddressLine2 *string `json:"orgAddressLine2,omitempty" tf:"org_address_line2"`
	OrgCity         *string `json:"orgCity" tf:"org_city"`
	OrgCountryCode  *string `json:"orgCountryCode" tf:"org_country_code"`
	OrgPostalCode   *string `json:"orgPostalCode" tf:"org_postal_code"`
	OrgState        *string `json:"orgState" tf:"org_state"`
}

type CertificateSpecOrganizationInformation struct {
	// Organization address
	OrgAddress []CertificateSpecOrganizationInformationOrgAddress `json:"orgAddress" tf:"org_address"`
	// +optional
	OrgFaxNumber *string `json:"orgFaxNumber,omitempty" tf:"org_fax_number"`
	// Organization name
	OrgOrganizationName *string `json:"orgOrganizationName" tf:"org_organization_name"`
	// Organization phone number
	OrgPhoneNumber *string `json:"orgPhoneNumber" tf:"org_phone_number"`
}

type CertificateSpecTechnicalContactTechAddress struct {
	// +optional
	TechAddressLine1 *string `json:"techAddressLine1,omitempty" tf:"tech_address_line1"`
	// +optional
	TechAddressLine2 *string `json:"techAddressLine2,omitempty" tf:"tech_address_line2"`
	// +optional
	TechCity *string `json:"techCity,omitempty" tf:"tech_city"`
	// +optional
	TechCountryCode *string `json:"techCountryCode,omitempty" tf:"tech_country_code"`
	// +optional
	TechPostalCode *string `json:"techPostalCode,omitempty" tf:"tech_postal_code"`
	// +optional
	TechState *string `json:"techState,omitempty" tf:"tech_state"`
}

type CertificateSpecTechnicalContact struct {
	// +optional
	TechAddress      []CertificateSpecTechnicalContactTechAddress `json:"techAddress,omitempty" tf:"tech_address"`
	TechEmailAddress *string                                      `json:"techEmailAddress" tf:"tech_email_address"`
	// +optional
	TechFaxNumber        *string `json:"techFaxNumber,omitempty" tf:"tech_fax_number"`
	TechFirstName        *string `json:"techFirstName" tf:"tech_first_name"`
	TechLastName         *string `json:"techLastName" tf:"tech_last_name"`
	TechOrganizationName *string `json:"techOrganizationName" tf:"tech_organization_name"`
	TechPhoneNumber      *string `json:"techPhoneNumber" tf:"tech_phone_number"`
	TechTitle            *string `json:"techTitle" tf:"tech_title"`
}

type CertificateSpec struct {
	State *CertificateSpecResource `json:"state,omitempty" tf:"-"`

	Resource CertificateSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type CertificateSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// administrative address same as organization flag
	// +optional
	AdministrativeAddressSameAsOrganizationFlag *bool `json:"administrativeAddressSameAsOrganizationFlag,omitempty" tf:"administrative_address_same_as_organization_flag"`
	// +optional
	AdministrativeContact *CertificateSpecAdministrativeContact `json:"administrativeContact,omitempty" tf:"administrative_contact"`
	// Administrative contact same as technical flag
	// +optional
	AdministrativeContactSameAsTechnicalFlag *bool `json:"administrativeContactSameAsTechnicalFlag,omitempty" tf:"administrative_contact_same_as_technical_flag"`
	// billing address same as organization flag
	// +optional
	BillingAddressSameAsOrganizationFlag *bool `json:"billingAddressSameAsOrganizationFlag,omitempty" tf:"billing_address_same_as_organization_flag"`
	// +optional
	BillingContact *CertificateSpecBillingContact `json:"billingContact,omitempty" tf:"billing_contact"`
	// billing contact
	// +optional
	BillingContactSameAsTechnicalFlag *bool `json:"billingContactSameAsTechnicalFlag,omitempty" tf:"billing_contact_same_as_technical_flag"`
	// certificate signing request info
	CertificateSigningRequest *string `json:"certificateSigningRequest" tf:"certificate_signing_request"`
	// Email address of the approver
	OrderApproverEmailAddress *string `json:"orderApproverEmailAddress" tf:"order_approver_email_address"`
	// Organization information
	OrganizationInformation *CertificateSpecOrganizationInformation `json:"organizationInformation" tf:"organization_information"`
	// Renewal flag
	// +optional
	RenewalFlag *bool `json:"renewalFlag,omitempty" tf:"renewal_flag"`
	// Server count
	ServerCount *int64 `json:"serverCount" tf:"server_count"`
	// server type
	ServerType *string `json:"serverType" tf:"server_type"`
	// ssl type
	SslType *string `json:"sslType" tf:"ssl_type"`
	// Technical contact info
	TechnicalContact *CertificateSpecTechnicalContact `json:"technicalContact" tf:"technical_contact"`
	// Technical contact same as org address flag
	// +optional
	TechnicalContactSameAsOrgAddressFlag *bool `json:"technicalContactSameAsOrgAddressFlag,omitempty" tf:"technical_contact_same_as_org_address_flag"`
	// vslidity of the ssl certificate in month
	ValidityMonths *int64 `json:"validityMonths" tf:"validity_months"`
}

type CertificateStatus struct {
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

// CertificateList is a list of Certificates
type CertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Certificate CRD objects
	Items []Certificate `json:"items,omitempty"`
}
