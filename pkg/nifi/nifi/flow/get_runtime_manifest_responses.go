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

// GetRuntimeManifestReader is a Reader for the GetRuntimeManifest structure.
type GetRuntimeManifestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRuntimeManifestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRuntimeManifestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRuntimeManifestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRuntimeManifestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRuntimeManifestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetRuntimeManifestConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /flow/runtime-manifest] getRuntimeManifest", response, response.Code())
	}
}

// NewGetRuntimeManifestOK creates a GetRuntimeManifestOK with default headers values
func NewGetRuntimeManifestOK() *GetRuntimeManifestOK {
	return &GetRuntimeManifestOK{}
}

/*
GetRuntimeManifestOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRuntimeManifestOK struct {
	Payload *models.RuntimeManifestEntity
}

// IsSuccess returns true when this get runtime manifest o k response has a 2xx status code
func (o *GetRuntimeManifestOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get runtime manifest o k response has a 3xx status code
func (o *GetRuntimeManifestOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get runtime manifest o k response has a 4xx status code
func (o *GetRuntimeManifestOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get runtime manifest o k response has a 5xx status code
func (o *GetRuntimeManifestOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get runtime manifest o k response a status code equal to that given
func (o *GetRuntimeManifestOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get runtime manifest o k response
func (o *GetRuntimeManifestOK) Code() int {
	return 200
}

func (o *GetRuntimeManifestOK) Error() string {
	return fmt.Sprintf("[GET /flow/runtime-manifest][%d] getRuntimeManifestOK  %+v", 200, o.Payload)
}

func (o *GetRuntimeManifestOK) String() string {
	return fmt.Sprintf("[GET /flow/runtime-manifest][%d] getRuntimeManifestOK  %+v", 200, o.Payload)
}

func (o *GetRuntimeManifestOK) GetPayload() *models.RuntimeManifestEntity {
	return o.Payload
}

func (o *GetRuntimeManifestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeManifestEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRuntimeManifestBadRequest creates a GetRuntimeManifestBadRequest with default headers values
func NewGetRuntimeManifestBadRequest() *GetRuntimeManifestBadRequest {
	return &GetRuntimeManifestBadRequest{}
}

/*
GetRuntimeManifestBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetRuntimeManifestBadRequest struct {
}

// IsSuccess returns true when this get runtime manifest bad request response has a 2xx status code
func (o *GetRuntimeManifestBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get runtime manifest bad request response has a 3xx status code
func (o *GetRuntimeManifestBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get runtime manifest bad request response has a 4xx status code
func (o *GetRuntimeManifestBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get runtime manifest bad request response has a 5xx status code
func (o *GetRuntimeManifestBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get runtime manifest bad request response a status code equal to that given
func (o *GetRuntimeManifestBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get runtime manifest bad request response
func (o *GetRuntimeManifestBadRequest) Code() int {
	return 400
}

func (o *GetRuntimeManifestBadRequest) Error() string {
	return fmt.Sprintf("[GET /flow/runtime-manifest][%d] getRuntimeManifestBadRequest ", 400)
}

func (o *GetRuntimeManifestBadRequest) String() string {
	return fmt.Sprintf("[GET /flow/runtime-manifest][%d] getRuntimeManifestBadRequest ", 400)
}

func (o *GetRuntimeManifestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRuntimeManifestUnauthorized creates a GetRuntimeManifestUnauthorized with default headers values
func NewGetRuntimeManifestUnauthorized() *GetRuntimeManifestUnauthorized {
	return &GetRuntimeManifestUnauthorized{}
}

/*
GetRuntimeManifestUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetRuntimeManifestUnauthorized struct {
}

// IsSuccess returns true when this get runtime manifest unauthorized response has a 2xx status code
func (o *GetRuntimeManifestUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get runtime manifest unauthorized response has a 3xx status code
func (o *GetRuntimeManifestUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get runtime manifest unauthorized response has a 4xx status code
func (o *GetRuntimeManifestUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get runtime manifest unauthorized response has a 5xx status code
func (o *GetRuntimeManifestUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get runtime manifest unauthorized response a status code equal to that given
func (o *GetRuntimeManifestUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get runtime manifest unauthorized response
func (o *GetRuntimeManifestUnauthorized) Code() int {
	return 401
}

func (o *GetRuntimeManifestUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flow/runtime-manifest][%d] getRuntimeManifestUnauthorized ", 401)
}

func (o *GetRuntimeManifestUnauthorized) String() string {
	return fmt.Sprintf("[GET /flow/runtime-manifest][%d] getRuntimeManifestUnauthorized ", 401)
}

func (o *GetRuntimeManifestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRuntimeManifestForbidden creates a GetRuntimeManifestForbidden with default headers values
func NewGetRuntimeManifestForbidden() *GetRuntimeManifestForbidden {
	return &GetRuntimeManifestForbidden{}
}

/*
GetRuntimeManifestForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetRuntimeManifestForbidden struct {
}

// IsSuccess returns true when this get runtime manifest forbidden response has a 2xx status code
func (o *GetRuntimeManifestForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get runtime manifest forbidden response has a 3xx status code
func (o *GetRuntimeManifestForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get runtime manifest forbidden response has a 4xx status code
func (o *GetRuntimeManifestForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get runtime manifest forbidden response has a 5xx status code
func (o *GetRuntimeManifestForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get runtime manifest forbidden response a status code equal to that given
func (o *GetRuntimeManifestForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get runtime manifest forbidden response
func (o *GetRuntimeManifestForbidden) Code() int {
	return 403
}

func (o *GetRuntimeManifestForbidden) Error() string {
	return fmt.Sprintf("[GET /flow/runtime-manifest][%d] getRuntimeManifestForbidden ", 403)
}

func (o *GetRuntimeManifestForbidden) String() string {
	return fmt.Sprintf("[GET /flow/runtime-manifest][%d] getRuntimeManifestForbidden ", 403)
}

func (o *GetRuntimeManifestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRuntimeManifestConflict creates a GetRuntimeManifestConflict with default headers values
func NewGetRuntimeManifestConflict() *GetRuntimeManifestConflict {
	return &GetRuntimeManifestConflict{}
}

/*
GetRuntimeManifestConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetRuntimeManifestConflict struct {
}

// IsSuccess returns true when this get runtime manifest conflict response has a 2xx status code
func (o *GetRuntimeManifestConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get runtime manifest conflict response has a 3xx status code
func (o *GetRuntimeManifestConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get runtime manifest conflict response has a 4xx status code
func (o *GetRuntimeManifestConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get runtime manifest conflict response has a 5xx status code
func (o *GetRuntimeManifestConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get runtime manifest conflict response a status code equal to that given
func (o *GetRuntimeManifestConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the get runtime manifest conflict response
func (o *GetRuntimeManifestConflict) Code() int {
	return 409
}

func (o *GetRuntimeManifestConflict) Error() string {
	return fmt.Sprintf("[GET /flow/runtime-manifest][%d] getRuntimeManifestConflict ", 409)
}

func (o *GetRuntimeManifestConflict) String() string {
	return fmt.Sprintf("[GET /flow/runtime-manifest][%d] getRuntimeManifestConflict ", 409)
}

func (o *GetRuntimeManifestConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
