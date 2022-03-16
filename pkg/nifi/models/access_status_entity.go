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

// AccessStatusEntity access status entity
//
// swagger:model AccessStatusEntity
type AccessStatusEntity struct {

	// access status
	AccessStatus *AccessStatusDTO `json:"accessStatus,omitempty"`
}

// Validate validates this access status entity
func (m *AccessStatusEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccessStatusEntity) validateAccessStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessStatus) { // not required
		return nil
	}

	if m.AccessStatus != nil {
		if err := m.AccessStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accessStatus")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this access status entity based on the context it is used
func (m *AccessStatusEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccessStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccessStatusEntity) contextValidateAccessStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.AccessStatus != nil {
		if err := m.AccessStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accessStatus")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccessStatusEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccessStatusEntity) UnmarshalBinary(b []byte) error {
	var res AccessStatusEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
