// Code generated by go-swagger; DO NOT EDIT.

package scan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ScanArtifactHandlerFunc turns a function with the right signature into a scan artifact handler
type ScanArtifactHandlerFunc func(ScanArtifactParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ScanArtifactHandlerFunc) Handle(params ScanArtifactParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ScanArtifactHandler interface for that can handle valid scan artifact params
type ScanArtifactHandler interface {
	Handle(ScanArtifactParams, interface{}) middleware.Responder
}

// NewScanArtifact creates a new http.Handler for the scan artifact operation
func NewScanArtifact(ctx *middleware.Context, handler ScanArtifactHandler) *ScanArtifact {
	return &ScanArtifact{Context: ctx, Handler: handler}
}

/*ScanArtifact swagger:route POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/scan scan scanArtifact

Scan the artifact

Scan the specified artifact

*/
type ScanArtifact struct {
	Context *middleware.Context
	Handler ScanArtifactHandler
}

func (o *ScanArtifact) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewScanArtifactParams()

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
