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

// ComponentValidationResultEntity component validation result entity
//
// swagger:model ComponentValidationResultEntity
type ComponentValidationResultEntity struct {

	// The bulletins for this component.
	Bulletins []*BulletinEntity `json:"bulletins"`

	// component
	Component *ComponentValidationResultDTO `json:"component,omitempty"`

	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged bool `json:"disconnectedNodeAcknowledged,omitempty"`

	// The id of the component.
	ID string `json:"id,omitempty"`

	// The permissions for this component.
	Permissions *PermissionsDTO `json:"permissions,omitempty"`

	// The position of this component in the UI if applicable.
	Position *PositionDTO `json:"position,omitempty"`

	// The revision for this request/response. The revision is required for any mutable flow requests and is included in all responses.
	Revision *RevisionDTO `json:"revision,omitempty"`

	// The URI for futures requests to the component.
	URI string `json:"uri,omitempty"`
}

// Validate validates this component validation result entity
func (m *ComponentValidationResultEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBulletins(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComponent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevision(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComponentValidationResultEntity) validateBulletins(formats strfmt.Registry) error {
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

func (m *ComponentValidationResultEntity) validateComponent(formats strfmt.Registry) error {
	if swag.IsZero(m.Component) { // not required
		return nil
	}

	if m.Component != nil {
		if err := m.Component.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("component")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("component")
			}
			return err
		}
	}

	return nil
}

func (m *ComponentValidationResultEntity) validatePermissions(formats strfmt.Registry) error {
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

func (m *ComponentValidationResultEntity) validatePosition(formats strfmt.Registry) error {
	if swag.IsZero(m.Position) { // not required
		return nil
	}

	if m.Position != nil {
		if err := m.Position.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("position")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("position")
			}
			return err
		}
	}

	return nil
}

func (m *ComponentValidationResultEntity) validateRevision(formats strfmt.Registry) error {
	if swag.IsZero(m.Revision) { // not required
		return nil
	}

	if m.Revision != nil {
		if err := m.Revision.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("revision")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("revision")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this component validation result entity based on the context it is used
func (m *ComponentValidationResultEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBulletins(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateComponent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePermissions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePosition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRevision(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComponentValidationResultEntity) contextValidateBulletins(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ComponentValidationResultEntity) contextValidateComponent(ctx context.Context, formats strfmt.Registry) error {

	if m.Component != nil {
		if err := m.Component.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("component")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("component")
			}
			return err
		}
	}

	return nil
}

func (m *ComponentValidationResultEntity) contextValidatePermissions(ctx context.Context, formats strfmt.Registry) error {

	if m.Permissions != nil {
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

func (m *ComponentValidationResultEntity) contextValidatePosition(ctx context.Context, formats strfmt.Registry) error {

	if m.Position != nil {
		if err := m.Position.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("position")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("position")
			}
			return err
		}
	}

	return nil
}

func (m *ComponentValidationResultEntity) contextValidateRevision(ctx context.Context, formats strfmt.Registry) error {

	if m.Revision != nil {
		if err := m.Revision.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("revision")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("revision")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ComponentValidationResultEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComponentValidationResultEntity) UnmarshalBinary(b []byte) error {
	var res ComponentValidationResultEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
