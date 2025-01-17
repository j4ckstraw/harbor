// Code generated by go-swagger; DO NOT EDIT.

package preheat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetExecutionHandlerFunc turns a function with the right signature into a get execution handler
type GetExecutionHandlerFunc func(GetExecutionParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetExecutionHandlerFunc) Handle(params GetExecutionParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetExecutionHandler interface for that can handle valid get execution params
type GetExecutionHandler interface {
	Handle(GetExecutionParams, interface{}) middleware.Responder
}

// NewGetExecution creates a new http.Handler for the get execution operation
func NewGetExecution(ctx *middleware.Context, handler GetExecutionHandler) *GetExecution {
	return &GetExecution{Context: ctx, Handler: handler}
}

/*GetExecution swagger:route GET /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id} preheat getExecution

Get a execution detail by id

Get a execution detail by id

*/
type GetExecution struct {
	Context *middleware.Context
	Handler GetExecutionHandler
}

func (o *GetExecution) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetExecutionParams()

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
