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
		return nil, runtime.NewAPIError("[PUT /flow/process-groups/{id}] scheduleComponents", response, response.Code())
	}
}

// NewScheduleComponentsOK creates a ScheduleComponentsOK with default headers values
func NewScheduleComponentsOK() *ScheduleComponentsOK {
	return &ScheduleComponentsOK{}
}

/*
ScheduleComponentsOK describes a response with status code 200, with default header values.

successful operation
*/
type ScheduleComponentsOK struct {
	Payload *models.ScheduleComponentsEntity
}

// IsSuccess returns true when this schedule components o k response has a 2xx status code
func (o *ScheduleComponentsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this schedule components o k response has a 3xx status code
func (o *ScheduleComponentsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schedule components o k response has a 4xx status code
func (o *ScheduleComponentsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this schedule components o k response has a 5xx status code
func (o *ScheduleComponentsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this schedule components o k response a status code equal to that given
func (o *ScheduleComponentsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the schedule components o k response
func (o *ScheduleComponentsOK) Code() int {
	return 200
}

func (o *ScheduleComponentsOK) Error() string {
	return fmt.Sprintf("[PUT /flow/process-groups/{id}][%d] scheduleComponentsOK  %+v", 200, o.Payload)
}

func (o *ScheduleComponentsOK) String() string {
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

/*
ScheduleComponentsBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type ScheduleComponentsBadRequest struct {
}

// IsSuccess returns true when this schedule components bad request response has a 2xx status code
func (o *ScheduleComponentsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this schedule components bad request response has a 3xx status code
func (o *ScheduleComponentsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schedule components bad request response has a 4xx status code
func (o *ScheduleComponentsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this schedule components bad request response has a 5xx status code
func (o *ScheduleComponentsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this schedule components bad request response a status code equal to that given
func (o *ScheduleComponentsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the schedule components bad request response
func (o *ScheduleComponentsBadRequest) Code() int {
	return 400
}

func (o *ScheduleComponentsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /flow/process-groups/{id}][%d] scheduleComponentsBadRequest ", 400)
}

func (o *ScheduleComponentsBadRequest) String() string {
	return fmt.Sprintf("[PUT /flow/process-groups/{id}][%d] scheduleComponentsBadRequest ", 400)
}

func (o *ScheduleComponentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewScheduleComponentsUnauthorized creates a ScheduleComponentsUnauthorized with default headers values
func NewScheduleComponentsUnauthorized() *ScheduleComponentsUnauthorized {
	return &ScheduleComponentsUnauthorized{}
}

/*
ScheduleComponentsUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type ScheduleComponentsUnauthorized struct {
}

// IsSuccess returns true when this schedule components unauthorized response has a 2xx status code
func (o *ScheduleComponentsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this schedule components unauthorized response has a 3xx status code
func (o *ScheduleComponentsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schedule components unauthorized response has a 4xx status code
func (o *ScheduleComponentsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this schedule components unauthorized response has a 5xx status code
func (o *ScheduleComponentsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this schedule components unauthorized response a status code equal to that given
func (o *ScheduleComponentsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the schedule components unauthorized response
func (o *ScheduleComponentsUnauthorized) Code() int {
	return 401
}

func (o *ScheduleComponentsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /flow/process-groups/{id}][%d] scheduleComponentsUnauthorized ", 401)
}

func (o *ScheduleComponentsUnauthorized) String() string {
	return fmt.Sprintf("[PUT /flow/process-groups/{id}][%d] scheduleComponentsUnauthorized ", 401)
}

func (o *ScheduleComponentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewScheduleComponentsForbidden creates a ScheduleComponentsForbidden with default headers values
func NewScheduleComponentsForbidden() *ScheduleComponentsForbidden {
	return &ScheduleComponentsForbidden{}
}

/*
ScheduleComponentsForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type ScheduleComponentsForbidden struct {
}

// IsSuccess returns true when this schedule components forbidden response has a 2xx status code
func (o *ScheduleComponentsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this schedule components forbidden response has a 3xx status code
func (o *ScheduleComponentsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schedule components forbidden response has a 4xx status code
func (o *ScheduleComponentsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this schedule components forbidden response has a 5xx status code
func (o *ScheduleComponentsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this schedule components forbidden response a status code equal to that given
func (o *ScheduleComponentsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the schedule components forbidden response
func (o *ScheduleComponentsForbidden) Code() int {
	return 403
}

func (o *ScheduleComponentsForbidden) Error() string {
	return fmt.Sprintf("[PUT /flow/process-groups/{id}][%d] scheduleComponentsForbidden ", 403)
}

func (o *ScheduleComponentsForbidden) String() string {
	return fmt.Sprintf("[PUT /flow/process-groups/{id}][%d] scheduleComponentsForbidden ", 403)
}

func (o *ScheduleComponentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewScheduleComponentsNotFound creates a ScheduleComponentsNotFound with default headers values
func NewScheduleComponentsNotFound() *ScheduleComponentsNotFound {
	return &ScheduleComponentsNotFound{}
}

/*
ScheduleComponentsNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type ScheduleComponentsNotFound struct {
}

// IsSuccess returns true when this schedule components not found response has a 2xx status code
func (o *ScheduleComponentsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this schedule components not found response has a 3xx status code
func (o *ScheduleComponentsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schedule components not found response has a 4xx status code
func (o *ScheduleComponentsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this schedule components not found response has a 5xx status code
func (o *ScheduleComponentsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this schedule components not found response a status code equal to that given
func (o *ScheduleComponentsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the schedule components not found response
func (o *ScheduleComponentsNotFound) Code() int {
	return 404
}

func (o *ScheduleComponentsNotFound) Error() string {
	return fmt.Sprintf("[PUT /flow/process-groups/{id}][%d] scheduleComponentsNotFound ", 404)
}

func (o *ScheduleComponentsNotFound) String() string {
	return fmt.Sprintf("[PUT /flow/process-groups/{id}][%d] scheduleComponentsNotFound ", 404)
}

func (o *ScheduleComponentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewScheduleComponentsConflict creates a ScheduleComponentsConflict with default headers values
func NewScheduleComponentsConflict() *ScheduleComponentsConflict {
	return &ScheduleComponentsConflict{}
}

/*
ScheduleComponentsConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type ScheduleComponentsConflict struct {
}

// IsSuccess returns true when this schedule components conflict response has a 2xx status code
func (o *ScheduleComponentsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this schedule components conflict response has a 3xx status code
func (o *ScheduleComponentsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schedule components conflict response has a 4xx status code
func (o *ScheduleComponentsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this schedule components conflict response has a 5xx status code
func (o *ScheduleComponentsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this schedule components conflict response a status code equal to that given
func (o *ScheduleComponentsConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the schedule components conflict response
func (o *ScheduleComponentsConflict) Code() int {
	return 409
}

func (o *ScheduleComponentsConflict) Error() string {
	return fmt.Sprintf("[PUT /flow/process-groups/{id}][%d] scheduleComponentsConflict ", 409)
}

func (o *ScheduleComponentsConflict) String() string {
	return fmt.Sprintf("[PUT /flow/process-groups/{id}][%d] scheduleComponentsConflict ", 409)
}

func (o *ScheduleComponentsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
