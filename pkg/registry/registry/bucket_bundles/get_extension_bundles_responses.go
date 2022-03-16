// Code generated by go-swagger; DO NOT EDIT.

package bucket_bundles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/registry/models"
)

// GetExtensionBundlesReader is a Reader for the GetExtensionBundles structure.
type GetExtensionBundlesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExtensionBundlesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExtensionBundlesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetExtensionBundlesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetExtensionBundlesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetExtensionBundlesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetExtensionBundlesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetExtensionBundlesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetExtensionBundlesOK creates a GetExtensionBundlesOK with default headers values
func NewGetExtensionBundlesOK() *GetExtensionBundlesOK {
	return &GetExtensionBundlesOK{}
}

/* GetExtensionBundlesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetExtensionBundlesOK struct {
	Payload []*models.ExtensionBundle
}

func (o *GetExtensionBundlesOK) Error() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}/bundles][%d] getExtensionBundlesOK  %+v", 200, o.Payload)
}
func (o *GetExtensionBundlesOK) GetPayload() []*models.ExtensionBundle {
	return o.Payload
}

func (o *GetExtensionBundlesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExtensionBundlesBadRequest creates a GetExtensionBundlesBadRequest with default headers values
func NewGetExtensionBundlesBadRequest() *GetExtensionBundlesBadRequest {
	return &GetExtensionBundlesBadRequest{}
}

/* GetExtensionBundlesBadRequest describes a response with status code 400, with default header values.

NiFi Registry was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetExtensionBundlesBadRequest struct {
}

func (o *GetExtensionBundlesBadRequest) Error() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}/bundles][%d] getExtensionBundlesBadRequest ", 400)
}

func (o *GetExtensionBundlesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExtensionBundlesUnauthorized creates a GetExtensionBundlesUnauthorized with default headers values
func NewGetExtensionBundlesUnauthorized() *GetExtensionBundlesUnauthorized {
	return &GetExtensionBundlesUnauthorized{}
}

/* GetExtensionBundlesUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetExtensionBundlesUnauthorized struct {
}

func (o *GetExtensionBundlesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}/bundles][%d] getExtensionBundlesUnauthorized ", 401)
}

func (o *GetExtensionBundlesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExtensionBundlesForbidden creates a GetExtensionBundlesForbidden with default headers values
func NewGetExtensionBundlesForbidden() *GetExtensionBundlesForbidden {
	return &GetExtensionBundlesForbidden{}
}

/* GetExtensionBundlesForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetExtensionBundlesForbidden struct {
}

func (o *GetExtensionBundlesForbidden) Error() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}/bundles][%d] getExtensionBundlesForbidden ", 403)
}

func (o *GetExtensionBundlesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExtensionBundlesNotFound creates a GetExtensionBundlesNotFound with default headers values
func NewGetExtensionBundlesNotFound() *GetExtensionBundlesNotFound {
	return &GetExtensionBundlesNotFound{}
}

/* GetExtensionBundlesNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetExtensionBundlesNotFound struct {
}

func (o *GetExtensionBundlesNotFound) Error() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}/bundles][%d] getExtensionBundlesNotFound ", 404)
}

func (o *GetExtensionBundlesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExtensionBundlesConflict creates a GetExtensionBundlesConflict with default headers values
func NewGetExtensionBundlesConflict() *GetExtensionBundlesConflict {
	return &GetExtensionBundlesConflict{}
}

/* GetExtensionBundlesConflict describes a response with status code 409, with default header values.

NiFi Registry was unable to complete the request because it assumes a server state that is not valid.
*/
type GetExtensionBundlesConflict struct {
}

func (o *GetExtensionBundlesConflict) Error() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}/bundles][%d] getExtensionBundlesConflict ", 409)
}

func (o *GetExtensionBundlesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}