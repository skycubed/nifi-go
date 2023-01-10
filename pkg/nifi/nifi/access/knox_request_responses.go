// Code generated by go-swagger; DO NOT EDIT.

package access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// KnoxRequestReader is a Reader for the KnoxRequest structure.
type KnoxRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KnoxRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewKnoxRequestDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewKnoxRequestDefault creates a KnoxRequestDefault with default headers values
func NewKnoxRequestDefault(code int) *KnoxRequestDefault {
	return &KnoxRequestDefault{
		_statusCode: code,
	}
}

/*
KnoxRequestDefault describes a response with status code -1, with default header values.

successful operation
*/
type KnoxRequestDefault struct {
	_statusCode int
}

// Code gets the status code for the knox request default response
func (o *KnoxRequestDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this knox request default response has a 2xx status code
func (o *KnoxRequestDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this knox request default response has a 3xx status code
func (o *KnoxRequestDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this knox request default response has a 4xx status code
func (o *KnoxRequestDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this knox request default response has a 5xx status code
func (o *KnoxRequestDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this knox request default response a status code equal to that given
func (o *KnoxRequestDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *KnoxRequestDefault) Error() string {
	return fmt.Sprintf("[GET /access/knox/request][%d] knoxRequest default ", o._statusCode)
}

func (o *KnoxRequestDefault) String() string {
	return fmt.Sprintf("[GET /access/knox/request][%d] knoxRequest default ", o._statusCode)
}

func (o *KnoxRequestDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
