// Code generated by go-swagger; DO NOT EDIT.

package access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// KnoxLogoutReader is a Reader for the KnoxLogout structure.
type KnoxLogoutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KnoxLogoutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewKnoxLogoutDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewKnoxLogoutDefault creates a KnoxLogoutDefault with default headers values
func NewKnoxLogoutDefault(code int) *KnoxLogoutDefault {
	return &KnoxLogoutDefault{
		_statusCode: code,
	}
}

/*
KnoxLogoutDefault describes a response with status code -1, with default header values.

successful operation
*/
type KnoxLogoutDefault struct {
	_statusCode int
}

// Code gets the status code for the knox logout default response
func (o *KnoxLogoutDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this knox logout default response has a 2xx status code
func (o *KnoxLogoutDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this knox logout default response has a 3xx status code
func (o *KnoxLogoutDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this knox logout default response has a 4xx status code
func (o *KnoxLogoutDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this knox logout default response has a 5xx status code
func (o *KnoxLogoutDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this knox logout default response a status code equal to that given
func (o *KnoxLogoutDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *KnoxLogoutDefault) Error() string {
	return fmt.Sprintf("[GET /access/knox/logout][%d] knoxLogout default ", o._statusCode)
}

func (o *KnoxLogoutDefault) String() string {
	return fmt.Sprintf("[GET /access/knox/logout][%d] knoxLogout default ", o._statusCode)
}

func (o *KnoxLogoutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
