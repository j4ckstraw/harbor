// Code generated by go-swagger; DO NOT EDIT.

package retention

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteRetentionHandlerFunc turns a function with the right signature into a delete retention handler
type DeleteRetentionHandlerFunc func(DeleteRetentionParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteRetentionHandlerFunc) Handle(params DeleteRetentionParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteRetentionHandler interface for that can handle valid delete retention params
type DeleteRetentionHandler interface {
	Handle(DeleteRetentionParams, interface{}) middleware.Responder
}

// NewDeleteRetention creates a new http.Handler for the delete retention operation
func NewDeleteRetention(ctx *middleware.Context, handler DeleteRetentionHandler) *DeleteRetention {
	return &DeleteRetention{Context: ctx, Handler: handler}
}

/*DeleteRetention swagger:route DELETE /retentions/{id} Retention deleteRetention

Delete Retention Policy

Delete Retention Policy, you can reference metadatas API for the policy model. You can check project metadatas to find whether a retention policy is already binded. This method should only be called when retention policy has already binded to project.

*/
type DeleteRetention struct {
	Context *middleware.Context
	Handler DeleteRetentionHandler
}

func (o *DeleteRetention) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteRetentionParams()

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
