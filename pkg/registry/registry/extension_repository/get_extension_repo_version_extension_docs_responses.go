// Code generated by go-swagger; DO NOT EDIT.

package extension_repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetExtensionRepoVersionExtensionDocsReader is a Reader for the GetExtensionRepoVersionExtensionDocs structure.
type GetExtensionRepoVersionExtensionDocsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExtensionRepoVersionExtensionDocsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExtensionRepoVersionExtensionDocsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetExtensionRepoVersionExtensionDocsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetExtensionRepoVersionExtensionDocsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetExtensionRepoVersionExtensionDocsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetExtensionRepoVersionExtensionDocsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetExtensionRepoVersionExtensionDocsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetExtensionRepoVersionExtensionDocsOK creates a GetExtensionRepoVersionExtensionDocsOK with default headers values
func NewGetExtensionRepoVersionExtensionDocsOK() *GetExtensionRepoVersionExtensionDocsOK {
	return &GetExtensionRepoVersionExtensionDocsOK{}
}

/* GetExtensionRepoVersionExtensionDocsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetExtensionRepoVersionExtensionDocsOK struct {
	Payload string
}

func (o *GetExtensionRepoVersionExtensionDocsOK) Error() string {
	return fmt.Sprintf("[GET /extension-repository/{bucketName}/{groupId}/{artifactId}/{version}/extensions/{name}/docs][%d] getExtensionRepoVersionExtensionDocsOK  %+v", 200, o.Payload)
}
func (o *GetExtensionRepoVersionExtensionDocsOK) GetPayload() string {
	return o.Payload
}

func (o *GetExtensionRepoVersionExtensionDocsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExtensionRepoVersionExtensionDocsBadRequest creates a GetExtensionRepoVersionExtensionDocsBadRequest with default headers values
func NewGetExtensionRepoVersionExtensionDocsBadRequest() *GetExtensionRepoVersionExtensionDocsBadRequest {
	return &GetExtensionRepoVersionExtensionDocsBadRequest{}
}

/* GetExtensionRepoVersionExtensionDocsBadRequest describes a response with status code 400, with default header values.

NiFi Registry was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetExtensionRepoVersionExtensionDocsBadRequest struct {
}

func (o *GetExtensionRepoVersionExtensionDocsBadRequest) Error() string {
	return fmt.Sprintf("[GET /extension-repository/{bucketName}/{groupId}/{artifactId}/{version}/extensions/{name}/docs][%d] getExtensionRepoVersionExtensionDocsBadRequest ", 400)
}

func (o *GetExtensionRepoVersionExtensionDocsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExtensionRepoVersionExtensionDocsUnauthorized creates a GetExtensionRepoVersionExtensionDocsUnauthorized with default headers values
func NewGetExtensionRepoVersionExtensionDocsUnauthorized() *GetExtensionRepoVersionExtensionDocsUnauthorized {
	return &GetExtensionRepoVersionExtensionDocsUnauthorized{}
}

/* GetExtensionRepoVersionExtensionDocsUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetExtensionRepoVersionExtensionDocsUnauthorized struct {
}

func (o *GetExtensionRepoVersionExtensionDocsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /extension-repository/{bucketName}/{groupId}/{artifactId}/{version}/extensions/{name}/docs][%d] getExtensionRepoVersionExtensionDocsUnauthorized ", 401)
}

func (o *GetExtensionRepoVersionExtensionDocsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExtensionRepoVersionExtensionDocsForbidden creates a GetExtensionRepoVersionExtensionDocsForbidden with default headers values
func NewGetExtensionRepoVersionExtensionDocsForbidden() *GetExtensionRepoVersionExtensionDocsForbidden {
	return &GetExtensionRepoVersionExtensionDocsForbidden{}
}

/* GetExtensionRepoVersionExtensionDocsForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetExtensionRepoVersionExtensionDocsForbidden struct {
}

func (o *GetExtensionRepoVersionExtensionDocsForbidden) Error() string {
	return fmt.Sprintf("[GET /extension-repository/{bucketName}/{groupId}/{artifactId}/{version}/extensions/{name}/docs][%d] getExtensionRepoVersionExtensionDocsForbidden ", 403)
}

func (o *GetExtensionRepoVersionExtensionDocsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExtensionRepoVersionExtensionDocsNotFound creates a GetExtensionRepoVersionExtensionDocsNotFound with default headers values
func NewGetExtensionRepoVersionExtensionDocsNotFound() *GetExtensionRepoVersionExtensionDocsNotFound {
	return &GetExtensionRepoVersionExtensionDocsNotFound{}
}

/* GetExtensionRepoVersionExtensionDocsNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetExtensionRepoVersionExtensionDocsNotFound struct {
}

func (o *GetExtensionRepoVersionExtensionDocsNotFound) Error() string {
	return fmt.Sprintf("[GET /extension-repository/{bucketName}/{groupId}/{artifactId}/{version}/extensions/{name}/docs][%d] getExtensionRepoVersionExtensionDocsNotFound ", 404)
}

func (o *GetExtensionRepoVersionExtensionDocsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExtensionRepoVersionExtensionDocsConflict creates a GetExtensionRepoVersionExtensionDocsConflict with default headers values
func NewGetExtensionRepoVersionExtensionDocsConflict() *GetExtensionRepoVersionExtensionDocsConflict {
	return &GetExtensionRepoVersionExtensionDocsConflict{}
}

/* GetExtensionRepoVersionExtensionDocsConflict describes a response with status code 409, with default header values.

NiFi Registry was unable to complete the request because it assumes a server state that is not valid.
*/
type GetExtensionRepoVersionExtensionDocsConflict struct {
}

func (o *GetExtensionRepoVersionExtensionDocsConflict) Error() string {
	return fmt.Sprintf("[GET /extension-repository/{bucketName}/{groupId}/{artifactId}/{version}/extensions/{name}/docs][%d] getExtensionRepoVersionExtensionDocsConflict ", 409)
}

func (o *GetExtensionRepoVersionExtensionDocsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
