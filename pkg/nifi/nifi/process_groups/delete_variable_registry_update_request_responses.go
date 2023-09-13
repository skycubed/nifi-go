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

// DeleteVariableRegistryUpdateRequestReader is a Reader for the DeleteVariableRegistryUpdateRequest structure.
type DeleteVariableRegistryUpdateRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVariableRegistryUpdateRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteVariableRegistryUpdateRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteVariableRegistryUpdateRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteVariableRegistryUpdateRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteVariableRegistryUpdateRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteVariableRegistryUpdateRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteVariableRegistryUpdateRequestConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /process-groups/{groupId}/variable-registry/update-requests/{updateId}] deleteVariableRegistryUpdateRequest", response, response.Code())
	}
}

// NewDeleteVariableRegistryUpdateRequestOK creates a DeleteVariableRegistryUpdateRequestOK with default headers values
func NewDeleteVariableRegistryUpdateRequestOK() *DeleteVariableRegistryUpdateRequestOK {
	return &DeleteVariableRegistryUpdateRequestOK{}
}

/*
DeleteVariableRegistryUpdateRequestOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteVariableRegistryUpdateRequestOK struct {
	Payload *models.VariableRegistryUpdateRequestEntity
}

// IsSuccess returns true when this delete variable registry update request o k response has a 2xx status code
func (o *DeleteVariableRegistryUpdateRequestOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete variable registry update request o k response has a 3xx status code
func (o *DeleteVariableRegistryUpdateRequestOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete variable registry update request o k response has a 4xx status code
func (o *DeleteVariableRegistryUpdateRequestOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete variable registry update request o k response has a 5xx status code
func (o *DeleteVariableRegistryUpdateRequestOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete variable registry update request o k response a status code equal to that given
func (o *DeleteVariableRegistryUpdateRequestOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete variable registry update request o k response
func (o *DeleteVariableRegistryUpdateRequestOK) Code() int {
	return 200
}

func (o *DeleteVariableRegistryUpdateRequestOK) Error() string {
	return fmt.Sprintf("[DELETE /process-groups/{groupId}/variable-registry/update-requests/{updateId}][%d] deleteVariableRegistryUpdateRequestOK  %+v", 200, o.Payload)
}

func (o *DeleteVariableRegistryUpdateRequestOK) String() string {
	return fmt.Sprintf("[DELETE /process-groups/{groupId}/variable-registry/update-requests/{updateId}][%d] deleteVariableRegistryUpdateRequestOK  %+v", 200, o.Payload)
}

func (o *DeleteVariableRegistryUpdateRequestOK) GetPayload() *models.VariableRegistryUpdateRequestEntity {
	return o.Payload
}

func (o *DeleteVariableRegistryUpdateRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VariableRegistryUpdateRequestEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVariableRegistryUpdateRequestBadRequest creates a DeleteVariableRegistryUpdateRequestBadRequest with default headers values
func NewDeleteVariableRegistryUpdateRequestBadRequest() *DeleteVariableRegistryUpdateRequestBadRequest {
	return &DeleteVariableRegistryUpdateRequestBadRequest{}
}

/*
DeleteVariableRegistryUpdateRequestBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type DeleteVariableRegistryUpdateRequestBadRequest struct {
}

// IsSuccess returns true when this delete variable registry update request bad request response has a 2xx status code
func (o *DeleteVariableRegistryUpdateRequestBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete variable registry update request bad request response has a 3xx status code
func (o *DeleteVariableRegistryUpdateRequestBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete variable registry update request bad request response has a 4xx status code
func (o *DeleteVariableRegistryUpdateRequestBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete variable registry update request bad request response has a 5xx status code
func (o *DeleteVariableRegistryUpdateRequestBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete variable registry update request bad request response a status code equal to that given
func (o *DeleteVariableRegistryUpdateRequestBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete variable registry update request bad request response
func (o *DeleteVariableRegistryUpdateRequestBadRequest) Code() int {
	return 400
}

func (o *DeleteVariableRegistryUpdateRequestBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /process-groups/{groupId}/variable-registry/update-requests/{updateId}][%d] deleteVariableRegistryUpdateRequestBadRequest ", 400)
}

func (o *DeleteVariableRegistryUpdateRequestBadRequest) String() string {
	return fmt.Sprintf("[DELETE /process-groups/{groupId}/variable-registry/update-requests/{updateId}][%d] deleteVariableRegistryUpdateRequestBadRequest ", 400)
}

func (o *DeleteVariableRegistryUpdateRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVariableRegistryUpdateRequestUnauthorized creates a DeleteVariableRegistryUpdateRequestUnauthorized with default headers values
func NewDeleteVariableRegistryUpdateRequestUnauthorized() *DeleteVariableRegistryUpdateRequestUnauthorized {
	return &DeleteVariableRegistryUpdateRequestUnauthorized{}
}

/*
DeleteVariableRegistryUpdateRequestUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type DeleteVariableRegistryUpdateRequestUnauthorized struct {
}

// IsSuccess returns true when this delete variable registry update request unauthorized response has a 2xx status code
func (o *DeleteVariableRegistryUpdateRequestUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete variable registry update request unauthorized response has a 3xx status code
func (o *DeleteVariableRegistryUpdateRequestUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete variable registry update request unauthorized response has a 4xx status code
func (o *DeleteVariableRegistryUpdateRequestUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete variable registry update request unauthorized response has a 5xx status code
func (o *DeleteVariableRegistryUpdateRequestUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete variable registry update request unauthorized response a status code equal to that given
func (o *DeleteVariableRegistryUpdateRequestUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete variable registry update request unauthorized response
func (o *DeleteVariableRegistryUpdateRequestUnauthorized) Code() int {
	return 401
}

func (o *DeleteVariableRegistryUpdateRequestUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /process-groups/{groupId}/variable-registry/update-requests/{updateId}][%d] deleteVariableRegistryUpdateRequestUnauthorized ", 401)
}

func (o *DeleteVariableRegistryUpdateRequestUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /process-groups/{groupId}/variable-registry/update-requests/{updateId}][%d] deleteVariableRegistryUpdateRequestUnauthorized ", 401)
}

func (o *DeleteVariableRegistryUpdateRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVariableRegistryUpdateRequestForbidden creates a DeleteVariableRegistryUpdateRequestForbidden with default headers values
func NewDeleteVariableRegistryUpdateRequestForbidden() *DeleteVariableRegistryUpdateRequestForbidden {
	return &DeleteVariableRegistryUpdateRequestForbidden{}
}

/*
DeleteVariableRegistryUpdateRequestForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type DeleteVariableRegistryUpdateRequestForbidden struct {
}

// IsSuccess returns true when this delete variable registry update request forbidden response has a 2xx status code
func (o *DeleteVariableRegistryUpdateRequestForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete variable registry update request forbidden response has a 3xx status code
func (o *DeleteVariableRegistryUpdateRequestForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete variable registry update request forbidden response has a 4xx status code
func (o *DeleteVariableRegistryUpdateRequestForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete variable registry update request forbidden response has a 5xx status code
func (o *DeleteVariableRegistryUpdateRequestForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete variable registry update request forbidden response a status code equal to that given
func (o *DeleteVariableRegistryUpdateRequestForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete variable registry update request forbidden response
func (o *DeleteVariableRegistryUpdateRequestForbidden) Code() int {
	return 403
}

func (o *DeleteVariableRegistryUpdateRequestForbidden) Error() string {
	return fmt.Sprintf("[DELETE /process-groups/{groupId}/variable-registry/update-requests/{updateId}][%d] deleteVariableRegistryUpdateRequestForbidden ", 403)
}

func (o *DeleteVariableRegistryUpdateRequestForbidden) String() string {
	return fmt.Sprintf("[DELETE /process-groups/{groupId}/variable-registry/update-requests/{updateId}][%d] deleteVariableRegistryUpdateRequestForbidden ", 403)
}

func (o *DeleteVariableRegistryUpdateRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVariableRegistryUpdateRequestNotFound creates a DeleteVariableRegistryUpdateRequestNotFound with default headers values
func NewDeleteVariableRegistryUpdateRequestNotFound() *DeleteVariableRegistryUpdateRequestNotFound {
	return &DeleteVariableRegistryUpdateRequestNotFound{}
}

/*
DeleteVariableRegistryUpdateRequestNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type DeleteVariableRegistryUpdateRequestNotFound struct {
}

// IsSuccess returns true when this delete variable registry update request not found response has a 2xx status code
func (o *DeleteVariableRegistryUpdateRequestNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete variable registry update request not found response has a 3xx status code
func (o *DeleteVariableRegistryUpdateRequestNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete variable registry update request not found response has a 4xx status code
func (o *DeleteVariableRegistryUpdateRequestNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete variable registry update request not found response has a 5xx status code
func (o *DeleteVariableRegistryUpdateRequestNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete variable registry update request not found response a status code equal to that given
func (o *DeleteVariableRegistryUpdateRequestNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete variable registry update request not found response
func (o *DeleteVariableRegistryUpdateRequestNotFound) Code() int {
	return 404
}

func (o *DeleteVariableRegistryUpdateRequestNotFound) Error() string {
	return fmt.Sprintf("[DELETE /process-groups/{groupId}/variable-registry/update-requests/{updateId}][%d] deleteVariableRegistryUpdateRequestNotFound ", 404)
}

func (o *DeleteVariableRegistryUpdateRequestNotFound) String() string {
	return fmt.Sprintf("[DELETE /process-groups/{groupId}/variable-registry/update-requests/{updateId}][%d] deleteVariableRegistryUpdateRequestNotFound ", 404)
}

func (o *DeleteVariableRegistryUpdateRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVariableRegistryUpdateRequestConflict creates a DeleteVariableRegistryUpdateRequestConflict with default headers values
func NewDeleteVariableRegistryUpdateRequestConflict() *DeleteVariableRegistryUpdateRequestConflict {
	return &DeleteVariableRegistryUpdateRequestConflict{}
}

/*
DeleteVariableRegistryUpdateRequestConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type DeleteVariableRegistryUpdateRequestConflict struct {
}

// IsSuccess returns true when this delete variable registry update request conflict response has a 2xx status code
func (o *DeleteVariableRegistryUpdateRequestConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete variable registry update request conflict response has a 3xx status code
func (o *DeleteVariableRegistryUpdateRequestConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete variable registry update request conflict response has a 4xx status code
func (o *DeleteVariableRegistryUpdateRequestConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete variable registry update request conflict response has a 5xx status code
func (o *DeleteVariableRegistryUpdateRequestConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this delete variable registry update request conflict response a status code equal to that given
func (o *DeleteVariableRegistryUpdateRequestConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the delete variable registry update request conflict response
func (o *DeleteVariableRegistryUpdateRequestConflict) Code() int {
	return 409
}

func (o *DeleteVariableRegistryUpdateRequestConflict) Error() string {
	return fmt.Sprintf("[DELETE /process-groups/{groupId}/variable-registry/update-requests/{updateId}][%d] deleteVariableRegistryUpdateRequestConflict ", 409)
}

func (o *DeleteVariableRegistryUpdateRequestConflict) String() string {
	return fmt.Sprintf("[DELETE /process-groups/{groupId}/variable-registry/update-requests/{updateId}][%d] deleteVariableRegistryUpdateRequestConflict ", 409)
}

func (o *DeleteVariableRegistryUpdateRequestConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
