// Code generated by go-swagger; DO NOT EDIT.

package provenance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// SubmitProvenanceRequestReader is a Reader for the SubmitProvenanceRequest structure.
type SubmitProvenanceRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubmitProvenanceRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSubmitProvenanceRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSubmitProvenanceRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSubmitProvenanceRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSubmitProvenanceRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewSubmitProvenanceRequestConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSubmitProvenanceRequestOK creates a SubmitProvenanceRequestOK with default headers values
func NewSubmitProvenanceRequestOK() *SubmitProvenanceRequestOK {
	return &SubmitProvenanceRequestOK{}
}

/* SubmitProvenanceRequestOK describes a response with status code 200, with default header values.

successful operation
*/
type SubmitProvenanceRequestOK struct {
	Payload *models.ProvenanceEntity
}

func (o *SubmitProvenanceRequestOK) Error() string {
	return fmt.Sprintf("[POST /provenance][%d] submitProvenanceRequestOK  %+v", 200, o.Payload)
}
func (o *SubmitProvenanceRequestOK) GetPayload() *models.ProvenanceEntity {
	return o.Payload
}

func (o *SubmitProvenanceRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProvenanceEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitProvenanceRequestBadRequest creates a SubmitProvenanceRequestBadRequest with default headers values
func NewSubmitProvenanceRequestBadRequest() *SubmitProvenanceRequestBadRequest {
	return &SubmitProvenanceRequestBadRequest{}
}

/* SubmitProvenanceRequestBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type SubmitProvenanceRequestBadRequest struct {
}

func (o *SubmitProvenanceRequestBadRequest) Error() string {
	return fmt.Sprintf("[POST /provenance][%d] submitProvenanceRequestBadRequest ", 400)
}

func (o *SubmitProvenanceRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSubmitProvenanceRequestUnauthorized creates a SubmitProvenanceRequestUnauthorized with default headers values
func NewSubmitProvenanceRequestUnauthorized() *SubmitProvenanceRequestUnauthorized {
	return &SubmitProvenanceRequestUnauthorized{}
}

/* SubmitProvenanceRequestUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type SubmitProvenanceRequestUnauthorized struct {
}

func (o *SubmitProvenanceRequestUnauthorized) Error() string {
	return fmt.Sprintf("[POST /provenance][%d] submitProvenanceRequestUnauthorized ", 401)
}

func (o *SubmitProvenanceRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSubmitProvenanceRequestForbidden creates a SubmitProvenanceRequestForbidden with default headers values
func NewSubmitProvenanceRequestForbidden() *SubmitProvenanceRequestForbidden {
	return &SubmitProvenanceRequestForbidden{}
}

/* SubmitProvenanceRequestForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type SubmitProvenanceRequestForbidden struct {
}

func (o *SubmitProvenanceRequestForbidden) Error() string {
	return fmt.Sprintf("[POST /provenance][%d] submitProvenanceRequestForbidden ", 403)
}

func (o *SubmitProvenanceRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSubmitProvenanceRequestConflict creates a SubmitProvenanceRequestConflict with default headers values
func NewSubmitProvenanceRequestConflict() *SubmitProvenanceRequestConflict {
	return &SubmitProvenanceRequestConflict{}
}

/* SubmitProvenanceRequestConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type SubmitProvenanceRequestConflict struct {
}

func (o *SubmitProvenanceRequestConflict) Error() string {
	return fmt.Sprintf("[POST /provenance][%d] submitProvenanceRequestConflict ", 409)
}

func (o *SubmitProvenanceRequestConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}