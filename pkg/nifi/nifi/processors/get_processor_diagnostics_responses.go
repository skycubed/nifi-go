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

// GetProcessorDiagnosticsReader is a Reader for the GetProcessorDiagnostics structure.
type GetProcessorDiagnosticsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProcessorDiagnosticsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProcessorDiagnosticsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProcessorDiagnosticsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetProcessorDiagnosticsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetProcessorDiagnosticsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProcessorDiagnosticsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetProcessorDiagnosticsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProcessorDiagnosticsOK creates a GetProcessorDiagnosticsOK with default headers values
func NewGetProcessorDiagnosticsOK() *GetProcessorDiagnosticsOK {
	return &GetProcessorDiagnosticsOK{}
}

/* GetProcessorDiagnosticsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetProcessorDiagnosticsOK struct {
	Payload *models.ProcessorEntity
}

func (o *GetProcessorDiagnosticsOK) Error() string {
	return fmt.Sprintf("[GET /processors/{id}/diagnostics][%d] getProcessorDiagnosticsOK  %+v", 200, o.Payload)
}
func (o *GetProcessorDiagnosticsOK) GetPayload() *models.ProcessorEntity {
	return o.Payload
}

func (o *GetProcessorDiagnosticsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProcessorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProcessorDiagnosticsBadRequest creates a GetProcessorDiagnosticsBadRequest with default headers values
func NewGetProcessorDiagnosticsBadRequest() *GetProcessorDiagnosticsBadRequest {
	return &GetProcessorDiagnosticsBadRequest{}
}

/* GetProcessorDiagnosticsBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetProcessorDiagnosticsBadRequest struct {
}

func (o *GetProcessorDiagnosticsBadRequest) Error() string {
	return fmt.Sprintf("[GET /processors/{id}/diagnostics][%d] getProcessorDiagnosticsBadRequest ", 400)
}

func (o *GetProcessorDiagnosticsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProcessorDiagnosticsUnauthorized creates a GetProcessorDiagnosticsUnauthorized with default headers values
func NewGetProcessorDiagnosticsUnauthorized() *GetProcessorDiagnosticsUnauthorized {
	return &GetProcessorDiagnosticsUnauthorized{}
}

/* GetProcessorDiagnosticsUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetProcessorDiagnosticsUnauthorized struct {
}

func (o *GetProcessorDiagnosticsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /processors/{id}/diagnostics][%d] getProcessorDiagnosticsUnauthorized ", 401)
}

func (o *GetProcessorDiagnosticsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProcessorDiagnosticsForbidden creates a GetProcessorDiagnosticsForbidden with default headers values
func NewGetProcessorDiagnosticsForbidden() *GetProcessorDiagnosticsForbidden {
	return &GetProcessorDiagnosticsForbidden{}
}

/* GetProcessorDiagnosticsForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetProcessorDiagnosticsForbidden struct {
}

func (o *GetProcessorDiagnosticsForbidden) Error() string {
	return fmt.Sprintf("[GET /processors/{id}/diagnostics][%d] getProcessorDiagnosticsForbidden ", 403)
}

func (o *GetProcessorDiagnosticsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProcessorDiagnosticsNotFound creates a GetProcessorDiagnosticsNotFound with default headers values
func NewGetProcessorDiagnosticsNotFound() *GetProcessorDiagnosticsNotFound {
	return &GetProcessorDiagnosticsNotFound{}
}

/* GetProcessorDiagnosticsNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetProcessorDiagnosticsNotFound struct {
}

func (o *GetProcessorDiagnosticsNotFound) Error() string {
	return fmt.Sprintf("[GET /processors/{id}/diagnostics][%d] getProcessorDiagnosticsNotFound ", 404)
}

func (o *GetProcessorDiagnosticsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProcessorDiagnosticsConflict creates a GetProcessorDiagnosticsConflict with default headers values
func NewGetProcessorDiagnosticsConflict() *GetProcessorDiagnosticsConflict {
	return &GetProcessorDiagnosticsConflict{}
}

/* GetProcessorDiagnosticsConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetProcessorDiagnosticsConflict struct {
}

func (o *GetProcessorDiagnosticsConflict) Error() string {
	return fmt.Sprintf("[GET /processors/{id}/diagnostics][%d] getProcessorDiagnosticsConflict ", 409)
}

func (o *GetProcessorDiagnosticsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
