// Code generated by go-swagger; DO NOT EDIT.

package processors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// SubmitProcessorVerificationRequestReader is a Reader for the SubmitProcessorVerificationRequest structure.
type SubmitProcessorVerificationRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubmitProcessorVerificationRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSubmitProcessorVerificationRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSubmitProcessorVerificationRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSubmitProcessorVerificationRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSubmitProcessorVerificationRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSubmitProcessorVerificationRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewSubmitProcessorVerificationRequestConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSubmitProcessorVerificationRequestOK creates a SubmitProcessorVerificationRequestOK with default headers values
func NewSubmitProcessorVerificationRequestOK() *SubmitProcessorVerificationRequestOK {
	return &SubmitProcessorVerificationRequestOK{}
}

/* SubmitProcessorVerificationRequestOK describes a response with status code 200, with default header values.

successful operation
*/
type SubmitProcessorVerificationRequestOK struct {
	Payload *models.VerifyConfigRequestEntity
}

func (o *SubmitProcessorVerificationRequestOK) Error() string {
	return fmt.Sprintf("[POST /processors/{id}/config/verification-requests][%d] submitProcessorVerificationRequestOK  %+v", 200, o.Payload)
}
func (o *SubmitProcessorVerificationRequestOK) GetPayload() *models.VerifyConfigRequestEntity {
	return o.Payload
}

func (o *SubmitProcessorVerificationRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VerifyConfigRequestEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitProcessorVerificationRequestBadRequest creates a SubmitProcessorVerificationRequestBadRequest with default headers values
func NewSubmitProcessorVerificationRequestBadRequest() *SubmitProcessorVerificationRequestBadRequest {
	return &SubmitProcessorVerificationRequestBadRequest{}
}

/* SubmitProcessorVerificationRequestBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type SubmitProcessorVerificationRequestBadRequest struct {
}

func (o *SubmitProcessorVerificationRequestBadRequest) Error() string {
	return fmt.Sprintf("[POST /processors/{id}/config/verification-requests][%d] submitProcessorVerificationRequestBadRequest ", 400)
}

func (o *SubmitProcessorVerificationRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSubmitProcessorVerificationRequestUnauthorized creates a SubmitProcessorVerificationRequestUnauthorized with default headers values
func NewSubmitProcessorVerificationRequestUnauthorized() *SubmitProcessorVerificationRequestUnauthorized {
	return &SubmitProcessorVerificationRequestUnauthorized{}
}

/* SubmitProcessorVerificationRequestUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type SubmitProcessorVerificationRequestUnauthorized struct {
}

func (o *SubmitProcessorVerificationRequestUnauthorized) Error() string {
	return fmt.Sprintf("[POST /processors/{id}/config/verification-requests][%d] submitProcessorVerificationRequestUnauthorized ", 401)
}

func (o *SubmitProcessorVerificationRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSubmitProcessorVerificationRequestForbidden creates a SubmitProcessorVerificationRequestForbidden with default headers values
func NewSubmitProcessorVerificationRequestForbidden() *SubmitProcessorVerificationRequestForbidden {
	return &SubmitProcessorVerificationRequestForbidden{}
}

/* SubmitProcessorVerificationRequestForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type SubmitProcessorVerificationRequestForbidden struct {
}

func (o *SubmitProcessorVerificationRequestForbidden) Error() string {
	return fmt.Sprintf("[POST /processors/{id}/config/verification-requests][%d] submitProcessorVerificationRequestForbidden ", 403)
}

func (o *SubmitProcessorVerificationRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSubmitProcessorVerificationRequestNotFound creates a SubmitProcessorVerificationRequestNotFound with default headers values
func NewSubmitProcessorVerificationRequestNotFound() *SubmitProcessorVerificationRequestNotFound {
	return &SubmitProcessorVerificationRequestNotFound{}
}

/* SubmitProcessorVerificationRequestNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type SubmitProcessorVerificationRequestNotFound struct {
}

func (o *SubmitProcessorVerificationRequestNotFound) Error() string {
	return fmt.Sprintf("[POST /processors/{id}/config/verification-requests][%d] submitProcessorVerificationRequestNotFound ", 404)
}

func (o *SubmitProcessorVerificationRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSubmitProcessorVerificationRequestConflict creates a SubmitProcessorVerificationRequestConflict with default headers values
func NewSubmitProcessorVerificationRequestConflict() *SubmitProcessorVerificationRequestConflict {
	return &SubmitProcessorVerificationRequestConflict{}
}

/* SubmitProcessorVerificationRequestConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type SubmitProcessorVerificationRequestConflict struct {
}

func (o *SubmitProcessorVerificationRequestConflict) Error() string {
	return fmt.Sprintf("[POST /processors/{id}/config/verification-requests][%d] submitProcessorVerificationRequestConflict ", 409)
}

func (o *SubmitProcessorVerificationRequestConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
