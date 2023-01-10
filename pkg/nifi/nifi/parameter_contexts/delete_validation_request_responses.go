// Code generated by go-swagger; DO NOT EDIT.

package parameter_contexts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// DeleteValidationRequestReader is a Reader for the DeleteValidationRequest structure.
type DeleteValidationRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteValidationRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteValidationRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteValidationRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteValidationRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteValidationRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteValidationRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteValidationRequestConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteValidationRequestOK creates a DeleteValidationRequestOK with default headers values
func NewDeleteValidationRequestOK() *DeleteValidationRequestOK {
	return &DeleteValidationRequestOK{}
}

/*
DeleteValidationRequestOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteValidationRequestOK struct {
	Payload *models.ParameterContextValidationRequestEntity
}

// IsSuccess returns true when this delete validation request o k response has a 2xx status code
func (o *DeleteValidationRequestOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete validation request o k response has a 3xx status code
func (o *DeleteValidationRequestOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete validation request o k response has a 4xx status code
func (o *DeleteValidationRequestOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete validation request o k response has a 5xx status code
func (o *DeleteValidationRequestOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete validation request o k response a status code equal to that given
func (o *DeleteValidationRequestOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteValidationRequestOK) Error() string {
	return fmt.Sprintf("[DELETE /parameter-contexts/{contextId}/validation-requests/{id}][%d] deleteValidationRequestOK  %+v", 200, o.Payload)
}

func (o *DeleteValidationRequestOK) String() string {
	return fmt.Sprintf("[DELETE /parameter-contexts/{contextId}/validation-requests/{id}][%d] deleteValidationRequestOK  %+v", 200, o.Payload)
}

func (o *DeleteValidationRequestOK) GetPayload() *models.ParameterContextValidationRequestEntity {
	return o.Payload
}

func (o *DeleteValidationRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ParameterContextValidationRequestEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteValidationRequestBadRequest creates a DeleteValidationRequestBadRequest with default headers values
func NewDeleteValidationRequestBadRequest() *DeleteValidationRequestBadRequest {
	return &DeleteValidationRequestBadRequest{}
}

/*
DeleteValidationRequestBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type DeleteValidationRequestBadRequest struct {
}

// IsSuccess returns true when this delete validation request bad request response has a 2xx status code
func (o *DeleteValidationRequestBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete validation request bad request response has a 3xx status code
func (o *DeleteValidationRequestBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete validation request bad request response has a 4xx status code
func (o *DeleteValidationRequestBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete validation request bad request response has a 5xx status code
func (o *DeleteValidationRequestBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete validation request bad request response a status code equal to that given
func (o *DeleteValidationRequestBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteValidationRequestBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /parameter-contexts/{contextId}/validation-requests/{id}][%d] deleteValidationRequestBadRequest ", 400)
}

func (o *DeleteValidationRequestBadRequest) String() string {
	return fmt.Sprintf("[DELETE /parameter-contexts/{contextId}/validation-requests/{id}][%d] deleteValidationRequestBadRequest ", 400)
}

func (o *DeleteValidationRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteValidationRequestUnauthorized creates a DeleteValidationRequestUnauthorized with default headers values
func NewDeleteValidationRequestUnauthorized() *DeleteValidationRequestUnauthorized {
	return &DeleteValidationRequestUnauthorized{}
}

/*
DeleteValidationRequestUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type DeleteValidationRequestUnauthorized struct {
}

// IsSuccess returns true when this delete validation request unauthorized response has a 2xx status code
func (o *DeleteValidationRequestUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete validation request unauthorized response has a 3xx status code
func (o *DeleteValidationRequestUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete validation request unauthorized response has a 4xx status code
func (o *DeleteValidationRequestUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete validation request unauthorized response has a 5xx status code
func (o *DeleteValidationRequestUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete validation request unauthorized response a status code equal to that given
func (o *DeleteValidationRequestUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteValidationRequestUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /parameter-contexts/{contextId}/validation-requests/{id}][%d] deleteValidationRequestUnauthorized ", 401)
}

func (o *DeleteValidationRequestUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /parameter-contexts/{contextId}/validation-requests/{id}][%d] deleteValidationRequestUnauthorized ", 401)
}

func (o *DeleteValidationRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteValidationRequestForbidden creates a DeleteValidationRequestForbidden with default headers values
func NewDeleteValidationRequestForbidden() *DeleteValidationRequestForbidden {
	return &DeleteValidationRequestForbidden{}
}

/*
DeleteValidationRequestForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type DeleteValidationRequestForbidden struct {
}

// IsSuccess returns true when this delete validation request forbidden response has a 2xx status code
func (o *DeleteValidationRequestForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete validation request forbidden response has a 3xx status code
func (o *DeleteValidationRequestForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete validation request forbidden response has a 4xx status code
func (o *DeleteValidationRequestForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete validation request forbidden response has a 5xx status code
func (o *DeleteValidationRequestForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete validation request forbidden response a status code equal to that given
func (o *DeleteValidationRequestForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteValidationRequestForbidden) Error() string {
	return fmt.Sprintf("[DELETE /parameter-contexts/{contextId}/validation-requests/{id}][%d] deleteValidationRequestForbidden ", 403)
}

func (o *DeleteValidationRequestForbidden) String() string {
	return fmt.Sprintf("[DELETE /parameter-contexts/{contextId}/validation-requests/{id}][%d] deleteValidationRequestForbidden ", 403)
}

func (o *DeleteValidationRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteValidationRequestNotFound creates a DeleteValidationRequestNotFound with default headers values
func NewDeleteValidationRequestNotFound() *DeleteValidationRequestNotFound {
	return &DeleteValidationRequestNotFound{}
}

/*
DeleteValidationRequestNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type DeleteValidationRequestNotFound struct {
}

// IsSuccess returns true when this delete validation request not found response has a 2xx status code
func (o *DeleteValidationRequestNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete validation request not found response has a 3xx status code
func (o *DeleteValidationRequestNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete validation request not found response has a 4xx status code
func (o *DeleteValidationRequestNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete validation request not found response has a 5xx status code
func (o *DeleteValidationRequestNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete validation request not found response a status code equal to that given
func (o *DeleteValidationRequestNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteValidationRequestNotFound) Error() string {
	return fmt.Sprintf("[DELETE /parameter-contexts/{contextId}/validation-requests/{id}][%d] deleteValidationRequestNotFound ", 404)
}

func (o *DeleteValidationRequestNotFound) String() string {
	return fmt.Sprintf("[DELETE /parameter-contexts/{contextId}/validation-requests/{id}][%d] deleteValidationRequestNotFound ", 404)
}

func (o *DeleteValidationRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteValidationRequestConflict creates a DeleteValidationRequestConflict with default headers values
func NewDeleteValidationRequestConflict() *DeleteValidationRequestConflict {
	return &DeleteValidationRequestConflict{}
}

/*
DeleteValidationRequestConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type DeleteValidationRequestConflict struct {
}

// IsSuccess returns true when this delete validation request conflict response has a 2xx status code
func (o *DeleteValidationRequestConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete validation request conflict response has a 3xx status code
func (o *DeleteValidationRequestConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete validation request conflict response has a 4xx status code
func (o *DeleteValidationRequestConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete validation request conflict response has a 5xx status code
func (o *DeleteValidationRequestConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this delete validation request conflict response a status code equal to that given
func (o *DeleteValidationRequestConflict) IsCode(code int) bool {
	return code == 409
}

func (o *DeleteValidationRequestConflict) Error() string {
	return fmt.Sprintf("[DELETE /parameter-contexts/{contextId}/validation-requests/{id}][%d] deleteValidationRequestConflict ", 409)
}

func (o *DeleteValidationRequestConflict) String() string {
	return fmt.Sprintf("[DELETE /parameter-contexts/{contextId}/validation-requests/{id}][%d] deleteValidationRequestConflict ", 409)
}

func (o *DeleteValidationRequestConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
