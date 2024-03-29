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
		return nil, runtime.NewAPIError("[GET /extension-repository/{bucketName}/{groupId}/{artifactId}/{version}/extensions/{name}/docs] getExtensionRepoVersionExtensionDocs", response, response.Code())
	}
}

// NewGetExtensionRepoVersionExtensionDocsOK creates a GetExtensionRepoVersionExtensionDocsOK with default headers values
func NewGetExtensionRepoVersionExtensionDocsOK() *GetExtensionRepoVersionExtensionDocsOK {
	return &GetExtensionRepoVersionExtensionDocsOK{}
}

/*
GetExtensionRepoVersionExtensionDocsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetExtensionRepoVersionExtensionDocsOK struct {
	Payload string
}

// IsSuccess returns true when this get extension repo version extension docs o k response has a 2xx status code
func (o *GetExtensionRepoVersionExtensionDocsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get extension repo version extension docs o k response has a 3xx status code
func (o *GetExtensionRepoVersionExtensionDocsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get extension repo version extension docs o k response has a 4xx status code
func (o *GetExtensionRepoVersionExtensionDocsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get extension repo version extension docs o k response has a 5xx status code
func (o *GetExtensionRepoVersionExtensionDocsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get extension repo version extension docs o k response a status code equal to that given
func (o *GetExtensionRepoVersionExtensionDocsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get extension repo version extension docs o k response
func (o *GetExtensionRepoVersionExtensionDocsOK) Code() int {
	return 200
}

func (o *GetExtensionRepoVersionExtensionDocsOK) Error() string {
	return fmt.Sprintf("[GET /extension-repository/{bucketName}/{groupId}/{artifactId}/{version}/extensions/{name}/docs][%d] getExtensionRepoVersionExtensionDocsOK  %+v", 200, o.Payload)
}

func (o *GetExtensionRepoVersionExtensionDocsOK) String() string {
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

/*
GetExtensionRepoVersionExtensionDocsBadRequest describes a response with status code 400, with default header values.

NiFi Registry was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetExtensionRepoVersionExtensionDocsBadRequest struct {
}

// IsSuccess returns true when this get extension repo version extension docs bad request response has a 2xx status code
func (o *GetExtensionRepoVersionExtensionDocsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get extension repo version extension docs bad request response has a 3xx status code
func (o *GetExtensionRepoVersionExtensionDocsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get extension repo version extension docs bad request response has a 4xx status code
func (o *GetExtensionRepoVersionExtensionDocsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get extension repo version extension docs bad request response has a 5xx status code
func (o *GetExtensionRepoVersionExtensionDocsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get extension repo version extension docs bad request response a status code equal to that given
func (o *GetExtensionRepoVersionExtensionDocsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get extension repo version extension docs bad request response
func (o *GetExtensionRepoVersionExtensionDocsBadRequest) Code() int {
	return 400
}

func (o *GetExtensionRepoVersionExtensionDocsBadRequest) Error() string {
	return fmt.Sprintf("[GET /extension-repository/{bucketName}/{groupId}/{artifactId}/{version}/extensions/{name}/docs][%d] getExtensionRepoVersionExtensionDocsBadRequest ", 400)
}

func (o *GetExtensionRepoVersionExtensionDocsBadRequest) String() string {
	return fmt.Sprintf("[GET /extension-repository/{bucketName}/{groupId}/{artifactId}/{version}/extensions/{name}/docs][%d] getExtensionRepoVersionExtensionDocsBadRequest ", 400)
}

func (o *GetExtensionRepoVersionExtensionDocsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExtensionRepoVersionExtensionDocsUnauthorized creates a GetExtensionRepoVersionExtensionDocsUnauthorized with default headers values
func NewGetExtensionRepoVersionExtensionDocsUnauthorized() *GetExtensionRepoVersionExtensionDocsUnauthorized {
	return &GetExtensionRepoVersionExtensionDocsUnauthorized{}
}

/*
GetExtensionRepoVersionExtensionDocsUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetExtensionRepoVersionExtensionDocsUnauthorized struct {
}

// IsSuccess returns true when this get extension repo version extension docs unauthorized response has a 2xx status code
func (o *GetExtensionRepoVersionExtensionDocsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get extension repo version extension docs unauthorized response has a 3xx status code
func (o *GetExtensionRepoVersionExtensionDocsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get extension repo version extension docs unauthorized response has a 4xx status code
func (o *GetExtensionRepoVersionExtensionDocsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get extension repo version extension docs unauthorized response has a 5xx status code
func (o *GetExtensionRepoVersionExtensionDocsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get extension repo version extension docs unauthorized response a status code equal to that given
func (o *GetExtensionRepoVersionExtensionDocsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get extension repo version extension docs unauthorized response
func (o *GetExtensionRepoVersionExtensionDocsUnauthorized) Code() int {
	return 401
}

func (o *GetExtensionRepoVersionExtensionDocsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /extension-repository/{bucketName}/{groupId}/{artifactId}/{version}/extensions/{name}/docs][%d] getExtensionRepoVersionExtensionDocsUnauthorized ", 401)
}

func (o *GetExtensionRepoVersionExtensionDocsUnauthorized) String() string {
	return fmt.Sprintf("[GET /extension-repository/{bucketName}/{groupId}/{artifactId}/{version}/extensions/{name}/docs][%d] getExtensionRepoVersionExtensionDocsUnauthorized ", 401)
}

func (o *GetExtensionRepoVersionExtensionDocsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExtensionRepoVersionExtensionDocsForbidden creates a GetExtensionRepoVersionExtensionDocsForbidden with default headers values
func NewGetExtensionRepoVersionExtensionDocsForbidden() *GetExtensionRepoVersionExtensionDocsForbidden {
	return &GetExtensionRepoVersionExtensionDocsForbidden{}
}

/*
GetExtensionRepoVersionExtensionDocsForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetExtensionRepoVersionExtensionDocsForbidden struct {
}

// IsSuccess returns true when this get extension repo version extension docs forbidden response has a 2xx status code
func (o *GetExtensionRepoVersionExtensionDocsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get extension repo version extension docs forbidden response has a 3xx status code
func (o *GetExtensionRepoVersionExtensionDocsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get extension repo version extension docs forbidden response has a 4xx status code
func (o *GetExtensionRepoVersionExtensionDocsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get extension repo version extension docs forbidden response has a 5xx status code
func (o *GetExtensionRepoVersionExtensionDocsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get extension repo version extension docs forbidden response a status code equal to that given
func (o *GetExtensionRepoVersionExtensionDocsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get extension repo version extension docs forbidden response
func (o *GetExtensionRepoVersionExtensionDocsForbidden) Code() int {
	return 403
}

func (o *GetExtensionRepoVersionExtensionDocsForbidden) Error() string {
	return fmt.Sprintf("[GET /extension-repository/{bucketName}/{groupId}/{artifactId}/{version}/extensions/{name}/docs][%d] getExtensionRepoVersionExtensionDocsForbidden ", 403)
}

func (o *GetExtensionRepoVersionExtensionDocsForbidden) String() string {
	return fmt.Sprintf("[GET /extension-repository/{bucketName}/{groupId}/{artifactId}/{version}/extensions/{name}/docs][%d] getExtensionRepoVersionExtensionDocsForbidden ", 403)
}

func (o *GetExtensionRepoVersionExtensionDocsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExtensionRepoVersionExtensionDocsNotFound creates a GetExtensionRepoVersionExtensionDocsNotFound with default headers values
func NewGetExtensionRepoVersionExtensionDocsNotFound() *GetExtensionRepoVersionExtensionDocsNotFound {
	return &GetExtensionRepoVersionExtensionDocsNotFound{}
}

/*
GetExtensionRepoVersionExtensionDocsNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetExtensionRepoVersionExtensionDocsNotFound struct {
}

// IsSuccess returns true when this get extension repo version extension docs not found response has a 2xx status code
func (o *GetExtensionRepoVersionExtensionDocsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get extension repo version extension docs not found response has a 3xx status code
func (o *GetExtensionRepoVersionExtensionDocsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get extension repo version extension docs not found response has a 4xx status code
func (o *GetExtensionRepoVersionExtensionDocsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get extension repo version extension docs not found response has a 5xx status code
func (o *GetExtensionRepoVersionExtensionDocsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get extension repo version extension docs not found response a status code equal to that given
func (o *GetExtensionRepoVersionExtensionDocsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get extension repo version extension docs not found response
func (o *GetExtensionRepoVersionExtensionDocsNotFound) Code() int {
	return 404
}

func (o *GetExtensionRepoVersionExtensionDocsNotFound) Error() string {
	return fmt.Sprintf("[GET /extension-repository/{bucketName}/{groupId}/{artifactId}/{version}/extensions/{name}/docs][%d] getExtensionRepoVersionExtensionDocsNotFound ", 404)
}

func (o *GetExtensionRepoVersionExtensionDocsNotFound) String() string {
	return fmt.Sprintf("[GET /extension-repository/{bucketName}/{groupId}/{artifactId}/{version}/extensions/{name}/docs][%d] getExtensionRepoVersionExtensionDocsNotFound ", 404)
}

func (o *GetExtensionRepoVersionExtensionDocsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExtensionRepoVersionExtensionDocsConflict creates a GetExtensionRepoVersionExtensionDocsConflict with default headers values
func NewGetExtensionRepoVersionExtensionDocsConflict() *GetExtensionRepoVersionExtensionDocsConflict {
	return &GetExtensionRepoVersionExtensionDocsConflict{}
}

/*
GetExtensionRepoVersionExtensionDocsConflict describes a response with status code 409, with default header values.

NiFi Registry was unable to complete the request because it assumes a server state that is not valid.
*/
type GetExtensionRepoVersionExtensionDocsConflict struct {
}

// IsSuccess returns true when this get extension repo version extension docs conflict response has a 2xx status code
func (o *GetExtensionRepoVersionExtensionDocsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get extension repo version extension docs conflict response has a 3xx status code
func (o *GetExtensionRepoVersionExtensionDocsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get extension repo version extension docs conflict response has a 4xx status code
func (o *GetExtensionRepoVersionExtensionDocsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get extension repo version extension docs conflict response has a 5xx status code
func (o *GetExtensionRepoVersionExtensionDocsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get extension repo version extension docs conflict response a status code equal to that given
func (o *GetExtensionRepoVersionExtensionDocsConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the get extension repo version extension docs conflict response
func (o *GetExtensionRepoVersionExtensionDocsConflict) Code() int {
	return 409
}

func (o *GetExtensionRepoVersionExtensionDocsConflict) Error() string {
	return fmt.Sprintf("[GET /extension-repository/{bucketName}/{groupId}/{artifactId}/{version}/extensions/{name}/docs][%d] getExtensionRepoVersionExtensionDocsConflict ", 409)
}

func (o *GetExtensionRepoVersionExtensionDocsConflict) String() string {
	return fmt.Sprintf("[GET /extension-repository/{bucketName}/{groupId}/{artifactId}/{version}/extensions/{name}/docs][%d] getExtensionRepoVersionExtensionDocsConflict ", 409)
}

func (o *GetExtensionRepoVersionExtensionDocsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
