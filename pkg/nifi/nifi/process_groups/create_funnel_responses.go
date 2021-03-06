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

// CreateFunnelReader is a Reader for the CreateFunnel structure.
type CreateFunnelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateFunnelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateFunnelCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateFunnelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateFunnelUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateFunnelForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateFunnelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateFunnelConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateFunnelCreated creates a CreateFunnelCreated with default headers values
func NewCreateFunnelCreated() *CreateFunnelCreated {
	return &CreateFunnelCreated{}
}

/* CreateFunnelCreated describes a response with status code 201, with default header values.

successful operation
*/
type CreateFunnelCreated struct {
	Payload *models.FunnelEntity
}

func (o *CreateFunnelCreated) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/funnels][%d] createFunnelCreated  %+v", 201, o.Payload)
}
func (o *CreateFunnelCreated) GetPayload() *models.FunnelEntity {
	return o.Payload
}

func (o *CreateFunnelCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FunnelEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFunnelBadRequest creates a CreateFunnelBadRequest with default headers values
func NewCreateFunnelBadRequest() *CreateFunnelBadRequest {
	return &CreateFunnelBadRequest{}
}

/* CreateFunnelBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type CreateFunnelBadRequest struct {
}

func (o *CreateFunnelBadRequest) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/funnels][%d] createFunnelBadRequest ", 400)
}

func (o *CreateFunnelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateFunnelUnauthorized creates a CreateFunnelUnauthorized with default headers values
func NewCreateFunnelUnauthorized() *CreateFunnelUnauthorized {
	return &CreateFunnelUnauthorized{}
}

/* CreateFunnelUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type CreateFunnelUnauthorized struct {
}

func (o *CreateFunnelUnauthorized) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/funnels][%d] createFunnelUnauthorized ", 401)
}

func (o *CreateFunnelUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateFunnelForbidden creates a CreateFunnelForbidden with default headers values
func NewCreateFunnelForbidden() *CreateFunnelForbidden {
	return &CreateFunnelForbidden{}
}

/* CreateFunnelForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type CreateFunnelForbidden struct {
}

func (o *CreateFunnelForbidden) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/funnels][%d] createFunnelForbidden ", 403)
}

func (o *CreateFunnelForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateFunnelNotFound creates a CreateFunnelNotFound with default headers values
func NewCreateFunnelNotFound() *CreateFunnelNotFound {
	return &CreateFunnelNotFound{}
}

/* CreateFunnelNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type CreateFunnelNotFound struct {
}

func (o *CreateFunnelNotFound) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/funnels][%d] createFunnelNotFound ", 404)
}

func (o *CreateFunnelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateFunnelConflict creates a CreateFunnelConflict with default headers values
func NewCreateFunnelConflict() *CreateFunnelConflict {
	return &CreateFunnelConflict{}
}

/* CreateFunnelConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type CreateFunnelConflict struct {
}

func (o *CreateFunnelConflict) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/funnels][%d] createFunnelConflict ", 409)
}

func (o *CreateFunnelConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
