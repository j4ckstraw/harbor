// Code generated by go-swagger; DO NOT EDIT.

package systeminfo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// GetCertOKCode is the HTTP code returned for type GetCertOK
const GetCertOKCode int = 200

/*GetCertOK Get default root certificate successfully.

swagger:response getCertOK
*/
type GetCertOK struct {
	/*To set the filename of the downloaded file.

	 */
	ContentDisposition string `json:"Content-Disposition"`

	/*
	  In: Body
	*/
	Payload io.ReadCloser `json:"body,omitempty"`
}

// NewGetCertOK creates GetCertOK with default headers values
func NewGetCertOK() *GetCertOK {

	return &GetCertOK{}
}

// WithContentDisposition adds the contentDisposition to the get cert o k response
func (o *GetCertOK) WithContentDisposition(contentDisposition string) *GetCertOK {
	o.ContentDisposition = contentDisposition
	return o
}

// SetContentDisposition sets the contentDisposition to the get cert o k response
func (o *GetCertOK) SetContentDisposition(contentDisposition string) {
	o.ContentDisposition = contentDisposition
}

// WithPayload adds the payload to the get cert o k response
func (o *GetCertOK) WithPayload(payload io.ReadCloser) *GetCertOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get cert o k response
func (o *GetCertOK) SetPayload(payload io.ReadCloser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCertOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Content-Disposition

	contentDisposition := o.ContentDisposition
	if contentDisposition != "" {
		rw.Header().Set("Content-Disposition", contentDisposition)
	}

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetCertNotFoundCode is the HTTP code returned for type GetCertNotFound
const GetCertNotFoundCode int = 404

/*GetCertNotFound Not found the default root certificate.

swagger:response getCertNotFound
*/
type GetCertNotFound struct {
}

// NewGetCertNotFound creates GetCertNotFound with default headers values
func NewGetCertNotFound() *GetCertNotFound {

	return &GetCertNotFound{}
}

// WriteResponse to the client
func (o *GetCertNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetCertInternalServerErrorCode is the HTTP code returned for type GetCertInternalServerError
const GetCertInternalServerErrorCode int = 500

/*GetCertInternalServerError Internal server error

swagger:response getCertInternalServerError
*/
type GetCertInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetCertInternalServerError creates GetCertInternalServerError with default headers values
func NewGetCertInternalServerError() *GetCertInternalServerError {

	return &GetCertInternalServerError{}
}

// WithXRequestID adds the xRequestId to the get cert internal server error response
func (o *GetCertInternalServerError) WithXRequestID(xRequestID string) *GetCertInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get cert internal server error response
func (o *GetCertInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get cert internal server error response
func (o *GetCertInternalServerError) WithPayload(payload *models.Errors) *GetCertInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get cert internal server error response
func (o *GetCertInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCertInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
