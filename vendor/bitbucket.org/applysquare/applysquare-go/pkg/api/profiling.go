package api

import (
	"fmt"
	"net/http"
	"runtime/pprof"

	"bitbucket.org/applysquare/applysquare-go/pkg/common"
)

func isDev(r *http.Request) bool {
	for _, p := range common.SplitAndTrim(r.Header.Get("X-A2-USER-PERMISSIONS"), ",") {
		if p == "dev" {
			return true
		}
	}
	return false
}

type noopResponseWriter struct{}

var noopRW http.ResponseWriter = &noopResponseWriter{}

func (n *noopResponseWriter) Header() http.Header {
	return http.Header{}
}

func (n *noopResponseWriter) Write(d []byte) (int, error) {
	return len(d), nil
}

func (n *noopResponseWriter) WriteHeader(s int) {}

type profilingMiddleware struct{}

func (p *profilingMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if r.URL.Query().Get("profiling") == "" {
		next(rw, r)
		return
	}
	if !isDev(r) {
		next(rw, r)
		return
	}

	rw.Header().Set("Content-Type", "application/octet-stream")
	if err := pprof.StartCPUProfile(rw); err != nil {
		// StartCPUProfile failed, so no writes yet.
		// Can change header back to text content
		// and send error code.
		rw.Header().Set("Content-Type", "text/plain; charset=utf-8")
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, "Could not enable CPU profiling: %s\n", err)
		return
	}
	defer pprof.StopCPUProfile()
	next(noopRW, r)
}
