// Code generated by go-swagger; DO NOT EDIT.

package process_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// CreateConnectionReader is a Reader for the CreateConnection structure.
type CreateConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateConnectionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateConnectionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateConnectionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateConnectionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateConnectionCreated creates a CreateConnectionCreated with default headers values
func NewCreateConnectionCreated() *CreateConnectionCreated {
	return &CreateConnectionCreated{}
}

/* CreateConnectionCreated describes a response with status code 201, with default header values.

successful operation
*/
type CreateConnectionCreated struct {
	Payload *models.ConnectionEntity
}

func (o *CreateConnectionCreated) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/connections][%d] createConnectionCreated  %+v", 201, o.Payload)
}
func (o *CreateConnectionCreated) GetPayload() *models.ConnectionEntity {
	return o.Payload
}

func (o *CreateConnectionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConnectionEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConnectionBadRequest creates a CreateConnectionBadRequest with default headers values
func NewCreateConnectionBadRequest() *CreateConnectionBadRequest {
	return &CreateConnectionBadRequest{}
}

/* CreateConnectionBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type CreateConnectionBadRequest struct {
}

func (o *CreateConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/connections][%d] createConnectionBadRequest ", 400)
}

func (o *CreateConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateConnectionUnauthorized creates a CreateConnectionUnauthorized with default headers values
func NewCreateConnectionUnauthorized() *CreateConnectionUnauthorized {
	return &CreateConnectionUnauthorized{}
}

/* CreateConnectionUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type CreateConnectionUnauthorized struct {
}

func (o *CreateConnectionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/connections][%d] createConnectionUnauthorized ", 401)
}

func (o *CreateConnectionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateConnectionForbidden creates a CreateConnectionForbidden with default headers values
func NewCreateConnectionForbidden() *CreateConnectionForbidden {
	return &CreateConnectionForbidden{}
}

/* CreateConnectionForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type CreateConnectionForbidden struct {
}

func (o *CreateConnectionForbidden) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/connections][%d] createConnectionForbidden ", 403)
}

func (o *CreateConnectionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateConnectionNotFound creates a CreateConnectionNotFound with default headers values
func NewCreateConnectionNotFound() *CreateConnectionNotFound {
	return &CreateConnectionNotFound{}
}

/* CreateConnectionNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type CreateConnectionNotFound struct {
}

func (o *CreateConnectionNotFound) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/connections][%d] createConnectionNotFound ", 404)
}

func (o *CreateConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateConnectionConflict creates a CreateConnectionConflict with default headers values
func NewCreateConnectionConflict() *CreateConnectionConflict {
	return &CreateConnectionConflict{}
}

/* CreateConnectionConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type CreateConnectionConflict struct {
}

func (o *CreateConnectionConflict) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/connections][%d] createConnectionConflict ", 409)
}

func (o *CreateConnectionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
