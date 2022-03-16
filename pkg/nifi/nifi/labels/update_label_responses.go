// Code generated by go-swagger; DO NOT EDIT.

package labels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// UpdateLabelReader is a Reader for the UpdateLabel structure.
type UpdateLabelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateLabelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateLabelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateLabelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateLabelUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateLabelForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateLabelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateLabelConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateLabelOK creates a UpdateLabelOK with default headers values
func NewUpdateLabelOK() *UpdateLabelOK {
	return &UpdateLabelOK{}
}

/* UpdateLabelOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateLabelOK struct {
	Payload *models.LabelEntity
}

func (o *UpdateLabelOK) Error() string {
	return fmt.Sprintf("[PUT /labels/{id}][%d] updateLabelOK  %+v", 200, o.Payload)
}
func (o *UpdateLabelOK) GetPayload() *models.LabelEntity {
	return o.Payload
}

func (o *UpdateLabelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LabelEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateLabelBadRequest creates a UpdateLabelBadRequest with default headers values
func NewUpdateLabelBadRequest() *UpdateLabelBadRequest {
	return &UpdateLabelBadRequest{}
}

/* UpdateLabelBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type UpdateLabelBadRequest struct {
}

func (o *UpdateLabelBadRequest) Error() string {
	return fmt.Sprintf("[PUT /labels/{id}][%d] updateLabelBadRequest ", 400)
}

func (o *UpdateLabelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateLabelUnauthorized creates a UpdateLabelUnauthorized with default headers values
func NewUpdateLabelUnauthorized() *UpdateLabelUnauthorized {
	return &UpdateLabelUnauthorized{}
}

/* UpdateLabelUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type UpdateLabelUnauthorized struct {
}

func (o *UpdateLabelUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /labels/{id}][%d] updateLabelUnauthorized ", 401)
}

func (o *UpdateLabelUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateLabelForbidden creates a UpdateLabelForbidden with default headers values
func NewUpdateLabelForbidden() *UpdateLabelForbidden {
	return &UpdateLabelForbidden{}
}

/* UpdateLabelForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type UpdateLabelForbidden struct {
}

func (o *UpdateLabelForbidden) Error() string {
	return fmt.Sprintf("[PUT /labels/{id}][%d] updateLabelForbidden ", 403)
}

func (o *UpdateLabelForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateLabelNotFound creates a UpdateLabelNotFound with default headers values
func NewUpdateLabelNotFound() *UpdateLabelNotFound {
	return &UpdateLabelNotFound{}
}

/* UpdateLabelNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type UpdateLabelNotFound struct {
}

func (o *UpdateLabelNotFound) Error() string {
	return fmt.Sprintf("[PUT /labels/{id}][%d] updateLabelNotFound ", 404)
}

func (o *UpdateLabelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateLabelConflict creates a UpdateLabelConflict with default headers values
func NewUpdateLabelConflict() *UpdateLabelConflict {
	return &UpdateLabelConflict{}
}

/* UpdateLabelConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type UpdateLabelConflict struct {
}

func (o *UpdateLabelConflict) Error() string {
	return fmt.Sprintf("[PUT /labels/{id}][%d] updateLabelConflict ", 409)
}

func (o *UpdateLabelConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
