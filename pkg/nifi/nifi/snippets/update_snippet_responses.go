// Code generated by go-swagger; DO NOT EDIT.

package snippets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// UpdateSnippetReader is a Reader for the UpdateSnippet structure.
type UpdateSnippetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSnippetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSnippetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateSnippetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateSnippetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateSnippetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateSnippetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateSnippetConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /snippets/{id}] updateSnippet", response, response.Code())
	}
}

// NewUpdateSnippetOK creates a UpdateSnippetOK with default headers values
func NewUpdateSnippetOK() *UpdateSnippetOK {
	return &UpdateSnippetOK{}
}

/*
UpdateSnippetOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateSnippetOK struct {
	Payload *models.SnippetEntity
}

// IsSuccess returns true when this update snippet o k response has a 2xx status code
func (o *UpdateSnippetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update snippet o k response has a 3xx status code
func (o *UpdateSnippetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update snippet o k response has a 4xx status code
func (o *UpdateSnippetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update snippet o k response has a 5xx status code
func (o *UpdateSnippetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update snippet o k response a status code equal to that given
func (o *UpdateSnippetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update snippet o k response
func (o *UpdateSnippetOK) Code() int {
	return 200
}

func (o *UpdateSnippetOK) Error() string {
	return fmt.Sprintf("[PUT /snippets/{id}][%d] updateSnippetOK  %+v", 200, o.Payload)
}

func (o *UpdateSnippetOK) String() string {
	return fmt.Sprintf("[PUT /snippets/{id}][%d] updateSnippetOK  %+v", 200, o.Payload)
}

func (o *UpdateSnippetOK) GetPayload() *models.SnippetEntity {
	return o.Payload
}

func (o *UpdateSnippetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnippetEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSnippetBadRequest creates a UpdateSnippetBadRequest with default headers values
func NewUpdateSnippetBadRequest() *UpdateSnippetBadRequest {
	return &UpdateSnippetBadRequest{}
}

/*
UpdateSnippetBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type UpdateSnippetBadRequest struct {
}

// IsSuccess returns true when this update snippet bad request response has a 2xx status code
func (o *UpdateSnippetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update snippet bad request response has a 3xx status code
func (o *UpdateSnippetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update snippet bad request response has a 4xx status code
func (o *UpdateSnippetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update snippet bad request response has a 5xx status code
func (o *UpdateSnippetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update snippet bad request response a status code equal to that given
func (o *UpdateSnippetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update snippet bad request response
func (o *UpdateSnippetBadRequest) Code() int {
	return 400
}

func (o *UpdateSnippetBadRequest) Error() string {
	return fmt.Sprintf("[PUT /snippets/{id}][%d] updateSnippetBadRequest ", 400)
}

func (o *UpdateSnippetBadRequest) String() string {
	return fmt.Sprintf("[PUT /snippets/{id}][%d] updateSnippetBadRequest ", 400)
}

func (o *UpdateSnippetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSnippetUnauthorized creates a UpdateSnippetUnauthorized with default headers values
func NewUpdateSnippetUnauthorized() *UpdateSnippetUnauthorized {
	return &UpdateSnippetUnauthorized{}
}

/*
UpdateSnippetUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type UpdateSnippetUnauthorized struct {
}

// IsSuccess returns true when this update snippet unauthorized response has a 2xx status code
func (o *UpdateSnippetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update snippet unauthorized response has a 3xx status code
func (o *UpdateSnippetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update snippet unauthorized response has a 4xx status code
func (o *UpdateSnippetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update snippet unauthorized response has a 5xx status code
func (o *UpdateSnippetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update snippet unauthorized response a status code equal to that given
func (o *UpdateSnippetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update snippet unauthorized response
func (o *UpdateSnippetUnauthorized) Code() int {
	return 401
}

func (o *UpdateSnippetUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /snippets/{id}][%d] updateSnippetUnauthorized ", 401)
}

func (o *UpdateSnippetUnauthorized) String() string {
	return fmt.Sprintf("[PUT /snippets/{id}][%d] updateSnippetUnauthorized ", 401)
}

func (o *UpdateSnippetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSnippetForbidden creates a UpdateSnippetForbidden with default headers values
func NewUpdateSnippetForbidden() *UpdateSnippetForbidden {
	return &UpdateSnippetForbidden{}
}

/*
UpdateSnippetForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type UpdateSnippetForbidden struct {
}

// IsSuccess returns true when this update snippet forbidden response has a 2xx status code
func (o *UpdateSnippetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update snippet forbidden response has a 3xx status code
func (o *UpdateSnippetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update snippet forbidden response has a 4xx status code
func (o *UpdateSnippetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update snippet forbidden response has a 5xx status code
func (o *UpdateSnippetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update snippet forbidden response a status code equal to that given
func (o *UpdateSnippetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update snippet forbidden response
func (o *UpdateSnippetForbidden) Code() int {
	return 403
}

func (o *UpdateSnippetForbidden) Error() string {
	return fmt.Sprintf("[PUT /snippets/{id}][%d] updateSnippetForbidden ", 403)
}

func (o *UpdateSnippetForbidden) String() string {
	return fmt.Sprintf("[PUT /snippets/{id}][%d] updateSnippetForbidden ", 403)
}

func (o *UpdateSnippetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSnippetNotFound creates a UpdateSnippetNotFound with default headers values
func NewUpdateSnippetNotFound() *UpdateSnippetNotFound {
	return &UpdateSnippetNotFound{}
}

/*
UpdateSnippetNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type UpdateSnippetNotFound struct {
}

// IsSuccess returns true when this update snippet not found response has a 2xx status code
func (o *UpdateSnippetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update snippet not found response has a 3xx status code
func (o *UpdateSnippetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update snippet not found response has a 4xx status code
func (o *UpdateSnippetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update snippet not found response has a 5xx status code
func (o *UpdateSnippetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update snippet not found response a status code equal to that given
func (o *UpdateSnippetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update snippet not found response
func (o *UpdateSnippetNotFound) Code() int {
	return 404
}

func (o *UpdateSnippetNotFound) Error() string {
	return fmt.Sprintf("[PUT /snippets/{id}][%d] updateSnippetNotFound ", 404)
}

func (o *UpdateSnippetNotFound) String() string {
	return fmt.Sprintf("[PUT /snippets/{id}][%d] updateSnippetNotFound ", 404)
}

func (o *UpdateSnippetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSnippetConflict creates a UpdateSnippetConflict with default headers values
func NewUpdateSnippetConflict() *UpdateSnippetConflict {
	return &UpdateSnippetConflict{}
}

/*
UpdateSnippetConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type UpdateSnippetConflict struct {
}

// IsSuccess returns true when this update snippet conflict response has a 2xx status code
func (o *UpdateSnippetConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update snippet conflict response has a 3xx status code
func (o *UpdateSnippetConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update snippet conflict response has a 4xx status code
func (o *UpdateSnippetConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update snippet conflict response has a 5xx status code
func (o *UpdateSnippetConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update snippet conflict response a status code equal to that given
func (o *UpdateSnippetConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the update snippet conflict response
func (o *UpdateSnippetConflict) Code() int {
	return 409
}

func (o *UpdateSnippetConflict) Error() string {
	return fmt.Sprintf("[PUT /snippets/{id}][%d] updateSnippetConflict ", 409)
}

func (o *UpdateSnippetConflict) String() string {
	return fmt.Sprintf("[PUT /snippets/{id}][%d] updateSnippetConflict ", 409)
}

func (o *UpdateSnippetConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
