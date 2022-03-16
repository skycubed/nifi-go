// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ControllerServiceReferencingComponentDTO controller service referencing component d t o
//
// swagger:model ControllerServiceReferencingComponentDTO
type ControllerServiceReferencingComponentDTO struct {

	// The number of active threads for the referencing component.
	ActiveThreadCount int32 `json:"activeThreadCount,omitempty"`

	// The descriptors for the component properties.
	Descriptors map[string]PropertyDescriptorDTO `json:"descriptors,omitempty"`

	// The group id for the component referencing a controller service. If this component is another controller service or a reporting task, this field is blank.
	GroupID string `json:"groupId,omitempty"`

	// The id of the component referencing a controller service.
	ID string `json:"id,omitempty"`

	// The name of the component referencing a controller service.
	Name string `json:"name,omitempty"`

	// The properties for the component.
	Properties map[string]string `json:"properties,omitempty"`

	// If the referencing component represents a controller service, this indicates whether it has already been represented in this hierarchy.
	ReferenceCycle bool `json:"referenceCycle,omitempty"`

	// The type of reference this is.
	// Enum: [Processor ControllerService ReportingTask]
	ReferenceType string `json:"referenceType,omitempty"`

	// If the referencing component represents a controller service, these are the components that reference it.
	// Unique: true
	ReferencingComponents []*ControllerServiceReferencingComponentEntity `json:"referencingComponents"`

	// The scheduled state of a processor or reporting task referencing a controller service. If this component is another controller service, this field represents the controller service state.
	State string `json:"state,omitempty"`

	// The type of the component referencing a controller service in simple Java class name format without package name.
	Type string `json:"type,omitempty"`

	// The validation errors for the component.
	ValidationErrors []string `json:"validationErrors"`
}

// Validate validates this controller service referencing component d t o
func (m *ControllerServiceReferencingComponentDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescriptors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReferenceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReferencingComponents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ControllerServiceReferencingComponentDTO) validateDescriptors(formats strfmt.Registry) error {
	if swag.IsZero(m.Descriptors) { // not required
		return nil
	}

	for k := range m.Descriptors {

		if err := validate.Required("descriptors"+"."+k, "body", m.Descriptors[k]); err != nil {
			return err
		}
		if val, ok := m.Descriptors[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("descriptors" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("descriptors" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

var controllerServiceReferencingComponentDTOTypeReferenceTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Processor","ControllerService","ReportingTask"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		controllerServiceReferencingComponentDTOTypeReferenceTypePropEnum = append(controllerServiceReferencingComponentDTOTypeReferenceTypePropEnum, v)
	}
}

const (

	// ControllerServiceReferencingComponentDTOReferenceTypeProcessor captures enum value "Processor"
	ControllerServiceReferencingComponentDTOReferenceTypeProcessor string = "Processor"

	// ControllerServiceReferencingComponentDTOReferenceTypeControllerService captures enum value "ControllerService"
	ControllerServiceReferencingComponentDTOReferenceTypeControllerService string = "ControllerService"

	// ControllerServiceReferencingComponentDTOReferenceTypeReportingTask captures enum value "ReportingTask"
	ControllerServiceReferencingComponentDTOReferenceTypeReportingTask string = "ReportingTask"
)

// prop value enum
func (m *ControllerServiceReferencingComponentDTO) validateReferenceTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, controllerServiceReferencingComponentDTOTypeReferenceTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ControllerServiceReferencingComponentDTO) validateReferenceType(formats strfmt.Registry) error {
	if swag.IsZero(m.ReferenceType) { // not required
		return nil
	}

	// value enum
	if err := m.validateReferenceTypeEnum("referenceType", "body", m.ReferenceType); err != nil {
		return err
	}

	return nil
}

func (m *ControllerServiceReferencingComponentDTO) validateReferencingComponents(formats strfmt.Registry) error {
	if swag.IsZero(m.ReferencingComponents) { // not required
		return nil
	}

	if err := validate.UniqueItems("referencingComponents", "body", m.ReferencingComponents); err != nil {
		return err
	}

	for i := 0; i < len(m.ReferencingComponents); i++ {
		if swag.IsZero(m.ReferencingComponents[i]) { // not required
			continue
		}

		if m.ReferencingComponents[i] != nil {
			if err := m.ReferencingComponents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("referencingComponents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("referencingComponents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this controller service referencing component d t o based on the context it is used
func (m *ControllerServiceReferencingComponentDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDescriptors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReferencingComponents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ControllerServiceReferencingComponentDTO) contextValidateDescriptors(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Descriptors {

		if val, ok := m.Descriptors[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ControllerServiceReferencingComponentDTO) contextValidateReferencingComponents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ReferencingComponents); i++ {

		if m.ReferencingComponents[i] != nil {
			if err := m.ReferencingComponents[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("referencingComponents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("referencingComponents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ControllerServiceReferencingComponentDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ControllerServiceReferencingComponentDTO) UnmarshalBinary(b []byte) error {
	var res ControllerServiceReferencingComponentDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
