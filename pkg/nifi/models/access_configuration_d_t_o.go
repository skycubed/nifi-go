// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AccessConfigurationDTO access configuration d t o
//
// swagger:model AccessConfigurationDTO
type AccessConfigurationDTO struct {

	// Indicates whether or not this NiFi supports user login.
	SupportsLogin bool `json:"supportsLogin,omitempty"`
}

// Validate validates this access configuration d t o
func (m *AccessConfigurationDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this access configuration d t o based on context it is used
func (m *AccessConfigurationDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccessConfigurationDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccessConfigurationDTO) UnmarshalBinary(b []byte) error {
	var res AccessConfigurationDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
