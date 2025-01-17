// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// ListAllRepositoriesOKCode is the HTTP code returned for type ListAllRepositoriesOK
const ListAllRepositoriesOKCode int = 200

/*ListAllRepositoriesOK Success

swagger:response listAllRepositoriesOK
*/
type ListAllRepositoriesOK struct {
	/*Link refers to the previous page and next page

	 */
	Link string `json:"Link"`
	/*The total count of repositories

	 */
	XTotalCount int64 `json:"X-Total-Count"`

	/*
	  In: Body
	*/
	Payload []*models.Repository `json:"body,omitempty"`
}

// NewListAllRepositoriesOK creates ListAllRepositoriesOK with default headers values
func NewListAllRepositoriesOK() *ListAllRepositoriesOK {

	return &ListAllRepositoriesOK{}
}

// WithLink adds the link to the list all repositories o k response
func (o *ListAllRepositoriesOK) WithLink(link string) *ListAllRepositoriesOK {
	o.Link = link
	return o
}

// SetLink sets the link to the list all repositories o k response
func (o *ListAllRepositoriesOK) SetLink(link string) {
	o.Link = link
}

// WithXTotalCount adds the xTotalCount to the list all repositories o k response
func (o *ListAllRepositoriesOK) WithXTotalCount(xTotalCount int64) *ListAllRepositoriesOK {
	o.XTotalCount = xTotalCount
	return o
}

// SetXTotalCount sets the xTotalCount to the list all repositories o k response
func (o *ListAllRepositoriesOK) SetXTotalCount(xTotalCount int64) {
	o.XTotalCount = xTotalCount
}

// WithPayload adds the payload to the list all repositories o k response
func (o *ListAllRepositoriesOK) WithPayload(payload []*models.Repository) *ListAllRepositoriesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list all repositories o k response
func (o *ListAllRepositoriesOK) SetPayload(payload []*models.Repository) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAllRepositoriesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
		payload = make([]*models.Repository, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListAllRepositoriesBadRequestCode is the HTTP code returned for type ListAllRepositoriesBadRequest
const ListAllRepositoriesBadRequestCode int = 400

/*ListAllRepositoriesBadRequest Bad request

swagger:response listAllRepositoriesBadRequest
*/
type ListAllRepositoriesBadRequest struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewListAllRepositoriesBadRequest creates ListAllRepositoriesBadRequest with default headers values
func NewListAllRepositoriesBadRequest() *ListAllRepositoriesBadRequest {

	return &ListAllRepositoriesBadRequest{}
}

// WithXRequestID adds the xRequestId to the list all repositories bad request response
func (o *ListAllRepositoriesBadRequest) WithXRequestID(xRequestID string) *ListAllRepositoriesBadRequest {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the list all repositories bad request response
func (o *ListAllRepositoriesBadRequest) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the list all repositories bad request response
func (o *ListAllRepositoriesBadRequest) WithPayload(payload *models.Errors) *ListAllRepositoriesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list all repositories bad request response
func (o *ListAllRepositoriesBadRequest) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAllRepositoriesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ListAllRepositoriesInternalServerErrorCode is the HTTP code returned for type ListAllRepositoriesInternalServerError
const ListAllRepositoriesInternalServerErrorCode int = 500

/*ListAllRepositoriesInternalServerError Internal server error

swagger:response listAllRepositoriesInternalServerError
*/
type ListAllRepositoriesInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewListAllRepositoriesInternalServerError creates ListAllRepositoriesInternalServerError with default headers values
func NewListAllRepositoriesInternalServerError() *ListAllRepositoriesInternalServerError {

	return &ListAllRepositoriesInternalServerError{}
}

// WithXRequestID adds the xRequestId to the list all repositories internal server error response
func (o *ListAllRepositoriesInternalServerError) WithXRequestID(xRequestID string) *ListAllRepositoriesInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the list all repositories internal server error response
func (o *ListAllRepositoriesInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the list all repositories internal server error response
func (o *ListAllRepositoriesInternalServerError) WithPayload(payload *models.Errors) *ListAllRepositoriesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list all repositories internal server error response
func (o *ListAllRepositoriesInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAllRepositoriesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
