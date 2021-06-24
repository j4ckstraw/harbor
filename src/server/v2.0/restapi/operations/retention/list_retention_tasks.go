// Code generated by go-swagger; DO NOT EDIT.

package retention

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListRetentionTasksHandlerFunc turns a function with the right signature into a list retention tasks handler
type ListRetentionTasksHandlerFunc func(ListRetentionTasksParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ListRetentionTasksHandlerFunc) Handle(params ListRetentionTasksParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ListRetentionTasksHandler interface for that can handle valid list retention tasks params
type ListRetentionTasksHandler interface {
	Handle(ListRetentionTasksParams, interface{}) middleware.Responder
}

// NewListRetentionTasks creates a new http.Handler for the list retention tasks operation
func NewListRetentionTasks(ctx *middleware.Context, handler ListRetentionTasksHandler) *ListRetentionTasks {
	return &ListRetentionTasks{Context: ctx, Handler: handler}
}

/*ListRetentionTasks swagger:route GET /retentions/{id}/executions/{eid}/tasks Retention listRetentionTasks

Get Retention tasks

Get Retention tasks, each repository as a task.

*/
type ListRetentionTasks struct {
	Context *middleware.Context
	Handler ListRetentionTasksHandler
}

func (o *ListRetentionTasks) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListRetentionTasksParams()

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
