// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewServiceBrokerAuthLoginParams creates a new ServiceBrokerAuthLoginParams object
// with the default values initialized.
func NewServiceBrokerAuthLoginParams() *ServiceBrokerAuthLoginParams {
	var (
		accessTypeDefault = string("online")
	)
	return &ServiceBrokerAuthLoginParams{
		AccessType: &accessTypeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewServiceBrokerAuthLoginParamsWithTimeout creates a new ServiceBrokerAuthLoginParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServiceBrokerAuthLoginParamsWithTimeout(timeout time.Duration) *ServiceBrokerAuthLoginParams {
	var (
		accessTypeDefault = string("online")
	)
	return &ServiceBrokerAuthLoginParams{
		AccessType: &accessTypeDefault,

		timeout: timeout,
	}
}

// NewServiceBrokerAuthLoginParamsWithContext creates a new ServiceBrokerAuthLoginParams object
// with the default values initialized, and the ability to set a context for a request
func NewServiceBrokerAuthLoginParamsWithContext(ctx context.Context) *ServiceBrokerAuthLoginParams {
	var (
		accessTypeDefault = string("online")
	)
	return &ServiceBrokerAuthLoginParams{
		AccessType: &accessTypeDefault,

		Context: ctx,
	}
}

// NewServiceBrokerAuthLoginParamsWithHTTPClient creates a new ServiceBrokerAuthLoginParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServiceBrokerAuthLoginParamsWithHTTPClient(client *http.Client) *ServiceBrokerAuthLoginParams {
	var (
		accessTypeDefault = string("online")
	)
	return &ServiceBrokerAuthLoginParams{
		AccessType: &accessTypeDefault,
		HTTPClient: client,
	}
}

/*ServiceBrokerAuthLoginParams contains all the parameters to send to the API endpoint
for the service broker auth login operation typically these are written to a http.Request
*/
type ServiceBrokerAuthLoginParams struct {

	/*AccessType
	  Determines if a refresh token is returned

	*/
	AccessType *string
	/*RedirectURL
	  The URL to redirect to after login/registration

	*/
	RedirectURL *string
	/*UserID
	  The user id of the user

	*/
	UserID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the service broker auth login params
func (o *ServiceBrokerAuthLoginParams) WithTimeout(timeout time.Duration) *ServiceBrokerAuthLoginParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service broker auth login params
func (o *ServiceBrokerAuthLoginParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service broker auth login params
func (o *ServiceBrokerAuthLoginParams) WithContext(ctx context.Context) *ServiceBrokerAuthLoginParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service broker auth login params
func (o *ServiceBrokerAuthLoginParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service broker auth login params
func (o *ServiceBrokerAuthLoginParams) WithHTTPClient(client *http.Client) *ServiceBrokerAuthLoginParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service broker auth login params
func (o *ServiceBrokerAuthLoginParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessType adds the accessType to the service broker auth login params
func (o *ServiceBrokerAuthLoginParams) WithAccessType(accessType *string) *ServiceBrokerAuthLoginParams {
	o.SetAccessType(accessType)
	return o
}

// SetAccessType adds the accessType to the service broker auth login params
func (o *ServiceBrokerAuthLoginParams) SetAccessType(accessType *string) {
	o.AccessType = accessType
}

// WithRedirectURL adds the redirectURL to the service broker auth login params
func (o *ServiceBrokerAuthLoginParams) WithRedirectURL(redirectURL *string) *ServiceBrokerAuthLoginParams {
	o.SetRedirectURL(redirectURL)
	return o
}

// SetRedirectURL adds the redirectUrl to the service broker auth login params
func (o *ServiceBrokerAuthLoginParams) SetRedirectURL(redirectURL *string) {
	o.RedirectURL = redirectURL
}

// WithUserID adds the userID to the service broker auth login params
func (o *ServiceBrokerAuthLoginParams) WithUserID(userID *string) *ServiceBrokerAuthLoginParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the service broker auth login params
func (o *ServiceBrokerAuthLoginParams) SetUserID(userID *string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *ServiceBrokerAuthLoginParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccessType != nil {

		// query param access_type
		var qrAccessType string
		if o.AccessType != nil {
			qrAccessType = *o.AccessType
		}
		qAccessType := qrAccessType
		if qAccessType != "" {
			if err := r.SetQueryParam("access_type", qAccessType); err != nil {
				return err
			}
		}

	}

	if o.RedirectURL != nil {

		// query param redirect_url
		var qrRedirectURL string
		if o.RedirectURL != nil {
			qrRedirectURL = *o.RedirectURL
		}
		qRedirectURL := qrRedirectURL
		if qRedirectURL != "" {
			if err := r.SetQueryParam("redirect_url", qRedirectURL); err != nil {
				return err
			}
		}

	}

	if o.UserID != nil {

		// query param user_id
		var qrUserID string
		if o.UserID != nil {
			qrUserID = *o.UserID
		}
		qUserID := qrUserID
		if qUserID != "" {
			if err := r.SetQueryParam("user_id", qUserID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
