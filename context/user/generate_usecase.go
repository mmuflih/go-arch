package user

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mmuflih/go-di-arch/config"
	"github.com/mmuflih/go-di-arch/domain/model"
)

/**
 * Created by Muhammad Muflih Kholidin
 * at 2020-10-04 18:33:51
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

type GenerateUsecase interface {
	ClaimToken(*model.User) GetTokenResponse
}

type generateUC struct {
	keys *config.Keys
}

func NewGenerateUsecase(keys *config.Keys) GenerateUsecase {
	return &generateUC{keys}
}

func (gu *generateUC) ClaimToken(u *model.User) GetTokenResponse {
	/* Create the token */
	token := jwt.New(jwt.SigningMethodHS512)

	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	expiredAt := time.Now().Add(time.Hour * (24 * 1500)).Unix()
	claims["user_id"] = u.ID
	claims["exp"] = expiredAt

	/* Sign the token with our secret */
	tokenString, _ := token.SignedString(gu.keys.Signature)

	auth := AccessToken{
		AccessToken: tokenString,
		TokenType:   "Bearer",
		ExpiresIn:   expiredAt,
	}
	return gu.createResponse(auth, u)
}

func (gu *generateUC) createResponse(auth AccessToken,
	u *model.User) GetTokenResponse {
	return GetTokenResponse{
		auth,
		NewMeResponse(u),
	}
}
