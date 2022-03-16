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

// Property property
//
// swagger:model Property
type Property struct {

	// The allowable values for this property
	AllowableValues []*AllowableValue `json:"allowableValues" xml:"allowableValues"`

	// The controller service required by this property, or null if none is required
	ControllerServiceDefinition *ControllerServiceDefinition `json:"controllerServiceDefinition,omitempty"`

	// The default value
	DefaultValue string `json:"defaultValue,omitempty"`

	// The description
	Description string `json:"description,omitempty"`

	// The display name
	DisplayName string `json:"displayName,omitempty"`

	// Whether or not the processor is dynamic
	Dynamic bool `json:"dynamic,omitempty"`

	// Whether or not the processor dynamically modifies the classpath
	DynamicallyModifiesClasspath bool `json:"dynamicallyModifiesClasspath,omitempty"`

	// The scope of expression language support
	// Enum: [NONE VARIABLE_REGISTRY FLOWFILE_ATTRIBUTES]
	ExpressionLanguageScope string `json:"expressionLanguageScope,omitempty"`

	// Whether or not expression language is supported
	ExpressionLanguageSupported bool `json:"expressionLanguageSupported,omitempty"`

	// The name of the property
	Name string `json:"name,omitempty"`

	// Whether or not the property is required
	Required bool `json:"required,omitempty"`

	// Whether or not the property is sensitive
	Sensitive bool `json:"sensitive,omitempty"`
}

// Validate validates this property
func (m *Property) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowableValues(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateControllerServiceDefinition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpressionLanguageScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Property) validateAllowableValues(formats strfmt.Registry) error {
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

func (m *Property) validateControllerServiceDefinition(formats strfmt.Registry) error {
	if swag.IsZero(m.ControllerServiceDefinition) { // not required
		return nil
	}

	if m.ControllerServiceDefinition != nil {
		if err := m.ControllerServiceDefinition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("controllerServiceDefinition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("controllerServiceDefinition")
			}
			return err
		}
	}

	return nil
}

var propertyTypeExpressionLanguageScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NONE","VARIABLE_REGISTRY","FLOWFILE_ATTRIBUTES"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		propertyTypeExpressionLanguageScopePropEnum = append(propertyTypeExpressionLanguageScopePropEnum, v)
	}
}

const (

	// PropertyExpressionLanguageScopeNONE captures enum value "NONE"
	PropertyExpressionLanguageScopeNONE string = "NONE"

	// PropertyExpressionLanguageScopeVARIABLEREGISTRY captures enum value "VARIABLE_REGISTRY"
	PropertyExpressionLanguageScopeVARIABLEREGISTRY string = "VARIABLE_REGISTRY"

	// PropertyExpressionLanguageScopeFLOWFILEATTRIBUTES captures enum value "FLOWFILE_ATTRIBUTES"
	PropertyExpressionLanguageScopeFLOWFILEATTRIBUTES string = "FLOWFILE_ATTRIBUTES"
)

// prop value enum
func (m *Property) validateExpressionLanguageScopeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, propertyTypeExpressionLanguageScopePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Property) validateExpressionLanguageScope(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpressionLanguageScope) { // not required
		return nil
	}

	// value enum
	if err := m.validateExpressionLanguageScopeEnum("expressionLanguageScope", "body", m.ExpressionLanguageScope); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this property based on the context it is used
func (m *Property) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAllowableValues(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateControllerServiceDefinition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Property) contextValidateAllowableValues(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AllowableValues); i++ {

		if m.AllowableValues[i] != nil {
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

func (m *Property) contextValidateControllerServiceDefinition(ctx context.Context, formats strfmt.Registry) error {

	if m.ControllerServiceDefinition != nil {
		if err := m.ControllerServiceDefinition.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("controllerServiceDefinition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("controllerServiceDefinition")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Property) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Property) UnmarshalBinary(b []byte) error {
	var res Property
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}