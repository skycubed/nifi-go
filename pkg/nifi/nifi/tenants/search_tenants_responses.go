// Code generated by go-swagger; DO NOT EDIT.

package tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// SearchTenantsReader is a Reader for the SearchTenants structure.
type SearchTenantsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchTenantsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchTenantsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchTenantsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSearchTenantsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSearchTenantsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchTenantsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewSearchTenantsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchTenantsOK creates a SearchTenantsOK with default headers values
func NewSearchTenantsOK() *SearchTenantsOK {
	return &SearchTenantsOK{}
}

/* SearchTenantsOK describes a response with status code 200, with default header values.

successful operation
*/
type SearchTenantsOK struct {
	Payload *models.TenantsEntity
}

func (o *SearchTenantsOK) Error() string {
	return fmt.Sprintf("[GET /tenants/search-results][%d] searchTenantsOK  %+v", 200, o.Payload)
}
func (o *SearchTenantsOK) GetPayload() *models.TenantsEntity {
	return o.Payload
}

func (o *SearchTenantsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TenantsEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchTenantsBadRequest creates a SearchTenantsBadRequest with default headers values
func NewSearchTenantsBadRequest() *SearchTenantsBadRequest {
	return &SearchTenantsBadRequest{}
}

/* SearchTenantsBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type SearchTenantsBadRequest struct {
}

func (o *SearchTenantsBadRequest) Error() string {
	return fmt.Sprintf("[GET /tenants/search-results][%d] searchTenantsBadRequest ", 400)
}

func (o *SearchTenantsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSearchTenantsUnauthorized creates a SearchTenantsUnauthorized with default headers values
func NewSearchTenantsUnauthorized() *SearchTenantsUnauthorized {
	return &SearchTenantsUnauthorized{}
}

/* SearchTenantsUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type SearchTenantsUnauthorized struct {
}

func (o *SearchTenantsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /tenants/search-results][%d] searchTenantsUnauthorized ", 401)
}

func (o *SearchTenantsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSearchTenantsForbidden creates a SearchTenantsForbidden with default headers values
func NewSearchTenantsForbidden() *SearchTenantsForbidden {
	return &SearchTenantsForbidden{}
}

/* SearchTenantsForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type SearchTenantsForbidden struct {
}

func (o *SearchTenantsForbidden) Error() string {
	return fmt.Sprintf("[GET /tenants/search-results][%d] searchTenantsForbidden ", 403)
}

func (o *SearchTenantsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSearchTenantsNotFound creates a SearchTenantsNotFound with default headers values
func NewSearchTenantsNotFound() *SearchTenantsNotFound {
	return &SearchTenantsNotFound{}
}

/* SearchTenantsNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type SearchTenantsNotFound struct {
}

func (o *SearchTenantsNotFound) Error() string {
	return fmt.Sprintf("[GET /tenants/search-results][%d] searchTenantsNotFound ", 404)
}

func (o *SearchTenantsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSearchTenantsConflict creates a SearchTenantsConflict with default headers values
func NewSearchTenantsConflict() *SearchTenantsConflict {
	return &SearchTenantsConflict{}
}

/* SearchTenantsConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type SearchTenantsConflict struct {
}

func (o *SearchTenantsConflict) Error() string {
	return fmt.Sprintf("[GET /tenants/search-results][%d] searchTenantsConflict ", 409)
}

func (o *SearchTenantsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
