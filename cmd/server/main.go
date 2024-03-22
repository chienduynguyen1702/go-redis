package server

import (
	"context"

	"github.com/chienduynguyen1702/go-redis/initialize"

	redisclient "github.com/chienduynguyen1702/go-redis/internal/redis-client"
)

func Init() {
	initialize.LoadEnvVarFile()
}

func main() {
	// Load environment variables

	// Create a new redis client
	redisClient := redisclient.NewRedisClient(context.Background())

}
