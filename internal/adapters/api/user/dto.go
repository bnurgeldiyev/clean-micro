package user

import (
	"github.com/gofrs/uuid"
	"time"
)

type User struct {
	ID       uuid.UUID `json:"id" db:"id"`
	Username string    `json:"username" db:"username"`
	CreateTS time.Time `json:"create_ts" db:"create_ts"`
	UpdateTS time.Time `json:"update_ts" db:"update_ts"`
}

type Auth struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
