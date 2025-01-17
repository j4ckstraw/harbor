// Code generated by go-swagger; DO NOT EDIT.

package system_cve_allowlist

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// PutSystemCVEAllowlistOKCode is the HTTP code returned for type PutSystemCVEAllowlistOK
const PutSystemCVEAllowlistOKCode int = 200

/*PutSystemCVEAllowlistOK Successfully updated the CVE allowlist.

swagger:response putSystemCveAllowlistOK
*/
type PutSystemCVEAllowlistOK struct {
}

// NewPutSystemCVEAllowlistOK creates PutSystemCVEAllowlistOK with default headers values
func NewPutSystemCVEAllowlistOK() *PutSystemCVEAllowlistOK {

	return &PutSystemCVEAllowlistOK{}
}

// WriteResponse to the client
func (o *PutSystemCVEAllowlistOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PutSystemCVEAllowlistUnauthorizedCode is the HTTP code returned for type PutSystemCVEAllowlistUnauthorized
const PutSystemCVEAllowlistUnauthorizedCode int = 401

/*PutSystemCVEAllowlistUnauthorized Unauthorized

swagger:response putSystemCveAllowlistUnauthorized
*/
type PutSystemCVEAllowlistUnauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewPutSystemCVEAllowlistUnauthorized creates PutSystemCVEAllowlistUnauthorized with default headers values
func NewPutSystemCVEAllowlistUnauthorized() *PutSystemCVEAllowlistUnauthorized {

	return &PutSystemCVEAllowlistUnauthorized{}
}

// WithXRequestID adds the xRequestId to the put system Cve allowlist unauthorized response
func (o *PutSystemCVEAllowlistUnauthorized) WithXRequestID(xRequestID string) *PutSystemCVEAllowlistUnauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the put system Cve allowlist unauthorized response
func (o *PutSystemCVEAllowlistUnauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the put system Cve allowlist unauthorized response
func (o *PutSystemCVEAllowlistUnauthorized) WithPayload(payload *models.Errors) *PutSystemCVEAllowlistUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put system Cve allowlist unauthorized response
func (o *PutSystemCVEAllowlistUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutSystemCVEAllowlistUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// PutSystemCVEAllowlistForbiddenCode is the HTTP code returned for type PutSystemCVEAllowlistForbidden
const PutSystemCVEAllowlistForbiddenCode int = 403

/*PutSystemCVEAllowlistForbidden Forbidden

swagger:response putSystemCveAllowlistForbidden
*/
type PutSystemCVEAllowlistForbidden struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewPutSystemCVEAllowlistForbidden creates PutSystemCVEAllowlistForbidden with default headers values
func NewPutSystemCVEAllowlistForbidden() *PutSystemCVEAllowlistForbidden {

	return &PutSystemCVEAllowlistForbidden{}
}

// WithXRequestID adds the xRequestId to the put system Cve allowlist forbidden response
func (o *PutSystemCVEAllowlistForbidden) WithXRequestID(xRequestID string) *PutSystemCVEAllowlistForbidden {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the put system Cve allowlist forbidden response
func (o *PutSystemCVEAllowlistForbidden) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the put system Cve allowlist forbidden response
func (o *PutSystemCVEAllowlistForbidden) WithPayload(payload *models.Errors) *PutSystemCVEAllowlistForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put system Cve allowlist forbidden response
func (o *PutSystemCVEAllowlistForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutSystemCVEAllowlistForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// PutSystemCVEAllowlistInternalServerErrorCode is the HTTP code returned for type PutSystemCVEAllowlistInternalServerError
const PutSystemCVEAllowlistInternalServerErrorCode int = 500

/*PutSystemCVEAllowlistInternalServerError Internal server error

swagger:response putSystemCveAllowlistInternalServerError
*/
type PutSystemCVEAllowlistInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewPutSystemCVEAllowlistInternalServerError creates PutSystemCVEAllowlistInternalServerError with default headers values
func NewPutSystemCVEAllowlistInternalServerError() *PutSystemCVEAllowlistInternalServerError {

	return &PutSystemCVEAllowlistInternalServerError{}
}

// WithXRequestID adds the xRequestId to the put system Cve allowlist internal server error response
func (o *PutSystemCVEAllowlistInternalServerError) WithXRequestID(xRequestID string) *PutSystemCVEAllowlistInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the put system Cve allowlist internal server error response
func (o *PutSystemCVEAllowlistInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the put system Cve allowlist internal server error response
func (o *PutSystemCVEAllowlistInternalServerError) WithPayload(payload *models.Errors) *PutSystemCVEAllowlistInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put system Cve allowlist internal server error response
func (o *PutSystemCVEAllowlistInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutSystemCVEAllowlistInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
