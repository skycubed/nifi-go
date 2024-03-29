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

// GetLineageReader is a Reader for the GetLineage structure.
type GetLineageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLineageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLineageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLineageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLineageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLineageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLineageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetLineageConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /provenance/lineage/{id}] getLineage", response, response.Code())
	}
}

// NewGetLineageOK creates a GetLineageOK with default headers values
func NewGetLineageOK() *GetLineageOK {
	return &GetLineageOK{}
}

/*
GetLineageOK describes a response with status code 200, with default header values.

successful operation
*/
type GetLineageOK struct {
	Payload *models.LineageEntity
}

// IsSuccess returns true when this get lineage o k response has a 2xx status code
func (o *GetLineageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get lineage o k response has a 3xx status code
func (o *GetLineageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get lineage o k response has a 4xx status code
func (o *GetLineageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get lineage o k response has a 5xx status code
func (o *GetLineageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get lineage o k response a status code equal to that given
func (o *GetLineageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get lineage o k response
func (o *GetLineageOK) Code() int {
	return 200
}

func (o *GetLineageOK) Error() string {
	return fmt.Sprintf("[GET /provenance/lineage/{id}][%d] getLineageOK  %+v", 200, o.Payload)
}

func (o *GetLineageOK) String() string {
	return fmt.Sprintf("[GET /provenance/lineage/{id}][%d] getLineageOK  %+v", 200, o.Payload)
}

func (o *GetLineageOK) GetPayload() *models.LineageEntity {
	return o.Payload
}

func (o *GetLineageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LineageEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLineageBadRequest creates a GetLineageBadRequest with default headers values
func NewGetLineageBadRequest() *GetLineageBadRequest {
	return &GetLineageBadRequest{}
}

/*
GetLineageBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetLineageBadRequest struct {
}

// IsSuccess returns true when this get lineage bad request response has a 2xx status code
func (o *GetLineageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get lineage bad request response has a 3xx status code
func (o *GetLineageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get lineage bad request response has a 4xx status code
func (o *GetLineageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get lineage bad request response has a 5xx status code
func (o *GetLineageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get lineage bad request response a status code equal to that given
func (o *GetLineageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get lineage bad request response
func (o *GetLineageBadRequest) Code() int {
	return 400
}

func (o *GetLineageBadRequest) Error() string {
	return fmt.Sprintf("[GET /provenance/lineage/{id}][%d] getLineageBadRequest ", 400)
}

func (o *GetLineageBadRequest) String() string {
	return fmt.Sprintf("[GET /provenance/lineage/{id}][%d] getLineageBadRequest ", 400)
}

func (o *GetLineageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLineageUnauthorized creates a GetLineageUnauthorized with default headers values
func NewGetLineageUnauthorized() *GetLineageUnauthorized {
	return &GetLineageUnauthorized{}
}

/*
GetLineageUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetLineageUnauthorized struct {
}

// IsSuccess returns true when this get lineage unauthorized response has a 2xx status code
func (o *GetLineageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get lineage unauthorized response has a 3xx status code
func (o *GetLineageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get lineage unauthorized response has a 4xx status code
func (o *GetLineageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get lineage unauthorized response has a 5xx status code
func (o *GetLineageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get lineage unauthorized response a status code equal to that given
func (o *GetLineageUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get lineage unauthorized response
func (o *GetLineageUnauthorized) Code() int {
	return 401
}

func (o *GetLineageUnauthorized) Error() string {
	return fmt.Sprintf("[GET /provenance/lineage/{id}][%d] getLineageUnauthorized ", 401)
}

func (o *GetLineageUnauthorized) String() string {
	return fmt.Sprintf("[GET /provenance/lineage/{id}][%d] getLineageUnauthorized ", 401)
}

func (o *GetLineageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLineageForbidden creates a GetLineageForbidden with default headers values
func NewGetLineageForbidden() *GetLineageForbidden {
	return &GetLineageForbidden{}
}

/*
GetLineageForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetLineageForbidden struct {
}

// IsSuccess returns true when this get lineage forbidden response has a 2xx status code
func (o *GetLineageForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get lineage forbidden response has a 3xx status code
func (o *GetLineageForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get lineage forbidden response has a 4xx status code
func (o *GetLineageForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get lineage forbidden response has a 5xx status code
func (o *GetLineageForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get lineage forbidden response a status code equal to that given
func (o *GetLineageForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get lineage forbidden response
func (o *GetLineageForbidden) Code() int {
	return 403
}

func (o *GetLineageForbidden) Error() string {
	return fmt.Sprintf("[GET /provenance/lineage/{id}][%d] getLineageForbidden ", 403)
}

func (o *GetLineageForbidden) String() string {
	return fmt.Sprintf("[GET /provenance/lineage/{id}][%d] getLineageForbidden ", 403)
}

func (o *GetLineageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLineageNotFound creates a GetLineageNotFound with default headers values
func NewGetLineageNotFound() *GetLineageNotFound {
	return &GetLineageNotFound{}
}

/*
GetLineageNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetLineageNotFound struct {
}

// IsSuccess returns true when this get lineage not found response has a 2xx status code
func (o *GetLineageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get lineage not found response has a 3xx status code
func (o *GetLineageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get lineage not found response has a 4xx status code
func (o *GetLineageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get lineage not found response has a 5xx status code
func (o *GetLineageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get lineage not found response a status code equal to that given
func (o *GetLineageNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get lineage not found response
func (o *GetLineageNotFound) Code() int {
	return 404
}

func (o *GetLineageNotFound) Error() string {
	return fmt.Sprintf("[GET /provenance/lineage/{id}][%d] getLineageNotFound ", 404)
}

func (o *GetLineageNotFound) String() string {
	return fmt.Sprintf("[GET /provenance/lineage/{id}][%d] getLineageNotFound ", 404)
}

func (o *GetLineageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLineageConflict creates a GetLineageConflict with default headers values
func NewGetLineageConflict() *GetLineageConflict {
	return &GetLineageConflict{}
}

/*
GetLineageConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetLineageConflict struct {
}

// IsSuccess returns true when this get lineage conflict response has a 2xx status code
func (o *GetLineageConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get lineage conflict response has a 3xx status code
func (o *GetLineageConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get lineage conflict response has a 4xx status code
func (o *GetLineageConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get lineage conflict response has a 5xx status code
func (o *GetLineageConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get lineage conflict response a status code equal to that given
func (o *GetLineageConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the get lineage conflict response
func (o *GetLineageConflict) Code() int {
	return 409
}

func (o *GetLineageConflict) Error() string {
	return fmt.Sprintf("[GET /provenance/lineage/{id}][%d] getLineageConflict ", 409)
}

func (o *GetLineageConflict) String() string {
	return fmt.Sprintf("[GET /provenance/lineage/{id}][%d] getLineageConflict ", 409)
}

func (o *GetLineageConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
