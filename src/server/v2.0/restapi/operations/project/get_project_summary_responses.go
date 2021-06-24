// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// GetProjectSummaryOKCode is the HTTP code returned for type GetProjectSummaryOK
const GetProjectSummaryOKCode int = 200

/*GetProjectSummaryOK Get summary of the project successfully.

swagger:response getProjectSummaryOK
*/
type GetProjectSummaryOK struct {

	/*
	  In: Body
	*/
	Payload *models.ProjectSummary `json:"body,omitempty"`
}

// NewGetProjectSummaryOK creates GetProjectSummaryOK with default headers values
func NewGetProjectSummaryOK() *GetProjectSummaryOK {

	return &GetProjectSummaryOK{}
}

// WithPayload adds the payload to the get project summary o k response
func (o *GetProjectSummaryOK) WithPayload(payload *models.ProjectSummary) *GetProjectSummaryOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get project summary o k response
func (o *GetProjectSummaryOK) SetPayload(payload *models.ProjectSummary) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectSummaryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetProjectSummaryBadRequestCode is the HTTP code returned for type GetProjectSummaryBadRequest
const GetProjectSummaryBadRequestCode int = 400

/*GetProjectSummaryBadRequest Bad request

swagger:response getProjectSummaryBadRequest
*/
type GetProjectSummaryBadRequest struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetProjectSummaryBadRequest creates GetProjectSummaryBadRequest with default headers values
func NewGetProjectSummaryBadRequest() *GetProjectSummaryBadRequest {

	return &GetProjectSummaryBadRequest{}
}

// WithXRequestID adds the xRequestId to the get project summary bad request response
func (o *GetProjectSummaryBadRequest) WithXRequestID(xRequestID string) *GetProjectSummaryBadRequest {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get project summary bad request response
func (o *GetProjectSummaryBadRequest) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get project summary bad request response
func (o *GetProjectSummaryBadRequest) WithPayload(payload *models.Errors) *GetProjectSummaryBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get project summary bad request response
func (o *GetProjectSummaryBadRequest) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectSummaryBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetProjectSummaryUnauthorizedCode is the HTTP code returned for type GetProjectSummaryUnauthorized
const GetProjectSummaryUnauthorizedCode int = 401

/*GetProjectSummaryUnauthorized Unauthorized

swagger:response getProjectSummaryUnauthorized
*/
type GetProjectSummaryUnauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetProjectSummaryUnauthorized creates GetProjectSummaryUnauthorized with default headers values
func NewGetProjectSummaryUnauthorized() *GetProjectSummaryUnauthorized {

	return &GetProjectSummaryUnauthorized{}
}

// WithXRequestID adds the xRequestId to the get project summary unauthorized response
func (o *GetProjectSummaryUnauthorized) WithXRequestID(xRequestID string) *GetProjectSummaryUnauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get project summary unauthorized response
func (o *GetProjectSummaryUnauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get project summary unauthorized response
func (o *GetProjectSummaryUnauthorized) WithPayload(payload *models.Errors) *GetProjectSummaryUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get project summary unauthorized response
func (o *GetProjectSummaryUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectSummaryUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetProjectSummaryForbiddenCode is the HTTP code returned for type GetProjectSummaryForbidden
const GetProjectSummaryForbiddenCode int = 403

/*GetProjectSummaryForbidden Forbidden

swagger:response getProjectSummaryForbidden
*/
type GetProjectSummaryForbidden struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetProjectSummaryForbidden creates GetProjectSummaryForbidden with default headers values
func NewGetProjectSummaryForbidden() *GetProjectSummaryForbidden {

	return &GetProjectSummaryForbidden{}
}

// WithXRequestID adds the xRequestId to the get project summary forbidden response
func (o *GetProjectSummaryForbidden) WithXRequestID(xRequestID string) *GetProjectSummaryForbidden {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get project summary forbidden response
func (o *GetProjectSummaryForbidden) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get project summary forbidden response
func (o *GetProjectSummaryForbidden) WithPayload(payload *models.Errors) *GetProjectSummaryForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get project summary forbidden response
func (o *GetProjectSummaryForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectSummaryForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetProjectSummaryNotFoundCode is the HTTP code returned for type GetProjectSummaryNotFound
const GetProjectSummaryNotFoundCode int = 404

/*GetProjectSummaryNotFound Not found

swagger:response getProjectSummaryNotFound
*/
type GetProjectSummaryNotFound struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetProjectSummaryNotFound creates GetProjectSummaryNotFound with default headers values
func NewGetProjectSummaryNotFound() *GetProjectSummaryNotFound {

	return &GetProjectSummaryNotFound{}
}

// WithXRequestID adds the xRequestId to the get project summary not found response
func (o *GetProjectSummaryNotFound) WithXRequestID(xRequestID string) *GetProjectSummaryNotFound {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get project summary not found response
func (o *GetProjectSummaryNotFound) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get project summary not found response
func (o *GetProjectSummaryNotFound) WithPayload(payload *models.Errors) *GetProjectSummaryNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get project summary not found response
func (o *GetProjectSummaryNotFound) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectSummaryNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetProjectSummaryInternalServerErrorCode is the HTTP code returned for type GetProjectSummaryInternalServerError
const GetProjectSummaryInternalServerErrorCode int = 500

/*GetProjectSummaryInternalServerError Internal server error

swagger:response getProjectSummaryInternalServerError
*/
type GetProjectSummaryInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetProjectSummaryInternalServerError creates GetProjectSummaryInternalServerError with default headers values
func NewGetProjectSummaryInternalServerError() *GetProjectSummaryInternalServerError {

	return &GetProjectSummaryInternalServerError{}
}

// WithXRequestID adds the xRequestId to the get project summary internal server error response
func (o *GetProjectSummaryInternalServerError) WithXRequestID(xRequestID string) *GetProjectSummaryInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get project summary internal server error response
func (o *GetProjectSummaryInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get project summary internal server error response
func (o *GetProjectSummaryInternalServerError) WithPayload(payload *models.Errors) *GetProjectSummaryInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get project summary internal server error response
func (o *GetProjectSummaryInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectSummaryInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
