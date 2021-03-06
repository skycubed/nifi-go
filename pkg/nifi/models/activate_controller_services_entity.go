// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ActivateControllerServicesEntity activate controller services entity
//
// swagger:model ActivateControllerServicesEntity
type ActivateControllerServicesEntity struct {

	// Optional services to schedule. If not specified, all authorized descendant controller services will be used.
	Components map[string]RevisionDTO `json:"components,omitempty"`

	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged bool `json:"disconnectedNodeAcknowledged,omitempty"`

	// The id of the ProcessGroup
	ID string `json:"id,omitempty"`

	// The desired state of the descendant components
	// Enum: [ENABLED DISABLED]
	State string `json:"state,omitempty"`
}

// Validate validates this activate controller services entity
func (m *ActivateControllerServicesEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComponents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActivateControllerServicesEntity) validateComponents(formats strfmt.Registry) error {
	if swag.IsZero(m.Components) { // not required
		return nil
	}

	for k := range m.Components {

		if err := validate.Required("components"+"."+k, "body", m.Components[k]); err != nil {
			return err
		}
		if val, ok := m.Components[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("components" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("components" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

var activateControllerServicesEntityTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ENABLED","DISABLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		activateControllerServicesEntityTypeStatePropEnum = append(activateControllerServicesEntityTypeStatePropEnum, v)
	}
}

const (

	// ActivateControllerServicesEntityStateENABLED captures enum value "ENABLED"
	ActivateControllerServicesEntityStateENABLED string = "ENABLED"

	// ActivateControllerServicesEntityStateDISABLED captures enum value "DISABLED"
	ActivateControllerServicesEntityStateDISABLED string = "DISABLED"
)

// prop value enum
func (m *ActivateControllerServicesEntity) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, activateControllerServicesEntityTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ActivateControllerServicesEntity) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this activate controller services entity based on the context it is used
func (m *ActivateControllerServicesEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComponents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActivateControllerServicesEntity) contextValidateComponents(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Components {

		if val, ok := m.Components[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ActivateControllerServicesEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActivateControllerServicesEntity) UnmarshalBinary(b []byte) error {
	var res ActivateControllerServicesEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
