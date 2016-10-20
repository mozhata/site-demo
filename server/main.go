package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/garyburd/redigo/redis"
	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
	"github.com/zykzhang/site-demo/server/database/sqldb"
)

func openLog() {
	flag.Lookup("logtostderr").Value.Set("true")
	flag.Parse()
}
func main() {
	openLog()

	checkMySQL()

	router := httprouter.New()

	router.GET("/", welcome)
	router.GET("/redis", checkRedis)
	router.GET("/andy/:name", hi)
	router.GET("/match/*filepath", matchAll)
	router.GET("/protected", basicAuth(protected, "kang", "123!"))

	// router.NotFound = notFound

	n := negroni.Classic()
	n.UseHandler(router)

	log.Fatal(http.ListenAndServe(":8080", n))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is not found handler")
}

func protected(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "this is protected ~")
}

func basicAuth(h httprouter.Handle, requiredUser, requiredPassword string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Get the Basic Authentication credentials
		user, password, hasAuth := r.BasicAuth()

		if hasAuth && user == requiredUser && password == requiredPassword {
			// Delegate request to the given handle
			h(w, r, ps)
		} else {
			// Request Basic Authentication otherwise
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}

func matchAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "the filepath is %s", ps.ByName("filepath"))
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

func checkMySQL() {
	env := sqldb.NewEnv()
	glog.Infoln(env.HasTable("local_auth"))
}
