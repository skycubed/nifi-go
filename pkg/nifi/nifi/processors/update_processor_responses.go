// Code generated by go-swagger; DO NOT EDIT.

package processors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// UpdateProcessorReader is a Reader for the UpdateProcessor structure.
type UpdateProcessorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProcessorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateProcessorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateProcessorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateProcessorUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateProcessorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateProcessorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateProcessorConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateProcessorOK creates a UpdateProcessorOK with default headers values
func NewUpdateProcessorOK() *UpdateProcessorOK {
	return &UpdateProcessorOK{}
}

/* UpdateProcessorOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateProcessorOK struct {
	Payload *models.ProcessorEntity
}

func (o *UpdateProcessorOK) Error() string {
	return fmt.Sprintf("[PUT /processors/{id}][%d] updateProcessorOK  %+v", 200, o.Payload)
}
func (o *UpdateProcessorOK) GetPayload() *models.ProcessorEntity {
	return o.Payload
}

func (o *UpdateProcessorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProcessorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProcessorBadRequest creates a UpdateProcessorBadRequest with default headers values
func NewUpdateProcessorBadRequest() *UpdateProcessorBadRequest {
	return &UpdateProcessorBadRequest{}
}

/* UpdateProcessorBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type UpdateProcessorBadRequest struct {
}

func (o *UpdateProcessorBadRequest) Error() string {
	return fmt.Sprintf("[PUT /processors/{id}][%d] updateProcessorBadRequest ", 400)
}

func (o *UpdateProcessorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateProcessorUnauthorized creates a UpdateProcessorUnauthorized with default headers values
func NewUpdateProcessorUnauthorized() *UpdateProcessorUnauthorized {
	return &UpdateProcessorUnauthorized{}
}

/* UpdateProcessorUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type UpdateProcessorUnauthorized struct {
}

func (o *UpdateProcessorUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /processors/{id}][%d] updateProcessorUnauthorized ", 401)
}

func (o *UpdateProcessorUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateProcessorForbidden creates a UpdateProcessorForbidden with default headers values
func NewUpdateProcessorForbidden() *UpdateProcessorForbidden {
	return &UpdateProcessorForbidden{}
}

/* UpdateProcessorForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type UpdateProcessorForbidden struct {
}

func (o *UpdateProcessorForbidden) Error() string {
	return fmt.Sprintf("[PUT /processors/{id}][%d] updateProcessorForbidden ", 403)
}

func (o *UpdateProcessorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateProcessorNotFound creates a UpdateProcessorNotFound with default headers values
func NewUpdateProcessorNotFound() *UpdateProcessorNotFound {
	return &UpdateProcessorNotFound{}
}

/* UpdateProcessorNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type UpdateProcessorNotFound struct {
}

func (o *UpdateProcessorNotFound) Error() string {
	return fmt.Sprintf("[PUT /processors/{id}][%d] updateProcessorNotFound ", 404)
}

func (o *UpdateProcessorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateProcessorConflict creates a UpdateProcessorConflict with default headers values
func NewUpdateProcessorConflict() *UpdateProcessorConflict {
	return &UpdateProcessorConflict{}
}

/* UpdateProcessorConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type UpdateProcessorConflict struct {
}

func (o *UpdateProcessorConflict) Error() string {
	return fmt.Sprintf("[PUT /processors/{id}][%d] updateProcessorConflict ", 409)
}

func (o *UpdateProcessorConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}