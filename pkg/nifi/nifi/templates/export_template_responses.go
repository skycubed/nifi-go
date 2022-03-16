// Code generated by go-swagger; DO NOT EDIT.

package templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExportTemplateReader is a Reader for the ExportTemplate structure.
type ExportTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExportTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExportTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewExportTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewExportTemplateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewExportTemplateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewExportTemplateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewExportTemplateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExportTemplateOK creates a ExportTemplateOK with default headers values
func NewExportTemplateOK() *ExportTemplateOK {
	return &ExportTemplateOK{}
}

/* ExportTemplateOK describes a response with status code 200, with default header values.

successful operation
*/
type ExportTemplateOK struct {
	Payload string
}

func (o *ExportTemplateOK) Error() string {
	return fmt.Sprintf("[GET /templates/{id}/download][%d] exportTemplateOK  %+v", 200, o.Payload)
}
func (o *ExportTemplateOK) GetPayload() string {
	return o.Payload
}

func (o *ExportTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExportTemplateBadRequest creates a ExportTemplateBadRequest with default headers values
func NewExportTemplateBadRequest() *ExportTemplateBadRequest {
	return &ExportTemplateBadRequest{}
}

/* ExportTemplateBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type ExportTemplateBadRequest struct {
}

func (o *ExportTemplateBadRequest) Error() string {
	return fmt.Sprintf("[GET /templates/{id}/download][%d] exportTemplateBadRequest ", 400)
}

func (o *ExportTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportTemplateUnauthorized creates a ExportTemplateUnauthorized with default headers values
func NewExportTemplateUnauthorized() *ExportTemplateUnauthorized {
	return &ExportTemplateUnauthorized{}
}

/* ExportTemplateUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type ExportTemplateUnauthorized struct {
}

func (o *ExportTemplateUnauthorized) Error() string {
	return fmt.Sprintf("[GET /templates/{id}/download][%d] exportTemplateUnauthorized ", 401)
}

func (o *ExportTemplateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportTemplateForbidden creates a ExportTemplateForbidden with default headers values
func NewExportTemplateForbidden() *ExportTemplateForbidden {
	return &ExportTemplateForbidden{}
}

/* ExportTemplateForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type ExportTemplateForbidden struct {
}

func (o *ExportTemplateForbidden) Error() string {
	return fmt.Sprintf("[GET /templates/{id}/download][%d] exportTemplateForbidden ", 403)
}

func (o *ExportTemplateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportTemplateNotFound creates a ExportTemplateNotFound with default headers values
func NewExportTemplateNotFound() *ExportTemplateNotFound {
	return &ExportTemplateNotFound{}
}

/* ExportTemplateNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type ExportTemplateNotFound struct {
}

func (o *ExportTemplateNotFound) Error() string {
	return fmt.Sprintf("[GET /templates/{id}/download][%d] exportTemplateNotFound ", 404)
}

func (o *ExportTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportTemplateConflict creates a ExportTemplateConflict with default headers values
func NewExportTemplateConflict() *ExportTemplateConflict {
	return &ExportTemplateConflict{}
}

/* ExportTemplateConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type ExportTemplateConflict struct {
}

func (o *ExportTemplateConflict) Error() string {
	return fmt.Sprintf("[GET /templates/{id}/download][%d] exportTemplateConflict ", 409)
}

func (o *ExportTemplateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
