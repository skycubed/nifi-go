// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PropertyDescriptor property descriptor
//
// swagger:model PropertyDescriptor
type PropertyDescriptor struct {

	// A list of the allowable values for the property
	AllowableValues []*PropertyAllowableValue `json:"allowableValues"`

	// The default value if a user-set value is not specified
	DefaultValue string `json:"defaultValue,omitempty"`

	// The dependencies that this property has on other properties
	Dependencies []*PropertyDependency `json:"dependencies"`

	// The description of what the property does
	Description string `json:"description,omitempty"`

	// The display name of the property key, if different from the name
	DisplayName string `json:"displayName,omitempty"`

	// Whether or not the descriptor is for a dynamically added property
	Dynamic bool `json:"dynamic,omitempty"`

	// The scope of expression language supported by this property
	// Enum: [NONE VARIABLE_REGISTRY FLOWFILE_ATTRIBUTES]
	ExpressionLanguageScope string `json:"expressionLanguageScope,omitempty"`

	// The description of the expression language scope supported by this property
	// Read Only: true
	ExpressionLanguageScopeDescription string `json:"expressionLanguageScopeDescription,omitempty"`

	// The name of the property key
	// Required: true
	Name *string `json:"name"`

	// Whether or not  the property is required for the component
	Required bool `json:"required,omitempty"`

	// Indicates that this property references external resources
	ResourceDefinition *PropertyResourceDefinition `json:"resourceDefinition,omitempty"`

	// Whether or not  the value of the property is considered sensitive (e.g., passwords and keys)
	Sensitive bool `json:"sensitive,omitempty"`

	// Indicates that this property is for selecting a controller service of the specified type
	TypeProvidedByValue *DefinedType `json:"typeProvidedByValue,omitempty"`

	// A regular expression that can be used to validate the value of this property
	ValidRegex string `json:"validRegex,omitempty"`

	// Name of the validator used for this property descriptor
	Validator string `json:"validator,omitempty"`
}

// Validate validates this property descriptor
func (m *PropertyDescriptor) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowableValues(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDependencies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpressionLanguageScope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceDefinition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypeProvidedByValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PropertyDescriptor) validateAllowableValues(formats strfmt.Registry) error {
	if swag.IsZero(m.AllowableValues) { // not required
		return nil
	}

	for i := 0; i < len(m.AllowableValues); i++ {
		if swag.IsZero(m.AllowableValues[i]) { // not required
			continue
		}

		if m.AllowableValues[i] != nil {
			if err := m.AllowableValues[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allowableValues" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("allowableValues" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PropertyDescriptor) validateDependencies(formats strfmt.Registry) error {
	if swag.IsZero(m.Dependencies) { // not required
		return nil
	}

	for i := 0; i < len(m.Dependencies); i++ {
		if swag.IsZero(m.Dependencies[i]) { // not required
			continue
		}

		if m.Dependencies[i] != nil {
			if err := m.Dependencies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dependencies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dependencies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var propertyDescriptorTypeExpressionLanguageScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NONE","VARIABLE_REGISTRY","FLOWFILE_ATTRIBUTES"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		propertyDescriptorTypeExpressionLanguageScopePropEnum = append(propertyDescriptorTypeExpressionLanguageScopePropEnum, v)
	}
}

const (

	// PropertyDescriptorExpressionLanguageScopeNONE captures enum value "NONE"
	PropertyDescriptorExpressionLanguageScopeNONE string = "NONE"

	// PropertyDescriptorExpressionLanguageScopeVARIABLEREGISTRY captures enum value "VARIABLE_REGISTRY"
	PropertyDescriptorExpressionLanguageScopeVARIABLEREGISTRY string = "VARIABLE_REGISTRY"

	// PropertyDescriptorExpressionLanguageScopeFLOWFILEATTRIBUTES captures enum value "FLOWFILE_ATTRIBUTES"
	PropertyDescriptorExpressionLanguageScopeFLOWFILEATTRIBUTES string = "FLOWFILE_ATTRIBUTES"
)

// prop value enum
func (m *PropertyDescriptor) validateExpressionLanguageScopeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, propertyDescriptorTypeExpressionLanguageScopePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PropertyDescriptor) validateExpressionLanguageScope(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpressionLanguageScope) { // not required
		return nil
	}

	// value enum
	if err := m.validateExpressionLanguageScopeEnum("expressionLanguageScope", "body", m.ExpressionLanguageScope); err != nil {
		return err
	}

	return nil
}

func (m *PropertyDescriptor) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PropertyDescriptor) validateResourceDefinition(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceDefinition) { // not required
		return nil
	}

	if m.ResourceDefinition != nil {
		if err := m.ResourceDefinition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceDefinition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourceDefinition")
			}
			return err
		}
	}

	return nil
}

func (m *PropertyDescriptor) validateTypeProvidedByValue(formats strfmt.Registry) error {
	if swag.IsZero(m.TypeProvidedByValue) { // not required
		return nil
	}

	if m.TypeProvidedByValue != nil {
		if err := m.TypeProvidedByValue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("typeProvidedByValue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("typeProvidedByValue")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this property descriptor based on the context it is used
func (m *PropertyDescriptor) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAllowableValues(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDependencies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExpressionLanguageScopeDescription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResourceDefinition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTypeProvidedByValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PropertyDescriptor) contextValidateAllowableValues(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AllowableValues); i++ {

		if m.AllowableValues[i] != nil {

			if swag.IsZero(m.AllowableValues[i]) { // not required
				return nil
			}

			if err := m.AllowableValues[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allowableValues" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("allowableValues" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PropertyDescriptor) contextValidateDependencies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Dependencies); i++ {

		if m.Dependencies[i] != nil {

			if swag.IsZero(m.Dependencies[i]) { // not required
				return nil
			}

			if err := m.Dependencies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dependencies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dependencies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PropertyDescriptor) contextValidateExpressionLanguageScopeDescription(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "expressionLanguageScopeDescription", "body", string(m.ExpressionLanguageScopeDescription)); err != nil {
		return err
	}

	return nil
}

func (m *PropertyDescriptor) contextValidateResourceDefinition(ctx context.Context, formats strfmt.Registry) error {

	if m.ResourceDefinition != nil {

		if swag.IsZero(m.ResourceDefinition) { // not required
			return nil
		}

		if err := m.ResourceDefinition.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceDefinition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourceDefinition")
			}
			return err
		}
	}

	return nil
}

func (m *PropertyDescriptor) contextValidateTypeProvidedByValue(ctx context.Context, formats strfmt.Registry) error {

	if m.TypeProvidedByValue != nil {

		if swag.IsZero(m.TypeProvidedByValue) { // not required
			return nil
		}

		if err := m.TypeProvidedByValue.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("typeProvidedByValue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("typeProvidedByValue")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PropertyDescriptor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PropertyDescriptor) UnmarshalBinary(b []byte) error {
	var res PropertyDescriptor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
