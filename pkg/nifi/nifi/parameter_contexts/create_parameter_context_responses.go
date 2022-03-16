// Code generated by go-swagger; DO NOT EDIT.

package parameter_contexts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// CreateParameterContextReader is a Reader for the CreateParameterContext structure.
type CreateParameterContextReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateParameterContextReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateParameterContextOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateParameterContextBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateParameterContextUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateParameterContextForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateParameterContextNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateParameterContextConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateParameterContextOK creates a CreateParameterContextOK with default headers values
func NewCreateParameterContextOK() *CreateParameterContextOK {
	return &CreateParameterContextOK{}
}

/* CreateParameterContextOK describes a response with status code 200, with default header values.

successful operation
*/
type CreateParameterContextOK struct {
	Payload *models.ParameterContextEntity
}

func (o *CreateParameterContextOK) Error() string {
	return fmt.Sprintf("[POST /parameter-contexts][%d] createParameterContextOK  %+v", 200, o.Payload)
}
func (o *CreateParameterContextOK) GetPayload() *models.ParameterContextEntity {
	return o.Payload
}

func (o *CreateParameterContextOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ParameterContextEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateParameterContextBadRequest creates a CreateParameterContextBadRequest with default headers values
func NewCreateParameterContextBadRequest() *CreateParameterContextBadRequest {
	return &CreateParameterContextBadRequest{}
}

/* CreateParameterContextBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type CreateParameterContextBadRequest struct {
}

func (o *CreateParameterContextBadRequest) Error() string {
	return fmt.Sprintf("[POST /parameter-contexts][%d] createParameterContextBadRequest ", 400)
}

func (o *CreateParameterContextBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateParameterContextUnauthorized creates a CreateParameterContextUnauthorized with default headers values
func NewCreateParameterContextUnauthorized() *CreateParameterContextUnauthorized {
	return &CreateParameterContextUnauthorized{}
}

/* CreateParameterContextUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type CreateParameterContextUnauthorized struct {
}

func (o *CreateParameterContextUnauthorized) Error() string {
	return fmt.Sprintf("[POST /parameter-contexts][%d] createParameterContextUnauthorized ", 401)
}

func (o *CreateParameterContextUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateParameterContextForbidden creates a CreateParameterContextForbidden with default headers values
func NewCreateParameterContextForbidden() *CreateParameterContextForbidden {
	return &CreateParameterContextForbidden{}
}

/* CreateParameterContextForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type CreateParameterContextForbidden struct {
}

func (o *CreateParameterContextForbidden) Error() string {
	return fmt.Sprintf("[POST /parameter-contexts][%d] createParameterContextForbidden ", 403)
}

func (o *CreateParameterContextForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateParameterContextNotFound creates a CreateParameterContextNotFound with default headers values
func NewCreateParameterContextNotFound() *CreateParameterContextNotFound {
	return &CreateParameterContextNotFound{}
}

/* CreateParameterContextNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type CreateParameterContextNotFound struct {
}

func (o *CreateParameterContextNotFound) Error() string {
	return fmt.Sprintf("[POST /parameter-contexts][%d] createParameterContextNotFound ", 404)
}

func (o *CreateParameterContextNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateParameterContextConflict creates a CreateParameterContextConflict with default headers values
func NewCreateParameterContextConflict() *CreateParameterContextConflict {
	return &CreateParameterContextConflict{}
}

/* CreateParameterContextConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type CreateParameterContextConflict struct {
}

func (o *CreateParameterContextConflict) Error() string {
	return fmt.Sprintf("[POST /parameter-contexts][%d] createParameterContextConflict ", 409)
}

func (o *CreateParameterContextConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
