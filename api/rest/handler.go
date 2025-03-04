package rest

import "net/http"

type IHTTPHandlerProvider interface {
	GetHTTPHandler() []*HTTPHandler
}
type HTTPHandler struct {
	// ACLBindings
	IHTTPHandlerProvider
	Version    uint8
	Method     string
	Path       string
	Middleware []func(http.Handler) http.Handler
	Func       func(*SessionContext)
}
