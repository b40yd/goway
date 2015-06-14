package goway

import (
	"fmt"
	"net/http"
	"path/filepath"
	"regexp"
)

const (
	CONNECT = "CONNECT"
	DELETE  = "DELETE"
	GET     = "GET"
	HEAD    = "HEAD"
	OPTIONS = "OPTIONS"
	PATCH   = "PATCH"
	POST    = "POST"
	PUT     = "PUT"
	TRACE   = "TRACE"
	ANY     = "*"
)

var Params map[string]string

type Router interface {
	Get(pattern string, handler Handler)
	// Patch adds a route for a HTTP PATCH request to the specified matching pattern.
	Patch(string, Handler)
	// Post adds a route for a HTTP POST request to the specified matching pattern.
	Post(string, Handler)
	// Put adds a route for a HTTP PUT request to the specified matching pattern.
	Put(string, Handler)
	// Delete adds a route for a HTTP DELETE request to the specified matching pattern.
	Delete(string, Handler)
	// Any adds a route for any HTTP method request to the specified matching pattern.
	Any(string, Handler)
	// AddRoute adds a route for a given HTTP method request to the specified matching pattern.
	Static(string, string)
	NotFound(handlers Handler)
}

type Routes struct {
	routes   []*router
	handlers []Handler
	notFound []Handler
	params  map[string]string
}

type router struct {
	method  string
	regex   *regexp.Regexp
	handler Handler
}

// Get adds a new Route for GET requests.
func (r *Routes) Get(pattern string, handler Handler) {
	r.AddRouter(GET, pattern, handler)
}

// Put adds a new Route for PUT requests.
func (r *Routes) Put(pattern string, handler Handler) {
	r.AddRouter(PUT, pattern, handler)
}

// Del adds a new Route for DELETE requests.
func (r *Routes) Delete(pattern string, handler Handler) {
	r.AddRouter(DELETE, pattern, handler)
}

// Patch adds a new Route for PATCH requests.
func (r *Routes) Patch(pattern string, handler Handler) {
	r.AddRouter(PATCH, pattern, handler)
}

// Post adds a new Route for POST requests.
func (r *Routes) Post(pattern string, handler Handler) {
	r.AddRouter(POST, pattern, handler)
}

// Post adds a new Route for POST requests.
func (r *Routes) Any(pattern string, handler Handler) {
	r.AddRouter(POST, pattern, handler)
}

// Adds a new Route for Static http requests. Serves
// static files from the specified directory
func (r *Routes) Static(pattern string, dir string) {
	//append a regex to the param to match everything
	// that comes after the prefix
	pattern = pattern + "(.+)"
	r.AddRouter(GET, pattern, func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Clean(r.URL.Path)
		path = filepath.Join(dir, path)
		http.ServeFile(w, r, path)
	})
}

func (r *Routes) NotFound(handler Handler) {
	r.notFound = []Handler{handler}
}

func (r *Routes) AddRouter(method string, pattern string, handler Handler) {

	var routeReg1 = regexp.MustCompile(`:[^/#?()\.\\]+`)
	var routeReg2 = regexp.MustCompile(`\*\*`)

	route := &router{}
	pattern = routeReg1.ReplaceAllStringFunc(pattern, func(m string) string {
		return fmt.Sprintf(`(?P<%s>[^/#?]+)`, m[1:])
	})
	var index int
	pattern = routeReg2.ReplaceAllStringFunc(pattern, func(m string) string {
		index++
		return fmt.Sprintf(`(?P<_%d>[^#?]*)`, index)
	})
	pattern += `\/?`
	route.regex = regexp.MustCompile(pattern)
	route.method = method
	route.handler = handler

	//and finally append to the list of Routes
	r.routes = append(r.routes, route)
}

// Match the HTTP method
func (routes *Routes) matchMethod(method string, r *router) bool {
	return r.method == "*" || method == r.method || (method == "HEAD" && r.method == "GET")
}

func (routes *Routes) matchFunc(r *router, matches []string, path string) (bool, map[string]string) {
	if len(matches) > 0 && matches[0] == path {
		params := make(map[string]string)
		for i, name := range r.regex.SubexpNames() {
			if len(name) > 0 {
				params[name] = matches[i]
			}
		}
		//fmt.Printf("%v   ",params)
		return true, params
	}
	return false, nil
}

func (routes *Routes) Match(method string, route *router, r *http.Request) (bool, Handler) {
	var ok bool = false
	//match router
	if routes.matchMethod(method, route) {
		matches := route.regex.FindStringSubmatch(r.URL.Path)
		ok, Params = routes.matchFunc(route, matches, r.URL.Path)
	}
	return ok, Params
}

func (routes *Routes) Handler(w http.ResponseWriter, r *http.Request, c Context) {
	// init response writer
	rw := &responseWriter{responseWriter: w}

	for _, route := range routes.routes {
		if ok, vals := routes.Match(r.Method, route, r); ok {
			handlers := make([]Handler, 0)
			routes.handlers = append(handlers, route.handler)
			c.Map(vals)
			break
		} else {
			// not found handler,set routes.notFound
			routes.handlers = routes.notFound //make([]Handler, 0)
		}
	}
	//
	ctx := &context{c, routes.handlers, func() {}, NewResponseWriter(w), 0}
	c.MapTo(ctx, (*Context)(nil))
	c.MapTo(rw, (*Router)(nil))
	ctx.bootstrap()

}

func NewRouter() *Routes {
	return &Routes{handlers: []Handler{http.NotFound},params: make(map[string]string)}
}
