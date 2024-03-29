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

// VersionedPort versioned port
//
// swagger:model VersionedPort
type VersionedPort struct {

	// Whether or not this port allows remote access for site-to-site
	AllowRemoteAccess bool `json:"allowRemoteAccess,omitempty"`

	// The user-supplied comments for the component
	Comments string `json:"comments,omitempty"`

	// component type
	// Enum: [CONNECTION PROCESSOR PROCESS_GROUP REMOTE_PROCESS_GROUP INPUT_PORT OUTPUT_PORT REMOTE_INPUT_PORT REMOTE_OUTPUT_PORT FUNNEL LABEL CONTROLLER_SERVICE]
	ComponentType string `json:"componentType,omitempty"`

	// The number of tasks that should be concurrently scheduled for the port.
	ConcurrentlySchedulableTaskCount int32 `json:"concurrentlySchedulableTaskCount,omitempty"`

	// The ID of the Process Group that this component belongs to
	GroupIdentifier string `json:"groupIdentifier,omitempty"`

	// The component's unique identifier
	Identifier string `json:"identifier,omitempty"`

	// The component's name
	Name string `json:"name,omitempty"`

	// The component's position on the graph
	Position *Position `json:"position,omitempty"`

	// The scheduled state of the component
	// Enum: [ENABLED DISABLED]
	ScheduledState string `json:"scheduledState,omitempty"`

	// The type of port.
	// Enum: [INPUT_PORT OUTPUT_PORT]
	Type string `json:"type,omitempty"`
}

// Validate validates this versioned port
func (m *VersionedPort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComponentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheduledState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var versionedPortTypeComponentTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CONNECTION","PROCESSOR","PROCESS_GROUP","REMOTE_PROCESS_GROUP","INPUT_PORT","OUTPUT_PORT","REMOTE_INPUT_PORT","REMOTE_OUTPUT_PORT","FUNNEL","LABEL","CONTROLLER_SERVICE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		versionedPortTypeComponentTypePropEnum = append(versionedPortTypeComponentTypePropEnum, v)
	}
}

const (

	// VersionedPortComponentTypeCONNECTION captures enum value "CONNECTION"
	VersionedPortComponentTypeCONNECTION string = "CONNECTION"

	// VersionedPortComponentTypePROCESSOR captures enum value "PROCESSOR"
	VersionedPortComponentTypePROCESSOR string = "PROCESSOR"

	// VersionedPortComponentTypePROCESSGROUP captures enum value "PROCESS_GROUP"
	VersionedPortComponentTypePROCESSGROUP string = "PROCESS_GROUP"

	// VersionedPortComponentTypeREMOTEPROCESSGROUP captures enum value "REMOTE_PROCESS_GROUP"
	VersionedPortComponentTypeREMOTEPROCESSGROUP string = "REMOTE_PROCESS_GROUP"

	// VersionedPortComponentTypeINPUTPORT captures enum value "INPUT_PORT"
	VersionedPortComponentTypeINPUTPORT string = "INPUT_PORT"

	// VersionedPortComponentTypeOUTPUTPORT captures enum value "OUTPUT_PORT"
	VersionedPortComponentTypeOUTPUTPORT string = "OUTPUT_PORT"

	// VersionedPortComponentTypeREMOTEINPUTPORT captures enum value "REMOTE_INPUT_PORT"
	VersionedPortComponentTypeREMOTEINPUTPORT string = "REMOTE_INPUT_PORT"

	// VersionedPortComponentTypeREMOTEOUTPUTPORT captures enum value "REMOTE_OUTPUT_PORT"
	VersionedPortComponentTypeREMOTEOUTPUTPORT string = "REMOTE_OUTPUT_PORT"

	// VersionedPortComponentTypeFUNNEL captures enum value "FUNNEL"
	VersionedPortComponentTypeFUNNEL string = "FUNNEL"

	// VersionedPortComponentTypeLABEL captures enum value "LABEL"
	VersionedPortComponentTypeLABEL string = "LABEL"

	// VersionedPortComponentTypeCONTROLLERSERVICE captures enum value "CONTROLLER_SERVICE"
	VersionedPortComponentTypeCONTROLLERSERVICE string = "CONTROLLER_SERVICE"
)

// prop value enum
func (m *VersionedPort) validateComponentTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, versionedPortTypeComponentTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VersionedPort) validateComponentType(formats strfmt.Registry) error {
	if swag.IsZero(m.ComponentType) { // not required
		return nil
	}

	// value enum
	if err := m.validateComponentTypeEnum("componentType", "body", m.ComponentType); err != nil {
		return err
	}

	return nil
}

func (m *VersionedPort) validatePosition(formats strfmt.Registry) error {
	if swag.IsZero(m.Position) { // not required
		return nil
	}

	if m.Position != nil {
		if err := m.Position.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("position")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("position")
			}
			return err
		}
	}

	return nil
}

var versionedPortTypeScheduledStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ENABLED","DISABLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		versionedPortTypeScheduledStatePropEnum = append(versionedPortTypeScheduledStatePropEnum, v)
	}
}

const (

	// VersionedPortScheduledStateENABLED captures enum value "ENABLED"
	VersionedPortScheduledStateENABLED string = "ENABLED"

	// VersionedPortScheduledStateDISABLED captures enum value "DISABLED"
	VersionedPortScheduledStateDISABLED string = "DISABLED"
)

// prop value enum
func (m *VersionedPort) validateScheduledStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, versionedPortTypeScheduledStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VersionedPort) validateScheduledState(formats strfmt.Registry) error {
	if swag.IsZero(m.ScheduledState) { // not required
		return nil
	}

	// value enum
	if err := m.validateScheduledStateEnum("scheduledState", "body", m.ScheduledState); err != nil {
		return err
	}

	return nil
}

var versionedPortTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["INPUT_PORT","OUTPUT_PORT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		versionedPortTypeTypePropEnum = append(versionedPortTypeTypePropEnum, v)
	}
}

const (

	// VersionedPortTypeINPUTPORT captures enum value "INPUT_PORT"
	VersionedPortTypeINPUTPORT string = "INPUT_PORT"

	// VersionedPortTypeOUTPUTPORT captures enum value "OUTPUT_PORT"
	VersionedPortTypeOUTPUTPORT string = "OUTPUT_PORT"
)

// prop value enum
func (m *VersionedPort) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, versionedPortTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VersionedPort) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this versioned port based on the context it is used
func (m *VersionedPort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePosition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VersionedPort) contextValidatePosition(ctx context.Context, formats strfmt.Registry) error {

	if m.Position != nil {

		if swag.IsZero(m.Position) { // not required
			return nil
		}

		if err := m.Position.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("position")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("position")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VersionedPort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VersionedPort) UnmarshalBinary(b []byte) error {
	var res VersionedPort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
