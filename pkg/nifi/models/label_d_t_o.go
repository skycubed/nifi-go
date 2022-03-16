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

// LabelDTO label d t o
//
// swagger:model LabelDTO
type LabelDTO struct {

	// The height of the label in pixels when at a 1:1 scale.
	Height float64 `json:"height,omitempty"`

	// The id of the component.
	ID string `json:"id,omitempty"`

	// The text that appears in the label.
	Label string `json:"label,omitempty"`

	// The id of parent process group of this component if applicable.
	ParentGroupID string `json:"parentGroupId,omitempty"`

	// The position of this component in the UI if applicable.
	Position *PositionDTO `json:"position,omitempty"`

	// The styles for this label (font-size : 12px, background-color : #eee, etc).
	Style map[string]string `json:"style,omitempty"`

	// The ID of the corresponding component that is under version control
	VersionedComponentID string `json:"versionedComponentId,omitempty"`

	// The width of the label in pixels when at a 1:1 scale.
	Width float64 `json:"width,omitempty"`
}

// Validate validates this label d t o
func (m *LabelDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LabelDTO) validatePosition(formats strfmt.Registry) error {
	if swag.IsZero(m.Position) { // not required
		return nil
	}

	if m.Position != nil {
		if err := m.Position.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("position")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("position")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this label d t o based on the context it is used
func (m *LabelDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePosition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LabelDTO) contextValidatePosition(ctx context.Context, formats strfmt.Registry) error {

	if m.Position != nil {
		if err := m.Position.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("position")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("position")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LabelDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LabelDTO) UnmarshalBinary(b []byte) error {
	var res LabelDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
