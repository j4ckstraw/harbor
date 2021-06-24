// Code generated by go-swagger; DO NOT EDIT.

package project_metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteProjectMetadataHandlerFunc turns a function with the right signature into a delete project metadata handler
type DeleteProjectMetadataHandlerFunc func(DeleteProjectMetadataParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteProjectMetadataHandlerFunc) Handle(params DeleteProjectMetadataParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteProjectMetadataHandler interface for that can handle valid delete project metadata params
type DeleteProjectMetadataHandler interface {
	Handle(DeleteProjectMetadataParams, interface{}) middleware.Responder
}

// NewDeleteProjectMetadata creates a new http.Handler for the delete project metadata operation
func NewDeleteProjectMetadata(ctx *middleware.Context, handler DeleteProjectMetadataHandler) *DeleteProjectMetadata {
	return &DeleteProjectMetadata{Context: ctx, Handler: handler}
}

/*DeleteProjectMetadata swagger:route DELETE /projects/{project_name_or_id}/metadatas/{meta_name} projectMetadata deleteProjectMetadata

Delete the specific metadata for the specific project

Delete the specific metadata for the specific project

*/
type DeleteProjectMetadata struct {
	Context *middleware.Context
	Handler DeleteProjectMetadataHandler
}

func (o *DeleteProjectMetadata) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteProjectMetadataParams()

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
