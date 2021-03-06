// Code generated by go-swagger; DO NOT EDIT.

package provenance_events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// GetInputContentReader is a Reader for the GetInputContent structure.
type GetInputContentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInputContentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInputContentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInputContentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetInputContentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetInputContentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetInputContentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetInputContentConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetInputContentOK creates a GetInputContentOK with default headers values
func NewGetInputContentOK() *GetInputContentOK {
	return &GetInputContentOK{}
}

/* GetInputContentOK describes a response with status code 200, with default header values.

successful operation
*/
type GetInputContentOK struct {
	Payload models.StreamingOutput
}

func (o *GetInputContentOK) Error() string {
	return fmt.Sprintf("[GET /provenance-events/{id}/content/input][%d] getInputContentOK  %+v", 200, o.Payload)
}
func (o *GetInputContentOK) GetPayload() models.StreamingOutput {
	return o.Payload
}

func (o *GetInputContentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInputContentBadRequest creates a GetInputContentBadRequest with default headers values
func NewGetInputContentBadRequest() *GetInputContentBadRequest {
	return &GetInputContentBadRequest{}
}

/* GetInputContentBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetInputContentBadRequest struct {
}

func (o *GetInputContentBadRequest) Error() string {
	return fmt.Sprintf("[GET /provenance-events/{id}/content/input][%d] getInputContentBadRequest ", 400)
}

func (o *GetInputContentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInputContentUnauthorized creates a GetInputContentUnauthorized with default headers values
func NewGetInputContentUnauthorized() *GetInputContentUnauthorized {
	return &GetInputContentUnauthorized{}
}

/* GetInputContentUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetInputContentUnauthorized struct {
}

func (o *GetInputContentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /provenance-events/{id}/content/input][%d] getInputContentUnauthorized ", 401)
}

func (o *GetInputContentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInputContentForbidden creates a GetInputContentForbidden with default headers values
func NewGetInputContentForbidden() *GetInputContentForbidden {
	return &GetInputContentForbidden{}
}

/* GetInputContentForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetInputContentForbidden struct {
}

func (o *GetInputContentForbidden) Error() string {
	return fmt.Sprintf("[GET /provenance-events/{id}/content/input][%d] getInputContentForbidden ", 403)
}

func (o *GetInputContentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInputContentNotFound creates a GetInputContentNotFound with default headers values
func NewGetInputContentNotFound() *GetInputContentNotFound {
	return &GetInputContentNotFound{}
}

/* GetInputContentNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetInputContentNotFound struct {
}

func (o *GetInputContentNotFound) Error() string {
	return fmt.Sprintf("[GET /provenance-events/{id}/content/input][%d] getInputContentNotFound ", 404)
}

func (o *GetInputContentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInputContentConflict creates a GetInputContentConflict with default headers values
func NewGetInputContentConflict() *GetInputContentConflict {
	return &GetInputContentConflict{}
}

/* GetInputContentConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetInputContentConflict struct {
}

func (o *GetInputContentConflict) Error() string {
	return fmt.Sprintf("[GET /provenance-events/{id}/content/input][%d] getInputContentConflict ", 409)
}

func (o *GetInputContentConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
