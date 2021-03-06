// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "adder": Application Controllers
//
// Command:
// $ goagen
// --design=github.com/yu-ichiko/goa-security-example/design
// --out=$(GOPATH)/src/github.com/yu-ichiko/goa-security-example
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// OperandsController is the controller interface for the Operands actions.
type OperandsController interface {
	goa.Muxer
	Get(*GetOperandsContext) error
}

// MountOperandsController "mounts" a Operands resource controller on the given service.
func MountOperandsController(service *goa.Service, ctrl OperandsController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetOperandsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Get(rctx)
	}
	h = handleSecurity("basic_auth", h)
	service.Mux.Handle("GET", "/", ctrl.MuxHandler("get", h, nil))
	service.LogInfo("mount", "ctrl", "Operands", "action", "Get", "route", "GET /", "security", "basic_auth")
}
