package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/luizalabs/teresa-api/models"
)

/*DeleteUserNoContent No content

swagger:response deleteUserNoContent
*/
type DeleteUserNoContent struct {
}

// NewDeleteUserNoContent creates DeleteUserNoContent with default headers values
func NewDeleteUserNoContent() *DeleteUserNoContent {
	return &DeleteUserNoContent{}
}

// WriteResponse to the client
func (o *DeleteUserNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

/*DeleteUserDefault Error

swagger:response deleteUserDefault
*/
type DeleteUserDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteUserDefault creates DeleteUserDefault with default headers values
func NewDeleteUserDefault(code int) *DeleteUserDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteUserDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete user default response
func (o *DeleteUserDefault) WithStatusCode(code int) *DeleteUserDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete user default response
func (o *DeleteUserDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete user default response
func (o *DeleteUserDefault) WithPayload(payload *models.Error) *DeleteUserDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user default response
func (o *DeleteUserDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
