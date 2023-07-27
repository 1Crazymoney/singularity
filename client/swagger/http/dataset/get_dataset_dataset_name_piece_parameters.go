// Code generated by go-swagger; DO NOT EDIT.

package dataset

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

// NewGetDatasetDatasetNamePieceParams creates a new GetDatasetDatasetNamePieceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDatasetDatasetNamePieceParams() *GetDatasetDatasetNamePieceParams {
	return &GetDatasetDatasetNamePieceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDatasetDatasetNamePieceParamsWithTimeout creates a new GetDatasetDatasetNamePieceParams object
// with the ability to set a timeout on a request.
func NewGetDatasetDatasetNamePieceParamsWithTimeout(timeout time.Duration) *GetDatasetDatasetNamePieceParams {
	return &GetDatasetDatasetNamePieceParams{
		timeout: timeout,
	}
}

// NewGetDatasetDatasetNamePieceParamsWithContext creates a new GetDatasetDatasetNamePieceParams object
// with the ability to set a context for a request.
func NewGetDatasetDatasetNamePieceParamsWithContext(ctx context.Context) *GetDatasetDatasetNamePieceParams {
	return &GetDatasetDatasetNamePieceParams{
		Context: ctx,
	}
}

// NewGetDatasetDatasetNamePieceParamsWithHTTPClient creates a new GetDatasetDatasetNamePieceParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDatasetDatasetNamePieceParamsWithHTTPClient(client *http.Client) *GetDatasetDatasetNamePieceParams {
	return &GetDatasetDatasetNamePieceParams{
		HTTPClient: client,
	}
}

/*
GetDatasetDatasetNamePieceParams contains all the parameters to send to the API endpoint

	for the get dataset dataset name piece operation.

	Typically these are written to a http.Request.
*/
type GetDatasetDatasetNamePieceParams struct {

	/* DatasetName.

	   Dataset name
	*/
	DatasetName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get dataset dataset name piece params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDatasetDatasetNamePieceParams) WithDefaults() *GetDatasetDatasetNamePieceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get dataset dataset name piece params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDatasetDatasetNamePieceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get dataset dataset name piece params
func (o *GetDatasetDatasetNamePieceParams) WithTimeout(timeout time.Duration) *GetDatasetDatasetNamePieceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get dataset dataset name piece params
func (o *GetDatasetDatasetNamePieceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get dataset dataset name piece params
func (o *GetDatasetDatasetNamePieceParams) WithContext(ctx context.Context) *GetDatasetDatasetNamePieceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get dataset dataset name piece params
func (o *GetDatasetDatasetNamePieceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get dataset dataset name piece params
func (o *GetDatasetDatasetNamePieceParams) WithHTTPClient(client *http.Client) *GetDatasetDatasetNamePieceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get dataset dataset name piece params
func (o *GetDatasetDatasetNamePieceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatasetName adds the datasetName to the get dataset dataset name piece params
func (o *GetDatasetDatasetNamePieceParams) WithDatasetName(datasetName string) *GetDatasetDatasetNamePieceParams {
	o.SetDatasetName(datasetName)
	return o
}

// SetDatasetName adds the datasetName to the get dataset dataset name piece params
func (o *GetDatasetDatasetNamePieceParams) SetDatasetName(datasetName string) {
	o.DatasetName = datasetName
}

// WriteToRequest writes these params to a swagger request
func (o *GetDatasetDatasetNamePieceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param datasetName
	if err := r.SetPathParam("datasetName", o.DatasetName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
