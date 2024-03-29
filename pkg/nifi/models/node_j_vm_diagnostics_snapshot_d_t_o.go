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

// NodeJVMDiagnosticsSnapshotDTO node j VM diagnostics snapshot d t o
//
// swagger:model NodeJVMDiagnosticsSnapshotDTO
type NodeJVMDiagnosticsSnapshotDTO struct {

	// The API address of the node
	Address string `json:"address,omitempty"`

	// The API port used to communicate with the node
	APIPort int32 `json:"apiPort,omitempty"`

	// The unique ID that identifies the node
	NodeID string `json:"nodeId,omitempty"`

	// The JVM Diagnostics Snapshot
	Snapshot *JVMDiagnosticsSnapshotDTO `json:"snapshot,omitempty"`
}

// Validate validates this node j VM diagnostics snapshot d t o
func (m *NodeJVMDiagnosticsSnapshotDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSnapshot(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeJVMDiagnosticsSnapshotDTO) validateSnapshot(formats strfmt.Registry) error {
	if swag.IsZero(m.Snapshot) { // not required
		return nil
	}

	if m.Snapshot != nil {
		if err := m.Snapshot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshot")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("snapshot")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this node j VM diagnostics snapshot d t o based on the context it is used
func (m *NodeJVMDiagnosticsSnapshotDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSnapshot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeJVMDiagnosticsSnapshotDTO) contextValidateSnapshot(ctx context.Context, formats strfmt.Registry) error {

	if m.Snapshot != nil {

		if swag.IsZero(m.Snapshot) { // not required
			return nil
		}

		if err := m.Snapshot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshot")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("snapshot")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NodeJVMDiagnosticsSnapshotDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeJVMDiagnosticsSnapshotDTO) UnmarshalBinary(b []byte) error {
	var res NodeJVMDiagnosticsSnapshotDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
