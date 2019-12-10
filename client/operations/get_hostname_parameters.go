// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewGetHostnameParams creates a new GetHostnameParams object
// with the default values initialized.
func NewGetHostnameParams() *GetHostnameParams {

	return &GetHostnameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetHostnameParamsWithTimeout creates a new GetHostnameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetHostnameParamsWithTimeout(timeout time.Duration) *GetHostnameParams {

	return &GetHostnameParams{

		timeout: timeout,
	}
}

// NewGetHostnameParamsWithContext creates a new GetHostnameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetHostnameParamsWithContext(ctx context.Context) *GetHostnameParams {

	return &GetHostnameParams{

		Context: ctx,
	}
}

// NewGetHostnameParamsWithHTTPClient creates a new GetHostnameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetHostnameParamsWithHTTPClient(client *http.Client) *GetHostnameParams {

	return &GetHostnameParams{
		HTTPClient: client,
	}
}

/*GetHostnameParams contains all the parameters to send to the API endpoint
for the get hostname operation typically these are written to a http.Request
*/
type GetHostnameParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get hostname params
func (o *GetHostnameParams) WithTimeout(timeout time.Duration) *GetHostnameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get hostname params
func (o *GetHostnameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get hostname params
func (o *GetHostnameParams) WithContext(ctx context.Context) *GetHostnameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get hostname params
func (o *GetHostnameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get hostname params
func (o *GetHostnameParams) WithHTTPClient(client *http.Client) *GetHostnameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get hostname params
func (o *GetHostnameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetHostnameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
