package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/luizalabs/paas/api/models"
)

/*UpdateAppOK App

swagger:response updateAppOK
*/
type UpdateAppOK struct {

	// In: body
	Payload *models.App `json:"body,omitempty"`
}

// NewUpdateAppOK creates UpdateAppOK with default headers values
func NewUpdateAppOK() *UpdateAppOK {
	return &UpdateAppOK{}
}

// WithPayload adds the payload to the update app o k response
func (o *UpdateAppOK) WithPayload(payload *models.App) *UpdateAppOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update app o k response
func (o *UpdateAppOK) SetPayload(payload *models.App) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateAppOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UpdateAppUnauthorized User not authorized

swagger:response updateAppUnauthorized
*/
type UpdateAppUnauthorized struct {

	// In: body
	Payload *models.Unauthorized `json:"body,omitempty"`
}

// NewUpdateAppUnauthorized creates UpdateAppUnauthorized with default headers values
func NewUpdateAppUnauthorized() *UpdateAppUnauthorized {
	return &UpdateAppUnauthorized{}
}

// WithPayload adds the payload to the update app unauthorized response
func (o *UpdateAppUnauthorized) WithPayload(payload *models.Unauthorized) *UpdateAppUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update app unauthorized response
func (o *UpdateAppUnauthorized) SetPayload(payload *models.Unauthorized) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateAppUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UpdateAppForbidden User does not have the credentials to access this resource


swagger:response updateAppForbidden
*/
type UpdateAppForbidden struct {

	// In: body
	Payload *models.Unauthorized `json:"body,omitempty"`
}

// NewUpdateAppForbidden creates UpdateAppForbidden with default headers values
func NewUpdateAppForbidden() *UpdateAppForbidden {
	return &UpdateAppForbidden{}
}

// WithPayload adds the payload to the update app forbidden response
func (o *UpdateAppForbidden) WithPayload(payload *models.Unauthorized) *UpdateAppForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update app forbidden response
func (o *UpdateAppForbidden) SetPayload(payload *models.Unauthorized) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateAppForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UpdateAppDefault Error

swagger:response updateAppDefault
*/
type UpdateAppDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateAppDefault creates UpdateAppDefault with default headers values
func NewUpdateAppDefault(code int) *UpdateAppDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateAppDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update app default response
func (o *UpdateAppDefault) WithStatusCode(code int) *UpdateAppDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update app default response
func (o *UpdateAppDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update app default response
func (o *UpdateAppDefault) WithPayload(payload *models.Error) *UpdateAppDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update app default response
func (o *UpdateAppDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateAppDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}