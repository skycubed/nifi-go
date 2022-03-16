// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProcessGroupReplaceRequestDTO process group replace request d t o
//
// swagger:model ProcessGroupReplaceRequestDTO
type ProcessGroupReplaceRequestDTO struct {

	// Whether or not this request has completed
	Complete bool `json:"complete,omitempty"`

	// An explanation of why this request failed, or null if this request has not failed
	FailureReason string `json:"failureReason,omitempty"`

	// The last time this request was updated.
	LastUpdated string `json:"lastUpdated,omitempty"`

	// The percentage complete for the request, between 0 and 100
	PercentCompleted int32 `json:"percentCompleted,omitempty"`

	// The unique ID of the Process Group being updated
	ProcessGroupID string `json:"processGroupId,omitempty"`

	// The unique ID of this request.
	RequestID string `json:"requestId,omitempty"`

	// The state of the request
	State string `json:"state,omitempty"`

	// The URI for future requests to this drop request.
	URI string `json:"uri,omitempty"`
}

// Validate validates this process group replace request d t o
func (m *ProcessGroupReplaceRequestDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this process group replace request d t o based on context it is used
func (m *ProcessGroupReplaceRequestDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProcessGroupReplaceRequestDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProcessGroupReplaceRequestDTO) UnmarshalBinary(b []byte) error {
	var res ProcessGroupReplaceRequestDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}