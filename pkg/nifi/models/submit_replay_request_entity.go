// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SubmitReplayRequestEntity submit replay request entity
//
// swagger:model SubmitReplayRequestEntity
type SubmitReplayRequestEntity struct {

	// The identifier of the node where to submit the replay request.
	ClusterNodeID string `json:"clusterNodeId,omitempty"`

	// The event identifier
	EventID int64 `json:"eventId,omitempty"`
}

// Validate validates this submit replay request entity
func (m *SubmitReplayRequestEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this submit replay request entity based on context it is used
func (m *SubmitReplayRequestEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SubmitReplayRequestEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubmitReplayRequestEntity) UnmarshalBinary(b []byte) error {
	var res SubmitReplayRequestEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
