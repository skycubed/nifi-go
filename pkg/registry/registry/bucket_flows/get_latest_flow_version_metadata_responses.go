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

// GetLatestFlowVersionMetadataReader is a Reader for the GetLatestFlowVersionMetadata structure.
type GetLatestFlowVersionMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLatestFlowVersionMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLatestFlowVersionMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetLatestFlowVersionMetadataUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLatestFlowVersionMetadataForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLatestFlowVersionMetadataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetLatestFlowVersionMetadataConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLatestFlowVersionMetadataOK creates a GetLatestFlowVersionMetadataOK with default headers values
func NewGetLatestFlowVersionMetadataOK() *GetLatestFlowVersionMetadataOK {
	return &GetLatestFlowVersionMetadataOK{}
}

/*
GetLatestFlowVersionMetadataOK describes a response with status code 200, with default header values.

successful operation
*/
type GetLatestFlowVersionMetadataOK struct {
	Payload *models.VersionedFlowSnapshotMetadata
}

// IsSuccess returns true when this get latest flow version metadata o k response has a 2xx status code
func (o *GetLatestFlowVersionMetadataOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get latest flow version metadata o k response has a 3xx status code
func (o *GetLatestFlowVersionMetadataOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get latest flow version metadata o k response has a 4xx status code
func (o *GetLatestFlowVersionMetadataOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get latest flow version metadata o k response has a 5xx status code
func (o *GetLatestFlowVersionMetadataOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get latest flow version metadata o k response a status code equal to that given
func (o *GetLatestFlowVersionMetadataOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetLatestFlowVersionMetadataOK) Error() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}/flows/{flowId}/versions/latest/metadata][%d] getLatestFlowVersionMetadataOK  %+v", 200, o.Payload)
}

func (o *GetLatestFlowVersionMetadataOK) String() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}/flows/{flowId}/versions/latest/metadata][%d] getLatestFlowVersionMetadataOK  %+v", 200, o.Payload)
}

func (o *GetLatestFlowVersionMetadataOK) GetPayload() *models.VersionedFlowSnapshotMetadata {
	return o.Payload
}

func (o *GetLatestFlowVersionMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VersionedFlowSnapshotMetadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLatestFlowVersionMetadataUnauthorized creates a GetLatestFlowVersionMetadataUnauthorized with default headers values
func NewGetLatestFlowVersionMetadataUnauthorized() *GetLatestFlowVersionMetadataUnauthorized {
	return &GetLatestFlowVersionMetadataUnauthorized{}
}

/*
GetLatestFlowVersionMetadataUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetLatestFlowVersionMetadataUnauthorized struct {
}

// IsSuccess returns true when this get latest flow version metadata unauthorized response has a 2xx status code
func (o *GetLatestFlowVersionMetadataUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get latest flow version metadata unauthorized response has a 3xx status code
func (o *GetLatestFlowVersionMetadataUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get latest flow version metadata unauthorized response has a 4xx status code
func (o *GetLatestFlowVersionMetadataUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get latest flow version metadata unauthorized response has a 5xx status code
func (o *GetLatestFlowVersionMetadataUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get latest flow version metadata unauthorized response a status code equal to that given
func (o *GetLatestFlowVersionMetadataUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetLatestFlowVersionMetadataUnauthorized) Error() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}/flows/{flowId}/versions/latest/metadata][%d] getLatestFlowVersionMetadataUnauthorized ", 401)
}

func (o *GetLatestFlowVersionMetadataUnauthorized) String() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}/flows/{flowId}/versions/latest/metadata][%d] getLatestFlowVersionMetadataUnauthorized ", 401)
}

func (o *GetLatestFlowVersionMetadataUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLatestFlowVersionMetadataForbidden creates a GetLatestFlowVersionMetadataForbidden with default headers values
func NewGetLatestFlowVersionMetadataForbidden() *GetLatestFlowVersionMetadataForbidden {
	return &GetLatestFlowVersionMetadataForbidden{}
}

/*
GetLatestFlowVersionMetadataForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetLatestFlowVersionMetadataForbidden struct {
}

// IsSuccess returns true when this get latest flow version metadata forbidden response has a 2xx status code
func (o *GetLatestFlowVersionMetadataForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get latest flow version metadata forbidden response has a 3xx status code
func (o *GetLatestFlowVersionMetadataForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get latest flow version metadata forbidden response has a 4xx status code
func (o *GetLatestFlowVersionMetadataForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get latest flow version metadata forbidden response has a 5xx status code
func (o *GetLatestFlowVersionMetadataForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get latest flow version metadata forbidden response a status code equal to that given
func (o *GetLatestFlowVersionMetadataForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetLatestFlowVersionMetadataForbidden) Error() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}/flows/{flowId}/versions/latest/metadata][%d] getLatestFlowVersionMetadataForbidden ", 403)
}

func (o *GetLatestFlowVersionMetadataForbidden) String() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}/flows/{flowId}/versions/latest/metadata][%d] getLatestFlowVersionMetadataForbidden ", 403)
}

func (o *GetLatestFlowVersionMetadataForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLatestFlowVersionMetadataNotFound creates a GetLatestFlowVersionMetadataNotFound with default headers values
func NewGetLatestFlowVersionMetadataNotFound() *GetLatestFlowVersionMetadataNotFound {
	return &GetLatestFlowVersionMetadataNotFound{}
}

/*
GetLatestFlowVersionMetadataNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetLatestFlowVersionMetadataNotFound struct {
}

// IsSuccess returns true when this get latest flow version metadata not found response has a 2xx status code
func (o *GetLatestFlowVersionMetadataNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get latest flow version metadata not found response has a 3xx status code
func (o *GetLatestFlowVersionMetadataNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get latest flow version metadata not found response has a 4xx status code
func (o *GetLatestFlowVersionMetadataNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get latest flow version metadata not found response has a 5xx status code
func (o *GetLatestFlowVersionMetadataNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get latest flow version metadata not found response a status code equal to that given
func (o *GetLatestFlowVersionMetadataNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetLatestFlowVersionMetadataNotFound) Error() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}/flows/{flowId}/versions/latest/metadata][%d] getLatestFlowVersionMetadataNotFound ", 404)
}

func (o *GetLatestFlowVersionMetadataNotFound) String() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}/flows/{flowId}/versions/latest/metadata][%d] getLatestFlowVersionMetadataNotFound ", 404)
}

func (o *GetLatestFlowVersionMetadataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLatestFlowVersionMetadataConflict creates a GetLatestFlowVersionMetadataConflict with default headers values
func NewGetLatestFlowVersionMetadataConflict() *GetLatestFlowVersionMetadataConflict {
	return &GetLatestFlowVersionMetadataConflict{}
}

/*
GetLatestFlowVersionMetadataConflict describes a response with status code 409, with default header values.

NiFi Registry was unable to complete the request because it assumes a server state that is not valid.
*/
type GetLatestFlowVersionMetadataConflict struct {
}

// IsSuccess returns true when this get latest flow version metadata conflict response has a 2xx status code
func (o *GetLatestFlowVersionMetadataConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get latest flow version metadata conflict response has a 3xx status code
func (o *GetLatestFlowVersionMetadataConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get latest flow version metadata conflict response has a 4xx status code
func (o *GetLatestFlowVersionMetadataConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get latest flow version metadata conflict response has a 5xx status code
func (o *GetLatestFlowVersionMetadataConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get latest flow version metadata conflict response a status code equal to that given
func (o *GetLatestFlowVersionMetadataConflict) IsCode(code int) bool {
	return code == 409
}

func (o *GetLatestFlowVersionMetadataConflict) Error() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}/flows/{flowId}/versions/latest/metadata][%d] getLatestFlowVersionMetadataConflict ", 409)
}

func (o *GetLatestFlowVersionMetadataConflict) String() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}/flows/{flowId}/versions/latest/metadata][%d] getLatestFlowVersionMetadataConflict ", 409)
}

func (o *GetLatestFlowVersionMetadataConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
