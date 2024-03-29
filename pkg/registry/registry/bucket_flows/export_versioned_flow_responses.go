// Code generated by go-swagger; DO NOT EDIT.

package bucket_flows

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/registry/models"
)

// ExportVersionedFlowReader is a Reader for the ExportVersionedFlow structure.
type ExportVersionedFlowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExportVersionedFlowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExportVersionedFlowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewExportVersionedFlowUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewExportVersionedFlowForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewExportVersionedFlowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewExportVersionedFlowConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /buckets/{bucketId}/flows/{flowId}/versions/{versionNumber}/export] exportVersionedFlow", response, response.Code())
	}
}

// NewExportVersionedFlowOK creates a ExportVersionedFlowOK with default headers values
func NewExportVersionedFlowOK() *ExportVersionedFlowOK {
	return &ExportVersionedFlowOK{}
}

/*
ExportVersionedFlowOK describes a response with status code 200, with default header values.

successful operation
*/
type ExportVersionedFlowOK struct {
	Payload *models.VersionedFlowSnapshot
}

// IsSuccess returns true when this export versioned flow o k response has a 2xx status code
func (o *ExportVersionedFlowOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this export versioned flow o k response has a 3xx status code
func (o *ExportVersionedFlowOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export versioned flow o k response has a 4xx status code
func (o *ExportVersionedFlowOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this export versioned flow o k response has a 5xx status code
func (o *ExportVersionedFlowOK) IsServerError() bool {
	return false
}

// IsCode returns true when this export versioned flow o k response a status code equal to that given
func (o *ExportVersionedFlowOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the export versioned flow o k response
func (o *ExportVersionedFlowOK) Code() int {
	return 200
}

func (o *ExportVersionedFlowOK) Error() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}/flows/{flowId}/versions/{versionNumber}/export][%d] exportVersionedFlowOK  %+v", 200, o.Payload)
}

func (o *ExportVersionedFlowOK) String() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}/flows/{flowId}/versions/{versionNumber}/export][%d] exportVersionedFlowOK  %+v", 200, o.Payload)
}

func (o *ExportVersionedFlowOK) GetPayload() *models.VersionedFlowSnapshot {
	return o.Payload
}

func (o *ExportVersionedFlowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VersionedFlowSnapshot)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExportVersionedFlowUnauthorized creates a ExportVersionedFlowUnauthorized with default headers values
func NewExportVersionedFlowUnauthorized() *ExportVersionedFlowUnauthorized {
	return &ExportVersionedFlowUnauthorized{}
}

/*
ExportVersionedFlowUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type ExportVersionedFlowUnauthorized struct {
}

// IsSuccess returns true when this export versioned flow unauthorized response has a 2xx status code
func (o *ExportVersionedFlowUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this export versioned flow unauthorized response has a 3xx status code
func (o *ExportVersionedFlowUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export versioned flow unauthorized response has a 4xx status code
func (o *ExportVersionedFlowUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this export versioned flow unauthorized response has a 5xx status code
func (o *ExportVersionedFlowUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this export versioned flow unauthorized response a status code equal to that given
func (o *ExportVersionedFlowUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the export versioned flow unauthorized response
func (o *ExportVersionedFlowUnauthorized) Code() int {
	return 401
}

func (o *ExportVersionedFlowUnauthorized) Error() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}/flows/{flowId}/versions/{versionNumber}/export][%d] exportVersionedFlowUnauthorized ", 401)
}

func (o *ExportVersionedFlowUnauthorized) String() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}/flows/{flowId}/versions/{versionNumber}/export][%d] exportVersionedFlowUnauthorized ", 401)
}

func (o *ExportVersionedFlowUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportVersionedFlowForbidden creates a ExportVersionedFlowForbidden with default headers values
func NewExportVersionedFlowForbidden() *ExportVersionedFlowForbidden {
	return &ExportVersionedFlowForbidden{}
}

/*
ExportVersionedFlowForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type ExportVersionedFlowForbidden struct {
}

// IsSuccess returns true when this export versioned flow forbidden response has a 2xx status code
func (o *ExportVersionedFlowForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this export versioned flow forbidden response has a 3xx status code
func (o *ExportVersionedFlowForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export versioned flow forbidden response has a 4xx status code
func (o *ExportVersionedFlowForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this export versioned flow forbidden response has a 5xx status code
func (o *ExportVersionedFlowForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this export versioned flow forbidden response a status code equal to that given
func (o *ExportVersionedFlowForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the export versioned flow forbidden response
func (o *ExportVersionedFlowForbidden) Code() int {
	return 403
}

func (o *ExportVersionedFlowForbidden) Error() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}/flows/{flowId}/versions/{versionNumber}/export][%d] exportVersionedFlowForbidden ", 403)
}

func (o *ExportVersionedFlowForbidden) String() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}/flows/{flowId}/versions/{versionNumber}/export][%d] exportVersionedFlowForbidden ", 403)
}

func (o *ExportVersionedFlowForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportVersionedFlowNotFound creates a ExportVersionedFlowNotFound with default headers values
func NewExportVersionedFlowNotFound() *ExportVersionedFlowNotFound {
	return &ExportVersionedFlowNotFound{}
}

/*
ExportVersionedFlowNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type ExportVersionedFlowNotFound struct {
}

// IsSuccess returns true when this export versioned flow not found response has a 2xx status code
func (o *ExportVersionedFlowNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this export versioned flow not found response has a 3xx status code
func (o *ExportVersionedFlowNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export versioned flow not found response has a 4xx status code
func (o *ExportVersionedFlowNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this export versioned flow not found response has a 5xx status code
func (o *ExportVersionedFlowNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this export versioned flow not found response a status code equal to that given
func (o *ExportVersionedFlowNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the export versioned flow not found response
func (o *ExportVersionedFlowNotFound) Code() int {
	return 404
}

func (o *ExportVersionedFlowNotFound) Error() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}/flows/{flowId}/versions/{versionNumber}/export][%d] exportVersionedFlowNotFound ", 404)
}

func (o *ExportVersionedFlowNotFound) String() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}/flows/{flowId}/versions/{versionNumber}/export][%d] exportVersionedFlowNotFound ", 404)
}

func (o *ExportVersionedFlowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportVersionedFlowConflict creates a ExportVersionedFlowConflict with default headers values
func NewExportVersionedFlowConflict() *ExportVersionedFlowConflict {
	return &ExportVersionedFlowConflict{}
}

/*
ExportVersionedFlowConflict describes a response with status code 409, with default header values.

NiFi Registry was unable to complete the request because it assumes a server state that is not valid.
*/
type ExportVersionedFlowConflict struct {
}

// IsSuccess returns true when this export versioned flow conflict response has a 2xx status code
func (o *ExportVersionedFlowConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this export versioned flow conflict response has a 3xx status code
func (o *ExportVersionedFlowConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this export versioned flow conflict response has a 4xx status code
func (o *ExportVersionedFlowConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this export versioned flow conflict response has a 5xx status code
func (o *ExportVersionedFlowConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this export versioned flow conflict response a status code equal to that given
func (o *ExportVersionedFlowConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the export versioned flow conflict response
func (o *ExportVersionedFlowConflict) Code() int {
	return 409
}

func (o *ExportVersionedFlowConflict) Error() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}/flows/{flowId}/versions/{versionNumber}/export][%d] exportVersionedFlowConflict ", 409)
}

func (o *ExportVersionedFlowConflict) String() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}/flows/{flowId}/versions/{versionNumber}/export][%d] exportVersionedFlowConflict ", 409)
}

func (o *ExportVersionedFlowConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
