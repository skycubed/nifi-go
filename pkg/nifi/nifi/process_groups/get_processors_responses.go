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

// GetProcessorsReader is a Reader for the GetProcessors structure.
type GetProcessorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProcessorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProcessorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProcessorsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetProcessorsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetProcessorsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProcessorsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetProcessorsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /process-groups/{id}/processors] getProcessors", response, response.Code())
	}
}

// NewGetProcessorsOK creates a GetProcessorsOK with default headers values
func NewGetProcessorsOK() *GetProcessorsOK {
	return &GetProcessorsOK{}
}

/*
GetProcessorsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetProcessorsOK struct {
	Payload *models.ProcessorsEntity
}

// IsSuccess returns true when this get processors o k response has a 2xx status code
func (o *GetProcessorsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get processors o k response has a 3xx status code
func (o *GetProcessorsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get processors o k response has a 4xx status code
func (o *GetProcessorsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get processors o k response has a 5xx status code
func (o *GetProcessorsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get processors o k response a status code equal to that given
func (o *GetProcessorsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get processors o k response
func (o *GetProcessorsOK) Code() int {
	return 200
}

func (o *GetProcessorsOK) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/processors][%d] getProcessorsOK  %+v", 200, o.Payload)
}

func (o *GetProcessorsOK) String() string {
	return fmt.Sprintf("[GET /process-groups/{id}/processors][%d] getProcessorsOK  %+v", 200, o.Payload)
}

func (o *GetProcessorsOK) GetPayload() *models.ProcessorsEntity {
	return o.Payload
}

func (o *GetProcessorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProcessorsEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProcessorsBadRequest creates a GetProcessorsBadRequest with default headers values
func NewGetProcessorsBadRequest() *GetProcessorsBadRequest {
	return &GetProcessorsBadRequest{}
}

/*
GetProcessorsBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetProcessorsBadRequest struct {
}

// IsSuccess returns true when this get processors bad request response has a 2xx status code
func (o *GetProcessorsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get processors bad request response has a 3xx status code
func (o *GetProcessorsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get processors bad request response has a 4xx status code
func (o *GetProcessorsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get processors bad request response has a 5xx status code
func (o *GetProcessorsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get processors bad request response a status code equal to that given
func (o *GetProcessorsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get processors bad request response
func (o *GetProcessorsBadRequest) Code() int {
	return 400
}

func (o *GetProcessorsBadRequest) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/processors][%d] getProcessorsBadRequest ", 400)
}

func (o *GetProcessorsBadRequest) String() string {
	return fmt.Sprintf("[GET /process-groups/{id}/processors][%d] getProcessorsBadRequest ", 400)
}

func (o *GetProcessorsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProcessorsUnauthorized creates a GetProcessorsUnauthorized with default headers values
func NewGetProcessorsUnauthorized() *GetProcessorsUnauthorized {
	return &GetProcessorsUnauthorized{}
}

/*
GetProcessorsUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetProcessorsUnauthorized struct {
}

// IsSuccess returns true when this get processors unauthorized response has a 2xx status code
func (o *GetProcessorsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get processors unauthorized response has a 3xx status code
func (o *GetProcessorsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get processors unauthorized response has a 4xx status code
func (o *GetProcessorsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get processors unauthorized response has a 5xx status code
func (o *GetProcessorsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get processors unauthorized response a status code equal to that given
func (o *GetProcessorsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get processors unauthorized response
func (o *GetProcessorsUnauthorized) Code() int {
	return 401
}

func (o *GetProcessorsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/processors][%d] getProcessorsUnauthorized ", 401)
}

func (o *GetProcessorsUnauthorized) String() string {
	return fmt.Sprintf("[GET /process-groups/{id}/processors][%d] getProcessorsUnauthorized ", 401)
}

func (o *GetProcessorsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProcessorsForbidden creates a GetProcessorsForbidden with default headers values
func NewGetProcessorsForbidden() *GetProcessorsForbidden {
	return &GetProcessorsForbidden{}
}

/*
GetProcessorsForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetProcessorsForbidden struct {
}

// IsSuccess returns true when this get processors forbidden response has a 2xx status code
func (o *GetProcessorsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get processors forbidden response has a 3xx status code
func (o *GetProcessorsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get processors forbidden response has a 4xx status code
func (o *GetProcessorsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get processors forbidden response has a 5xx status code
func (o *GetProcessorsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get processors forbidden response a status code equal to that given
func (o *GetProcessorsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get processors forbidden response
func (o *GetProcessorsForbidden) Code() int {
	return 403
}

func (o *GetProcessorsForbidden) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/processors][%d] getProcessorsForbidden ", 403)
}

func (o *GetProcessorsForbidden) String() string {
	return fmt.Sprintf("[GET /process-groups/{id}/processors][%d] getProcessorsForbidden ", 403)
}

func (o *GetProcessorsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProcessorsNotFound creates a GetProcessorsNotFound with default headers values
func NewGetProcessorsNotFound() *GetProcessorsNotFound {
	return &GetProcessorsNotFound{}
}

/*
GetProcessorsNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetProcessorsNotFound struct {
}

// IsSuccess returns true when this get processors not found response has a 2xx status code
func (o *GetProcessorsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get processors not found response has a 3xx status code
func (o *GetProcessorsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get processors not found response has a 4xx status code
func (o *GetProcessorsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get processors not found response has a 5xx status code
func (o *GetProcessorsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get processors not found response a status code equal to that given
func (o *GetProcessorsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get processors not found response
func (o *GetProcessorsNotFound) Code() int {
	return 404
}

func (o *GetProcessorsNotFound) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/processors][%d] getProcessorsNotFound ", 404)
}

func (o *GetProcessorsNotFound) String() string {
	return fmt.Sprintf("[GET /process-groups/{id}/processors][%d] getProcessorsNotFound ", 404)
}

func (o *GetProcessorsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProcessorsConflict creates a GetProcessorsConflict with default headers values
func NewGetProcessorsConflict() *GetProcessorsConflict {
	return &GetProcessorsConflict{}
}

/*
GetProcessorsConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetProcessorsConflict struct {
}

// IsSuccess returns true when this get processors conflict response has a 2xx status code
func (o *GetProcessorsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get processors conflict response has a 3xx status code
func (o *GetProcessorsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get processors conflict response has a 4xx status code
func (o *GetProcessorsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get processors conflict response has a 5xx status code
func (o *GetProcessorsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get processors conflict response a status code equal to that given
func (o *GetProcessorsConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the get processors conflict response
func (o *GetProcessorsConflict) Code() int {
	return 409
}

func (o *GetProcessorsConflict) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/processors][%d] getProcessorsConflict ", 409)
}

func (o *GetProcessorsConflict) String() string {
	return fmt.Sprintf("[GET /process-groups/{id}/processors][%d] getProcessorsConflict ", 409)
}

func (o *GetProcessorsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
