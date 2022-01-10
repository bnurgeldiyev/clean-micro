package cache

import (
	"clean-micro/pkg/helpers"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gofrs/uuid"
	"time"
)

const userRefreshToken = "user:refreshToken:"

func (r *RedisService) UserSetRefreshToken(ctx context.Context, id, refreshToken uuid.UUID, expiry time.Duration) (err error) {
	key := fmt.Sprintf("%v%v", userRefreshToken, id)

	err = r.client.Set(ctx, key, refreshToken.String(), expiry).Err()

	return
}

func (r *RedisService) UserGetRefreshToken(ctx context.Context, id uuid.UUID) (refreshToken uuid.UUID, err error) {
	key := fmt.Sprintf("%v%v", userRefreshToken, id)

	tokenStr, err := r.client.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return
	}

	if err == redis.Nil {
		return uuid.Nil, nil
	}

	refreshToken, err = helpers.ConvertStringToUUID(tokenStr)

	return
}
