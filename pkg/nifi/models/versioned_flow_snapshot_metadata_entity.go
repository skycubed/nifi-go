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

// VersionedFlowSnapshotMetadataEntity versioned flow snapshot metadata entity
//
// swagger:model VersionedFlowSnapshotMetadataEntity
type VersionedFlowSnapshotMetadataEntity struct {

	// The ID of the Registry that this flow belongs to
	RegistryID string `json:"registryId,omitempty"`

	// The collection of versioned flow snapshot metadata
	VersionedFlowSnapshotMetadata *VersionedFlowSnapshotMetadata `json:"versionedFlowSnapshotMetadata,omitempty"`
}

// Validate validates this versioned flow snapshot metadata entity
func (m *VersionedFlowSnapshotMetadataEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVersionedFlowSnapshotMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VersionedFlowSnapshotMetadataEntity) validateVersionedFlowSnapshotMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.VersionedFlowSnapshotMetadata) { // not required
		return nil
	}

	if m.VersionedFlowSnapshotMetadata != nil {
		if err := m.VersionedFlowSnapshotMetadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("versionedFlowSnapshotMetadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("versionedFlowSnapshotMetadata")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this versioned flow snapshot metadata entity based on the context it is used
func (m *VersionedFlowSnapshotMetadataEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVersionedFlowSnapshotMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VersionedFlowSnapshotMetadataEntity) contextValidateVersionedFlowSnapshotMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.VersionedFlowSnapshotMetadata != nil {

		if swag.IsZero(m.VersionedFlowSnapshotMetadata) { // not required
			return nil
		}

		if err := m.VersionedFlowSnapshotMetadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("versionedFlowSnapshotMetadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("versionedFlowSnapshotMetadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VersionedFlowSnapshotMetadataEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VersionedFlowSnapshotMetadataEntity) UnmarshalBinary(b []byte) error {
	var res VersionedFlowSnapshotMetadataEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
