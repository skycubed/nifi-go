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

// CurrentUser current user
//
// swagger:model CurrentUser
type CurrentUser struct {

	// Indicates if the current user is anonymous
	// Read Only: true
	Anonymous *bool `json:"anonymous,omitempty"`

	// The identity of the current user
	// Read Only: true
	Identity string `json:"identity,omitempty"`

	// Indicates if the NiFi Registry instance supports logging in
	LoginSupported bool `json:"loginSupported,omitempty"`

	// Indicates if the NiFi Registry instance supports logging in with an OIDC provider
	OidcloginSupported bool `json:"oidcloginSupported,omitempty"`

	// The access that the current user has to top level resources
	// Read Only: true
	ResourcePermissions *ResourcePermissions `json:"resourcePermissions,omitempty"`
}

// Validate validates this current user
func (m *CurrentUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourcePermissions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CurrentUser) validateResourcePermissions(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourcePermissions) { // not required
		return nil
	}

	if m.ResourcePermissions != nil {
		if err := m.ResourcePermissions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourcePermissions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourcePermissions")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this current user based on the context it is used
func (m *CurrentUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAnonymous(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIdentity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResourcePermissions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CurrentUser) contextValidateAnonymous(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "anonymous", "body", m.Anonymous); err != nil {
		return err
	}

	return nil
}

func (m *CurrentUser) contextValidateIdentity(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "identity", "body", string(m.Identity)); err != nil {
		return err
	}

	return nil
}

func (m *CurrentUser) contextValidateResourcePermissions(ctx context.Context, formats strfmt.Registry) error {

	if m.ResourcePermissions != nil {

		if swag.IsZero(m.ResourcePermissions) { // not required
			return nil
		}

		if err := m.ResourcePermissions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourcePermissions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourcePermissions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CurrentUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CurrentUser) UnmarshalBinary(b []byte) error {
	var res CurrentUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
