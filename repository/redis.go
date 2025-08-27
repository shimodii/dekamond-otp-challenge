package repository

import (
	"os"

	"github.com/redis/go-redis/v9"
)

var RedisDB *redis.Client

func OpenRedis() {

	redisdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "",
		DB:       0,
	})

	RedisDB = redisdb
}
