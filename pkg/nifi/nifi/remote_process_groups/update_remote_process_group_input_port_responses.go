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

// UpdateRemoteProcessGroupInputPortReader is a Reader for the UpdateRemoteProcessGroupInputPort structure.
type UpdateRemoteProcessGroupInputPortReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRemoteProcessGroupInputPortReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRemoteProcessGroupInputPortOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateRemoteProcessGroupInputPortBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateRemoteProcessGroupInputPortUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRemoteProcessGroupInputPortForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateRemoteProcessGroupInputPortNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateRemoteProcessGroupInputPortConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /remote-process-groups/{id}/input-ports/{port-id}] updateRemoteProcessGroupInputPort", response, response.Code())
	}
}

// NewUpdateRemoteProcessGroupInputPortOK creates a UpdateRemoteProcessGroupInputPortOK with default headers values
func NewUpdateRemoteProcessGroupInputPortOK() *UpdateRemoteProcessGroupInputPortOK {
	return &UpdateRemoteProcessGroupInputPortOK{}
}

/*
UpdateRemoteProcessGroupInputPortOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateRemoteProcessGroupInputPortOK struct {
	Payload *models.RemoteProcessGroupPortEntity
}

// IsSuccess returns true when this update remote process group input port o k response has a 2xx status code
func (o *UpdateRemoteProcessGroupInputPortOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update remote process group input port o k response has a 3xx status code
func (o *UpdateRemoteProcessGroupInputPortOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update remote process group input port o k response has a 4xx status code
func (o *UpdateRemoteProcessGroupInputPortOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update remote process group input port o k response has a 5xx status code
func (o *UpdateRemoteProcessGroupInputPortOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update remote process group input port o k response a status code equal to that given
func (o *UpdateRemoteProcessGroupInputPortOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update remote process group input port o k response
func (o *UpdateRemoteProcessGroupInputPortOK) Code() int {
	return 200
}

func (o *UpdateRemoteProcessGroupInputPortOK) Error() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/input-ports/{port-id}][%d] updateRemoteProcessGroupInputPortOK  %+v", 200, o.Payload)
}

func (o *UpdateRemoteProcessGroupInputPortOK) String() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/input-ports/{port-id}][%d] updateRemoteProcessGroupInputPortOK  %+v", 200, o.Payload)
}

func (o *UpdateRemoteProcessGroupInputPortOK) GetPayload() *models.RemoteProcessGroupPortEntity {
	return o.Payload
}

func (o *UpdateRemoteProcessGroupInputPortOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RemoteProcessGroupPortEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRemoteProcessGroupInputPortBadRequest creates a UpdateRemoteProcessGroupInputPortBadRequest with default headers values
func NewUpdateRemoteProcessGroupInputPortBadRequest() *UpdateRemoteProcessGroupInputPortBadRequest {
	return &UpdateRemoteProcessGroupInputPortBadRequest{}
}

/*
UpdateRemoteProcessGroupInputPortBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type UpdateRemoteProcessGroupInputPortBadRequest struct {
}

// IsSuccess returns true when this update remote process group input port bad request response has a 2xx status code
func (o *UpdateRemoteProcessGroupInputPortBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update remote process group input port bad request response has a 3xx status code
func (o *UpdateRemoteProcessGroupInputPortBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update remote process group input port bad request response has a 4xx status code
func (o *UpdateRemoteProcessGroupInputPortBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update remote process group input port bad request response has a 5xx status code
func (o *UpdateRemoteProcessGroupInputPortBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update remote process group input port bad request response a status code equal to that given
func (o *UpdateRemoteProcessGroupInputPortBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update remote process group input port bad request response
func (o *UpdateRemoteProcessGroupInputPortBadRequest) Code() int {
	return 400
}

func (o *UpdateRemoteProcessGroupInputPortBadRequest) Error() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/input-ports/{port-id}][%d] updateRemoteProcessGroupInputPortBadRequest ", 400)
}

func (o *UpdateRemoteProcessGroupInputPortBadRequest) String() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/input-ports/{port-id}][%d] updateRemoteProcessGroupInputPortBadRequest ", 400)
}

func (o *UpdateRemoteProcessGroupInputPortBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRemoteProcessGroupInputPortUnauthorized creates a UpdateRemoteProcessGroupInputPortUnauthorized with default headers values
func NewUpdateRemoteProcessGroupInputPortUnauthorized() *UpdateRemoteProcessGroupInputPortUnauthorized {
	return &UpdateRemoteProcessGroupInputPortUnauthorized{}
}

/*
UpdateRemoteProcessGroupInputPortUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type UpdateRemoteProcessGroupInputPortUnauthorized struct {
}

// IsSuccess returns true when this update remote process group input port unauthorized response has a 2xx status code
func (o *UpdateRemoteProcessGroupInputPortUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update remote process group input port unauthorized response has a 3xx status code
func (o *UpdateRemoteProcessGroupInputPortUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update remote process group input port unauthorized response has a 4xx status code
func (o *UpdateRemoteProcessGroupInputPortUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update remote process group input port unauthorized response has a 5xx status code
func (o *UpdateRemoteProcessGroupInputPortUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update remote process group input port unauthorized response a status code equal to that given
func (o *UpdateRemoteProcessGroupInputPortUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update remote process group input port unauthorized response
func (o *UpdateRemoteProcessGroupInputPortUnauthorized) Code() int {
	return 401
}

func (o *UpdateRemoteProcessGroupInputPortUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/input-ports/{port-id}][%d] updateRemoteProcessGroupInputPortUnauthorized ", 401)
}

func (o *UpdateRemoteProcessGroupInputPortUnauthorized) String() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/input-ports/{port-id}][%d] updateRemoteProcessGroupInputPortUnauthorized ", 401)
}

func (o *UpdateRemoteProcessGroupInputPortUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRemoteProcessGroupInputPortForbidden creates a UpdateRemoteProcessGroupInputPortForbidden with default headers values
func NewUpdateRemoteProcessGroupInputPortForbidden() *UpdateRemoteProcessGroupInputPortForbidden {
	return &UpdateRemoteProcessGroupInputPortForbidden{}
}

/*
UpdateRemoteProcessGroupInputPortForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type UpdateRemoteProcessGroupInputPortForbidden struct {
}

// IsSuccess returns true when this update remote process group input port forbidden response has a 2xx status code
func (o *UpdateRemoteProcessGroupInputPortForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update remote process group input port forbidden response has a 3xx status code
func (o *UpdateRemoteProcessGroupInputPortForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update remote process group input port forbidden response has a 4xx status code
func (o *UpdateRemoteProcessGroupInputPortForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update remote process group input port forbidden response has a 5xx status code
func (o *UpdateRemoteProcessGroupInputPortForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update remote process group input port forbidden response a status code equal to that given
func (o *UpdateRemoteProcessGroupInputPortForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update remote process group input port forbidden response
func (o *UpdateRemoteProcessGroupInputPortForbidden) Code() int {
	return 403
}

func (o *UpdateRemoteProcessGroupInputPortForbidden) Error() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/input-ports/{port-id}][%d] updateRemoteProcessGroupInputPortForbidden ", 403)
}

func (o *UpdateRemoteProcessGroupInputPortForbidden) String() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/input-ports/{port-id}][%d] updateRemoteProcessGroupInputPortForbidden ", 403)
}

func (o *UpdateRemoteProcessGroupInputPortForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRemoteProcessGroupInputPortNotFound creates a UpdateRemoteProcessGroupInputPortNotFound with default headers values
func NewUpdateRemoteProcessGroupInputPortNotFound() *UpdateRemoteProcessGroupInputPortNotFound {
	return &UpdateRemoteProcessGroupInputPortNotFound{}
}

/*
UpdateRemoteProcessGroupInputPortNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type UpdateRemoteProcessGroupInputPortNotFound struct {
}

// IsSuccess returns true when this update remote process group input port not found response has a 2xx status code
func (o *UpdateRemoteProcessGroupInputPortNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update remote process group input port not found response has a 3xx status code
func (o *UpdateRemoteProcessGroupInputPortNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update remote process group input port not found response has a 4xx status code
func (o *UpdateRemoteProcessGroupInputPortNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update remote process group input port not found response has a 5xx status code
func (o *UpdateRemoteProcessGroupInputPortNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update remote process group input port not found response a status code equal to that given
func (o *UpdateRemoteProcessGroupInputPortNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update remote process group input port not found response
func (o *UpdateRemoteProcessGroupInputPortNotFound) Code() int {
	return 404
}

func (o *UpdateRemoteProcessGroupInputPortNotFound) Error() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/input-ports/{port-id}][%d] updateRemoteProcessGroupInputPortNotFound ", 404)
}

func (o *UpdateRemoteProcessGroupInputPortNotFound) String() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/input-ports/{port-id}][%d] updateRemoteProcessGroupInputPortNotFound ", 404)
}

func (o *UpdateRemoteProcessGroupInputPortNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRemoteProcessGroupInputPortConflict creates a UpdateRemoteProcessGroupInputPortConflict with default headers values
func NewUpdateRemoteProcessGroupInputPortConflict() *UpdateRemoteProcessGroupInputPortConflict {
	return &UpdateRemoteProcessGroupInputPortConflict{}
}

/*
UpdateRemoteProcessGroupInputPortConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type UpdateRemoteProcessGroupInputPortConflict struct {
}

// IsSuccess returns true when this update remote process group input port conflict response has a 2xx status code
func (o *UpdateRemoteProcessGroupInputPortConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update remote process group input port conflict response has a 3xx status code
func (o *UpdateRemoteProcessGroupInputPortConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update remote process group input port conflict response has a 4xx status code
func (o *UpdateRemoteProcessGroupInputPortConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update remote process group input port conflict response has a 5xx status code
func (o *UpdateRemoteProcessGroupInputPortConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update remote process group input port conflict response a status code equal to that given
func (o *UpdateRemoteProcessGroupInputPortConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the update remote process group input port conflict response
func (o *UpdateRemoteProcessGroupInputPortConflict) Code() int {
	return 409
}

func (o *UpdateRemoteProcessGroupInputPortConflict) Error() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/input-ports/{port-id}][%d] updateRemoteProcessGroupInputPortConflict ", 409)
}

func (o *UpdateRemoteProcessGroupInputPortConflict) String() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/input-ports/{port-id}][%d] updateRemoteProcessGroupInputPortConflict ", 409)
}

func (o *UpdateRemoteProcessGroupInputPortConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
