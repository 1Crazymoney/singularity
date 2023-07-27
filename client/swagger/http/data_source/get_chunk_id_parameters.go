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

// NewGetChunkIDParams creates a new GetChunkIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetChunkIDParams() *GetChunkIDParams {
	return &GetChunkIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetChunkIDParamsWithTimeout creates a new GetChunkIDParams object
// with the ability to set a timeout on a request.
func NewGetChunkIDParamsWithTimeout(timeout time.Duration) *GetChunkIDParams {
	return &GetChunkIDParams{
		timeout: timeout,
	}
}

// NewGetChunkIDParamsWithContext creates a new GetChunkIDParams object
// with the ability to set a context for a request.
func NewGetChunkIDParamsWithContext(ctx context.Context) *GetChunkIDParams {
	return &GetChunkIDParams{
		Context: ctx,
	}
}

// NewGetChunkIDParamsWithHTTPClient creates a new GetChunkIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetChunkIDParamsWithHTTPClient(client *http.Client) *GetChunkIDParams {
	return &GetChunkIDParams{
		HTTPClient: client,
	}
}

/*
GetChunkIDParams contains all the parameters to send to the API endpoint

	for the get chunk ID operation.

	Typically these are written to a http.Request.
*/
type GetChunkIDParams struct {

	/* ID.

	   Chunk ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get chunk ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetChunkIDParams) WithDefaults() *GetChunkIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get chunk ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetChunkIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get chunk ID params
func (o *GetChunkIDParams) WithTimeout(timeout time.Duration) *GetChunkIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get chunk ID params
func (o *GetChunkIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get chunk ID params
func (o *GetChunkIDParams) WithContext(ctx context.Context) *GetChunkIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get chunk ID params
func (o *GetChunkIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get chunk ID params
func (o *GetChunkIDParams) WithHTTPClient(client *http.Client) *GetChunkIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get chunk ID params
func (o *GetChunkIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get chunk ID params
func (o *GetChunkIDParams) WithID(id string) *GetChunkIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get chunk ID params
func (o *GetChunkIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetChunkIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
