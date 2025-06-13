package http

import "time"

// HandlerFunc is a function type that handles HTTP requests.
type HandlerFunc func(*Request, *Response) error

// Server represents an HTTP server.
type Server struct {
	Addr         string        // Address to listen on
	Router       *Router       // Router for handling requests
	ReadTimeout  time.Duration // Timeout for reading requests
	WriteTimeout time.Duration // Timeout for writing responses
}

// NewServer creates a new HTTP server.
