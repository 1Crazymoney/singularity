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

// PostStorageS3ArvancloudReader is a Reader for the PostStorageS3Arvancloud structure.
type PostStorageS3ArvancloudReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostStorageS3ArvancloudReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostStorageS3ArvancloudOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostStorageS3ArvancloudBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostStorageS3ArvancloudInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/s3/arvancloud] PostStorageS3Arvancloud", response, response.Code())
	}
}

// NewPostStorageS3ArvancloudOK creates a PostStorageS3ArvancloudOK with default headers values
func NewPostStorageS3ArvancloudOK() *PostStorageS3ArvancloudOK {
	return &PostStorageS3ArvancloudOK{}
}

/*
PostStorageS3ArvancloudOK describes a response with status code 200, with default header values.

OK
*/
type PostStorageS3ArvancloudOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this post storage s3 arvancloud o k response has a 2xx status code
func (o *PostStorageS3ArvancloudOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post storage s3 arvancloud o k response has a 3xx status code
func (o *PostStorageS3ArvancloudOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage s3 arvancloud o k response has a 4xx status code
func (o *PostStorageS3ArvancloudOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post storage s3 arvancloud o k response has a 5xx status code
func (o *PostStorageS3ArvancloudOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post storage s3 arvancloud o k response a status code equal to that given
func (o *PostStorageS3ArvancloudOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post storage s3 arvancloud o k response
func (o *PostStorageS3ArvancloudOK) Code() int {
	return 200
}

func (o *PostStorageS3ArvancloudOK) Error() string {
	return fmt.Sprintf("[POST /storage/s3/arvancloud][%d] postStorageS3ArvancloudOK  %+v", 200, o.Payload)
}

func (o *PostStorageS3ArvancloudOK) String() string {
	return fmt.Sprintf("[POST /storage/s3/arvancloud][%d] postStorageS3ArvancloudOK  %+v", 200, o.Payload)
}

func (o *PostStorageS3ArvancloudOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *PostStorageS3ArvancloudOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStorageS3ArvancloudBadRequest creates a PostStorageS3ArvancloudBadRequest with default headers values
func NewPostStorageS3ArvancloudBadRequest() *PostStorageS3ArvancloudBadRequest {
	return &PostStorageS3ArvancloudBadRequest{}
}

/*
PostStorageS3ArvancloudBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostStorageS3ArvancloudBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post storage s3 arvancloud bad request response has a 2xx status code
func (o *PostStorageS3ArvancloudBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post storage s3 arvancloud bad request response has a 3xx status code
func (o *PostStorageS3ArvancloudBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage s3 arvancloud bad request response has a 4xx status code
func (o *PostStorageS3ArvancloudBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post storage s3 arvancloud bad request response has a 5xx status code
func (o *PostStorageS3ArvancloudBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post storage s3 arvancloud bad request response a status code equal to that given
func (o *PostStorageS3ArvancloudBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post storage s3 arvancloud bad request response
func (o *PostStorageS3ArvancloudBadRequest) Code() int {
	return 400
}

func (o *PostStorageS3ArvancloudBadRequest) Error() string {
	return fmt.Sprintf("[POST /storage/s3/arvancloud][%d] postStorageS3ArvancloudBadRequest  %+v", 400, o.Payload)
}

func (o *PostStorageS3ArvancloudBadRequest) String() string {
	return fmt.Sprintf("[POST /storage/s3/arvancloud][%d] postStorageS3ArvancloudBadRequest  %+v", 400, o.Payload)
}

func (o *PostStorageS3ArvancloudBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostStorageS3ArvancloudBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStorageS3ArvancloudInternalServerError creates a PostStorageS3ArvancloudInternalServerError with default headers values
func NewPostStorageS3ArvancloudInternalServerError() *PostStorageS3ArvancloudInternalServerError {
	return &PostStorageS3ArvancloudInternalServerError{}
}

/*
PostStorageS3ArvancloudInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostStorageS3ArvancloudInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post storage s3 arvancloud internal server error response has a 2xx status code
func (o *PostStorageS3ArvancloudInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post storage s3 arvancloud internal server error response has a 3xx status code
func (o *PostStorageS3ArvancloudInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage s3 arvancloud internal server error response has a 4xx status code
func (o *PostStorageS3ArvancloudInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post storage s3 arvancloud internal server error response has a 5xx status code
func (o *PostStorageS3ArvancloudInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post storage s3 arvancloud internal server error response a status code equal to that given
func (o *PostStorageS3ArvancloudInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post storage s3 arvancloud internal server error response
func (o *PostStorageS3ArvancloudInternalServerError) Code() int {
	return 500
}

func (o *PostStorageS3ArvancloudInternalServerError) Error() string {
	return fmt.Sprintf("[POST /storage/s3/arvancloud][%d] postStorageS3ArvancloudInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStorageS3ArvancloudInternalServerError) String() string {
	return fmt.Sprintf("[POST /storage/s3/arvancloud][%d] postStorageS3ArvancloudInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStorageS3ArvancloudInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostStorageS3ArvancloudInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}