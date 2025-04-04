package router

import (
	"net/http"

	"github.com/ARtorias742/myrouter/matcher"
)

type Route struct {
	method  string
	path    string
	handler http.HandlerFunc
}

type Router struct {
	routes  []Route
	matcher *matcher.Matcher
}

func NewRouter() *Router {
	return &Router{
		routes:  make([]Route, 0),
		matcher: matcher.NewMatcher(),
	}
}

func (r *Router) Handle(method, path string, handler http.HandlerFunc) {
	r.routes = append(r.routes, Route{
		method:  method,
		path:    path,
		handler: handler,
	})
}

func (r *Router) Get(path string, handler http.HandlerFunc) {
	r.Handle(http.MethodGet, path, handler)
}

func (r *Router) Post(path string, handler http.HandlerFunc) {
	r.Handle(http.MethodPost, path, handler)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if route.method == req.Method && r.matcher.Match(route.path, req.URL.Path) {
			route.handler(w, req)
			return
		}
	}
	http.NotFound(w, req)
}
