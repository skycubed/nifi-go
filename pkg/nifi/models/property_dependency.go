// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PropertyDependency property dependency
//
// swagger:model PropertyDependency
type PropertyDependency struct {

	// The values that satisfy the dependency
	DependentValues []string `json:"dependentValues"`

	// The name of the property that is depended upon
	PropertyDisplayName string `json:"propertyDisplayName,omitempty"`

	// The name of the property that is depended upon
	PropertyName string `json:"propertyName,omitempty"`
}

// Validate validates this property dependency
func (m *PropertyDependency) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this property dependency based on context it is used
func (m *PropertyDependency) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PropertyDependency) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PropertyDependency) UnmarshalBinary(b []byte) error {
	var res PropertyDependency
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}