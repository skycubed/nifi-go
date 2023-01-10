// Code generated by go-swagger; DO NOT EDIT.

package access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// OidcLogoutCallbackReader is a Reader for the OidcLogoutCallback structure.
type OidcLogoutCallbackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OidcLogoutCallbackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewOidcLogoutCallbackDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewOidcLogoutCallbackDefault creates a OidcLogoutCallbackDefault with default headers values
func NewOidcLogoutCallbackDefault(code int) *OidcLogoutCallbackDefault {
	return &OidcLogoutCallbackDefault{
		_statusCode: code,
	}
}

/*
OidcLogoutCallbackDefault describes a response with status code -1, with default header values.

successful operation
*/
type OidcLogoutCallbackDefault struct {
	_statusCode int
}

// Code gets the status code for the oidc logout callback default response
func (o *OidcLogoutCallbackDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this oidc logout callback default response has a 2xx status code
func (o *OidcLogoutCallbackDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this oidc logout callback default response has a 3xx status code
func (o *OidcLogoutCallbackDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this oidc logout callback default response has a 4xx status code
func (o *OidcLogoutCallbackDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this oidc logout callback default response has a 5xx status code
func (o *OidcLogoutCallbackDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this oidc logout callback default response a status code equal to that given
func (o *OidcLogoutCallbackDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *OidcLogoutCallbackDefault) Error() string {
	return fmt.Sprintf("[GET /access/oidc/logout/callback][%d] oidcLogoutCallback default ", o._statusCode)
}

func (o *OidcLogoutCallbackDefault) String() string {
	return fmt.Sprintf("[GET /access/oidc/logout/callback][%d] oidcLogoutCallback default ", o._statusCode)
}

func (o *OidcLogoutCallbackDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}