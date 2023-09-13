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

// GetInputPortStatusReader is a Reader for the GetInputPortStatus structure.
type GetInputPortStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInputPortStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInputPortStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInputPortStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetInputPortStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetInputPortStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetInputPortStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetInputPortStatusConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /flow/input-ports/{id}/status] getInputPortStatus", response, response.Code())
	}
}

// NewGetInputPortStatusOK creates a GetInputPortStatusOK with default headers values
func NewGetInputPortStatusOK() *GetInputPortStatusOK {
	return &GetInputPortStatusOK{}
}

/*
GetInputPortStatusOK describes a response with status code 200, with default header values.

successful operation
*/
type GetInputPortStatusOK struct {
	Payload *models.PortStatusEntity
}

// IsSuccess returns true when this get input port status o k response has a 2xx status code
func (o *GetInputPortStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get input port status o k response has a 3xx status code
func (o *GetInputPortStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get input port status o k response has a 4xx status code
func (o *GetInputPortStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get input port status o k response has a 5xx status code
func (o *GetInputPortStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get input port status o k response a status code equal to that given
func (o *GetInputPortStatusOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get input port status o k response
func (o *GetInputPortStatusOK) Code() int {
	return 200
}

func (o *GetInputPortStatusOK) Error() string {
	return fmt.Sprintf("[GET /flow/input-ports/{id}/status][%d] getInputPortStatusOK  %+v", 200, o.Payload)
}

func (o *GetInputPortStatusOK) String() string {
	return fmt.Sprintf("[GET /flow/input-ports/{id}/status][%d] getInputPortStatusOK  %+v", 200, o.Payload)
}

func (o *GetInputPortStatusOK) GetPayload() *models.PortStatusEntity {
	return o.Payload
}

func (o *GetInputPortStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortStatusEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInputPortStatusBadRequest creates a GetInputPortStatusBadRequest with default headers values
func NewGetInputPortStatusBadRequest() *GetInputPortStatusBadRequest {
	return &GetInputPortStatusBadRequest{}
}

/*
GetInputPortStatusBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetInputPortStatusBadRequest struct {
}

// IsSuccess returns true when this get input port status bad request response has a 2xx status code
func (o *GetInputPortStatusBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get input port status bad request response has a 3xx status code
func (o *GetInputPortStatusBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get input port status bad request response has a 4xx status code
func (o *GetInputPortStatusBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get input port status bad request response has a 5xx status code
func (o *GetInputPortStatusBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get input port status bad request response a status code equal to that given
func (o *GetInputPortStatusBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get input port status bad request response
func (o *GetInputPortStatusBadRequest) Code() int {
	return 400
}

func (o *GetInputPortStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /flow/input-ports/{id}/status][%d] getInputPortStatusBadRequest ", 400)
}

func (o *GetInputPortStatusBadRequest) String() string {
	return fmt.Sprintf("[GET /flow/input-ports/{id}/status][%d] getInputPortStatusBadRequest ", 400)
}

func (o *GetInputPortStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInputPortStatusUnauthorized creates a GetInputPortStatusUnauthorized with default headers values
func NewGetInputPortStatusUnauthorized() *GetInputPortStatusUnauthorized {
	return &GetInputPortStatusUnauthorized{}
}

/*
GetInputPortStatusUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetInputPortStatusUnauthorized struct {
}

// IsSuccess returns true when this get input port status unauthorized response has a 2xx status code
func (o *GetInputPortStatusUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get input port status unauthorized response has a 3xx status code
func (o *GetInputPortStatusUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get input port status unauthorized response has a 4xx status code
func (o *GetInputPortStatusUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get input port status unauthorized response has a 5xx status code
func (o *GetInputPortStatusUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get input port status unauthorized response a status code equal to that given
func (o *GetInputPortStatusUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get input port status unauthorized response
func (o *GetInputPortStatusUnauthorized) Code() int {
	return 401
}

func (o *GetInputPortStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flow/input-ports/{id}/status][%d] getInputPortStatusUnauthorized ", 401)
}

func (o *GetInputPortStatusUnauthorized) String() string {
	return fmt.Sprintf("[GET /flow/input-ports/{id}/status][%d] getInputPortStatusUnauthorized ", 401)
}

func (o *GetInputPortStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInputPortStatusForbidden creates a GetInputPortStatusForbidden with default headers values
func NewGetInputPortStatusForbidden() *GetInputPortStatusForbidden {
	return &GetInputPortStatusForbidden{}
}

/*
GetInputPortStatusForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetInputPortStatusForbidden struct {
}

// IsSuccess returns true when this get input port status forbidden response has a 2xx status code
func (o *GetInputPortStatusForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get input port status forbidden response has a 3xx status code
func (o *GetInputPortStatusForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get input port status forbidden response has a 4xx status code
func (o *GetInputPortStatusForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get input port status forbidden response has a 5xx status code
func (o *GetInputPortStatusForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get input port status forbidden response a status code equal to that given
func (o *GetInputPortStatusForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get input port status forbidden response
func (o *GetInputPortStatusForbidden) Code() int {
	return 403
}

func (o *GetInputPortStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /flow/input-ports/{id}/status][%d] getInputPortStatusForbidden ", 403)
}

func (o *GetInputPortStatusForbidden) String() string {
	return fmt.Sprintf("[GET /flow/input-ports/{id}/status][%d] getInputPortStatusForbidden ", 403)
}

func (o *GetInputPortStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInputPortStatusNotFound creates a GetInputPortStatusNotFound with default headers values
func NewGetInputPortStatusNotFound() *GetInputPortStatusNotFound {
	return &GetInputPortStatusNotFound{}
}

/*
GetInputPortStatusNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetInputPortStatusNotFound struct {
}

// IsSuccess returns true when this get input port status not found response has a 2xx status code
func (o *GetInputPortStatusNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get input port status not found response has a 3xx status code
func (o *GetInputPortStatusNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get input port status not found response has a 4xx status code
func (o *GetInputPortStatusNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get input port status not found response has a 5xx status code
func (o *GetInputPortStatusNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get input port status not found response a status code equal to that given
func (o *GetInputPortStatusNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get input port status not found response
func (o *GetInputPortStatusNotFound) Code() int {
	return 404
}

func (o *GetInputPortStatusNotFound) Error() string {
	return fmt.Sprintf("[GET /flow/input-ports/{id}/status][%d] getInputPortStatusNotFound ", 404)
}

func (o *GetInputPortStatusNotFound) String() string {
	return fmt.Sprintf("[GET /flow/input-ports/{id}/status][%d] getInputPortStatusNotFound ", 404)
}

func (o *GetInputPortStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInputPortStatusConflict creates a GetInputPortStatusConflict with default headers values
func NewGetInputPortStatusConflict() *GetInputPortStatusConflict {
	return &GetInputPortStatusConflict{}
}

/*
GetInputPortStatusConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetInputPortStatusConflict struct {
}

// IsSuccess returns true when this get input port status conflict response has a 2xx status code
func (o *GetInputPortStatusConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get input port status conflict response has a 3xx status code
func (o *GetInputPortStatusConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get input port status conflict response has a 4xx status code
func (o *GetInputPortStatusConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get input port status conflict response has a 5xx status code
func (o *GetInputPortStatusConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get input port status conflict response a status code equal to that given
func (o *GetInputPortStatusConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the get input port status conflict response
func (o *GetInputPortStatusConflict) Code() int {
	return 409
}

func (o *GetInputPortStatusConflict) Error() string {
	return fmt.Sprintf("[GET /flow/input-ports/{id}/status][%d] getInputPortStatusConflict ", 409)
}

func (o *GetInputPortStatusConflict) String() string {
	return fmt.Sprintf("[GET /flow/input-ports/{id}/status][%d] getInputPortStatusConflict ", 409)
}

func (o *GetInputPortStatusConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
