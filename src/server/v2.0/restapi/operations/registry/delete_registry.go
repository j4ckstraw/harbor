// Code generated by go-swagger; DO NOT EDIT.

package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteRegistryHandlerFunc turns a function with the right signature into a delete registry handler
type DeleteRegistryHandlerFunc func(DeleteRegistryParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteRegistryHandlerFunc) Handle(params DeleteRegistryParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteRegistryHandler interface for that can handle valid delete registry params
type DeleteRegistryHandler interface {
	Handle(DeleteRegistryParams, interface{}) middleware.Responder
}

// NewDeleteRegistry creates a new http.Handler for the delete registry operation
func NewDeleteRegistry(ctx *middleware.Context, handler DeleteRegistryHandler) *DeleteRegistry {
	return &DeleteRegistry{Context: ctx, Handler: handler}
}

/*DeleteRegistry swagger:route DELETE /registries/{id} registry deleteRegistry

Delete the specific registry

Delete the specific registry

*/
type DeleteRegistry struct {
	Context *middleware.Context
	Handler DeleteRegistryHandler
}

func (o *DeleteRegistry) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteRegistryParams()

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
