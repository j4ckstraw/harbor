// Code generated by go-swagger; DO NOT EDIT.

package artifact

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateTagHandlerFunc turns a function with the right signature into a create tag handler
type CreateTagHandlerFunc func(CreateTagParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateTagHandlerFunc) Handle(params CreateTagParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateTagHandler interface for that can handle valid create tag params
type CreateTagHandler interface {
	Handle(CreateTagParams, interface{}) middleware.Responder
}

// NewCreateTag creates a new http.Handler for the create tag operation
func NewCreateTag(ctx *middleware.Context, handler CreateTagHandler) *CreateTag {
	return &CreateTag{Context: ctx, Handler: handler}
}

/*CreateTag swagger:route POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags artifact createTag

Create tag

Create a tag for the specified artifact

*/
type CreateTag struct {
	Context *middleware.Context
	Handler CreateTagHandler
}

func (o *CreateTag) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateTagParams()

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
