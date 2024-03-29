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

// ConnectionStatusSnapshotDTO connection status snapshot d t o
//
// swagger:model ConnectionStatusSnapshotDTO
type ConnectionStatusSnapshotDTO struct {

	// The size of the FlowFiles that have come into the connection in the last 5 minutes.
	BytesIn int64 `json:"bytesIn,omitempty"`

	// The number of bytes that have left the connection in the last 5 minutes.
	BytesOut int64 `json:"bytesOut,omitempty"`

	// The size of the FlowFiles that are currently queued in the connection.
	BytesQueued int64 `json:"bytesQueued,omitempty"`

	// The id of the destination of the connection.
	DestinationID string `json:"destinationId,omitempty"`

	// The name of the destination of the connection.
	DestinationName string `json:"destinationName,omitempty"`

	// The number of FlowFiles that have come into the connection in the last 5 minutes.
	FlowFilesIn int32 `json:"flowFilesIn,omitempty"`

	// The number of FlowFiles that have left the connection in the last 5 minutes.
	FlowFilesOut int32 `json:"flowFilesOut,omitempty"`

	// The number of FlowFiles that are currently queued in the connection.
	FlowFilesQueued int32 `json:"flowFilesQueued,omitempty"`

	// The id of the process group the connection belongs to.
	GroupID string `json:"groupId,omitempty"`

	// The id of the connection.
	ID string `json:"id,omitempty"`

	// The input count/size for the connection in the last 5 minutes, pretty printed.
	Input string `json:"input,omitempty"`

	// The name of the connection.
	Name string `json:"name,omitempty"`

	// The output count/sie for the connection in the last 5 minutes, pretty printed.
	Output string `json:"output,omitempty"`

	// Connection percent use regarding queued flow files size and backpressure threshold if configured.
	PercentUseBytes int32 `json:"percentUseBytes,omitempty"`

	// Connection percent use regarding queued flow files count and backpressure threshold if configured.
	PercentUseCount int32 `json:"percentUseCount,omitempty"`

	// Predictions, if available, for this connection (null if not available)
	Predictions *ConnectionStatusPredictionsSnapshotDTO `json:"predictions,omitempty"`

	// The total count and size of queued flowfiles formatted.
	Queued string `json:"queued,omitempty"`

	// The number of flowfiles that are queued, pretty printed.
	QueuedCount string `json:"queuedCount,omitempty"`

	// The total size of flowfiles that are queued formatted.
	QueuedSize string `json:"queuedSize,omitempty"`

	// The id of the source of the connection.
	SourceID string `json:"sourceId,omitempty"`

	// The name of the source of the connection.
	SourceName string `json:"sourceName,omitempty"`
}

// Validate validates this connection status snapshot d t o
func (m *ConnectionStatusSnapshotDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePredictions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConnectionStatusSnapshotDTO) validatePredictions(formats strfmt.Registry) error {
	if swag.IsZero(m.Predictions) { // not required
		return nil
	}

	if m.Predictions != nil {
		if err := m.Predictions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("predictions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("predictions")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this connection status snapshot d t o based on the context it is used
func (m *ConnectionStatusSnapshotDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePredictions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConnectionStatusSnapshotDTO) contextValidatePredictions(ctx context.Context, formats strfmt.Registry) error {

	if m.Predictions != nil {

		if swag.IsZero(m.Predictions) { // not required
			return nil
		}

		if err := m.Predictions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("predictions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("predictions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConnectionStatusSnapshotDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConnectionStatusSnapshotDTO) UnmarshalBinary(b []byte) error {
	var res ConnectionStatusSnapshotDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
