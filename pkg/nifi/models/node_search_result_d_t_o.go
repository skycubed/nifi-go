// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NodeSearchResultDTO node search result d t o
//
// swagger:model NodeSearchResultDTO
type NodeSearchResultDTO struct {

	// The address of the node that matched the search.
	Address string `json:"address,omitempty"`

	// The id of the node that matched the search.
	ID string `json:"id,omitempty"`
}

// Validate validates this node search result d t o
func (m *NodeSearchResultDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this node search result d t o based on context it is used
func (m *NodeSearchResultDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NodeSearchResultDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeSearchResultDTO) UnmarshalBinary(b []byte) error {
	var res NodeSearchResultDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
