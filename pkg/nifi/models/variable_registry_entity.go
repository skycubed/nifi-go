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

// VariableRegistryEntity variable registry entity
//
// swagger:model VariableRegistryEntity
type VariableRegistryEntity struct {

	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged bool `json:"disconnectedNodeAcknowledged,omitempty"`

	// The revision of the Process Group that the Variable Registry belongs to
	ProcessGroupRevision *RevisionDTO `json:"processGroupRevision,omitempty"`

	// The Variable Registry.
	VariableRegistry *VariableRegistryDTO `json:"variableRegistry,omitempty"`
}

// Validate validates this variable registry entity
func (m *VariableRegistryEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProcessGroupRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVariableRegistry(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VariableRegistryEntity) validateProcessGroupRevision(formats strfmt.Registry) error {
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

func (m *VariableRegistryEntity) validateVariableRegistry(formats strfmt.Registry) error {
	if swag.IsZero(m.VariableRegistry) { // not required
		return nil
	}

	if m.VariableRegistry != nil {
		if err := m.VariableRegistry.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("variableRegistry")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("variableRegistry")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this variable registry entity based on the context it is used
func (m *VariableRegistryEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProcessGroupRevision(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVariableRegistry(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VariableRegistryEntity) contextValidateProcessGroupRevision(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VariableRegistryEntity) contextValidateVariableRegistry(ctx context.Context, formats strfmt.Registry) error {

	if m.VariableRegistry != nil {

		if swag.IsZero(m.VariableRegistry) { // not required
			return nil
		}

		if err := m.VariableRegistry.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("variableRegistry")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("variableRegistry")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VariableRegistryEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VariableRegistryEntity) UnmarshalBinary(b []byte) error {
	var res VariableRegistryEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
