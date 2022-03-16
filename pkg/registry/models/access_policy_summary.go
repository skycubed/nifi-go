// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AccessPolicySummary access policy summary
//
// swagger:model AccessPolicySummary
type AccessPolicySummary struct {

	// The action associated with this access policy.
	// Required: true
	// Enum: [read write delete]
	Action *string `json:"action"`

	// Indicates if this access policy is configurable, based on which Authorizer has been configured to manage it.
	// Read Only: true
	Configurable *bool `json:"configurable,omitempty"`

	// The id of the policy. Set by server at creation time.
	// Read Only: true
	Identifier string `json:"identifier,omitempty"`

	// The resource for this access policy.
	// Required: true
	Resource *string `json:"resource"`

	// The revision of this entity used for optimistic-locking during updates.
	// Read Only: true
	Revision *RevisionInfo `json:"revision,omitempty"`
}

// Validate validates this access policy summary
func (m *AccessPolicySummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResource(formats); err != nil {
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

var accessPolicySummaryTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["read","write","delete"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accessPolicySummaryTypeActionPropEnum = append(accessPolicySummaryTypeActionPropEnum, v)
	}
}

const (

	// AccessPolicySummaryActionRead captures enum value "read"
	AccessPolicySummaryActionRead string = "read"

	// AccessPolicySummaryActionWrite captures enum value "write"
	AccessPolicySummaryActionWrite string = "write"

	// AccessPolicySummaryActionDelete captures enum value "delete"
	AccessPolicySummaryActionDelete string = "delete"
)

// prop value enum
func (m *AccessPolicySummary) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, accessPolicySummaryTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AccessPolicySummary) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	// value enum
	if err := m.validateActionEnum("action", "body", *m.Action); err != nil {
		return err
	}

	return nil
}

func (m *AccessPolicySummary) validateResource(formats strfmt.Registry) error {

	if err := validate.Required("resource", "body", m.Resource); err != nil {
		return err
	}

	return nil
}

func (m *AccessPolicySummary) validateRevision(formats strfmt.Registry) error {
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

// ContextValidate validate this access policy summary based on the context it is used
func (m *AccessPolicySummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfigurable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIdentifier(ctx, formats); err != nil {
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

func (m *AccessPolicySummary) contextValidateConfigurable(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "configurable", "body", m.Configurable); err != nil {
		return err
	}

	return nil
}

func (m *AccessPolicySummary) contextValidateIdentifier(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "identifier", "body", string(m.Identifier)); err != nil {
		return err
	}

	return nil
}

func (m *AccessPolicySummary) contextValidateRevision(ctx context.Context, formats strfmt.Registry) error {

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
func (m *AccessPolicySummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccessPolicySummary) UnmarshalBinary(b []byte) error {
	var res AccessPolicySummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
