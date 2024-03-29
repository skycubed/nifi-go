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

// GetRevertRequestReader is a Reader for the GetRevertRequest structure.
type GetRevertRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRevertRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRevertRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRevertRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRevertRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRevertRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRevertRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetRevertRequestConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /versions/revert-requests/{id}] getRevertRequest", response, response.Code())
	}
}

// NewGetRevertRequestOK creates a GetRevertRequestOK with default headers values
func NewGetRevertRequestOK() *GetRevertRequestOK {
	return &GetRevertRequestOK{}
}

/*
GetRevertRequestOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRevertRequestOK struct {
	Payload *models.VersionedFlowUpdateRequestEntity
}

// IsSuccess returns true when this get revert request o k response has a 2xx status code
func (o *GetRevertRequestOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get revert request o k response has a 3xx status code
func (o *GetRevertRequestOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get revert request o k response has a 4xx status code
func (o *GetRevertRequestOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get revert request o k response has a 5xx status code
func (o *GetRevertRequestOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get revert request o k response a status code equal to that given
func (o *GetRevertRequestOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get revert request o k response
func (o *GetRevertRequestOK) Code() int {
	return 200
}

func (o *GetRevertRequestOK) Error() string {
	return fmt.Sprintf("[GET /versions/revert-requests/{id}][%d] getRevertRequestOK  %+v", 200, o.Payload)
}

func (o *GetRevertRequestOK) String() string {
	return fmt.Sprintf("[GET /versions/revert-requests/{id}][%d] getRevertRequestOK  %+v", 200, o.Payload)
}

func (o *GetRevertRequestOK) GetPayload() *models.VersionedFlowUpdateRequestEntity {
	return o.Payload
}

func (o *GetRevertRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VersionedFlowUpdateRequestEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRevertRequestBadRequest creates a GetRevertRequestBadRequest with default headers values
func NewGetRevertRequestBadRequest() *GetRevertRequestBadRequest {
	return &GetRevertRequestBadRequest{}
}

/*
GetRevertRequestBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetRevertRequestBadRequest struct {
}

// IsSuccess returns true when this get revert request bad request response has a 2xx status code
func (o *GetRevertRequestBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get revert request bad request response has a 3xx status code
func (o *GetRevertRequestBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get revert request bad request response has a 4xx status code
func (o *GetRevertRequestBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get revert request bad request response has a 5xx status code
func (o *GetRevertRequestBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get revert request bad request response a status code equal to that given
func (o *GetRevertRequestBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get revert request bad request response
func (o *GetRevertRequestBadRequest) Code() int {
	return 400
}

func (o *GetRevertRequestBadRequest) Error() string {
	return fmt.Sprintf("[GET /versions/revert-requests/{id}][%d] getRevertRequestBadRequest ", 400)
}

func (o *GetRevertRequestBadRequest) String() string {
	return fmt.Sprintf("[GET /versions/revert-requests/{id}][%d] getRevertRequestBadRequest ", 400)
}

func (o *GetRevertRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRevertRequestUnauthorized creates a GetRevertRequestUnauthorized with default headers values
func NewGetRevertRequestUnauthorized() *GetRevertRequestUnauthorized {
	return &GetRevertRequestUnauthorized{}
}

/*
GetRevertRequestUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetRevertRequestUnauthorized struct {
}

// IsSuccess returns true when this get revert request unauthorized response has a 2xx status code
func (o *GetRevertRequestUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get revert request unauthorized response has a 3xx status code
func (o *GetRevertRequestUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get revert request unauthorized response has a 4xx status code
func (o *GetRevertRequestUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get revert request unauthorized response has a 5xx status code
func (o *GetRevertRequestUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get revert request unauthorized response a status code equal to that given
func (o *GetRevertRequestUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get revert request unauthorized response
func (o *GetRevertRequestUnauthorized) Code() int {
	return 401
}

func (o *GetRevertRequestUnauthorized) Error() string {
	return fmt.Sprintf("[GET /versions/revert-requests/{id}][%d] getRevertRequestUnauthorized ", 401)
}

func (o *GetRevertRequestUnauthorized) String() string {
	return fmt.Sprintf("[GET /versions/revert-requests/{id}][%d] getRevertRequestUnauthorized ", 401)
}

func (o *GetRevertRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRevertRequestForbidden creates a GetRevertRequestForbidden with default headers values
func NewGetRevertRequestForbidden() *GetRevertRequestForbidden {
	return &GetRevertRequestForbidden{}
}

/*
GetRevertRequestForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetRevertRequestForbidden struct {
}

// IsSuccess returns true when this get revert request forbidden response has a 2xx status code
func (o *GetRevertRequestForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get revert request forbidden response has a 3xx status code
func (o *GetRevertRequestForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get revert request forbidden response has a 4xx status code
func (o *GetRevertRequestForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get revert request forbidden response has a 5xx status code
func (o *GetRevertRequestForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get revert request forbidden response a status code equal to that given
func (o *GetRevertRequestForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get revert request forbidden response
func (o *GetRevertRequestForbidden) Code() int {
	return 403
}

func (o *GetRevertRequestForbidden) Error() string {
	return fmt.Sprintf("[GET /versions/revert-requests/{id}][%d] getRevertRequestForbidden ", 403)
}

func (o *GetRevertRequestForbidden) String() string {
	return fmt.Sprintf("[GET /versions/revert-requests/{id}][%d] getRevertRequestForbidden ", 403)
}

func (o *GetRevertRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRevertRequestNotFound creates a GetRevertRequestNotFound with default headers values
func NewGetRevertRequestNotFound() *GetRevertRequestNotFound {
	return &GetRevertRequestNotFound{}
}

/*
GetRevertRequestNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetRevertRequestNotFound struct {
}

// IsSuccess returns true when this get revert request not found response has a 2xx status code
func (o *GetRevertRequestNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get revert request not found response has a 3xx status code
func (o *GetRevertRequestNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get revert request not found response has a 4xx status code
func (o *GetRevertRequestNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get revert request not found response has a 5xx status code
func (o *GetRevertRequestNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get revert request not found response a status code equal to that given
func (o *GetRevertRequestNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get revert request not found response
func (o *GetRevertRequestNotFound) Code() int {
	return 404
}

func (o *GetRevertRequestNotFound) Error() string {
	return fmt.Sprintf("[GET /versions/revert-requests/{id}][%d] getRevertRequestNotFound ", 404)
}

func (o *GetRevertRequestNotFound) String() string {
	return fmt.Sprintf("[GET /versions/revert-requests/{id}][%d] getRevertRequestNotFound ", 404)
}

func (o *GetRevertRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRevertRequestConflict creates a GetRevertRequestConflict with default headers values
func NewGetRevertRequestConflict() *GetRevertRequestConflict {
	return &GetRevertRequestConflict{}
}

/*
GetRevertRequestConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetRevertRequestConflict struct {
}

// IsSuccess returns true when this get revert request conflict response has a 2xx status code
func (o *GetRevertRequestConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get revert request conflict response has a 3xx status code
func (o *GetRevertRequestConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get revert request conflict response has a 4xx status code
func (o *GetRevertRequestConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get revert request conflict response has a 5xx status code
func (o *GetRevertRequestConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get revert request conflict response a status code equal to that given
func (o *GetRevertRequestConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the get revert request conflict response
func (o *GetRevertRequestConflict) Code() int {
	return 409
}

func (o *GetRevertRequestConflict) Error() string {
	return fmt.Sprintf("[GET /versions/revert-requests/{id}][%d] getRevertRequestConflict ", 409)
}

func (o *GetRevertRequestConflict) String() string {
	return fmt.Sprintf("[GET /versions/revert-requests/{id}][%d] getRevertRequestConflict ", 409)
}

func (o *GetRevertRequestConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
