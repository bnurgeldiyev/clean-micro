package cache

import (
	"context"
	"github.com/gofrs/uuid"
	"time"
)

type Service interface {
	UserSetRefreshToken(ctx context.Context, id, refreshToken uuid.UUID, expiry time.Duration) (err error)
	UserGetRefreshToken(ctx context.Context, id uuid.UUID) (refreshToken uuid.UUID, err error)
}
