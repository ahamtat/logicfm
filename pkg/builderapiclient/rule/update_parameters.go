// Code generated by go-swagger; DO NOT EDIT.

package rule

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
	"github.com/go-openapi/swag"

	"github.com/ahamtat/logicfm/pkg/models"
)

// NewUpdateParams creates a new UpdateParams object
// with the default values initialized.
func NewUpdateParams() *UpdateParams {
	var ()
	return &UpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateParamsWithTimeout creates a new UpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateParamsWithTimeout(timeout time.Duration) *UpdateParams {
	var ()
	return &UpdateParams{

		timeout: timeout,
	}
}

// NewUpdateParamsWithContext creates a new UpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateParamsWithContext(ctx context.Context) *UpdateParams {
	var ()
	return &UpdateParams{

		Context: ctx,
	}
}

// NewUpdateParamsWithHTTPClient creates a new UpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateParamsWithHTTPClient(client *http.Client) *UpdateParams {
	var ()
	return &UpdateParams{
		HTTPClient: client,
	}
}

/*UpdateParams contains all the parameters to send to the API endpoint
for the update operation typically these are written to a http.Request
*/
type UpdateParams struct {

	/*ID*/
	ID int64
	/*Rule*/
	Rule *models.Rule

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update params
func (o *UpdateParams) WithTimeout(timeout time.Duration) *UpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update params
func (o *UpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update params
func (o *UpdateParams) WithContext(ctx context.Context) *UpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update params
func (o *UpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update params
func (o *UpdateParams) WithHTTPClient(client *http.Client) *UpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update params
func (o *UpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update params
func (o *UpdateParams) WithID(id int64) *UpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update params
func (o *UpdateParams) SetID(id int64) {
	o.ID = id
}

// WithRule adds the rule to the update params
func (o *UpdateParams) WithRule(rule *models.Rule) *UpdateParams {
	o.SetRule(rule)
	return o
}

// SetRule adds the rule to the update params
func (o *UpdateParams) SetRule(rule *models.Rule) {
	o.Rule = rule
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if o.Rule != nil {
		if err := r.SetBodyParam(o.Rule); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
