// Code generated by go-swagger; DO NOT EDIT.

package preheat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListExecutionsHandlerFunc turns a function with the right signature into a list executions handler
type ListExecutionsHandlerFunc func(ListExecutionsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ListExecutionsHandlerFunc) Handle(params ListExecutionsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ListExecutionsHandler interface for that can handle valid list executions params
type ListExecutionsHandler interface {
	Handle(ListExecutionsParams, interface{}) middleware.Responder
}

// NewListExecutions creates a new http.Handler for the list executions operation
func NewListExecutions(ctx *middleware.Context, handler ListExecutionsHandler) *ListExecutions {
	return &ListExecutions{Context: ctx, Handler: handler}
}

/*ListExecutions swagger:route GET /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions preheat listExecutions

List executions for the given policy

List executions for the given policy

*/
type ListExecutions struct {
	Context *middleware.Context
	Handler ListExecutionsHandler
}

func (o *ListExecutions) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListExecutionsParams()

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
