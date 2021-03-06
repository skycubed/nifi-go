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

// GetProcessGroupStatusHistoryReader is a Reader for the GetProcessGroupStatusHistory structure.
type GetProcessGroupStatusHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProcessGroupStatusHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProcessGroupStatusHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProcessGroupStatusHistoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetProcessGroupStatusHistoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetProcessGroupStatusHistoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProcessGroupStatusHistoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetProcessGroupStatusHistoryConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProcessGroupStatusHistoryOK creates a GetProcessGroupStatusHistoryOK with default headers values
func NewGetProcessGroupStatusHistoryOK() *GetProcessGroupStatusHistoryOK {
	return &GetProcessGroupStatusHistoryOK{}
}

/* GetProcessGroupStatusHistoryOK describes a response with status code 200, with default header values.

successful operation
*/
type GetProcessGroupStatusHistoryOK struct {
	Payload *models.StatusHistoryEntity
}

func (o *GetProcessGroupStatusHistoryOK) Error() string {
	return fmt.Sprintf("[GET /flow/process-groups/{id}/status/history][%d] getProcessGroupStatusHistoryOK  %+v", 200, o.Payload)
}
func (o *GetProcessGroupStatusHistoryOK) GetPayload() *models.StatusHistoryEntity {
	return o.Payload
}

func (o *GetProcessGroupStatusHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StatusHistoryEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProcessGroupStatusHistoryBadRequest creates a GetProcessGroupStatusHistoryBadRequest with default headers values
func NewGetProcessGroupStatusHistoryBadRequest() *GetProcessGroupStatusHistoryBadRequest {
	return &GetProcessGroupStatusHistoryBadRequest{}
}

/* GetProcessGroupStatusHistoryBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetProcessGroupStatusHistoryBadRequest struct {
}

func (o *GetProcessGroupStatusHistoryBadRequest) Error() string {
	return fmt.Sprintf("[GET /flow/process-groups/{id}/status/history][%d] getProcessGroupStatusHistoryBadRequest ", 400)
}

func (o *GetProcessGroupStatusHistoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProcessGroupStatusHistoryUnauthorized creates a GetProcessGroupStatusHistoryUnauthorized with default headers values
func NewGetProcessGroupStatusHistoryUnauthorized() *GetProcessGroupStatusHistoryUnauthorized {
	return &GetProcessGroupStatusHistoryUnauthorized{}
}

/* GetProcessGroupStatusHistoryUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetProcessGroupStatusHistoryUnauthorized struct {
}

func (o *GetProcessGroupStatusHistoryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flow/process-groups/{id}/status/history][%d] getProcessGroupStatusHistoryUnauthorized ", 401)
}

func (o *GetProcessGroupStatusHistoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProcessGroupStatusHistoryForbidden creates a GetProcessGroupStatusHistoryForbidden with default headers values
func NewGetProcessGroupStatusHistoryForbidden() *GetProcessGroupStatusHistoryForbidden {
	return &GetProcessGroupStatusHistoryForbidden{}
}

/* GetProcessGroupStatusHistoryForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetProcessGroupStatusHistoryForbidden struct {
}

func (o *GetProcessGroupStatusHistoryForbidden) Error() string {
	return fmt.Sprintf("[GET /flow/process-groups/{id}/status/history][%d] getProcessGroupStatusHistoryForbidden ", 403)
}

func (o *GetProcessGroupStatusHistoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProcessGroupStatusHistoryNotFound creates a GetProcessGroupStatusHistoryNotFound with default headers values
func NewGetProcessGroupStatusHistoryNotFound() *GetProcessGroupStatusHistoryNotFound {
	return &GetProcessGroupStatusHistoryNotFound{}
}

/* GetProcessGroupStatusHistoryNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetProcessGroupStatusHistoryNotFound struct {
}

func (o *GetProcessGroupStatusHistoryNotFound) Error() string {
	return fmt.Sprintf("[GET /flow/process-groups/{id}/status/history][%d] getProcessGroupStatusHistoryNotFound ", 404)
}

func (o *GetProcessGroupStatusHistoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProcessGroupStatusHistoryConflict creates a GetProcessGroupStatusHistoryConflict with default headers values
func NewGetProcessGroupStatusHistoryConflict() *GetProcessGroupStatusHistoryConflict {
	return &GetProcessGroupStatusHistoryConflict{}
}

/* GetProcessGroupStatusHistoryConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetProcessGroupStatusHistoryConflict struct {
}

func (o *GetProcessGroupStatusHistoryConflict) Error() string {
	return fmt.Sprintf("[GET /flow/process-groups/{id}/status/history][%d] getProcessGroupStatusHistoryConflict ", 409)
}

func (o *GetProcessGroupStatusHistoryConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
