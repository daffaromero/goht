package http

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// ParseRequest parses an HTTP request from a buffered reader.
func ParseRequest(reader *bufio.Reader) (*Request, error) {
	// Read the request line
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("failed to read request line: %w", err)
	}
	requestLine = strings.TrimSpace(requestLine)

	// Parse the request line
	parts := strings.SplitN(requestLine, " ", 3)
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid request line: %s", requestLine)
	}

	req := &Request{
		Method:  parts[0],
		Path:    parts[1],
		Version: parts[2],
		Headers: NewHeaders(),
	}

	// Parse headers
	if err := parseHeaders(reader, req.Headers); err != nil {
		return nil, fmt.Errorf("failed to parse headers: %w", err)
	}

	// Parse body if Content-Length is set
	if err := parseBody(reader, req); err != nil {
		return nil, fmt.Errorf("failed to parse body: %w", err)
	}

	return req, nil
}

// parseHeaders parses HTTP headers from the buffered reader.
func parseHeaders(reader *bufio.Reader, headers Headers) error {
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		line = strings.TrimSpace(line)

		// Stop reading headers on an empty line
		if line == "" {
			break
		}

		// Split header into key and value
		colonIndex := strings.Index(line, ":")
		if colonIndex == -1 {
			continue // Skip malformed headers
		}

		key := strings.TrimSpace(line[:colonIndex])
		value := strings.TrimSpace(line[colonIndex+1:])
		headers.Set(key, value)
	}

	return nil
}

// parseBody parses HTTP request body based on the Content-Length header.
func parseBody(reader *bufio.Reader, req *Request) error {
	contentLengthStr := req.Headers.Get(HeaderContentLength)
	if contentLengthStr == "" {
		return nil // No body to parse if Content-Length is not set
	}

	contentLength, err := strconv.Atoi(contentLengthStr)
	if err != nil {
		return fmt.Errorf("invalid Content-Length value: %s", contentLengthStr)
	}

	if contentLength < 0 {
		return fmt.Errorf("invalid Content-Length: %s", contentLengthStr)
	}

	req.Body = make([]byte, contentLength)
	if _, err := io.ReadFull(reader, req.Body); err != nil {
		return fmt.Errorf("failed to read request body: %w", err)
	}

	return nil
}

// ParseResponse parses an HTTP response from a buffered reader.
func ParseResponse(reader *bufio.Reader) (*Response, error) {
	// Read the status line
	statusLine, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("failed to read status line: %w", err)
	}
	statusLine = strings.TrimSpace(statusLine)

	// Parse the status line
	parts := strings.SplitN(statusLine, " ", 3)
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid status line: %s", statusLine)
	}

	statusCode, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, fmt.Errorf("invalid status code: %s", parts[1])
	}

	resp := &Response{
		Version:    parts[0],
		StatusCode: statusCode,
		Headers:    NewHeaders(),
	}

	// Parse headers
	if err := parseHeaders(reader, resp.Headers); err != nil {
		return nil, fmt.Errorf("failed to parse headers: %w", err)
	}

	// Parse body if Content-Length is set
	if err := parseResponseBody(reader, resp); err != nil {
		return nil, fmt.Errorf("failed to parse body: %w", err)
	}

	return resp, nil
}

func parseResponseBody(reader *bufio.Reader, resp *Response) error {
	contentLengthStr := resp.Headers.Get(HeaderContentLength)
	if contentLengthStr == "" {
		return nil // No body to parse if Content-Length is not set
	}

	contentLength, err := strconv.Atoi(contentLengthStr)
	if err != nil {
		return fmt.Errorf("invalid Content-Length value: %s", contentLengthStr)
	}

	if contentLength < 0 {
		return fmt.Errorf("invalid Content-Length: %s", contentLengthStr)
	}

	resp.Body = make([]byte, contentLength)
	if _, err := io.ReadFull(reader, resp.Body); err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	return nil
}
