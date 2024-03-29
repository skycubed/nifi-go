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

// ImportTemplateReader is a Reader for the ImportTemplate structure.
type ImportTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImportTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewImportTemplateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImportTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewImportTemplateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImportTemplateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewImportTemplateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /process-groups/{id}/templates/import] importTemplate", response, response.Code())
	}
}

// NewImportTemplateCreated creates a ImportTemplateCreated with default headers values
func NewImportTemplateCreated() *ImportTemplateCreated {
	return &ImportTemplateCreated{}
}

/*
ImportTemplateCreated describes a response with status code 201, with default header values.

successful operation
*/
type ImportTemplateCreated struct {
	Payload *models.TemplateEntity
}

// IsSuccess returns true when this import template created response has a 2xx status code
func (o *ImportTemplateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this import template created response has a 3xx status code
func (o *ImportTemplateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this import template created response has a 4xx status code
func (o *ImportTemplateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this import template created response has a 5xx status code
func (o *ImportTemplateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this import template created response a status code equal to that given
func (o *ImportTemplateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the import template created response
func (o *ImportTemplateCreated) Code() int {
	return 201
}

func (o *ImportTemplateCreated) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/templates/import][%d] importTemplateCreated  %+v", 201, o.Payload)
}

func (o *ImportTemplateCreated) String() string {
	return fmt.Sprintf("[POST /process-groups/{id}/templates/import][%d] importTemplateCreated  %+v", 201, o.Payload)
}

func (o *ImportTemplateCreated) GetPayload() *models.TemplateEntity {
	return o.Payload
}

func (o *ImportTemplateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TemplateEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportTemplateBadRequest creates a ImportTemplateBadRequest with default headers values
func NewImportTemplateBadRequest() *ImportTemplateBadRequest {
	return &ImportTemplateBadRequest{}
}

/*
ImportTemplateBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type ImportTemplateBadRequest struct {
}

// IsSuccess returns true when this import template bad request response has a 2xx status code
func (o *ImportTemplateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this import template bad request response has a 3xx status code
func (o *ImportTemplateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this import template bad request response has a 4xx status code
func (o *ImportTemplateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this import template bad request response has a 5xx status code
func (o *ImportTemplateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this import template bad request response a status code equal to that given
func (o *ImportTemplateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the import template bad request response
func (o *ImportTemplateBadRequest) Code() int {
	return 400
}

func (o *ImportTemplateBadRequest) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/templates/import][%d] importTemplateBadRequest ", 400)
}

func (o *ImportTemplateBadRequest) String() string {
	return fmt.Sprintf("[POST /process-groups/{id}/templates/import][%d] importTemplateBadRequest ", 400)
}

func (o *ImportTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportTemplateUnauthorized creates a ImportTemplateUnauthorized with default headers values
func NewImportTemplateUnauthorized() *ImportTemplateUnauthorized {
	return &ImportTemplateUnauthorized{}
}

/*
ImportTemplateUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type ImportTemplateUnauthorized struct {
}

// IsSuccess returns true when this import template unauthorized response has a 2xx status code
func (o *ImportTemplateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this import template unauthorized response has a 3xx status code
func (o *ImportTemplateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this import template unauthorized response has a 4xx status code
func (o *ImportTemplateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this import template unauthorized response has a 5xx status code
func (o *ImportTemplateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this import template unauthorized response a status code equal to that given
func (o *ImportTemplateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the import template unauthorized response
func (o *ImportTemplateUnauthorized) Code() int {
	return 401
}

func (o *ImportTemplateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/templates/import][%d] importTemplateUnauthorized ", 401)
}

func (o *ImportTemplateUnauthorized) String() string {
	return fmt.Sprintf("[POST /process-groups/{id}/templates/import][%d] importTemplateUnauthorized ", 401)
}

func (o *ImportTemplateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportTemplateForbidden creates a ImportTemplateForbidden with default headers values
func NewImportTemplateForbidden() *ImportTemplateForbidden {
	return &ImportTemplateForbidden{}
}

/*
ImportTemplateForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type ImportTemplateForbidden struct {
}

// IsSuccess returns true when this import template forbidden response has a 2xx status code
func (o *ImportTemplateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this import template forbidden response has a 3xx status code
func (o *ImportTemplateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this import template forbidden response has a 4xx status code
func (o *ImportTemplateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this import template forbidden response has a 5xx status code
func (o *ImportTemplateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this import template forbidden response a status code equal to that given
func (o *ImportTemplateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the import template forbidden response
func (o *ImportTemplateForbidden) Code() int {
	return 403
}

func (o *ImportTemplateForbidden) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/templates/import][%d] importTemplateForbidden ", 403)
}

func (o *ImportTemplateForbidden) String() string {
	return fmt.Sprintf("[POST /process-groups/{id}/templates/import][%d] importTemplateForbidden ", 403)
}

func (o *ImportTemplateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportTemplateConflict creates a ImportTemplateConflict with default headers values
func NewImportTemplateConflict() *ImportTemplateConflict {
	return &ImportTemplateConflict{}
}

/*
ImportTemplateConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type ImportTemplateConflict struct {
}

// IsSuccess returns true when this import template conflict response has a 2xx status code
func (o *ImportTemplateConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this import template conflict response has a 3xx status code
func (o *ImportTemplateConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this import template conflict response has a 4xx status code
func (o *ImportTemplateConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this import template conflict response has a 5xx status code
func (o *ImportTemplateConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this import template conflict response a status code equal to that given
func (o *ImportTemplateConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the import template conflict response
func (o *ImportTemplateConflict) Code() int {
	return 409
}

func (o *ImportTemplateConflict) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/templates/import][%d] importTemplateConflict ", 409)
}

func (o *ImportTemplateConflict) String() string {
	return fmt.Sprintf("[POST /process-groups/{id}/templates/import][%d] importTemplateConflict ", 409)
}

func (o *ImportTemplateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
