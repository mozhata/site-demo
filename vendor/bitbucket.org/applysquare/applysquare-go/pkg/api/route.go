package api

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func NewRoute(name string, method string, pattern string, handler http.HandlerFunc) Route {
	return Route{
		Name:        name,
		Method:      method,
		Pattern:     pattern,
		HandlerFunc: handler,
	}
}
