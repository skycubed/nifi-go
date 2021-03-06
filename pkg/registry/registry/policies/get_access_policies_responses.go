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

// GetAccessPoliciesReader is a Reader for the GetAccessPolicies structure.
type GetAccessPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccessPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccessPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAccessPoliciesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAccessPoliciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAccessPoliciesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAccessPoliciesOK creates a GetAccessPoliciesOK with default headers values
func NewGetAccessPoliciesOK() *GetAccessPoliciesOK {
	return &GetAccessPoliciesOK{}
}

/* GetAccessPoliciesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAccessPoliciesOK struct {
	Payload []*models.AccessPolicy
}

func (o *GetAccessPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /policies][%d] getAccessPoliciesOK  %+v", 200, o.Payload)
}
func (o *GetAccessPoliciesOK) GetPayload() []*models.AccessPolicy {
	return o.Payload
}

func (o *GetAccessPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccessPoliciesUnauthorized creates a GetAccessPoliciesUnauthorized with default headers values
func NewGetAccessPoliciesUnauthorized() *GetAccessPoliciesUnauthorized {
	return &GetAccessPoliciesUnauthorized{}
}

/* GetAccessPoliciesUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetAccessPoliciesUnauthorized struct {
}

func (o *GetAccessPoliciesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /policies][%d] getAccessPoliciesUnauthorized ", 401)
}

func (o *GetAccessPoliciesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAccessPoliciesForbidden creates a GetAccessPoliciesForbidden with default headers values
func NewGetAccessPoliciesForbidden() *GetAccessPoliciesForbidden {
	return &GetAccessPoliciesForbidden{}
}

/* GetAccessPoliciesForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetAccessPoliciesForbidden struct {
}

func (o *GetAccessPoliciesForbidden) Error() string {
	return fmt.Sprintf("[GET /policies][%d] getAccessPoliciesForbidden ", 403)
}

func (o *GetAccessPoliciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAccessPoliciesConflict creates a GetAccessPoliciesConflict with default headers values
func NewGetAccessPoliciesConflict() *GetAccessPoliciesConflict {
	return &GetAccessPoliciesConflict{}
}

/* GetAccessPoliciesConflict describes a response with status code 409, with default header values.

NiFi Registry was unable to complete the request because it assumes a server state that is not valid.
*/
type GetAccessPoliciesConflict struct {
}

func (o *GetAccessPoliciesConflict) Error() string {
	return fmt.Sprintf("[GET /policies][%d] getAccessPoliciesConflict ", 409)
}

func (o *GetAccessPoliciesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
