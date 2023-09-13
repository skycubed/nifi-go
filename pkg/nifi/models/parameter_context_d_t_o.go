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

// ParameterContextDTO parameter context d t o
//
// swagger:model ParameterContextDTO
type ParameterContextDTO struct {

	// The Process Groups that are bound to this Parameter Context
	// Unique: true
	BoundProcessGroups []*ProcessGroupEntity `json:"boundProcessGroups"`

	// The Description of the Parameter Context.
	Description string `json:"description,omitempty"`

	// The ID the Parameter Context.
	ID string `json:"id,omitempty"`

	// A list of references of Parameter Contexts from which this one inherits parameters
	InheritedParameterContexts []*ParameterContextReferenceEntity `json:"inheritedParameterContexts"`

	// The Name of the Parameter Context.
	Name string `json:"name,omitempty"`

	// Optional configuration for a Parameter Provider
	ParameterProviderConfiguration *ParameterProviderConfigurationEntity `json:"parameterProviderConfiguration,omitempty"`

	// The Parameters for the Parameter Context
	// Unique: true
	Parameters []*ParameterEntity `json:"parameters"`
}

// Validate validates this parameter context d t o
func (m *ParameterContextDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBoundProcessGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInheritedParameterContexts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameterProviderConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ParameterContextDTO) validateBoundProcessGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.BoundProcessGroups) { // not required
		return nil
	}

	if err := validate.UniqueItems("boundProcessGroups", "body", m.BoundProcessGroups); err != nil {
		return err
	}

	for i := 0; i < len(m.BoundProcessGroups); i++ {
		if swag.IsZero(m.BoundProcessGroups[i]) { // not required
			continue
		}

		if m.BoundProcessGroups[i] != nil {
			if err := m.BoundProcessGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("boundProcessGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("boundProcessGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ParameterContextDTO) validateInheritedParameterContexts(formats strfmt.Registry) error {
	if swag.IsZero(m.InheritedParameterContexts) { // not required
		return nil
	}

	for i := 0; i < len(m.InheritedParameterContexts); i++ {
		if swag.IsZero(m.InheritedParameterContexts[i]) { // not required
			continue
		}

		if m.InheritedParameterContexts[i] != nil {
			if err := m.InheritedParameterContexts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inheritedParameterContexts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("inheritedParameterContexts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ParameterContextDTO) validateParameterProviderConfiguration(formats strfmt.Registry) error {
	if swag.IsZero(m.ParameterProviderConfiguration) { // not required
		return nil
	}

	if m.ParameterProviderConfiguration != nil {
		if err := m.ParameterProviderConfiguration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parameterProviderConfiguration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parameterProviderConfiguration")
			}
			return err
		}
	}

	return nil
}

func (m *ParameterContextDTO) validateParameters(formats strfmt.Registry) error {
	if swag.IsZero(m.Parameters) { // not required
		return nil
	}

	if err := validate.UniqueItems("parameters", "body", m.Parameters); err != nil {
		return err
	}

	for i := 0; i < len(m.Parameters); i++ {
		if swag.IsZero(m.Parameters[i]) { // not required
			continue
		}

		if m.Parameters[i] != nil {
			if err := m.Parameters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this parameter context d t o based on the context it is used
func (m *ParameterContextDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBoundProcessGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInheritedParameterContexts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParameterProviderConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParameters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ParameterContextDTO) contextValidateBoundProcessGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BoundProcessGroups); i++ {

		if m.BoundProcessGroups[i] != nil {

			if swag.IsZero(m.BoundProcessGroups[i]) { // not required
				return nil
			}

			if err := m.BoundProcessGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("boundProcessGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("boundProcessGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ParameterContextDTO) contextValidateInheritedParameterContexts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InheritedParameterContexts); i++ {

		if m.InheritedParameterContexts[i] != nil {

			if swag.IsZero(m.InheritedParameterContexts[i]) { // not required
				return nil
			}

			if err := m.InheritedParameterContexts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inheritedParameterContexts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("inheritedParameterContexts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ParameterContextDTO) contextValidateParameterProviderConfiguration(ctx context.Context, formats strfmt.Registry) error {

	if m.ParameterProviderConfiguration != nil {

		if swag.IsZero(m.ParameterProviderConfiguration) { // not required
			return nil
		}

		if err := m.ParameterProviderConfiguration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parameterProviderConfiguration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parameterProviderConfiguration")
			}
			return err
		}
	}

	return nil
}

func (m *ParameterContextDTO) contextValidateParameters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Parameters); i++ {

		if m.Parameters[i] != nil {

			if swag.IsZero(m.Parameters[i]) { // not required
				return nil
			}

			if err := m.Parameters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ParameterContextDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ParameterContextDTO) UnmarshalBinary(b []byte) error {
	var res ParameterContextDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
