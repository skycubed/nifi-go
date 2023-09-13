// Code generated by go-swagger; DO NOT EDIT.

package process_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// UploadProcessGroupReader is a Reader for the UploadProcessGroup structure.
type UploadProcessGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadProcessGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUploadProcessGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUploadProcessGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUploadProcessGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUploadProcessGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUploadProcessGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUploadProcessGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /process-groups/{id}/process-groups/upload] uploadProcessGroup", response, response.Code())
	}
}

// NewUploadProcessGroupOK creates a UploadProcessGroupOK with default headers values
func NewUploadProcessGroupOK() *UploadProcessGroupOK {
	return &UploadProcessGroupOK{}
}

/*
UploadProcessGroupOK describes a response with status code 200, with default header values.

successful operation
*/
type UploadProcessGroupOK struct {
	Payload *models.ProcessGroupEntity
}

// IsSuccess returns true when this upload process group o k response has a 2xx status code
func (o *UploadProcessGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this upload process group o k response has a 3xx status code
func (o *UploadProcessGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload process group o k response has a 4xx status code
func (o *UploadProcessGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this upload process group o k response has a 5xx status code
func (o *UploadProcessGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this upload process group o k response a status code equal to that given
func (o *UploadProcessGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the upload process group o k response
func (o *UploadProcessGroupOK) Code() int {
	return 200
}

func (o *UploadProcessGroupOK) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/process-groups/upload][%d] uploadProcessGroupOK  %+v", 200, o.Payload)
}

func (o *UploadProcessGroupOK) String() string {
	return fmt.Sprintf("[POST /process-groups/{id}/process-groups/upload][%d] uploadProcessGroupOK  %+v", 200, o.Payload)
}

func (o *UploadProcessGroupOK) GetPayload() *models.ProcessGroupEntity {
	return o.Payload
}

func (o *UploadProcessGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProcessGroupEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadProcessGroupBadRequest creates a UploadProcessGroupBadRequest with default headers values
func NewUploadProcessGroupBadRequest() *UploadProcessGroupBadRequest {
	return &UploadProcessGroupBadRequest{}
}

/*
UploadProcessGroupBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type UploadProcessGroupBadRequest struct {
}

// IsSuccess returns true when this upload process group bad request response has a 2xx status code
func (o *UploadProcessGroupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload process group bad request response has a 3xx status code
func (o *UploadProcessGroupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload process group bad request response has a 4xx status code
func (o *UploadProcessGroupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this upload process group bad request response has a 5xx status code
func (o *UploadProcessGroupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this upload process group bad request response a status code equal to that given
func (o *UploadProcessGroupBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the upload process group bad request response
func (o *UploadProcessGroupBadRequest) Code() int {
	return 400
}

func (o *UploadProcessGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/process-groups/upload][%d] uploadProcessGroupBadRequest ", 400)
}

func (o *UploadProcessGroupBadRequest) String() string {
	return fmt.Sprintf("[POST /process-groups/{id}/process-groups/upload][%d] uploadProcessGroupBadRequest ", 400)
}

func (o *UploadProcessGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUploadProcessGroupUnauthorized creates a UploadProcessGroupUnauthorized with default headers values
func NewUploadProcessGroupUnauthorized() *UploadProcessGroupUnauthorized {
	return &UploadProcessGroupUnauthorized{}
}

/*
UploadProcessGroupUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type UploadProcessGroupUnauthorized struct {
}

// IsSuccess returns true when this upload process group unauthorized response has a 2xx status code
func (o *UploadProcessGroupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload process group unauthorized response has a 3xx status code
func (o *UploadProcessGroupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload process group unauthorized response has a 4xx status code
func (o *UploadProcessGroupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this upload process group unauthorized response has a 5xx status code
func (o *UploadProcessGroupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this upload process group unauthorized response a status code equal to that given
func (o *UploadProcessGroupUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the upload process group unauthorized response
func (o *UploadProcessGroupUnauthorized) Code() int {
	return 401
}

func (o *UploadProcessGroupUnauthorized) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/process-groups/upload][%d] uploadProcessGroupUnauthorized ", 401)
}

func (o *UploadProcessGroupUnauthorized) String() string {
	return fmt.Sprintf("[POST /process-groups/{id}/process-groups/upload][%d] uploadProcessGroupUnauthorized ", 401)
}

func (o *UploadProcessGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUploadProcessGroupForbidden creates a UploadProcessGroupForbidden with default headers values
func NewUploadProcessGroupForbidden() *UploadProcessGroupForbidden {
	return &UploadProcessGroupForbidden{}
}

/*
UploadProcessGroupForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type UploadProcessGroupForbidden struct {
}

// IsSuccess returns true when this upload process group forbidden response has a 2xx status code
func (o *UploadProcessGroupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload process group forbidden response has a 3xx status code
func (o *UploadProcessGroupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload process group forbidden response has a 4xx status code
func (o *UploadProcessGroupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this upload process group forbidden response has a 5xx status code
func (o *UploadProcessGroupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this upload process group forbidden response a status code equal to that given
func (o *UploadProcessGroupForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the upload process group forbidden response
func (o *UploadProcessGroupForbidden) Code() int {
	return 403
}

func (o *UploadProcessGroupForbidden) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/process-groups/upload][%d] uploadProcessGroupForbidden ", 403)
}

func (o *UploadProcessGroupForbidden) String() string {
	return fmt.Sprintf("[POST /process-groups/{id}/process-groups/upload][%d] uploadProcessGroupForbidden ", 403)
}

func (o *UploadProcessGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUploadProcessGroupNotFound creates a UploadProcessGroupNotFound with default headers values
func NewUploadProcessGroupNotFound() *UploadProcessGroupNotFound {
	return &UploadProcessGroupNotFound{}
}

/*
UploadProcessGroupNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type UploadProcessGroupNotFound struct {
}

// IsSuccess returns true when this upload process group not found response has a 2xx status code
func (o *UploadProcessGroupNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload process group not found response has a 3xx status code
func (o *UploadProcessGroupNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload process group not found response has a 4xx status code
func (o *UploadProcessGroupNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this upload process group not found response has a 5xx status code
func (o *UploadProcessGroupNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this upload process group not found response a status code equal to that given
func (o *UploadProcessGroupNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the upload process group not found response
func (o *UploadProcessGroupNotFound) Code() int {
	return 404
}

func (o *UploadProcessGroupNotFound) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/process-groups/upload][%d] uploadProcessGroupNotFound ", 404)
}

func (o *UploadProcessGroupNotFound) String() string {
	return fmt.Sprintf("[POST /process-groups/{id}/process-groups/upload][%d] uploadProcessGroupNotFound ", 404)
}

func (o *UploadProcessGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUploadProcessGroupConflict creates a UploadProcessGroupConflict with default headers values
func NewUploadProcessGroupConflict() *UploadProcessGroupConflict {
	return &UploadProcessGroupConflict{}
}

/*
UploadProcessGroupConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type UploadProcessGroupConflict struct {
}

// IsSuccess returns true when this upload process group conflict response has a 2xx status code
func (o *UploadProcessGroupConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload process group conflict response has a 3xx status code
func (o *UploadProcessGroupConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload process group conflict response has a 4xx status code
func (o *UploadProcessGroupConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this upload process group conflict response has a 5xx status code
func (o *UploadProcessGroupConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this upload process group conflict response a status code equal to that given
func (o *UploadProcessGroupConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the upload process group conflict response
func (o *UploadProcessGroupConflict) Code() int {
	return 409
}

func (o *UploadProcessGroupConflict) Error() string {
	return fmt.Sprintf("[POST /process-groups/{id}/process-groups/upload][%d] uploadProcessGroupConflict ", 409)
}

func (o *UploadProcessGroupConflict) String() string {
	return fmt.Sprintf("[POST /process-groups/{id}/process-groups/upload][%d] uploadProcessGroupConflict ", 409)
}

func (o *UploadProcessGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
