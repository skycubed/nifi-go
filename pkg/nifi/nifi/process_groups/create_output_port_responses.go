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

// CreateOutputPortReader is a Reader for the CreateOutputPort structure.
type CreateOutputPortReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOutputPortReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateOutputPortCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateOutputPortBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateOutputPortUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateOutputPortForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateOutputPortNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateOutputPortConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /process-groups/{id}/output-ports] createOutputPort", response, response.Code())
	}
}

// NewCreateOutputPortCreated creates a CreateOutputPortCreated with default headers values
func NewCreateOutputPortCreated() *CreateOutputPortCreated {
	return &CreateOutputPortCreated{}
}

/*
CreateOutputPortCreated describes a response with status code 201, with default header values.

successful operation
*/
type CreateOutputPortCreated struct {
	Payload *models.PortEntity
}

// IsSuccess returns true when this create output port created response has a 2xx status code
func (o *CreateOutputPortCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create output port created response has a 3xx status code
func (o *CreateOutputPortCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create output port created response has a 4xx status code
func (o *CreateOutputPortCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create output port created response has a 5xx status code
func (o *CreateOutputPortCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create output port created response a status code equal to that given
func (o *CreateOutputPortCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create output port created response
func (o *CreateOutputPortCreated) Code() int {
	return 201
}

func (o *CreateOutputPortCreated) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/output-ports][%d] createOutputPortCreated  %+v", 201, o.Payload)
}

func (o *CreateOutputPortCreated) String() string {
	return fmt.Sprintf("[POST /process-groups/{id}/output-ports][%d] createOutputPortCreated  %+v", 201, o.Payload)
}

func (o *CreateOutputPortCreated) GetPayload() *models.PortEntity {
	return o.Payload
}

func (o *CreateOutputPortCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOutputPortBadRequest creates a CreateOutputPortBadRequest with default headers values
func NewCreateOutputPortBadRequest() *CreateOutputPortBadRequest {
	return &CreateOutputPortBadRequest{}
}

/*
CreateOutputPortBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type CreateOutputPortBadRequest struct {
}

// IsSuccess returns true when this create output port bad request response has a 2xx status code
func (o *CreateOutputPortBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create output port bad request response has a 3xx status code
func (o *CreateOutputPortBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create output port bad request response has a 4xx status code
func (o *CreateOutputPortBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create output port bad request response has a 5xx status code
func (o *CreateOutputPortBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create output port bad request response a status code equal to that given
func (o *CreateOutputPortBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create output port bad request response
func (o *CreateOutputPortBadRequest) Code() int {
	return 400
}

func (o *CreateOutputPortBadRequest) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/output-ports][%d] createOutputPortBadRequest ", 400)
}

func (o *CreateOutputPortBadRequest) String() string {
	return fmt.Sprintf("[POST /process-groups/{id}/output-ports][%d] createOutputPortBadRequest ", 400)
}

func (o *CreateOutputPortBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateOutputPortUnauthorized creates a CreateOutputPortUnauthorized with default headers values
func NewCreateOutputPortUnauthorized() *CreateOutputPortUnauthorized {
	return &CreateOutputPortUnauthorized{}
}

/*
CreateOutputPortUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type CreateOutputPortUnauthorized struct {
}

// IsSuccess returns true when this create output port unauthorized response has a 2xx status code
func (o *CreateOutputPortUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create output port unauthorized response has a 3xx status code
func (o *CreateOutputPortUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create output port unauthorized response has a 4xx status code
func (o *CreateOutputPortUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create output port unauthorized response has a 5xx status code
func (o *CreateOutputPortUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create output port unauthorized response a status code equal to that given
func (o *CreateOutputPortUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create output port unauthorized response
func (o *CreateOutputPortUnauthorized) Code() int {
	return 401
}

func (o *CreateOutputPortUnauthorized) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/output-ports][%d] createOutputPortUnauthorized ", 401)
}

func (o *CreateOutputPortUnauthorized) String() string {
	return fmt.Sprintf("[POST /process-groups/{id}/output-ports][%d] createOutputPortUnauthorized ", 401)
}

func (o *CreateOutputPortUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateOutputPortForbidden creates a CreateOutputPortForbidden with default headers values
func NewCreateOutputPortForbidden() *CreateOutputPortForbidden {
	return &CreateOutputPortForbidden{}
}

/*
CreateOutputPortForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type CreateOutputPortForbidden struct {
}

// IsSuccess returns true when this create output port forbidden response has a 2xx status code
func (o *CreateOutputPortForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create output port forbidden response has a 3xx status code
func (o *CreateOutputPortForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create output port forbidden response has a 4xx status code
func (o *CreateOutputPortForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create output port forbidden response has a 5xx status code
func (o *CreateOutputPortForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create output port forbidden response a status code equal to that given
func (o *CreateOutputPortForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create output port forbidden response
func (o *CreateOutputPortForbidden) Code() int {
	return 403
}

func (o *CreateOutputPortForbidden) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/output-ports][%d] createOutputPortForbidden ", 403)
}

func (o *CreateOutputPortForbidden) String() string {
	return fmt.Sprintf("[POST /process-groups/{id}/output-ports][%d] createOutputPortForbidden ", 403)
}

func (o *CreateOutputPortForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateOutputPortNotFound creates a CreateOutputPortNotFound with default headers values
func NewCreateOutputPortNotFound() *CreateOutputPortNotFound {
	return &CreateOutputPortNotFound{}
}

/*
CreateOutputPortNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type CreateOutputPortNotFound struct {
}

// IsSuccess returns true when this create output port not found response has a 2xx status code
func (o *CreateOutputPortNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create output port not found response has a 3xx status code
func (o *CreateOutputPortNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create output port not found response has a 4xx status code
func (o *CreateOutputPortNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create output port not found response has a 5xx status code
func (o *CreateOutputPortNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create output port not found response a status code equal to that given
func (o *CreateOutputPortNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create output port not found response
func (o *CreateOutputPortNotFound) Code() int {
	return 404
}

func (o *CreateOutputPortNotFound) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/output-ports][%d] createOutputPortNotFound ", 404)
}

func (o *CreateOutputPortNotFound) String() string {
	return fmt.Sprintf("[POST /process-groups/{id}/output-ports][%d] createOutputPortNotFound ", 404)
}

func (o *CreateOutputPortNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateOutputPortConflict creates a CreateOutputPortConflict with default headers values
func NewCreateOutputPortConflict() *CreateOutputPortConflict {
	return &CreateOutputPortConflict{}
}

/*
CreateOutputPortConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type CreateOutputPortConflict struct {
}

// IsSuccess returns true when this create output port conflict response has a 2xx status code
func (o *CreateOutputPortConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create output port conflict response has a 3xx status code
func (o *CreateOutputPortConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create output port conflict response has a 4xx status code
func (o *CreateOutputPortConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create output port conflict response has a 5xx status code
func (o *CreateOutputPortConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create output port conflict response a status code equal to that given
func (o *CreateOutputPortConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create output port conflict response
func (o *CreateOutputPortConflict) Code() int {
	return 409
}

func (o *CreateOutputPortConflict) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/output-ports][%d] createOutputPortConflict ", 409)
}

func (o *CreateOutputPortConflict) String() string {
	return fmt.Sprintf("[POST /process-groups/{id}/output-ports][%d] createOutputPortConflict ", 409)
}

func (o *CreateOutputPortConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
