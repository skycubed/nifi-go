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

// DeleteFlowRegistryClientReader is a Reader for the DeleteFlowRegistryClient structure.
type DeleteFlowRegistryClientReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFlowRegistryClientReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteFlowRegistryClientOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteFlowRegistryClientBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteFlowRegistryClientUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteFlowRegistryClientForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteFlowRegistryClientNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteFlowRegistryClientConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /controller/registry-clients/{id}] deleteFlowRegistryClient", response, response.Code())
	}
}

// NewDeleteFlowRegistryClientOK creates a DeleteFlowRegistryClientOK with default headers values
func NewDeleteFlowRegistryClientOK() *DeleteFlowRegistryClientOK {
	return &DeleteFlowRegistryClientOK{}
}

/*
DeleteFlowRegistryClientOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteFlowRegistryClientOK struct {
	Payload *models.FlowRegistryClientEntity
}

// IsSuccess returns true when this delete flow registry client o k response has a 2xx status code
func (o *DeleteFlowRegistryClientOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete flow registry client o k response has a 3xx status code
func (o *DeleteFlowRegistryClientOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete flow registry client o k response has a 4xx status code
func (o *DeleteFlowRegistryClientOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete flow registry client o k response has a 5xx status code
func (o *DeleteFlowRegistryClientOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete flow registry client o k response a status code equal to that given
func (o *DeleteFlowRegistryClientOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete flow registry client o k response
func (o *DeleteFlowRegistryClientOK) Code() int {
	return 200
}

func (o *DeleteFlowRegistryClientOK) Error() string {
	return fmt.Sprintf("[DELETE /controller/registry-clients/{id}][%d] deleteFlowRegistryClientOK  %+v", 200, o.Payload)
}

func (o *DeleteFlowRegistryClientOK) String() string {
	return fmt.Sprintf("[DELETE /controller/registry-clients/{id}][%d] deleteFlowRegistryClientOK  %+v", 200, o.Payload)
}

func (o *DeleteFlowRegistryClientOK) GetPayload() *models.FlowRegistryClientEntity {
	return o.Payload
}

func (o *DeleteFlowRegistryClientOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FlowRegistryClientEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowRegistryClientBadRequest creates a DeleteFlowRegistryClientBadRequest with default headers values
func NewDeleteFlowRegistryClientBadRequest() *DeleteFlowRegistryClientBadRequest {
	return &DeleteFlowRegistryClientBadRequest{}
}

/*
DeleteFlowRegistryClientBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type DeleteFlowRegistryClientBadRequest struct {
}

// IsSuccess returns true when this delete flow registry client bad request response has a 2xx status code
func (o *DeleteFlowRegistryClientBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete flow registry client bad request response has a 3xx status code
func (o *DeleteFlowRegistryClientBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete flow registry client bad request response has a 4xx status code
func (o *DeleteFlowRegistryClientBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete flow registry client bad request response has a 5xx status code
func (o *DeleteFlowRegistryClientBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete flow registry client bad request response a status code equal to that given
func (o *DeleteFlowRegistryClientBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete flow registry client bad request response
func (o *DeleteFlowRegistryClientBadRequest) Code() int {
	return 400
}

func (o *DeleteFlowRegistryClientBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /controller/registry-clients/{id}][%d] deleteFlowRegistryClientBadRequest ", 400)
}

func (o *DeleteFlowRegistryClientBadRequest) String() string {
	return fmt.Sprintf("[DELETE /controller/registry-clients/{id}][%d] deleteFlowRegistryClientBadRequest ", 400)
}

func (o *DeleteFlowRegistryClientBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFlowRegistryClientUnauthorized creates a DeleteFlowRegistryClientUnauthorized with default headers values
func NewDeleteFlowRegistryClientUnauthorized() *DeleteFlowRegistryClientUnauthorized {
	return &DeleteFlowRegistryClientUnauthorized{}
}

/*
DeleteFlowRegistryClientUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type DeleteFlowRegistryClientUnauthorized struct {
}

// IsSuccess returns true when this delete flow registry client unauthorized response has a 2xx status code
func (o *DeleteFlowRegistryClientUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete flow registry client unauthorized response has a 3xx status code
func (o *DeleteFlowRegistryClientUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete flow registry client unauthorized response has a 4xx status code
func (o *DeleteFlowRegistryClientUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete flow registry client unauthorized response has a 5xx status code
func (o *DeleteFlowRegistryClientUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete flow registry client unauthorized response a status code equal to that given
func (o *DeleteFlowRegistryClientUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete flow registry client unauthorized response
func (o *DeleteFlowRegistryClientUnauthorized) Code() int {
	return 401
}

func (o *DeleteFlowRegistryClientUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /controller/registry-clients/{id}][%d] deleteFlowRegistryClientUnauthorized ", 401)
}

func (o *DeleteFlowRegistryClientUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /controller/registry-clients/{id}][%d] deleteFlowRegistryClientUnauthorized ", 401)
}

func (o *DeleteFlowRegistryClientUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFlowRegistryClientForbidden creates a DeleteFlowRegistryClientForbidden with default headers values
func NewDeleteFlowRegistryClientForbidden() *DeleteFlowRegistryClientForbidden {
	return &DeleteFlowRegistryClientForbidden{}
}

/*
DeleteFlowRegistryClientForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type DeleteFlowRegistryClientForbidden struct {
}

// IsSuccess returns true when this delete flow registry client forbidden response has a 2xx status code
func (o *DeleteFlowRegistryClientForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete flow registry client forbidden response has a 3xx status code
func (o *DeleteFlowRegistryClientForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete flow registry client forbidden response has a 4xx status code
func (o *DeleteFlowRegistryClientForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete flow registry client forbidden response has a 5xx status code
func (o *DeleteFlowRegistryClientForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete flow registry client forbidden response a status code equal to that given
func (o *DeleteFlowRegistryClientForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete flow registry client forbidden response
func (o *DeleteFlowRegistryClientForbidden) Code() int {
	return 403
}

func (o *DeleteFlowRegistryClientForbidden) Error() string {
	return fmt.Sprintf("[DELETE /controller/registry-clients/{id}][%d] deleteFlowRegistryClientForbidden ", 403)
}

func (o *DeleteFlowRegistryClientForbidden) String() string {
	return fmt.Sprintf("[DELETE /controller/registry-clients/{id}][%d] deleteFlowRegistryClientForbidden ", 403)
}

func (o *DeleteFlowRegistryClientForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFlowRegistryClientNotFound creates a DeleteFlowRegistryClientNotFound with default headers values
func NewDeleteFlowRegistryClientNotFound() *DeleteFlowRegistryClientNotFound {
	return &DeleteFlowRegistryClientNotFound{}
}

/*
DeleteFlowRegistryClientNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type DeleteFlowRegistryClientNotFound struct {
}

// IsSuccess returns true when this delete flow registry client not found response has a 2xx status code
func (o *DeleteFlowRegistryClientNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete flow registry client not found response has a 3xx status code
func (o *DeleteFlowRegistryClientNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete flow registry client not found response has a 4xx status code
func (o *DeleteFlowRegistryClientNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete flow registry client not found response has a 5xx status code
func (o *DeleteFlowRegistryClientNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete flow registry client not found response a status code equal to that given
func (o *DeleteFlowRegistryClientNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete flow registry client not found response
func (o *DeleteFlowRegistryClientNotFound) Code() int {
	return 404
}

func (o *DeleteFlowRegistryClientNotFound) Error() string {
	return fmt.Sprintf("[DELETE /controller/registry-clients/{id}][%d] deleteFlowRegistryClientNotFound ", 404)
}

func (o *DeleteFlowRegistryClientNotFound) String() string {
	return fmt.Sprintf("[DELETE /controller/registry-clients/{id}][%d] deleteFlowRegistryClientNotFound ", 404)
}

func (o *DeleteFlowRegistryClientNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFlowRegistryClientConflict creates a DeleteFlowRegistryClientConflict with default headers values
func NewDeleteFlowRegistryClientConflict() *DeleteFlowRegistryClientConflict {
	return &DeleteFlowRegistryClientConflict{}
}

/*
DeleteFlowRegistryClientConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type DeleteFlowRegistryClientConflict struct {
}

// IsSuccess returns true when this delete flow registry client conflict response has a 2xx status code
func (o *DeleteFlowRegistryClientConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete flow registry client conflict response has a 3xx status code
func (o *DeleteFlowRegistryClientConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete flow registry client conflict response has a 4xx status code
func (o *DeleteFlowRegistryClientConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete flow registry client conflict response has a 5xx status code
func (o *DeleteFlowRegistryClientConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this delete flow registry client conflict response a status code equal to that given
func (o *DeleteFlowRegistryClientConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the delete flow registry client conflict response
func (o *DeleteFlowRegistryClientConflict) Code() int {
	return 409
}

func (o *DeleteFlowRegistryClientConflict) Error() string {
	return fmt.Sprintf("[DELETE /controller/registry-clients/{id}][%d] deleteFlowRegistryClientConflict ", 409)
}

func (o *DeleteFlowRegistryClientConflict) String() string {
	return fmt.Sprintf("[DELETE /controller/registry-clients/{id}][%d] deleteFlowRegistryClientConflict ", 409)
}

func (o *DeleteFlowRegistryClientConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
