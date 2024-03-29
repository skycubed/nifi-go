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

// CreateExtensionBundleVersionReader is a Reader for the CreateExtensionBundleVersion structure.
type CreateExtensionBundleVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateExtensionBundleVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateExtensionBundleVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateExtensionBundleVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateExtensionBundleVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateExtensionBundleVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateExtensionBundleVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateExtensionBundleVersionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /buckets/{bucketId}/bundles/{bundleType}] createExtensionBundleVersion", response, response.Code())
	}
}

// NewCreateExtensionBundleVersionOK creates a CreateExtensionBundleVersionOK with default headers values
func NewCreateExtensionBundleVersionOK() *CreateExtensionBundleVersionOK {
	return &CreateExtensionBundleVersionOK{}
}

/*
CreateExtensionBundleVersionOK describes a response with status code 200, with default header values.

successful operation
*/
type CreateExtensionBundleVersionOK struct {
	Payload *models.BundleVersion
}

// IsSuccess returns true when this create extension bundle version o k response has a 2xx status code
func (o *CreateExtensionBundleVersionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create extension bundle version o k response has a 3xx status code
func (o *CreateExtensionBundleVersionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create extension bundle version o k response has a 4xx status code
func (o *CreateExtensionBundleVersionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create extension bundle version o k response has a 5xx status code
func (o *CreateExtensionBundleVersionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create extension bundle version o k response a status code equal to that given
func (o *CreateExtensionBundleVersionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create extension bundle version o k response
func (o *CreateExtensionBundleVersionOK) Code() int {
	return 200
}

func (o *CreateExtensionBundleVersionOK) Error() string {
	return fmt.Sprintf("[POST /buckets/{bucketId}/bundles/{bundleType}][%d] createExtensionBundleVersionOK  %+v", 200, o.Payload)
}

func (o *CreateExtensionBundleVersionOK) String() string {
	return fmt.Sprintf("[POST /buckets/{bucketId}/bundles/{bundleType}][%d] createExtensionBundleVersionOK  %+v", 200, o.Payload)
}

func (o *CreateExtensionBundleVersionOK) GetPayload() *models.BundleVersion {
	return o.Payload
}

func (o *CreateExtensionBundleVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BundleVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateExtensionBundleVersionBadRequest creates a CreateExtensionBundleVersionBadRequest with default headers values
func NewCreateExtensionBundleVersionBadRequest() *CreateExtensionBundleVersionBadRequest {
	return &CreateExtensionBundleVersionBadRequest{}
}

/*
CreateExtensionBundleVersionBadRequest describes a response with status code 400, with default header values.

NiFi Registry was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type CreateExtensionBundleVersionBadRequest struct {
}

// IsSuccess returns true when this create extension bundle version bad request response has a 2xx status code
func (o *CreateExtensionBundleVersionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create extension bundle version bad request response has a 3xx status code
func (o *CreateExtensionBundleVersionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create extension bundle version bad request response has a 4xx status code
func (o *CreateExtensionBundleVersionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create extension bundle version bad request response has a 5xx status code
func (o *CreateExtensionBundleVersionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create extension bundle version bad request response a status code equal to that given
func (o *CreateExtensionBundleVersionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create extension bundle version bad request response
func (o *CreateExtensionBundleVersionBadRequest) Code() int {
	return 400
}

func (o *CreateExtensionBundleVersionBadRequest) Error() string {
	return fmt.Sprintf("[POST /buckets/{bucketId}/bundles/{bundleType}][%d] createExtensionBundleVersionBadRequest ", 400)
}

func (o *CreateExtensionBundleVersionBadRequest) String() string {
	return fmt.Sprintf("[POST /buckets/{bucketId}/bundles/{bundleType}][%d] createExtensionBundleVersionBadRequest ", 400)
}

func (o *CreateExtensionBundleVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateExtensionBundleVersionUnauthorized creates a CreateExtensionBundleVersionUnauthorized with default headers values
func NewCreateExtensionBundleVersionUnauthorized() *CreateExtensionBundleVersionUnauthorized {
	return &CreateExtensionBundleVersionUnauthorized{}
}

/*
CreateExtensionBundleVersionUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type CreateExtensionBundleVersionUnauthorized struct {
}

// IsSuccess returns true when this create extension bundle version unauthorized response has a 2xx status code
func (o *CreateExtensionBundleVersionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create extension bundle version unauthorized response has a 3xx status code
func (o *CreateExtensionBundleVersionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create extension bundle version unauthorized response has a 4xx status code
func (o *CreateExtensionBundleVersionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create extension bundle version unauthorized response has a 5xx status code
func (o *CreateExtensionBundleVersionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create extension bundle version unauthorized response a status code equal to that given
func (o *CreateExtensionBundleVersionUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create extension bundle version unauthorized response
func (o *CreateExtensionBundleVersionUnauthorized) Code() int {
	return 401
}

func (o *CreateExtensionBundleVersionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /buckets/{bucketId}/bundles/{bundleType}][%d] createExtensionBundleVersionUnauthorized ", 401)
}

func (o *CreateExtensionBundleVersionUnauthorized) String() string {
	return fmt.Sprintf("[POST /buckets/{bucketId}/bundles/{bundleType}][%d] createExtensionBundleVersionUnauthorized ", 401)
}

func (o *CreateExtensionBundleVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateExtensionBundleVersionForbidden creates a CreateExtensionBundleVersionForbidden with default headers values
func NewCreateExtensionBundleVersionForbidden() *CreateExtensionBundleVersionForbidden {
	return &CreateExtensionBundleVersionForbidden{}
}

/*
CreateExtensionBundleVersionForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type CreateExtensionBundleVersionForbidden struct {
}

// IsSuccess returns true when this create extension bundle version forbidden response has a 2xx status code
func (o *CreateExtensionBundleVersionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create extension bundle version forbidden response has a 3xx status code
func (o *CreateExtensionBundleVersionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create extension bundle version forbidden response has a 4xx status code
func (o *CreateExtensionBundleVersionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create extension bundle version forbidden response has a 5xx status code
func (o *CreateExtensionBundleVersionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create extension bundle version forbidden response a status code equal to that given
func (o *CreateExtensionBundleVersionForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create extension bundle version forbidden response
func (o *CreateExtensionBundleVersionForbidden) Code() int {
	return 403
}

func (o *CreateExtensionBundleVersionForbidden) Error() string {
	return fmt.Sprintf("[POST /buckets/{bucketId}/bundles/{bundleType}][%d] createExtensionBundleVersionForbidden ", 403)
}

func (o *CreateExtensionBundleVersionForbidden) String() string {
	return fmt.Sprintf("[POST /buckets/{bucketId}/bundles/{bundleType}][%d] createExtensionBundleVersionForbidden ", 403)
}

func (o *CreateExtensionBundleVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateExtensionBundleVersionNotFound creates a CreateExtensionBundleVersionNotFound with default headers values
func NewCreateExtensionBundleVersionNotFound() *CreateExtensionBundleVersionNotFound {
	return &CreateExtensionBundleVersionNotFound{}
}

/*
CreateExtensionBundleVersionNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type CreateExtensionBundleVersionNotFound struct {
}

// IsSuccess returns true when this create extension bundle version not found response has a 2xx status code
func (o *CreateExtensionBundleVersionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create extension bundle version not found response has a 3xx status code
func (o *CreateExtensionBundleVersionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create extension bundle version not found response has a 4xx status code
func (o *CreateExtensionBundleVersionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create extension bundle version not found response has a 5xx status code
func (o *CreateExtensionBundleVersionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create extension bundle version not found response a status code equal to that given
func (o *CreateExtensionBundleVersionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create extension bundle version not found response
func (o *CreateExtensionBundleVersionNotFound) Code() int {
	return 404
}

func (o *CreateExtensionBundleVersionNotFound) Error() string {
	return fmt.Sprintf("[POST /buckets/{bucketId}/bundles/{bundleType}][%d] createExtensionBundleVersionNotFound ", 404)
}

func (o *CreateExtensionBundleVersionNotFound) String() string {
	return fmt.Sprintf("[POST /buckets/{bucketId}/bundles/{bundleType}][%d] createExtensionBundleVersionNotFound ", 404)
}

func (o *CreateExtensionBundleVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateExtensionBundleVersionConflict creates a CreateExtensionBundleVersionConflict with default headers values
func NewCreateExtensionBundleVersionConflict() *CreateExtensionBundleVersionConflict {
	return &CreateExtensionBundleVersionConflict{}
}

/*
CreateExtensionBundleVersionConflict describes a response with status code 409, with default header values.

NiFi Registry was unable to complete the request because it assumes a server state that is not valid.
*/
type CreateExtensionBundleVersionConflict struct {
}

// IsSuccess returns true when this create extension bundle version conflict response has a 2xx status code
func (o *CreateExtensionBundleVersionConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create extension bundle version conflict response has a 3xx status code
func (o *CreateExtensionBundleVersionConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create extension bundle version conflict response has a 4xx status code
func (o *CreateExtensionBundleVersionConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create extension bundle version conflict response has a 5xx status code
func (o *CreateExtensionBundleVersionConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create extension bundle version conflict response a status code equal to that given
func (o *CreateExtensionBundleVersionConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create extension bundle version conflict response
func (o *CreateExtensionBundleVersionConflict) Code() int {
	return 409
}

func (o *CreateExtensionBundleVersionConflict) Error() string {
	return fmt.Sprintf("[POST /buckets/{bucketId}/bundles/{bundleType}][%d] createExtensionBundleVersionConflict ", 409)
}

func (o *CreateExtensionBundleVersionConflict) String() string {
	return fmt.Sprintf("[POST /buckets/{bucketId}/bundles/{bundleType}][%d] createExtensionBundleVersionConflict ", 409)
}

func (o *CreateExtensionBundleVersionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
