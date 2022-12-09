// Code generated by go-swagger; DO NOT EDIT.

package org

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteOrgOKCode is the HTTP code returned for type DeleteOrgOK
const DeleteOrgOKCode int = 200

/*DeleteOrgOK deletion success

swagger:response deleteOrgOK
*/
type DeleteOrgOK struct {

	/*
	  In: Body
	*/
	Payload *DeleteOrgOKBody `json:"body,omitempty"`
}

// NewDeleteOrgOK creates DeleteOrgOK with default headers values
func NewDeleteOrgOK() *DeleteOrgOK {

	return &DeleteOrgOK{}
}

// WithPayload adds the payload to the delete org o k response
func (o *DeleteOrgOK) WithPayload(payload *DeleteOrgOKBody) *DeleteOrgOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete org o k response
func (o *DeleteOrgOK) SetPayload(payload *DeleteOrgOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteOrgOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteOrgUnauthorizedCode is the HTTP code returned for type DeleteOrgUnauthorized
const DeleteOrgUnauthorizedCode int = 401

/*DeleteOrgUnauthorized Unauthorized

swagger:response deleteOrgUnauthorized
*/
type DeleteOrgUnauthorized struct {
}

// NewDeleteOrgUnauthorized creates DeleteOrgUnauthorized with default headers values
func NewDeleteOrgUnauthorized() *DeleteOrgUnauthorized {

	return &DeleteOrgUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteOrgUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// DeleteOrgForbiddenCode is the HTTP code returned for type DeleteOrgForbidden
const DeleteOrgForbiddenCode int = 403

/*DeleteOrgForbidden Forbidden

swagger:response deleteOrgForbidden
*/
type DeleteOrgForbidden struct {
}

// NewDeleteOrgForbidden creates DeleteOrgForbidden with default headers values
func NewDeleteOrgForbidden() *DeleteOrgForbidden {

	return &DeleteOrgForbidden{}
}

// WriteResponse to the client
func (o *DeleteOrgForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// DeleteOrgNotFoundCode is the HTTP code returned for type DeleteOrgNotFound
const DeleteOrgNotFoundCode int = 404

/*DeleteOrgNotFound resource not found

swagger:response deleteOrgNotFound
*/
type DeleteOrgNotFound struct {

	/*
	  In: Body
	*/
	Payload *DeleteOrgNotFoundBody `json:"body,omitempty"`
}

// NewDeleteOrgNotFound creates DeleteOrgNotFound with default headers values
func NewDeleteOrgNotFound() *DeleteOrgNotFound {

	return &DeleteOrgNotFound{}
}

// WithPayload adds the payload to the delete org not found response
func (o *DeleteOrgNotFound) WithPayload(payload *DeleteOrgNotFoundBody) *DeleteOrgNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete org not found response
func (o *DeleteOrgNotFound) SetPayload(payload *DeleteOrgNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteOrgNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}