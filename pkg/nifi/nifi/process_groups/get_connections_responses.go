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

// GetConnectionsReader is a Reader for the GetConnections structure.
type GetConnectionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConnectionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConnectionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConnectionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConnectionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConnectionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConnectionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetConnectionsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConnectionsOK creates a GetConnectionsOK with default headers values
func NewGetConnectionsOK() *GetConnectionsOK {
	return &GetConnectionsOK{}
}

/*
GetConnectionsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetConnectionsOK struct {
	Payload *models.ConnectionsEntity
}

// IsSuccess returns true when this get connections o k response has a 2xx status code
func (o *GetConnectionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get connections o k response has a 3xx status code
func (o *GetConnectionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get connections o k response has a 4xx status code
func (o *GetConnectionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get connections o k response has a 5xx status code
func (o *GetConnectionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get connections o k response a status code equal to that given
func (o *GetConnectionsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetConnectionsOK) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/connections][%d] getConnectionsOK  %+v", 200, o.Payload)
}

func (o *GetConnectionsOK) String() string {
	return fmt.Sprintf("[GET /process-groups/{id}/connections][%d] getConnectionsOK  %+v", 200, o.Payload)
}

func (o *GetConnectionsOK) GetPayload() *models.ConnectionsEntity {
	return o.Payload
}

func (o *GetConnectionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConnectionsEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConnectionsBadRequest creates a GetConnectionsBadRequest with default headers values
func NewGetConnectionsBadRequest() *GetConnectionsBadRequest {
	return &GetConnectionsBadRequest{}
}

/*
GetConnectionsBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetConnectionsBadRequest struct {
}

// IsSuccess returns true when this get connections bad request response has a 2xx status code
func (o *GetConnectionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get connections bad request response has a 3xx status code
func (o *GetConnectionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get connections bad request response has a 4xx status code
func (o *GetConnectionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get connections bad request response has a 5xx status code
func (o *GetConnectionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get connections bad request response a status code equal to that given
func (o *GetConnectionsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetConnectionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/connections][%d] getConnectionsBadRequest ", 400)
}

func (o *GetConnectionsBadRequest) String() string {
	return fmt.Sprintf("[GET /process-groups/{id}/connections][%d] getConnectionsBadRequest ", 400)
}

func (o *GetConnectionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetConnectionsUnauthorized creates a GetConnectionsUnauthorized with default headers values
func NewGetConnectionsUnauthorized() *GetConnectionsUnauthorized {
	return &GetConnectionsUnauthorized{}
}

/*
GetConnectionsUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetConnectionsUnauthorized struct {
}

// IsSuccess returns true when this get connections unauthorized response has a 2xx status code
func (o *GetConnectionsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get connections unauthorized response has a 3xx status code
func (o *GetConnectionsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get connections unauthorized response has a 4xx status code
func (o *GetConnectionsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get connections unauthorized response has a 5xx status code
func (o *GetConnectionsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get connections unauthorized response a status code equal to that given
func (o *GetConnectionsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetConnectionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/connections][%d] getConnectionsUnauthorized ", 401)
}

func (o *GetConnectionsUnauthorized) String() string {
	return fmt.Sprintf("[GET /process-groups/{id}/connections][%d] getConnectionsUnauthorized ", 401)
}

func (o *GetConnectionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetConnectionsForbidden creates a GetConnectionsForbidden with default headers values
func NewGetConnectionsForbidden() *GetConnectionsForbidden {
	return &GetConnectionsForbidden{}
}

/*
GetConnectionsForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetConnectionsForbidden struct {
}

// IsSuccess returns true when this get connections forbidden response has a 2xx status code
func (o *GetConnectionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get connections forbidden response has a 3xx status code
func (o *GetConnectionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get connections forbidden response has a 4xx status code
func (o *GetConnectionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get connections forbidden response has a 5xx status code
func (o *GetConnectionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get connections forbidden response a status code equal to that given
func (o *GetConnectionsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetConnectionsForbidden) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/connections][%d] getConnectionsForbidden ", 403)
}

func (o *GetConnectionsForbidden) String() string {
	return fmt.Sprintf("[GET /process-groups/{id}/connections][%d] getConnectionsForbidden ", 403)
}

func (o *GetConnectionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetConnectionsNotFound creates a GetConnectionsNotFound with default headers values
func NewGetConnectionsNotFound() *GetConnectionsNotFound {
	return &GetConnectionsNotFound{}
}

/*
GetConnectionsNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetConnectionsNotFound struct {
}

// IsSuccess returns true when this get connections not found response has a 2xx status code
func (o *GetConnectionsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get connections not found response has a 3xx status code
func (o *GetConnectionsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get connections not found response has a 4xx status code
func (o *GetConnectionsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get connections not found response has a 5xx status code
func (o *GetConnectionsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get connections not found response a status code equal to that given
func (o *GetConnectionsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetConnectionsNotFound) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/connections][%d] getConnectionsNotFound ", 404)
}

func (o *GetConnectionsNotFound) String() string {
	return fmt.Sprintf("[GET /process-groups/{id}/connections][%d] getConnectionsNotFound ", 404)
}

func (o *GetConnectionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetConnectionsConflict creates a GetConnectionsConflict with default headers values
func NewGetConnectionsConflict() *GetConnectionsConflict {
	return &GetConnectionsConflict{}
}

/*
GetConnectionsConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetConnectionsConflict struct {
}

// IsSuccess returns true when this get connections conflict response has a 2xx status code
func (o *GetConnectionsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get connections conflict response has a 3xx status code
func (o *GetConnectionsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get connections conflict response has a 4xx status code
func (o *GetConnectionsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get connections conflict response has a 5xx status code
func (o *GetConnectionsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get connections conflict response a status code equal to that given
func (o *GetConnectionsConflict) IsCode(code int) bool {
	return code == 409
}

func (o *GetConnectionsConflict) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/connections][%d] getConnectionsConflict ", 409)
}

func (o *GetConnectionsConflict) String() string {
	return fmt.Sprintf("[GET /process-groups/{id}/connections][%d] getConnectionsConflict ", 409)
}

func (o *GetConnectionsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
