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

// InputPortsEntity input ports entity
//
// swagger:model InputPortsEntity
type InputPortsEntity struct {

	// input ports
	// Unique: true
	InputPorts []*PortEntity `json:"inputPorts"`
}

// Validate validates this input ports entity
func (m *InputPortsEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInputPorts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InputPortsEntity) validateInputPorts(formats strfmt.Registry) error {
	if swag.IsZero(m.InputPorts) { // not required
		return nil
	}

	if err := validate.UniqueItems("inputPorts", "body", m.InputPorts); err != nil {
		return err
	}

	for i := 0; i < len(m.InputPorts); i++ {
		if swag.IsZero(m.InputPorts[i]) { // not required
			continue
		}

		if m.InputPorts[i] != nil {
			if err := m.InputPorts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inputPorts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("inputPorts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this input ports entity based on the context it is used
func (m *InputPortsEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInputPorts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InputPortsEntity) contextValidateInputPorts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InputPorts); i++ {

		if m.InputPorts[i] != nil {
			if err := m.InputPorts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inputPorts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("inputPorts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InputPortsEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InputPortsEntity) UnmarshalBinary(b []byte) error {
	var res InputPortsEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
