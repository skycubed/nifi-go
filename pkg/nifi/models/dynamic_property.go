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

// DynamicProperty dynamic property
//
// swagger:model DynamicProperty
type DynamicProperty struct {

	// The description of the dynamic property
	Description string `json:"description,omitempty"`

	// The scope of the expression language support
	// Enum: [NONE VARIABLE_REGISTRY FLOWFILE_ATTRIBUTES]
	ExpressionLanguageScope string `json:"expressionLanguageScope,omitempty"`

	// The description of the dynamic property name
	Name string `json:"name,omitempty"`

	// The description of the dynamic property value
	Value string `json:"value,omitempty"`
}

// Validate validates this dynamic property
func (m *DynamicProperty) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpressionLanguageScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var dynamicPropertyTypeExpressionLanguageScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NONE","VARIABLE_REGISTRY","FLOWFILE_ATTRIBUTES"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dynamicPropertyTypeExpressionLanguageScopePropEnum = append(dynamicPropertyTypeExpressionLanguageScopePropEnum, v)
	}
}

const (

	// DynamicPropertyExpressionLanguageScopeNONE captures enum value "NONE"
	DynamicPropertyExpressionLanguageScopeNONE string = "NONE"

	// DynamicPropertyExpressionLanguageScopeVARIABLEREGISTRY captures enum value "VARIABLE_REGISTRY"
	DynamicPropertyExpressionLanguageScopeVARIABLEREGISTRY string = "VARIABLE_REGISTRY"

	// DynamicPropertyExpressionLanguageScopeFLOWFILEATTRIBUTES captures enum value "FLOWFILE_ATTRIBUTES"
	DynamicPropertyExpressionLanguageScopeFLOWFILEATTRIBUTES string = "FLOWFILE_ATTRIBUTES"
)

// prop value enum
func (m *DynamicProperty) validateExpressionLanguageScopeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dynamicPropertyTypeExpressionLanguageScopePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DynamicProperty) validateExpressionLanguageScope(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpressionLanguageScope) { // not required
		return nil
	}

	// value enum
	if err := m.validateExpressionLanguageScopeEnum("expressionLanguageScope", "body", m.ExpressionLanguageScope); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this dynamic property based on context it is used
func (m *DynamicProperty) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DynamicProperty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DynamicProperty) UnmarshalBinary(b []byte) error {
	var res DynamicProperty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}