// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ExtensionRepoBucket extension repo bucket
//
// swagger:model ExtensionRepoBucket
type ExtensionRepoBucket struct {

	// The name of the bucket
	BucketName string `json:"bucketName,omitempty"`

	// An WebLink to this entity.
	// Read Only: true
	Link *JaxbLink `json:"link,omitempty"`
}

// Validate validates this extension repo bucket
func (m *ExtensionRepoBucket) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLink(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExtensionRepoBucket) validateLink(formats strfmt.Registry) error {
	if swag.IsZero(m.Link) { // not required
		return nil
	}

	if m.Link != nil {
		if err := m.Link.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("link")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("link")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this extension repo bucket based on the context it is used
func (m *ExtensionRepoBucket) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLink(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExtensionRepoBucket) contextValidateLink(ctx context.Context, formats strfmt.Registry) error {

	if m.Link != nil {

		if swag.IsZero(m.Link) { // not required
			return nil
		}

		if err := m.Link.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("link")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("link")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExtensionRepoBucket) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExtensionRepoBucket) UnmarshalBinary(b []byte) error {
	var res ExtensionRepoBucket
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
