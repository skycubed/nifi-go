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

// TemplatesEntity templates entity
//
// swagger:model TemplatesEntity
type TemplatesEntity struct {

	// When this content was generated.
	Generated string `json:"generated,omitempty"`

	// templates
	// Unique: true
	Templates []*TemplateEntity `json:"templates"`
}

// Validate validates this templates entity
func (m *TemplatesEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTemplates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TemplatesEntity) validateTemplates(formats strfmt.Registry) error {
	if swag.IsZero(m.Templates) { // not required
		return nil
	}

	if err := validate.UniqueItems("templates", "body", m.Templates); err != nil {
		return err
	}

	for i := 0; i < len(m.Templates); i++ {
		if swag.IsZero(m.Templates[i]) { // not required
			continue
		}

		if m.Templates[i] != nil {
			if err := m.Templates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("templates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("templates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this templates entity based on the context it is used
func (m *TemplatesEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTemplates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TemplatesEntity) contextValidateTemplates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Templates); i++ {

		if m.Templates[i] != nil {
			if err := m.Templates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("templates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("templates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TemplatesEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TemplatesEntity) UnmarshalBinary(b []byte) error {
	var res TemplatesEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}