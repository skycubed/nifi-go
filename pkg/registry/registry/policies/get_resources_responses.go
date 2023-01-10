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

// GetResourcesReader is a Reader for the GetResources structure.
type GetResourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetResourcesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetResourcesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetResourcesOK creates a GetResourcesOK with default headers values
func NewGetResourcesOK() *GetResourcesOK {
	return &GetResourcesOK{}
}

/*
GetResourcesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetResourcesOK struct {
	Payload []*models.Resource
}

// IsSuccess returns true when this get resources o k response has a 2xx status code
func (o *GetResourcesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get resources o k response has a 3xx status code
func (o *GetResourcesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resources o k response has a 4xx status code
func (o *GetResourcesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get resources o k response has a 5xx status code
func (o *GetResourcesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get resources o k response a status code equal to that given
func (o *GetResourcesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetResourcesOK) Error() string {
	return fmt.Sprintf("[GET /policies/resources][%d] getResourcesOK  %+v", 200, o.Payload)
}

func (o *GetResourcesOK) String() string {
	return fmt.Sprintf("[GET /policies/resources][%d] getResourcesOK  %+v", 200, o.Payload)
}

func (o *GetResourcesOK) GetPayload() []*models.Resource {
	return o.Payload
}

func (o *GetResourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourcesUnauthorized creates a GetResourcesUnauthorized with default headers values
func NewGetResourcesUnauthorized() *GetResourcesUnauthorized {
	return &GetResourcesUnauthorized{}
}

/*
GetResourcesUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetResourcesUnauthorized struct {
}

// IsSuccess returns true when this get resources unauthorized response has a 2xx status code
func (o *GetResourcesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get resources unauthorized response has a 3xx status code
func (o *GetResourcesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resources unauthorized response has a 4xx status code
func (o *GetResourcesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get resources unauthorized response has a 5xx status code
func (o *GetResourcesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get resources unauthorized response a status code equal to that given
func (o *GetResourcesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetResourcesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /policies/resources][%d] getResourcesUnauthorized ", 401)
}

func (o *GetResourcesUnauthorized) String() string {
	return fmt.Sprintf("[GET /policies/resources][%d] getResourcesUnauthorized ", 401)
}

func (o *GetResourcesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetResourcesForbidden creates a GetResourcesForbidden with default headers values
func NewGetResourcesForbidden() *GetResourcesForbidden {
	return &GetResourcesForbidden{}
}

/*
GetResourcesForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetResourcesForbidden struct {
}

// IsSuccess returns true when this get resources forbidden response has a 2xx status code
func (o *GetResourcesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get resources forbidden response has a 3xx status code
func (o *GetResourcesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resources forbidden response has a 4xx status code
func (o *GetResourcesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get resources forbidden response has a 5xx status code
func (o *GetResourcesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get resources forbidden response a status code equal to that given
func (o *GetResourcesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetResourcesForbidden) Error() string {
	return fmt.Sprintf("[GET /policies/resources][%d] getResourcesForbidden ", 403)
}

func (o *GetResourcesForbidden) String() string {
	return fmt.Sprintf("[GET /policies/resources][%d] getResourcesForbidden ", 403)
}

func (o *GetResourcesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
