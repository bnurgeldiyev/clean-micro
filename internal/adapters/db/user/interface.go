package user

import (
	"clean-micro/internal/domain/user"
	"github.com/jmoiron/sqlx"
)

type userStorage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) user.Storage {
	return &userStorage{
		db: db,
	}
}
