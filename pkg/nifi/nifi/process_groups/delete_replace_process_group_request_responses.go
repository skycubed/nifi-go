// Code generated by go-swagger; DO NOT EDIT.

package process_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// DeleteReplaceProcessGroupRequestReader is a Reader for the DeleteReplaceProcessGroupRequest structure.
type DeleteReplaceProcessGroupRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteReplaceProcessGroupRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteReplaceProcessGroupRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteReplaceProcessGroupRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteReplaceProcessGroupRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteReplaceProcessGroupRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteReplaceProcessGroupRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteReplaceProcessGroupRequestConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /process-groups/replace-requests/{id}] deleteReplaceProcessGroupRequest", response, response.Code())
	}
}

// NewDeleteReplaceProcessGroupRequestOK creates a DeleteReplaceProcessGroupRequestOK with default headers values
func NewDeleteReplaceProcessGroupRequestOK() *DeleteReplaceProcessGroupRequestOK {
	return &DeleteReplaceProcessGroupRequestOK{}
}

/*
DeleteReplaceProcessGroupRequestOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteReplaceProcessGroupRequestOK struct {
	Payload *models.ProcessGroupReplaceRequestEntity
}

// IsSuccess returns true when this delete replace process group request o k response has a 2xx status code
func (o *DeleteReplaceProcessGroupRequestOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete replace process group request o k response has a 3xx status code
func (o *DeleteReplaceProcessGroupRequestOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete replace process group request o k response has a 4xx status code
func (o *DeleteReplaceProcessGroupRequestOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete replace process group request o k response has a 5xx status code
func (o *DeleteReplaceProcessGroupRequestOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete replace process group request o k response a status code equal to that given
func (o *DeleteReplaceProcessGroupRequestOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete replace process group request o k response
func (o *DeleteReplaceProcessGroupRequestOK) Code() int {
	return 200
}

func (o *DeleteReplaceProcessGroupRequestOK) Error() string {
	return fmt.Sprintf("[DELETE /process-groups/replace-requests/{id}][%d] deleteReplaceProcessGroupRequestOK  %+v", 200, o.Payload)
}

func (o *DeleteReplaceProcessGroupRequestOK) String() string {
	return fmt.Sprintf("[DELETE /process-groups/replace-requests/{id}][%d] deleteReplaceProcessGroupRequestOK  %+v", 200, o.Payload)
}

func (o *DeleteReplaceProcessGroupRequestOK) GetPayload() *models.ProcessGroupReplaceRequestEntity {
	return o.Payload
}

func (o *DeleteReplaceProcessGroupRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProcessGroupReplaceRequestEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteReplaceProcessGroupRequestBadRequest creates a DeleteReplaceProcessGroupRequestBadRequest with default headers values
func NewDeleteReplaceProcessGroupRequestBadRequest() *DeleteReplaceProcessGroupRequestBadRequest {
	return &DeleteReplaceProcessGroupRequestBadRequest{}
}

/*
DeleteReplaceProcessGroupRequestBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type DeleteReplaceProcessGroupRequestBadRequest struct {
}

// IsSuccess returns true when this delete replace process group request bad request response has a 2xx status code
func (o *DeleteReplaceProcessGroupRequestBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete replace process group request bad request response has a 3xx status code
func (o *DeleteReplaceProcessGroupRequestBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete replace process group request bad request response has a 4xx status code
func (o *DeleteReplaceProcessGroupRequestBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete replace process group request bad request response has a 5xx status code
func (o *DeleteReplaceProcessGroupRequestBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete replace process group request bad request response a status code equal to that given
func (o *DeleteReplaceProcessGroupRequestBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete replace process group request bad request response
func (o *DeleteReplaceProcessGroupRequestBadRequest) Code() int {
	return 400
}

func (o *DeleteReplaceProcessGroupRequestBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /process-groups/replace-requests/{id}][%d] deleteReplaceProcessGroupRequestBadRequest ", 400)
}

func (o *DeleteReplaceProcessGroupRequestBadRequest) String() string {
	return fmt.Sprintf("[DELETE /process-groups/replace-requests/{id}][%d] deleteReplaceProcessGroupRequestBadRequest ", 400)
}

func (o *DeleteReplaceProcessGroupRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteReplaceProcessGroupRequestUnauthorized creates a DeleteReplaceProcessGroupRequestUnauthorized with default headers values
func NewDeleteReplaceProcessGroupRequestUnauthorized() *DeleteReplaceProcessGroupRequestUnauthorized {
	return &DeleteReplaceProcessGroupRequestUnauthorized{}
}

/*
DeleteReplaceProcessGroupRequestUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type DeleteReplaceProcessGroupRequestUnauthorized struct {
}

// IsSuccess returns true when this delete replace process group request unauthorized response has a 2xx status code
func (o *DeleteReplaceProcessGroupRequestUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete replace process group request unauthorized response has a 3xx status code
func (o *DeleteReplaceProcessGroupRequestUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete replace process group request unauthorized response has a 4xx status code
func (o *DeleteReplaceProcessGroupRequestUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete replace process group request unauthorized response has a 5xx status code
func (o *DeleteReplaceProcessGroupRequestUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete replace process group request unauthorized response a status code equal to that given
func (o *DeleteReplaceProcessGroupRequestUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete replace process group request unauthorized response
func (o *DeleteReplaceProcessGroupRequestUnauthorized) Code() int {
	return 401
}

func (o *DeleteReplaceProcessGroupRequestUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /process-groups/replace-requests/{id}][%d] deleteReplaceProcessGroupRequestUnauthorized ", 401)
}

func (o *DeleteReplaceProcessGroupRequestUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /process-groups/replace-requests/{id}][%d] deleteReplaceProcessGroupRequestUnauthorized ", 401)
}

func (o *DeleteReplaceProcessGroupRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteReplaceProcessGroupRequestForbidden creates a DeleteReplaceProcessGroupRequestForbidden with default headers values
func NewDeleteReplaceProcessGroupRequestForbidden() *DeleteReplaceProcessGroupRequestForbidden {
	return &DeleteReplaceProcessGroupRequestForbidden{}
}

/*
DeleteReplaceProcessGroupRequestForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type DeleteReplaceProcessGroupRequestForbidden struct {
}

// IsSuccess returns true when this delete replace process group request forbidden response has a 2xx status code
func (o *DeleteReplaceProcessGroupRequestForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete replace process group request forbidden response has a 3xx status code
func (o *DeleteReplaceProcessGroupRequestForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete replace process group request forbidden response has a 4xx status code
func (o *DeleteReplaceProcessGroupRequestForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete replace process group request forbidden response has a 5xx status code
func (o *DeleteReplaceProcessGroupRequestForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete replace process group request forbidden response a status code equal to that given
func (o *DeleteReplaceProcessGroupRequestForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete replace process group request forbidden response
func (o *DeleteReplaceProcessGroupRequestForbidden) Code() int {
	return 403
}

func (o *DeleteReplaceProcessGroupRequestForbidden) Error() string {
	return fmt.Sprintf("[DELETE /process-groups/replace-requests/{id}][%d] deleteReplaceProcessGroupRequestForbidden ", 403)
}

func (o *DeleteReplaceProcessGroupRequestForbidden) String() string {
	return fmt.Sprintf("[DELETE /process-groups/replace-requests/{id}][%d] deleteReplaceProcessGroupRequestForbidden ", 403)
}

func (o *DeleteReplaceProcessGroupRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteReplaceProcessGroupRequestNotFound creates a DeleteReplaceProcessGroupRequestNotFound with default headers values
func NewDeleteReplaceProcessGroupRequestNotFound() *DeleteReplaceProcessGroupRequestNotFound {
	return &DeleteReplaceProcessGroupRequestNotFound{}
}

/*
DeleteReplaceProcessGroupRequestNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type DeleteReplaceProcessGroupRequestNotFound struct {
}

// IsSuccess returns true when this delete replace process group request not found response has a 2xx status code
func (o *DeleteReplaceProcessGroupRequestNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete replace process group request not found response has a 3xx status code
func (o *DeleteReplaceProcessGroupRequestNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete replace process group request not found response has a 4xx status code
func (o *DeleteReplaceProcessGroupRequestNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete replace process group request not found response has a 5xx status code
func (o *DeleteReplaceProcessGroupRequestNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete replace process group request not found response a status code equal to that given
func (o *DeleteReplaceProcessGroupRequestNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete replace process group request not found response
func (o *DeleteReplaceProcessGroupRequestNotFound) Code() int {
	return 404
}

func (o *DeleteReplaceProcessGroupRequestNotFound) Error() string {
	return fmt.Sprintf("[DELETE /process-groups/replace-requests/{id}][%d] deleteReplaceProcessGroupRequestNotFound ", 404)
}

func (o *DeleteReplaceProcessGroupRequestNotFound) String() string {
	return fmt.Sprintf("[DELETE /process-groups/replace-requests/{id}][%d] deleteReplaceProcessGroupRequestNotFound ", 404)
}

func (o *DeleteReplaceProcessGroupRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteReplaceProcessGroupRequestConflict creates a DeleteReplaceProcessGroupRequestConflict with default headers values
func NewDeleteReplaceProcessGroupRequestConflict() *DeleteReplaceProcessGroupRequestConflict {
	return &DeleteReplaceProcessGroupRequestConflict{}
}

/*
DeleteReplaceProcessGroupRequestConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type DeleteReplaceProcessGroupRequestConflict struct {
}

// IsSuccess returns true when this delete replace process group request conflict response has a 2xx status code
func (o *DeleteReplaceProcessGroupRequestConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete replace process group request conflict response has a 3xx status code
func (o *DeleteReplaceProcessGroupRequestConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete replace process group request conflict response has a 4xx status code
func (o *DeleteReplaceProcessGroupRequestConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete replace process group request conflict response has a 5xx status code
func (o *DeleteReplaceProcessGroupRequestConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this delete replace process group request conflict response a status code equal to that given
func (o *DeleteReplaceProcessGroupRequestConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the delete replace process group request conflict response
func (o *DeleteReplaceProcessGroupRequestConflict) Code() int {
	return 409
}

func (o *DeleteReplaceProcessGroupRequestConflict) Error() string {
	return fmt.Sprintf("[DELETE /process-groups/replace-requests/{id}][%d] deleteReplaceProcessGroupRequestConflict ", 409)
}

func (o *DeleteReplaceProcessGroupRequestConflict) String() string {
	return fmt.Sprintf("[DELETE /process-groups/replace-requests/{id}][%d] deleteReplaceProcessGroupRequestConflict ", 409)
}

func (o *DeleteReplaceProcessGroupRequestConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
