// Code generated by go-swagger; DO NOT EDIT.

package controller_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// GetControllerServiceReferencesReader is a Reader for the GetControllerServiceReferences structure.
type GetControllerServiceReferencesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetControllerServiceReferencesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetControllerServiceReferencesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetControllerServiceReferencesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetControllerServiceReferencesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetControllerServiceReferencesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetControllerServiceReferencesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetControllerServiceReferencesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /controller-services/{id}/references] getControllerServiceReferences", response, response.Code())
	}
}

// NewGetControllerServiceReferencesOK creates a GetControllerServiceReferencesOK with default headers values
func NewGetControllerServiceReferencesOK() *GetControllerServiceReferencesOK {
	return &GetControllerServiceReferencesOK{}
}

/*
GetControllerServiceReferencesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetControllerServiceReferencesOK struct {
	Payload *models.ControllerServiceReferencingComponentsEntity
}

// IsSuccess returns true when this get controller service references o k response has a 2xx status code
func (o *GetControllerServiceReferencesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get controller service references o k response has a 3xx status code
func (o *GetControllerServiceReferencesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get controller service references o k response has a 4xx status code
func (o *GetControllerServiceReferencesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get controller service references o k response has a 5xx status code
func (o *GetControllerServiceReferencesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get controller service references o k response a status code equal to that given
func (o *GetControllerServiceReferencesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get controller service references o k response
func (o *GetControllerServiceReferencesOK) Code() int {
	return 200
}

func (o *GetControllerServiceReferencesOK) Error() string {
	return fmt.Sprintf("[GET /controller-services/{id}/references][%d] getControllerServiceReferencesOK  %+v", 200, o.Payload)
}

func (o *GetControllerServiceReferencesOK) String() string {
	return fmt.Sprintf("[GET /controller-services/{id}/references][%d] getControllerServiceReferencesOK  %+v", 200, o.Payload)
}

func (o *GetControllerServiceReferencesOK) GetPayload() *models.ControllerServiceReferencingComponentsEntity {
	return o.Payload
}

func (o *GetControllerServiceReferencesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ControllerServiceReferencingComponentsEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetControllerServiceReferencesBadRequest creates a GetControllerServiceReferencesBadRequest with default headers values
func NewGetControllerServiceReferencesBadRequest() *GetControllerServiceReferencesBadRequest {
	return &GetControllerServiceReferencesBadRequest{}
}

/*
GetControllerServiceReferencesBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetControllerServiceReferencesBadRequest struct {
}

// IsSuccess returns true when this get controller service references bad request response has a 2xx status code
func (o *GetControllerServiceReferencesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get controller service references bad request response has a 3xx status code
func (o *GetControllerServiceReferencesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get controller service references bad request response has a 4xx status code
func (o *GetControllerServiceReferencesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get controller service references bad request response has a 5xx status code
func (o *GetControllerServiceReferencesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get controller service references bad request response a status code equal to that given
func (o *GetControllerServiceReferencesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get controller service references bad request response
func (o *GetControllerServiceReferencesBadRequest) Code() int {
	return 400
}

func (o *GetControllerServiceReferencesBadRequest) Error() string {
	return fmt.Sprintf("[GET /controller-services/{id}/references][%d] getControllerServiceReferencesBadRequest ", 400)
}

func (o *GetControllerServiceReferencesBadRequest) String() string {
	return fmt.Sprintf("[GET /controller-services/{id}/references][%d] getControllerServiceReferencesBadRequest ", 400)
}

func (o *GetControllerServiceReferencesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetControllerServiceReferencesUnauthorized creates a GetControllerServiceReferencesUnauthorized with default headers values
func NewGetControllerServiceReferencesUnauthorized() *GetControllerServiceReferencesUnauthorized {
	return &GetControllerServiceReferencesUnauthorized{}
}

/*
GetControllerServiceReferencesUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetControllerServiceReferencesUnauthorized struct {
}

// IsSuccess returns true when this get controller service references unauthorized response has a 2xx status code
func (o *GetControllerServiceReferencesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get controller service references unauthorized response has a 3xx status code
func (o *GetControllerServiceReferencesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get controller service references unauthorized response has a 4xx status code
func (o *GetControllerServiceReferencesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get controller service references unauthorized response has a 5xx status code
func (o *GetControllerServiceReferencesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get controller service references unauthorized response a status code equal to that given
func (o *GetControllerServiceReferencesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get controller service references unauthorized response
func (o *GetControllerServiceReferencesUnauthorized) Code() int {
	return 401
}

func (o *GetControllerServiceReferencesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /controller-services/{id}/references][%d] getControllerServiceReferencesUnauthorized ", 401)
}

func (o *GetControllerServiceReferencesUnauthorized) String() string {
	return fmt.Sprintf("[GET /controller-services/{id}/references][%d] getControllerServiceReferencesUnauthorized ", 401)
}

func (o *GetControllerServiceReferencesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetControllerServiceReferencesForbidden creates a GetControllerServiceReferencesForbidden with default headers values
func NewGetControllerServiceReferencesForbidden() *GetControllerServiceReferencesForbidden {
	return &GetControllerServiceReferencesForbidden{}
}

/*
GetControllerServiceReferencesForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetControllerServiceReferencesForbidden struct {
}

// IsSuccess returns true when this get controller service references forbidden response has a 2xx status code
func (o *GetControllerServiceReferencesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get controller service references forbidden response has a 3xx status code
func (o *GetControllerServiceReferencesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get controller service references forbidden response has a 4xx status code
func (o *GetControllerServiceReferencesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get controller service references forbidden response has a 5xx status code
func (o *GetControllerServiceReferencesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get controller service references forbidden response a status code equal to that given
func (o *GetControllerServiceReferencesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get controller service references forbidden response
func (o *GetControllerServiceReferencesForbidden) Code() int {
	return 403
}

func (o *GetControllerServiceReferencesForbidden) Error() string {
	return fmt.Sprintf("[GET /controller-services/{id}/references][%d] getControllerServiceReferencesForbidden ", 403)
}

func (o *GetControllerServiceReferencesForbidden) String() string {
	return fmt.Sprintf("[GET /controller-services/{id}/references][%d] getControllerServiceReferencesForbidden ", 403)
}

func (o *GetControllerServiceReferencesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetControllerServiceReferencesNotFound creates a GetControllerServiceReferencesNotFound with default headers values
func NewGetControllerServiceReferencesNotFound() *GetControllerServiceReferencesNotFound {
	return &GetControllerServiceReferencesNotFound{}
}

/*
GetControllerServiceReferencesNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetControllerServiceReferencesNotFound struct {
}

// IsSuccess returns true when this get controller service references not found response has a 2xx status code
func (o *GetControllerServiceReferencesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get controller service references not found response has a 3xx status code
func (o *GetControllerServiceReferencesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get controller service references not found response has a 4xx status code
func (o *GetControllerServiceReferencesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get controller service references not found response has a 5xx status code
func (o *GetControllerServiceReferencesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get controller service references not found response a status code equal to that given
func (o *GetControllerServiceReferencesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get controller service references not found response
func (o *GetControllerServiceReferencesNotFound) Code() int {
	return 404
}

func (o *GetControllerServiceReferencesNotFound) Error() string {
	return fmt.Sprintf("[GET /controller-services/{id}/references][%d] getControllerServiceReferencesNotFound ", 404)
}

func (o *GetControllerServiceReferencesNotFound) String() string {
	return fmt.Sprintf("[GET /controller-services/{id}/references][%d] getControllerServiceReferencesNotFound ", 404)
}

func (o *GetControllerServiceReferencesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetControllerServiceReferencesConflict creates a GetControllerServiceReferencesConflict with default headers values
func NewGetControllerServiceReferencesConflict() *GetControllerServiceReferencesConflict {
	return &GetControllerServiceReferencesConflict{}
}

/*
GetControllerServiceReferencesConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetControllerServiceReferencesConflict struct {
}

// IsSuccess returns true when this get controller service references conflict response has a 2xx status code
func (o *GetControllerServiceReferencesConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get controller service references conflict response has a 3xx status code
func (o *GetControllerServiceReferencesConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get controller service references conflict response has a 4xx status code
func (o *GetControllerServiceReferencesConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get controller service references conflict response has a 5xx status code
func (o *GetControllerServiceReferencesConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get controller service references conflict response a status code equal to that given
func (o *GetControllerServiceReferencesConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the get controller service references conflict response
func (o *GetControllerServiceReferencesConflict) Code() int {
	return 409
}

func (o *GetControllerServiceReferencesConflict) Error() string {
	return fmt.Sprintf("[GET /controller-services/{id}/references][%d] getControllerServiceReferencesConflict ", 409)
}

func (o *GetControllerServiceReferencesConflict) String() string {
	return fmt.Sprintf("[GET /controller-services/{id}/references][%d] getControllerServiceReferencesConflict ", 409)
}

func (o *GetControllerServiceReferencesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
