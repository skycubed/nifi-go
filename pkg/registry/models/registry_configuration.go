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

// RegistryConfiguration registry configuration
//
// swagger:model RegistryConfiguration
type RegistryConfiguration struct {

	// Whether this NiFi Registry supports a configurable authorizer.
	// Read Only: true
	SupportsConfigurableAuthorizer *bool `json:"supportsConfigurableAuthorizer,omitempty"`

	// Whether this NiFi Registry supports configurable users and groups.
	// Read Only: true
	SupportsConfigurableUsersAndGroups *bool `json:"supportsConfigurableUsersAndGroups,omitempty"`

	// Whether this NiFi Registry supports a managed authorizer. Managed authorizers can visualize users, groups, and policies in the UI.
	// Read Only: true
	SupportsManagedAuthorizer *bool `json:"supportsManagedAuthorizer,omitempty"`
}

// Validate validates this registry configuration
func (m *RegistryConfiguration) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this registry configuration based on the context it is used
func (m *RegistryConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSupportsConfigurableAuthorizer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSupportsConfigurableUsersAndGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSupportsManagedAuthorizer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegistryConfiguration) contextValidateSupportsConfigurableAuthorizer(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "supportsConfigurableAuthorizer", "body", m.SupportsConfigurableAuthorizer); err != nil {
		return err
	}

	return nil
}

func (m *RegistryConfiguration) contextValidateSupportsConfigurableUsersAndGroups(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "supportsConfigurableUsersAndGroups", "body", m.SupportsConfigurableUsersAndGroups); err != nil {
		return err
	}

	return nil
}

func (m *RegistryConfiguration) contextValidateSupportsManagedAuthorizer(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "supportsManagedAuthorizer", "body", m.SupportsManagedAuthorizer); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RegistryConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegistryConfiguration) UnmarshalBinary(b []byte) error {
	var res RegistryConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
