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

// ReportingTaskDefinition reporting task definition
//
// swagger:model ReportingTaskDefinition
type ReportingTaskDefinition struct {

	// The artifact name of the bundle that provides the referenced type.
	Artifact string `json:"artifact,omitempty"`

	// The build metadata for this component
	BuildInfo *BuildInfo `json:"buildInfo,omitempty"`

	// The default scheduling period for each scheduling strategy. The scheduling period is expected to be a time period, such as "30 sec".
	DefaultSchedulingPeriodBySchedulingStrategy map[string]string `json:"defaultSchedulingPeriodBySchedulingStrategy,omitempty"`

	// The default scheduling strategy for the reporting task.
	DefaultSchedulingStrategy string `json:"defaultSchedulingStrategy,omitempty"`

	// Whether or not the component has been deprecated
	Deprecated bool `json:"deprecated,omitempty"`

	// If this component has been deprecated, this optional field can be used to provide an explanation
	DeprecationReason string `json:"deprecationReason,omitempty"`

	// Explicit restrictions that indicate a require permission to use the component
	// Unique: true
	ExplicitRestrictions []*Restriction `json:"explicitRestrictions"`

	// The group name of the bundle that provides the referenced type.
	Group string `json:"group,omitempty"`

	// Descriptions of configuration properties applicable to this component.
	PropertyDescriptors map[string]PropertyDescriptor `json:"propertyDescriptors,omitempty"`

	// If this type represents a provider for an interface, this lists the APIs it implements
	ProvidedAPIImplementations []*DefinedType `json:"providedApiImplementations"`

	// Whether or not the component has a general restriction
	Restricted bool `json:"restricted,omitempty"`

	// An optional description of the general restriction
	RestrictedExplanation string `json:"restrictedExplanation,omitempty"`

	// stateful
	Stateful *Stateful `json:"stateful,omitempty"`

	// The supported scheduling strategies, such as TIME_DRIVER or CRON.
	SupportedSchedulingStrategies []string `json:"supportedSchedulingStrategies"`

	// Whether or not this component makes use of dynamic (user-set) properties.
	SupportsDynamicProperties bool `json:"supportsDynamicProperties,omitempty"`

	// The tags associated with this type
	// Unique: true
	Tags []string `json:"tags"`

	// The fully-qualified class type
	// Required: true
	Type *string `json:"type"`

	// The description of the type.
	TypeDescription string `json:"typeDescription,omitempty"`

	// The version of the bundle that provides the referenced type.
	Version string `json:"version,omitempty"`
}

// Validate validates this reporting task definition
func (m *ReportingTaskDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuildInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExplicitRestrictions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePropertyDescriptors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvidedAPIImplementations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStateful(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
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

func (m *ReportingTaskDefinition) validateBuildInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.BuildInfo) { // not required
		return nil
	}

	if m.BuildInfo != nil {
		if err := m.BuildInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("buildInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("buildInfo")
			}
			return err
		}
	}

	return nil
}

func (m *ReportingTaskDefinition) validateExplicitRestrictions(formats strfmt.Registry) error {
	if swag.IsZero(m.ExplicitRestrictions) { // not required
		return nil
	}

	if err := validate.UniqueItems("explicitRestrictions", "body", m.ExplicitRestrictions); err != nil {
		return err
	}

	for i := 0; i < len(m.ExplicitRestrictions); i++ {
		if swag.IsZero(m.ExplicitRestrictions[i]) { // not required
			continue
		}

		if m.ExplicitRestrictions[i] != nil {
			if err := m.ExplicitRestrictions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("explicitRestrictions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("explicitRestrictions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReportingTaskDefinition) validatePropertyDescriptors(formats strfmt.Registry) error {
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

func (m *ReportingTaskDefinition) validateProvidedAPIImplementations(formats strfmt.Registry) error {
	if swag.IsZero(m.ProvidedAPIImplementations) { // not required
		return nil
	}

	for i := 0; i < len(m.ProvidedAPIImplementations); i++ {
		if swag.IsZero(m.ProvidedAPIImplementations[i]) { // not required
			continue
		}

		if m.ProvidedAPIImplementations[i] != nil {
			if err := m.ProvidedAPIImplementations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("providedApiImplementations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("providedApiImplementations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReportingTaskDefinition) validateStateful(formats strfmt.Registry) error {
	if swag.IsZero(m.Stateful) { // not required
		return nil
	}

	if m.Stateful != nil {
		if err := m.Stateful.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stateful")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stateful")
			}
			return err
		}
	}

	return nil
}

func (m *ReportingTaskDefinition) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	if err := validate.UniqueItems("tags", "body", m.Tags); err != nil {
		return err
	}

	return nil
}

func (m *ReportingTaskDefinition) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this reporting task definition based on the context it is used
func (m *ReportingTaskDefinition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBuildInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExplicitRestrictions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePropertyDescriptors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProvidedAPIImplementations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStateful(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportingTaskDefinition) contextValidateBuildInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.BuildInfo != nil {

		if swag.IsZero(m.BuildInfo) { // not required
			return nil
		}

		if err := m.BuildInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("buildInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("buildInfo")
			}
			return err
		}
	}

	return nil
}

func (m *ReportingTaskDefinition) contextValidateExplicitRestrictions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExplicitRestrictions); i++ {

		if m.ExplicitRestrictions[i] != nil {

			if swag.IsZero(m.ExplicitRestrictions[i]) { // not required
				return nil
			}

			if err := m.ExplicitRestrictions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("explicitRestrictions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("explicitRestrictions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReportingTaskDefinition) contextValidatePropertyDescriptors(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.PropertyDescriptors {

		if val, ok := m.PropertyDescriptors[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ReportingTaskDefinition) contextValidateProvidedAPIImplementations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ProvidedAPIImplementations); i++ {

		if m.ProvidedAPIImplementations[i] != nil {

			if swag.IsZero(m.ProvidedAPIImplementations[i]) { // not required
				return nil
			}

			if err := m.ProvidedAPIImplementations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("providedApiImplementations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("providedApiImplementations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReportingTaskDefinition) contextValidateStateful(ctx context.Context, formats strfmt.Registry) error {

	if m.Stateful != nil {

		if swag.IsZero(m.Stateful) { // not required
			return nil
		}

		if err := m.Stateful.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stateful")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stateful")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReportingTaskDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportingTaskDefinition) UnmarshalBinary(b []byte) error {
	var res ReportingTaskDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
