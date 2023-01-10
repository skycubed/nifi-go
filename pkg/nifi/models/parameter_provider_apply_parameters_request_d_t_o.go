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

// ParameterProviderApplyParametersRequestDTO parameter provider apply parameters request d t o
//
// swagger:model ParameterProviderApplyParametersRequestDTO
type ParameterProviderApplyParametersRequestDTO struct {

	// Whether or not the request is completed
	Complete bool `json:"complete,omitempty"`

	// The reason for the request failing, or null if the request has not failed
	FailureReason string `json:"failureReason,omitempty"`

	// The timestamp of when the request was last updated
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"lastUpdated,omitempty"`

	// The Parameter Contexts updated by this Parameter Provider. This may not be populated until the request has successfully completed.
	ParameterContextUpdates []*ParameterContextUpdateEntity `json:"parameterContextUpdates"`

	// The Parameter Provider that is being operated on. This may not be populated until the request has successfully completed.
	ParameterProvider *ParameterProviderDTO `json:"parameterProvider,omitempty"`

	// A value between 0 and 100 (inclusive) indicating how close the request is to completion
	PercentCompleted int32 `json:"percentCompleted,omitempty"`

	// The components that are referenced by the update.
	// Unique: true
	ReferencingComponents []*AffectedComponentEntity `json:"referencingComponents"`

	// The ID of the request
	RequestID string `json:"requestId,omitempty"`

	// A description of the current state of the request
	State string `json:"state,omitempty"`

	// The timestamp of when the request was submitted
	// Format: date-time
	SubmissionTime strfmt.DateTime `json:"submissionTime,omitempty"`

	// The steps that are required in order to complete the request, along with the status of each
	UpdateSteps []*ParameterProviderApplyParametersUpdateStepDTO `json:"updateSteps"`

	// The URI for the request
	URI string `json:"uri,omitempty"`
}

// Validate validates this parameter provider apply parameters request d t o
func (m *ParameterProviderApplyParametersRequestDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameterContextUpdates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameterProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReferencingComponents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubmissionTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateSteps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ParameterProviderApplyParametersRequestDTO) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("lastUpdated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ParameterProviderApplyParametersRequestDTO) validateParameterContextUpdates(formats strfmt.Registry) error {
	if swag.IsZero(m.ParameterContextUpdates) { // not required
		return nil
	}

	for i := 0; i < len(m.ParameterContextUpdates); i++ {
		if swag.IsZero(m.ParameterContextUpdates[i]) { // not required
			continue
		}

		if m.ParameterContextUpdates[i] != nil {
			if err := m.ParameterContextUpdates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameterContextUpdates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parameterContextUpdates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ParameterProviderApplyParametersRequestDTO) validateParameterProvider(formats strfmt.Registry) error {
	if swag.IsZero(m.ParameterProvider) { // not required
		return nil
	}

	if m.ParameterProvider != nil {
		if err := m.ParameterProvider.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parameterProvider")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parameterProvider")
			}
			return err
		}
	}

	return nil
}

func (m *ParameterProviderApplyParametersRequestDTO) validateReferencingComponents(formats strfmt.Registry) error {
	if swag.IsZero(m.ReferencingComponents) { // not required
		return nil
	}

	if err := validate.UniqueItems("referencingComponents", "body", m.ReferencingComponents); err != nil {
		return err
	}

	for i := 0; i < len(m.ReferencingComponents); i++ {
		if swag.IsZero(m.ReferencingComponents[i]) { // not required
			continue
		}

		if m.ReferencingComponents[i] != nil {
			if err := m.ReferencingComponents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("referencingComponents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("referencingComponents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ParameterProviderApplyParametersRequestDTO) validateSubmissionTime(formats strfmt.Registry) error {
	if swag.IsZero(m.SubmissionTime) { // not required
		return nil
	}

	if err := validate.FormatOf("submissionTime", "body", "date-time", m.SubmissionTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ParameterProviderApplyParametersRequestDTO) validateUpdateSteps(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateSteps) { // not required
		return nil
	}

	for i := 0; i < len(m.UpdateSteps); i++ {
		if swag.IsZero(m.UpdateSteps[i]) { // not required
			continue
		}

		if m.UpdateSteps[i] != nil {
			if err := m.UpdateSteps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateSteps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateSteps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this parameter provider apply parameters request d t o based on the context it is used
func (m *ParameterProviderApplyParametersRequestDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateParameterContextUpdates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParameterProvider(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReferencingComponents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdateSteps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ParameterProviderApplyParametersRequestDTO) contextValidateParameterContextUpdates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ParameterContextUpdates); i++ {

		if m.ParameterContextUpdates[i] != nil {
			if err := m.ParameterContextUpdates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameterContextUpdates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parameterContextUpdates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ParameterProviderApplyParametersRequestDTO) contextValidateParameterProvider(ctx context.Context, formats strfmt.Registry) error {

	if m.ParameterProvider != nil {
		if err := m.ParameterProvider.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parameterProvider")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parameterProvider")
			}
			return err
		}
	}

	return nil
}

func (m *ParameterProviderApplyParametersRequestDTO) contextValidateReferencingComponents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ReferencingComponents); i++ {

		if m.ReferencingComponents[i] != nil {
			if err := m.ReferencingComponents[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("referencingComponents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("referencingComponents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ParameterProviderApplyParametersRequestDTO) contextValidateUpdateSteps(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.UpdateSteps); i++ {

		if m.UpdateSteps[i] != nil {
			if err := m.UpdateSteps[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateSteps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateSteps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ParameterProviderApplyParametersRequestDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ParameterProviderApplyParametersRequestDTO) UnmarshalBinary(b []byte) error {
	var res ParameterProviderApplyParametersRequestDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}