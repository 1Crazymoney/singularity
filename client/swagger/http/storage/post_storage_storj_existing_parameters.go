// Code generated by go-swagger; DO NOT EDIT.

package storage

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

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// NewPostStorageStorjExistingParams creates a new PostStorageStorjExistingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostStorageStorjExistingParams() *PostStorageStorjExistingParams {
	return &PostStorageStorjExistingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostStorageStorjExistingParamsWithTimeout creates a new PostStorageStorjExistingParams object
// with the ability to set a timeout on a request.
func NewPostStorageStorjExistingParamsWithTimeout(timeout time.Duration) *PostStorageStorjExistingParams {
	return &PostStorageStorjExistingParams{
		timeout: timeout,
	}
}

// NewPostStorageStorjExistingParamsWithContext creates a new PostStorageStorjExistingParams object
// with the ability to set a context for a request.
func NewPostStorageStorjExistingParamsWithContext(ctx context.Context) *PostStorageStorjExistingParams {
	return &PostStorageStorjExistingParams{
		Context: ctx,
	}
}

// NewPostStorageStorjExistingParamsWithHTTPClient creates a new PostStorageStorjExistingParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostStorageStorjExistingParamsWithHTTPClient(client *http.Client) *PostStorageStorjExistingParams {
	return &PostStorageStorjExistingParams{
		HTTPClient: client,
	}
}

/*
PostStorageStorjExistingParams contains all the parameters to send to the API endpoint

	for the post storage storj existing operation.

	Typically these are written to a http.Request.
*/
type PostStorageStorjExistingParams struct {

	/* Request.

	   Request body
	*/
	Request *models.StorageCreateStorjExistingStorageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post storage storj existing params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostStorageStorjExistingParams) WithDefaults() *PostStorageStorjExistingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post storage storj existing params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostStorageStorjExistingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post storage storj existing params
func (o *PostStorageStorjExistingParams) WithTimeout(timeout time.Duration) *PostStorageStorjExistingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post storage storj existing params
func (o *PostStorageStorjExistingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post storage storj existing params
func (o *PostStorageStorjExistingParams) WithContext(ctx context.Context) *PostStorageStorjExistingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post storage storj existing params
func (o *PostStorageStorjExistingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post storage storj existing params
func (o *PostStorageStorjExistingParams) WithHTTPClient(client *http.Client) *PostStorageStorjExistingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post storage storj existing params
func (o *PostStorageStorjExistingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post storage storj existing params
func (o *PostStorageStorjExistingParams) WithRequest(request *models.StorageCreateStorjExistingStorageRequest) *PostStorageStorjExistingParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post storage storj existing params
func (o *PostStorageStorjExistingParams) SetRequest(request *models.StorageCreateStorjExistingStorageRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostStorageStorjExistingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}