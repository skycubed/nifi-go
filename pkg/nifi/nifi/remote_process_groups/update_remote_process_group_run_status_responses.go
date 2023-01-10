// Code generated by go-swagger; DO NOT EDIT.

package remote_process_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// UpdateRemoteProcessGroupRunStatusReader is a Reader for the UpdateRemoteProcessGroupRunStatus structure.
type UpdateRemoteProcessGroupRunStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRemoteProcessGroupRunStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRemoteProcessGroupRunStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateRemoteProcessGroupRunStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateRemoteProcessGroupRunStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRemoteProcessGroupRunStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateRemoteProcessGroupRunStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateRemoteProcessGroupRunStatusConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRemoteProcessGroupRunStatusOK creates a UpdateRemoteProcessGroupRunStatusOK with default headers values
func NewUpdateRemoteProcessGroupRunStatusOK() *UpdateRemoteProcessGroupRunStatusOK {
	return &UpdateRemoteProcessGroupRunStatusOK{}
}

/*
UpdateRemoteProcessGroupRunStatusOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateRemoteProcessGroupRunStatusOK struct {
	Payload *models.RemoteProcessGroupEntity
}

// IsSuccess returns true when this update remote process group run status o k response has a 2xx status code
func (o *UpdateRemoteProcessGroupRunStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update remote process group run status o k response has a 3xx status code
func (o *UpdateRemoteProcessGroupRunStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update remote process group run status o k response has a 4xx status code
func (o *UpdateRemoteProcessGroupRunStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update remote process group run status o k response has a 5xx status code
func (o *UpdateRemoteProcessGroupRunStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update remote process group run status o k response a status code equal to that given
func (o *UpdateRemoteProcessGroupRunStatusOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateRemoteProcessGroupRunStatusOK) Error() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/run-status][%d] updateRemoteProcessGroupRunStatusOK  %+v", 200, o.Payload)
}

func (o *UpdateRemoteProcessGroupRunStatusOK) String() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/run-status][%d] updateRemoteProcessGroupRunStatusOK  %+v", 200, o.Payload)
}

func (o *UpdateRemoteProcessGroupRunStatusOK) GetPayload() *models.RemoteProcessGroupEntity {
	return o.Payload
}

func (o *UpdateRemoteProcessGroupRunStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RemoteProcessGroupEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRemoteProcessGroupRunStatusBadRequest creates a UpdateRemoteProcessGroupRunStatusBadRequest with default headers values
func NewUpdateRemoteProcessGroupRunStatusBadRequest() *UpdateRemoteProcessGroupRunStatusBadRequest {
	return &UpdateRemoteProcessGroupRunStatusBadRequest{}
}

/*
UpdateRemoteProcessGroupRunStatusBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type UpdateRemoteProcessGroupRunStatusBadRequest struct {
}

// IsSuccess returns true when this update remote process group run status bad request response has a 2xx status code
func (o *UpdateRemoteProcessGroupRunStatusBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update remote process group run status bad request response has a 3xx status code
func (o *UpdateRemoteProcessGroupRunStatusBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update remote process group run status bad request response has a 4xx status code
func (o *UpdateRemoteProcessGroupRunStatusBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update remote process group run status bad request response has a 5xx status code
func (o *UpdateRemoteProcessGroupRunStatusBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update remote process group run status bad request response a status code equal to that given
func (o *UpdateRemoteProcessGroupRunStatusBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UpdateRemoteProcessGroupRunStatusBadRequest) Error() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/run-status][%d] updateRemoteProcessGroupRunStatusBadRequest ", 400)
}

func (o *UpdateRemoteProcessGroupRunStatusBadRequest) String() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/run-status][%d] updateRemoteProcessGroupRunStatusBadRequest ", 400)
}

func (o *UpdateRemoteProcessGroupRunStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRemoteProcessGroupRunStatusUnauthorized creates a UpdateRemoteProcessGroupRunStatusUnauthorized with default headers values
func NewUpdateRemoteProcessGroupRunStatusUnauthorized() *UpdateRemoteProcessGroupRunStatusUnauthorized {
	return &UpdateRemoteProcessGroupRunStatusUnauthorized{}
}

/*
UpdateRemoteProcessGroupRunStatusUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type UpdateRemoteProcessGroupRunStatusUnauthorized struct {
}

// IsSuccess returns true when this update remote process group run status unauthorized response has a 2xx status code
func (o *UpdateRemoteProcessGroupRunStatusUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update remote process group run status unauthorized response has a 3xx status code
func (o *UpdateRemoteProcessGroupRunStatusUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update remote process group run status unauthorized response has a 4xx status code
func (o *UpdateRemoteProcessGroupRunStatusUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update remote process group run status unauthorized response has a 5xx status code
func (o *UpdateRemoteProcessGroupRunStatusUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update remote process group run status unauthorized response a status code equal to that given
func (o *UpdateRemoteProcessGroupRunStatusUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateRemoteProcessGroupRunStatusUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/run-status][%d] updateRemoteProcessGroupRunStatusUnauthorized ", 401)
}

func (o *UpdateRemoteProcessGroupRunStatusUnauthorized) String() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/run-status][%d] updateRemoteProcessGroupRunStatusUnauthorized ", 401)
}

func (o *UpdateRemoteProcessGroupRunStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRemoteProcessGroupRunStatusForbidden creates a UpdateRemoteProcessGroupRunStatusForbidden with default headers values
func NewUpdateRemoteProcessGroupRunStatusForbidden() *UpdateRemoteProcessGroupRunStatusForbidden {
	return &UpdateRemoteProcessGroupRunStatusForbidden{}
}

/*
UpdateRemoteProcessGroupRunStatusForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type UpdateRemoteProcessGroupRunStatusForbidden struct {
}

// IsSuccess returns true when this update remote process group run status forbidden response has a 2xx status code
func (o *UpdateRemoteProcessGroupRunStatusForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update remote process group run status forbidden response has a 3xx status code
func (o *UpdateRemoteProcessGroupRunStatusForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update remote process group run status forbidden response has a 4xx status code
func (o *UpdateRemoteProcessGroupRunStatusForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update remote process group run status forbidden response has a 5xx status code
func (o *UpdateRemoteProcessGroupRunStatusForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update remote process group run status forbidden response a status code equal to that given
func (o *UpdateRemoteProcessGroupRunStatusForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateRemoteProcessGroupRunStatusForbidden) Error() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/run-status][%d] updateRemoteProcessGroupRunStatusForbidden ", 403)
}

func (o *UpdateRemoteProcessGroupRunStatusForbidden) String() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/run-status][%d] updateRemoteProcessGroupRunStatusForbidden ", 403)
}

func (o *UpdateRemoteProcessGroupRunStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRemoteProcessGroupRunStatusNotFound creates a UpdateRemoteProcessGroupRunStatusNotFound with default headers values
func NewUpdateRemoteProcessGroupRunStatusNotFound() *UpdateRemoteProcessGroupRunStatusNotFound {
	return &UpdateRemoteProcessGroupRunStatusNotFound{}
}

/*
UpdateRemoteProcessGroupRunStatusNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type UpdateRemoteProcessGroupRunStatusNotFound struct {
}

// IsSuccess returns true when this update remote process group run status not found response has a 2xx status code
func (o *UpdateRemoteProcessGroupRunStatusNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update remote process group run status not found response has a 3xx status code
func (o *UpdateRemoteProcessGroupRunStatusNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update remote process group run status not found response has a 4xx status code
func (o *UpdateRemoteProcessGroupRunStatusNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update remote process group run status not found response has a 5xx status code
func (o *UpdateRemoteProcessGroupRunStatusNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update remote process group run status not found response a status code equal to that given
func (o *UpdateRemoteProcessGroupRunStatusNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateRemoteProcessGroupRunStatusNotFound) Error() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/run-status][%d] updateRemoteProcessGroupRunStatusNotFound ", 404)
}

func (o *UpdateRemoteProcessGroupRunStatusNotFound) String() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/run-status][%d] updateRemoteProcessGroupRunStatusNotFound ", 404)
}

func (o *UpdateRemoteProcessGroupRunStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRemoteProcessGroupRunStatusConflict creates a UpdateRemoteProcessGroupRunStatusConflict with default headers values
func NewUpdateRemoteProcessGroupRunStatusConflict() *UpdateRemoteProcessGroupRunStatusConflict {
	return &UpdateRemoteProcessGroupRunStatusConflict{}
}

/*
UpdateRemoteProcessGroupRunStatusConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type UpdateRemoteProcessGroupRunStatusConflict struct {
}

// IsSuccess returns true when this update remote process group run status conflict response has a 2xx status code
func (o *UpdateRemoteProcessGroupRunStatusConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update remote process group run status conflict response has a 3xx status code
func (o *UpdateRemoteProcessGroupRunStatusConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update remote process group run status conflict response has a 4xx status code
func (o *UpdateRemoteProcessGroupRunStatusConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update remote process group run status conflict response has a 5xx status code
func (o *UpdateRemoteProcessGroupRunStatusConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update remote process group run status conflict response a status code equal to that given
func (o *UpdateRemoteProcessGroupRunStatusConflict) IsCode(code int) bool {
	return code == 409
}

func (o *UpdateRemoteProcessGroupRunStatusConflict) Error() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/run-status][%d] updateRemoteProcessGroupRunStatusConflict ", 409)
}

func (o *UpdateRemoteProcessGroupRunStatusConflict) String() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/run-status][%d] updateRemoteProcessGroupRunStatusConflict ", 409)
}

func (o *UpdateRemoteProcessGroupRunStatusConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
