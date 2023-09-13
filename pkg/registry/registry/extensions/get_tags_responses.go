// Code generated by go-swagger; DO NOT EDIT.

package extensions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/registry/models"
)

// GetTagsReader is a Reader for the GetTags structure.
type GetTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTagsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTagsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTagsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTagsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetTagsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /extensions/tags] getTags", response, response.Code())
	}
}

// NewGetTagsOK creates a GetTagsOK with default headers values
func NewGetTagsOK() *GetTagsOK {
	return &GetTagsOK{}
}

/*
GetTagsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetTagsOK struct {
	Payload []*models.TagCount
}

// IsSuccess returns true when this get tags o k response has a 2xx status code
func (o *GetTagsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get tags o k response has a 3xx status code
func (o *GetTagsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tags o k response has a 4xx status code
func (o *GetTagsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get tags o k response has a 5xx status code
func (o *GetTagsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get tags o k response a status code equal to that given
func (o *GetTagsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get tags o k response
func (o *GetTagsOK) Code() int {
	return 200
}

func (o *GetTagsOK) Error() string {
	return fmt.Sprintf("[GET /extensions/tags][%d] getTagsOK  %+v", 200, o.Payload)
}

func (o *GetTagsOK) String() string {
	return fmt.Sprintf("[GET /extensions/tags][%d] getTagsOK  %+v", 200, o.Payload)
}

func (o *GetTagsOK) GetPayload() []*models.TagCount {
	return o.Payload
}

func (o *GetTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTagsBadRequest creates a GetTagsBadRequest with default headers values
func NewGetTagsBadRequest() *GetTagsBadRequest {
	return &GetTagsBadRequest{}
}

/*
GetTagsBadRequest describes a response with status code 400, with default header values.

NiFi Registry was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetTagsBadRequest struct {
}

// IsSuccess returns true when this get tags bad request response has a 2xx status code
func (o *GetTagsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tags bad request response has a 3xx status code
func (o *GetTagsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tags bad request response has a 4xx status code
func (o *GetTagsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tags bad request response has a 5xx status code
func (o *GetTagsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get tags bad request response a status code equal to that given
func (o *GetTagsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get tags bad request response
func (o *GetTagsBadRequest) Code() int {
	return 400
}

func (o *GetTagsBadRequest) Error() string {
	return fmt.Sprintf("[GET /extensions/tags][%d] getTagsBadRequest ", 400)
}

func (o *GetTagsBadRequest) String() string {
	return fmt.Sprintf("[GET /extensions/tags][%d] getTagsBadRequest ", 400)
}

func (o *GetTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTagsUnauthorized creates a GetTagsUnauthorized with default headers values
func NewGetTagsUnauthorized() *GetTagsUnauthorized {
	return &GetTagsUnauthorized{}
}

/*
GetTagsUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetTagsUnauthorized struct {
}

// IsSuccess returns true when this get tags unauthorized response has a 2xx status code
func (o *GetTagsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tags unauthorized response has a 3xx status code
func (o *GetTagsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tags unauthorized response has a 4xx status code
func (o *GetTagsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tags unauthorized response has a 5xx status code
func (o *GetTagsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get tags unauthorized response a status code equal to that given
func (o *GetTagsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get tags unauthorized response
func (o *GetTagsUnauthorized) Code() int {
	return 401
}

func (o *GetTagsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /extensions/tags][%d] getTagsUnauthorized ", 401)
}

func (o *GetTagsUnauthorized) String() string {
	return fmt.Sprintf("[GET /extensions/tags][%d] getTagsUnauthorized ", 401)
}

func (o *GetTagsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTagsForbidden creates a GetTagsForbidden with default headers values
func NewGetTagsForbidden() *GetTagsForbidden {
	return &GetTagsForbidden{}
}

/*
GetTagsForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetTagsForbidden struct {
}

// IsSuccess returns true when this get tags forbidden response has a 2xx status code
func (o *GetTagsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tags forbidden response has a 3xx status code
func (o *GetTagsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tags forbidden response has a 4xx status code
func (o *GetTagsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tags forbidden response has a 5xx status code
func (o *GetTagsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get tags forbidden response a status code equal to that given
func (o *GetTagsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get tags forbidden response
func (o *GetTagsForbidden) Code() int {
	return 403
}

func (o *GetTagsForbidden) Error() string {
	return fmt.Sprintf("[GET /extensions/tags][%d] getTagsForbidden ", 403)
}

func (o *GetTagsForbidden) String() string {
	return fmt.Sprintf("[GET /extensions/tags][%d] getTagsForbidden ", 403)
}

func (o *GetTagsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTagsNotFound creates a GetTagsNotFound with default headers values
func NewGetTagsNotFound() *GetTagsNotFound {
	return &GetTagsNotFound{}
}

/*
GetTagsNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetTagsNotFound struct {
}

// IsSuccess returns true when this get tags not found response has a 2xx status code
func (o *GetTagsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tags not found response has a 3xx status code
func (o *GetTagsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tags not found response has a 4xx status code
func (o *GetTagsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tags not found response has a 5xx status code
func (o *GetTagsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get tags not found response a status code equal to that given
func (o *GetTagsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get tags not found response
func (o *GetTagsNotFound) Code() int {
	return 404
}

func (o *GetTagsNotFound) Error() string {
	return fmt.Sprintf("[GET /extensions/tags][%d] getTagsNotFound ", 404)
}

func (o *GetTagsNotFound) String() string {
	return fmt.Sprintf("[GET /extensions/tags][%d] getTagsNotFound ", 404)
}

func (o *GetTagsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTagsConflict creates a GetTagsConflict with default headers values
func NewGetTagsConflict() *GetTagsConflict {
	return &GetTagsConflict{}
}

/*
GetTagsConflict describes a response with status code 409, with default header values.

NiFi Registry was unable to complete the request because it assumes a server state that is not valid.
*/
type GetTagsConflict struct {
}

// IsSuccess returns true when this get tags conflict response has a 2xx status code
func (o *GetTagsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tags conflict response has a 3xx status code
func (o *GetTagsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tags conflict response has a 4xx status code
func (o *GetTagsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tags conflict response has a 5xx status code
func (o *GetTagsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get tags conflict response a status code equal to that given
func (o *GetTagsConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the get tags conflict response
func (o *GetTagsConflict) Code() int {
	return 409
}

func (o *GetTagsConflict) Error() string {
	return fmt.Sprintf("[GET /extensions/tags][%d] getTagsConflict ", 409)
}

func (o *GetTagsConflict) String() string {
	return fmt.Sprintf("[GET /extensions/tags][%d] getTagsConflict ", 409)
}

func (o *GetTagsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
