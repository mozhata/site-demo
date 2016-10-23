package route

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

// Route basic route
type Route struct {
	Method  string
	Pattern string
	Handle  httprouter.Handle
}

func BuildHandler(routes []Route) http.Handler {
	router := httprouter.New()
	for _, route := range routes {
		router.Handle(route.Method, route.Pattern, route.Handle)
	}

	// use middleware
	n := negroni.Classic()
	n.UseHandler(router)
	return n
}
