package route

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

// Route basic route
type Route struct {
	Pattern string
	Method  string
	Handle  httprouter.Handle
}

func NewRoute(pattern, method string, handle httprouter.Handle) *Route {
	return &Route{
		Pattern: pattern,
		Method:  method,
		Handle:  handle,
	}
}

func BuildHandler(routes []*Route) http.Handler {
	router := httprouter.New()
	for _, route := range routes {
		router.Handle(route.Method, route.Pattern, route.Handle)
	}

	// TODO: add serverfile route

	// use middleware
	n := negroni.Classic()
	n.UseHandler(router)
	return n
}
