package trace

import (
	"net/http"

	"github.com/gorilla/context"
)

const key = 1443109103

// SetToRequest attach a trace to the http request.
func SetToRequest(tr T, r *http.Request) {
	context.Set(r, key, tr)
}

// FromRequest returns the attached trace from http request.
func FromRequest(r *http.Request) T {
	return context.Get(r, key).(T)
}
