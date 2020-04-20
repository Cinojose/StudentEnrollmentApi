// Code generated by go-swagger; DO NOT EDIT.

package student

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// FindStudentsByStatusHandlerFunc turns a function with the right signature into a find students by status handler
type FindStudentsByStatusHandlerFunc func(FindStudentsByStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FindStudentsByStatusHandlerFunc) Handle(params FindStudentsByStatusParams) middleware.Responder {
	return fn(params)
}

// FindStudentsByStatusHandler interface for that can handle valid find students by status params
type FindStudentsByStatusHandler interface {
	Handle(FindStudentsByStatusParams) middleware.Responder
}

// NewFindStudentsByStatus creates a new http.Handler for the find students by status operation
func NewFindStudentsByStatus(ctx *middleware.Context, handler FindStudentsByStatusHandler) *FindStudentsByStatus {
	return &FindStudentsByStatus{Context: ctx, Handler: handler}
}

/*FindStudentsByStatus swagger:route GET /fetchStudents student findStudentsByStatus

Finds students by multiple options

Multiple status values can be provided with comma separated strings

*/
type FindStudentsByStatus struct {
	Context *middleware.Context
	Handler FindStudentsByStatusHandler
}

func (o *FindStudentsByStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewFindStudentsByStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
