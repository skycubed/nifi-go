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

// UpdateRemoteProcessGroupOutputPortReader is a Reader for the UpdateRemoteProcessGroupOutputPort structure.
type UpdateRemoteProcessGroupOutputPortReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRemoteProcessGroupOutputPortReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRemoteProcessGroupOutputPortOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateRemoteProcessGroupOutputPortBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateRemoteProcessGroupOutputPortUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRemoteProcessGroupOutputPortForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateRemoteProcessGroupOutputPortNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateRemoteProcessGroupOutputPortConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRemoteProcessGroupOutputPortOK creates a UpdateRemoteProcessGroupOutputPortOK with default headers values
func NewUpdateRemoteProcessGroupOutputPortOK() *UpdateRemoteProcessGroupOutputPortOK {
	return &UpdateRemoteProcessGroupOutputPortOK{}
}

/*
UpdateRemoteProcessGroupOutputPortOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateRemoteProcessGroupOutputPortOK struct {
	Payload *models.RemoteProcessGroupPortEntity
}

// IsSuccess returns true when this update remote process group output port o k response has a 2xx status code
func (o *UpdateRemoteProcessGroupOutputPortOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update remote process group output port o k response has a 3xx status code
func (o *UpdateRemoteProcessGroupOutputPortOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update remote process group output port o k response has a 4xx status code
func (o *UpdateRemoteProcessGroupOutputPortOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update remote process group output port o k response has a 5xx status code
func (o *UpdateRemoteProcessGroupOutputPortOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update remote process group output port o k response a status code equal to that given
func (o *UpdateRemoteProcessGroupOutputPortOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateRemoteProcessGroupOutputPortOK) Error() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/output-ports/{port-id}][%d] updateRemoteProcessGroupOutputPortOK  %+v", 200, o.Payload)
}

func (o *UpdateRemoteProcessGroupOutputPortOK) String() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/output-ports/{port-id}][%d] updateRemoteProcessGroupOutputPortOK  %+v", 200, o.Payload)
}

func (o *UpdateRemoteProcessGroupOutputPortOK) GetPayload() *models.RemoteProcessGroupPortEntity {
	return o.Payload
}

func (o *UpdateRemoteProcessGroupOutputPortOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RemoteProcessGroupPortEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRemoteProcessGroupOutputPortBadRequest creates a UpdateRemoteProcessGroupOutputPortBadRequest with default headers values
func NewUpdateRemoteProcessGroupOutputPortBadRequest() *UpdateRemoteProcessGroupOutputPortBadRequest {
	return &UpdateRemoteProcessGroupOutputPortBadRequest{}
}

/*
UpdateRemoteProcessGroupOutputPortBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type UpdateRemoteProcessGroupOutputPortBadRequest struct {
}

// IsSuccess returns true when this update remote process group output port bad request response has a 2xx status code
func (o *UpdateRemoteProcessGroupOutputPortBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update remote process group output port bad request response has a 3xx status code
func (o *UpdateRemoteProcessGroupOutputPortBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update remote process group output port bad request response has a 4xx status code
func (o *UpdateRemoteProcessGroupOutputPortBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update remote process group output port bad request response has a 5xx status code
func (o *UpdateRemoteProcessGroupOutputPortBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update remote process group output port bad request response a status code equal to that given
func (o *UpdateRemoteProcessGroupOutputPortBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UpdateRemoteProcessGroupOutputPortBadRequest) Error() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/output-ports/{port-id}][%d] updateRemoteProcessGroupOutputPortBadRequest ", 400)
}

func (o *UpdateRemoteProcessGroupOutputPortBadRequest) String() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/output-ports/{port-id}][%d] updateRemoteProcessGroupOutputPortBadRequest ", 400)
}

func (o *UpdateRemoteProcessGroupOutputPortBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRemoteProcessGroupOutputPortUnauthorized creates a UpdateRemoteProcessGroupOutputPortUnauthorized with default headers values
func NewUpdateRemoteProcessGroupOutputPortUnauthorized() *UpdateRemoteProcessGroupOutputPortUnauthorized {
	return &UpdateRemoteProcessGroupOutputPortUnauthorized{}
}

/*
UpdateRemoteProcessGroupOutputPortUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type UpdateRemoteProcessGroupOutputPortUnauthorized struct {
}

// IsSuccess returns true when this update remote process group output port unauthorized response has a 2xx status code
func (o *UpdateRemoteProcessGroupOutputPortUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update remote process group output port unauthorized response has a 3xx status code
func (o *UpdateRemoteProcessGroupOutputPortUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update remote process group output port unauthorized response has a 4xx status code
func (o *UpdateRemoteProcessGroupOutputPortUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update remote process group output port unauthorized response has a 5xx status code
func (o *UpdateRemoteProcessGroupOutputPortUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update remote process group output port unauthorized response a status code equal to that given
func (o *UpdateRemoteProcessGroupOutputPortUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateRemoteProcessGroupOutputPortUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/output-ports/{port-id}][%d] updateRemoteProcessGroupOutputPortUnauthorized ", 401)
}

func (o *UpdateRemoteProcessGroupOutputPortUnauthorized) String() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/output-ports/{port-id}][%d] updateRemoteProcessGroupOutputPortUnauthorized ", 401)
}

func (o *UpdateRemoteProcessGroupOutputPortUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRemoteProcessGroupOutputPortForbidden creates a UpdateRemoteProcessGroupOutputPortForbidden with default headers values
func NewUpdateRemoteProcessGroupOutputPortForbidden() *UpdateRemoteProcessGroupOutputPortForbidden {
	return &UpdateRemoteProcessGroupOutputPortForbidden{}
}

/*
UpdateRemoteProcessGroupOutputPortForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type UpdateRemoteProcessGroupOutputPortForbidden struct {
}

// IsSuccess returns true when this update remote process group output port forbidden response has a 2xx status code
func (o *UpdateRemoteProcessGroupOutputPortForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update remote process group output port forbidden response has a 3xx status code
func (o *UpdateRemoteProcessGroupOutputPortForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update remote process group output port forbidden response has a 4xx status code
func (o *UpdateRemoteProcessGroupOutputPortForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update remote process group output port forbidden response has a 5xx status code
func (o *UpdateRemoteProcessGroupOutputPortForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update remote process group output port forbidden response a status code equal to that given
func (o *UpdateRemoteProcessGroupOutputPortForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateRemoteProcessGroupOutputPortForbidden) Error() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/output-ports/{port-id}][%d] updateRemoteProcessGroupOutputPortForbidden ", 403)
}

func (o *UpdateRemoteProcessGroupOutputPortForbidden) String() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/output-ports/{port-id}][%d] updateRemoteProcessGroupOutputPortForbidden ", 403)
}

func (o *UpdateRemoteProcessGroupOutputPortForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRemoteProcessGroupOutputPortNotFound creates a UpdateRemoteProcessGroupOutputPortNotFound with default headers values
func NewUpdateRemoteProcessGroupOutputPortNotFound() *UpdateRemoteProcessGroupOutputPortNotFound {
	return &UpdateRemoteProcessGroupOutputPortNotFound{}
}

/*
UpdateRemoteProcessGroupOutputPortNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type UpdateRemoteProcessGroupOutputPortNotFound struct {
}

// IsSuccess returns true when this update remote process group output port not found response has a 2xx status code
func (o *UpdateRemoteProcessGroupOutputPortNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update remote process group output port not found response has a 3xx status code
func (o *UpdateRemoteProcessGroupOutputPortNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update remote process group output port not found response has a 4xx status code
func (o *UpdateRemoteProcessGroupOutputPortNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update remote process group output port not found response has a 5xx status code
func (o *UpdateRemoteProcessGroupOutputPortNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update remote process group output port not found response a status code equal to that given
func (o *UpdateRemoteProcessGroupOutputPortNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateRemoteProcessGroupOutputPortNotFound) Error() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/output-ports/{port-id}][%d] updateRemoteProcessGroupOutputPortNotFound ", 404)
}

func (o *UpdateRemoteProcessGroupOutputPortNotFound) String() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/output-ports/{port-id}][%d] updateRemoteProcessGroupOutputPortNotFound ", 404)
}

func (o *UpdateRemoteProcessGroupOutputPortNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRemoteProcessGroupOutputPortConflict creates a UpdateRemoteProcessGroupOutputPortConflict with default headers values
func NewUpdateRemoteProcessGroupOutputPortConflict() *UpdateRemoteProcessGroupOutputPortConflict {
	return &UpdateRemoteProcessGroupOutputPortConflict{}
}

/*
UpdateRemoteProcessGroupOutputPortConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type UpdateRemoteProcessGroupOutputPortConflict struct {
}

// IsSuccess returns true when this update remote process group output port conflict response has a 2xx status code
func (o *UpdateRemoteProcessGroupOutputPortConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update remote process group output port conflict response has a 3xx status code
func (o *UpdateRemoteProcessGroupOutputPortConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update remote process group output port conflict response has a 4xx status code
func (o *UpdateRemoteProcessGroupOutputPortConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update remote process group output port conflict response has a 5xx status code
func (o *UpdateRemoteProcessGroupOutputPortConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update remote process group output port conflict response a status code equal to that given
func (o *UpdateRemoteProcessGroupOutputPortConflict) IsCode(code int) bool {
	return code == 409
}

func (o *UpdateRemoteProcessGroupOutputPortConflict) Error() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/output-ports/{port-id}][%d] updateRemoteProcessGroupOutputPortConflict ", 409)
}

func (o *UpdateRemoteProcessGroupOutputPortConflict) String() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}/output-ports/{port-id}][%d] updateRemoteProcessGroupOutputPortConflict ", 409)
}

func (o *UpdateRemoteProcessGroupOutputPortConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
