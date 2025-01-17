// Code generated by go-swagger; DO NOT EDIT.

package usergroup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// ListUserGroupsOKCode is the HTTP code returned for type ListUserGroupsOK
const ListUserGroupsOKCode int = 200

/*ListUserGroupsOK Get user group successfully.

swagger:response listUserGroupsOK
*/
type ListUserGroupsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.UserGroup `json:"body,omitempty"`
}

// NewListUserGroupsOK creates ListUserGroupsOK with default headers values
func NewListUserGroupsOK() *ListUserGroupsOK {

	return &ListUserGroupsOK{}
}

// WithPayload adds the payload to the list user groups o k response
func (o *ListUserGroupsOK) WithPayload(payload []*models.UserGroup) *ListUserGroupsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list user groups o k response
func (o *ListUserGroupsOK) SetPayload(payload []*models.UserGroup) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListUserGroupsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.UserGroup, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListUserGroupsUnauthorizedCode is the HTTP code returned for type ListUserGroupsUnauthorized
const ListUserGroupsUnauthorizedCode int = 401

/*ListUserGroupsUnauthorized Unauthorized

swagger:response listUserGroupsUnauthorized
*/
type ListUserGroupsUnauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewListUserGroupsUnauthorized creates ListUserGroupsUnauthorized with default headers values
func NewListUserGroupsUnauthorized() *ListUserGroupsUnauthorized {

	return &ListUserGroupsUnauthorized{}
}

// WithXRequestID adds the xRequestId to the list user groups unauthorized response
func (o *ListUserGroupsUnauthorized) WithXRequestID(xRequestID string) *ListUserGroupsUnauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the list user groups unauthorized response
func (o *ListUserGroupsUnauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the list user groups unauthorized response
func (o *ListUserGroupsUnauthorized) WithPayload(payload *models.Errors) *ListUserGroupsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list user groups unauthorized response
func (o *ListUserGroupsUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListUserGroupsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListUserGroupsForbiddenCode is the HTTP code returned for type ListUserGroupsForbidden
const ListUserGroupsForbiddenCode int = 403

/*ListUserGroupsForbidden Forbidden

swagger:response listUserGroupsForbidden
*/
type ListUserGroupsForbidden struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewListUserGroupsForbidden creates ListUserGroupsForbidden with default headers values
func NewListUserGroupsForbidden() *ListUserGroupsForbidden {

	return &ListUserGroupsForbidden{}
}

// WithXRequestID adds the xRequestId to the list user groups forbidden response
func (o *ListUserGroupsForbidden) WithXRequestID(xRequestID string) *ListUserGroupsForbidden {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the list user groups forbidden response
func (o *ListUserGroupsForbidden) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the list user groups forbidden response
func (o *ListUserGroupsForbidden) WithPayload(payload *models.Errors) *ListUserGroupsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list user groups forbidden response
func (o *ListUserGroupsForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListUserGroupsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListUserGroupsInternalServerErrorCode is the HTTP code returned for type ListUserGroupsInternalServerError
const ListUserGroupsInternalServerErrorCode int = 500

/*ListUserGroupsInternalServerError Internal server error

swagger:response listUserGroupsInternalServerError
*/
type ListUserGroupsInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewListUserGroupsInternalServerError creates ListUserGroupsInternalServerError with default headers values
func NewListUserGroupsInternalServerError() *ListUserGroupsInternalServerError {

	return &ListUserGroupsInternalServerError{}
}

// WithXRequestID adds the xRequestId to the list user groups internal server error response
func (o *ListUserGroupsInternalServerError) WithXRequestID(xRequestID string) *ListUserGroupsInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the list user groups internal server error response
func (o *ListUserGroupsInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the list user groups internal server error response
func (o *ListUserGroupsInternalServerError) WithPayload(payload *models.Errors) *ListUserGroupsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list user groups internal server error response
func (o *ListUserGroupsInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListUserGroupsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
