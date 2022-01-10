package cache

import (
	"clean-micro/pkg/redisdb"
	"github.com/go-redis/redis/v8"
)

type RedisService struct {
	Service
	client *redis.Client
}

func NewRedisService(redisConn string, redisDb int, maxIdle int, maxActive int) (r *RedisService) {

	client := redisdb.NewRedisService(redisConn, redisDb, maxIdle, maxActive)

	return &RedisService{
		client: client,
	}
}
