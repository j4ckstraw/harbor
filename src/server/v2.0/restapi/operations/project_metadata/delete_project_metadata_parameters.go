// Code generated by go-swagger; DO NOT EDIT.

package project_metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewDeleteProjectMetadataParams creates a new DeleteProjectMetadataParams object
// with the default values initialized.
func NewDeleteProjectMetadataParams() DeleteProjectMetadataParams {

	var (
		// initialize parameters with default values

		xIsResourceNameDefault = bool(false)
	)

	return DeleteProjectMetadataParams{
		XIsResourceName: &xIsResourceNameDefault,
	}
}

// DeleteProjectMetadataParams contains all the bound params for the delete project metadata operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteProjectMetadata
type DeleteProjectMetadataParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

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
	/*The name of metadata.
	  Required: true
	  In: path
	*/
	MetaName string
	/*The name or id of the project
	  Required: true
	  In: path
	*/
	ProjectNameOrID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteProjectMetadataParams() beforehand.
func (o *DeleteProjectMetadataParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindXIsResourceName(r.Header[http.CanonicalHeaderKey("X-Is-Resource-Name")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindXRequestID(r.Header[http.CanonicalHeaderKey("X-Request-Id")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	rMetaName, rhkMetaName, _ := route.Params.GetOK("meta_name")
	if err := o.bindMetaName(rMetaName, rhkMetaName, route.Formats); err != nil {
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
func (o *DeleteProjectMetadataParams) bindXIsResourceName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewDeleteProjectMetadataParams()
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
func (o *DeleteProjectMetadataParams) bindXRequestID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *DeleteProjectMetadataParams) validateXRequestID(formats strfmt.Registry) error {

	if err := validate.MinLength("X-Request-Id", "header", (*o.XRequestID), 1); err != nil {
		return err
	}

	return nil
}

// bindMetaName binds and validates parameter MetaName from path.
func (o *DeleteProjectMetadataParams) bindMetaName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.MetaName = raw

	return nil
}

// bindProjectNameOrID binds and validates parameter ProjectNameOrID from path.
func (o *DeleteProjectMetadataParams) bindProjectNameOrID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ProjectNameOrID = raw

	return nil
}
