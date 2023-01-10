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
)

// ControllerBulletinsEntity controller bulletins entity
//
// swagger:model ControllerBulletinsEntity
type ControllerBulletinsEntity struct {

	// System level bulletins to be reported to the user.
	Bulletins []*BulletinEntity `json:"bulletins"`

	// Controller service bulletins to be reported to the user.
	ControllerServiceBulletins []*BulletinEntity `json:"controllerServiceBulletins"`

	// Flow registry client bulletins to be reported to the user.
	FlowRegistryClientBulletins []*BulletinEntity `json:"flowRegistryClientBulletins"`

	// Parameter provider bulletins to be reported to the user.
	ParameterProviderBulletins []*BulletinEntity `json:"parameterProviderBulletins"`

	// Reporting task bulletins to be reported to the user.
	ReportingTaskBulletins []*BulletinEntity `json:"reportingTaskBulletins"`
}

// Validate validates this controller bulletins entity
func (m *ControllerBulletinsEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBulletins(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateControllerServiceBulletins(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlowRegistryClientBulletins(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameterProviderBulletins(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReportingTaskBulletins(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ControllerBulletinsEntity) validateBulletins(formats strfmt.Registry) error {
	if swag.IsZero(m.Bulletins) { // not required
		return nil
	}

	for i := 0; i < len(m.Bulletins); i++ {
		if swag.IsZero(m.Bulletins[i]) { // not required
			continue
		}

		if m.Bulletins[i] != nil {
			if err := m.Bulletins[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bulletins" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("bulletins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ControllerBulletinsEntity) validateControllerServiceBulletins(formats strfmt.Registry) error {
	if swag.IsZero(m.ControllerServiceBulletins) { // not required
		return nil
	}

	for i := 0; i < len(m.ControllerServiceBulletins); i++ {
		if swag.IsZero(m.ControllerServiceBulletins[i]) { // not required
			continue
		}

		if m.ControllerServiceBulletins[i] != nil {
			if err := m.ControllerServiceBulletins[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("controllerServiceBulletins" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("controllerServiceBulletins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ControllerBulletinsEntity) validateFlowRegistryClientBulletins(formats strfmt.Registry) error {
	if swag.IsZero(m.FlowRegistryClientBulletins) { // not required
		return nil
	}

	for i := 0; i < len(m.FlowRegistryClientBulletins); i++ {
		if swag.IsZero(m.FlowRegistryClientBulletins[i]) { // not required
			continue
		}

		if m.FlowRegistryClientBulletins[i] != nil {
			if err := m.FlowRegistryClientBulletins[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("flowRegistryClientBulletins" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("flowRegistryClientBulletins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ControllerBulletinsEntity) validateParameterProviderBulletins(formats strfmt.Registry) error {
	if swag.IsZero(m.ParameterProviderBulletins) { // not required
		return nil
	}

	for i := 0; i < len(m.ParameterProviderBulletins); i++ {
		if swag.IsZero(m.ParameterProviderBulletins[i]) { // not required
			continue
		}

		if m.ParameterProviderBulletins[i] != nil {
			if err := m.ParameterProviderBulletins[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameterProviderBulletins" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parameterProviderBulletins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ControllerBulletinsEntity) validateReportingTaskBulletins(formats strfmt.Registry) error {
	if swag.IsZero(m.ReportingTaskBulletins) { // not required
		return nil
	}

	for i := 0; i < len(m.ReportingTaskBulletins); i++ {
		if swag.IsZero(m.ReportingTaskBulletins[i]) { // not required
			continue
		}

		if m.ReportingTaskBulletins[i] != nil {
			if err := m.ReportingTaskBulletins[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reportingTaskBulletins" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("reportingTaskBulletins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this controller bulletins entity based on the context it is used
func (m *ControllerBulletinsEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBulletins(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateControllerServiceBulletins(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFlowRegistryClientBulletins(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParameterProviderBulletins(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReportingTaskBulletins(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ControllerBulletinsEntity) contextValidateBulletins(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Bulletins); i++ {

		if m.Bulletins[i] != nil {
			if err := m.Bulletins[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bulletins" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("bulletins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ControllerBulletinsEntity) contextValidateControllerServiceBulletins(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ControllerServiceBulletins); i++ {

		if m.ControllerServiceBulletins[i] != nil {
			if err := m.ControllerServiceBulletins[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("controllerServiceBulletins" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("controllerServiceBulletins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ControllerBulletinsEntity) contextValidateFlowRegistryClientBulletins(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FlowRegistryClientBulletins); i++ {

		if m.FlowRegistryClientBulletins[i] != nil {
			if err := m.FlowRegistryClientBulletins[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("flowRegistryClientBulletins" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("flowRegistryClientBulletins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ControllerBulletinsEntity) contextValidateParameterProviderBulletins(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ParameterProviderBulletins); i++ {

		if m.ParameterProviderBulletins[i] != nil {
			if err := m.ParameterProviderBulletins[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameterProviderBulletins" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parameterProviderBulletins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ControllerBulletinsEntity) contextValidateReportingTaskBulletins(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ReportingTaskBulletins); i++ {

		if m.ReportingTaskBulletins[i] != nil {
			if err := m.ReportingTaskBulletins[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reportingTaskBulletins" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("reportingTaskBulletins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ControllerBulletinsEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ControllerBulletinsEntity) UnmarshalBinary(b []byte) error {
	var res ControllerBulletinsEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
