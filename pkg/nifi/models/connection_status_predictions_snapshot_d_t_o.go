// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConnectionStatusPredictionsSnapshotDTO connection status predictions snapshot d t o
//
// swagger:model ConnectionStatusPredictionsSnapshotDTO
type ConnectionStatusPredictionsSnapshotDTO struct {

	// The predicted total number of bytes in the queue at the next configured interval.
	PredictedBytesAtNextInterval int64 `json:"predictedBytesAtNextInterval,omitempty"`

	// The predicted number of queued objects at the next configured interval.
	PredictedCountAtNextInterval int32 `json:"predictedCountAtNextInterval,omitempty"`

	// The predicted number of milliseconds before the connection will have backpressure applied, based on the total number of bytes in the queue.
	PredictedMillisUntilBytesBackpressure int64 `json:"predictedMillisUntilBytesBackpressure,omitempty"`

	// The predicted number of milliseconds before the connection will have backpressure applied, based on the queued count.
	PredictedMillisUntilCountBackpressure int64 `json:"predictedMillisUntilCountBackpressure,omitempty"`

	// Predicted connection percent use regarding queued flow files size and backpressure threshold if configured.
	PredictedPercentBytes int32 `json:"predictedPercentBytes,omitempty"`

	// Predicted connection percent use regarding queued flow files count and backpressure threshold if configured.
	PredictedPercentCount int32 `json:"predictedPercentCount,omitempty"`

	// The configured interval (in seconds) for predicting connection queue count and size (and percent usage).
	PredictionIntervalSeconds int32 `json:"predictionIntervalSeconds,omitempty"`
}

// Validate validates this connection status predictions snapshot d t o
func (m *ConnectionStatusPredictionsSnapshotDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this connection status predictions snapshot d t o based on context it is used
func (m *ConnectionStatusPredictionsSnapshotDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConnectionStatusPredictionsSnapshotDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConnectionStatusPredictionsSnapshotDTO) UnmarshalBinary(b []byte) error {
	var res ConnectionStatusPredictionsSnapshotDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
