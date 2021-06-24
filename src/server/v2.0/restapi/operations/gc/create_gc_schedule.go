// Code generated by go-swagger; DO NOT EDIT.

package gc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateGCScheduleHandlerFunc turns a function with the right signature into a create GC schedule handler
type CreateGCScheduleHandlerFunc func(CreateGCScheduleParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateGCScheduleHandlerFunc) Handle(params CreateGCScheduleParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateGCScheduleHandler interface for that can handle valid create GC schedule params
type CreateGCScheduleHandler interface {
	Handle(CreateGCScheduleParams, interface{}) middleware.Responder
}

// NewCreateGCSchedule creates a new http.Handler for the create GC schedule operation
func NewCreateGCSchedule(ctx *middleware.Context, handler CreateGCScheduleHandler) *CreateGCSchedule {
	return &CreateGCSchedule{Context: ctx, Handler: handler}
}

/*CreateGCSchedule swagger:route POST /system/gc/schedule gc createGcSchedule

Create a gc schedule.

This endpoint is for update gc schedule.


*/
type CreateGCSchedule struct {
	Context *middleware.Context
	Handler CreateGCScheduleHandler
}

func (o *CreateGCSchedule) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateGCScheduleParams()

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
