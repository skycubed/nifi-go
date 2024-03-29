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

// ControllerServiceRunStatusEntity controller service run status entity
//
// swagger:model ControllerServiceRunStatusEntity
type ControllerServiceRunStatusEntity struct {

	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged bool `json:"disconnectedNodeAcknowledged,omitempty"`

	// The revision for this request/response. The revision is required for any mutable flow requests and is included in all responses.
	Revision *RevisionDTO `json:"revision,omitempty"`

	// The run status of the ControllerService.
	// Enum: [ENABLED DISABLED]
	State string `json:"state,omitempty"`
}

// Validate validates this controller service run status entity
func (m *ControllerServiceRunStatusEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRevision(formats); err != nil {
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

func (m *ControllerServiceRunStatusEntity) validateRevision(formats strfmt.Registry) error {
	if swag.IsZero(m.Revision) { // not required
		return nil
	}

	if m.Revision != nil {
		if err := m.Revision.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("revision")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("revision")
			}
			return err
		}
	}

	return nil
}

var controllerServiceRunStatusEntityTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ENABLED","DISABLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		controllerServiceRunStatusEntityTypeStatePropEnum = append(controllerServiceRunStatusEntityTypeStatePropEnum, v)
	}
}

const (

	// ControllerServiceRunStatusEntityStateENABLED captures enum value "ENABLED"
	ControllerServiceRunStatusEntityStateENABLED string = "ENABLED"

	// ControllerServiceRunStatusEntityStateDISABLED captures enum value "DISABLED"
	ControllerServiceRunStatusEntityStateDISABLED string = "DISABLED"
)

// prop value enum
func (m *ControllerServiceRunStatusEntity) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, controllerServiceRunStatusEntityTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ControllerServiceRunStatusEntity) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this controller service run status entity based on the context it is used
func (m *ControllerServiceRunStatusEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRevision(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ControllerServiceRunStatusEntity) contextValidateRevision(ctx context.Context, formats strfmt.Registry) error {

	if m.Revision != nil {

		if swag.IsZero(m.Revision) { // not required
			return nil
		}

		if err := m.Revision.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("revision")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("revision")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ControllerServiceRunStatusEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ControllerServiceRunStatusEntity) UnmarshalBinary(b []byte) error {
	var res ControllerServiceRunStatusEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
