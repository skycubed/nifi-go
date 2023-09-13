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

// CreateFlowReader is a Reader for the CreateFlow structure.
type CreateFlowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateFlowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateFlowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateFlowBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateFlowUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateFlowForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateFlowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateFlowConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /buckets/{bucketId}/flows] createFlow", response, response.Code())
	}
}

// NewCreateFlowOK creates a CreateFlowOK with default headers values
func NewCreateFlowOK() *CreateFlowOK {
	return &CreateFlowOK{}
}

/*
CreateFlowOK describes a response with status code 200, with default header values.

successful operation
*/
type CreateFlowOK struct {
	Payload *models.VersionedFlow
}

// IsSuccess returns true when this create flow o k response has a 2xx status code
func (o *CreateFlowOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create flow o k response has a 3xx status code
func (o *CreateFlowOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create flow o k response has a 4xx status code
func (o *CreateFlowOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create flow o k response has a 5xx status code
func (o *CreateFlowOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create flow o k response a status code equal to that given
func (o *CreateFlowOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create flow o k response
func (o *CreateFlowOK) Code() int {
	return 200
}

func (o *CreateFlowOK) Error() string {
	return fmt.Sprintf("[POST /buckets/{bucketId}/flows][%d] createFlowOK  %+v", 200, o.Payload)
}

func (o *CreateFlowOK) String() string {
	return fmt.Sprintf("[POST /buckets/{bucketId}/flows][%d] createFlowOK  %+v", 200, o.Payload)
}

func (o *CreateFlowOK) GetPayload() *models.VersionedFlow {
	return o.Payload
}

func (o *CreateFlowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VersionedFlow)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFlowBadRequest creates a CreateFlowBadRequest with default headers values
func NewCreateFlowBadRequest() *CreateFlowBadRequest {
	return &CreateFlowBadRequest{}
}

/*
CreateFlowBadRequest describes a response with status code 400, with default header values.

NiFi Registry was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type CreateFlowBadRequest struct {
}

// IsSuccess returns true when this create flow bad request response has a 2xx status code
func (o *CreateFlowBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create flow bad request response has a 3xx status code
func (o *CreateFlowBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create flow bad request response has a 4xx status code
func (o *CreateFlowBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create flow bad request response has a 5xx status code
func (o *CreateFlowBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create flow bad request response a status code equal to that given
func (o *CreateFlowBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create flow bad request response
func (o *CreateFlowBadRequest) Code() int {
	return 400
}

func (o *CreateFlowBadRequest) Error() string {
	return fmt.Sprintf("[POST /buckets/{bucketId}/flows][%d] createFlowBadRequest ", 400)
}

func (o *CreateFlowBadRequest) String() string {
	return fmt.Sprintf("[POST /buckets/{bucketId}/flows][%d] createFlowBadRequest ", 400)
}

func (o *CreateFlowBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateFlowUnauthorized creates a CreateFlowUnauthorized with default headers values
func NewCreateFlowUnauthorized() *CreateFlowUnauthorized {
	return &CreateFlowUnauthorized{}
}

/*
CreateFlowUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type CreateFlowUnauthorized struct {
}

// IsSuccess returns true when this create flow unauthorized response has a 2xx status code
func (o *CreateFlowUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create flow unauthorized response has a 3xx status code
func (o *CreateFlowUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create flow unauthorized response has a 4xx status code
func (o *CreateFlowUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create flow unauthorized response has a 5xx status code
func (o *CreateFlowUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create flow unauthorized response a status code equal to that given
func (o *CreateFlowUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create flow unauthorized response
func (o *CreateFlowUnauthorized) Code() int {
	return 401
}

func (o *CreateFlowUnauthorized) Error() string {
	return fmt.Sprintf("[POST /buckets/{bucketId}/flows][%d] createFlowUnauthorized ", 401)
}

func (o *CreateFlowUnauthorized) String() string {
	return fmt.Sprintf("[POST /buckets/{bucketId}/flows][%d] createFlowUnauthorized ", 401)
}

func (o *CreateFlowUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateFlowForbidden creates a CreateFlowForbidden with default headers values
func NewCreateFlowForbidden() *CreateFlowForbidden {
	return &CreateFlowForbidden{}
}

/*
CreateFlowForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type CreateFlowForbidden struct {
}

// IsSuccess returns true when this create flow forbidden response has a 2xx status code
func (o *CreateFlowForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create flow forbidden response has a 3xx status code
func (o *CreateFlowForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create flow forbidden response has a 4xx status code
func (o *CreateFlowForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create flow forbidden response has a 5xx status code
func (o *CreateFlowForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create flow forbidden response a status code equal to that given
func (o *CreateFlowForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create flow forbidden response
func (o *CreateFlowForbidden) Code() int {
	return 403
}

func (o *CreateFlowForbidden) Error() string {
	return fmt.Sprintf("[POST /buckets/{bucketId}/flows][%d] createFlowForbidden ", 403)
}

func (o *CreateFlowForbidden) String() string {
	return fmt.Sprintf("[POST /buckets/{bucketId}/flows][%d] createFlowForbidden ", 403)
}

func (o *CreateFlowForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateFlowNotFound creates a CreateFlowNotFound with default headers values
func NewCreateFlowNotFound() *CreateFlowNotFound {
	return &CreateFlowNotFound{}
}

/*
CreateFlowNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type CreateFlowNotFound struct {
}

// IsSuccess returns true when this create flow not found response has a 2xx status code
func (o *CreateFlowNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create flow not found response has a 3xx status code
func (o *CreateFlowNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create flow not found response has a 4xx status code
func (o *CreateFlowNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create flow not found response has a 5xx status code
func (o *CreateFlowNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create flow not found response a status code equal to that given
func (o *CreateFlowNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create flow not found response
func (o *CreateFlowNotFound) Code() int {
	return 404
}

func (o *CreateFlowNotFound) Error() string {
	return fmt.Sprintf("[POST /buckets/{bucketId}/flows][%d] createFlowNotFound ", 404)
}

func (o *CreateFlowNotFound) String() string {
	return fmt.Sprintf("[POST /buckets/{bucketId}/flows][%d] createFlowNotFound ", 404)
}

func (o *CreateFlowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateFlowConflict creates a CreateFlowConflict with default headers values
func NewCreateFlowConflict() *CreateFlowConflict {
	return &CreateFlowConflict{}
}

/*
CreateFlowConflict describes a response with status code 409, with default header values.

NiFi Registry was unable to complete the request because it assumes a server state that is not valid.
*/
type CreateFlowConflict struct {
}

// IsSuccess returns true when this create flow conflict response has a 2xx status code
func (o *CreateFlowConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create flow conflict response has a 3xx status code
func (o *CreateFlowConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create flow conflict response has a 4xx status code
func (o *CreateFlowConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create flow conflict response has a 5xx status code
func (o *CreateFlowConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create flow conflict response a status code equal to that given
func (o *CreateFlowConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create flow conflict response
func (o *CreateFlowConflict) Code() int {
	return 409
}

func (o *CreateFlowConflict) Error() string {
	return fmt.Sprintf("[POST /buckets/{bucketId}/flows][%d] createFlowConflict ", 409)
}

func (o *CreateFlowConflict) String() string {
	return fmt.Sprintf("[POST /buckets/{bucketId}/flows][%d] createFlowConflict ", 409)
}

func (o *CreateFlowConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
