// Code generated by go-swagger; DO NOT EDIT.

package webhook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListWebhookPoliciesOfProjectHandlerFunc turns a function with the right signature into a list webhook policies of project handler
type ListWebhookPoliciesOfProjectHandlerFunc func(ListWebhookPoliciesOfProjectParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ListWebhookPoliciesOfProjectHandlerFunc) Handle(params ListWebhookPoliciesOfProjectParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ListWebhookPoliciesOfProjectHandler interface for that can handle valid list webhook policies of project params
type ListWebhookPoliciesOfProjectHandler interface {
	Handle(ListWebhookPoliciesOfProjectParams, interface{}) middleware.Responder
}

// NewListWebhookPoliciesOfProject creates a new http.Handler for the list webhook policies of project operation
func NewListWebhookPoliciesOfProject(ctx *middleware.Context, handler ListWebhookPoliciesOfProjectHandler) *ListWebhookPoliciesOfProject {
	return &ListWebhookPoliciesOfProject{Context: ctx, Handler: handler}
}

/*ListWebhookPoliciesOfProject swagger:route GET /projects/{project_name_or_id}/webhook/policies webhook listWebhookPoliciesOfProject

List project webhook policies.

This endpoint returns webhook policies of a project.


*/
type ListWebhookPoliciesOfProject struct {
	Context *middleware.Context
	Handler ListWebhookPoliciesOfProjectHandler
}

func (o *ListWebhookPoliciesOfProject) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListWebhookPoliciesOfProjectParams()

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
