// Code generated by go-swagger; DO NOT EDIT.

package webhook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// LastTriggerOKCode is the HTTP code returned for type LastTriggerOK
const LastTriggerOKCode int = 200

/*LastTriggerOK Test webhook connection successfully.

swagger:response lastTriggerOK
*/
type LastTriggerOK struct {

	/*
	  In: Body
	*/
	Payload []*models.WebhookLastTrigger `json:"body,omitempty"`
}

// NewLastTriggerOK creates LastTriggerOK with default headers values
func NewLastTriggerOK() *LastTriggerOK {

	return &LastTriggerOK{}
}

// WithPayload adds the payload to the last trigger o k response
func (o *LastTriggerOK) WithPayload(payload []*models.WebhookLastTrigger) *LastTriggerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the last trigger o k response
func (o *LastTriggerOK) SetPayload(payload []*models.WebhookLastTrigger) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LastTriggerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.WebhookLastTrigger, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// LastTriggerBadRequestCode is the HTTP code returned for type LastTriggerBadRequest
const LastTriggerBadRequestCode int = 400

/*LastTriggerBadRequest Bad request

swagger:response lastTriggerBadRequest
*/
type LastTriggerBadRequest struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewLastTriggerBadRequest creates LastTriggerBadRequest with default headers values
func NewLastTriggerBadRequest() *LastTriggerBadRequest {

	return &LastTriggerBadRequest{}
}

// WithXRequestID adds the xRequestId to the last trigger bad request response
func (o *LastTriggerBadRequest) WithXRequestID(xRequestID string) *LastTriggerBadRequest {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the last trigger bad request response
func (o *LastTriggerBadRequest) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the last trigger bad request response
func (o *LastTriggerBadRequest) WithPayload(payload *models.Errors) *LastTriggerBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the last trigger bad request response
func (o *LastTriggerBadRequest) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LastTriggerBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// LastTriggerUnauthorizedCode is the HTTP code returned for type LastTriggerUnauthorized
const LastTriggerUnauthorizedCode int = 401

/*LastTriggerUnauthorized Unauthorized

swagger:response lastTriggerUnauthorized
*/
type LastTriggerUnauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewLastTriggerUnauthorized creates LastTriggerUnauthorized with default headers values
func NewLastTriggerUnauthorized() *LastTriggerUnauthorized {

	return &LastTriggerUnauthorized{}
}

// WithXRequestID adds the xRequestId to the last trigger unauthorized response
func (o *LastTriggerUnauthorized) WithXRequestID(xRequestID string) *LastTriggerUnauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the last trigger unauthorized response
func (o *LastTriggerUnauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the last trigger unauthorized response
func (o *LastTriggerUnauthorized) WithPayload(payload *models.Errors) *LastTriggerUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the last trigger unauthorized response
func (o *LastTriggerUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LastTriggerUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// LastTriggerForbiddenCode is the HTTP code returned for type LastTriggerForbidden
const LastTriggerForbiddenCode int = 403

/*LastTriggerForbidden Forbidden

swagger:response lastTriggerForbidden
*/
type LastTriggerForbidden struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewLastTriggerForbidden creates LastTriggerForbidden with default headers values
func NewLastTriggerForbidden() *LastTriggerForbidden {

	return &LastTriggerForbidden{}
}

// WithXRequestID adds the xRequestId to the last trigger forbidden response
func (o *LastTriggerForbidden) WithXRequestID(xRequestID string) *LastTriggerForbidden {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the last trigger forbidden response
func (o *LastTriggerForbidden) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the last trigger forbidden response
func (o *LastTriggerForbidden) WithPayload(payload *models.Errors) *LastTriggerForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the last trigger forbidden response
func (o *LastTriggerForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LastTriggerForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// LastTriggerInternalServerErrorCode is the HTTP code returned for type LastTriggerInternalServerError
const LastTriggerInternalServerErrorCode int = 500

/*LastTriggerInternalServerError Internal server error

swagger:response lastTriggerInternalServerError
*/
type LastTriggerInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewLastTriggerInternalServerError creates LastTriggerInternalServerError with default headers values
func NewLastTriggerInternalServerError() *LastTriggerInternalServerError {

	return &LastTriggerInternalServerError{}
}

// WithXRequestID adds the xRequestId to the last trigger internal server error response
func (o *LastTriggerInternalServerError) WithXRequestID(xRequestID string) *LastTriggerInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the last trigger internal server error response
func (o *LastTriggerInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the last trigger internal server error response
func (o *LastTriggerInternalServerError) WithPayload(payload *models.Errors) *LastTriggerInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the last trigger internal server error response
func (o *LastTriggerInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LastTriggerInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
