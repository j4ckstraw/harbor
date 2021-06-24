// Code generated by go-swagger; DO NOT EDIT.

package configure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetConfigurationsHandlerFunc turns a function with the right signature into a get configurations handler
type GetConfigurationsHandlerFunc func(GetConfigurationsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetConfigurationsHandlerFunc) Handle(params GetConfigurationsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetConfigurationsHandler interface for that can handle valid get configurations params
type GetConfigurationsHandler interface {
	Handle(GetConfigurationsParams, interface{}) middleware.Responder
}

// NewGetConfigurations creates a new http.Handler for the get configurations operation
func NewGetConfigurations(ctx *middleware.Context, handler GetConfigurationsHandler) *GetConfigurations {
	return &GetConfigurations{Context: ctx, Handler: handler}
}

/*GetConfigurations swagger:route GET /configurations configure getConfigurations

Get system configurations.

This endpoint is for retrieving system configurations that only provides for admin user.


*/
type GetConfigurations struct {
	Context *middleware.Context
	Handler GetConfigurationsHandler
}

func (o *GetConfigurations) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetConfigurationsParams()

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
