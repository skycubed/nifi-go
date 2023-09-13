// Code generated by go-swagger; DO NOT EDIT.

package items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/registry/models"
)

// GetAvailableBucketItemFieldsReader is a Reader for the GetAvailableBucketItemFields structure.
type GetAvailableBucketItemFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAvailableBucketItemFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAvailableBucketItemFieldsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /items/fields] getAvailableBucketItemFields", response, response.Code())
	}
}

// NewGetAvailableBucketItemFieldsOK creates a GetAvailableBucketItemFieldsOK with default headers values
func NewGetAvailableBucketItemFieldsOK() *GetAvailableBucketItemFieldsOK {
	return &GetAvailableBucketItemFieldsOK{}
}

/*
GetAvailableBucketItemFieldsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAvailableBucketItemFieldsOK struct {
	Payload *models.Fields
}

// IsSuccess returns true when this get available bucket item fields o k response has a 2xx status code
func (o *GetAvailableBucketItemFieldsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get available bucket item fields o k response has a 3xx status code
func (o *GetAvailableBucketItemFieldsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get available bucket item fields o k response has a 4xx status code
func (o *GetAvailableBucketItemFieldsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get available bucket item fields o k response has a 5xx status code
func (o *GetAvailableBucketItemFieldsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get available bucket item fields o k response a status code equal to that given
func (o *GetAvailableBucketItemFieldsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get available bucket item fields o k response
func (o *GetAvailableBucketItemFieldsOK) Code() int {
	return 200
}

func (o *GetAvailableBucketItemFieldsOK) Error() string {
	return fmt.Sprintf("[GET /items/fields][%d] getAvailableBucketItemFieldsOK  %+v", 200, o.Payload)
}

func (o *GetAvailableBucketItemFieldsOK) String() string {
	return fmt.Sprintf("[GET /items/fields][%d] getAvailableBucketItemFieldsOK  %+v", 200, o.Payload)
}

func (o *GetAvailableBucketItemFieldsOK) GetPayload() *models.Fields {
	return o.Payload
}

func (o *GetAvailableBucketItemFieldsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Fields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
