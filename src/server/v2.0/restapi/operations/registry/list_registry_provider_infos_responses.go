// Code generated by go-swagger; DO NOT EDIT.

package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// ListRegistryProviderInfosOKCode is the HTTP code returned for type ListRegistryProviderInfosOK
const ListRegistryProviderInfosOKCode int = 200

/*ListRegistryProviderInfosOK Success.

swagger:response listRegistryProviderInfosOK
*/
type ListRegistryProviderInfosOK struct {

	/*
	  In: Body
	*/
	Payload map[string]models.RegistryProviderInfo `json:"body,omitempty"`
}

// NewListRegistryProviderInfosOK creates ListRegistryProviderInfosOK with default headers values
func NewListRegistryProviderInfosOK() *ListRegistryProviderInfosOK {

	return &ListRegistryProviderInfosOK{}
}

// WithPayload adds the payload to the list registry provider infos o k response
func (o *ListRegistryProviderInfosOK) WithPayload(payload map[string]models.RegistryProviderInfo) *ListRegistryProviderInfosOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list registry provider infos o k response
func (o *ListRegistryProviderInfosOK) SetPayload(payload map[string]models.RegistryProviderInfo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListRegistryProviderInfosOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty map
		payload = make(map[string]models.RegistryProviderInfo, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListRegistryProviderInfosUnauthorizedCode is the HTTP code returned for type ListRegistryProviderInfosUnauthorized
const ListRegistryProviderInfosUnauthorizedCode int = 401

/*ListRegistryProviderInfosUnauthorized Unauthorized

swagger:response listRegistryProviderInfosUnauthorized
*/
type ListRegistryProviderInfosUnauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewListRegistryProviderInfosUnauthorized creates ListRegistryProviderInfosUnauthorized with default headers values
func NewListRegistryProviderInfosUnauthorized() *ListRegistryProviderInfosUnauthorized {

	return &ListRegistryProviderInfosUnauthorized{}
}

// WithXRequestID adds the xRequestId to the list registry provider infos unauthorized response
func (o *ListRegistryProviderInfosUnauthorized) WithXRequestID(xRequestID string) *ListRegistryProviderInfosUnauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the list registry provider infos unauthorized response
func (o *ListRegistryProviderInfosUnauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the list registry provider infos unauthorized response
func (o *ListRegistryProviderInfosUnauthorized) WithPayload(payload *models.Errors) *ListRegistryProviderInfosUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list registry provider infos unauthorized response
func (o *ListRegistryProviderInfosUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListRegistryProviderInfosUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ListRegistryProviderInfosForbiddenCode is the HTTP code returned for type ListRegistryProviderInfosForbidden
const ListRegistryProviderInfosForbiddenCode int = 403

/*ListRegistryProviderInfosForbidden Forbidden

swagger:response listRegistryProviderInfosForbidden
*/
type ListRegistryProviderInfosForbidden struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewListRegistryProviderInfosForbidden creates ListRegistryProviderInfosForbidden with default headers values
func NewListRegistryProviderInfosForbidden() *ListRegistryProviderInfosForbidden {

	return &ListRegistryProviderInfosForbidden{}
}

// WithXRequestID adds the xRequestId to the list registry provider infos forbidden response
func (o *ListRegistryProviderInfosForbidden) WithXRequestID(xRequestID string) *ListRegistryProviderInfosForbidden {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the list registry provider infos forbidden response
func (o *ListRegistryProviderInfosForbidden) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the list registry provider infos forbidden response
func (o *ListRegistryProviderInfosForbidden) WithPayload(payload *models.Errors) *ListRegistryProviderInfosForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list registry provider infos forbidden response
func (o *ListRegistryProviderInfosForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListRegistryProviderInfosForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ListRegistryProviderInfosInternalServerErrorCode is the HTTP code returned for type ListRegistryProviderInfosInternalServerError
const ListRegistryProviderInfosInternalServerErrorCode int = 500

/*ListRegistryProviderInfosInternalServerError Internal server error

swagger:response listRegistryProviderInfosInternalServerError
*/
type ListRegistryProviderInfosInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewListRegistryProviderInfosInternalServerError creates ListRegistryProviderInfosInternalServerError with default headers values
func NewListRegistryProviderInfosInternalServerError() *ListRegistryProviderInfosInternalServerError {

	return &ListRegistryProviderInfosInternalServerError{}
}

// WithXRequestID adds the xRequestId to the list registry provider infos internal server error response
func (o *ListRegistryProviderInfosInternalServerError) WithXRequestID(xRequestID string) *ListRegistryProviderInfosInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the list registry provider infos internal server error response
func (o *ListRegistryProviderInfosInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the list registry provider infos internal server error response
func (o *ListRegistryProviderInfosInternalServerError) WithPayload(payload *models.Errors) *ListRegistryProviderInfosInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list registry provider infos internal server error response
func (o *ListRegistryProviderInfosInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListRegistryProviderInfosInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
