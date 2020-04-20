// Code generated by go-swagger; DO NOT EDIT.

package student

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteStudentOKCode is the HTTP code returned for type DeleteStudentOK
const DeleteStudentOKCode int = 200

/*DeleteStudentOK Student deleted

swagger:response deleteStudentOK
*/
type DeleteStudentOK struct {
}

// NewDeleteStudentOK creates DeleteStudentOK with default headers values
func NewDeleteStudentOK() *DeleteStudentOK {

	return &DeleteStudentOK{}
}

// WriteResponse to the client
func (o *DeleteStudentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeleteStudentBadRequestCode is the HTTP code returned for type DeleteStudentBadRequest
const DeleteStudentBadRequestCode int = 400

/*DeleteStudentBadRequest Invalid ID supplied

swagger:response deleteStudentBadRequest
*/
type DeleteStudentBadRequest struct {
}

// NewDeleteStudentBadRequest creates DeleteStudentBadRequest with default headers values
func NewDeleteStudentBadRequest() *DeleteStudentBadRequest {

	return &DeleteStudentBadRequest{}
}

// WriteResponse to the client
func (o *DeleteStudentBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// DeleteStudentNotFoundCode is the HTTP code returned for type DeleteStudentNotFound
const DeleteStudentNotFoundCode int = 404

/*DeleteStudentNotFound Pet not found

swagger:response deleteStudentNotFound
*/
type DeleteStudentNotFound struct {
}

// NewDeleteStudentNotFound creates DeleteStudentNotFound with default headers values
func NewDeleteStudentNotFound() *DeleteStudentNotFound {

	return &DeleteStudentNotFound{}
}

// WriteResponse to the client
func (o *DeleteStudentNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}