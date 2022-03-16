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

// ScheduleComponentsEntity schedule components entity
//
// swagger:model ScheduleComponentsEntity
type ScheduleComponentsEntity struct {

	// Optional components to schedule. If not specified, all authorized descendant components will be used.
	Components map[string]RevisionDTO `json:"components,omitempty"`

	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged bool `json:"disconnectedNodeAcknowledged,omitempty"`

	// The id of the ProcessGroup
	ID string `json:"id,omitempty"`

	// The desired state of the descendant components
	// Enum: [RUNNING STOPPED ENABLED DISABLED]
	State string `json:"state,omitempty"`
}

// Validate validates this schedule components entity
func (m *ScheduleComponentsEntity) Validate(formats strfmt.Registry) error {
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

func (m *ScheduleComponentsEntity) validateComponents(formats strfmt.Registry) error {
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

var scheduleComponentsEntityTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RUNNING","STOPPED","ENABLED","DISABLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		scheduleComponentsEntityTypeStatePropEnum = append(scheduleComponentsEntityTypeStatePropEnum, v)
	}
}

const (

	// ScheduleComponentsEntityStateRUNNING captures enum value "RUNNING"
	ScheduleComponentsEntityStateRUNNING string = "RUNNING"

	// ScheduleComponentsEntityStateSTOPPED captures enum value "STOPPED"
	ScheduleComponentsEntityStateSTOPPED string = "STOPPED"

	// ScheduleComponentsEntityStateENABLED captures enum value "ENABLED"
	ScheduleComponentsEntityStateENABLED string = "ENABLED"

	// ScheduleComponentsEntityStateDISABLED captures enum value "DISABLED"
	ScheduleComponentsEntityStateDISABLED string = "DISABLED"
)

// prop value enum
func (m *ScheduleComponentsEntity) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, scheduleComponentsEntityTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ScheduleComponentsEntity) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this schedule components entity based on the context it is used
func (m *ScheduleComponentsEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComponents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScheduleComponentsEntity) contextValidateComponents(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ScheduleComponentsEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScheduleComponentsEntity) UnmarshalBinary(b []byte) error {
	var res ScheduleComponentsEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
