package user

import (
	"context"
	"github.com/gofrs/uuid"
)

type Storage interface {
	GetByUsername(ctx context.Context, username string) (item *User, err error)
	GetByID(ctx context.Context, id uuid.UUID) (item *User, err error)
	GetPasswordByUsername(ctx context.Context, username string) (pwdHash string, err error)
	UpdatePasswordByID(ctx context.Context, pwdHash string, id uuid.UUID) (err error)
	UpdateUsernameByID(ctx context.Context, username string, id uuid.UUID) (err error)
	CreateUser(ctx context.Context, username, pwdHash string) (err error)
	DeleteByID(ctx context.Context, id uuid.UUID) (err error)
}
