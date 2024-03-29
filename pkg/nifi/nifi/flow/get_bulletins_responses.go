// Code generated by go-swagger; DO NOT EDIT.

package flow

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// GetBulletinsReader is a Reader for the GetBulletins structure.
type GetBulletinsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBulletinsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBulletinsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetBulletinsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetBulletinsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetBulletinsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetBulletinsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetBulletinsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /flow/controller/bulletins] getBulletins", response, response.Code())
	}
}

// NewGetBulletinsOK creates a GetBulletinsOK with default headers values
func NewGetBulletinsOK() *GetBulletinsOK {
	return &GetBulletinsOK{}
}

/*
GetBulletinsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetBulletinsOK struct {
	Payload *models.ControllerBulletinsEntity
}

// IsSuccess returns true when this get bulletins o k response has a 2xx status code
func (o *GetBulletinsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get bulletins o k response has a 3xx status code
func (o *GetBulletinsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bulletins o k response has a 4xx status code
func (o *GetBulletinsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get bulletins o k response has a 5xx status code
func (o *GetBulletinsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get bulletins o k response a status code equal to that given
func (o *GetBulletinsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get bulletins o k response
func (o *GetBulletinsOK) Code() int {
	return 200
}

func (o *GetBulletinsOK) Error() string {
	return fmt.Sprintf("[GET /flow/controller/bulletins][%d] getBulletinsOK  %+v", 200, o.Payload)
}

func (o *GetBulletinsOK) String() string {
	return fmt.Sprintf("[GET /flow/controller/bulletins][%d] getBulletinsOK  %+v", 200, o.Payload)
}

func (o *GetBulletinsOK) GetPayload() *models.ControllerBulletinsEntity {
	return o.Payload
}

func (o *GetBulletinsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ControllerBulletinsEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBulletinsBadRequest creates a GetBulletinsBadRequest with default headers values
func NewGetBulletinsBadRequest() *GetBulletinsBadRequest {
	return &GetBulletinsBadRequest{}
}

/*
GetBulletinsBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetBulletinsBadRequest struct {
}

// IsSuccess returns true when this get bulletins bad request response has a 2xx status code
func (o *GetBulletinsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get bulletins bad request response has a 3xx status code
func (o *GetBulletinsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bulletins bad request response has a 4xx status code
func (o *GetBulletinsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get bulletins bad request response has a 5xx status code
func (o *GetBulletinsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get bulletins bad request response a status code equal to that given
func (o *GetBulletinsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get bulletins bad request response
func (o *GetBulletinsBadRequest) Code() int {
	return 400
}

func (o *GetBulletinsBadRequest) Error() string {
	return fmt.Sprintf("[GET /flow/controller/bulletins][%d] getBulletinsBadRequest ", 400)
}

func (o *GetBulletinsBadRequest) String() string {
	return fmt.Sprintf("[GET /flow/controller/bulletins][%d] getBulletinsBadRequest ", 400)
}

func (o *GetBulletinsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBulletinsUnauthorized creates a GetBulletinsUnauthorized with default headers values
func NewGetBulletinsUnauthorized() *GetBulletinsUnauthorized {
	return &GetBulletinsUnauthorized{}
}

/*
GetBulletinsUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetBulletinsUnauthorized struct {
}

// IsSuccess returns true when this get bulletins unauthorized response has a 2xx status code
func (o *GetBulletinsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get bulletins unauthorized response has a 3xx status code
func (o *GetBulletinsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bulletins unauthorized response has a 4xx status code
func (o *GetBulletinsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get bulletins unauthorized response has a 5xx status code
func (o *GetBulletinsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get bulletins unauthorized response a status code equal to that given
func (o *GetBulletinsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get bulletins unauthorized response
func (o *GetBulletinsUnauthorized) Code() int {
	return 401
}

func (o *GetBulletinsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /flow/controller/bulletins][%d] getBulletinsUnauthorized ", 401)
}

func (o *GetBulletinsUnauthorized) String() string {
	return fmt.Sprintf("[GET /flow/controller/bulletins][%d] getBulletinsUnauthorized ", 401)
}

func (o *GetBulletinsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBulletinsForbidden creates a GetBulletinsForbidden with default headers values
func NewGetBulletinsForbidden() *GetBulletinsForbidden {
	return &GetBulletinsForbidden{}
}

/*
GetBulletinsForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetBulletinsForbidden struct {
}

// IsSuccess returns true when this get bulletins forbidden response has a 2xx status code
func (o *GetBulletinsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get bulletins forbidden response has a 3xx status code
func (o *GetBulletinsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bulletins forbidden response has a 4xx status code
func (o *GetBulletinsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get bulletins forbidden response has a 5xx status code
func (o *GetBulletinsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get bulletins forbidden response a status code equal to that given
func (o *GetBulletinsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get bulletins forbidden response
func (o *GetBulletinsForbidden) Code() int {
	return 403
}

func (o *GetBulletinsForbidden) Error() string {
	return fmt.Sprintf("[GET /flow/controller/bulletins][%d] getBulletinsForbidden ", 403)
}

func (o *GetBulletinsForbidden) String() string {
	return fmt.Sprintf("[GET /flow/controller/bulletins][%d] getBulletinsForbidden ", 403)
}

func (o *GetBulletinsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBulletinsNotFound creates a GetBulletinsNotFound with default headers values
func NewGetBulletinsNotFound() *GetBulletinsNotFound {
	return &GetBulletinsNotFound{}
}

/*
GetBulletinsNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetBulletinsNotFound struct {
}

// IsSuccess returns true when this get bulletins not found response has a 2xx status code
func (o *GetBulletinsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get bulletins not found response has a 3xx status code
func (o *GetBulletinsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bulletins not found response has a 4xx status code
func (o *GetBulletinsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get bulletins not found response has a 5xx status code
func (o *GetBulletinsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get bulletins not found response a status code equal to that given
func (o *GetBulletinsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get bulletins not found response
func (o *GetBulletinsNotFound) Code() int {
	return 404
}

func (o *GetBulletinsNotFound) Error() string {
	return fmt.Sprintf("[GET /flow/controller/bulletins][%d] getBulletinsNotFound ", 404)
}

func (o *GetBulletinsNotFound) String() string {
	return fmt.Sprintf("[GET /flow/controller/bulletins][%d] getBulletinsNotFound ", 404)
}

func (o *GetBulletinsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBulletinsConflict creates a GetBulletinsConflict with default headers values
func NewGetBulletinsConflict() *GetBulletinsConflict {
	return &GetBulletinsConflict{}
}

/*
GetBulletinsConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetBulletinsConflict struct {
}

// IsSuccess returns true when this get bulletins conflict response has a 2xx status code
func (o *GetBulletinsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get bulletins conflict response has a 3xx status code
func (o *GetBulletinsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bulletins conflict response has a 4xx status code
func (o *GetBulletinsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get bulletins conflict response has a 5xx status code
func (o *GetBulletinsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get bulletins conflict response a status code equal to that given
func (o *GetBulletinsConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the get bulletins conflict response
func (o *GetBulletinsConflict) Code() int {
	return 409
}

func (o *GetBulletinsConflict) Error() string {
	return fmt.Sprintf("[GET /flow/controller/bulletins][%d] getBulletinsConflict ", 409)
}

func (o *GetBulletinsConflict) String() string {
	return fmt.Sprintf("[GET /flow/controller/bulletins][%d] getBulletinsConflict ", 409)
}

func (o *GetBulletinsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
