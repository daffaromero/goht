package http

// Constants for HTTP methods
const (
	MethodGet     = "GET"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodDelete  = "DELETE"
	MethodPatch   = "PATCH"
	MethodHead    = "HEAD"
	MethodOptions = "OPTIONS"
)

// Constants for common HTTP versions
const (
	VersionHTTP10     = "HTTP/1.0"
	VersionHTTP11     = "HTTP/1.1"
	VersionHTTP20     = "HTTP/2.0"
	VersionHTTP30     = "HTTP/3.0"
	VersionHTTP11Plus = "HTTP/1.1+"   // Represents any version greater than or equal to HTTP/1.1
	VersionHTTP20Plus = "HTTP/2.0+"   // Represents any version greater than or equal to HTTP/2.0
	VersionHTTP30Plus = "HTTP/3.0+"   // Represents any version greater than or equal to HTTP/3.0
	VersionHTTPLatest = "HTTP/Latest" // Represents the latest HTTP version
)

// Request represents an HTTP request.
type Request struct {
	Method  string  // HTTP method (GET, POST, etc.)
	Path    string  // Request path
	Version string  // HTTP version (e.g., "HTTP/1.1")
	Headers Headers // HTTP headers
	Body    []byte  // Request body
}

// NewRequest creates a new HTTP request with the specified method, path, and version.
func NewRequest(method, path, version string) *Request {
	return &Request{
		Method:  method,
		Path:    path,
		Version: version,
		Headers: NewHeaders(),
	}
}
