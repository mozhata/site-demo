/*
this package is just use as a test
*/
package check

import (
	"fmt"
	"net/http"

	"github.com/garyburd/redigo/redis"
	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"
	"github.com/zykzhang/site-demo/server/database/sqldb"
	"github.com/zykzhang/site-demo/skeleton/route"
)

func CheckMySQL() {
	env := sqldb.NewEnv()
	glog.Infoln(env.HasTable("local_auth"))
}

func NewRoutes() []*route.Route {
	return []*route.Route{
		route.NewRoute(
			"/",
			"GET",
			welcome,
		),
		route.NewRoute(
			"/redis",
			"GET",
			checkRedis,
		),
		route.NewRoute(
			"/any/:name",
			"GET",
			hi,
		),
		route.NewRoute(
			"/match/*filepath",
			"GET",
			matchAll,
		),
		route.NewRoute(
			"/protected",
			"GET",
			basicAuth(protected, "kang", "123!"),
		),
	}
}

func welcome(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "welcome, this is powered by %s", "httprouter")
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

func hi(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello %s", ps.ByName("name"))
}

func matchAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "the filepath is %s", ps.ByName("filepath"))
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
func protected(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "this is protected ~")
}
