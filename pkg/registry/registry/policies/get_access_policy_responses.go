// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/registry/models"
)

// GetAccessPolicyReader is a Reader for the GetAccessPolicy structure.
type GetAccessPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccessPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccessPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAccessPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAccessPolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAccessPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAccessPolicyConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAccessPolicyOK creates a GetAccessPolicyOK with default headers values
func NewGetAccessPolicyOK() *GetAccessPolicyOK {
	return &GetAccessPolicyOK{}
}

/* GetAccessPolicyOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAccessPolicyOK struct {
	Payload *models.AccessPolicy
}

func (o *GetAccessPolicyOK) Error() string {
	return fmt.Sprintf("[GET /policies/{id}][%d] getAccessPolicyOK  %+v", 200, o.Payload)
}
func (o *GetAccessPolicyOK) GetPayload() *models.AccessPolicy {
	return o.Payload
}

func (o *GetAccessPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccessPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccessPolicyUnauthorized creates a GetAccessPolicyUnauthorized with default headers values
func NewGetAccessPolicyUnauthorized() *GetAccessPolicyUnauthorized {
	return &GetAccessPolicyUnauthorized{}
}

/* GetAccessPolicyUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetAccessPolicyUnauthorized struct {
}

func (o *GetAccessPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[GET /policies/{id}][%d] getAccessPolicyUnauthorized ", 401)
}

func (o *GetAccessPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAccessPolicyForbidden creates a GetAccessPolicyForbidden with default headers values
func NewGetAccessPolicyForbidden() *GetAccessPolicyForbidden {
	return &GetAccessPolicyForbidden{}
}

/* GetAccessPolicyForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetAccessPolicyForbidden struct {
}

func (o *GetAccessPolicyForbidden) Error() string {
	return fmt.Sprintf("[GET /policies/{id}][%d] getAccessPolicyForbidden ", 403)
}

func (o *GetAccessPolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAccessPolicyNotFound creates a GetAccessPolicyNotFound with default headers values
func NewGetAccessPolicyNotFound() *GetAccessPolicyNotFound {
	return &GetAccessPolicyNotFound{}
}

/* GetAccessPolicyNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetAccessPolicyNotFound struct {
}

func (o *GetAccessPolicyNotFound) Error() string {
	return fmt.Sprintf("[GET /policies/{id}][%d] getAccessPolicyNotFound ", 404)
}

func (o *GetAccessPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAccessPolicyConflict creates a GetAccessPolicyConflict with default headers values
func NewGetAccessPolicyConflict() *GetAccessPolicyConflict {
	return &GetAccessPolicyConflict{}
}

/* GetAccessPolicyConflict describes a response with status code 409, with default header values.

NiFi Registry was unable to complete the request because it assumes a server state that is not valid.
*/
type GetAccessPolicyConflict struct {
}

func (o *GetAccessPolicyConflict) Error() string {
	return fmt.Sprintf("[GET /policies/{id}][%d] getAccessPolicyConflict ", 409)
}

func (o *GetAccessPolicyConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
