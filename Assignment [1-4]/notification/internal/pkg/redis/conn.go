package redis

import (
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func InitRDBConn() *redis.Client {
	maxRetries, _ := strconv.Atoi(os.Getenv("MAX_RETRIES"))

	return redis.NewClient(&redis.Options{
		Addr:       "redis:6379",
		Password:   "",
		DB:         0,
		MaxRetries: maxRetries,
	})
}
