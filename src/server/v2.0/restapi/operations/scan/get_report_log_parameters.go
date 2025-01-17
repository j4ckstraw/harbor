// Code generated by go-swagger; DO NOT EDIT.

package scan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewGetReportLogParams creates a new GetReportLogParams object
// no default values defined in spec.
func NewGetReportLogParams() GetReportLogParams {

	return GetReportLogParams{}
}

// GetReportLogParams contains all the bound params for the get report log operation
// typically these are obtained from a http.Request
//
// swagger:parameters getReportLog
type GetReportLogParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*An unique ID for the request
	  Min Length: 1
	  In: header
	*/
	XRequestID *string
	/*The name of the project
	  Required: true
	  In: path
	*/
	ProjectName string
	/*The reference of the artifact, can be digest or tag
	  Required: true
	  In: path
	*/
	Reference string
	/*The report id to get the log
	  Required: true
	  In: path
	*/
	ReportID string
	/*The name of the repository. If it contains slash, encode it with URL encoding. e.g. a/b -> a%252Fb
	  Required: true
	  In: path
	*/
	RepositoryName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetReportLogParams() beforehand.
func (o *GetReportLogParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindXRequestID(r.Header[http.CanonicalHeaderKey("X-Request-Id")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	rProjectName, rhkProjectName, _ := route.Params.GetOK("project_name")
	if err := o.bindProjectName(rProjectName, rhkProjectName, route.Formats); err != nil {
		res = append(res, err)
	}

	rReference, rhkReference, _ := route.Params.GetOK("reference")
	if err := o.bindReference(rReference, rhkReference, route.Formats); err != nil {
		res = append(res, err)
	}

	rReportID, rhkReportID, _ := route.Params.GetOK("report_id")
	if err := o.bindReportID(rReportID, rhkReportID, route.Formats); err != nil {
		res = append(res, err)
	}

	rRepositoryName, rhkRepositoryName, _ := route.Params.GetOK("repository_name")
	if err := o.bindRepositoryName(rRepositoryName, rhkRepositoryName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindXRequestID binds and validates parameter XRequestID from header.
func (o *GetReportLogParams) bindXRequestID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *GetReportLogParams) validateXRequestID(formats strfmt.Registry) error {

	if err := validate.MinLength("X-Request-Id", "header", (*o.XRequestID), 1); err != nil {
		return err
	}

	return nil
}

// bindProjectName binds and validates parameter ProjectName from path.
func (o *GetReportLogParams) bindProjectName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ProjectName = raw

	return nil
}

// bindReference binds and validates parameter Reference from path.
func (o *GetReportLogParams) bindReference(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Reference = raw

	return nil
}

// bindReportID binds and validates parameter ReportID from path.
func (o *GetReportLogParams) bindReportID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ReportID = raw

	return nil
}

// bindRepositoryName binds and validates parameter RepositoryName from path.
func (o *GetReportLogParams) bindRepositoryName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.RepositoryName = raw

	return nil
}
