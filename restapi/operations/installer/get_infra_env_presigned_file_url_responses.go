// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// GetInfraEnvPresignedFileURLOKCode is the HTTP code returned for type GetInfraEnvPresignedFileURLOK
const GetInfraEnvPresignedFileURLOKCode int = 200

/*GetInfraEnvPresignedFileURLOK Success.

swagger:response getInfraEnvPresignedFileUrlOK
*/
type GetInfraEnvPresignedFileURLOK struct {

	/*
	  In: Body
	*/
	Payload *models.PresignedURL `json:"body,omitempty"`
}

// NewGetInfraEnvPresignedFileURLOK creates GetInfraEnvPresignedFileURLOK with default headers values
func NewGetInfraEnvPresignedFileURLOK() *GetInfraEnvPresignedFileURLOK {

	return &GetInfraEnvPresignedFileURLOK{}
}

// WithPayload adds the payload to the get infra env presigned file Url o k response
func (o *GetInfraEnvPresignedFileURLOK) WithPayload(payload *models.PresignedURL) *GetInfraEnvPresignedFileURLOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get infra env presigned file Url o k response
func (o *GetInfraEnvPresignedFileURLOK) SetPayload(payload *models.PresignedURL) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInfraEnvPresignedFileURLOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetInfraEnvPresignedFileURLUnauthorizedCode is the HTTP code returned for type GetInfraEnvPresignedFileURLUnauthorized
const GetInfraEnvPresignedFileURLUnauthorizedCode int = 401

/*GetInfraEnvPresignedFileURLUnauthorized Unauthorized.

swagger:response getInfraEnvPresignedFileUrlUnauthorized
*/
type GetInfraEnvPresignedFileURLUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewGetInfraEnvPresignedFileURLUnauthorized creates GetInfraEnvPresignedFileURLUnauthorized with default headers values
func NewGetInfraEnvPresignedFileURLUnauthorized() *GetInfraEnvPresignedFileURLUnauthorized {

	return &GetInfraEnvPresignedFileURLUnauthorized{}
}

// WithPayload adds the payload to the get infra env presigned file Url unauthorized response
func (o *GetInfraEnvPresignedFileURLUnauthorized) WithPayload(payload *models.InfraError) *GetInfraEnvPresignedFileURLUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get infra env presigned file Url unauthorized response
func (o *GetInfraEnvPresignedFileURLUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInfraEnvPresignedFileURLUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetInfraEnvPresignedFileURLForbiddenCode is the HTTP code returned for type GetInfraEnvPresignedFileURLForbidden
const GetInfraEnvPresignedFileURLForbiddenCode int = 403

/*GetInfraEnvPresignedFileURLForbidden Forbidden.

swagger:response getInfraEnvPresignedFileUrlForbidden
*/
type GetInfraEnvPresignedFileURLForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewGetInfraEnvPresignedFileURLForbidden creates GetInfraEnvPresignedFileURLForbidden with default headers values
func NewGetInfraEnvPresignedFileURLForbidden() *GetInfraEnvPresignedFileURLForbidden {

	return &GetInfraEnvPresignedFileURLForbidden{}
}

// WithPayload adds the payload to the get infra env presigned file Url forbidden response
func (o *GetInfraEnvPresignedFileURLForbidden) WithPayload(payload *models.InfraError) *GetInfraEnvPresignedFileURLForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get infra env presigned file Url forbidden response
func (o *GetInfraEnvPresignedFileURLForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInfraEnvPresignedFileURLForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetInfraEnvPresignedFileURLNotFoundCode is the HTTP code returned for type GetInfraEnvPresignedFileURLNotFound
const GetInfraEnvPresignedFileURLNotFoundCode int = 404

/*GetInfraEnvPresignedFileURLNotFound Error.

swagger:response getInfraEnvPresignedFileUrlNotFound
*/
type GetInfraEnvPresignedFileURLNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetInfraEnvPresignedFileURLNotFound creates GetInfraEnvPresignedFileURLNotFound with default headers values
func NewGetInfraEnvPresignedFileURLNotFound() *GetInfraEnvPresignedFileURLNotFound {

	return &GetInfraEnvPresignedFileURLNotFound{}
}

// WithPayload adds the payload to the get infra env presigned file Url not found response
func (o *GetInfraEnvPresignedFileURLNotFound) WithPayload(payload *models.Error) *GetInfraEnvPresignedFileURLNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get infra env presigned file Url not found response
func (o *GetInfraEnvPresignedFileURLNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInfraEnvPresignedFileURLNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetInfraEnvPresignedFileURLMethodNotAllowedCode is the HTTP code returned for type GetInfraEnvPresignedFileURLMethodNotAllowed
const GetInfraEnvPresignedFileURLMethodNotAllowedCode int = 405

/*GetInfraEnvPresignedFileURLMethodNotAllowed Method Not Allowed.

swagger:response getInfraEnvPresignedFileUrlMethodNotAllowed
*/
type GetInfraEnvPresignedFileURLMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetInfraEnvPresignedFileURLMethodNotAllowed creates GetInfraEnvPresignedFileURLMethodNotAllowed with default headers values
func NewGetInfraEnvPresignedFileURLMethodNotAllowed() *GetInfraEnvPresignedFileURLMethodNotAllowed {

	return &GetInfraEnvPresignedFileURLMethodNotAllowed{}
}

// WithPayload adds the payload to the get infra env presigned file Url method not allowed response
func (o *GetInfraEnvPresignedFileURLMethodNotAllowed) WithPayload(payload *models.Error) *GetInfraEnvPresignedFileURLMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get infra env presigned file Url method not allowed response
func (o *GetInfraEnvPresignedFileURLMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInfraEnvPresignedFileURLMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetInfraEnvPresignedFileURLInternalServerErrorCode is the HTTP code returned for type GetInfraEnvPresignedFileURLInternalServerError
const GetInfraEnvPresignedFileURLInternalServerErrorCode int = 500

/*GetInfraEnvPresignedFileURLInternalServerError Error.

swagger:response getInfraEnvPresignedFileUrlInternalServerError
*/
type GetInfraEnvPresignedFileURLInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetInfraEnvPresignedFileURLInternalServerError creates GetInfraEnvPresignedFileURLInternalServerError with default headers values
func NewGetInfraEnvPresignedFileURLInternalServerError() *GetInfraEnvPresignedFileURLInternalServerError {

	return &GetInfraEnvPresignedFileURLInternalServerError{}
}

// WithPayload adds the payload to the get infra env presigned file Url internal server error response
func (o *GetInfraEnvPresignedFileURLInternalServerError) WithPayload(payload *models.Error) *GetInfraEnvPresignedFileURLInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get infra env presigned file Url internal server error response
func (o *GetInfraEnvPresignedFileURLInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInfraEnvPresignedFileURLInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetInfraEnvPresignedFileURLNotImplementedCode is the HTTP code returned for type GetInfraEnvPresignedFileURLNotImplemented
const GetInfraEnvPresignedFileURLNotImplementedCode int = 501

/*GetInfraEnvPresignedFileURLNotImplemented Not implemented.

swagger:response getInfraEnvPresignedFileUrlNotImplemented
*/
type GetInfraEnvPresignedFileURLNotImplemented struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetInfraEnvPresignedFileURLNotImplemented creates GetInfraEnvPresignedFileURLNotImplemented with default headers values
func NewGetInfraEnvPresignedFileURLNotImplemented() *GetInfraEnvPresignedFileURLNotImplemented {

	return &GetInfraEnvPresignedFileURLNotImplemented{}
}

// WithPayload adds the payload to the get infra env presigned file Url not implemented response
func (o *GetInfraEnvPresignedFileURLNotImplemented) WithPayload(payload *models.Error) *GetInfraEnvPresignedFileURLNotImplemented {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get infra env presigned file Url not implemented response
func (o *GetInfraEnvPresignedFileURLNotImplemented) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInfraEnvPresignedFileURLNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetInfraEnvPresignedFileURLServiceUnavailableCode is the HTTP code returned for type GetInfraEnvPresignedFileURLServiceUnavailable
const GetInfraEnvPresignedFileURLServiceUnavailableCode int = 503

/*GetInfraEnvPresignedFileURLServiceUnavailable Unavailable.

swagger:response getInfraEnvPresignedFileUrlServiceUnavailable
*/
type GetInfraEnvPresignedFileURLServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetInfraEnvPresignedFileURLServiceUnavailable creates GetInfraEnvPresignedFileURLServiceUnavailable with default headers values
func NewGetInfraEnvPresignedFileURLServiceUnavailable() *GetInfraEnvPresignedFileURLServiceUnavailable {

	return &GetInfraEnvPresignedFileURLServiceUnavailable{}
}

// WithPayload adds the payload to the get infra env presigned file Url service unavailable response
func (o *GetInfraEnvPresignedFileURLServiceUnavailable) WithPayload(payload *models.Error) *GetInfraEnvPresignedFileURLServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get infra env presigned file Url service unavailable response
func (o *GetInfraEnvPresignedFileURLServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInfraEnvPresignedFileURLServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
