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

// CreateFlowFileListingReader is a Reader for the CreateFlowFileListing structure.
type CreateFlowFileListingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateFlowFileListingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateFlowFileListingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateFlowFileListingAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateFlowFileListingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateFlowFileListingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateFlowFileListingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateFlowFileListingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateFlowFileListingConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateFlowFileListingOK creates a CreateFlowFileListingOK with default headers values
func NewCreateFlowFileListingOK() *CreateFlowFileListingOK {
	return &CreateFlowFileListingOK{}
}

/*
CreateFlowFileListingOK describes a response with status code 200, with default header values.

successful operation
*/
type CreateFlowFileListingOK struct {
	Payload *models.ListingRequestEntity
}

// IsSuccess returns true when this create flow file listing o k response has a 2xx status code
func (o *CreateFlowFileListingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create flow file listing o k response has a 3xx status code
func (o *CreateFlowFileListingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create flow file listing o k response has a 4xx status code
func (o *CreateFlowFileListingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create flow file listing o k response has a 5xx status code
func (o *CreateFlowFileListingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create flow file listing o k response a status code equal to that given
func (o *CreateFlowFileListingOK) IsCode(code int) bool {
	return code == 200
}

func (o *CreateFlowFileListingOK) Error() string {
	return fmt.Sprintf("[POST /flowfile-queues/{id}/listing-requests][%d] createFlowFileListingOK  %+v", 200, o.Payload)
}

func (o *CreateFlowFileListingOK) String() string {
	return fmt.Sprintf("[POST /flowfile-queues/{id}/listing-requests][%d] createFlowFileListingOK  %+v", 200, o.Payload)
}

func (o *CreateFlowFileListingOK) GetPayload() *models.ListingRequestEntity {
	return o.Payload
}

func (o *CreateFlowFileListingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListingRequestEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFlowFileListingAccepted creates a CreateFlowFileListingAccepted with default headers values
func NewCreateFlowFileListingAccepted() *CreateFlowFileListingAccepted {
	return &CreateFlowFileListingAccepted{}
}

/*
CreateFlowFileListingAccepted describes a response with status code 202, with default header values.

The request has been accepted. A HTTP response header will contain the URI where the response can be polled.
*/
type CreateFlowFileListingAccepted struct {
}

// IsSuccess returns true when this create flow file listing accepted response has a 2xx status code
func (o *CreateFlowFileListingAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create flow file listing accepted response has a 3xx status code
func (o *CreateFlowFileListingAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create flow file listing accepted response has a 4xx status code
func (o *CreateFlowFileListingAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this create flow file listing accepted response has a 5xx status code
func (o *CreateFlowFileListingAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this create flow file listing accepted response a status code equal to that given
func (o *CreateFlowFileListingAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *CreateFlowFileListingAccepted) Error() string {
	return fmt.Sprintf("[POST /flowfile-queues/{id}/listing-requests][%d] createFlowFileListingAccepted ", 202)
}

func (o *CreateFlowFileListingAccepted) String() string {
	return fmt.Sprintf("[POST /flowfile-queues/{id}/listing-requests][%d] createFlowFileListingAccepted ", 202)
}

func (o *CreateFlowFileListingAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateFlowFileListingBadRequest creates a CreateFlowFileListingBadRequest with default headers values
func NewCreateFlowFileListingBadRequest() *CreateFlowFileListingBadRequest {
	return &CreateFlowFileListingBadRequest{}
}

/*
CreateFlowFileListingBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type CreateFlowFileListingBadRequest struct {
}

// IsSuccess returns true when this create flow file listing bad request response has a 2xx status code
func (o *CreateFlowFileListingBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create flow file listing bad request response has a 3xx status code
func (o *CreateFlowFileListingBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create flow file listing bad request response has a 4xx status code
func (o *CreateFlowFileListingBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create flow file listing bad request response has a 5xx status code
func (o *CreateFlowFileListingBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create flow file listing bad request response a status code equal to that given
func (o *CreateFlowFileListingBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateFlowFileListingBadRequest) Error() string {
	return fmt.Sprintf("[POST /flowfile-queues/{id}/listing-requests][%d] createFlowFileListingBadRequest ", 400)
}

func (o *CreateFlowFileListingBadRequest) String() string {
	return fmt.Sprintf("[POST /flowfile-queues/{id}/listing-requests][%d] createFlowFileListingBadRequest ", 400)
}

func (o *CreateFlowFileListingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateFlowFileListingUnauthorized creates a CreateFlowFileListingUnauthorized with default headers values
func NewCreateFlowFileListingUnauthorized() *CreateFlowFileListingUnauthorized {
	return &CreateFlowFileListingUnauthorized{}
}

/*
CreateFlowFileListingUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type CreateFlowFileListingUnauthorized struct {
}

// IsSuccess returns true when this create flow file listing unauthorized response has a 2xx status code
func (o *CreateFlowFileListingUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create flow file listing unauthorized response has a 3xx status code
func (o *CreateFlowFileListingUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create flow file listing unauthorized response has a 4xx status code
func (o *CreateFlowFileListingUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create flow file listing unauthorized response has a 5xx status code
func (o *CreateFlowFileListingUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create flow file listing unauthorized response a status code equal to that given
func (o *CreateFlowFileListingUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateFlowFileListingUnauthorized) Error() string {
	return fmt.Sprintf("[POST /flowfile-queues/{id}/listing-requests][%d] createFlowFileListingUnauthorized ", 401)
}

func (o *CreateFlowFileListingUnauthorized) String() string {
	return fmt.Sprintf("[POST /flowfile-queues/{id}/listing-requests][%d] createFlowFileListingUnauthorized ", 401)
}

func (o *CreateFlowFileListingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateFlowFileListingForbidden creates a CreateFlowFileListingForbidden with default headers values
func NewCreateFlowFileListingForbidden() *CreateFlowFileListingForbidden {
	return &CreateFlowFileListingForbidden{}
}

/*
CreateFlowFileListingForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type CreateFlowFileListingForbidden struct {
}

// IsSuccess returns true when this create flow file listing forbidden response has a 2xx status code
func (o *CreateFlowFileListingForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create flow file listing forbidden response has a 3xx status code
func (o *CreateFlowFileListingForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create flow file listing forbidden response has a 4xx status code
func (o *CreateFlowFileListingForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create flow file listing forbidden response has a 5xx status code
func (o *CreateFlowFileListingForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create flow file listing forbidden response a status code equal to that given
func (o *CreateFlowFileListingForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateFlowFileListingForbidden) Error() string {
	return fmt.Sprintf("[POST /flowfile-queues/{id}/listing-requests][%d] createFlowFileListingForbidden ", 403)
}

func (o *CreateFlowFileListingForbidden) String() string {
	return fmt.Sprintf("[POST /flowfile-queues/{id}/listing-requests][%d] createFlowFileListingForbidden ", 403)
}

func (o *CreateFlowFileListingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateFlowFileListingNotFound creates a CreateFlowFileListingNotFound with default headers values
func NewCreateFlowFileListingNotFound() *CreateFlowFileListingNotFound {
	return &CreateFlowFileListingNotFound{}
}

/*
CreateFlowFileListingNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type CreateFlowFileListingNotFound struct {
}

// IsSuccess returns true when this create flow file listing not found response has a 2xx status code
func (o *CreateFlowFileListingNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create flow file listing not found response has a 3xx status code
func (o *CreateFlowFileListingNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create flow file listing not found response has a 4xx status code
func (o *CreateFlowFileListingNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create flow file listing not found response has a 5xx status code
func (o *CreateFlowFileListingNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create flow file listing not found response a status code equal to that given
func (o *CreateFlowFileListingNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CreateFlowFileListingNotFound) Error() string {
	return fmt.Sprintf("[POST /flowfile-queues/{id}/listing-requests][%d] createFlowFileListingNotFound ", 404)
}

func (o *CreateFlowFileListingNotFound) String() string {
	return fmt.Sprintf("[POST /flowfile-queues/{id}/listing-requests][%d] createFlowFileListingNotFound ", 404)
}

func (o *CreateFlowFileListingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateFlowFileListingConflict creates a CreateFlowFileListingConflict with default headers values
func NewCreateFlowFileListingConflict() *CreateFlowFileListingConflict {
	return &CreateFlowFileListingConflict{}
}

/*
CreateFlowFileListingConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type CreateFlowFileListingConflict struct {
}

// IsSuccess returns true when this create flow file listing conflict response has a 2xx status code
func (o *CreateFlowFileListingConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create flow file listing conflict response has a 3xx status code
func (o *CreateFlowFileListingConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create flow file listing conflict response has a 4xx status code
func (o *CreateFlowFileListingConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create flow file listing conflict response has a 5xx status code
func (o *CreateFlowFileListingConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create flow file listing conflict response a status code equal to that given
func (o *CreateFlowFileListingConflict) IsCode(code int) bool {
	return code == 409
}

func (o *CreateFlowFileListingConflict) Error() string {
	return fmt.Sprintf("[POST /flowfile-queues/{id}/listing-requests][%d] createFlowFileListingConflict ", 409)
}

func (o *CreateFlowFileListingConflict) String() string {
	return fmt.Sprintf("[POST /flowfile-queues/{id}/listing-requests][%d] createFlowFileListingConflict ", 409)
}

func (o *CreateFlowFileListingConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
