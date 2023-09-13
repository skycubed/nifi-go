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

// VersionedFlowSnapshotEntity versioned flow snapshot entity
//
// swagger:model VersionedFlowSnapshotEntity
type VersionedFlowSnapshotEntity struct {

	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged bool `json:"disconnectedNodeAcknowledged,omitempty"`

	// The Revision of the Process Group under Version Control
	ProcessGroupRevision *RevisionDTO `json:"processGroupRevision,omitempty"`

	// The ID of the Registry that this flow belongs to
	RegistryID string `json:"registryId,omitempty"`

	// If the Process Group to be updated has a child or descendant Process Group that is also under Version Control, this specifies whether or not the contents of that child/descendant Process Group should be updated.
	UpdateDescendantVersionedFlows bool `json:"updateDescendantVersionedFlows,omitempty"`

	// The versioned flow snapshot
	VersionedFlowSnapshot *VersionedFlowSnapshot `json:"versionedFlowSnapshot,omitempty"`
}

// Validate validates this versioned flow snapshot entity
func (m *VersionedFlowSnapshotEntity) Validate(formats strfmt.Registry) error {
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

func (m *VersionedFlowSnapshotEntity) validateProcessGroupRevision(formats strfmt.Registry) error {
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

func (m *VersionedFlowSnapshotEntity) validateVersionedFlowSnapshot(formats strfmt.Registry) error {
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

// ContextValidate validate this versioned flow snapshot entity based on the context it is used
func (m *VersionedFlowSnapshotEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *VersionedFlowSnapshotEntity) contextValidateProcessGroupRevision(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VersionedFlowSnapshotEntity) contextValidateVersionedFlowSnapshot(ctx context.Context, formats strfmt.Registry) error {

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
func (m *VersionedFlowSnapshotEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VersionedFlowSnapshotEntity) UnmarshalBinary(b []byte) error {
	var res VersionedFlowSnapshotEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
