// Code generated by go-swagger; DO NOT EDIT.

package file

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new file API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for file API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetFileID(params *GetFileIDParams, opts ...ClientOption) (*GetFileIDOK, error)

	GetFileIDDeals(params *GetFileIDDealsParams, opts ...ClientOption) (*GetFileIDDealsOK, error)

	PostFileIDPrepareToPack(params *PostFileIDPrepareToPackParams, opts ...ClientOption) (*PostFileIDPrepareToPackOK, error)

	PostPreparationIDSourceNameFile(params *PostPreparationIDSourceNameFileParams, opts ...ClientOption) (*PostPreparationIDSourceNameFileOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetFileID gets details about a file
*/
func (a *Client) GetFileID(params *GetFileIDParams, opts ...ClientOption) (*GetFileIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFileIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetFileID",
		Method:             "GET",
		PathPattern:        "/file/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetFileIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFileIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetFileID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetFileIDDeals gets all deals that have been made for a file
*/
func (a *Client) GetFileIDDeals(params *GetFileIDDealsParams, opts ...ClientOption) (*GetFileIDDealsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFileIDDealsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetFileIDDeals",
		Method:             "GET",
		PathPattern:        "/file/{id}/deals",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetFileIDDealsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFileIDDealsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetFileIDDeals: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostFileIDPrepareToPack prepares job for a given item
*/
func (a *Client) PostFileIDPrepareToPack(params *PostFileIDPrepareToPackParams, opts ...ClientOption) (*PostFileIDPrepareToPackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostFileIDPrepareToPackParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostFileIDPrepareToPack",
		Method:             "POST",
		PathPattern:        "/file/{id}/prepare_to_pack",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostFileIDPrepareToPackReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostFileIDPrepareToPackOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostFileIDPrepareToPack: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostPreparationIDSourceNameFile pushes a file to be queued

Tells Singularity that something is ready to be grabbed for data preparation
*/
func (a *Client) PostPreparationIDSourceNameFile(params *PostPreparationIDSourceNameFileParams, opts ...ClientOption) (*PostPreparationIDSourceNameFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPreparationIDSourceNameFileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostPreparationIDSourceNameFile",
		Method:             "POST",
		PathPattern:        "/preparation/{id}/source/{name}/file",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostPreparationIDSourceNameFileReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostPreparationIDSourceNameFileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostPreparationIDSourceNameFile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}