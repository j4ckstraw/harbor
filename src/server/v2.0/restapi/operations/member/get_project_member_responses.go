// Code generated by go-swagger; DO NOT EDIT.

package member

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// GetProjectMemberOKCode is the HTTP code returned for type GetProjectMemberOK
const GetProjectMemberOKCode int = 200

/*GetProjectMemberOK Project member retrieved successfully.

swagger:response getProjectMemberOK
*/
type GetProjectMemberOK struct {

	/*
	  In: Body
	*/
	Payload *models.ProjectMemberEntity `json:"body,omitempty"`
}

// NewGetProjectMemberOK creates GetProjectMemberOK with default headers values
func NewGetProjectMemberOK() *GetProjectMemberOK {

	return &GetProjectMemberOK{}
}

// WithPayload adds the payload to the get project member o k response
func (o *GetProjectMemberOK) WithPayload(payload *models.ProjectMemberEntity) *GetProjectMemberOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get project member o k response
func (o *GetProjectMemberOK) SetPayload(payload *models.ProjectMemberEntity) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectMemberOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetProjectMemberBadRequestCode is the HTTP code returned for type GetProjectMemberBadRequest
const GetProjectMemberBadRequestCode int = 400

/*GetProjectMemberBadRequest Bad request

swagger:response getProjectMemberBadRequest
*/
type GetProjectMemberBadRequest struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetProjectMemberBadRequest creates GetProjectMemberBadRequest with default headers values
func NewGetProjectMemberBadRequest() *GetProjectMemberBadRequest {

	return &GetProjectMemberBadRequest{}
}

// WithXRequestID adds the xRequestId to the get project member bad request response
func (o *GetProjectMemberBadRequest) WithXRequestID(xRequestID string) *GetProjectMemberBadRequest {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get project member bad request response
func (o *GetProjectMemberBadRequest) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get project member bad request response
func (o *GetProjectMemberBadRequest) WithPayload(payload *models.Errors) *GetProjectMemberBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get project member bad request response
func (o *GetProjectMemberBadRequest) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectMemberBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetProjectMemberUnauthorizedCode is the HTTP code returned for type GetProjectMemberUnauthorized
const GetProjectMemberUnauthorizedCode int = 401

/*GetProjectMemberUnauthorized Unauthorized

swagger:response getProjectMemberUnauthorized
*/
type GetProjectMemberUnauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetProjectMemberUnauthorized creates GetProjectMemberUnauthorized with default headers values
func NewGetProjectMemberUnauthorized() *GetProjectMemberUnauthorized {

	return &GetProjectMemberUnauthorized{}
}

// WithXRequestID adds the xRequestId to the get project member unauthorized response
func (o *GetProjectMemberUnauthorized) WithXRequestID(xRequestID string) *GetProjectMemberUnauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get project member unauthorized response
func (o *GetProjectMemberUnauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get project member unauthorized response
func (o *GetProjectMemberUnauthorized) WithPayload(payload *models.Errors) *GetProjectMemberUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get project member unauthorized response
func (o *GetProjectMemberUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectMemberUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetProjectMemberForbiddenCode is the HTTP code returned for type GetProjectMemberForbidden
const GetProjectMemberForbiddenCode int = 403

/*GetProjectMemberForbidden Forbidden

swagger:response getProjectMemberForbidden
*/
type GetProjectMemberForbidden struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetProjectMemberForbidden creates GetProjectMemberForbidden with default headers values
func NewGetProjectMemberForbidden() *GetProjectMemberForbidden {

	return &GetProjectMemberForbidden{}
}

// WithXRequestID adds the xRequestId to the get project member forbidden response
func (o *GetProjectMemberForbidden) WithXRequestID(xRequestID string) *GetProjectMemberForbidden {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get project member forbidden response
func (o *GetProjectMemberForbidden) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get project member forbidden response
func (o *GetProjectMemberForbidden) WithPayload(payload *models.Errors) *GetProjectMemberForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get project member forbidden response
func (o *GetProjectMemberForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectMemberForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetProjectMemberNotFoundCode is the HTTP code returned for type GetProjectMemberNotFound
const GetProjectMemberNotFoundCode int = 404

/*GetProjectMemberNotFound Not found

swagger:response getProjectMemberNotFound
*/
type GetProjectMemberNotFound struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetProjectMemberNotFound creates GetProjectMemberNotFound with default headers values
func NewGetProjectMemberNotFound() *GetProjectMemberNotFound {

	return &GetProjectMemberNotFound{}
}

// WithXRequestID adds the xRequestId to the get project member not found response
func (o *GetProjectMemberNotFound) WithXRequestID(xRequestID string) *GetProjectMemberNotFound {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get project member not found response
func (o *GetProjectMemberNotFound) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get project member not found response
func (o *GetProjectMemberNotFound) WithPayload(payload *models.Errors) *GetProjectMemberNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get project member not found response
func (o *GetProjectMemberNotFound) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectMemberNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetProjectMemberInternalServerErrorCode is the HTTP code returned for type GetProjectMemberInternalServerError
const GetProjectMemberInternalServerErrorCode int = 500

/*GetProjectMemberInternalServerError Internal server error

swagger:response getProjectMemberInternalServerError
*/
type GetProjectMemberInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetProjectMemberInternalServerError creates GetProjectMemberInternalServerError with default headers values
func NewGetProjectMemberInternalServerError() *GetProjectMemberInternalServerError {

	return &GetProjectMemberInternalServerError{}
}

// WithXRequestID adds the xRequestId to the get project member internal server error response
func (o *GetProjectMemberInternalServerError) WithXRequestID(xRequestID string) *GetProjectMemberInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get project member internal server error response
func (o *GetProjectMemberInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get project member internal server error response
func (o *GetProjectMemberInternalServerError) WithPayload(payload *models.Errors) *GetProjectMemberInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get project member internal server error response
func (o *GetProjectMemberInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectMemberInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
