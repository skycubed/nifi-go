// Code generated by go-swagger; DO NOT EDIT.

package site_to_site

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// GetPeersReader is a Reader for the GetPeers structure.
type GetPeersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPeersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPeersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPeersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPeersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPeersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetPeersConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPeersOK creates a GetPeersOK with default headers values
func NewGetPeersOK() *GetPeersOK {
	return &GetPeersOK{}
}

/* GetPeersOK describes a response with status code 200, with default header values.

successful operation
*/
type GetPeersOK struct {
	Payload *models.PeersEntity
}

func (o *GetPeersOK) Error() string {
	return fmt.Sprintf("[GET /site-to-site/peers][%d] getPeersOK  %+v", 200, o.Payload)
}
func (o *GetPeersOK) GetPayload() *models.PeersEntity {
	return o.Payload
}

func (o *GetPeersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PeersEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPeersBadRequest creates a GetPeersBadRequest with default headers values
func NewGetPeersBadRequest() *GetPeersBadRequest {
	return &GetPeersBadRequest{}
}

/* GetPeersBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetPeersBadRequest struct {
}

func (o *GetPeersBadRequest) Error() string {
	return fmt.Sprintf("[GET /site-to-site/peers][%d] getPeersBadRequest ", 400)
}

func (o *GetPeersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPeersUnauthorized creates a GetPeersUnauthorized with default headers values
func NewGetPeersUnauthorized() *GetPeersUnauthorized {
	return &GetPeersUnauthorized{}
}

/* GetPeersUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetPeersUnauthorized struct {
}

func (o *GetPeersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /site-to-site/peers][%d] getPeersUnauthorized ", 401)
}

func (o *GetPeersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPeersForbidden creates a GetPeersForbidden with default headers values
func NewGetPeersForbidden() *GetPeersForbidden {
	return &GetPeersForbidden{}
}

/* GetPeersForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetPeersForbidden struct {
}

func (o *GetPeersForbidden) Error() string {
	return fmt.Sprintf("[GET /site-to-site/peers][%d] getPeersForbidden ", 403)
}

func (o *GetPeersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPeersConflict creates a GetPeersConflict with default headers values
func NewGetPeersConflict() *GetPeersConflict {
	return &GetPeersConflict{}
}

/* GetPeersConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetPeersConflict struct {
}

func (o *GetPeersConflict) Error() string {
	return fmt.Sprintf("[GET /site-to-site/peers][%d] getPeersConflict ", 409)
}

func (o *GetPeersConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
