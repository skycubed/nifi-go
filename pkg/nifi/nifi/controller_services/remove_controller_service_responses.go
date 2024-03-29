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

// RemoveControllerServiceReader is a Reader for the RemoveControllerService structure.
type RemoveControllerServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveControllerServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveControllerServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveControllerServiceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRemoveControllerServiceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRemoveControllerServiceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRemoveControllerServiceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewRemoveControllerServiceConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /controller-services/{id}] removeControllerService", response, response.Code())
	}
}

// NewRemoveControllerServiceOK creates a RemoveControllerServiceOK with default headers values
func NewRemoveControllerServiceOK() *RemoveControllerServiceOK {
	return &RemoveControllerServiceOK{}
}

/*
RemoveControllerServiceOK describes a response with status code 200, with default header values.

successful operation
*/
type RemoveControllerServiceOK struct {
	Payload *models.ControllerServiceEntity
}

// IsSuccess returns true when this remove controller service o k response has a 2xx status code
func (o *RemoveControllerServiceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this remove controller service o k response has a 3xx status code
func (o *RemoveControllerServiceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove controller service o k response has a 4xx status code
func (o *RemoveControllerServiceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove controller service o k response has a 5xx status code
func (o *RemoveControllerServiceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this remove controller service o k response a status code equal to that given
func (o *RemoveControllerServiceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the remove controller service o k response
func (o *RemoveControllerServiceOK) Code() int {
	return 200
}

func (o *RemoveControllerServiceOK) Error() string {
	return fmt.Sprintf("[DELETE /controller-services/{id}][%d] removeControllerServiceOK  %+v", 200, o.Payload)
}

func (o *RemoveControllerServiceOK) String() string {
	return fmt.Sprintf("[DELETE /controller-services/{id}][%d] removeControllerServiceOK  %+v", 200, o.Payload)
}

func (o *RemoveControllerServiceOK) GetPayload() *models.ControllerServiceEntity {
	return o.Payload
}

func (o *RemoveControllerServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ControllerServiceEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveControllerServiceBadRequest creates a RemoveControllerServiceBadRequest with default headers values
func NewRemoveControllerServiceBadRequest() *RemoveControllerServiceBadRequest {
	return &RemoveControllerServiceBadRequest{}
}

/*
RemoveControllerServiceBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type RemoveControllerServiceBadRequest struct {
}

// IsSuccess returns true when this remove controller service bad request response has a 2xx status code
func (o *RemoveControllerServiceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove controller service bad request response has a 3xx status code
func (o *RemoveControllerServiceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove controller service bad request response has a 4xx status code
func (o *RemoveControllerServiceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove controller service bad request response has a 5xx status code
func (o *RemoveControllerServiceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this remove controller service bad request response a status code equal to that given
func (o *RemoveControllerServiceBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the remove controller service bad request response
func (o *RemoveControllerServiceBadRequest) Code() int {
	return 400
}

func (o *RemoveControllerServiceBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /controller-services/{id}][%d] removeControllerServiceBadRequest ", 400)
}

func (o *RemoveControllerServiceBadRequest) String() string {
	return fmt.Sprintf("[DELETE /controller-services/{id}][%d] removeControllerServiceBadRequest ", 400)
}

func (o *RemoveControllerServiceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveControllerServiceUnauthorized creates a RemoveControllerServiceUnauthorized with default headers values
func NewRemoveControllerServiceUnauthorized() *RemoveControllerServiceUnauthorized {
	return &RemoveControllerServiceUnauthorized{}
}

/*
RemoveControllerServiceUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type RemoveControllerServiceUnauthorized struct {
}

// IsSuccess returns true when this remove controller service unauthorized response has a 2xx status code
func (o *RemoveControllerServiceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove controller service unauthorized response has a 3xx status code
func (o *RemoveControllerServiceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove controller service unauthorized response has a 4xx status code
func (o *RemoveControllerServiceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove controller service unauthorized response has a 5xx status code
func (o *RemoveControllerServiceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this remove controller service unauthorized response a status code equal to that given
func (o *RemoveControllerServiceUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the remove controller service unauthorized response
func (o *RemoveControllerServiceUnauthorized) Code() int {
	return 401
}

func (o *RemoveControllerServiceUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /controller-services/{id}][%d] removeControllerServiceUnauthorized ", 401)
}

func (o *RemoveControllerServiceUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /controller-services/{id}][%d] removeControllerServiceUnauthorized ", 401)
}

func (o *RemoveControllerServiceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveControllerServiceForbidden creates a RemoveControllerServiceForbidden with default headers values
func NewRemoveControllerServiceForbidden() *RemoveControllerServiceForbidden {
	return &RemoveControllerServiceForbidden{}
}

/*
RemoveControllerServiceForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type RemoveControllerServiceForbidden struct {
}

// IsSuccess returns true when this remove controller service forbidden response has a 2xx status code
func (o *RemoveControllerServiceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove controller service forbidden response has a 3xx status code
func (o *RemoveControllerServiceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove controller service forbidden response has a 4xx status code
func (o *RemoveControllerServiceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove controller service forbidden response has a 5xx status code
func (o *RemoveControllerServiceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this remove controller service forbidden response a status code equal to that given
func (o *RemoveControllerServiceForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the remove controller service forbidden response
func (o *RemoveControllerServiceForbidden) Code() int {
	return 403
}

func (o *RemoveControllerServiceForbidden) Error() string {
	return fmt.Sprintf("[DELETE /controller-services/{id}][%d] removeControllerServiceForbidden ", 403)
}

func (o *RemoveControllerServiceForbidden) String() string {
	return fmt.Sprintf("[DELETE /controller-services/{id}][%d] removeControllerServiceForbidden ", 403)
}

func (o *RemoveControllerServiceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveControllerServiceNotFound creates a RemoveControllerServiceNotFound with default headers values
func NewRemoveControllerServiceNotFound() *RemoveControllerServiceNotFound {
	return &RemoveControllerServiceNotFound{}
}

/*
RemoveControllerServiceNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type RemoveControllerServiceNotFound struct {
}

// IsSuccess returns true when this remove controller service not found response has a 2xx status code
func (o *RemoveControllerServiceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove controller service not found response has a 3xx status code
func (o *RemoveControllerServiceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove controller service not found response has a 4xx status code
func (o *RemoveControllerServiceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove controller service not found response has a 5xx status code
func (o *RemoveControllerServiceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this remove controller service not found response a status code equal to that given
func (o *RemoveControllerServiceNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the remove controller service not found response
func (o *RemoveControllerServiceNotFound) Code() int {
	return 404
}

func (o *RemoveControllerServiceNotFound) Error() string {
	return fmt.Sprintf("[DELETE /controller-services/{id}][%d] removeControllerServiceNotFound ", 404)
}

func (o *RemoveControllerServiceNotFound) String() string {
	return fmt.Sprintf("[DELETE /controller-services/{id}][%d] removeControllerServiceNotFound ", 404)
}

func (o *RemoveControllerServiceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveControllerServiceConflict creates a RemoveControllerServiceConflict with default headers values
func NewRemoveControllerServiceConflict() *RemoveControllerServiceConflict {
	return &RemoveControllerServiceConflict{}
}

/*
RemoveControllerServiceConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type RemoveControllerServiceConflict struct {
}

// IsSuccess returns true when this remove controller service conflict response has a 2xx status code
func (o *RemoveControllerServiceConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove controller service conflict response has a 3xx status code
func (o *RemoveControllerServiceConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove controller service conflict response has a 4xx status code
func (o *RemoveControllerServiceConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove controller service conflict response has a 5xx status code
func (o *RemoveControllerServiceConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this remove controller service conflict response a status code equal to that given
func (o *RemoveControllerServiceConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the remove controller service conflict response
func (o *RemoveControllerServiceConflict) Code() int {
	return 409
}

func (o *RemoveControllerServiceConflict) Error() string {
	return fmt.Sprintf("[DELETE /controller-services/{id}][%d] removeControllerServiceConflict ", 409)
}

func (o *RemoveControllerServiceConflict) String() string {
	return fmt.Sprintf("[DELETE /controller-services/{id}][%d] removeControllerServiceConflict ", 409)
}

func (o *RemoveControllerServiceConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
