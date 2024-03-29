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

// ExtensionRepoVersionSummary extension repo version summary
//
// swagger:model ExtensionRepoVersionSummary
type ExtensionRepoVersionSummary struct {

	// The artifact id
	ArtifactID string `json:"artifactId,omitempty"`

	// The identity of the user that created this version
	Author string `json:"author,omitempty"`

	// The bucket name
	BucketName string `json:"bucketName,omitempty"`

	// The group id
	GroupID string `json:"groupId,omitempty"`

	// An WebLink to this entity.
	// Read Only: true
	Link *JaxbLink `json:"link,omitempty"`

	// The timestamp of when this version was created
	Timestamp int64 `json:"timestamp,omitempty"`

	// The version
	Version string `json:"version,omitempty"`
}

// Validate validates this extension repo version summary
func (m *ExtensionRepoVersionSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLink(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExtensionRepoVersionSummary) validateLink(formats strfmt.Registry) error {
	if swag.IsZero(m.Link) { // not required
		return nil
	}

	if m.Link != nil {
		if err := m.Link.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("link")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("link")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this extension repo version summary based on the context it is used
func (m *ExtensionRepoVersionSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLink(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExtensionRepoVersionSummary) contextValidateLink(ctx context.Context, formats strfmt.Registry) error {

	if m.Link != nil {

		if swag.IsZero(m.Link) { // not required
			return nil
		}

		if err := m.Link.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("link")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("link")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExtensionRepoVersionSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExtensionRepoVersionSummary) UnmarshalBinary(b []byte) error {
	var res ExtensionRepoVersionSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
