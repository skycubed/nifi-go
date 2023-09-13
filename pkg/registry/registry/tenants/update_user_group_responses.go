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

// UpdateUserGroupReader is a Reader for the UpdateUserGroup structure.
type UpdateUserGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateUserGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateUserGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateUserGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateUserGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateUserGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateUserGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /tenants/user-groups/{id}] updateUserGroup", response, response.Code())
	}
}

// NewUpdateUserGroupOK creates a UpdateUserGroupOK with default headers values
func NewUpdateUserGroupOK() *UpdateUserGroupOK {
	return &UpdateUserGroupOK{}
}

/*
UpdateUserGroupOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateUserGroupOK struct {
	Payload *models.UserGroup
}

// IsSuccess returns true when this update user group o k response has a 2xx status code
func (o *UpdateUserGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update user group o k response has a 3xx status code
func (o *UpdateUserGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user group o k response has a 4xx status code
func (o *UpdateUserGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update user group o k response has a 5xx status code
func (o *UpdateUserGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update user group o k response a status code equal to that given
func (o *UpdateUserGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update user group o k response
func (o *UpdateUserGroupOK) Code() int {
	return 200
}

func (o *UpdateUserGroupOK) Error() string {
	return fmt.Sprintf("[PUT /tenants/user-groups/{id}][%d] updateUserGroupOK  %+v", 200, o.Payload)
}

func (o *UpdateUserGroupOK) String() string {
	return fmt.Sprintf("[PUT /tenants/user-groups/{id}][%d] updateUserGroupOK  %+v", 200, o.Payload)
}

func (o *UpdateUserGroupOK) GetPayload() *models.UserGroup {
	return o.Payload
}

func (o *UpdateUserGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserGroupBadRequest creates a UpdateUserGroupBadRequest with default headers values
func NewUpdateUserGroupBadRequest() *UpdateUserGroupBadRequest {
	return &UpdateUserGroupBadRequest{}
}

/*
UpdateUserGroupBadRequest describes a response with status code 400, with default header values.

NiFi Registry was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type UpdateUserGroupBadRequest struct {
}

// IsSuccess returns true when this update user group bad request response has a 2xx status code
func (o *UpdateUserGroupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update user group bad request response has a 3xx status code
func (o *UpdateUserGroupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user group bad request response has a 4xx status code
func (o *UpdateUserGroupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update user group bad request response has a 5xx status code
func (o *UpdateUserGroupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update user group bad request response a status code equal to that given
func (o *UpdateUserGroupBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update user group bad request response
func (o *UpdateUserGroupBadRequest) Code() int {
	return 400
}

func (o *UpdateUserGroupBadRequest) Error() string {
	return fmt.Sprintf("[PUT /tenants/user-groups/{id}][%d] updateUserGroupBadRequest ", 400)
}

func (o *UpdateUserGroupBadRequest) String() string {
	return fmt.Sprintf("[PUT /tenants/user-groups/{id}][%d] updateUserGroupBadRequest ", 400)
}

func (o *UpdateUserGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUserGroupUnauthorized creates a UpdateUserGroupUnauthorized with default headers values
func NewUpdateUserGroupUnauthorized() *UpdateUserGroupUnauthorized {
	return &UpdateUserGroupUnauthorized{}
}

/*
UpdateUserGroupUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type UpdateUserGroupUnauthorized struct {
}

// IsSuccess returns true when this update user group unauthorized response has a 2xx status code
func (o *UpdateUserGroupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update user group unauthorized response has a 3xx status code
func (o *UpdateUserGroupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user group unauthorized response has a 4xx status code
func (o *UpdateUserGroupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update user group unauthorized response has a 5xx status code
func (o *UpdateUserGroupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update user group unauthorized response a status code equal to that given
func (o *UpdateUserGroupUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update user group unauthorized response
func (o *UpdateUserGroupUnauthorized) Code() int {
	return 401
}

func (o *UpdateUserGroupUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /tenants/user-groups/{id}][%d] updateUserGroupUnauthorized ", 401)
}

func (o *UpdateUserGroupUnauthorized) String() string {
	return fmt.Sprintf("[PUT /tenants/user-groups/{id}][%d] updateUserGroupUnauthorized ", 401)
}

func (o *UpdateUserGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUserGroupForbidden creates a UpdateUserGroupForbidden with default headers values
func NewUpdateUserGroupForbidden() *UpdateUserGroupForbidden {
	return &UpdateUserGroupForbidden{}
}

/*
UpdateUserGroupForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type UpdateUserGroupForbidden struct {
}

// IsSuccess returns true when this update user group forbidden response has a 2xx status code
func (o *UpdateUserGroupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update user group forbidden response has a 3xx status code
func (o *UpdateUserGroupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user group forbidden response has a 4xx status code
func (o *UpdateUserGroupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update user group forbidden response has a 5xx status code
func (o *UpdateUserGroupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update user group forbidden response a status code equal to that given
func (o *UpdateUserGroupForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update user group forbidden response
func (o *UpdateUserGroupForbidden) Code() int {
	return 403
}

func (o *UpdateUserGroupForbidden) Error() string {
	return fmt.Sprintf("[PUT /tenants/user-groups/{id}][%d] updateUserGroupForbidden ", 403)
}

func (o *UpdateUserGroupForbidden) String() string {
	return fmt.Sprintf("[PUT /tenants/user-groups/{id}][%d] updateUserGroupForbidden ", 403)
}

func (o *UpdateUserGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUserGroupNotFound creates a UpdateUserGroupNotFound with default headers values
func NewUpdateUserGroupNotFound() *UpdateUserGroupNotFound {
	return &UpdateUserGroupNotFound{}
}

/*
UpdateUserGroupNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type UpdateUserGroupNotFound struct {
}

// IsSuccess returns true when this update user group not found response has a 2xx status code
func (o *UpdateUserGroupNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update user group not found response has a 3xx status code
func (o *UpdateUserGroupNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user group not found response has a 4xx status code
func (o *UpdateUserGroupNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update user group not found response has a 5xx status code
func (o *UpdateUserGroupNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update user group not found response a status code equal to that given
func (o *UpdateUserGroupNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update user group not found response
func (o *UpdateUserGroupNotFound) Code() int {
	return 404
}

func (o *UpdateUserGroupNotFound) Error() string {
	return fmt.Sprintf("[PUT /tenants/user-groups/{id}][%d] updateUserGroupNotFound ", 404)
}

func (o *UpdateUserGroupNotFound) String() string {
	return fmt.Sprintf("[PUT /tenants/user-groups/{id}][%d] updateUserGroupNotFound ", 404)
}

func (o *UpdateUserGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUserGroupConflict creates a UpdateUserGroupConflict with default headers values
func NewUpdateUserGroupConflict() *UpdateUserGroupConflict {
	return &UpdateUserGroupConflict{}
}

/*
UpdateUserGroupConflict describes a response with status code 409, with default header values.

NiFi Registry was unable to complete the request because it assumes a server state that is not valid.
*/
type UpdateUserGroupConflict struct {
}

// IsSuccess returns true when this update user group conflict response has a 2xx status code
func (o *UpdateUserGroupConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update user group conflict response has a 3xx status code
func (o *UpdateUserGroupConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user group conflict response has a 4xx status code
func (o *UpdateUserGroupConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update user group conflict response has a 5xx status code
func (o *UpdateUserGroupConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update user group conflict response a status code equal to that given
func (o *UpdateUserGroupConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the update user group conflict response
func (o *UpdateUserGroupConflict) Code() int {
	return 409
}

func (o *UpdateUserGroupConflict) Error() string {
	return fmt.Sprintf("[PUT /tenants/user-groups/{id}][%d] updateUserGroupConflict ", 409)
}

func (o *UpdateUserGroupConflict) String() string {
	return fmt.Sprintf("[PUT /tenants/user-groups/{id}][%d] updateUserGroupConflict ", 409)
}

func (o *UpdateUserGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
