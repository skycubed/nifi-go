// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// JmxMetricsResultDTO jmx metrics result d t o
//
// swagger:model JmxMetricsResultDTO
type JmxMetricsResultDTO struct {

	// The attribute name of the metrics bean's attribute.
	AttributeName string `json:"attributeName,omitempty"`

	// The attribute value of the the metrics bean's attribute
	AttributeValue interface{} `json:"attributeValue,omitempty"`

	// The bean name of the metrics bean.
	BeanName string `json:"beanName,omitempty"`
}

// Validate validates this jmx metrics result d t o
func (m *JmxMetricsResultDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this jmx metrics result d t o based on context it is used
func (m *JmxMetricsResultDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *JmxMetricsResultDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JmxMetricsResultDTO) UnmarshalBinary(b []byte) error {
	var res JmxMetricsResultDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}