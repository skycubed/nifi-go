// Code generated by go-swagger; DO NOT EDIT.

package snippets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// DeleteSnippetReader is a Reader for the DeleteSnippet structure.
type DeleteSnippetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSnippetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSnippetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteSnippetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteSnippetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteSnippetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteSnippetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteSnippetConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteSnippetOK creates a DeleteSnippetOK with default headers values
func NewDeleteSnippetOK() *DeleteSnippetOK {
	return &DeleteSnippetOK{}
}

/* DeleteSnippetOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteSnippetOK struct {
	Payload *models.SnippetEntity
}

func (o *DeleteSnippetOK) Error() string {
	return fmt.Sprintf("[DELETE /snippets/{id}][%d] deleteSnippetOK  %+v", 200, o.Payload)
}
func (o *DeleteSnippetOK) GetPayload() *models.SnippetEntity {
	return o.Payload
}

func (o *DeleteSnippetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnippetEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSnippetBadRequest creates a DeleteSnippetBadRequest with default headers values
func NewDeleteSnippetBadRequest() *DeleteSnippetBadRequest {
	return &DeleteSnippetBadRequest{}
}

/* DeleteSnippetBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type DeleteSnippetBadRequest struct {
}

func (o *DeleteSnippetBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /snippets/{id}][%d] deleteSnippetBadRequest ", 400)
}

func (o *DeleteSnippetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSnippetUnauthorized creates a DeleteSnippetUnauthorized with default headers values
func NewDeleteSnippetUnauthorized() *DeleteSnippetUnauthorized {
	return &DeleteSnippetUnauthorized{}
}

/* DeleteSnippetUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type DeleteSnippetUnauthorized struct {
}

func (o *DeleteSnippetUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /snippets/{id}][%d] deleteSnippetUnauthorized ", 401)
}

func (o *DeleteSnippetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSnippetForbidden creates a DeleteSnippetForbidden with default headers values
func NewDeleteSnippetForbidden() *DeleteSnippetForbidden {
	return &DeleteSnippetForbidden{}
}

/* DeleteSnippetForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type DeleteSnippetForbidden struct {
}

func (o *DeleteSnippetForbidden) Error() string {
	return fmt.Sprintf("[DELETE /snippets/{id}][%d] deleteSnippetForbidden ", 403)
}

func (o *DeleteSnippetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSnippetNotFound creates a DeleteSnippetNotFound with default headers values
func NewDeleteSnippetNotFound() *DeleteSnippetNotFound {
	return &DeleteSnippetNotFound{}
}

/* DeleteSnippetNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type DeleteSnippetNotFound struct {
}

func (o *DeleteSnippetNotFound) Error() string {
	return fmt.Sprintf("[DELETE /snippets/{id}][%d] deleteSnippetNotFound ", 404)
}

func (o *DeleteSnippetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSnippetConflict creates a DeleteSnippetConflict with default headers values
func NewDeleteSnippetConflict() *DeleteSnippetConflict {
	return &DeleteSnippetConflict{}
}

/* DeleteSnippetConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type DeleteSnippetConflict struct {
}

func (o *DeleteSnippetConflict) Error() string {
	return fmt.Sprintf("[DELETE /snippets/{id}][%d] deleteSnippetConflict ", 409)
}

func (o *DeleteSnippetConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}