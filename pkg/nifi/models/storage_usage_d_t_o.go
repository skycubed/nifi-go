// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorageUsageDTO storage usage d t o
//
// swagger:model StorageUsageDTO
type StorageUsageDTO struct {

	// Amount of free space.
	FreeSpace string `json:"freeSpace,omitempty"`

	// The number of bytes of free space.
	FreeSpaceBytes int64 `json:"freeSpaceBytes,omitempty"`

	// The identifier of this storage location. The identifier will correspond to the identifier keyed in the storage configuration.
	Identifier string `json:"identifier,omitempty"`

	// Amount of total space.
	TotalSpace string `json:"totalSpace,omitempty"`

	// The number of bytes of total space.
	TotalSpaceBytes int64 `json:"totalSpaceBytes,omitempty"`

	// Amount of used space.
	UsedSpace string `json:"usedSpace,omitempty"`

	// The number of bytes of used space.
	UsedSpaceBytes int64 `json:"usedSpaceBytes,omitempty"`

	// Utilization of this storage location.
	Utilization string `json:"utilization,omitempty"`
}

// Validate validates this storage usage d t o
func (m *StorageUsageDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this storage usage d t o based on context it is used
func (m *StorageUsageDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StorageUsageDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageUsageDTO) UnmarshalBinary(b []byte) error {
	var res StorageUsageDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}