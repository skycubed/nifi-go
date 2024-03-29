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

// ProvenanceDTO provenance d t o
//
// swagger:model ProvenanceDTO
type ProvenanceDTO struct {

	// The timestamp when the query will expire.
	Expiration string `json:"expiration,omitempty"`

	// Whether the query has finished.
	Finished bool `json:"finished,omitempty"`

	// The id of the provenance query.
	ID string `json:"id,omitempty"`

	// The current percent complete.
	PercentCompleted int32 `json:"percentCompleted,omitempty"`

	// The provenance request.
	Request *ProvenanceRequestDTO `json:"request,omitempty"`

	// The provenance results.
	Results *ProvenanceResultsDTO `json:"results,omitempty"`

	// The timestamp when the query was submitted.
	SubmissionTime string `json:"submissionTime,omitempty"`

	// The URI for this query. Used for obtaining/deleting the request at a later time
	URI string `json:"uri,omitempty"`
}

// Validate validates this provenance d t o
func (m *ProvenanceDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRequest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProvenanceDTO) validateRequest(formats strfmt.Registry) error {
	if swag.IsZero(m.Request) { // not required
		return nil
	}

	if m.Request != nil {
		if err := m.Request.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("request")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("request")
			}
			return err
		}
	}

	return nil
}

func (m *ProvenanceDTO) validateResults(formats strfmt.Registry) error {
	if swag.IsZero(m.Results) { // not required
		return nil
	}

	if m.Results != nil {
		if err := m.Results.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("results")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("results")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this provenance d t o based on the context it is used
func (m *ProvenanceDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRequest(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProvenanceDTO) contextValidateRequest(ctx context.Context, formats strfmt.Registry) error {

	if m.Request != nil {

		if swag.IsZero(m.Request) { // not required
			return nil
		}

		if err := m.Request.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("request")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("request")
			}
			return err
		}
	}

	return nil
}

func (m *ProvenanceDTO) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	if m.Results != nil {

		if swag.IsZero(m.Results) { // not required
			return nil
		}

		if err := m.Results.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("results")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("results")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProvenanceDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProvenanceDTO) UnmarshalBinary(b []byte) error {
	var res ProvenanceDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
