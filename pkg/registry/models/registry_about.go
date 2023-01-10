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

// RegistryAbout registry about
//
// swagger:model RegistryAbout
type RegistryAbout struct {

	// The version string for this Nifi Registry
	// Read Only: true
	RegistryAboutVersion string `json:"registryAboutVersion,omitempty"`
}

// Validate validates this registry about
func (m *RegistryAbout) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this registry about based on the context it is used
func (m *RegistryAbout) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRegistryAboutVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegistryAbout) contextValidateRegistryAboutVersion(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "registryAboutVersion", "body", string(m.RegistryAboutVersion)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RegistryAbout) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegistryAbout) UnmarshalBinary(b []byte) error {
	var res RegistryAbout
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
