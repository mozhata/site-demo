package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/garyburd/redigo/redis"
	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"
)

func openLog() {
	flag.Lookup("logtostderr").Value.Set("true")
	flag.Parse()
}
func main() {
	openLog()

	router := httprouter.New()

	router.GET("/", welcome)
	router.GET("/andy/:name", hi)
	router.GET("/redis", checkRedis)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func welcome(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "welcome, this is powered by %s", "httprouter")
}

func hi(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello %s", ps.ByName("name"))
}

func checkRedis(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	conn, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	v, err := conn.Do("SET", "name", "red")
	if err != nil {
		glog.Error(err)
	}
	glog.Infoln("v: ", v)

	v, err = redis.String(conn.Do("GET", "name"))
	if err != nil {
		glog.Error(err)
	}
	fmt.Fprintf(w, "check redis, the value got is %s", v)
}
