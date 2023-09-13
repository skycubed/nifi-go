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

// FlowSnippetDTO flow snippet d t o
//
// swagger:model FlowSnippetDTO
type FlowSnippetDTO struct {

	// The connections in this flow snippet.
	// Unique: true
	Connections []*ConnectionDTO `json:"connections"`

	// The controller services in this flow snippet.
	// Unique: true
	ControllerServices []*ControllerServiceDTO `json:"controllerServices"`

	// The funnels in this flow snippet.
	// Unique: true
	Funnels []*FunnelDTO `json:"funnels"`

	// The input ports in this flow snippet.
	// Unique: true
	InputPorts []*PortDTO `json:"inputPorts"`

	// The labels in this flow snippet.
	// Unique: true
	Labels []*LabelDTO `json:"labels"`

	// The output ports in this flow snippet.
	// Unique: true
	OutputPorts []*PortDTO `json:"outputPorts"`

	// The process groups in this flow snippet.
	// Unique: true
	ProcessGroups []*ProcessGroupDTO `json:"processGroups"`

	// The processors in this flow snippet.
	// Unique: true
	Processors []*ProcessorDTO `json:"processors"`

	// The remote process groups in this flow snippet.
	// Unique: true
	RemoteProcessGroups []*RemoteProcessGroupDTO `json:"remoteProcessGroups"`
}

// Validate validates this flow snippet d t o
func (m *FlowSnippetDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnections(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateControllerServices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFunnels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInputPorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutputPorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemoteProcessGroups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FlowSnippetDTO) validateConnections(formats strfmt.Registry) error {
	if swag.IsZero(m.Connections) { // not required
		return nil
	}

	if err := validate.UniqueItems("connections", "body", m.Connections); err != nil {
		return err
	}

	for i := 0; i < len(m.Connections); i++ {
		if swag.IsZero(m.Connections[i]) { // not required
			continue
		}

		if m.Connections[i] != nil {
			if err := m.Connections[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connections" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("connections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FlowSnippetDTO) validateControllerServices(formats strfmt.Registry) error {
	if swag.IsZero(m.ControllerServices) { // not required
		return nil
	}

	if err := validate.UniqueItems("controllerServices", "body", m.ControllerServices); err != nil {
		return err
	}

	for i := 0; i < len(m.ControllerServices); i++ {
		if swag.IsZero(m.ControllerServices[i]) { // not required
			continue
		}

		if m.ControllerServices[i] != nil {
			if err := m.ControllerServices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("controllerServices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("controllerServices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FlowSnippetDTO) validateFunnels(formats strfmt.Registry) error {
	if swag.IsZero(m.Funnels) { // not required
		return nil
	}

	if err := validate.UniqueItems("funnels", "body", m.Funnels); err != nil {
		return err
	}

	for i := 0; i < len(m.Funnels); i++ {
		if swag.IsZero(m.Funnels[i]) { // not required
			continue
		}

		if m.Funnels[i] != nil {
			if err := m.Funnels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("funnels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("funnels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FlowSnippetDTO) validateInputPorts(formats strfmt.Registry) error {
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

func (m *FlowSnippetDTO) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	if err := validate.UniqueItems("labels", "body", m.Labels); err != nil {
		return err
	}

	for i := 0; i < len(m.Labels); i++ {
		if swag.IsZero(m.Labels[i]) { // not required
			continue
		}

		if m.Labels[i] != nil {
			if err := m.Labels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FlowSnippetDTO) validateOutputPorts(formats strfmt.Registry) error {
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

func (m *FlowSnippetDTO) validateProcessGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.ProcessGroups) { // not required
		return nil
	}

	if err := validate.UniqueItems("processGroups", "body", m.ProcessGroups); err != nil {
		return err
	}

	for i := 0; i < len(m.ProcessGroups); i++ {
		if swag.IsZero(m.ProcessGroups[i]) { // not required
			continue
		}

		if m.ProcessGroups[i] != nil {
			if err := m.ProcessGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("processGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("processGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FlowSnippetDTO) validateProcessors(formats strfmt.Registry) error {
	if swag.IsZero(m.Processors) { // not required
		return nil
	}

	if err := validate.UniqueItems("processors", "body", m.Processors); err != nil {
		return err
	}

	for i := 0; i < len(m.Processors); i++ {
		if swag.IsZero(m.Processors[i]) { // not required
			continue
		}

		if m.Processors[i] != nil {
			if err := m.Processors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("processors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("processors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FlowSnippetDTO) validateRemoteProcessGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.RemoteProcessGroups) { // not required
		return nil
	}

	if err := validate.UniqueItems("remoteProcessGroups", "body", m.RemoteProcessGroups); err != nil {
		return err
	}

	for i := 0; i < len(m.RemoteProcessGroups); i++ {
		if swag.IsZero(m.RemoteProcessGroups[i]) { // not required
			continue
		}

		if m.RemoteProcessGroups[i] != nil {
			if err := m.RemoteProcessGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("remoteProcessGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("remoteProcessGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this flow snippet d t o based on the context it is used
func (m *FlowSnippetDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnections(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateControllerServices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFunnels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInputPorts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOutputPorts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProcessGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProcessors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRemoteProcessGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FlowSnippetDTO) contextValidateConnections(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Connections); i++ {

		if m.Connections[i] != nil {

			if swag.IsZero(m.Connections[i]) { // not required
				return nil
			}

			if err := m.Connections[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connections" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("connections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FlowSnippetDTO) contextValidateControllerServices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ControllerServices); i++ {

		if m.ControllerServices[i] != nil {

			if swag.IsZero(m.ControllerServices[i]) { // not required
				return nil
			}

			if err := m.ControllerServices[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("controllerServices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("controllerServices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FlowSnippetDTO) contextValidateFunnels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Funnels); i++ {

		if m.Funnels[i] != nil {

			if swag.IsZero(m.Funnels[i]) { // not required
				return nil
			}

			if err := m.Funnels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("funnels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("funnels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FlowSnippetDTO) contextValidateInputPorts(ctx context.Context, formats strfmt.Registry) error {

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

func (m *FlowSnippetDTO) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Labels); i++ {

		if m.Labels[i] != nil {

			if swag.IsZero(m.Labels[i]) { // not required
				return nil
			}

			if err := m.Labels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FlowSnippetDTO) contextValidateOutputPorts(ctx context.Context, formats strfmt.Registry) error {

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

func (m *FlowSnippetDTO) contextValidateProcessGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ProcessGroups); i++ {

		if m.ProcessGroups[i] != nil {

			if swag.IsZero(m.ProcessGroups[i]) { // not required
				return nil
			}

			if err := m.ProcessGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("processGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("processGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FlowSnippetDTO) contextValidateProcessors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Processors); i++ {

		if m.Processors[i] != nil {

			if swag.IsZero(m.Processors[i]) { // not required
				return nil
			}

			if err := m.Processors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("processors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("processors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FlowSnippetDTO) contextValidateRemoteProcessGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RemoteProcessGroups); i++ {

		if m.RemoteProcessGroups[i] != nil {

			if swag.IsZero(m.RemoteProcessGroups[i]) { // not required
				return nil
			}

			if err := m.RemoteProcessGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("remoteProcessGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("remoteProcessGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FlowSnippetDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlowSnippetDTO) UnmarshalBinary(b []byte) error {
	var res FlowSnippetDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
