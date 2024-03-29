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

// ProvenanceOptionsEntity provenance options entity
//
// swagger:model ProvenanceOptionsEntity
type ProvenanceOptionsEntity struct {

	// provenance options
	ProvenanceOptions *ProvenanceOptionsDTO `json:"provenanceOptions,omitempty"`
}

// Validate validates this provenance options entity
func (m *ProvenanceOptionsEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProvenanceOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProvenanceOptionsEntity) validateProvenanceOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.ProvenanceOptions) { // not required
		return nil
	}

	if m.ProvenanceOptions != nil {
		if err := m.ProvenanceOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("provenanceOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("provenanceOptions")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this provenance options entity based on the context it is used
func (m *ProvenanceOptionsEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProvenanceOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProvenanceOptionsEntity) contextValidateProvenanceOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.ProvenanceOptions != nil {

		if swag.IsZero(m.ProvenanceOptions) { // not required
			return nil
		}

		if err := m.ProvenanceOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("provenanceOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("provenanceOptions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProvenanceOptionsEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProvenanceOptionsEntity) UnmarshalBinary(b []byte) error {
	var res ProvenanceOptionsEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
