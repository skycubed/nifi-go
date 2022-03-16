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

// GlobalGetFlowVersionsReader is a Reader for the GlobalGetFlowVersions structure.
type GlobalGetFlowVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GlobalGetFlowVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGlobalGetFlowVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGlobalGetFlowVersionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGlobalGetFlowVersionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGlobalGetFlowVersionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGlobalGetFlowVersionsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGlobalGetFlowVersionsOK creates a GlobalGetFlowVersionsOK with default headers values
func NewGlobalGetFlowVersionsOK() *GlobalGetFlowVersionsOK {
	return &GlobalGetFlowVersionsOK{}
}

/* GlobalGetFlowVersionsOK describes a response with status code 200, with default header values.

successful operation
*/
type GlobalGetFlowVersionsOK struct {
	Payload []*models.VersionedFlowSnapshotMetadata
}

func (o *GlobalGetFlowVersionsOK) Error() string {
	return fmt.Sprintf("[GET /flows/{flowId}/versions][%d] globalGetFlowVersionsOK  %+v", 200, o.Payload)
}
func (o *GlobalGetFlowVersionsOK) GetPayload() []*models.VersionedFlowSnapshotMetadata {
	return o.Payload
}

func (o *GlobalGetFlowVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGlobalGetFlowVersionsUnauthorized creates a GlobalGetFlowVersionsUnauthorized with default headers values
func NewGlobalGetFlowVersionsUnauthorized() *GlobalGetFlowVersionsUnauthorized {
	return &GlobalGetFlowVersionsUnauthorized{}
}

/* GlobalGetFlowVersionsUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GlobalGetFlowVersionsUnauthorized struct {
}

func (o *GlobalGetFlowVersionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flows/{flowId}/versions][%d] globalGetFlowVersionsUnauthorized ", 401)
}

func (o *GlobalGetFlowVersionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGlobalGetFlowVersionsForbidden creates a GlobalGetFlowVersionsForbidden with default headers values
func NewGlobalGetFlowVersionsForbidden() *GlobalGetFlowVersionsForbidden {
	return &GlobalGetFlowVersionsForbidden{}
}

/* GlobalGetFlowVersionsForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GlobalGetFlowVersionsForbidden struct {
}

func (o *GlobalGetFlowVersionsForbidden) Error() string {
	return fmt.Sprintf("[GET /flows/{flowId}/versions][%d] globalGetFlowVersionsForbidden ", 403)
}

func (o *GlobalGetFlowVersionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGlobalGetFlowVersionsNotFound creates a GlobalGetFlowVersionsNotFound with default headers values
func NewGlobalGetFlowVersionsNotFound() *GlobalGetFlowVersionsNotFound {
	return &GlobalGetFlowVersionsNotFound{}
}

/* GlobalGetFlowVersionsNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GlobalGetFlowVersionsNotFound struct {
}

func (o *GlobalGetFlowVersionsNotFound) Error() string {
	return fmt.Sprintf("[GET /flows/{flowId}/versions][%d] globalGetFlowVersionsNotFound ", 404)
}

func (o *GlobalGetFlowVersionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGlobalGetFlowVersionsConflict creates a GlobalGetFlowVersionsConflict with default headers values
func NewGlobalGetFlowVersionsConflict() *GlobalGetFlowVersionsConflict {
	return &GlobalGetFlowVersionsConflict{}
}

/* GlobalGetFlowVersionsConflict describes a response with status code 409, with default header values.

NiFi Registry was unable to complete the request because it assumes a server state that is not valid.
*/
type GlobalGetFlowVersionsConflict struct {
}

func (o *GlobalGetFlowVersionsConflict) Error() string {
	return fmt.Sprintf("[GET /flows/{flowId}/versions][%d] globalGetFlowVersionsConflict ", 409)
}

func (o *GlobalGetFlowVersionsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}