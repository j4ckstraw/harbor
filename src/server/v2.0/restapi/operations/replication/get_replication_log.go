// Code generated by go-swagger; DO NOT EDIT.

package replication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetReplicationLogHandlerFunc turns a function with the right signature into a get replication log handler
type GetReplicationLogHandlerFunc func(GetReplicationLogParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetReplicationLogHandlerFunc) Handle(params GetReplicationLogParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetReplicationLogHandler interface for that can handle valid get replication log params
type GetReplicationLogHandler interface {
	Handle(GetReplicationLogParams, interface{}) middleware.Responder
}

// NewGetReplicationLog creates a new http.Handler for the get replication log operation
func NewGetReplicationLog(ctx *middleware.Context, handler GetReplicationLogHandler) *GetReplicationLog {
	return &GetReplicationLog{Context: ctx, Handler: handler}
}

/*GetReplicationLog swagger:route GET /replication/executions/{id}/tasks/{task_id}/log replication getReplicationLog

Get the log of the specific replication task

Get the log of the specific replication task

*/
type GetReplicationLog struct {
	Context *middleware.Context
	Handler GetReplicationLogHandler
}

func (o *GetReplicationLog) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetReplicationLogParams()

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
