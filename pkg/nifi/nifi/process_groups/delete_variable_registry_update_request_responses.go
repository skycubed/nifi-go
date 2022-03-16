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

// DeleteVariableRegistryUpdateRequestReader is a Reader for the DeleteVariableRegistryUpdateRequest structure.
type DeleteVariableRegistryUpdateRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVariableRegistryUpdateRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteVariableRegistryUpdateRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteVariableRegistryUpdateRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteVariableRegistryUpdateRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteVariableRegistryUpdateRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteVariableRegistryUpdateRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteVariableRegistryUpdateRequestConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteVariableRegistryUpdateRequestOK creates a DeleteVariableRegistryUpdateRequestOK with default headers values
func NewDeleteVariableRegistryUpdateRequestOK() *DeleteVariableRegistryUpdateRequestOK {
	return &DeleteVariableRegistryUpdateRequestOK{}
}

/* DeleteVariableRegistryUpdateRequestOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteVariableRegistryUpdateRequestOK struct {
	Payload *models.VariableRegistryUpdateRequestEntity
}

func (o *DeleteVariableRegistryUpdateRequestOK) Error() string {
	return fmt.Sprintf("[DELETE /process-groups/{groupId}/variable-registry/update-requests/{updateId}][%d] deleteVariableRegistryUpdateRequestOK  %+v", 200, o.Payload)
}
func (o *DeleteVariableRegistryUpdateRequestOK) GetPayload() *models.VariableRegistryUpdateRequestEntity {
	return o.Payload
}

func (o *DeleteVariableRegistryUpdateRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VariableRegistryUpdateRequestEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVariableRegistryUpdateRequestBadRequest creates a DeleteVariableRegistryUpdateRequestBadRequest with default headers values
func NewDeleteVariableRegistryUpdateRequestBadRequest() *DeleteVariableRegistryUpdateRequestBadRequest {
	return &DeleteVariableRegistryUpdateRequestBadRequest{}
}

/* DeleteVariableRegistryUpdateRequestBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type DeleteVariableRegistryUpdateRequestBadRequest struct {
}

func (o *DeleteVariableRegistryUpdateRequestBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /process-groups/{groupId}/variable-registry/update-requests/{updateId}][%d] deleteVariableRegistryUpdateRequestBadRequest ", 400)
}

func (o *DeleteVariableRegistryUpdateRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVariableRegistryUpdateRequestUnauthorized creates a DeleteVariableRegistryUpdateRequestUnauthorized with default headers values
func NewDeleteVariableRegistryUpdateRequestUnauthorized() *DeleteVariableRegistryUpdateRequestUnauthorized {
	return &DeleteVariableRegistryUpdateRequestUnauthorized{}
}

/* DeleteVariableRegistryUpdateRequestUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type DeleteVariableRegistryUpdateRequestUnauthorized struct {
}

func (o *DeleteVariableRegistryUpdateRequestUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /process-groups/{groupId}/variable-registry/update-requests/{updateId}][%d] deleteVariableRegistryUpdateRequestUnauthorized ", 401)
}

func (o *DeleteVariableRegistryUpdateRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVariableRegistryUpdateRequestForbidden creates a DeleteVariableRegistryUpdateRequestForbidden with default headers values
func NewDeleteVariableRegistryUpdateRequestForbidden() *DeleteVariableRegistryUpdateRequestForbidden {
	return &DeleteVariableRegistryUpdateRequestForbidden{}
}

/* DeleteVariableRegistryUpdateRequestForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type DeleteVariableRegistryUpdateRequestForbidden struct {
}

func (o *DeleteVariableRegistryUpdateRequestForbidden) Error() string {
	return fmt.Sprintf("[DELETE /process-groups/{groupId}/variable-registry/update-requests/{updateId}][%d] deleteVariableRegistryUpdateRequestForbidden ", 403)
}

func (o *DeleteVariableRegistryUpdateRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVariableRegistryUpdateRequestNotFound creates a DeleteVariableRegistryUpdateRequestNotFound with default headers values
func NewDeleteVariableRegistryUpdateRequestNotFound() *DeleteVariableRegistryUpdateRequestNotFound {
	return &DeleteVariableRegistryUpdateRequestNotFound{}
}

/* DeleteVariableRegistryUpdateRequestNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type DeleteVariableRegistryUpdateRequestNotFound struct {
}

func (o *DeleteVariableRegistryUpdateRequestNotFound) Error() string {
	return fmt.Sprintf("[DELETE /process-groups/{groupId}/variable-registry/update-requests/{updateId}][%d] deleteVariableRegistryUpdateRequestNotFound ", 404)
}

func (o *DeleteVariableRegistryUpdateRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVariableRegistryUpdateRequestConflict creates a DeleteVariableRegistryUpdateRequestConflict with default headers values
func NewDeleteVariableRegistryUpdateRequestConflict() *DeleteVariableRegistryUpdateRequestConflict {
	return &DeleteVariableRegistryUpdateRequestConflict{}
}

/* DeleteVariableRegistryUpdateRequestConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type DeleteVariableRegistryUpdateRequestConflict struct {
}

func (o *DeleteVariableRegistryUpdateRequestConflict) Error() string {
	return fmt.Sprintf("[DELETE /process-groups/{groupId}/variable-registry/update-requests/{updateId}][%d] deleteVariableRegistryUpdateRequestConflict ", 409)
}

func (o *DeleteVariableRegistryUpdateRequestConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}