// Code generated by go-swagger; DO NOT EDIT.

package scanner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteScannerHandlerFunc turns a function with the right signature into a delete scanner handler
type DeleteScannerHandlerFunc func(DeleteScannerParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteScannerHandlerFunc) Handle(params DeleteScannerParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteScannerHandler interface for that can handle valid delete scanner params
type DeleteScannerHandler interface {
	Handle(DeleteScannerParams, interface{}) middleware.Responder
}

// NewDeleteScanner creates a new http.Handler for the delete scanner operation
func NewDeleteScanner(ctx *middleware.Context, handler DeleteScannerHandler) *DeleteScanner {
	return &DeleteScanner{Context: ctx, Handler: handler}
}

/*DeleteScanner swagger:route DELETE /scanners/{registration_id} scanner deleteScanner

Delete a scanner registration

Deletes the specified scanner registration.


*/
type DeleteScanner struct {
	Context *middleware.Context
	Handler DeleteScannerHandler
}

func (o *DeleteScanner) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteScannerParams()

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
