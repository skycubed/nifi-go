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

// StatusHistoryEntity status history entity
//
// swagger:model StatusHistoryEntity
type StatusHistoryEntity struct {

	// Indicates whether the user can read a given resource.
	CanRead bool `json:"canRead,omitempty"`

	// status history
	StatusHistory *StatusHistoryDTO `json:"statusHistory,omitempty"`
}

// Validate validates this status history entity
func (m *StatusHistoryEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatusHistory(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StatusHistoryEntity) validateStatusHistory(formats strfmt.Registry) error {
	if swag.IsZero(m.StatusHistory) { // not required
		return nil
	}

	if m.StatusHistory != nil {
		if err := m.StatusHistory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statusHistory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("statusHistory")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this status history entity based on the context it is used
func (m *StatusHistoryEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatusHistory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StatusHistoryEntity) contextValidateStatusHistory(ctx context.Context, formats strfmt.Registry) error {

	if m.StatusHistory != nil {

		if swag.IsZero(m.StatusHistory) { // not required
			return nil
		}

		if err := m.StatusHistory.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statusHistory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("statusHistory")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StatusHistoryEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatusHistoryEntity) UnmarshalBinary(b []byte) error {
	var res StatusHistoryEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
