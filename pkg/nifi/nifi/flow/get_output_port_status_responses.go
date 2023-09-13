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

// GetOutputPortStatusReader is a Reader for the GetOutputPortStatus structure.
type GetOutputPortStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutputPortStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutputPortStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutputPortStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutputPortStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutputPortStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutputPortStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetOutputPortStatusConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /flow/output-ports/{id}/status] getOutputPortStatus", response, response.Code())
	}
}

// NewGetOutputPortStatusOK creates a GetOutputPortStatusOK with default headers values
func NewGetOutputPortStatusOK() *GetOutputPortStatusOK {
	return &GetOutputPortStatusOK{}
}

/*
GetOutputPortStatusOK describes a response with status code 200, with default header values.

successful operation
*/
type GetOutputPortStatusOK struct {
	Payload *models.PortStatusEntity
}

// IsSuccess returns true when this get output port status o k response has a 2xx status code
func (o *GetOutputPortStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get output port status o k response has a 3xx status code
func (o *GetOutputPortStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get output port status o k response has a 4xx status code
func (o *GetOutputPortStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get output port status o k response has a 5xx status code
func (o *GetOutputPortStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get output port status o k response a status code equal to that given
func (o *GetOutputPortStatusOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get output port status o k response
func (o *GetOutputPortStatusOK) Code() int {
	return 200
}

func (o *GetOutputPortStatusOK) Error() string {
	return fmt.Sprintf("[GET /flow/output-ports/{id}/status][%d] getOutputPortStatusOK  %+v", 200, o.Payload)
}

func (o *GetOutputPortStatusOK) String() string {
	return fmt.Sprintf("[GET /flow/output-ports/{id}/status][%d] getOutputPortStatusOK  %+v", 200, o.Payload)
}

func (o *GetOutputPortStatusOK) GetPayload() *models.PortStatusEntity {
	return o.Payload
}

func (o *GetOutputPortStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortStatusEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutputPortStatusBadRequest creates a GetOutputPortStatusBadRequest with default headers values
func NewGetOutputPortStatusBadRequest() *GetOutputPortStatusBadRequest {
	return &GetOutputPortStatusBadRequest{}
}

/*
GetOutputPortStatusBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetOutputPortStatusBadRequest struct {
}

// IsSuccess returns true when this get output port status bad request response has a 2xx status code
func (o *GetOutputPortStatusBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get output port status bad request response has a 3xx status code
func (o *GetOutputPortStatusBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get output port status bad request response has a 4xx status code
func (o *GetOutputPortStatusBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get output port status bad request response has a 5xx status code
func (o *GetOutputPortStatusBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get output port status bad request response a status code equal to that given
func (o *GetOutputPortStatusBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get output port status bad request response
func (o *GetOutputPortStatusBadRequest) Code() int {
	return 400
}

func (o *GetOutputPortStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /flow/output-ports/{id}/status][%d] getOutputPortStatusBadRequest ", 400)
}

func (o *GetOutputPortStatusBadRequest) String() string {
	return fmt.Sprintf("[GET /flow/output-ports/{id}/status][%d] getOutputPortStatusBadRequest ", 400)
}

func (o *GetOutputPortStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOutputPortStatusUnauthorized creates a GetOutputPortStatusUnauthorized with default headers values
func NewGetOutputPortStatusUnauthorized() *GetOutputPortStatusUnauthorized {
	return &GetOutputPortStatusUnauthorized{}
}

/*
GetOutputPortStatusUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetOutputPortStatusUnauthorized struct {
}

// IsSuccess returns true when this get output port status unauthorized response has a 2xx status code
func (o *GetOutputPortStatusUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get output port status unauthorized response has a 3xx status code
func (o *GetOutputPortStatusUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get output port status unauthorized response has a 4xx status code
func (o *GetOutputPortStatusUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get output port status unauthorized response has a 5xx status code
func (o *GetOutputPortStatusUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get output port status unauthorized response a status code equal to that given
func (o *GetOutputPortStatusUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get output port status unauthorized response
func (o *GetOutputPortStatusUnauthorized) Code() int {
	return 401
}

func (o *GetOutputPortStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flow/output-ports/{id}/status][%d] getOutputPortStatusUnauthorized ", 401)
}

func (o *GetOutputPortStatusUnauthorized) String() string {
	return fmt.Sprintf("[GET /flow/output-ports/{id}/status][%d] getOutputPortStatusUnauthorized ", 401)
}

func (o *GetOutputPortStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOutputPortStatusForbidden creates a GetOutputPortStatusForbidden with default headers values
func NewGetOutputPortStatusForbidden() *GetOutputPortStatusForbidden {
	return &GetOutputPortStatusForbidden{}
}

/*
GetOutputPortStatusForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetOutputPortStatusForbidden struct {
}

// IsSuccess returns true when this get output port status forbidden response has a 2xx status code
func (o *GetOutputPortStatusForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get output port status forbidden response has a 3xx status code
func (o *GetOutputPortStatusForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get output port status forbidden response has a 4xx status code
func (o *GetOutputPortStatusForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get output port status forbidden response has a 5xx status code
func (o *GetOutputPortStatusForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get output port status forbidden response a status code equal to that given
func (o *GetOutputPortStatusForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get output port status forbidden response
func (o *GetOutputPortStatusForbidden) Code() int {
	return 403
}

func (o *GetOutputPortStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /flow/output-ports/{id}/status][%d] getOutputPortStatusForbidden ", 403)
}

func (o *GetOutputPortStatusForbidden) String() string {
	return fmt.Sprintf("[GET /flow/output-ports/{id}/status][%d] getOutputPortStatusForbidden ", 403)
}

func (o *GetOutputPortStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOutputPortStatusNotFound creates a GetOutputPortStatusNotFound with default headers values
func NewGetOutputPortStatusNotFound() *GetOutputPortStatusNotFound {
	return &GetOutputPortStatusNotFound{}
}

/*
GetOutputPortStatusNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetOutputPortStatusNotFound struct {
}

// IsSuccess returns true when this get output port status not found response has a 2xx status code
func (o *GetOutputPortStatusNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get output port status not found response has a 3xx status code
func (o *GetOutputPortStatusNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get output port status not found response has a 4xx status code
func (o *GetOutputPortStatusNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get output port status not found response has a 5xx status code
func (o *GetOutputPortStatusNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get output port status not found response a status code equal to that given
func (o *GetOutputPortStatusNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get output port status not found response
func (o *GetOutputPortStatusNotFound) Code() int {
	return 404
}

func (o *GetOutputPortStatusNotFound) Error() string {
	return fmt.Sprintf("[GET /flow/output-ports/{id}/status][%d] getOutputPortStatusNotFound ", 404)
}

func (o *GetOutputPortStatusNotFound) String() string {
	return fmt.Sprintf("[GET /flow/output-ports/{id}/status][%d] getOutputPortStatusNotFound ", 404)
}

func (o *GetOutputPortStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOutputPortStatusConflict creates a GetOutputPortStatusConflict with default headers values
func NewGetOutputPortStatusConflict() *GetOutputPortStatusConflict {
	return &GetOutputPortStatusConflict{}
}

/*
GetOutputPortStatusConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetOutputPortStatusConflict struct {
}

// IsSuccess returns true when this get output port status conflict response has a 2xx status code
func (o *GetOutputPortStatusConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get output port status conflict response has a 3xx status code
func (o *GetOutputPortStatusConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get output port status conflict response has a 4xx status code
func (o *GetOutputPortStatusConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get output port status conflict response has a 5xx status code
func (o *GetOutputPortStatusConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get output port status conflict response a status code equal to that given
func (o *GetOutputPortStatusConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the get output port status conflict response
func (o *GetOutputPortStatusConflict) Code() int {
	return 409
}

func (o *GetOutputPortStatusConflict) Error() string {
	return fmt.Sprintf("[GET /flow/output-ports/{id}/status][%d] getOutputPortStatusConflict ", 409)
}

func (o *GetOutputPortStatusConflict) String() string {
	return fmt.Sprintf("[GET /flow/output-ports/{id}/status][%d] getOutputPortStatusConflict ", 409)
}

func (o *GetOutputPortStatusConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
