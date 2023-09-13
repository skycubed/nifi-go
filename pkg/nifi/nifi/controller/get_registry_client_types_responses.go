// Code generated by go-swagger; DO NOT EDIT.

package controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// GetRegistryClientTypesReader is a Reader for the GetRegistryClientTypes structure.
type GetRegistryClientTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRegistryClientTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRegistryClientTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRegistryClientTypesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRegistryClientTypesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRegistryClientTypesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetRegistryClientTypesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /controller/registry-types] getRegistryClientTypes", response, response.Code())
	}
}

// NewGetRegistryClientTypesOK creates a GetRegistryClientTypesOK with default headers values
func NewGetRegistryClientTypesOK() *GetRegistryClientTypesOK {
	return &GetRegistryClientTypesOK{}
}

/*
GetRegistryClientTypesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRegistryClientTypesOK struct {
	Payload *models.FlowRegistryClientTypesEntity
}

// IsSuccess returns true when this get registry client types o k response has a 2xx status code
func (o *GetRegistryClientTypesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get registry client types o k response has a 3xx status code
func (o *GetRegistryClientTypesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get registry client types o k response has a 4xx status code
func (o *GetRegistryClientTypesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get registry client types o k response has a 5xx status code
func (o *GetRegistryClientTypesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get registry client types o k response a status code equal to that given
func (o *GetRegistryClientTypesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get registry client types o k response
func (o *GetRegistryClientTypesOK) Code() int {
	return 200
}

func (o *GetRegistryClientTypesOK) Error() string {
	return fmt.Sprintf("[GET /controller/registry-types][%d] getRegistryClientTypesOK  %+v", 200, o.Payload)
}

func (o *GetRegistryClientTypesOK) String() string {
	return fmt.Sprintf("[GET /controller/registry-types][%d] getRegistryClientTypesOK  %+v", 200, o.Payload)
}

func (o *GetRegistryClientTypesOK) GetPayload() *models.FlowRegistryClientTypesEntity {
	return o.Payload
}

func (o *GetRegistryClientTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FlowRegistryClientTypesEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegistryClientTypesBadRequest creates a GetRegistryClientTypesBadRequest with default headers values
func NewGetRegistryClientTypesBadRequest() *GetRegistryClientTypesBadRequest {
	return &GetRegistryClientTypesBadRequest{}
}

/*
GetRegistryClientTypesBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetRegistryClientTypesBadRequest struct {
}

// IsSuccess returns true when this get registry client types bad request response has a 2xx status code
func (o *GetRegistryClientTypesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get registry client types bad request response has a 3xx status code
func (o *GetRegistryClientTypesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get registry client types bad request response has a 4xx status code
func (o *GetRegistryClientTypesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get registry client types bad request response has a 5xx status code
func (o *GetRegistryClientTypesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get registry client types bad request response a status code equal to that given
func (o *GetRegistryClientTypesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get registry client types bad request response
func (o *GetRegistryClientTypesBadRequest) Code() int {
	return 400
}

func (o *GetRegistryClientTypesBadRequest) Error() string {
	return fmt.Sprintf("[GET /controller/registry-types][%d] getRegistryClientTypesBadRequest ", 400)
}

func (o *GetRegistryClientTypesBadRequest) String() string {
	return fmt.Sprintf("[GET /controller/registry-types][%d] getRegistryClientTypesBadRequest ", 400)
}

func (o *GetRegistryClientTypesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRegistryClientTypesUnauthorized creates a GetRegistryClientTypesUnauthorized with default headers values
func NewGetRegistryClientTypesUnauthorized() *GetRegistryClientTypesUnauthorized {
	return &GetRegistryClientTypesUnauthorized{}
}

/*
GetRegistryClientTypesUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetRegistryClientTypesUnauthorized struct {
}

// IsSuccess returns true when this get registry client types unauthorized response has a 2xx status code
func (o *GetRegistryClientTypesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get registry client types unauthorized response has a 3xx status code
func (o *GetRegistryClientTypesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get registry client types unauthorized response has a 4xx status code
func (o *GetRegistryClientTypesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get registry client types unauthorized response has a 5xx status code
func (o *GetRegistryClientTypesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get registry client types unauthorized response a status code equal to that given
func (o *GetRegistryClientTypesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get registry client types unauthorized response
func (o *GetRegistryClientTypesUnauthorized) Code() int {
	return 401
}

func (o *GetRegistryClientTypesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /controller/registry-types][%d] getRegistryClientTypesUnauthorized ", 401)
}

func (o *GetRegistryClientTypesUnauthorized) String() string {
	return fmt.Sprintf("[GET /controller/registry-types][%d] getRegistryClientTypesUnauthorized ", 401)
}

func (o *GetRegistryClientTypesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRegistryClientTypesForbidden creates a GetRegistryClientTypesForbidden with default headers values
func NewGetRegistryClientTypesForbidden() *GetRegistryClientTypesForbidden {
	return &GetRegistryClientTypesForbidden{}
}

/*
GetRegistryClientTypesForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetRegistryClientTypesForbidden struct {
}

// IsSuccess returns true when this get registry client types forbidden response has a 2xx status code
func (o *GetRegistryClientTypesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get registry client types forbidden response has a 3xx status code
func (o *GetRegistryClientTypesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get registry client types forbidden response has a 4xx status code
func (o *GetRegistryClientTypesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get registry client types forbidden response has a 5xx status code
func (o *GetRegistryClientTypesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get registry client types forbidden response a status code equal to that given
func (o *GetRegistryClientTypesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get registry client types forbidden response
func (o *GetRegistryClientTypesForbidden) Code() int {
	return 403
}

func (o *GetRegistryClientTypesForbidden) Error() string {
	return fmt.Sprintf("[GET /controller/registry-types][%d] getRegistryClientTypesForbidden ", 403)
}

func (o *GetRegistryClientTypesForbidden) String() string {
	return fmt.Sprintf("[GET /controller/registry-types][%d] getRegistryClientTypesForbidden ", 403)
}

func (o *GetRegistryClientTypesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRegistryClientTypesConflict creates a GetRegistryClientTypesConflict with default headers values
func NewGetRegistryClientTypesConflict() *GetRegistryClientTypesConflict {
	return &GetRegistryClientTypesConflict{}
}

/*
GetRegistryClientTypesConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetRegistryClientTypesConflict struct {
}

// IsSuccess returns true when this get registry client types conflict response has a 2xx status code
func (o *GetRegistryClientTypesConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get registry client types conflict response has a 3xx status code
func (o *GetRegistryClientTypesConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get registry client types conflict response has a 4xx status code
func (o *GetRegistryClientTypesConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get registry client types conflict response has a 5xx status code
func (o *GetRegistryClientTypesConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get registry client types conflict response a status code equal to that given
func (o *GetRegistryClientTypesConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the get registry client types conflict response
func (o *GetRegistryClientTypesConflict) Code() int {
	return 409
}

func (o *GetRegistryClientTypesConflict) Error() string {
	return fmt.Sprintf("[GET /controller/registry-types][%d] getRegistryClientTypesConflict ", 409)
}

func (o *GetRegistryClientTypesConflict) String() string {
	return fmt.Sprintf("[GET /controller/registry-types][%d] getRegistryClientTypesConflict ", 409)
}

func (o *GetRegistryClientTypesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
