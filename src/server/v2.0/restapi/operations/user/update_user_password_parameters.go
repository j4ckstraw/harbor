// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewUpdateUserPasswordParams creates a new UpdateUserPasswordParams object
// no default values defined in spec.
func NewUpdateUserPasswordParams() UpdateUserPasswordParams {

	return UpdateUserPasswordParams{}
}

// UpdateUserPasswordParams contains all the bound params for the update user password operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateUserPassword
type UpdateUserPasswordParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*An unique ID for the request
	  Min Length: 1
	  In: header
	*/
	XRequestID *string
	/*Password to be updated, the attribute 'old_password' is optional when the API is called by the system administrator.
	  Required: true
	  In: body
	*/
	Password *models.PasswordReq
	/*
	  Required: true
	  In: path
	*/
	UserID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateUserPasswordParams() beforehand.
func (o *UpdateUserPasswordParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindXRequestID(r.Header[http.CanonicalHeaderKey("X-Request-Id")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.PasswordReq
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("password", "body", ""))
			} else {
				res = append(res, errors.NewParseError("password", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Password = &body
			}
		}
	} else {
		res = append(res, errors.Required("password", "body", ""))
	}
	rUserID, rhkUserID, _ := route.Params.GetOK("user_id")
	if err := o.bindUserID(rUserID, rhkUserID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindXRequestID binds and validates parameter XRequestID from header.
func (o *UpdateUserPasswordParams) bindXRequestID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *UpdateUserPasswordParams) validateXRequestID(formats strfmt.Registry) error {

	if err := validate.MinLength("X-Request-Id", "header", (*o.XRequestID), 1); err != nil {
		return err
	}

	return nil
}

// bindUserID binds and validates parameter UserID from path.
func (o *UpdateUserPasswordParams) bindUserID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("user_id", "path", "int64", raw)
	}
	o.UserID = value

	return nil
}
