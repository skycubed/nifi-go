// Code generated by go-swagger; DO NOT EDIT.

package reporting_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// AnalyzeConfigurationReader is a Reader for the AnalyzeConfiguration structure.
type AnalyzeConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AnalyzeConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAnalyzeConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAnalyzeConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAnalyzeConfigurationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAnalyzeConfigurationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAnalyzeConfigurationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewAnalyzeConfigurationConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /reporting-tasks/{id}/config/analysis] analyzeConfiguration", response, response.Code())
	}
}

// NewAnalyzeConfigurationOK creates a AnalyzeConfigurationOK with default headers values
func NewAnalyzeConfigurationOK() *AnalyzeConfigurationOK {
	return &AnalyzeConfigurationOK{}
}

/*
AnalyzeConfigurationOK describes a response with status code 200, with default header values.

successful operation
*/
type AnalyzeConfigurationOK struct {
	Payload *models.ConfigurationAnalysisEntity
}

// IsSuccess returns true when this analyze configuration o k response has a 2xx status code
func (o *AnalyzeConfigurationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this analyze configuration o k response has a 3xx status code
func (o *AnalyzeConfigurationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this analyze configuration o k response has a 4xx status code
func (o *AnalyzeConfigurationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this analyze configuration o k response has a 5xx status code
func (o *AnalyzeConfigurationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this analyze configuration o k response a status code equal to that given
func (o *AnalyzeConfigurationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the analyze configuration o k response
func (o *AnalyzeConfigurationOK) Code() int {
	return 200
}

func (o *AnalyzeConfigurationOK) Error() string {
	return fmt.Sprintf("[POST /reporting-tasks/{id}/config/analysis][%d] analyzeConfigurationOK  %+v", 200, o.Payload)
}

func (o *AnalyzeConfigurationOK) String() string {
	return fmt.Sprintf("[POST /reporting-tasks/{id}/config/analysis][%d] analyzeConfigurationOK  %+v", 200, o.Payload)
}

func (o *AnalyzeConfigurationOK) GetPayload() *models.ConfigurationAnalysisEntity {
	return o.Payload
}

func (o *AnalyzeConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigurationAnalysisEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAnalyzeConfigurationBadRequest creates a AnalyzeConfigurationBadRequest with default headers values
func NewAnalyzeConfigurationBadRequest() *AnalyzeConfigurationBadRequest {
	return &AnalyzeConfigurationBadRequest{}
}

/*
AnalyzeConfigurationBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type AnalyzeConfigurationBadRequest struct {
}

// IsSuccess returns true when this analyze configuration bad request response has a 2xx status code
func (o *AnalyzeConfigurationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this analyze configuration bad request response has a 3xx status code
func (o *AnalyzeConfigurationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this analyze configuration bad request response has a 4xx status code
func (o *AnalyzeConfigurationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this analyze configuration bad request response has a 5xx status code
func (o *AnalyzeConfigurationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this analyze configuration bad request response a status code equal to that given
func (o *AnalyzeConfigurationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the analyze configuration bad request response
func (o *AnalyzeConfigurationBadRequest) Code() int {
	return 400
}

func (o *AnalyzeConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[POST /reporting-tasks/{id}/config/analysis][%d] analyzeConfigurationBadRequest ", 400)
}

func (o *AnalyzeConfigurationBadRequest) String() string {
	return fmt.Sprintf("[POST /reporting-tasks/{id}/config/analysis][%d] analyzeConfigurationBadRequest ", 400)
}

func (o *AnalyzeConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAnalyzeConfigurationUnauthorized creates a AnalyzeConfigurationUnauthorized with default headers values
func NewAnalyzeConfigurationUnauthorized() *AnalyzeConfigurationUnauthorized {
	return &AnalyzeConfigurationUnauthorized{}
}

/*
AnalyzeConfigurationUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type AnalyzeConfigurationUnauthorized struct {
}

// IsSuccess returns true when this analyze configuration unauthorized response has a 2xx status code
func (o *AnalyzeConfigurationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this analyze configuration unauthorized response has a 3xx status code
func (o *AnalyzeConfigurationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this analyze configuration unauthorized response has a 4xx status code
func (o *AnalyzeConfigurationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this analyze configuration unauthorized response has a 5xx status code
func (o *AnalyzeConfigurationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this analyze configuration unauthorized response a status code equal to that given
func (o *AnalyzeConfigurationUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the analyze configuration unauthorized response
func (o *AnalyzeConfigurationUnauthorized) Code() int {
	return 401
}

func (o *AnalyzeConfigurationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /reporting-tasks/{id}/config/analysis][%d] analyzeConfigurationUnauthorized ", 401)
}

func (o *AnalyzeConfigurationUnauthorized) String() string {
	return fmt.Sprintf("[POST /reporting-tasks/{id}/config/analysis][%d] analyzeConfigurationUnauthorized ", 401)
}

func (o *AnalyzeConfigurationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAnalyzeConfigurationForbidden creates a AnalyzeConfigurationForbidden with default headers values
func NewAnalyzeConfigurationForbidden() *AnalyzeConfigurationForbidden {
	return &AnalyzeConfigurationForbidden{}
}

/*
AnalyzeConfigurationForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type AnalyzeConfigurationForbidden struct {
}

// IsSuccess returns true when this analyze configuration forbidden response has a 2xx status code
func (o *AnalyzeConfigurationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this analyze configuration forbidden response has a 3xx status code
func (o *AnalyzeConfigurationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this analyze configuration forbidden response has a 4xx status code
func (o *AnalyzeConfigurationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this analyze configuration forbidden response has a 5xx status code
func (o *AnalyzeConfigurationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this analyze configuration forbidden response a status code equal to that given
func (o *AnalyzeConfigurationForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the analyze configuration forbidden response
func (o *AnalyzeConfigurationForbidden) Code() int {
	return 403
}

func (o *AnalyzeConfigurationForbidden) Error() string {
	return fmt.Sprintf("[POST /reporting-tasks/{id}/config/analysis][%d] analyzeConfigurationForbidden ", 403)
}

func (o *AnalyzeConfigurationForbidden) String() string {
	return fmt.Sprintf("[POST /reporting-tasks/{id}/config/analysis][%d] analyzeConfigurationForbidden ", 403)
}

func (o *AnalyzeConfigurationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAnalyzeConfigurationNotFound creates a AnalyzeConfigurationNotFound with default headers values
func NewAnalyzeConfigurationNotFound() *AnalyzeConfigurationNotFound {
	return &AnalyzeConfigurationNotFound{}
}

/*
AnalyzeConfigurationNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type AnalyzeConfigurationNotFound struct {
}

// IsSuccess returns true when this analyze configuration not found response has a 2xx status code
func (o *AnalyzeConfigurationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this analyze configuration not found response has a 3xx status code
func (o *AnalyzeConfigurationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this analyze configuration not found response has a 4xx status code
func (o *AnalyzeConfigurationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this analyze configuration not found response has a 5xx status code
func (o *AnalyzeConfigurationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this analyze configuration not found response a status code equal to that given
func (o *AnalyzeConfigurationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the analyze configuration not found response
func (o *AnalyzeConfigurationNotFound) Code() int {
	return 404
}

func (o *AnalyzeConfigurationNotFound) Error() string {
	return fmt.Sprintf("[POST /reporting-tasks/{id}/config/analysis][%d] analyzeConfigurationNotFound ", 404)
}

func (o *AnalyzeConfigurationNotFound) String() string {
	return fmt.Sprintf("[POST /reporting-tasks/{id}/config/analysis][%d] analyzeConfigurationNotFound ", 404)
}

func (o *AnalyzeConfigurationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAnalyzeConfigurationConflict creates a AnalyzeConfigurationConflict with default headers values
func NewAnalyzeConfigurationConflict() *AnalyzeConfigurationConflict {
	return &AnalyzeConfigurationConflict{}
}

/*
AnalyzeConfigurationConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type AnalyzeConfigurationConflict struct {
}

// IsSuccess returns true when this analyze configuration conflict response has a 2xx status code
func (o *AnalyzeConfigurationConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this analyze configuration conflict response has a 3xx status code
func (o *AnalyzeConfigurationConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this analyze configuration conflict response has a 4xx status code
func (o *AnalyzeConfigurationConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this analyze configuration conflict response has a 5xx status code
func (o *AnalyzeConfigurationConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this analyze configuration conflict response a status code equal to that given
func (o *AnalyzeConfigurationConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the analyze configuration conflict response
func (o *AnalyzeConfigurationConflict) Code() int {
	return 409
}

func (o *AnalyzeConfigurationConflict) Error() string {
	return fmt.Sprintf("[POST /reporting-tasks/{id}/config/analysis][%d] analyzeConfigurationConflict ", 409)
}

func (o *AnalyzeConfigurationConflict) String() string {
	return fmt.Sprintf("[POST /reporting-tasks/{id}/config/analysis][%d] analyzeConfigurationConflict ", 409)
}

func (o *AnalyzeConfigurationConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
