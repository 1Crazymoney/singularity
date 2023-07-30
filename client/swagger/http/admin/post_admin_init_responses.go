// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// PostAdminInitReader is a Reader for the PostAdminInit structure.
type PostAdminInitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAdminInitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostAdminInitNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewPostAdminInitInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /admin/init] PostAdminInit", response, response.Code())
	}
}

// NewPostAdminInitNoContent creates a PostAdminInitNoContent with default headers values
func NewPostAdminInitNoContent() *PostAdminInitNoContent {
	return &PostAdminInitNoContent{}
}

/*
PostAdminInitNoContent describes a response with status code 204, with default header values.

No Content
*/
type PostAdminInitNoContent struct {
}

// IsSuccess returns true when this post admin init no content response has a 2xx status code
func (o *PostAdminInitNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post admin init no content response has a 3xx status code
func (o *PostAdminInitNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post admin init no content response has a 4xx status code
func (o *PostAdminInitNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this post admin init no content response has a 5xx status code
func (o *PostAdminInitNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this post admin init no content response a status code equal to that given
func (o *PostAdminInitNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the post admin init no content response
func (o *PostAdminInitNoContent) Code() int {
	return 204
}

func (o *PostAdminInitNoContent) Error() string {
	return fmt.Sprintf("[POST /admin/init][%d] postAdminInitNoContent ", 204)
}

func (o *PostAdminInitNoContent) String() string {
	return fmt.Sprintf("[POST /admin/init][%d] postAdminInitNoContent ", 204)
}

func (o *PostAdminInitNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAdminInitInternalServerError creates a PostAdminInitInternalServerError with default headers values
func NewPostAdminInitInternalServerError() *PostAdminInitInternalServerError {
	return &PostAdminInitInternalServerError{}
}

/*
PostAdminInitInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostAdminInitInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post admin init internal server error response has a 2xx status code
func (o *PostAdminInitInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post admin init internal server error response has a 3xx status code
func (o *PostAdminInitInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post admin init internal server error response has a 4xx status code
func (o *PostAdminInitInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post admin init internal server error response has a 5xx status code
func (o *PostAdminInitInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post admin init internal server error response a status code equal to that given
func (o *PostAdminInitInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post admin init internal server error response
func (o *PostAdminInitInternalServerError) Code() int {
	return 500
}

func (o *PostAdminInitInternalServerError) Error() string {
	return fmt.Sprintf("[POST /admin/init][%d] postAdminInitInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAdminInitInternalServerError) String() string {
	return fmt.Sprintf("[POST /admin/init][%d] postAdminInitInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAdminInitInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostAdminInitInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}