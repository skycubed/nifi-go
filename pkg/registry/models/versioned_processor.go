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

// VersionedProcessor versioned processor
//
// swagger:model VersionedProcessor
type VersionedProcessor struct {

	// The annotation data for the processor used to relay configuration between a custom UI and the procesosr.
	AnnotationData string `json:"annotationData"`

	// The names of all relationships that cause a flow file to be terminated if the relationship is not connected elsewhere. This property differs from the 'isAutoTerminate' property of the RelationshipDTO in that the RelationshipDTO is meant to depict the current configuration, whereas this property can be set in a DTO when updating a Processor in order to change which Relationships should be auto-terminated.
	// Unique: true
	AutoTerminatedRelationships []string `json:"autoTerminatedRelationships"`

	// The level at which the processor will report bulletins.
	BulletinLevel string `json:"bulletinLevel,omitempty"`

	// Information about the bundle from which the component came
	Bundle *Bundle `json:"bundle,omitempty"`

	// The user-supplied comments for the component
	Comments string `json:"comments"`

	// component type
	// Enum: [CONNECTION PROCESSOR PROCESS_GROUP REMOTE_PROCESS_GROUP INPUT_PORT OUTPUT_PORT REMOTE_INPUT_PORT REMOTE_OUTPUT_PORT FUNNEL LABEL CONTROLLER_SERVICE]
	ComponentType string `json:"componentType,omitempty"`

	// The number of tasks that should be concurrently schedule for the processor. If the processor doesn't allow parallol processing then any positive input will be ignored.
	ConcurrentlySchedulableTaskCount int32 `json:"concurrentlySchedulableTaskCount,omitempty"`

	// Indicates the node where the process will execute.
	ExecutionNode string `json:"executionNode,omitempty"`

	// The ID of the Process Group that this component belongs to
	GroupIdentifier string `json:"groupIdentifier,omitempty"`

	// The component's unique identifier
	Identifier string `json:"identifier,omitempty"`

	// The component's name
	Name string `json:"name,omitempty"`

	// The amout of time that is used when the process penalizes a flowfile.
	PenaltyDuration string `json:"penaltyDuration,omitempty"`

	// The component's position on the graph
	Position *Position `json:"position,omitempty"`

	// The properties for the processor. Properties whose value is not set will only contain the property name.
	Properties map[string]string `json:"properties,omitempty"`

	// The property descriptors for the processor.
	PropertyDescriptors map[string]VersionedPropertyDescriptor `json:"propertyDescriptors,omitempty"`

	// The run duration for the processor in milliseconds.
	RunDurationMillis int64 `json:"runDurationMillis"`

	// The scheduled state of the component
	// Enum: [ENABLED DISABLED]
	ScheduledState string `json:"scheduledState,omitempty"`

	// The frequency with which to schedule the processor. The format of the value will depend on th value of schedulingStrategy.
	SchedulingPeriod string `json:"schedulingPeriod,omitempty"`

	// Indcates whether the prcessor should be scheduled to run in event or timer driven mode.
	SchedulingStrategy string `json:"schedulingStrategy,omitempty"`

	// Stylistic data for rendering in a UI
	Style map[string]string `json:"style"`

	// The type of Processor
	Type string `json:"type,omitempty"`

	// The amount of time that must elapse before this processor is scheduled again after yielding.
	YieldDuration string `json:"yieldDuration,omitempty"`
}

// Validate validates this versioned processor
func (m *VersionedProcessor) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoTerminatedRelationships(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBundle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComponentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePropertyDescriptors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheduledState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VersionedProcessor) validateAutoTerminatedRelationships(formats strfmt.Registry) error {
	if swag.IsZero(m.AutoTerminatedRelationships) { // not required
		return nil
	}

	if err := validate.UniqueItems("autoTerminatedRelationships", "body", m.AutoTerminatedRelationships); err != nil {
		return err
	}

	return nil
}

func (m *VersionedProcessor) validateBundle(formats strfmt.Registry) error {
	if swag.IsZero(m.Bundle) { // not required
		return nil
	}

	if m.Bundle != nil {
		if err := m.Bundle.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bundle")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bundle")
			}
			return err
		}
	}

	return nil
}

var versionedProcessorTypeComponentTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CONNECTION","PROCESSOR","PROCESS_GROUP","REMOTE_PROCESS_GROUP","INPUT_PORT","OUTPUT_PORT","REMOTE_INPUT_PORT","REMOTE_OUTPUT_PORT","FUNNEL","LABEL","CONTROLLER_SERVICE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		versionedProcessorTypeComponentTypePropEnum = append(versionedProcessorTypeComponentTypePropEnum, v)
	}
}

const (

	// VersionedProcessorComponentTypeCONNECTION captures enum value "CONNECTION"
	VersionedProcessorComponentTypeCONNECTION string = "CONNECTION"

	// VersionedProcessorComponentTypePROCESSOR captures enum value "PROCESSOR"
	VersionedProcessorComponentTypePROCESSOR string = "PROCESSOR"

	// VersionedProcessorComponentTypePROCESSGROUP captures enum value "PROCESS_GROUP"
	VersionedProcessorComponentTypePROCESSGROUP string = "PROCESS_GROUP"

	// VersionedProcessorComponentTypeREMOTEPROCESSGROUP captures enum value "REMOTE_PROCESS_GROUP"
	VersionedProcessorComponentTypeREMOTEPROCESSGROUP string = "REMOTE_PROCESS_GROUP"

	// VersionedProcessorComponentTypeINPUTPORT captures enum value "INPUT_PORT"
	VersionedProcessorComponentTypeINPUTPORT string = "INPUT_PORT"

	// VersionedProcessorComponentTypeOUTPUTPORT captures enum value "OUTPUT_PORT"
	VersionedProcessorComponentTypeOUTPUTPORT string = "OUTPUT_PORT"

	// VersionedProcessorComponentTypeREMOTEINPUTPORT captures enum value "REMOTE_INPUT_PORT"
	VersionedProcessorComponentTypeREMOTEINPUTPORT string = "REMOTE_INPUT_PORT"

	// VersionedProcessorComponentTypeREMOTEOUTPUTPORT captures enum value "REMOTE_OUTPUT_PORT"
	VersionedProcessorComponentTypeREMOTEOUTPUTPORT string = "REMOTE_OUTPUT_PORT"

	// VersionedProcessorComponentTypeFUNNEL captures enum value "FUNNEL"
	VersionedProcessorComponentTypeFUNNEL string = "FUNNEL"

	// VersionedProcessorComponentTypeLABEL captures enum value "LABEL"
	VersionedProcessorComponentTypeLABEL string = "LABEL"

	// VersionedProcessorComponentTypeCONTROLLERSERVICE captures enum value "CONTROLLER_SERVICE"
	VersionedProcessorComponentTypeCONTROLLERSERVICE string = "CONTROLLER_SERVICE"
)

// prop value enum
func (m *VersionedProcessor) validateComponentTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, versionedProcessorTypeComponentTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VersionedProcessor) validateComponentType(formats strfmt.Registry) error {
	if swag.IsZero(m.ComponentType) { // not required
		return nil
	}

	// value enum
	if err := m.validateComponentTypeEnum("componentType", "body", m.ComponentType); err != nil {
		return err
	}

	return nil
}

func (m *VersionedProcessor) validatePosition(formats strfmt.Registry) error {
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

func (m *VersionedProcessor) validatePropertyDescriptors(formats strfmt.Registry) error {
	if swag.IsZero(m.PropertyDescriptors) { // not required
		return nil
	}

	for k := range m.PropertyDescriptors {

		if err := validate.Required("propertyDescriptors"+"."+k, "body", m.PropertyDescriptors[k]); err != nil {
			return err
		}
		if val, ok := m.PropertyDescriptors[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("propertyDescriptors" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("propertyDescriptors" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

var versionedProcessorTypeScheduledStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ENABLED","DISABLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		versionedProcessorTypeScheduledStatePropEnum = append(versionedProcessorTypeScheduledStatePropEnum, v)
	}
}

const (

	// VersionedProcessorScheduledStateENABLED captures enum value "ENABLED"
	VersionedProcessorScheduledStateENABLED string = "ENABLED"

	// VersionedProcessorScheduledStateDISABLED captures enum value "DISABLED"
	VersionedProcessorScheduledStateDISABLED string = "DISABLED"
)

// prop value enum
func (m *VersionedProcessor) validateScheduledStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, versionedProcessorTypeScheduledStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VersionedProcessor) validateScheduledState(formats strfmt.Registry) error {
	if swag.IsZero(m.ScheduledState) { // not required
		return nil
	}

	// value enum
	if err := m.validateScheduledStateEnum("scheduledState", "body", m.ScheduledState); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this versioned processor based on the context it is used
func (m *VersionedProcessor) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBundle(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePosition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePropertyDescriptors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VersionedProcessor) contextValidateBundle(ctx context.Context, formats strfmt.Registry) error {

	if m.Bundle != nil {

		if swag.IsZero(m.Bundle) { // not required
			return nil
		}

		if err := m.Bundle.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bundle")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bundle")
			}
			return err
		}
	}

	return nil
}

func (m *VersionedProcessor) contextValidatePosition(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VersionedProcessor) contextValidatePropertyDescriptors(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.PropertyDescriptors {

		if val, ok := m.PropertyDescriptors[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VersionedProcessor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VersionedProcessor) UnmarshalBinary(b []byte) error {
	var res VersionedProcessor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
