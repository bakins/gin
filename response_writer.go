package gin

import (
	"net/http"
	"bufio"
        "errors"
        "net"
)

type (
	ResponseWriter interface {
		http.ResponseWriter
		Status() int
		Written() bool

		// private
		setStatus(int)
	}

	responseWriter struct {
		http.ResponseWriter
		status  int
		written bool
	}
)

func (w *responseWriter) reset(writer http.ResponseWriter) {
	w.ResponseWriter = writer
	w.status = 0
	w.written = false
}

func (w *responseWriter) setStatus(code int) {
	w.status = code
}

func (w *responseWriter) WriteHeader(code int) {
	w.status = code
	w.written = true
	w.ResponseWriter.WriteHeader(code)
}

func (w *responseWriter) Status() int {
	return w.status
}

func (w *responseWriter) Written() bool {
	return w.written
}

// allow connection hijacking
func (w *ResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hijacker, ok := w.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, errors.New("the ResponseWriter doesn't support the Hijacker interface")
	}
	return hijacker.Hijack()
}
