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

// ProcessGroupStatusEntity process group status entity
//
// swagger:model ProcessGroupStatusEntity
type ProcessGroupStatusEntity struct {

	// Indicates whether the user can read a given resource.
	CanRead bool `json:"canRead,omitempty"`

	// process group status
	ProcessGroupStatus *ProcessGroupStatusDTO `json:"processGroupStatus,omitempty"`
}

// Validate validates this process group status entity
func (m *ProcessGroupStatusEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProcessGroupStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProcessGroupStatusEntity) validateProcessGroupStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.ProcessGroupStatus) { // not required
		return nil
	}

	if m.ProcessGroupStatus != nil {
		if err := m.ProcessGroupStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("processGroupStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("processGroupStatus")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this process group status entity based on the context it is used
func (m *ProcessGroupStatusEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProcessGroupStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProcessGroupStatusEntity) contextValidateProcessGroupStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.ProcessGroupStatus != nil {

		if swag.IsZero(m.ProcessGroupStatus) { // not required
			return nil
		}

		if err := m.ProcessGroupStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("processGroupStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("processGroupStatus")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProcessGroupStatusEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProcessGroupStatusEntity) UnmarshalBinary(b []byte) error {
	var res ProcessGroupStatusEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
