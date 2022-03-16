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

// ScheduleComponentsReader is a Reader for the ScheduleComponents structure.
type ScheduleComponentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ScheduleComponentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewScheduleComponentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewScheduleComponentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewScheduleComponentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewScheduleComponentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewScheduleComponentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewScheduleComponentsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewScheduleComponentsOK creates a ScheduleComponentsOK with default headers values
func NewScheduleComponentsOK() *ScheduleComponentsOK {
	return &ScheduleComponentsOK{}
}

/* ScheduleComponentsOK describes a response with status code 200, with default header values.

successful operation
*/
type ScheduleComponentsOK struct {
	Payload *models.ScheduleComponentsEntity
}

func (o *ScheduleComponentsOK) Error() string {
	return fmt.Sprintf("[PUT /flow/process-groups/{id}][%d] scheduleComponentsOK  %+v", 200, o.Payload)
}
func (o *ScheduleComponentsOK) GetPayload() *models.ScheduleComponentsEntity {
	return o.Payload
}

func (o *ScheduleComponentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScheduleComponentsEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewScheduleComponentsBadRequest creates a ScheduleComponentsBadRequest with default headers values
func NewScheduleComponentsBadRequest() *ScheduleComponentsBadRequest {
	return &ScheduleComponentsBadRequest{}
}

/* ScheduleComponentsBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type ScheduleComponentsBadRequest struct {
}

func (o *ScheduleComponentsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /flow/process-groups/{id}][%d] scheduleComponentsBadRequest ", 400)
}

func (o *ScheduleComponentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewScheduleComponentsUnauthorized creates a ScheduleComponentsUnauthorized with default headers values
func NewScheduleComponentsUnauthorized() *ScheduleComponentsUnauthorized {
	return &ScheduleComponentsUnauthorized{}
}

/* ScheduleComponentsUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type ScheduleComponentsUnauthorized struct {
}

func (o *ScheduleComponentsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /flow/process-groups/{id}][%d] scheduleComponentsUnauthorized ", 401)
}

func (o *ScheduleComponentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewScheduleComponentsForbidden creates a ScheduleComponentsForbidden with default headers values
func NewScheduleComponentsForbidden() *ScheduleComponentsForbidden {
	return &ScheduleComponentsForbidden{}
}

/* ScheduleComponentsForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type ScheduleComponentsForbidden struct {
}

func (o *ScheduleComponentsForbidden) Error() string {
	return fmt.Sprintf("[PUT /flow/process-groups/{id}][%d] scheduleComponentsForbidden ", 403)
}

func (o *ScheduleComponentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewScheduleComponentsNotFound creates a ScheduleComponentsNotFound with default headers values
func NewScheduleComponentsNotFound() *ScheduleComponentsNotFound {
	return &ScheduleComponentsNotFound{}
}

/* ScheduleComponentsNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type ScheduleComponentsNotFound struct {
}

func (o *ScheduleComponentsNotFound) Error() string {
	return fmt.Sprintf("[PUT /flow/process-groups/{id}][%d] scheduleComponentsNotFound ", 404)
}

func (o *ScheduleComponentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewScheduleComponentsConflict creates a ScheduleComponentsConflict with default headers values
func NewScheduleComponentsConflict() *ScheduleComponentsConflict {
	return &ScheduleComponentsConflict{}
}

/* ScheduleComponentsConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type ScheduleComponentsConflict struct {
}

func (o *ScheduleComponentsConflict) Error() string {
	return fmt.Sprintf("[PUT /flow/process-groups/{id}][%d] scheduleComponentsConflict ", 409)
}

func (o *ScheduleComponentsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}