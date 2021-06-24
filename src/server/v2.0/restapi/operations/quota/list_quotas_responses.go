// Code generated by go-swagger; DO NOT EDIT.

package quota

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// ListQuotasOKCode is the HTTP code returned for type ListQuotasOK
const ListQuotasOKCode int = 200

/*ListQuotasOK Successfully retrieved the quotas.

swagger:response listQuotasOK
*/
type ListQuotasOK struct {
	/*Link refers to the previous page and next page

	 */
	Link string `json:"Link"`
	/*The total count of access logs

	 */
	XTotalCount int64 `json:"X-Total-Count"`

	/*
	  In: Body
	*/
	Payload []*models.Quota `json:"body,omitempty"`
}

// NewListQuotasOK creates ListQuotasOK with default headers values
func NewListQuotasOK() *ListQuotasOK {

	return &ListQuotasOK{}
}

// WithLink adds the link to the list quotas o k response
func (o *ListQuotasOK) WithLink(link string) *ListQuotasOK {
	o.Link = link
	return o
}

// SetLink sets the link to the list quotas o k response
func (o *ListQuotasOK) SetLink(link string) {
	o.Link = link
}

// WithXTotalCount adds the xTotalCount to the list quotas o k response
func (o *ListQuotasOK) WithXTotalCount(xTotalCount int64) *ListQuotasOK {
	o.XTotalCount = xTotalCount
	return o
}

// SetXTotalCount sets the xTotalCount to the list quotas o k response
func (o *ListQuotasOK) SetXTotalCount(xTotalCount int64) {
	o.XTotalCount = xTotalCount
}

// WithPayload adds the payload to the list quotas o k response
func (o *ListQuotasOK) WithPayload(payload []*models.Quota) *ListQuotasOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list quotas o k response
func (o *ListQuotasOK) SetPayload(payload []*models.Quota) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListQuotasOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
		payload = make([]*models.Quota, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListQuotasUnauthorizedCode is the HTTP code returned for type ListQuotasUnauthorized
const ListQuotasUnauthorizedCode int = 401

/*ListQuotasUnauthorized Unauthorized

swagger:response listQuotasUnauthorized
*/
type ListQuotasUnauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewListQuotasUnauthorized creates ListQuotasUnauthorized with default headers values
func NewListQuotasUnauthorized() *ListQuotasUnauthorized {

	return &ListQuotasUnauthorized{}
}

// WithXRequestID adds the xRequestId to the list quotas unauthorized response
func (o *ListQuotasUnauthorized) WithXRequestID(xRequestID string) *ListQuotasUnauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the list quotas unauthorized response
func (o *ListQuotasUnauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the list quotas unauthorized response
func (o *ListQuotasUnauthorized) WithPayload(payload *models.Errors) *ListQuotasUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list quotas unauthorized response
func (o *ListQuotasUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListQuotasUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ListQuotasForbiddenCode is the HTTP code returned for type ListQuotasForbidden
const ListQuotasForbiddenCode int = 403

/*ListQuotasForbidden Forbidden

swagger:response listQuotasForbidden
*/
type ListQuotasForbidden struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewListQuotasForbidden creates ListQuotasForbidden with default headers values
func NewListQuotasForbidden() *ListQuotasForbidden {

	return &ListQuotasForbidden{}
}

// WithXRequestID adds the xRequestId to the list quotas forbidden response
func (o *ListQuotasForbidden) WithXRequestID(xRequestID string) *ListQuotasForbidden {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the list quotas forbidden response
func (o *ListQuotasForbidden) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the list quotas forbidden response
func (o *ListQuotasForbidden) WithPayload(payload *models.Errors) *ListQuotasForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list quotas forbidden response
func (o *ListQuotasForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListQuotasForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ListQuotasInternalServerErrorCode is the HTTP code returned for type ListQuotasInternalServerError
const ListQuotasInternalServerErrorCode int = 500

/*ListQuotasInternalServerError Internal server error

swagger:response listQuotasInternalServerError
*/
type ListQuotasInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewListQuotasInternalServerError creates ListQuotasInternalServerError with default headers values
func NewListQuotasInternalServerError() *ListQuotasInternalServerError {

	return &ListQuotasInternalServerError{}
}

// WithXRequestID adds the xRequestId to the list quotas internal server error response
func (o *ListQuotasInternalServerError) WithXRequestID(xRequestID string) *ListQuotasInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the list quotas internal server error response
func (o *ListQuotasInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the list quotas internal server error response
func (o *ListQuotasInternalServerError) WithPayload(payload *models.Errors) *ListQuotasInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list quotas internal server error response
func (o *ListQuotasInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListQuotasInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
