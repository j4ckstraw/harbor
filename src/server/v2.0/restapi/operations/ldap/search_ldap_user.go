// Code generated by go-swagger; DO NOT EDIT.

package ldap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// SearchLdapUserHandlerFunc turns a function with the right signature into a search ldap user handler
type SearchLdapUserHandlerFunc func(SearchLdapUserParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn SearchLdapUserHandlerFunc) Handle(params SearchLdapUserParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// SearchLdapUserHandler interface for that can handle valid search ldap user params
type SearchLdapUserHandler interface {
	Handle(SearchLdapUserParams, interface{}) middleware.Responder
}

// NewSearchLdapUser creates a new http.Handler for the search ldap user operation
func NewSearchLdapUser(ctx *middleware.Context, handler SearchLdapUserHandler) *SearchLdapUser {
	return &SearchLdapUser{Context: ctx, Handler: handler}
}

/*SearchLdapUser swagger:route GET /ldap/users/search Ldap searchLdapUser

Search available ldap users.

This endpoint searches the available ldap users based on related configuration parameters. Support searched by input ladp configuration, load configuration from the system and specific filter.


*/
type SearchLdapUser struct {
	Context *middleware.Context
	Handler SearchLdapUserHandler
}

func (o *SearchLdapUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSearchLdapUserParams()

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
