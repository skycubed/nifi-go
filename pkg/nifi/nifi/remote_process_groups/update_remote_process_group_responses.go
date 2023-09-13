// Code generated by go-swagger; DO NOT EDIT.

package remote_process_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skycubed/nifi-go/pkg/nifi/models"
)

// UpdateRemoteProcessGroupReader is a Reader for the UpdateRemoteProcessGroup structure.
type UpdateRemoteProcessGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRemoteProcessGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRemoteProcessGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateRemoteProcessGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateRemoteProcessGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRemoteProcessGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateRemoteProcessGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateRemoteProcessGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /remote-process-groups/{id}] updateRemoteProcessGroup", response, response.Code())
	}
}

// NewUpdateRemoteProcessGroupOK creates a UpdateRemoteProcessGroupOK with default headers values
func NewUpdateRemoteProcessGroupOK() *UpdateRemoteProcessGroupOK {
	return &UpdateRemoteProcessGroupOK{}
}

/*
UpdateRemoteProcessGroupOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateRemoteProcessGroupOK struct {
	Payload *models.RemoteProcessGroupEntity
}

// IsSuccess returns true when this update remote process group o k response has a 2xx status code
func (o *UpdateRemoteProcessGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update remote process group o k response has a 3xx status code
func (o *UpdateRemoteProcessGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update remote process group o k response has a 4xx status code
func (o *UpdateRemoteProcessGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update remote process group o k response has a 5xx status code
func (o *UpdateRemoteProcessGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update remote process group o k response a status code equal to that given
func (o *UpdateRemoteProcessGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update remote process group o k response
func (o *UpdateRemoteProcessGroupOK) Code() int {
	return 200
}

func (o *UpdateRemoteProcessGroupOK) Error() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}][%d] updateRemoteProcessGroupOK  %+v", 200, o.Payload)
}

func (o *UpdateRemoteProcessGroupOK) String() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}][%d] updateRemoteProcessGroupOK  %+v", 200, o.Payload)
}

func (o *UpdateRemoteProcessGroupOK) GetPayload() *models.RemoteProcessGroupEntity {
	return o.Payload
}

func (o *UpdateRemoteProcessGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RemoteProcessGroupEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRemoteProcessGroupBadRequest creates a UpdateRemoteProcessGroupBadRequest with default headers values
func NewUpdateRemoteProcessGroupBadRequest() *UpdateRemoteProcessGroupBadRequest {
	return &UpdateRemoteProcessGroupBadRequest{}
}

/*
UpdateRemoteProcessGroupBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type UpdateRemoteProcessGroupBadRequest struct {
}

// IsSuccess returns true when this update remote process group bad request response has a 2xx status code
func (o *UpdateRemoteProcessGroupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update remote process group bad request response has a 3xx status code
func (o *UpdateRemoteProcessGroupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update remote process group bad request response has a 4xx status code
func (o *UpdateRemoteProcessGroupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update remote process group bad request response has a 5xx status code
func (o *UpdateRemoteProcessGroupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update remote process group bad request response a status code equal to that given
func (o *UpdateRemoteProcessGroupBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update remote process group bad request response
func (o *UpdateRemoteProcessGroupBadRequest) Code() int {
	return 400
}

func (o *UpdateRemoteProcessGroupBadRequest) Error() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}][%d] updateRemoteProcessGroupBadRequest ", 400)
}

func (o *UpdateRemoteProcessGroupBadRequest) String() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}][%d] updateRemoteProcessGroupBadRequest ", 400)
}

func (o *UpdateRemoteProcessGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRemoteProcessGroupUnauthorized creates a UpdateRemoteProcessGroupUnauthorized with default headers values
func NewUpdateRemoteProcessGroupUnauthorized() *UpdateRemoteProcessGroupUnauthorized {
	return &UpdateRemoteProcessGroupUnauthorized{}
}

/*
UpdateRemoteProcessGroupUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type UpdateRemoteProcessGroupUnauthorized struct {
}

// IsSuccess returns true when this update remote process group unauthorized response has a 2xx status code
func (o *UpdateRemoteProcessGroupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update remote process group unauthorized response has a 3xx status code
func (o *UpdateRemoteProcessGroupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update remote process group unauthorized response has a 4xx status code
func (o *UpdateRemoteProcessGroupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update remote process group unauthorized response has a 5xx status code
func (o *UpdateRemoteProcessGroupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update remote process group unauthorized response a status code equal to that given
func (o *UpdateRemoteProcessGroupUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update remote process group unauthorized response
func (o *UpdateRemoteProcessGroupUnauthorized) Code() int {
	return 401
}

func (o *UpdateRemoteProcessGroupUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}][%d] updateRemoteProcessGroupUnauthorized ", 401)
}

func (o *UpdateRemoteProcessGroupUnauthorized) String() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}][%d] updateRemoteProcessGroupUnauthorized ", 401)
}

func (o *UpdateRemoteProcessGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRemoteProcessGroupForbidden creates a UpdateRemoteProcessGroupForbidden with default headers values
func NewUpdateRemoteProcessGroupForbidden() *UpdateRemoteProcessGroupForbidden {
	return &UpdateRemoteProcessGroupForbidden{}
}

/*
UpdateRemoteProcessGroupForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type UpdateRemoteProcessGroupForbidden struct {
}

// IsSuccess returns true when this update remote process group forbidden response has a 2xx status code
func (o *UpdateRemoteProcessGroupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update remote process group forbidden response has a 3xx status code
func (o *UpdateRemoteProcessGroupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update remote process group forbidden response has a 4xx status code
func (o *UpdateRemoteProcessGroupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update remote process group forbidden response has a 5xx status code
func (o *UpdateRemoteProcessGroupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update remote process group forbidden response a status code equal to that given
func (o *UpdateRemoteProcessGroupForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update remote process group forbidden response
func (o *UpdateRemoteProcessGroupForbidden) Code() int {
	return 403
}

func (o *UpdateRemoteProcessGroupForbidden) Error() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}][%d] updateRemoteProcessGroupForbidden ", 403)
}

func (o *UpdateRemoteProcessGroupForbidden) String() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}][%d] updateRemoteProcessGroupForbidden ", 403)
}

func (o *UpdateRemoteProcessGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRemoteProcessGroupNotFound creates a UpdateRemoteProcessGroupNotFound with default headers values
func NewUpdateRemoteProcessGroupNotFound() *UpdateRemoteProcessGroupNotFound {
	return &UpdateRemoteProcessGroupNotFound{}
}

/*
UpdateRemoteProcessGroupNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type UpdateRemoteProcessGroupNotFound struct {
}

// IsSuccess returns true when this update remote process group not found response has a 2xx status code
func (o *UpdateRemoteProcessGroupNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update remote process group not found response has a 3xx status code
func (o *UpdateRemoteProcessGroupNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update remote process group not found response has a 4xx status code
func (o *UpdateRemoteProcessGroupNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update remote process group not found response has a 5xx status code
func (o *UpdateRemoteProcessGroupNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update remote process group not found response a status code equal to that given
func (o *UpdateRemoteProcessGroupNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update remote process group not found response
func (o *UpdateRemoteProcessGroupNotFound) Code() int {
	return 404
}

func (o *UpdateRemoteProcessGroupNotFound) Error() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}][%d] updateRemoteProcessGroupNotFound ", 404)
}

func (o *UpdateRemoteProcessGroupNotFound) String() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}][%d] updateRemoteProcessGroupNotFound ", 404)
}

func (o *UpdateRemoteProcessGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRemoteProcessGroupConflict creates a UpdateRemoteProcessGroupConflict with default headers values
func NewUpdateRemoteProcessGroupConflict() *UpdateRemoteProcessGroupConflict {
	return &UpdateRemoteProcessGroupConflict{}
}

/*
UpdateRemoteProcessGroupConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type UpdateRemoteProcessGroupConflict struct {
}

// IsSuccess returns true when this update remote process group conflict response has a 2xx status code
func (o *UpdateRemoteProcessGroupConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update remote process group conflict response has a 3xx status code
func (o *UpdateRemoteProcessGroupConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update remote process group conflict response has a 4xx status code
func (o *UpdateRemoteProcessGroupConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update remote process group conflict response has a 5xx status code
func (o *UpdateRemoteProcessGroupConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update remote process group conflict response a status code equal to that given
func (o *UpdateRemoteProcessGroupConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the update remote process group conflict response
func (o *UpdateRemoteProcessGroupConflict) Code() int {
	return 409
}

func (o *UpdateRemoteProcessGroupConflict) Error() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}][%d] updateRemoteProcessGroupConflict ", 409)
}

func (o *UpdateRemoteProcessGroupConflict) String() string {
	return fmt.Sprintf("[PUT /remote-process-groups/{id}][%d] updateRemoteProcessGroupConflict ", 409)
}

func (o *UpdateRemoteProcessGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
