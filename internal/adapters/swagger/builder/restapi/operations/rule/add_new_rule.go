// Code generated by go-swagger; DO NOT EDIT.

package rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AddNewRuleHandlerFunc turns a function with the right signature into a add new rule handler
type AddNewRuleHandlerFunc func(AddNewRuleParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddNewRuleHandlerFunc) Handle(params AddNewRuleParams) middleware.Responder {
	return fn(params)
}

// AddNewRuleHandler interface for that can handle valid add new rule params
type AddNewRuleHandler interface {
	Handle(AddNewRuleParams) middleware.Responder
}

// NewAddNewRule creates a new http.Handler for the add new rule operation
func NewAddNewRule(ctx *middleware.Context, handler AddNewRuleHandler) *AddNewRule {
	return &AddNewRule{Context: ctx, Handler: handler}
}

/*AddNewRule swagger:route POST /rule/add/{musrvId} rule addNewRule

build rule and query structures

Endpoint is used to generate microservice rule and query structures


*/
type AddNewRule struct {
	Context *middleware.Context
	Handler AddNewRuleHandler
}

func (o *AddNewRule) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddNewRuleParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
