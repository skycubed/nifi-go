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

// CreateDropRequestReader is a Reader for the CreateDropRequest structure.
type CreateDropRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDropRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateDropRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateDropRequestAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateDropRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateDropRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateDropRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateDropRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateDropRequestConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateDropRequestOK creates a CreateDropRequestOK with default headers values
func NewCreateDropRequestOK() *CreateDropRequestOK {
	return &CreateDropRequestOK{}
}

/* CreateDropRequestOK describes a response with status code 200, with default header values.

successful operation
*/
type CreateDropRequestOK struct {
	Payload *models.DropRequestEntity
}

func (o *CreateDropRequestOK) Error() string {
	return fmt.Sprintf("[POST /flowfile-queues/{id}/drop-requests][%d] createDropRequestOK  %+v", 200, o.Payload)
}
func (o *CreateDropRequestOK) GetPayload() *models.DropRequestEntity {
	return o.Payload
}

func (o *CreateDropRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DropRequestEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDropRequestAccepted creates a CreateDropRequestAccepted with default headers values
func NewCreateDropRequestAccepted() *CreateDropRequestAccepted {
	return &CreateDropRequestAccepted{}
}

/* CreateDropRequestAccepted describes a response with status code 202, with default header values.

The request has been accepted. A HTTP response header will contain the URI where the response can be polled.
*/
type CreateDropRequestAccepted struct {
}

func (o *CreateDropRequestAccepted) Error() string {
	return fmt.Sprintf("[POST /flowfile-queues/{id}/drop-requests][%d] createDropRequestAccepted ", 202)
}

func (o *CreateDropRequestAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateDropRequestBadRequest creates a CreateDropRequestBadRequest with default headers values
func NewCreateDropRequestBadRequest() *CreateDropRequestBadRequest {
	return &CreateDropRequestBadRequest{}
}

/* CreateDropRequestBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type CreateDropRequestBadRequest struct {
}

func (o *CreateDropRequestBadRequest) Error() string {
	return fmt.Sprintf("[POST /flowfile-queues/{id}/drop-requests][%d] createDropRequestBadRequest ", 400)
}

func (o *CreateDropRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateDropRequestUnauthorized creates a CreateDropRequestUnauthorized with default headers values
func NewCreateDropRequestUnauthorized() *CreateDropRequestUnauthorized {
	return &CreateDropRequestUnauthorized{}
}

/* CreateDropRequestUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type CreateDropRequestUnauthorized struct {
}

func (o *CreateDropRequestUnauthorized) Error() string {
	return fmt.Sprintf("[POST /flowfile-queues/{id}/drop-requests][%d] createDropRequestUnauthorized ", 401)
}

func (o *CreateDropRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateDropRequestForbidden creates a CreateDropRequestForbidden with default headers values
func NewCreateDropRequestForbidden() *CreateDropRequestForbidden {
	return &CreateDropRequestForbidden{}
}

/* CreateDropRequestForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type CreateDropRequestForbidden struct {
}

func (o *CreateDropRequestForbidden) Error() string {
	return fmt.Sprintf("[POST /flowfile-queues/{id}/drop-requests][%d] createDropRequestForbidden ", 403)
}

func (o *CreateDropRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateDropRequestNotFound creates a CreateDropRequestNotFound with default headers values
func NewCreateDropRequestNotFound() *CreateDropRequestNotFound {
	return &CreateDropRequestNotFound{}
}

/* CreateDropRequestNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type CreateDropRequestNotFound struct {
}

func (o *CreateDropRequestNotFound) Error() string {
	return fmt.Sprintf("[POST /flowfile-queues/{id}/drop-requests][%d] createDropRequestNotFound ", 404)
}

func (o *CreateDropRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateDropRequestConflict creates a CreateDropRequestConflict with default headers values
func NewCreateDropRequestConflict() *CreateDropRequestConflict {
	return &CreateDropRequestConflict{}
}

/* CreateDropRequestConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type CreateDropRequestConflict struct {
}

func (o *CreateDropRequestConflict) Error() string {
	return fmt.Sprintf("[POST /flowfile-queues/{id}/drop-requests][%d] createDropRequestConflict ", 409)
}

func (o *CreateDropRequestConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
