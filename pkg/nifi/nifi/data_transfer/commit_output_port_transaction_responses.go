// Code generated by go-swagger; DO NOT EDIT.

package data_transfer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// CommitOutputPortTransactionReader is a Reader for the CommitOutputPortTransaction structure.
type CommitOutputPortTransactionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CommitOutputPortTransactionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCommitOutputPortTransactionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCommitOutputPortTransactionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCommitOutputPortTransactionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCommitOutputPortTransactionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCommitOutputPortTransactionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCommitOutputPortTransactionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewCommitOutputPortTransactionServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCommitOutputPortTransactionOK creates a CommitOutputPortTransactionOK with default headers values
func NewCommitOutputPortTransactionOK() *CommitOutputPortTransactionOK {
	return &CommitOutputPortTransactionOK{}
}

/*
CommitOutputPortTransactionOK describes a response with status code 200, with default header values.

successful operation
*/
type CommitOutputPortTransactionOK struct {
	Payload *models.TransactionResultEntity
}

// IsSuccess returns true when this commit output port transaction o k response has a 2xx status code
func (o *CommitOutputPortTransactionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this commit output port transaction o k response has a 3xx status code
func (o *CommitOutputPortTransactionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this commit output port transaction o k response has a 4xx status code
func (o *CommitOutputPortTransactionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this commit output port transaction o k response has a 5xx status code
func (o *CommitOutputPortTransactionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this commit output port transaction o k response a status code equal to that given
func (o *CommitOutputPortTransactionOK) IsCode(code int) bool {
	return code == 200
}

func (o *CommitOutputPortTransactionOK) Error() string {
	return fmt.Sprintf("[DELETE /data-transfer/output-ports/{portId}/transactions/{transactionId}][%d] commitOutputPortTransactionOK  %+v", 200, o.Payload)
}

func (o *CommitOutputPortTransactionOK) String() string {
	return fmt.Sprintf("[DELETE /data-transfer/output-ports/{portId}/transactions/{transactionId}][%d] commitOutputPortTransactionOK  %+v", 200, o.Payload)
}

func (o *CommitOutputPortTransactionOK) GetPayload() *models.TransactionResultEntity {
	return o.Payload
}

func (o *CommitOutputPortTransactionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TransactionResultEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCommitOutputPortTransactionBadRequest creates a CommitOutputPortTransactionBadRequest with default headers values
func NewCommitOutputPortTransactionBadRequest() *CommitOutputPortTransactionBadRequest {
	return &CommitOutputPortTransactionBadRequest{}
}

/*
CommitOutputPortTransactionBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type CommitOutputPortTransactionBadRequest struct {
}

// IsSuccess returns true when this commit output port transaction bad request response has a 2xx status code
func (o *CommitOutputPortTransactionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this commit output port transaction bad request response has a 3xx status code
func (o *CommitOutputPortTransactionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this commit output port transaction bad request response has a 4xx status code
func (o *CommitOutputPortTransactionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this commit output port transaction bad request response has a 5xx status code
func (o *CommitOutputPortTransactionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this commit output port transaction bad request response a status code equal to that given
func (o *CommitOutputPortTransactionBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CommitOutputPortTransactionBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /data-transfer/output-ports/{portId}/transactions/{transactionId}][%d] commitOutputPortTransactionBadRequest ", 400)
}

func (o *CommitOutputPortTransactionBadRequest) String() string {
	return fmt.Sprintf("[DELETE /data-transfer/output-ports/{portId}/transactions/{transactionId}][%d] commitOutputPortTransactionBadRequest ", 400)
}

func (o *CommitOutputPortTransactionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCommitOutputPortTransactionUnauthorized creates a CommitOutputPortTransactionUnauthorized with default headers values
func NewCommitOutputPortTransactionUnauthorized() *CommitOutputPortTransactionUnauthorized {
	return &CommitOutputPortTransactionUnauthorized{}
}

/*
CommitOutputPortTransactionUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type CommitOutputPortTransactionUnauthorized struct {
}

// IsSuccess returns true when this commit output port transaction unauthorized response has a 2xx status code
func (o *CommitOutputPortTransactionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this commit output port transaction unauthorized response has a 3xx status code
func (o *CommitOutputPortTransactionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this commit output port transaction unauthorized response has a 4xx status code
func (o *CommitOutputPortTransactionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this commit output port transaction unauthorized response has a 5xx status code
func (o *CommitOutputPortTransactionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this commit output port transaction unauthorized response a status code equal to that given
func (o *CommitOutputPortTransactionUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CommitOutputPortTransactionUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /data-transfer/output-ports/{portId}/transactions/{transactionId}][%d] commitOutputPortTransactionUnauthorized ", 401)
}

func (o *CommitOutputPortTransactionUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /data-transfer/output-ports/{portId}/transactions/{transactionId}][%d] commitOutputPortTransactionUnauthorized ", 401)
}

func (o *CommitOutputPortTransactionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCommitOutputPortTransactionForbidden creates a CommitOutputPortTransactionForbidden with default headers values
func NewCommitOutputPortTransactionForbidden() *CommitOutputPortTransactionForbidden {
	return &CommitOutputPortTransactionForbidden{}
}

/*
CommitOutputPortTransactionForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type CommitOutputPortTransactionForbidden struct {
}

// IsSuccess returns true when this commit output port transaction forbidden response has a 2xx status code
func (o *CommitOutputPortTransactionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this commit output port transaction forbidden response has a 3xx status code
func (o *CommitOutputPortTransactionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this commit output port transaction forbidden response has a 4xx status code
func (o *CommitOutputPortTransactionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this commit output port transaction forbidden response has a 5xx status code
func (o *CommitOutputPortTransactionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this commit output port transaction forbidden response a status code equal to that given
func (o *CommitOutputPortTransactionForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CommitOutputPortTransactionForbidden) Error() string {
	return fmt.Sprintf("[DELETE /data-transfer/output-ports/{portId}/transactions/{transactionId}][%d] commitOutputPortTransactionForbidden ", 403)
}

func (o *CommitOutputPortTransactionForbidden) String() string {
	return fmt.Sprintf("[DELETE /data-transfer/output-ports/{portId}/transactions/{transactionId}][%d] commitOutputPortTransactionForbidden ", 403)
}

func (o *CommitOutputPortTransactionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCommitOutputPortTransactionNotFound creates a CommitOutputPortTransactionNotFound with default headers values
func NewCommitOutputPortTransactionNotFound() *CommitOutputPortTransactionNotFound {
	return &CommitOutputPortTransactionNotFound{}
}

/*
CommitOutputPortTransactionNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type CommitOutputPortTransactionNotFound struct {
}

// IsSuccess returns true when this commit output port transaction not found response has a 2xx status code
func (o *CommitOutputPortTransactionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this commit output port transaction not found response has a 3xx status code
func (o *CommitOutputPortTransactionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this commit output port transaction not found response has a 4xx status code
func (o *CommitOutputPortTransactionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this commit output port transaction not found response has a 5xx status code
func (o *CommitOutputPortTransactionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this commit output port transaction not found response a status code equal to that given
func (o *CommitOutputPortTransactionNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CommitOutputPortTransactionNotFound) Error() string {
	return fmt.Sprintf("[DELETE /data-transfer/output-ports/{portId}/transactions/{transactionId}][%d] commitOutputPortTransactionNotFound ", 404)
}

func (o *CommitOutputPortTransactionNotFound) String() string {
	return fmt.Sprintf("[DELETE /data-transfer/output-ports/{portId}/transactions/{transactionId}][%d] commitOutputPortTransactionNotFound ", 404)
}

func (o *CommitOutputPortTransactionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCommitOutputPortTransactionConflict creates a CommitOutputPortTransactionConflict with default headers values
func NewCommitOutputPortTransactionConflict() *CommitOutputPortTransactionConflict {
	return &CommitOutputPortTransactionConflict{}
}

/*
CommitOutputPortTransactionConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type CommitOutputPortTransactionConflict struct {
}

// IsSuccess returns true when this commit output port transaction conflict response has a 2xx status code
func (o *CommitOutputPortTransactionConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this commit output port transaction conflict response has a 3xx status code
func (o *CommitOutputPortTransactionConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this commit output port transaction conflict response has a 4xx status code
func (o *CommitOutputPortTransactionConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this commit output port transaction conflict response has a 5xx status code
func (o *CommitOutputPortTransactionConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this commit output port transaction conflict response a status code equal to that given
func (o *CommitOutputPortTransactionConflict) IsCode(code int) bool {
	return code == 409
}

func (o *CommitOutputPortTransactionConflict) Error() string {
	return fmt.Sprintf("[DELETE /data-transfer/output-ports/{portId}/transactions/{transactionId}][%d] commitOutputPortTransactionConflict ", 409)
}

func (o *CommitOutputPortTransactionConflict) String() string {
	return fmt.Sprintf("[DELETE /data-transfer/output-ports/{portId}/transactions/{transactionId}][%d] commitOutputPortTransactionConflict ", 409)
}

func (o *CommitOutputPortTransactionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCommitOutputPortTransactionServiceUnavailable creates a CommitOutputPortTransactionServiceUnavailable with default headers values
func NewCommitOutputPortTransactionServiceUnavailable() *CommitOutputPortTransactionServiceUnavailable {
	return &CommitOutputPortTransactionServiceUnavailable{}
}

/*
CommitOutputPortTransactionServiceUnavailable describes a response with status code 503, with default header values.

NiFi instance is not ready for serving request, or temporarily overloaded. Retrying the same request later may be successful
*/
type CommitOutputPortTransactionServiceUnavailable struct {
}

// IsSuccess returns true when this commit output port transaction service unavailable response has a 2xx status code
func (o *CommitOutputPortTransactionServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this commit output port transaction service unavailable response has a 3xx status code
func (o *CommitOutputPortTransactionServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this commit output port transaction service unavailable response has a 4xx status code
func (o *CommitOutputPortTransactionServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this commit output port transaction service unavailable response has a 5xx status code
func (o *CommitOutputPortTransactionServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this commit output port transaction service unavailable response a status code equal to that given
func (o *CommitOutputPortTransactionServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *CommitOutputPortTransactionServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /data-transfer/output-ports/{portId}/transactions/{transactionId}][%d] commitOutputPortTransactionServiceUnavailable ", 503)
}

func (o *CommitOutputPortTransactionServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /data-transfer/output-ports/{portId}/transactions/{transactionId}][%d] commitOutputPortTransactionServiceUnavailable ", 503)
}

func (o *CommitOutputPortTransactionServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
