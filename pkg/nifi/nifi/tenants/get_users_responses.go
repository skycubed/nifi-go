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

// GetUsersReader is a Reader for the GetUsers structure.
type GetUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUsersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUsersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUsersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetUsersConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /tenants/users] getUsers", response, response.Code())
	}
}

// NewGetUsersOK creates a GetUsersOK with default headers values
func NewGetUsersOK() *GetUsersOK {
	return &GetUsersOK{}
}

/*
GetUsersOK describes a response with status code 200, with default header values.

successful operation
*/
type GetUsersOK struct {
	Payload *models.UsersEntity
}

// IsSuccess returns true when this get users o k response has a 2xx status code
func (o *GetUsersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get users o k response has a 3xx status code
func (o *GetUsersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get users o k response has a 4xx status code
func (o *GetUsersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get users o k response has a 5xx status code
func (o *GetUsersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get users o k response a status code equal to that given
func (o *GetUsersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get users o k response
func (o *GetUsersOK) Code() int {
	return 200
}

func (o *GetUsersOK) Error() string {
	return fmt.Sprintf("[GET /tenants/users][%d] getUsersOK  %+v", 200, o.Payload)
}

func (o *GetUsersOK) String() string {
	return fmt.Sprintf("[GET /tenants/users][%d] getUsersOK  %+v", 200, o.Payload)
}

func (o *GetUsersOK) GetPayload() *models.UsersEntity {
	return o.Payload
}

func (o *GetUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UsersEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersBadRequest creates a GetUsersBadRequest with default headers values
func NewGetUsersBadRequest() *GetUsersBadRequest {
	return &GetUsersBadRequest{}
}

/*
GetUsersBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetUsersBadRequest struct {
}

// IsSuccess returns true when this get users bad request response has a 2xx status code
func (o *GetUsersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get users bad request response has a 3xx status code
func (o *GetUsersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get users bad request response has a 4xx status code
func (o *GetUsersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get users bad request response has a 5xx status code
func (o *GetUsersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get users bad request response a status code equal to that given
func (o *GetUsersBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get users bad request response
func (o *GetUsersBadRequest) Code() int {
	return 400
}

func (o *GetUsersBadRequest) Error() string {
	return fmt.Sprintf("[GET /tenants/users][%d] getUsersBadRequest ", 400)
}

func (o *GetUsersBadRequest) String() string {
	return fmt.Sprintf("[GET /tenants/users][%d] getUsersBadRequest ", 400)
}

func (o *GetUsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsersUnauthorized creates a GetUsersUnauthorized with default headers values
func NewGetUsersUnauthorized() *GetUsersUnauthorized {
	return &GetUsersUnauthorized{}
}

/*
GetUsersUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetUsersUnauthorized struct {
}

// IsSuccess returns true when this get users unauthorized response has a 2xx status code
func (o *GetUsersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get users unauthorized response has a 3xx status code
func (o *GetUsersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get users unauthorized response has a 4xx status code
func (o *GetUsersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get users unauthorized response has a 5xx status code
func (o *GetUsersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get users unauthorized response a status code equal to that given
func (o *GetUsersUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get users unauthorized response
func (o *GetUsersUnauthorized) Code() int {
	return 401
}

func (o *GetUsersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /tenants/users][%d] getUsersUnauthorized ", 401)
}

func (o *GetUsersUnauthorized) String() string {
	return fmt.Sprintf("[GET /tenants/users][%d] getUsersUnauthorized ", 401)
}

func (o *GetUsersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsersForbidden creates a GetUsersForbidden with default headers values
func NewGetUsersForbidden() *GetUsersForbidden {
	return &GetUsersForbidden{}
}

/*
GetUsersForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetUsersForbidden struct {
}

// IsSuccess returns true when this get users forbidden response has a 2xx status code
func (o *GetUsersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get users forbidden response has a 3xx status code
func (o *GetUsersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get users forbidden response has a 4xx status code
func (o *GetUsersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get users forbidden response has a 5xx status code
func (o *GetUsersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get users forbidden response a status code equal to that given
func (o *GetUsersForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get users forbidden response
func (o *GetUsersForbidden) Code() int {
	return 403
}

func (o *GetUsersForbidden) Error() string {
	return fmt.Sprintf("[GET /tenants/users][%d] getUsersForbidden ", 403)
}

func (o *GetUsersForbidden) String() string {
	return fmt.Sprintf("[GET /tenants/users][%d] getUsersForbidden ", 403)
}

func (o *GetUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsersNotFound creates a GetUsersNotFound with default headers values
func NewGetUsersNotFound() *GetUsersNotFound {
	return &GetUsersNotFound{}
}

/*
GetUsersNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetUsersNotFound struct {
}

// IsSuccess returns true when this get users not found response has a 2xx status code
func (o *GetUsersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get users not found response has a 3xx status code
func (o *GetUsersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get users not found response has a 4xx status code
func (o *GetUsersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get users not found response has a 5xx status code
func (o *GetUsersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get users not found response a status code equal to that given
func (o *GetUsersNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get users not found response
func (o *GetUsersNotFound) Code() int {
	return 404
}

func (o *GetUsersNotFound) Error() string {
	return fmt.Sprintf("[GET /tenants/users][%d] getUsersNotFound ", 404)
}

func (o *GetUsersNotFound) String() string {
	return fmt.Sprintf("[GET /tenants/users][%d] getUsersNotFound ", 404)
}

func (o *GetUsersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsersConflict creates a GetUsersConflict with default headers values
func NewGetUsersConflict() *GetUsersConflict {
	return &GetUsersConflict{}
}

/*
GetUsersConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetUsersConflict struct {
}

// IsSuccess returns true when this get users conflict response has a 2xx status code
func (o *GetUsersConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get users conflict response has a 3xx status code
func (o *GetUsersConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get users conflict response has a 4xx status code
func (o *GetUsersConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get users conflict response has a 5xx status code
func (o *GetUsersConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get users conflict response a status code equal to that given
func (o *GetUsersConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the get users conflict response
func (o *GetUsersConflict) Code() int {
	return 409
}

func (o *GetUsersConflict) Error() string {
	return fmt.Sprintf("[GET /tenants/users][%d] getUsersConflict ", 409)
}

func (o *GetUsersConflict) String() string {
	return fmt.Sprintf("[GET /tenants/users][%d] getUsersConflict ", 409)
}

func (o *GetUsersConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
