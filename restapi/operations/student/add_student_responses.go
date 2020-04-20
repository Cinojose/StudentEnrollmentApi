// Code generated by go-swagger; DO NOT EDIT.

package student

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// AddStudentOKCode is the HTTP code returned for type AddStudentOK
const AddStudentOKCode int = 200

/*AddStudentOK Student updated

swagger:response addStudentOK
*/
type AddStudentOK struct {
}

// NewAddStudentOK creates AddStudentOK with default headers values
func NewAddStudentOK() *AddStudentOK {

	return &AddStudentOK{}
}

// WriteResponse to the client
func (o *AddStudentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// AddStudentMethodNotAllowedCode is the HTTP code returned for type AddStudentMethodNotAllowed
const AddStudentMethodNotAllowedCode int = 405

/*AddStudentMethodNotAllowed Method not allowed

swagger:response addStudentMethodNotAllowed
*/
type AddStudentMethodNotAllowed struct {
}

// NewAddStudentMethodNotAllowed creates AddStudentMethodNotAllowed with default headers values
func NewAddStudentMethodNotAllowed() *AddStudentMethodNotAllowed {

	return &AddStudentMethodNotAllowed{}
}

// WriteResponse to the client
func (o *AddStudentMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(405)
}

// AddStudentInternalServerErrorCode is the HTTP code returned for type AddStudentInternalServerError
const AddStudentInternalServerErrorCode int = 500

/*AddStudentInternalServerError Invalid Input

swagger:response addStudentInternalServerError
*/
type AddStudentInternalServerError struct {
}

// NewAddStudentInternalServerError creates AddStudentInternalServerError with default headers values
func NewAddStudentInternalServerError() *AddStudentInternalServerError {

	return &AddStudentInternalServerError{}
}

// WriteResponse to the client
func (o *AddStudentInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}