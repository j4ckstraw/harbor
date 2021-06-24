// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// UpdateUserPasswordOKCode is the HTTP code returned for type UpdateUserPasswordOK
const UpdateUserPasswordOKCode int = 200

/*UpdateUserPasswordOK Success

swagger:response updateUserPasswordOK
*/
type UpdateUserPasswordOK struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`
}

// NewUpdateUserPasswordOK creates UpdateUserPasswordOK with default headers values
func NewUpdateUserPasswordOK() *UpdateUserPasswordOK {

	return &UpdateUserPasswordOK{}
}

// WithXRequestID adds the xRequestId to the update user password o k response
func (o *UpdateUserPasswordOK) WithXRequestID(xRequestID string) *UpdateUserPasswordOK {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the update user password o k response
func (o *UpdateUserPasswordOK) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WriteResponse to the client
func (o *UpdateUserPasswordOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// UpdateUserPasswordBadRequestCode is the HTTP code returned for type UpdateUserPasswordBadRequest
const UpdateUserPasswordBadRequestCode int = 400

/*UpdateUserPasswordBadRequest Invalid user ID; Password does not meet requirement

swagger:response updateUserPasswordBadRequest
*/
type UpdateUserPasswordBadRequest struct {
}

// NewUpdateUserPasswordBadRequest creates UpdateUserPasswordBadRequest with default headers values
func NewUpdateUserPasswordBadRequest() *UpdateUserPasswordBadRequest {

	return &UpdateUserPasswordBadRequest{}
}

// WriteResponse to the client
func (o *UpdateUserPasswordBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// UpdateUserPasswordUnauthorizedCode is the HTTP code returned for type UpdateUserPasswordUnauthorized
const UpdateUserPasswordUnauthorizedCode int = 401

/*UpdateUserPasswordUnauthorized Unauthorized

swagger:response updateUserPasswordUnauthorized
*/
type UpdateUserPasswordUnauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdateUserPasswordUnauthorized creates UpdateUserPasswordUnauthorized with default headers values
func NewUpdateUserPasswordUnauthorized() *UpdateUserPasswordUnauthorized {

	return &UpdateUserPasswordUnauthorized{}
}

// WithXRequestID adds the xRequestId to the update user password unauthorized response
func (o *UpdateUserPasswordUnauthorized) WithXRequestID(xRequestID string) *UpdateUserPasswordUnauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the update user password unauthorized response
func (o *UpdateUserPasswordUnauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the update user password unauthorized response
func (o *UpdateUserPasswordUnauthorized) WithPayload(payload *models.Errors) *UpdateUserPasswordUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user password unauthorized response
func (o *UpdateUserPasswordUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserPasswordUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// UpdateUserPasswordForbiddenCode is the HTTP code returned for type UpdateUserPasswordForbidden
const UpdateUserPasswordForbiddenCode int = 403

/*UpdateUserPasswordForbidden The caller does not have permission to update the password of the user with given ID, or the old password in request body is not correct.

swagger:response updateUserPasswordForbidden
*/
type UpdateUserPasswordForbidden struct {
}

// NewUpdateUserPasswordForbidden creates UpdateUserPasswordForbidden with default headers values
func NewUpdateUserPasswordForbidden() *UpdateUserPasswordForbidden {

	return &UpdateUserPasswordForbidden{}
}

// WriteResponse to the client
func (o *UpdateUserPasswordForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// UpdateUserPasswordInternalServerErrorCode is the HTTP code returned for type UpdateUserPasswordInternalServerError
const UpdateUserPasswordInternalServerErrorCode int = 500

/*UpdateUserPasswordInternalServerError Internal server error

swagger:response updateUserPasswordInternalServerError
*/
type UpdateUserPasswordInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdateUserPasswordInternalServerError creates UpdateUserPasswordInternalServerError with default headers values
func NewUpdateUserPasswordInternalServerError() *UpdateUserPasswordInternalServerError {

	return &UpdateUserPasswordInternalServerError{}
}

// WithXRequestID adds the xRequestId to the update user password internal server error response
func (o *UpdateUserPasswordInternalServerError) WithXRequestID(xRequestID string) *UpdateUserPasswordInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the update user password internal server error response
func (o *UpdateUserPasswordInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the update user password internal server error response
func (o *UpdateUserPasswordInternalServerError) WithPayload(payload *models.Errors) *UpdateUserPasswordInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user password internal server error response
func (o *UpdateUserPasswordInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserPasswordInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
