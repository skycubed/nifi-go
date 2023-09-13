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

// ControllerServicesEntity controller services entity
//
// swagger:model ControllerServicesEntity
type ControllerServicesEntity struct {

	// controller services
	// Unique: true
	ControllerServices []*ControllerServiceEntity `json:"controllerServices"`

	// The current time on the system.
	CurrentTime string `json:"currentTime,omitempty"`
}

// Validate validates this controller services entity
func (m *ControllerServicesEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateControllerServices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ControllerServicesEntity) validateControllerServices(formats strfmt.Registry) error {
	if swag.IsZero(m.ControllerServices) { // not required
		return nil
	}

	if err := validate.UniqueItems("controllerServices", "body", m.ControllerServices); err != nil {
		return err
	}

	for i := 0; i < len(m.ControllerServices); i++ {
		if swag.IsZero(m.ControllerServices[i]) { // not required
			continue
		}

		if m.ControllerServices[i] != nil {
			if err := m.ControllerServices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("controllerServices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("controllerServices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this controller services entity based on the context it is used
func (m *ControllerServicesEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateControllerServices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ControllerServicesEntity) contextValidateControllerServices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ControllerServices); i++ {

		if m.ControllerServices[i] != nil {

			if swag.IsZero(m.ControllerServices[i]) { // not required
				return nil
			}

			if err := m.ControllerServices[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("controllerServices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("controllerServices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ControllerServicesEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ControllerServicesEntity) UnmarshalBinary(b []byte) error {
	var res ControllerServicesEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
