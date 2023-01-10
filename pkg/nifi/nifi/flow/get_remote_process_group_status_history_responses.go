// Code generated by go-swagger; DO NOT EDIT.

package flow

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// GetRemoteProcessGroupStatusHistoryReader is a Reader for the GetRemoteProcessGroupStatusHistory structure.
type GetRemoteProcessGroupStatusHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRemoteProcessGroupStatusHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRemoteProcessGroupStatusHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRemoteProcessGroupStatusHistoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRemoteProcessGroupStatusHistoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRemoteProcessGroupStatusHistoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRemoteProcessGroupStatusHistoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetRemoteProcessGroupStatusHistoryConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRemoteProcessGroupStatusHistoryOK creates a GetRemoteProcessGroupStatusHistoryOK with default headers values
func NewGetRemoteProcessGroupStatusHistoryOK() *GetRemoteProcessGroupStatusHistoryOK {
	return &GetRemoteProcessGroupStatusHistoryOK{}
}

/*
GetRemoteProcessGroupStatusHistoryOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRemoteProcessGroupStatusHistoryOK struct {
	Payload *models.StatusHistoryEntity
}

// IsSuccess returns true when this get remote process group status history o k response has a 2xx status code
func (o *GetRemoteProcessGroupStatusHistoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get remote process group status history o k response has a 3xx status code
func (o *GetRemoteProcessGroupStatusHistoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get remote process group status history o k response has a 4xx status code
func (o *GetRemoteProcessGroupStatusHistoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get remote process group status history o k response has a 5xx status code
func (o *GetRemoteProcessGroupStatusHistoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get remote process group status history o k response a status code equal to that given
func (o *GetRemoteProcessGroupStatusHistoryOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRemoteProcessGroupStatusHistoryOK) Error() string {
	return fmt.Sprintf("[GET /flow/remote-process-groups/{id}/status/history][%d] getRemoteProcessGroupStatusHistoryOK  %+v", 200, o.Payload)
}

func (o *GetRemoteProcessGroupStatusHistoryOK) String() string {
	return fmt.Sprintf("[GET /flow/remote-process-groups/{id}/status/history][%d] getRemoteProcessGroupStatusHistoryOK  %+v", 200, o.Payload)
}

func (o *GetRemoteProcessGroupStatusHistoryOK) GetPayload() *models.StatusHistoryEntity {
	return o.Payload
}

func (o *GetRemoteProcessGroupStatusHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StatusHistoryEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRemoteProcessGroupStatusHistoryBadRequest creates a GetRemoteProcessGroupStatusHistoryBadRequest with default headers values
func NewGetRemoteProcessGroupStatusHistoryBadRequest() *GetRemoteProcessGroupStatusHistoryBadRequest {
	return &GetRemoteProcessGroupStatusHistoryBadRequest{}
}

/*
GetRemoteProcessGroupStatusHistoryBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetRemoteProcessGroupStatusHistoryBadRequest struct {
}

// IsSuccess returns true when this get remote process group status history bad request response has a 2xx status code
func (o *GetRemoteProcessGroupStatusHistoryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get remote process group status history bad request response has a 3xx status code
func (o *GetRemoteProcessGroupStatusHistoryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get remote process group status history bad request response has a 4xx status code
func (o *GetRemoteProcessGroupStatusHistoryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get remote process group status history bad request response has a 5xx status code
func (o *GetRemoteProcessGroupStatusHistoryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get remote process group status history bad request response a status code equal to that given
func (o *GetRemoteProcessGroupStatusHistoryBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetRemoteProcessGroupStatusHistoryBadRequest) Error() string {
	return fmt.Sprintf("[GET /flow/remote-process-groups/{id}/status/history][%d] getRemoteProcessGroupStatusHistoryBadRequest ", 400)
}

func (o *GetRemoteProcessGroupStatusHistoryBadRequest) String() string {
	return fmt.Sprintf("[GET /flow/remote-process-groups/{id}/status/history][%d] getRemoteProcessGroupStatusHistoryBadRequest ", 400)
}

func (o *GetRemoteProcessGroupStatusHistoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRemoteProcessGroupStatusHistoryUnauthorized creates a GetRemoteProcessGroupStatusHistoryUnauthorized with default headers values
func NewGetRemoteProcessGroupStatusHistoryUnauthorized() *GetRemoteProcessGroupStatusHistoryUnauthorized {
	return &GetRemoteProcessGroupStatusHistoryUnauthorized{}
}

/*
GetRemoteProcessGroupStatusHistoryUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetRemoteProcessGroupStatusHistoryUnauthorized struct {
}

// IsSuccess returns true when this get remote process group status history unauthorized response has a 2xx status code
func (o *GetRemoteProcessGroupStatusHistoryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get remote process group status history unauthorized response has a 3xx status code
func (o *GetRemoteProcessGroupStatusHistoryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get remote process group status history unauthorized response has a 4xx status code
func (o *GetRemoteProcessGroupStatusHistoryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get remote process group status history unauthorized response has a 5xx status code
func (o *GetRemoteProcessGroupStatusHistoryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get remote process group status history unauthorized response a status code equal to that given
func (o *GetRemoteProcessGroupStatusHistoryUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetRemoteProcessGroupStatusHistoryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flow/remote-process-groups/{id}/status/history][%d] getRemoteProcessGroupStatusHistoryUnauthorized ", 401)
}

func (o *GetRemoteProcessGroupStatusHistoryUnauthorized) String() string {
	return fmt.Sprintf("[GET /flow/remote-process-groups/{id}/status/history][%d] getRemoteProcessGroupStatusHistoryUnauthorized ", 401)
}

func (o *GetRemoteProcessGroupStatusHistoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRemoteProcessGroupStatusHistoryForbidden creates a GetRemoteProcessGroupStatusHistoryForbidden with default headers values
func NewGetRemoteProcessGroupStatusHistoryForbidden() *GetRemoteProcessGroupStatusHistoryForbidden {
	return &GetRemoteProcessGroupStatusHistoryForbidden{}
}

/*
GetRemoteProcessGroupStatusHistoryForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetRemoteProcessGroupStatusHistoryForbidden struct {
}

// IsSuccess returns true when this get remote process group status history forbidden response has a 2xx status code
func (o *GetRemoteProcessGroupStatusHistoryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get remote process group status history forbidden response has a 3xx status code
func (o *GetRemoteProcessGroupStatusHistoryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get remote process group status history forbidden response has a 4xx status code
func (o *GetRemoteProcessGroupStatusHistoryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get remote process group status history forbidden response has a 5xx status code
func (o *GetRemoteProcessGroupStatusHistoryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get remote process group status history forbidden response a status code equal to that given
func (o *GetRemoteProcessGroupStatusHistoryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetRemoteProcessGroupStatusHistoryForbidden) Error() string {
	return fmt.Sprintf("[GET /flow/remote-process-groups/{id}/status/history][%d] getRemoteProcessGroupStatusHistoryForbidden ", 403)
}

func (o *GetRemoteProcessGroupStatusHistoryForbidden) String() string {
	return fmt.Sprintf("[GET /flow/remote-process-groups/{id}/status/history][%d] getRemoteProcessGroupStatusHistoryForbidden ", 403)
}

func (o *GetRemoteProcessGroupStatusHistoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRemoteProcessGroupStatusHistoryNotFound creates a GetRemoteProcessGroupStatusHistoryNotFound with default headers values
func NewGetRemoteProcessGroupStatusHistoryNotFound() *GetRemoteProcessGroupStatusHistoryNotFound {
	return &GetRemoteProcessGroupStatusHistoryNotFound{}
}

/*
GetRemoteProcessGroupStatusHistoryNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetRemoteProcessGroupStatusHistoryNotFound struct {
}

// IsSuccess returns true when this get remote process group status history not found response has a 2xx status code
func (o *GetRemoteProcessGroupStatusHistoryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get remote process group status history not found response has a 3xx status code
func (o *GetRemoteProcessGroupStatusHistoryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get remote process group status history not found response has a 4xx status code
func (o *GetRemoteProcessGroupStatusHistoryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get remote process group status history not found response has a 5xx status code
func (o *GetRemoteProcessGroupStatusHistoryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get remote process group status history not found response a status code equal to that given
func (o *GetRemoteProcessGroupStatusHistoryNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetRemoteProcessGroupStatusHistoryNotFound) Error() string {
	return fmt.Sprintf("[GET /flow/remote-process-groups/{id}/status/history][%d] getRemoteProcessGroupStatusHistoryNotFound ", 404)
}

func (o *GetRemoteProcessGroupStatusHistoryNotFound) String() string {
	return fmt.Sprintf("[GET /flow/remote-process-groups/{id}/status/history][%d] getRemoteProcessGroupStatusHistoryNotFound ", 404)
}

func (o *GetRemoteProcessGroupStatusHistoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRemoteProcessGroupStatusHistoryConflict creates a GetRemoteProcessGroupStatusHistoryConflict with default headers values
func NewGetRemoteProcessGroupStatusHistoryConflict() *GetRemoteProcessGroupStatusHistoryConflict {
	return &GetRemoteProcessGroupStatusHistoryConflict{}
}

/*
GetRemoteProcessGroupStatusHistoryConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetRemoteProcessGroupStatusHistoryConflict struct {
}

// IsSuccess returns true when this get remote process group status history conflict response has a 2xx status code
func (o *GetRemoteProcessGroupStatusHistoryConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get remote process group status history conflict response has a 3xx status code
func (o *GetRemoteProcessGroupStatusHistoryConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get remote process group status history conflict response has a 4xx status code
func (o *GetRemoteProcessGroupStatusHistoryConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get remote process group status history conflict response has a 5xx status code
func (o *GetRemoteProcessGroupStatusHistoryConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get remote process group status history conflict response a status code equal to that given
func (o *GetRemoteProcessGroupStatusHistoryConflict) IsCode(code int) bool {
	return code == 409
}

func (o *GetRemoteProcessGroupStatusHistoryConflict) Error() string {
	return fmt.Sprintf("[GET /flow/remote-process-groups/{id}/status/history][%d] getRemoteProcessGroupStatusHistoryConflict ", 409)
}

func (o *GetRemoteProcessGroupStatusHistoryConflict) String() string {
	return fmt.Sprintf("[GET /flow/remote-process-groups/{id}/status/history][%d] getRemoteProcessGroupStatusHistoryConflict ", 409)
}

func (o *GetRemoteProcessGroupStatusHistoryConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
