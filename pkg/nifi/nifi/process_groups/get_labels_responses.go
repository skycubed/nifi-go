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

// GetLabelsReader is a Reader for the GetLabels structure.
type GetLabelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLabelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLabelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLabelsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLabelsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLabelsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLabelsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetLabelsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLabelsOK creates a GetLabelsOK with default headers values
func NewGetLabelsOK() *GetLabelsOK {
	return &GetLabelsOK{}
}

/*
GetLabelsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetLabelsOK struct {
	Payload *models.LabelsEntity
}

// IsSuccess returns true when this get labels o k response has a 2xx status code
func (o *GetLabelsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get labels o k response has a 3xx status code
func (o *GetLabelsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get labels o k response has a 4xx status code
func (o *GetLabelsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get labels o k response has a 5xx status code
func (o *GetLabelsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get labels o k response a status code equal to that given
func (o *GetLabelsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetLabelsOK) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/labels][%d] getLabelsOK  %+v", 200, o.Payload)
}

func (o *GetLabelsOK) String() string {
	return fmt.Sprintf("[GET /process-groups/{id}/labels][%d] getLabelsOK  %+v", 200, o.Payload)
}

func (o *GetLabelsOK) GetPayload() *models.LabelsEntity {
	return o.Payload
}

func (o *GetLabelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LabelsEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLabelsBadRequest creates a GetLabelsBadRequest with default headers values
func NewGetLabelsBadRequest() *GetLabelsBadRequest {
	return &GetLabelsBadRequest{}
}

/*
GetLabelsBadRequest describes a response with status code 400, with default header values.

NiFi was unable to complete the request because it was invalid. The request should not be retried without modification.
*/
type GetLabelsBadRequest struct {
}

// IsSuccess returns true when this get labels bad request response has a 2xx status code
func (o *GetLabelsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get labels bad request response has a 3xx status code
func (o *GetLabelsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get labels bad request response has a 4xx status code
func (o *GetLabelsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get labels bad request response has a 5xx status code
func (o *GetLabelsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get labels bad request response a status code equal to that given
func (o *GetLabelsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetLabelsBadRequest) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/labels][%d] getLabelsBadRequest ", 400)
}

func (o *GetLabelsBadRequest) String() string {
	return fmt.Sprintf("[GET /process-groups/{id}/labels][%d] getLabelsBadRequest ", 400)
}

func (o *GetLabelsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLabelsUnauthorized creates a GetLabelsUnauthorized with default headers values
func NewGetLabelsUnauthorized() *GetLabelsUnauthorized {
	return &GetLabelsUnauthorized{}
}

/*
GetLabelsUnauthorized describes a response with status code 401, with default header values.

Client could not be authenticated.
*/
type GetLabelsUnauthorized struct {
}

// IsSuccess returns true when this get labels unauthorized response has a 2xx status code
func (o *GetLabelsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get labels unauthorized response has a 3xx status code
func (o *GetLabelsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get labels unauthorized response has a 4xx status code
func (o *GetLabelsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get labels unauthorized response has a 5xx status code
func (o *GetLabelsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get labels unauthorized response a status code equal to that given
func (o *GetLabelsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetLabelsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/labels][%d] getLabelsUnauthorized ", 401)
}

func (o *GetLabelsUnauthorized) String() string {
	return fmt.Sprintf("[GET /process-groups/{id}/labels][%d] getLabelsUnauthorized ", 401)
}

func (o *GetLabelsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLabelsForbidden creates a GetLabelsForbidden with default headers values
func NewGetLabelsForbidden() *GetLabelsForbidden {
	return &GetLabelsForbidden{}
}

/*
GetLabelsForbidden describes a response with status code 403, with default header values.

Client is not authorized to make this request.
*/
type GetLabelsForbidden struct {
}

// IsSuccess returns true when this get labels forbidden response has a 2xx status code
func (o *GetLabelsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get labels forbidden response has a 3xx status code
func (o *GetLabelsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get labels forbidden response has a 4xx status code
func (o *GetLabelsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get labels forbidden response has a 5xx status code
func (o *GetLabelsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get labels forbidden response a status code equal to that given
func (o *GetLabelsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetLabelsForbidden) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/labels][%d] getLabelsForbidden ", 403)
}

func (o *GetLabelsForbidden) String() string {
	return fmt.Sprintf("[GET /process-groups/{id}/labels][%d] getLabelsForbidden ", 403)
}

func (o *GetLabelsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLabelsNotFound creates a GetLabelsNotFound with default headers values
func NewGetLabelsNotFound() *GetLabelsNotFound {
	return &GetLabelsNotFound{}
}

/*
GetLabelsNotFound describes a response with status code 404, with default header values.

The specified resource could not be found.
*/
type GetLabelsNotFound struct {
}

// IsSuccess returns true when this get labels not found response has a 2xx status code
func (o *GetLabelsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get labels not found response has a 3xx status code
func (o *GetLabelsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get labels not found response has a 4xx status code
func (o *GetLabelsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get labels not found response has a 5xx status code
func (o *GetLabelsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get labels not found response a status code equal to that given
func (o *GetLabelsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetLabelsNotFound) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/labels][%d] getLabelsNotFound ", 404)
}

func (o *GetLabelsNotFound) String() string {
	return fmt.Sprintf("[GET /process-groups/{id}/labels][%d] getLabelsNotFound ", 404)
}

func (o *GetLabelsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLabelsConflict creates a GetLabelsConflict with default headers values
func NewGetLabelsConflict() *GetLabelsConflict {
	return &GetLabelsConflict{}
}

/*
GetLabelsConflict describes a response with status code 409, with default header values.

The request was valid but NiFi was not in the appropriate state to process it. Retrying the same request later may be successful.
*/
type GetLabelsConflict struct {
}

// IsSuccess returns true when this get labels conflict response has a 2xx status code
func (o *GetLabelsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get labels conflict response has a 3xx status code
func (o *GetLabelsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get labels conflict response has a 4xx status code
func (o *GetLabelsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get labels conflict response has a 5xx status code
func (o *GetLabelsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get labels conflict response a status code equal to that given
func (o *GetLabelsConflict) IsCode(code int) bool {
	return code == 409
}

func (o *GetLabelsConflict) Error() string {
	return fmt.Sprintf("[GET /process-groups/{id}/labels][%d] getLabelsConflict ", 409)
}

func (o *GetLabelsConflict) String() string {
	return fmt.Sprintf("[GET /process-groups/{id}/labels][%d] getLabelsConflict ", 409)
}

func (o *GetLabelsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
