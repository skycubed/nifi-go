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

// RemoveUserGroupReader is a Reader for the RemoveUserGroup structure.
type RemoveUserGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveUserGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveUserGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveUserGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRemoveUserGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRemoveUserGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRemoveUserGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewRemoveUserGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRemoveUserGroupOK creates a RemoveUserGroupOK with default headers values
func NewRemoveUserGroupOK() *RemoveUserGroupOK {
	return &RemoveUserGroupOK{}
}

/*
RemoveUserGroupOK describes a response with status code 200, with default header values.

successful operation
*/
type RemoveUserGroupOK struct {
	Payload *models.UserGroup
}

// IsSuccess returns true when this remove user group o k response has a 2xx status code
func (o *RemoveUserGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this remove user group o k response has a 3xx status code
func (o *RemoveUserGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove user group o k response has a 4xx status code
func (o *RemoveUserGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove user group o k response has a 5xx status code
func (o *RemoveUserGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this remove user group o k response a status code equal to that given
func (o *RemoveUserGroupOK) IsCode(code int) bool {
	return code == 200
}

func (o *RemoveUserGroupOK) Error() string {
	return fmt.Sprintf("[DELETE /tenants/user-groups/{id}][%d] removeUserGroupOK  %+v", 200, o.Payload)
}

func (o *RemoveUserGroupOK) String() string {
	return fmt.Sprintf("[DELETE /tenants/user-groups/{id}][%d] removeUserGroupOK  %+v", 200, o.Payload)
}

func (o *RemoveUserGroupOK) GetPayload() *models.UserGroup {
	return o.Payload
}

func (o *RemoveUserGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveUserGroupBadRequest creates a RemoveUserGroupBadRequest with default headers values
func NewRemoveUserGroupBadRequest() *RemoveUserGroupBadRequest {
	return &RemoveUserGroupBadRequest{}
}

/*
RemoveUserGroupBadRequest describes a response with status code 400, with default header values.

NiFi Registry was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type RemoveUserGroupBadRequest struct {
}

// IsSuccess returns true when this remove user group bad request response has a 2xx status code
func (o *RemoveUserGroupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove user group bad request response has a 3xx status code
func (o *RemoveUserGroupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove user group bad request response has a 4xx status code
func (o *RemoveUserGroupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove user group bad request response has a 5xx status code
func (o *RemoveUserGroupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this remove user group bad request response a status code equal to that given
func (o *RemoveUserGroupBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *RemoveUserGroupBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /tenants/user-groups/{id}][%d] removeUserGroupBadRequest ", 400)
}

func (o *RemoveUserGroupBadRequest) String() string {
	return fmt.Sprintf("[DELETE /tenants/user-groups/{id}][%d] removeUserGroupBadRequest ", 400)
}

func (o *RemoveUserGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveUserGroupUnauthorized creates a RemoveUserGroupUnauthorized with default headers values
func NewRemoveUserGroupUnauthorized() *RemoveUserGroupUnauthorized {
	return &RemoveUserGroupUnauthorized{}
}

/*
RemoveUserGroupUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type RemoveUserGroupUnauthorized struct {
}

// IsSuccess returns true when this remove user group unauthorized response has a 2xx status code
func (o *RemoveUserGroupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove user group unauthorized response has a 3xx status code
func (o *RemoveUserGroupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove user group unauthorized response has a 4xx status code
func (o *RemoveUserGroupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove user group unauthorized response has a 5xx status code
func (o *RemoveUserGroupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this remove user group unauthorized response a status code equal to that given
func (o *RemoveUserGroupUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *RemoveUserGroupUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /tenants/user-groups/{id}][%d] removeUserGroupUnauthorized ", 401)
}

func (o *RemoveUserGroupUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /tenants/user-groups/{id}][%d] removeUserGroupUnauthorized ", 401)
}

func (o *RemoveUserGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveUserGroupForbidden creates a RemoveUserGroupForbidden with default headers values
func NewRemoveUserGroupForbidden() *RemoveUserGroupForbidden {
	return &RemoveUserGroupForbidden{}
}

/*
RemoveUserGroupForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type RemoveUserGroupForbidden struct {
}

// IsSuccess returns true when this remove user group forbidden response has a 2xx status code
func (o *RemoveUserGroupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove user group forbidden response has a 3xx status code
func (o *RemoveUserGroupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove user group forbidden response has a 4xx status code
func (o *RemoveUserGroupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove user group forbidden response has a 5xx status code
func (o *RemoveUserGroupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this remove user group forbidden response a status code equal to that given
func (o *RemoveUserGroupForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *RemoveUserGroupForbidden) Error() string {
	return fmt.Sprintf("[DELETE /tenants/user-groups/{id}][%d] removeUserGroupForbidden ", 403)
}

func (o *RemoveUserGroupForbidden) String() string {
	return fmt.Sprintf("[DELETE /tenants/user-groups/{id}][%d] removeUserGroupForbidden ", 403)
}

func (o *RemoveUserGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveUserGroupNotFound creates a RemoveUserGroupNotFound with default headers values
func NewRemoveUserGroupNotFound() *RemoveUserGroupNotFound {
	return &RemoveUserGroupNotFound{}
}

/*
RemoveUserGroupNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type RemoveUserGroupNotFound struct {
}

// IsSuccess returns true when this remove user group not found response has a 2xx status code
func (o *RemoveUserGroupNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove user group not found response has a 3xx status code
func (o *RemoveUserGroupNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove user group not found response has a 4xx status code
func (o *RemoveUserGroupNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove user group not found response has a 5xx status code
func (o *RemoveUserGroupNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this remove user group not found response a status code equal to that given
func (o *RemoveUserGroupNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *RemoveUserGroupNotFound) Error() string {
	return fmt.Sprintf("[DELETE /tenants/user-groups/{id}][%d] removeUserGroupNotFound ", 404)
}

func (o *RemoveUserGroupNotFound) String() string {
	return fmt.Sprintf("[DELETE /tenants/user-groups/{id}][%d] removeUserGroupNotFound ", 404)
}

func (o *RemoveUserGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveUserGroupConflict creates a RemoveUserGroupConflict with default headers values
func NewRemoveUserGroupConflict() *RemoveUserGroupConflict {
	return &RemoveUserGroupConflict{}
}

/*
RemoveUserGroupConflict describes a response with status code 409, with default header values.

NiFi Registry was unable to complete the request because it assumes a server state that is not valid.
*/
type RemoveUserGroupConflict struct {
}

// IsSuccess returns true when this remove user group conflict response has a 2xx status code
func (o *RemoveUserGroupConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove user group conflict response has a 3xx status code
func (o *RemoveUserGroupConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove user group conflict response has a 4xx status code
func (o *RemoveUserGroupConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove user group conflict response has a 5xx status code
func (o *RemoveUserGroupConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this remove user group conflict response a status code equal to that given
func (o *RemoveUserGroupConflict) IsCode(code int) bool {
	return code == 409
}

func (o *RemoveUserGroupConflict) Error() string {
	return fmt.Sprintf("[DELETE /tenants/user-groups/{id}][%d] removeUserGroupConflict ", 409)
}

func (o *RemoveUserGroupConflict) String() string {
	return fmt.Sprintf("[DELETE /tenants/user-groups/{id}][%d] removeUserGroupConflict ", 409)
}

func (o *RemoveUserGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
