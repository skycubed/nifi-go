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

// FlowComparisonEntity flow comparison entity
//
// swagger:model FlowComparisonEntity
type FlowComparisonEntity struct {

	// The list of differences for each component in the flow that is not the same between the two flows
	// Unique: true
	ComponentDifferences []*ComponentDifferenceDTO `json:"componentDifferences"`
}

// Validate validates this flow comparison entity
func (m *FlowComparisonEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComponentDifferences(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FlowComparisonEntity) validateComponentDifferences(formats strfmt.Registry) error {
	if swag.IsZero(m.ComponentDifferences) { // not required
		return nil
	}

	if err := validate.UniqueItems("componentDifferences", "body", m.ComponentDifferences); err != nil {
		return err
	}

	for i := 0; i < len(m.ComponentDifferences); i++ {
		if swag.IsZero(m.ComponentDifferences[i]) { // not required
			continue
		}

		if m.ComponentDifferences[i] != nil {
			if err := m.ComponentDifferences[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("componentDifferences" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("componentDifferences" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this flow comparison entity based on the context it is used
func (m *FlowComparisonEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComponentDifferences(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FlowComparisonEntity) contextValidateComponentDifferences(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ComponentDifferences); i++ {

		if m.ComponentDifferences[i] != nil {

			if swag.IsZero(m.ComponentDifferences[i]) { // not required
				return nil
			}

			if err := m.ComponentDifferences[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("componentDifferences" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("componentDifferences" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FlowComparisonEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlowComparisonEntity) UnmarshalBinary(b []byte) error {
	var res FlowComparisonEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
