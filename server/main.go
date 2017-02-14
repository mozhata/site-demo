package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/mozhata/site-demo/server/check"
	"github.com/mozhata/site-demo/skeleton/route"
)

func openLog() {
	flag.Lookup("logtostderr").Value.Set("true")
	flag.Parse()
}
func main() {
	openLog()

	check.CheckMySQL()

	routes := check.NewRoutes()
	handler := route.BuildHandler(routes)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
