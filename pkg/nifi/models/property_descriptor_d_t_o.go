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

// PropertyDescriptorDTO property descriptor d t o
//
// swagger:model PropertyDescriptorDTO
type PropertyDescriptorDTO struct {

	// Allowable values for the property. If empty then the allowed values are not constrained.
	AllowableValues []*AllowableValueEntity `json:"allowableValues"`

	// The default value for the property.
	DefaultValue string `json:"defaultValue,omitempty"`

	// A list of dependencies that must be met in order for this Property to be relevant. If any of these dependencies is not met, the property described by this Property Descriptor is not relevant.
	Dependencies []*PropertyDependencyDTO `json:"dependencies"`

	// The description for the property. Used to relay additional details to a user or provide a mechanism of documenting intent.
	Description string `json:"description,omitempty"`

	// The human readable name for the property.
	DisplayName string `json:"displayName,omitempty"`

	// Whether the property is dynamic (user-defined).
	Dynamic bool `json:"dynamic,omitempty"`

	// Scope of the Expression Language evaluation for the property.
	ExpressionLanguageScope string `json:"expressionLanguageScope,omitempty"`

	// If the property identifies a controller service this returns the fully qualified type.
	IdentifiesControllerService string `json:"identifiesControllerService,omitempty"`

	// If the property identifies a controller service this returns the bundle of the type, null otherwise.
	IdentifiesControllerServiceBundle *BundleDTO `json:"identifiesControllerServiceBundle,omitempty"`

	// The name for the property.
	Name string `json:"name,omitempty"`

	// Whether the property is required.
	Required bool `json:"required,omitempty"`

	// Whether the property is sensitive and protected whenever stored or represented.
	Sensitive bool `json:"sensitive,omitempty"`

	// Whether the property supports expression language.
	SupportsEl bool `json:"supportsEl,omitempty"`
}

// Validate validates this property descriptor d t o
func (m *PropertyDescriptorDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowableValues(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDependencies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentifiesControllerServiceBundle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PropertyDescriptorDTO) validateAllowableValues(formats strfmt.Registry) error {
	if swag.IsZero(m.AllowableValues) { // not required
		return nil
	}

	for i := 0; i < len(m.AllowableValues); i++ {
		if swag.IsZero(m.AllowableValues[i]) { // not required
			continue
		}

		if m.AllowableValues[i] != nil {
			if err := m.AllowableValues[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allowableValues" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("allowableValues" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PropertyDescriptorDTO) validateDependencies(formats strfmt.Registry) error {
	if swag.IsZero(m.Dependencies) { // not required
		return nil
	}

	for i := 0; i < len(m.Dependencies); i++ {
		if swag.IsZero(m.Dependencies[i]) { // not required
			continue
		}

		if m.Dependencies[i] != nil {
			if err := m.Dependencies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dependencies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dependencies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PropertyDescriptorDTO) validateIdentifiesControllerServiceBundle(formats strfmt.Registry) error {
	if swag.IsZero(m.IdentifiesControllerServiceBundle) { // not required
		return nil
	}

	if m.IdentifiesControllerServiceBundle != nil {
		if err := m.IdentifiesControllerServiceBundle.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identifiesControllerServiceBundle")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("identifiesControllerServiceBundle")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this property descriptor d t o based on the context it is used
func (m *PropertyDescriptorDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAllowableValues(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDependencies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIdentifiesControllerServiceBundle(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PropertyDescriptorDTO) contextValidateAllowableValues(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AllowableValues); i++ {

		if m.AllowableValues[i] != nil {

			if swag.IsZero(m.AllowableValues[i]) { // not required
				return nil
			}

			if err := m.AllowableValues[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allowableValues" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("allowableValues" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PropertyDescriptorDTO) contextValidateDependencies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Dependencies); i++ {

		if m.Dependencies[i] != nil {

			if swag.IsZero(m.Dependencies[i]) { // not required
				return nil
			}

			if err := m.Dependencies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dependencies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dependencies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PropertyDescriptorDTO) contextValidateIdentifiesControllerServiceBundle(ctx context.Context, formats strfmt.Registry) error {

	if m.IdentifiesControllerServiceBundle != nil {

		if swag.IsZero(m.IdentifiesControllerServiceBundle) { // not required
			return nil
		}

		if err := m.IdentifiesControllerServiceBundle.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identifiesControllerServiceBundle")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("identifiesControllerServiceBundle")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PropertyDescriptorDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PropertyDescriptorDTO) UnmarshalBinary(b []byte) error {
	var res PropertyDescriptorDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
