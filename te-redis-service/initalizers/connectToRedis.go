package initalizers

import (
	"os"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func ConnectToRedis() {
	redisURL := os.Getenv("REDIS")
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		panic(err)
	}
	Redis = redis.NewClient(opt)
}
