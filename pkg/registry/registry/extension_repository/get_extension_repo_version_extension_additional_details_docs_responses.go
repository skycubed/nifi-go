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

// GetExtensionRepoVersionExtensionAdditionalDetailsDocsReader is a Reader for the GetExtensionRepoVersionExtensionAdditionalDetailsDocs structure.
type GetExtensionRepoVersionExtensionAdditionalDetailsDocsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExtensionRepoVersionExtensionAdditionalDetailsDocsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExtensionRepoVersionExtensionAdditionalDetailsDocsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetExtensionRepoVersionExtensionAdditionalDetailsDocsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetExtensionRepoVersionExtensionAdditionalDetailsDocsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetExtensionRepoVersionExtensionAdditionalDetailsDocsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetExtensionRepoVersionExtensionAdditionalDetailsDocsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetExtensionRepoVersionExtensionAdditionalDetailsDocsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetExtensionRepoVersionExtensionAdditionalDetailsDocsOK creates a GetExtensionRepoVersionExtensionAdditionalDetailsDocsOK with default headers values
func NewGetExtensionRepoVersionExtensionAdditionalDetailsDocsOK() *GetExtensionRepoVersionExtensionAdditionalDetailsDocsOK {
	return &GetExtensionRepoVersionExtensionAdditionalDetailsDocsOK{}
}

/* GetExtensionRepoVersionExtensionAdditionalDetailsDocsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetExtensionRepoVersionExtensionAdditionalDetailsDocsOK struct {
	Payload string
}

func (o *GetExtensionRepoVersionExtensionAdditionalDetailsDocsOK) Error() string {
	return fmt.Sprintf("[GET /extension-repository/{bucketName}/{groupId}/{artifactId}/{version}/extensions/{name}/docs/additional-details][%d] getExtensionRepoVersionExtensionAdditionalDetailsDocsOK  %+v", 200, o.Payload)
}
func (o *GetExtensionRepoVersionExtensionAdditionalDetailsDocsOK) GetPayload() string {
	return o.Payload
}

func (o *GetExtensionRepoVersionExtensionAdditionalDetailsDocsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExtensionRepoVersionExtensionAdditionalDetailsDocsBadRequest creates a GetExtensionRepoVersionExtensionAdditionalDetailsDocsBadRequest with default headers values
func NewGetExtensionRepoVersionExtensionAdditionalDetailsDocsBadRequest() *GetExtensionRepoVersionExtensionAdditionalDetailsDocsBadRequest {
	return &GetExtensionRepoVersionExtensionAdditionalDetailsDocsBadRequest{}
}

/* GetExtensionRepoVersionExtensionAdditionalDetailsDocsBadRequest describes a response with status code 400, with default header values.

NiFi Registry was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetExtensionRepoVersionExtensionAdditionalDetailsDocsBadRequest struct {
}

func (o *GetExtensionRepoVersionExtensionAdditionalDetailsDocsBadRequest) Error() string {
	return fmt.Sprintf("[GET /extension-repository/{bucketName}/{groupId}/{artifactId}/{version}/extensions/{name}/docs/additional-details][%d] getExtensionRepoVersionExtensionAdditionalDetailsDocsBadRequest ", 400)
}

func (o *GetExtensionRepoVersionExtensionAdditionalDetailsDocsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExtensionRepoVersionExtensionAdditionalDetailsDocsUnauthorized creates a GetExtensionRepoVersionExtensionAdditionalDetailsDocsUnauthorized with default headers values
func NewGetExtensionRepoVersionExtensionAdditionalDetailsDocsUnauthorized() *GetExtensionRepoVersionExtensionAdditionalDetailsDocsUnauthorized {
	return &GetExtensionRepoVersionExtensionAdditionalDetailsDocsUnauthorized{}
}

/* GetExtensionRepoVersionExtensionAdditionalDetailsDocsUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetExtensionRepoVersionExtensionAdditionalDetailsDocsUnauthorized struct {
}

func (o *GetExtensionRepoVersionExtensionAdditionalDetailsDocsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /extension-repository/{bucketName}/{groupId}/{artifactId}/{version}/extensions/{name}/docs/additional-details][%d] getExtensionRepoVersionExtensionAdditionalDetailsDocsUnauthorized ", 401)
}

func (o *GetExtensionRepoVersionExtensionAdditionalDetailsDocsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExtensionRepoVersionExtensionAdditionalDetailsDocsForbidden creates a GetExtensionRepoVersionExtensionAdditionalDetailsDocsForbidden with default headers values
func NewGetExtensionRepoVersionExtensionAdditionalDetailsDocsForbidden() *GetExtensionRepoVersionExtensionAdditionalDetailsDocsForbidden {
	return &GetExtensionRepoVersionExtensionAdditionalDetailsDocsForbidden{}
}

/* GetExtensionRepoVersionExtensionAdditionalDetailsDocsForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetExtensionRepoVersionExtensionAdditionalDetailsDocsForbidden struct {
}

func (o *GetExtensionRepoVersionExtensionAdditionalDetailsDocsForbidden) Error() string {
	return fmt.Sprintf("[GET /extension-repository/{bucketName}/{groupId}/{artifactId}/{version}/extensions/{name}/docs/additional-details][%d] getExtensionRepoVersionExtensionAdditionalDetailsDocsForbidden ", 403)
}

func (o *GetExtensionRepoVersionExtensionAdditionalDetailsDocsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExtensionRepoVersionExtensionAdditionalDetailsDocsNotFound creates a GetExtensionRepoVersionExtensionAdditionalDetailsDocsNotFound with default headers values
func NewGetExtensionRepoVersionExtensionAdditionalDetailsDocsNotFound() *GetExtensionRepoVersionExtensionAdditionalDetailsDocsNotFound {
	return &GetExtensionRepoVersionExtensionAdditionalDetailsDocsNotFound{}
}

/* GetExtensionRepoVersionExtensionAdditionalDetailsDocsNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetExtensionRepoVersionExtensionAdditionalDetailsDocsNotFound struct {
}

func (o *GetExtensionRepoVersionExtensionAdditionalDetailsDocsNotFound) Error() string {
	return fmt.Sprintf("[GET /extension-repository/{bucketName}/{groupId}/{artifactId}/{version}/extensions/{name}/docs/additional-details][%d] getExtensionRepoVersionExtensionAdditionalDetailsDocsNotFound ", 404)
}

func (o *GetExtensionRepoVersionExtensionAdditionalDetailsDocsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExtensionRepoVersionExtensionAdditionalDetailsDocsConflict creates a GetExtensionRepoVersionExtensionAdditionalDetailsDocsConflict with default headers values
func NewGetExtensionRepoVersionExtensionAdditionalDetailsDocsConflict() *GetExtensionRepoVersionExtensionAdditionalDetailsDocsConflict {
	return &GetExtensionRepoVersionExtensionAdditionalDetailsDocsConflict{}
}

/* GetExtensionRepoVersionExtensionAdditionalDetailsDocsConflict describes a response with status code 409, with default header values.

NiFi Registry was unable to complete the request because it assumes a server state that is not valid.
*/
type GetExtensionRepoVersionExtensionAdditionalDetailsDocsConflict struct {
}

func (o *GetExtensionRepoVersionExtensionAdditionalDetailsDocsConflict) Error() string {
	return fmt.Sprintf("[GET /extension-repository/{bucketName}/{groupId}/{artifactId}/{version}/extensions/{name}/docs/additional-details][%d] getExtensionRepoVersionExtensionAdditionalDetailsDocsConflict ", 409)
}

func (o *GetExtensionRepoVersionExtensionAdditionalDetailsDocsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}