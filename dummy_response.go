package testkit // import "github.com/novakit/testkit"

import (
	"bytes"
	"net/http"
)

// DummyResponse dummy http.ResponseWriter for testing purpose
type DummyResponse struct {
	*bytes.Buffer
	StatusCode int

	header http.Header
}

// Header implements http.ResponseWriter
func (d *DummyResponse) Header() http.Header {
	return d.header
}

// WriteHeader implements http.ResponseWriter
func (d *DummyResponse) WriteHeader(statusCode int) {
	d.StatusCode = statusCode
}

// Write override Buffer.Write with StatusCode setup
func (d *DummyResponse) Write(p []byte) (int, error) {
	if d.StatusCode == 0 {
		d.StatusCode = http.StatusOK
	}
	return d.Buffer.Write(p)
}

// NewDummyResponse create a dummy response
func NewDummyResponse() *DummyResponse {
	return &DummyResponse{
		Buffer: &bytes.Buffer{},
		header: http.Header{},
	}
}
