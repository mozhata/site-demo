package main

import (
	"github.com/golang/glog"
	"github.com/zykzhang/site-demo/cmd/livereload/lib"
)

func main() {
	projectDir := lib.GetProjectDir()

	r := lib.NewMethodRegistry()

	// r.Register("ass_dev", func() {
	// 	// Meant to be run in dev docker env.
	// 	runner := internal.NewBuilderRunner(projectDir, "ass", "cmd/ass/main.go", []string{
	// 		"-v=1",
	// 		"--alsologtostderr",
	// 		"--couchdb=http://couchdb:5984",
	// 		"--elasticsearch=elasticsearch:9200",
	// 		"--addr=:3001",
	// 		"--sql_driver=postgres",
	// 		`--sql_params=user=postgres password=postgres dbname=postgres host=postgres port=5432 sslmode=disable`,
	// 		"serve",
	// 	})

	r.Register("dev", func() {
		glog.Infoln("this is a placeholder")
	})


	glog.Infoln(projectDir, r)
}
