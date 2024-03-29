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

// GetRemoteProcessGroupsReader is a Reader for the GetRemoteProcessGroups structure.
type GetRemoteProcessGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRemoteProcessGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRemoteProcessGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRemoteProcessGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRemoteProcessGroupsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRemoteProcessGroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRemoteProcessGroupsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetRemoteProcessGroupsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /process-groups/{id}/remote-process-groups] getRemoteProcessGroups", response, response.Code())
	}
}

// NewGetRemoteProcessGroupsOK creates a GetRemoteProcessGroupsOK with default headers values
func NewGetRemoteProcessGroupsOK() *GetRemoteProcessGroupsOK {
	return &GetRemoteProcessGroupsOK{}
}

/*
GetRemoteProcessGroupsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRemoteProcessGroupsOK struct {
	Payload *models.RemoteProcessGroupsEntity
}

// IsSuccess returns true when this get remote process groups o k response has a 2xx status code
func (o *GetRemoteProcessGroupsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get remote process groups o k response has a 3xx status code
func (o *GetRemoteProcessGroupsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get remote process groups o k response has a 4xx status code
func (o *GetRemoteProcessGroupsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get remote process groups o k response has a 5xx status code
func (o *GetRemoteProcessGroupsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get remote process groups o k response a status code equal to that given
func (o *GetRemoteProcessGroupsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get remote process groups o k response
func (o *GetRemoteProcessGroupsOK) Code() int {
	return 200
}

func (o *GetRemoteProcessGroupsOK) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/remote-process-groups][%d] getRemoteProcessGroupsOK  %+v", 200, o.Payload)
}

func (o *GetRemoteProcessGroupsOK) String() string {
	return fmt.Sprintf("[GET /process-groups/{id}/remote-process-groups][%d] getRemoteProcessGroupsOK  %+v", 200, o.Payload)
}

func (o *GetRemoteProcessGroupsOK) GetPayload() *models.RemoteProcessGroupsEntity {
	return o.Payload
}

func (o *GetRemoteProcessGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RemoteProcessGroupsEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRemoteProcessGroupsBadRequest creates a GetRemoteProcessGroupsBadRequest with default headers values
func NewGetRemoteProcessGroupsBadRequest() *GetRemoteProcessGroupsBadRequest {
	return &GetRemoteProcessGroupsBadRequest{}
}

/*
GetRemoteProcessGroupsBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetRemoteProcessGroupsBadRequest struct {
}

// IsSuccess returns true when this get remote process groups bad request response has a 2xx status code
func (o *GetRemoteProcessGroupsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get remote process groups bad request response has a 3xx status code
func (o *GetRemoteProcessGroupsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get remote process groups bad request response has a 4xx status code
func (o *GetRemoteProcessGroupsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get remote process groups bad request response has a 5xx status code
func (o *GetRemoteProcessGroupsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get remote process groups bad request response a status code equal to that given
func (o *GetRemoteProcessGroupsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get remote process groups bad request response
func (o *GetRemoteProcessGroupsBadRequest) Code() int {
	return 400
}

func (o *GetRemoteProcessGroupsBadRequest) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/remote-process-groups][%d] getRemoteProcessGroupsBadRequest ", 400)
}

func (o *GetRemoteProcessGroupsBadRequest) String() string {
	return fmt.Sprintf("[GET /process-groups/{id}/remote-process-groups][%d] getRemoteProcessGroupsBadRequest ", 400)
}

func (o *GetRemoteProcessGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRemoteProcessGroupsUnauthorized creates a GetRemoteProcessGroupsUnauthorized with default headers values
func NewGetRemoteProcessGroupsUnauthorized() *GetRemoteProcessGroupsUnauthorized {
	return &GetRemoteProcessGroupsUnauthorized{}
}

/*
GetRemoteProcessGroupsUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetRemoteProcessGroupsUnauthorized struct {
}

// IsSuccess returns true when this get remote process groups unauthorized response has a 2xx status code
func (o *GetRemoteProcessGroupsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get remote process groups unauthorized response has a 3xx status code
func (o *GetRemoteProcessGroupsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get remote process groups unauthorized response has a 4xx status code
func (o *GetRemoteProcessGroupsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get remote process groups unauthorized response has a 5xx status code
func (o *GetRemoteProcessGroupsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get remote process groups unauthorized response a status code equal to that given
func (o *GetRemoteProcessGroupsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get remote process groups unauthorized response
func (o *GetRemoteProcessGroupsUnauthorized) Code() int {
	return 401
}

func (o *GetRemoteProcessGroupsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/remote-process-groups][%d] getRemoteProcessGroupsUnauthorized ", 401)
}

func (o *GetRemoteProcessGroupsUnauthorized) String() string {
	return fmt.Sprintf("[GET /process-groups/{id}/remote-process-groups][%d] getRemoteProcessGroupsUnauthorized ", 401)
}

func (o *GetRemoteProcessGroupsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRemoteProcessGroupsForbidden creates a GetRemoteProcessGroupsForbidden with default headers values
func NewGetRemoteProcessGroupsForbidden() *GetRemoteProcessGroupsForbidden {
	return &GetRemoteProcessGroupsForbidden{}
}

/*
GetRemoteProcessGroupsForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetRemoteProcessGroupsForbidden struct {
}

// IsSuccess returns true when this get remote process groups forbidden response has a 2xx status code
func (o *GetRemoteProcessGroupsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get remote process groups forbidden response has a 3xx status code
func (o *GetRemoteProcessGroupsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get remote process groups forbidden response has a 4xx status code
func (o *GetRemoteProcessGroupsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get remote process groups forbidden response has a 5xx status code
func (o *GetRemoteProcessGroupsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get remote process groups forbidden response a status code equal to that given
func (o *GetRemoteProcessGroupsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get remote process groups forbidden response
func (o *GetRemoteProcessGroupsForbidden) Code() int {
	return 403
}

func (o *GetRemoteProcessGroupsForbidden) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/remote-process-groups][%d] getRemoteProcessGroupsForbidden ", 403)
}

func (o *GetRemoteProcessGroupsForbidden) String() string {
	return fmt.Sprintf("[GET /process-groups/{id}/remote-process-groups][%d] getRemoteProcessGroupsForbidden ", 403)
}

func (o *GetRemoteProcessGroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRemoteProcessGroupsNotFound creates a GetRemoteProcessGroupsNotFound with default headers values
func NewGetRemoteProcessGroupsNotFound() *GetRemoteProcessGroupsNotFound {
	return &GetRemoteProcessGroupsNotFound{}
}

/*
GetRemoteProcessGroupsNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetRemoteProcessGroupsNotFound struct {
}

// IsSuccess returns true when this get remote process groups not found response has a 2xx status code
func (o *GetRemoteProcessGroupsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get remote process groups not found response has a 3xx status code
func (o *GetRemoteProcessGroupsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get remote process groups not found response has a 4xx status code
func (o *GetRemoteProcessGroupsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get remote process groups not found response has a 5xx status code
func (o *GetRemoteProcessGroupsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get remote process groups not found response a status code equal to that given
func (o *GetRemoteProcessGroupsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get remote process groups not found response
func (o *GetRemoteProcessGroupsNotFound) Code() int {
	return 404
}

func (o *GetRemoteProcessGroupsNotFound) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/remote-process-groups][%d] getRemoteProcessGroupsNotFound ", 404)
}

func (o *GetRemoteProcessGroupsNotFound) String() string {
	return fmt.Sprintf("[GET /process-groups/{id}/remote-process-groups][%d] getRemoteProcessGroupsNotFound ", 404)
}

func (o *GetRemoteProcessGroupsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRemoteProcessGroupsConflict creates a GetRemoteProcessGroupsConflict with default headers values
func NewGetRemoteProcessGroupsConflict() *GetRemoteProcessGroupsConflict {
	return &GetRemoteProcessGroupsConflict{}
}

/*
GetRemoteProcessGroupsConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetRemoteProcessGroupsConflict struct {
}

// IsSuccess returns true when this get remote process groups conflict response has a 2xx status code
func (o *GetRemoteProcessGroupsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get remote process groups conflict response has a 3xx status code
func (o *GetRemoteProcessGroupsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get remote process groups conflict response has a 4xx status code
func (o *GetRemoteProcessGroupsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get remote process groups conflict response has a 5xx status code
func (o *GetRemoteProcessGroupsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get remote process groups conflict response a status code equal to that given
func (o *GetRemoteProcessGroupsConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the get remote process groups conflict response
func (o *GetRemoteProcessGroupsConflict) Code() int {
	return 409
}

func (o *GetRemoteProcessGroupsConflict) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/remote-process-groups][%d] getRemoteProcessGroupsConflict ", 409)
}

func (o *GetRemoteProcessGroupsConflict) String() string {
	return fmt.Sprintf("[GET /process-groups/{id}/remote-process-groups][%d] getRemoteProcessGroupsConflict ", 409)
}

func (o *GetRemoteProcessGroupsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
