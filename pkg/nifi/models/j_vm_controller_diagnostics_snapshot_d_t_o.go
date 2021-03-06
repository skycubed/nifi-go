// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// JVMControllerDiagnosticsSnapshotDTO j VM controller diagnostics snapshot d t o
//
// swagger:model JVMControllerDiagnosticsSnapshotDTO
type JVMControllerDiagnosticsSnapshotDTO struct {

	// Whether or not this node is cluster coordinator
	ClusterCoordinator bool `json:"clusterCoordinator,omitempty"`

	// The maximum number of event-driven threads
	MaxEventDrivenThreads int32 `json:"maxEventDrivenThreads,omitempty"`

	// The maximum number of timer-driven threads
	MaxTimerDrivenThreads int32 `json:"maxTimerDrivenThreads,omitempty"`

	// Whether or not this node is primary node
	PrimaryNode bool `json:"primaryNode,omitempty"`
}

// Validate validates this j VM controller diagnostics snapshot d t o
func (m *JVMControllerDiagnosticsSnapshotDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this j VM controller diagnostics snapshot d t o based on context it is used
func (m *JVMControllerDiagnosticsSnapshotDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *JVMControllerDiagnosticsSnapshotDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JVMControllerDiagnosticsSnapshotDTO) UnmarshalBinary(b []byte) error {
	var res JVMControllerDiagnosticsSnapshotDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
