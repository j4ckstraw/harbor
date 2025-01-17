// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// UpdateUserProfileOKCode is the HTTP code returned for type UpdateUserProfileOK
const UpdateUserProfileOKCode int = 200

/*UpdateUserProfileOK Success

swagger:response updateUserProfileOK
*/
type UpdateUserProfileOK struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`
}

// NewUpdateUserProfileOK creates UpdateUserProfileOK with default headers values
func NewUpdateUserProfileOK() *UpdateUserProfileOK {

	return &UpdateUserProfileOK{}
}

// WithXRequestID adds the xRequestId to the update user profile o k response
func (o *UpdateUserProfileOK) WithXRequestID(xRequestID string) *UpdateUserProfileOK {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the update user profile o k response
func (o *UpdateUserProfileOK) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WriteResponse to the client
func (o *UpdateUserProfileOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// UpdateUserProfileUnauthorizedCode is the HTTP code returned for type UpdateUserProfileUnauthorized
const UpdateUserProfileUnauthorizedCode int = 401

/*UpdateUserProfileUnauthorized Unauthorized

swagger:response updateUserProfileUnauthorized
*/
type UpdateUserProfileUnauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdateUserProfileUnauthorized creates UpdateUserProfileUnauthorized with default headers values
func NewUpdateUserProfileUnauthorized() *UpdateUserProfileUnauthorized {

	return &UpdateUserProfileUnauthorized{}
}

// WithXRequestID adds the xRequestId to the update user profile unauthorized response
func (o *UpdateUserProfileUnauthorized) WithXRequestID(xRequestID string) *UpdateUserProfileUnauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the update user profile unauthorized response
func (o *UpdateUserProfileUnauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the update user profile unauthorized response
func (o *UpdateUserProfileUnauthorized) WithPayload(payload *models.Errors) *UpdateUserProfileUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user profile unauthorized response
func (o *UpdateUserProfileUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserProfileUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// UpdateUserProfileForbiddenCode is the HTTP code returned for type UpdateUserProfileForbidden
const UpdateUserProfileForbiddenCode int = 403

/*UpdateUserProfileForbidden Forbidden

swagger:response updateUserProfileForbidden
*/
type UpdateUserProfileForbidden struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdateUserProfileForbidden creates UpdateUserProfileForbidden with default headers values
func NewUpdateUserProfileForbidden() *UpdateUserProfileForbidden {

	return &UpdateUserProfileForbidden{}
}

// WithXRequestID adds the xRequestId to the update user profile forbidden response
func (o *UpdateUserProfileForbidden) WithXRequestID(xRequestID string) *UpdateUserProfileForbidden {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the update user profile forbidden response
func (o *UpdateUserProfileForbidden) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the update user profile forbidden response
func (o *UpdateUserProfileForbidden) WithPayload(payload *models.Errors) *UpdateUserProfileForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user profile forbidden response
func (o *UpdateUserProfileForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserProfileForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// UpdateUserProfileNotFoundCode is the HTTP code returned for type UpdateUserProfileNotFound
const UpdateUserProfileNotFoundCode int = 404

/*UpdateUserProfileNotFound Not found

swagger:response updateUserProfileNotFound
*/
type UpdateUserProfileNotFound struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdateUserProfileNotFound creates UpdateUserProfileNotFound with default headers values
func NewUpdateUserProfileNotFound() *UpdateUserProfileNotFound {

	return &UpdateUserProfileNotFound{}
}

// WithXRequestID adds the xRequestId to the update user profile not found response
func (o *UpdateUserProfileNotFound) WithXRequestID(xRequestID string) *UpdateUserProfileNotFound {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the update user profile not found response
func (o *UpdateUserProfileNotFound) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the update user profile not found response
func (o *UpdateUserProfileNotFound) WithPayload(payload *models.Errors) *UpdateUserProfileNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user profile not found response
func (o *UpdateUserProfileNotFound) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserProfileNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// UpdateUserProfileInternalServerErrorCode is the HTTP code returned for type UpdateUserProfileInternalServerError
const UpdateUserProfileInternalServerErrorCode int = 500

/*UpdateUserProfileInternalServerError Internal server error

swagger:response updateUserProfileInternalServerError
*/
type UpdateUserProfileInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdateUserProfileInternalServerError creates UpdateUserProfileInternalServerError with default headers values
func NewUpdateUserProfileInternalServerError() *UpdateUserProfileInternalServerError {

	return &UpdateUserProfileInternalServerError{}
}

// WithXRequestID adds the xRequestId to the update user profile internal server error response
func (o *UpdateUserProfileInternalServerError) WithXRequestID(xRequestID string) *UpdateUserProfileInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the update user profile internal server error response
func (o *UpdateUserProfileInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the update user profile internal server error response
func (o *UpdateUserProfileInternalServerError) WithPayload(payload *models.Errors) *UpdateUserProfileInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user profile internal server error response
func (o *UpdateUserProfileInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserProfileInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
