package api

import "net/http"

type HandlerForTest struct {
	H           http.Handler
	UserId      string
	Permissions string
}

func (h *HandlerForTest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.Header.Add("X-A2-USER-UUID", h.UserId)
	if h.Permissions != "" {
		r.Header.Add("X-A2-USER-PERMISSIONS", h.Permissions)
	}
	h.H.ServeHTTP(w, r)
}
