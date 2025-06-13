package http

import (
	"fmt"
	"io"
	"strconv"
)

// Response represents an HTTP response.
type Response struct {
	Version    string  // HTTP version (e.g., "HTTP/1.1")
	StatusCode int     // HTTP status code (e.g., 200, 404)
	Headers    Headers // HTTP headers
	Body       []byte  // Response body
}

// NewResponse creates a new HTTP response.
func NewResponse(statusCode int) *Response {
	return &Response{
		Version:    VersionHTTP11, // Default to HTTP/1.1
		StatusCode: statusCode,
		Headers:    NewHeaders(),
	}
}

// Write writes the HTTP response to the provided writer.
func (r *Response) Write(w io.Writer) error {
	// Write the status line
	statusLine := fmt.Sprintf("%s %d %s\r\n",
		r.Version, r.StatusCode, StatusText(r.StatusCode))
	if _, err := w.Write([]byte(statusLine)); err != nil {
		return err
	}

	// Set the Content-Length header if not already set
	if r.Headers.Get(HeaderContentLength) == "" {
		r.Headers.Set(HeaderContentLength, strconv.Itoa(len(r.Body)))
	}

	// Write headers
	for key, value := range r.Headers {
		header := fmt.Sprintf("%s: %s\r\n", key, value)
		if _, err := w.Write([]byte(header)); err != nil {
			return err
		}
	}

	// Write a blank line to separate headers from the body
	if _, err := w.Write([]byte("\r\n")); err != nil {
		return err
	}

	// Write the body
	if len(r.Body) > 0 {
		if _, err := w.Write(r.Body); err != nil {
			return err
		}
	}

	return nil
}
