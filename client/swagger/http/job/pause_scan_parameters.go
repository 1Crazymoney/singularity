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

// NewPauseScanParams creates a new PauseScanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPauseScanParams() *PauseScanParams {
	return &PauseScanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPauseScanParamsWithTimeout creates a new PauseScanParams object
// with the ability to set a timeout on a request.
func NewPauseScanParamsWithTimeout(timeout time.Duration) *PauseScanParams {
	return &PauseScanParams{
		timeout: timeout,
	}
}

// NewPauseScanParamsWithContext creates a new PauseScanParams object
// with the ability to set a context for a request.
func NewPauseScanParamsWithContext(ctx context.Context) *PauseScanParams {
	return &PauseScanParams{
		Context: ctx,
	}
}

// NewPauseScanParamsWithHTTPClient creates a new PauseScanParams object
// with the ability to set a custom HTTPClient for a request.
func NewPauseScanParamsWithHTTPClient(client *http.Client) *PauseScanParams {
	return &PauseScanParams{
		HTTPClient: client,
	}
}

/*
PauseScanParams contains all the parameters to send to the API endpoint

	for the pause scan operation.

	Typically these are written to a http.Request.
*/
type PauseScanParams struct {

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

// WithDefaults hydrates default values in the pause scan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PauseScanParams) WithDefaults() *PauseScanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pause scan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PauseScanParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pause scan params
func (o *PauseScanParams) WithTimeout(timeout time.Duration) *PauseScanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pause scan params
func (o *PauseScanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pause scan params
func (o *PauseScanParams) WithContext(ctx context.Context) *PauseScanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pause scan params
func (o *PauseScanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pause scan params
func (o *PauseScanParams) WithHTTPClient(client *http.Client) *PauseScanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pause scan params
func (o *PauseScanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the pause scan params
func (o *PauseScanParams) WithID(id string) *PauseScanParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the pause scan params
func (o *PauseScanParams) SetID(id string) {
	o.ID = id
}

// WithName adds the name to the pause scan params
func (o *PauseScanParams) WithName(name string) *PauseScanParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the pause scan params
func (o *PauseScanParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PauseScanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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