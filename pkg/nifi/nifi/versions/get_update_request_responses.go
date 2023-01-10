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

// GetUpdateRequestReader is a Reader for the GetUpdateRequest structure.
type GetUpdateRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUpdateRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUpdateRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUpdateRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUpdateRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUpdateRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUpdateRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetUpdateRequestConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUpdateRequestOK creates a GetUpdateRequestOK with default headers values
func NewGetUpdateRequestOK() *GetUpdateRequestOK {
	return &GetUpdateRequestOK{}
}

/*
GetUpdateRequestOK describes a response with status code 200, with default header values.

successful operation
*/
type GetUpdateRequestOK struct {
	Payload *models.VersionedFlowUpdateRequestEntity
}

// IsSuccess returns true when this get update request o k response has a 2xx status code
func (o *GetUpdateRequestOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get update request o k response has a 3xx status code
func (o *GetUpdateRequestOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get update request o k response has a 4xx status code
func (o *GetUpdateRequestOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get update request o k response has a 5xx status code
func (o *GetUpdateRequestOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get update request o k response a status code equal to that given
func (o *GetUpdateRequestOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetUpdateRequestOK) Error() string {
	return fmt.Sprintf("[GET /versions/update-requests/{id}][%d] getUpdateRequestOK  %+v", 200, o.Payload)
}

func (o *GetUpdateRequestOK) String() string {
	return fmt.Sprintf("[GET /versions/update-requests/{id}][%d] getUpdateRequestOK  %+v", 200, o.Payload)
}

func (o *GetUpdateRequestOK) GetPayload() *models.VersionedFlowUpdateRequestEntity {
	return o.Payload
}

func (o *GetUpdateRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VersionedFlowUpdateRequestEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUpdateRequestBadRequest creates a GetUpdateRequestBadRequest with default headers values
func NewGetUpdateRequestBadRequest() *GetUpdateRequestBadRequest {
	return &GetUpdateRequestBadRequest{}
}

/*
GetUpdateRequestBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetUpdateRequestBadRequest struct {
}

// IsSuccess returns true when this get update request bad request response has a 2xx status code
func (o *GetUpdateRequestBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get update request bad request response has a 3xx status code
func (o *GetUpdateRequestBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get update request bad request response has a 4xx status code
func (o *GetUpdateRequestBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get update request bad request response has a 5xx status code
func (o *GetUpdateRequestBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get update request bad request response a status code equal to that given
func (o *GetUpdateRequestBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetUpdateRequestBadRequest) Error() string {
	return fmt.Sprintf("[GET /versions/update-requests/{id}][%d] getUpdateRequestBadRequest ", 400)
}

func (o *GetUpdateRequestBadRequest) String() string {
	return fmt.Sprintf("[GET /versions/update-requests/{id}][%d] getUpdateRequestBadRequest ", 400)
}

func (o *GetUpdateRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUpdateRequestUnauthorized creates a GetUpdateRequestUnauthorized with default headers values
func NewGetUpdateRequestUnauthorized() *GetUpdateRequestUnauthorized {
	return &GetUpdateRequestUnauthorized{}
}

/*
GetUpdateRequestUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetUpdateRequestUnauthorized struct {
}

// IsSuccess returns true when this get update request unauthorized response has a 2xx status code
func (o *GetUpdateRequestUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get update request unauthorized response has a 3xx status code
func (o *GetUpdateRequestUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get update request unauthorized response has a 4xx status code
func (o *GetUpdateRequestUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get update request unauthorized response has a 5xx status code
func (o *GetUpdateRequestUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get update request unauthorized response a status code equal to that given
func (o *GetUpdateRequestUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetUpdateRequestUnauthorized) Error() string {
	return fmt.Sprintf("[GET /versions/update-requests/{id}][%d] getUpdateRequestUnauthorized ", 401)
}

func (o *GetUpdateRequestUnauthorized) String() string {
	return fmt.Sprintf("[GET /versions/update-requests/{id}][%d] getUpdateRequestUnauthorized ", 401)
}

func (o *GetUpdateRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUpdateRequestForbidden creates a GetUpdateRequestForbidden with default headers values
func NewGetUpdateRequestForbidden() *GetUpdateRequestForbidden {
	return &GetUpdateRequestForbidden{}
}

/*
GetUpdateRequestForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetUpdateRequestForbidden struct {
}

// IsSuccess returns true when this get update request forbidden response has a 2xx status code
func (o *GetUpdateRequestForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get update request forbidden response has a 3xx status code
func (o *GetUpdateRequestForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get update request forbidden response has a 4xx status code
func (o *GetUpdateRequestForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get update request forbidden response has a 5xx status code
func (o *GetUpdateRequestForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get update request forbidden response a status code equal to that given
func (o *GetUpdateRequestForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetUpdateRequestForbidden) Error() string {
	return fmt.Sprintf("[GET /versions/update-requests/{id}][%d] getUpdateRequestForbidden ", 403)
}

func (o *GetUpdateRequestForbidden) String() string {
	return fmt.Sprintf("[GET /versions/update-requests/{id}][%d] getUpdateRequestForbidden ", 403)
}

func (o *GetUpdateRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUpdateRequestNotFound creates a GetUpdateRequestNotFound with default headers values
func NewGetUpdateRequestNotFound() *GetUpdateRequestNotFound {
	return &GetUpdateRequestNotFound{}
}

/*
GetUpdateRequestNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetUpdateRequestNotFound struct {
}

// IsSuccess returns true when this get update request not found response has a 2xx status code
func (o *GetUpdateRequestNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get update request not found response has a 3xx status code
func (o *GetUpdateRequestNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get update request not found response has a 4xx status code
func (o *GetUpdateRequestNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get update request not found response has a 5xx status code
func (o *GetUpdateRequestNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get update request not found response a status code equal to that given
func (o *GetUpdateRequestNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetUpdateRequestNotFound) Error() string {
	return fmt.Sprintf("[GET /versions/update-requests/{id}][%d] getUpdateRequestNotFound ", 404)
}

func (o *GetUpdateRequestNotFound) String() string {
	return fmt.Sprintf("[GET /versions/update-requests/{id}][%d] getUpdateRequestNotFound ", 404)
}

func (o *GetUpdateRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUpdateRequestConflict creates a GetUpdateRequestConflict with default headers values
func NewGetUpdateRequestConflict() *GetUpdateRequestConflict {
	return &GetUpdateRequestConflict{}
}

/*
GetUpdateRequestConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetUpdateRequestConflict struct {
}

// IsSuccess returns true when this get update request conflict response has a 2xx status code
func (o *GetUpdateRequestConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get update request conflict response has a 3xx status code
func (o *GetUpdateRequestConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get update request conflict response has a 4xx status code
func (o *GetUpdateRequestConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get update request conflict response has a 5xx status code
func (o *GetUpdateRequestConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get update request conflict response a status code equal to that given
func (o *GetUpdateRequestConflict) IsCode(code int) bool {
	return code == 409
}

func (o *GetUpdateRequestConflict) Error() string {
	return fmt.Sprintf("[GET /versions/update-requests/{id}][%d] getUpdateRequestConflict ", 409)
}

func (o *GetUpdateRequestConflict) String() string {
	return fmt.Sprintf("[GET /versions/update-requests/{id}][%d] getUpdateRequestConflict ", 409)
}

func (o *GetUpdateRequestConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
