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

// ClusteSummaryEntity cluste summary entity
//
// swagger:model ClusteSummaryEntity
type ClusteSummaryEntity struct {

	// cluster summary
	ClusterSummary *ClusterSummaryDTO `json:"clusterSummary,omitempty"`
}

// Validate validates this cluste summary entity
func (m *ClusteSummaryEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterSummary(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusteSummaryEntity) validateClusterSummary(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterSummary) { // not required
		return nil
	}

	if m.ClusterSummary != nil {
		if err := m.ClusterSummary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterSummary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusterSummary")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluste summary entity based on the context it is used
func (m *ClusteSummaryEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusterSummary(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusteSummaryEntity) contextValidateClusterSummary(ctx context.Context, formats strfmt.Registry) error {

	if m.ClusterSummary != nil {
		if err := m.ClusterSummary.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterSummary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusterSummary")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusteSummaryEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusteSummaryEntity) UnmarshalBinary(b []byte) error {
	var res ClusteSummaryEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
