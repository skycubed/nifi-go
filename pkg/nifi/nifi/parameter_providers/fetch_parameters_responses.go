// Code generated by go-swagger; DO NOT EDIT.

package parameter_providers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// FetchParametersReader is a Reader for the FetchParameters structure.
type FetchParametersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FetchParametersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFetchParametersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewFetchParametersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewFetchParametersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFetchParametersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFetchParametersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewFetchParametersConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFetchParametersOK creates a FetchParametersOK with default headers values
func NewFetchParametersOK() *FetchParametersOK {
	return &FetchParametersOK{}
}

/*
FetchParametersOK describes a response with status code 200, with default header values.

successful operation
*/
type FetchParametersOK struct {
	Payload *models.ParameterProviderEntity
}

// IsSuccess returns true when this fetch parameters o k response has a 2xx status code
func (o *FetchParametersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this fetch parameters o k response has a 3xx status code
func (o *FetchParametersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fetch parameters o k response has a 4xx status code
func (o *FetchParametersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this fetch parameters o k response has a 5xx status code
func (o *FetchParametersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this fetch parameters o k response a status code equal to that given
func (o *FetchParametersOK) IsCode(code int) bool {
	return code == 200
}

func (o *FetchParametersOK) Error() string {
	return fmt.Sprintf("[POST /parameter-providers/{id}/parameters/fetch-requests][%d] fetchParametersOK  %+v", 200, o.Payload)
}

func (o *FetchParametersOK) String() string {
	return fmt.Sprintf("[POST /parameter-providers/{id}/parameters/fetch-requests][%d] fetchParametersOK  %+v", 200, o.Payload)
}

func (o *FetchParametersOK) GetPayload() *models.ParameterProviderEntity {
	return o.Payload
}

func (o *FetchParametersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ParameterProviderEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFetchParametersBadRequest creates a FetchParametersBadRequest with default headers values
func NewFetchParametersBadRequest() *FetchParametersBadRequest {
	return &FetchParametersBadRequest{}
}

/*
FetchParametersBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type FetchParametersBadRequest struct {
}

// IsSuccess returns true when this fetch parameters bad request response has a 2xx status code
func (o *FetchParametersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this fetch parameters bad request response has a 3xx status code
func (o *FetchParametersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fetch parameters bad request response has a 4xx status code
func (o *FetchParametersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this fetch parameters bad request response has a 5xx status code
func (o *FetchParametersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this fetch parameters bad request response a status code equal to that given
func (o *FetchParametersBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *FetchParametersBadRequest) Error() string {
	return fmt.Sprintf("[POST /parameter-providers/{id}/parameters/fetch-requests][%d] fetchParametersBadRequest ", 400)
}

func (o *FetchParametersBadRequest) String() string {
	return fmt.Sprintf("[POST /parameter-providers/{id}/parameters/fetch-requests][%d] fetchParametersBadRequest ", 400)
}

func (o *FetchParametersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFetchParametersUnauthorized creates a FetchParametersUnauthorized with default headers values
func NewFetchParametersUnauthorized() *FetchParametersUnauthorized {
	return &FetchParametersUnauthorized{}
}

/*
FetchParametersUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type FetchParametersUnauthorized struct {
}

// IsSuccess returns true when this fetch parameters unauthorized response has a 2xx status code
func (o *FetchParametersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this fetch parameters unauthorized response has a 3xx status code
func (o *FetchParametersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fetch parameters unauthorized response has a 4xx status code
func (o *FetchParametersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this fetch parameters unauthorized response has a 5xx status code
func (o *FetchParametersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this fetch parameters unauthorized response a status code equal to that given
func (o *FetchParametersUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *FetchParametersUnauthorized) Error() string {
	return fmt.Sprintf("[POST /parameter-providers/{id}/parameters/fetch-requests][%d] fetchParametersUnauthorized ", 401)
}

func (o *FetchParametersUnauthorized) String() string {
	return fmt.Sprintf("[POST /parameter-providers/{id}/parameters/fetch-requests][%d] fetchParametersUnauthorized ", 401)
}

func (o *FetchParametersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFetchParametersForbidden creates a FetchParametersForbidden with default headers values
func NewFetchParametersForbidden() *FetchParametersForbidden {
	return &FetchParametersForbidden{}
}

/*
FetchParametersForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type FetchParametersForbidden struct {
}

// IsSuccess returns true when this fetch parameters forbidden response has a 2xx status code
func (o *FetchParametersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this fetch parameters forbidden response has a 3xx status code
func (o *FetchParametersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fetch parameters forbidden response has a 4xx status code
func (o *FetchParametersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this fetch parameters forbidden response has a 5xx status code
func (o *FetchParametersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this fetch parameters forbidden response a status code equal to that given
func (o *FetchParametersForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *FetchParametersForbidden) Error() string {
	return fmt.Sprintf("[POST /parameter-providers/{id}/parameters/fetch-requests][%d] fetchParametersForbidden ", 403)
}

func (o *FetchParametersForbidden) String() string {
	return fmt.Sprintf("[POST /parameter-providers/{id}/parameters/fetch-requests][%d] fetchParametersForbidden ", 403)
}

func (o *FetchParametersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFetchParametersNotFound creates a FetchParametersNotFound with default headers values
func NewFetchParametersNotFound() *FetchParametersNotFound {
	return &FetchParametersNotFound{}
}

/*
FetchParametersNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type FetchParametersNotFound struct {
}

// IsSuccess returns true when this fetch parameters not found response has a 2xx status code
func (o *FetchParametersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this fetch parameters not found response has a 3xx status code
func (o *FetchParametersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fetch parameters not found response has a 4xx status code
func (o *FetchParametersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this fetch parameters not found response has a 5xx status code
func (o *FetchParametersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this fetch parameters not found response a status code equal to that given
func (o *FetchParametersNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *FetchParametersNotFound) Error() string {
	return fmt.Sprintf("[POST /parameter-providers/{id}/parameters/fetch-requests][%d] fetchParametersNotFound ", 404)
}

func (o *FetchParametersNotFound) String() string {
	return fmt.Sprintf("[POST /parameter-providers/{id}/parameters/fetch-requests][%d] fetchParametersNotFound ", 404)
}

func (o *FetchParametersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFetchParametersConflict creates a FetchParametersConflict with default headers values
func NewFetchParametersConflict() *FetchParametersConflict {
	return &FetchParametersConflict{}
}

/*
FetchParametersConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type FetchParametersConflict struct {
}

// IsSuccess returns true when this fetch parameters conflict response has a 2xx status code
func (o *FetchParametersConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this fetch parameters conflict response has a 3xx status code
func (o *FetchParametersConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fetch parameters conflict response has a 4xx status code
func (o *FetchParametersConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this fetch parameters conflict response has a 5xx status code
func (o *FetchParametersConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this fetch parameters conflict response a status code equal to that given
func (o *FetchParametersConflict) IsCode(code int) bool {
	return code == 409
}

func (o *FetchParametersConflict) Error() string {
	return fmt.Sprintf("[POST /parameter-providers/{id}/parameters/fetch-requests][%d] fetchParametersConflict ", 409)
}

func (o *FetchParametersConflict) String() string {
	return fmt.Sprintf("[POST /parameter-providers/{id}/parameters/fetch-requests][%d] fetchParametersConflict ", 409)
}

func (o *FetchParametersConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}