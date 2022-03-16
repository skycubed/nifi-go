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

// UpdateFlowReader is a Reader for the UpdateFlow structure.
type UpdateFlowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateFlowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateFlowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateFlowBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateFlowUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateFlowForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateFlowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateFlowConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateFlowOK creates a UpdateFlowOK with default headers values
func NewUpdateFlowOK() *UpdateFlowOK {
	return &UpdateFlowOK{}
}

/* UpdateFlowOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateFlowOK struct {
	Payload *models.VersionedFlow
}

func (o *UpdateFlowOK) Error() string {
	return fmt.Sprintf("[PUT /buckets/{bucketId}/flows/{flowId}][%d] updateFlowOK  %+v", 200, o.Payload)
}
func (o *UpdateFlowOK) GetPayload() *models.VersionedFlow {
	return o.Payload
}

func (o *UpdateFlowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VersionedFlow)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFlowBadRequest creates a UpdateFlowBadRequest with default headers values
func NewUpdateFlowBadRequest() *UpdateFlowBadRequest {
	return &UpdateFlowBadRequest{}
}

/* UpdateFlowBadRequest describes a response with status code 400, with default header values.

NiFi Registry was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type UpdateFlowBadRequest struct {
}

func (o *UpdateFlowBadRequest) Error() string {
	return fmt.Sprintf("[PUT /buckets/{bucketId}/flows/{flowId}][%d] updateFlowBadRequest ", 400)
}

func (o *UpdateFlowBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateFlowUnauthorized creates a UpdateFlowUnauthorized with default headers values
func NewUpdateFlowUnauthorized() *UpdateFlowUnauthorized {
	return &UpdateFlowUnauthorized{}
}

/* UpdateFlowUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type UpdateFlowUnauthorized struct {
}

func (o *UpdateFlowUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /buckets/{bucketId}/flows/{flowId}][%d] updateFlowUnauthorized ", 401)
}

func (o *UpdateFlowUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateFlowForbidden creates a UpdateFlowForbidden with default headers values
func NewUpdateFlowForbidden() *UpdateFlowForbidden {
	return &UpdateFlowForbidden{}
}

/* UpdateFlowForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type UpdateFlowForbidden struct {
}

func (o *UpdateFlowForbidden) Error() string {
	return fmt.Sprintf("[PUT /buckets/{bucketId}/flows/{flowId}][%d] updateFlowForbidden ", 403)
}

func (o *UpdateFlowForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateFlowNotFound creates a UpdateFlowNotFound with default headers values
func NewUpdateFlowNotFound() *UpdateFlowNotFound {
	return &UpdateFlowNotFound{}
}

/* UpdateFlowNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type UpdateFlowNotFound struct {
}

func (o *UpdateFlowNotFound) Error() string {
	return fmt.Sprintf("[PUT /buckets/{bucketId}/flows/{flowId}][%d] updateFlowNotFound ", 404)
}

func (o *UpdateFlowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateFlowConflict creates a UpdateFlowConflict with default headers values
func NewUpdateFlowConflict() *UpdateFlowConflict {
	return &UpdateFlowConflict{}
}

/* UpdateFlowConflict describes a response with status code 409, with default header values.

NiFi Registry was unable to complete the request because it assumes a server state that is not valid.
*/
type UpdateFlowConflict struct {
}

func (o *UpdateFlowConflict) Error() string {
	return fmt.Sprintf("[PUT /buckets/{bucketId}/flows/{flowId}][%d] updateFlowConflict ", 409)
}

func (o *UpdateFlowConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}