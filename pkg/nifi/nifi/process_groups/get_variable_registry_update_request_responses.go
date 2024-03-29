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

// GetVariableRegistryUpdateRequestReader is a Reader for the GetVariableRegistryUpdateRequest structure.
type GetVariableRegistryUpdateRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVariableRegistryUpdateRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVariableRegistryUpdateRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVariableRegistryUpdateRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetVariableRegistryUpdateRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetVariableRegistryUpdateRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVariableRegistryUpdateRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetVariableRegistryUpdateRequestConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /process-groups/{groupId}/variable-registry/update-requests/{updateId}] getVariableRegistryUpdateRequest", response, response.Code())
	}
}

// NewGetVariableRegistryUpdateRequestOK creates a GetVariableRegistryUpdateRequestOK with default headers values
func NewGetVariableRegistryUpdateRequestOK() *GetVariableRegistryUpdateRequestOK {
	return &GetVariableRegistryUpdateRequestOK{}
}

/*
GetVariableRegistryUpdateRequestOK describes a response with status code 200, with default header values.

successful operation
*/
type GetVariableRegistryUpdateRequestOK struct {
	Payload *models.VariableRegistryUpdateRequestEntity
}

// IsSuccess returns true when this get variable registry update request o k response has a 2xx status code
func (o *GetVariableRegistryUpdateRequestOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get variable registry update request o k response has a 3xx status code
func (o *GetVariableRegistryUpdateRequestOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get variable registry update request o k response has a 4xx status code
func (o *GetVariableRegistryUpdateRequestOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get variable registry update request o k response has a 5xx status code
func (o *GetVariableRegistryUpdateRequestOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get variable registry update request o k response a status code equal to that given
func (o *GetVariableRegistryUpdateRequestOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get variable registry update request o k response
func (o *GetVariableRegistryUpdateRequestOK) Code() int {
	return 200
}

func (o *GetVariableRegistryUpdateRequestOK) Error() string {
	return fmt.Sprintf("[GET /process-groups/{groupId}/variable-registry/update-requests/{updateId}][%d] getVariableRegistryUpdateRequestOK  %+v", 200, o.Payload)
}

func (o *GetVariableRegistryUpdateRequestOK) String() string {
	return fmt.Sprintf("[GET /process-groups/{groupId}/variable-registry/update-requests/{updateId}][%d] getVariableRegistryUpdateRequestOK  %+v", 200, o.Payload)
}

func (o *GetVariableRegistryUpdateRequestOK) GetPayload() *models.VariableRegistryUpdateRequestEntity {
	return o.Payload
}

func (o *GetVariableRegistryUpdateRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VariableRegistryUpdateRequestEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVariableRegistryUpdateRequestBadRequest creates a GetVariableRegistryUpdateRequestBadRequest with default headers values
func NewGetVariableRegistryUpdateRequestBadRequest() *GetVariableRegistryUpdateRequestBadRequest {
	return &GetVariableRegistryUpdateRequestBadRequest{}
}

/*
GetVariableRegistryUpdateRequestBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetVariableRegistryUpdateRequestBadRequest struct {
}

// IsSuccess returns true when this get variable registry update request bad request response has a 2xx status code
func (o *GetVariableRegistryUpdateRequestBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get variable registry update request bad request response has a 3xx status code
func (o *GetVariableRegistryUpdateRequestBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get variable registry update request bad request response has a 4xx status code
func (o *GetVariableRegistryUpdateRequestBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get variable registry update request bad request response has a 5xx status code
func (o *GetVariableRegistryUpdateRequestBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get variable registry update request bad request response a status code equal to that given
func (o *GetVariableRegistryUpdateRequestBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get variable registry update request bad request response
func (o *GetVariableRegistryUpdateRequestBadRequest) Code() int {
	return 400
}

func (o *GetVariableRegistryUpdateRequestBadRequest) Error() string {
	return fmt.Sprintf("[GET /process-groups/{groupId}/variable-registry/update-requests/{updateId}][%d] getVariableRegistryUpdateRequestBadRequest ", 400)
}

func (o *GetVariableRegistryUpdateRequestBadRequest) String() string {
	return fmt.Sprintf("[GET /process-groups/{groupId}/variable-registry/update-requests/{updateId}][%d] getVariableRegistryUpdateRequestBadRequest ", 400)
}

func (o *GetVariableRegistryUpdateRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVariableRegistryUpdateRequestUnauthorized creates a GetVariableRegistryUpdateRequestUnauthorized with default headers values
func NewGetVariableRegistryUpdateRequestUnauthorized() *GetVariableRegistryUpdateRequestUnauthorized {
	return &GetVariableRegistryUpdateRequestUnauthorized{}
}

/*
GetVariableRegistryUpdateRequestUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetVariableRegistryUpdateRequestUnauthorized struct {
}

// IsSuccess returns true when this get variable registry update request unauthorized response has a 2xx status code
func (o *GetVariableRegistryUpdateRequestUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get variable registry update request unauthorized response has a 3xx status code
func (o *GetVariableRegistryUpdateRequestUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get variable registry update request unauthorized response has a 4xx status code
func (o *GetVariableRegistryUpdateRequestUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get variable registry update request unauthorized response has a 5xx status code
func (o *GetVariableRegistryUpdateRequestUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get variable registry update request unauthorized response a status code equal to that given
func (o *GetVariableRegistryUpdateRequestUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get variable registry update request unauthorized response
func (o *GetVariableRegistryUpdateRequestUnauthorized) Code() int {
	return 401
}

func (o *GetVariableRegistryUpdateRequestUnauthorized) Error() string {
	return fmt.Sprintf("[GET /process-groups/{groupId}/variable-registry/update-requests/{updateId}][%d] getVariableRegistryUpdateRequestUnauthorized ", 401)
}

func (o *GetVariableRegistryUpdateRequestUnauthorized) String() string {
	return fmt.Sprintf("[GET /process-groups/{groupId}/variable-registry/update-requests/{updateId}][%d] getVariableRegistryUpdateRequestUnauthorized ", 401)
}

func (o *GetVariableRegistryUpdateRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVariableRegistryUpdateRequestForbidden creates a GetVariableRegistryUpdateRequestForbidden with default headers values
func NewGetVariableRegistryUpdateRequestForbidden() *GetVariableRegistryUpdateRequestForbidden {
	return &GetVariableRegistryUpdateRequestForbidden{}
}

/*
GetVariableRegistryUpdateRequestForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetVariableRegistryUpdateRequestForbidden struct {
}

// IsSuccess returns true when this get variable registry update request forbidden response has a 2xx status code
func (o *GetVariableRegistryUpdateRequestForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get variable registry update request forbidden response has a 3xx status code
func (o *GetVariableRegistryUpdateRequestForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get variable registry update request forbidden response has a 4xx status code
func (o *GetVariableRegistryUpdateRequestForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get variable registry update request forbidden response has a 5xx status code
func (o *GetVariableRegistryUpdateRequestForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get variable registry update request forbidden response a status code equal to that given
func (o *GetVariableRegistryUpdateRequestForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get variable registry update request forbidden response
func (o *GetVariableRegistryUpdateRequestForbidden) Code() int {
	return 403
}

func (o *GetVariableRegistryUpdateRequestForbidden) Error() string {
	return fmt.Sprintf("[GET /process-groups/{groupId}/variable-registry/update-requests/{updateId}][%d] getVariableRegistryUpdateRequestForbidden ", 403)
}

func (o *GetVariableRegistryUpdateRequestForbidden) String() string {
	return fmt.Sprintf("[GET /process-groups/{groupId}/variable-registry/update-requests/{updateId}][%d] getVariableRegistryUpdateRequestForbidden ", 403)
}

func (o *GetVariableRegistryUpdateRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVariableRegistryUpdateRequestNotFound creates a GetVariableRegistryUpdateRequestNotFound with default headers values
func NewGetVariableRegistryUpdateRequestNotFound() *GetVariableRegistryUpdateRequestNotFound {
	return &GetVariableRegistryUpdateRequestNotFound{}
}

/*
GetVariableRegistryUpdateRequestNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetVariableRegistryUpdateRequestNotFound struct {
}

// IsSuccess returns true when this get variable registry update request not found response has a 2xx status code
func (o *GetVariableRegistryUpdateRequestNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get variable registry update request not found response has a 3xx status code
func (o *GetVariableRegistryUpdateRequestNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get variable registry update request not found response has a 4xx status code
func (o *GetVariableRegistryUpdateRequestNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get variable registry update request not found response has a 5xx status code
func (o *GetVariableRegistryUpdateRequestNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get variable registry update request not found response a status code equal to that given
func (o *GetVariableRegistryUpdateRequestNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get variable registry update request not found response
func (o *GetVariableRegistryUpdateRequestNotFound) Code() int {
	return 404
}

func (o *GetVariableRegistryUpdateRequestNotFound) Error() string {
	return fmt.Sprintf("[GET /process-groups/{groupId}/variable-registry/update-requests/{updateId}][%d] getVariableRegistryUpdateRequestNotFound ", 404)
}

func (o *GetVariableRegistryUpdateRequestNotFound) String() string {
	return fmt.Sprintf("[GET /process-groups/{groupId}/variable-registry/update-requests/{updateId}][%d] getVariableRegistryUpdateRequestNotFound ", 404)
}

func (o *GetVariableRegistryUpdateRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVariableRegistryUpdateRequestConflict creates a GetVariableRegistryUpdateRequestConflict with default headers values
func NewGetVariableRegistryUpdateRequestConflict() *GetVariableRegistryUpdateRequestConflict {
	return &GetVariableRegistryUpdateRequestConflict{}
}

/*
GetVariableRegistryUpdateRequestConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetVariableRegistryUpdateRequestConflict struct {
}

// IsSuccess returns true when this get variable registry update request conflict response has a 2xx status code
func (o *GetVariableRegistryUpdateRequestConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get variable registry update request conflict response has a 3xx status code
func (o *GetVariableRegistryUpdateRequestConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get variable registry update request conflict response has a 4xx status code
func (o *GetVariableRegistryUpdateRequestConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get variable registry update request conflict response has a 5xx status code
func (o *GetVariableRegistryUpdateRequestConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get variable registry update request conflict response a status code equal to that given
func (o *GetVariableRegistryUpdateRequestConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the get variable registry update request conflict response
func (o *GetVariableRegistryUpdateRequestConflict) Code() int {
	return 409
}

func (o *GetVariableRegistryUpdateRequestConflict) Error() string {
	return fmt.Sprintf("[GET /process-groups/{groupId}/variable-registry/update-requests/{updateId}][%d] getVariableRegistryUpdateRequestConflict ", 409)
}

func (o *GetVariableRegistryUpdateRequestConflict) String() string {
	return fmt.Sprintf("[GET /process-groups/{groupId}/variable-registry/update-requests/{updateId}][%d] getVariableRegistryUpdateRequestConflict ", 409)
}

func (o *GetVariableRegistryUpdateRequestConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
