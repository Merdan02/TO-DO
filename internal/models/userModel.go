package models

import (
	"time"
)

type UserModel struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password_hash"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
