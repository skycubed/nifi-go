// Code generated by go-swagger; DO NOT EDIT.

package flows

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/registry/models"
)

// GlobalGetFlowReader is a Reader for the GlobalGetFlow structure.
type GlobalGetFlowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GlobalGetFlowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGlobalGetFlowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGlobalGetFlowBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGlobalGetFlowUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGlobalGetFlowForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGlobalGetFlowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGlobalGetFlowConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /flows/{flowId}] globalGetFlow", response, response.Code())
	}
}

// NewGlobalGetFlowOK creates a GlobalGetFlowOK with default headers values
func NewGlobalGetFlowOK() *GlobalGetFlowOK {
	return &GlobalGetFlowOK{}
}

/*
GlobalGetFlowOK describes a response with status code 200, with default header values.

successful operation
*/
type GlobalGetFlowOK struct {
	Payload *models.VersionedFlow
}

// IsSuccess returns true when this global get flow o k response has a 2xx status code
func (o *GlobalGetFlowOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this global get flow o k response has a 3xx status code
func (o *GlobalGetFlowOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this global get flow o k response has a 4xx status code
func (o *GlobalGetFlowOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this global get flow o k response has a 5xx status code
func (o *GlobalGetFlowOK) IsServerError() bool {
	return false
}

// IsCode returns true when this global get flow o k response a status code equal to that given
func (o *GlobalGetFlowOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the global get flow o k response
func (o *GlobalGetFlowOK) Code() int {
	return 200
}

func (o *GlobalGetFlowOK) Error() string {
	return fmt.Sprintf("[GET /flows/{flowId}][%d] globalGetFlowOK  %+v", 200, o.Payload)
}

func (o *GlobalGetFlowOK) String() string {
	return fmt.Sprintf("[GET /flows/{flowId}][%d] globalGetFlowOK  %+v", 200, o.Payload)
}

func (o *GlobalGetFlowOK) GetPayload() *models.VersionedFlow {
	return o.Payload
}

func (o *GlobalGetFlowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VersionedFlow)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGlobalGetFlowBadRequest creates a GlobalGetFlowBadRequest with default headers values
func NewGlobalGetFlowBadRequest() *GlobalGetFlowBadRequest {
	return &GlobalGetFlowBadRequest{}
}

/*
GlobalGetFlowBadRequest describes a response with status code 400, with default header values.

NiFi Registry was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GlobalGetFlowBadRequest struct {
}

// IsSuccess returns true when this global get flow bad request response has a 2xx status code
func (o *GlobalGetFlowBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this global get flow bad request response has a 3xx status code
func (o *GlobalGetFlowBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this global get flow bad request response has a 4xx status code
func (o *GlobalGetFlowBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this global get flow bad request response has a 5xx status code
func (o *GlobalGetFlowBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this global get flow bad request response a status code equal to that given
func (o *GlobalGetFlowBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the global get flow bad request response
func (o *GlobalGetFlowBadRequest) Code() int {
	return 400
}

func (o *GlobalGetFlowBadRequest) Error() string {
	return fmt.Sprintf("[GET /flows/{flowId}][%d] globalGetFlowBadRequest ", 400)
}

func (o *GlobalGetFlowBadRequest) String() string {
	return fmt.Sprintf("[GET /flows/{flowId}][%d] globalGetFlowBadRequest ", 400)
}

func (o *GlobalGetFlowBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGlobalGetFlowUnauthorized creates a GlobalGetFlowUnauthorized with default headers values
func NewGlobalGetFlowUnauthorized() *GlobalGetFlowUnauthorized {
	return &GlobalGetFlowUnauthorized{}
}

/*
GlobalGetFlowUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GlobalGetFlowUnauthorized struct {
}

// IsSuccess returns true when this global get flow unauthorized response has a 2xx status code
func (o *GlobalGetFlowUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this global get flow unauthorized response has a 3xx status code
func (o *GlobalGetFlowUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this global get flow unauthorized response has a 4xx status code
func (o *GlobalGetFlowUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this global get flow unauthorized response has a 5xx status code
func (o *GlobalGetFlowUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this global get flow unauthorized response a status code equal to that given
func (o *GlobalGetFlowUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the global get flow unauthorized response
func (o *GlobalGetFlowUnauthorized) Code() int {
	return 401
}

func (o *GlobalGetFlowUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flows/{flowId}][%d] globalGetFlowUnauthorized ", 401)
}

func (o *GlobalGetFlowUnauthorized) String() string {
	return fmt.Sprintf("[GET /flows/{flowId}][%d] globalGetFlowUnauthorized ", 401)
}

func (o *GlobalGetFlowUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGlobalGetFlowForbidden creates a GlobalGetFlowForbidden with default headers values
func NewGlobalGetFlowForbidden() *GlobalGetFlowForbidden {
	return &GlobalGetFlowForbidden{}
}

/*
GlobalGetFlowForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GlobalGetFlowForbidden struct {
}

// IsSuccess returns true when this global get flow forbidden response has a 2xx status code
func (o *GlobalGetFlowForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this global get flow forbidden response has a 3xx status code
func (o *GlobalGetFlowForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this global get flow forbidden response has a 4xx status code
func (o *GlobalGetFlowForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this global get flow forbidden response has a 5xx status code
func (o *GlobalGetFlowForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this global get flow forbidden response a status code equal to that given
func (o *GlobalGetFlowForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the global get flow forbidden response
func (o *GlobalGetFlowForbidden) Code() int {
	return 403
}

func (o *GlobalGetFlowForbidden) Error() string {
	return fmt.Sprintf("[GET /flows/{flowId}][%d] globalGetFlowForbidden ", 403)
}

func (o *GlobalGetFlowForbidden) String() string {
	return fmt.Sprintf("[GET /flows/{flowId}][%d] globalGetFlowForbidden ", 403)
}

func (o *GlobalGetFlowForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGlobalGetFlowNotFound creates a GlobalGetFlowNotFound with default headers values
func NewGlobalGetFlowNotFound() *GlobalGetFlowNotFound {
	return &GlobalGetFlowNotFound{}
}

/*
GlobalGetFlowNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GlobalGetFlowNotFound struct {
}

// IsSuccess returns true when this global get flow not found response has a 2xx status code
func (o *GlobalGetFlowNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this global get flow not found response has a 3xx status code
func (o *GlobalGetFlowNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this global get flow not found response has a 4xx status code
func (o *GlobalGetFlowNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this global get flow not found response has a 5xx status code
func (o *GlobalGetFlowNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this global get flow not found response a status code equal to that given
func (o *GlobalGetFlowNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the global get flow not found response
func (o *GlobalGetFlowNotFound) Code() int {
	return 404
}

func (o *GlobalGetFlowNotFound) Error() string {
	return fmt.Sprintf("[GET /flows/{flowId}][%d] globalGetFlowNotFound ", 404)
}

func (o *GlobalGetFlowNotFound) String() string {
	return fmt.Sprintf("[GET /flows/{flowId}][%d] globalGetFlowNotFound ", 404)
}

func (o *GlobalGetFlowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGlobalGetFlowConflict creates a GlobalGetFlowConflict with default headers values
func NewGlobalGetFlowConflict() *GlobalGetFlowConflict {
	return &GlobalGetFlowConflict{}
}

/*
GlobalGetFlowConflict describes a response with status code 409, with default header values.

NiFi Registry was unable to complete the request because it assumes a server state that is not valid.
*/
type GlobalGetFlowConflict struct {
}

// IsSuccess returns true when this global get flow conflict response has a 2xx status code
func (o *GlobalGetFlowConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this global get flow conflict response has a 3xx status code
func (o *GlobalGetFlowConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this global get flow conflict response has a 4xx status code
func (o *GlobalGetFlowConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this global get flow conflict response has a 5xx status code
func (o *GlobalGetFlowConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this global get flow conflict response a status code equal to that given
func (o *GlobalGetFlowConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the global get flow conflict response
func (o *GlobalGetFlowConflict) Code() int {
	return 409
}

func (o *GlobalGetFlowConflict) Error() string {
	return fmt.Sprintf("[GET /flows/{flowId}][%d] globalGetFlowConflict ", 409)
}

func (o *GlobalGetFlowConflict) String() string {
	return fmt.Sprintf("[GET /flows/{flowId}][%d] globalGetFlowConflict ", 409)
}

func (o *GlobalGetFlowConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
