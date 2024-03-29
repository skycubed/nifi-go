// Code generated by go-swagger; DO NOT EDIT.

package output_ports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// GetOutputPortReader is a Reader for the GetOutputPort structure.
type GetOutputPortReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutputPortReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutputPortOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutputPortBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutputPortUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutputPortForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutputPortNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetOutputPortConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /output-ports/{id}] getOutputPort", response, response.Code())
	}
}

// NewGetOutputPortOK creates a GetOutputPortOK with default headers values
func NewGetOutputPortOK() *GetOutputPortOK {
	return &GetOutputPortOK{}
}

/*
GetOutputPortOK describes a response with status code 200, with default header values.

successful operation
*/
type GetOutputPortOK struct {
	Payload *models.PortEntity
}

// IsSuccess returns true when this get output port o k response has a 2xx status code
func (o *GetOutputPortOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get output port o k response has a 3xx status code
func (o *GetOutputPortOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get output port o k response has a 4xx status code
func (o *GetOutputPortOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get output port o k response has a 5xx status code
func (o *GetOutputPortOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get output port o k response a status code equal to that given
func (o *GetOutputPortOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get output port o k response
func (o *GetOutputPortOK) Code() int {
	return 200
}

func (o *GetOutputPortOK) Error() string {
	return fmt.Sprintf("[GET /output-ports/{id}][%d] getOutputPortOK  %+v", 200, o.Payload)
}

func (o *GetOutputPortOK) String() string {
	return fmt.Sprintf("[GET /output-ports/{id}][%d] getOutputPortOK  %+v", 200, o.Payload)
}

func (o *GetOutputPortOK) GetPayload() *models.PortEntity {
	return o.Payload
}

func (o *GetOutputPortOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutputPortBadRequest creates a GetOutputPortBadRequest with default headers values
func NewGetOutputPortBadRequest() *GetOutputPortBadRequest {
	return &GetOutputPortBadRequest{}
}

/*
GetOutputPortBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetOutputPortBadRequest struct {
}

// IsSuccess returns true when this get output port bad request response has a 2xx status code
func (o *GetOutputPortBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get output port bad request response has a 3xx status code
func (o *GetOutputPortBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get output port bad request response has a 4xx status code
func (o *GetOutputPortBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get output port bad request response has a 5xx status code
func (o *GetOutputPortBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get output port bad request response a status code equal to that given
func (o *GetOutputPortBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get output port bad request response
func (o *GetOutputPortBadRequest) Code() int {
	return 400
}

func (o *GetOutputPortBadRequest) Error() string {
	return fmt.Sprintf("[GET /output-ports/{id}][%d] getOutputPortBadRequest ", 400)
}

func (o *GetOutputPortBadRequest) String() string {
	return fmt.Sprintf("[GET /output-ports/{id}][%d] getOutputPortBadRequest ", 400)
}

func (o *GetOutputPortBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOutputPortUnauthorized creates a GetOutputPortUnauthorized with default headers values
func NewGetOutputPortUnauthorized() *GetOutputPortUnauthorized {
	return &GetOutputPortUnauthorized{}
}

/*
GetOutputPortUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetOutputPortUnauthorized struct {
}

// IsSuccess returns true when this get output port unauthorized response has a 2xx status code
func (o *GetOutputPortUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get output port unauthorized response has a 3xx status code
func (o *GetOutputPortUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get output port unauthorized response has a 4xx status code
func (o *GetOutputPortUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get output port unauthorized response has a 5xx status code
func (o *GetOutputPortUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get output port unauthorized response a status code equal to that given
func (o *GetOutputPortUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get output port unauthorized response
func (o *GetOutputPortUnauthorized) Code() int {
	return 401
}

func (o *GetOutputPortUnauthorized) Error() string {
	return fmt.Sprintf("[GET /output-ports/{id}][%d] getOutputPortUnauthorized ", 401)
}

func (o *GetOutputPortUnauthorized) String() string {
	return fmt.Sprintf("[GET /output-ports/{id}][%d] getOutputPortUnauthorized ", 401)
}

func (o *GetOutputPortUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOutputPortForbidden creates a GetOutputPortForbidden with default headers values
func NewGetOutputPortForbidden() *GetOutputPortForbidden {
	return &GetOutputPortForbidden{}
}

/*
GetOutputPortForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetOutputPortForbidden struct {
}

// IsSuccess returns true when this get output port forbidden response has a 2xx status code
func (o *GetOutputPortForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get output port forbidden response has a 3xx status code
func (o *GetOutputPortForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get output port forbidden response has a 4xx status code
func (o *GetOutputPortForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get output port forbidden response has a 5xx status code
func (o *GetOutputPortForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get output port forbidden response a status code equal to that given
func (o *GetOutputPortForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get output port forbidden response
func (o *GetOutputPortForbidden) Code() int {
	return 403
}

func (o *GetOutputPortForbidden) Error() string {
	return fmt.Sprintf("[GET /output-ports/{id}][%d] getOutputPortForbidden ", 403)
}

func (o *GetOutputPortForbidden) String() string {
	return fmt.Sprintf("[GET /output-ports/{id}][%d] getOutputPortForbidden ", 403)
}

func (o *GetOutputPortForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOutputPortNotFound creates a GetOutputPortNotFound with default headers values
func NewGetOutputPortNotFound() *GetOutputPortNotFound {
	return &GetOutputPortNotFound{}
}

/*
GetOutputPortNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetOutputPortNotFound struct {
}

// IsSuccess returns true when this get output port not found response has a 2xx status code
func (o *GetOutputPortNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get output port not found response has a 3xx status code
func (o *GetOutputPortNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get output port not found response has a 4xx status code
func (o *GetOutputPortNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get output port not found response has a 5xx status code
func (o *GetOutputPortNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get output port not found response a status code equal to that given
func (o *GetOutputPortNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get output port not found response
func (o *GetOutputPortNotFound) Code() int {
	return 404
}

func (o *GetOutputPortNotFound) Error() string {
	return fmt.Sprintf("[GET /output-ports/{id}][%d] getOutputPortNotFound ", 404)
}

func (o *GetOutputPortNotFound) String() string {
	return fmt.Sprintf("[GET /output-ports/{id}][%d] getOutputPortNotFound ", 404)
}

func (o *GetOutputPortNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOutputPortConflict creates a GetOutputPortConflict with default headers values
func NewGetOutputPortConflict() *GetOutputPortConflict {
	return &GetOutputPortConflict{}
}

/*
GetOutputPortConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetOutputPortConflict struct {
}

// IsSuccess returns true when this get output port conflict response has a 2xx status code
func (o *GetOutputPortConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get output port conflict response has a 3xx status code
func (o *GetOutputPortConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get output port conflict response has a 4xx status code
func (o *GetOutputPortConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get output port conflict response has a 5xx status code
func (o *GetOutputPortConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get output port conflict response a status code equal to that given
func (o *GetOutputPortConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the get output port conflict response
func (o *GetOutputPortConflict) Code() int {
	return 409
}

func (o *GetOutputPortConflict) Error() string {
	return fmt.Sprintf("[GET /output-ports/{id}][%d] getOutputPortConflict ", 409)
}

func (o *GetOutputPortConflict) String() string {
	return fmt.Sprintf("[GET /output-ports/{id}][%d] getOutputPortConflict ", 409)
}

func (o *GetOutputPortConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
