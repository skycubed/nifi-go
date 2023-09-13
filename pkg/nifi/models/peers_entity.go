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

// PeersEntity peers entity
//
// swagger:model PeersEntity
type PeersEntity struct {

	// peers
	Peers []*PeerDTO `json:"peers"`
}

// Validate validates this peers entity
func (m *PeersEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePeers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PeersEntity) validatePeers(formats strfmt.Registry) error {
	if swag.IsZero(m.Peers) { // not required
		return nil
	}

	for i := 0; i < len(m.Peers); i++ {
		if swag.IsZero(m.Peers[i]) { // not required
			continue
		}

		if m.Peers[i] != nil {
			if err := m.Peers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("peers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("peers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this peers entity based on the context it is used
func (m *PeersEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePeers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PeersEntity) contextValidatePeers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Peers); i++ {

		if m.Peers[i] != nil {

			if swag.IsZero(m.Peers[i]) { // not required
				return nil
			}

			if err := m.Peers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("peers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("peers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PeersEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PeersEntity) UnmarshalBinary(b []byte) error {
	var res PeersEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
