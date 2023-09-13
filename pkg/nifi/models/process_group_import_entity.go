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

// ProcessGroupImportEntity process group import entity
//
// swagger:model ProcessGroupImportEntity
type ProcessGroupImportEntity struct {

	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged bool `json:"disconnectedNodeAcknowledged,omitempty"`

	// The Revision for the Process Group
	ProcessGroupRevision *RevisionDTO `json:"processGroupRevision,omitempty"`

	// The Versioned Flow Snapshot to import
	VersionedFlowSnapshot *VersionedFlowSnapshot `json:"versionedFlowSnapshot,omitempty"`
}

// Validate validates this process group import entity
func (m *ProcessGroupImportEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProcessGroupRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionedFlowSnapshot(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProcessGroupImportEntity) validateProcessGroupRevision(formats strfmt.Registry) error {
	if swag.IsZero(m.ProcessGroupRevision) { // not required
		return nil
	}

	if m.ProcessGroupRevision != nil {
		if err := m.ProcessGroupRevision.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("processGroupRevision")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("processGroupRevision")
			}
			return err
		}
	}

	return nil
}

func (m *ProcessGroupImportEntity) validateVersionedFlowSnapshot(formats strfmt.Registry) error {
	if swag.IsZero(m.VersionedFlowSnapshot) { // not required
		return nil
	}

	if m.VersionedFlowSnapshot != nil {
		if err := m.VersionedFlowSnapshot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("versionedFlowSnapshot")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("versionedFlowSnapshot")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this process group import entity based on the context it is used
func (m *ProcessGroupImportEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProcessGroupRevision(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVersionedFlowSnapshot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProcessGroupImportEntity) contextValidateProcessGroupRevision(ctx context.Context, formats strfmt.Registry) error {

	if m.ProcessGroupRevision != nil {

		if swag.IsZero(m.ProcessGroupRevision) { // not required
			return nil
		}

		if err := m.ProcessGroupRevision.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("processGroupRevision")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("processGroupRevision")
			}
			return err
		}
	}

	return nil
}

func (m *ProcessGroupImportEntity) contextValidateVersionedFlowSnapshot(ctx context.Context, formats strfmt.Registry) error {

	if m.VersionedFlowSnapshot != nil {

		if swag.IsZero(m.VersionedFlowSnapshot) { // not required
			return nil
		}

		if err := m.VersionedFlowSnapshot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("versionedFlowSnapshot")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("versionedFlowSnapshot")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProcessGroupImportEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProcessGroupImportEntity) UnmarshalBinary(b []byte) error {
	var res ProcessGroupImportEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
