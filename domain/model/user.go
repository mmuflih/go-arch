package model

import (
	"time"

	"github.com/mmuflih/datelib"
)

type User struct {
	ID        uint64           `json:"id"`
	Name      string           `json:"name"`
	Sex       *string          `json:"sex"`
	BirthDate *time.Time       `json:"birth_date"`
	LastLogin time.Time        `json:"last_login"`
	AvatarURL string           `json:"avatar_url"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	DeletedAt datelib.NullTime `json:"deleted_at"`
}

func NewUser(name string) *User {
	now := time.Now()
	return &User{
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
		Sex:       nil,
		LastLogin: now,
	}
}
