// Code generated by go-swagger; DO NOT EDIT.

package immutable

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// NewUpdateImmuRuleParams creates a new UpdateImmuRuleParams object
// with the default values initialized.
func NewUpdateImmuRuleParams() UpdateImmuRuleParams {

	var (
		// initialize parameters with default values

		xIsResourceNameDefault = bool(false)
	)

	return UpdateImmuRuleParams{
		XIsResourceName: &xIsResourceNameDefault,
	}
}

// UpdateImmuRuleParams contains all the bound params for the update immu rule operation
// typically these are obtained from a http.Request
//
// swagger:parameters UpdateImmuRule
type UpdateImmuRuleParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	ImmutableRule *models.ImmutableRule
	/*The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name.
	  In: header
	  Default: false
	*/
	XIsResourceName *bool
	/*An unique ID for the request
	  Min Length: 1
	  In: header
	*/
	XRequestID *string
	/*The ID of the immutable rule
	  Required: true
	  In: path
	*/
	ImmutableRuleID int64
	/*The name or id of the project
	  Required: true
	  In: path
	*/
	ProjectNameOrID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateImmuRuleParams() beforehand.
func (o *UpdateImmuRuleParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.ImmutableRule
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("immutableRule", "body", ""))
			} else {
				res = append(res, errors.NewParseError("immutableRule", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.ImmutableRule = &body
			}
		}
	} else {
		res = append(res, errors.Required("immutableRule", "body", ""))
	}
	if err := o.bindXIsResourceName(r.Header[http.CanonicalHeaderKey("X-Is-Resource-Name")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindXRequestID(r.Header[http.CanonicalHeaderKey("X-Request-Id")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	rImmutableRuleID, rhkImmutableRuleID, _ := route.Params.GetOK("immutable_rule_id")
	if err := o.bindImmutableRuleID(rImmutableRuleID, rhkImmutableRuleID, route.Formats); err != nil {
		res = append(res, err)
	}

	rProjectNameOrID, rhkProjectNameOrID, _ := route.Params.GetOK("project_name_or_id")
	if err := o.bindProjectNameOrID(rProjectNameOrID, rhkProjectNameOrID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindXIsResourceName binds and validates parameter XIsResourceName from header.
func (o *UpdateImmuRuleParams) bindXIsResourceName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewUpdateImmuRuleParams()
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("X-Is-Resource-Name", "header", "bool", raw)
	}
	o.XIsResourceName = &value

	return nil
}

// bindXRequestID binds and validates parameter XRequestID from header.
func (o *UpdateImmuRuleParams) bindXRequestID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.XRequestID = &raw

	if err := o.validateXRequestID(formats); err != nil {
		return err
	}

	return nil
}

// validateXRequestID carries on validations for parameter XRequestID
func (o *UpdateImmuRuleParams) validateXRequestID(formats strfmt.Registry) error {

	if err := validate.MinLength("X-Request-Id", "header", (*o.XRequestID), 1); err != nil {
		return err
	}

	return nil
}

// bindImmutableRuleID binds and validates parameter ImmutableRuleID from path.
func (o *UpdateImmuRuleParams) bindImmutableRuleID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("immutable_rule_id", "path", "int64", raw)
	}
	o.ImmutableRuleID = value

	return nil
}

// bindProjectNameOrID binds and validates parameter ProjectNameOrID from path.
func (o *UpdateImmuRuleParams) bindProjectNameOrID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ProjectNameOrID = raw

	return nil
}
