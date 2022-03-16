// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Restriction restriction
//
// swagger:model Restriction
type Restriction struct {

	// The explanation of this restriction
	Explanation string `json:"explanation,omitempty"`

	// The permission required for this restriction
	RequiredPermission string `json:"requiredPermission,omitempty"`
}

// Validate validates this restriction
func (m *Restriction) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this restriction based on context it is used
func (m *Restriction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Restriction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Restriction) UnmarshalBinary(b []byte) error {
	var res Restriction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
