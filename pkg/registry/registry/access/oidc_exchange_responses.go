// Code generated by go-swagger; DO NOT EDIT.

package access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// OidcExchangeReader is a Reader for the OidcExchange structure.
type OidcExchangeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OidcExchangeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOidcExchangeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /access/oidc/exchange] oidcExchange", response, response.Code())
	}
}

// NewOidcExchangeOK creates a OidcExchangeOK with default headers values
func NewOidcExchangeOK() *OidcExchangeOK {
	return &OidcExchangeOK{}
}

/*
OidcExchangeOK describes a response with status code 200, with default header values.

successful operation
*/
type OidcExchangeOK struct {
	Payload string
}

// IsSuccess returns true when this oidc exchange o k response has a 2xx status code
func (o *OidcExchangeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this oidc exchange o k response has a 3xx status code
func (o *OidcExchangeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this oidc exchange o k response has a 4xx status code
func (o *OidcExchangeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this oidc exchange o k response has a 5xx status code
func (o *OidcExchangeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this oidc exchange o k response a status code equal to that given
func (o *OidcExchangeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the oidc exchange o k response
func (o *OidcExchangeOK) Code() int {
	return 200
}

func (o *OidcExchangeOK) Error() string {
	return fmt.Sprintf("[POST /access/oidc/exchange][%d] oidcExchangeOK  %+v", 200, o.Payload)
}

func (o *OidcExchangeOK) String() string {
	return fmt.Sprintf("[POST /access/oidc/exchange][%d] oidcExchangeOK  %+v", 200, o.Payload)
}

func (o *OidcExchangeOK) GetPayload() string {
	return o.Payload
}

func (o *OidcExchangeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
