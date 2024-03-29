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

// GetFunnelsReader is a Reader for the GetFunnels structure.
type GetFunnelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFunnelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFunnelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFunnelsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFunnelsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFunnelsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFunnelsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetFunnelsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /process-groups/{id}/funnels] getFunnels", response, response.Code())
	}
}

// NewGetFunnelsOK creates a GetFunnelsOK with default headers values
func NewGetFunnelsOK() *GetFunnelsOK {
	return &GetFunnelsOK{}
}

/*
GetFunnelsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetFunnelsOK struct {
	Payload *models.FunnelsEntity
}

// IsSuccess returns true when this get funnels o k response has a 2xx status code
func (o *GetFunnelsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get funnels o k response has a 3xx status code
func (o *GetFunnelsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get funnels o k response has a 4xx status code
func (o *GetFunnelsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get funnels o k response has a 5xx status code
func (o *GetFunnelsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get funnels o k response a status code equal to that given
func (o *GetFunnelsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get funnels o k response
func (o *GetFunnelsOK) Code() int {
	return 200
}

func (o *GetFunnelsOK) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/funnels][%d] getFunnelsOK  %+v", 200, o.Payload)
}

func (o *GetFunnelsOK) String() string {
	return fmt.Sprintf("[GET /process-groups/{id}/funnels][%d] getFunnelsOK  %+v", 200, o.Payload)
}

func (o *GetFunnelsOK) GetPayload() *models.FunnelsEntity {
	return o.Payload
}

func (o *GetFunnelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FunnelsEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFunnelsBadRequest creates a GetFunnelsBadRequest with default headers values
func NewGetFunnelsBadRequest() *GetFunnelsBadRequest {
	return &GetFunnelsBadRequest{}
}

/*
GetFunnelsBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetFunnelsBadRequest struct {
}

// IsSuccess returns true when this get funnels bad request response has a 2xx status code
func (o *GetFunnelsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get funnels bad request response has a 3xx status code
func (o *GetFunnelsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get funnels bad request response has a 4xx status code
func (o *GetFunnelsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get funnels bad request response has a 5xx status code
func (o *GetFunnelsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get funnels bad request response a status code equal to that given
func (o *GetFunnelsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get funnels bad request response
func (o *GetFunnelsBadRequest) Code() int {
	return 400
}

func (o *GetFunnelsBadRequest) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/funnels][%d] getFunnelsBadRequest ", 400)
}

func (o *GetFunnelsBadRequest) String() string {
	return fmt.Sprintf("[GET /process-groups/{id}/funnels][%d] getFunnelsBadRequest ", 400)
}

func (o *GetFunnelsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFunnelsUnauthorized creates a GetFunnelsUnauthorized with default headers values
func NewGetFunnelsUnauthorized() *GetFunnelsUnauthorized {
	return &GetFunnelsUnauthorized{}
}

/*
GetFunnelsUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetFunnelsUnauthorized struct {
}

// IsSuccess returns true when this get funnels unauthorized response has a 2xx status code
func (o *GetFunnelsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get funnels unauthorized response has a 3xx status code
func (o *GetFunnelsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get funnels unauthorized response has a 4xx status code
func (o *GetFunnelsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get funnels unauthorized response has a 5xx status code
func (o *GetFunnelsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get funnels unauthorized response a status code equal to that given
func (o *GetFunnelsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get funnels unauthorized response
func (o *GetFunnelsUnauthorized) Code() int {
	return 401
}

func (o *GetFunnelsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/funnels][%d] getFunnelsUnauthorized ", 401)
}

func (o *GetFunnelsUnauthorized) String() string {
	return fmt.Sprintf("[GET /process-groups/{id}/funnels][%d] getFunnelsUnauthorized ", 401)
}

func (o *GetFunnelsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFunnelsForbidden creates a GetFunnelsForbidden with default headers values
func NewGetFunnelsForbidden() *GetFunnelsForbidden {
	return &GetFunnelsForbidden{}
}

/*
GetFunnelsForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetFunnelsForbidden struct {
}

// IsSuccess returns true when this get funnels forbidden response has a 2xx status code
func (o *GetFunnelsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get funnels forbidden response has a 3xx status code
func (o *GetFunnelsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get funnels forbidden response has a 4xx status code
func (o *GetFunnelsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get funnels forbidden response has a 5xx status code
func (o *GetFunnelsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get funnels forbidden response a status code equal to that given
func (o *GetFunnelsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get funnels forbidden response
func (o *GetFunnelsForbidden) Code() int {
	return 403
}

func (o *GetFunnelsForbidden) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/funnels][%d] getFunnelsForbidden ", 403)
}

func (o *GetFunnelsForbidden) String() string {
	return fmt.Sprintf("[GET /process-groups/{id}/funnels][%d] getFunnelsForbidden ", 403)
}

func (o *GetFunnelsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFunnelsNotFound creates a GetFunnelsNotFound with default headers values
func NewGetFunnelsNotFound() *GetFunnelsNotFound {
	return &GetFunnelsNotFound{}
}

/*
GetFunnelsNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetFunnelsNotFound struct {
}

// IsSuccess returns true when this get funnels not found response has a 2xx status code
func (o *GetFunnelsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get funnels not found response has a 3xx status code
func (o *GetFunnelsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get funnels not found response has a 4xx status code
func (o *GetFunnelsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get funnels not found response has a 5xx status code
func (o *GetFunnelsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get funnels not found response a status code equal to that given
func (o *GetFunnelsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get funnels not found response
func (o *GetFunnelsNotFound) Code() int {
	return 404
}

func (o *GetFunnelsNotFound) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/funnels][%d] getFunnelsNotFound ", 404)
}

func (o *GetFunnelsNotFound) String() string {
	return fmt.Sprintf("[GET /process-groups/{id}/funnels][%d] getFunnelsNotFound ", 404)
}

func (o *GetFunnelsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFunnelsConflict creates a GetFunnelsConflict with default headers values
func NewGetFunnelsConflict() *GetFunnelsConflict {
	return &GetFunnelsConflict{}
}

/*
GetFunnelsConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetFunnelsConflict struct {
}

// IsSuccess returns true when this get funnels conflict response has a 2xx status code
func (o *GetFunnelsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get funnels conflict response has a 3xx status code
func (o *GetFunnelsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get funnels conflict response has a 4xx status code
func (o *GetFunnelsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get funnels conflict response has a 5xx status code
func (o *GetFunnelsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get funnels conflict response a status code equal to that given
func (o *GetFunnelsConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the get funnels conflict response
func (o *GetFunnelsConflict) Code() int {
	return 409
}

func (o *GetFunnelsConflict) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/funnels][%d] getFunnelsConflict ", 409)
}

func (o *GetFunnelsConflict) String() string {
	return fmt.Sprintf("[GET /process-groups/{id}/funnels][%d] getFunnelsConflict ", 409)
}

func (o *GetFunnelsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
