// Code generated by go-swagger; DO NOT EDIT.

package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// ListRegistryProviderTypesOKCode is the HTTP code returned for type ListRegistryProviderTypesOK
const ListRegistryProviderTypesOKCode int = 200

/*ListRegistryProviderTypesOK Success.

swagger:response listRegistryProviderTypesOK
*/
type ListRegistryProviderTypesOK struct {

	/*
	  In: Body
	*/
	Payload []string `json:"body,omitempty"`
}

// NewListRegistryProviderTypesOK creates ListRegistryProviderTypesOK with default headers values
func NewListRegistryProviderTypesOK() *ListRegistryProviderTypesOK {

	return &ListRegistryProviderTypesOK{}
}

// WithPayload adds the payload to the list registry provider types o k response
func (o *ListRegistryProviderTypesOK) WithPayload(payload []string) *ListRegistryProviderTypesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list registry provider types o k response
func (o *ListRegistryProviderTypesOK) SetPayload(payload []string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListRegistryProviderTypesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]string, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListRegistryProviderTypesUnauthorizedCode is the HTTP code returned for type ListRegistryProviderTypesUnauthorized
const ListRegistryProviderTypesUnauthorizedCode int = 401

/*ListRegistryProviderTypesUnauthorized Unauthorized

swagger:response listRegistryProviderTypesUnauthorized
*/
type ListRegistryProviderTypesUnauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewListRegistryProviderTypesUnauthorized creates ListRegistryProviderTypesUnauthorized with default headers values
func NewListRegistryProviderTypesUnauthorized() *ListRegistryProviderTypesUnauthorized {

	return &ListRegistryProviderTypesUnauthorized{}
}

// WithXRequestID adds the xRequestId to the list registry provider types unauthorized response
func (o *ListRegistryProviderTypesUnauthorized) WithXRequestID(xRequestID string) *ListRegistryProviderTypesUnauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the list registry provider types unauthorized response
func (o *ListRegistryProviderTypesUnauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the list registry provider types unauthorized response
func (o *ListRegistryProviderTypesUnauthorized) WithPayload(payload *models.Errors) *ListRegistryProviderTypesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list registry provider types unauthorized response
func (o *ListRegistryProviderTypesUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListRegistryProviderTypesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ListRegistryProviderTypesForbiddenCode is the HTTP code returned for type ListRegistryProviderTypesForbidden
const ListRegistryProviderTypesForbiddenCode int = 403

/*ListRegistryProviderTypesForbidden Forbidden

swagger:response listRegistryProviderTypesForbidden
*/
type ListRegistryProviderTypesForbidden struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewListRegistryProviderTypesForbidden creates ListRegistryProviderTypesForbidden with default headers values
func NewListRegistryProviderTypesForbidden() *ListRegistryProviderTypesForbidden {

	return &ListRegistryProviderTypesForbidden{}
}

// WithXRequestID adds the xRequestId to the list registry provider types forbidden response
func (o *ListRegistryProviderTypesForbidden) WithXRequestID(xRequestID string) *ListRegistryProviderTypesForbidden {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the list registry provider types forbidden response
func (o *ListRegistryProviderTypesForbidden) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the list registry provider types forbidden response
func (o *ListRegistryProviderTypesForbidden) WithPayload(payload *models.Errors) *ListRegistryProviderTypesForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list registry provider types forbidden response
func (o *ListRegistryProviderTypesForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListRegistryProviderTypesForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ListRegistryProviderTypesInternalServerErrorCode is the HTTP code returned for type ListRegistryProviderTypesInternalServerError
const ListRegistryProviderTypesInternalServerErrorCode int = 500

/*ListRegistryProviderTypesInternalServerError Internal server error

swagger:response listRegistryProviderTypesInternalServerError
*/
type ListRegistryProviderTypesInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewListRegistryProviderTypesInternalServerError creates ListRegistryProviderTypesInternalServerError with default headers values
func NewListRegistryProviderTypesInternalServerError() *ListRegistryProviderTypesInternalServerError {

	return &ListRegistryProviderTypesInternalServerError{}
}

// WithXRequestID adds the xRequestId to the list registry provider types internal server error response
func (o *ListRegistryProviderTypesInternalServerError) WithXRequestID(xRequestID string) *ListRegistryProviderTypesInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the list registry provider types internal server error response
func (o *ListRegistryProviderTypesInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the list registry provider types internal server error response
func (o *ListRegistryProviderTypesInternalServerError) WithPayload(payload *models.Errors) *ListRegistryProviderTypesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list registry provider types internal server error response
func (o *ListRegistryProviderTypesInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListRegistryProviderTypesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
