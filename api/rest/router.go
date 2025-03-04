package rest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Router interface {
	http.Handler
	Route(pattern string, callback func(r Router))
	Register(handlers []IHTTPHandlerProvider)
	Use(middlewares ...func(http.Handler) http.Handler)
}

type GRouter struct {
	r chi.Router
}

func NewGRouter() *GRouter {
	return &GRouter{r: chi.NewRouter()}
}

func (g *GRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.r.ServeHTTP(w, r)
}

func (g *GRouter) Route(pattern string, callback func(r Router)) {
	g.r.Route(pattern, func(r chi.Router) {
		callback(&GRouter{r: r})
	})
}

func (g *GRouter) Use(middlewares ...func(http.Handler) http.Handler) {
	g.r.Use(middlewares...)
}

func (g *GRouter) Register(handlers []IHTTPHandlerProvider) {
	for _, handler := range handlers {
		for _, hlr := range handler.GetHTTPHandler() {
			path := fmt.Sprintf("/api/v%d/%s", hlr.Version, hlr.Path)
			log.Printf("Adding handler with Path - %s, and Method - %s", path, hlr.Method)

			wrapperFunc := APIWrapper(hlr.Func)
			switch hlr.Method {
			case http.MethodGet:
				g.r.Get(path, wrapperFunc)
			case http.MethodPost:
				g.r.Post(path, wrapperFunc)
			case http.MethodPut:
				g.r.Put(path, wrapperFunc)
			case http.MethodDelete:
				g.r.Delete(path, wrapperFunc)
			case http.MethodPatch:
				g.r.Patch(path, wrapperFunc)
			default:
				log.Fatal("Invalid method")
			}
		}
	}
}

func APIWrapper(restHandlerFunc func(c *SessionContext)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c := NewSessionContext(r, w)
		restHandlerFunc(c)
	}
}
