// Code generated by go-swagger; DO NOT EDIT.

package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// GetRegistryInfoOKCode is the HTTP code returned for type GetRegistryInfoOK
const GetRegistryInfoOKCode int = 200

/*GetRegistryInfoOK Success

swagger:response getRegistryInfoOK
*/
type GetRegistryInfoOK struct {

	/*
	  In: Body
	*/
	Payload *models.RegistryInfo `json:"body,omitempty"`
}

// NewGetRegistryInfoOK creates GetRegistryInfoOK with default headers values
func NewGetRegistryInfoOK() *GetRegistryInfoOK {

	return &GetRegistryInfoOK{}
}

// WithPayload adds the payload to the get registry info o k response
func (o *GetRegistryInfoOK) WithPayload(payload *models.RegistryInfo) *GetRegistryInfoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get registry info o k response
func (o *GetRegistryInfoOK) SetPayload(payload *models.RegistryInfo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRegistryInfoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRegistryInfoUnauthorizedCode is the HTTP code returned for type GetRegistryInfoUnauthorized
const GetRegistryInfoUnauthorizedCode int = 401

/*GetRegistryInfoUnauthorized Unauthorized

swagger:response getRegistryInfoUnauthorized
*/
type GetRegistryInfoUnauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetRegistryInfoUnauthorized creates GetRegistryInfoUnauthorized with default headers values
func NewGetRegistryInfoUnauthorized() *GetRegistryInfoUnauthorized {

	return &GetRegistryInfoUnauthorized{}
}

// WithXRequestID adds the xRequestId to the get registry info unauthorized response
func (o *GetRegistryInfoUnauthorized) WithXRequestID(xRequestID string) *GetRegistryInfoUnauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get registry info unauthorized response
func (o *GetRegistryInfoUnauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get registry info unauthorized response
func (o *GetRegistryInfoUnauthorized) WithPayload(payload *models.Errors) *GetRegistryInfoUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get registry info unauthorized response
func (o *GetRegistryInfoUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRegistryInfoUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetRegistryInfoForbiddenCode is the HTTP code returned for type GetRegistryInfoForbidden
const GetRegistryInfoForbiddenCode int = 403

/*GetRegistryInfoForbidden Forbidden

swagger:response getRegistryInfoForbidden
*/
type GetRegistryInfoForbidden struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetRegistryInfoForbidden creates GetRegistryInfoForbidden with default headers values
func NewGetRegistryInfoForbidden() *GetRegistryInfoForbidden {

	return &GetRegistryInfoForbidden{}
}

// WithXRequestID adds the xRequestId to the get registry info forbidden response
func (o *GetRegistryInfoForbidden) WithXRequestID(xRequestID string) *GetRegistryInfoForbidden {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get registry info forbidden response
func (o *GetRegistryInfoForbidden) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get registry info forbidden response
func (o *GetRegistryInfoForbidden) WithPayload(payload *models.Errors) *GetRegistryInfoForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get registry info forbidden response
func (o *GetRegistryInfoForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRegistryInfoForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetRegistryInfoNotFoundCode is the HTTP code returned for type GetRegistryInfoNotFound
const GetRegistryInfoNotFoundCode int = 404

/*GetRegistryInfoNotFound Not found

swagger:response getRegistryInfoNotFound
*/
type GetRegistryInfoNotFound struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetRegistryInfoNotFound creates GetRegistryInfoNotFound with default headers values
func NewGetRegistryInfoNotFound() *GetRegistryInfoNotFound {

	return &GetRegistryInfoNotFound{}
}

// WithXRequestID adds the xRequestId to the get registry info not found response
func (o *GetRegistryInfoNotFound) WithXRequestID(xRequestID string) *GetRegistryInfoNotFound {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get registry info not found response
func (o *GetRegistryInfoNotFound) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get registry info not found response
func (o *GetRegistryInfoNotFound) WithPayload(payload *models.Errors) *GetRegistryInfoNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get registry info not found response
func (o *GetRegistryInfoNotFound) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRegistryInfoNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetRegistryInfoInternalServerErrorCode is the HTTP code returned for type GetRegistryInfoInternalServerError
const GetRegistryInfoInternalServerErrorCode int = 500

/*GetRegistryInfoInternalServerError Internal server error

swagger:response getRegistryInfoInternalServerError
*/
type GetRegistryInfoInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetRegistryInfoInternalServerError creates GetRegistryInfoInternalServerError with default headers values
func NewGetRegistryInfoInternalServerError() *GetRegistryInfoInternalServerError {

	return &GetRegistryInfoInternalServerError{}
}

// WithXRequestID adds the xRequestId to the get registry info internal server error response
func (o *GetRegistryInfoInternalServerError) WithXRequestID(xRequestID string) *GetRegistryInfoInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get registry info internal server error response
func (o *GetRegistryInfoInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get registry info internal server error response
func (o *GetRegistryInfoInternalServerError) WithPayload(payload *models.Errors) *GetRegistryInfoInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get registry info internal server error response
func (o *GetRegistryInfoInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRegistryInfoInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
