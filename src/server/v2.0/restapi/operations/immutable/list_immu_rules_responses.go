// Code generated by go-swagger; DO NOT EDIT.

package immutable

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// ListImmuRulesOKCode is the HTTP code returned for type ListImmuRulesOK
const ListImmuRulesOKCode int = 200

/*ListImmuRulesOK Success

swagger:response listImmuRulesOK
*/
type ListImmuRulesOK struct {
	/*Link refers to the previous page and next page

	 */
	Link string `json:"Link"`
	/*The total count of immutable tag

	 */
	XTotalCount int64 `json:"X-Total-Count"`

	/*
	  In: Body
	*/
	Payload []*models.ImmutableRule `json:"body,omitempty"`
}

// NewListImmuRulesOK creates ListImmuRulesOK with default headers values
func NewListImmuRulesOK() *ListImmuRulesOK {

	return &ListImmuRulesOK{}
}

// WithLink adds the link to the list immu rules o k response
func (o *ListImmuRulesOK) WithLink(link string) *ListImmuRulesOK {
	o.Link = link
	return o
}

// SetLink sets the link to the list immu rules o k response
func (o *ListImmuRulesOK) SetLink(link string) {
	o.Link = link
}

// WithXTotalCount adds the xTotalCount to the list immu rules o k response
func (o *ListImmuRulesOK) WithXTotalCount(xTotalCount int64) *ListImmuRulesOK {
	o.XTotalCount = xTotalCount
	return o
}

// SetXTotalCount sets the xTotalCount to the list immu rules o k response
func (o *ListImmuRulesOK) SetXTotalCount(xTotalCount int64) {
	o.XTotalCount = xTotalCount
}

// WithPayload adds the payload to the list immu rules o k response
func (o *ListImmuRulesOK) WithPayload(payload []*models.ImmutableRule) *ListImmuRulesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list immu rules o k response
func (o *ListImmuRulesOK) SetPayload(payload []*models.ImmutableRule) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListImmuRulesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
		payload = make([]*models.ImmutableRule, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListImmuRulesBadRequestCode is the HTTP code returned for type ListImmuRulesBadRequest
const ListImmuRulesBadRequestCode int = 400

/*ListImmuRulesBadRequest Bad request

swagger:response listImmuRulesBadRequest
*/
type ListImmuRulesBadRequest struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewListImmuRulesBadRequest creates ListImmuRulesBadRequest with default headers values
func NewListImmuRulesBadRequest() *ListImmuRulesBadRequest {

	return &ListImmuRulesBadRequest{}
}

// WithXRequestID adds the xRequestId to the list immu rules bad request response
func (o *ListImmuRulesBadRequest) WithXRequestID(xRequestID string) *ListImmuRulesBadRequest {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the list immu rules bad request response
func (o *ListImmuRulesBadRequest) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the list immu rules bad request response
func (o *ListImmuRulesBadRequest) WithPayload(payload *models.Errors) *ListImmuRulesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list immu rules bad request response
func (o *ListImmuRulesBadRequest) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListImmuRulesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ListImmuRulesUnauthorizedCode is the HTTP code returned for type ListImmuRulesUnauthorized
const ListImmuRulesUnauthorizedCode int = 401

/*ListImmuRulesUnauthorized Unauthorized

swagger:response listImmuRulesUnauthorized
*/
type ListImmuRulesUnauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewListImmuRulesUnauthorized creates ListImmuRulesUnauthorized with default headers values
func NewListImmuRulesUnauthorized() *ListImmuRulesUnauthorized {

	return &ListImmuRulesUnauthorized{}
}

// WithXRequestID adds the xRequestId to the list immu rules unauthorized response
func (o *ListImmuRulesUnauthorized) WithXRequestID(xRequestID string) *ListImmuRulesUnauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the list immu rules unauthorized response
func (o *ListImmuRulesUnauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the list immu rules unauthorized response
func (o *ListImmuRulesUnauthorized) WithPayload(payload *models.Errors) *ListImmuRulesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list immu rules unauthorized response
func (o *ListImmuRulesUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListImmuRulesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ListImmuRulesForbiddenCode is the HTTP code returned for type ListImmuRulesForbidden
const ListImmuRulesForbiddenCode int = 403

/*ListImmuRulesForbidden Forbidden

swagger:response listImmuRulesForbidden
*/
type ListImmuRulesForbidden struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewListImmuRulesForbidden creates ListImmuRulesForbidden with default headers values
func NewListImmuRulesForbidden() *ListImmuRulesForbidden {

	return &ListImmuRulesForbidden{}
}

// WithXRequestID adds the xRequestId to the list immu rules forbidden response
func (o *ListImmuRulesForbidden) WithXRequestID(xRequestID string) *ListImmuRulesForbidden {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the list immu rules forbidden response
func (o *ListImmuRulesForbidden) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the list immu rules forbidden response
func (o *ListImmuRulesForbidden) WithPayload(payload *models.Errors) *ListImmuRulesForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list immu rules forbidden response
func (o *ListImmuRulesForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListImmuRulesForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ListImmuRulesInternalServerErrorCode is the HTTP code returned for type ListImmuRulesInternalServerError
const ListImmuRulesInternalServerErrorCode int = 500

/*ListImmuRulesInternalServerError Internal server error

swagger:response listImmuRulesInternalServerError
*/
type ListImmuRulesInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewListImmuRulesInternalServerError creates ListImmuRulesInternalServerError with default headers values
func NewListImmuRulesInternalServerError() *ListImmuRulesInternalServerError {

	return &ListImmuRulesInternalServerError{}
}

// WithXRequestID adds the xRequestId to the list immu rules internal server error response
func (o *ListImmuRulesInternalServerError) WithXRequestID(xRequestID string) *ListImmuRulesInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the list immu rules internal server error response
func (o *ListImmuRulesInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the list immu rules internal server error response
func (o *ListImmuRulesInternalServerError) WithPayload(payload *models.Errors) *ListImmuRulesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list immu rules internal server error response
func (o *ListImmuRulesInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListImmuRulesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
