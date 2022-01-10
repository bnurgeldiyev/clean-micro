package user

import (
	"context"
	"github.com/gofrs/uuid"
)

type Service interface {
	Auth(ctx context.Context, username, password string) (item *Auth, err error)
	Create(ctx context.Context, username, password string) (err error)
	Delete(ctx context.Context, username string) (err error)
	UpdateUsername(ctx context.Context, password, oldUsername, newUsername string) (err error)
	UpdatePassword(ctx context.Context, username, oldPassword, newPassword string) (err error)
	Access(ctx context.Context, accessToken string) (username string, err error)
	UpdateAccessToken(ctx context.Context, accessToken string, refreshToken uuid.UUID) (token string, err error)
}
