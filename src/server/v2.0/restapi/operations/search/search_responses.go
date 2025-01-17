// Code generated by go-swagger; DO NOT EDIT.

package search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// SearchOKCode is the HTTP code returned for type SearchOK
const SearchOKCode int = 200

/*SearchOK An array of search results

swagger:response searchOK
*/
type SearchOK struct {

	/*
	  In: Body
	*/
	Payload *models.Search `json:"body,omitempty"`
}

// NewSearchOK creates SearchOK with default headers values
func NewSearchOK() *SearchOK {

	return &SearchOK{}
}

// WithPayload adds the payload to the search o k response
func (o *SearchOK) WithPayload(payload *models.Search) *SearchOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the search o k response
func (o *SearchOK) SetPayload(payload *models.Search) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SearchOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SearchInternalServerErrorCode is the HTTP code returned for type SearchInternalServerError
const SearchInternalServerErrorCode int = 500

/*SearchInternalServerError Internal server error

swagger:response searchInternalServerError
*/
type SearchInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewSearchInternalServerError creates SearchInternalServerError with default headers values
func NewSearchInternalServerError() *SearchInternalServerError {

	return &SearchInternalServerError{}
}

// WithXRequestID adds the xRequestId to the search internal server error response
func (o *SearchInternalServerError) WithXRequestID(xRequestID string) *SearchInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the search internal server error response
func (o *SearchInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the search internal server error response
func (o *SearchInternalServerError) WithPayload(payload *models.Errors) *SearchInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the search internal server error response
func (o *SearchInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SearchInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
