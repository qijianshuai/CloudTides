// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AddProjectHandlerFunc turns a function with the right signature into a add project handler
type AddProjectHandlerFunc func(AddProjectParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddProjectHandlerFunc) Handle(params AddProjectParams) middleware.Responder {
	return fn(params)
}

// AddProjectHandler interface for that can handle valid add project params
type AddProjectHandler interface {
	Handle(AddProjectParams) middleware.Responder
}

// NewAddProject creates a new http.Handler for the add project operation
func NewAddProject(ctx *middleware.Context, handler AddProjectHandler) *AddProject {
	return &AddProject{Context: ctx, Handler: handler}
}

/* AddProject swagger:route POST /project project addProject

add boinc projects

*/
type AddProject struct {
	Context *middleware.Context
	Handler AddProjectHandler
}

func (o *AddProject) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddProjectParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// AddProjectBody add project body
//
// swagger:model AddProjectBody
type AddProjectBody struct {

	// has account manager
	HasAccountManager bool `json:"hasAccountManager,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this add project body
func (o *AddProjectBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add project body based on context it is used
func (o *AddProjectBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddProjectBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddProjectBody) UnmarshalBinary(b []byte) error {
	var res AddProjectBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// AddProjectOKBody add project o k body
//
// swagger:model AddProjectOKBody
type AddProjectOKBody struct {

	// id
	ID int64 `json:"id,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this add project o k body
func (o *AddProjectOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add project o k body based on context it is used
func (o *AddProjectOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddProjectOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddProjectOKBody) UnmarshalBinary(b []byte) error {
	var res AddProjectOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
