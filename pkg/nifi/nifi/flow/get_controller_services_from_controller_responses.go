// Code generated by go-swagger; DO NOT EDIT.

package flow

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// GetControllerServicesFromControllerReader is a Reader for the GetControllerServicesFromController structure.
type GetControllerServicesFromControllerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetControllerServicesFromControllerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetControllerServicesFromControllerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetControllerServicesFromControllerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetControllerServicesFromControllerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetControllerServicesFromControllerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetControllerServicesFromControllerConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetControllerServicesFromControllerOK creates a GetControllerServicesFromControllerOK with default headers values
func NewGetControllerServicesFromControllerOK() *GetControllerServicesFromControllerOK {
	return &GetControllerServicesFromControllerOK{}
}

/* GetControllerServicesFromControllerOK describes a response with status code 200, with default header values.

successful operation
*/
type GetControllerServicesFromControllerOK struct {
	Payload *models.ControllerServicesEntity
}

func (o *GetControllerServicesFromControllerOK) Error() string {
	return fmt.Sprintf("[GET /flow/controller/controller-services][%d] getControllerServicesFromControllerOK  %+v", 200, o.Payload)
}
func (o *GetControllerServicesFromControllerOK) GetPayload() *models.ControllerServicesEntity {
	return o.Payload
}

func (o *GetControllerServicesFromControllerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ControllerServicesEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetControllerServicesFromControllerBadRequest creates a GetControllerServicesFromControllerBadRequest with default headers values
func NewGetControllerServicesFromControllerBadRequest() *GetControllerServicesFromControllerBadRequest {
	return &GetControllerServicesFromControllerBadRequest{}
}

/* GetControllerServicesFromControllerBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetControllerServicesFromControllerBadRequest struct {
}

func (o *GetControllerServicesFromControllerBadRequest) Error() string {
	return fmt.Sprintf("[GET /flow/controller/controller-services][%d] getControllerServicesFromControllerBadRequest ", 400)
}

func (o *GetControllerServicesFromControllerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetControllerServicesFromControllerUnauthorized creates a GetControllerServicesFromControllerUnauthorized with default headers values
func NewGetControllerServicesFromControllerUnauthorized() *GetControllerServicesFromControllerUnauthorized {
	return &GetControllerServicesFromControllerUnauthorized{}
}

/* GetControllerServicesFromControllerUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetControllerServicesFromControllerUnauthorized struct {
}

func (o *GetControllerServicesFromControllerUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flow/controller/controller-services][%d] getControllerServicesFromControllerUnauthorized ", 401)
}

func (o *GetControllerServicesFromControllerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetControllerServicesFromControllerForbidden creates a GetControllerServicesFromControllerForbidden with default headers values
func NewGetControllerServicesFromControllerForbidden() *GetControllerServicesFromControllerForbidden {
	return &GetControllerServicesFromControllerForbidden{}
}

/* GetControllerServicesFromControllerForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetControllerServicesFromControllerForbidden struct {
}

func (o *GetControllerServicesFromControllerForbidden) Error() string {
	return fmt.Sprintf("[GET /flow/controller/controller-services][%d] getControllerServicesFromControllerForbidden ", 403)
}

func (o *GetControllerServicesFromControllerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetControllerServicesFromControllerConflict creates a GetControllerServicesFromControllerConflict with default headers values
func NewGetControllerServicesFromControllerConflict() *GetControllerServicesFromControllerConflict {
	return &GetControllerServicesFromControllerConflict{}
}

/* GetControllerServicesFromControllerConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetControllerServicesFromControllerConflict struct {
}

func (o *GetControllerServicesFromControllerConflict) Error() string {
	return fmt.Sprintf("[GET /flow/controller/controller-services][%d] getControllerServicesFromControllerConflict ", 409)
}

func (o *GetControllerServicesFromControllerConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
