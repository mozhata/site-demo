package my-redis

import "gopkg.in/redis.v3"

func NewClient() *redis.Client {
	return redis.NewClient(&redis.Options{
	Addr: "redis:6377",
  Password: "",
  DB: 0,
	})
}
