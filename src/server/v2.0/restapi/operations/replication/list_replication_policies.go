// Code generated by go-swagger; DO NOT EDIT.

package replication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListReplicationPoliciesHandlerFunc turns a function with the right signature into a list replication policies handler
type ListReplicationPoliciesHandlerFunc func(ListReplicationPoliciesParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ListReplicationPoliciesHandlerFunc) Handle(params ListReplicationPoliciesParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ListReplicationPoliciesHandler interface for that can handle valid list replication policies params
type ListReplicationPoliciesHandler interface {
	Handle(ListReplicationPoliciesParams, interface{}) middleware.Responder
}

// NewListReplicationPolicies creates a new http.Handler for the list replication policies operation
func NewListReplicationPolicies(ctx *middleware.Context, handler ListReplicationPoliciesHandler) *ListReplicationPolicies {
	return &ListReplicationPolicies{Context: ctx, Handler: handler}
}

/*ListReplicationPolicies swagger:route GET /replication/policies replication listReplicationPolicies

List replication policies

List replication policies

*/
type ListReplicationPolicies struct {
	Context *middleware.Context
	Handler ListReplicationPoliciesHandler
}

func (o *ListReplicationPolicies) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListReplicationPoliciesParams()

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
