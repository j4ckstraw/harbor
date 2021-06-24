// Code generated by go-swagger; DO NOT EDIT.

package artifact

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// RemoveLabelHandlerFunc turns a function with the right signature into a remove label handler
type RemoveLabelHandlerFunc func(RemoveLabelParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn RemoveLabelHandlerFunc) Handle(params RemoveLabelParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// RemoveLabelHandler interface for that can handle valid remove label params
type RemoveLabelHandler interface {
	Handle(RemoveLabelParams, interface{}) middleware.Responder
}

// NewRemoveLabel creates a new http.Handler for the remove label operation
func NewRemoveLabel(ctx *middleware.Context, handler RemoveLabelHandler) *RemoveLabel {
	return &RemoveLabel{Context: ctx, Handler: handler}
}

/*RemoveLabel swagger:route DELETE /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/labels/{label_id} artifact removeLabel

Remove label from artifact

Remove the label from the specified artiact.

*/
type RemoveLabel struct {
	Context *middleware.Context
	Handler RemoveLabelHandler
}

func (o *RemoveLabel) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRemoveLabelParams()

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
