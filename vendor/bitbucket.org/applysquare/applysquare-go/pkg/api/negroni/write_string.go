package negroni

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
)

type slowStringWriter struct {
	http.ResponseWriter
}

func (w *slowStringWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hijacker, ok := w.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, fmt.Errorf("the ResponseWriter doesn't support the Hijacker interface")
	}
	return hijacker.Hijack()
}

func (w *slowStringWriter) WriteString(s string) (int, error) {
	return w.ResponseWriter.Write([]byte(s))
}
