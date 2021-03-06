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

// GetGlobalExtensionRepoVersionSha256Reader is a Reader for the GetGlobalExtensionRepoVersionSha256 structure.
type GetGlobalExtensionRepoVersionSha256Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGlobalExtensionRepoVersionSha256Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGlobalExtensionRepoVersionSha256OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGlobalExtensionRepoVersionSha256BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetGlobalExtensionRepoVersionSha256Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGlobalExtensionRepoVersionSha256Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGlobalExtensionRepoVersionSha256NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetGlobalExtensionRepoVersionSha256Conflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGlobalExtensionRepoVersionSha256OK creates a GetGlobalExtensionRepoVersionSha256OK with default headers values
func NewGetGlobalExtensionRepoVersionSha256OK() *GetGlobalExtensionRepoVersionSha256OK {
	return &GetGlobalExtensionRepoVersionSha256OK{}
}

/* GetGlobalExtensionRepoVersionSha256OK describes a response with status code 200, with default header values.

successful operation
*/
type GetGlobalExtensionRepoVersionSha256OK struct {
	Payload string
}

func (o *GetGlobalExtensionRepoVersionSha256OK) Error() string {
	return fmt.Sprintf("[GET /extension-repository/{groupId}/{artifactId}/{version}/sha256][%d] getGlobalExtensionRepoVersionSha256OK  %+v", 200, o.Payload)
}
func (o *GetGlobalExtensionRepoVersionSha256OK) GetPayload() string {
	return o.Payload
}

func (o *GetGlobalExtensionRepoVersionSha256OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGlobalExtensionRepoVersionSha256BadRequest creates a GetGlobalExtensionRepoVersionSha256BadRequest with default headers values
func NewGetGlobalExtensionRepoVersionSha256BadRequest() *GetGlobalExtensionRepoVersionSha256BadRequest {
	return &GetGlobalExtensionRepoVersionSha256BadRequest{}
}

/* GetGlobalExtensionRepoVersionSha256BadRequest describes a response with status code 400, with default header values.

NiFi Registry was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetGlobalExtensionRepoVersionSha256BadRequest struct {
}

func (o *GetGlobalExtensionRepoVersionSha256BadRequest) Error() string {
	return fmt.Sprintf("[GET /extension-repository/{groupId}/{artifactId}/{version}/sha256][%d] getGlobalExtensionRepoVersionSha256BadRequest ", 400)
}

func (o *GetGlobalExtensionRepoVersionSha256BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGlobalExtensionRepoVersionSha256Unauthorized creates a GetGlobalExtensionRepoVersionSha256Unauthorized with default headers values
func NewGetGlobalExtensionRepoVersionSha256Unauthorized() *GetGlobalExtensionRepoVersionSha256Unauthorized {
	return &GetGlobalExtensionRepoVersionSha256Unauthorized{}
}

/* GetGlobalExtensionRepoVersionSha256Unauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetGlobalExtensionRepoVersionSha256Unauthorized struct {
}

func (o *GetGlobalExtensionRepoVersionSha256Unauthorized) Error() string {
	return fmt.Sprintf("[GET /extension-repository/{groupId}/{artifactId}/{version}/sha256][%d] getGlobalExtensionRepoVersionSha256Unauthorized ", 401)
}

func (o *GetGlobalExtensionRepoVersionSha256Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGlobalExtensionRepoVersionSha256Forbidden creates a GetGlobalExtensionRepoVersionSha256Forbidden with default headers values
func NewGetGlobalExtensionRepoVersionSha256Forbidden() *GetGlobalExtensionRepoVersionSha256Forbidden {
	return &GetGlobalExtensionRepoVersionSha256Forbidden{}
}

/* GetGlobalExtensionRepoVersionSha256Forbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetGlobalExtensionRepoVersionSha256Forbidden struct {
}

func (o *GetGlobalExtensionRepoVersionSha256Forbidden) Error() string {
	return fmt.Sprintf("[GET /extension-repository/{groupId}/{artifactId}/{version}/sha256][%d] getGlobalExtensionRepoVersionSha256Forbidden ", 403)
}

func (o *GetGlobalExtensionRepoVersionSha256Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGlobalExtensionRepoVersionSha256NotFound creates a GetGlobalExtensionRepoVersionSha256NotFound with default headers values
func NewGetGlobalExtensionRepoVersionSha256NotFound() *GetGlobalExtensionRepoVersionSha256NotFound {
	return &GetGlobalExtensionRepoVersionSha256NotFound{}
}

/* GetGlobalExtensionRepoVersionSha256NotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetGlobalExtensionRepoVersionSha256NotFound struct {
}

func (o *GetGlobalExtensionRepoVersionSha256NotFound) Error() string {
	return fmt.Sprintf("[GET /extension-repository/{groupId}/{artifactId}/{version}/sha256][%d] getGlobalExtensionRepoVersionSha256NotFound ", 404)
}

func (o *GetGlobalExtensionRepoVersionSha256NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGlobalExtensionRepoVersionSha256Conflict creates a GetGlobalExtensionRepoVersionSha256Conflict with default headers values
func NewGetGlobalExtensionRepoVersionSha256Conflict() *GetGlobalExtensionRepoVersionSha256Conflict {
	return &GetGlobalExtensionRepoVersionSha256Conflict{}
}

/* GetGlobalExtensionRepoVersionSha256Conflict describes a response with status code 409, with default header values.

NiFi Registry was unable to complete the request because it assumes a server state that is not valid.
*/
type GetGlobalExtensionRepoVersionSha256Conflict struct {
}

func (o *GetGlobalExtensionRepoVersionSha256Conflict) Error() string {
	return fmt.Sprintf("[GET /extension-repository/{groupId}/{artifactId}/{version}/sha256][%d] getGlobalExtensionRepoVersionSha256Conflict ", 409)
}

func (o *GetGlobalExtensionRepoVersionSha256Conflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
