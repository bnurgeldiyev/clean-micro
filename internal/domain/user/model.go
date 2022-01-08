package user

import (
	"github.com/gofrs/uuid"
	"time"
)

type User struct {
	ID       uuid.UUID `json:"id" db:"id"`
	Username string    `json:"username" db:"username"`
	Password string    `json:"password" db:"password"`
	CreateTS time.Time `json:"create_ts" db:"create_ts"`
	UpdateTS time.Time `json:"update_ts" db:"update_ts"`
}
