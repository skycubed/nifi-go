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

// VersionControlComponentMappingEntity version control component mapping entity
//
// swagger:model VersionControlComponentMappingEntity
type VersionControlComponentMappingEntity struct {

	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged bool `json:"disconnectedNodeAcknowledged,omitempty"`

	// The revision of the Process Group
	ProcessGroupRevision *RevisionDTO `json:"processGroupRevision,omitempty"`

	// The mapping of Versioned Component Identifiers to instance ID's
	VersionControlComponentMapping map[string]string `json:"versionControlComponentMapping,omitempty"`

	// The Version Control information
	VersionControlInformation *VersionControlInformationDTO `json:"versionControlInformation,omitempty"`
}

// Validate validates this version control component mapping entity
func (m *VersionControlComponentMappingEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProcessGroupRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionControlInformation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VersionControlComponentMappingEntity) validateProcessGroupRevision(formats strfmt.Registry) error {
	if swag.IsZero(m.ProcessGroupRevision) { // not required
		return nil
	}

	if m.ProcessGroupRevision != nil {
		if err := m.ProcessGroupRevision.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("processGroupRevision")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("processGroupRevision")
			}
			return err
		}
	}

	return nil
}

func (m *VersionControlComponentMappingEntity) validateVersionControlInformation(formats strfmt.Registry) error {
	if swag.IsZero(m.VersionControlInformation) { // not required
		return nil
	}

	if m.VersionControlInformation != nil {
		if err := m.VersionControlInformation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("versionControlInformation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("versionControlInformation")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this version control component mapping entity based on the context it is used
func (m *VersionControlComponentMappingEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProcessGroupRevision(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVersionControlInformation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VersionControlComponentMappingEntity) contextValidateProcessGroupRevision(ctx context.Context, formats strfmt.Registry) error {

	if m.ProcessGroupRevision != nil {

		if swag.IsZero(m.ProcessGroupRevision) { // not required
			return nil
		}

		if err := m.ProcessGroupRevision.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("processGroupRevision")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("processGroupRevision")
			}
			return err
		}
	}

	return nil
}

func (m *VersionControlComponentMappingEntity) contextValidateVersionControlInformation(ctx context.Context, formats strfmt.Registry) error {

	if m.VersionControlInformation != nil {

		if swag.IsZero(m.VersionControlInformation) { // not required
			return nil
		}

		if err := m.VersionControlInformation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("versionControlInformation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("versionControlInformation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VersionControlComponentMappingEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VersionControlComponentMappingEntity) UnmarshalBinary(b []byte) error {
	var res VersionControlComponentMappingEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
