package http

import "strings"

// Route represents a single route in the HTTP router.
type Route struct {
	Method  string        // HTTP method (GET, POST, etc.)
	Pattern string        // URL pattern for the route
	Handler HandlerFunc // Function to handle the request
}

// Router represents an HTTP router that maps routes to handlers.
type Router struct {
	Routes []Route // List of routes
}

// NewRouter creates a new HTTP router.
func NewRouter() *Router {
	return &Router{
		Routes: make([]Route, 0),
	}
}

// Handle registers a new route with the router.
func (r *Router) Handle(method, pattern string, handler HandlerFunc) {
	r.Routes = append(r.Routes, Route{
		Method:  method,
		Pattern: pattern,
		Handler: handler,
	})
}

// GET registers a GET handler
func (r *Router) GET(pattern string, handler HandlerFunc) {
	r.Handle(MethodGet, pattern, handler)
}

// POST registers a POST handler
func (r *Router) POST(pattern string, handler HandlerFunc) {
	r.Handle(MethodPost, pattern, handler)
}

// PUT registers a PUT handler
func (r *Router) PUT(pattern string, handler HandlerFunc) {
	r.Handle(MethodPut, pattern, handler)
}

// DELETE registers a DELETE handler
func (r *Router) DELETE(pattern string, handler HandlerFunc) {
	r.Handle(MethodDelete, pattern, handler)
}

// PATCH registers a PATCH handler
func (r *Router) PATCH(pattern string, handler HandlerFunc) {
	r.Handle(MethodPatch, pattern, handler)
}

// HEAD registers a HEAD handler
func (r *Router) HEAD(pattern string, handler HandlerFunc) {
	r.Handle(MethodHead, pattern, handler)
}

// OPTIONS registers an OPTIONS handler
func (r *Router) OPTIONS(pattern string, handler HandlerFunc) {
	r.Handle(MethodOptions, pattern, handler)
}

// ServeHTTP serves an HTTP request by finding the appropriate route and executing its handler.
func (r *Router) ServeHTTP(req *Request, resp *Response) error {
	for _, route := range r.Routes {

	}
}

// matchRoute checks if the request matches the route pattern.
func (r *Router) matchRoute(route Route, req *Request) bool {
	// Check if the method matches
	if route.Method != req.Method {
		return false
	}

	// Check if the path matches the pattern
	if route.Pattern != req.Path {
		return false
	}

	return true
}

// matchPattern matches URL patterns with wildcards and named parameters
func (r *Router) matchPattern(pattern, path string) (map[string]string, bool) {
		// Remove query parameters from the path
	if queryIndex := strings.Index(path, "?"); queryIndex != -1 {
		path = path[:queryIndex]
	}
	
	// Exact matches
	if pattern == path {
		return make(map[string]string), true
	}

	// Handle wildcard patterns (e.g., "/api/*")
	patternParts := strings.Split(strings.Trim(pattern, "/"), "/")
	pathParts := strings.Split(strings.Trim(path, "/"), "/")


	// Handle root path
	if pattern == "/" && path == "/" {
		return make(map[string]string), true
	}

	if len(patternParts) != len(pathParts) {
		return nil, false
	}

	params := make(map[string]string)

	for i, patternPart := range patternParts {
		pathPart := pathParts[i]

		// Named parameters (starts with ':')
		if strings.HasPrefix(patternpart, ":") {
			paramName := patternPart[1:] // Remove the leading ':'
			params[paramName] = pathPart
			continue
		}

		// If the pattern part is not a wildcard and does not match the path part, return false
		if patternPart != pathPart {
			return false
		}
	}
}