// Code generated by go-swagger; DO NOT EDIT.

package flow

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GenerateClientIDReader is a Reader for the GenerateClientID structure.
type GenerateClientIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GenerateClientIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGenerateClientIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGenerateClientIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGenerateClientIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGenerateClientIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGenerateClientIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGenerateClientIDOK creates a GenerateClientIDOK with default headers values
func NewGenerateClientIDOK() *GenerateClientIDOK {
	return &GenerateClientIDOK{}
}

/*
GenerateClientIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GenerateClientIDOK struct {
	Payload string
}

// IsSuccess returns true when this generate client Id o k response has a 2xx status code
func (o *GenerateClientIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this generate client Id o k response has a 3xx status code
func (o *GenerateClientIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate client Id o k response has a 4xx status code
func (o *GenerateClientIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this generate client Id o k response has a 5xx status code
func (o *GenerateClientIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this generate client Id o k response a status code equal to that given
func (o *GenerateClientIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GenerateClientIDOK) Error() string {
	return fmt.Sprintf("[GET /flow/client-id][%d] generateClientIdOK  %+v", 200, o.Payload)
}

func (o *GenerateClientIDOK) String() string {
	return fmt.Sprintf("[GET /flow/client-id][%d] generateClientIdOK  %+v", 200, o.Payload)
}

func (o *GenerateClientIDOK) GetPayload() string {
	return o.Payload
}

func (o *GenerateClientIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateClientIDBadRequest creates a GenerateClientIDBadRequest with default headers values
func NewGenerateClientIDBadRequest() *GenerateClientIDBadRequest {
	return &GenerateClientIDBadRequest{}
}

/*
GenerateClientIDBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GenerateClientIDBadRequest struct {
}

// IsSuccess returns true when this generate client Id bad request response has a 2xx status code
func (o *GenerateClientIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this generate client Id bad request response has a 3xx status code
func (o *GenerateClientIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate client Id bad request response has a 4xx status code
func (o *GenerateClientIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this generate client Id bad request response has a 5xx status code
func (o *GenerateClientIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this generate client Id bad request response a status code equal to that given
func (o *GenerateClientIDBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GenerateClientIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /flow/client-id][%d] generateClientIdBadRequest ", 400)
}

func (o *GenerateClientIDBadRequest) String() string {
	return fmt.Sprintf("[GET /flow/client-id][%d] generateClientIdBadRequest ", 400)
}

func (o *GenerateClientIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGenerateClientIDUnauthorized creates a GenerateClientIDUnauthorized with default headers values
func NewGenerateClientIDUnauthorized() *GenerateClientIDUnauthorized {
	return &GenerateClientIDUnauthorized{}
}

/*
GenerateClientIDUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GenerateClientIDUnauthorized struct {
}

// IsSuccess returns true when this generate client Id unauthorized response has a 2xx status code
func (o *GenerateClientIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this generate client Id unauthorized response has a 3xx status code
func (o *GenerateClientIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate client Id unauthorized response has a 4xx status code
func (o *GenerateClientIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this generate client Id unauthorized response has a 5xx status code
func (o *GenerateClientIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this generate client Id unauthorized response a status code equal to that given
func (o *GenerateClientIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GenerateClientIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flow/client-id][%d] generateClientIdUnauthorized ", 401)
}

func (o *GenerateClientIDUnauthorized) String() string {
	return fmt.Sprintf("[GET /flow/client-id][%d] generateClientIdUnauthorized ", 401)
}

func (o *GenerateClientIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGenerateClientIDForbidden creates a GenerateClientIDForbidden with default headers values
func NewGenerateClientIDForbidden() *GenerateClientIDForbidden {
	return &GenerateClientIDForbidden{}
}

/*
GenerateClientIDForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GenerateClientIDForbidden struct {
}

// IsSuccess returns true when this generate client Id forbidden response has a 2xx status code
func (o *GenerateClientIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this generate client Id forbidden response has a 3xx status code
func (o *GenerateClientIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate client Id forbidden response has a 4xx status code
func (o *GenerateClientIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this generate client Id forbidden response has a 5xx status code
func (o *GenerateClientIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this generate client Id forbidden response a status code equal to that given
func (o *GenerateClientIDForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GenerateClientIDForbidden) Error() string {
	return fmt.Sprintf("[GET /flow/client-id][%d] generateClientIdForbidden ", 403)
}

func (o *GenerateClientIDForbidden) String() string {
	return fmt.Sprintf("[GET /flow/client-id][%d] generateClientIdForbidden ", 403)
}

func (o *GenerateClientIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGenerateClientIDConflict creates a GenerateClientIDConflict with default headers values
func NewGenerateClientIDConflict() *GenerateClientIDConflict {
	return &GenerateClientIDConflict{}
}

/*
GenerateClientIDConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GenerateClientIDConflict struct {
}

// IsSuccess returns true when this generate client Id conflict response has a 2xx status code
func (o *GenerateClientIDConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this generate client Id conflict response has a 3xx status code
func (o *GenerateClientIDConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate client Id conflict response has a 4xx status code
func (o *GenerateClientIDConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this generate client Id conflict response has a 5xx status code
func (o *GenerateClientIDConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this generate client Id conflict response a status code equal to that given
func (o *GenerateClientIDConflict) IsCode(code int) bool {
	return code == 409
}

func (o *GenerateClientIDConflict) Error() string {
	return fmt.Sprintf("[GET /flow/client-id][%d] generateClientIdConflict ", 409)
}

func (o *GenerateClientIDConflict) String() string {
	return fmt.Sprintf("[GET /flow/client-id][%d] generateClientIdConflict ", 409)
}

func (o *GenerateClientIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
