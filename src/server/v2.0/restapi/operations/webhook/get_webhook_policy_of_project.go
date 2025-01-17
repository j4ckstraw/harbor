// Code generated by go-swagger; DO NOT EDIT.

package webhook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetWebhookPolicyOfProjectHandlerFunc turns a function with the right signature into a get webhook policy of project handler
type GetWebhookPolicyOfProjectHandlerFunc func(GetWebhookPolicyOfProjectParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetWebhookPolicyOfProjectHandlerFunc) Handle(params GetWebhookPolicyOfProjectParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetWebhookPolicyOfProjectHandler interface for that can handle valid get webhook policy of project params
type GetWebhookPolicyOfProjectHandler interface {
	Handle(GetWebhookPolicyOfProjectParams, interface{}) middleware.Responder
}

// NewGetWebhookPolicyOfProject creates a new http.Handler for the get webhook policy of project operation
func NewGetWebhookPolicyOfProject(ctx *middleware.Context, handler GetWebhookPolicyOfProjectHandler) *GetWebhookPolicyOfProject {
	return &GetWebhookPolicyOfProject{Context: ctx, Handler: handler}
}

/*GetWebhookPolicyOfProject swagger:route GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id} webhook getWebhookPolicyOfProject

Get project webhook policy

This endpoint returns specified webhook policy of a project.


*/
type GetWebhookPolicyOfProject struct {
	Context *middleware.Context
	Handler GetWebhookPolicyOfProjectHandler
}

func (o *GetWebhookPolicyOfProject) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetWebhookPolicyOfProjectParams()

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
