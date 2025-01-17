// Code generated by go-swagger; DO NOT EDIT.

package robotv1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateRobotV1HandlerFunc turns a function with the right signature into a update robot v1 handler
type UpdateRobotV1HandlerFunc func(UpdateRobotV1Params, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateRobotV1HandlerFunc) Handle(params UpdateRobotV1Params, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UpdateRobotV1Handler interface for that can handle valid update robot v1 params
type UpdateRobotV1Handler interface {
	Handle(UpdateRobotV1Params, interface{}) middleware.Responder
}

// NewUpdateRobotV1 creates a new http.Handler for the update robot v1 operation
func NewUpdateRobotV1(ctx *middleware.Context, handler UpdateRobotV1Handler) *UpdateRobotV1 {
	return &UpdateRobotV1{Context: ctx, Handler: handler}
}

/*UpdateRobotV1 swagger:route PUT /projects/{project_name_or_id}/robots/{robot_id} robotv1 updateRobotV1

Update status of robot account.

Used to disable/enable a specified robot account.

*/
type UpdateRobotV1 struct {
	Context *middleware.Context
	Handler UpdateRobotV1Handler
}

func (o *UpdateRobotV1) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateRobotV1Params()

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
