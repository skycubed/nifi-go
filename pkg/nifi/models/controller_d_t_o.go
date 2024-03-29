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
	"github.com/go-openapi/validate"
)

// ControllerDTO controller d t o
//
// swagger:model ControllerDTO
type ControllerDTO struct {

	// The number of active remote ports contained in the NiFi.
	ActiveRemotePortCount int32 `json:"activeRemotePortCount,omitempty"`

	// The comments for the NiFi.
	Comments string `json:"comments,omitempty"`

	// The number of disabled components in the NiFi.
	DisabledCount int32 `json:"disabledCount,omitempty"`

	// The id of the NiFi.
	ID string `json:"id,omitempty"`

	// The number of inactive remote ports contained in the NiFi.
	InactiveRemotePortCount int32 `json:"inactiveRemotePortCount,omitempty"`

	// The number of input ports contained in the NiFi.
	InputPortCount int32 `json:"inputPortCount,omitempty"`

	// The input ports available to send data to for the NiFi.
	// Unique: true
	InputPorts []*PortDTO `json:"inputPorts"`

	// If clustered, the id of the Cluster Manager, otherwise the id of the NiFi.
	InstanceID string `json:"instanceId,omitempty"`

	// The number of invalid components in the NiFi.
	InvalidCount int32 `json:"invalidCount,omitempty"`

	// The name of the NiFi.
	Name string `json:"name,omitempty"`

	// The number of output ports in the NiFi.
	OutputPortCount int32 `json:"outputPortCount,omitempty"`

	// The output ports available to received data from the NiFi.
	// Unique: true
	OutputPorts []*PortDTO `json:"outputPorts"`

	// The HTTP(S) Port on which this instance is listening for Remote Transfers of Flow Files. If this instance is not configured to receive Flow Files from remote instances, this will be null.
	RemoteSiteHTTPListeningPort int32 `json:"remoteSiteHttpListeningPort,omitempty"`

	// The Socket Port on which this instance is listening for Remote Transfers of Flow Files. If this instance is not configured to receive Flow Files from remote instances, this will be null.
	RemoteSiteListeningPort int32 `json:"remoteSiteListeningPort,omitempty"`

	// The number of running components in the NiFi.
	RunningCount int32 `json:"runningCount,omitempty"`

	// Indicates whether or not Site-to-Site communications with this instance is secure (2-way authentication).
	SiteToSiteSecure bool `json:"siteToSiteSecure,omitempty"`

	// The number of stopped components in the NiFi.
	StoppedCount int32 `json:"stoppedCount,omitempty"`
}

// Validate validates this controller d t o
func (m *ControllerDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInputPorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutputPorts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ControllerDTO) validateInputPorts(formats strfmt.Registry) error {
	if swag.IsZero(m.InputPorts) { // not required
		return nil
	}

	if err := validate.UniqueItems("inputPorts", "body", m.InputPorts); err != nil {
		return err
	}

	for i := 0; i < len(m.InputPorts); i++ {
		if swag.IsZero(m.InputPorts[i]) { // not required
			continue
		}

		if m.InputPorts[i] != nil {
			if err := m.InputPorts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inputPorts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("inputPorts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ControllerDTO) validateOutputPorts(formats strfmt.Registry) error {
	if swag.IsZero(m.OutputPorts) { // not required
		return nil
	}

	if err := validate.UniqueItems("outputPorts", "body", m.OutputPorts); err != nil {
		return err
	}

	for i := 0; i < len(m.OutputPorts); i++ {
		if swag.IsZero(m.OutputPorts[i]) { // not required
			continue
		}

		if m.OutputPorts[i] != nil {
			if err := m.OutputPorts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("outputPorts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("outputPorts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this controller d t o based on the context it is used
func (m *ControllerDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInputPorts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOutputPorts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ControllerDTO) contextValidateInputPorts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InputPorts); i++ {

		if m.InputPorts[i] != nil {

			if swag.IsZero(m.InputPorts[i]) { // not required
				return nil
			}

			if err := m.InputPorts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inputPorts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("inputPorts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ControllerDTO) contextValidateOutputPorts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OutputPorts); i++ {

		if m.OutputPorts[i] != nil {

			if swag.IsZero(m.OutputPorts[i]) { // not required
				return nil
			}

			if err := m.OutputPorts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("outputPorts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("outputPorts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ControllerDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ControllerDTO) UnmarshalBinary(b []byte) error {
	var res ControllerDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
