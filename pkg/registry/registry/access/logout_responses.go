// Code generated by go-swagger; DO NOT EDIT.

package access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// LogoutReader is a Reader for the Logout structure.
type LogoutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LogoutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLogoutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewLogoutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewLogoutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLogoutOK creates a LogoutOK with default headers values
func NewLogoutOK() *LogoutOK {
	return &LogoutOK{}
}

/*
LogoutOK describes a response with status code 200, with default header values.

User was logged out successfully.
*/
type LogoutOK struct {
}

// IsSuccess returns true when this logout o k response has a 2xx status code
func (o *LogoutOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this logout o k response has a 3xx status code
func (o *LogoutOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this logout o k response has a 4xx status code
func (o *LogoutOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this logout o k response has a 5xx status code
func (o *LogoutOK) IsServerError() bool {
	return false
}

// IsCode returns true when this logout o k response a status code equal to that given
func (o *LogoutOK) IsCode(code int) bool {
	return code == 200
}

func (o *LogoutOK) Error() string {
	return fmt.Sprintf("[DELETE /access/logout][%d] logoutOK ", 200)
}

func (o *LogoutOK) String() string {
	return fmt.Sprintf("[DELETE /access/logout][%d] logoutOK ", 200)
}

func (o *LogoutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLogoutUnauthorized creates a LogoutUnauthorized with default headers values
func NewLogoutUnauthorized() *LogoutUnauthorized {
	return &LogoutUnauthorized{}
}

/*
LogoutUnauthorized describes a response with status code 401, with default header values.

Authentication token provided was empty or not in the correct JWT format.
*/
type LogoutUnauthorized struct {
}

// IsSuccess returns true when this logout unauthorized response has a 2xx status code
func (o *LogoutUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this logout unauthorized response has a 3xx status code
func (o *LogoutUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this logout unauthorized response has a 4xx status code
func (o *LogoutUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this logout unauthorized response has a 5xx status code
func (o *LogoutUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this logout unauthorized response a status code equal to that given
func (o *LogoutUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *LogoutUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /access/logout][%d] logoutUnauthorized ", 401)
}

func (o *LogoutUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /access/logout][%d] logoutUnauthorized ", 401)
}

func (o *LogoutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLogoutInternalServerError creates a LogoutInternalServerError with default headers values
func NewLogoutInternalServerError() *LogoutInternalServerError {
	return &LogoutInternalServerError{}
}

/*
LogoutInternalServerError describes a response with status code 500, with default header values.

Client failed to log out.
*/
type LogoutInternalServerError struct {
}

// IsSuccess returns true when this logout internal server error response has a 2xx status code
func (o *LogoutInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this logout internal server error response has a 3xx status code
func (o *LogoutInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this logout internal server error response has a 4xx status code
func (o *LogoutInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this logout internal server error response has a 5xx status code
func (o *LogoutInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this logout internal server error response a status code equal to that given
func (o *LogoutInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *LogoutInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /access/logout][%d] logoutInternalServerError ", 500)
}

func (o *LogoutInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /access/logout][%d] logoutInternalServerError ", 500)
}

func (o *LogoutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}