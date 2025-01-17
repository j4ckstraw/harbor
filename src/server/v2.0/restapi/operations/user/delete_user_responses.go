// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// DeleteUserOKCode is the HTTP code returned for type DeleteUserOK
const DeleteUserOKCode int = 200

/*DeleteUserOK Success

swagger:response deleteUserOK
*/
type DeleteUserOK struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`
}

// NewDeleteUserOK creates DeleteUserOK with default headers values
func NewDeleteUserOK() *DeleteUserOK {

	return &DeleteUserOK{}
}

// WithXRequestID adds the xRequestId to the delete user o k response
func (o *DeleteUserOK) WithXRequestID(xRequestID string) *DeleteUserOK {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the delete user o k response
func (o *DeleteUserOK) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WriteResponse to the client
func (o *DeleteUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeleteUserUnauthorizedCode is the HTTP code returned for type DeleteUserUnauthorized
const DeleteUserUnauthorizedCode int = 401

/*DeleteUserUnauthorized Unauthorized

swagger:response deleteUserUnauthorized
*/
type DeleteUserUnauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewDeleteUserUnauthorized creates DeleteUserUnauthorized with default headers values
func NewDeleteUserUnauthorized() *DeleteUserUnauthorized {

	return &DeleteUserUnauthorized{}
}

// WithXRequestID adds the xRequestId to the delete user unauthorized response
func (o *DeleteUserUnauthorized) WithXRequestID(xRequestID string) *DeleteUserUnauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the delete user unauthorized response
func (o *DeleteUserUnauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the delete user unauthorized response
func (o *DeleteUserUnauthorized) WithPayload(payload *models.Errors) *DeleteUserUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user unauthorized response
func (o *DeleteUserUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// DeleteUserForbiddenCode is the HTTP code returned for type DeleteUserForbidden
const DeleteUserForbiddenCode int = 403

/*DeleteUserForbidden Forbidden

swagger:response deleteUserForbidden
*/
type DeleteUserForbidden struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewDeleteUserForbidden creates DeleteUserForbidden with default headers values
func NewDeleteUserForbidden() *DeleteUserForbidden {

	return &DeleteUserForbidden{}
}

// WithXRequestID adds the xRequestId to the delete user forbidden response
func (o *DeleteUserForbidden) WithXRequestID(xRequestID string) *DeleteUserForbidden {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the delete user forbidden response
func (o *DeleteUserForbidden) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the delete user forbidden response
func (o *DeleteUserForbidden) WithPayload(payload *models.Errors) *DeleteUserForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user forbidden response
func (o *DeleteUserForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// DeleteUserNotFoundCode is the HTTP code returned for type DeleteUserNotFound
const DeleteUserNotFoundCode int = 404

/*DeleteUserNotFound Not found

swagger:response deleteUserNotFound
*/
type DeleteUserNotFound struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewDeleteUserNotFound creates DeleteUserNotFound with default headers values
func NewDeleteUserNotFound() *DeleteUserNotFound {

	return &DeleteUserNotFound{}
}

// WithXRequestID adds the xRequestId to the delete user not found response
func (o *DeleteUserNotFound) WithXRequestID(xRequestID string) *DeleteUserNotFound {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the delete user not found response
func (o *DeleteUserNotFound) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the delete user not found response
func (o *DeleteUserNotFound) WithPayload(payload *models.Errors) *DeleteUserNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user not found response
func (o *DeleteUserNotFound) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteUserInternalServerErrorCode is the HTTP code returned for type DeleteUserInternalServerError
const DeleteUserInternalServerErrorCode int = 500

/*DeleteUserInternalServerError Internal server error

swagger:response deleteUserInternalServerError
*/
type DeleteUserInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewDeleteUserInternalServerError creates DeleteUserInternalServerError with default headers values
func NewDeleteUserInternalServerError() *DeleteUserInternalServerError {

	return &DeleteUserInternalServerError{}
}

// WithXRequestID adds the xRequestId to the delete user internal server error response
func (o *DeleteUserInternalServerError) WithXRequestID(xRequestID string) *DeleteUserInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the delete user internal server error response
func (o *DeleteUserInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the delete user internal server error response
func (o *DeleteUserInternalServerError) WithPayload(payload *models.Errors) *DeleteUserInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user internal server error response
func (o *DeleteUserInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
