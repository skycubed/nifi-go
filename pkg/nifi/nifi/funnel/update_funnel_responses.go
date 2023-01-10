// Code generated by go-swagger; DO NOT EDIT.

package funnel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// UpdateFunnelReader is a Reader for the UpdateFunnel structure.
type UpdateFunnelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateFunnelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateFunnelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateFunnelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateFunnelUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateFunnelForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateFunnelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateFunnelConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateFunnelOK creates a UpdateFunnelOK with default headers values
func NewUpdateFunnelOK() *UpdateFunnelOK {
	return &UpdateFunnelOK{}
}

/*
UpdateFunnelOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateFunnelOK struct {
	Payload *models.FunnelEntity
}

// IsSuccess returns true when this update funnel o k response has a 2xx status code
func (o *UpdateFunnelOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update funnel o k response has a 3xx status code
func (o *UpdateFunnelOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update funnel o k response has a 4xx status code
func (o *UpdateFunnelOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update funnel o k response has a 5xx status code
func (o *UpdateFunnelOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update funnel o k response a status code equal to that given
func (o *UpdateFunnelOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateFunnelOK) Error() string {
	return fmt.Sprintf("[PUT /funnels/{id}][%d] updateFunnelOK  %+v", 200, o.Payload)
}

func (o *UpdateFunnelOK) String() string {
	return fmt.Sprintf("[PUT /funnels/{id}][%d] updateFunnelOK  %+v", 200, o.Payload)
}

func (o *UpdateFunnelOK) GetPayload() *models.FunnelEntity {
	return o.Payload
}

func (o *UpdateFunnelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FunnelEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFunnelBadRequest creates a UpdateFunnelBadRequest with default headers values
func NewUpdateFunnelBadRequest() *UpdateFunnelBadRequest {
	return &UpdateFunnelBadRequest{}
}

/*
UpdateFunnelBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type UpdateFunnelBadRequest struct {
}

// IsSuccess returns true when this update funnel bad request response has a 2xx status code
func (o *UpdateFunnelBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update funnel bad request response has a 3xx status code
func (o *UpdateFunnelBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update funnel bad request response has a 4xx status code
func (o *UpdateFunnelBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update funnel bad request response has a 5xx status code
func (o *UpdateFunnelBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update funnel bad request response a status code equal to that given
func (o *UpdateFunnelBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UpdateFunnelBadRequest) Error() string {
	return fmt.Sprintf("[PUT /funnels/{id}][%d] updateFunnelBadRequest ", 400)
}

func (o *UpdateFunnelBadRequest) String() string {
	return fmt.Sprintf("[PUT /funnels/{id}][%d] updateFunnelBadRequest ", 400)
}

func (o *UpdateFunnelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateFunnelUnauthorized creates a UpdateFunnelUnauthorized with default headers values
func NewUpdateFunnelUnauthorized() *UpdateFunnelUnauthorized {
	return &UpdateFunnelUnauthorized{}
}

/*
UpdateFunnelUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type UpdateFunnelUnauthorized struct {
}

// IsSuccess returns true when this update funnel unauthorized response has a 2xx status code
func (o *UpdateFunnelUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update funnel unauthorized response has a 3xx status code
func (o *UpdateFunnelUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update funnel unauthorized response has a 4xx status code
func (o *UpdateFunnelUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update funnel unauthorized response has a 5xx status code
func (o *UpdateFunnelUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update funnel unauthorized response a status code equal to that given
func (o *UpdateFunnelUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateFunnelUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /funnels/{id}][%d] updateFunnelUnauthorized ", 401)
}

func (o *UpdateFunnelUnauthorized) String() string {
	return fmt.Sprintf("[PUT /funnels/{id}][%d] updateFunnelUnauthorized ", 401)
}

func (o *UpdateFunnelUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateFunnelForbidden creates a UpdateFunnelForbidden with default headers values
func NewUpdateFunnelForbidden() *UpdateFunnelForbidden {
	return &UpdateFunnelForbidden{}
}

/*
UpdateFunnelForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type UpdateFunnelForbidden struct {
}

// IsSuccess returns true when this update funnel forbidden response has a 2xx status code
func (o *UpdateFunnelForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update funnel forbidden response has a 3xx status code
func (o *UpdateFunnelForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update funnel forbidden response has a 4xx status code
func (o *UpdateFunnelForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update funnel forbidden response has a 5xx status code
func (o *UpdateFunnelForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update funnel forbidden response a status code equal to that given
func (o *UpdateFunnelForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateFunnelForbidden) Error() string {
	return fmt.Sprintf("[PUT /funnels/{id}][%d] updateFunnelForbidden ", 403)
}

func (o *UpdateFunnelForbidden) String() string {
	return fmt.Sprintf("[PUT /funnels/{id}][%d] updateFunnelForbidden ", 403)
}

func (o *UpdateFunnelForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateFunnelNotFound creates a UpdateFunnelNotFound with default headers values
func NewUpdateFunnelNotFound() *UpdateFunnelNotFound {
	return &UpdateFunnelNotFound{}
}

/*
UpdateFunnelNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type UpdateFunnelNotFound struct {
}

// IsSuccess returns true when this update funnel not found response has a 2xx status code
func (o *UpdateFunnelNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update funnel not found response has a 3xx status code
func (o *UpdateFunnelNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update funnel not found response has a 4xx status code
func (o *UpdateFunnelNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update funnel not found response has a 5xx status code
func (o *UpdateFunnelNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update funnel not found response a status code equal to that given
func (o *UpdateFunnelNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateFunnelNotFound) Error() string {
	return fmt.Sprintf("[PUT /funnels/{id}][%d] updateFunnelNotFound ", 404)
}

func (o *UpdateFunnelNotFound) String() string {
	return fmt.Sprintf("[PUT /funnels/{id}][%d] updateFunnelNotFound ", 404)
}

func (o *UpdateFunnelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateFunnelConflict creates a UpdateFunnelConflict with default headers values
func NewUpdateFunnelConflict() *UpdateFunnelConflict {
	return &UpdateFunnelConflict{}
}

/*
UpdateFunnelConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type UpdateFunnelConflict struct {
}

// IsSuccess returns true when this update funnel conflict response has a 2xx status code
func (o *UpdateFunnelConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update funnel conflict response has a 3xx status code
func (o *UpdateFunnelConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update funnel conflict response has a 4xx status code
func (o *UpdateFunnelConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update funnel conflict response has a 5xx status code
func (o *UpdateFunnelConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update funnel conflict response a status code equal to that given
func (o *UpdateFunnelConflict) IsCode(code int) bool {
	return code == 409
}

func (o *UpdateFunnelConflict) Error() string {
	return fmt.Sprintf("[PUT /funnels/{id}][%d] updateFunnelConflict ", 409)
}

func (o *UpdateFunnelConflict) String() string {
	return fmt.Sprintf("[PUT /funnels/{id}][%d] updateFunnelConflict ", 409)
}

func (o *UpdateFunnelConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
