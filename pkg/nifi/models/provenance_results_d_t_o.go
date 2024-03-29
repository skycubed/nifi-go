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
	"github.com/go-openapi/validate"
)

// ProvenanceResultsDTO provenance results d t o
//
// swagger:model ProvenanceResultsDTO
type ProvenanceResultsDTO struct {

	// Any errors that occurred while performing the provenance request.
	// Unique: true
	Errors []string `json:"errors"`

	// Then the search was performed.
	Generated string `json:"generated,omitempty"`

	// The oldest event available in the provenance repository.
	OldestEvent string `json:"oldestEvent,omitempty"`

	// The provenance events that matched the search criteria.
	ProvenanceEvents []*ProvenanceEventDTO `json:"provenanceEvents"`

	// The time offset of the server that's used for event time.
	TimeOffset int32 `json:"timeOffset,omitempty"`

	// The total number of results formatted.
	Total string `json:"total,omitempty"`

	// The total number of results.
	TotalCount int64 `json:"totalCount,omitempty"`
}

// Validate validates this provenance results d t o
func (m *ProvenanceResultsDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvenanceEvents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProvenanceResultsDTO) validateErrors(formats strfmt.Registry) error {
	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	if err := validate.UniqueItems("errors", "body", m.Errors); err != nil {
		return err
	}

	return nil
}

func (m *ProvenanceResultsDTO) validateProvenanceEvents(formats strfmt.Registry) error {
	if swag.IsZero(m.ProvenanceEvents) { // not required
		return nil
	}

	for i := 0; i < len(m.ProvenanceEvents); i++ {
		if swag.IsZero(m.ProvenanceEvents[i]) { // not required
			continue
		}

		if m.ProvenanceEvents[i] != nil {
			if err := m.ProvenanceEvents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("provenanceEvents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("provenanceEvents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this provenance results d t o based on the context it is used
func (m *ProvenanceResultsDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProvenanceEvents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProvenanceResultsDTO) contextValidateProvenanceEvents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ProvenanceEvents); i++ {

		if m.ProvenanceEvents[i] != nil {

			if swag.IsZero(m.ProvenanceEvents[i]) { // not required
				return nil
			}

			if err := m.ProvenanceEvents[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("provenanceEvents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("provenanceEvents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProvenanceResultsDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProvenanceResultsDTO) UnmarshalBinary(b []byte) error {
	var res ProvenanceResultsDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
