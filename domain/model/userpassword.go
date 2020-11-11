package model

import (
	"time"
)

/**
 * Created by Muhammad Muflih Kholidin
 * at 2020-09-25 17:37:18
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

type UserPassword struct {
	ID           uint64    `json:"id"`
	UserID       uint64    `json:"user_id"`
	Password     string    `json:"password"`
	Active       bool      `json:"active"`
	ResetToken   *string   `json:"reset_token"`
	TokenExpired *uint64   `json:"token_expired"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func NewUserPassword(userID uint64, password string) *UserPassword {
	now := time.Now()
	return &UserPassword{
		0, userID, password, true, nil, nil, now, now,
	}
}
