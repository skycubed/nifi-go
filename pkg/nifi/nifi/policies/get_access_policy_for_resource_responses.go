// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// GetAccessPolicyForResourceReader is a Reader for the GetAccessPolicyForResource structure.
type GetAccessPolicyForResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccessPolicyForResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccessPolicyForResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAccessPolicyForResourceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAccessPolicyForResourceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAccessPolicyForResourceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAccessPolicyForResourceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetAccessPolicyForResourceConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /policies/{action}/{resource}] getAccessPolicyForResource", response, response.Code())
	}
}

// NewGetAccessPolicyForResourceOK creates a GetAccessPolicyForResourceOK with default headers values
func NewGetAccessPolicyForResourceOK() *GetAccessPolicyForResourceOK {
	return &GetAccessPolicyForResourceOK{}
}

/*
GetAccessPolicyForResourceOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAccessPolicyForResourceOK struct {
	Payload *models.AccessPolicyEntity
}

// IsSuccess returns true when this get access policy for resource o k response has a 2xx status code
func (o *GetAccessPolicyForResourceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get access policy for resource o k response has a 3xx status code
func (o *GetAccessPolicyForResourceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get access policy for resource o k response has a 4xx status code
func (o *GetAccessPolicyForResourceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get access policy for resource o k response has a 5xx status code
func (o *GetAccessPolicyForResourceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get access policy for resource o k response a status code equal to that given
func (o *GetAccessPolicyForResourceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get access policy for resource o k response
func (o *GetAccessPolicyForResourceOK) Code() int {
	return 200
}

func (o *GetAccessPolicyForResourceOK) Error() string {
	return fmt.Sprintf("[GET /policies/{action}/{resource}][%d] getAccessPolicyForResourceOK  %+v", 200, o.Payload)
}

func (o *GetAccessPolicyForResourceOK) String() string {
	return fmt.Sprintf("[GET /policies/{action}/{resource}][%d] getAccessPolicyForResourceOK  %+v", 200, o.Payload)
}

func (o *GetAccessPolicyForResourceOK) GetPayload() *models.AccessPolicyEntity {
	return o.Payload
}

func (o *GetAccessPolicyForResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccessPolicyEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccessPolicyForResourceBadRequest creates a GetAccessPolicyForResourceBadRequest with default headers values
func NewGetAccessPolicyForResourceBadRequest() *GetAccessPolicyForResourceBadRequest {
	return &GetAccessPolicyForResourceBadRequest{}
}

/*
GetAccessPolicyForResourceBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetAccessPolicyForResourceBadRequest struct {
}

// IsSuccess returns true when this get access policy for resource bad request response has a 2xx status code
func (o *GetAccessPolicyForResourceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get access policy for resource bad request response has a 3xx status code
func (o *GetAccessPolicyForResourceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get access policy for resource bad request response has a 4xx status code
func (o *GetAccessPolicyForResourceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get access policy for resource bad request response has a 5xx status code
func (o *GetAccessPolicyForResourceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get access policy for resource bad request response a status code equal to that given
func (o *GetAccessPolicyForResourceBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get access policy for resource bad request response
func (o *GetAccessPolicyForResourceBadRequest) Code() int {
	return 400
}

func (o *GetAccessPolicyForResourceBadRequest) Error() string {
	return fmt.Sprintf("[GET /policies/{action}/{resource}][%d] getAccessPolicyForResourceBadRequest ", 400)
}

func (o *GetAccessPolicyForResourceBadRequest) String() string {
	return fmt.Sprintf("[GET /policies/{action}/{resource}][%d] getAccessPolicyForResourceBadRequest ", 400)
}

func (o *GetAccessPolicyForResourceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAccessPolicyForResourceUnauthorized creates a GetAccessPolicyForResourceUnauthorized with default headers values
func NewGetAccessPolicyForResourceUnauthorized() *GetAccessPolicyForResourceUnauthorized {
	return &GetAccessPolicyForResourceUnauthorized{}
}

/*
GetAccessPolicyForResourceUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetAccessPolicyForResourceUnauthorized struct {
}

// IsSuccess returns true when this get access policy for resource unauthorized response has a 2xx status code
func (o *GetAccessPolicyForResourceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get access policy for resource unauthorized response has a 3xx status code
func (o *GetAccessPolicyForResourceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get access policy for resource unauthorized response has a 4xx status code
func (o *GetAccessPolicyForResourceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get access policy for resource unauthorized response has a 5xx status code
func (o *GetAccessPolicyForResourceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get access policy for resource unauthorized response a status code equal to that given
func (o *GetAccessPolicyForResourceUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get access policy for resource unauthorized response
func (o *GetAccessPolicyForResourceUnauthorized) Code() int {
	return 401
}

func (o *GetAccessPolicyForResourceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /policies/{action}/{resource}][%d] getAccessPolicyForResourceUnauthorized ", 401)
}

func (o *GetAccessPolicyForResourceUnauthorized) String() string {
	return fmt.Sprintf("[GET /policies/{action}/{resource}][%d] getAccessPolicyForResourceUnauthorized ", 401)
}

func (o *GetAccessPolicyForResourceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAccessPolicyForResourceForbidden creates a GetAccessPolicyForResourceForbidden with default headers values
func NewGetAccessPolicyForResourceForbidden() *GetAccessPolicyForResourceForbidden {
	return &GetAccessPolicyForResourceForbidden{}
}

/*
GetAccessPolicyForResourceForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetAccessPolicyForResourceForbidden struct {
}

// IsSuccess returns true when this get access policy for resource forbidden response has a 2xx status code
func (o *GetAccessPolicyForResourceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get access policy for resource forbidden response has a 3xx status code
func (o *GetAccessPolicyForResourceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get access policy for resource forbidden response has a 4xx status code
func (o *GetAccessPolicyForResourceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get access policy for resource forbidden response has a 5xx status code
func (o *GetAccessPolicyForResourceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get access policy for resource forbidden response a status code equal to that given
func (o *GetAccessPolicyForResourceForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get access policy for resource forbidden response
func (o *GetAccessPolicyForResourceForbidden) Code() int {
	return 403
}

func (o *GetAccessPolicyForResourceForbidden) Error() string {
	return fmt.Sprintf("[GET /policies/{action}/{resource}][%d] getAccessPolicyForResourceForbidden ", 403)
}

func (o *GetAccessPolicyForResourceForbidden) String() string {
	return fmt.Sprintf("[GET /policies/{action}/{resource}][%d] getAccessPolicyForResourceForbidden ", 403)
}

func (o *GetAccessPolicyForResourceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAccessPolicyForResourceNotFound creates a GetAccessPolicyForResourceNotFound with default headers values
func NewGetAccessPolicyForResourceNotFound() *GetAccessPolicyForResourceNotFound {
	return &GetAccessPolicyForResourceNotFound{}
}

/*
GetAccessPolicyForResourceNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetAccessPolicyForResourceNotFound struct {
}

// IsSuccess returns true when this get access policy for resource not found response has a 2xx status code
func (o *GetAccessPolicyForResourceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get access policy for resource not found response has a 3xx status code
func (o *GetAccessPolicyForResourceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get access policy for resource not found response has a 4xx status code
func (o *GetAccessPolicyForResourceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get access policy for resource not found response has a 5xx status code
func (o *GetAccessPolicyForResourceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get access policy for resource not found response a status code equal to that given
func (o *GetAccessPolicyForResourceNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get access policy for resource not found response
func (o *GetAccessPolicyForResourceNotFound) Code() int {
	return 404
}

func (o *GetAccessPolicyForResourceNotFound) Error() string {
	return fmt.Sprintf("[GET /policies/{action}/{resource}][%d] getAccessPolicyForResourceNotFound ", 404)
}

func (o *GetAccessPolicyForResourceNotFound) String() string {
	return fmt.Sprintf("[GET /policies/{action}/{resource}][%d] getAccessPolicyForResourceNotFound ", 404)
}

func (o *GetAccessPolicyForResourceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAccessPolicyForResourceConflict creates a GetAccessPolicyForResourceConflict with default headers values
func NewGetAccessPolicyForResourceConflict() *GetAccessPolicyForResourceConflict {
	return &GetAccessPolicyForResourceConflict{}
}

/*
GetAccessPolicyForResourceConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetAccessPolicyForResourceConflict struct {
}

// IsSuccess returns true when this get access policy for resource conflict response has a 2xx status code
func (o *GetAccessPolicyForResourceConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get access policy for resource conflict response has a 3xx status code
func (o *GetAccessPolicyForResourceConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get access policy for resource conflict response has a 4xx status code
func (o *GetAccessPolicyForResourceConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get access policy for resource conflict response has a 5xx status code
func (o *GetAccessPolicyForResourceConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get access policy for resource conflict response a status code equal to that given
func (o *GetAccessPolicyForResourceConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the get access policy for resource conflict response
func (o *GetAccessPolicyForResourceConflict) Code() int {
	return 409
}

func (o *GetAccessPolicyForResourceConflict) Error() string {
	return fmt.Sprintf("[GET /policies/{action}/{resource}][%d] getAccessPolicyForResourceConflict ", 409)
}

func (o *GetAccessPolicyForResourceConflict) String() string {
	return fmt.Sprintf("[GET /policies/{action}/{resource}][%d] getAccessPolicyForResourceConflict ", 409)
}

func (o *GetAccessPolicyForResourceConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
