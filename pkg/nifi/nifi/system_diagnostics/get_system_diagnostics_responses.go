// Code generated by go-swagger; DO NOT EDIT.

package system_diagnostics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// GetSystemDiagnosticsReader is a Reader for the GetSystemDiagnostics structure.
type GetSystemDiagnosticsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSystemDiagnosticsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSystemDiagnosticsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSystemDiagnosticsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSystemDiagnosticsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /system-diagnostics] getSystemDiagnostics", response, response.Code())
	}
}

// NewGetSystemDiagnosticsOK creates a GetSystemDiagnosticsOK with default headers values
func NewGetSystemDiagnosticsOK() *GetSystemDiagnosticsOK {
	return &GetSystemDiagnosticsOK{}
}

/*
GetSystemDiagnosticsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetSystemDiagnosticsOK struct {
	Payload *models.SystemDiagnosticsEntity
}

// IsSuccess returns true when this get system diagnostics o k response has a 2xx status code
func (o *GetSystemDiagnosticsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get system diagnostics o k response has a 3xx status code
func (o *GetSystemDiagnosticsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get system diagnostics o k response has a 4xx status code
func (o *GetSystemDiagnosticsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get system diagnostics o k response has a 5xx status code
func (o *GetSystemDiagnosticsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get system diagnostics o k response a status code equal to that given
func (o *GetSystemDiagnosticsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get system diagnostics o k response
func (o *GetSystemDiagnosticsOK) Code() int {
	return 200
}

func (o *GetSystemDiagnosticsOK) Error() string {
	return fmt.Sprintf("[GET /system-diagnostics][%d] getSystemDiagnosticsOK  %+v", 200, o.Payload)
}

func (o *GetSystemDiagnosticsOK) String() string {
	return fmt.Sprintf("[GET /system-diagnostics][%d] getSystemDiagnosticsOK  %+v", 200, o.Payload)
}

func (o *GetSystemDiagnosticsOK) GetPayload() *models.SystemDiagnosticsEntity {
	return o.Payload
}

func (o *GetSystemDiagnosticsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemDiagnosticsEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSystemDiagnosticsUnauthorized creates a GetSystemDiagnosticsUnauthorized with default headers values
func NewGetSystemDiagnosticsUnauthorized() *GetSystemDiagnosticsUnauthorized {
	return &GetSystemDiagnosticsUnauthorized{}
}

/*
GetSystemDiagnosticsUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetSystemDiagnosticsUnauthorized struct {
}

// IsSuccess returns true when this get system diagnostics unauthorized response has a 2xx status code
func (o *GetSystemDiagnosticsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get system diagnostics unauthorized response has a 3xx status code
func (o *GetSystemDiagnosticsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get system diagnostics unauthorized response has a 4xx status code
func (o *GetSystemDiagnosticsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get system diagnostics unauthorized response has a 5xx status code
func (o *GetSystemDiagnosticsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get system diagnostics unauthorized response a status code equal to that given
func (o *GetSystemDiagnosticsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get system diagnostics unauthorized response
func (o *GetSystemDiagnosticsUnauthorized) Code() int {
	return 401
}

func (o *GetSystemDiagnosticsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /system-diagnostics][%d] getSystemDiagnosticsUnauthorized ", 401)
}

func (o *GetSystemDiagnosticsUnauthorized) String() string {
	return fmt.Sprintf("[GET /system-diagnostics][%d] getSystemDiagnosticsUnauthorized ", 401)
}

func (o *GetSystemDiagnosticsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSystemDiagnosticsForbidden creates a GetSystemDiagnosticsForbidden with default headers values
func NewGetSystemDiagnosticsForbidden() *GetSystemDiagnosticsForbidden {
	return &GetSystemDiagnosticsForbidden{}
}

/*
GetSystemDiagnosticsForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetSystemDiagnosticsForbidden struct {
}

// IsSuccess returns true when this get system diagnostics forbidden response has a 2xx status code
func (o *GetSystemDiagnosticsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get system diagnostics forbidden response has a 3xx status code
func (o *GetSystemDiagnosticsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get system diagnostics forbidden response has a 4xx status code
func (o *GetSystemDiagnosticsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get system diagnostics forbidden response has a 5xx status code
func (o *GetSystemDiagnosticsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get system diagnostics forbidden response a status code equal to that given
func (o *GetSystemDiagnosticsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get system diagnostics forbidden response
func (o *GetSystemDiagnosticsForbidden) Code() int {
	return 403
}

func (o *GetSystemDiagnosticsForbidden) Error() string {
	return fmt.Sprintf("[GET /system-diagnostics][%d] getSystemDiagnosticsForbidden ", 403)
}

func (o *GetSystemDiagnosticsForbidden) String() string {
	return fmt.Sprintf("[GET /system-diagnostics][%d] getSystemDiagnosticsForbidden ", 403)
}

func (o *GetSystemDiagnosticsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
