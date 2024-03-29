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

// GetParameterContextsReader is a Reader for the GetParameterContexts structure.
type GetParameterContextsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetParameterContextsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetParameterContextsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetParameterContextsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetParameterContextsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetParameterContextsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetParameterContextsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /flow/parameter-contexts] getParameterContexts", response, response.Code())
	}
}

// NewGetParameterContextsOK creates a GetParameterContextsOK with default headers values
func NewGetParameterContextsOK() *GetParameterContextsOK {
	return &GetParameterContextsOK{}
}

/*
GetParameterContextsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetParameterContextsOK struct {
	Payload *models.ParameterContextsEntity
}

// IsSuccess returns true when this get parameter contexts o k response has a 2xx status code
func (o *GetParameterContextsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get parameter contexts o k response has a 3xx status code
func (o *GetParameterContextsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get parameter contexts o k response has a 4xx status code
func (o *GetParameterContextsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get parameter contexts o k response has a 5xx status code
func (o *GetParameterContextsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get parameter contexts o k response a status code equal to that given
func (o *GetParameterContextsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get parameter contexts o k response
func (o *GetParameterContextsOK) Code() int {
	return 200
}

func (o *GetParameterContextsOK) Error() string {
	return fmt.Sprintf("[GET /flow/parameter-contexts][%d] getParameterContextsOK  %+v", 200, o.Payload)
}

func (o *GetParameterContextsOK) String() string {
	return fmt.Sprintf("[GET /flow/parameter-contexts][%d] getParameterContextsOK  %+v", 200, o.Payload)
}

func (o *GetParameterContextsOK) GetPayload() *models.ParameterContextsEntity {
	return o.Payload
}

func (o *GetParameterContextsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ParameterContextsEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetParameterContextsBadRequest creates a GetParameterContextsBadRequest with default headers values
func NewGetParameterContextsBadRequest() *GetParameterContextsBadRequest {
	return &GetParameterContextsBadRequest{}
}

/*
GetParameterContextsBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetParameterContextsBadRequest struct {
}

// IsSuccess returns true when this get parameter contexts bad request response has a 2xx status code
func (o *GetParameterContextsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get parameter contexts bad request response has a 3xx status code
func (o *GetParameterContextsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get parameter contexts bad request response has a 4xx status code
func (o *GetParameterContextsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get parameter contexts bad request response has a 5xx status code
func (o *GetParameterContextsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get parameter contexts bad request response a status code equal to that given
func (o *GetParameterContextsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get parameter contexts bad request response
func (o *GetParameterContextsBadRequest) Code() int {
	return 400
}

func (o *GetParameterContextsBadRequest) Error() string {
	return fmt.Sprintf("[GET /flow/parameter-contexts][%d] getParameterContextsBadRequest ", 400)
}

func (o *GetParameterContextsBadRequest) String() string {
	return fmt.Sprintf("[GET /flow/parameter-contexts][%d] getParameterContextsBadRequest ", 400)
}

func (o *GetParameterContextsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetParameterContextsUnauthorized creates a GetParameterContextsUnauthorized with default headers values
func NewGetParameterContextsUnauthorized() *GetParameterContextsUnauthorized {
	return &GetParameterContextsUnauthorized{}
}

/*
GetParameterContextsUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetParameterContextsUnauthorized struct {
}

// IsSuccess returns true when this get parameter contexts unauthorized response has a 2xx status code
func (o *GetParameterContextsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get parameter contexts unauthorized response has a 3xx status code
func (o *GetParameterContextsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get parameter contexts unauthorized response has a 4xx status code
func (o *GetParameterContextsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get parameter contexts unauthorized response has a 5xx status code
func (o *GetParameterContextsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get parameter contexts unauthorized response a status code equal to that given
func (o *GetParameterContextsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get parameter contexts unauthorized response
func (o *GetParameterContextsUnauthorized) Code() int {
	return 401
}

func (o *GetParameterContextsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flow/parameter-contexts][%d] getParameterContextsUnauthorized ", 401)
}

func (o *GetParameterContextsUnauthorized) String() string {
	return fmt.Sprintf("[GET /flow/parameter-contexts][%d] getParameterContextsUnauthorized ", 401)
}

func (o *GetParameterContextsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetParameterContextsForbidden creates a GetParameterContextsForbidden with default headers values
func NewGetParameterContextsForbidden() *GetParameterContextsForbidden {
	return &GetParameterContextsForbidden{}
}

/*
GetParameterContextsForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetParameterContextsForbidden struct {
}

// IsSuccess returns true when this get parameter contexts forbidden response has a 2xx status code
func (o *GetParameterContextsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get parameter contexts forbidden response has a 3xx status code
func (o *GetParameterContextsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get parameter contexts forbidden response has a 4xx status code
func (o *GetParameterContextsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get parameter contexts forbidden response has a 5xx status code
func (o *GetParameterContextsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get parameter contexts forbidden response a status code equal to that given
func (o *GetParameterContextsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get parameter contexts forbidden response
func (o *GetParameterContextsForbidden) Code() int {
	return 403
}

func (o *GetParameterContextsForbidden) Error() string {
	return fmt.Sprintf("[GET /flow/parameter-contexts][%d] getParameterContextsForbidden ", 403)
}

func (o *GetParameterContextsForbidden) String() string {
	return fmt.Sprintf("[GET /flow/parameter-contexts][%d] getParameterContextsForbidden ", 403)
}

func (o *GetParameterContextsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetParameterContextsConflict creates a GetParameterContextsConflict with default headers values
func NewGetParameterContextsConflict() *GetParameterContextsConflict {
	return &GetParameterContextsConflict{}
}

/*
GetParameterContextsConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetParameterContextsConflict struct {
}

// IsSuccess returns true when this get parameter contexts conflict response has a 2xx status code
func (o *GetParameterContextsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get parameter contexts conflict response has a 3xx status code
func (o *GetParameterContextsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get parameter contexts conflict response has a 4xx status code
func (o *GetParameterContextsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get parameter contexts conflict response has a 5xx status code
func (o *GetParameterContextsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get parameter contexts conflict response a status code equal to that given
func (o *GetParameterContextsConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the get parameter contexts conflict response
func (o *GetParameterContextsConflict) Code() int {
	return 409
}

func (o *GetParameterContextsConflict) Error() string {
	return fmt.Sprintf("[GET /flow/parameter-contexts][%d] getParameterContextsConflict ", 409)
}

func (o *GetParameterContextsConflict) String() string {
	return fmt.Sprintf("[GET /flow/parameter-contexts][%d] getParameterContextsConflict ", 409)
}

func (o *GetParameterContextsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
