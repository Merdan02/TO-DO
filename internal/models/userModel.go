package models

import (
	"github.com/google/uuid"
	"time"
)

type UserModel struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password_hash"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
