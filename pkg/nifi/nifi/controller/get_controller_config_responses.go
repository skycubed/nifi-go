// Code generated by go-swagger; DO NOT EDIT.

package controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// GetControllerConfigReader is a Reader for the GetControllerConfig structure.
type GetControllerConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetControllerConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetControllerConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetControllerConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetControllerConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetControllerConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetControllerConfigConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetControllerConfigOK creates a GetControllerConfigOK with default headers values
func NewGetControllerConfigOK() *GetControllerConfigOK {
	return &GetControllerConfigOK{}
}

/* GetControllerConfigOK describes a response with status code 200, with default header values.

successful operation
*/
type GetControllerConfigOK struct {
	Payload *models.ControllerConfigurationEntity
}

func (o *GetControllerConfigOK) Error() string {
	return fmt.Sprintf("[GET /controller/config][%d] getControllerConfigOK  %+v", 200, o.Payload)
}
func (o *GetControllerConfigOK) GetPayload() *models.ControllerConfigurationEntity {
	return o.Payload
}

func (o *GetControllerConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ControllerConfigurationEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetControllerConfigBadRequest creates a GetControllerConfigBadRequest with default headers values
func NewGetControllerConfigBadRequest() *GetControllerConfigBadRequest {
	return &GetControllerConfigBadRequest{}
}

/* GetControllerConfigBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetControllerConfigBadRequest struct {
}

func (o *GetControllerConfigBadRequest) Error() string {
	return fmt.Sprintf("[GET /controller/config][%d] getControllerConfigBadRequest ", 400)
}

func (o *GetControllerConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetControllerConfigUnauthorized creates a GetControllerConfigUnauthorized with default headers values
func NewGetControllerConfigUnauthorized() *GetControllerConfigUnauthorized {
	return &GetControllerConfigUnauthorized{}
}

/* GetControllerConfigUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetControllerConfigUnauthorized struct {
}

func (o *GetControllerConfigUnauthorized) Error() string {
	return fmt.Sprintf("[GET /controller/config][%d] getControllerConfigUnauthorized ", 401)
}

func (o *GetControllerConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetControllerConfigForbidden creates a GetControllerConfigForbidden with default headers values
func NewGetControllerConfigForbidden() *GetControllerConfigForbidden {
	return &GetControllerConfigForbidden{}
}

/* GetControllerConfigForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetControllerConfigForbidden struct {
}

func (o *GetControllerConfigForbidden) Error() string {
	return fmt.Sprintf("[GET /controller/config][%d] getControllerConfigForbidden ", 403)
}

func (o *GetControllerConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetControllerConfigConflict creates a GetControllerConfigConflict with default headers values
func NewGetControllerConfigConflict() *GetControllerConfigConflict {
	return &GetControllerConfigConflict{}
}

/* GetControllerConfigConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetControllerConfigConflict struct {
}

func (o *GetControllerConfigConflict) Error() string {
	return fmt.Sprintf("[GET /controller/config][%d] getControllerConfigConflict ", 409)
}

func (o *GetControllerConfigConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
