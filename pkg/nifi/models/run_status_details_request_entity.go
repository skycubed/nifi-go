// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RunStatusDetailsRequestEntity run status details request entity
//
// swagger:model RunStatusDetailsRequestEntity
type RunStatusDetailsRequestEntity struct {

	// The IDs of all processors whose run status details should be provided
	// Unique: true
	ProcessorIds []string `json:"processorIds"`
}

// Validate validates this run status details request entity
func (m *RunStatusDetailsRequestEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProcessorIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RunStatusDetailsRequestEntity) validateProcessorIds(formats strfmt.Registry) error {
	if swag.IsZero(m.ProcessorIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("processorIds", "body", m.ProcessorIds); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this run status details request entity based on context it is used
func (m *RunStatusDetailsRequestEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RunStatusDetailsRequestEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RunStatusDetailsRequestEntity) UnmarshalBinary(b []byte) error {
	var res RunStatusDetailsRequestEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
