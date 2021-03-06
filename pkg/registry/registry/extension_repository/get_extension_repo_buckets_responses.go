// Code generated by go-swagger; DO NOT EDIT.

package extension_repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/registry/models"
)

// GetExtensionRepoBucketsReader is a Reader for the GetExtensionRepoBuckets structure.
type GetExtensionRepoBucketsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExtensionRepoBucketsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExtensionRepoBucketsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetExtensionRepoBucketsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetExtensionRepoBucketsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetExtensionRepoBucketsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetExtensionRepoBucketsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetExtensionRepoBucketsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetExtensionRepoBucketsOK creates a GetExtensionRepoBucketsOK with default headers values
func NewGetExtensionRepoBucketsOK() *GetExtensionRepoBucketsOK {
	return &GetExtensionRepoBucketsOK{}
}

/* GetExtensionRepoBucketsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetExtensionRepoBucketsOK struct {
	Payload []*models.ExtensionRepoBucket
}

func (o *GetExtensionRepoBucketsOK) Error() string {
	return fmt.Sprintf("[GET /extension-repository][%d] getExtensionRepoBucketsOK  %+v", 200, o.Payload)
}
func (o *GetExtensionRepoBucketsOK) GetPayload() []*models.ExtensionRepoBucket {
	return o.Payload
}

func (o *GetExtensionRepoBucketsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExtensionRepoBucketsBadRequest creates a GetExtensionRepoBucketsBadRequest with default headers values
func NewGetExtensionRepoBucketsBadRequest() *GetExtensionRepoBucketsBadRequest {
	return &GetExtensionRepoBucketsBadRequest{}
}

/* GetExtensionRepoBucketsBadRequest describes a response with status code 400, with default header values.

NiFi Registry was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetExtensionRepoBucketsBadRequest struct {
}

func (o *GetExtensionRepoBucketsBadRequest) Error() string {
	return fmt.Sprintf("[GET /extension-repository][%d] getExtensionRepoBucketsBadRequest ", 400)
}

func (o *GetExtensionRepoBucketsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExtensionRepoBucketsUnauthorized creates a GetExtensionRepoBucketsUnauthorized with default headers values
func NewGetExtensionRepoBucketsUnauthorized() *GetExtensionRepoBucketsUnauthorized {
	return &GetExtensionRepoBucketsUnauthorized{}
}

/* GetExtensionRepoBucketsUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetExtensionRepoBucketsUnauthorized struct {
}

func (o *GetExtensionRepoBucketsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /extension-repository][%d] getExtensionRepoBucketsUnauthorized ", 401)
}

func (o *GetExtensionRepoBucketsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExtensionRepoBucketsForbidden creates a GetExtensionRepoBucketsForbidden with default headers values
func NewGetExtensionRepoBucketsForbidden() *GetExtensionRepoBucketsForbidden {
	return &GetExtensionRepoBucketsForbidden{}
}

/* GetExtensionRepoBucketsForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetExtensionRepoBucketsForbidden struct {
}

func (o *GetExtensionRepoBucketsForbidden) Error() string {
	return fmt.Sprintf("[GET /extension-repository][%d] getExtensionRepoBucketsForbidden ", 403)
}

func (o *GetExtensionRepoBucketsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExtensionRepoBucketsNotFound creates a GetExtensionRepoBucketsNotFound with default headers values
func NewGetExtensionRepoBucketsNotFound() *GetExtensionRepoBucketsNotFound {
	return &GetExtensionRepoBucketsNotFound{}
}

/* GetExtensionRepoBucketsNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetExtensionRepoBucketsNotFound struct {
}

func (o *GetExtensionRepoBucketsNotFound) Error() string {
	return fmt.Sprintf("[GET /extension-repository][%d] getExtensionRepoBucketsNotFound ", 404)
}

func (o *GetExtensionRepoBucketsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExtensionRepoBucketsConflict creates a GetExtensionRepoBucketsConflict with default headers values
func NewGetExtensionRepoBucketsConflict() *GetExtensionRepoBucketsConflict {
	return &GetExtensionRepoBucketsConflict{}
}

/* GetExtensionRepoBucketsConflict describes a response with status code 409, with default header values.

NiFi Registry was unable to complete the request because it assumes a server state that is not valid.
*/
type GetExtensionRepoBucketsConflict struct {
}

func (o *GetExtensionRepoBucketsConflict) Error() string {
	return fmt.Sprintf("[GET /extension-repository][%d] getExtensionRepoBucketsConflict ", 409)
}

func (o *GetExtensionRepoBucketsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
