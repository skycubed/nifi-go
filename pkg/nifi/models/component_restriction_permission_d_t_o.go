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

// ComponentRestrictionPermissionDTO component restriction permission d t o
//
// swagger:model ComponentRestrictionPermissionDTO
type ComponentRestrictionPermissionDTO struct {

	// The permissions for this component restriction. Note: the read permission are not used and will always be false.
	Permissions *PermissionsDTO `json:"permissions,omitempty"`

	// The required permission necessary for this restriction.
	RequiredPermission *RequiredPermissionDTO `json:"requiredPermission,omitempty"`
}

// Validate validates this component restriction permission d t o
func (m *ComponentRestrictionPermissionDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequiredPermission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComponentRestrictionPermissionDTO) validatePermissions(formats strfmt.Registry) error {
	if swag.IsZero(m.Permissions) { // not required
		return nil
	}

	if m.Permissions != nil {
		if err := m.Permissions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("permissions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("permissions")
			}
			return err
		}
	}

	return nil
}

func (m *ComponentRestrictionPermissionDTO) validateRequiredPermission(formats strfmt.Registry) error {
	if swag.IsZero(m.RequiredPermission) { // not required
		return nil
	}

	if m.RequiredPermission != nil {
		if err := m.RequiredPermission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requiredPermission")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("requiredPermission")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this component restriction permission d t o based on the context it is used
func (m *ComponentRestrictionPermissionDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePermissions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRequiredPermission(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComponentRestrictionPermissionDTO) contextValidatePermissions(ctx context.Context, formats strfmt.Registry) error {

	if m.Permissions != nil {

		if swag.IsZero(m.Permissions) { // not required
			return nil
		}

		if err := m.Permissions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("permissions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("permissions")
			}
			return err
		}
	}

	return nil
}

func (m *ComponentRestrictionPermissionDTO) contextValidateRequiredPermission(ctx context.Context, formats strfmt.Registry) error {

	if m.RequiredPermission != nil {

		if swag.IsZero(m.RequiredPermission) { // not required
			return nil
		}

		if err := m.RequiredPermission.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requiredPermission")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("requiredPermission")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ComponentRestrictionPermissionDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComponentRestrictionPermissionDTO) UnmarshalBinary(b []byte) error {
	var res ComponentRestrictionPermissionDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
