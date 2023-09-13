// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ReportingTaskRunStatusEntity reporting task run status entity
//
// swagger:model ReportingTaskRunStatusEntity
type ReportingTaskRunStatusEntity struct {

	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged bool `json:"disconnectedNodeAcknowledged,omitempty"`

	// The revision for this request/response. The revision is required for any mutable flow requests and is included in all responses.
	Revision *RevisionDTO `json:"revision,omitempty"`

	// The run status of the ReportingTask.
	// Enum: [RUNNING STOPPED]
	State string `json:"state,omitempty"`
}

// Validate validates this reporting task run status entity
func (m *ReportingTaskRunStatusEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportingTaskRunStatusEntity) validateRevision(formats strfmt.Registry) error {
	if swag.IsZero(m.Revision) { // not required
		return nil
	}

	if m.Revision != nil {
		if err := m.Revision.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("revision")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("revision")
			}
			return err
		}
	}

	return nil
}

var reportingTaskRunStatusEntityTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RUNNING","STOPPED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		reportingTaskRunStatusEntityTypeStatePropEnum = append(reportingTaskRunStatusEntityTypeStatePropEnum, v)
	}
}

const (

	// ReportingTaskRunStatusEntityStateRUNNING captures enum value "RUNNING"
	ReportingTaskRunStatusEntityStateRUNNING string = "RUNNING"

	// ReportingTaskRunStatusEntityStateSTOPPED captures enum value "STOPPED"
	ReportingTaskRunStatusEntityStateSTOPPED string = "STOPPED"
)

// prop value enum
func (m *ReportingTaskRunStatusEntity) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, reportingTaskRunStatusEntityTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ReportingTaskRunStatusEntity) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this reporting task run status entity based on the context it is used
func (m *ReportingTaskRunStatusEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRevision(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportingTaskRunStatusEntity) contextValidateRevision(ctx context.Context, formats strfmt.Registry) error {

	if m.Revision != nil {

		if swag.IsZero(m.Revision) { // not required
			return nil
		}

		if err := m.Revision.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("revision")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("revision")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReportingTaskRunStatusEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportingTaskRunStatusEntity) UnmarshalBinary(b []byte) error {
	var res ReportingTaskRunStatusEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
