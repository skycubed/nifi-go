// Code generated by go-swagger; DO NOT EDIT.

package tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// GetUserGroupReader is a Reader for the GetUserGroup structure.
type GetUserGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUserGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetUserGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserGroupOK creates a GetUserGroupOK with default headers values
func NewGetUserGroupOK() *GetUserGroupOK {
	return &GetUserGroupOK{}
}

/* GetUserGroupOK describes a response with status code 200, with default header values.

successful operation
*/
type GetUserGroupOK struct {
	Payload *models.UserGroupEntity
}

func (o *GetUserGroupOK) Error() string {
	return fmt.Sprintf("[GET /tenants/user-groups/{id}][%d] getUserGroupOK  %+v", 200, o.Payload)
}
func (o *GetUserGroupOK) GetPayload() *models.UserGroupEntity {
	return o.Payload
}

func (o *GetUserGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserGroupEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGroupBadRequest creates a GetUserGroupBadRequest with default headers values
func NewGetUserGroupBadRequest() *GetUserGroupBadRequest {
	return &GetUserGroupBadRequest{}
}

/* GetUserGroupBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetUserGroupBadRequest struct {
}

func (o *GetUserGroupBadRequest) Error() string {
	return fmt.Sprintf("[GET /tenants/user-groups/{id}][%d] getUserGroupBadRequest ", 400)
}

func (o *GetUserGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserGroupUnauthorized creates a GetUserGroupUnauthorized with default headers values
func NewGetUserGroupUnauthorized() *GetUserGroupUnauthorized {
	return &GetUserGroupUnauthorized{}
}

/* GetUserGroupUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetUserGroupUnauthorized struct {
}

func (o *GetUserGroupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /tenants/user-groups/{id}][%d] getUserGroupUnauthorized ", 401)
}

func (o *GetUserGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserGroupForbidden creates a GetUserGroupForbidden with default headers values
func NewGetUserGroupForbidden() *GetUserGroupForbidden {
	return &GetUserGroupForbidden{}
}

/* GetUserGroupForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetUserGroupForbidden struct {
}

func (o *GetUserGroupForbidden) Error() string {
	return fmt.Sprintf("[GET /tenants/user-groups/{id}][%d] getUserGroupForbidden ", 403)
}

func (o *GetUserGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserGroupNotFound creates a GetUserGroupNotFound with default headers values
func NewGetUserGroupNotFound() *GetUserGroupNotFound {
	return &GetUserGroupNotFound{}
}

/* GetUserGroupNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetUserGroupNotFound struct {
}

func (o *GetUserGroupNotFound) Error() string {
	return fmt.Sprintf("[GET /tenants/user-groups/{id}][%d] getUserGroupNotFound ", 404)
}

func (o *GetUserGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserGroupConflict creates a GetUserGroupConflict with default headers values
func NewGetUserGroupConflict() *GetUserGroupConflict {
	return &GetUserGroupConflict{}
}

/* GetUserGroupConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetUserGroupConflict struct {
}

func (o *GetUserGroupConflict) Error() string {
	return fmt.Sprintf("[GET /tenants/user-groups/{id}][%d] getUserGroupConflict ", 409)
}

func (o *GetUserGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
