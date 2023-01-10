// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RegisteredFlowSnapshot registered flow snapshot
//
// swagger:model RegisteredFlowSnapshot
type RegisteredFlowSnapshot struct {

	// bucket
	Bucket *FlowRegistryBucket `json:"bucket,omitempty"`

	// external controller services
	ExternalControllerServices map[string]ExternalControllerServiceReference `json:"externalControllerServices,omitempty"`

	// flow
	Flow *RegisteredFlow `json:"flow,omitempty"`

	// flow contents
	FlowContents *VersionedProcessGroup `json:"flowContents,omitempty"`

	// flow encoding version
	FlowEncodingVersion string `json:"flowEncodingVersion,omitempty"`

	// latest
	Latest bool `json:"latest,omitempty"`

	// parameter contexts
	ParameterContexts map[string]VersionedParameterContext `json:"parameterContexts,omitempty"`

	// parameter providers
	ParameterProviders map[string]ParameterProviderReference `json:"parameterProviders,omitempty"`

	// snapshot metadata
	SnapshotMetadata *RegisteredFlowSnapshotMetadata `json:"snapshotMetadata,omitempty"`
}

// Validate validates this registered flow snapshot
func (m *RegisteredFlowSnapshot) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBucket(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalControllerServices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlowContents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameterContexts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameterProviders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegisteredFlowSnapshot) validateBucket(formats strfmt.Registry) error {
	if swag.IsZero(m.Bucket) { // not required
		return nil
	}

	if m.Bucket != nil {
		if err := m.Bucket.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bucket")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bucket")
			}
			return err
		}
	}

	return nil
}

func (m *RegisteredFlowSnapshot) validateExternalControllerServices(formats strfmt.Registry) error {
	if swag.IsZero(m.ExternalControllerServices) { // not required
		return nil
	}

	for k := range m.ExternalControllerServices {

		if err := validate.Required("externalControllerServices"+"."+k, "body", m.ExternalControllerServices[k]); err != nil {
			return err
		}
		if val, ok := m.ExternalControllerServices[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("externalControllerServices" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("externalControllerServices" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *RegisteredFlowSnapshot) validateFlow(formats strfmt.Registry) error {
	if swag.IsZero(m.Flow) { // not required
		return nil
	}

	if m.Flow != nil {
		if err := m.Flow.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flow")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("flow")
			}
			return err
		}
	}

	return nil
}

func (m *RegisteredFlowSnapshot) validateFlowContents(formats strfmt.Registry) error {
	if swag.IsZero(m.FlowContents) { // not required
		return nil
	}

	if m.FlowContents != nil {
		if err := m.FlowContents.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flowContents")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("flowContents")
			}
			return err
		}
	}

	return nil
}

func (m *RegisteredFlowSnapshot) validateParameterContexts(formats strfmt.Registry) error {
	if swag.IsZero(m.ParameterContexts) { // not required
		return nil
	}

	for k := range m.ParameterContexts {

		if err := validate.Required("parameterContexts"+"."+k, "body", m.ParameterContexts[k]); err != nil {
			return err
		}
		if val, ok := m.ParameterContexts[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameterContexts" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parameterContexts" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *RegisteredFlowSnapshot) validateParameterProviders(formats strfmt.Registry) error {
	if swag.IsZero(m.ParameterProviders) { // not required
		return nil
	}

	for k := range m.ParameterProviders {

		if err := validate.Required("parameterProviders"+"."+k, "body", m.ParameterProviders[k]); err != nil {
			return err
		}
		if val, ok := m.ParameterProviders[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameterProviders" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parameterProviders" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *RegisteredFlowSnapshot) validateSnapshotMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.SnapshotMetadata) { // not required
		return nil
	}

	if m.SnapshotMetadata != nil {
		if err := m.SnapshotMetadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshotMetadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("snapshotMetadata")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this registered flow snapshot based on the context it is used
func (m *RegisteredFlowSnapshot) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBucket(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExternalControllerServices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFlow(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFlowContents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParameterContexts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParameterProviders(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSnapshotMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegisteredFlowSnapshot) contextValidateBucket(ctx context.Context, formats strfmt.Registry) error {

	if m.Bucket != nil {
		if err := m.Bucket.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bucket")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bucket")
			}
			return err
		}
	}

	return nil
}

func (m *RegisteredFlowSnapshot) contextValidateExternalControllerServices(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.ExternalControllerServices {

		if val, ok := m.ExternalControllerServices[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *RegisteredFlowSnapshot) contextValidateFlow(ctx context.Context, formats strfmt.Registry) error {

	if m.Flow != nil {
		if err := m.Flow.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flow")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("flow")
			}
			return err
		}
	}

	return nil
}

func (m *RegisteredFlowSnapshot) contextValidateFlowContents(ctx context.Context, formats strfmt.Registry) error {

	if m.FlowContents != nil {
		if err := m.FlowContents.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flowContents")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("flowContents")
			}
			return err
		}
	}

	return nil
}

func (m *RegisteredFlowSnapshot) contextValidateParameterContexts(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.ParameterContexts {

		if val, ok := m.ParameterContexts[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *RegisteredFlowSnapshot) contextValidateParameterProviders(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.ParameterProviders {

		if val, ok := m.ParameterProviders[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *RegisteredFlowSnapshot) contextValidateSnapshotMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.SnapshotMetadata != nil {
		if err := m.SnapshotMetadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshotMetadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("snapshotMetadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RegisteredFlowSnapshot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegisteredFlowSnapshot) UnmarshalBinary(b []byte) error {
	var res RegisteredFlowSnapshot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}