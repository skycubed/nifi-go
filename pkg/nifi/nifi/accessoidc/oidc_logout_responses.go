// Code generated by go-swagger; DO NOT EDIT.

package accessoidc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// OidcLogoutReader is a Reader for the OidcLogout structure.
type OidcLogoutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OidcLogoutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewOidcLogoutDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewOidcLogoutDefault creates a OidcLogoutDefault with default headers values
func NewOidcLogoutDefault(code int) *OidcLogoutDefault {
	return &OidcLogoutDefault{
		_statusCode: code,
	}
}

/*
OidcLogoutDefault describes a response with status code -1, with default header values.

successful operation
*/
type OidcLogoutDefault struct {
	_statusCode int
}

// Code gets the status code for the oidc logout default response
func (o *OidcLogoutDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this oidc logout default response has a 2xx status code
func (o *OidcLogoutDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this oidc logout default response has a 3xx status code
func (o *OidcLogoutDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this oidc logout default response has a 4xx status code
func (o *OidcLogoutDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this oidc logout default response has a 5xx status code
func (o *OidcLogoutDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this oidc logout default response a status code equal to that given
func (o *OidcLogoutDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *OidcLogoutDefault) Error() string {
	return fmt.Sprintf("[GET /access/oidc/logout][%d] oidcLogout default ", o._statusCode)
}

func (o *OidcLogoutDefault) String() string {
	return fmt.Sprintf("[GET /access/oidc/logout][%d] oidcLogout default ", o._statusCode)
}

func (o *OidcLogoutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
