// Code generated by go-swagger; DO NOT EDIT.

package output_ports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// UpdateOutputPortReader is a Reader for the UpdateOutputPort structure.
type UpdateOutputPortReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOutputPortReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOutputPortOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateOutputPortBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateOutputPortUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateOutputPortForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateOutputPortNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateOutputPortConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateOutputPortOK creates a UpdateOutputPortOK with default headers values
func NewUpdateOutputPortOK() *UpdateOutputPortOK {
	return &UpdateOutputPortOK{}
}

/* UpdateOutputPortOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateOutputPortOK struct {
	Payload *models.PortEntity
}

func (o *UpdateOutputPortOK) Error() string {
	return fmt.Sprintf("[PUT /output-ports/{id}][%d] updateOutputPortOK  %+v", 200, o.Payload)
}
func (o *UpdateOutputPortOK) GetPayload() *models.PortEntity {
	return o.Payload
}

func (o *UpdateOutputPortOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOutputPortBadRequest creates a UpdateOutputPortBadRequest with default headers values
func NewUpdateOutputPortBadRequest() *UpdateOutputPortBadRequest {
	return &UpdateOutputPortBadRequest{}
}

/* UpdateOutputPortBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type UpdateOutputPortBadRequest struct {
}

func (o *UpdateOutputPortBadRequest) Error() string {
	return fmt.Sprintf("[PUT /output-ports/{id}][%d] updateOutputPortBadRequest ", 400)
}

func (o *UpdateOutputPortBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateOutputPortUnauthorized creates a UpdateOutputPortUnauthorized with default headers values
func NewUpdateOutputPortUnauthorized() *UpdateOutputPortUnauthorized {
	return &UpdateOutputPortUnauthorized{}
}

/* UpdateOutputPortUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type UpdateOutputPortUnauthorized struct {
}

func (o *UpdateOutputPortUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /output-ports/{id}][%d] updateOutputPortUnauthorized ", 401)
}

func (o *UpdateOutputPortUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateOutputPortForbidden creates a UpdateOutputPortForbidden with default headers values
func NewUpdateOutputPortForbidden() *UpdateOutputPortForbidden {
	return &UpdateOutputPortForbidden{}
}

/* UpdateOutputPortForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type UpdateOutputPortForbidden struct {
}

func (o *UpdateOutputPortForbidden) Error() string {
	return fmt.Sprintf("[PUT /output-ports/{id}][%d] updateOutputPortForbidden ", 403)
}

func (o *UpdateOutputPortForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateOutputPortNotFound creates a UpdateOutputPortNotFound with default headers values
func NewUpdateOutputPortNotFound() *UpdateOutputPortNotFound {
	return &UpdateOutputPortNotFound{}
}

/* UpdateOutputPortNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type UpdateOutputPortNotFound struct {
}

func (o *UpdateOutputPortNotFound) Error() string {
	return fmt.Sprintf("[PUT /output-ports/{id}][%d] updateOutputPortNotFound ", 404)
}

func (o *UpdateOutputPortNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateOutputPortConflict creates a UpdateOutputPortConflict with default headers values
func NewUpdateOutputPortConflict() *UpdateOutputPortConflict {
	return &UpdateOutputPortConflict{}
}

/* UpdateOutputPortConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type UpdateOutputPortConflict struct {
}

func (o *UpdateOutputPortConflict) Error() string {
	return fmt.Sprintf("[PUT /output-ports/{id}][%d] updateOutputPortConflict ", 409)
}

func (o *UpdateOutputPortConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
