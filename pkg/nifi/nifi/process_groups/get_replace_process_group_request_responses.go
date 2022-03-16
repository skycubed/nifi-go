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

// GetReplaceProcessGroupRequestReader is a Reader for the GetReplaceProcessGroupRequest structure.
type GetReplaceProcessGroupRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReplaceProcessGroupRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReplaceProcessGroupRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetReplaceProcessGroupRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetReplaceProcessGroupRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetReplaceProcessGroupRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetReplaceProcessGroupRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetReplaceProcessGroupRequestConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetReplaceProcessGroupRequestOK creates a GetReplaceProcessGroupRequestOK with default headers values
func NewGetReplaceProcessGroupRequestOK() *GetReplaceProcessGroupRequestOK {
	return &GetReplaceProcessGroupRequestOK{}
}

/* GetReplaceProcessGroupRequestOK describes a response with status code 200, with default header values.

successful operation
*/
type GetReplaceProcessGroupRequestOK struct {
	Payload *models.ProcessGroupReplaceRequestEntity
}

func (o *GetReplaceProcessGroupRequestOK) Error() string {
	return fmt.Sprintf("[GET /process-groups/replace-requests/{id}][%d] getReplaceProcessGroupRequestOK  %+v", 200, o.Payload)
}
func (o *GetReplaceProcessGroupRequestOK) GetPayload() *models.ProcessGroupReplaceRequestEntity {
	return o.Payload
}

func (o *GetReplaceProcessGroupRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProcessGroupReplaceRequestEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReplaceProcessGroupRequestBadRequest creates a GetReplaceProcessGroupRequestBadRequest with default headers values
func NewGetReplaceProcessGroupRequestBadRequest() *GetReplaceProcessGroupRequestBadRequest {
	return &GetReplaceProcessGroupRequestBadRequest{}
}

/* GetReplaceProcessGroupRequestBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetReplaceProcessGroupRequestBadRequest struct {
}

func (o *GetReplaceProcessGroupRequestBadRequest) Error() string {
	return fmt.Sprintf("[GET /process-groups/replace-requests/{id}][%d] getReplaceProcessGroupRequestBadRequest ", 400)
}

func (o *GetReplaceProcessGroupRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetReplaceProcessGroupRequestUnauthorized creates a GetReplaceProcessGroupRequestUnauthorized with default headers values
func NewGetReplaceProcessGroupRequestUnauthorized() *GetReplaceProcessGroupRequestUnauthorized {
	return &GetReplaceProcessGroupRequestUnauthorized{}
}

/* GetReplaceProcessGroupRequestUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetReplaceProcessGroupRequestUnauthorized struct {
}

func (o *GetReplaceProcessGroupRequestUnauthorized) Error() string {
	return fmt.Sprintf("[GET /process-groups/replace-requests/{id}][%d] getReplaceProcessGroupRequestUnauthorized ", 401)
}

func (o *GetReplaceProcessGroupRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetReplaceProcessGroupRequestForbidden creates a GetReplaceProcessGroupRequestForbidden with default headers values
func NewGetReplaceProcessGroupRequestForbidden() *GetReplaceProcessGroupRequestForbidden {
	return &GetReplaceProcessGroupRequestForbidden{}
}

/* GetReplaceProcessGroupRequestForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetReplaceProcessGroupRequestForbidden struct {
}

func (o *GetReplaceProcessGroupRequestForbidden) Error() string {
	return fmt.Sprintf("[GET /process-groups/replace-requests/{id}][%d] getReplaceProcessGroupRequestForbidden ", 403)
}

func (o *GetReplaceProcessGroupRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetReplaceProcessGroupRequestNotFound creates a GetReplaceProcessGroupRequestNotFound with default headers values
func NewGetReplaceProcessGroupRequestNotFound() *GetReplaceProcessGroupRequestNotFound {
	return &GetReplaceProcessGroupRequestNotFound{}
}

/* GetReplaceProcessGroupRequestNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetReplaceProcessGroupRequestNotFound struct {
}

func (o *GetReplaceProcessGroupRequestNotFound) Error() string {
	return fmt.Sprintf("[GET /process-groups/replace-requests/{id}][%d] getReplaceProcessGroupRequestNotFound ", 404)
}

func (o *GetReplaceProcessGroupRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetReplaceProcessGroupRequestConflict creates a GetReplaceProcessGroupRequestConflict with default headers values
func NewGetReplaceProcessGroupRequestConflict() *GetReplaceProcessGroupRequestConflict {
	return &GetReplaceProcessGroupRequestConflict{}
}

/* GetReplaceProcessGroupRequestConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetReplaceProcessGroupRequestConflict struct {
}

func (o *GetReplaceProcessGroupRequestConflict) Error() string {
	return fmt.Sprintf("[GET /process-groups/replace-requests/{id}][%d] getReplaceProcessGroupRequestConflict ", 409)
}

func (o *GetReplaceProcessGroupRequestConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}