package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/influxdata/chronograf/models"
)

/*GetLayoutsIDOK Returns the specified layout containing `cells`.

swagger:response getLayoutsIdOK
*/
type GetLayoutsIDOK struct {

	// In: body
	Payload *models.Layout `json:"body,omitempty"`
}

// NewGetLayoutsIDOK creates GetLayoutsIDOK with default headers values
func NewGetLayoutsIDOK() *GetLayoutsIDOK {
	return &GetLayoutsIDOK{}
}

// WithPayload adds the payload to the get layouts Id o k response
func (o *GetLayoutsIDOK) WithPayload(payload *models.Layout) *GetLayoutsIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get layouts Id o k response
func (o *GetLayoutsIDOK) SetPayload(payload *models.Layout) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLayoutsIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetLayoutsIDNotFound Unknown layout id

swagger:response getLayoutsIdNotFound
*/
type GetLayoutsIDNotFound struct {

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetLayoutsIDNotFound creates GetLayoutsIDNotFound with default headers values
func NewGetLayoutsIDNotFound() *GetLayoutsIDNotFound {
	return &GetLayoutsIDNotFound{}
}

// WithPayload adds the payload to the get layouts Id not found response
func (o *GetLayoutsIDNotFound) WithPayload(payload *models.Error) *GetLayoutsIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get layouts Id not found response
func (o *GetLayoutsIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLayoutsIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetLayoutsIDDefault Unexpected internal service error

swagger:response getLayoutsIdDefault
*/
type GetLayoutsIDDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetLayoutsIDDefault creates GetLayoutsIDDefault with default headers values
func NewGetLayoutsIDDefault(code int) *GetLayoutsIDDefault {
	if code <= 0 {
		code = 500
	}

	return &GetLayoutsIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get layouts ID default response
func (o *GetLayoutsIDDefault) WithStatusCode(code int) *GetLayoutsIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get layouts ID default response
func (o *GetLayoutsIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get layouts ID default response
func (o *GetLayoutsIDDefault) WithPayload(payload *models.Error) *GetLayoutsIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get layouts ID default response
func (o *GetLayoutsIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLayoutsIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}