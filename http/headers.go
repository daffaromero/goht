package http

import "strings"

// Constants for common HTTP headers.
const (
	HeaderContentType   = "Content-Type"
	HeaderContentLength = "Content-Length"
	HeaderHost          = "Host"
	HeaderUserAgent     = "User-Agent"
	HeaderAccept        = "Accept"
	HeaderAuthorization = "Authorization"
	HeaderCacheControl  = "Cache-Control"
	HeaderSetCookie     = "Set-Cookie"
	HeaderCookie        = "Cookie"
	HeaderLocation      = "Location"
	HeaderConnection    = "Connection"
)

// Headers is a map of HTTP headers.
type Headers map[string]string

// NewHeaders creates a new Headers instance.
func NewHeaders() Headers {
	return make(Headers)
}

// Get retrieves the value of a header by its key, returning the value in lowercase.
func (h Headers) Get(key string) string {
	if value, exists := h[key]; exists {
		return strings.ToLower(value)
	}
	return ""
}

// Set sets the value of a header by its key. If the key already exists, it updates the value.
func (h Headers) Set(key, value string) {
	h[strings.ToLower(key)] = value
}

// Add adds a new header to the map. If the key already exists, it appends the value.
func (h Headers) Add(key, value string) {
	key = strings.ToLower(key)
	if existingValue, exists := h[key]; exists {
		h[key] = existingValue + ", " + value
	} else {
		h[key] = value
	}
}

// Del removes a header from the map by its key.
func (h Headers) Del(key string) {
	delete(h, strings.ToLower(key))
}
