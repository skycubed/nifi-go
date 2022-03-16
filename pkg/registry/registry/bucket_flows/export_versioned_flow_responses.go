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
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExportVersionedFlowOK creates a ExportVersionedFlowOK with default headers values
func NewExportVersionedFlowOK() *ExportVersionedFlowOK {
	return &ExportVersionedFlowOK{}
}

/* ExportVersionedFlowOK describes a response with status code 200, with default header values.

successful operation
*/
type ExportVersionedFlowOK struct {
	Payload *models.VersionedFlowSnapshot
}

func (o *ExportVersionedFlowOK) Error() string {
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

/* ExportVersionedFlowUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type ExportVersionedFlowUnauthorized struct {
}

func (o *ExportVersionedFlowUnauthorized) Error() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}/flows/{flowId}/versions/{versionNumber}/export][%d] exportVersionedFlowUnauthorized ", 401)
}

func (o *ExportVersionedFlowUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportVersionedFlowForbidden creates a ExportVersionedFlowForbidden with default headers values
func NewExportVersionedFlowForbidden() *ExportVersionedFlowForbidden {
	return &ExportVersionedFlowForbidden{}
}

/* ExportVersionedFlowForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type ExportVersionedFlowForbidden struct {
}

func (o *ExportVersionedFlowForbidden) Error() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}/flows/{flowId}/versions/{versionNumber}/export][%d] exportVersionedFlowForbidden ", 403)
}

func (o *ExportVersionedFlowForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportVersionedFlowNotFound creates a ExportVersionedFlowNotFound with default headers values
func NewExportVersionedFlowNotFound() *ExportVersionedFlowNotFound {
	return &ExportVersionedFlowNotFound{}
}

/* ExportVersionedFlowNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type ExportVersionedFlowNotFound struct {
}

func (o *ExportVersionedFlowNotFound) Error() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}/flows/{flowId}/versions/{versionNumber}/export][%d] exportVersionedFlowNotFound ", 404)
}

func (o *ExportVersionedFlowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportVersionedFlowConflict creates a ExportVersionedFlowConflict with default headers values
func NewExportVersionedFlowConflict() *ExportVersionedFlowConflict {
	return &ExportVersionedFlowConflict{}
}

/* ExportVersionedFlowConflict describes a response with status code 409, with default header values.

NiFi Registry was unable to complete the request because it assumes a server state that is not valid.
*/
type ExportVersionedFlowConflict struct {
}

func (o *ExportVersionedFlowConflict) Error() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}/flows/{flowId}/versions/{versionNumber}/export][%d] exportVersionedFlowConflict ", 409)
}

func (o *ExportVersionedFlowConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
