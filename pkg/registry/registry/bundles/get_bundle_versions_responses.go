// Code generated by go-swagger; DO NOT EDIT.

package bundles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/registry/models"
)

// GetBundleVersionsReader is a Reader for the GetBundleVersions structure.
type GetBundleVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBundleVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBundleVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetBundleVersionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBundleVersionsOK creates a GetBundleVersionsOK with default headers values
func NewGetBundleVersionsOK() *GetBundleVersionsOK {
	return &GetBundleVersionsOK{}
}

/* GetBundleVersionsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetBundleVersionsOK struct {
	Payload []*models.BundleVersionMetadata
}

func (o *GetBundleVersionsOK) Error() string {
	return fmt.Sprintf("[GET /bundles/versions][%d] getBundleVersionsOK  %+v", 200, o.Payload)
}
func (o *GetBundleVersionsOK) GetPayload() []*models.BundleVersionMetadata {
	return o.Payload
}

func (o *GetBundleVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBundleVersionsUnauthorized creates a GetBundleVersionsUnauthorized with default headers values
func NewGetBundleVersionsUnauthorized() *GetBundleVersionsUnauthorized {
	return &GetBundleVersionsUnauthorized{}
}

/* GetBundleVersionsUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetBundleVersionsUnauthorized struct {
}

func (o *GetBundleVersionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /bundles/versions][%d] getBundleVersionsUnauthorized ", 401)
}

func (o *GetBundleVersionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
