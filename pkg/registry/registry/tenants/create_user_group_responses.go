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

// CreateUserGroupReader is a Reader for the CreateUserGroup structure.
type CreateUserGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateUserGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateUserGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateUserGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateUserGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateUserGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateUserGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateUserGroupOK creates a CreateUserGroupOK with default headers values
func NewCreateUserGroupOK() *CreateUserGroupOK {
	return &CreateUserGroupOK{}
}

/*
CreateUserGroupOK describes a response with status code 200, with default header values.

successful operation
*/
type CreateUserGroupOK struct {
	Payload *models.UserGroup
}

// IsSuccess returns true when this create user group o k response has a 2xx status code
func (o *CreateUserGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create user group o k response has a 3xx status code
func (o *CreateUserGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user group o k response has a 4xx status code
func (o *CreateUserGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create user group o k response has a 5xx status code
func (o *CreateUserGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create user group o k response a status code equal to that given
func (o *CreateUserGroupOK) IsCode(code int) bool {
	return code == 200
}

func (o *CreateUserGroupOK) Error() string {
	return fmt.Sprintf("[POST /tenants/user-groups][%d] createUserGroupOK  %+v", 200, o.Payload)
}

func (o *CreateUserGroupOK) String() string {
	return fmt.Sprintf("[POST /tenants/user-groups][%d] createUserGroupOK  %+v", 200, o.Payload)
}

func (o *CreateUserGroupOK) GetPayload() *models.UserGroup {
	return o.Payload
}

func (o *CreateUserGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserGroupBadRequest creates a CreateUserGroupBadRequest with default headers values
func NewCreateUserGroupBadRequest() *CreateUserGroupBadRequest {
	return &CreateUserGroupBadRequest{}
}

/*
CreateUserGroupBadRequest describes a response with status code 400, with default header values.

NiFi Registry was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type CreateUserGroupBadRequest struct {
}

// IsSuccess returns true when this create user group bad request response has a 2xx status code
func (o *CreateUserGroupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user group bad request response has a 3xx status code
func (o *CreateUserGroupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user group bad request response has a 4xx status code
func (o *CreateUserGroupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create user group bad request response has a 5xx status code
func (o *CreateUserGroupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create user group bad request response a status code equal to that given
func (o *CreateUserGroupBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateUserGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /tenants/user-groups][%d] createUserGroupBadRequest ", 400)
}

func (o *CreateUserGroupBadRequest) String() string {
	return fmt.Sprintf("[POST /tenants/user-groups][%d] createUserGroupBadRequest ", 400)
}

func (o *CreateUserGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUserGroupUnauthorized creates a CreateUserGroupUnauthorized with default headers values
func NewCreateUserGroupUnauthorized() *CreateUserGroupUnauthorized {
	return &CreateUserGroupUnauthorized{}
}

/*
CreateUserGroupUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type CreateUserGroupUnauthorized struct {
}

// IsSuccess returns true when this create user group unauthorized response has a 2xx status code
func (o *CreateUserGroupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user group unauthorized response has a 3xx status code
func (o *CreateUserGroupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user group unauthorized response has a 4xx status code
func (o *CreateUserGroupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create user group unauthorized response has a 5xx status code
func (o *CreateUserGroupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create user group unauthorized response a status code equal to that given
func (o *CreateUserGroupUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateUserGroupUnauthorized) Error() string {
	return fmt.Sprintf("[POST /tenants/user-groups][%d] createUserGroupUnauthorized ", 401)
}

func (o *CreateUserGroupUnauthorized) String() string {
	return fmt.Sprintf("[POST /tenants/user-groups][%d] createUserGroupUnauthorized ", 401)
}

func (o *CreateUserGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUserGroupForbidden creates a CreateUserGroupForbidden with default headers values
func NewCreateUserGroupForbidden() *CreateUserGroupForbidden {
	return &CreateUserGroupForbidden{}
}

/*
CreateUserGroupForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type CreateUserGroupForbidden struct {
}

// IsSuccess returns true when this create user group forbidden response has a 2xx status code
func (o *CreateUserGroupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user group forbidden response has a 3xx status code
func (o *CreateUserGroupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user group forbidden response has a 4xx status code
func (o *CreateUserGroupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create user group forbidden response has a 5xx status code
func (o *CreateUserGroupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create user group forbidden response a status code equal to that given
func (o *CreateUserGroupForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateUserGroupForbidden) Error() string {
	return fmt.Sprintf("[POST /tenants/user-groups][%d] createUserGroupForbidden ", 403)
}

func (o *CreateUserGroupForbidden) String() string {
	return fmt.Sprintf("[POST /tenants/user-groups][%d] createUserGroupForbidden ", 403)
}

func (o *CreateUserGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUserGroupNotFound creates a CreateUserGroupNotFound with default headers values
func NewCreateUserGroupNotFound() *CreateUserGroupNotFound {
	return &CreateUserGroupNotFound{}
}

/*
CreateUserGroupNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type CreateUserGroupNotFound struct {
}

// IsSuccess returns true when this create user group not found response has a 2xx status code
func (o *CreateUserGroupNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user group not found response has a 3xx status code
func (o *CreateUserGroupNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user group not found response has a 4xx status code
func (o *CreateUserGroupNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create user group not found response has a 5xx status code
func (o *CreateUserGroupNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create user group not found response a status code equal to that given
func (o *CreateUserGroupNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CreateUserGroupNotFound) Error() string {
	return fmt.Sprintf("[POST /tenants/user-groups][%d] createUserGroupNotFound ", 404)
}

func (o *CreateUserGroupNotFound) String() string {
	return fmt.Sprintf("[POST /tenants/user-groups][%d] createUserGroupNotFound ", 404)
}

func (o *CreateUserGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUserGroupConflict creates a CreateUserGroupConflict with default headers values
func NewCreateUserGroupConflict() *CreateUserGroupConflict {
	return &CreateUserGroupConflict{}
}

/*
CreateUserGroupConflict describes a response with status code 409, with default header values.

NiFi Registry was unable to complete the request because it assumes a server state that is not valid.
*/
type CreateUserGroupConflict struct {
}

// IsSuccess returns true when this create user group conflict response has a 2xx status code
func (o *CreateUserGroupConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user group conflict response has a 3xx status code
func (o *CreateUserGroupConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user group conflict response has a 4xx status code
func (o *CreateUserGroupConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create user group conflict response has a 5xx status code
func (o *CreateUserGroupConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create user group conflict response a status code equal to that given
func (o *CreateUserGroupConflict) IsCode(code int) bool {
	return code == 409
}

func (o *CreateUserGroupConflict) Error() string {
	return fmt.Sprintf("[POST /tenants/user-groups][%d] createUserGroupConflict ", 409)
}

func (o *CreateUserGroupConflict) String() string {
	return fmt.Sprintf("[POST /tenants/user-groups][%d] createUserGroupConflict ", 409)
}

func (o *CreateUserGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
