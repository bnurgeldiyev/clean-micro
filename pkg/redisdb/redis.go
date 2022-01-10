package redisdb

import "github.com/go-redis/redis/v8"

func NewRedisService(redisConn string, redisDb int, maxIdle int, maxActive int) (client *redis.Client) {

	client = redis.NewClient(&redis.Options{
		Addr:         redisConn,
		DB:           redisDb,
		PoolSize:     maxActive,
		MinIdleConns: maxIdle,
	})

	return
}
