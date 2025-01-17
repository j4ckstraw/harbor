// Code generated by go-swagger; DO NOT EDIT.

package retention

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// DeleteRetentionOKCode is the HTTP code returned for type DeleteRetentionOK
const DeleteRetentionOKCode int = 200

/*DeleteRetentionOK Update Retention Policy successfully.

swagger:response deleteRetentionOK
*/
type DeleteRetentionOK struct {
}

// NewDeleteRetentionOK creates DeleteRetentionOK with default headers values
func NewDeleteRetentionOK() *DeleteRetentionOK {

	return &DeleteRetentionOK{}
}

// WriteResponse to the client
func (o *DeleteRetentionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeleteRetentionUnauthorizedCode is the HTTP code returned for type DeleteRetentionUnauthorized
const DeleteRetentionUnauthorizedCode int = 401

/*DeleteRetentionUnauthorized Unauthorized

swagger:response deleteRetentionUnauthorized
*/
type DeleteRetentionUnauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewDeleteRetentionUnauthorized creates DeleteRetentionUnauthorized with default headers values
func NewDeleteRetentionUnauthorized() *DeleteRetentionUnauthorized {

	return &DeleteRetentionUnauthorized{}
}

// WithXRequestID adds the xRequestId to the delete retention unauthorized response
func (o *DeleteRetentionUnauthorized) WithXRequestID(xRequestID string) *DeleteRetentionUnauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the delete retention unauthorized response
func (o *DeleteRetentionUnauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the delete retention unauthorized response
func (o *DeleteRetentionUnauthorized) WithPayload(payload *models.Errors) *DeleteRetentionUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete retention unauthorized response
func (o *DeleteRetentionUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteRetentionUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// DeleteRetentionForbiddenCode is the HTTP code returned for type DeleteRetentionForbidden
const DeleteRetentionForbiddenCode int = 403

/*DeleteRetentionForbidden Forbidden

swagger:response deleteRetentionForbidden
*/
type DeleteRetentionForbidden struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewDeleteRetentionForbidden creates DeleteRetentionForbidden with default headers values
func NewDeleteRetentionForbidden() *DeleteRetentionForbidden {

	return &DeleteRetentionForbidden{}
}

// WithXRequestID adds the xRequestId to the delete retention forbidden response
func (o *DeleteRetentionForbidden) WithXRequestID(xRequestID string) *DeleteRetentionForbidden {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the delete retention forbidden response
func (o *DeleteRetentionForbidden) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the delete retention forbidden response
func (o *DeleteRetentionForbidden) WithPayload(payload *models.Errors) *DeleteRetentionForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete retention forbidden response
func (o *DeleteRetentionForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteRetentionForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// DeleteRetentionInternalServerErrorCode is the HTTP code returned for type DeleteRetentionInternalServerError
const DeleteRetentionInternalServerErrorCode int = 500

/*DeleteRetentionInternalServerError Internal server error

swagger:response deleteRetentionInternalServerError
*/
type DeleteRetentionInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewDeleteRetentionInternalServerError creates DeleteRetentionInternalServerError with default headers values
func NewDeleteRetentionInternalServerError() *DeleteRetentionInternalServerError {

	return &DeleteRetentionInternalServerError{}
}

// WithXRequestID adds the xRequestId to the delete retention internal server error response
func (o *DeleteRetentionInternalServerError) WithXRequestID(xRequestID string) *DeleteRetentionInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the delete retention internal server error response
func (o *DeleteRetentionInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the delete retention internal server error response
func (o *DeleteRetentionInternalServerError) WithPayload(payload *models.Errors) *DeleteRetentionInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete retention internal server error response
func (o *DeleteRetentionInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteRetentionInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
