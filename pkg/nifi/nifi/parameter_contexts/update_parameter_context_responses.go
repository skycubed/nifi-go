// Code generated by go-swagger; DO NOT EDIT.

package parameter_contexts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// UpdateParameterContextReader is a Reader for the UpdateParameterContext structure.
type UpdateParameterContextReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateParameterContextReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateParameterContextOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateParameterContextBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateParameterContextUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateParameterContextForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateParameterContextNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateParameterContextConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /parameter-contexts/{id}] updateParameterContext", response, response.Code())
	}
}

// NewUpdateParameterContextOK creates a UpdateParameterContextOK with default headers values
func NewUpdateParameterContextOK() *UpdateParameterContextOK {
	return &UpdateParameterContextOK{}
}

/*
UpdateParameterContextOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateParameterContextOK struct {
	Payload *models.ParameterContextEntity
}

// IsSuccess returns true when this update parameter context o k response has a 2xx status code
func (o *UpdateParameterContextOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update parameter context o k response has a 3xx status code
func (o *UpdateParameterContextOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update parameter context o k response has a 4xx status code
func (o *UpdateParameterContextOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update parameter context o k response has a 5xx status code
func (o *UpdateParameterContextOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update parameter context o k response a status code equal to that given
func (o *UpdateParameterContextOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update parameter context o k response
func (o *UpdateParameterContextOK) Code() int {
	return 200
}

func (o *UpdateParameterContextOK) Error() string {
	return fmt.Sprintf("[PUT /parameter-contexts/{id}][%d] updateParameterContextOK  %+v", 200, o.Payload)
}

func (o *UpdateParameterContextOK) String() string {
	return fmt.Sprintf("[PUT /parameter-contexts/{id}][%d] updateParameterContextOK  %+v", 200, o.Payload)
}

func (o *UpdateParameterContextOK) GetPayload() *models.ParameterContextEntity {
	return o.Payload
}

func (o *UpdateParameterContextOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ParameterContextEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateParameterContextBadRequest creates a UpdateParameterContextBadRequest with default headers values
func NewUpdateParameterContextBadRequest() *UpdateParameterContextBadRequest {
	return &UpdateParameterContextBadRequest{}
}

/*
UpdateParameterContextBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type UpdateParameterContextBadRequest struct {
}

// IsSuccess returns true when this update parameter context bad request response has a 2xx status code
func (o *UpdateParameterContextBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update parameter context bad request response has a 3xx status code
func (o *UpdateParameterContextBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update parameter context bad request response has a 4xx status code
func (o *UpdateParameterContextBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update parameter context bad request response has a 5xx status code
func (o *UpdateParameterContextBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update parameter context bad request response a status code equal to that given
func (o *UpdateParameterContextBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update parameter context bad request response
func (o *UpdateParameterContextBadRequest) Code() int {
	return 400
}

func (o *UpdateParameterContextBadRequest) Error() string {
	return fmt.Sprintf("[PUT /parameter-contexts/{id}][%d] updateParameterContextBadRequest ", 400)
}

func (o *UpdateParameterContextBadRequest) String() string {
	return fmt.Sprintf("[PUT /parameter-contexts/{id}][%d] updateParameterContextBadRequest ", 400)
}

func (o *UpdateParameterContextBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateParameterContextUnauthorized creates a UpdateParameterContextUnauthorized with default headers values
func NewUpdateParameterContextUnauthorized() *UpdateParameterContextUnauthorized {
	return &UpdateParameterContextUnauthorized{}
}

/*
UpdateParameterContextUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type UpdateParameterContextUnauthorized struct {
}

// IsSuccess returns true when this update parameter context unauthorized response has a 2xx status code
func (o *UpdateParameterContextUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update parameter context unauthorized response has a 3xx status code
func (o *UpdateParameterContextUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update parameter context unauthorized response has a 4xx status code
func (o *UpdateParameterContextUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update parameter context unauthorized response has a 5xx status code
func (o *UpdateParameterContextUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update parameter context unauthorized response a status code equal to that given
func (o *UpdateParameterContextUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update parameter context unauthorized response
func (o *UpdateParameterContextUnauthorized) Code() int {
	return 401
}

func (o *UpdateParameterContextUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /parameter-contexts/{id}][%d] updateParameterContextUnauthorized ", 401)
}

func (o *UpdateParameterContextUnauthorized) String() string {
	return fmt.Sprintf("[PUT /parameter-contexts/{id}][%d] updateParameterContextUnauthorized ", 401)
}

func (o *UpdateParameterContextUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateParameterContextForbidden creates a UpdateParameterContextForbidden with default headers values
func NewUpdateParameterContextForbidden() *UpdateParameterContextForbidden {
	return &UpdateParameterContextForbidden{}
}

/*
UpdateParameterContextForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type UpdateParameterContextForbidden struct {
}

// IsSuccess returns true when this update parameter context forbidden response has a 2xx status code
func (o *UpdateParameterContextForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update parameter context forbidden response has a 3xx status code
func (o *UpdateParameterContextForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update parameter context forbidden response has a 4xx status code
func (o *UpdateParameterContextForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update parameter context forbidden response has a 5xx status code
func (o *UpdateParameterContextForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update parameter context forbidden response a status code equal to that given
func (o *UpdateParameterContextForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update parameter context forbidden response
func (o *UpdateParameterContextForbidden) Code() int {
	return 403
}

func (o *UpdateParameterContextForbidden) Error() string {
	return fmt.Sprintf("[PUT /parameter-contexts/{id}][%d] updateParameterContextForbidden ", 403)
}

func (o *UpdateParameterContextForbidden) String() string {
	return fmt.Sprintf("[PUT /parameter-contexts/{id}][%d] updateParameterContextForbidden ", 403)
}

func (o *UpdateParameterContextForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateParameterContextNotFound creates a UpdateParameterContextNotFound with default headers values
func NewUpdateParameterContextNotFound() *UpdateParameterContextNotFound {
	return &UpdateParameterContextNotFound{}
}

/*
UpdateParameterContextNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type UpdateParameterContextNotFound struct {
}

// IsSuccess returns true when this update parameter context not found response has a 2xx status code
func (o *UpdateParameterContextNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update parameter context not found response has a 3xx status code
func (o *UpdateParameterContextNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update parameter context not found response has a 4xx status code
func (o *UpdateParameterContextNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update parameter context not found response has a 5xx status code
func (o *UpdateParameterContextNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update parameter context not found response a status code equal to that given
func (o *UpdateParameterContextNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update parameter context not found response
func (o *UpdateParameterContextNotFound) Code() int {
	return 404
}

func (o *UpdateParameterContextNotFound) Error() string {
	return fmt.Sprintf("[PUT /parameter-contexts/{id}][%d] updateParameterContextNotFound ", 404)
}

func (o *UpdateParameterContextNotFound) String() string {
	return fmt.Sprintf("[PUT /parameter-contexts/{id}][%d] updateParameterContextNotFound ", 404)
}

func (o *UpdateParameterContextNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateParameterContextConflict creates a UpdateParameterContextConflict with default headers values
func NewUpdateParameterContextConflict() *UpdateParameterContextConflict {
	return &UpdateParameterContextConflict{}
}

/*
UpdateParameterContextConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type UpdateParameterContextConflict struct {
}

// IsSuccess returns true when this update parameter context conflict response has a 2xx status code
func (o *UpdateParameterContextConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update parameter context conflict response has a 3xx status code
func (o *UpdateParameterContextConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update parameter context conflict response has a 4xx status code
func (o *UpdateParameterContextConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update parameter context conflict response has a 5xx status code
func (o *UpdateParameterContextConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update parameter context conflict response a status code equal to that given
func (o *UpdateParameterContextConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the update parameter context conflict response
func (o *UpdateParameterContextConflict) Code() int {
	return 409
}

func (o *UpdateParameterContextConflict) Error() string {
	return fmt.Sprintf("[PUT /parameter-contexts/{id}][%d] updateParameterContextConflict ", 409)
}

func (o *UpdateParameterContextConflict) String() string {
	return fmt.Sprintf("[PUT /parameter-contexts/{id}][%d] updateParameterContextConflict ", 409)
}

func (o *UpdateParameterContextConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
