// Code generated by go-swagger; DO NOT EDIT.

package provenance_events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// GetProvenanceEventReader is a Reader for the GetProvenanceEvent structure.
type GetProvenanceEventReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProvenanceEventReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProvenanceEventOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProvenanceEventBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetProvenanceEventUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetProvenanceEventForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProvenanceEventNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetProvenanceEventConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProvenanceEventOK creates a GetProvenanceEventOK with default headers values
func NewGetProvenanceEventOK() *GetProvenanceEventOK {
	return &GetProvenanceEventOK{}
}

/* GetProvenanceEventOK describes a response with status code 200, with default header values.

successful operation
*/
type GetProvenanceEventOK struct {
	Payload *models.ProvenanceEventEntity
}

func (o *GetProvenanceEventOK) Error() string {
	return fmt.Sprintf("[GET /provenance-events/{id}][%d] getProvenanceEventOK  %+v", 200, o.Payload)
}
func (o *GetProvenanceEventOK) GetPayload() *models.ProvenanceEventEntity {
	return o.Payload
}

func (o *GetProvenanceEventOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProvenanceEventEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProvenanceEventBadRequest creates a GetProvenanceEventBadRequest with default headers values
func NewGetProvenanceEventBadRequest() *GetProvenanceEventBadRequest {
	return &GetProvenanceEventBadRequest{}
}

/* GetProvenanceEventBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetProvenanceEventBadRequest struct {
}

func (o *GetProvenanceEventBadRequest) Error() string {
	return fmt.Sprintf("[GET /provenance-events/{id}][%d] getProvenanceEventBadRequest ", 400)
}

func (o *GetProvenanceEventBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProvenanceEventUnauthorized creates a GetProvenanceEventUnauthorized with default headers values
func NewGetProvenanceEventUnauthorized() *GetProvenanceEventUnauthorized {
	return &GetProvenanceEventUnauthorized{}
}

/* GetProvenanceEventUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetProvenanceEventUnauthorized struct {
}

func (o *GetProvenanceEventUnauthorized) Error() string {
	return fmt.Sprintf("[GET /provenance-events/{id}][%d] getProvenanceEventUnauthorized ", 401)
}

func (o *GetProvenanceEventUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProvenanceEventForbidden creates a GetProvenanceEventForbidden with default headers values
func NewGetProvenanceEventForbidden() *GetProvenanceEventForbidden {
	return &GetProvenanceEventForbidden{}
}

/* GetProvenanceEventForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetProvenanceEventForbidden struct {
}

func (o *GetProvenanceEventForbidden) Error() string {
	return fmt.Sprintf("[GET /provenance-events/{id}][%d] getProvenanceEventForbidden ", 403)
}

func (o *GetProvenanceEventForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProvenanceEventNotFound creates a GetProvenanceEventNotFound with default headers values
func NewGetProvenanceEventNotFound() *GetProvenanceEventNotFound {
	return &GetProvenanceEventNotFound{}
}

/* GetProvenanceEventNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetProvenanceEventNotFound struct {
}

func (o *GetProvenanceEventNotFound) Error() string {
	return fmt.Sprintf("[GET /provenance-events/{id}][%d] getProvenanceEventNotFound ", 404)
}

func (o *GetProvenanceEventNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProvenanceEventConflict creates a GetProvenanceEventConflict with default headers values
func NewGetProvenanceEventConflict() *GetProvenanceEventConflict {
	return &GetProvenanceEventConflict{}
}

/* GetProvenanceEventConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetProvenanceEventConflict struct {
}

func (o *GetProvenanceEventConflict) Error() string {
	return fmt.Sprintf("[GET /provenance-events/{id}][%d] getProvenanceEventConflict ", 409)
}

func (o *GetProvenanceEventConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
