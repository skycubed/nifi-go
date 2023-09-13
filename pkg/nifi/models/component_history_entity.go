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

// ComponentHistoryEntity component history entity
//
// swagger:model ComponentHistoryEntity
type ComponentHistoryEntity struct {

	// component history
	ComponentHistory *ComponentHistoryDTO `json:"componentHistory,omitempty"`
}

// Validate validates this component history entity
func (m *ComponentHistoryEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComponentHistory(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComponentHistoryEntity) validateComponentHistory(formats strfmt.Registry) error {
	if swag.IsZero(m.ComponentHistory) { // not required
		return nil
	}

	if m.ComponentHistory != nil {
		if err := m.ComponentHistory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("componentHistory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("componentHistory")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this component history entity based on the context it is used
func (m *ComponentHistoryEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComponentHistory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComponentHistoryEntity) contextValidateComponentHistory(ctx context.Context, formats strfmt.Registry) error {

	if m.ComponentHistory != nil {

		if swag.IsZero(m.ComponentHistory) { // not required
			return nil
		}

		if err := m.ComponentHistory.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("componentHistory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("componentHistory")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ComponentHistoryEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComponentHistoryEntity) UnmarshalBinary(b []byte) error {
	var res ComponentHistoryEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
