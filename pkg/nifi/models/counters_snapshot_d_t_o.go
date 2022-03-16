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

// CountersSnapshotDTO counters snapshot d t o
//
// swagger:model CountersSnapshotDTO
type CountersSnapshotDTO struct {

	// All counters in the NiFi.
	Counters []*CounterDTO `json:"counters"`

	// The timestamp when the report was generated.
	Generated string `json:"generated,omitempty"`
}

// Validate validates this counters snapshot d t o
func (m *CountersSnapshotDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCounters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CountersSnapshotDTO) validateCounters(formats strfmt.Registry) error {
	if swag.IsZero(m.Counters) { // not required
		return nil
	}

	for i := 0; i < len(m.Counters); i++ {
		if swag.IsZero(m.Counters[i]) { // not required
			continue
		}

		if m.Counters[i] != nil {
			if err := m.Counters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("counters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("counters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this counters snapshot d t o based on the context it is used
func (m *CountersSnapshotDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCounters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CountersSnapshotDTO) contextValidateCounters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Counters); i++ {

		if m.Counters[i] != nil {
			if err := m.Counters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("counters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("counters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CountersSnapshotDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CountersSnapshotDTO) UnmarshalBinary(b []byte) error {
	var res CountersSnapshotDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
