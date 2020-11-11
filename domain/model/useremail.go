package model

import "time"

/**
 * Created by Muhammad Muflih Kholidin
 * at 2020-09-25 17:37:10
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

type UserEmail struct {
	ID        uint64    `json:"id"`
	UserID    uint64    `json:"user_id"`
	Email     string    `json:"email"`
	Primary   bool      `json:"primary"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUserEmail(userID uint64, email string) *UserEmail {
	now := time.Now()
	return &UserEmail{
		0, userID, email, true, false, now, now,
	}
}
