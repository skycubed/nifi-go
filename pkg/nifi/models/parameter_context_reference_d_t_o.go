// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ParameterContextReferenceDTO parameter context reference d t o
//
// swagger:model ParameterContextReferenceDTO
type ParameterContextReferenceDTO struct {

	// The ID of the Parameter Context
	ID string `json:"id,omitempty"`

	// The name of the Parameter Context
	Name string `json:"name,omitempty"`
}

// Validate validates this parameter context reference d t o
func (m *ParameterContextReferenceDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this parameter context reference d t o based on context it is used
func (m *ParameterContextReferenceDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ParameterContextReferenceDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ParameterContextReferenceDTO) UnmarshalBinary(b []byte) error {
	var res ParameterContextReferenceDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}