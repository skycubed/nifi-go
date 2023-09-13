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

// ProcessorTypesEntity processor types entity
//
// swagger:model ProcessorTypesEntity
type ProcessorTypesEntity struct {

	// processor types
	// Unique: true
	ProcessorTypes []*DocumentedTypeDTO `json:"processorTypes"`
}

// Validate validates this processor types entity
func (m *ProcessorTypesEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProcessorTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProcessorTypesEntity) validateProcessorTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.ProcessorTypes) { // not required
		return nil
	}

	if err := validate.UniqueItems("processorTypes", "body", m.ProcessorTypes); err != nil {
		return err
	}

	for i := 0; i < len(m.ProcessorTypes); i++ {
		if swag.IsZero(m.ProcessorTypes[i]) { // not required
			continue
		}

		if m.ProcessorTypes[i] != nil {
			if err := m.ProcessorTypes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("processorTypes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("processorTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this processor types entity based on the context it is used
func (m *ProcessorTypesEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProcessorTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProcessorTypesEntity) contextValidateProcessorTypes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ProcessorTypes); i++ {

		if m.ProcessorTypes[i] != nil {

			if swag.IsZero(m.ProcessorTypes[i]) { // not required
				return nil
			}

			if err := m.ProcessorTypes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("processorTypes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("processorTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProcessorTypesEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProcessorTypesEntity) UnmarshalBinary(b []byte) error {
	var res ProcessorTypesEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
