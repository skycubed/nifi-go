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

// DeleteRevertRequestReader is a Reader for the DeleteRevertRequest structure.
type DeleteRevertRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRevertRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRevertRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRevertRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteRevertRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteRevertRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRevertRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteRevertRequestConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRevertRequestOK creates a DeleteRevertRequestOK with default headers values
func NewDeleteRevertRequestOK() *DeleteRevertRequestOK {
	return &DeleteRevertRequestOK{}
}

/*
DeleteRevertRequestOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteRevertRequestOK struct {
	Payload *models.VersionedFlowUpdateRequestEntity
}

// IsSuccess returns true when this delete revert request o k response has a 2xx status code
func (o *DeleteRevertRequestOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete revert request o k response has a 3xx status code
func (o *DeleteRevertRequestOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete revert request o k response has a 4xx status code
func (o *DeleteRevertRequestOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete revert request o k response has a 5xx status code
func (o *DeleteRevertRequestOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete revert request o k response a status code equal to that given
func (o *DeleteRevertRequestOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteRevertRequestOK) Error() string {
	return fmt.Sprintf("[DELETE /versions/revert-requests/{id}][%d] deleteRevertRequestOK  %+v", 200, o.Payload)
}

func (o *DeleteRevertRequestOK) String() string {
	return fmt.Sprintf("[DELETE /versions/revert-requests/{id}][%d] deleteRevertRequestOK  %+v", 200, o.Payload)
}

func (o *DeleteRevertRequestOK) GetPayload() *models.VersionedFlowUpdateRequestEntity {
	return o.Payload
}

func (o *DeleteRevertRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VersionedFlowUpdateRequestEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRevertRequestBadRequest creates a DeleteRevertRequestBadRequest with default headers values
func NewDeleteRevertRequestBadRequest() *DeleteRevertRequestBadRequest {
	return &DeleteRevertRequestBadRequest{}
}

/*
DeleteRevertRequestBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type DeleteRevertRequestBadRequest struct {
}

// IsSuccess returns true when this delete revert request bad request response has a 2xx status code
func (o *DeleteRevertRequestBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete revert request bad request response has a 3xx status code
func (o *DeleteRevertRequestBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete revert request bad request response has a 4xx status code
func (o *DeleteRevertRequestBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete revert request bad request response has a 5xx status code
func (o *DeleteRevertRequestBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete revert request bad request response a status code equal to that given
func (o *DeleteRevertRequestBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteRevertRequestBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /versions/revert-requests/{id}][%d] deleteRevertRequestBadRequest ", 400)
}

func (o *DeleteRevertRequestBadRequest) String() string {
	return fmt.Sprintf("[DELETE /versions/revert-requests/{id}][%d] deleteRevertRequestBadRequest ", 400)
}

func (o *DeleteRevertRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRevertRequestUnauthorized creates a DeleteRevertRequestUnauthorized with default headers values
func NewDeleteRevertRequestUnauthorized() *DeleteRevertRequestUnauthorized {
	return &DeleteRevertRequestUnauthorized{}
}

/*
DeleteRevertRequestUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type DeleteRevertRequestUnauthorized struct {
}

// IsSuccess returns true when this delete revert request unauthorized response has a 2xx status code
func (o *DeleteRevertRequestUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete revert request unauthorized response has a 3xx status code
func (o *DeleteRevertRequestUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete revert request unauthorized response has a 4xx status code
func (o *DeleteRevertRequestUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete revert request unauthorized response has a 5xx status code
func (o *DeleteRevertRequestUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete revert request unauthorized response a status code equal to that given
func (o *DeleteRevertRequestUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteRevertRequestUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /versions/revert-requests/{id}][%d] deleteRevertRequestUnauthorized ", 401)
}

func (o *DeleteRevertRequestUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /versions/revert-requests/{id}][%d] deleteRevertRequestUnauthorized ", 401)
}

func (o *DeleteRevertRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRevertRequestForbidden creates a DeleteRevertRequestForbidden with default headers values
func NewDeleteRevertRequestForbidden() *DeleteRevertRequestForbidden {
	return &DeleteRevertRequestForbidden{}
}

/*
DeleteRevertRequestForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type DeleteRevertRequestForbidden struct {
}

// IsSuccess returns true when this delete revert request forbidden response has a 2xx status code
func (o *DeleteRevertRequestForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete revert request forbidden response has a 3xx status code
func (o *DeleteRevertRequestForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete revert request forbidden response has a 4xx status code
func (o *DeleteRevertRequestForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete revert request forbidden response has a 5xx status code
func (o *DeleteRevertRequestForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete revert request forbidden response a status code equal to that given
func (o *DeleteRevertRequestForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteRevertRequestForbidden) Error() string {
	return fmt.Sprintf("[DELETE /versions/revert-requests/{id}][%d] deleteRevertRequestForbidden ", 403)
}

func (o *DeleteRevertRequestForbidden) String() string {
	return fmt.Sprintf("[DELETE /versions/revert-requests/{id}][%d] deleteRevertRequestForbidden ", 403)
}

func (o *DeleteRevertRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRevertRequestNotFound creates a DeleteRevertRequestNotFound with default headers values
func NewDeleteRevertRequestNotFound() *DeleteRevertRequestNotFound {
	return &DeleteRevertRequestNotFound{}
}

/*
DeleteRevertRequestNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type DeleteRevertRequestNotFound struct {
}

// IsSuccess returns true when this delete revert request not found response has a 2xx status code
func (o *DeleteRevertRequestNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete revert request not found response has a 3xx status code
func (o *DeleteRevertRequestNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete revert request not found response has a 4xx status code
func (o *DeleteRevertRequestNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete revert request not found response has a 5xx status code
func (o *DeleteRevertRequestNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete revert request not found response a status code equal to that given
func (o *DeleteRevertRequestNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteRevertRequestNotFound) Error() string {
	return fmt.Sprintf("[DELETE /versions/revert-requests/{id}][%d] deleteRevertRequestNotFound ", 404)
}

func (o *DeleteRevertRequestNotFound) String() string {
	return fmt.Sprintf("[DELETE /versions/revert-requests/{id}][%d] deleteRevertRequestNotFound ", 404)
}

func (o *DeleteRevertRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRevertRequestConflict creates a DeleteRevertRequestConflict with default headers values
func NewDeleteRevertRequestConflict() *DeleteRevertRequestConflict {
	return &DeleteRevertRequestConflict{}
}

/*
DeleteRevertRequestConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type DeleteRevertRequestConflict struct {
}

// IsSuccess returns true when this delete revert request conflict response has a 2xx status code
func (o *DeleteRevertRequestConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete revert request conflict response has a 3xx status code
func (o *DeleteRevertRequestConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete revert request conflict response has a 4xx status code
func (o *DeleteRevertRequestConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete revert request conflict response has a 5xx status code
func (o *DeleteRevertRequestConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this delete revert request conflict response a status code equal to that given
func (o *DeleteRevertRequestConflict) IsCode(code int) bool {
	return code == 409
}

func (o *DeleteRevertRequestConflict) Error() string {
	return fmt.Sprintf("[DELETE /versions/revert-requests/{id}][%d] deleteRevertRequestConflict ", 409)
}

func (o *DeleteRevertRequestConflict) String() string {
	return fmt.Sprintf("[DELETE /versions/revert-requests/{id}][%d] deleteRevertRequestConflict ", 409)
}

func (o *DeleteRevertRequestConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
