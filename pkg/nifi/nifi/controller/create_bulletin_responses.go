// Code generated by go-swagger; DO NOT EDIT.

package controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// CreateBulletinReader is a Reader for the CreateBulletin structure.
type CreateBulletinReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateBulletinReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateBulletinOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateBulletinBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateBulletinUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateBulletinForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateBulletinConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /controller/bulletin] createBulletin", response, response.Code())
	}
}

// NewCreateBulletinOK creates a CreateBulletinOK with default headers values
func NewCreateBulletinOK() *CreateBulletinOK {
	return &CreateBulletinOK{}
}

/*
CreateBulletinOK describes a response with status code 200, with default header values.

successful operation
*/
type CreateBulletinOK struct {
	Payload *models.BulletinEntity
}

// IsSuccess returns true when this create bulletin o k response has a 2xx status code
func (o *CreateBulletinOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create bulletin o k response has a 3xx status code
func (o *CreateBulletinOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create bulletin o k response has a 4xx status code
func (o *CreateBulletinOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create bulletin o k response has a 5xx status code
func (o *CreateBulletinOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create bulletin o k response a status code equal to that given
func (o *CreateBulletinOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create bulletin o k response
func (o *CreateBulletinOK) Code() int {
	return 200
}

func (o *CreateBulletinOK) Error() string {
	return fmt.Sprintf("[POST /controller/bulletin][%d] createBulletinOK  %+v", 200, o.Payload)
}

func (o *CreateBulletinOK) String() string {
	return fmt.Sprintf("[POST /controller/bulletin][%d] createBulletinOK  %+v", 200, o.Payload)
}

func (o *CreateBulletinOK) GetPayload() *models.BulletinEntity {
	return o.Payload
}

func (o *CreateBulletinOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BulletinEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBulletinBadRequest creates a CreateBulletinBadRequest with default headers values
func NewCreateBulletinBadRequest() *CreateBulletinBadRequest {
	return &CreateBulletinBadRequest{}
}

/*
CreateBulletinBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type CreateBulletinBadRequest struct {
}

// IsSuccess returns true when this create bulletin bad request response has a 2xx status code
func (o *CreateBulletinBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create bulletin bad request response has a 3xx status code
func (o *CreateBulletinBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create bulletin bad request response has a 4xx status code
func (o *CreateBulletinBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create bulletin bad request response has a 5xx status code
func (o *CreateBulletinBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create bulletin bad request response a status code equal to that given
func (o *CreateBulletinBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create bulletin bad request response
func (o *CreateBulletinBadRequest) Code() int {
	return 400
}

func (o *CreateBulletinBadRequest) Error() string {
	return fmt.Sprintf("[POST /controller/bulletin][%d] createBulletinBadRequest ", 400)
}

func (o *CreateBulletinBadRequest) String() string {
	return fmt.Sprintf("[POST /controller/bulletin][%d] createBulletinBadRequest ", 400)
}

func (o *CreateBulletinBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateBulletinUnauthorized creates a CreateBulletinUnauthorized with default headers values
func NewCreateBulletinUnauthorized() *CreateBulletinUnauthorized {
	return &CreateBulletinUnauthorized{}
}

/*
CreateBulletinUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type CreateBulletinUnauthorized struct {
}

// IsSuccess returns true when this create bulletin unauthorized response has a 2xx status code
func (o *CreateBulletinUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create bulletin unauthorized response has a 3xx status code
func (o *CreateBulletinUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create bulletin unauthorized response has a 4xx status code
func (o *CreateBulletinUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create bulletin unauthorized response has a 5xx status code
func (o *CreateBulletinUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create bulletin unauthorized response a status code equal to that given
func (o *CreateBulletinUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create bulletin unauthorized response
func (o *CreateBulletinUnauthorized) Code() int {
	return 401
}

func (o *CreateBulletinUnauthorized) Error() string {
	return fmt.Sprintf("[POST /controller/bulletin][%d] createBulletinUnauthorized ", 401)
}

func (o *CreateBulletinUnauthorized) String() string {
	return fmt.Sprintf("[POST /controller/bulletin][%d] createBulletinUnauthorized ", 401)
}

func (o *CreateBulletinUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateBulletinForbidden creates a CreateBulletinForbidden with default headers values
func NewCreateBulletinForbidden() *CreateBulletinForbidden {
	return &CreateBulletinForbidden{}
}

/*
CreateBulletinForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type CreateBulletinForbidden struct {
}

// IsSuccess returns true when this create bulletin forbidden response has a 2xx status code
func (o *CreateBulletinForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create bulletin forbidden response has a 3xx status code
func (o *CreateBulletinForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create bulletin forbidden response has a 4xx status code
func (o *CreateBulletinForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create bulletin forbidden response has a 5xx status code
func (o *CreateBulletinForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create bulletin forbidden response a status code equal to that given
func (o *CreateBulletinForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create bulletin forbidden response
func (o *CreateBulletinForbidden) Code() int {
	return 403
}

func (o *CreateBulletinForbidden) Error() string {
	return fmt.Sprintf("[POST /controller/bulletin][%d] createBulletinForbidden ", 403)
}

func (o *CreateBulletinForbidden) String() string {
	return fmt.Sprintf("[POST /controller/bulletin][%d] createBulletinForbidden ", 403)
}

func (o *CreateBulletinForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateBulletinConflict creates a CreateBulletinConflict with default headers values
func NewCreateBulletinConflict() *CreateBulletinConflict {
	return &CreateBulletinConflict{}
}

/*
CreateBulletinConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type CreateBulletinConflict struct {
}

// IsSuccess returns true when this create bulletin conflict response has a 2xx status code
func (o *CreateBulletinConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create bulletin conflict response has a 3xx status code
func (o *CreateBulletinConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create bulletin conflict response has a 4xx status code
func (o *CreateBulletinConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create bulletin conflict response has a 5xx status code
func (o *CreateBulletinConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create bulletin conflict response a status code equal to that given
func (o *CreateBulletinConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create bulletin conflict response
func (o *CreateBulletinConflict) Code() int {
	return 409
}

func (o *CreateBulletinConflict) Error() string {
	return fmt.Sprintf("[POST /controller/bulletin][%d] createBulletinConflict ", 409)
}

func (o *CreateBulletinConflict) String() string {
	return fmt.Sprintf("[POST /controller/bulletin][%d] createBulletinConflict ", 409)
}

func (o *CreateBulletinConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
