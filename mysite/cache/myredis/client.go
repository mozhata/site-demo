package myredis

import "github.com/garyburd/redigo/redis"

// import "gopkg.in/redis.v3"

// func NewClient() *redis.Client {
// 	return redis.NewClient(&redis.Options{
// 		Addr:     "redis:6379",
// 		Password: "",
// 		DB:       0,
// 	})
// }

func NewConn() redis.Conn {
	conn, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	return conn
}
