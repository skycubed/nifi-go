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

// PrioritizerTypesEntity prioritizer types entity
//
// swagger:model PrioritizerTypesEntity
type PrioritizerTypesEntity struct {

	// prioritizer types
	// Unique: true
	PrioritizerTypes []*DocumentedTypeDTO `json:"prioritizerTypes"`
}

// Validate validates this prioritizer types entity
func (m *PrioritizerTypesEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrioritizerTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrioritizerTypesEntity) validatePrioritizerTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.PrioritizerTypes) { // not required
		return nil
	}

	if err := validate.UniqueItems("prioritizerTypes", "body", m.PrioritizerTypes); err != nil {
		return err
	}

	for i := 0; i < len(m.PrioritizerTypes); i++ {
		if swag.IsZero(m.PrioritizerTypes[i]) { // not required
			continue
		}

		if m.PrioritizerTypes[i] != nil {
			if err := m.PrioritizerTypes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("prioritizerTypes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("prioritizerTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this prioritizer types entity based on the context it is used
func (m *PrioritizerTypesEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePrioritizerTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrioritizerTypesEntity) contextValidatePrioritizerTypes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PrioritizerTypes); i++ {

		if m.PrioritizerTypes[i] != nil {

			if swag.IsZero(m.PrioritizerTypes[i]) { // not required
				return nil
			}

			if err := m.PrioritizerTypes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("prioritizerTypes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("prioritizerTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PrioritizerTypesEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrioritizerTypesEntity) UnmarshalBinary(b []byte) error {
	var res PrioritizerTypesEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
