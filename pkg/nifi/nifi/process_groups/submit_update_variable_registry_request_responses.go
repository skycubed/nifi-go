// Code generated by go-swagger; DO NOT EDIT.

package process_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// SubmitUpdateVariableRegistryRequestReader is a Reader for the SubmitUpdateVariableRegistryRequest structure.
type SubmitUpdateVariableRegistryRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubmitUpdateVariableRegistryRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSubmitUpdateVariableRegistryRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSubmitUpdateVariableRegistryRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSubmitUpdateVariableRegistryRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSubmitUpdateVariableRegistryRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSubmitUpdateVariableRegistryRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewSubmitUpdateVariableRegistryRequestConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSubmitUpdateVariableRegistryRequestOK creates a SubmitUpdateVariableRegistryRequestOK with default headers values
func NewSubmitUpdateVariableRegistryRequestOK() *SubmitUpdateVariableRegistryRequestOK {
	return &SubmitUpdateVariableRegistryRequestOK{}
}

/* SubmitUpdateVariableRegistryRequestOK describes a response with status code 200, with default header values.

successful operation
*/
type SubmitUpdateVariableRegistryRequestOK struct {
	Payload *models.VariableRegistryUpdateRequestEntity
}

func (o *SubmitUpdateVariableRegistryRequestOK) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/variable-registry/update-requests][%d] submitUpdateVariableRegistryRequestOK  %+v", 200, o.Payload)
}
func (o *SubmitUpdateVariableRegistryRequestOK) GetPayload() *models.VariableRegistryUpdateRequestEntity {
	return o.Payload
}

func (o *SubmitUpdateVariableRegistryRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VariableRegistryUpdateRequestEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitUpdateVariableRegistryRequestBadRequest creates a SubmitUpdateVariableRegistryRequestBadRequest with default headers values
func NewSubmitUpdateVariableRegistryRequestBadRequest() *SubmitUpdateVariableRegistryRequestBadRequest {
	return &SubmitUpdateVariableRegistryRequestBadRequest{}
}

/* SubmitUpdateVariableRegistryRequestBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type SubmitUpdateVariableRegistryRequestBadRequest struct {
}

func (o *SubmitUpdateVariableRegistryRequestBadRequest) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/variable-registry/update-requests][%d] submitUpdateVariableRegistryRequestBadRequest ", 400)
}

func (o *SubmitUpdateVariableRegistryRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSubmitUpdateVariableRegistryRequestUnauthorized creates a SubmitUpdateVariableRegistryRequestUnauthorized with default headers values
func NewSubmitUpdateVariableRegistryRequestUnauthorized() *SubmitUpdateVariableRegistryRequestUnauthorized {
	return &SubmitUpdateVariableRegistryRequestUnauthorized{}
}

/* SubmitUpdateVariableRegistryRequestUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type SubmitUpdateVariableRegistryRequestUnauthorized struct {
}

func (o *SubmitUpdateVariableRegistryRequestUnauthorized) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/variable-registry/update-requests][%d] submitUpdateVariableRegistryRequestUnauthorized ", 401)
}

func (o *SubmitUpdateVariableRegistryRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSubmitUpdateVariableRegistryRequestForbidden creates a SubmitUpdateVariableRegistryRequestForbidden with default headers values
func NewSubmitUpdateVariableRegistryRequestForbidden() *SubmitUpdateVariableRegistryRequestForbidden {
	return &SubmitUpdateVariableRegistryRequestForbidden{}
}

/* SubmitUpdateVariableRegistryRequestForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type SubmitUpdateVariableRegistryRequestForbidden struct {
}

func (o *SubmitUpdateVariableRegistryRequestForbidden) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/variable-registry/update-requests][%d] submitUpdateVariableRegistryRequestForbidden ", 403)
}

func (o *SubmitUpdateVariableRegistryRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSubmitUpdateVariableRegistryRequestNotFound creates a SubmitUpdateVariableRegistryRequestNotFound with default headers values
func NewSubmitUpdateVariableRegistryRequestNotFound() *SubmitUpdateVariableRegistryRequestNotFound {
	return &SubmitUpdateVariableRegistryRequestNotFound{}
}

/* SubmitUpdateVariableRegistryRequestNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type SubmitUpdateVariableRegistryRequestNotFound struct {
}

func (o *SubmitUpdateVariableRegistryRequestNotFound) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/variable-registry/update-requests][%d] submitUpdateVariableRegistryRequestNotFound ", 404)
}

func (o *SubmitUpdateVariableRegistryRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSubmitUpdateVariableRegistryRequestConflict creates a SubmitUpdateVariableRegistryRequestConflict with default headers values
func NewSubmitUpdateVariableRegistryRequestConflict() *SubmitUpdateVariableRegistryRequestConflict {
	return &SubmitUpdateVariableRegistryRequestConflict{}
}

/* SubmitUpdateVariableRegistryRequestConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type SubmitUpdateVariableRegistryRequestConflict struct {
}

func (o *SubmitUpdateVariableRegistryRequestConflict) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/variable-registry/update-requests][%d] submitUpdateVariableRegistryRequestConflict ", 409)
}

func (o *SubmitUpdateVariableRegistryRequestConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
