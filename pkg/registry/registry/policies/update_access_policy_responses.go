// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/registry/models"
)

// UpdateAccessPolicyReader is a Reader for the UpdateAccessPolicy structure.
type UpdateAccessPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAccessPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAccessPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateAccessPolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateAccessPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateAccessPolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateAccessPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateAccessPolicyConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /policies/{id}] updateAccessPolicy", response, response.Code())
	}
}

// NewUpdateAccessPolicyOK creates a UpdateAccessPolicyOK with default headers values
func NewUpdateAccessPolicyOK() *UpdateAccessPolicyOK {
	return &UpdateAccessPolicyOK{}
}

/*
UpdateAccessPolicyOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateAccessPolicyOK struct {
	Payload *models.AccessPolicy
}

// IsSuccess returns true when this update access policy o k response has a 2xx status code
func (o *UpdateAccessPolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update access policy o k response has a 3xx status code
func (o *UpdateAccessPolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update access policy o k response has a 4xx status code
func (o *UpdateAccessPolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update access policy o k response has a 5xx status code
func (o *UpdateAccessPolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update access policy o k response a status code equal to that given
func (o *UpdateAccessPolicyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update access policy o k response
func (o *UpdateAccessPolicyOK) Code() int {
	return 200
}

func (o *UpdateAccessPolicyOK) Error() string {
	return fmt.Sprintf("[PUT /policies/{id}][%d] updateAccessPolicyOK  %+v", 200, o.Payload)
}

func (o *UpdateAccessPolicyOK) String() string {
	return fmt.Sprintf("[PUT /policies/{id}][%d] updateAccessPolicyOK  %+v", 200, o.Payload)
}

func (o *UpdateAccessPolicyOK) GetPayload() *models.AccessPolicy {
	return o.Payload
}

func (o *UpdateAccessPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccessPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAccessPolicyBadRequest creates a UpdateAccessPolicyBadRequest with default headers values
func NewUpdateAccessPolicyBadRequest() *UpdateAccessPolicyBadRequest {
	return &UpdateAccessPolicyBadRequest{}
}

/*
UpdateAccessPolicyBadRequest describes a response with status code 400, with default header values.

NiFi Registry was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type UpdateAccessPolicyBadRequest struct {
}

// IsSuccess returns true when this update access policy bad request response has a 2xx status code
func (o *UpdateAccessPolicyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update access policy bad request response has a 3xx status code
func (o *UpdateAccessPolicyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update access policy bad request response has a 4xx status code
func (o *UpdateAccessPolicyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update access policy bad request response has a 5xx status code
func (o *UpdateAccessPolicyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update access policy bad request response a status code equal to that given
func (o *UpdateAccessPolicyBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update access policy bad request response
func (o *UpdateAccessPolicyBadRequest) Code() int {
	return 400
}

func (o *UpdateAccessPolicyBadRequest) Error() string {
	return fmt.Sprintf("[PUT /policies/{id}][%d] updateAccessPolicyBadRequest ", 400)
}

func (o *UpdateAccessPolicyBadRequest) String() string {
	return fmt.Sprintf("[PUT /policies/{id}][%d] updateAccessPolicyBadRequest ", 400)
}

func (o *UpdateAccessPolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAccessPolicyUnauthorized creates a UpdateAccessPolicyUnauthorized with default headers values
func NewUpdateAccessPolicyUnauthorized() *UpdateAccessPolicyUnauthorized {
	return &UpdateAccessPolicyUnauthorized{}
}

/*
UpdateAccessPolicyUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type UpdateAccessPolicyUnauthorized struct {
}

// IsSuccess returns true when this update access policy unauthorized response has a 2xx status code
func (o *UpdateAccessPolicyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update access policy unauthorized response has a 3xx status code
func (o *UpdateAccessPolicyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update access policy unauthorized response has a 4xx status code
func (o *UpdateAccessPolicyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update access policy unauthorized response has a 5xx status code
func (o *UpdateAccessPolicyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update access policy unauthorized response a status code equal to that given
func (o *UpdateAccessPolicyUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update access policy unauthorized response
func (o *UpdateAccessPolicyUnauthorized) Code() int {
	return 401
}

func (o *UpdateAccessPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /policies/{id}][%d] updateAccessPolicyUnauthorized ", 401)
}

func (o *UpdateAccessPolicyUnauthorized) String() string {
	return fmt.Sprintf("[PUT /policies/{id}][%d] updateAccessPolicyUnauthorized ", 401)
}

func (o *UpdateAccessPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAccessPolicyForbidden creates a UpdateAccessPolicyForbidden with default headers values
func NewUpdateAccessPolicyForbidden() *UpdateAccessPolicyForbidden {
	return &UpdateAccessPolicyForbidden{}
}

/*
UpdateAccessPolicyForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type UpdateAccessPolicyForbidden struct {
}

// IsSuccess returns true when this update access policy forbidden response has a 2xx status code
func (o *UpdateAccessPolicyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update access policy forbidden response has a 3xx status code
func (o *UpdateAccessPolicyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update access policy forbidden response has a 4xx status code
func (o *UpdateAccessPolicyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update access policy forbidden response has a 5xx status code
func (o *UpdateAccessPolicyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update access policy forbidden response a status code equal to that given
func (o *UpdateAccessPolicyForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update access policy forbidden response
func (o *UpdateAccessPolicyForbidden) Code() int {
	return 403
}

func (o *UpdateAccessPolicyForbidden) Error() string {
	return fmt.Sprintf("[PUT /policies/{id}][%d] updateAccessPolicyForbidden ", 403)
}

func (o *UpdateAccessPolicyForbidden) String() string {
	return fmt.Sprintf("[PUT /policies/{id}][%d] updateAccessPolicyForbidden ", 403)
}

func (o *UpdateAccessPolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAccessPolicyNotFound creates a UpdateAccessPolicyNotFound with default headers values
func NewUpdateAccessPolicyNotFound() *UpdateAccessPolicyNotFound {
	return &UpdateAccessPolicyNotFound{}
}

/*
UpdateAccessPolicyNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type UpdateAccessPolicyNotFound struct {
}

// IsSuccess returns true when this update access policy not found response has a 2xx status code
func (o *UpdateAccessPolicyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update access policy not found response has a 3xx status code
func (o *UpdateAccessPolicyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update access policy not found response has a 4xx status code
func (o *UpdateAccessPolicyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update access policy not found response has a 5xx status code
func (o *UpdateAccessPolicyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update access policy not found response a status code equal to that given
func (o *UpdateAccessPolicyNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update access policy not found response
func (o *UpdateAccessPolicyNotFound) Code() int {
	return 404
}

func (o *UpdateAccessPolicyNotFound) Error() string {
	return fmt.Sprintf("[PUT /policies/{id}][%d] updateAccessPolicyNotFound ", 404)
}

func (o *UpdateAccessPolicyNotFound) String() string {
	return fmt.Sprintf("[PUT /policies/{id}][%d] updateAccessPolicyNotFound ", 404)
}

func (o *UpdateAccessPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAccessPolicyConflict creates a UpdateAccessPolicyConflict with default headers values
func NewUpdateAccessPolicyConflict() *UpdateAccessPolicyConflict {
	return &UpdateAccessPolicyConflict{}
}

/*
UpdateAccessPolicyConflict describes a response with status code 409, with default header values.

NiFi Registry was unable to complete the request because it assumes a server state that is not valid. The NiFi Registry might not be configured to use a ConfigurableAccessPolicyProvider.
*/
type UpdateAccessPolicyConflict struct {
}

// IsSuccess returns true when this update access policy conflict response has a 2xx status code
func (o *UpdateAccessPolicyConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update access policy conflict response has a 3xx status code
func (o *UpdateAccessPolicyConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update access policy conflict response has a 4xx status code
func (o *UpdateAccessPolicyConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update access policy conflict response has a 5xx status code
func (o *UpdateAccessPolicyConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update access policy conflict response a status code equal to that given
func (o *UpdateAccessPolicyConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the update access policy conflict response
func (o *UpdateAccessPolicyConflict) Code() int {
	return 409
}

func (o *UpdateAccessPolicyConflict) Error() string {
	return fmt.Sprintf("[PUT /policies/{id}][%d] updateAccessPolicyConflict ", 409)
}

func (o *UpdateAccessPolicyConflict) String() string {
	return fmt.Sprintf("[PUT /policies/{id}][%d] updateAccessPolicyConflict ", 409)
}

func (o *UpdateAccessPolicyConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
