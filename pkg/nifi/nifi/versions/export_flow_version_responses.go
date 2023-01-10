// Code generated by go-swagger; DO NOT EDIT.

package versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExportFlowVersionReader is a Reader for the ExportFlowVersion structure.
type ExportFlowVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExportFlowVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExportFlowVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewExportFlowVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewExportFlowVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewExportFlowVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewExportFlowVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewExportFlowVersionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExportFlowVersionOK creates a ExportFlowVersionOK with default headers values
func NewExportFlowVersionOK() *ExportFlowVersionOK {
	return &ExportFlowVersionOK{}
}

/*
ExportFlowVersionOK describes a response with status code 200, with default header values.

successful operation
*/
type ExportFlowVersionOK struct {
	Payload string
}

// IsSuccess returns true when this export flow version o k response has a 2xx status code
func (o *ExportFlowVersionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this export flow version o k response has a 3xx status code
func (o *ExportFlowVersionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export flow version o k response has a 4xx status code
func (o *ExportFlowVersionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this export flow version o k response has a 5xx status code
func (o *ExportFlowVersionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this export flow version o k response a status code equal to that given
func (o *ExportFlowVersionOK) IsCode(code int) bool {
	return code == 200
}

func (o *ExportFlowVersionOK) Error() string {
	return fmt.Sprintf("[GET /versions/process-groups/{id}/download][%d] exportFlowVersionOK  %+v", 200, o.Payload)
}

func (o *ExportFlowVersionOK) String() string {
	return fmt.Sprintf("[GET /versions/process-groups/{id}/download][%d] exportFlowVersionOK  %+v", 200, o.Payload)
}

func (o *ExportFlowVersionOK) GetPayload() string {
	return o.Payload
}

func (o *ExportFlowVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExportFlowVersionBadRequest creates a ExportFlowVersionBadRequest with default headers values
func NewExportFlowVersionBadRequest() *ExportFlowVersionBadRequest {
	return &ExportFlowVersionBadRequest{}
}

/*
ExportFlowVersionBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type ExportFlowVersionBadRequest struct {
}

// IsSuccess returns true when this export flow version bad request response has a 2xx status code
func (o *ExportFlowVersionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this export flow version bad request response has a 3xx status code
func (o *ExportFlowVersionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export flow version bad request response has a 4xx status code
func (o *ExportFlowVersionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this export flow version bad request response has a 5xx status code
func (o *ExportFlowVersionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this export flow version bad request response a status code equal to that given
func (o *ExportFlowVersionBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ExportFlowVersionBadRequest) Error() string {
	return fmt.Sprintf("[GET /versions/process-groups/{id}/download][%d] exportFlowVersionBadRequest ", 400)
}

func (o *ExportFlowVersionBadRequest) String() string {
	return fmt.Sprintf("[GET /versions/process-groups/{id}/download][%d] exportFlowVersionBadRequest ", 400)
}

func (o *ExportFlowVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportFlowVersionUnauthorized creates a ExportFlowVersionUnauthorized with default headers values
func NewExportFlowVersionUnauthorized() *ExportFlowVersionUnauthorized {
	return &ExportFlowVersionUnauthorized{}
}

/*
ExportFlowVersionUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type ExportFlowVersionUnauthorized struct {
}

// IsSuccess returns true when this export flow version unauthorized response has a 2xx status code
func (o *ExportFlowVersionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this export flow version unauthorized response has a 3xx status code
func (o *ExportFlowVersionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export flow version unauthorized response has a 4xx status code
func (o *ExportFlowVersionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this export flow version unauthorized response has a 5xx status code
func (o *ExportFlowVersionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this export flow version unauthorized response a status code equal to that given
func (o *ExportFlowVersionUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ExportFlowVersionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /versions/process-groups/{id}/download][%d] exportFlowVersionUnauthorized ", 401)
}

func (o *ExportFlowVersionUnauthorized) String() string {
	return fmt.Sprintf("[GET /versions/process-groups/{id}/download][%d] exportFlowVersionUnauthorized ", 401)
}

func (o *ExportFlowVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportFlowVersionForbidden creates a ExportFlowVersionForbidden with default headers values
func NewExportFlowVersionForbidden() *ExportFlowVersionForbidden {
	return &ExportFlowVersionForbidden{}
}

/*
ExportFlowVersionForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type ExportFlowVersionForbidden struct {
}

// IsSuccess returns true when this export flow version forbidden response has a 2xx status code
func (o *ExportFlowVersionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this export flow version forbidden response has a 3xx status code
func (o *ExportFlowVersionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export flow version forbidden response has a 4xx status code
func (o *ExportFlowVersionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this export flow version forbidden response has a 5xx status code
func (o *ExportFlowVersionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this export flow version forbidden response a status code equal to that given
func (o *ExportFlowVersionForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ExportFlowVersionForbidden) Error() string {
	return fmt.Sprintf("[GET /versions/process-groups/{id}/download][%d] exportFlowVersionForbidden ", 403)
}

func (o *ExportFlowVersionForbidden) String() string {
	return fmt.Sprintf("[GET /versions/process-groups/{id}/download][%d] exportFlowVersionForbidden ", 403)
}

func (o *ExportFlowVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportFlowVersionNotFound creates a ExportFlowVersionNotFound with default headers values
func NewExportFlowVersionNotFound() *ExportFlowVersionNotFound {
	return &ExportFlowVersionNotFound{}
}

/*
ExportFlowVersionNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type ExportFlowVersionNotFound struct {
}

// IsSuccess returns true when this export flow version not found response has a 2xx status code
func (o *ExportFlowVersionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this export flow version not found response has a 3xx status code
func (o *ExportFlowVersionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export flow version not found response has a 4xx status code
func (o *ExportFlowVersionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this export flow version not found response has a 5xx status code
func (o *ExportFlowVersionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this export flow version not found response a status code equal to that given
func (o *ExportFlowVersionNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ExportFlowVersionNotFound) Error() string {
	return fmt.Sprintf("[GET /versions/process-groups/{id}/download][%d] exportFlowVersionNotFound ", 404)
}

func (o *ExportFlowVersionNotFound) String() string {
	return fmt.Sprintf("[GET /versions/process-groups/{id}/download][%d] exportFlowVersionNotFound ", 404)
}

func (o *ExportFlowVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportFlowVersionConflict creates a ExportFlowVersionConflict with default headers values
func NewExportFlowVersionConflict() *ExportFlowVersionConflict {
	return &ExportFlowVersionConflict{}
}

/*
ExportFlowVersionConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type ExportFlowVersionConflict struct {
}

// IsSuccess returns true when this export flow version conflict response has a 2xx status code
func (o *ExportFlowVersionConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this export flow version conflict response has a 3xx status code
func (o *ExportFlowVersionConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export flow version conflict response has a 4xx status code
func (o *ExportFlowVersionConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this export flow version conflict response has a 5xx status code
func (o *ExportFlowVersionConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this export flow version conflict response a status code equal to that given
func (o *ExportFlowVersionConflict) IsCode(code int) bool {
	return code == 409
}

func (o *ExportFlowVersionConflict) Error() string {
	return fmt.Sprintf("[GET /versions/process-groups/{id}/download][%d] exportFlowVersionConflict ", 409)
}

func (o *ExportFlowVersionConflict) String() string {
	return fmt.Sprintf("[GET /versions/process-groups/{id}/download][%d] exportFlowVersionConflict ", 409)
}

func (o *ExportFlowVersionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
