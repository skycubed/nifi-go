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

// DeleteRegistryClientReader is a Reader for the DeleteRegistryClient structure.
type DeleteRegistryClientReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRegistryClientReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRegistryClientOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRegistryClientBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteRegistryClientUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteRegistryClientForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRegistryClientNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteRegistryClientConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /controller/registry-clients/{id}] deleteRegistryClient", response, response.Code())
	}
}

// NewDeleteRegistryClientOK creates a DeleteRegistryClientOK with default headers values
func NewDeleteRegistryClientOK() *DeleteRegistryClientOK {
	return &DeleteRegistryClientOK{}
}

/*
DeleteRegistryClientOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteRegistryClientOK struct {
	Payload *models.RegistryClientEntity
}

// IsSuccess returns true when this delete registry client o k response has a 2xx status code
func (o *DeleteRegistryClientOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete registry client o k response has a 3xx status code
func (o *DeleteRegistryClientOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete registry client o k response has a 4xx status code
func (o *DeleteRegistryClientOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete registry client o k response has a 5xx status code
func (o *DeleteRegistryClientOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete registry client o k response a status code equal to that given
func (o *DeleteRegistryClientOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete registry client o k response
func (o *DeleteRegistryClientOK) Code() int {
	return 200
}

func (o *DeleteRegistryClientOK) Error() string {
	return fmt.Sprintf("[DELETE /controller/registry-clients/{id}][%d] deleteRegistryClientOK  %+v", 200, o.Payload)
}

func (o *DeleteRegistryClientOK) String() string {
	return fmt.Sprintf("[DELETE /controller/registry-clients/{id}][%d] deleteRegistryClientOK  %+v", 200, o.Payload)
}

func (o *DeleteRegistryClientOK) GetPayload() *models.RegistryClientEntity {
	return o.Payload
}

func (o *DeleteRegistryClientOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegistryClientEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRegistryClientBadRequest creates a DeleteRegistryClientBadRequest with default headers values
func NewDeleteRegistryClientBadRequest() *DeleteRegistryClientBadRequest {
	return &DeleteRegistryClientBadRequest{}
}

/*
DeleteRegistryClientBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type DeleteRegistryClientBadRequest struct {
}

// IsSuccess returns true when this delete registry client bad request response has a 2xx status code
func (o *DeleteRegistryClientBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete registry client bad request response has a 3xx status code
func (o *DeleteRegistryClientBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete registry client bad request response has a 4xx status code
func (o *DeleteRegistryClientBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete registry client bad request response has a 5xx status code
func (o *DeleteRegistryClientBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete registry client bad request response a status code equal to that given
func (o *DeleteRegistryClientBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete registry client bad request response
func (o *DeleteRegistryClientBadRequest) Code() int {
	return 400
}

func (o *DeleteRegistryClientBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /controller/registry-clients/{id}][%d] deleteRegistryClientBadRequest ", 400)
}

func (o *DeleteRegistryClientBadRequest) String() string {
	return fmt.Sprintf("[DELETE /controller/registry-clients/{id}][%d] deleteRegistryClientBadRequest ", 400)
}

func (o *DeleteRegistryClientBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRegistryClientUnauthorized creates a DeleteRegistryClientUnauthorized with default headers values
func NewDeleteRegistryClientUnauthorized() *DeleteRegistryClientUnauthorized {
	return &DeleteRegistryClientUnauthorized{}
}

/*
DeleteRegistryClientUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type DeleteRegistryClientUnauthorized struct {
}

// IsSuccess returns true when this delete registry client unauthorized response has a 2xx status code
func (o *DeleteRegistryClientUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete registry client unauthorized response has a 3xx status code
func (o *DeleteRegistryClientUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete registry client unauthorized response has a 4xx status code
func (o *DeleteRegistryClientUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete registry client unauthorized response has a 5xx status code
func (o *DeleteRegistryClientUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete registry client unauthorized response a status code equal to that given
func (o *DeleteRegistryClientUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete registry client unauthorized response
func (o *DeleteRegistryClientUnauthorized) Code() int {
	return 401
}

func (o *DeleteRegistryClientUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /controller/registry-clients/{id}][%d] deleteRegistryClientUnauthorized ", 401)
}

func (o *DeleteRegistryClientUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /controller/registry-clients/{id}][%d] deleteRegistryClientUnauthorized ", 401)
}

func (o *DeleteRegistryClientUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRegistryClientForbidden creates a DeleteRegistryClientForbidden with default headers values
func NewDeleteRegistryClientForbidden() *DeleteRegistryClientForbidden {
	return &DeleteRegistryClientForbidden{}
}

/*
DeleteRegistryClientForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type DeleteRegistryClientForbidden struct {
}

// IsSuccess returns true when this delete registry client forbidden response has a 2xx status code
func (o *DeleteRegistryClientForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete registry client forbidden response has a 3xx status code
func (o *DeleteRegistryClientForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete registry client forbidden response has a 4xx status code
func (o *DeleteRegistryClientForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete registry client forbidden response has a 5xx status code
func (o *DeleteRegistryClientForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete registry client forbidden response a status code equal to that given
func (o *DeleteRegistryClientForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete registry client forbidden response
func (o *DeleteRegistryClientForbidden) Code() int {
	return 403
}

func (o *DeleteRegistryClientForbidden) Error() string {
	return fmt.Sprintf("[DELETE /controller/registry-clients/{id}][%d] deleteRegistryClientForbidden ", 403)
}

func (o *DeleteRegistryClientForbidden) String() string {
	return fmt.Sprintf("[DELETE /controller/registry-clients/{id}][%d] deleteRegistryClientForbidden ", 403)
}

func (o *DeleteRegistryClientForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRegistryClientNotFound creates a DeleteRegistryClientNotFound with default headers values
func NewDeleteRegistryClientNotFound() *DeleteRegistryClientNotFound {
	return &DeleteRegistryClientNotFound{}
}

/*
DeleteRegistryClientNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type DeleteRegistryClientNotFound struct {
}

// IsSuccess returns true when this delete registry client not found response has a 2xx status code
func (o *DeleteRegistryClientNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete registry client not found response has a 3xx status code
func (o *DeleteRegistryClientNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete registry client not found response has a 4xx status code
func (o *DeleteRegistryClientNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete registry client not found response has a 5xx status code
func (o *DeleteRegistryClientNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete registry client not found response a status code equal to that given
func (o *DeleteRegistryClientNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete registry client not found response
func (o *DeleteRegistryClientNotFound) Code() int {
	return 404
}

func (o *DeleteRegistryClientNotFound) Error() string {
	return fmt.Sprintf("[DELETE /controller/registry-clients/{id}][%d] deleteRegistryClientNotFound ", 404)
}

func (o *DeleteRegistryClientNotFound) String() string {
	return fmt.Sprintf("[DELETE /controller/registry-clients/{id}][%d] deleteRegistryClientNotFound ", 404)
}

func (o *DeleteRegistryClientNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRegistryClientConflict creates a DeleteRegistryClientConflict with default headers values
func NewDeleteRegistryClientConflict() *DeleteRegistryClientConflict {
	return &DeleteRegistryClientConflict{}
}

/*
DeleteRegistryClientConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type DeleteRegistryClientConflict struct {
}

// IsSuccess returns true when this delete registry client conflict response has a 2xx status code
func (o *DeleteRegistryClientConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete registry client conflict response has a 3xx status code
func (o *DeleteRegistryClientConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete registry client conflict response has a 4xx status code
func (o *DeleteRegistryClientConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete registry client conflict response has a 5xx status code
func (o *DeleteRegistryClientConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this delete registry client conflict response a status code equal to that given
func (o *DeleteRegistryClientConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the delete registry client conflict response
func (o *DeleteRegistryClientConflict) Code() int {
	return 409
}

func (o *DeleteRegistryClientConflict) Error() string {
	return fmt.Sprintf("[DELETE /controller/registry-clients/{id}][%d] deleteRegistryClientConflict ", 409)
}

func (o *DeleteRegistryClientConflict) String() string {
	return fmt.Sprintf("[DELETE /controller/registry-clients/{id}][%d] deleteRegistryClientConflict ", 409)
}

func (o *DeleteRegistryClientConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
