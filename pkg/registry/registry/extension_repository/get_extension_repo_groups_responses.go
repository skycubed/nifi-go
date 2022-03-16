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

// GetExtensionRepoGroupsReader is a Reader for the GetExtensionRepoGroups structure.
type GetExtensionRepoGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExtensionRepoGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExtensionRepoGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetExtensionRepoGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetExtensionRepoGroupsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetExtensionRepoGroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetExtensionRepoGroupsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetExtensionRepoGroupsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetExtensionRepoGroupsOK creates a GetExtensionRepoGroupsOK with default headers values
func NewGetExtensionRepoGroupsOK() *GetExtensionRepoGroupsOK {
	return &GetExtensionRepoGroupsOK{}
}

/* GetExtensionRepoGroupsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetExtensionRepoGroupsOK struct {
	Payload []*models.ExtensionRepoGroup
}

func (o *GetExtensionRepoGroupsOK) Error() string {
	return fmt.Sprintf("[GET /extension-repository/{bucketName}][%d] getExtensionRepoGroupsOK  %+v", 200, o.Payload)
}
func (o *GetExtensionRepoGroupsOK) GetPayload() []*models.ExtensionRepoGroup {
	return o.Payload
}

func (o *GetExtensionRepoGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExtensionRepoGroupsBadRequest creates a GetExtensionRepoGroupsBadRequest with default headers values
func NewGetExtensionRepoGroupsBadRequest() *GetExtensionRepoGroupsBadRequest {
	return &GetExtensionRepoGroupsBadRequest{}
}

/* GetExtensionRepoGroupsBadRequest describes a response with status code 400, with default header values.

NiFi Registry was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetExtensionRepoGroupsBadRequest struct {
}

func (o *GetExtensionRepoGroupsBadRequest) Error() string {
	return fmt.Sprintf("[GET /extension-repository/{bucketName}][%d] getExtensionRepoGroupsBadRequest ", 400)
}

func (o *GetExtensionRepoGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExtensionRepoGroupsUnauthorized creates a GetExtensionRepoGroupsUnauthorized with default headers values
func NewGetExtensionRepoGroupsUnauthorized() *GetExtensionRepoGroupsUnauthorized {
	return &GetExtensionRepoGroupsUnauthorized{}
}

/* GetExtensionRepoGroupsUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetExtensionRepoGroupsUnauthorized struct {
}

func (o *GetExtensionRepoGroupsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /extension-repository/{bucketName}][%d] getExtensionRepoGroupsUnauthorized ", 401)
}

func (o *GetExtensionRepoGroupsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExtensionRepoGroupsForbidden creates a GetExtensionRepoGroupsForbidden with default headers values
func NewGetExtensionRepoGroupsForbidden() *GetExtensionRepoGroupsForbidden {
	return &GetExtensionRepoGroupsForbidden{}
}

/* GetExtensionRepoGroupsForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetExtensionRepoGroupsForbidden struct {
}

func (o *GetExtensionRepoGroupsForbidden) Error() string {
	return fmt.Sprintf("[GET /extension-repository/{bucketName}][%d] getExtensionRepoGroupsForbidden ", 403)
}

func (o *GetExtensionRepoGroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExtensionRepoGroupsNotFound creates a GetExtensionRepoGroupsNotFound with default headers values
func NewGetExtensionRepoGroupsNotFound() *GetExtensionRepoGroupsNotFound {
	return &GetExtensionRepoGroupsNotFound{}
}

/* GetExtensionRepoGroupsNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetExtensionRepoGroupsNotFound struct {
}

func (o *GetExtensionRepoGroupsNotFound) Error() string {
	return fmt.Sprintf("[GET /extension-repository/{bucketName}][%d] getExtensionRepoGroupsNotFound ", 404)
}

func (o *GetExtensionRepoGroupsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExtensionRepoGroupsConflict creates a GetExtensionRepoGroupsConflict with default headers values
func NewGetExtensionRepoGroupsConflict() *GetExtensionRepoGroupsConflict {
	return &GetExtensionRepoGroupsConflict{}
}

/* GetExtensionRepoGroupsConflict describes a response with status code 409, with default header values.

NiFi Registry was unable to complete the request because it assumes a server state that is not valid.
*/
type GetExtensionRepoGroupsConflict struct {
}

func (o *GetExtensionRepoGroupsConflict) Error() string {
	return fmt.Sprintf("[GET /extension-repository/{bucketName}][%d] getExtensionRepoGroupsConflict ", 409)
}

func (o *GetExtensionRepoGroupsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}