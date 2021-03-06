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

// GlobalGetExtensionBundleReader is a Reader for the GlobalGetExtensionBundle structure.
type GlobalGetExtensionBundleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GlobalGetExtensionBundleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGlobalGetExtensionBundleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGlobalGetExtensionBundleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGlobalGetExtensionBundleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGlobalGetExtensionBundleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGlobalGetExtensionBundleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGlobalGetExtensionBundleConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGlobalGetExtensionBundleOK creates a GlobalGetExtensionBundleOK with default headers values
func NewGlobalGetExtensionBundleOK() *GlobalGetExtensionBundleOK {
	return &GlobalGetExtensionBundleOK{}
}

/* GlobalGetExtensionBundleOK describes a response with status code 200, with default header values.

successful operation
*/
type GlobalGetExtensionBundleOK struct {
	Payload *models.ExtensionBundle
}

func (o *GlobalGetExtensionBundleOK) Error() string {
	return fmt.Sprintf("[GET /bundles/{bundleId}][%d] globalGetExtensionBundleOK  %+v", 200, o.Payload)
}
func (o *GlobalGetExtensionBundleOK) GetPayload() *models.ExtensionBundle {
	return o.Payload
}

func (o *GlobalGetExtensionBundleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExtensionBundle)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGlobalGetExtensionBundleBadRequest creates a GlobalGetExtensionBundleBadRequest with default headers values
func NewGlobalGetExtensionBundleBadRequest() *GlobalGetExtensionBundleBadRequest {
	return &GlobalGetExtensionBundleBadRequest{}
}

/* GlobalGetExtensionBundleBadRequest describes a response with status code 400, with default header values.

NiFi Registry was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GlobalGetExtensionBundleBadRequest struct {
}

func (o *GlobalGetExtensionBundleBadRequest) Error() string {
	return fmt.Sprintf("[GET /bundles/{bundleId}][%d] globalGetExtensionBundleBadRequest ", 400)
}

func (o *GlobalGetExtensionBundleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGlobalGetExtensionBundleUnauthorized creates a GlobalGetExtensionBundleUnauthorized with default headers values
func NewGlobalGetExtensionBundleUnauthorized() *GlobalGetExtensionBundleUnauthorized {
	return &GlobalGetExtensionBundleUnauthorized{}
}

/* GlobalGetExtensionBundleUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GlobalGetExtensionBundleUnauthorized struct {
}

func (o *GlobalGetExtensionBundleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /bundles/{bundleId}][%d] globalGetExtensionBundleUnauthorized ", 401)
}

func (o *GlobalGetExtensionBundleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGlobalGetExtensionBundleForbidden creates a GlobalGetExtensionBundleForbidden with default headers values
func NewGlobalGetExtensionBundleForbidden() *GlobalGetExtensionBundleForbidden {
	return &GlobalGetExtensionBundleForbidden{}
}

/* GlobalGetExtensionBundleForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GlobalGetExtensionBundleForbidden struct {
}

func (o *GlobalGetExtensionBundleForbidden) Error() string {
	return fmt.Sprintf("[GET /bundles/{bundleId}][%d] globalGetExtensionBundleForbidden ", 403)
}

func (o *GlobalGetExtensionBundleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGlobalGetExtensionBundleNotFound creates a GlobalGetExtensionBundleNotFound with default headers values
func NewGlobalGetExtensionBundleNotFound() *GlobalGetExtensionBundleNotFound {
	return &GlobalGetExtensionBundleNotFound{}
}

/* GlobalGetExtensionBundleNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GlobalGetExtensionBundleNotFound struct {
}

func (o *GlobalGetExtensionBundleNotFound) Error() string {
	return fmt.Sprintf("[GET /bundles/{bundleId}][%d] globalGetExtensionBundleNotFound ", 404)
}

func (o *GlobalGetExtensionBundleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGlobalGetExtensionBundleConflict creates a GlobalGetExtensionBundleConflict with default headers values
func NewGlobalGetExtensionBundleConflict() *GlobalGetExtensionBundleConflict {
	return &GlobalGetExtensionBundleConflict{}
}

/* GlobalGetExtensionBundleConflict describes a response with status code 409, with default header values.

NiFi Registry was unable to complete the request because it assumes a server state that is not valid.
*/
type GlobalGetExtensionBundleConflict struct {
}

func (o *GlobalGetExtensionBundleConflict) Error() string {
	return fmt.Sprintf("[GET /bundles/{bundleId}][%d] globalGetExtensionBundleConflict ", 409)
}

func (o *GlobalGetExtensionBundleConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
