// Code generated by go-swagger; DO NOT EDIT.

package access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateAccessTokenReader is a Reader for the CreateAccessToken structure.
type CreateAccessTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAccessTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateAccessTokenCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateAccessTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateAccessTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateAccessTokenConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateAccessTokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /access/token] createAccessToken", response, response.Code())
	}
}

// NewCreateAccessTokenCreated creates a CreateAccessTokenCreated with default headers values
func NewCreateAccessTokenCreated() *CreateAccessTokenCreated {
	return &CreateAccessTokenCreated{}
}

/*
CreateAccessTokenCreated describes a response with status code 201, with default header values.

successful operation
*/
type CreateAccessTokenCreated struct {
	Payload string
}

// IsSuccess returns true when this create access token created response has a 2xx status code
func (o *CreateAccessTokenCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create access token created response has a 3xx status code
func (o *CreateAccessTokenCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create access token created response has a 4xx status code
func (o *CreateAccessTokenCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create access token created response has a 5xx status code
func (o *CreateAccessTokenCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create access token created response a status code equal to that given
func (o *CreateAccessTokenCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create access token created response
func (o *CreateAccessTokenCreated) Code() int {
	return 201
}

func (o *CreateAccessTokenCreated) Error() string {
	return fmt.Sprintf("[POST /access/token][%d] createAccessTokenCreated  %+v", 201, o.Payload)
}

func (o *CreateAccessTokenCreated) String() string {
	return fmt.Sprintf("[POST /access/token][%d] createAccessTokenCreated  %+v", 201, o.Payload)
}

func (o *CreateAccessTokenCreated) GetPayload() string {
	return o.Payload
}

func (o *CreateAccessTokenCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAccessTokenBadRequest creates a CreateAccessTokenBadRequest with default headers values
func NewCreateAccessTokenBadRequest() *CreateAccessTokenBadRequest {
	return &CreateAccessTokenBadRequest{}
}

/*
CreateAccessTokenBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type CreateAccessTokenBadRequest struct {
}

// IsSuccess returns true when this create access token bad request response has a 2xx status code
func (o *CreateAccessTokenBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create access token bad request response has a 3xx status code
func (o *CreateAccessTokenBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create access token bad request response has a 4xx status code
func (o *CreateAccessTokenBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create access token bad request response has a 5xx status code
func (o *CreateAccessTokenBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create access token bad request response a status code equal to that given
func (o *CreateAccessTokenBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create access token bad request response
func (o *CreateAccessTokenBadRequest) Code() int {
	return 400
}

func (o *CreateAccessTokenBadRequest) Error() string {
	return fmt.Sprintf("[POST /access/token][%d] createAccessTokenBadRequest ", 400)
}

func (o *CreateAccessTokenBadRequest) String() string {
	return fmt.Sprintf("[POST /access/token][%d] createAccessTokenBadRequest ", 400)
}

func (o *CreateAccessTokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAccessTokenForbidden creates a CreateAccessTokenForbidden with default headers values
func NewCreateAccessTokenForbidden() *CreateAccessTokenForbidden {
	return &CreateAccessTokenForbidden{}
}

/*
CreateAccessTokenForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type CreateAccessTokenForbidden struct {
}

// IsSuccess returns true when this create access token forbidden response has a 2xx status code
func (o *CreateAccessTokenForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create access token forbidden response has a 3xx status code
func (o *CreateAccessTokenForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create access token forbidden response has a 4xx status code
func (o *CreateAccessTokenForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create access token forbidden response has a 5xx status code
func (o *CreateAccessTokenForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create access token forbidden response a status code equal to that given
func (o *CreateAccessTokenForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create access token forbidden response
func (o *CreateAccessTokenForbidden) Code() int {
	return 403
}

func (o *CreateAccessTokenForbidden) Error() string {
	return fmt.Sprintf("[POST /access/token][%d] createAccessTokenForbidden ", 403)
}

func (o *CreateAccessTokenForbidden) String() string {
	return fmt.Sprintf("[POST /access/token][%d] createAccessTokenForbidden ", 403)
}

func (o *CreateAccessTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAccessTokenConflict creates a CreateAccessTokenConflict with default headers values
func NewCreateAccessTokenConflict() *CreateAccessTokenConflict {
	return &CreateAccessTokenConflict{}
}

/*
CreateAccessTokenConflict describes a response with status code 409, with default header values.

Unable to create access token because NiFi is not in the appropriate state. (i.e. may not be configured to support username/password login.
*/
type CreateAccessTokenConflict struct {
}

// IsSuccess returns true when this create access token conflict response has a 2xx status code
func (o *CreateAccessTokenConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create access token conflict response has a 3xx status code
func (o *CreateAccessTokenConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create access token conflict response has a 4xx status code
func (o *CreateAccessTokenConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create access token conflict response has a 5xx status code
func (o *CreateAccessTokenConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create access token conflict response a status code equal to that given
func (o *CreateAccessTokenConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create access token conflict response
func (o *CreateAccessTokenConflict) Code() int {
	return 409
}

func (o *CreateAccessTokenConflict) Error() string {
	return fmt.Sprintf("[POST /access/token][%d] createAccessTokenConflict ", 409)
}

func (o *CreateAccessTokenConflict) String() string {
	return fmt.Sprintf("[POST /access/token][%d] createAccessTokenConflict ", 409)
}

func (o *CreateAccessTokenConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAccessTokenInternalServerError creates a CreateAccessTokenInternalServerError with default headers values
func NewCreateAccessTokenInternalServerError() *CreateAccessTokenInternalServerError {
	return &CreateAccessTokenInternalServerError{}
}

/*
CreateAccessTokenInternalServerError describes a response with status code 500, with default header values.

Unable to create access token because an unexpected error occurred.
*/
type CreateAccessTokenInternalServerError struct {
}

// IsSuccess returns true when this create access token internal server error response has a 2xx status code
func (o *CreateAccessTokenInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create access token internal server error response has a 3xx status code
func (o *CreateAccessTokenInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create access token internal server error response has a 4xx status code
func (o *CreateAccessTokenInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create access token internal server error response has a 5xx status code
func (o *CreateAccessTokenInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create access token internal server error response a status code equal to that given
func (o *CreateAccessTokenInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create access token internal server error response
func (o *CreateAccessTokenInternalServerError) Code() int {
	return 500
}

func (o *CreateAccessTokenInternalServerError) Error() string {
	return fmt.Sprintf("[POST /access/token][%d] createAccessTokenInternalServerError ", 500)
}

func (o *CreateAccessTokenInternalServerError) String() string {
	return fmt.Sprintf("[POST /access/token][%d] createAccessTokenInternalServerError ", 500)
}

func (o *CreateAccessTokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
