// Code generated by go-swagger; DO NOT EDIT.

package flow

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetControllerServiceTypesParams creates a new GetControllerServiceTypesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetControllerServiceTypesParams() *GetControllerServiceTypesParams {
	return &GetControllerServiceTypesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetControllerServiceTypesParamsWithTimeout creates a new GetControllerServiceTypesParams object
// with the ability to set a timeout on a request.
func NewGetControllerServiceTypesParamsWithTimeout(timeout time.Duration) *GetControllerServiceTypesParams {
	return &GetControllerServiceTypesParams{
		timeout: timeout,
	}
}

// NewGetControllerServiceTypesParamsWithContext creates a new GetControllerServiceTypesParams object
// with the ability to set a context for a request.
func NewGetControllerServiceTypesParamsWithContext(ctx context.Context) *GetControllerServiceTypesParams {
	return &GetControllerServiceTypesParams{
		Context: ctx,
	}
}

// NewGetControllerServiceTypesParamsWithHTTPClient creates a new GetControllerServiceTypesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetControllerServiceTypesParamsWithHTTPClient(client *http.Client) *GetControllerServiceTypesParams {
	return &GetControllerServiceTypesParams{
		HTTPClient: client,
	}
}

/*
GetControllerServiceTypesParams contains all the parameters to send to the API endpoint

	for the get controller service types operation.

	Typically these are written to a http.Request.
*/
type GetControllerServiceTypesParams struct {

	/* BundleArtifactFilter.

	   If specified, will only return types that are a member of this bundle artifact.
	*/
	BundleArtifactFilter *string

	/* BundleGroupFilter.

	   If specified, will only return types that are a member of this bundle group.
	*/
	BundleGroupFilter *string

	/* ServiceBundleArtifact.

	   If serviceType specified, is the bundle artifact of the serviceType.
	*/
	ServiceBundleArtifact *string

	/* ServiceBundleGroup.

	   If serviceType specified, is the bundle group of the serviceType.
	*/
	ServiceBundleGroup *string

	/* ServiceBundleVersion.

	   If serviceType specified, is the bundle version of the serviceType.
	*/
	ServiceBundleVersion *string

	/* ServiceType.

	   If specified, will only return controller services that are compatible with this type of service.
	*/
	ServiceType *string

	/* TypeFilter.

	   If specified, will only return types whose fully qualified classname matches.
	*/
	TypeFilter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get controller service types params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetControllerServiceTypesParams) WithDefaults() *GetControllerServiceTypesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get controller service types params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetControllerServiceTypesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get controller service types params
func (o *GetControllerServiceTypesParams) WithTimeout(timeout time.Duration) *GetControllerServiceTypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get controller service types params
func (o *GetControllerServiceTypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get controller service types params
func (o *GetControllerServiceTypesParams) WithContext(ctx context.Context) *GetControllerServiceTypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get controller service types params
func (o *GetControllerServiceTypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get controller service types params
func (o *GetControllerServiceTypesParams) WithHTTPClient(client *http.Client) *GetControllerServiceTypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get controller service types params
func (o *GetControllerServiceTypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBundleArtifactFilter adds the bundleArtifactFilter to the get controller service types params
func (o *GetControllerServiceTypesParams) WithBundleArtifactFilter(bundleArtifactFilter *string) *GetControllerServiceTypesParams {
	o.SetBundleArtifactFilter(bundleArtifactFilter)
	return o
}

// SetBundleArtifactFilter adds the bundleArtifactFilter to the get controller service types params
func (o *GetControllerServiceTypesParams) SetBundleArtifactFilter(bundleArtifactFilter *string) {
	o.BundleArtifactFilter = bundleArtifactFilter
}

// WithBundleGroupFilter adds the bundleGroupFilter to the get controller service types params
func (o *GetControllerServiceTypesParams) WithBundleGroupFilter(bundleGroupFilter *string) *GetControllerServiceTypesParams {
	o.SetBundleGroupFilter(bundleGroupFilter)
	return o
}

// SetBundleGroupFilter adds the bundleGroupFilter to the get controller service types params
func (o *GetControllerServiceTypesParams) SetBundleGroupFilter(bundleGroupFilter *string) {
	o.BundleGroupFilter = bundleGroupFilter
}

// WithServiceBundleArtifact adds the serviceBundleArtifact to the get controller service types params
func (o *GetControllerServiceTypesParams) WithServiceBundleArtifact(serviceBundleArtifact *string) *GetControllerServiceTypesParams {
	o.SetServiceBundleArtifact(serviceBundleArtifact)
	return o
}

// SetServiceBundleArtifact adds the serviceBundleArtifact to the get controller service types params
func (o *GetControllerServiceTypesParams) SetServiceBundleArtifact(serviceBundleArtifact *string) {
	o.ServiceBundleArtifact = serviceBundleArtifact
}

// WithServiceBundleGroup adds the serviceBundleGroup to the get controller service types params
func (o *GetControllerServiceTypesParams) WithServiceBundleGroup(serviceBundleGroup *string) *GetControllerServiceTypesParams {
	o.SetServiceBundleGroup(serviceBundleGroup)
	return o
}

// SetServiceBundleGroup adds the serviceBundleGroup to the get controller service types params
func (o *GetControllerServiceTypesParams) SetServiceBundleGroup(serviceBundleGroup *string) {
	o.ServiceBundleGroup = serviceBundleGroup
}

// WithServiceBundleVersion adds the serviceBundleVersion to the get controller service types params
func (o *GetControllerServiceTypesParams) WithServiceBundleVersion(serviceBundleVersion *string) *GetControllerServiceTypesParams {
	o.SetServiceBundleVersion(serviceBundleVersion)
	return o
}

// SetServiceBundleVersion adds the serviceBundleVersion to the get controller service types params
func (o *GetControllerServiceTypesParams) SetServiceBundleVersion(serviceBundleVersion *string) {
	o.ServiceBundleVersion = serviceBundleVersion
}

// WithServiceType adds the serviceType to the get controller service types params
func (o *GetControllerServiceTypesParams) WithServiceType(serviceType *string) *GetControllerServiceTypesParams {
	o.SetServiceType(serviceType)
	return o
}

// SetServiceType adds the serviceType to the get controller service types params
func (o *GetControllerServiceTypesParams) SetServiceType(serviceType *string) {
	o.ServiceType = serviceType
}

// WithTypeFilter adds the typeFilter to the get controller service types params
func (o *GetControllerServiceTypesParams) WithTypeFilter(typeFilter *string) *GetControllerServiceTypesParams {
	o.SetTypeFilter(typeFilter)
	return o
}

// SetTypeFilter adds the typeFilter to the get controller service types params
func (o *GetControllerServiceTypesParams) SetTypeFilter(typeFilter *string) {
	o.TypeFilter = typeFilter
}

// WriteToRequest writes these params to a swagger request
func (o *GetControllerServiceTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BundleArtifactFilter != nil {

		// query param bundleArtifactFilter
		var qrBundleArtifactFilter string

		if o.BundleArtifactFilter != nil {
			qrBundleArtifactFilter = *o.BundleArtifactFilter
		}
		qBundleArtifactFilter := qrBundleArtifactFilter
		if qBundleArtifactFilter != "" {

			if err := r.SetQueryParam("bundleArtifactFilter", qBundleArtifactFilter); err != nil {
				return err
			}
		}
	}

	if o.BundleGroupFilter != nil {

		// query param bundleGroupFilter
		var qrBundleGroupFilter string

		if o.BundleGroupFilter != nil {
			qrBundleGroupFilter = *o.BundleGroupFilter
		}
		qBundleGroupFilter := qrBundleGroupFilter
		if qBundleGroupFilter != "" {

			if err := r.SetQueryParam("bundleGroupFilter", qBundleGroupFilter); err != nil {
				return err
			}
		}
	}

	if o.ServiceBundleArtifact != nil {

		// query param serviceBundleArtifact
		var qrServiceBundleArtifact string

		if o.ServiceBundleArtifact != nil {
			qrServiceBundleArtifact = *o.ServiceBundleArtifact
		}
		qServiceBundleArtifact := qrServiceBundleArtifact
		if qServiceBundleArtifact != "" {

			if err := r.SetQueryParam("serviceBundleArtifact", qServiceBundleArtifact); err != nil {
				return err
			}
		}
	}

	if o.ServiceBundleGroup != nil {

		// query param serviceBundleGroup
		var qrServiceBundleGroup string

		if o.ServiceBundleGroup != nil {
			qrServiceBundleGroup = *o.ServiceBundleGroup
		}
		qServiceBundleGroup := qrServiceBundleGroup
		if qServiceBundleGroup != "" {

			if err := r.SetQueryParam("serviceBundleGroup", qServiceBundleGroup); err != nil {
				return err
			}
		}
	}

	if o.ServiceBundleVersion != nil {

		// query param serviceBundleVersion
		var qrServiceBundleVersion string

		if o.ServiceBundleVersion != nil {
			qrServiceBundleVersion = *o.ServiceBundleVersion
		}
		qServiceBundleVersion := qrServiceBundleVersion
		if qServiceBundleVersion != "" {

			if err := r.SetQueryParam("serviceBundleVersion", qServiceBundleVersion); err != nil {
				return err
			}
		}
	}

	if o.ServiceType != nil {

		// query param serviceType
		var qrServiceType string

		if o.ServiceType != nil {
			qrServiceType = *o.ServiceType
		}
		qServiceType := qrServiceType
		if qServiceType != "" {

			if err := r.SetQueryParam("serviceType", qServiceType); err != nil {
				return err
			}
		}
	}

	if o.TypeFilter != nil {

		// query param typeFilter
		var qrTypeFilter string

		if o.TypeFilter != nil {
			qrTypeFilter = *o.TypeFilter
		}
		qTypeFilter := qrTypeFilter
		if qTypeFilter != "" {

			if err := r.SetQueryParam("typeFilter", qTypeFilter); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
