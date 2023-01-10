// Code generated by go-swagger; DO NOT EDIT.

package versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// SaveToFlowRegistryReader is a Reader for the SaveToFlowRegistry structure.
type SaveToFlowRegistryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SaveToFlowRegistryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSaveToFlowRegistryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSaveToFlowRegistryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSaveToFlowRegistryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSaveToFlowRegistryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSaveToFlowRegistryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewSaveToFlowRegistryConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSaveToFlowRegistryOK creates a SaveToFlowRegistryOK with default headers values
func NewSaveToFlowRegistryOK() *SaveToFlowRegistryOK {
	return &SaveToFlowRegistryOK{}
}

/*
SaveToFlowRegistryOK describes a response with status code 200, with default header values.

successful operation
*/
type SaveToFlowRegistryOK struct {
	Payload *models.VersionControlInformationEntity
}

// IsSuccess returns true when this save to flow registry o k response has a 2xx status code
func (o *SaveToFlowRegistryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this save to flow registry o k response has a 3xx status code
func (o *SaveToFlowRegistryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this save to flow registry o k response has a 4xx status code
func (o *SaveToFlowRegistryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this save to flow registry o k response has a 5xx status code
func (o *SaveToFlowRegistryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this save to flow registry o k response a status code equal to that given
func (o *SaveToFlowRegistryOK) IsCode(code int) bool {
	return code == 200
}

func (o *SaveToFlowRegistryOK) Error() string {
	return fmt.Sprintf("[POST /versions/process-groups/{id}][%d] saveToFlowRegistryOK  %+v", 200, o.Payload)
}

func (o *SaveToFlowRegistryOK) String() string {
	return fmt.Sprintf("[POST /versions/process-groups/{id}][%d] saveToFlowRegistryOK  %+v", 200, o.Payload)
}

func (o *SaveToFlowRegistryOK) GetPayload() *models.VersionControlInformationEntity {
	return o.Payload
}

func (o *SaveToFlowRegistryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VersionControlInformationEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSaveToFlowRegistryBadRequest creates a SaveToFlowRegistryBadRequest with default headers values
func NewSaveToFlowRegistryBadRequest() *SaveToFlowRegistryBadRequest {
	return &SaveToFlowRegistryBadRequest{}
}

/*
SaveToFlowRegistryBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type SaveToFlowRegistryBadRequest struct {
}

// IsSuccess returns true when this save to flow registry bad request response has a 2xx status code
func (o *SaveToFlowRegistryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this save to flow registry bad request response has a 3xx status code
func (o *SaveToFlowRegistryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this save to flow registry bad request response has a 4xx status code
func (o *SaveToFlowRegistryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this save to flow registry bad request response has a 5xx status code
func (o *SaveToFlowRegistryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this save to flow registry bad request response a status code equal to that given
func (o *SaveToFlowRegistryBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SaveToFlowRegistryBadRequest) Error() string {
	return fmt.Sprintf("[POST /versions/process-groups/{id}][%d] saveToFlowRegistryBadRequest ", 400)
}

func (o *SaveToFlowRegistryBadRequest) String() string {
	return fmt.Sprintf("[POST /versions/process-groups/{id}][%d] saveToFlowRegistryBadRequest ", 400)
}

func (o *SaveToFlowRegistryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSaveToFlowRegistryUnauthorized creates a SaveToFlowRegistryUnauthorized with default headers values
func NewSaveToFlowRegistryUnauthorized() *SaveToFlowRegistryUnauthorized {
	return &SaveToFlowRegistryUnauthorized{}
}

/*
SaveToFlowRegistryUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type SaveToFlowRegistryUnauthorized struct {
}

// IsSuccess returns true when this save to flow registry unauthorized response has a 2xx status code
func (o *SaveToFlowRegistryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this save to flow registry unauthorized response has a 3xx status code
func (o *SaveToFlowRegistryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this save to flow registry unauthorized response has a 4xx status code
func (o *SaveToFlowRegistryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this save to flow registry unauthorized response has a 5xx status code
func (o *SaveToFlowRegistryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this save to flow registry unauthorized response a status code equal to that given
func (o *SaveToFlowRegistryUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *SaveToFlowRegistryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /versions/process-groups/{id}][%d] saveToFlowRegistryUnauthorized ", 401)
}

func (o *SaveToFlowRegistryUnauthorized) String() string {
	return fmt.Sprintf("[POST /versions/process-groups/{id}][%d] saveToFlowRegistryUnauthorized ", 401)
}

func (o *SaveToFlowRegistryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSaveToFlowRegistryForbidden creates a SaveToFlowRegistryForbidden with default headers values
func NewSaveToFlowRegistryForbidden() *SaveToFlowRegistryForbidden {
	return &SaveToFlowRegistryForbidden{}
}

/*
SaveToFlowRegistryForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type SaveToFlowRegistryForbidden struct {
}

// IsSuccess returns true when this save to flow registry forbidden response has a 2xx status code
func (o *SaveToFlowRegistryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this save to flow registry forbidden response has a 3xx status code
func (o *SaveToFlowRegistryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this save to flow registry forbidden response has a 4xx status code
func (o *SaveToFlowRegistryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this save to flow registry forbidden response has a 5xx status code
func (o *SaveToFlowRegistryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this save to flow registry forbidden response a status code equal to that given
func (o *SaveToFlowRegistryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SaveToFlowRegistryForbidden) Error() string {
	return fmt.Sprintf("[POST /versions/process-groups/{id}][%d] saveToFlowRegistryForbidden ", 403)
}

func (o *SaveToFlowRegistryForbidden) String() string {
	return fmt.Sprintf("[POST /versions/process-groups/{id}][%d] saveToFlowRegistryForbidden ", 403)
}

func (o *SaveToFlowRegistryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSaveToFlowRegistryNotFound creates a SaveToFlowRegistryNotFound with default headers values
func NewSaveToFlowRegistryNotFound() *SaveToFlowRegistryNotFound {
	return &SaveToFlowRegistryNotFound{}
}

/*
SaveToFlowRegistryNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type SaveToFlowRegistryNotFound struct {
}

// IsSuccess returns true when this save to flow registry not found response has a 2xx status code
func (o *SaveToFlowRegistryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this save to flow registry not found response has a 3xx status code
func (o *SaveToFlowRegistryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this save to flow registry not found response has a 4xx status code
func (o *SaveToFlowRegistryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this save to flow registry not found response has a 5xx status code
func (o *SaveToFlowRegistryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this save to flow registry not found response a status code equal to that given
func (o *SaveToFlowRegistryNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SaveToFlowRegistryNotFound) Error() string {
	return fmt.Sprintf("[POST /versions/process-groups/{id}][%d] saveToFlowRegistryNotFound ", 404)
}

func (o *SaveToFlowRegistryNotFound) String() string {
	return fmt.Sprintf("[POST /versions/process-groups/{id}][%d] saveToFlowRegistryNotFound ", 404)
}

func (o *SaveToFlowRegistryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSaveToFlowRegistryConflict creates a SaveToFlowRegistryConflict with default headers values
func NewSaveToFlowRegistryConflict() *SaveToFlowRegistryConflict {
	return &SaveToFlowRegistryConflict{}
}

/*
SaveToFlowRegistryConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type SaveToFlowRegistryConflict struct {
}

// IsSuccess returns true when this save to flow registry conflict response has a 2xx status code
func (o *SaveToFlowRegistryConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this save to flow registry conflict response has a 3xx status code
func (o *SaveToFlowRegistryConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this save to flow registry conflict response has a 4xx status code
func (o *SaveToFlowRegistryConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this save to flow registry conflict response has a 5xx status code
func (o *SaveToFlowRegistryConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this save to flow registry conflict response a status code equal to that given
func (o *SaveToFlowRegistryConflict) IsCode(code int) bool {
	return code == 409
}

func (o *SaveToFlowRegistryConflict) Error() string {
	return fmt.Sprintf("[POST /versions/process-groups/{id}][%d] saveToFlowRegistryConflict ", 409)
}

func (o *SaveToFlowRegistryConflict) String() string {
	return fmt.Sprintf("[POST /versions/process-groups/{id}][%d] saveToFlowRegistryConflict ", 409)
}

func (o *SaveToFlowRegistryConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
