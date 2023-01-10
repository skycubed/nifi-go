// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BuildInfo build info
//
// swagger:model BuildInfo
type BuildInfo struct {

	// The compiler used for the build
	Compiler string `json:"compiler,omitempty"`

	// The compiler flags used for the build.
	CompilerFlags string `json:"compilerFlags,omitempty"`

	// The SCM revision id of the source code used for this build.
	Revision string `json:"revision,omitempty"`

	// The target architecture of the built component.
	TargetArch string `json:"targetArch,omitempty"`

	// The timestamp (milliseconds since Epoch) of the build.
	Timestamp int64 `json:"timestamp,omitempty"`

	// The version number of the built component.
	Version string `json:"version,omitempty"`
}

// Validate validates this build info
func (m *BuildInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this build info based on context it is used
func (m *BuildInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BuildInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuildInfo) UnmarshalBinary(b []byte) error {
	var res BuildInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
