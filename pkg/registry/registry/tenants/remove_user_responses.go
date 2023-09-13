// Code generated by go-swagger; DO NOT EDIT.

package tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/registry/models"
)

// RemoveUserReader is a Reader for the RemoveUser structure.
type RemoveUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRemoveUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRemoveUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRemoveUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewRemoveUserConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /tenants/users/{id}] removeUser", response, response.Code())
	}
}

// NewRemoveUserOK creates a RemoveUserOK with default headers values
func NewRemoveUserOK() *RemoveUserOK {
	return &RemoveUserOK{}
}

/*
RemoveUserOK describes a response with status code 200, with default header values.

successful operation
*/
type RemoveUserOK struct {
	Payload *models.User
}

// IsSuccess returns true when this remove user o k response has a 2xx status code
func (o *RemoveUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this remove user o k response has a 3xx status code
func (o *RemoveUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove user o k response has a 4xx status code
func (o *RemoveUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove user o k response has a 5xx status code
func (o *RemoveUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this remove user o k response a status code equal to that given
func (o *RemoveUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the remove user o k response
func (o *RemoveUserOK) Code() int {
	return 200
}

func (o *RemoveUserOK) Error() string {
	return fmt.Sprintf("[DELETE /tenants/users/{id}][%d] removeUserOK  %+v", 200, o.Payload)
}

func (o *RemoveUserOK) String() string {
	return fmt.Sprintf("[DELETE /tenants/users/{id}][%d] removeUserOK  %+v", 200, o.Payload)
}

func (o *RemoveUserOK) GetPayload() *models.User {
	return o.Payload
}

func (o *RemoveUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveUserBadRequest creates a RemoveUserBadRequest with default headers values
func NewRemoveUserBadRequest() *RemoveUserBadRequest {
	return &RemoveUserBadRequest{}
}

/*
RemoveUserBadRequest describes a response with status code 400, with default header values.

NiFi Registry was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type RemoveUserBadRequest struct {
}

// IsSuccess returns true when this remove user bad request response has a 2xx status code
func (o *RemoveUserBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove user bad request response has a 3xx status code
func (o *RemoveUserBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove user bad request response has a 4xx status code
func (o *RemoveUserBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove user bad request response has a 5xx status code
func (o *RemoveUserBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this remove user bad request response a status code equal to that given
func (o *RemoveUserBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the remove user bad request response
func (o *RemoveUserBadRequest) Code() int {
	return 400
}

func (o *RemoveUserBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /tenants/users/{id}][%d] removeUserBadRequest ", 400)
}

func (o *RemoveUserBadRequest) String() string {
	return fmt.Sprintf("[DELETE /tenants/users/{id}][%d] removeUserBadRequest ", 400)
}

func (o *RemoveUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveUserUnauthorized creates a RemoveUserUnauthorized with default headers values
func NewRemoveUserUnauthorized() *RemoveUserUnauthorized {
	return &RemoveUserUnauthorized{}
}

/*
RemoveUserUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type RemoveUserUnauthorized struct {
}

// IsSuccess returns true when this remove user unauthorized response has a 2xx status code
func (o *RemoveUserUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove user unauthorized response has a 3xx status code
func (o *RemoveUserUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove user unauthorized response has a 4xx status code
func (o *RemoveUserUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove user unauthorized response has a 5xx status code
func (o *RemoveUserUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this remove user unauthorized response a status code equal to that given
func (o *RemoveUserUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the remove user unauthorized response
func (o *RemoveUserUnauthorized) Code() int {
	return 401
}

func (o *RemoveUserUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /tenants/users/{id}][%d] removeUserUnauthorized ", 401)
}

func (o *RemoveUserUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /tenants/users/{id}][%d] removeUserUnauthorized ", 401)
}

func (o *RemoveUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveUserForbidden creates a RemoveUserForbidden with default headers values
func NewRemoveUserForbidden() *RemoveUserForbidden {
	return &RemoveUserForbidden{}
}

/*
RemoveUserForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type RemoveUserForbidden struct {
}

// IsSuccess returns true when this remove user forbidden response has a 2xx status code
func (o *RemoveUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove user forbidden response has a 3xx status code
func (o *RemoveUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove user forbidden response has a 4xx status code
func (o *RemoveUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove user forbidden response has a 5xx status code
func (o *RemoveUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this remove user forbidden response a status code equal to that given
func (o *RemoveUserForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the remove user forbidden response
func (o *RemoveUserForbidden) Code() int {
	return 403
}

func (o *RemoveUserForbidden) Error() string {
	return fmt.Sprintf("[DELETE /tenants/users/{id}][%d] removeUserForbidden ", 403)
}

func (o *RemoveUserForbidden) String() string {
	return fmt.Sprintf("[DELETE /tenants/users/{id}][%d] removeUserForbidden ", 403)
}

func (o *RemoveUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveUserNotFound creates a RemoveUserNotFound with default headers values
func NewRemoveUserNotFound() *RemoveUserNotFound {
	return &RemoveUserNotFound{}
}

/*
RemoveUserNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type RemoveUserNotFound struct {
}

// IsSuccess returns true when this remove user not found response has a 2xx status code
func (o *RemoveUserNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove user not found response has a 3xx status code
func (o *RemoveUserNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove user not found response has a 4xx status code
func (o *RemoveUserNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove user not found response has a 5xx status code
func (o *RemoveUserNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this remove user not found response a status code equal to that given
func (o *RemoveUserNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the remove user not found response
func (o *RemoveUserNotFound) Code() int {
	return 404
}

func (o *RemoveUserNotFound) Error() string {
	return fmt.Sprintf("[DELETE /tenants/users/{id}][%d] removeUserNotFound ", 404)
}

func (o *RemoveUserNotFound) String() string {
	return fmt.Sprintf("[DELETE /tenants/users/{id}][%d] removeUserNotFound ", 404)
}

func (o *RemoveUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveUserConflict creates a RemoveUserConflict with default headers values
func NewRemoveUserConflict() *RemoveUserConflict {
	return &RemoveUserConflict{}
}

/*
RemoveUserConflict describes a response with status code 409, with default header values.

NiFi Registry was unable to complete the request because it assumes a server state that is not valid.
*/
type RemoveUserConflict struct {
}

// IsSuccess returns true when this remove user conflict response has a 2xx status code
func (o *RemoveUserConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove user conflict response has a 3xx status code
func (o *RemoveUserConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove user conflict response has a 4xx status code
func (o *RemoveUserConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove user conflict response has a 5xx status code
func (o *RemoveUserConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this remove user conflict response a status code equal to that given
func (o *RemoveUserConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the remove user conflict response
func (o *RemoveUserConflict) Code() int {
	return 409
}

func (o *RemoveUserConflict) Error() string {
	return fmt.Sprintf("[DELETE /tenants/users/{id}][%d] removeUserConflict ", 409)
}

func (o *RemoveUserConflict) String() string {
	return fmt.Sprintf("[DELETE /tenants/users/{id}][%d] removeUserConflict ", 409)
}

func (o *RemoveUserConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
