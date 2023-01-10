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

// GetListingRequestReader is a Reader for the GetListingRequest structure.
type GetListingRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetListingRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetListingRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetListingRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetListingRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetListingRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetListingRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetListingRequestConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetListingRequestOK creates a GetListingRequestOK with default headers values
func NewGetListingRequestOK() *GetListingRequestOK {
	return &GetListingRequestOK{}
}

/*
GetListingRequestOK describes a response with status code 200, with default header values.

successful operation
*/
type GetListingRequestOK struct {
	Payload *models.ListingRequestEntity
}

// IsSuccess returns true when this get listing request o k response has a 2xx status code
func (o *GetListingRequestOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get listing request o k response has a 3xx status code
func (o *GetListingRequestOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get listing request o k response has a 4xx status code
func (o *GetListingRequestOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get listing request o k response has a 5xx status code
func (o *GetListingRequestOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get listing request o k response a status code equal to that given
func (o *GetListingRequestOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetListingRequestOK) Error() string {
	return fmt.Sprintf("[GET /flowfile-queues/{id}/listing-requests/{listing-request-id}][%d] getListingRequestOK  %+v", 200, o.Payload)
}

func (o *GetListingRequestOK) String() string {
	return fmt.Sprintf("[GET /flowfile-queues/{id}/listing-requests/{listing-request-id}][%d] getListingRequestOK  %+v", 200, o.Payload)
}

func (o *GetListingRequestOK) GetPayload() *models.ListingRequestEntity {
	return o.Payload
}

func (o *GetListingRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListingRequestEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetListingRequestBadRequest creates a GetListingRequestBadRequest with default headers values
func NewGetListingRequestBadRequest() *GetListingRequestBadRequest {
	return &GetListingRequestBadRequest{}
}

/*
GetListingRequestBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetListingRequestBadRequest struct {
}

// IsSuccess returns true when this get listing request bad request response has a 2xx status code
func (o *GetListingRequestBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get listing request bad request response has a 3xx status code
func (o *GetListingRequestBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get listing request bad request response has a 4xx status code
func (o *GetListingRequestBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get listing request bad request response has a 5xx status code
func (o *GetListingRequestBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get listing request bad request response a status code equal to that given
func (o *GetListingRequestBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetListingRequestBadRequest) Error() string {
	return fmt.Sprintf("[GET /flowfile-queues/{id}/listing-requests/{listing-request-id}][%d] getListingRequestBadRequest ", 400)
}

func (o *GetListingRequestBadRequest) String() string {
	return fmt.Sprintf("[GET /flowfile-queues/{id}/listing-requests/{listing-request-id}][%d] getListingRequestBadRequest ", 400)
}

func (o *GetListingRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListingRequestUnauthorized creates a GetListingRequestUnauthorized with default headers values
func NewGetListingRequestUnauthorized() *GetListingRequestUnauthorized {
	return &GetListingRequestUnauthorized{}
}

/*
GetListingRequestUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetListingRequestUnauthorized struct {
}

// IsSuccess returns true when this get listing request unauthorized response has a 2xx status code
func (o *GetListingRequestUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get listing request unauthorized response has a 3xx status code
func (o *GetListingRequestUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get listing request unauthorized response has a 4xx status code
func (o *GetListingRequestUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get listing request unauthorized response has a 5xx status code
func (o *GetListingRequestUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get listing request unauthorized response a status code equal to that given
func (o *GetListingRequestUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetListingRequestUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flowfile-queues/{id}/listing-requests/{listing-request-id}][%d] getListingRequestUnauthorized ", 401)
}

func (o *GetListingRequestUnauthorized) String() string {
	return fmt.Sprintf("[GET /flowfile-queues/{id}/listing-requests/{listing-request-id}][%d] getListingRequestUnauthorized ", 401)
}

func (o *GetListingRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListingRequestForbidden creates a GetListingRequestForbidden with default headers values
func NewGetListingRequestForbidden() *GetListingRequestForbidden {
	return &GetListingRequestForbidden{}
}

/*
GetListingRequestForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetListingRequestForbidden struct {
}

// IsSuccess returns true when this get listing request forbidden response has a 2xx status code
func (o *GetListingRequestForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get listing request forbidden response has a 3xx status code
func (o *GetListingRequestForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get listing request forbidden response has a 4xx status code
func (o *GetListingRequestForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get listing request forbidden response has a 5xx status code
func (o *GetListingRequestForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get listing request forbidden response a status code equal to that given
func (o *GetListingRequestForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetListingRequestForbidden) Error() string {
	return fmt.Sprintf("[GET /flowfile-queues/{id}/listing-requests/{listing-request-id}][%d] getListingRequestForbidden ", 403)
}

func (o *GetListingRequestForbidden) String() string {
	return fmt.Sprintf("[GET /flowfile-queues/{id}/listing-requests/{listing-request-id}][%d] getListingRequestForbidden ", 403)
}

func (o *GetListingRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListingRequestNotFound creates a GetListingRequestNotFound with default headers values
func NewGetListingRequestNotFound() *GetListingRequestNotFound {
	return &GetListingRequestNotFound{}
}

/*
GetListingRequestNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetListingRequestNotFound struct {
}

// IsSuccess returns true when this get listing request not found response has a 2xx status code
func (o *GetListingRequestNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get listing request not found response has a 3xx status code
func (o *GetListingRequestNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get listing request not found response has a 4xx status code
func (o *GetListingRequestNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get listing request not found response has a 5xx status code
func (o *GetListingRequestNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get listing request not found response a status code equal to that given
func (o *GetListingRequestNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetListingRequestNotFound) Error() string {
	return fmt.Sprintf("[GET /flowfile-queues/{id}/listing-requests/{listing-request-id}][%d] getListingRequestNotFound ", 404)
}

func (o *GetListingRequestNotFound) String() string {
	return fmt.Sprintf("[GET /flowfile-queues/{id}/listing-requests/{listing-request-id}][%d] getListingRequestNotFound ", 404)
}

func (o *GetListingRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListingRequestConflict creates a GetListingRequestConflict with default headers values
func NewGetListingRequestConflict() *GetListingRequestConflict {
	return &GetListingRequestConflict{}
}

/*
GetListingRequestConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetListingRequestConflict struct {
}

// IsSuccess returns true when this get listing request conflict response has a 2xx status code
func (o *GetListingRequestConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get listing request conflict response has a 3xx status code
func (o *GetListingRequestConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get listing request conflict response has a 4xx status code
func (o *GetListingRequestConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get listing request conflict response has a 5xx status code
func (o *GetListingRequestConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get listing request conflict response a status code equal to that given
func (o *GetListingRequestConflict) IsCode(code int) bool {
	return code == 409
}

func (o *GetListingRequestConflict) Error() string {
	return fmt.Sprintf("[GET /flowfile-queues/{id}/listing-requests/{listing-request-id}][%d] getListingRequestConflict ", 409)
}

func (o *GetListingRequestConflict) String() string {
	return fmt.Sprintf("[GET /flowfile-queues/{id}/listing-requests/{listing-request-id}][%d] getListingRequestConflict ", 409)
}

func (o *GetListingRequestConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
