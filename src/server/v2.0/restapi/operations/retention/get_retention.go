// Code generated by go-swagger; DO NOT EDIT.

package retention

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetRetentionHandlerFunc turns a function with the right signature into a get retention handler
type GetRetentionHandlerFunc func(GetRetentionParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRetentionHandlerFunc) Handle(params GetRetentionParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetRetentionHandler interface for that can handle valid get retention params
type GetRetentionHandler interface {
	Handle(GetRetentionParams, interface{}) middleware.Responder
}

// NewGetRetention creates a new http.Handler for the get retention operation
func NewGetRetention(ctx *middleware.Context, handler GetRetentionHandler) *GetRetention {
	return &GetRetention{Context: ctx, Handler: handler}
}

/*GetRetention swagger:route GET /retentions/{id} Retention getRetention

Get Retention Policy

Get Retention Policy.

*/
type GetRetention struct {
	Context *middleware.Context
	Handler GetRetentionHandler
}

func (o *GetRetention) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRetentionParams()

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
