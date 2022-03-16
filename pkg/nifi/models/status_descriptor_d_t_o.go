// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StatusDescriptorDTO status descriptor d t o
//
// swagger:model StatusDescriptorDTO
type StatusDescriptorDTO struct {

	// The description of the status field.
	Description string `json:"description,omitempty"`

	// The name of the status field.
	Field string `json:"field,omitempty"`

	// The formatter for the status descriptor.
	Formatter string `json:"formatter,omitempty"`

	// The label for the status field.
	Label string `json:"label,omitempty"`
}

// Validate validates this status descriptor d t o
func (m *StatusDescriptorDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this status descriptor d t o based on context it is used
func (m *StatusDescriptorDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StatusDescriptorDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatusDescriptorDTO) UnmarshalBinary(b []byte) error {
	var res StatusDescriptorDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
