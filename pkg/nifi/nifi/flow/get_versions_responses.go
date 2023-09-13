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

// GetVersionsReader is a Reader for the GetVersions structure.
type GetVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVersionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetVersionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetVersionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVersionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetVersionsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /flow/registries/{registry-id}/buckets/{bucket-id}/flows/{flow-id}/versions] getVersions", response, response.Code())
	}
}

// NewGetVersionsOK creates a GetVersionsOK with default headers values
func NewGetVersionsOK() *GetVersionsOK {
	return &GetVersionsOK{}
}

/*
GetVersionsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetVersionsOK struct {
	Payload *models.VersionedFlowSnapshotMetadataSetEntity
}

// IsSuccess returns true when this get versions o k response has a 2xx status code
func (o *GetVersionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get versions o k response has a 3xx status code
func (o *GetVersionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get versions o k response has a 4xx status code
func (o *GetVersionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get versions o k response has a 5xx status code
func (o *GetVersionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get versions o k response a status code equal to that given
func (o *GetVersionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get versions o k response
func (o *GetVersionsOK) Code() int {
	return 200
}

func (o *GetVersionsOK) Error() string {
	return fmt.Sprintf("[GET /flow/registries/{registry-id}/buckets/{bucket-id}/flows/{flow-id}/versions][%d] getVersionsOK  %+v", 200, o.Payload)
}

func (o *GetVersionsOK) String() string {
	return fmt.Sprintf("[GET /flow/registries/{registry-id}/buckets/{bucket-id}/flows/{flow-id}/versions][%d] getVersionsOK  %+v", 200, o.Payload)
}

func (o *GetVersionsOK) GetPayload() *models.VersionedFlowSnapshotMetadataSetEntity {
	return o.Payload
}

func (o *GetVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VersionedFlowSnapshotMetadataSetEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVersionsBadRequest creates a GetVersionsBadRequest with default headers values
func NewGetVersionsBadRequest() *GetVersionsBadRequest {
	return &GetVersionsBadRequest{}
}

/*
GetVersionsBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetVersionsBadRequest struct {
}

// IsSuccess returns true when this get versions bad request response has a 2xx status code
func (o *GetVersionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get versions bad request response has a 3xx status code
func (o *GetVersionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get versions bad request response has a 4xx status code
func (o *GetVersionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get versions bad request response has a 5xx status code
func (o *GetVersionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get versions bad request response a status code equal to that given
func (o *GetVersionsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get versions bad request response
func (o *GetVersionsBadRequest) Code() int {
	return 400
}

func (o *GetVersionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /flow/registries/{registry-id}/buckets/{bucket-id}/flows/{flow-id}/versions][%d] getVersionsBadRequest ", 400)
}

func (o *GetVersionsBadRequest) String() string {
	return fmt.Sprintf("[GET /flow/registries/{registry-id}/buckets/{bucket-id}/flows/{flow-id}/versions][%d] getVersionsBadRequest ", 400)
}

func (o *GetVersionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVersionsUnauthorized creates a GetVersionsUnauthorized with default headers values
func NewGetVersionsUnauthorized() *GetVersionsUnauthorized {
	return &GetVersionsUnauthorized{}
}

/*
GetVersionsUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetVersionsUnauthorized struct {
}

// IsSuccess returns true when this get versions unauthorized response has a 2xx status code
func (o *GetVersionsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get versions unauthorized response has a 3xx status code
func (o *GetVersionsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get versions unauthorized response has a 4xx status code
func (o *GetVersionsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get versions unauthorized response has a 5xx status code
func (o *GetVersionsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get versions unauthorized response a status code equal to that given
func (o *GetVersionsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get versions unauthorized response
func (o *GetVersionsUnauthorized) Code() int {
	return 401
}

func (o *GetVersionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flow/registries/{registry-id}/buckets/{bucket-id}/flows/{flow-id}/versions][%d] getVersionsUnauthorized ", 401)
}

func (o *GetVersionsUnauthorized) String() string {
	return fmt.Sprintf("[GET /flow/registries/{registry-id}/buckets/{bucket-id}/flows/{flow-id}/versions][%d] getVersionsUnauthorized ", 401)
}

func (o *GetVersionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVersionsForbidden creates a GetVersionsForbidden with default headers values
func NewGetVersionsForbidden() *GetVersionsForbidden {
	return &GetVersionsForbidden{}
}

/*
GetVersionsForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetVersionsForbidden struct {
}

// IsSuccess returns true when this get versions forbidden response has a 2xx status code
func (o *GetVersionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get versions forbidden response has a 3xx status code
func (o *GetVersionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get versions forbidden response has a 4xx status code
func (o *GetVersionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get versions forbidden response has a 5xx status code
func (o *GetVersionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get versions forbidden response a status code equal to that given
func (o *GetVersionsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get versions forbidden response
func (o *GetVersionsForbidden) Code() int {
	return 403
}

func (o *GetVersionsForbidden) Error() string {
	return fmt.Sprintf("[GET /flow/registries/{registry-id}/buckets/{bucket-id}/flows/{flow-id}/versions][%d] getVersionsForbidden ", 403)
}

func (o *GetVersionsForbidden) String() string {
	return fmt.Sprintf("[GET /flow/registries/{registry-id}/buckets/{bucket-id}/flows/{flow-id}/versions][%d] getVersionsForbidden ", 403)
}

func (o *GetVersionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVersionsNotFound creates a GetVersionsNotFound with default headers values
func NewGetVersionsNotFound() *GetVersionsNotFound {
	return &GetVersionsNotFound{}
}

/*
GetVersionsNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetVersionsNotFound struct {
}

// IsSuccess returns true when this get versions not found response has a 2xx status code
func (o *GetVersionsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get versions not found response has a 3xx status code
func (o *GetVersionsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get versions not found response has a 4xx status code
func (o *GetVersionsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get versions not found response has a 5xx status code
func (o *GetVersionsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get versions not found response a status code equal to that given
func (o *GetVersionsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get versions not found response
func (o *GetVersionsNotFound) Code() int {
	return 404
}

func (o *GetVersionsNotFound) Error() string {
	return fmt.Sprintf("[GET /flow/registries/{registry-id}/buckets/{bucket-id}/flows/{flow-id}/versions][%d] getVersionsNotFound ", 404)
}

func (o *GetVersionsNotFound) String() string {
	return fmt.Sprintf("[GET /flow/registries/{registry-id}/buckets/{bucket-id}/flows/{flow-id}/versions][%d] getVersionsNotFound ", 404)
}

func (o *GetVersionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVersionsConflict creates a GetVersionsConflict with default headers values
func NewGetVersionsConflict() *GetVersionsConflict {
	return &GetVersionsConflict{}
}

/*
GetVersionsConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetVersionsConflict struct {
}

// IsSuccess returns true when this get versions conflict response has a 2xx status code
func (o *GetVersionsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get versions conflict response has a 3xx status code
func (o *GetVersionsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get versions conflict response has a 4xx status code
func (o *GetVersionsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get versions conflict response has a 5xx status code
func (o *GetVersionsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get versions conflict response a status code equal to that given
func (o *GetVersionsConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the get versions conflict response
func (o *GetVersionsConflict) Code() int {
	return 409
}

func (o *GetVersionsConflict) Error() string {
	return fmt.Sprintf("[GET /flow/registries/{registry-id}/buckets/{bucket-id}/flows/{flow-id}/versions][%d] getVersionsConflict ", 409)
}

func (o *GetVersionsConflict) String() string {
	return fmt.Sprintf("[GET /flow/registries/{registry-id}/buckets/{bucket-id}/flows/{flow-id}/versions][%d] getVersionsConflict ", 409)
}

func (o *GetVersionsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
