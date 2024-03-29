// Code generated by go-swagger; DO NOT EDIT.

package flowfile_queues

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// GetDropRequestReader is a Reader for the GetDropRequest structure.
type GetDropRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDropRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDropRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDropRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDropRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDropRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDropRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetDropRequestConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /flowfile-queues/{id}/drop-requests/{drop-request-id}] getDropRequest", response, response.Code())
	}
}

// NewGetDropRequestOK creates a GetDropRequestOK with default headers values
func NewGetDropRequestOK() *GetDropRequestOK {
	return &GetDropRequestOK{}
}

/*
GetDropRequestOK describes a response with status code 200, with default header values.

successful operation
*/
type GetDropRequestOK struct {
	Payload *models.DropRequestEntity
}

// IsSuccess returns true when this get drop request o k response has a 2xx status code
func (o *GetDropRequestOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get drop request o k response has a 3xx status code
func (o *GetDropRequestOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get drop request o k response has a 4xx status code
func (o *GetDropRequestOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get drop request o k response has a 5xx status code
func (o *GetDropRequestOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get drop request o k response a status code equal to that given
func (o *GetDropRequestOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get drop request o k response
func (o *GetDropRequestOK) Code() int {
	return 200
}

func (o *GetDropRequestOK) Error() string {
	return fmt.Sprintf("[GET /flowfile-queues/{id}/drop-requests/{drop-request-id}][%d] getDropRequestOK  %+v", 200, o.Payload)
}

func (o *GetDropRequestOK) String() string {
	return fmt.Sprintf("[GET /flowfile-queues/{id}/drop-requests/{drop-request-id}][%d] getDropRequestOK  %+v", 200, o.Payload)
}

func (o *GetDropRequestOK) GetPayload() *models.DropRequestEntity {
	return o.Payload
}

func (o *GetDropRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DropRequestEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDropRequestBadRequest creates a GetDropRequestBadRequest with default headers values
func NewGetDropRequestBadRequest() *GetDropRequestBadRequest {
	return &GetDropRequestBadRequest{}
}

/*
GetDropRequestBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetDropRequestBadRequest struct {
}

// IsSuccess returns true when this get drop request bad request response has a 2xx status code
func (o *GetDropRequestBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get drop request bad request response has a 3xx status code
func (o *GetDropRequestBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get drop request bad request response has a 4xx status code
func (o *GetDropRequestBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get drop request bad request response has a 5xx status code
func (o *GetDropRequestBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get drop request bad request response a status code equal to that given
func (o *GetDropRequestBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get drop request bad request response
func (o *GetDropRequestBadRequest) Code() int {
	return 400
}

func (o *GetDropRequestBadRequest) Error() string {
	return fmt.Sprintf("[GET /flowfile-queues/{id}/drop-requests/{drop-request-id}][%d] getDropRequestBadRequest ", 400)
}

func (o *GetDropRequestBadRequest) String() string {
	return fmt.Sprintf("[GET /flowfile-queues/{id}/drop-requests/{drop-request-id}][%d] getDropRequestBadRequest ", 400)
}

func (o *GetDropRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDropRequestUnauthorized creates a GetDropRequestUnauthorized with default headers values
func NewGetDropRequestUnauthorized() *GetDropRequestUnauthorized {
	return &GetDropRequestUnauthorized{}
}

/*
GetDropRequestUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetDropRequestUnauthorized struct {
}

// IsSuccess returns true when this get drop request unauthorized response has a 2xx status code
func (o *GetDropRequestUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get drop request unauthorized response has a 3xx status code
func (o *GetDropRequestUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get drop request unauthorized response has a 4xx status code
func (o *GetDropRequestUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get drop request unauthorized response has a 5xx status code
func (o *GetDropRequestUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get drop request unauthorized response a status code equal to that given
func (o *GetDropRequestUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get drop request unauthorized response
func (o *GetDropRequestUnauthorized) Code() int {
	return 401
}

func (o *GetDropRequestUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flowfile-queues/{id}/drop-requests/{drop-request-id}][%d] getDropRequestUnauthorized ", 401)
}

func (o *GetDropRequestUnauthorized) String() string {
	return fmt.Sprintf("[GET /flowfile-queues/{id}/drop-requests/{drop-request-id}][%d] getDropRequestUnauthorized ", 401)
}

func (o *GetDropRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDropRequestForbidden creates a GetDropRequestForbidden with default headers values
func NewGetDropRequestForbidden() *GetDropRequestForbidden {
	return &GetDropRequestForbidden{}
}

/*
GetDropRequestForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetDropRequestForbidden struct {
}

// IsSuccess returns true when this get drop request forbidden response has a 2xx status code
func (o *GetDropRequestForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get drop request forbidden response has a 3xx status code
func (o *GetDropRequestForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get drop request forbidden response has a 4xx status code
func (o *GetDropRequestForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get drop request forbidden response has a 5xx status code
func (o *GetDropRequestForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get drop request forbidden response a status code equal to that given
func (o *GetDropRequestForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get drop request forbidden response
func (o *GetDropRequestForbidden) Code() int {
	return 403
}

func (o *GetDropRequestForbidden) Error() string {
	return fmt.Sprintf("[GET /flowfile-queues/{id}/drop-requests/{drop-request-id}][%d] getDropRequestForbidden ", 403)
}

func (o *GetDropRequestForbidden) String() string {
	return fmt.Sprintf("[GET /flowfile-queues/{id}/drop-requests/{drop-request-id}][%d] getDropRequestForbidden ", 403)
}

func (o *GetDropRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDropRequestNotFound creates a GetDropRequestNotFound with default headers values
func NewGetDropRequestNotFound() *GetDropRequestNotFound {
	return &GetDropRequestNotFound{}
}

/*
GetDropRequestNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetDropRequestNotFound struct {
}

// IsSuccess returns true when this get drop request not found response has a 2xx status code
func (o *GetDropRequestNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get drop request not found response has a 3xx status code
func (o *GetDropRequestNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get drop request not found response has a 4xx status code
func (o *GetDropRequestNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get drop request not found response has a 5xx status code
func (o *GetDropRequestNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get drop request not found response a status code equal to that given
func (o *GetDropRequestNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get drop request not found response
func (o *GetDropRequestNotFound) Code() int {
	return 404
}

func (o *GetDropRequestNotFound) Error() string {
	return fmt.Sprintf("[GET /flowfile-queues/{id}/drop-requests/{drop-request-id}][%d] getDropRequestNotFound ", 404)
}

func (o *GetDropRequestNotFound) String() string {
	return fmt.Sprintf("[GET /flowfile-queues/{id}/drop-requests/{drop-request-id}][%d] getDropRequestNotFound ", 404)
}

func (o *GetDropRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDropRequestConflict creates a GetDropRequestConflict with default headers values
func NewGetDropRequestConflict() *GetDropRequestConflict {
	return &GetDropRequestConflict{}
}

/*
GetDropRequestConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetDropRequestConflict struct {
}

// IsSuccess returns true when this get drop request conflict response has a 2xx status code
func (o *GetDropRequestConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get drop request conflict response has a 3xx status code
func (o *GetDropRequestConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get drop request conflict response has a 4xx status code
func (o *GetDropRequestConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get drop request conflict response has a 5xx status code
func (o *GetDropRequestConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get drop request conflict response a status code equal to that given
func (o *GetDropRequestConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the get drop request conflict response
func (o *GetDropRequestConflict) Code() int {
	return 409
}

func (o *GetDropRequestConflict) Error() string {
	return fmt.Sprintf("[GET /flowfile-queues/{id}/drop-requests/{drop-request-id}][%d] getDropRequestConflict ", 409)
}

func (o *GetDropRequestConflict) String() string {
	return fmt.Sprintf("[GET /flowfile-queues/{id}/drop-requests/{drop-request-id}][%d] getDropRequestConflict ", 409)
}

func (o *GetDropRequestConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
