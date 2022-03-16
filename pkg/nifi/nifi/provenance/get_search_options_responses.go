// Code generated by go-swagger; DO NOT EDIT.

package provenance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// GetSearchOptionsReader is a Reader for the GetSearchOptions structure.
type GetSearchOptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSearchOptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSearchOptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSearchOptionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetSearchOptionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSearchOptionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetSearchOptionsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSearchOptionsOK creates a GetSearchOptionsOK with default headers values
func NewGetSearchOptionsOK() *GetSearchOptionsOK {
	return &GetSearchOptionsOK{}
}

/* GetSearchOptionsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetSearchOptionsOK struct {
	Payload *models.ProvenanceOptionsEntity
}

func (o *GetSearchOptionsOK) Error() string {
	return fmt.Sprintf("[GET /provenance/search-options][%d] getSearchOptionsOK  %+v", 200, o.Payload)
}
func (o *GetSearchOptionsOK) GetPayload() *models.ProvenanceOptionsEntity {
	return o.Payload
}

func (o *GetSearchOptionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProvenanceOptionsEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSearchOptionsBadRequest creates a GetSearchOptionsBadRequest with default headers values
func NewGetSearchOptionsBadRequest() *GetSearchOptionsBadRequest {
	return &GetSearchOptionsBadRequest{}
}

/* GetSearchOptionsBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetSearchOptionsBadRequest struct {
}

func (o *GetSearchOptionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /provenance/search-options][%d] getSearchOptionsBadRequest ", 400)
}

func (o *GetSearchOptionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSearchOptionsUnauthorized creates a GetSearchOptionsUnauthorized with default headers values
func NewGetSearchOptionsUnauthorized() *GetSearchOptionsUnauthorized {
	return &GetSearchOptionsUnauthorized{}
}

/* GetSearchOptionsUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetSearchOptionsUnauthorized struct {
}

func (o *GetSearchOptionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /provenance/search-options][%d] getSearchOptionsUnauthorized ", 401)
}

func (o *GetSearchOptionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSearchOptionsForbidden creates a GetSearchOptionsForbidden with default headers values
func NewGetSearchOptionsForbidden() *GetSearchOptionsForbidden {
	return &GetSearchOptionsForbidden{}
}

/* GetSearchOptionsForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetSearchOptionsForbidden struct {
}

func (o *GetSearchOptionsForbidden) Error() string {
	return fmt.Sprintf("[GET /provenance/search-options][%d] getSearchOptionsForbidden ", 403)
}

func (o *GetSearchOptionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSearchOptionsConflict creates a GetSearchOptionsConflict with default headers values
func NewGetSearchOptionsConflict() *GetSearchOptionsConflict {
	return &GetSearchOptionsConflict{}
}

/* GetSearchOptionsConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetSearchOptionsConflict struct {
}

func (o *GetSearchOptionsConflict) Error() string {
	return fmt.Sprintf("[GET /provenance/search-options][%d] getSearchOptionsConflict ", 409)
}

func (o *GetSearchOptionsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
