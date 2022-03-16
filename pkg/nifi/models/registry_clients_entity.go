// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RegistryClientsEntity registry clients entity
//
// swagger:model RegistryClientsEntity
type RegistryClientsEntity struct {

	// registries
	// Unique: true
	Registries []*RegistryClientEntity `json:"registries"`
}

// Validate validates this registry clients entity
func (m *RegistryClientsEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRegistries(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegistryClientsEntity) validateRegistries(formats strfmt.Registry) error {
	if swag.IsZero(m.Registries) { // not required
		return nil
	}

	if err := validate.UniqueItems("registries", "body", m.Registries); err != nil {
		return err
	}

	for i := 0; i < len(m.Registries); i++ {
		if swag.IsZero(m.Registries[i]) { // not required
			continue
		}

		if m.Registries[i] != nil {
			if err := m.Registries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("registries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("registries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this registry clients entity based on the context it is used
func (m *RegistryClientsEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRegistries(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegistryClientsEntity) contextValidateRegistries(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Registries); i++ {

		if m.Registries[i] != nil {
			if err := m.Registries[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("registries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("registries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RegistryClientsEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegistryClientsEntity) UnmarshalBinary(b []byte) error {
	var res RegistryClientsEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}