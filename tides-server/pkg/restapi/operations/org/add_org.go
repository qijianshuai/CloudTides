// Code generated by go-swagger; DO NOT EDIT.

package org

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AddOrgHandlerFunc turns a function with the right signature into a add org handler
type AddOrgHandlerFunc func(AddOrgParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddOrgHandlerFunc) Handle(params AddOrgParams) middleware.Responder {
	return fn(params)
}

// AddOrgHandler interface for that can handle valid add org params
type AddOrgHandler interface {
	Handle(AddOrgParams) middleware.Responder
}

// NewAddOrg creates a new http.Handler for the add org operation
func NewAddOrg(ctx *middleware.Context, handler AddOrgHandler) *AddOrg {
	return &AddOrg{Context: ctx, Handler: handler}
}

/* AddOrg swagger:route POST /org org addOrg

add Org

*/
type AddOrg struct {
	Context *middleware.Context
	Handler AddOrgHandler
}

func (o *AddOrg) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAddOrgParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// AddOrgBody add org body
//
// swagger:model AddOrgBody
type AddOrgBody struct {

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this add org body
func (o *AddOrgBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add org body based on context it is used
func (o *AddOrgBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddOrgBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddOrgBody) UnmarshalBinary(b []byte) error {
	var res AddOrgBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// AddOrgForbiddenBody add org forbidden body
//
// swagger:model AddOrgForbiddenBody
type AddOrgForbiddenBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this add org forbidden body
func (o *AddOrgForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add org forbidden body based on context it is used
func (o *AddOrgForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddOrgForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddOrgForbiddenBody) UnmarshalBinary(b []byte) error {
	var res AddOrgForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// AddOrgOKBody add org o k body
//
// swagger:model AddOrgOKBody
type AddOrgOKBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this add org o k body
func (o *AddOrgOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add org o k body based on context it is used
func (o *AddOrgOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddOrgOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddOrgOKBody) UnmarshalBinary(b []byte) error {
	var res AddOrgOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// AddOrgUnauthorizedBody add org unauthorized body
//
// swagger:model AddOrgUnauthorizedBody
type AddOrgUnauthorizedBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this add org unauthorized body
func (o *AddOrgUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add org unauthorized body based on context it is used
func (o *AddOrgUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddOrgUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddOrgUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res AddOrgUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
