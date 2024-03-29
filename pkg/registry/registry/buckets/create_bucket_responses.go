// Code generated by go-swagger; DO NOT EDIT.

package buckets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/registry/models"
)

// CreateBucketReader is a Reader for the CreateBucket structure.
type CreateBucketReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateBucketReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateBucketOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateBucketBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateBucketUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateBucketForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /buckets] createBucket", response, response.Code())
	}
}

// NewCreateBucketOK creates a CreateBucketOK with default headers values
func NewCreateBucketOK() *CreateBucketOK {
	return &CreateBucketOK{}
}

/*
CreateBucketOK describes a response with status code 200, with default header values.

successful operation
*/
type CreateBucketOK struct {
	Payload *models.Bucket
}

// IsSuccess returns true when this create bucket o k response has a 2xx status code
func (o *CreateBucketOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create bucket o k response has a 3xx status code
func (o *CreateBucketOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create bucket o k response has a 4xx status code
func (o *CreateBucketOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create bucket o k response has a 5xx status code
func (o *CreateBucketOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create bucket o k response a status code equal to that given
func (o *CreateBucketOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create bucket o k response
func (o *CreateBucketOK) Code() int {
	return 200
}

func (o *CreateBucketOK) Error() string {
	return fmt.Sprintf("[POST /buckets][%d] createBucketOK  %+v", 200, o.Payload)
}

func (o *CreateBucketOK) String() string {
	return fmt.Sprintf("[POST /buckets][%d] createBucketOK  %+v", 200, o.Payload)
}

func (o *CreateBucketOK) GetPayload() *models.Bucket {
	return o.Payload
}

func (o *CreateBucketOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Bucket)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBucketBadRequest creates a CreateBucketBadRequest with default headers values
func NewCreateBucketBadRequest() *CreateBucketBadRequest {
	return &CreateBucketBadRequest{}
}

/*
CreateBucketBadRequest describes a response with status code 400, with default header values.

NiFi Registry was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type CreateBucketBadRequest struct {
}

// IsSuccess returns true when this create bucket bad request response has a 2xx status code
func (o *CreateBucketBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create bucket bad request response has a 3xx status code
func (o *CreateBucketBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create bucket bad request response has a 4xx status code
func (o *CreateBucketBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create bucket bad request response has a 5xx status code
func (o *CreateBucketBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create bucket bad request response a status code equal to that given
func (o *CreateBucketBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create bucket bad request response
func (o *CreateBucketBadRequest) Code() int {
	return 400
}

func (o *CreateBucketBadRequest) Error() string {
	return fmt.Sprintf("[POST /buckets][%d] createBucketBadRequest ", 400)
}

func (o *CreateBucketBadRequest) String() string {
	return fmt.Sprintf("[POST /buckets][%d] createBucketBadRequest ", 400)
}

func (o *CreateBucketBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateBucketUnauthorized creates a CreateBucketUnauthorized with default headers values
func NewCreateBucketUnauthorized() *CreateBucketUnauthorized {
	return &CreateBucketUnauthorized{}
}

/*
CreateBucketUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type CreateBucketUnauthorized struct {
}

// IsSuccess returns true when this create bucket unauthorized response has a 2xx status code
func (o *CreateBucketUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create bucket unauthorized response has a 3xx status code
func (o *CreateBucketUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create bucket unauthorized response has a 4xx status code
func (o *CreateBucketUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create bucket unauthorized response has a 5xx status code
func (o *CreateBucketUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create bucket unauthorized response a status code equal to that given
func (o *CreateBucketUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create bucket unauthorized response
func (o *CreateBucketUnauthorized) Code() int {
	return 401
}

func (o *CreateBucketUnauthorized) Error() string {
	return fmt.Sprintf("[POST /buckets][%d] createBucketUnauthorized ", 401)
}

func (o *CreateBucketUnauthorized) String() string {
	return fmt.Sprintf("[POST /buckets][%d] createBucketUnauthorized ", 401)
}

func (o *CreateBucketUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateBucketForbidden creates a CreateBucketForbidden with default headers values
func NewCreateBucketForbidden() *CreateBucketForbidden {
	return &CreateBucketForbidden{}
}

/*
CreateBucketForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type CreateBucketForbidden struct {
}

// IsSuccess returns true when this create bucket forbidden response has a 2xx status code
func (o *CreateBucketForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create bucket forbidden response has a 3xx status code
func (o *CreateBucketForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create bucket forbidden response has a 4xx status code
func (o *CreateBucketForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create bucket forbidden response has a 5xx status code
func (o *CreateBucketForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create bucket forbidden response a status code equal to that given
func (o *CreateBucketForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create bucket forbidden response
func (o *CreateBucketForbidden) Code() int {
	return 403
}

func (o *CreateBucketForbidden) Error() string {
	return fmt.Sprintf("[POST /buckets][%d] createBucketForbidden ", 403)
}

func (o *CreateBucketForbidden) String() string {
	return fmt.Sprintf("[POST /buckets][%d] createBucketForbidden ", 403)
}

func (o *CreateBucketForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
