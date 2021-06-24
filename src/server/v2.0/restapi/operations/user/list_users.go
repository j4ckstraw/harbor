// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListUsersHandlerFunc turns a function with the right signature into a list users handler
type ListUsersHandlerFunc func(ListUsersParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ListUsersHandlerFunc) Handle(params ListUsersParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ListUsersHandler interface for that can handle valid list users params
type ListUsersHandler interface {
	Handle(ListUsersParams, interface{}) middleware.Responder
}

// NewListUsers creates a new http.Handler for the list users operation
func NewListUsers(ctx *middleware.Context, handler ListUsersHandler) *ListUsers {
	return &ListUsers{Context: ctx, Handler: handler}
}

/*ListUsers swagger:route GET /users user listUsers

List users

*/
type ListUsers struct {
	Context *middleware.Context
	Handler ListUsersHandler
}

func (o *ListUsers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListUsersParams()

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
