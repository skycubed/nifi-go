// Code generated by go-swagger; DO NOT EDIT.

package reporting_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// ClearStateReader is a Reader for the ClearState structure.
type ClearStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClearStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClearStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewClearStateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewClearStateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewClearStateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewClearStateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewClearStateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewClearStateOK creates a ClearStateOK with default headers values
func NewClearStateOK() *ClearStateOK {
	return &ClearStateOK{}
}

/*
ClearStateOK describes a response with status code 200, with default header values.

successful operation
*/
type ClearStateOK struct {
	Payload *models.ComponentStateEntity
}

// IsSuccess returns true when this clear state o k response has a 2xx status code
func (o *ClearStateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this clear state o k response has a 3xx status code
func (o *ClearStateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this clear state o k response has a 4xx status code
func (o *ClearStateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this clear state o k response has a 5xx status code
func (o *ClearStateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this clear state o k response a status code equal to that given
func (o *ClearStateOK) IsCode(code int) bool {
	return code == 200
}

func (o *ClearStateOK) Error() string {
	return fmt.Sprintf("[POST /reporting-tasks/{id}/state/clear-requests][%d] clearStateOK  %+v", 200, o.Payload)
}

func (o *ClearStateOK) String() string {
	return fmt.Sprintf("[POST /reporting-tasks/{id}/state/clear-requests][%d] clearStateOK  %+v", 200, o.Payload)
}

func (o *ClearStateOK) GetPayload() *models.ComponentStateEntity {
	return o.Payload
}

func (o *ClearStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComponentStateEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClearStateBadRequest creates a ClearStateBadRequest with default headers values
func NewClearStateBadRequest() *ClearStateBadRequest {
	return &ClearStateBadRequest{}
}

/*
ClearStateBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type ClearStateBadRequest struct {
}

// IsSuccess returns true when this clear state bad request response has a 2xx status code
func (o *ClearStateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this clear state bad request response has a 3xx status code
func (o *ClearStateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this clear state bad request response has a 4xx status code
func (o *ClearStateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this clear state bad request response has a 5xx status code
func (o *ClearStateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this clear state bad request response a status code equal to that given
func (o *ClearStateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ClearStateBadRequest) Error() string {
	return fmt.Sprintf("[POST /reporting-tasks/{id}/state/clear-requests][%d] clearStateBadRequest ", 400)
}

func (o *ClearStateBadRequest) String() string {
	return fmt.Sprintf("[POST /reporting-tasks/{id}/state/clear-requests][%d] clearStateBadRequest ", 400)
}

func (o *ClearStateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewClearStateUnauthorized creates a ClearStateUnauthorized with default headers values
func NewClearStateUnauthorized() *ClearStateUnauthorized {
	return &ClearStateUnauthorized{}
}

/*
ClearStateUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type ClearStateUnauthorized struct {
}

// IsSuccess returns true when this clear state unauthorized response has a 2xx status code
func (o *ClearStateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this clear state unauthorized response has a 3xx status code
func (o *ClearStateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this clear state unauthorized response has a 4xx status code
func (o *ClearStateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this clear state unauthorized response has a 5xx status code
func (o *ClearStateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this clear state unauthorized response a status code equal to that given
func (o *ClearStateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ClearStateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /reporting-tasks/{id}/state/clear-requests][%d] clearStateUnauthorized ", 401)
}

func (o *ClearStateUnauthorized) String() string {
	return fmt.Sprintf("[POST /reporting-tasks/{id}/state/clear-requests][%d] clearStateUnauthorized ", 401)
}

func (o *ClearStateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewClearStateForbidden creates a ClearStateForbidden with default headers values
func NewClearStateForbidden() *ClearStateForbidden {
	return &ClearStateForbidden{}
}

/*
ClearStateForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type ClearStateForbidden struct {
}

// IsSuccess returns true when this clear state forbidden response has a 2xx status code
func (o *ClearStateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this clear state forbidden response has a 3xx status code
func (o *ClearStateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this clear state forbidden response has a 4xx status code
func (o *ClearStateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this clear state forbidden response has a 5xx status code
func (o *ClearStateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this clear state forbidden response a status code equal to that given
func (o *ClearStateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ClearStateForbidden) Error() string {
	return fmt.Sprintf("[POST /reporting-tasks/{id}/state/clear-requests][%d] clearStateForbidden ", 403)
}

func (o *ClearStateForbidden) String() string {
	return fmt.Sprintf("[POST /reporting-tasks/{id}/state/clear-requests][%d] clearStateForbidden ", 403)
}

func (o *ClearStateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewClearStateNotFound creates a ClearStateNotFound with default headers values
func NewClearStateNotFound() *ClearStateNotFound {
	return &ClearStateNotFound{}
}

/*
ClearStateNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type ClearStateNotFound struct {
}

// IsSuccess returns true when this clear state not found response has a 2xx status code
func (o *ClearStateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this clear state not found response has a 3xx status code
func (o *ClearStateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this clear state not found response has a 4xx status code
func (o *ClearStateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this clear state not found response has a 5xx status code
func (o *ClearStateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this clear state not found response a status code equal to that given
func (o *ClearStateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ClearStateNotFound) Error() string {
	return fmt.Sprintf("[POST /reporting-tasks/{id}/state/clear-requests][%d] clearStateNotFound ", 404)
}

func (o *ClearStateNotFound) String() string {
	return fmt.Sprintf("[POST /reporting-tasks/{id}/state/clear-requests][%d] clearStateNotFound ", 404)
}

func (o *ClearStateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewClearStateConflict creates a ClearStateConflict with default headers values
func NewClearStateConflict() *ClearStateConflict {
	return &ClearStateConflict{}
}

/*
ClearStateConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type ClearStateConflict struct {
}

// IsSuccess returns true when this clear state conflict response has a 2xx status code
func (o *ClearStateConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this clear state conflict response has a 3xx status code
func (o *ClearStateConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this clear state conflict response has a 4xx status code
func (o *ClearStateConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this clear state conflict response has a 5xx status code
func (o *ClearStateConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this clear state conflict response a status code equal to that given
func (o *ClearStateConflict) IsCode(code int) bool {
	return code == 409
}

func (o *ClearStateConflict) Error() string {
	return fmt.Sprintf("[POST /reporting-tasks/{id}/state/clear-requests][%d] clearStateConflict ", 409)
}

func (o *ClearStateConflict) String() string {
	return fmt.Sprintf("[POST /reporting-tasks/{id}/state/clear-requests][%d] clearStateConflict ", 409)
}

func (o *ClearStateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
