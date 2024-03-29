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

// GetProcessorTypesReader is a Reader for the GetProcessorTypes structure.
type GetProcessorTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProcessorTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProcessorTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProcessorTypesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetProcessorTypesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetProcessorTypesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetProcessorTypesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /flow/processor-types] getProcessorTypes", response, response.Code())
	}
}

// NewGetProcessorTypesOK creates a GetProcessorTypesOK with default headers values
func NewGetProcessorTypesOK() *GetProcessorTypesOK {
	return &GetProcessorTypesOK{}
}

/*
GetProcessorTypesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetProcessorTypesOK struct {
	Payload *models.ProcessorTypesEntity
}

// IsSuccess returns true when this get processor types o k response has a 2xx status code
func (o *GetProcessorTypesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get processor types o k response has a 3xx status code
func (o *GetProcessorTypesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get processor types o k response has a 4xx status code
func (o *GetProcessorTypesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get processor types o k response has a 5xx status code
func (o *GetProcessorTypesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get processor types o k response a status code equal to that given
func (o *GetProcessorTypesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get processor types o k response
func (o *GetProcessorTypesOK) Code() int {
	return 200
}

func (o *GetProcessorTypesOK) Error() string {
	return fmt.Sprintf("[GET /flow/processor-types][%d] getProcessorTypesOK  %+v", 200, o.Payload)
}

func (o *GetProcessorTypesOK) String() string {
	return fmt.Sprintf("[GET /flow/processor-types][%d] getProcessorTypesOK  %+v", 200, o.Payload)
}

func (o *GetProcessorTypesOK) GetPayload() *models.ProcessorTypesEntity {
	return o.Payload
}

func (o *GetProcessorTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProcessorTypesEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProcessorTypesBadRequest creates a GetProcessorTypesBadRequest with default headers values
func NewGetProcessorTypesBadRequest() *GetProcessorTypesBadRequest {
	return &GetProcessorTypesBadRequest{}
}

/*
GetProcessorTypesBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetProcessorTypesBadRequest struct {
}

// IsSuccess returns true when this get processor types bad request response has a 2xx status code
func (o *GetProcessorTypesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get processor types bad request response has a 3xx status code
func (o *GetProcessorTypesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get processor types bad request response has a 4xx status code
func (o *GetProcessorTypesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get processor types bad request response has a 5xx status code
func (o *GetProcessorTypesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get processor types bad request response a status code equal to that given
func (o *GetProcessorTypesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get processor types bad request response
func (o *GetProcessorTypesBadRequest) Code() int {
	return 400
}

func (o *GetProcessorTypesBadRequest) Error() string {
	return fmt.Sprintf("[GET /flow/processor-types][%d] getProcessorTypesBadRequest ", 400)
}

func (o *GetProcessorTypesBadRequest) String() string {
	return fmt.Sprintf("[GET /flow/processor-types][%d] getProcessorTypesBadRequest ", 400)
}

func (o *GetProcessorTypesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProcessorTypesUnauthorized creates a GetProcessorTypesUnauthorized with default headers values
func NewGetProcessorTypesUnauthorized() *GetProcessorTypesUnauthorized {
	return &GetProcessorTypesUnauthorized{}
}

/*
GetProcessorTypesUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetProcessorTypesUnauthorized struct {
}

// IsSuccess returns true when this get processor types unauthorized response has a 2xx status code
func (o *GetProcessorTypesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get processor types unauthorized response has a 3xx status code
func (o *GetProcessorTypesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get processor types unauthorized response has a 4xx status code
func (o *GetProcessorTypesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get processor types unauthorized response has a 5xx status code
func (o *GetProcessorTypesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get processor types unauthorized response a status code equal to that given
func (o *GetProcessorTypesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get processor types unauthorized response
func (o *GetProcessorTypesUnauthorized) Code() int {
	return 401
}

func (o *GetProcessorTypesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flow/processor-types][%d] getProcessorTypesUnauthorized ", 401)
}

func (o *GetProcessorTypesUnauthorized) String() string {
	return fmt.Sprintf("[GET /flow/processor-types][%d] getProcessorTypesUnauthorized ", 401)
}

func (o *GetProcessorTypesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProcessorTypesForbidden creates a GetProcessorTypesForbidden with default headers values
func NewGetProcessorTypesForbidden() *GetProcessorTypesForbidden {
	return &GetProcessorTypesForbidden{}
}

/*
GetProcessorTypesForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetProcessorTypesForbidden struct {
}

// IsSuccess returns true when this get processor types forbidden response has a 2xx status code
func (o *GetProcessorTypesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get processor types forbidden response has a 3xx status code
func (o *GetProcessorTypesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get processor types forbidden response has a 4xx status code
func (o *GetProcessorTypesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get processor types forbidden response has a 5xx status code
func (o *GetProcessorTypesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get processor types forbidden response a status code equal to that given
func (o *GetProcessorTypesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get processor types forbidden response
func (o *GetProcessorTypesForbidden) Code() int {
	return 403
}

func (o *GetProcessorTypesForbidden) Error() string {
	return fmt.Sprintf("[GET /flow/processor-types][%d] getProcessorTypesForbidden ", 403)
}

func (o *GetProcessorTypesForbidden) String() string {
	return fmt.Sprintf("[GET /flow/processor-types][%d] getProcessorTypesForbidden ", 403)
}

func (o *GetProcessorTypesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProcessorTypesConflict creates a GetProcessorTypesConflict with default headers values
func NewGetProcessorTypesConflict() *GetProcessorTypesConflict {
	return &GetProcessorTypesConflict{}
}

/*
GetProcessorTypesConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetProcessorTypesConflict struct {
}

// IsSuccess returns true when this get processor types conflict response has a 2xx status code
func (o *GetProcessorTypesConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get processor types conflict response has a 3xx status code
func (o *GetProcessorTypesConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get processor types conflict response has a 4xx status code
func (o *GetProcessorTypesConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get processor types conflict response has a 5xx status code
func (o *GetProcessorTypesConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get processor types conflict response a status code equal to that given
func (o *GetProcessorTypesConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the get processor types conflict response
func (o *GetProcessorTypesConflict) Code() int {
	return 409
}

func (o *GetProcessorTypesConflict) Error() string {
	return fmt.Sprintf("[GET /flow/processor-types][%d] getProcessorTypesConflict ", 409)
}

func (o *GetProcessorTypesConflict) String() string {
	return fmt.Sprintf("[GET /flow/processor-types][%d] getProcessorTypesConflict ", 409)
}

func (o *GetProcessorTypesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
