// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ReplayLastEventSnapshotDTO replay last event snapshot d t o
//
// swagger:model ReplayLastEventSnapshotDTO
type ReplayLastEventSnapshotDTO struct {

	// Whether or not an event was available. This may not be populated if there was a failure.
	EventAvailable bool `json:"eventAvailable,omitempty"`

	// The IDs of the events that were successfully replayed
	EventsReplayed []int64 `json:"eventsReplayed"`

	// If unable to replay an event, specifies why the event could not be replayed
	FailureExplanation string `json:"failureExplanation,omitempty"`
}

// Validate validates this replay last event snapshot d t o
func (m *ReplayLastEventSnapshotDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this replay last event snapshot d t o based on context it is used
func (m *ReplayLastEventSnapshotDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReplayLastEventSnapshotDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReplayLastEventSnapshotDTO) UnmarshalBinary(b []byte) error {
	var res ReplayLastEventSnapshotDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
