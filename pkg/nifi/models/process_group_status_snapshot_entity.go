// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProcessGroupStatusSnapshotEntity process group status snapshot entity
//
// swagger:model ProcessGroupStatusSnapshotEntity
type ProcessGroupStatusSnapshotEntity struct {

	// Indicates whether the user can read a given resource.
	CanRead bool `json:"canRead,omitempty"`

	// The id of the process group.
	ID string `json:"id,omitempty"`

	// process group status snapshot
	ProcessGroupStatusSnapshot *ProcessGroupStatusSnapshotDTO `json:"processGroupStatusSnapshot,omitempty"`
}

// Validate validates this process group status snapshot entity
func (m *ProcessGroupStatusSnapshotEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProcessGroupStatusSnapshot(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProcessGroupStatusSnapshotEntity) validateProcessGroupStatusSnapshot(formats strfmt.Registry) error {
	if swag.IsZero(m.ProcessGroupStatusSnapshot) { // not required
		return nil
	}

	if m.ProcessGroupStatusSnapshot != nil {
		if err := m.ProcessGroupStatusSnapshot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("processGroupStatusSnapshot")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("processGroupStatusSnapshot")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this process group status snapshot entity based on the context it is used
func (m *ProcessGroupStatusSnapshotEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProcessGroupStatusSnapshot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProcessGroupStatusSnapshotEntity) contextValidateProcessGroupStatusSnapshot(ctx context.Context, formats strfmt.Registry) error {

	if m.ProcessGroupStatusSnapshot != nil {

		if swag.IsZero(m.ProcessGroupStatusSnapshot) { // not required
			return nil
		}

		if err := m.ProcessGroupStatusSnapshot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("processGroupStatusSnapshot")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("processGroupStatusSnapshot")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProcessGroupStatusSnapshotEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProcessGroupStatusSnapshotEntity) UnmarshalBinary(b []byte) error {
	var res ProcessGroupStatusSnapshotEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
