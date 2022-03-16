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

// GetFlowVersionReader is a Reader for the GetFlowVersion structure.
type GetFlowVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFlowVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFlowVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFlowVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFlowVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFlowVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFlowVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetFlowVersionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFlowVersionOK creates a GetFlowVersionOK with default headers values
func NewGetFlowVersionOK() *GetFlowVersionOK {
	return &GetFlowVersionOK{}
}

/* GetFlowVersionOK describes a response with status code 200, with default header values.

successful operation
*/
type GetFlowVersionOK struct {
	Payload *models.VersionedFlowSnapshot
}

func (o *GetFlowVersionOK) Error() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}/flows/{flowId}/versions/{versionNumber}][%d] getFlowVersionOK  %+v", 200, o.Payload)
}
func (o *GetFlowVersionOK) GetPayload() *models.VersionedFlowSnapshot {
	return o.Payload
}

func (o *GetFlowVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VersionedFlowSnapshot)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowVersionBadRequest creates a GetFlowVersionBadRequest with default headers values
func NewGetFlowVersionBadRequest() *GetFlowVersionBadRequest {
	return &GetFlowVersionBadRequest{}
}

/* GetFlowVersionBadRequest describes a response with status code 400, with default header values.

NiFi Registry was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetFlowVersionBadRequest struct {
}

func (o *GetFlowVersionBadRequest) Error() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}/flows/{flowId}/versions/{versionNumber}][%d] getFlowVersionBadRequest ", 400)
}

func (o *GetFlowVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFlowVersionUnauthorized creates a GetFlowVersionUnauthorized with default headers values
func NewGetFlowVersionUnauthorized() *GetFlowVersionUnauthorized {
	return &GetFlowVersionUnauthorized{}
}

/* GetFlowVersionUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetFlowVersionUnauthorized struct {
}

func (o *GetFlowVersionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}/flows/{flowId}/versions/{versionNumber}][%d] getFlowVersionUnauthorized ", 401)
}

func (o *GetFlowVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFlowVersionForbidden creates a GetFlowVersionForbidden with default headers values
func NewGetFlowVersionForbidden() *GetFlowVersionForbidden {
	return &GetFlowVersionForbidden{}
}

/* GetFlowVersionForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetFlowVersionForbidden struct {
}

func (o *GetFlowVersionForbidden) Error() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}/flows/{flowId}/versions/{versionNumber}][%d] getFlowVersionForbidden ", 403)
}

func (o *GetFlowVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFlowVersionNotFound creates a GetFlowVersionNotFound with default headers values
func NewGetFlowVersionNotFound() *GetFlowVersionNotFound {
	return &GetFlowVersionNotFound{}
}

/* GetFlowVersionNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetFlowVersionNotFound struct {
}

func (o *GetFlowVersionNotFound) Error() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}/flows/{flowId}/versions/{versionNumber}][%d] getFlowVersionNotFound ", 404)
}

func (o *GetFlowVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFlowVersionConflict creates a GetFlowVersionConflict with default headers values
func NewGetFlowVersionConflict() *GetFlowVersionConflict {
	return &GetFlowVersionConflict{}
}

/* GetFlowVersionConflict describes a response with status code 409, with default header values.

NiFi Registry was unable to complete the request because it assumes a server state that is not valid.
*/
type GetFlowVersionConflict struct {
}

func (o *GetFlowVersionConflict) Error() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}/flows/{flowId}/versions/{versionNumber}][%d] getFlowVersionConflict ", 409)
}

func (o *GetFlowVersionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}