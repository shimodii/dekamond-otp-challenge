package repository

import (
	"github.com/redis/go-redis/v9"
)

var RedisDB *redis.Client

func OpenRedis() {
	redisdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	RedisDB = redisdb
}
