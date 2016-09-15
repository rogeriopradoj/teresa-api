package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/luizalabs/tapi/models"
)

/*PartialUpdateAppOK App

swagger:response partialUpdateAppOK
*/
type PartialUpdateAppOK struct {

	// In: body
	Payload *models.App `json:"body,omitempty"`
}

// NewPartialUpdateAppOK creates PartialUpdateAppOK with default headers values
func NewPartialUpdateAppOK() *PartialUpdateAppOK {
	return &PartialUpdateAppOK{}
}

// WithPayload adds the payload to the partial update app o k response
func (o *PartialUpdateAppOK) WithPayload(payload *models.App) *PartialUpdateAppOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the partial update app o k response
func (o *PartialUpdateAppOK) SetPayload(payload *models.App) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PartialUpdateAppOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PartialUpdateAppDefault Unexpected error

swagger:response partialUpdateAppDefault
*/
type PartialUpdateAppDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewPartialUpdateAppDefault creates PartialUpdateAppDefault with default headers values
func NewPartialUpdateAppDefault(code int) *PartialUpdateAppDefault {
	if code <= 0 {
		code = 500
	}

	return &PartialUpdateAppDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the partial update app default response
func (o *PartialUpdateAppDefault) WithStatusCode(code int) *PartialUpdateAppDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the partial update app default response
func (o *PartialUpdateAppDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the partial update app default response
func (o *PartialUpdateAppDefault) WithPayload(payload *models.Error) *PartialUpdateAppDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the partial update app default response
func (o *PartialUpdateAppDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PartialUpdateAppDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
