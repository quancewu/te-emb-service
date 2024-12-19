package initalizers

import (
	"context"
	"os"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

var redisLock sync.RWMutex

var ctx = context.Background()

func ConnectToRedis() {
	redis_url := os.Getenv("REDIS")
	opt, err := redis.ParseURL(redis_url)
	if err != nil {
		panic(err)
	}

	Redis = redis.NewClient(opt)
}

func SetSessionIdToRedis(sessionId string, userInfo []byte) error {
	redisLock.Lock()
	defer redisLock.Unlock()
	err := Redis.Set(ctx, sessionId, userInfo, 24*60*time.Minute).Err()
	if err != nil {
		// log.Printf("Key %s can't save into Redis cache\n", sessionId)
		return err
	} else {
		_, err := Redis.Expire(ctx, sessionId, 24*60*time.Minute).Result()
		return err
	}
}

func GetSessionIdFromRedis(sessionId string) ([]byte, error) {
	redisLock.Lock()
	defer redisLock.Unlock()
	userInfo, err := Redis.Get(ctx, sessionId).Result()
	if err != nil {
		// log.Printf("Key %s can't get from Redis cache\n", sessionId)
		return []byte(""), err
	} else {
		return []byte(userInfo), nil
	}
}
