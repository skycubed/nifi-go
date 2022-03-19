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
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateAccessTokenCreated creates a CreateAccessTokenCreated with default headers values
func NewCreateAccessTokenCreated() *CreateAccessTokenCreated {
	return &CreateAccessTokenCreated{}
}

/* CreateAccessTokenCreated describes a response with status code 201, with default header values.

successful operation
*/
type CreateAccessTokenCreated struct {
	Payload string
}

func (o *CreateAccessTokenCreated) Error() string {
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

/* CreateAccessTokenBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type CreateAccessTokenBadRequest struct {
}

func (o *CreateAccessTokenBadRequest) Error() string {
	return fmt.Sprintf("[POST /access/token][%d] createAccessTokenBadRequest ", 400)
}

func (o *CreateAccessTokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAccessTokenForbidden creates a CreateAccessTokenForbidden with default headers values
func NewCreateAccessTokenForbidden() *CreateAccessTokenForbidden {
	return &CreateAccessTokenForbidden{}
}

/* CreateAccessTokenForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type CreateAccessTokenForbidden struct {
}

func (o *CreateAccessTokenForbidden) Error() string {
	return fmt.Sprintf("[POST /access/token][%d] createAccessTokenForbidden ", 403)
}

func (o *CreateAccessTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAccessTokenConflict creates a CreateAccessTokenConflict with default headers values
func NewCreateAccessTokenConflict() *CreateAccessTokenConflict {
	return &CreateAccessTokenConflict{}
}

/* CreateAccessTokenConflict describes a response with status code 409, with default header values.

Unable to create access token because NiFi is not in the appropriate state. (i.e. may not be configured to support username/password login.
*/
type CreateAccessTokenConflict struct {
}

func (o *CreateAccessTokenConflict) Error() string {
	return fmt.Sprintf("[POST /access/token][%d] createAccessTokenConflict ", 409)
}

func (o *CreateAccessTokenConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAccessTokenInternalServerError creates a CreateAccessTokenInternalServerError with default headers values
func NewCreateAccessTokenInternalServerError() *CreateAccessTokenInternalServerError {
	return &CreateAccessTokenInternalServerError{}
}

/* CreateAccessTokenInternalServerError describes a response with status code 500, with default header values.

Unable to create access token because an unexpected error occurred.
*/
type CreateAccessTokenInternalServerError struct {
}

func (o *CreateAccessTokenInternalServerError) Error() string {
	return fmt.Sprintf("[POST /access/token][%d] createAccessTokenInternalServerError ", 500)
}

func (o *CreateAccessTokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
