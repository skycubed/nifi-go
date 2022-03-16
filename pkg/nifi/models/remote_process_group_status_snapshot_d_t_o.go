// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RemoteProcessGroupStatusSnapshotDTO remote process group status snapshot d t o
//
// swagger:model RemoteProcessGroupStatusSnapshotDTO
type RemoteProcessGroupStatusSnapshotDTO struct {

	// The number of active threads for the remote process group.
	ActiveThreadCount int32 `json:"activeThreadCount,omitempty"`

	// The size of the FlowFiles received from the remote process group in the last 5 minutes.
	BytesReceived int64 `json:"bytesReceived,omitempty"`

	// The size of the FlowFiles sent to the remote process group in the last 5 minutes.
	BytesSent int64 `json:"bytesSent,omitempty"`

	// The number of FlowFiles received from the remote process group in the last 5 minutes.
	FlowFilesReceived int32 `json:"flowFilesReceived,omitempty"`

	// The number of FlowFiles sent to the remote process group in the last 5 minutes.
	FlowFilesSent int32 `json:"flowFilesSent,omitempty"`

	// The id of the parent process group the remote process group resides in.
	GroupID string `json:"groupId,omitempty"`

	// The id of the remote process group.
	ID string `json:"id,omitempty"`

	// The name of the remote process group.
	Name string `json:"name,omitempty"`

	// The count/size of the flowfiles received from the remote process group in the last 5 minutes.
	Received string `json:"received,omitempty"`

	// The count/size of the flowfiles sent to the remote process group in the last 5 minutes.
	Sent string `json:"sent,omitempty"`

	// The URI of the target system.
	TargetURI string `json:"targetUri,omitempty"`

	// The transmission status of the remote process group.
	TransmissionStatus string `json:"transmissionStatus,omitempty"`
}

// Validate validates this remote process group status snapshot d t o
func (m *RemoteProcessGroupStatusSnapshotDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this remote process group status snapshot d t o based on context it is used
func (m *RemoteProcessGroupStatusSnapshotDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RemoteProcessGroupStatusSnapshotDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemoteProcessGroupStatusSnapshotDTO) UnmarshalBinary(b []byte) error {
	var res RemoteProcessGroupStatusSnapshotDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
