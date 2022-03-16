// Code generated by go-swagger; DO NOT EDIT.

package bundles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/registry/models"
)

// GlobalGetBundleVersionExtensionsReader is a Reader for the GlobalGetBundleVersionExtensions structure.
type GlobalGetBundleVersionExtensionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GlobalGetBundleVersionExtensionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGlobalGetBundleVersionExtensionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGlobalGetBundleVersionExtensionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGlobalGetBundleVersionExtensionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGlobalGetBundleVersionExtensionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGlobalGetBundleVersionExtensionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGlobalGetBundleVersionExtensionsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGlobalGetBundleVersionExtensionsOK creates a GlobalGetBundleVersionExtensionsOK with default headers values
func NewGlobalGetBundleVersionExtensionsOK() *GlobalGetBundleVersionExtensionsOK {
	return &GlobalGetBundleVersionExtensionsOK{}
}

/* GlobalGetBundleVersionExtensionsOK describes a response with status code 200, with default header values.

successful operation
*/
type GlobalGetBundleVersionExtensionsOK struct {
	Payload []*models.ExtensionMetadata
}

func (o *GlobalGetBundleVersionExtensionsOK) Error() string {
	return fmt.Sprintf("[GET /bundles/{bundleId}/versions/{version}/extensions][%d] globalGetBundleVersionExtensionsOK  %+v", 200, o.Payload)
}
func (o *GlobalGetBundleVersionExtensionsOK) GetPayload() []*models.ExtensionMetadata {
	return o.Payload
}

func (o *GlobalGetBundleVersionExtensionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGlobalGetBundleVersionExtensionsBadRequest creates a GlobalGetBundleVersionExtensionsBadRequest with default headers values
func NewGlobalGetBundleVersionExtensionsBadRequest() *GlobalGetBundleVersionExtensionsBadRequest {
	return &GlobalGetBundleVersionExtensionsBadRequest{}
}

/* GlobalGetBundleVersionExtensionsBadRequest describes a response with status code 400, with default header values.

NiFi Registry was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GlobalGetBundleVersionExtensionsBadRequest struct {
}

func (o *GlobalGetBundleVersionExtensionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /bundles/{bundleId}/versions/{version}/extensions][%d] globalGetBundleVersionExtensionsBadRequest ", 400)
}

func (o *GlobalGetBundleVersionExtensionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGlobalGetBundleVersionExtensionsUnauthorized creates a GlobalGetBundleVersionExtensionsUnauthorized with default headers values
func NewGlobalGetBundleVersionExtensionsUnauthorized() *GlobalGetBundleVersionExtensionsUnauthorized {
	return &GlobalGetBundleVersionExtensionsUnauthorized{}
}

/* GlobalGetBundleVersionExtensionsUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GlobalGetBundleVersionExtensionsUnauthorized struct {
}

func (o *GlobalGetBundleVersionExtensionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /bundles/{bundleId}/versions/{version}/extensions][%d] globalGetBundleVersionExtensionsUnauthorized ", 401)
}

func (o *GlobalGetBundleVersionExtensionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGlobalGetBundleVersionExtensionsForbidden creates a GlobalGetBundleVersionExtensionsForbidden with default headers values
func NewGlobalGetBundleVersionExtensionsForbidden() *GlobalGetBundleVersionExtensionsForbidden {
	return &GlobalGetBundleVersionExtensionsForbidden{}
}

/* GlobalGetBundleVersionExtensionsForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GlobalGetBundleVersionExtensionsForbidden struct {
}

func (o *GlobalGetBundleVersionExtensionsForbidden) Error() string {
	return fmt.Sprintf("[GET /bundles/{bundleId}/versions/{version}/extensions][%d] globalGetBundleVersionExtensionsForbidden ", 403)
}

func (o *GlobalGetBundleVersionExtensionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGlobalGetBundleVersionExtensionsNotFound creates a GlobalGetBundleVersionExtensionsNotFound with default headers values
func NewGlobalGetBundleVersionExtensionsNotFound() *GlobalGetBundleVersionExtensionsNotFound {
	return &GlobalGetBundleVersionExtensionsNotFound{}
}

/* GlobalGetBundleVersionExtensionsNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GlobalGetBundleVersionExtensionsNotFound struct {
}

func (o *GlobalGetBundleVersionExtensionsNotFound) Error() string {
	return fmt.Sprintf("[GET /bundles/{bundleId}/versions/{version}/extensions][%d] globalGetBundleVersionExtensionsNotFound ", 404)
}

func (o *GlobalGetBundleVersionExtensionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGlobalGetBundleVersionExtensionsConflict creates a GlobalGetBundleVersionExtensionsConflict with default headers values
func NewGlobalGetBundleVersionExtensionsConflict() *GlobalGetBundleVersionExtensionsConflict {
	return &GlobalGetBundleVersionExtensionsConflict{}
}

/* GlobalGetBundleVersionExtensionsConflict describes a response with status code 409, with default header values.

NiFi Registry was unable to complete the request because it assumes a server state that is not valid.
*/
type GlobalGetBundleVersionExtensionsConflict struct {
}

func (o *GlobalGetBundleVersionExtensionsConflict) Error() string {
	return fmt.Sprintf("[GET /bundles/{bundleId}/versions/{version}/extensions][%d] globalGetBundleVersionExtensionsConflict ", 409)
}

func (o *GlobalGetBundleVersionExtensionsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
