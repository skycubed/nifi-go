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

// VersionedFlowEntity versioned flow entity
//
// swagger:model VersionedFlowEntity
type VersionedFlowEntity struct {

	// The versioned flow
	VersionedFlow *VersionedFlowDTO `json:"versionedFlow,omitempty"`
}

// Validate validates this versioned flow entity
func (m *VersionedFlowEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVersionedFlow(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VersionedFlowEntity) validateVersionedFlow(formats strfmt.Registry) error {
	if swag.IsZero(m.VersionedFlow) { // not required
		return nil
	}

	if m.VersionedFlow != nil {
		if err := m.VersionedFlow.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("versionedFlow")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("versionedFlow")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this versioned flow entity based on the context it is used
func (m *VersionedFlowEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVersionedFlow(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VersionedFlowEntity) contextValidateVersionedFlow(ctx context.Context, formats strfmt.Registry) error {

	if m.VersionedFlow != nil {

		if swag.IsZero(m.VersionedFlow) { // not required
			return nil
		}

		if err := m.VersionedFlow.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("versionedFlow")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("versionedFlow")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VersionedFlowEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VersionedFlowEntity) UnmarshalBinary(b []byte) error {
	var res VersionedFlowEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
