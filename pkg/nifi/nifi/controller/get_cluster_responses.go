// Code generated by go-swagger; DO NOT EDIT.

package controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// GetClusterReader is a Reader for the GetCluster structure.
type GetClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetClusterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetClusterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetClusterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetClusterConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetClusterOK creates a GetClusterOK with default headers values
func NewGetClusterOK() *GetClusterOK {
	return &GetClusterOK{}
}

/*
GetClusterOK describes a response with status code 200, with default header values.

successful operation
*/
type GetClusterOK struct {
	Payload *models.ClusterEntity
}

// IsSuccess returns true when this get cluster o k response has a 2xx status code
func (o *GetClusterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cluster o k response has a 3xx status code
func (o *GetClusterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster o k response has a 4xx status code
func (o *GetClusterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cluster o k response has a 5xx status code
func (o *GetClusterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster o k response a status code equal to that given
func (o *GetClusterOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetClusterOK) Error() string {
	return fmt.Sprintf("[GET /controller/cluster][%d] getClusterOK  %+v", 200, o.Payload)
}

func (o *GetClusterOK) String() string {
	return fmt.Sprintf("[GET /controller/cluster][%d] getClusterOK  %+v", 200, o.Payload)
}

func (o *GetClusterOK) GetPayload() *models.ClusterEntity {
	return o.Payload
}

func (o *GetClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterBadRequest creates a GetClusterBadRequest with default headers values
func NewGetClusterBadRequest() *GetClusterBadRequest {
	return &GetClusterBadRequest{}
}

/*
GetClusterBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetClusterBadRequest struct {
}

// IsSuccess returns true when this get cluster bad request response has a 2xx status code
func (o *GetClusterBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster bad request response has a 3xx status code
func (o *GetClusterBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster bad request response has a 4xx status code
func (o *GetClusterBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cluster bad request response has a 5xx status code
func (o *GetClusterBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster bad request response a status code equal to that given
func (o *GetClusterBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetClusterBadRequest) Error() string {
	return fmt.Sprintf("[GET /controller/cluster][%d] getClusterBadRequest ", 400)
}

func (o *GetClusterBadRequest) String() string {
	return fmt.Sprintf("[GET /controller/cluster][%d] getClusterBadRequest ", 400)
}

func (o *GetClusterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClusterUnauthorized creates a GetClusterUnauthorized with default headers values
func NewGetClusterUnauthorized() *GetClusterUnauthorized {
	return &GetClusterUnauthorized{}
}

/*
GetClusterUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetClusterUnauthorized struct {
}

// IsSuccess returns true when this get cluster unauthorized response has a 2xx status code
func (o *GetClusterUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster unauthorized response has a 3xx status code
func (o *GetClusterUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster unauthorized response has a 4xx status code
func (o *GetClusterUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cluster unauthorized response has a 5xx status code
func (o *GetClusterUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster unauthorized response a status code equal to that given
func (o *GetClusterUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetClusterUnauthorized) Error() string {
	return fmt.Sprintf("[GET /controller/cluster][%d] getClusterUnauthorized ", 401)
}

func (o *GetClusterUnauthorized) String() string {
	return fmt.Sprintf("[GET /controller/cluster][%d] getClusterUnauthorized ", 401)
}

func (o *GetClusterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClusterForbidden creates a GetClusterForbidden with default headers values
func NewGetClusterForbidden() *GetClusterForbidden {
	return &GetClusterForbidden{}
}

/*
GetClusterForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetClusterForbidden struct {
}

// IsSuccess returns true when this get cluster forbidden response has a 2xx status code
func (o *GetClusterForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster forbidden response has a 3xx status code
func (o *GetClusterForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster forbidden response has a 4xx status code
func (o *GetClusterForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cluster forbidden response has a 5xx status code
func (o *GetClusterForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster forbidden response a status code equal to that given
func (o *GetClusterForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetClusterForbidden) Error() string {
	return fmt.Sprintf("[GET /controller/cluster][%d] getClusterForbidden ", 403)
}

func (o *GetClusterForbidden) String() string {
	return fmt.Sprintf("[GET /controller/cluster][%d] getClusterForbidden ", 403)
}

func (o *GetClusterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClusterConflict creates a GetClusterConflict with default headers values
func NewGetClusterConflict() *GetClusterConflict {
	return &GetClusterConflict{}
}

/*
GetClusterConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetClusterConflict struct {
}

// IsSuccess returns true when this get cluster conflict response has a 2xx status code
func (o *GetClusterConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster conflict response has a 3xx status code
func (o *GetClusterConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster conflict response has a 4xx status code
func (o *GetClusterConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cluster conflict response has a 5xx status code
func (o *GetClusterConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster conflict response a status code equal to that given
func (o *GetClusterConflict) IsCode(code int) bool {
	return code == 409
}

func (o *GetClusterConflict) Error() string {
	return fmt.Sprintf("[GET /controller/cluster][%d] getClusterConflict ", 409)
}

func (o *GetClusterConflict) String() string {
	return fmt.Sprintf("[GET /controller/cluster][%d] getClusterConflict ", 409)
}

func (o *GetClusterConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
