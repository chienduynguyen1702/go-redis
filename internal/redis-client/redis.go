package redisclient

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

// RedisClient is a struct that holds a redis client
type RedisClient struct {
	Client *redis.Client
}

// NewRedisClient creates a new redis client
func NewRedisClient(ctx context.Context) *redis.Client {
	redisAddr := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	redisProtocol, _ := strconv.Atoi(os.Getenv("REDIS_PROTOCOL"))

	redisOptions := &redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
		Protocol: redisProtocol,
	}

	rdb := redis.NewClient(redisOptions)

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Error connecting to redis")
	}

	log.Println("Connected to redis")
	return rdb

}
