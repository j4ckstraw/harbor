// Code generated by go-swagger; DO NOT EDIT.

package scanner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// GetScannerMetadataOKCode is the HTTP code returned for type GetScannerMetadataOK
const GetScannerMetadataOKCode int = 200

/*GetScannerMetadataOK The metadata of the specified scanner adapter

swagger:response getScannerMetadataOK
*/
type GetScannerMetadataOK struct {

	/*
	  In: Body
	*/
	Payload *models.ScannerAdapterMetadata `json:"body,omitempty"`
}

// NewGetScannerMetadataOK creates GetScannerMetadataOK with default headers values
func NewGetScannerMetadataOK() *GetScannerMetadataOK {

	return &GetScannerMetadataOK{}
}

// WithPayload adds the payload to the get scanner metadata o k response
func (o *GetScannerMetadataOK) WithPayload(payload *models.ScannerAdapterMetadata) *GetScannerMetadataOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get scanner metadata o k response
func (o *GetScannerMetadataOK) SetPayload(payload *models.ScannerAdapterMetadata) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetScannerMetadataOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetScannerMetadataUnauthorizedCode is the HTTP code returned for type GetScannerMetadataUnauthorized
const GetScannerMetadataUnauthorizedCode int = 401

/*GetScannerMetadataUnauthorized Unauthorized

swagger:response getScannerMetadataUnauthorized
*/
type GetScannerMetadataUnauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetScannerMetadataUnauthorized creates GetScannerMetadataUnauthorized with default headers values
func NewGetScannerMetadataUnauthorized() *GetScannerMetadataUnauthorized {

	return &GetScannerMetadataUnauthorized{}
}

// WithXRequestID adds the xRequestId to the get scanner metadata unauthorized response
func (o *GetScannerMetadataUnauthorized) WithXRequestID(xRequestID string) *GetScannerMetadataUnauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get scanner metadata unauthorized response
func (o *GetScannerMetadataUnauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get scanner metadata unauthorized response
func (o *GetScannerMetadataUnauthorized) WithPayload(payload *models.Errors) *GetScannerMetadataUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get scanner metadata unauthorized response
func (o *GetScannerMetadataUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetScannerMetadataUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetScannerMetadataForbiddenCode is the HTTP code returned for type GetScannerMetadataForbidden
const GetScannerMetadataForbiddenCode int = 403

/*GetScannerMetadataForbidden Forbidden

swagger:response getScannerMetadataForbidden
*/
type GetScannerMetadataForbidden struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetScannerMetadataForbidden creates GetScannerMetadataForbidden with default headers values
func NewGetScannerMetadataForbidden() *GetScannerMetadataForbidden {

	return &GetScannerMetadataForbidden{}
}

// WithXRequestID adds the xRequestId to the get scanner metadata forbidden response
func (o *GetScannerMetadataForbidden) WithXRequestID(xRequestID string) *GetScannerMetadataForbidden {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get scanner metadata forbidden response
func (o *GetScannerMetadataForbidden) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get scanner metadata forbidden response
func (o *GetScannerMetadataForbidden) WithPayload(payload *models.Errors) *GetScannerMetadataForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get scanner metadata forbidden response
func (o *GetScannerMetadataForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetScannerMetadataForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetScannerMetadataInternalServerErrorCode is the HTTP code returned for type GetScannerMetadataInternalServerError
const GetScannerMetadataInternalServerErrorCode int = 500

/*GetScannerMetadataInternalServerError Internal server error

swagger:response getScannerMetadataInternalServerError
*/
type GetScannerMetadataInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetScannerMetadataInternalServerError creates GetScannerMetadataInternalServerError with default headers values
func NewGetScannerMetadataInternalServerError() *GetScannerMetadataInternalServerError {

	return &GetScannerMetadataInternalServerError{}
}

// WithXRequestID adds the xRequestId to the get scanner metadata internal server error response
func (o *GetScannerMetadataInternalServerError) WithXRequestID(xRequestID string) *GetScannerMetadataInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get scanner metadata internal server error response
func (o *GetScannerMetadataInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get scanner metadata internal server error response
func (o *GetScannerMetadataInternalServerError) WithPayload(payload *models.Errors) *GetScannerMetadataInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get scanner metadata internal server error response
func (o *GetScannerMetadataInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetScannerMetadataInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
