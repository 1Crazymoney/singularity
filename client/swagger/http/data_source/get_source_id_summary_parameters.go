// Code generated by go-swagger; DO NOT EDIT.

package data_source

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

// NewGetSourceIDSummaryParams creates a new GetSourceIDSummaryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSourceIDSummaryParams() *GetSourceIDSummaryParams {
	return &GetSourceIDSummaryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSourceIDSummaryParamsWithTimeout creates a new GetSourceIDSummaryParams object
// with the ability to set a timeout on a request.
func NewGetSourceIDSummaryParamsWithTimeout(timeout time.Duration) *GetSourceIDSummaryParams {
	return &GetSourceIDSummaryParams{
		timeout: timeout,
	}
}

// NewGetSourceIDSummaryParamsWithContext creates a new GetSourceIDSummaryParams object
// with the ability to set a context for a request.
func NewGetSourceIDSummaryParamsWithContext(ctx context.Context) *GetSourceIDSummaryParams {
	return &GetSourceIDSummaryParams{
		Context: ctx,
	}
}

// NewGetSourceIDSummaryParamsWithHTTPClient creates a new GetSourceIDSummaryParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSourceIDSummaryParamsWithHTTPClient(client *http.Client) *GetSourceIDSummaryParams {
	return &GetSourceIDSummaryParams{
		HTTPClient: client,
	}
}

/*
GetSourceIDSummaryParams contains all the parameters to send to the API endpoint

	for the get source ID summary operation.

	Typically these are written to a http.Request.
*/
type GetSourceIDSummaryParams struct {

	/* ID.

	   Source ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get source ID summary params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSourceIDSummaryParams) WithDefaults() *GetSourceIDSummaryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get source ID summary params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSourceIDSummaryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get source ID summary params
func (o *GetSourceIDSummaryParams) WithTimeout(timeout time.Duration) *GetSourceIDSummaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get source ID summary params
func (o *GetSourceIDSummaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get source ID summary params
func (o *GetSourceIDSummaryParams) WithContext(ctx context.Context) *GetSourceIDSummaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get source ID summary params
func (o *GetSourceIDSummaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get source ID summary params
func (o *GetSourceIDSummaryParams) WithHTTPClient(client *http.Client) *GetSourceIDSummaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get source ID summary params
func (o *GetSourceIDSummaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get source ID summary params
func (o *GetSourceIDSummaryParams) WithID(id string) *GetSourceIDSummaryParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get source ID summary params
func (o *GetSourceIDSummaryParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetSourceIDSummaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
