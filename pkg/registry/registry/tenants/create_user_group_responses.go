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

/* CreateUserGroupOK describes a response with status code 200, with default header values.

successful operation
*/
type CreateUserGroupOK struct {
	Payload *models.UserGroup
}

func (o *CreateUserGroupOK) Error() string {
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

/* CreateUserGroupBadRequest describes a response with status code 400, with default header values.

NiFi Registry was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type CreateUserGroupBadRequest struct {
}

func (o *CreateUserGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /tenants/user-groups][%d] createUserGroupBadRequest ", 400)
}

func (o *CreateUserGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUserGroupUnauthorized creates a CreateUserGroupUnauthorized with default headers values
func NewCreateUserGroupUnauthorized() *CreateUserGroupUnauthorized {
	return &CreateUserGroupUnauthorized{}
}

/* CreateUserGroupUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type CreateUserGroupUnauthorized struct {
}

func (o *CreateUserGroupUnauthorized) Error() string {
	return fmt.Sprintf("[POST /tenants/user-groups][%d] createUserGroupUnauthorized ", 401)
}

func (o *CreateUserGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUserGroupForbidden creates a CreateUserGroupForbidden with default headers values
func NewCreateUserGroupForbidden() *CreateUserGroupForbidden {
	return &CreateUserGroupForbidden{}
}

/* CreateUserGroupForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type CreateUserGroupForbidden struct {
}

func (o *CreateUserGroupForbidden) Error() string {
	return fmt.Sprintf("[POST /tenants/user-groups][%d] createUserGroupForbidden ", 403)
}

func (o *CreateUserGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUserGroupNotFound creates a CreateUserGroupNotFound with default headers values
func NewCreateUserGroupNotFound() *CreateUserGroupNotFound {
	return &CreateUserGroupNotFound{}
}

/* CreateUserGroupNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type CreateUserGroupNotFound struct {
}

func (o *CreateUserGroupNotFound) Error() string {
	return fmt.Sprintf("[POST /tenants/user-groups][%d] createUserGroupNotFound ", 404)
}

func (o *CreateUserGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUserGroupConflict creates a CreateUserGroupConflict with default headers values
func NewCreateUserGroupConflict() *CreateUserGroupConflict {
	return &CreateUserGroupConflict{}
}

/* CreateUserGroupConflict describes a response with status code 409, with default header values.

NiFi Registry was unable to complete the request because it assumes a server state that is not valid.
*/
type CreateUserGroupConflict struct {
}

func (o *CreateUserGroupConflict) Error() string {
	return fmt.Sprintf("[POST /tenants/user-groups][%d] createUserGroupConflict ", 409)
}

func (o *CreateUserGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
