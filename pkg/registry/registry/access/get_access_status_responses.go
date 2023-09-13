// Code generated by go-swagger; DO NOT EDIT.

package access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/registry/models"
)

// GetAccessStatusReader is a Reader for the GetAccessStatus structure.
type GetAccessStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccessStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccessStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewGetAccessStatusConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /access] getAccessStatus", response, response.Code())
	}
}

// NewGetAccessStatusOK creates a GetAccessStatusOK with default headers values
func NewGetAccessStatusOK() *GetAccessStatusOK {
	return &GetAccessStatusOK{}
}

/*
GetAccessStatusOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAccessStatusOK struct {
	Payload *models.CurrentUser
}

// IsSuccess returns true when this get access status o k response has a 2xx status code
func (o *GetAccessStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get access status o k response has a 3xx status code
func (o *GetAccessStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get access status o k response has a 4xx status code
func (o *GetAccessStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get access status o k response has a 5xx status code
func (o *GetAccessStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get access status o k response a status code equal to that given
func (o *GetAccessStatusOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get access status o k response
func (o *GetAccessStatusOK) Code() int {
	return 200
}

func (o *GetAccessStatusOK) Error() string {
	return fmt.Sprintf("[GET /access][%d] getAccessStatusOK  %+v", 200, o.Payload)
}

func (o *GetAccessStatusOK) String() string {
	return fmt.Sprintf("[GET /access][%d] getAccessStatusOK  %+v", 200, o.Payload)
}

func (o *GetAccessStatusOK) GetPayload() *models.CurrentUser {
	return o.Payload
}

func (o *GetAccessStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CurrentUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccessStatusConflict creates a GetAccessStatusConflict with default headers values
func NewGetAccessStatusConflict() *GetAccessStatusConflict {
	return &GetAccessStatusConflict{}
}

/*
GetAccessStatusConflict describes a response with status code 409, with default header values.

NiFi Registry was unable to complete the request because it assumes a server state that is not valid. The NiFi Registry might be running unsecured.
*/
type GetAccessStatusConflict struct {
}

// IsSuccess returns true when this get access status conflict response has a 2xx status code
func (o *GetAccessStatusConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get access status conflict response has a 3xx status code
func (o *GetAccessStatusConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get access status conflict response has a 4xx status code
func (o *GetAccessStatusConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get access status conflict response has a 5xx status code
func (o *GetAccessStatusConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get access status conflict response a status code equal to that given
func (o *GetAccessStatusConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the get access status conflict response
func (o *GetAccessStatusConflict) Code() int {
	return 409
}

func (o *GetAccessStatusConflict) Error() string {
	return fmt.Sprintf("[GET /access][%d] getAccessStatusConflict ", 409)
}

func (o *GetAccessStatusConflict) String() string {
	return fmt.Sprintf("[GET /access][%d] getAccessStatusConflict ", 409)
}

func (o *GetAccessStatusConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
