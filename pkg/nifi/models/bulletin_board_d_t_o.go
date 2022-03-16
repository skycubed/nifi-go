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

// BulletinBoardDTO bulletin board d t o
//
// swagger:model BulletinBoardDTO
type BulletinBoardDTO struct {

	// The bulletins in the bulletin board, that matches the supplied request.
	Bulletins []*BulletinEntity `json:"bulletins"`

	// The timestamp when this report was generated.
	Generated string `json:"generated,omitempty"`
}

// Validate validates this bulletin board d t o
func (m *BulletinBoardDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBulletins(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BulletinBoardDTO) validateBulletins(formats strfmt.Registry) error {
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

// ContextValidate validate this bulletin board d t o based on the context it is used
func (m *BulletinBoardDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBulletins(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BulletinBoardDTO) contextValidateBulletins(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *BulletinBoardDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BulletinBoardDTO) UnmarshalBinary(b []byte) error {
	var res BulletinBoardDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
