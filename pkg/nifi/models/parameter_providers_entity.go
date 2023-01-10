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

// ParameterProvidersEntity parameter providers entity
//
// swagger:model ParameterProvidersEntity
type ParameterProvidersEntity struct {

	// parameter providers
	// Unique: true
	ParameterProviders []*ParameterProviderEntity `json:"parameterProviders"`
}

// Validate validates this parameter providers entity
func (m *ParameterProvidersEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParameterProviders(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ParameterProvidersEntity) validateParameterProviders(formats strfmt.Registry) error {
	if swag.IsZero(m.ParameterProviders) { // not required
		return nil
	}

	if err := validate.UniqueItems("parameterProviders", "body", m.ParameterProviders); err != nil {
		return err
	}

	for i := 0; i < len(m.ParameterProviders); i++ {
		if swag.IsZero(m.ParameterProviders[i]) { // not required
			continue
		}

		if m.ParameterProviders[i] != nil {
			if err := m.ParameterProviders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameterProviders" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parameterProviders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this parameter providers entity based on the context it is used
func (m *ParameterProvidersEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateParameterProviders(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ParameterProvidersEntity) contextValidateParameterProviders(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ParameterProviders); i++ {

		if m.ParameterProviders[i] != nil {
			if err := m.ParameterProviders[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameterProviders" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parameterProviders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ParameterProvidersEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ParameterProvidersEntity) UnmarshalBinary(b []byte) error {
	var res ParameterProvidersEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}