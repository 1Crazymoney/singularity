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

// NewPostStorageS3HuaweiobsParams creates a new PostStorageS3HuaweiobsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostStorageS3HuaweiobsParams() *PostStorageS3HuaweiobsParams {
	return &PostStorageS3HuaweiobsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostStorageS3HuaweiobsParamsWithTimeout creates a new PostStorageS3HuaweiobsParams object
// with the ability to set a timeout on a request.
func NewPostStorageS3HuaweiobsParamsWithTimeout(timeout time.Duration) *PostStorageS3HuaweiobsParams {
	return &PostStorageS3HuaweiobsParams{
		timeout: timeout,
	}
}

// NewPostStorageS3HuaweiobsParamsWithContext creates a new PostStorageS3HuaweiobsParams object
// with the ability to set a context for a request.
func NewPostStorageS3HuaweiobsParamsWithContext(ctx context.Context) *PostStorageS3HuaweiobsParams {
	return &PostStorageS3HuaweiobsParams{
		Context: ctx,
	}
}

// NewPostStorageS3HuaweiobsParamsWithHTTPClient creates a new PostStorageS3HuaweiobsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostStorageS3HuaweiobsParamsWithHTTPClient(client *http.Client) *PostStorageS3HuaweiobsParams {
	return &PostStorageS3HuaweiobsParams{
		HTTPClient: client,
	}
}

/*
PostStorageS3HuaweiobsParams contains all the parameters to send to the API endpoint

	for the post storage s3 huaweiobs operation.

	Typically these are written to a http.Request.
*/
type PostStorageS3HuaweiobsParams struct {

	/* Request.

	   Request body
	*/
	Request *models.StorageCreateS3HuaweiOBSStorageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post storage s3 huaweiobs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostStorageS3HuaweiobsParams) WithDefaults() *PostStorageS3HuaweiobsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post storage s3 huaweiobs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostStorageS3HuaweiobsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post storage s3 huaweiobs params
func (o *PostStorageS3HuaweiobsParams) WithTimeout(timeout time.Duration) *PostStorageS3HuaweiobsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post storage s3 huaweiobs params
func (o *PostStorageS3HuaweiobsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post storage s3 huaweiobs params
func (o *PostStorageS3HuaweiobsParams) WithContext(ctx context.Context) *PostStorageS3HuaweiobsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post storage s3 huaweiobs params
func (o *PostStorageS3HuaweiobsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post storage s3 huaweiobs params
func (o *PostStorageS3HuaweiobsParams) WithHTTPClient(client *http.Client) *PostStorageS3HuaweiobsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post storage s3 huaweiobs params
func (o *PostStorageS3HuaweiobsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post storage s3 huaweiobs params
func (o *PostStorageS3HuaweiobsParams) WithRequest(request *models.StorageCreateS3HuaweiOBSStorageRequest) *PostStorageS3HuaweiobsParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post storage s3 huaweiobs params
func (o *PostStorageS3HuaweiobsParams) SetRequest(request *models.StorageCreateS3HuaweiOBSStorageRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostStorageS3HuaweiobsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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