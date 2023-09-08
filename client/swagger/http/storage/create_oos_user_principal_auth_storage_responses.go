// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// CreateOosUserPrincipalAuthStorageReader is a Reader for the CreateOosUserPrincipalAuthStorage structure.
type CreateOosUserPrincipalAuthStorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOosUserPrincipalAuthStorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateOosUserPrincipalAuthStorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateOosUserPrincipalAuthStorageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateOosUserPrincipalAuthStorageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/oos/user_principal_auth] CreateOosUser_principal_authStorage", response, response.Code())
	}
}

// NewCreateOosUserPrincipalAuthStorageOK creates a CreateOosUserPrincipalAuthStorageOK with default headers values
func NewCreateOosUserPrincipalAuthStorageOK() *CreateOosUserPrincipalAuthStorageOK {
	return &CreateOosUserPrincipalAuthStorageOK{}
}

/*
CreateOosUserPrincipalAuthStorageOK describes a response with status code 200, with default header values.

OK
*/
type CreateOosUserPrincipalAuthStorageOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this create oos user principal auth storage o k response has a 2xx status code
func (o *CreateOosUserPrincipalAuthStorageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create oos user principal auth storage o k response has a 3xx status code
func (o *CreateOosUserPrincipalAuthStorageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create oos user principal auth storage o k response has a 4xx status code
func (o *CreateOosUserPrincipalAuthStorageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create oos user principal auth storage o k response has a 5xx status code
func (o *CreateOosUserPrincipalAuthStorageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create oos user principal auth storage o k response a status code equal to that given
func (o *CreateOosUserPrincipalAuthStorageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create oos user principal auth storage o k response
func (o *CreateOosUserPrincipalAuthStorageOK) Code() int {
	return 200
}

func (o *CreateOosUserPrincipalAuthStorageOK) Error() string {
	return fmt.Sprintf("[POST /storage/oos/user_principal_auth][%d] createOosUserPrincipalAuthStorageOK  %+v", 200, o.Payload)
}

func (o *CreateOosUserPrincipalAuthStorageOK) String() string {
	return fmt.Sprintf("[POST /storage/oos/user_principal_auth][%d] createOosUserPrincipalAuthStorageOK  %+v", 200, o.Payload)
}

func (o *CreateOosUserPrincipalAuthStorageOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *CreateOosUserPrincipalAuthStorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOosUserPrincipalAuthStorageBadRequest creates a CreateOosUserPrincipalAuthStorageBadRequest with default headers values
func NewCreateOosUserPrincipalAuthStorageBadRequest() *CreateOosUserPrincipalAuthStorageBadRequest {
	return &CreateOosUserPrincipalAuthStorageBadRequest{}
}

/*
CreateOosUserPrincipalAuthStorageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateOosUserPrincipalAuthStorageBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create oos user principal auth storage bad request response has a 2xx status code
func (o *CreateOosUserPrincipalAuthStorageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create oos user principal auth storage bad request response has a 3xx status code
func (o *CreateOosUserPrincipalAuthStorageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create oos user principal auth storage bad request response has a 4xx status code
func (o *CreateOosUserPrincipalAuthStorageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create oos user principal auth storage bad request response has a 5xx status code
func (o *CreateOosUserPrincipalAuthStorageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create oos user principal auth storage bad request response a status code equal to that given
func (o *CreateOosUserPrincipalAuthStorageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create oos user principal auth storage bad request response
func (o *CreateOosUserPrincipalAuthStorageBadRequest) Code() int {
	return 400
}

func (o *CreateOosUserPrincipalAuthStorageBadRequest) Error() string {
	return fmt.Sprintf("[POST /storage/oos/user_principal_auth][%d] createOosUserPrincipalAuthStorageBadRequest  %+v", 400, o.Payload)
}

func (o *CreateOosUserPrincipalAuthStorageBadRequest) String() string {
	return fmt.Sprintf("[POST /storage/oos/user_principal_auth][%d] createOosUserPrincipalAuthStorageBadRequest  %+v", 400, o.Payload)
}

func (o *CreateOosUserPrincipalAuthStorageBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateOosUserPrincipalAuthStorageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOosUserPrincipalAuthStorageInternalServerError creates a CreateOosUserPrincipalAuthStorageInternalServerError with default headers values
func NewCreateOosUserPrincipalAuthStorageInternalServerError() *CreateOosUserPrincipalAuthStorageInternalServerError {
	return &CreateOosUserPrincipalAuthStorageInternalServerError{}
}

/*
CreateOosUserPrincipalAuthStorageInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateOosUserPrincipalAuthStorageInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create oos user principal auth storage internal server error response has a 2xx status code
func (o *CreateOosUserPrincipalAuthStorageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create oos user principal auth storage internal server error response has a 3xx status code
func (o *CreateOosUserPrincipalAuthStorageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create oos user principal auth storage internal server error response has a 4xx status code
func (o *CreateOosUserPrincipalAuthStorageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create oos user principal auth storage internal server error response has a 5xx status code
func (o *CreateOosUserPrincipalAuthStorageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create oos user principal auth storage internal server error response a status code equal to that given
func (o *CreateOosUserPrincipalAuthStorageInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create oos user principal auth storage internal server error response
func (o *CreateOosUserPrincipalAuthStorageInternalServerError) Code() int {
	return 500
}

func (o *CreateOosUserPrincipalAuthStorageInternalServerError) Error() string {
	return fmt.Sprintf("[POST /storage/oos/user_principal_auth][%d] createOosUserPrincipalAuthStorageInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateOosUserPrincipalAuthStorageInternalServerError) String() string {
	return fmt.Sprintf("[POST /storage/oos/user_principal_auth][%d] createOosUserPrincipalAuthStorageInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateOosUserPrincipalAuthStorageInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateOosUserPrincipalAuthStorageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}