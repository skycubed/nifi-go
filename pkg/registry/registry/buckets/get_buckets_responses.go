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

// GetBucketsReader is a Reader for the GetBuckets structure.
type GetBucketsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBucketsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBucketsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetBucketsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /buckets] getBuckets", response, response.Code())
	}
}

// NewGetBucketsOK creates a GetBucketsOK with default headers values
func NewGetBucketsOK() *GetBucketsOK {
	return &GetBucketsOK{}
}

/*
GetBucketsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetBucketsOK struct {
	Payload []*models.Bucket
}

// IsSuccess returns true when this get buckets o k response has a 2xx status code
func (o *GetBucketsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get buckets o k response has a 3xx status code
func (o *GetBucketsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get buckets o k response has a 4xx status code
func (o *GetBucketsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get buckets o k response has a 5xx status code
func (o *GetBucketsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get buckets o k response a status code equal to that given
func (o *GetBucketsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get buckets o k response
func (o *GetBucketsOK) Code() int {
	return 200
}

func (o *GetBucketsOK) Error() string {
	return fmt.Sprintf("[GET /buckets][%d] getBucketsOK  %+v", 200, o.Payload)
}

func (o *GetBucketsOK) String() string {
	return fmt.Sprintf("[GET /buckets][%d] getBucketsOK  %+v", 200, o.Payload)
}

func (o *GetBucketsOK) GetPayload() []*models.Bucket {
	return o.Payload
}

func (o *GetBucketsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBucketsUnauthorized creates a GetBucketsUnauthorized with default headers values
func NewGetBucketsUnauthorized() *GetBucketsUnauthorized {
	return &GetBucketsUnauthorized{}
}

/*
GetBucketsUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetBucketsUnauthorized struct {
}

// IsSuccess returns true when this get buckets unauthorized response has a 2xx status code
func (o *GetBucketsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get buckets unauthorized response has a 3xx status code
func (o *GetBucketsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get buckets unauthorized response has a 4xx status code
func (o *GetBucketsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get buckets unauthorized response has a 5xx status code
func (o *GetBucketsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get buckets unauthorized response a status code equal to that given
func (o *GetBucketsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get buckets unauthorized response
func (o *GetBucketsUnauthorized) Code() int {
	return 401
}

func (o *GetBucketsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /buckets][%d] getBucketsUnauthorized ", 401)
}

func (o *GetBucketsUnauthorized) String() string {
	return fmt.Sprintf("[GET /buckets][%d] getBucketsUnauthorized ", 401)
}

func (o *GetBucketsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
