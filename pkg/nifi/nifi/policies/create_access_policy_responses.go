// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// CreateAccessPolicyReader is a Reader for the CreateAccessPolicy structure.
type CreateAccessPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAccessPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateAccessPolicyCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateAccessPolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateAccessPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateAccessPolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateAccessPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateAccessPolicyConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /policies] createAccessPolicy", response, response.Code())
	}
}

// NewCreateAccessPolicyCreated creates a CreateAccessPolicyCreated with default headers values
func NewCreateAccessPolicyCreated() *CreateAccessPolicyCreated {
	return &CreateAccessPolicyCreated{}
}

/*
CreateAccessPolicyCreated describes a response with status code 201, with default header values.

successful operation
*/
type CreateAccessPolicyCreated struct {
	Payload *models.AccessPolicyEntity
}

// IsSuccess returns true when this create access policy created response has a 2xx status code
func (o *CreateAccessPolicyCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create access policy created response has a 3xx status code
func (o *CreateAccessPolicyCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create access policy created response has a 4xx status code
func (o *CreateAccessPolicyCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create access policy created response has a 5xx status code
func (o *CreateAccessPolicyCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create access policy created response a status code equal to that given
func (o *CreateAccessPolicyCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create access policy created response
func (o *CreateAccessPolicyCreated) Code() int {
	return 201
}

func (o *CreateAccessPolicyCreated) Error() string {
	return fmt.Sprintf("[POST /policies][%d] createAccessPolicyCreated  %+v", 201, o.Payload)
}

func (o *CreateAccessPolicyCreated) String() string {
	return fmt.Sprintf("[POST /policies][%d] createAccessPolicyCreated  %+v", 201, o.Payload)
}

func (o *CreateAccessPolicyCreated) GetPayload() *models.AccessPolicyEntity {
	return o.Payload
}

func (o *CreateAccessPolicyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccessPolicyEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAccessPolicyBadRequest creates a CreateAccessPolicyBadRequest with default headers values
func NewCreateAccessPolicyBadRequest() *CreateAccessPolicyBadRequest {
	return &CreateAccessPolicyBadRequest{}
}

/*
CreateAccessPolicyBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type CreateAccessPolicyBadRequest struct {
}

// IsSuccess returns true when this create access policy bad request response has a 2xx status code
func (o *CreateAccessPolicyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create access policy bad request response has a 3xx status code
func (o *CreateAccessPolicyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create access policy bad request response has a 4xx status code
func (o *CreateAccessPolicyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create access policy bad request response has a 5xx status code
func (o *CreateAccessPolicyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create access policy bad request response a status code equal to that given
func (o *CreateAccessPolicyBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create access policy bad request response
func (o *CreateAccessPolicyBadRequest) Code() int {
	return 400
}

func (o *CreateAccessPolicyBadRequest) Error() string {
	return fmt.Sprintf("[POST /policies][%d] createAccessPolicyBadRequest ", 400)
}

func (o *CreateAccessPolicyBadRequest) String() string {
	return fmt.Sprintf("[POST /policies][%d] createAccessPolicyBadRequest ", 400)
}

func (o *CreateAccessPolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAccessPolicyUnauthorized creates a CreateAccessPolicyUnauthorized with default headers values
func NewCreateAccessPolicyUnauthorized() *CreateAccessPolicyUnauthorized {
	return &CreateAccessPolicyUnauthorized{}
}

/*
CreateAccessPolicyUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type CreateAccessPolicyUnauthorized struct {
}

// IsSuccess returns true when this create access policy unauthorized response has a 2xx status code
func (o *CreateAccessPolicyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create access policy unauthorized response has a 3xx status code
func (o *CreateAccessPolicyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create access policy unauthorized response has a 4xx status code
func (o *CreateAccessPolicyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create access policy unauthorized response has a 5xx status code
func (o *CreateAccessPolicyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create access policy unauthorized response a status code equal to that given
func (o *CreateAccessPolicyUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create access policy unauthorized response
func (o *CreateAccessPolicyUnauthorized) Code() int {
	return 401
}

func (o *CreateAccessPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[POST /policies][%d] createAccessPolicyUnauthorized ", 401)
}

func (o *CreateAccessPolicyUnauthorized) String() string {
	return fmt.Sprintf("[POST /policies][%d] createAccessPolicyUnauthorized ", 401)
}

func (o *CreateAccessPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAccessPolicyForbidden creates a CreateAccessPolicyForbidden with default headers values
func NewCreateAccessPolicyForbidden() *CreateAccessPolicyForbidden {
	return &CreateAccessPolicyForbidden{}
}

/*
CreateAccessPolicyForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type CreateAccessPolicyForbidden struct {
}

// IsSuccess returns true when this create access policy forbidden response has a 2xx status code
func (o *CreateAccessPolicyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create access policy forbidden response has a 3xx status code
func (o *CreateAccessPolicyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create access policy forbidden response has a 4xx status code
func (o *CreateAccessPolicyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create access policy forbidden response has a 5xx status code
func (o *CreateAccessPolicyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create access policy forbidden response a status code equal to that given
func (o *CreateAccessPolicyForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create access policy forbidden response
func (o *CreateAccessPolicyForbidden) Code() int {
	return 403
}

func (o *CreateAccessPolicyForbidden) Error() string {
	return fmt.Sprintf("[POST /policies][%d] createAccessPolicyForbidden ", 403)
}

func (o *CreateAccessPolicyForbidden) String() string {
	return fmt.Sprintf("[POST /policies][%d] createAccessPolicyForbidden ", 403)
}

func (o *CreateAccessPolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAccessPolicyNotFound creates a CreateAccessPolicyNotFound with default headers values
func NewCreateAccessPolicyNotFound() *CreateAccessPolicyNotFound {
	return &CreateAccessPolicyNotFound{}
}

/*
CreateAccessPolicyNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type CreateAccessPolicyNotFound struct {
}

// IsSuccess returns true when this create access policy not found response has a 2xx status code
func (o *CreateAccessPolicyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create access policy not found response has a 3xx status code
func (o *CreateAccessPolicyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create access policy not found response has a 4xx status code
func (o *CreateAccessPolicyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create access policy not found response has a 5xx status code
func (o *CreateAccessPolicyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create access policy not found response a status code equal to that given
func (o *CreateAccessPolicyNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create access policy not found response
func (o *CreateAccessPolicyNotFound) Code() int {
	return 404
}

func (o *CreateAccessPolicyNotFound) Error() string {
	return fmt.Sprintf("[POST /policies][%d] createAccessPolicyNotFound ", 404)
}

func (o *CreateAccessPolicyNotFound) String() string {
	return fmt.Sprintf("[POST /policies][%d] createAccessPolicyNotFound ", 404)
}

func (o *CreateAccessPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAccessPolicyConflict creates a CreateAccessPolicyConflict with default headers values
func NewCreateAccessPolicyConflict() *CreateAccessPolicyConflict {
	return &CreateAccessPolicyConflict{}
}

/*
CreateAccessPolicyConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type CreateAccessPolicyConflict struct {
}

// IsSuccess returns true when this create access policy conflict response has a 2xx status code
func (o *CreateAccessPolicyConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create access policy conflict response has a 3xx status code
func (o *CreateAccessPolicyConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create access policy conflict response has a 4xx status code
func (o *CreateAccessPolicyConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create access policy conflict response has a 5xx status code
func (o *CreateAccessPolicyConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create access policy conflict response a status code equal to that given
func (o *CreateAccessPolicyConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create access policy conflict response
func (o *CreateAccessPolicyConflict) Code() int {
	return 409
}

func (o *CreateAccessPolicyConflict) Error() string {
	return fmt.Sprintf("[POST /policies][%d] createAccessPolicyConflict ", 409)
}

func (o *CreateAccessPolicyConflict) String() string {
	return fmt.Sprintf("[POST /policies][%d] createAccessPolicyConflict ", 409)
}

func (o *CreateAccessPolicyConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
