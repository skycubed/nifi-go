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

// GetBucketReader is a Reader for the GetBucket structure.
type GetBucketReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBucketReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBucketOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetBucketUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetBucketForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetBucketNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /buckets/{bucketId}] getBucket", response, response.Code())
	}
}

// NewGetBucketOK creates a GetBucketOK with default headers values
func NewGetBucketOK() *GetBucketOK {
	return &GetBucketOK{}
}

/*
GetBucketOK describes a response with status code 200, with default header values.

successful operation
*/
type GetBucketOK struct {
	Payload *models.Bucket
}

// IsSuccess returns true when this get bucket o k response has a 2xx status code
func (o *GetBucketOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get bucket o k response has a 3xx status code
func (o *GetBucketOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bucket o k response has a 4xx status code
func (o *GetBucketOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get bucket o k response has a 5xx status code
func (o *GetBucketOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get bucket o k response a status code equal to that given
func (o *GetBucketOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get bucket o k response
func (o *GetBucketOK) Code() int {
	return 200
}

func (o *GetBucketOK) Error() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}][%d] getBucketOK  %+v", 200, o.Payload)
}

func (o *GetBucketOK) String() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}][%d] getBucketOK  %+v", 200, o.Payload)
}

func (o *GetBucketOK) GetPayload() *models.Bucket {
	return o.Payload
}

func (o *GetBucketOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Bucket)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBucketUnauthorized creates a GetBucketUnauthorized with default headers values
func NewGetBucketUnauthorized() *GetBucketUnauthorized {
	return &GetBucketUnauthorized{}
}

/*
GetBucketUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetBucketUnauthorized struct {
}

// IsSuccess returns true when this get bucket unauthorized response has a 2xx status code
func (o *GetBucketUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get bucket unauthorized response has a 3xx status code
func (o *GetBucketUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bucket unauthorized response has a 4xx status code
func (o *GetBucketUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get bucket unauthorized response has a 5xx status code
func (o *GetBucketUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get bucket unauthorized response a status code equal to that given
func (o *GetBucketUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get bucket unauthorized response
func (o *GetBucketUnauthorized) Code() int {
	return 401
}

func (o *GetBucketUnauthorized) Error() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}][%d] getBucketUnauthorized ", 401)
}

func (o *GetBucketUnauthorized) String() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}][%d] getBucketUnauthorized ", 401)
}

func (o *GetBucketUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBucketForbidden creates a GetBucketForbidden with default headers values
func NewGetBucketForbidden() *GetBucketForbidden {
	return &GetBucketForbidden{}
}

/*
GetBucketForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetBucketForbidden struct {
}

// IsSuccess returns true when this get bucket forbidden response has a 2xx status code
func (o *GetBucketForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get bucket forbidden response has a 3xx status code
func (o *GetBucketForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bucket forbidden response has a 4xx status code
func (o *GetBucketForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get bucket forbidden response has a 5xx status code
func (o *GetBucketForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get bucket forbidden response a status code equal to that given
func (o *GetBucketForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get bucket forbidden response
func (o *GetBucketForbidden) Code() int {
	return 403
}

func (o *GetBucketForbidden) Error() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}][%d] getBucketForbidden ", 403)
}

func (o *GetBucketForbidden) String() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}][%d] getBucketForbidden ", 403)
}

func (o *GetBucketForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBucketNotFound creates a GetBucketNotFound with default headers values
func NewGetBucketNotFound() *GetBucketNotFound {
	return &GetBucketNotFound{}
}

/*
GetBucketNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetBucketNotFound struct {
}

// IsSuccess returns true when this get bucket not found response has a 2xx status code
func (o *GetBucketNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get bucket not found response has a 3xx status code
func (o *GetBucketNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bucket not found response has a 4xx status code
func (o *GetBucketNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get bucket not found response has a 5xx status code
func (o *GetBucketNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get bucket not found response a status code equal to that given
func (o *GetBucketNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get bucket not found response
func (o *GetBucketNotFound) Code() int {
	return 404
}

func (o *GetBucketNotFound) Error() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}][%d] getBucketNotFound ", 404)
}

func (o *GetBucketNotFound) String() string {
	return fmt.Sprintf("[GET /buckets/{bucketId}][%d] getBucketNotFound ", 404)
}

func (o *GetBucketNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
