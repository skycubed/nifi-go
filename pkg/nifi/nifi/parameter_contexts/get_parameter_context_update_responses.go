// Code generated by go-swagger; DO NOT EDIT.

package parameter_contexts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// GetParameterContextUpdateReader is a Reader for the GetParameterContextUpdate structure.
type GetParameterContextUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetParameterContextUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetParameterContextUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetParameterContextUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetParameterContextUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetParameterContextUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetParameterContextUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetParameterContextUpdateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /parameter-contexts/{contextId}/update-requests/{requestId}] getParameterContextUpdate", response, response.Code())
	}
}

// NewGetParameterContextUpdateOK creates a GetParameterContextUpdateOK with default headers values
func NewGetParameterContextUpdateOK() *GetParameterContextUpdateOK {
	return &GetParameterContextUpdateOK{}
}

/*
GetParameterContextUpdateOK describes a response with status code 200, with default header values.

successful operation
*/
type GetParameterContextUpdateOK struct {
	Payload *models.ParameterContextUpdateRequestEntity
}

// IsSuccess returns true when this get parameter context update o k response has a 2xx status code
func (o *GetParameterContextUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get parameter context update o k response has a 3xx status code
func (o *GetParameterContextUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get parameter context update o k response has a 4xx status code
func (o *GetParameterContextUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get parameter context update o k response has a 5xx status code
func (o *GetParameterContextUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get parameter context update o k response a status code equal to that given
func (o *GetParameterContextUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get parameter context update o k response
func (o *GetParameterContextUpdateOK) Code() int {
	return 200
}

func (o *GetParameterContextUpdateOK) Error() string {
	return fmt.Sprintf("[GET /parameter-contexts/{contextId}/update-requests/{requestId}][%d] getParameterContextUpdateOK  %+v", 200, o.Payload)
}

func (o *GetParameterContextUpdateOK) String() string {
	return fmt.Sprintf("[GET /parameter-contexts/{contextId}/update-requests/{requestId}][%d] getParameterContextUpdateOK  %+v", 200, o.Payload)
}

func (o *GetParameterContextUpdateOK) GetPayload() *models.ParameterContextUpdateRequestEntity {
	return o.Payload
}

func (o *GetParameterContextUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ParameterContextUpdateRequestEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetParameterContextUpdateBadRequest creates a GetParameterContextUpdateBadRequest with default headers values
func NewGetParameterContextUpdateBadRequest() *GetParameterContextUpdateBadRequest {
	return &GetParameterContextUpdateBadRequest{}
}

/*
GetParameterContextUpdateBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetParameterContextUpdateBadRequest struct {
}

// IsSuccess returns true when this get parameter context update bad request response has a 2xx status code
func (o *GetParameterContextUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get parameter context update bad request response has a 3xx status code
func (o *GetParameterContextUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get parameter context update bad request response has a 4xx status code
func (o *GetParameterContextUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get parameter context update bad request response has a 5xx status code
func (o *GetParameterContextUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get parameter context update bad request response a status code equal to that given
func (o *GetParameterContextUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get parameter context update bad request response
func (o *GetParameterContextUpdateBadRequest) Code() int {
	return 400
}

func (o *GetParameterContextUpdateBadRequest) Error() string {
	return fmt.Sprintf("[GET /parameter-contexts/{contextId}/update-requests/{requestId}][%d] getParameterContextUpdateBadRequest ", 400)
}

func (o *GetParameterContextUpdateBadRequest) String() string {
	return fmt.Sprintf("[GET /parameter-contexts/{contextId}/update-requests/{requestId}][%d] getParameterContextUpdateBadRequest ", 400)
}

func (o *GetParameterContextUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetParameterContextUpdateUnauthorized creates a GetParameterContextUpdateUnauthorized with default headers values
func NewGetParameterContextUpdateUnauthorized() *GetParameterContextUpdateUnauthorized {
	return &GetParameterContextUpdateUnauthorized{}
}

/*
GetParameterContextUpdateUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetParameterContextUpdateUnauthorized struct {
}

// IsSuccess returns true when this get parameter context update unauthorized response has a 2xx status code
func (o *GetParameterContextUpdateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get parameter context update unauthorized response has a 3xx status code
func (o *GetParameterContextUpdateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get parameter context update unauthorized response has a 4xx status code
func (o *GetParameterContextUpdateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get parameter context update unauthorized response has a 5xx status code
func (o *GetParameterContextUpdateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get parameter context update unauthorized response a status code equal to that given
func (o *GetParameterContextUpdateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get parameter context update unauthorized response
func (o *GetParameterContextUpdateUnauthorized) Code() int {
	return 401
}

func (o *GetParameterContextUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[GET /parameter-contexts/{contextId}/update-requests/{requestId}][%d] getParameterContextUpdateUnauthorized ", 401)
}

func (o *GetParameterContextUpdateUnauthorized) String() string {
	return fmt.Sprintf("[GET /parameter-contexts/{contextId}/update-requests/{requestId}][%d] getParameterContextUpdateUnauthorized ", 401)
}

func (o *GetParameterContextUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetParameterContextUpdateForbidden creates a GetParameterContextUpdateForbidden with default headers values
func NewGetParameterContextUpdateForbidden() *GetParameterContextUpdateForbidden {
	return &GetParameterContextUpdateForbidden{}
}

/*
GetParameterContextUpdateForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetParameterContextUpdateForbidden struct {
}

// IsSuccess returns true when this get parameter context update forbidden response has a 2xx status code
func (o *GetParameterContextUpdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get parameter context update forbidden response has a 3xx status code
func (o *GetParameterContextUpdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get parameter context update forbidden response has a 4xx status code
func (o *GetParameterContextUpdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get parameter context update forbidden response has a 5xx status code
func (o *GetParameterContextUpdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get parameter context update forbidden response a status code equal to that given
func (o *GetParameterContextUpdateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get parameter context update forbidden response
func (o *GetParameterContextUpdateForbidden) Code() int {
	return 403
}

func (o *GetParameterContextUpdateForbidden) Error() string {
	return fmt.Sprintf("[GET /parameter-contexts/{contextId}/update-requests/{requestId}][%d] getParameterContextUpdateForbidden ", 403)
}

func (o *GetParameterContextUpdateForbidden) String() string {
	return fmt.Sprintf("[GET /parameter-contexts/{contextId}/update-requests/{requestId}][%d] getParameterContextUpdateForbidden ", 403)
}

func (o *GetParameterContextUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetParameterContextUpdateNotFound creates a GetParameterContextUpdateNotFound with default headers values
func NewGetParameterContextUpdateNotFound() *GetParameterContextUpdateNotFound {
	return &GetParameterContextUpdateNotFound{}
}

/*
GetParameterContextUpdateNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetParameterContextUpdateNotFound struct {
}

// IsSuccess returns true when this get parameter context update not found response has a 2xx status code
func (o *GetParameterContextUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get parameter context update not found response has a 3xx status code
func (o *GetParameterContextUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get parameter context update not found response has a 4xx status code
func (o *GetParameterContextUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get parameter context update not found response has a 5xx status code
func (o *GetParameterContextUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get parameter context update not found response a status code equal to that given
func (o *GetParameterContextUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get parameter context update not found response
func (o *GetParameterContextUpdateNotFound) Code() int {
	return 404
}

func (o *GetParameterContextUpdateNotFound) Error() string {
	return fmt.Sprintf("[GET /parameter-contexts/{contextId}/update-requests/{requestId}][%d] getParameterContextUpdateNotFound ", 404)
}

func (o *GetParameterContextUpdateNotFound) String() string {
	return fmt.Sprintf("[GET /parameter-contexts/{contextId}/update-requests/{requestId}][%d] getParameterContextUpdateNotFound ", 404)
}

func (o *GetParameterContextUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetParameterContextUpdateConflict creates a GetParameterContextUpdateConflict with default headers values
func NewGetParameterContextUpdateConflict() *GetParameterContextUpdateConflict {
	return &GetParameterContextUpdateConflict{}
}

/*
GetParameterContextUpdateConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetParameterContextUpdateConflict struct {
}

// IsSuccess returns true when this get parameter context update conflict response has a 2xx status code
func (o *GetParameterContextUpdateConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get parameter context update conflict response has a 3xx status code
func (o *GetParameterContextUpdateConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get parameter context update conflict response has a 4xx status code
func (o *GetParameterContextUpdateConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get parameter context update conflict response has a 5xx status code
func (o *GetParameterContextUpdateConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get parameter context update conflict response a status code equal to that given
func (o *GetParameterContextUpdateConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the get parameter context update conflict response
func (o *GetParameterContextUpdateConflict) Code() int {
	return 409
}

func (o *GetParameterContextUpdateConflict) Error() string {
	return fmt.Sprintf("[GET /parameter-contexts/{contextId}/update-requests/{requestId}][%d] getParameterContextUpdateConflict ", 409)
}

func (o *GetParameterContextUpdateConflict) String() string {
	return fmt.Sprintf("[GET /parameter-contexts/{contextId}/update-requests/{requestId}][%d] getParameterContextUpdateConflict ", 409)
}

func (o *GetParameterContextUpdateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
