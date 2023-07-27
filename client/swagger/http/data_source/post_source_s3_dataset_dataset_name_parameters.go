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

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// NewPostSourceS3DatasetDatasetNameParams creates a new PostSourceS3DatasetDatasetNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostSourceS3DatasetDatasetNameParams() *PostSourceS3DatasetDatasetNameParams {
	return &PostSourceS3DatasetDatasetNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostSourceS3DatasetDatasetNameParamsWithTimeout creates a new PostSourceS3DatasetDatasetNameParams object
// with the ability to set a timeout on a request.
func NewPostSourceS3DatasetDatasetNameParamsWithTimeout(timeout time.Duration) *PostSourceS3DatasetDatasetNameParams {
	return &PostSourceS3DatasetDatasetNameParams{
		timeout: timeout,
	}
}

// NewPostSourceS3DatasetDatasetNameParamsWithContext creates a new PostSourceS3DatasetDatasetNameParams object
// with the ability to set a context for a request.
func NewPostSourceS3DatasetDatasetNameParamsWithContext(ctx context.Context) *PostSourceS3DatasetDatasetNameParams {
	return &PostSourceS3DatasetDatasetNameParams{
		Context: ctx,
	}
}

// NewPostSourceS3DatasetDatasetNameParamsWithHTTPClient creates a new PostSourceS3DatasetDatasetNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostSourceS3DatasetDatasetNameParamsWithHTTPClient(client *http.Client) *PostSourceS3DatasetDatasetNameParams {
	return &PostSourceS3DatasetDatasetNameParams{
		HTTPClient: client,
	}
}

/*
PostSourceS3DatasetDatasetNameParams contains all the parameters to send to the API endpoint

	for the post source s3 dataset dataset name operation.

	Typically these are written to a http.Request.
*/
type PostSourceS3DatasetDatasetNameParams struct {

	/* DatasetName.

	   Dataset name
	*/
	DatasetName string

	/* Request.

	   Request body
	*/
	Request *models.DatasourceS3Request

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post source s3 dataset dataset name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSourceS3DatasetDatasetNameParams) WithDefaults() *PostSourceS3DatasetDatasetNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post source s3 dataset dataset name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSourceS3DatasetDatasetNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post source s3 dataset dataset name params
func (o *PostSourceS3DatasetDatasetNameParams) WithTimeout(timeout time.Duration) *PostSourceS3DatasetDatasetNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post source s3 dataset dataset name params
func (o *PostSourceS3DatasetDatasetNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post source s3 dataset dataset name params
func (o *PostSourceS3DatasetDatasetNameParams) WithContext(ctx context.Context) *PostSourceS3DatasetDatasetNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post source s3 dataset dataset name params
func (o *PostSourceS3DatasetDatasetNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post source s3 dataset dataset name params
func (o *PostSourceS3DatasetDatasetNameParams) WithHTTPClient(client *http.Client) *PostSourceS3DatasetDatasetNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post source s3 dataset dataset name params
func (o *PostSourceS3DatasetDatasetNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatasetName adds the datasetName to the post source s3 dataset dataset name params
func (o *PostSourceS3DatasetDatasetNameParams) WithDatasetName(datasetName string) *PostSourceS3DatasetDatasetNameParams {
	o.SetDatasetName(datasetName)
	return o
}

// SetDatasetName adds the datasetName to the post source s3 dataset dataset name params
func (o *PostSourceS3DatasetDatasetNameParams) SetDatasetName(datasetName string) {
	o.DatasetName = datasetName
}

// WithRequest adds the request to the post source s3 dataset dataset name params
func (o *PostSourceS3DatasetDatasetNameParams) WithRequest(request *models.DatasourceS3Request) *PostSourceS3DatasetDatasetNameParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post source s3 dataset dataset name params
func (o *PostSourceS3DatasetDatasetNameParams) SetRequest(request *models.DatasourceS3Request) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostSourceS3DatasetDatasetNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param datasetName
	if err := r.SetPathParam("datasetName", o.DatasetName); err != nil {
		return err
	}
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
