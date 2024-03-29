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

// ReportingTasksEntity reporting tasks entity
//
// swagger:model ReportingTasksEntity
type ReportingTasksEntity struct {

	// reporting tasks
	// Unique: true
	ReportingTasks []*ReportingTaskEntity `json:"reportingTasks"`
}

// Validate validates this reporting tasks entity
func (m *ReportingTasksEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReportingTasks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportingTasksEntity) validateReportingTasks(formats strfmt.Registry) error {
	if swag.IsZero(m.ReportingTasks) { // not required
		return nil
	}

	if err := validate.UniqueItems("reportingTasks", "body", m.ReportingTasks); err != nil {
		return err
	}

	for i := 0; i < len(m.ReportingTasks); i++ {
		if swag.IsZero(m.ReportingTasks[i]) { // not required
			continue
		}

		if m.ReportingTasks[i] != nil {
			if err := m.ReportingTasks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reportingTasks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("reportingTasks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this reporting tasks entity based on the context it is used
func (m *ReportingTasksEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateReportingTasks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportingTasksEntity) contextValidateReportingTasks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ReportingTasks); i++ {

		if m.ReportingTasks[i] != nil {

			if swag.IsZero(m.ReportingTasks[i]) { // not required
				return nil
			}

			if err := m.ReportingTasks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reportingTasks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("reportingTasks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReportingTasksEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportingTasksEntity) UnmarshalBinary(b []byte) error {
	var res ReportingTasksEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
