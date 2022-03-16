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

/* KnoxLogoutDefault describes a response with status code -1, with default header values.

successful operation
*/
type KnoxLogoutDefault struct {
	_statusCode int
}

// Code gets the status code for the knox logout default response
func (o *KnoxLogoutDefault) Code() int {
	return o._statusCode
}

func (o *KnoxLogoutDefault) Error() string {
	return fmt.Sprintf("[GET /access/knox/logout][%d] knoxLogout default ", o._statusCode)
}

func (o *KnoxLogoutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}