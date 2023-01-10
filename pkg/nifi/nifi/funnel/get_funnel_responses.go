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

// GetFunnelReader is a Reader for the GetFunnel structure.
type GetFunnelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFunnelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFunnelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFunnelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFunnelUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFunnelForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFunnelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetFunnelConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFunnelOK creates a GetFunnelOK with default headers values
func NewGetFunnelOK() *GetFunnelOK {
	return &GetFunnelOK{}
}

/*
GetFunnelOK describes a response with status code 200, with default header values.

successful operation
*/
type GetFunnelOK struct {
	Payload *models.FunnelEntity
}

// IsSuccess returns true when this get funnel o k response has a 2xx status code
func (o *GetFunnelOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get funnel o k response has a 3xx status code
func (o *GetFunnelOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get funnel o k response has a 4xx status code
func (o *GetFunnelOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get funnel o k response has a 5xx status code
func (o *GetFunnelOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get funnel o k response a status code equal to that given
func (o *GetFunnelOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetFunnelOK) Error() string {
	return fmt.Sprintf("[GET /funnels/{id}][%d] getFunnelOK  %+v", 200, o.Payload)
}

func (o *GetFunnelOK) String() string {
	return fmt.Sprintf("[GET /funnels/{id}][%d] getFunnelOK  %+v", 200, o.Payload)
}

func (o *GetFunnelOK) GetPayload() *models.FunnelEntity {
	return o.Payload
}

func (o *GetFunnelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FunnelEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFunnelBadRequest creates a GetFunnelBadRequest with default headers values
func NewGetFunnelBadRequest() *GetFunnelBadRequest {
	return &GetFunnelBadRequest{}
}

/*
GetFunnelBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetFunnelBadRequest struct {
}

// IsSuccess returns true when this get funnel bad request response has a 2xx status code
func (o *GetFunnelBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get funnel bad request response has a 3xx status code
func (o *GetFunnelBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get funnel bad request response has a 4xx status code
func (o *GetFunnelBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get funnel bad request response has a 5xx status code
func (o *GetFunnelBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get funnel bad request response a status code equal to that given
func (o *GetFunnelBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetFunnelBadRequest) Error() string {
	return fmt.Sprintf("[GET /funnels/{id}][%d] getFunnelBadRequest ", 400)
}

func (o *GetFunnelBadRequest) String() string {
	return fmt.Sprintf("[GET /funnels/{id}][%d] getFunnelBadRequest ", 400)
}

func (o *GetFunnelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFunnelUnauthorized creates a GetFunnelUnauthorized with default headers values
func NewGetFunnelUnauthorized() *GetFunnelUnauthorized {
	return &GetFunnelUnauthorized{}
}

/*
GetFunnelUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetFunnelUnauthorized struct {
}

// IsSuccess returns true when this get funnel unauthorized response has a 2xx status code
func (o *GetFunnelUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get funnel unauthorized response has a 3xx status code
func (o *GetFunnelUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get funnel unauthorized response has a 4xx status code
func (o *GetFunnelUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get funnel unauthorized response has a 5xx status code
func (o *GetFunnelUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get funnel unauthorized response a status code equal to that given
func (o *GetFunnelUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetFunnelUnauthorized) Error() string {
	return fmt.Sprintf("[GET /funnels/{id}][%d] getFunnelUnauthorized ", 401)
}

func (o *GetFunnelUnauthorized) String() string {
	return fmt.Sprintf("[GET /funnels/{id}][%d] getFunnelUnauthorized ", 401)
}

func (o *GetFunnelUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFunnelForbidden creates a GetFunnelForbidden with default headers values
func NewGetFunnelForbidden() *GetFunnelForbidden {
	return &GetFunnelForbidden{}
}

/*
GetFunnelForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetFunnelForbidden struct {
}

// IsSuccess returns true when this get funnel forbidden response has a 2xx status code
func (o *GetFunnelForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get funnel forbidden response has a 3xx status code
func (o *GetFunnelForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get funnel forbidden response has a 4xx status code
func (o *GetFunnelForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get funnel forbidden response has a 5xx status code
func (o *GetFunnelForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get funnel forbidden response a status code equal to that given
func (o *GetFunnelForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetFunnelForbidden) Error() string {
	return fmt.Sprintf("[GET /funnels/{id}][%d] getFunnelForbidden ", 403)
}

func (o *GetFunnelForbidden) String() string {
	return fmt.Sprintf("[GET /funnels/{id}][%d] getFunnelForbidden ", 403)
}

func (o *GetFunnelForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFunnelNotFound creates a GetFunnelNotFound with default headers values
func NewGetFunnelNotFound() *GetFunnelNotFound {
	return &GetFunnelNotFound{}
}

/*
GetFunnelNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetFunnelNotFound struct {
}

// IsSuccess returns true when this get funnel not found response has a 2xx status code
func (o *GetFunnelNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get funnel not found response has a 3xx status code
func (o *GetFunnelNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get funnel not found response has a 4xx status code
func (o *GetFunnelNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get funnel not found response has a 5xx status code
func (o *GetFunnelNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get funnel not found response a status code equal to that given
func (o *GetFunnelNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetFunnelNotFound) Error() string {
	return fmt.Sprintf("[GET /funnels/{id}][%d] getFunnelNotFound ", 404)
}

func (o *GetFunnelNotFound) String() string {
	return fmt.Sprintf("[GET /funnels/{id}][%d] getFunnelNotFound ", 404)
}

func (o *GetFunnelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFunnelConflict creates a GetFunnelConflict with default headers values
func NewGetFunnelConflict() *GetFunnelConflict {
	return &GetFunnelConflict{}
}

/*
GetFunnelConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetFunnelConflict struct {
}

// IsSuccess returns true when this get funnel conflict response has a 2xx status code
func (o *GetFunnelConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get funnel conflict response has a 3xx status code
func (o *GetFunnelConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get funnel conflict response has a 4xx status code
func (o *GetFunnelConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get funnel conflict response has a 5xx status code
func (o *GetFunnelConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get funnel conflict response a status code equal to that given
func (o *GetFunnelConflict) IsCode(code int) bool {
	return code == 409
}

func (o *GetFunnelConflict) Error() string {
	return fmt.Sprintf("[GET /funnels/{id}][%d] getFunnelConflict ", 409)
}

func (o *GetFunnelConflict) String() string {
	return fmt.Sprintf("[GET /funnels/{id}][%d] getFunnelConflict ", 409)
}

func (o *GetFunnelConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
