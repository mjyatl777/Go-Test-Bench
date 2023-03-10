// Code generated by go-swagger; DO NOT EDIT.

package xss

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// XSSGetQueryReflectedXSSHandlerFunc turns a function with the right signature into a XSS get query reflected XSS handler
type XSSGetQueryReflectedXSSHandlerFunc func(XSSGetQueryReflectedXSSParams) middleware.Responder

// Handle executing the request and returning a response
func (fn XSSGetQueryReflectedXSSHandlerFunc) Handle(params XSSGetQueryReflectedXSSParams) middleware.Responder {
	return fn(params)
}

// XSSGetQueryReflectedXSSHandler interface for that can handle valid XSS get query reflected XSS params
type XSSGetQueryReflectedXSSHandler interface {
	Handle(XSSGetQueryReflectedXSSParams) middleware.Responder
}

// NewXSSGetQueryReflectedXSS creates a new http.Handler for the XSS get query reflected XSS operation
func NewXSSGetQueryReflectedXSS(ctx *middleware.Context, handler XSSGetQueryReflectedXSSHandler) *XSSGetQueryReflectedXSS {
	return &XSSGetQueryReflectedXSS{Context: ctx, Handler: handler}
}

/*
	XSSGetQueryReflectedXSS swagger:route GET /xss/reflectedXss/query/{safety} xss xssGetQueryReflectedXss

demonstrates Reflected XSS via query, with vulnerable function reflectedXss
*/
type XSSGetQueryReflectedXSS struct {
	Context *middleware.Context
	Handler XSSGetQueryReflectedXSSHandler
}

func (o *XSSGetQueryReflectedXSS) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewXSSGetQueryReflectedXSSParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
