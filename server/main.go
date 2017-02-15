package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/mozhata/site-demo/server/check"
	"github.com/mozhata/site-demo/skeleton/route"
)

func init() {
	flag.Set("logtostderr", "true")
	flag.Parse()
}

func main() {
	check.CheckMySQL()

	routes := check.NewRoutes()
	handler := route.BuildHandler(routes)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
