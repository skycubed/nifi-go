// Code generated by go-swagger; DO NOT EDIT.

package accesssaml

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SamlLocalLogoutReader is a Reader for the SamlLocalLogout structure.
type SamlLocalLogoutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SamlLocalLogoutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewSamlLocalLogoutDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewSamlLocalLogoutDefault creates a SamlLocalLogoutDefault with default headers values
func NewSamlLocalLogoutDefault(code int) *SamlLocalLogoutDefault {
	return &SamlLocalLogoutDefault{
		_statusCode: code,
	}
}

/*
SamlLocalLogoutDefault describes a response with status code -1, with default header values.

successful operation
*/
type SamlLocalLogoutDefault struct {
	_statusCode int
}

// IsSuccess returns true when this saml local logout default response has a 2xx status code
func (o *SamlLocalLogoutDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this saml local logout default response has a 3xx status code
func (o *SamlLocalLogoutDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this saml local logout default response has a 4xx status code
func (o *SamlLocalLogoutDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this saml local logout default response has a 5xx status code
func (o *SamlLocalLogoutDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this saml local logout default response a status code equal to that given
func (o *SamlLocalLogoutDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the saml local logout default response
func (o *SamlLocalLogoutDefault) Code() int {
	return o._statusCode
}

func (o *SamlLocalLogoutDefault) Error() string {
	return fmt.Sprintf("[GET /access/saml/local-logout][%d] samlLocalLogout default ", o._statusCode)
}

func (o *SamlLocalLogoutDefault) String() string {
	return fmt.Sprintf("[GET /access/saml/local-logout][%d] samlLocalLogout default ", o._statusCode)
}

func (o *SamlLocalLogoutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
