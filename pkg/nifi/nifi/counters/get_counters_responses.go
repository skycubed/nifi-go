// Code generated by go-swagger; DO NOT EDIT.

package counters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// GetCountersReader is a Reader for the GetCounters structure.
type GetCountersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCountersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCountersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCountersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCountersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCountersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetCountersConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCountersOK creates a GetCountersOK with default headers values
func NewGetCountersOK() *GetCountersOK {
	return &GetCountersOK{}
}

/* GetCountersOK describes a response with status code 200, with default header values.

successful operation
*/
type GetCountersOK struct {
	Payload *models.CountersEntity
}

func (o *GetCountersOK) Error() string {
	return fmt.Sprintf("[GET /counters][%d] getCountersOK  %+v", 200, o.Payload)
}
func (o *GetCountersOK) GetPayload() *models.CountersEntity {
	return o.Payload
}

func (o *GetCountersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CountersEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCountersBadRequest creates a GetCountersBadRequest with default headers values
func NewGetCountersBadRequest() *GetCountersBadRequest {
	return &GetCountersBadRequest{}
}

/* GetCountersBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetCountersBadRequest struct {
}

func (o *GetCountersBadRequest) Error() string {
	return fmt.Sprintf("[GET /counters][%d] getCountersBadRequest ", 400)
}

func (o *GetCountersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCountersUnauthorized creates a GetCountersUnauthorized with default headers values
func NewGetCountersUnauthorized() *GetCountersUnauthorized {
	return &GetCountersUnauthorized{}
}

/* GetCountersUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetCountersUnauthorized struct {
}

func (o *GetCountersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /counters][%d] getCountersUnauthorized ", 401)
}

func (o *GetCountersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCountersForbidden creates a GetCountersForbidden with default headers values
func NewGetCountersForbidden() *GetCountersForbidden {
	return &GetCountersForbidden{}
}

/* GetCountersForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetCountersForbidden struct {
}

func (o *GetCountersForbidden) Error() string {
	return fmt.Sprintf("[GET /counters][%d] getCountersForbidden ", 403)
}

func (o *GetCountersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCountersConflict creates a GetCountersConflict with default headers values
func NewGetCountersConflict() *GetCountersConflict {
	return &GetCountersConflict{}
}

/* GetCountersConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetCountersConflict struct {
}

func (o *GetCountersConflict) Error() string {
	return fmt.Sprintf("[GET /counters][%d] getCountersConflict ", 409)
}

func (o *GetCountersConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}