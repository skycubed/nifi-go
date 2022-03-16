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

// Permissions permissions
//
// swagger:model Permissions
type Permissions struct {

	// Indicates whether the user can delete a given resource.
	// Read Only: true
	CanDelete *bool `json:"canDelete,omitempty"`

	// Indicates whether the user can read a given resource.
	// Read Only: true
	CanRead *bool `json:"canRead,omitempty"`

	// Indicates whether the user can write a given resource.
	// Read Only: true
	CanWrite *bool `json:"canWrite,omitempty"`
}

// Validate validates this permissions
func (m *Permissions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this permissions based on the context it is used
func (m *Permissions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCanDelete(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCanRead(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCanWrite(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Permissions) contextValidateCanDelete(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "canDelete", "body", m.CanDelete); err != nil {
		return err
	}

	return nil
}

func (m *Permissions) contextValidateCanRead(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "canRead", "body", m.CanRead); err != nil {
		return err
	}

	return nil
}

func (m *Permissions) contextValidateCanWrite(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "canWrite", "body", m.CanWrite); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Permissions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Permissions) UnmarshalBinary(b []byte) error {
	var res Permissions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}