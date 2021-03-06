// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VersionedParameter versioned parameter
//
// swagger:model VersionedParameter
type VersionedParameter struct {

	// The description of the param
	Description string `json:"description,omitempty"`

	// The name of the parameter
	Name string `json:"name,omitempty"`

	// Whether or not the parameter value is sensitive
	Sensitive bool `json:"sensitive,omitempty"`

	// The value of the parameter
	Value string `json:"value,omitempty"`
}

// Validate validates this versioned parameter
func (m *VersionedParameter) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this versioned parameter based on context it is used
func (m *VersionedParameter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VersionedParameter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VersionedParameter) UnmarshalBinary(b []byte) error {
	var res VersionedParameter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
