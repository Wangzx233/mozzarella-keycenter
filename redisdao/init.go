package redisdao

import (
	"github.com/go-redis/redis"
	"log"
)

var rdb *redis.Client

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "root", // no password set
		DB:       0,      // use default DB
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		log.Println(err)
	}
}
