// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewStartPackAllParams creates a new StartPackAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStartPackAllParams() *StartPackAllParams {
	return &StartPackAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStartPackAllParamsWithTimeout creates a new StartPackAllParams object
// with the ability to set a timeout on a request.
func NewStartPackAllParamsWithTimeout(timeout time.Duration) *StartPackAllParams {
	return &StartPackAllParams{
		timeout: timeout,
	}
}

// NewStartPackAllParamsWithContext creates a new StartPackAllParams object
// with the ability to set a context for a request.
func NewStartPackAllParamsWithContext(ctx context.Context) *StartPackAllParams {
	return &StartPackAllParams{
		Context: ctx,
	}
}

// NewStartPackAllParamsWithHTTPClient creates a new StartPackAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewStartPackAllParamsWithHTTPClient(client *http.Client) *StartPackAllParams {
	return &StartPackAllParams{
		HTTPClient: client,
	}
}

/*
StartPackAllParams contains all the parameters to send to the API endpoint

	for the start pack all operation.

	Typically these are written to a http.Request.
*/
type StartPackAllParams struct {

	/* ID.

	   Preparation ID or name
	*/
	ID string

	/* Name.

	   Storage ID or name
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the start pack all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartPackAllParams) WithDefaults() *StartPackAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the start pack all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartPackAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the start pack all params
func (o *StartPackAllParams) WithTimeout(timeout time.Duration) *StartPackAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start pack all params
func (o *StartPackAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start pack all params
func (o *StartPackAllParams) WithContext(ctx context.Context) *StartPackAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start pack all params
func (o *StartPackAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start pack all params
func (o *StartPackAllParams) WithHTTPClient(client *http.Client) *StartPackAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start pack all params
func (o *StartPackAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the start pack all params
func (o *StartPackAllParams) WithID(id string) *StartPackAllParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the start pack all params
func (o *StartPackAllParams) SetID(id string) {
	o.ID = id
}

// WithName adds the name to the start pack all params
func (o *StartPackAllParams) WithName(name string) *StartPackAllParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the start pack all params
func (o *StartPackAllParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *StartPackAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}