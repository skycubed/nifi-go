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

// ConnectionStatisticsDTO connection statistics d t o
//
// swagger:model ConnectionStatisticsDTO
type ConnectionStatisticsDTO struct {

	// The status snapshot that represents the aggregate stats of the cluster
	AggregateSnapshot *ConnectionStatisticsSnapshotDTO `json:"aggregateSnapshot,omitempty"`

	// The ID of the connection
	ID string `json:"id,omitempty"`

	// A list of status snapshots for each node
	NodeSnapshots []*NodeConnectionStatisticsSnapshotDTO `json:"nodeSnapshots"`

	// The timestamp of when the stats were last refreshed
	StatsLastRefreshed string `json:"statsLastRefreshed,omitempty"`
}

// Validate validates this connection statistics d t o
func (m *ConnectionStatisticsDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAggregateSnapshot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeSnapshots(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConnectionStatisticsDTO) validateAggregateSnapshot(formats strfmt.Registry) error {
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

func (m *ConnectionStatisticsDTO) validateNodeSnapshots(formats strfmt.Registry) error {
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

// ContextValidate validate this connection statistics d t o based on the context it is used
func (m *ConnectionStatisticsDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *ConnectionStatisticsDTO) contextValidateAggregateSnapshot(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ConnectionStatisticsDTO) contextValidateNodeSnapshots(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ConnectionStatisticsDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConnectionStatisticsDTO) UnmarshalBinary(b []byte) error {
	var res ConnectionStatisticsDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
