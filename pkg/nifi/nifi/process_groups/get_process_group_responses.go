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

// GetProcessGroupReader is a Reader for the GetProcessGroup structure.
type GetProcessGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProcessGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProcessGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProcessGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetProcessGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetProcessGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProcessGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetProcessGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /process-groups/{id}] getProcessGroup", response, response.Code())
	}
}

// NewGetProcessGroupOK creates a GetProcessGroupOK with default headers values
func NewGetProcessGroupOK() *GetProcessGroupOK {
	return &GetProcessGroupOK{}
}

/*
GetProcessGroupOK describes a response with status code 200, with default header values.

successful operation
*/
type GetProcessGroupOK struct {
	Payload *models.ProcessGroupEntity
}

// IsSuccess returns true when this get process group o k response has a 2xx status code
func (o *GetProcessGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get process group o k response has a 3xx status code
func (o *GetProcessGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get process group o k response has a 4xx status code
func (o *GetProcessGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get process group o k response has a 5xx status code
func (o *GetProcessGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get process group o k response a status code equal to that given
func (o *GetProcessGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get process group o k response
func (o *GetProcessGroupOK) Code() int {
	return 200
}

func (o *GetProcessGroupOK) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}][%d] getProcessGroupOK  %+v", 200, o.Payload)
}

func (o *GetProcessGroupOK) String() string {
	return fmt.Sprintf("[GET /process-groups/{id}][%d] getProcessGroupOK  %+v", 200, o.Payload)
}

func (o *GetProcessGroupOK) GetPayload() *models.ProcessGroupEntity {
	return o.Payload
}

func (o *GetProcessGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProcessGroupEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProcessGroupBadRequest creates a GetProcessGroupBadRequest with default headers values
func NewGetProcessGroupBadRequest() *GetProcessGroupBadRequest {
	return &GetProcessGroupBadRequest{}
}

/*
GetProcessGroupBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetProcessGroupBadRequest struct {
}

// IsSuccess returns true when this get process group bad request response has a 2xx status code
func (o *GetProcessGroupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get process group bad request response has a 3xx status code
func (o *GetProcessGroupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get process group bad request response has a 4xx status code
func (o *GetProcessGroupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get process group bad request response has a 5xx status code
func (o *GetProcessGroupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get process group bad request response a status code equal to that given
func (o *GetProcessGroupBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get process group bad request response
func (o *GetProcessGroupBadRequest) Code() int {
	return 400
}

func (o *GetProcessGroupBadRequest) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}][%d] getProcessGroupBadRequest ", 400)
}

func (o *GetProcessGroupBadRequest) String() string {
	return fmt.Sprintf("[GET /process-groups/{id}][%d] getProcessGroupBadRequest ", 400)
}

func (o *GetProcessGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProcessGroupUnauthorized creates a GetProcessGroupUnauthorized with default headers values
func NewGetProcessGroupUnauthorized() *GetProcessGroupUnauthorized {
	return &GetProcessGroupUnauthorized{}
}

/*
GetProcessGroupUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetProcessGroupUnauthorized struct {
}

// IsSuccess returns true when this get process group unauthorized response has a 2xx status code
func (o *GetProcessGroupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get process group unauthorized response has a 3xx status code
func (o *GetProcessGroupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get process group unauthorized response has a 4xx status code
func (o *GetProcessGroupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get process group unauthorized response has a 5xx status code
func (o *GetProcessGroupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get process group unauthorized response a status code equal to that given
func (o *GetProcessGroupUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get process group unauthorized response
func (o *GetProcessGroupUnauthorized) Code() int {
	return 401
}

func (o *GetProcessGroupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}][%d] getProcessGroupUnauthorized ", 401)
}

func (o *GetProcessGroupUnauthorized) String() string {
	return fmt.Sprintf("[GET /process-groups/{id}][%d] getProcessGroupUnauthorized ", 401)
}

func (o *GetProcessGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProcessGroupForbidden creates a GetProcessGroupForbidden with default headers values
func NewGetProcessGroupForbidden() *GetProcessGroupForbidden {
	return &GetProcessGroupForbidden{}
}

/*
GetProcessGroupForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetProcessGroupForbidden struct {
}

// IsSuccess returns true when this get process group forbidden response has a 2xx status code
func (o *GetProcessGroupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get process group forbidden response has a 3xx status code
func (o *GetProcessGroupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get process group forbidden response has a 4xx status code
func (o *GetProcessGroupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get process group forbidden response has a 5xx status code
func (o *GetProcessGroupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get process group forbidden response a status code equal to that given
func (o *GetProcessGroupForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get process group forbidden response
func (o *GetProcessGroupForbidden) Code() int {
	return 403
}

func (o *GetProcessGroupForbidden) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}][%d] getProcessGroupForbidden ", 403)
}

func (o *GetProcessGroupForbidden) String() string {
	return fmt.Sprintf("[GET /process-groups/{id}][%d] getProcessGroupForbidden ", 403)
}

func (o *GetProcessGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProcessGroupNotFound creates a GetProcessGroupNotFound with default headers values
func NewGetProcessGroupNotFound() *GetProcessGroupNotFound {
	return &GetProcessGroupNotFound{}
}

/*
GetProcessGroupNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetProcessGroupNotFound struct {
}

// IsSuccess returns true when this get process group not found response has a 2xx status code
func (o *GetProcessGroupNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get process group not found response has a 3xx status code
func (o *GetProcessGroupNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get process group not found response has a 4xx status code
func (o *GetProcessGroupNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get process group not found response has a 5xx status code
func (o *GetProcessGroupNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get process group not found response a status code equal to that given
func (o *GetProcessGroupNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get process group not found response
func (o *GetProcessGroupNotFound) Code() int {
	return 404
}

func (o *GetProcessGroupNotFound) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}][%d] getProcessGroupNotFound ", 404)
}

func (o *GetProcessGroupNotFound) String() string {
	return fmt.Sprintf("[GET /process-groups/{id}][%d] getProcessGroupNotFound ", 404)
}

func (o *GetProcessGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProcessGroupConflict creates a GetProcessGroupConflict with default headers values
func NewGetProcessGroupConflict() *GetProcessGroupConflict {
	return &GetProcessGroupConflict{}
}

/*
GetProcessGroupConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetProcessGroupConflict struct {
}

// IsSuccess returns true when this get process group conflict response has a 2xx status code
func (o *GetProcessGroupConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get process group conflict response has a 3xx status code
func (o *GetProcessGroupConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get process group conflict response has a 4xx status code
func (o *GetProcessGroupConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get process group conflict response has a 5xx status code
func (o *GetProcessGroupConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get process group conflict response a status code equal to that given
func (o *GetProcessGroupConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the get process group conflict response
func (o *GetProcessGroupConflict) Code() int {
	return 409
}

func (o *GetProcessGroupConflict) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}][%d] getProcessGroupConflict ", 409)
}

func (o *GetProcessGroupConflict) String() string {
	return fmt.Sprintf("[GET /process-groups/{id}][%d] getProcessGroupConflict ", 409)
}

func (o *GetProcessGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
