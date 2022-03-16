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

// UploadProcessGroupReader is a Reader for the UploadProcessGroup structure.
type UploadProcessGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadProcessGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUploadProcessGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUploadProcessGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUploadProcessGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUploadProcessGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUploadProcessGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUploadProcessGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUploadProcessGroupOK creates a UploadProcessGroupOK with default headers values
func NewUploadProcessGroupOK() *UploadProcessGroupOK {
	return &UploadProcessGroupOK{}
}

/* UploadProcessGroupOK describes a response with status code 200, with default header values.

successful operation
*/
type UploadProcessGroupOK struct {
	Payload *models.ProcessGroupEntity
}

func (o *UploadProcessGroupOK) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/process-groups/upload][%d] uploadProcessGroupOK  %+v", 200, o.Payload)
}
func (o *UploadProcessGroupOK) GetPayload() *models.ProcessGroupEntity {
	return o.Payload
}

func (o *UploadProcessGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProcessGroupEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadProcessGroupBadRequest creates a UploadProcessGroupBadRequest with default headers values
func NewUploadProcessGroupBadRequest() *UploadProcessGroupBadRequest {
	return &UploadProcessGroupBadRequest{}
}

/* UploadProcessGroupBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type UploadProcessGroupBadRequest struct {
}

func (o *UploadProcessGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/process-groups/upload][%d] uploadProcessGroupBadRequest ", 400)
}

func (o *UploadProcessGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUploadProcessGroupUnauthorized creates a UploadProcessGroupUnauthorized with default headers values
func NewUploadProcessGroupUnauthorized() *UploadProcessGroupUnauthorized {
	return &UploadProcessGroupUnauthorized{}
}

/* UploadProcessGroupUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type UploadProcessGroupUnauthorized struct {
}

func (o *UploadProcessGroupUnauthorized) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/process-groups/upload][%d] uploadProcessGroupUnauthorized ", 401)
}

func (o *UploadProcessGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUploadProcessGroupForbidden creates a UploadProcessGroupForbidden with default headers values
func NewUploadProcessGroupForbidden() *UploadProcessGroupForbidden {
	return &UploadProcessGroupForbidden{}
}

/* UploadProcessGroupForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type UploadProcessGroupForbidden struct {
}

func (o *UploadProcessGroupForbidden) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/process-groups/upload][%d] uploadProcessGroupForbidden ", 403)
}

func (o *UploadProcessGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUploadProcessGroupNotFound creates a UploadProcessGroupNotFound with default headers values
func NewUploadProcessGroupNotFound() *UploadProcessGroupNotFound {
	return &UploadProcessGroupNotFound{}
}

/* UploadProcessGroupNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type UploadProcessGroupNotFound struct {
}

func (o *UploadProcessGroupNotFound) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/process-groups/upload][%d] uploadProcessGroupNotFound ", 404)
}

func (o *UploadProcessGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUploadProcessGroupConflict creates a UploadProcessGroupConflict with default headers values
func NewUploadProcessGroupConflict() *UploadProcessGroupConflict {
	return &UploadProcessGroupConflict{}
}

/* UploadProcessGroupConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type UploadProcessGroupConflict struct {
}

func (o *UploadProcessGroupConflict) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/process-groups/upload][%d] uploadProcessGroupConflict ", 409)
}

func (o *UploadProcessGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}