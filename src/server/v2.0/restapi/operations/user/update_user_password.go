// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateUserPasswordHandlerFunc turns a function with the right signature into a update user password handler
type UpdateUserPasswordHandlerFunc func(UpdateUserPasswordParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateUserPasswordHandlerFunc) Handle(params UpdateUserPasswordParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UpdateUserPasswordHandler interface for that can handle valid update user password params
type UpdateUserPasswordHandler interface {
	Handle(UpdateUserPasswordParams, interface{}) middleware.Responder
}

// NewUpdateUserPassword creates a new http.Handler for the update user password operation
func NewUpdateUserPassword(ctx *middleware.Context, handler UpdateUserPasswordHandler) *UpdateUserPassword {
	return &UpdateUserPassword{Context: ctx, Handler: handler}
}

/*UpdateUserPassword swagger:route PUT /users/{user_id}/password user updateUserPassword

Change the password on a user that already exists.

This endpoint is for user to update password. Users with the admin role can change any user's password. Regular users can change only their own password.


*/
type UpdateUserPassword struct {
	Context *middleware.Context
	Handler UpdateUserPasswordHandler
}

func (o *UpdateUserPassword) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateUserPasswordParams()

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
