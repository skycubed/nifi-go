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

// GetControllerServiceTypesReader is a Reader for the GetControllerServiceTypes structure.
type GetControllerServiceTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetControllerServiceTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetControllerServiceTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetControllerServiceTypesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetControllerServiceTypesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetControllerServiceTypesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetControllerServiceTypesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetControllerServiceTypesOK creates a GetControllerServiceTypesOK with default headers values
func NewGetControllerServiceTypesOK() *GetControllerServiceTypesOK {
	return &GetControllerServiceTypesOK{}
}

/*
GetControllerServiceTypesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetControllerServiceTypesOK struct {
	Payload *models.ControllerServiceTypesEntity
}

// IsSuccess returns true when this get controller service types o k response has a 2xx status code
func (o *GetControllerServiceTypesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get controller service types o k response has a 3xx status code
func (o *GetControllerServiceTypesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get controller service types o k response has a 4xx status code
func (o *GetControllerServiceTypesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get controller service types o k response has a 5xx status code
func (o *GetControllerServiceTypesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get controller service types o k response a status code equal to that given
func (o *GetControllerServiceTypesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetControllerServiceTypesOK) Error() string {
	return fmt.Sprintf("[GET /flow/controller-service-types][%d] getControllerServiceTypesOK  %+v", 200, o.Payload)
}

func (o *GetControllerServiceTypesOK) String() string {
	return fmt.Sprintf("[GET /flow/controller-service-types][%d] getControllerServiceTypesOK  %+v", 200, o.Payload)
}

func (o *GetControllerServiceTypesOK) GetPayload() *models.ControllerServiceTypesEntity {
	return o.Payload
}

func (o *GetControllerServiceTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ControllerServiceTypesEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetControllerServiceTypesBadRequest creates a GetControllerServiceTypesBadRequest with default headers values
func NewGetControllerServiceTypesBadRequest() *GetControllerServiceTypesBadRequest {
	return &GetControllerServiceTypesBadRequest{}
}

/*
GetControllerServiceTypesBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetControllerServiceTypesBadRequest struct {
}

// IsSuccess returns true when this get controller service types bad request response has a 2xx status code
func (o *GetControllerServiceTypesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get controller service types bad request response has a 3xx status code
func (o *GetControllerServiceTypesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get controller service types bad request response has a 4xx status code
func (o *GetControllerServiceTypesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get controller service types bad request response has a 5xx status code
func (o *GetControllerServiceTypesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get controller service types bad request response a status code equal to that given
func (o *GetControllerServiceTypesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetControllerServiceTypesBadRequest) Error() string {
	return fmt.Sprintf("[GET /flow/controller-service-types][%d] getControllerServiceTypesBadRequest ", 400)
}

func (o *GetControllerServiceTypesBadRequest) String() string {
	return fmt.Sprintf("[GET /flow/controller-service-types][%d] getControllerServiceTypesBadRequest ", 400)
}

func (o *GetControllerServiceTypesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetControllerServiceTypesUnauthorized creates a GetControllerServiceTypesUnauthorized with default headers values
func NewGetControllerServiceTypesUnauthorized() *GetControllerServiceTypesUnauthorized {
	return &GetControllerServiceTypesUnauthorized{}
}

/*
GetControllerServiceTypesUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetControllerServiceTypesUnauthorized struct {
}

// IsSuccess returns true when this get controller service types unauthorized response has a 2xx status code
func (o *GetControllerServiceTypesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get controller service types unauthorized response has a 3xx status code
func (o *GetControllerServiceTypesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get controller service types unauthorized response has a 4xx status code
func (o *GetControllerServiceTypesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get controller service types unauthorized response has a 5xx status code
func (o *GetControllerServiceTypesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get controller service types unauthorized response a status code equal to that given
func (o *GetControllerServiceTypesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetControllerServiceTypesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flow/controller-service-types][%d] getControllerServiceTypesUnauthorized ", 401)
}

func (o *GetControllerServiceTypesUnauthorized) String() string {
	return fmt.Sprintf("[GET /flow/controller-service-types][%d] getControllerServiceTypesUnauthorized ", 401)
}

func (o *GetControllerServiceTypesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetControllerServiceTypesForbidden creates a GetControllerServiceTypesForbidden with default headers values
func NewGetControllerServiceTypesForbidden() *GetControllerServiceTypesForbidden {
	return &GetControllerServiceTypesForbidden{}
}

/*
GetControllerServiceTypesForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetControllerServiceTypesForbidden struct {
}

// IsSuccess returns true when this get controller service types forbidden response has a 2xx status code
func (o *GetControllerServiceTypesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get controller service types forbidden response has a 3xx status code
func (o *GetControllerServiceTypesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get controller service types forbidden response has a 4xx status code
func (o *GetControllerServiceTypesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get controller service types forbidden response has a 5xx status code
func (o *GetControllerServiceTypesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get controller service types forbidden response a status code equal to that given
func (o *GetControllerServiceTypesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetControllerServiceTypesForbidden) Error() string {
	return fmt.Sprintf("[GET /flow/controller-service-types][%d] getControllerServiceTypesForbidden ", 403)
}

func (o *GetControllerServiceTypesForbidden) String() string {
	return fmt.Sprintf("[GET /flow/controller-service-types][%d] getControllerServiceTypesForbidden ", 403)
}

func (o *GetControllerServiceTypesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetControllerServiceTypesConflict creates a GetControllerServiceTypesConflict with default headers values
func NewGetControllerServiceTypesConflict() *GetControllerServiceTypesConflict {
	return &GetControllerServiceTypesConflict{}
}

/*
GetControllerServiceTypesConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetControllerServiceTypesConflict struct {
}

// IsSuccess returns true when this get controller service types conflict response has a 2xx status code
func (o *GetControllerServiceTypesConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get controller service types conflict response has a 3xx status code
func (o *GetControllerServiceTypesConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get controller service types conflict response has a 4xx status code
func (o *GetControllerServiceTypesConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get controller service types conflict response has a 5xx status code
func (o *GetControllerServiceTypesConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get controller service types conflict response a status code equal to that given
func (o *GetControllerServiceTypesConflict) IsCode(code int) bool {
	return code == 409
}

func (o *GetControllerServiceTypesConflict) Error() string {
	return fmt.Sprintf("[GET /flow/controller-service-types][%d] getControllerServiceTypesConflict ", 409)
}

func (o *GetControllerServiceTypesConflict) String() string {
	return fmt.Sprintf("[GET /flow/controller-service-types][%d] getControllerServiceTypesConflict ", 409)
}

func (o *GetControllerServiceTypesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
