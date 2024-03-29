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

// BucketEntity bucket entity
//
// swagger:model BucketEntity
type BucketEntity struct {

	// bucket
	Bucket *BucketDTO `json:"bucket,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// permissions
	Permissions *PermissionsDTO `json:"permissions,omitempty"`
}

// Validate validates this bucket entity
func (m *BucketEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBucket(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BucketEntity) validateBucket(formats strfmt.Registry) error {
	if swag.IsZero(m.Bucket) { // not required
		return nil
	}

	if m.Bucket != nil {
		if err := m.Bucket.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bucket")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bucket")
			}
			return err
		}
	}

	return nil
}

func (m *BucketEntity) validatePermissions(formats strfmt.Registry) error {
	if swag.IsZero(m.Permissions) { // not required
		return nil
	}

	if m.Permissions != nil {
		if err := m.Permissions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("permissions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("permissions")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this bucket entity based on the context it is used
func (m *BucketEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBucket(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePermissions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BucketEntity) contextValidateBucket(ctx context.Context, formats strfmt.Registry) error {

	if m.Bucket != nil {

		if swag.IsZero(m.Bucket) { // not required
			return nil
		}

		if err := m.Bucket.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bucket")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bucket")
			}
			return err
		}
	}

	return nil
}

func (m *BucketEntity) contextValidatePermissions(ctx context.Context, formats strfmt.Registry) error {

	if m.Permissions != nil {

		if swag.IsZero(m.Permissions) { // not required
			return nil
		}

		if err := m.Permissions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("permissions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("permissions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BucketEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BucketEntity) UnmarshalBinary(b []byte) error {
	var res BucketEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
