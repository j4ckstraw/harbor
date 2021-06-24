// Code generated by go-swagger; DO NOT EDIT.

package robotv1

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

// NewListRobotV1Params creates a new ListRobotV1Params object
// with the default values initialized.
func NewListRobotV1Params() ListRobotV1Params {

	var (
		// initialize parameters with default values

		xIsResourceNameDefault = bool(false)

		pageDefault     = int64(1)
		pageSizeDefault = int64(10)
	)

	return ListRobotV1Params{
		XIsResourceName: &xIsResourceNameDefault,

		Page: &pageDefault,

		PageSize: &pageSizeDefault,
	}
}

// ListRobotV1Params contains all the bound params for the list robot v1 operation
// typically these are obtained from a http.Request
//
// swagger:parameters ListRobotV1
type ListRobotV1Params struct {

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
	/*The name or id of the project
	  Required: true
	  In: path
	*/
	ProjectNameOrID string
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
// To ensure default values, the struct must have been initialized with NewListRobotV1Params() beforehand.
func (o *ListRobotV1Params) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if err := o.bindXIsResourceName(r.Header[http.CanonicalHeaderKey("X-Is-Resource-Name")], true, route.Formats); err != nil {
		res = append(res, err)
	}

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

	rProjectNameOrID, rhkProjectNameOrID, _ := route.Params.GetOK("project_name_or_id")
	if err := o.bindProjectNameOrID(rProjectNameOrID, rhkProjectNameOrID, route.Formats); err != nil {
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

// bindXIsResourceName binds and validates parameter XIsResourceName from header.
func (o *ListRobotV1Params) bindXIsResourceName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewListRobotV1Params()
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
func (o *ListRobotV1Params) bindXRequestID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *ListRobotV1Params) validateXRequestID(formats strfmt.Registry) error {

	if err := validate.MinLength("X-Request-Id", "header", (*o.XRequestID), 1); err != nil {
		return err
	}

	return nil
}

// bindPage binds and validates parameter Page from query.
func (o *ListRobotV1Params) bindPage(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewListRobotV1Params()
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
func (o *ListRobotV1Params) bindPageSize(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewListRobotV1Params()
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
func (o *ListRobotV1Params) validatePageSize(formats strfmt.Registry) error {

	if err := validate.MaximumInt("page_size", "query", int64(*o.PageSize), 100, false); err != nil {
		return err
	}

	return nil
}

// bindProjectNameOrID binds and validates parameter ProjectNameOrID from path.
func (o *ListRobotV1Params) bindProjectNameOrID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ProjectNameOrID = raw

	return nil
}

// bindQ binds and validates parameter Q from query.
func (o *ListRobotV1Params) bindQ(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *ListRobotV1Params) bindSort(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
