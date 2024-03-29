// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RemoteProcessGroupStatusDTO remote process group status d t o
//
// swagger:model RemoteProcessGroupStatusDTO
type RemoteProcessGroupStatusDTO struct {

	// A status snapshot that represents the aggregate stats of all nodes in the cluster. If the NiFi instance is a standalone instance, rather than a cluster, this represents the stats of the single instance.
	AggregateSnapshot *RemoteProcessGroupStatusSnapshotDTO `json:"aggregateSnapshot,omitempty"`

	// The unique ID of the process group that the Processor belongs to
	GroupID string `json:"groupId,omitempty"`

	// The unique ID of the Processor
	ID string `json:"id,omitempty"`

	// The name of the remote process group.
	Name string `json:"name,omitempty"`

	// A status snapshot for each node in the cluster. If the NiFi instance is a standalone instance, rather than a cluster, this may be null.
	NodeSnapshots []*NodeRemoteProcessGroupStatusSnapshotDTO `json:"nodeSnapshots"`

	// The time the status for the process group was last refreshed.
	StatsLastRefreshed string `json:"statsLastRefreshed,omitempty"`

	// The URI of the target system.
	TargetURI string `json:"targetUri,omitempty"`

	// The transmission status of the remote process group.
	TransmissionStatus string `json:"transmissionStatus,omitempty"`

	// Indicates whether the component is valid, invalid, or still in the process of validating (i.e., it is unknown whether or not the component is valid)
	// Enum: [VALID INVALID VALIDATING]
	ValidationStatus string `json:"validationStatus,omitempty"`
}

// Validate validates this remote process group status d t o
func (m *RemoteProcessGroupStatusDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAggregateSnapshot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeSnapshots(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidationStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemoteProcessGroupStatusDTO) validateAggregateSnapshot(formats strfmt.Registry) error {
	if swag.IsZero(m.AggregateSnapshot) { // not required
		return nil
	}

	if m.AggregateSnapshot != nil {
		if err := m.AggregateSnapshot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregateSnapshot")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aggregateSnapshot")
			}
			return err
		}
	}

	return nil
}

func (m *RemoteProcessGroupStatusDTO) validateNodeSnapshots(formats strfmt.Registry) error {
	if swag.IsZero(m.NodeSnapshots) { // not required
		return nil
	}

	for i := 0; i < len(m.NodeSnapshots); i++ {
		if swag.IsZero(m.NodeSnapshots[i]) { // not required
			continue
		}

		if m.NodeSnapshots[i] != nil {
			if err := m.NodeSnapshots[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodeSnapshots" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nodeSnapshots" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var remoteProcessGroupStatusDTOTypeValidationStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["VALID","INVALID","VALIDATING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		remoteProcessGroupStatusDTOTypeValidationStatusPropEnum = append(remoteProcessGroupStatusDTOTypeValidationStatusPropEnum, v)
	}
}

const (

	// RemoteProcessGroupStatusDTOValidationStatusVALID captures enum value "VALID"
	RemoteProcessGroupStatusDTOValidationStatusVALID string = "VALID"

	// RemoteProcessGroupStatusDTOValidationStatusINVALID captures enum value "INVALID"
	RemoteProcessGroupStatusDTOValidationStatusINVALID string = "INVALID"

	// RemoteProcessGroupStatusDTOValidationStatusVALIDATING captures enum value "VALIDATING"
	RemoteProcessGroupStatusDTOValidationStatusVALIDATING string = "VALIDATING"
)

// prop value enum
func (m *RemoteProcessGroupStatusDTO) validateValidationStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, remoteProcessGroupStatusDTOTypeValidationStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RemoteProcessGroupStatusDTO) validateValidationStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.ValidationStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateValidationStatusEnum("validationStatus", "body", m.ValidationStatus); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this remote process group status d t o based on the context it is used
func (m *RemoteProcessGroupStatusDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAggregateSnapshot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNodeSnapshots(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemoteProcessGroupStatusDTO) contextValidateAggregateSnapshot(ctx context.Context, formats strfmt.Registry) error {

	if m.AggregateSnapshot != nil {

		if swag.IsZero(m.AggregateSnapshot) { // not required
			return nil
		}

		if err := m.AggregateSnapshot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregateSnapshot")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aggregateSnapshot")
			}
			return err
		}
	}

	return nil
}

func (m *RemoteProcessGroupStatusDTO) contextValidateNodeSnapshots(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NodeSnapshots); i++ {

		if m.NodeSnapshots[i] != nil {

			if swag.IsZero(m.NodeSnapshots[i]) { // not required
				return nil
			}

			if err := m.NodeSnapshots[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodeSnapshots" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nodeSnapshots" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RemoteProcessGroupStatusDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemoteProcessGroupStatusDTO) UnmarshalBinary(b []byte) error {
	var res RemoteProcessGroupStatusDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
