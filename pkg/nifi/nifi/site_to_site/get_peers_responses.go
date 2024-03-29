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
		return nil, runtime.NewAPIError("[GET /site-to-site/peers] getPeers", response, response.Code())
	}
}

// NewGetPeersOK creates a GetPeersOK with default headers values
func NewGetPeersOK() *GetPeersOK {
	return &GetPeersOK{}
}

/*
GetPeersOK describes a response with status code 200, with default header values.

successful operation
*/
type GetPeersOK struct {
	Payload *models.PeersEntity
}

// IsSuccess returns true when this get peers o k response has a 2xx status code
func (o *GetPeersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get peers o k response has a 3xx status code
func (o *GetPeersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get peers o k response has a 4xx status code
func (o *GetPeersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get peers o k response has a 5xx status code
func (o *GetPeersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get peers o k response a status code equal to that given
func (o *GetPeersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get peers o k response
func (o *GetPeersOK) Code() int {
	return 200
}

func (o *GetPeersOK) Error() string {
	return fmt.Sprintf("[GET /site-to-site/peers][%d] getPeersOK  %+v", 200, o.Payload)
}

func (o *GetPeersOK) String() string {
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

/*
GetPeersBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetPeersBadRequest struct {
}

// IsSuccess returns true when this get peers bad request response has a 2xx status code
func (o *GetPeersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get peers bad request response has a 3xx status code
func (o *GetPeersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get peers bad request response has a 4xx status code
func (o *GetPeersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get peers bad request response has a 5xx status code
func (o *GetPeersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get peers bad request response a status code equal to that given
func (o *GetPeersBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get peers bad request response
func (o *GetPeersBadRequest) Code() int {
	return 400
}

func (o *GetPeersBadRequest) Error() string {
	return fmt.Sprintf("[GET /site-to-site/peers][%d] getPeersBadRequest ", 400)
}

func (o *GetPeersBadRequest) String() string {
	return fmt.Sprintf("[GET /site-to-site/peers][%d] getPeersBadRequest ", 400)
}

func (o *GetPeersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPeersUnauthorized creates a GetPeersUnauthorized with default headers values
func NewGetPeersUnauthorized() *GetPeersUnauthorized {
	return &GetPeersUnauthorized{}
}

/*
GetPeersUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetPeersUnauthorized struct {
}

// IsSuccess returns true when this get peers unauthorized response has a 2xx status code
func (o *GetPeersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get peers unauthorized response has a 3xx status code
func (o *GetPeersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get peers unauthorized response has a 4xx status code
func (o *GetPeersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get peers unauthorized response has a 5xx status code
func (o *GetPeersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get peers unauthorized response a status code equal to that given
func (o *GetPeersUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get peers unauthorized response
func (o *GetPeersUnauthorized) Code() int {
	return 401
}

func (o *GetPeersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /site-to-site/peers][%d] getPeersUnauthorized ", 401)
}

func (o *GetPeersUnauthorized) String() string {
	return fmt.Sprintf("[GET /site-to-site/peers][%d] getPeersUnauthorized ", 401)
}

func (o *GetPeersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPeersForbidden creates a GetPeersForbidden with default headers values
func NewGetPeersForbidden() *GetPeersForbidden {
	return &GetPeersForbidden{}
}

/*
GetPeersForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetPeersForbidden struct {
}

// IsSuccess returns true when this get peers forbidden response has a 2xx status code
func (o *GetPeersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get peers forbidden response has a 3xx status code
func (o *GetPeersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get peers forbidden response has a 4xx status code
func (o *GetPeersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get peers forbidden response has a 5xx status code
func (o *GetPeersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get peers forbidden response a status code equal to that given
func (o *GetPeersForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get peers forbidden response
func (o *GetPeersForbidden) Code() int {
	return 403
}

func (o *GetPeersForbidden) Error() string {
	return fmt.Sprintf("[GET /site-to-site/peers][%d] getPeersForbidden ", 403)
}

func (o *GetPeersForbidden) String() string {
	return fmt.Sprintf("[GET /site-to-site/peers][%d] getPeersForbidden ", 403)
}

func (o *GetPeersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPeersConflict creates a GetPeersConflict with default headers values
func NewGetPeersConflict() *GetPeersConflict {
	return &GetPeersConflict{}
}

/*
GetPeersConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetPeersConflict struct {
}

// IsSuccess returns true when this get peers conflict response has a 2xx status code
func (o *GetPeersConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get peers conflict response has a 3xx status code
func (o *GetPeersConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get peers conflict response has a 4xx status code
func (o *GetPeersConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get peers conflict response has a 5xx status code
func (o *GetPeersConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get peers conflict response a status code equal to that given
func (o *GetPeersConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the get peers conflict response
func (o *GetPeersConflict) Code() int {
	return 409
}

func (o *GetPeersConflict) Error() string {
	return fmt.Sprintf("[GET /site-to-site/peers][%d] getPeersConflict ", 409)
}

func (o *GetPeersConflict) String() string {
	return fmt.Sprintf("[GET /site-to-site/peers][%d] getPeersConflict ", 409)
}

func (o *GetPeersConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
