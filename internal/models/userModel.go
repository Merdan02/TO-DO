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
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password_hash"`
}

type UserResponse struct {
	ID       int       `json:"id"`
	Username string    `json:"username"`
	Role     string    `json:"role"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}
