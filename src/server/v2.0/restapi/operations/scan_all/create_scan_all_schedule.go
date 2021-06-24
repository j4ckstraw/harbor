// Code generated by go-swagger; DO NOT EDIT.

package scan_all

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateScanAllScheduleHandlerFunc turns a function with the right signature into a create scan all schedule handler
type CreateScanAllScheduleHandlerFunc func(CreateScanAllScheduleParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateScanAllScheduleHandlerFunc) Handle(params CreateScanAllScheduleParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateScanAllScheduleHandler interface for that can handle valid create scan all schedule params
type CreateScanAllScheduleHandler interface {
	Handle(CreateScanAllScheduleParams, interface{}) middleware.Responder
}

// NewCreateScanAllSchedule creates a new http.Handler for the create scan all schedule operation
func NewCreateScanAllSchedule(ctx *middleware.Context, handler CreateScanAllScheduleHandler) *CreateScanAllSchedule {
	return &CreateScanAllSchedule{Context: ctx, Handler: handler}
}

/*CreateScanAllSchedule swagger:route POST /system/scanAll/schedule scanAll createScanAllSchedule

Create a schedule or a manual trigger for the scan all job.

This endpoint is for creating a schedule or a manual trigger for the scan all job, which scans all of images in Harbor.

*/
type CreateScanAllSchedule struct {
	Context *middleware.Context
	Handler CreateScanAllScheduleHandler
}

func (o *CreateScanAllSchedule) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateScanAllScheduleParams()

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
