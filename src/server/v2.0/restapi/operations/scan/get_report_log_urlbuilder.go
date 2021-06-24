// Code generated by go-swagger; DO NOT EDIT.

package scan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// GetReportLogURL generates an URL for the get report log operation
type GetReportLogURL struct {
	ProjectName    string
	Reference      string
	ReportID       string
	RepositoryName string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetReportLogURL) WithBasePath(bp string) *GetReportLogURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetReportLogURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetReportLogURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/scan/{report_id}/log"

	projectName := o.ProjectName
	if projectName != "" {
		_path = strings.Replace(_path, "{project_name}", projectName, -1)
	} else {
		return nil, errors.New("projectName is required on GetReportLogURL")
	}

	reference := o.Reference
	if reference != "" {
		_path = strings.Replace(_path, "{reference}", reference, -1)
	} else {
		return nil, errors.New("reference is required on GetReportLogURL")
	}

	reportID := o.ReportID
	if reportID != "" {
		_path = strings.Replace(_path, "{report_id}", reportID, -1)
	} else {
		return nil, errors.New("reportId is required on GetReportLogURL")
	}

	repositoryName := o.RepositoryName
	if repositoryName != "" {
		_path = strings.Replace(_path, "{repository_name}", repositoryName, -1)
	} else {
		return nil, errors.New("repositoryName is required on GetReportLogURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/v2.0"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetReportLogURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetReportLogURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetReportLogURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetReportLogURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetReportLogURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetReportLogURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
