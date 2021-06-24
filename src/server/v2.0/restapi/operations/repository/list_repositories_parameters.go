// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewListRepositoriesParams creates a new ListRepositoriesParams object
// with the default values initialized.
func NewListRepositoriesParams() ListRepositoriesParams {

	var (
		// initialize parameters with default values

		pageDefault     = int64(1)
		pageSizeDefault = int64(10)
	)

	return ListRepositoriesParams{
		Page: &pageDefault,

		PageSize: &pageSizeDefault,
	}
}

// ListRepositoriesParams contains all the bound params for the list repositories operation
// typically these are obtained from a http.Request
//
// swagger:parameters listRepositories
type ListRepositoriesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*An unique ID for the request
	  Min Length: 1
	  In: header
	*/
	XRequestID *string
	/*The page number
	  In: query
	  Default: 1
	*/
	Page *int64
	/*The size of per page
	  Maximum: 100
	  In: query
	  Default: 10
	*/
	PageSize *int64
	/*The name of the project
	  Required: true
	  In: path
	*/
	ProjectName string
	/*Query string to query resources. Supported query patterns are "exact match(k=v)", "fuzzy match(k=~v)", "range(k=[min~max])", "list with union releationship(k={v1 v2 v3})" and "list with intersetion relationship(k=(v1 v2 v3))". The value of range and list can be string(enclosed by " or '), integer or time(in format "2020-04-09 02:36:00"). All of these query patterns should be put in the query string "q=xxx" and splitted by ",". e.g. q=k1=v1,k2=~v2,k3=[min~max]
	  In: query
	*/
	Q *string
	/*Sort the resource list in ascending or descending order. e.g. sort by field1 in ascending orderr and field2 in descending order with "sort=field1,-field2"
	  In: query
	*/
	Sort *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewListRepositoriesParams() beforehand.
func (o *ListRepositoriesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if err := o.bindXRequestID(r.Header[http.CanonicalHeaderKey("X-Request-Id")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	qPage, qhkPage, _ := qs.GetOK("page")
	if err := o.bindPage(qPage, qhkPage, route.Formats); err != nil {
		res = append(res, err)
	}

	qPageSize, qhkPageSize, _ := qs.GetOK("page_size")
	if err := o.bindPageSize(qPageSize, qhkPageSize, route.Formats); err != nil {
		res = append(res, err)
	}

	rProjectName, rhkProjectName, _ := route.Params.GetOK("project_name")
	if err := o.bindProjectName(rProjectName, rhkProjectName, route.Formats); err != nil {
		res = append(res, err)
	}

	qQ, qhkQ, _ := qs.GetOK("q")
	if err := o.bindQ(qQ, qhkQ, route.Formats); err != nil {
		res = append(res, err)
	}

	qSort, qhkSort, _ := qs.GetOK("sort")
	if err := o.bindSort(qSort, qhkSort, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindXRequestID binds and validates parameter XRequestID from header.
func (o *ListRepositoriesParams) bindXRequestID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *ListRepositoriesParams) validateXRequestID(formats strfmt.Registry) error {

	if err := validate.MinLength("X-Request-Id", "header", (*o.XRequestID), 1); err != nil {
		return err
	}

	return nil
}

// bindPage binds and validates parameter Page from query.
func (o *ListRepositoriesParams) bindPage(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewListRepositoriesParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("page", "query", "int64", raw)
	}
	o.Page = &value

	return nil
}

// bindPageSize binds and validates parameter PageSize from query.
func (o *ListRepositoriesParams) bindPageSize(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewListRepositoriesParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("page_size", "query", "int64", raw)
	}
	o.PageSize = &value

	if err := o.validatePageSize(formats); err != nil {
		return err
	}

	return nil
}

// validatePageSize carries on validations for parameter PageSize
func (o *ListRepositoriesParams) validatePageSize(formats strfmt.Registry) error {

	if err := validate.MaximumInt("page_size", "query", int64(*o.PageSize), 100, false); err != nil {
		return err
	}

	return nil
}

// bindProjectName binds and validates parameter ProjectName from path.
func (o *ListRepositoriesParams) bindProjectName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ProjectName = raw

	return nil
}

// bindQ binds and validates parameter Q from query.
func (o *ListRepositoriesParams) bindQ(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Q = &raw

	return nil
}

// bindSort binds and validates parameter Sort from query.
func (o *ListRepositoriesParams) bindSort(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Sort = &raw

	return nil
}
