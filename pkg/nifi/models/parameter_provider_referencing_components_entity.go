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

// ParameterProviderReferencingComponentsEntity parameter provider referencing components entity
//
// swagger:model ParameterProviderReferencingComponentsEntity
type ParameterProviderReferencingComponentsEntity struct {

	// parameter provider referencing components
	// Unique: true
	ParameterProviderReferencingComponents []*ParameterProviderReferencingComponentEntity `json:"parameterProviderReferencingComponents"`
}

// Validate validates this parameter provider referencing components entity
func (m *ParameterProviderReferencingComponentsEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParameterProviderReferencingComponents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ParameterProviderReferencingComponentsEntity) validateParameterProviderReferencingComponents(formats strfmt.Registry) error {
	if swag.IsZero(m.ParameterProviderReferencingComponents) { // not required
		return nil
	}

	if err := validate.UniqueItems("parameterProviderReferencingComponents", "body", m.ParameterProviderReferencingComponents); err != nil {
		return err
	}

	for i := 0; i < len(m.ParameterProviderReferencingComponents); i++ {
		if swag.IsZero(m.ParameterProviderReferencingComponents[i]) { // not required
			continue
		}

		if m.ParameterProviderReferencingComponents[i] != nil {
			if err := m.ParameterProviderReferencingComponents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameterProviderReferencingComponents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parameterProviderReferencingComponents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this parameter provider referencing components entity based on the context it is used
func (m *ParameterProviderReferencingComponentsEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateParameterProviderReferencingComponents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ParameterProviderReferencingComponentsEntity) contextValidateParameterProviderReferencingComponents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ParameterProviderReferencingComponents); i++ {

		if m.ParameterProviderReferencingComponents[i] != nil {
			if err := m.ParameterProviderReferencingComponents[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameterProviderReferencingComponents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parameterProviderReferencingComponents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ParameterProviderReferencingComponentsEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ParameterProviderReferencingComponentsEntity) UnmarshalBinary(b []byte) error {
	var res ParameterProviderReferencingComponentsEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}