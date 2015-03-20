package goway

import (
	"net/http"
)

const (
	contentJsonType = "application/json"
)

type ResponseWriter interface {
	Header() http.Header
	Write([]byte) (int, error)
	WriteHeader(int)
	IsWrite() bool
	Status() int
	Size() int
}

type responseWriter struct {
	responseWriter http.ResponseWriter
	size           int
	status         int
}

// Created an new ResponseWriter
func NewResponseWriter(rw http.ResponseWriter) ResponseWriter {
	return &responseWriter{rw, 0, 0}
}

// Header returns the header map that will be sent by
// WriteHeader. Changing the header after a call to
// WriteHeader (or Write) has no effect unless the modified
// headers were declared as trailers by setting the
// "Trailer" header before the call to WriteHeader.
func (rw *responseWriter) Header() http.Header {
	return rw.responseWriter.Header()
}

// Write writes the data to the connection as part of an HTTP reply.
// If WriteHeader has not yet been called, Write calls WriteHeader(http.StatusOK)
// before writing the data.  If the Header does not contain a
// Content-Type line, Write adds a Content-Type set to the result of passing
// the initial 512 bytes of written data to DetectContentType.
// and sets `started` to true
func (rw *responseWriter) Write(p []byte) (int, error) {
	if !rw.IsWrite() {
		// The status will be StatusOK if WriteHeader has not been called yet
		rw.WriteHeader(http.StatusOK)
	}
	size, err := rw.responseWriter.Write(p)
	rw.size += size
	return size, err
}

// WriteHeader sends an HTTP response header with status code.
// If WriteHeader is not called explicitly, the first call to Write
// will trigger an implicit WriteHeader(http.StatusOK).
// Thus explicit calls to WriteHeader are mainly used to
// send error codes.
func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.responseWriter.WriteHeader(statusCode)
	rw.status = statusCode

}

// HTTP status
func (rw *responseWriter) Status() int {
	return rw.status
}

// HTTP content length size
func (rw *responseWriter) Size() int {
	return rw.size
}

// whether a write status for response writer
func (rw *responseWriter) IsWrite() bool {
	return rw.status != 0
}
