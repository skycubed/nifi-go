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
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLogOutOK creates a LogOutOK with default headers values
func NewLogOutOK() *LogOutOK {
	return &LogOutOK{}
}

/* LogOutOK describes a response with status code 200, with default header values.

User was logged out successfully.
*/
type LogOutOK struct {
}

func (o *LogOutOK) Error() string {
	return fmt.Sprintf("[DELETE /access/logout][%d] logOutOK ", 200)
}

func (o *LogOutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLogOutUnauthorized creates a LogOutUnauthorized with default headers values
func NewLogOutUnauthorized() *LogOutUnauthorized {
	return &LogOutUnauthorized{}
}

/* LogOutUnauthorized describes a response with status code 401, with default header values.

Authentication token provided was empty or not in the correct JWT format.
*/
type LogOutUnauthorized struct {
}

func (o *LogOutUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /access/logout][%d] logOutUnauthorized ", 401)
}

func (o *LogOutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLogOutInternalServerError creates a LogOutInternalServerError with default headers values
func NewLogOutInternalServerError() *LogOutInternalServerError {
	return &LogOutInternalServerError{}
}

/* LogOutInternalServerError describes a response with status code 500, with default header values.

Client failed to log out.
*/
type LogOutInternalServerError struct {
}

func (o *LogOutInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /access/logout][%d] logOutInternalServerError ", 500)
}

func (o *LogOutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
