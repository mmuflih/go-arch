package user

import (
	"time"

	"github.com/mmuflih/datelib"
	"github.com/mmuflih/go-di-arch/domain/model"
)

/**
 * Created by Muhammad Muflih Kholidin
 * at 2020-09-28 23:14:27
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

type MeResponse struct {
	ID        uint64  `json:"id"`
	Name      string  `json:"name"`
	Sex       *string `json:"sex"`
	BirthDate string  `json:"birth_date"`
	LastLogin string  `json:"last_login"`
	AvatarURL string  `json:"avatar_url"`
}

func NewMeResponse(u *model.User) MeResponse {
	birthDate := ""
	if u.BirthDate != nil {
		birthDate = u.BirthDate.Format(datelib.YMD)
	}
	return MeResponse{
		u.ID, u.Name, u.Sex, birthDate,
		u.LastLogin.Format(time.RFC3339), u.AvatarURL,
	}
}

type GetTokenResponse struct {
	Auth AccessToken `json:"auth"`
	User MeResponse  `json:"user"`
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"type_type"`
	ExpiresIn   int64  `json:"expires_in"`
}
