package router

import (
	"net/http"

	"github.com/ARtorias742/matcher"
	"github.com/ARtorias742/middleware"
)

type Route struct {
	name        string
	method      string
	path        string
	host        string
	handler     http.HandlerFunc
	middlewares []middleware.Middleware
}

type Router struct {
	routes  []*Route
	matcher *matcher.Matcher
}

func NewRouter() *Router {
	return &Router{
		routes:  make([]*Route, 0),
		matcher: matcher.NewMatcher(),
	}
}

func (r *Router) Handle(method, path string, handler http.HandlerFunc) *Route {
	route := &Route{
		method:  method,
		path:    path,
		handler: handler,
	}
	r.routes = append(r.routes, route)
	return route
}

func (r *Router) Get(path string, handler http.HandlerFunc) *Route {
	return r.Handle(http.MethodGet, path, handler)
}

func (r *Router) Post(path string, handler http.HandlerFunc) *Route {
	return r.Handle(http.MethodPost, path, handler)
}

// Move Name, Host, and Use to Route type for chaining
func (rt *Route) Name(name string) *Route {
	rt.name = name
	return rt
}

func (rt *Route) Host(host string) *Route {
	rt.host = host
	return rt
}

func (rt *Route) Use(mw ...middleware.Middleware) *Route {
	rt.middlewares = append(rt.middlewares, mw...)
	return rt
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if route.method != req.Method {
			continue
		}
		if route.host != "" && route.host != req.Host {
			continue
		}
		params, matched := r.matcher.Match(route.path, req.URL.Path)
		if !matched {
			continue
		}

		// Apply middleware chain
		handler := route.handler
		for i := len(route.middlewares) - 1; i >= 0; i-- {
			handler = route.middlewares[i](handler)
		}

		// Set path parameters in request context if needed
		ctx := matcher.WithParams(req.Context(), params)
		handler.ServeHTTP(w, req.WithContext(ctx))
		return
	}
	http.NotFound(w, req)
}

func (r *Router) FindByName(name string) *Route {
	for _, route := range r.routes {
		if route.name == name {
			return route
		}
	}
	return nil
}
