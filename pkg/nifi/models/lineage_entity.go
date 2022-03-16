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

// LineageEntity lineage entity
//
// swagger:model LineageEntity
type LineageEntity struct {

	// lineage
	Lineage *LineageDTO `json:"lineage,omitempty"`
}

// Validate validates this lineage entity
func (m *LineageEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLineage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LineageEntity) validateLineage(formats strfmt.Registry) error {
	if swag.IsZero(m.Lineage) { // not required
		return nil
	}

	if m.Lineage != nil {
		if err := m.Lineage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lineage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lineage")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this lineage entity based on the context it is used
func (m *LineageEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLineage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LineageEntity) contextValidateLineage(ctx context.Context, formats strfmt.Registry) error {

	if m.Lineage != nil {
		if err := m.Lineage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lineage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lineage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LineageEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LineageEntity) UnmarshalBinary(b []byte) error {
	var res LineageEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
