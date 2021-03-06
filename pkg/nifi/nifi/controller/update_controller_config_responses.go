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

// UpdateControllerConfigReader is a Reader for the UpdateControllerConfig structure.
type UpdateControllerConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateControllerConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateControllerConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateControllerConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateControllerConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateControllerConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateControllerConfigConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateControllerConfigOK creates a UpdateControllerConfigOK with default headers values
func NewUpdateControllerConfigOK() *UpdateControllerConfigOK {
	return &UpdateControllerConfigOK{}
}

/* UpdateControllerConfigOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateControllerConfigOK struct {
	Payload *models.ControllerConfigurationEntity
}

func (o *UpdateControllerConfigOK) Error() string {
	return fmt.Sprintf("[PUT /controller/config][%d] updateControllerConfigOK  %+v", 200, o.Payload)
}
func (o *UpdateControllerConfigOK) GetPayload() *models.ControllerConfigurationEntity {
	return o.Payload
}

func (o *UpdateControllerConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ControllerConfigurationEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateControllerConfigBadRequest creates a UpdateControllerConfigBadRequest with default headers values
func NewUpdateControllerConfigBadRequest() *UpdateControllerConfigBadRequest {
	return &UpdateControllerConfigBadRequest{}
}

/* UpdateControllerConfigBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type UpdateControllerConfigBadRequest struct {
}

func (o *UpdateControllerConfigBadRequest) Error() string {
	return fmt.Sprintf("[PUT /controller/config][%d] updateControllerConfigBadRequest ", 400)
}

func (o *UpdateControllerConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateControllerConfigUnauthorized creates a UpdateControllerConfigUnauthorized with default headers values
func NewUpdateControllerConfigUnauthorized() *UpdateControllerConfigUnauthorized {
	return &UpdateControllerConfigUnauthorized{}
}

/* UpdateControllerConfigUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type UpdateControllerConfigUnauthorized struct {
}

func (o *UpdateControllerConfigUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /controller/config][%d] updateControllerConfigUnauthorized ", 401)
}

func (o *UpdateControllerConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateControllerConfigForbidden creates a UpdateControllerConfigForbidden with default headers values
func NewUpdateControllerConfigForbidden() *UpdateControllerConfigForbidden {
	return &UpdateControllerConfigForbidden{}
}

/* UpdateControllerConfigForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type UpdateControllerConfigForbidden struct {
}

func (o *UpdateControllerConfigForbidden) Error() string {
	return fmt.Sprintf("[PUT /controller/config][%d] updateControllerConfigForbidden ", 403)
}

func (o *UpdateControllerConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateControllerConfigConflict creates a UpdateControllerConfigConflict with default headers values
func NewUpdateControllerConfigConflict() *UpdateControllerConfigConflict {
	return &UpdateControllerConfigConflict{}
}

/* UpdateControllerConfigConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type UpdateControllerConfigConflict struct {
}

func (o *UpdateControllerConfigConflict) Error() string {
	return fmt.Sprintf("[PUT /controller/config][%d] updateControllerConfigConflict ", 409)
}

func (o *UpdateControllerConfigConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
