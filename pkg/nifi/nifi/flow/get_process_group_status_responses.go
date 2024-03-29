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

// GetProcessGroupStatusReader is a Reader for the GetProcessGroupStatus structure.
type GetProcessGroupStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProcessGroupStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProcessGroupStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProcessGroupStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetProcessGroupStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetProcessGroupStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProcessGroupStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetProcessGroupStatusConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /flow/process-groups/{id}/status] getProcessGroupStatus", response, response.Code())
	}
}

// NewGetProcessGroupStatusOK creates a GetProcessGroupStatusOK with default headers values
func NewGetProcessGroupStatusOK() *GetProcessGroupStatusOK {
	return &GetProcessGroupStatusOK{}
}

/*
GetProcessGroupStatusOK describes a response with status code 200, with default header values.

successful operation
*/
type GetProcessGroupStatusOK struct {
	Payload *models.ProcessGroupStatusEntity
}

// IsSuccess returns true when this get process group status o k response has a 2xx status code
func (o *GetProcessGroupStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get process group status o k response has a 3xx status code
func (o *GetProcessGroupStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get process group status o k response has a 4xx status code
func (o *GetProcessGroupStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get process group status o k response has a 5xx status code
func (o *GetProcessGroupStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get process group status o k response a status code equal to that given
func (o *GetProcessGroupStatusOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get process group status o k response
func (o *GetProcessGroupStatusOK) Code() int {
	return 200
}

func (o *GetProcessGroupStatusOK) Error() string {
	return fmt.Sprintf("[GET /flow/process-groups/{id}/status][%d] getProcessGroupStatusOK  %+v", 200, o.Payload)
}

func (o *GetProcessGroupStatusOK) String() string {
	return fmt.Sprintf("[GET /flow/process-groups/{id}/status][%d] getProcessGroupStatusOK  %+v", 200, o.Payload)
}

func (o *GetProcessGroupStatusOK) GetPayload() *models.ProcessGroupStatusEntity {
	return o.Payload
}

func (o *GetProcessGroupStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProcessGroupStatusEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProcessGroupStatusBadRequest creates a GetProcessGroupStatusBadRequest with default headers values
func NewGetProcessGroupStatusBadRequest() *GetProcessGroupStatusBadRequest {
	return &GetProcessGroupStatusBadRequest{}
}

/*
GetProcessGroupStatusBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetProcessGroupStatusBadRequest struct {
}

// IsSuccess returns true when this get process group status bad request response has a 2xx status code
func (o *GetProcessGroupStatusBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get process group status bad request response has a 3xx status code
func (o *GetProcessGroupStatusBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get process group status bad request response has a 4xx status code
func (o *GetProcessGroupStatusBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get process group status bad request response has a 5xx status code
func (o *GetProcessGroupStatusBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get process group status bad request response a status code equal to that given
func (o *GetProcessGroupStatusBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get process group status bad request response
func (o *GetProcessGroupStatusBadRequest) Code() int {
	return 400
}

func (o *GetProcessGroupStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /flow/process-groups/{id}/status][%d] getProcessGroupStatusBadRequest ", 400)
}

func (o *GetProcessGroupStatusBadRequest) String() string {
	return fmt.Sprintf("[GET /flow/process-groups/{id}/status][%d] getProcessGroupStatusBadRequest ", 400)
}

func (o *GetProcessGroupStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProcessGroupStatusUnauthorized creates a GetProcessGroupStatusUnauthorized with default headers values
func NewGetProcessGroupStatusUnauthorized() *GetProcessGroupStatusUnauthorized {
	return &GetProcessGroupStatusUnauthorized{}
}

/*
GetProcessGroupStatusUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetProcessGroupStatusUnauthorized struct {
}

// IsSuccess returns true when this get process group status unauthorized response has a 2xx status code
func (o *GetProcessGroupStatusUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get process group status unauthorized response has a 3xx status code
func (o *GetProcessGroupStatusUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get process group status unauthorized response has a 4xx status code
func (o *GetProcessGroupStatusUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get process group status unauthorized response has a 5xx status code
func (o *GetProcessGroupStatusUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get process group status unauthorized response a status code equal to that given
func (o *GetProcessGroupStatusUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get process group status unauthorized response
func (o *GetProcessGroupStatusUnauthorized) Code() int {
	return 401
}

func (o *GetProcessGroupStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flow/process-groups/{id}/status][%d] getProcessGroupStatusUnauthorized ", 401)
}

func (o *GetProcessGroupStatusUnauthorized) String() string {
	return fmt.Sprintf("[GET /flow/process-groups/{id}/status][%d] getProcessGroupStatusUnauthorized ", 401)
}

func (o *GetProcessGroupStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProcessGroupStatusForbidden creates a GetProcessGroupStatusForbidden with default headers values
func NewGetProcessGroupStatusForbidden() *GetProcessGroupStatusForbidden {
	return &GetProcessGroupStatusForbidden{}
}

/*
GetProcessGroupStatusForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetProcessGroupStatusForbidden struct {
}

// IsSuccess returns true when this get process group status forbidden response has a 2xx status code
func (o *GetProcessGroupStatusForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get process group status forbidden response has a 3xx status code
func (o *GetProcessGroupStatusForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get process group status forbidden response has a 4xx status code
func (o *GetProcessGroupStatusForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get process group status forbidden response has a 5xx status code
func (o *GetProcessGroupStatusForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get process group status forbidden response a status code equal to that given
func (o *GetProcessGroupStatusForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get process group status forbidden response
func (o *GetProcessGroupStatusForbidden) Code() int {
	return 403
}

func (o *GetProcessGroupStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /flow/process-groups/{id}/status][%d] getProcessGroupStatusForbidden ", 403)
}

func (o *GetProcessGroupStatusForbidden) String() string {
	return fmt.Sprintf("[GET /flow/process-groups/{id}/status][%d] getProcessGroupStatusForbidden ", 403)
}

func (o *GetProcessGroupStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProcessGroupStatusNotFound creates a GetProcessGroupStatusNotFound with default headers values
func NewGetProcessGroupStatusNotFound() *GetProcessGroupStatusNotFound {
	return &GetProcessGroupStatusNotFound{}
}

/*
GetProcessGroupStatusNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetProcessGroupStatusNotFound struct {
}

// IsSuccess returns true when this get process group status not found response has a 2xx status code
func (o *GetProcessGroupStatusNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get process group status not found response has a 3xx status code
func (o *GetProcessGroupStatusNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get process group status not found response has a 4xx status code
func (o *GetProcessGroupStatusNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get process group status not found response has a 5xx status code
func (o *GetProcessGroupStatusNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get process group status not found response a status code equal to that given
func (o *GetProcessGroupStatusNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get process group status not found response
func (o *GetProcessGroupStatusNotFound) Code() int {
	return 404
}

func (o *GetProcessGroupStatusNotFound) Error() string {
	return fmt.Sprintf("[GET /flow/process-groups/{id}/status][%d] getProcessGroupStatusNotFound ", 404)
}

func (o *GetProcessGroupStatusNotFound) String() string {
	return fmt.Sprintf("[GET /flow/process-groups/{id}/status][%d] getProcessGroupStatusNotFound ", 404)
}

func (o *GetProcessGroupStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProcessGroupStatusConflict creates a GetProcessGroupStatusConflict with default headers values
func NewGetProcessGroupStatusConflict() *GetProcessGroupStatusConflict {
	return &GetProcessGroupStatusConflict{}
}

/*
GetProcessGroupStatusConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetProcessGroupStatusConflict struct {
}

// IsSuccess returns true when this get process group status conflict response has a 2xx status code
func (o *GetProcessGroupStatusConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get process group status conflict response has a 3xx status code
func (o *GetProcessGroupStatusConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get process group status conflict response has a 4xx status code
func (o *GetProcessGroupStatusConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get process group status conflict response has a 5xx status code
func (o *GetProcessGroupStatusConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get process group status conflict response a status code equal to that given
func (o *GetProcessGroupStatusConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the get process group status conflict response
func (o *GetProcessGroupStatusConflict) Code() int {
	return 409
}

func (o *GetProcessGroupStatusConflict) Error() string {
	return fmt.Sprintf("[GET /flow/process-groups/{id}/status][%d] getProcessGroupStatusConflict ", 409)
}

func (o *GetProcessGroupStatusConflict) String() string {
	return fmt.Sprintf("[GET /flow/process-groups/{id}/status][%d] getProcessGroupStatusConflict ", 409)
}

func (o *GetProcessGroupStatusConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
