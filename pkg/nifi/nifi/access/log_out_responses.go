// Code generated by go-swagger; DO NOT EDIT.

package access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// LogOutReader is a Reader for the LogOut structure.
type LogOutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LogOutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLogOutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewLogOutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewLogOutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /access/logout] logOut", response, response.Code())
	}
}

// NewLogOutOK creates a LogOutOK with default headers values
func NewLogOutOK() *LogOutOK {
	return &LogOutOK{}
}

/*
LogOutOK describes a response with status code 200, with default header values.

User was logged out successfully.
*/
type LogOutOK struct {
}

// IsSuccess returns true when this log out o k response has a 2xx status code
func (o *LogOutOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this log out o k response has a 3xx status code
func (o *LogOutOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this log out o k response has a 4xx status code
func (o *LogOutOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this log out o k response has a 5xx status code
func (o *LogOutOK) IsServerError() bool {
	return false
}

// IsCode returns true when this log out o k response a status code equal to that given
func (o *LogOutOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the log out o k response
func (o *LogOutOK) Code() int {
	return 200
}

func (o *LogOutOK) Error() string {
	return fmt.Sprintf("[DELETE /access/logout][%d] logOutOK ", 200)
}

func (o *LogOutOK) String() string {
	return fmt.Sprintf("[DELETE /access/logout][%d] logOutOK ", 200)
}

func (o *LogOutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLogOutUnauthorized creates a LogOutUnauthorized with default headers values
func NewLogOutUnauthorized() *LogOutUnauthorized {
	return &LogOutUnauthorized{}
}

/*
LogOutUnauthorized describes a response with status code 401, with default header values.

Authentication token provided was empty or not in the correct JWT format.
*/
type LogOutUnauthorized struct {
}

// IsSuccess returns true when this log out unauthorized response has a 2xx status code
func (o *LogOutUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this log out unauthorized response has a 3xx status code
func (o *LogOutUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this log out unauthorized response has a 4xx status code
func (o *LogOutUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this log out unauthorized response has a 5xx status code
func (o *LogOutUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this log out unauthorized response a status code equal to that given
func (o *LogOutUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the log out unauthorized response
func (o *LogOutUnauthorized) Code() int {
	return 401
}

func (o *LogOutUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /access/logout][%d] logOutUnauthorized ", 401)
}

func (o *LogOutUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /access/logout][%d] logOutUnauthorized ", 401)
}

func (o *LogOutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLogOutInternalServerError creates a LogOutInternalServerError with default headers values
func NewLogOutInternalServerError() *LogOutInternalServerError {
	return &LogOutInternalServerError{}
}

/*
LogOutInternalServerError describes a response with status code 500, with default header values.

Client failed to log out.
*/
type LogOutInternalServerError struct {
}

// IsSuccess returns true when this log out internal server error response has a 2xx status code
func (o *LogOutInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this log out internal server error response has a 3xx status code
func (o *LogOutInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this log out internal server error response has a 4xx status code
func (o *LogOutInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this log out internal server error response has a 5xx status code
func (o *LogOutInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this log out internal server error response a status code equal to that given
func (o *LogOutInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the log out internal server error response
func (o *LogOutInternalServerError) Code() int {
	return 500
}

func (o *LogOutInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /access/logout][%d] logOutInternalServerError ", 500)
}

func (o *LogOutInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /access/logout][%d] logOutInternalServerError ", 500)
}

func (o *LogOutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
