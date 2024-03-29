// Code generated by go-swagger; DO NOT EDIT.

package versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// DeleteUpdateRequestReader is a Reader for the DeleteUpdateRequest structure.
type DeleteUpdateRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUpdateRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteUpdateRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteUpdateRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteUpdateRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteUpdateRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteUpdateRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteUpdateRequestConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /versions/update-requests/{id}] deleteUpdateRequest", response, response.Code())
	}
}

// NewDeleteUpdateRequestOK creates a DeleteUpdateRequestOK with default headers values
func NewDeleteUpdateRequestOK() *DeleteUpdateRequestOK {
	return &DeleteUpdateRequestOK{}
}

/*
DeleteUpdateRequestOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteUpdateRequestOK struct {
	Payload *models.VersionedFlowUpdateRequestEntity
}

// IsSuccess returns true when this delete update request o k response has a 2xx status code
func (o *DeleteUpdateRequestOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete update request o k response has a 3xx status code
func (o *DeleteUpdateRequestOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete update request o k response has a 4xx status code
func (o *DeleteUpdateRequestOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete update request o k response has a 5xx status code
func (o *DeleteUpdateRequestOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete update request o k response a status code equal to that given
func (o *DeleteUpdateRequestOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete update request o k response
func (o *DeleteUpdateRequestOK) Code() int {
	return 200
}

func (o *DeleteUpdateRequestOK) Error() string {
	return fmt.Sprintf("[DELETE /versions/update-requests/{id}][%d] deleteUpdateRequestOK  %+v", 200, o.Payload)
}

func (o *DeleteUpdateRequestOK) String() string {
	return fmt.Sprintf("[DELETE /versions/update-requests/{id}][%d] deleteUpdateRequestOK  %+v", 200, o.Payload)
}

func (o *DeleteUpdateRequestOK) GetPayload() *models.VersionedFlowUpdateRequestEntity {
	return o.Payload
}

func (o *DeleteUpdateRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VersionedFlowUpdateRequestEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUpdateRequestBadRequest creates a DeleteUpdateRequestBadRequest with default headers values
func NewDeleteUpdateRequestBadRequest() *DeleteUpdateRequestBadRequest {
	return &DeleteUpdateRequestBadRequest{}
}

/*
DeleteUpdateRequestBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type DeleteUpdateRequestBadRequest struct {
}

// IsSuccess returns true when this delete update request bad request response has a 2xx status code
func (o *DeleteUpdateRequestBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete update request bad request response has a 3xx status code
func (o *DeleteUpdateRequestBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete update request bad request response has a 4xx status code
func (o *DeleteUpdateRequestBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete update request bad request response has a 5xx status code
func (o *DeleteUpdateRequestBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete update request bad request response a status code equal to that given
func (o *DeleteUpdateRequestBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete update request bad request response
func (o *DeleteUpdateRequestBadRequest) Code() int {
	return 400
}

func (o *DeleteUpdateRequestBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /versions/update-requests/{id}][%d] deleteUpdateRequestBadRequest ", 400)
}

func (o *DeleteUpdateRequestBadRequest) String() string {
	return fmt.Sprintf("[DELETE /versions/update-requests/{id}][%d] deleteUpdateRequestBadRequest ", 400)
}

func (o *DeleteUpdateRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUpdateRequestUnauthorized creates a DeleteUpdateRequestUnauthorized with default headers values
func NewDeleteUpdateRequestUnauthorized() *DeleteUpdateRequestUnauthorized {
	return &DeleteUpdateRequestUnauthorized{}
}

/*
DeleteUpdateRequestUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type DeleteUpdateRequestUnauthorized struct {
}

// IsSuccess returns true when this delete update request unauthorized response has a 2xx status code
func (o *DeleteUpdateRequestUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete update request unauthorized response has a 3xx status code
func (o *DeleteUpdateRequestUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete update request unauthorized response has a 4xx status code
func (o *DeleteUpdateRequestUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete update request unauthorized response has a 5xx status code
func (o *DeleteUpdateRequestUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete update request unauthorized response a status code equal to that given
func (o *DeleteUpdateRequestUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete update request unauthorized response
func (o *DeleteUpdateRequestUnauthorized) Code() int {
	return 401
}

func (o *DeleteUpdateRequestUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /versions/update-requests/{id}][%d] deleteUpdateRequestUnauthorized ", 401)
}

func (o *DeleteUpdateRequestUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /versions/update-requests/{id}][%d] deleteUpdateRequestUnauthorized ", 401)
}

func (o *DeleteUpdateRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUpdateRequestForbidden creates a DeleteUpdateRequestForbidden with default headers values
func NewDeleteUpdateRequestForbidden() *DeleteUpdateRequestForbidden {
	return &DeleteUpdateRequestForbidden{}
}

/*
DeleteUpdateRequestForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type DeleteUpdateRequestForbidden struct {
}

// IsSuccess returns true when this delete update request forbidden response has a 2xx status code
func (o *DeleteUpdateRequestForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete update request forbidden response has a 3xx status code
func (o *DeleteUpdateRequestForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete update request forbidden response has a 4xx status code
func (o *DeleteUpdateRequestForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete update request forbidden response has a 5xx status code
func (o *DeleteUpdateRequestForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete update request forbidden response a status code equal to that given
func (o *DeleteUpdateRequestForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete update request forbidden response
func (o *DeleteUpdateRequestForbidden) Code() int {
	return 403
}

func (o *DeleteUpdateRequestForbidden) Error() string {
	return fmt.Sprintf("[DELETE /versions/update-requests/{id}][%d] deleteUpdateRequestForbidden ", 403)
}

func (o *DeleteUpdateRequestForbidden) String() string {
	return fmt.Sprintf("[DELETE /versions/update-requests/{id}][%d] deleteUpdateRequestForbidden ", 403)
}

func (o *DeleteUpdateRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUpdateRequestNotFound creates a DeleteUpdateRequestNotFound with default headers values
func NewDeleteUpdateRequestNotFound() *DeleteUpdateRequestNotFound {
	return &DeleteUpdateRequestNotFound{}
}

/*
DeleteUpdateRequestNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type DeleteUpdateRequestNotFound struct {
}

// IsSuccess returns true when this delete update request not found response has a 2xx status code
func (o *DeleteUpdateRequestNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete update request not found response has a 3xx status code
func (o *DeleteUpdateRequestNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete update request not found response has a 4xx status code
func (o *DeleteUpdateRequestNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete update request not found response has a 5xx status code
func (o *DeleteUpdateRequestNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete update request not found response a status code equal to that given
func (o *DeleteUpdateRequestNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete update request not found response
func (o *DeleteUpdateRequestNotFound) Code() int {
	return 404
}

func (o *DeleteUpdateRequestNotFound) Error() string {
	return fmt.Sprintf("[DELETE /versions/update-requests/{id}][%d] deleteUpdateRequestNotFound ", 404)
}

func (o *DeleteUpdateRequestNotFound) String() string {
	return fmt.Sprintf("[DELETE /versions/update-requests/{id}][%d] deleteUpdateRequestNotFound ", 404)
}

func (o *DeleteUpdateRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUpdateRequestConflict creates a DeleteUpdateRequestConflict with default headers values
func NewDeleteUpdateRequestConflict() *DeleteUpdateRequestConflict {
	return &DeleteUpdateRequestConflict{}
}

/*
DeleteUpdateRequestConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type DeleteUpdateRequestConflict struct {
}

// IsSuccess returns true when this delete update request conflict response has a 2xx status code
func (o *DeleteUpdateRequestConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete update request conflict response has a 3xx status code
func (o *DeleteUpdateRequestConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete update request conflict response has a 4xx status code
func (o *DeleteUpdateRequestConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete update request conflict response has a 5xx status code
func (o *DeleteUpdateRequestConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this delete update request conflict response a status code equal to that given
func (o *DeleteUpdateRequestConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the delete update request conflict response
func (o *DeleteUpdateRequestConflict) Code() int {
	return 409
}

func (o *DeleteUpdateRequestConflict) Error() string {
	return fmt.Sprintf("[DELETE /versions/update-requests/{id}][%d] deleteUpdateRequestConflict ", 409)
}

func (o *DeleteUpdateRequestConflict) String() string {
	return fmt.Sprintf("[DELETE /versions/update-requests/{id}][%d] deleteUpdateRequestConflict ", 409)
}

func (o *DeleteUpdateRequestConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
