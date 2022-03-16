// Code generated by go-swagger; DO NOT EDIT.

package data_transfer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ReceiveFlowFilesReader is a Reader for the ReceiveFlowFiles structure.
type ReceiveFlowFilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReceiveFlowFilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReceiveFlowFilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReceiveFlowFilesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewReceiveFlowFilesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewReceiveFlowFilesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReceiveFlowFilesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewReceiveFlowFilesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewReceiveFlowFilesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReceiveFlowFilesOK creates a ReceiveFlowFilesOK with default headers values
func NewReceiveFlowFilesOK() *ReceiveFlowFilesOK {
	return &ReceiveFlowFilesOK{}
}

/* ReceiveFlowFilesOK describes a response with status code 200, with default header values.

successful operation
*/
type ReceiveFlowFilesOK struct {
	Payload string
}

func (o *ReceiveFlowFilesOK) Error() string {
	return fmt.Sprintf("[POST /data-transfer/input-ports/{portId}/transactions/{transactionId}/flow-files][%d] receiveFlowFilesOK  %+v", 200, o.Payload)
}
func (o *ReceiveFlowFilesOK) GetPayload() string {
	return o.Payload
}

func (o *ReceiveFlowFilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReceiveFlowFilesBadRequest creates a ReceiveFlowFilesBadRequest with default headers values
func NewReceiveFlowFilesBadRequest() *ReceiveFlowFilesBadRequest {
	return &ReceiveFlowFilesBadRequest{}
}

/* ReceiveFlowFilesBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type ReceiveFlowFilesBadRequest struct {
}

func (o *ReceiveFlowFilesBadRequest) Error() string {
	return fmt.Sprintf("[POST /data-transfer/input-ports/{portId}/transactions/{transactionId}/flow-files][%d] receiveFlowFilesBadRequest ", 400)
}

func (o *ReceiveFlowFilesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReceiveFlowFilesUnauthorized creates a ReceiveFlowFilesUnauthorized with default headers values
func NewReceiveFlowFilesUnauthorized() *ReceiveFlowFilesUnauthorized {
	return &ReceiveFlowFilesUnauthorized{}
}

/* ReceiveFlowFilesUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type ReceiveFlowFilesUnauthorized struct {
}

func (o *ReceiveFlowFilesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /data-transfer/input-ports/{portId}/transactions/{transactionId}/flow-files][%d] receiveFlowFilesUnauthorized ", 401)
}

func (o *ReceiveFlowFilesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReceiveFlowFilesForbidden creates a ReceiveFlowFilesForbidden with default headers values
func NewReceiveFlowFilesForbidden() *ReceiveFlowFilesForbidden {
	return &ReceiveFlowFilesForbidden{}
}

/* ReceiveFlowFilesForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type ReceiveFlowFilesForbidden struct {
}

func (o *ReceiveFlowFilesForbidden) Error() string {
	return fmt.Sprintf("[POST /data-transfer/input-ports/{portId}/transactions/{transactionId}/flow-files][%d] receiveFlowFilesForbidden ", 403)
}

func (o *ReceiveFlowFilesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReceiveFlowFilesNotFound creates a ReceiveFlowFilesNotFound with default headers values
func NewReceiveFlowFilesNotFound() *ReceiveFlowFilesNotFound {
	return &ReceiveFlowFilesNotFound{}
}

/* ReceiveFlowFilesNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type ReceiveFlowFilesNotFound struct {
}

func (o *ReceiveFlowFilesNotFound) Error() string {
	return fmt.Sprintf("[POST /data-transfer/input-ports/{portId}/transactions/{transactionId}/flow-files][%d] receiveFlowFilesNotFound ", 404)
}

func (o *ReceiveFlowFilesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReceiveFlowFilesConflict creates a ReceiveFlowFilesConflict with default headers values
func NewReceiveFlowFilesConflict() *ReceiveFlowFilesConflict {
	return &ReceiveFlowFilesConflict{}
}

/* ReceiveFlowFilesConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type ReceiveFlowFilesConflict struct {
}

func (o *ReceiveFlowFilesConflict) Error() string {
	return fmt.Sprintf("[POST /data-transfer/input-ports/{portId}/transactions/{transactionId}/flow-files][%d] receiveFlowFilesConflict ", 409)
}

func (o *ReceiveFlowFilesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReceiveFlowFilesServiceUnavailable creates a ReceiveFlowFilesServiceUnavailable with default headers values
func NewReceiveFlowFilesServiceUnavailable() *ReceiveFlowFilesServiceUnavailable {
	return &ReceiveFlowFilesServiceUnavailable{}
}

/* ReceiveFlowFilesServiceUnavailable describes a response with status code 503, with default header values.

NiFi instance is not ready for serving request, or temporarily overloaded. Retrying the same request later may be successful
*/
type ReceiveFlowFilesServiceUnavailable struct {
}

func (o *ReceiveFlowFilesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /data-transfer/input-ports/{portId}/transactions/{transactionId}/flow-files][%d] receiveFlowFilesServiceUnavailable ", 503)
}

func (o *ReceiveFlowFilesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}