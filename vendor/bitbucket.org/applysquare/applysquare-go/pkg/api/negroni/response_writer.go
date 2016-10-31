package negroni

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
)

// ResponseWriter is a wrapper around http.ResponseWriter that provides extra information about
// the response. It is recommended that middleware handlers use this construct to wrap a responsewriter
// if the functionality calls for it.
type ResponseWriter interface {
	http.ResponseWriter
	http.Flusher
	// Status returns the status code of the response or 0 if the response has not been written.
	Status() int
	// Written returns whether or not the ResponseWriter has been written.
	Written() bool
	// Size returns the size of the response body.
	Size() int
	// Before allows for a function to be called before the ResponseWriter has been written to. This is
	// useful for setting headers or any other operations that must happen before a response has been written.
	Before(func(ResponseWriter))
	// For high performance write string.
	WriteString(string) (int, error)
}

type beforeFunc func(ResponseWriter)

// NewResponseWriter creates a ResponseWriter that wraps an http.ResponseWriter
func NewResponseWriter(rw http.ResponseWriter) ResponseWriter {
	if srw, ok := rw.(sResponseWriter); ok {
		return &responseWriter{srw, 0, 0, nil}
	}
	log.Print("response writer doesn't support WriteString, performance will suffer.")
	return &responseWriter{&slowStringWriter{rw}, 0, 0, nil}
}

type sResponseWriter interface {
	http.ResponseWriter
	WriteString(string) (int, error)
}

type responseWriter struct {
	sResponseWriter
	status      int
	size        int
	beforeFuncs []beforeFunc
}

func (rw *responseWriter) WriteHeader(s int) {
	rw.status = s
	rw.callBefore()
	rw.sResponseWriter.WriteHeader(s)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	if !rw.Written() {
		// The status will be StatusOK if WriteHeader has not been called yet
		rw.WriteHeader(http.StatusOK)
	}
	size, err := rw.sResponseWriter.Write(b)
	rw.size += size
	return size, err
}

type stringWriter interface {
	WriteString(string) (int, error)
}

func (rw *responseWriter) WriteString(s string) (int, error) {
	if !rw.Written() {
		// The status will be StatusOK if WriteHeader has not been called yet
		rw.WriteHeader(http.StatusOK)
	}
	size, err := rw.sResponseWriter.WriteString(s)
	rw.size += size
	return size, err
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) Size() int {
	return rw.size
}

func (rw *responseWriter) Written() bool {
	return rw.status != 0
}

func (rw *responseWriter) Before(before func(ResponseWriter)) {
	rw.beforeFuncs = append(rw.beforeFuncs, before)
}

func (rw *responseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hijacker, ok := rw.sResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, fmt.Errorf("the ResponseWriter doesn't support the Hijacker interface")
	}
	return hijacker.Hijack()
}

func (rw *responseWriter) CloseNotify() <-chan bool {
	return rw.sResponseWriter.(http.CloseNotifier).CloseNotify()
}

func (rw *responseWriter) callBefore() {
	for i := len(rw.beforeFuncs) - 1; i >= 0; i-- {
		rw.beforeFuncs[i](rw)
	}
}

func (rw *responseWriter) Flush() {
	flusher, ok := rw.sResponseWriter.(http.Flusher)
	if ok {
		flusher.Flush()
	}
}
