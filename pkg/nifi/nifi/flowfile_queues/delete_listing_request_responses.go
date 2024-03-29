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

// DeleteListingRequestReader is a Reader for the DeleteListingRequest structure.
type DeleteListingRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteListingRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteListingRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteListingRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteListingRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteListingRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteListingRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteListingRequestConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /flowfile-queues/{id}/listing-requests/{listing-request-id}] deleteListingRequest", response, response.Code())
	}
}

// NewDeleteListingRequestOK creates a DeleteListingRequestOK with default headers values
func NewDeleteListingRequestOK() *DeleteListingRequestOK {
	return &DeleteListingRequestOK{}
}

/*
DeleteListingRequestOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteListingRequestOK struct {
	Payload *models.ListingRequestEntity
}

// IsSuccess returns true when this delete listing request o k response has a 2xx status code
func (o *DeleteListingRequestOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete listing request o k response has a 3xx status code
func (o *DeleteListingRequestOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete listing request o k response has a 4xx status code
func (o *DeleteListingRequestOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete listing request o k response has a 5xx status code
func (o *DeleteListingRequestOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete listing request o k response a status code equal to that given
func (o *DeleteListingRequestOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete listing request o k response
func (o *DeleteListingRequestOK) Code() int {
	return 200
}

func (o *DeleteListingRequestOK) Error() string {
	return fmt.Sprintf("[DELETE /flowfile-queues/{id}/listing-requests/{listing-request-id}][%d] deleteListingRequestOK  %+v", 200, o.Payload)
}

func (o *DeleteListingRequestOK) String() string {
	return fmt.Sprintf("[DELETE /flowfile-queues/{id}/listing-requests/{listing-request-id}][%d] deleteListingRequestOK  %+v", 200, o.Payload)
}

func (o *DeleteListingRequestOK) GetPayload() *models.ListingRequestEntity {
	return o.Payload
}

func (o *DeleteListingRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListingRequestEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteListingRequestBadRequest creates a DeleteListingRequestBadRequest with default headers values
func NewDeleteListingRequestBadRequest() *DeleteListingRequestBadRequest {
	return &DeleteListingRequestBadRequest{}
}

/*
DeleteListingRequestBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type DeleteListingRequestBadRequest struct {
}

// IsSuccess returns true when this delete listing request bad request response has a 2xx status code
func (o *DeleteListingRequestBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete listing request bad request response has a 3xx status code
func (o *DeleteListingRequestBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete listing request bad request response has a 4xx status code
func (o *DeleteListingRequestBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete listing request bad request response has a 5xx status code
func (o *DeleteListingRequestBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete listing request bad request response a status code equal to that given
func (o *DeleteListingRequestBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete listing request bad request response
func (o *DeleteListingRequestBadRequest) Code() int {
	return 400
}

func (o *DeleteListingRequestBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /flowfile-queues/{id}/listing-requests/{listing-request-id}][%d] deleteListingRequestBadRequest ", 400)
}

func (o *DeleteListingRequestBadRequest) String() string {
	return fmt.Sprintf("[DELETE /flowfile-queues/{id}/listing-requests/{listing-request-id}][%d] deleteListingRequestBadRequest ", 400)
}

func (o *DeleteListingRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteListingRequestUnauthorized creates a DeleteListingRequestUnauthorized with default headers values
func NewDeleteListingRequestUnauthorized() *DeleteListingRequestUnauthorized {
	return &DeleteListingRequestUnauthorized{}
}

/*
DeleteListingRequestUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type DeleteListingRequestUnauthorized struct {
}

// IsSuccess returns true when this delete listing request unauthorized response has a 2xx status code
func (o *DeleteListingRequestUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete listing request unauthorized response has a 3xx status code
func (o *DeleteListingRequestUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete listing request unauthorized response has a 4xx status code
func (o *DeleteListingRequestUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete listing request unauthorized response has a 5xx status code
func (o *DeleteListingRequestUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete listing request unauthorized response a status code equal to that given
func (o *DeleteListingRequestUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete listing request unauthorized response
func (o *DeleteListingRequestUnauthorized) Code() int {
	return 401
}

func (o *DeleteListingRequestUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /flowfile-queues/{id}/listing-requests/{listing-request-id}][%d] deleteListingRequestUnauthorized ", 401)
}

func (o *DeleteListingRequestUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /flowfile-queues/{id}/listing-requests/{listing-request-id}][%d] deleteListingRequestUnauthorized ", 401)
}

func (o *DeleteListingRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteListingRequestForbidden creates a DeleteListingRequestForbidden with default headers values
func NewDeleteListingRequestForbidden() *DeleteListingRequestForbidden {
	return &DeleteListingRequestForbidden{}
}

/*
DeleteListingRequestForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type DeleteListingRequestForbidden struct {
}

// IsSuccess returns true when this delete listing request forbidden response has a 2xx status code
func (o *DeleteListingRequestForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete listing request forbidden response has a 3xx status code
func (o *DeleteListingRequestForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete listing request forbidden response has a 4xx status code
func (o *DeleteListingRequestForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete listing request forbidden response has a 5xx status code
func (o *DeleteListingRequestForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete listing request forbidden response a status code equal to that given
func (o *DeleteListingRequestForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete listing request forbidden response
func (o *DeleteListingRequestForbidden) Code() int {
	return 403
}

func (o *DeleteListingRequestForbidden) Error() string {
	return fmt.Sprintf("[DELETE /flowfile-queues/{id}/listing-requests/{listing-request-id}][%d] deleteListingRequestForbidden ", 403)
}

func (o *DeleteListingRequestForbidden) String() string {
	return fmt.Sprintf("[DELETE /flowfile-queues/{id}/listing-requests/{listing-request-id}][%d] deleteListingRequestForbidden ", 403)
}

func (o *DeleteListingRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteListingRequestNotFound creates a DeleteListingRequestNotFound with default headers values
func NewDeleteListingRequestNotFound() *DeleteListingRequestNotFound {
	return &DeleteListingRequestNotFound{}
}

/*
DeleteListingRequestNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type DeleteListingRequestNotFound struct {
}

// IsSuccess returns true when this delete listing request not found response has a 2xx status code
func (o *DeleteListingRequestNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete listing request not found response has a 3xx status code
func (o *DeleteListingRequestNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete listing request not found response has a 4xx status code
func (o *DeleteListingRequestNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete listing request not found response has a 5xx status code
func (o *DeleteListingRequestNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete listing request not found response a status code equal to that given
func (o *DeleteListingRequestNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete listing request not found response
func (o *DeleteListingRequestNotFound) Code() int {
	return 404
}

func (o *DeleteListingRequestNotFound) Error() string {
	return fmt.Sprintf("[DELETE /flowfile-queues/{id}/listing-requests/{listing-request-id}][%d] deleteListingRequestNotFound ", 404)
}

func (o *DeleteListingRequestNotFound) String() string {
	return fmt.Sprintf("[DELETE /flowfile-queues/{id}/listing-requests/{listing-request-id}][%d] deleteListingRequestNotFound ", 404)
}

func (o *DeleteListingRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteListingRequestConflict creates a DeleteListingRequestConflict with default headers values
func NewDeleteListingRequestConflict() *DeleteListingRequestConflict {
	return &DeleteListingRequestConflict{}
}

/*
DeleteListingRequestConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type DeleteListingRequestConflict struct {
}

// IsSuccess returns true when this delete listing request conflict response has a 2xx status code
func (o *DeleteListingRequestConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete listing request conflict response has a 3xx status code
func (o *DeleteListingRequestConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete listing request conflict response has a 4xx status code
func (o *DeleteListingRequestConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete listing request conflict response has a 5xx status code
func (o *DeleteListingRequestConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this delete listing request conflict response a status code equal to that given
func (o *DeleteListingRequestConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the delete listing request conflict response
func (o *DeleteListingRequestConflict) Code() int {
	return 409
}

func (o *DeleteListingRequestConflict) Error() string {
	return fmt.Sprintf("[DELETE /flowfile-queues/{id}/listing-requests/{listing-request-id}][%d] deleteListingRequestConflict ", 409)
}

func (o *DeleteListingRequestConflict) String() string {
	return fmt.Sprintf("[DELETE /flowfile-queues/{id}/listing-requests/{listing-request-id}][%d] deleteListingRequestConflict ", 409)
}

func (o *DeleteListingRequestConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
