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

// InstantiateTemplateReader is a Reader for the InstantiateTemplate structure.
type InstantiateTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InstantiateTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewInstantiateTemplateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInstantiateTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewInstantiateTemplateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewInstantiateTemplateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewInstantiateTemplateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewInstantiateTemplateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /process-groups/{id}/template-instance] instantiateTemplate", response, response.Code())
	}
}

// NewInstantiateTemplateCreated creates a InstantiateTemplateCreated with default headers values
func NewInstantiateTemplateCreated() *InstantiateTemplateCreated {
	return &InstantiateTemplateCreated{}
}

/*
InstantiateTemplateCreated describes a response with status code 201, with default header values.

successful operation
*/
type InstantiateTemplateCreated struct {
	Payload *models.FlowEntity
}

// IsSuccess returns true when this instantiate template created response has a 2xx status code
func (o *InstantiateTemplateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this instantiate template created response has a 3xx status code
func (o *InstantiateTemplateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this instantiate template created response has a 4xx status code
func (o *InstantiateTemplateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this instantiate template created response has a 5xx status code
func (o *InstantiateTemplateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this instantiate template created response a status code equal to that given
func (o *InstantiateTemplateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the instantiate template created response
func (o *InstantiateTemplateCreated) Code() int {
	return 201
}

func (o *InstantiateTemplateCreated) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/template-instance][%d] instantiateTemplateCreated  %+v", 201, o.Payload)
}

func (o *InstantiateTemplateCreated) String() string {
	return fmt.Sprintf("[POST /process-groups/{id}/template-instance][%d] instantiateTemplateCreated  %+v", 201, o.Payload)
}

func (o *InstantiateTemplateCreated) GetPayload() *models.FlowEntity {
	return o.Payload
}

func (o *InstantiateTemplateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FlowEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstantiateTemplateBadRequest creates a InstantiateTemplateBadRequest with default headers values
func NewInstantiateTemplateBadRequest() *InstantiateTemplateBadRequest {
	return &InstantiateTemplateBadRequest{}
}

/*
InstantiateTemplateBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type InstantiateTemplateBadRequest struct {
}

// IsSuccess returns true when this instantiate template bad request response has a 2xx status code
func (o *InstantiateTemplateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this instantiate template bad request response has a 3xx status code
func (o *InstantiateTemplateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this instantiate template bad request response has a 4xx status code
func (o *InstantiateTemplateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this instantiate template bad request response has a 5xx status code
func (o *InstantiateTemplateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this instantiate template bad request response a status code equal to that given
func (o *InstantiateTemplateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the instantiate template bad request response
func (o *InstantiateTemplateBadRequest) Code() int {
	return 400
}

func (o *InstantiateTemplateBadRequest) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/template-instance][%d] instantiateTemplateBadRequest ", 400)
}

func (o *InstantiateTemplateBadRequest) String() string {
	return fmt.Sprintf("[POST /process-groups/{id}/template-instance][%d] instantiateTemplateBadRequest ", 400)
}

func (o *InstantiateTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInstantiateTemplateUnauthorized creates a InstantiateTemplateUnauthorized with default headers values
func NewInstantiateTemplateUnauthorized() *InstantiateTemplateUnauthorized {
	return &InstantiateTemplateUnauthorized{}
}

/*
InstantiateTemplateUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type InstantiateTemplateUnauthorized struct {
}

// IsSuccess returns true when this instantiate template unauthorized response has a 2xx status code
func (o *InstantiateTemplateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this instantiate template unauthorized response has a 3xx status code
func (o *InstantiateTemplateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this instantiate template unauthorized response has a 4xx status code
func (o *InstantiateTemplateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this instantiate template unauthorized response has a 5xx status code
func (o *InstantiateTemplateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this instantiate template unauthorized response a status code equal to that given
func (o *InstantiateTemplateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the instantiate template unauthorized response
func (o *InstantiateTemplateUnauthorized) Code() int {
	return 401
}

func (o *InstantiateTemplateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/template-instance][%d] instantiateTemplateUnauthorized ", 401)
}

func (o *InstantiateTemplateUnauthorized) String() string {
	return fmt.Sprintf("[POST /process-groups/{id}/template-instance][%d] instantiateTemplateUnauthorized ", 401)
}

func (o *InstantiateTemplateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInstantiateTemplateForbidden creates a InstantiateTemplateForbidden with default headers values
func NewInstantiateTemplateForbidden() *InstantiateTemplateForbidden {
	return &InstantiateTemplateForbidden{}
}

/*
InstantiateTemplateForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type InstantiateTemplateForbidden struct {
}

// IsSuccess returns true when this instantiate template forbidden response has a 2xx status code
func (o *InstantiateTemplateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this instantiate template forbidden response has a 3xx status code
func (o *InstantiateTemplateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this instantiate template forbidden response has a 4xx status code
func (o *InstantiateTemplateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this instantiate template forbidden response has a 5xx status code
func (o *InstantiateTemplateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this instantiate template forbidden response a status code equal to that given
func (o *InstantiateTemplateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the instantiate template forbidden response
func (o *InstantiateTemplateForbidden) Code() int {
	return 403
}

func (o *InstantiateTemplateForbidden) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/template-instance][%d] instantiateTemplateForbidden ", 403)
}

func (o *InstantiateTemplateForbidden) String() string {
	return fmt.Sprintf("[POST /process-groups/{id}/template-instance][%d] instantiateTemplateForbidden ", 403)
}

func (o *InstantiateTemplateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInstantiateTemplateNotFound creates a InstantiateTemplateNotFound with default headers values
func NewInstantiateTemplateNotFound() *InstantiateTemplateNotFound {
	return &InstantiateTemplateNotFound{}
}

/*
InstantiateTemplateNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type InstantiateTemplateNotFound struct {
}

// IsSuccess returns true when this instantiate template not found response has a 2xx status code
func (o *InstantiateTemplateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this instantiate template not found response has a 3xx status code
func (o *InstantiateTemplateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this instantiate template not found response has a 4xx status code
func (o *InstantiateTemplateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this instantiate template not found response has a 5xx status code
func (o *InstantiateTemplateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this instantiate template not found response a status code equal to that given
func (o *InstantiateTemplateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the instantiate template not found response
func (o *InstantiateTemplateNotFound) Code() int {
	return 404
}

func (o *InstantiateTemplateNotFound) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/template-instance][%d] instantiateTemplateNotFound ", 404)
}

func (o *InstantiateTemplateNotFound) String() string {
	return fmt.Sprintf("[POST /process-groups/{id}/template-instance][%d] instantiateTemplateNotFound ", 404)
}

func (o *InstantiateTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInstantiateTemplateConflict creates a InstantiateTemplateConflict with default headers values
func NewInstantiateTemplateConflict() *InstantiateTemplateConflict {
	return &InstantiateTemplateConflict{}
}

/*
InstantiateTemplateConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type InstantiateTemplateConflict struct {
}

// IsSuccess returns true when this instantiate template conflict response has a 2xx status code
func (o *InstantiateTemplateConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this instantiate template conflict response has a 3xx status code
func (o *InstantiateTemplateConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this instantiate template conflict response has a 4xx status code
func (o *InstantiateTemplateConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this instantiate template conflict response has a 5xx status code
func (o *InstantiateTemplateConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this instantiate template conflict response a status code equal to that given
func (o *InstantiateTemplateConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the instantiate template conflict response
func (o *InstantiateTemplateConflict) Code() int {
	return 409
}

func (o *InstantiateTemplateConflict) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/template-instance][%d] instantiateTemplateConflict ", 409)
}

func (o *InstantiateTemplateConflict) String() string {
	return fmt.Sprintf("[POST /process-groups/{id}/template-instance][%d] instantiateTemplateConflict ", 409)
}

func (o *InstantiateTemplateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
