// Code generated by go-swagger; DO NOT EDIT.

package rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ahamtat/logicfm/internal/adapters/swagger/builder/models"
)

// DeleteRuleNoContentCode is the HTTP code returned for type DeleteRuleNoContent
const DeleteRuleNoContentCode int = 204

/*DeleteRuleNoContent Deleted

swagger:response deleteRuleNoContent
*/
type DeleteRuleNoContent struct {
}

// NewDeleteRuleNoContent creates DeleteRuleNoContent with default headers values
func NewDeleteRuleNoContent() *DeleteRuleNoContent {

	return &DeleteRuleNoContent{}
}

// WriteResponse to the client
func (o *DeleteRuleNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

/*DeleteRuleDefault error

swagger:response deleteRuleDefault
*/
type DeleteRuleDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteRuleDefault creates DeleteRuleDefault with default headers values
func NewDeleteRuleDefault(code int) *DeleteRuleDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteRuleDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete rule default response
func (o *DeleteRuleDefault) WithStatusCode(code int) *DeleteRuleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete rule default response
func (o *DeleteRuleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete rule default response
func (o *DeleteRuleDefault) WithPayload(payload *models.Error) *DeleteRuleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete rule default response
func (o *DeleteRuleDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteRuleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
