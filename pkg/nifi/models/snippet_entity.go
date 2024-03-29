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

// SnippetEntity snippet entity
//
// swagger:model SnippetEntity
type SnippetEntity struct {

	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged bool `json:"disconnectedNodeAcknowledged,omitempty"`

	// The snippet.
	Snippet *SnippetDTO `json:"snippet,omitempty"`
}

// Validate validates this snippet entity
func (m *SnippetEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSnippet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnippetEntity) validateSnippet(formats strfmt.Registry) error {
	if swag.IsZero(m.Snippet) { // not required
		return nil
	}

	if m.Snippet != nil {
		if err := m.Snippet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snippet")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("snippet")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snippet entity based on the context it is used
func (m *SnippetEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSnippet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnippetEntity) contextValidateSnippet(ctx context.Context, formats strfmt.Registry) error {

	if m.Snippet != nil {

		if swag.IsZero(m.Snippet) { // not required
			return nil
		}

		if err := m.Snippet.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snippet")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("snippet")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnippetEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnippetEntity) UnmarshalBinary(b []byte) error {
	var res SnippetEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
