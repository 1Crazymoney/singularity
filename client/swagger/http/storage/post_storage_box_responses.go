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

// PostStorageBoxReader is a Reader for the PostStorageBox structure.
type PostStorageBoxReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostStorageBoxReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostStorageBoxOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostStorageBoxBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostStorageBoxInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/box] PostStorageBox", response, response.Code())
	}
}

// NewPostStorageBoxOK creates a PostStorageBoxOK with default headers values
func NewPostStorageBoxOK() *PostStorageBoxOK {
	return &PostStorageBoxOK{}
}

/*
PostStorageBoxOK describes a response with status code 200, with default header values.

OK
*/
type PostStorageBoxOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this post storage box o k response has a 2xx status code
func (o *PostStorageBoxOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post storage box o k response has a 3xx status code
func (o *PostStorageBoxOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage box o k response has a 4xx status code
func (o *PostStorageBoxOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post storage box o k response has a 5xx status code
func (o *PostStorageBoxOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post storage box o k response a status code equal to that given
func (o *PostStorageBoxOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post storage box o k response
func (o *PostStorageBoxOK) Code() int {
	return 200
}

func (o *PostStorageBoxOK) Error() string {
	return fmt.Sprintf("[POST /storage/box][%d] postStorageBoxOK  %+v", 200, o.Payload)
}

func (o *PostStorageBoxOK) String() string {
	return fmt.Sprintf("[POST /storage/box][%d] postStorageBoxOK  %+v", 200, o.Payload)
}

func (o *PostStorageBoxOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *PostStorageBoxOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStorageBoxBadRequest creates a PostStorageBoxBadRequest with default headers values
func NewPostStorageBoxBadRequest() *PostStorageBoxBadRequest {
	return &PostStorageBoxBadRequest{}
}

/*
PostStorageBoxBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostStorageBoxBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post storage box bad request response has a 2xx status code
func (o *PostStorageBoxBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post storage box bad request response has a 3xx status code
func (o *PostStorageBoxBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage box bad request response has a 4xx status code
func (o *PostStorageBoxBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post storage box bad request response has a 5xx status code
func (o *PostStorageBoxBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post storage box bad request response a status code equal to that given
func (o *PostStorageBoxBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post storage box bad request response
func (o *PostStorageBoxBadRequest) Code() int {
	return 400
}

func (o *PostStorageBoxBadRequest) Error() string {
	return fmt.Sprintf("[POST /storage/box][%d] postStorageBoxBadRequest  %+v", 400, o.Payload)
}

func (o *PostStorageBoxBadRequest) String() string {
	return fmt.Sprintf("[POST /storage/box][%d] postStorageBoxBadRequest  %+v", 400, o.Payload)
}

func (o *PostStorageBoxBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostStorageBoxBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStorageBoxInternalServerError creates a PostStorageBoxInternalServerError with default headers values
func NewPostStorageBoxInternalServerError() *PostStorageBoxInternalServerError {
	return &PostStorageBoxInternalServerError{}
}

/*
PostStorageBoxInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostStorageBoxInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post storage box internal server error response has a 2xx status code
func (o *PostStorageBoxInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post storage box internal server error response has a 3xx status code
func (o *PostStorageBoxInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage box internal server error response has a 4xx status code
func (o *PostStorageBoxInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post storage box internal server error response has a 5xx status code
func (o *PostStorageBoxInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post storage box internal server error response a status code equal to that given
func (o *PostStorageBoxInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post storage box internal server error response
func (o *PostStorageBoxInternalServerError) Code() int {
	return 500
}

func (o *PostStorageBoxInternalServerError) Error() string {
	return fmt.Sprintf("[POST /storage/box][%d] postStorageBoxInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStorageBoxInternalServerError) String() string {
	return fmt.Sprintf("[POST /storage/box][%d] postStorageBoxInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStorageBoxInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostStorageBoxInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}