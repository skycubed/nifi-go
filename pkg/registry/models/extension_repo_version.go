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

// ExtensionRepoVersion extension repo version
//
// swagger:model ExtensionRepoVersion
type ExtensionRepoVersion struct {

	// The WebLink to download this version of the extension bundle.
	// Read Only: true
	DownloadLink *JaxbLink `json:"downloadLink,omitempty"`

	// The WebLink to view the metadata about the extensions contained in the extension bundle.
	// Read Only: true
	ExtensionsLink *JaxbLink `json:"extensionsLink,omitempty"`

	// The WebLink to retrieve the SHA-256 digest for this version of the extension bundle.
	// Read Only: true
	Sha256Link *JaxbLink `json:"sha256Link,omitempty"`

	// Indicates if the client supplied a SHA-256 when uploading this version of the extension bundle.
	// Read Only: true
	Sha256Supplied *JaxbLink `json:"sha256Supplied,omitempty"`
}

// Validate validates this extension repo version
func (m *ExtensionRepoVersion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDownloadLink(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtensionsLink(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSha256Link(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSha256Supplied(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExtensionRepoVersion) validateDownloadLink(formats strfmt.Registry) error {
	if swag.IsZero(m.DownloadLink) { // not required
		return nil
	}

	if m.DownloadLink != nil {
		if err := m.DownloadLink.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("downloadLink")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("downloadLink")
			}
			return err
		}
	}

	return nil
}

func (m *ExtensionRepoVersion) validateExtensionsLink(formats strfmt.Registry) error {
	if swag.IsZero(m.ExtensionsLink) { // not required
		return nil
	}

	if m.ExtensionsLink != nil {
		if err := m.ExtensionsLink.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("extensionsLink")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("extensionsLink")
			}
			return err
		}
	}

	return nil
}

func (m *ExtensionRepoVersion) validateSha256Link(formats strfmt.Registry) error {
	if swag.IsZero(m.Sha256Link) { // not required
		return nil
	}

	if m.Sha256Link != nil {
		if err := m.Sha256Link.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sha256Link")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sha256Link")
			}
			return err
		}
	}

	return nil
}

func (m *ExtensionRepoVersion) validateSha256Supplied(formats strfmt.Registry) error {
	if swag.IsZero(m.Sha256Supplied) { // not required
		return nil
	}

	if m.Sha256Supplied != nil {
		if err := m.Sha256Supplied.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sha256Supplied")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sha256Supplied")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this extension repo version based on the context it is used
func (m *ExtensionRepoVersion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDownloadLink(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExtensionsLink(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSha256Link(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSha256Supplied(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExtensionRepoVersion) contextValidateDownloadLink(ctx context.Context, formats strfmt.Registry) error {

	if m.DownloadLink != nil {
		if err := m.DownloadLink.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("downloadLink")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("downloadLink")
			}
			return err
		}
	}

	return nil
}

func (m *ExtensionRepoVersion) contextValidateExtensionsLink(ctx context.Context, formats strfmt.Registry) error {

	if m.ExtensionsLink != nil {
		if err := m.ExtensionsLink.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("extensionsLink")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("extensionsLink")
			}
			return err
		}
	}

	return nil
}

func (m *ExtensionRepoVersion) contextValidateSha256Link(ctx context.Context, formats strfmt.Registry) error {

	if m.Sha256Link != nil {
		if err := m.Sha256Link.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sha256Link")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sha256Link")
			}
			return err
		}
	}

	return nil
}

func (m *ExtensionRepoVersion) contextValidateSha256Supplied(ctx context.Context, formats strfmt.Registry) error {

	if m.Sha256Supplied != nil {
		if err := m.Sha256Supplied.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sha256Supplied")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sha256Supplied")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExtensionRepoVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExtensionRepoVersion) UnmarshalBinary(b []byte) error {
	var res ExtensionRepoVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
