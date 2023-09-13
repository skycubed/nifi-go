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

// ConfigurationAnalysisEntity configuration analysis entity
//
// swagger:model ConfigurationAnalysisEntity
type ConfigurationAnalysisEntity struct {

	// The configuration analysis
	ConfigurationAnalysis *ConfigurationAnalysisDTO `json:"configurationAnalysis,omitempty"`
}

// Validate validates this configuration analysis entity
func (m *ConfigurationAnalysisEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfigurationAnalysis(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigurationAnalysisEntity) validateConfigurationAnalysis(formats strfmt.Registry) error {
	if swag.IsZero(m.ConfigurationAnalysis) { // not required
		return nil
	}

	if m.ConfigurationAnalysis != nil {
		if err := m.ConfigurationAnalysis.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configurationAnalysis")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("configurationAnalysis")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configuration analysis entity based on the context it is used
func (m *ConfigurationAnalysisEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfigurationAnalysis(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigurationAnalysisEntity) contextValidateConfigurationAnalysis(ctx context.Context, formats strfmt.Registry) error {

	if m.ConfigurationAnalysis != nil {

		if swag.IsZero(m.ConfigurationAnalysis) { // not required
			return nil
		}

		if err := m.ConfigurationAnalysis.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configurationAnalysis")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("configurationAnalysis")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigurationAnalysisEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigurationAnalysisEntity) UnmarshalBinary(b []byte) error {
	var res ConfigurationAnalysisEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
