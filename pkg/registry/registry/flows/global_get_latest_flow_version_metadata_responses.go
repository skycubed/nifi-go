// Code generated by go-swagger; DO NOT EDIT.

package flows

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/registry/models"
)

// GlobalGetLatestFlowVersionMetadataReader is a Reader for the GlobalGetLatestFlowVersionMetadata structure.
type GlobalGetLatestFlowVersionMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GlobalGetLatestFlowVersionMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGlobalGetLatestFlowVersionMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGlobalGetLatestFlowVersionMetadataUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGlobalGetLatestFlowVersionMetadataForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGlobalGetLatestFlowVersionMetadataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGlobalGetLatestFlowVersionMetadataConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /flows/{flowId}/versions/latest/metadata] globalGetLatestFlowVersionMetadata", response, response.Code())
	}
}

// NewGlobalGetLatestFlowVersionMetadataOK creates a GlobalGetLatestFlowVersionMetadataOK with default headers values
func NewGlobalGetLatestFlowVersionMetadataOK() *GlobalGetLatestFlowVersionMetadataOK {
	return &GlobalGetLatestFlowVersionMetadataOK{}
}

/*
GlobalGetLatestFlowVersionMetadataOK describes a response with status code 200, with default header values.

successful operation
*/
type GlobalGetLatestFlowVersionMetadataOK struct {
	Payload *models.VersionedFlowSnapshotMetadata
}

// IsSuccess returns true when this global get latest flow version metadata o k response has a 2xx status code
func (o *GlobalGetLatestFlowVersionMetadataOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this global get latest flow version metadata o k response has a 3xx status code
func (o *GlobalGetLatestFlowVersionMetadataOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this global get latest flow version metadata o k response has a 4xx status code
func (o *GlobalGetLatestFlowVersionMetadataOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this global get latest flow version metadata o k response has a 5xx status code
func (o *GlobalGetLatestFlowVersionMetadataOK) IsServerError() bool {
	return false
}

// IsCode returns true when this global get latest flow version metadata o k response a status code equal to that given
func (o *GlobalGetLatestFlowVersionMetadataOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the global get latest flow version metadata o k response
func (o *GlobalGetLatestFlowVersionMetadataOK) Code() int {
	return 200
}

func (o *GlobalGetLatestFlowVersionMetadataOK) Error() string {
	return fmt.Sprintf("[GET /flows/{flowId}/versions/latest/metadata][%d] globalGetLatestFlowVersionMetadataOK  %+v", 200, o.Payload)
}

func (o *GlobalGetLatestFlowVersionMetadataOK) String() string {
	return fmt.Sprintf("[GET /flows/{flowId}/versions/latest/metadata][%d] globalGetLatestFlowVersionMetadataOK  %+v", 200, o.Payload)
}

func (o *GlobalGetLatestFlowVersionMetadataOK) GetPayload() *models.VersionedFlowSnapshotMetadata {
	return o.Payload
}

func (o *GlobalGetLatestFlowVersionMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VersionedFlowSnapshotMetadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGlobalGetLatestFlowVersionMetadataUnauthorized creates a GlobalGetLatestFlowVersionMetadataUnauthorized with default headers values
func NewGlobalGetLatestFlowVersionMetadataUnauthorized() *GlobalGetLatestFlowVersionMetadataUnauthorized {
	return &GlobalGetLatestFlowVersionMetadataUnauthorized{}
}

/*
GlobalGetLatestFlowVersionMetadataUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GlobalGetLatestFlowVersionMetadataUnauthorized struct {
}

// IsSuccess returns true when this global get latest flow version metadata unauthorized response has a 2xx status code
func (o *GlobalGetLatestFlowVersionMetadataUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this global get latest flow version metadata unauthorized response has a 3xx status code
func (o *GlobalGetLatestFlowVersionMetadataUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this global get latest flow version metadata unauthorized response has a 4xx status code
func (o *GlobalGetLatestFlowVersionMetadataUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this global get latest flow version metadata unauthorized response has a 5xx status code
func (o *GlobalGetLatestFlowVersionMetadataUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this global get latest flow version metadata unauthorized response a status code equal to that given
func (o *GlobalGetLatestFlowVersionMetadataUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the global get latest flow version metadata unauthorized response
func (o *GlobalGetLatestFlowVersionMetadataUnauthorized) Code() int {
	return 401
}

func (o *GlobalGetLatestFlowVersionMetadataUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flows/{flowId}/versions/latest/metadata][%d] globalGetLatestFlowVersionMetadataUnauthorized ", 401)
}

func (o *GlobalGetLatestFlowVersionMetadataUnauthorized) String() string {
	return fmt.Sprintf("[GET /flows/{flowId}/versions/latest/metadata][%d] globalGetLatestFlowVersionMetadataUnauthorized ", 401)
}

func (o *GlobalGetLatestFlowVersionMetadataUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGlobalGetLatestFlowVersionMetadataForbidden creates a GlobalGetLatestFlowVersionMetadataForbidden with default headers values
func NewGlobalGetLatestFlowVersionMetadataForbidden() *GlobalGetLatestFlowVersionMetadataForbidden {
	return &GlobalGetLatestFlowVersionMetadataForbidden{}
}

/*
GlobalGetLatestFlowVersionMetadataForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GlobalGetLatestFlowVersionMetadataForbidden struct {
}

// IsSuccess returns true when this global get latest flow version metadata forbidden response has a 2xx status code
func (o *GlobalGetLatestFlowVersionMetadataForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this global get latest flow version metadata forbidden response has a 3xx status code
func (o *GlobalGetLatestFlowVersionMetadataForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this global get latest flow version metadata forbidden response has a 4xx status code
func (o *GlobalGetLatestFlowVersionMetadataForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this global get latest flow version metadata forbidden response has a 5xx status code
func (o *GlobalGetLatestFlowVersionMetadataForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this global get latest flow version metadata forbidden response a status code equal to that given
func (o *GlobalGetLatestFlowVersionMetadataForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the global get latest flow version metadata forbidden response
func (o *GlobalGetLatestFlowVersionMetadataForbidden) Code() int {
	return 403
}

func (o *GlobalGetLatestFlowVersionMetadataForbidden) Error() string {
	return fmt.Sprintf("[GET /flows/{flowId}/versions/latest/metadata][%d] globalGetLatestFlowVersionMetadataForbidden ", 403)
}

func (o *GlobalGetLatestFlowVersionMetadataForbidden) String() string {
	return fmt.Sprintf("[GET /flows/{flowId}/versions/latest/metadata][%d] globalGetLatestFlowVersionMetadataForbidden ", 403)
}

func (o *GlobalGetLatestFlowVersionMetadataForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGlobalGetLatestFlowVersionMetadataNotFound creates a GlobalGetLatestFlowVersionMetadataNotFound with default headers values
func NewGlobalGetLatestFlowVersionMetadataNotFound() *GlobalGetLatestFlowVersionMetadataNotFound {
	return &GlobalGetLatestFlowVersionMetadataNotFound{}
}

/*
GlobalGetLatestFlowVersionMetadataNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GlobalGetLatestFlowVersionMetadataNotFound struct {
}

// IsSuccess returns true when this global get latest flow version metadata not found response has a 2xx status code
func (o *GlobalGetLatestFlowVersionMetadataNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this global get latest flow version metadata not found response has a 3xx status code
func (o *GlobalGetLatestFlowVersionMetadataNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this global get latest flow version metadata not found response has a 4xx status code
func (o *GlobalGetLatestFlowVersionMetadataNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this global get latest flow version metadata not found response has a 5xx status code
func (o *GlobalGetLatestFlowVersionMetadataNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this global get latest flow version metadata not found response a status code equal to that given
func (o *GlobalGetLatestFlowVersionMetadataNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the global get latest flow version metadata not found response
func (o *GlobalGetLatestFlowVersionMetadataNotFound) Code() int {
	return 404
}

func (o *GlobalGetLatestFlowVersionMetadataNotFound) Error() string {
	return fmt.Sprintf("[GET /flows/{flowId}/versions/latest/metadata][%d] globalGetLatestFlowVersionMetadataNotFound ", 404)
}

func (o *GlobalGetLatestFlowVersionMetadataNotFound) String() string {
	return fmt.Sprintf("[GET /flows/{flowId}/versions/latest/metadata][%d] globalGetLatestFlowVersionMetadataNotFound ", 404)
}

func (o *GlobalGetLatestFlowVersionMetadataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGlobalGetLatestFlowVersionMetadataConflict creates a GlobalGetLatestFlowVersionMetadataConflict with default headers values
func NewGlobalGetLatestFlowVersionMetadataConflict() *GlobalGetLatestFlowVersionMetadataConflict {
	return &GlobalGetLatestFlowVersionMetadataConflict{}
}

/*
GlobalGetLatestFlowVersionMetadataConflict describes a response with status code 409, with default header values.

NiFi Registry was unable to complete the request because it assumes a server state that is not valid.
*/
type GlobalGetLatestFlowVersionMetadataConflict struct {
}

// IsSuccess returns true when this global get latest flow version metadata conflict response has a 2xx status code
func (o *GlobalGetLatestFlowVersionMetadataConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this global get latest flow version metadata conflict response has a 3xx status code
func (o *GlobalGetLatestFlowVersionMetadataConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this global get latest flow version metadata conflict response has a 4xx status code
func (o *GlobalGetLatestFlowVersionMetadataConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this global get latest flow version metadata conflict response has a 5xx status code
func (o *GlobalGetLatestFlowVersionMetadataConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this global get latest flow version metadata conflict response a status code equal to that given
func (o *GlobalGetLatestFlowVersionMetadataConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the global get latest flow version metadata conflict response
func (o *GlobalGetLatestFlowVersionMetadataConflict) Code() int {
	return 409
}

func (o *GlobalGetLatestFlowVersionMetadataConflict) Error() string {
	return fmt.Sprintf("[GET /flows/{flowId}/versions/latest/metadata][%d] globalGetLatestFlowVersionMetadataConflict ", 409)
}

func (o *GlobalGetLatestFlowVersionMetadataConflict) String() string {
	return fmt.Sprintf("[GET /flows/{flowId}/versions/latest/metadata][%d] globalGetLatestFlowVersionMetadataConflict ", 409)
}

func (o *GlobalGetLatestFlowVersionMetadataConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
