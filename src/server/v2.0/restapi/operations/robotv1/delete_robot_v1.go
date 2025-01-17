// Code generated by go-swagger; DO NOT EDIT.

package robotv1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteRobotV1HandlerFunc turns a function with the right signature into a delete robot v1 handler
type DeleteRobotV1HandlerFunc func(DeleteRobotV1Params, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteRobotV1HandlerFunc) Handle(params DeleteRobotV1Params, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteRobotV1Handler interface for that can handle valid delete robot v1 params
type DeleteRobotV1Handler interface {
	Handle(DeleteRobotV1Params, interface{}) middleware.Responder
}

// NewDeleteRobotV1 creates a new http.Handler for the delete robot v1 operation
func NewDeleteRobotV1(ctx *middleware.Context, handler DeleteRobotV1Handler) *DeleteRobotV1 {
	return &DeleteRobotV1{Context: ctx, Handler: handler}
}

/*DeleteRobotV1 swagger:route DELETE /projects/{project_name_or_id}/robots/{robot_id} robotv1 deleteRobotV1

Delete a robot account

This endpoint deletes specific robot account information by robot ID.

*/
type DeleteRobotV1 struct {
	Context *middleware.Context
	Handler DeleteRobotV1Handler
}

func (o *DeleteRobotV1) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteRobotV1Params()

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
