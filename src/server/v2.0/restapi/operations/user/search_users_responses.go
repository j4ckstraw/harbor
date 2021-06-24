// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// SearchUsersOKCode is the HTTP code returned for type SearchUsersOK
const SearchUsersOKCode int = 200

/*SearchUsersOK Search users by username successfully.

swagger:response searchUsersOK
*/
type SearchUsersOK struct {
	/*Link to previous page and next page

	 */
	Link string `json:"Link"`
	/*The total count of available items

	 */
	XTotalCount int64 `json:"X-Total-Count"`

	/*
	  In: Body
	*/
	Payload []*models.UserSearchRespItem `json:"body,omitempty"`
}

// NewSearchUsersOK creates SearchUsersOK with default headers values
func NewSearchUsersOK() *SearchUsersOK {

	return &SearchUsersOK{}
}

// WithLink adds the link to the search users o k response
func (o *SearchUsersOK) WithLink(link string) *SearchUsersOK {
	o.Link = link
	return o
}

// SetLink sets the link to the search users o k response
func (o *SearchUsersOK) SetLink(link string) {
	o.Link = link
}

// WithXTotalCount adds the xTotalCount to the search users o k response
func (o *SearchUsersOK) WithXTotalCount(xTotalCount int64) *SearchUsersOK {
	o.XTotalCount = xTotalCount
	return o
}

// SetXTotalCount sets the xTotalCount to the search users o k response
func (o *SearchUsersOK) SetXTotalCount(xTotalCount int64) {
	o.XTotalCount = xTotalCount
}

// WithPayload adds the payload to the search users o k response
func (o *SearchUsersOK) WithPayload(payload []*models.UserSearchRespItem) *SearchUsersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the search users o k response
func (o *SearchUsersOK) SetPayload(payload []*models.UserSearchRespItem) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SearchUsersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Link

	link := o.Link
	if link != "" {
		rw.Header().Set("Link", link)
	}

	// response header X-Total-Count

	xTotalCount := swag.FormatInt64(o.XTotalCount)
	if xTotalCount != "" {
		rw.Header().Set("X-Total-Count", xTotalCount)
	}

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.UserSearchRespItem, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// SearchUsersUnauthorizedCode is the HTTP code returned for type SearchUsersUnauthorized
const SearchUsersUnauthorizedCode int = 401

/*SearchUsersUnauthorized Unauthorized

swagger:response searchUsersUnauthorized
*/
type SearchUsersUnauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewSearchUsersUnauthorized creates SearchUsersUnauthorized with default headers values
func NewSearchUsersUnauthorized() *SearchUsersUnauthorized {

	return &SearchUsersUnauthorized{}
}

// WithXRequestID adds the xRequestId to the search users unauthorized response
func (o *SearchUsersUnauthorized) WithXRequestID(xRequestID string) *SearchUsersUnauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the search users unauthorized response
func (o *SearchUsersUnauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the search users unauthorized response
func (o *SearchUsersUnauthorized) WithPayload(payload *models.Errors) *SearchUsersUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the search users unauthorized response
func (o *SearchUsersUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SearchUsersUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// SearchUsersInternalServerErrorCode is the HTTP code returned for type SearchUsersInternalServerError
const SearchUsersInternalServerErrorCode int = 500

/*SearchUsersInternalServerError Internal server error

swagger:response searchUsersInternalServerError
*/
type SearchUsersInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewSearchUsersInternalServerError creates SearchUsersInternalServerError with default headers values
func NewSearchUsersInternalServerError() *SearchUsersInternalServerError {

	return &SearchUsersInternalServerError{}
}

// WithXRequestID adds the xRequestId to the search users internal server error response
func (o *SearchUsersInternalServerError) WithXRequestID(xRequestID string) *SearchUsersInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the search users internal server error response
func (o *SearchUsersInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the search users internal server error response
func (o *SearchUsersInternalServerError) WithPayload(payload *models.Errors) *SearchUsersInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the search users internal server error response
func (o *SearchUsersInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SearchUsersInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
