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

// GetExtensionRepoVersionsReader is a Reader for the GetExtensionRepoVersions structure.
type GetExtensionRepoVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExtensionRepoVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExtensionRepoVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetExtensionRepoVersionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetExtensionRepoVersionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetExtensionRepoVersionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetExtensionRepoVersionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetExtensionRepoVersionsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetExtensionRepoVersionsOK creates a GetExtensionRepoVersionsOK with default headers values
func NewGetExtensionRepoVersionsOK() *GetExtensionRepoVersionsOK {
	return &GetExtensionRepoVersionsOK{}
}

/* GetExtensionRepoVersionsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetExtensionRepoVersionsOK struct {
	Payload []*models.ExtensionRepoVersionSummary
}

func (o *GetExtensionRepoVersionsOK) Error() string {
	return fmt.Sprintf("[GET /extension-repository/{bucketName}/{groupId}/{artifactId}][%d] getExtensionRepoVersionsOK  %+v", 200, o.Payload)
}
func (o *GetExtensionRepoVersionsOK) GetPayload() []*models.ExtensionRepoVersionSummary {
	return o.Payload
}

func (o *GetExtensionRepoVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExtensionRepoVersionsBadRequest creates a GetExtensionRepoVersionsBadRequest with default headers values
func NewGetExtensionRepoVersionsBadRequest() *GetExtensionRepoVersionsBadRequest {
	return &GetExtensionRepoVersionsBadRequest{}
}

/* GetExtensionRepoVersionsBadRequest describes a response with status code 400, with default header values.

NiFi Registry was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetExtensionRepoVersionsBadRequest struct {
}

func (o *GetExtensionRepoVersionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /extension-repository/{bucketName}/{groupId}/{artifactId}][%d] getExtensionRepoVersionsBadRequest ", 400)
}

func (o *GetExtensionRepoVersionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExtensionRepoVersionsUnauthorized creates a GetExtensionRepoVersionsUnauthorized with default headers values
func NewGetExtensionRepoVersionsUnauthorized() *GetExtensionRepoVersionsUnauthorized {
	return &GetExtensionRepoVersionsUnauthorized{}
}

/* GetExtensionRepoVersionsUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetExtensionRepoVersionsUnauthorized struct {
}

func (o *GetExtensionRepoVersionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /extension-repository/{bucketName}/{groupId}/{artifactId}][%d] getExtensionRepoVersionsUnauthorized ", 401)
}

func (o *GetExtensionRepoVersionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExtensionRepoVersionsForbidden creates a GetExtensionRepoVersionsForbidden with default headers values
func NewGetExtensionRepoVersionsForbidden() *GetExtensionRepoVersionsForbidden {
	return &GetExtensionRepoVersionsForbidden{}
}

/* GetExtensionRepoVersionsForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetExtensionRepoVersionsForbidden struct {
}

func (o *GetExtensionRepoVersionsForbidden) Error() string {
	return fmt.Sprintf("[GET /extension-repository/{bucketName}/{groupId}/{artifactId}][%d] getExtensionRepoVersionsForbidden ", 403)
}

func (o *GetExtensionRepoVersionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExtensionRepoVersionsNotFound creates a GetExtensionRepoVersionsNotFound with default headers values
func NewGetExtensionRepoVersionsNotFound() *GetExtensionRepoVersionsNotFound {
	return &GetExtensionRepoVersionsNotFound{}
}

/* GetExtensionRepoVersionsNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetExtensionRepoVersionsNotFound struct {
}

func (o *GetExtensionRepoVersionsNotFound) Error() string {
	return fmt.Sprintf("[GET /extension-repository/{bucketName}/{groupId}/{artifactId}][%d] getExtensionRepoVersionsNotFound ", 404)
}

func (o *GetExtensionRepoVersionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExtensionRepoVersionsConflict creates a GetExtensionRepoVersionsConflict with default headers values
func NewGetExtensionRepoVersionsConflict() *GetExtensionRepoVersionsConflict {
	return &GetExtensionRepoVersionsConflict{}
}

/* GetExtensionRepoVersionsConflict describes a response with status code 409, with default header values.

NiFi Registry was unable to complete the request because it assumes a server state that is not valid.
*/
type GetExtensionRepoVersionsConflict struct {
}

func (o *GetExtensionRepoVersionsConflict) Error() string {
	return fmt.Sprintf("[GET /extension-repository/{bucketName}/{groupId}/{artifactId}][%d] getExtensionRepoVersionsConflict ", 409)
}

func (o *GetExtensionRepoVersionsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
