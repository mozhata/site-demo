package api

import (
	"net/http"

	"golang.org/x/net/trace"

	"bitbucket.org/applysquare/applysquare-go/pkg/api/negroni"
	"github.com/gorilla/mux"
)

// GetServingHandler returns the handler to serve given root.
func GetServingHandler(routes []Route) http.Handler {
	// Use negroni for middleware management.
	n := negroni.New(
		// TODO(guye): Setup a Raven middleware.
		negroni.NewRecovery(),
		negroni.NewLogger(),
		&profilingMiddleware{},
	)

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	// Handlers for debugging.

	// This is critical to determine whether the server has started. Do not remove it.
	router.HandleFunc("/debug/heartbeat", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok!"))
	})

	// For request trace.
	router.HandleFunc("/debug/requests", func(w http.ResponseWriter, r *http.Request) {
		if !isDev(r) {
			http.Error(w, "not allowed", http.StatusUnauthorized)
			return
		}
		trace.Render(w, r, true)
	})

	// For eventlog trace.
	router.HandleFunc("/debug/events", func(w http.ResponseWriter, r *http.Request) {
		if !isDev(r) {
			http.Error(w, "not allowed", http.StatusUnauthorized)
			return
		}
		trace.RenderEvents(w, r, true)
	})

	// Static.
	router.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", http.FileServer(http.Dir("./dist"))))

	n.UseHandler(router)
	return n
}
