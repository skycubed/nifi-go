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

// GetInputPortsReader is a Reader for the GetInputPorts structure.
type GetInputPortsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInputPortsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInputPortsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInputPortsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetInputPortsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetInputPortsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetInputPortsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetInputPortsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetInputPortsOK creates a GetInputPortsOK with default headers values
func NewGetInputPortsOK() *GetInputPortsOK {
	return &GetInputPortsOK{}
}

/* GetInputPortsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetInputPortsOK struct {
	Payload *models.InputPortsEntity
}

func (o *GetInputPortsOK) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/input-ports][%d] getInputPortsOK  %+v", 200, o.Payload)
}
func (o *GetInputPortsOK) GetPayload() *models.InputPortsEntity {
	return o.Payload
}

func (o *GetInputPortsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InputPortsEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInputPortsBadRequest creates a GetInputPortsBadRequest with default headers values
func NewGetInputPortsBadRequest() *GetInputPortsBadRequest {
	return &GetInputPortsBadRequest{}
}

/* GetInputPortsBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetInputPortsBadRequest struct {
}

func (o *GetInputPortsBadRequest) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/input-ports][%d] getInputPortsBadRequest ", 400)
}

func (o *GetInputPortsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInputPortsUnauthorized creates a GetInputPortsUnauthorized with default headers values
func NewGetInputPortsUnauthorized() *GetInputPortsUnauthorized {
	return &GetInputPortsUnauthorized{}
}

/* GetInputPortsUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetInputPortsUnauthorized struct {
}

func (o *GetInputPortsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/input-ports][%d] getInputPortsUnauthorized ", 401)
}

func (o *GetInputPortsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInputPortsForbidden creates a GetInputPortsForbidden with default headers values
func NewGetInputPortsForbidden() *GetInputPortsForbidden {
	return &GetInputPortsForbidden{}
}

/* GetInputPortsForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetInputPortsForbidden struct {
}

func (o *GetInputPortsForbidden) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/input-ports][%d] getInputPortsForbidden ", 403)
}

func (o *GetInputPortsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInputPortsNotFound creates a GetInputPortsNotFound with default headers values
func NewGetInputPortsNotFound() *GetInputPortsNotFound {
	return &GetInputPortsNotFound{}
}

/* GetInputPortsNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetInputPortsNotFound struct {
}

func (o *GetInputPortsNotFound) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/input-ports][%d] getInputPortsNotFound ", 404)
}

func (o *GetInputPortsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInputPortsConflict creates a GetInputPortsConflict with default headers values
func NewGetInputPortsConflict() *GetInputPortsConflict {
	return &GetInputPortsConflict{}
}

/* GetInputPortsConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetInputPortsConflict struct {
}

func (o *GetInputPortsConflict) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/input-ports][%d] getInputPortsConflict ", 409)
}

func (o *GetInputPortsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
