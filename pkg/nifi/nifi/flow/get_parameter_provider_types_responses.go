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

// GetParameterProviderTypesReader is a Reader for the GetParameterProviderTypes structure.
type GetParameterProviderTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetParameterProviderTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetParameterProviderTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetParameterProviderTypesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetParameterProviderTypesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetParameterProviderTypesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetParameterProviderTypesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetParameterProviderTypesOK creates a GetParameterProviderTypesOK with default headers values
func NewGetParameterProviderTypesOK() *GetParameterProviderTypesOK {
	return &GetParameterProviderTypesOK{}
}

/*
GetParameterProviderTypesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetParameterProviderTypesOK struct {
	Payload *models.ParameterProviderTypesEntity
}

// IsSuccess returns true when this get parameter provider types o k response has a 2xx status code
func (o *GetParameterProviderTypesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get parameter provider types o k response has a 3xx status code
func (o *GetParameterProviderTypesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get parameter provider types o k response has a 4xx status code
func (o *GetParameterProviderTypesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get parameter provider types o k response has a 5xx status code
func (o *GetParameterProviderTypesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get parameter provider types o k response a status code equal to that given
func (o *GetParameterProviderTypesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetParameterProviderTypesOK) Error() string {
	return fmt.Sprintf("[GET /flow/parameter-provider-types][%d] getParameterProviderTypesOK  %+v", 200, o.Payload)
}

func (o *GetParameterProviderTypesOK) String() string {
	return fmt.Sprintf("[GET /flow/parameter-provider-types][%d] getParameterProviderTypesOK  %+v", 200, o.Payload)
}

func (o *GetParameterProviderTypesOK) GetPayload() *models.ParameterProviderTypesEntity {
	return o.Payload
}

func (o *GetParameterProviderTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ParameterProviderTypesEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetParameterProviderTypesBadRequest creates a GetParameterProviderTypesBadRequest with default headers values
func NewGetParameterProviderTypesBadRequest() *GetParameterProviderTypesBadRequest {
	return &GetParameterProviderTypesBadRequest{}
}

/*
GetParameterProviderTypesBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetParameterProviderTypesBadRequest struct {
}

// IsSuccess returns true when this get parameter provider types bad request response has a 2xx status code
func (o *GetParameterProviderTypesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get parameter provider types bad request response has a 3xx status code
func (o *GetParameterProviderTypesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get parameter provider types bad request response has a 4xx status code
func (o *GetParameterProviderTypesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get parameter provider types bad request response has a 5xx status code
func (o *GetParameterProviderTypesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get parameter provider types bad request response a status code equal to that given
func (o *GetParameterProviderTypesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetParameterProviderTypesBadRequest) Error() string {
	return fmt.Sprintf("[GET /flow/parameter-provider-types][%d] getParameterProviderTypesBadRequest ", 400)
}

func (o *GetParameterProviderTypesBadRequest) String() string {
	return fmt.Sprintf("[GET /flow/parameter-provider-types][%d] getParameterProviderTypesBadRequest ", 400)
}

func (o *GetParameterProviderTypesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetParameterProviderTypesUnauthorized creates a GetParameterProviderTypesUnauthorized with default headers values
func NewGetParameterProviderTypesUnauthorized() *GetParameterProviderTypesUnauthorized {
	return &GetParameterProviderTypesUnauthorized{}
}

/*
GetParameterProviderTypesUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetParameterProviderTypesUnauthorized struct {
}

// IsSuccess returns true when this get parameter provider types unauthorized response has a 2xx status code
func (o *GetParameterProviderTypesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get parameter provider types unauthorized response has a 3xx status code
func (o *GetParameterProviderTypesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get parameter provider types unauthorized response has a 4xx status code
func (o *GetParameterProviderTypesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get parameter provider types unauthorized response has a 5xx status code
func (o *GetParameterProviderTypesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get parameter provider types unauthorized response a status code equal to that given
func (o *GetParameterProviderTypesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetParameterProviderTypesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flow/parameter-provider-types][%d] getParameterProviderTypesUnauthorized ", 401)
}

func (o *GetParameterProviderTypesUnauthorized) String() string {
	return fmt.Sprintf("[GET /flow/parameter-provider-types][%d] getParameterProviderTypesUnauthorized ", 401)
}

func (o *GetParameterProviderTypesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetParameterProviderTypesForbidden creates a GetParameterProviderTypesForbidden with default headers values
func NewGetParameterProviderTypesForbidden() *GetParameterProviderTypesForbidden {
	return &GetParameterProviderTypesForbidden{}
}

/*
GetParameterProviderTypesForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetParameterProviderTypesForbidden struct {
}

// IsSuccess returns true when this get parameter provider types forbidden response has a 2xx status code
func (o *GetParameterProviderTypesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get parameter provider types forbidden response has a 3xx status code
func (o *GetParameterProviderTypesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get parameter provider types forbidden response has a 4xx status code
func (o *GetParameterProviderTypesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get parameter provider types forbidden response has a 5xx status code
func (o *GetParameterProviderTypesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get parameter provider types forbidden response a status code equal to that given
func (o *GetParameterProviderTypesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetParameterProviderTypesForbidden) Error() string {
	return fmt.Sprintf("[GET /flow/parameter-provider-types][%d] getParameterProviderTypesForbidden ", 403)
}

func (o *GetParameterProviderTypesForbidden) String() string {
	return fmt.Sprintf("[GET /flow/parameter-provider-types][%d] getParameterProviderTypesForbidden ", 403)
}

func (o *GetParameterProviderTypesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetParameterProviderTypesConflict creates a GetParameterProviderTypesConflict with default headers values
func NewGetParameterProviderTypesConflict() *GetParameterProviderTypesConflict {
	return &GetParameterProviderTypesConflict{}
}

/*
GetParameterProviderTypesConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetParameterProviderTypesConflict struct {
}

// IsSuccess returns true when this get parameter provider types conflict response has a 2xx status code
func (o *GetParameterProviderTypesConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get parameter provider types conflict response has a 3xx status code
func (o *GetParameterProviderTypesConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get parameter provider types conflict response has a 4xx status code
func (o *GetParameterProviderTypesConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get parameter provider types conflict response has a 5xx status code
func (o *GetParameterProviderTypesConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get parameter provider types conflict response a status code equal to that given
func (o *GetParameterProviderTypesConflict) IsCode(code int) bool {
	return code == 409
}

func (o *GetParameterProviderTypesConflict) Error() string {
	return fmt.Sprintf("[GET /flow/parameter-provider-types][%d] getParameterProviderTypesConflict ", 409)
}

func (o *GetParameterProviderTypesConflict) String() string {
	return fmt.Sprintf("[GET /flow/parameter-provider-types][%d] getParameterProviderTypesConflict ", 409)
}

func (o *GetParameterProviderTypesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
