// Code generated by go-swagger; DO NOT EDIT.

package member

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetProjectMemberHandlerFunc turns a function with the right signature into a get project member handler
type GetProjectMemberHandlerFunc func(GetProjectMemberParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetProjectMemberHandlerFunc) Handle(params GetProjectMemberParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetProjectMemberHandler interface for that can handle valid get project member params
type GetProjectMemberHandler interface {
	Handle(GetProjectMemberParams, interface{}) middleware.Responder
}

// NewGetProjectMember creates a new http.Handler for the get project member operation
func NewGetProjectMember(ctx *middleware.Context, handler GetProjectMemberHandler) *GetProjectMember {
	return &GetProjectMember{Context: ctx, Handler: handler}
}

/*GetProjectMember swagger:route GET /projects/{project_name_or_id}/members/{mid} member getProjectMember

Get the project member information

Get the project member information

*/
type GetProjectMember struct {
	Context *middleware.Context
	Handler GetProjectMemberHandler
}

func (o *GetProjectMember) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetProjectMemberParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
