// Code generated by go-swagger; DO NOT EDIT.

package project_metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// GetProjectMetadataURL generates an URL for the get project metadata operation
type GetProjectMetadataURL struct {
	MetaName        string
	ProjectNameOrID string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetProjectMetadataURL) WithBasePath(bp string) *GetProjectMetadataURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetProjectMetadataURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetProjectMetadataURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/projects/{project_name_or_id}/metadatas/{meta_name}"

	metaName := o.MetaName
	if metaName != "" {
		_path = strings.Replace(_path, "{meta_name}", metaName, -1)
	} else {
		return nil, errors.New("metaName is required on GetProjectMetadataURL")
	}

	projectNameOrID := o.ProjectNameOrID
	if projectNameOrID != "" {
		_path = strings.Replace(_path, "{project_name_or_id}", projectNameOrID, -1)
	} else {
		return nil, errors.New("projectNameOrId is required on GetProjectMetadataURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/v2.0"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetProjectMetadataURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetProjectMetadataURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetProjectMetadataURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetProjectMetadataURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetProjectMetadataURL")
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
func (o *GetProjectMetadataURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
